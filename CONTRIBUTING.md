# Contributing

AIHostCheck welcomes focused bug reports, real-device results, documentation improvements, and new evidence-backed collectors. For substantial changes, open an issue first so the evidence source, cross-OS behavior, report contract, and privacy impact can be agreed before implementation.

Use Go 1.22 or newer. Keep the dependency-free standard-library baseline unless a proposal demonstrates a clear security or portability benefit. Before submitting:

```sh
gofmt -w .
go test ./...
go vet ./...
```

Explain the evidence source and privacy impact of every new capability. Add OS-specific implementations under build tags, preserve all six status meanings, update the JSON Schema when changing the contract, and avoid speculative detection. By contributing, you agree that your work is licensed under Apache-2.0.

Bug reports must not contain credentials, private paths, or an unreviewed diagnostic report. Include only the smallest redacted output needed to reproduce a problem. Security vulnerabilities belong in a [private security advisory](https://github.com/raydthanh/aihostcheck/security/advisories/new), not a public issue.
