# Stdlib Strategy and Shim Decision Matrix

Terms used in this document:

- **Shim**: glue code that preserves Haxe API/semantics when lowering to Go. It can be compiler-emitted or runtime-backed.
- **`hxrt`**: shared Go runtime helper package copied into generated output (`<go_output>/hxrt`).
- **staged stdlib (`std/_std`)**: target-specific stdlib override modules that replace selected upstream std behavior in a controlled way.
- **portable contract**: portability-first semantic baseline documented in `docs/portable-canonical-contract.md`.

Reference glossary: `docs/glossary.md`.

## Scope

`reflaxe.go` currently combines three mechanisms:

- runtime helpers in `runtime/hxrt/hxrt.go`
- compiler-emitted stdlib shims in `src/reflaxe/go/GoCompiler.hx`
- staged stdlib sources under `std/_std` (wired by `src/reflaxe/go/CompilerBootstrap.hx`)

This document records which compiler-core shims should stay, which should migrate, and why.
For runtime package internals and call-flow wiring, see `docs/hxrt-runtime.md`.
Execution history and validation evidence are tracked in `docs/stdlib-shim-migration-log.md`.

This document covers current shim ownership decisions. The long-range parity program (portable full stdlib coverage + portable/native facade boundary governance) is tracked in `docs/portable-stdlib-parity-program.md`.

## Why This Approach (and Whether It Is Go-Specific)

Short answer:

- The hybrid model is a general compiler-target pattern, not Go-only.
- Go amplifies the need for it because several Haxe semantics do not map directly to native Go behavior.

### Target-agnostic reasons to start hybrid

1. Time-to-quality
   - A full stdlib rewrite in target language before shipping parity is usually too large and too risky.
2. Verification-first iteration
   - Hybrid ownership lets teams land small slices with snapshot/semantic gates and migrate safely.
3. Ownership flexibility
   - Runtime helpers, compiler lowering, and staged stdlib code can move independently as evidence improves.

### Why Go makes this especially pragmatic

1. Exception model mismatch
   - Haxe throw/catch semantics require controlled panic/recover bridging.
2. String/nullability/dynamic representation differences
   - Pointer-string conventions and `Std.string`-compatible behavior need centralized helpers.
3. Reflection and resolver edge cases
   - Serializer/unserializer contracts require typed metadata-aware lowering plus runtime behavior.
4. Profile-dependent output policy
   - `portable|metal` profile guarantees are compile-time policy decisions, not runtime-only toggles.

### Design rule in this repo

Use the simplest ownership that preserves parity and maintainability:

- runtime (`hxrt`) for reusable target-runtime behavior,
- compiler shims for compile-time metadata/profile-sensitive contracts,
- staged `std/_std` as migration destination once parity is proven.

Family precedent matters here:

- `haxe.rust` keeps library-style surfaces like `StringTools` in staged std/runtime layers (`std/StringTools.cross.hx`, `std/hxrt/string/NativeString.hx`) instead of compiler-emitted decl blobs.
- `haxe.elixir` keeps library-style surfaces like `StringTools` and `DateTools` in target-gated std overrides (`std/*.cross.hx`, `std/_std/**`) instead of compiler-core shims.
- `haxe.go` should follow the same default. Compiler ownership needs a concrete reason such as compile-time metadata coupling, profile-sensitive lowering, or a representation boundary that staged std/runtime cannot express cleanly.

Portable and native surfaces are distinct:

- portable contract targets full Haxe stdlib parity;
- `go.*` is an explicit native facade and is not portability-safe.

## Alternatives Reviewed

| Alternative | Strength | Current blocker |
| --- | --- | --- |
| Compiler-lowered builtins/intrinsics only | Minimal generated wrappers for hot paths | Only viable when behavior is pure and profile-invariant; many stdlib surfaces still require serializer metadata, exception wiring, or dynamic conversion policy. |
| Externs + external Go runtime package | Clean boundary and reuse potential | Externs are type-only and `ignoreExterns: true` is currently required for deterministic emission in `src/reflaxe/go/CompilerInit.hx`. |
| Raw `__go__` in Haxe std/app code | Minimal indirection for target-native calls | Still the wrong default for app/examples, but now valid in framework-owned low-level abstraction islands via `@:goAllowRaw` + `reflaxe.go.macros.GoInjection.__go__`. Keep imports typed with extern metadata; do not use raw injection as a substitute for `@:go.import`/`@:go.name`. |
| Vendored stdlib-only (`std/_std`) | Most idiomatic long-term ownership model | Behavior-heavy contracts still depend on compiler context (serializer metadata, socket readiness/deadline behavior, profile-aware lowering). |

## Decision Matrix

`Compiler LOC` values below are from shim function spans in `src/reflaxe/go/GoCompiler.hx` (measured on 2026-02-19).

| Shim group | Primary surfaces | Compiler LOC | Highest CI tier | Decision | Reason | Follow-up |
| --- | --- | ---: | --- | --- | --- | --- |
| `json` | `haxe.Json`, `haxe.format.JsonParser/JsonPrinter` | 38 | Snapshot | Migrated (staged std/_std + runtime-owned behavior) | Staged std overrides (`std/_std/**`) now own JSON API surfaces; behavior delegates to `hxrt.JsonParse`/`hxrt.JsonStringify`. | `haxe.go-7zy.10`, `haxe.go-cgk.5` |
| `sys` | `Sys`, `sys.io.File`, `sys.io.Process` | 89 | Snapshot | Migrated (runtime-owned wrappers) | Behavior now lives in `hxrt.Sys*`/`hxrt.File*`/`hxrt.Process*`; compiler shim generation is reduced to thin wrapper/type-shape forwarding. | `haxe.go-7zy.11` (completed 2026-02-19) |
| `io` | `haxe.io.Bytes`, buffers, input/output base wiring | 108 | Snapshot + semantic-diff dependency | Keep (for now, with selective helper emission) | Shared representation boundary used by crypto/http/serializer flows; inherited Input/Output helper declarations are now emitted only when helper usage is detected. | `haxe.go-czm` (in progress) |
| `ds` | `haxe.ds.*Map`, `List`, enum maps | 149 | Snapshot + semantic-diff dependency | Keep (for now) | Serializer and HTTP contracts rely on deterministic generated map/list shapes. | - |
| `http` | `sys.Http` request/callback/proxy contract | 542 | Semantic-diff | Keep | Behavior includes callback choreography and deterministic request handling under test contract. | - |
| `stdlib_symbols` | `Std`, `Date`, `Math`, `Reflect`, crypto/xml/zip, filesystem subset | 706 | Semantic-diff | Split: keep compiler-context-sensitive core, migrate library surfaces | Keep compiler ownership for `Std`, `Reflect`, `Type`, `Xml`, and bytes/representation-sensitive crypto-zip paths where lowering context still matters. `StringTools`, `DateTools`, `haxe.io.Path`, and direct `haxe.SysTools` helper ownership now live in staged std. | `haxe.go-7zy.12`, `haxe.go-14as.33`, `haxe.go-14as.34`, `haxe.go-14as.35`, `haxe.go-14as.36`, `haxe.go-14as.25` |
| `regex_serializer` | `EReg`, `haxe.Serializer`, `haxe.Unserializer` | 2460 | Semantic-diff | Keep | High behavior density and project metadata coupling (resolver semantics, token stream, reflection). | - |
| `net_socket` | `sys.net.Host`, `sys.net.Socket` | 2958 | Semantic-diff | Keep | Deadline/select/shutdown readiness behavior is target-specific and currently best enforced in one compiler-controlled path. | - |

## Explicit Decision Records

These are the canonical per-surface decisions for shim ownership and alternatives.

| Record | Surface | Decision | Alternatives reviewed | Evidence |
| --- | --- | --- | --- | --- |
| `SDR-001` | `json` (`haxe.Json`, `haxe.format.Json*`) | Move API ownership to staged std (`std/_std`) and keep behavior in `hxrt` (`JsonParse`/`JsonStringify`) | Compiler shim, direct lower-call rewrites, extern/runtime package | Snapshot parity + migration log (`haxe.go-7zy.10`, `haxe.go-cgk.5`) |
| `SDR-002` | `sys` (`Sys`, `sys.io.File`, `sys.io.Process`) | Keep thin compiler wrappers, move behavior to `hxrt` (`Sys*`, `File*`, `Process*`) | Compiler shim, direct externs, `std/_std` | Snapshot parity + migration log (`haxe.go-7zy.11`) |
| `SDR-003` | `io` (`haxe.io.Bytes*`, stream helpers, encoding edges) | Keep compiler shims for now; allow selective helper emission + targeted runtime helpers | Compiler builtin lowering, extern/runtime package, `std/_std` | Semantic-diff contracts + shim-vs-direct benchmark (`test:perf:stdlib-shims`) |
| `SDR-004` | `ds` (`haxe.ds.*Map`, `List`) | Keep compiler-owned shape generation until typed null/reflect parity is proven in a replacement path | `std/_std`, extern-backed containers, pure runtime wrappers | Semantic-diff contracts (`ds_maps_list_contract`, map/list follow-ups) |
| `SDR-005` | `http` (`sys.Http`) | Keep compiler shims (behavioral choreography, callbacks, proxy and payload conversion) | Extern-only wrappers, raw-native app code, `std/_std` | Semantic-diff contracts (`http_request_callbacks_contract`, proxy/custom request tests) |
| `SDR-006` | `regex_serializer` (`EReg`, serializer/unserializer stack) | Keep compiler shims; revisit only with equivalent metadata-aware runtime path | Runtime-only package, extern-only, `std/_std` | Serializer and regex semantic-diff suite |
| `SDR-007` | `net_socket` (`sys.net.Host`, `sys.net.Socket`) | Keep compiler shims for deadline/shutdown/readiness semantics | Extern-only wrappers, `std/_std` | Socket semantic-diff contracts (loopback + advanced) |
| `SDR-008` | library-expressible `stdlib_symbols` surfaces (`StringTools`, `DateTools` helpers, `haxe.io.Path`, direct `haxe.SysTools`) | Default to staged std/runtime ownership; do not grow compiler-resident helper logic unless staged std/runtime is proven insufficient. `StringTools`, `DateTools`, `haxe.io.Path`, and `haxe.SysTools` are the completed migrations in this bucket so far. | Keep in compiler core, extern-only wrappers | Cross-target precedent (`haxe.rust` `std/StringTools.cross.hx` + runtime helper; `haxe.elixir` `std/StringTools.cross.hx`, `std/DateTools.cross.hx`) plus local migration beads (`haxe.go-14as.33`-`haxe.go-14as.36`, `haxe.go-14as.25`) |
| `SDR-009` | `haxe.Constraints` + `haxe.Rest` direct abstraction surfaces | Split ownership: staged std for `haxe.Constraints` typing/native bridge metadata, compiler lowering for native-slice `haxe.Rest` operations | Keep as compile-only debt, full compiler shim blobs, staged std abstract emission | Direct parity contracts (`haxe_constraints_contract`, `haxe_rest_contract`) plus snapshot `stdlib/haxe_constraints_rest_direct` |
| `SDR-010` | stack/main-loop `haxe.misc` surfaces (`haxe.CallStack`, `haxe.NativeStackTrace`, `haxe.EntryPoint`, `haxe.MainLoop`, `haxe.Timer`) | Use staged std for deterministic stack fallbacks and source-owned std inclusion for event-loop classes; classify the tranche as target-sensitive snapshot coverage, not portable semantic-diff parity | Leave as generic compile-only debt, force semantic-diff on event-loop scheduling, grow compiler-owned stack shims | Snapshot contract `stdlib/haxe_stack_loop_target_sensitive`, staged std overrides `std/haxe/CallStack.cross.hx` + `std/haxe/NativeStackTrace.cross.hx`, cross-target precedent from `haxe.rust/std/haxe/CallStack.cross.hx` |
| `SDR-011` | legacy text surfaces (`haxe.Utf8`, `haxe.Ucs2`) | Use staged std for deprecated `haxe.Utf8` helper semantics; keep `haxe.Ucs2` as explicit target-sensitive platform exclusion under snapshot coverage | Grow compiler-owned text shims, leave both as anonymous compile-only debt | `haxe_utf8_contract`, `stdlib/haxe_utf8_basic`, `stdlib/haxe_ucs2_platform_exclusion`, staged std override `std/haxe/Utf8.cross.hx` |

Review trigger for all records: revisit when an alternative path proves equal/better parity and performance under the same harness gates.

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

Measured at `2026-02-24T01:22:42Z` on `darwin/arm64` (`Apple M2 Pro`):

| Path | ns/op | B/op | allocs/op | Code-shape LOC (call path) |
| --- | ---: | ---: | ---: | ---: |
| Generated shim (`haxe__crypto__Base64_encode` + bytes conversion helpers) | 64.05 | 112 | 3 | 33 |
| Direct Go (`base64.StdEncoding.EncodeToString`) | 57.09 | 96 | 2 | 3 |
| Delta | +12.19% | +16 | +1 | +30 |

Interpretation:

- overhead is primarily representation conversion (`[]int` <-> `[]byte`) rather than base64 algorithm cost
- this supports keeping the compatibility shim while targeting focused conversion-path optimization (`haxe.go-7zy.12`)

## Migration Sequence

1. Move `json` out of compiler core first (`haxe.go-7zy.10`) and then promote staged std ownership (`haxe.go-cgk.5`) because it is the thinnest shim and lowest risk.
2. Move `sys` wrappers second (`haxe.go-7zy.11`, completed 2026-02-19) once snapshot parity remains stable.
3. Keep behavior-heavy shim groups in compiler core until an equivalent `std/_std` path proves equal parity under semantic-diff coverage.

## Revisit Triggers

Re-open keep decisions when one of these becomes true:

1. `std/_std` path reaches equal or better parity for the same fixtures.
2. Runtime package extraction can preserve profile/lowering policy without semantic drift.
3. A compiler shim becomes pure forwarding with no compiler-context decisions.
