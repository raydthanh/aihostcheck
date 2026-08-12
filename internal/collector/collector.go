package collector

import (
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/raydthanh/aihostcheck/internal/model"
	"github.com/raydthanh/aihostcheck/internal/probe"
)

func Collect(timeout time.Duration, version string) model.Report {
	c := platformCapabilities(timeout)
	c["os"] = model.Capability{Status: model.Detected, Value: runtime.GOOS, Details: map[string]string{"architecture": runtime.GOARCH}, Evidence: []model.Evidence{{Source: "go_runtime", Detail: "GOOS and GOARCH of running binary"}}}
	pythonCandidates := []commandSpec{{name: "python3", args: []string{"--version"}}, {name: "python", args: []string{"--version"}}}
	if runtime.GOOS == "windows" {
		pythonCandidates = []commandSpec{{name: "py.exe", args: []string{"--version"}}, {name: "python.exe", args: []string{"--version"}}, {name: "python3.exe", args: []string{"--version"}}}
	}
	c["python"] = firstAvailable(timeout, pythonCandidates)

	commands := map[string]commandSpec{
		"nodejs":        {name: "node", args: []string{"--version"}},
		"go":            {name: "go", args: []string{"version"}},
		"java":          {name: "java", args: []string{"-version"}},
		"git":           {name: "git", args: []string{"--version"}},
		"docker":        {name: "docker", args: []string{"--version"}},
		"podman":        {name: "podman", args: []string{"--version"}},
		"nvidia_driver": {name: "nvidia-smi", args: []string{"--query-gpu=driver_version", "--format=csv,noheader"}},
	}
	for key, p := range commands {
		c[key] = probe.Command(timeout, p.name, p.args...)
	}
	c["cuda"] = cudaCapability(timeout)
	return model.Report{SchemaVersion: model.SchemaVersion, GeneratedAt: time.Now().UTC().Format(time.RFC3339), ToolVersion: version, Platform: runtime.GOOS + "/" + runtime.GOARCH, Capabilities: c}
}

type commandSpec struct {
	name string
	args []string
}

func firstAvailable(timeout time.Duration, candidates []commandSpec) model.Capability {
	var inconclusive []model.Evidence
	for _, candidate := range candidates {
		result := probe.Command(timeout, candidate.name, candidate.args...)
		if result.Status == model.Detected {
			result.Details = map[string]string{"executable": candidate.name}
			return result
		}
		if result.Status != model.NotDetected {
			inconclusive = append(inconclusive, result.Evidence...)
		}
	}
	if len(inconclusive) > 0 {
		return model.Capability{Status: model.Unknown, Evidence: inconclusive}
	}
	return model.Capability{Status: model.NotDetected, Evidence: []model.Evidence{{Source: "executable_lookup", Detail: "no supported executable candidate found in PATH"}}}
}

func allAvailable(timeout time.Duration, candidates []commandSpec) model.Capability {
	var values []string
	var evidence []model.Evidence
	var inconclusive []model.Evidence
	for _, candidate := range candidates {
		result := probe.Command(timeout, candidate.name, candidate.args...)
		if result.Status == model.Detected {
			values = append(values, candidate.name+": "+result.Value)
			evidence = append(evidence, result.Evidence...)
		} else if result.Status != model.NotDetected {
			inconclusive = append(inconclusive, result.Evidence...)
		}
	}
	if len(values) == 0 {
		if len(inconclusive) > 0 {
			return model.Capability{Status: model.Unknown, Evidence: inconclusive}
		}
		return model.Capability{Status: model.NotDetected, Evidence: []model.Evidence{{Source: "executable_lookup", Detail: "no supported candidate found in PATH"}}}
	}
	return model.Capability{Status: model.Detected, Value: strings.Join(values, "; "), Evidence: evidence, Details: map[string]string{"count": strconv.Itoa(len(values))}}
}

func cudaCapability(timeout time.Duration) model.Capability {
	r := probe.Run(timeout, "nvcc", "--version")
	if r.Status != model.Detected {
		return model.Capability{Status: r.Status, Evidence: r.Evidence}
	}
	for _, line := range strings.Split(r.Output, "\n") {
		if strings.Contains(strings.ToLower(line), "release") {
			return model.Capability{Status: model.Detected, Value: strings.TrimSpace(line), Evidence: r.Evidence}
		}
	}
	return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "command", Detail: "nvcc completed without a parseable release line"}}}
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
