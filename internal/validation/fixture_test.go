package validation

import (
	"strings"
	"testing"
)

const validFixture = `{
  "fixture_version": "1.0.0",
  "tool_version": "v0.1.0",
  "report_schema_version": "1.0.0",
  "tested_at": "2026-08-13",
  "platform": {
    "os": "linux",
    "architecture": "amd64",
    "execution_environment": "physical_host"
  },
  "capabilities": {
    "os": {
      "outcome": "verified",
      "observed_status": "detected",
      "reviewed_fields": ["status", "value", "evidence"]
    },
    "cuda": {
      "outcome": "unavailable",
      "observed_status": "not_detected",
      "note": "No compatible hardware was available for an independent check."
    }
  },
  "limitations": ["GPU-specific behavior was not evaluated."],
  "review": {
    "report_reviewed_locally": true,
    "raw_report_committed": false,
    "sensitive_data_removed": true
  }
}`

func TestDecodeFixture(t *testing.T) {
	fixture, err := DecodeFixture(strings.NewReader(validFixture))
	if err != nil {
		t.Fatalf("DecodeFixture() error = %v", err)
	}
	if fixture.Platform.OS != "linux" {
		t.Fatalf("DecodeFixture().Platform.OS = %q, want linux", fixture.Platform.OS)
	}
}

func TestDecodeFixtureRejectsUnknownFields(t *testing.T) {
	input := strings.Replace(validFixture, `"tool_version": "v0.1.0",`, `"tool_version": "v0.1.0", "hostname": "private-host",`, 1)
	if _, err := DecodeFixture(strings.NewReader(input)); err == nil {
		t.Fatal("DecodeFixture() accepted an unknown privacy-sensitive field")
	}
}

func TestIncorrectResultRequiresIssueAndRegressionTest(t *testing.T) {
	input := strings.Replace(
		validFixture,
		`"outcome": "verified",`,
		`"outcome": "incorrect",`,
		1,
	)
	if _, err := DecodeFixture(strings.NewReader(input)); err == nil || !strings.Contains(err.Error(), "tracking_issue") {
		t.Fatalf("DecodeFixture() error = %v, want missing tracking_issue", err)
	}
}

func TestIncorrectResultAcceptsCanonicalReferences(t *testing.T) {
	input := strings.Replace(
		validFixture,
		`"outcome": "verified",`,
		`"outcome": "incorrect",
      "tracking_issue": "https://github.com/raydthanh/aihostcheck/issues/42",
      "regression_test": "internal/collector/platform_windows_test.go",`,
		1,
	)
	if _, err := DecodeFixture(strings.NewReader(input)); err != nil {
		t.Fatalf("DecodeFixture() error = %v, want valid incorrect result", err)
	}
}

func TestUnavailableResultRequiresLimitationNote(t *testing.T) {
	input := strings.Replace(validFixture, `,
      "note": "No compatible hardware was available for an independent check."`, "", 1)
	if _, err := DecodeFixture(strings.NewReader(input)); err == nil || !strings.Contains(err.Error(), "requires a note") {
		t.Fatalf("DecodeFixture() error = %v, want missing limitation note", err)
	}
}

func TestReviewAttestationPreventsRawReportCommit(t *testing.T) {
	input := strings.Replace(validFixture, `"raw_report_committed": false`, `"raw_report_committed": true`, 1)
	if _, err := DecodeFixture(strings.NewReader(input)); err == nil || !strings.Contains(err.Error(), "must be false") {
		t.Fatalf("DecodeFixture() error = %v, want raw-report rejection", err)
	}
}

func TestDecodeFixtureRejectsTrailingDocument(t *testing.T) {
	if _, err := DecodeFixture(strings.NewReader(validFixture + `{}`)); err == nil {
		t.Fatal("DecodeFixture() accepted multiple JSON documents")
	}
}
