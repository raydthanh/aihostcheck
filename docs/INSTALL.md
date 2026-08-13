# Install AIHostCheck

AIHostCheck is distributed as a standalone executable. It does not require Python, Node.js, Docker, or administrator privileges.

## 1. Choose a download

Open the [latest GitHub release](https://github.com/raydthanh/aihostcheck/releases/latest) and choose the archive that matches the computer:

| Computer | Download suffix |
| --- | --- |
| Windows on most Intel or AMD computers | `windows_amd64.zip` |
| Windows on ARM | `windows_arm64.zip` |
| macOS on Apple silicon (M1 or newer) | `darwin_arm64.tar.gz` |
| macOS on Intel | `darwin_amd64.tar.gz` |
| Linux on most Intel or AMD computers | `linux_amd64.tar.gz` |
| Linux on ARM64 | `linux_arm64.tar.gz` |

Each archive also contains the license, report schema, and a copy of this installation guide.

## 2. Verify the download

Download `checksums.txt` from the same release. Verify the archive's SHA-256
value before running it. The commands below fail instead of asking you to
compare long values visually.

### Windows PowerShell

```powershell
$Archive = "aihostcheck_0.2.1_windows_amd64.zip"
$Match = @(Select-String -Path .\checksums.txt -Pattern ("  " + [regex]::Escape($Archive) + "$"))
if ($Match.Count -ne 1) { throw "Expected one checksum entry" }
$Expected = ($Match[0].Line -split '\s+')[0]
$Actual = (Get-FileHash ".\$Archive" -Algorithm SHA256).Hash.ToLowerInvariant()
if ($Actual -ne $Expected.ToLowerInvariant()) { throw "Checksum mismatch" }
"Checksum verified: $Archive"
```

### macOS

```sh
archive=aihostcheck_0.2.1_darwin_arm64.tar.gz
checksum_line=$(awk -v name="$archive" '$2 == name { print }' checksums.txt)
test -n "$checksum_line"
printf '%s\n' "$checksum_line" | shasum -a 256 -c -
```

### Linux

```sh
archive=aihostcheck_0.2.1_linux_amd64.tar.gz
checksum_line=$(awk -v name="$archive" '$2 == name { print }' checksums.txt)
test -n "$checksum_line"
printf '%s\n' "$checksum_line" | sha256sum -c -
```

The command must print `OK` (macOS/Linux) or `Checksum verified` (Windows).
An empty or missing checksum line is a failure; do not run the archive. Replace
`0.2.1` and the platform suffix with the downloaded filename.

### Verify build provenance when available

AIHostCheck v0.1.0 predates artifact attestations. Releases published after
provenance support is enabled can also be authenticated with GitHub CLI:

```sh
gh attestation verify aihostcheck_VERSION_OS_ARCHIVE \
  --repo raydthanh/aihostcheck
```

This check must identify `raydthanh/aihostcheck` as the source. Provenance is an
additional origin check; keep the SHA-256 verification step. See the
[distribution and signing policy](DISTRIBUTION.md) for its security boundary.

## 3. Extract and run

### Windows PowerShell

```powershell
Expand-Archive .\aihostcheck_0.2.1_windows_amd64.zip
cd .\aihostcheck_0.2.1_windows_amd64
.\aihostcheck.exe --version
.\aihostcheck.exe --json > aihostcheck-report.json
```

### macOS

```sh
tar -xzf aihostcheck_0.2.1_darwin_arm64.tar.gz
cd aihostcheck_0.2.1_darwin_arm64
./aihostcheck --version
./aihostcheck --json > aihostcheck-report.json
```

### Linux

```sh
tar -xzf aihostcheck_0.2.1_linux_amd64.tar.gz
cd aihostcheck_0.2.1_linux_amd64
./aihostcheck --version
./aihostcheck --json > aihostcheck-report.json
```

Inspect `aihostcheck-report.json` before sharing it with an AI agent or another person. AIHostCheck does not upload the report.

Continue with [Use a report with GPT or an AI coding agent](USING_WITH_AI.md) for a safe prompt template and guidance on interpreting `unknown`, `not_detected`, and evidence fields.

## Platform security notices

Early AIHostCheck binaries are not code-signed. Windows SmartScreen or macOS Gatekeeper may therefore warn or block them even when their checksum is correct. Do not disable system-wide security protections. If local policy does not allow an unsigned binary, build from source instead or wait for signed releases.

## Build from source

Go 1.22 or newer is required:

```sh
git clone https://github.com/raydthanh/aihostcheck.git
cd aihostcheck
go build -o aihostcheck ./cmd/aihostcheck
./aihostcheck --version
```

On Windows, use `go build -o aihostcheck.exe ./cmd/aihostcheck` and run `.\aihostcheck.exe --version`.

## Troubleshooting

- `unknown` does not mean a capability is absent. It means AIHostCheck could not establish the fact safely with the available evidence.
- `not_detected` means a supported check completed and did not find the capability.
- Run without administrator or root privileges; elevation is neither required nor recommended.
- Report installation or detection problems through [GitHub Issues](https://github.com/raydthanh/aihostcheck/issues). Do not include secrets or an unreviewed diagnostic report.
