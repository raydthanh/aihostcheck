package collector

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/raydthanh/aihostcheck/internal/model"
)

func TestCollectContract(t *testing.T) {
	r := Collect(2*time.Second, "test")
	if r.SchemaVersion != model.SchemaVersion || r.Platform == "" {
		t.Fatalf("invalid metadata: %#v", r)
	}
	for _, key := range []string{"os", "cpu", "memory", "root_storage", "shell", "python", "nodejs", "go", "java", "git", "docker", "podman", "package_manager", "gpu", "nvidia_driver", "cuda"} {
		c, ok := r.Capabilities[key]
		if !ok {
			t.Errorf("missing %s", key)
			continue
		}
		if c.Status == "" {
			t.Errorf("empty status for %s", key)
		}
	}
	if _, err := json.Marshal(r); err != nil {
		t.Fatal(err)
	}
}
