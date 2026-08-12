# Contributing

Use Go 1.22 or newer. Keep the dependency-free standard-library baseline unless a proposal demonstrates a clear security or portability benefit. Before submitting:

```sh
gofmt -w .
go test ./...
go vet ./...
```

Explain the evidence source and privacy impact of every new capability. Add OS-specific implementations under build tags, preserve all six status meanings, update the JSON Schema when changing the contract, and avoid speculative detection. By contributing, you agree that your work is licensed under Apache-2.0.
