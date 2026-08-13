# Distribution and signing policy

This document records the distribution paths AIHostCheck has evaluated and the
trust model for release artifacts. It is a maintenance decision, not a claim
that every listed channel exists today.

## Current decision

GitHub Releases remains the only binary distribution channel for the next
release. Each release contains six native archives and a mandatory
`checksums.txt`. The release workflow tests the complete user path on native
Windows, macOS, and Linux runners: select the matching archive, verify SHA-256,
extract it, run `--version`, and generate a JSON report without elevation.

This baseline is deliberately conservative:

- it has no installer that edits `PATH`, the registry, shell profiles, package
  databases, or system directories;
- one workflow builds every artifact and refuses publication unless all six
  packages and all three native installation checks pass;
- third-party actions in the release workflow are pinned to reviewed full
  commit SHAs rather than movable version tags;
- users can inspect the archive contents before running the executable;
- package-manager metadata cannot drift away from the GitHub release because no
  additional channel is advertised yet.

Checksums detect corruption and substitution relative to the downloaded
checksum list. They do not, by themselves, authenticate who produced that list.
For releases built after provenance support is enabled, GitHub Actions also
creates a keyless artifact attestation for every archive. Users can verify an
archive's repository and workflow origin with `gh attestation verify`; this does
not replace the local checksum check.

## Evaluated installation paths

| Platform | Candidate | Maintenance and security tradeoff | Decision |
| --- | --- | --- | --- |
| Windows | [WinGet Community Repository](https://learn.microsoft.com/en-us/windows/package-manager/package/) | Good discovery and upgrades, but every release needs validated external manifest metadata and a stable publisher/package identity. Community submissions may be reviewed or rejected. | Defer until the release process has a second successfully verified version and manifest updates can be automated and tested. |
| macOS | [Homebrew tap](https://docs.brew.sh/How-to-Create-and-Maintain-a-Tap) | One-command install and upgrade, but an upstream tap is normally a separate `homebrew-*` repository. Formula URLs and SHA-256 values must change with every release, and Intel/Apple-silicon behavior must be tested. It does not remove Gatekeeper requirements for unsigned binaries. | Defer rather than create an unmaintained second repository. |
| Linux | Homebrew on Linux or direct `.deb`/`.rpm` repositories | A tap shares the macOS maintenance cost. Native distro packages improve integration but require multiple formats, repository metadata, signing identities, and upgrade testing across distributions. | Keep the standalone archive until real-user evidence identifies a priority distribution and the project can maintain its repository. |
| All | `curl ... | sh` or equivalent remote installer | Low friction, but executes network-fetched mutable code and can silently change shell or system configuration. It creates a larger trust and recovery surface than this early project can justify. | Not accepted for the current stage. |
| All | Build from a tagged source revision | Reviewable and independent of binary trust, but requires Go and does not serve non-Go users as the primary path. | Supported as a documented fallback, not the default installation. |

Package-manager work should start only after a public issue identifies the
target channel, update automation, test matrix, rollback behavior, and owner.
Popularity alone is not sufficient if the channel cannot be kept current.

## Artifact identity and provenance

The artifact-producing identity is:

- repository: `github.com/raydthanh/aihostcheck`;
- workflow: `.github/workflows/release.yml` in that repository;
- source: the immutable commit targeted by a semantic-version release;
- authorization: a same-repository `release/vMAJOR.MINOR.PATCH` pull request
  merged into `main`, or an existing semantic-version tag used by the documented
  maintainer fallback.

The publish job receives `contents: write`, `id-token: write`, and
`attestations: write` only after packaging and native installation verification
succeed. Other jobs remain read-only.

Provenance uses GitHub's
[`actions/attest`](https://github.com/actions/attest) action. It obtains a
short-lived Sigstore signing certificate from the workflow's GitHub OIDC
identity and records the attestation through GitHub. AIHostCheck therefore has
no persistent provenance private key to copy, export, or store as a repository
secret.

Verify an attested archive online with GitHub CLI:

```sh
gh attestation verify aihostcheck_VERSION_OS_ARCHIVE \
  --repo raydthanh/aihostcheck
```

The verification must identify the expected repository and release workflow.
An attestation proves build origin and digest binding; it does not prove that
the program is bug-free, and it is not a Windows Authenticode signature or an
Apple Developer ID signature.

## Key custody, rotation, and recovery

### Keyless provenance

- **Custody:** GitHub and Sigstore issue short-lived certificates from the
  workflow identity. No maintainer stores a long-lived provenance private key.
- **Authorization changes:** edits to the release workflow are reviewed like
  source code and must pass pull-request CI. Write permissions stay scoped to
  the publish job.
- **Rotation:** GitHub/Sigstore rotate their service keys and trust roots. The
  project reviews the full commit SHA pinned for `actions/attest` before each
  release and upgrades it through a tested pull request.
- **Recovery:** if the repository, workflow, or dependency is suspected to be
  compromised, stop publication, identify affected workflow runs and releases,
  remove or mark affected artifacts, repair the workflow, and publish rebuilt
  artifacts under a new version. Never silently replace an archive under an
  existing version. Delete invalid attestations through GitHub's supported
  lifecycle process and publish an incident note.

### Future operating-system code signing

Windows Authenticode and Apple Developer ID signing are not enabled. Before
either is added, a separate proposal must document:

1. the legal publisher identity shown to users;
2. certificate enrollment and renewal ownership;
3. hardware-backed or managed-service private-key custody with least privilege;
4. a second authorized recovery maintainer or documented account-recovery path;
5. rotation before expiration and immediate revocation after compromise;
6. timestamping, notarization where applicable, CI secret boundaries, and a
   signed test artifact verified on the target OS;
7. incident communication and a rule against replacing old release assets.

Until that proposal and a verified signed release exist, documentation must say
the binaries are not OS code-signed. A checksum or GitHub attestation must never
be described as suppressing SmartScreen or Gatekeeper warnings.
