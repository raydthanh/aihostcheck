package model

import (
	"encoding/json"
	"testing"
)

func TestAssessSchemaCompatibility(t *testing.T) {
	tests := []struct {
		name       string
		consumer   string
		report     string
		expected   SchemaCompatibility
	}{
		{name: "equal", consumer: "1.2.0", report: "1.2.0", expected: SchemaExact},
		{name: "same minor newer patch", consumer: "1.2.0", report: "1.2.4", expected: SchemaExact},
		{name: "older minor", consumer: "1.2.0", report: "1.1.7", expected: SchemaOlderMinor},
		{name: "newer minor", consumer: "1.2.0", report: "1.3.0", expected: SchemaNewerMinor},
		{name: "different major", consumer: "1.2.0", report: "2.0.0", expected: SchemaIncompatibleMajor},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := AssessSchemaCompatibility(test.consumer, test.report)
			if err != nil {
				t.Fatalf("AssessSchemaCompatibility() error = %v", err)
			}
			if actual != test.expected {
				t.Fatalf("AssessSchemaCompatibility() = %q, want %q", actual, test.expected)
			}
		})
	}
}

func TestAssessSchemaCompatibilityRejectsInvalidVersions(t *testing.T) {
	for _, version := range []string{"", "1", "1.0", "v1.0.0", "+1.0.0", "1.01.0", "1.0.0-rc.1"} {
		t.Run(version, func(t *testing.T) {
			if _, err := AssessSchemaCompatibility(SchemaVersion, version); err == nil {
				t.Fatalf("AssessSchemaCompatibility() accepted invalid version %q", version)
			}
		})
	}
}

func TestReportSerializesIndependentProvenance(t *testing.T) {
	report := Report{
		SchemaVersion: "1.0.0",
		GeneratedAt:   "2026-08-13T00:00:00Z",
		ToolVersion:   "0.2.0",
		Platform:      "linux/amd64",
		Capabilities: map[string]Capability{
			"os": {
				Status: Detected,
				Value:  "linux",
				Evidence: []Evidence{{
					Source: "go_runtime",
					Detail: "GOOS and GOARCH of running binary",
				}},
			},
		},
	}

	encoded, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	var document map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &document); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	for _, field := range []string{"schema_version", "generated_at", "tool_version", "platform", "capabilities"} {
		if _, ok := document[field]; !ok {
			t.Errorf("serialized report is missing provenance field %q", field)
		}
	}
}
