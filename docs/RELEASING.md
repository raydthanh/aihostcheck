# Release process

AIHostCheck uses controlled release pull requests and GitHub Actions to build releases. Release assets must never be assembled manually on a maintainer workstation.

## Version policy

- Use `release/vMAJOR.MINOR.PATCH` branch names, for example `release/v0.1.0`. The workflow creates the matching tag only after that release pull request is merged.
- During `0.x`, minor releases may add or revise pre-1.0 behavior; avoid breaking the JSON report schema without changing `schema_version` as documented in the README.
- Use prerelease suffixes such as `v0.2.0-rc.1` only for builds that should not be presented as stable.

## Maintainer checklist

1. Confirm CI is green on `main`.
2. Create a same-repository branch named `release/vMAJOR.MINOR.PATCH` from that exact `main` commit.
3. On the release branch, move relevant entries from `Unreleased` in `CHANGELOG.md` into a dated version section, then open a pull request.
4. If `schema_version` changed, confirm the normative schema, compatibility tests, examples, and `docs/migrations/` notes changed together.
5. Confirm normal CI and all six dry-run packaging jobs pass on the release pull request.
6. Merge the release pull request. Its closed-and-merged event creates the matching tag and GitHub Release from the resulting `main` commit. Fork pull requests and ordinary branch names cannot publish.
7. Confirm all six release packaging jobs and the native Windows, macOS, and
   Linux installation-verification jobs pass.
8. Confirm the provenance-attestation step succeeds before publication and the
   GitHub release contains six archives plus `checksums.txt`.
9. Download at least one archive, verify its checksum and attestation, run
   `--version`, and generate a JSON report without elevated privileges.
10. Confirm the website and installation documentation link to the new release.

The release workflow validates the version, embeds it in each binary, cross-compiles with `CGO_ENABLED=0`, normalizes archive timestamps and ownership metadata, and publishes only after every operating-system and architecture package succeeds. Existing externally created semantic-version tags remain supported as a maintainer fallback.

The [distribution and signing policy](DISTRIBUTION.md) defines the artifact
identity, scoped permissions, keyless provenance, recovery procedure, and the
requirements that must be met before operating-system code signing is added.
