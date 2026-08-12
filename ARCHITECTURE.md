# Architecture

`cmd/aihostcheck` handles CLI presentation. `internal/collector` owns a common capability catalog plus build-tagged OS collectors. `internal/probe` is the single bounded command-execution boundary. `internal/model` owns the versioned report types, and `internal/report` renders terminal output.

Collectors must be read-only and return evidence, not inference. Prefer standard-library/runtime or documented OS APIs. A command probe must name a binary directly, use hard-coded arguments, impose a timeout and output bound, and never include user-controlled text. A platform collector may use a fixed, read-only PowerShell/CIM query only when Windows has no practical supported inventory executable; the script must be embedded in source, run without profiles or interaction, and contain no interpolated input. New data collection must pass the privacy review in `PRIVACY.md` and have tests. Keep OS-specific code behind build tags so every supported target compiles independently.
