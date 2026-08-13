# AIHostCheck roadmap

This roadmap describes intended outcomes, not promises of adoption, funding, or delivery dates. Work moves forward when it preserves AIHostCheck's core guarantees: read-only collection, explicit uncertainty, reviewable evidence, cross-OS behavior, and no telemetry.

The public issue tracker is the execution layer for this document. A roadmap item is complete only when its acceptance evidence is visible in code, tests, documentation, or a release.

## Shipped foundation — v0.1

- Native Go CLI for Windows, macOS, and Linux.
- Human-readable terminal output and a versioned JSON report.
- Evidence-backed capability states, including `unknown`, `unsupported`, and `permission_denied`.
- Read-only probes for host, runtime, package-manager, container, GPU, NVIDIA, and CUDA signals.
- Six release archives with SHA-256 checksums.
- Native CI on Windows, macOS, and Linux.
- Privacy, security, architecture, installation, release, and AI-workflow documentation.

## Next release — reliability and distribution

The next release should make the existing foundation easier to trust on real machines and easier to install without weakening its privacy boundary.

### Real-device validation matrix

Validate reports on representative Windows, macOS, and Linux hosts beyond hosted CI runners.

Completion evidence:

- A documented, privacy-safe fixture format for validation conclusions; raw or
  redacted diagnostic reports are not repository fixtures.
- Coverage across supported operating systems and AMD64/ARM64 where hardware is available.
- Regression tests for every confirmed incorrect detection.
- Limitations recorded when hardware or permissions prevent a definitive result.

The fixture contract and automated validation are shipped. Device coverage is
still pending and will be counted only from reviewed real-host evidence in
`validation/fixtures/`; hosted CI and documentation examples do not count.

### Safer, easier installation

Reduce the gap between downloading a release and verifying that it is authentic.

Completion evidence:

- Reproducible release instructions remain public.
- Checksums stay mandatory for all native packages.
- At least one low-friction installation path is evaluated per operating-system family.
- Artifact signing is added only after a maintainable key and identity process is documented.

### Report provenance and compatibility

Make it easier for an AI consumer to decide whether a report is current and compatible.

Completion evidence:

- Tool and schema versions remain independently visible.
- Compatibility rules are covered by tests and examples.
- Schema changes include migration notes and representative reports.

## Following release — AI workflow integration

This phase turns the report contract into reusable workflow components without requiring a hosted account or uploading reports.

### Local review and redaction assistance

- Offer a local-first way to inspect a report before sharing it.
- Highlight fields that can still be sensitive in a user's context.
- Keep report processing on the user's machine by default.

### Agent adapters and capability profiles

- Define small adapters that transform a report into bounded context for AI coding workflows.
- Document how consumers must handle unknown and failed evidence.
- Add capability profiles for common tasks such as container development, Python projects, and GPU workloads.

### Deeper AI and accelerator diagnostics

- Improve evidence for GPU visibility, NVIDIA driver/toolkit compatibility, and common local AI runtimes.
- Prefer documented APIs and fixed read-only commands.
- Never infer that hardware is absent when collection is unsupported or blocked.

## Exploration, not commitments

These ideas require user evidence before they enter a release plan:

- Editor and coding-agent integrations.
- Support-bundle generation for maintainers and remote troubleshooting.
- A plugin model for specialized diagnostic packs.
- Additional accelerator vendors and AI runtime ecosystems.

## Explicit non-goals for the current stages

- No remote command execution.
- No background monitoring or telemetry.
- No automatic upload of diagnostic reports.
- No account requirement or hosted report database.
- No tool that silently installs, removes, or reconfigures software.
- No claim that a capability is absent when evidence is merely unavailable.

## How infrastructure support helps

Infrastructure credits and developer tooling would be used only for work that produces public project evidence:

| Support area | Project use | Public evidence |
| --- | --- | --- |
| Cross-platform compute | Broader OS and architecture validation | Test matrix, fixtures, and regression results |
| Code signing and secure delivery | Verifiable native artifacts | Documented signing policy and signed releases |
| Hosting and observability | Reliable documentation and website | Public status, error fixes, and deployment history |
| AI developer tooling | Test adapters against multiple agent workflows | Open adapters, examples, and compatibility notes |

Support would accelerate validation and distribution; it would not be presented as product adoption, a customer relationship, or an endorsement.

## How priorities are decided

Roadmap work is promoted when it satisfies at least one of these signals:

1. A reproducible user report exposes an incorrect or ambiguous diagnosis.
2. A common AI workflow cannot safely choose commands from the current report.
3. Real-device validation reveals an OS or hardware coverage gap.
4. A contributor proposes a privacy-bounded evidence source with cross-OS value.

Feature requests should use the repository's issue form and explain the problem, evidence source, privacy impact, and expected AI-consumer behavior.
