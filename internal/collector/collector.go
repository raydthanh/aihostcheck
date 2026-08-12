package collector

import (
	"runtime"
	"time"

	"github.com/raydthanh/aihostcheck/internal/model"
	"github.com/raydthanh/aihostcheck/internal/probe"
)

func Collect(timeout time.Duration, version string) model.Report {
	c := platformCapabilities(timeout)
	c["os"] = model.Capability{Status: model.Detected, Value: runtime.GOOS, Details: map[string]string{"architecture": runtime.GOARCH}, Evidence: []model.Evidence{{Source: "go_runtime", Detail: "GOOS and GOARCH of running binary"}}}
	commands := []struct {
		key, command string
		args         []string
	}{
		{"python", "python3", []string{"--version"}}, {"nodejs", "node", []string{"--version"}},
		{"go", "go", []string{"version"}}, {"java", "java", []string{"-version"}},
		{"git", "git", []string{"--version"}}, {"docker", "docker", []string{"--version"}},
		{"podman", "podman", []string{"--version"}}, {"nvidia_driver", "nvidia-smi", []string{"--query-gpu=driver_version", "--format=csv,noheader"}},
		{"cuda", "nvcc", []string{"--version"}},
	}
	for _, p := range commands {
		c[p.key] = probe.Command(timeout, p.command, p.args...)
	}
	return model.Report{SchemaVersion: model.SchemaVersion, GeneratedAt: time.Now().UTC().Format(time.RFC3339), ToolVersion: version, Platform: runtime.GOOS + "/" + runtime.GOARCH, Capabilities: c}
}
