# AIHostCheck

AIHostCheck is a small, read-only, cross-OS diagnostic layer that gives people, GPTs, and AI coding agents evidence about the host **before** they propose commands. It is a native Go CLI—not a system-info demo and not a website.

> [aihostcheck.bond](https://aihostcheck.bond) is only a companion site. This GitHub repository is the source of truth for code, releases, schemas, and security policy.

## Quick start

```sh
go build -o aihostcheck ./cmd/aihostcheck
./aihostcheck          # stable human-readable table
./aihostcheck --json   # versioned machine contract
```

No probe uses the network. AIHostCheck has no telemetry and never uploads a report.

## What it checks

The initial collector covers OS/version/architecture, logical CPU count, physical RAM, root storage, shell, Python, Node.js, Go, Java, Git, Docker, Podman, common package managers, GPU visibility, NVIDIA driver, and CUDA compiler. Platform collectors are compiled natively for Linux, macOS, and Windows. A missing executable is reported rather than guessed.

## Safety and privacy

Collection is read-only. The tool intentionally does **not** collect username, hostname, IP addresses, credentials, full environment-variable sets, personal files, or process command lines. The only environment value read is `SHELL` on Unix. Commands have fixed arguments, execute directly without a command shell, time out (3 seconds by default), and capture at most 32 KiB. Review output before sharing it: versions and hardware may still be sensitive in your context. See [PRIVACY.md](PRIVACY.md) and [SECURITY.md](SECURITY.md).

## Report contract

Every JSON document contains `schema_version`, UTC generation time, tool version, target platform, and a capability map. Each capability has one of `detected`, `not_detected`, `unknown`, `unsupported`, `error`, or `permission_denied`; detected claims include evidence identifying the source. The normative schema is [`schema/report.schema.json`](schema/report.schema.json). Additive fields require a minor schema version; removals or semantic changes require a major version.

## Scope and limitations

This foundation favors trustworthy presence/version evidence over exhaustive inventory. GPU discovery uses native tools available on the host and may be `not_detected` when optional utilities are absent. It does not contact Docker/Podman daemons. Windows and macOS behavior is continuously compiled and tested by native GitHub-hosted runners, but hardware-specific paths need broader real-device testing.

See [ARCHITECTURE.md](ARCHITECTURE.md) to extend collectors and [CONTRIBUTING.md](CONTRIBUTING.md) to contribute. Licensed under Apache-2.0.
