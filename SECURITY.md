# Security Policy

## Reporting a vulnerability

Please use GitHub's private security-advisory reporting for this repository rather than a public issue. Include affected versions, reproduction steps, and impact. Maintainers will acknowledge a report as soon as practical and coordinate disclosure after a fix.

## Security model

AIHostCheck is read-only and offline. Probes execute resolved binaries directly with fixed arguments, bounded output, and a per-command timeout. Reports are untrusted input when consumed elsewhere: validate them against the bundled JSON Schema and never interpolate values into shell commands. Running the tool with elevated privileges is neither required nor recommended.
