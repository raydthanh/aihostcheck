# Release process

AIHostCheck uses semantic version tags and GitHub Actions to build releases. Release assets must never be assembled manually on a maintainer workstation.

## Version policy

- Use `vMAJOR.MINOR.PATCH` tags, for example `v0.1.0`.
- During `0.x`, minor releases may add or revise pre-1.0 behavior; avoid breaking the JSON report schema without changing `schema_version` as documented in the README.
- Use prerelease suffixes such as `v0.2.0-rc.1` only for builds that should not be presented as stable.

## Maintainer checklist

1. Confirm CI is green on `main`.
2. Move relevant entries from `Unreleased` in `CHANGELOG.md` into a dated version section.
3. Merge the changelog change and confirm CI again.
4. Create and push an annotated semantic-version tag from that exact `main` commit.
5. Confirm all six packaging jobs pass and the GitHub release contains six archives plus `checksums.txt`.
6. Download at least one archive, verify its checksum, run `--version`, and generate a JSON report without elevated privileges.
7. Confirm the website and installation documentation link to the new release.

The release workflow validates tag format, embeds the tag in each binary, cross-compiles with `CGO_ENABLED=0`, and publishes only after every operating-system and architecture package succeeds.
