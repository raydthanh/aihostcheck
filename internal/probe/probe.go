package probe

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/raydthanh/aihostcheck/internal/model"
)

const outputLimit = 32 * 1024

type limitedBuffer struct{ bytes.Buffer }

// Result is the bounded output and status of a directly executed command.
// Output is redacted before it leaves this package.
type Result struct {
	Status   model.Status
	Output   string
	Evidence []model.Evidence
}

func (b *limitedBuffer) Write(p []byte) (int, error) {
	n := len(p)
	remaining := outputLimit - b.Len()
	if remaining > 0 {
		if len(p) > remaining {
			p = p[:remaining]
		}
		_, _ = b.Buffer.Write(p)
	}
	return n, nil
}

// Run executes an executable directly. Callers must provide fixed arguments and
// must never pass user-controlled values.
func Run(timeout time.Duration, name string, args ...string) Result {
	path, err := exec.LookPath(name)
	if err != nil {
		return Result{Status: model.NotDetected, Evidence: []model.Evidence{{Source: "executable_lookup", Detail: name + " not found in PATH"}}}
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var out limitedBuffer
	cmd := exec.CommandContext(ctx, path, args...)
	cmd.Stdout, cmd.Stderr = &out, &out
	err = cmd.Run()
	value := redactHome(strings.TrimSpace(out.String()))
	if ctx.Err() == context.DeadlineExceeded {
		return Result{Status: model.Error, Evidence: []model.Evidence{{Source: "command", Detail: name + " timed out"}}}
	}
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return Result{Status: model.NotDetected, Evidence: []model.Evidence{{Source: "executable_lookup", Detail: name + " disappeared before execution"}}}
		}
		if errors.Is(err, os.ErrPermission) {
			return Result{Status: model.PermissionDenied, Evidence: []model.Evidence{{Source: "command", Detail: name + " could not be executed due to permissions"}}}
		}
		return Result{Status: model.Error, Output: value, Evidence: []model.Evidence{{Source: "command", Detail: name + " exited unsuccessfully"}}}
	}
	return Result{Status: model.Detected, Output: value, Evidence: []model.Evidence{{Source: "command", Detail: name + " completed successfully"}}}
}

// Command converts a successful command into a single-line capability value.
func Command(timeout time.Duration, name string, args ...string) model.Capability {
	r := Run(timeout, name, args...)
	if r.Status != model.Detected {
		return model.Capability{Status: r.Status, Evidence: r.Evidence}
	}
	value := firstNonEmptyLine(r.Output)
	if value == "" {
		return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "command", Detail: name + " completed without usable output"}}}
	}
	return model.Capability{Status: model.Detected, Value: value, Evidence: r.Evidence}
}

func firstNonEmptyLine(s string) string {
	for _, line := range strings.Split(s, "\n") {
		if line = strings.TrimSpace(line); line != "" {
			return line
		}
	}
	return ""
}

func redactHome(s string) string {
	home, err := os.UserHomeDir()
	if err == nil && home != "" {
		s = strings.ReplaceAll(s, home, "~")
	}
	return s
}
