# Privacy

AIHostCheck works locally and offline. It contains no networking, analytics, telemetry, report upload, or persistent identifier. Reports exclude usernames, hostnames, IP addresses, credentials, complete environment-variable sets, personal files, and process command lines. On Unix, only `SHELL` is read to describe the active shell. System files read by the Linux collector are `/etc/os-release` and `/proc/meminfo`.

Some reported hardware and software versions can help fingerprint a machine. Users should inspect JSON before sharing it. A contribution that adds a field must document necessity, source, sensitivity, retention (normally none), failure behavior, and whether a less identifying alternative exists.
