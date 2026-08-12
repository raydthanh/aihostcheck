package collector

import (
	"encoding/json"
	"os"
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
		if len(c.Evidence) == 0 {
			t.Errorf("missing evidence for %s", key)
		}
		if c.Status == model.Detected && c.Value == "" {
			t.Errorf("detected capability %s has no value", key)
		}
	}
	if _, err := json.Marshal(r); err != nil {
		t.Fatal(err)
	}
}

func TestSchemaRequiresCoreCapabilities(t *testing.T) {
	b, err := os.ReadFile("../../schema/report.schema.json")
	if err != nil {
		t.Fatal(err)
	}
	var schema struct {
		Properties struct {
			Capabilities struct {
				Required []string `json:"required"`
			} `json:"capabilities"`
		} `json:"properties"`
	}
	if err := json.Unmarshal(b, &schema); err != nil {
		t.Fatalf("invalid JSON Schema document: %v", err)
	}
	required := make(map[string]bool, len(schema.Properties.Capabilities.Required))
	for _, key := range schema.Properties.Capabilities.Required {
		required[key] = true
	}
	for _, key := range []string{"os", "os_version", "cpu", "memory", "root_storage", "shell", "python", "nodejs", "package_manager", "gpu", "nvidia_driver", "cuda"} {
		if !required[key] {
			t.Errorf("schema does not require core capability %s", key)
		}
	}
}
