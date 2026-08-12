# AIHostCheck

[![CI](https://github.com/raydthanh/aihostcheck/actions/workflows/ci.yml/badge.svg)](https://github.com/raydthanh/aihostcheck/actions/workflows/ci.yml)
[![Website CI](https://github.com/raydthanh/aihostcheck/actions/workflows/website.yml/badge.svg)](https://github.com/raydthanh/aihostcheck/actions/workflows/website.yml)
[![Release](https://img.shields.io/github/v/release/raydthanh/aihostcheck)](https://github.com/raydthanh/aihostcheck/releases/latest)
[![License](https://img.shields.io/github/license/raydthanh/aihostcheck)](LICENSE)

AIHostCheck is a small, read-only, cross-OS diagnostic layer that gives people, GPTs, and AI coding agents evidence about the host **before** they propose commands. It is a native Go CLI—not a system-info demo and not a website.

> [aihostcheck.bond](https://aihostcheck.bond) is only a companion site. This GitHub repository is the source of truth for code, releases, schemas, and security policy.

## Quick start

Download the standalone archive for Windows, macOS, or Linux from [GitHub Releases](https://github.com/raydthanh/aihostcheck/releases). See the [installation and checksum guide](docs/INSTALL.md) for platform-specific steps.

To build the current source with Go 1.22 or newer:

```sh
go build -o aihostcheck ./cmd/aihostcheck
./aihostcheck          # human-readable table
./aihostcheck --json   # versioned machine contract
./aihostcheck --version
```

No probe uses the network. AIHostCheck has no telemetry and never uploads a report.

**Project status:** v0.1.0 is a working early release, not a claim of established adoption. The [public roadmap](docs/ROADMAP.md) separates shipped evidence, planned work, exploratory ideas, and explicit non-goals. Roadmap items move through the issue tracker and count as complete only when their code, tests, documentation, or release evidence is public.

**Community field test:** developers can follow the bounded [field-test protocol](docs/FIELD_TESTING.md) to validate the release on a real machine. Participation requires no account, telemetry, or full report submission; reviewed findings become public engineering evidence and regression tests.

Before giving a JSON report to GPT or another AI agent, inspect the file and share it as environment context—not as a command to execute. The [AI workflow guide](docs/USING_WITH_AI.md) includes a safe, reusable prompt and explains how an agent should interpret evidence and unknown states.

## What it checks

The initial collector covers OS/version/architecture, logical CPU count, physical RAM, system storage, shell evidence, Python, Node.js, Go, Java, Git, Docker, Podman, common package managers, GPU visibility, NVIDIA driver, and CUDA compiler. Platform collectors are compiled natively for Linux, macOS, and Windows. Python and package-manager probes use platform-specific candidates, and multiple detected package managers are preserved for AI consumers.

## Safety and privacy

Collection is read-only. The tool intentionally does **not** collect username, hostname, IP addresses, credentials, full environment-variable sets, personal files, or process command lines. On Unix, the report keeps only the basename of `SHELL`; command output is scrubbed of the current home-directory prefix. Commands have fixed arguments, time out (3 seconds by default), and capture at most 32 KiB. Windows GPU inventory uses one fixed, non-interactive, read-only CIM query with no user input. Review output before sharing it: versions and hardware may still be sensitive in your context. See [PRIVACY.md](PRIVACY.md) and [SECURITY.md](SECURITY.md).

## Report contract

Every JSON document contains `schema_version`, UTC generation time, tool version, target platform, and a capability map. Each capability has one of `detected`, `not_detected`, `unknown`, `unsupported`, `error`, or `permission_denied`; detected claims include evidence identifying the source. The normative schema is [`schema/report.schema.json`](schema/report.schema.json). Additive fields require a minor schema version; removals or semantic changes require a major version.

## Scope and limitations

This foundation favors trustworthy presence/version evidence over exhaustive inventory. If a required inventory utility is unavailable, GPU status is `unknown` rather than a false claim that no GPU exists. It does not contact Docker/Podman daemons. Windows and macOS behavior is continuously compiled and tested by native GitHub-hosted runners, but hardware-specific paths need broader real-device testing.

See the [roadmap](docs/ROADMAP.md) for measurable next outcomes, [ARCHITECTURE.md](ARCHITECTURE.md) to extend collectors, [CONTRIBUTING.md](CONTRIBUTING.md) to contribute, and [CHANGELOG.md](CHANGELOG.md) for project history. Maintainer releases follow the documented [release process](docs/RELEASING.md). Licensed under Apache-2.0.
