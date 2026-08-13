# Reviewed real-device fixtures

This directory accepts only minimal fixtures created after a person has run
AIHostCheck on a real host and reviewed the full report locally. It intentionally
contains no template JSON: an example must never be mistaken for validation
evidence.

Every `*.json` file is checked by `go test ./...` using both strict decoding and
the semantic rules in `internal/validation`. Before adding one, follow
[the validation protocol](../README.md) and its privacy checklist.

The absence of a fixture means that environment has not yet been validated. Do
not infer coverage from CI builds, a synthetic example, or a directory name.
