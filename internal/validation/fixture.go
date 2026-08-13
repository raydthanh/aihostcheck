package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/raydthanh/aihostcheck/internal/model"
)

const FixtureVersion = "1.0.0"

type Outcome string

const (
	Verified    Outcome = "verified"
	Incorrect   Outcome = "incorrect"
	Blocked     Outcome = "blocked"
	Unsupported Outcome = "unsupported"
	Unavailable Outcome = "unavailable"
	NotChecked  Outcome = "not_checked"
)

type Platform struct {
	OS                   string `json:"os"`
	Architecture         string `json:"architecture"`
	ExecutionEnvironment string `json:"execution_environment"`
}

type CapabilityCheck struct {
	Outcome        Outcome      `json:"outcome"`
	ObservedStatus model.Status `json:"observed_status"`
	ReviewedFields []string     `json:"reviewed_fields,omitempty"`
	Note           string       `json:"note,omitempty"`
	TrackingIssue  string       `json:"tracking_issue,omitempty"`
	RegressionTest string       `json:"regression_test,omitempty"`
}

type ReviewAttestation struct {
	ReportReviewedLocally bool `json:"report_reviewed_locally"`
	RawReportCommitted    bool `json:"raw_report_committed"`
	SensitiveDataRemoved  bool `json:"sensitive_data_removed"`
}

type Fixture struct {
	FixtureVersion      string                     `json:"fixture_version"`
	ToolVersion         string                     `json:"tool_version"`
	ReportSchemaVersion string                     `json:"report_schema_version"`
	TestedAt            string                     `json:"tested_at"`
	Platform            Platform                   `json:"platform"`
	Capabilities        map[string]CapabilityCheck `json:"capabilities"`
	Limitations         []string                   `json:"limitations,omitempty"`
	Review              ReviewAttestation          `json:"review"`
}

// DecodeFixture rejects unknown fields and trailing JSON so that repository
// fixtures stay within the intentionally small, privacy-bounded contract.
func DecodeFixture(reader io.Reader) (Fixture, error) {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	var fixture Fixture
	if err := decoder.Decode(&fixture); err != nil {
		return Fixture{}, fmt.Errorf("decode validation fixture: %w", err)
	}

	var trailing any
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		if err == nil {
			return Fixture{}, errors.New("decode validation fixture: multiple JSON values")
		}
		return Fixture{}, fmt.Errorf("decode validation fixture: trailing data: %w", err)
	}

	if err := fixture.Validate(); err != nil {
		return Fixture{}, err
	}
	return fixture, nil
}

func (fixture Fixture) Validate() error {
	if fixture.FixtureVersion != FixtureVersion {
		return fmt.Errorf("fixture_version must be %q", FixtureVersion)
	}
	if strings.TrimSpace(fixture.ToolVersion) == "" {
		return errors.New("tool_version is required")
	}
	if _, err := model.AssessSchemaCompatibility(model.SchemaVersion, fixture.ReportSchemaVersion); err != nil {
		return fmt.Errorf("report_schema_version: %w", err)
	}
	if parsed, err := time.Parse("2006-01-02", fixture.TestedAt); err != nil || parsed.Format("2006-01-02") != fixture.TestedAt {
		return errors.New("tested_at must be a calendar date in YYYY-MM-DD format")
	}

	if !oneOf(fixture.Platform.OS, "windows", "macos", "linux") {
		return errors.New("platform.os must be windows, macos, or linux")
	}
	if !oneOf(fixture.Platform.Architecture, "amd64", "arm64") {
		return errors.New("platform.architecture must be amd64 or arm64")
	}
	if !oneOf(fixture.Platform.ExecutionEnvironment, "physical_host", "virtual_machine", "wsl", "container", "other") {
		return errors.New("platform.execution_environment is invalid")
	}
	if len(fixture.Capabilities) == 0 {
		return errors.New("at least one capability result is required")
	}

	for name, check := range fixture.Capabilities {
		if strings.TrimSpace(name) == "" {
			return errors.New("capability names must not be empty")
		}
		if err := validateCapabilityCheck(name, check); err != nil {
			return err
		}
	}

	for index, limitation := range fixture.Limitations {
		if strings.TrimSpace(limitation) == "" {
			return fmt.Errorf("limitations[%d] must not be empty", index)
		}
	}
	if !fixture.Review.ReportReviewedLocally {
		return errors.New("review.report_reviewed_locally must be true")
	}
	if fixture.Review.RawReportCommitted {
		return errors.New("review.raw_report_committed must be false")
	}
	if !fixture.Review.SensitiveDataRemoved {
		return errors.New("review.sensitive_data_removed must be true")
	}
	return nil
}

func validateCapabilityCheck(name string, check CapabilityCheck) error {
	if !oneOf(string(check.Outcome), string(Verified), string(Incorrect), string(Blocked), string(Unsupported), string(Unavailable), string(NotChecked)) {
		return fmt.Errorf("capability %q has invalid outcome %q", name, check.Outcome)
	}
	if !oneOf(string(check.ObservedStatus), string(model.Detected), string(model.NotDetected), string(model.Unknown), string(model.Unsupported), string(model.Error), string(model.PermissionDenied)) {
		return fmt.Errorf("capability %q has invalid observed_status %q", name, check.ObservedStatus)
	}

	seenFields := make(map[string]struct{}, len(check.ReviewedFields))
	for _, field := range check.ReviewedFields {
		if !oneOf(field, "status", "value", "evidence", "details") {
			return fmt.Errorf("capability %q has invalid reviewed field %q", name, field)
		}
		if _, exists := seenFields[field]; exists {
			return fmt.Errorf("capability %q repeats reviewed field %q", name, field)
		}
		seenFields[field] = struct{}{}
	}

	if (check.Outcome == Verified || check.Outcome == Incorrect) && len(check.ReviewedFields) == 0 {
		return fmt.Errorf("capability %q must identify at least one reviewed field when %s", name, check.Outcome)
	}
	if check.Outcome == Incorrect {
		if !validTrackingIssue(check.TrackingIssue) {
			return fmt.Errorf("capability %q requires a canonical AIHostCheck tracking_issue when incorrect", name)
		}
		if !validRegressionTestPath(check.RegressionTest) {
			return fmt.Errorf("capability %q requires a repository-relative *_test.go regression_test when incorrect", name)
		}
	}
	if check.Outcome != Verified && check.Outcome != Incorrect && strings.TrimSpace(check.Note) == "" {
		return fmt.Errorf("capability %q requires a note for outcome %q", name, check.Outcome)
	}
	return nil
}

func validTrackingIssue(value string) bool {
	const prefix = "https://github.com/raydthanh/aihostcheck/issues/"
	if !strings.HasPrefix(value, prefix) {
		return false
	}
	number, err := strconv.Atoi(strings.TrimPrefix(value, prefix))
	return err == nil && number > 0
}

func validRegressionTestPath(value string) bool {
	return value != "" &&
		!strings.HasPrefix(value, "/") &&
		!strings.Contains(value, "\\") &&
		path.Clean(value) == value &&
		strings.HasSuffix(value, "_test.go")
}

func oneOf(value string, allowed ...string) bool {
	for _, candidate := range allowed {
		if value == candidate {
			return true
		}
	}
	return false
}
