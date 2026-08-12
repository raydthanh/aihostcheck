package probe

import (
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/raydthanh/aihostcheck/internal/model"
)

func TestCommandNotDetected(t *testing.T) {
	r := Command(time.Second, "aihostcheck-command-that-does-not-exist")
	if r.Status != model.NotDetected {
		t.Fatalf("status = %q", r.Status)
	}
}

func TestCommandOutput(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("test helper differs on Windows")
	}
	r := Command(time.Second, "printf", "first\nsecond")
	if r.Status != model.Detected || r.Value != "first" {
		t.Fatalf("result = %#v", r)
	}
}

func TestLimitedBuffer(t *testing.T) {
	var b limitedBuffer
	p := []byte(strings.Repeat("x", outputLimit+100))
	n, err := b.Write(p)
	if err != nil || n != len(p) || b.Len() != outputLimit {
		t.Fatalf("n=%d len=%d err=%v", n, b.Len(), err)
	}
}
