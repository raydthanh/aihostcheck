//go:build linux || darwin

package collector

import (
	"fmt"
	"os"
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
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	c["cpu"] = model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d logical CPUs", runtime.NumCPU()), Evidence: []model.Evidence{{Source: "go_runtime", Detail: "runtime.NumCPU"}}}
	c["memory"] = memoryCapability()
	var s syscall.Statfs_t
	if err := syscall.Statfs("/", &s); err == nil {
		c["root_storage"] = model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d bytes total, %d bytes available", uint64(s.Blocks)*uint64(s.Bsize), uint64(s.Bavail)*uint64(s.Bsize)), Evidence: []model.Evidence{{Source: "statfs", Detail: "root filesystem"}}}
	} else {
		c["root_storage"] = model.Capability{Status: model.Error, Evidence: []model.Evidence{{Source: "statfs", Detail: "root filesystem query failed"}}}
	}
	shell := os.Getenv("SHELL")
	if shell != "" {
		c["shell"] = model.Capability{Status: model.Detected, Value: shell, Evidence: []model.Evidence{{Source: "environment", Detail: "SHELL variable only"}}}
	} else {
		c["shell"] = model.Capability{Status: model.Unknown}
	}
	if runtime.GOOS == "darwin" {
		c["os_version"] = probe.Command(timeout, "sw_vers", "-productVersion")
		c["package_manager"] = probe.Command(timeout, "brew", "--version")
		c["gpu"] = probe.Command(timeout, "system_profiler", "SPDisplaysDataType")
	} else {
		c["os_version"] = osRelease()
		c["package_manager"] = firstDetected(timeout, []string{"apt-get", "dnf", "yum", "pacman", "apk"})
		c["gpu"] = probe.Command(timeout, "lspci", "-nn", "-d", "::0300")
	}
	return c
}

func memoryCapability() model.Capability {
	b, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		if runtime.GOOS == "darwin" {
			return probe.Command(2*time.Second, "sysctl", "-n", "hw.memsize")
		}
		return model.Capability{Status: model.Unknown}
	}
	for _, line := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(line, "MemTotal:") {
			f := strings.Fields(line)
			if len(f) >= 2 {
				kb, _ := strconv.ParseUint(f[1], 10, 64)
				return model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d bytes", kb*1024), Evidence: []model.Evidence{{Source: "file", Detail: "/proc/meminfo MemTotal"}}}
			}
		}
	}
	return model.Capability{Status: model.Unknown}
}

func osRelease() model.Capability {
	b, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return model.Capability{Status: model.Unknown}
	}
	for _, l := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(l, "PRETTY_NAME=") {
			return model.Capability{Status: model.Detected, Value: strings.Trim(strings.TrimPrefix(l, "PRETTY_NAME="), "\""), Evidence: []model.Evidence{{Source: "file", Detail: "/etc/os-release PRETTY_NAME"}}}
		}
	}
	return model.Capability{Status: model.Unknown}
}
func firstDetected(t time.Duration, names []string) model.Capability {
	for _, n := range names {
		r := probe.Command(t, n, "--version")
		if r.Status == model.Detected {
			r.Details = map[string]string{"manager": n}
			return r
		}
	}
	return model.Capability{Status: model.NotDetected, Evidence: []model.Evidence{{Source: "executable_lookup", Detail: "known package managers not found"}}}
}
