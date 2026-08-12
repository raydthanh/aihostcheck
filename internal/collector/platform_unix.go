//go:build linux || darwin

package collector

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/raydthanh/aihostcheck/internal/model"
	"github.com/raydthanh/aihostcheck/internal/probe"
)

func platformCapabilities(timeout time.Duration) map[string]model.Capability {
	c := map[string]model.Capability{}
	c["cpu"] = model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d logical CPUs", runtime.NumCPU()), Evidence: []model.Evidence{{Source: "go_runtime", Detail: "runtime.NumCPU"}}}
	c["memory"] = memoryCapability(timeout)
	var s syscall.Statfs_t
	if err := syscall.Statfs("/", &s); err == nil {
		c["root_storage"] = model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d bytes total, %d bytes available", uint64(s.Blocks)*uint64(s.Bsize), uint64(s.Bavail)*uint64(s.Bsize)), Evidence: []model.Evidence{{Source: "statfs", Detail: "root filesystem"}}}
	} else {
		c["root_storage"] = model.Capability{Status: model.Error, Evidence: []model.Evidence{{Source: "statfs", Detail: "root filesystem query failed"}}}
	}
	shell := os.Getenv("SHELL")
	if shell != "" {
		c["shell"] = model.Capability{Status: model.Detected, Value: filepath.Base(shell), Details: map[string]string{"kind": "login_shell"}, Evidence: []model.Evidence{{Source: "environment", Detail: "basename of SHELL; this identifies the login shell, not necessarily the current parent process"}}}
	} else {
		c["shell"] = model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "environment", Detail: "SHELL is not set"}}}
	}
	if runtime.GOOS == "darwin" {
		c["os_version"] = probe.Command(timeout, "sw_vers", "-productVersion")
		c["package_manager"] = allAvailable(timeout, []commandSpec{{name: "brew", args: []string{"--version"}}, {name: "port", args: []string{"version"}}})
		c["gpu"] = macGPUCapability(timeout)
	} else {
		c["os_version"] = osRelease()
		c["package_manager"] = allAvailable(timeout, []commandSpec{{name: "apt-get", args: []string{"--version"}}, {name: "dnf", args: []string{"--version"}}, {name: "yum", args: []string{"--version"}}, {name: "pacman", args: []string{"--version"}}, {name: "apk", args: []string{"--version"}}})
		c["gpu"] = linuxGPUCapability(timeout)
	}
	return c
}

func memoryCapability(timeout time.Duration) model.Capability {
	b, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		if runtime.GOOS == "darwin" {
			return probe.Command(timeout, "sysctl", "-n", "hw.memsize")
		}
		return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "file", Detail: "/proc/meminfo could not be read"}}}
	}
	for _, line := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(line, "MemTotal:") {
			f := strings.Fields(line)
			if len(f) >= 2 {
				kb, err := strconv.ParseUint(f[1], 10, 64)
				if err != nil {
					return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "file", Detail: "/proc/meminfo MemTotal was not parseable"}}}
				}
				return model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d bytes", kb*1024), Evidence: []model.Evidence{{Source: "file", Detail: "/proc/meminfo MemTotal"}}}
			}
		}
	}
	return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "file", Detail: "/proc/meminfo did not contain MemTotal"}}}
}

func osRelease() model.Capability {
	b, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "file", Detail: "/etc/os-release could not be read"}}}
	}
	for _, l := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(l, "PRETTY_NAME=") {
			return model.Capability{Status: model.Detected, Value: strings.Trim(strings.TrimPrefix(l, "PRETTY_NAME="), "\""), Evidence: []model.Evidence{{Source: "file", Detail: "/etc/os-release PRETTY_NAME"}}}
		}
	}
	return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "file", Detail: "/etc/os-release did not contain PRETTY_NAME"}}}
}

func linuxGPUCapability(timeout time.Duration) model.Capability {
	var lines []string
	var evidence []model.Evidence
	for _, class := range []string{"::0300", "::0302"} {
		r := probe.Run(timeout, "lspci", "-nn", "-d", class)
		if r.Status == model.NotDetected {
			return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "executable_lookup", Detail: "lspci is unavailable, so GPU presence was not determined"}}}
		}
		if r.Status != model.Detected {
			return model.Capability{Status: r.Status, Evidence: r.Evidence}
		}
		lines = append(lines, nonEmptyLines(r.Output)...)
		evidence = append(evidence, r.Evidence...)
	}
	if len(lines) == 0 {
		return model.Capability{Status: model.NotDetected, Evidence: evidence}
	}
	return model.Capability{Status: model.Detected, Value: strings.Join(lines, "; "), Evidence: evidence}
}

func macGPUCapability(timeout time.Duration) model.Capability {
	r := probe.Run(timeout, "system_profiler", "-json", "SPDisplaysDataType")
	if r.Status == model.NotDetected {
		return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "executable_lookup", Detail: "system_profiler is unavailable, so GPU presence was not determined"}}}
	}
	if r.Status != model.Detected {
		return model.Capability{Status: r.Status, Evidence: r.Evidence}
	}
	models, err := parseMacGPUs(r.Output)
	if err != nil {
		return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "command", Detail: "system_profiler returned unparseable JSON"}}}
	}
	if len(models) == 0 {
		return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "command", Detail: "system_profiler returned no parseable GPU model"}}}
	}
	return model.Capability{Status: model.Detected, Value: strings.Join(models, "; "), Evidence: r.Evidence}
}

func parseMacGPUs(output string) ([]string, error) {
	var payload struct {
		Displays []map[string]any `json:"SPDisplaysDataType"`
	}
	if err := json.Unmarshal([]byte(output), &payload); err != nil {
		return nil, err
	}
	var models []string
	seen := map[string]bool{}
	for _, display := range payload.Displays {
		for _, key := range []string{"sppci_model", "_name"} {
			value, ok := display[key].(string)
			value = strings.TrimSpace(value)
			if ok && value != "" && !seen[value] {
				seen[value] = true
				models = append(models, value)
				break
			}
		}
	}
	return models, nil
}

func nonEmptyLines(output string) []string {
	var lines []string
	for _, line := range strings.Split(output, "\n") {
		if line = strings.TrimSpace(line); line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
