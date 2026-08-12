package probe

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"strings"
	"time"

	"github.com/raydthanh/aihostcheck/internal/model"
)

const outputLimit = 32 * 1024

type limitedBuffer struct{ bytes.Buffer }

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

// Command runs an executable directly, never through a shell. Callers supply fixed arguments.
func Command(timeout time.Duration, name string, args ...string) model.Capability {
	path, err := exec.LookPath(name)
	if err != nil {
		return model.Capability{Status: model.NotDetected, Evidence: []model.Evidence{{Source: "executable_lookup", Detail: name + " not found in PATH"}}}
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var out limitedBuffer
	cmd := exec.CommandContext(ctx, path, args...)
	cmd.Stdout, cmd.Stderr = &out, &out
	err = cmd.Run()
	value := strings.TrimSpace(out.String())
	if ctx.Err() == context.DeadlineExceeded {
		return model.Capability{Status: model.Error, Evidence: []model.Evidence{{Source: "command", Detail: name + " timed out"}}}
	}
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return model.Capability{Status: model.NotDetected}
		}
		return model.Capability{Status: model.Error, Evidence: []model.Evidence{{Source: "command", Detail: name + " exited unsuccessfully"}}}
	}
	return model.Capability{Status: model.Detected, Value: firstLine(value), Evidence: []model.Evidence{{Source: "command", Detail: name + " completed successfully"}}}
}

func firstLine(s string) string {
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		return strings.TrimSpace(s[:i])
	}
	return s
}
