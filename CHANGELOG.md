# Changelog

All notable changes to AIHostCheck will be documented in this file. The project follows [Semantic Versioning](https://semver.org/) for releases and versions its machine-readable report contract separately through `schema_version`.

## Unreleased

### Fixed

- Increase the default per-command probe timeout from 3 to 15 seconds so a
  cold, non-interactive Windows PowerShell/CIM GPU inventory does not produce
  a false `error`; the explicit `--timeout` override remains available.

### Validation

- Record the first privacy-safe real-device fixture from a Windows AMD64
  physical host, linking the confirmed GPU timeout to its public issue and
  regression test without retaining the report or GPU identity.

## [0.2.0] - 2026-08-13

### Added

- Safe AI-consumer workflow guidance and structured community issue and pull-request templates.
- Report provenance and schema compatibility rules with a tested reference implementation and migration requirements.
- A versioned, privacy-safe real-device validation fixture contract, strict
  repository checks, and a maintainer review protocol that excludes raw reports.
- Native installation-path checks for Windows, macOS, and Linux, plus a public
  distribution/signing policy and keyless provenance for future releases.

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

[Unreleased]: https://github.com/raydthanh/aihostcheck/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/raydthanh/aihostcheck/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/raydthanh/aihostcheck/releases/tag/v0.1.0
