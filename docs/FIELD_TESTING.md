# Community field test

AIHostCheck v0.1.0 is ready for a small, public field test on real developer
machines. This is not a request for stars, endorsements, or unreviewed system
reports. It is a request for evidence: does the tool describe a host accurately
enough for an AI coding assistant to choose relevant instructions without
guessing?

The first release already builds and runs on Windows, macOS, and Linux. Native
CI verifies those operating-system paths, but hosted runners cannot represent
the variety of shells, package managers, GPUs, permissions, and toolchains found
on real machines. Community testing is how the project can discover those gaps
without adding telemetry or collecting reports automatically.

## Who can help

Any developer can participate. The most useful test environments currently
include:

- Windows with PowerShell, Command Prompt, or Git Bash;
- macOS on Intel or Apple silicon;
- Linux across different distributions, shells, and package managers;
- hosts with multiple Python or Node.js installations;
- Docker, Podman, NVIDIA, CUDA, or local-AI tooling;
- restricted corporate, educational, virtualized, or containerized hosts.

A plain laptop with no GPU is still a useful test machine. Correctly reporting
that evidence is just as important as detecting a specialized setup.

## Test protocol

1. Download the archive for your platform from the
   [v0.1.0 release](https://github.com/raydthanh/aihostcheck/releases/tag/v0.1.0).
2. Verify the SHA-256 checksum and follow the platform instructions in
   [INSTALL.md](INSTALL.md).
3. Run the human-readable report, then run `aihostcheck --json`.
4. Review the output locally. Compare detected versions and capabilities with
   what you already know about the machine.
5. If something is incorrect, ambiguous, missing, or unexpectedly difficult,
   open a [bug report](https://github.com/raydthanh/aihostcheck/issues/new?template=bug_report.yml).

You do not need to publish a report to say that the test passed. A useful issue
contains the operating-system family, AIHostCheck version, expected behavior,
actual behavior, and the smallest redacted excerpt needed to reproduce the
problem.

If you want a reviewed result to become part of the public validation matrix,
follow the [real-device validation protocol](../validation/README.md). Its
minimal fixture records only conclusions and limitations; it never contains the
diagnostic values or evidence reviewed on your host. A fixture counts as
coverage only after CI and maintainer privacy review pass.

## What to examine

Please pay particular attention to:

- incorrect operating-system, architecture, CPU, RAM, or storage evidence;
- the active shell being reported ambiguously;
- the wrong Python, Node.js, Java, Go, or Git executable winning on `PATH`;
- installed package managers, containers, GPU tooling, or CUDA being missed;
- `unknown`, `unsupported`, `permission_denied`, and `error` states being
  represented accurately instead of treated as absence;
- installation, checksum, or documentation steps that are hard to follow.

## Privacy boundary

AIHostCheck is read-only, performs no network probe, has no telemetry, and does
not upload reports. Even so, hardware and version information can be sensitive
in some contexts.

Before sharing any excerpt:

- inspect it manually;
- remove usernames, hostnames, organization names, private paths, IP addresses,
  identifiers, and unrelated values;
- never paste credentials, environment-variable sets, private files, or an
  unreviewed full report;
- use a private [security advisory](https://github.com/raydthanh/aihostcheck/security/advisories/new)
  rather than a public issue for a vulnerability.

Do not commit a full report even after redaction. The repository accepts only
the bounded fixture fields listed in the
[validation schema](../schema/validation-fixture.schema.json), and strict tests
reject unexpected JSON fields.

The complete data boundary is documented in [PRIVACY.md](../PRIVACY.md), and
the public validation objective is tracked in
[issue #13](https://github.com/raydthanh/aihostcheck/issues/13).

## What counts as success

This field test is successful when it produces public engineering evidence,
not a particular number of stars:

- confirmed behavior across representative real hosts;
- reproducible gaps that can become regression tests;
- clearer limitations where evidence is unavailable;
- safer installation and reporting instructions;
- feedback from people who evaluated the tool rather than only viewing the
  landing page.

Thank you for testing critically. A precise report that finds a limitation is
more valuable to this stage of the project than an unqualified endorsement.
