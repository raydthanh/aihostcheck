package validation

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRepositoryFixtures(t *testing.T) {
	paths, err := filepath.Glob(filepath.Join("..", "..", "validation", "fixtures", "*.json"))
	if err != nil {
		t.Fatalf("glob validation fixtures: %v", err)
	}

	for _, path := range paths {
		path := path
		t.Run(filepath.Base(path), func(t *testing.T) {
			file, err := os.Open(path)
			if err != nil {
				t.Fatalf("open fixture: %v", err)
			}
			defer file.Close()

			if _, err := DecodeFixture(file); err != nil {
				t.Fatalf("invalid repository fixture: %v", err)
			}
		})
	}
}

func TestFixtureSchemaIsValidJSON(t *testing.T) {
	path := filepath.Join("..", "..", "schema", "validation-fixture.schema.json")
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("open fixture schema: %v", err)
	}
	defer file.Close()

	var document map[string]any
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&document); err != nil {
		t.Fatalf("decode fixture schema: %v", err)
	}
}
