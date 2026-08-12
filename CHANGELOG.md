# Changelog

All notable changes to AIHostCheck will be documented in this file. The project follows [Semantic Versioning](https://semver.org/) for releases and versions its machine-readable report contract separately through `schema_version`.

## Unreleased

## [0.1.0] - 2026-08-12

### Added

- Initial read-only diagnostic core for Windows, macOS, and Linux.
- Human-readable and versioned JSON reports for AI-assisted development workflows.
- Privacy-aware evidence collection for OS, CPU, RAM, storage, shells, runtimes, package managers, containers, GPU, NVIDIA driver, and CUDA.
- Cross-OS CI, report schema, contributor documentation, and security/privacy policies.
- Automated release packaging for six OS/architecture combinations with SHA-256 checksums.
- Controlled release publishing that verifies the complete asset bundle before granting write access.

### Fixed

- Release publication explicitly targets the repository and does not depend on local Git metadata.

[Unreleased]: https://github.com/raydthanh/aihostcheck/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/raydthanh/aihostcheck/releases/tag/v0.1.0
