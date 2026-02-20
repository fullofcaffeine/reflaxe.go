# Stdlib Strategy and Shim Decision Matrix

## Scope

`reflaxe.go` currently combines three mechanisms:

- runtime helpers in `runtime/hxrt/hxrt.go`
- compiler-emitted stdlib shims in `src/reflaxe/go/GoCompiler.hx`
- staged stdlib sources under `std/_std` (wired by `src/reflaxe/go/CompilerBootstrap.hx`)

This document records which compiler-core shims should stay, which should migrate, and why.
For runtime package internals and call-flow wiring, see `docs/hxrt-runtime.md`.
Execution history and validation evidence are tracked in `docs/stdlib-shim-migration-log.md`.

## Alternatives Reviewed

| Alternative | Strength | Current blocker |
| --- | --- | --- |
| Externs + external Go runtime package | Clean boundary and reuse potential | Externs are type-only and `ignoreExterns: true` is currently required for deterministic emission in `src/reflaxe/go/CompilerInit.hx`. |
| Raw `__go__` in Haxe std/app code | Minimal indirection for target-native calls | Violates strict policy in app/examples (`src/reflaxe/go/macros/StrictModeEnforcer.hx`, `src/reflaxe/go/macros/BoundaryEnforcer.hx`) and harms portability/readability. |
| Vendored stdlib-only (`std/_std`) | Most idiomatic long-term ownership model | Behavior-heavy contracts still depend on compiler context (serializer metadata, socket readiness/deadline behavior, profile-aware lowering). |

## Decision Matrix

`Compiler LOC` values below are from shim function spans in `src/reflaxe/go/GoCompiler.hx` (measured on 2026-02-19).

| Shim group | Primary surfaces | Compiler LOC | Highest CI tier | Decision | Reason | Follow-up |
| --- | --- | ---: | --- | --- | --- | --- |
| `json` | `haxe.Json`, `haxe.format.JsonParser/JsonPrinter` | 38 | Snapshot | Migrated (runtime-lowered) | Compiler-emitted JSON declarations removed; calls now lower directly to `hxrt.JsonParse`/`hxrt.JsonStringify`. | `haxe.go-7zy.10` |
| `sys` | `Sys`, `sys.io.File`, `sys.io.Process` | 89 | Snapshot | Migrated (runtime-owned wrappers) | Behavior now lives in `hxrt.Sys*`/`hxrt.File*`/`hxrt.Process*`; compiler shim generation is reduced to thin wrapper/type-shape forwarding. | `haxe.go-7zy.11` (completed 2026-02-19) |
| `io` | `haxe.io.Bytes`, buffers, input/output base wiring | 108 | Snapshot + semantic-diff dependency | Keep (for now, with selective helper emission) | Shared representation boundary used by crypto/http/serializer flows; inherited Input/Output helper declarations are now emitted only when helper usage is detected. | `haxe.go-czm` (in progress) |
| `ds` | `haxe.ds.*Map`, `List`, enum maps | 149 | Snapshot + semantic-diff dependency | Keep (for now) | Serializer and HTTP contracts rely on deterministic generated map/list shapes. | - |
| `http` | `sys.Http` request/callback/proxy contract | 542 | Semantic-diff | Keep | Behavior includes callback choreography and deterministic request handling under test contract. | - |
| `stdlib_symbols` | `Std`, `StringTools`, `Date`, `Math`, `Reflect`, crypto/xml/zip, filesystem subset | 706 | Semantic-diff | Keep + optimize (landed) | Broad compatibility layer remains in compiler core; bytes conversion path now uses cached raw representation to cut repeated conversion overhead. | `haxe.go-7zy.12` |
| `regex_serializer` | `EReg`, `haxe.Serializer`, `haxe.Unserializer` | 2460 | Semantic-diff | Keep | High behavior density and project metadata coupling (resolver semantics, token stream, reflection). | - |
| `net_socket` | `sys.net.Host`, `sys.net.Socket` | 2958 | Semantic-diff | Keep | Deadline/select/shutdown readiness behavior is target-specific and currently best enforced in one compiler-controlled path. | - |

## Ownership Boundary (Post `haxe.go-7zy.11`)

- `runtime/hxrt/hxrt.go` owns `Sys`/`sys.io.File`/`sys.io.Process` behavior (OS args/cwd, file reads/writes, process launch/stdout/close).
- `src/reflaxe/go/GoCompiler.hx` owns lowering and generated type-shape wrappers only for this surface.
- `lowerSysStdlibShimDecls` must remain forwarding-only unless a behavior change is intentionally re-centralized and justified with parity/perf evidence.

## Measured Tradeoff: Shim vs Simpler Path

Representative surface: `haxe.crypto.Base64.encode` in `stdlib_symbols`.

Repro command:

```bash
npm run test:perf:stdlib-shims
```

Artifacts:

- `.cache/perf-stdlib-shim-review/report.json`
- `.cache/perf-stdlib-shim-review/report.md`

Measured at `2026-02-19T22:50:44Z` on `darwin/arm64` (`Apple M2 Pro`):

| Path | ns/op | B/op | allocs/op | Code-shape LOC (call path) |
| --- | ---: | ---: | ---: | ---: |
| Generated shim (`haxe__crypto__Base64_encode` + bytes conversion helpers) | 71.52 | 112 | 3 | 28 |
| Direct Go (`base64.StdEncoding.EncodeToString`) | 46.55 | 96 | 2 | 3 |
| Delta | +53.64% | +16 | +1 | +25 |

Interpretation:

- overhead is primarily representation conversion (`[]int` <-> `[]byte`) rather than base64 algorithm cost
- this supports keeping the compatibility shim while targeting focused conversion-path optimization (`haxe.go-7zy.12`)

## Migration Sequence

1. Move `json` out of compiler core first (`haxe.go-7zy.10`) because it is the thinnest shim and lowest risk.
2. Move `sys` wrappers second (`haxe.go-7zy.11`, completed 2026-02-19) once snapshot parity remains stable.
3. Keep behavior-heavy shim groups in compiler core until an equivalent `std/_std` path proves equal parity under semantic-diff coverage.

## Revisit Triggers

Re-open keep decisions when one of these becomes true:

1. `std/_std` path reaches equal or better parity for the same fixtures.
2. Runtime package extraction can preserve profile/lowering policy without semantic drift.
3. A compiler shim becomes pure forwarding with no compiler-context decisions.
