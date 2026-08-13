# Real-device validation protocol

This directory turns local, manual field testing into small public evidence
without committing a diagnostic report. A fixture records what a tester checked
and the outcome; it omits the values and evidence that could fingerprint the
host.

The fixture contract is versioned independently at
[`schema/validation-fixture.schema.json`](../schema/validation-fixture.schema.json).
Repository fixtures are also decoded strictly and checked semantically by
`internal/validation` on Windows, macOS, and Linux CI.

## What a fixture may contain

- AIHostCheck and report-schema versions;
- test date, OS family, architecture, and broad execution-environment class;
- capability names, the status emitted by AIHostCheck, fields reviewed locally,
  and one of the outcomes below;
- short limitations that contain no machine identifiers;
- links to a public issue and regression test when a detection was incorrect.

| Outcome | Meaning |
| --- | --- |
| `verified` | The named fields were compared locally with independent evidence and were correct. |
| `incorrect` | At least one reviewed field was wrong; a tracking issue and regression test are required. |
| `blocked` | Permissions or policy prevented a conclusive check. |
| `unsupported` | AIHostCheck explicitly does not support this check in the tested environment. |
| `unavailable` | Hardware, tooling, or independent evidence was unavailable to the tester. |
| `not_checked` | The tester deliberately did not evaluate this capability. |

An outcome describes the human validation result. `observed_status` records the
machine-readable status produced by AIHostCheck. They are separate so that an
AI consumer does not confuse unavailable validation evidence with
`not_detected`.

## Submission steps

1. Run the released AIHostCheck binary and inspect its JSON report only on the
   test host.
2. Independently check the capabilities you can verify. Do not publish the
   commands or their output unless a minimal redacted excerpt is needed for a
   bug report.
3. Create a fixture from the structure below. Include only reviewed conclusions;
   do not copy values, evidence details, or free-form report data.
4. For every `incorrect` result, first open a redacted issue and add a regression
   test that fails without the fix. Record both references in the fixture.
5. Save the fixture under `validation/fixtures/` with a non-identifying name,
   run `gofmt -w .`, `go test ./...`, and `go vet ./...`, then open a pull
   request.

This example is documentation only and does not count as device coverage:

```json
{
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
}
```

## Privacy review checklist

Before committing a fixture, verify every item:

- [ ] The full report stayed on the test host and was manually reviewed there.
- [ ] No raw or redacted full report is included.
- [ ] There are no usernames, hostnames, organization names, IP or MAC
      addresses, device identifiers, serial numbers, credentials, personal
      paths, environment-variable sets, or process command lines.
- [ ] Notes and limitations describe only why validation was incomplete; they do
      not reproduce collected values or identify the tester or machine.
- [ ] `verified` is used only for fields checked against independent evidence.
- [ ] Every incorrect detection has a public issue and regression test.
- [ ] Missing evidence remains `blocked`, `unsupported`, `unavailable`, or
      `not_checked`; it is not rewritten as a successful absence check.

Maintainers must review the fixture as potentially sensitive even when all
attestation fields are set correctly. If uncertain, leave the result in the
issue without committing a fixture.
