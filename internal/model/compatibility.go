package model

import (
	"fmt"
	"strconv"
	"strings"
)

// SchemaCompatibility describes how a report schema relates to the schema a
// consumer implements. It does not imply that individual capability evidence
// is sufficient for a particular task.
type SchemaCompatibility string

const (
	SchemaExact             SchemaCompatibility = "exact"
	SchemaOlderMinor        SchemaCompatibility = "compatible_older_minor"
	SchemaNewerMinor        SchemaCompatibility = "requires_tolerant_reader"
	SchemaIncompatibleMajor SchemaCompatibility = "incompatible_major"
)

// AssessSchemaCompatibility applies AIHostCheck's schema-version policy.
// Versions must be stable semantic versions in MAJOR.MINOR.PATCH form.
func AssessSchemaCompatibility(consumerVersion, reportVersion string) (SchemaCompatibility, error) {
	consumer, err := parseSchemaVersion(consumerVersion)
	if err != nil {
		return "", fmt.Errorf("consumer schema version: %w", err)
	}
	report, err := parseSchemaVersion(reportVersion)
	if err != nil {
		return "", fmt.Errorf("report schema version: %w", err)
	}

	if consumer.major != report.major {
		return SchemaIncompatibleMajor, nil
	}
	if consumer.minor > report.minor {
		return SchemaOlderMinor, nil
	}
	if consumer.minor < report.minor {
		return SchemaNewerMinor, nil
	}
	return SchemaExact, nil
}

type schemaVersion struct {
	major int
	minor int
	patch int
}

func parseSchemaVersion(value string) (schemaVersion, error) {
	parts := strings.Split(value, ".")
	if len(parts) != 3 {
		return schemaVersion{}, fmt.Errorf("%q must use MAJOR.MINOR.PATCH", value)
	}

	numbers := make([]int, len(parts))
	for i, part := range parts {
		if part == "" || (len(part) > 1 && part[0] == '0') {
			return schemaVersion{}, fmt.Errorf("%q is not a stable semantic version", value)
		}
		for _, character := range part {
			if character < '0' || character > '9' {
				return schemaVersion{}, fmt.Errorf("%q is not a stable semantic version", value)
			}
		}
		n, err := strconv.Atoi(part)
		if err != nil || n < 0 {
			return schemaVersion{}, fmt.Errorf("%q is not a stable semantic version", value)
		}
		numbers[i] = n
	}

	return schemaVersion{major: numbers[0], minor: numbers[1], patch: numbers[2]}, nil
}
