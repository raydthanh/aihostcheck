# Report provenance and compatibility

AIHostCheck versions its executable and its machine-readable report contract
independently. Consumers must inspect both instead of inferring schema behavior
from the CLI release number.

| Field | Meaning | Consumer use |
| --- | --- | --- |
| `schema_version` | Semantic version of the JSON contract. | Decide whether the document can be interpreted safely. |
| `tool_version` | AIHostCheck release that produced the report. | Identify collector behavior and known release limitations. |
| `generated_at` | UTC time at which collection began. | Decide whether host evidence may be stale. |
| `platform` | OS and architecture of the running binary. | Detect an unexpected or wrong-platform report. |

None of these fields proves that a capability is available. Capability status
and evidence remain authoritative for environment decisions.

## Compatibility rule

Assume a consumer implements schema `1.2.0`:

| Report schema | Classification | Required behavior |
| --- | --- | --- |
| `1.2.0` or another `1.2.x` | Exact schema line | Read known fields and preserve all status meanings. Patch changes do not change contract semantics. |
| `1.1.x` | Compatible older minor | Read fields that exist; do not treat fields introduced in `1.2` as absent capabilities. |
| `1.3.x` | Newer minor | Continue only if the reader deliberately tolerates unknown additive fields and capabilities. Otherwise upgrade the reader. |
| `2.x.x` | Incompatible major | Do not interpret capability semantics. Upgrade the reader or generate a report with a supported contract. |

Versions must use stable `MAJOR.MINOR.PATCH` notation. A missing, malformed, or
prerelease schema version is invalid. The reference implementation is
`model.AssessSchemaCompatibility` and is covered by table-driven tests.

### Representative report envelopes

Equal schema line:

```json
{"schema_version":"1.2.0","generated_at":"2026-08-13T00:00:00Z","tool_version":"0.2.0","platform":"linux/amd64","capabilities":{}}
```

Older minor schema:

```json
{"schema_version":"1.1.4","generated_at":"2026-08-12T00:00:00Z","tool_version":"0.1.4","platform":"windows/amd64","capabilities":{}}
```

Newer minor schema requiring a tolerant reader:

```json
{"schema_version":"1.3.0","generated_at":"2026-08-13T00:00:00Z","tool_version":"0.3.0","platform":"darwin/arm64","capabilities":{"new_capability":{"status":"unknown","evidence":[{"source":"example","detail":"additive capability unknown to the 1.2 consumer"}]}}}
```

Different major schema:

```json
{"schema_version":"2.0.0","generated_at":"2026-08-13T00:00:00Z","tool_version":"1.0.0","platform":"linux/arm64","capabilities":{}}
```

These short envelopes demonstrate version decisions; a complete report must
also satisfy the normative JSON Schema bundled with its release.

## When to request a fresh report

There is no universal age after which a host report becomes false. Request a
fresh report when:

- the OS, architecture, shell, PATH, runtime, package manager, container tool,
  driver, toolkit, or hardware relevant to the task may have changed;
- `generated_at` is missing, invalid, unexpectedly in the future, or predates a
  known host change;
- `tool_version` identifies a release with a relevant fixed limitation;
- a required capability is `unknown`, `unsupported`, `permission_denied`, or
  `error` and a newer collector can obtain safer evidence;
- the report's schema major is incompatible with the consumer.

Generating a fresh report does not justify elevation. Run without administrator
or root privileges unless a documented diagnostic step explicitly requires it.

## Schema-change and migration policy

- A patch version may clarify descriptions or tighten validation without
  changing field meaning.
- A minor version may add optional fields, capabilities, or status-preserving
  evidence. Consumers must not reinterpret missing additions as
  `not_detected`.
- A major version is required for removed fields, renamed fields, changed status
  meaning, or other incompatible semantics.
- Every semantic schema change must add migration notes under
  `docs/migrations/`, update the normative JSON Schema and examples, and add
  tests before release.

Migration notes must state who is affected, how old and new documents differ,
whether collection must be repeated, and how all six capability statuses retain
their meaning.
