# Stdlib Strategy and Shim Decision Matrix

Terms used in this document:

- **Shim**: glue code that preserves Haxe API/semantics when lowering to Go. It can be compiler-emitted or runtime-backed.
- **`hxrt`**: shared Go runtime helper package copied into generated output (`<go_output>/hxrt`).
- **staged stdlib (`std/go/_std`)**: ordinary target-specific source modules that replace selected upstream std behavior; package staging alone turns them into `.cross.hx` artifacts.
- **portable contract**: portability-first semantic baseline documented in `docs/portable-canonical-contract.md`.

Reference glossary: `docs/glossary.md`.

Canonical ownership rule: `docs/ownership-rubric.md`.

## Scope

`reflaxe.go` currently combines three mechanisms:

- runtime helpers in `runtime/hxrt/hxrt.go`
- compiler-emitted stdlib shims in `src/reflaxe/go/GoCompiler.hx`
- staged stdlib sources under `std/go/_std` (selected before typing by `haxe_libraries/reflaxe.go.hxml`)

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
4. Build-policy-dependent output decisions
   - Authority, specialization, fallback, strictness, and runtime choices are
     typed compile-time policies; source semantics come from portable or
     explicit native APIs/boundaries.

### Design rule in this repo

Use the simplest ownership that preserves parity and maintainability:

- runtime (`hxrt`) for reusable target-runtime behavior,
- compiler shims for compile-time metadata or policy-sensitive contracts,
- staged `std/go/_std` as the source migration destination once parity is proven.

Family precedent matters here:

- `haxe.rust` keeps library-style surfaces like `StringTools` in staged std/runtime layers (`std/StringTools.cross.hx`, `std/hxrt/string/NativeString.hx`) instead of compiler-emitted decl blobs.
- `haxe.elixir` keeps library-style surfaces like `StringTools` and `DateTools` in target-gated std overrides (`std/*.cross.hx`, `std/_std/**`) instead of compiler-core shims.
- `haxe.go` should follow the same default. Compiler ownership needs a concrete
  reason such as compile-time metadata coupling, policy-sensitive lowering, or
  a representation boundary that staged std/runtime cannot express cleanly.

Portable and native surfaces are distinct:

- portable contract targets full Haxe stdlib parity;
- `go.*` is an explicit native facade and is not portability-safe.

This document is the decision matrix for current shim groups.
The repo-wide rule for deciding ownership lives in `docs/ownership-rubric.md`.

## Alternatives Reviewed

| Alternative | Strength | Current blocker |
| --- | --- | --- |
| Compiler-lowered builtins/intrinsics only | Minimal generated wrappers for hot paths | Only viable when behavior is pure and profile-invariant; many stdlib surfaces still require serializer metadata, exception wiring, or dynamic conversion policy. |
| Externs + external Go runtime package | Clean boundary and reuse potential | Externs are type-only and `ignoreExterns: true` is currently required for deterministic emission in `src/reflaxe/go/CompilerInit.hx`. |
| Raw `__go__` in Haxe std/app code | Minimal indirection for target-native calls | Still the wrong default for app/examples, but now valid in framework-owned low-level abstraction islands via `@:goAllowRaw` + `reflaxe.go.macros.GoInjection.__go__`. Keep imports typed with extern metadata; do not use raw injection as a substitute for `@:go.import`/`@:go.name`. |
| Canonical target `_std` (`std/go/_std`) | Most idiomatic long-term ownership model | Behavior-heavy contracts still depend on compiler context (serializer metadata, socket readiness/deadline behavior, profile-aware lowering). |

## Decision Matrix

`Compiler LOC` values below are from shim function spans in `src/reflaxe/go/GoCompiler.hx` (measured on 2026-02-19).

| Shim group | Primary surfaces | Compiler LOC | Highest CI tier | Decision | Reason | Follow-up |
| --- | --- | ---: | --- | --- | --- | --- |
| `json` | `haxe.Json`, `haxe.format.JsonParser/JsonPrinter` | 38 | Snapshot | Migrated (canonical staged std + runtime-owned behavior) | Staged std source under `std/go/_std/**` owns JSON API surfaces; behavior delegates to `hxrt.JsonParse`/`hxrt.JsonStringify`. | `haxe.go-7zy.10`, `haxe.go-cgk.5` |
| `sys` | `Sys`, `sys.io.File`, `sys.io.Process` | 89 | Snapshot | Migrated (runtime-owned behavior) | Behavior now lives in `hxrt.Sys*`/`hxrt.File*`/`hxrt.Process*`; compiler shim generation retains type shape, Haxe error/EOF translation, and shared IO-helper adapters. `Sys.sleep` follows the same split: a typed generated adapter delegates seconds-to-duration conversion and blocking to `hxrt.SysSleep`. Direct `sys.io.FileInput` / `FileOutput` / `FileSeek` parity also rides on this layer through runtime-backed open/read/write/seek/tell/eof helpers. | `haxe.go-7zy.11` (completed 2026-02-19), `haxe.go-14as.17`, `haxe_go-vfp.8.7.1` |
| `io` | `haxe.io.Bytes`, buffers, input/output base wiring | 108 | Snapshot + semantic-diff dependency | Split: keep representation-sensitive core, migrate algorithmic helpers when parity is proven | Shared representation boundary still matters for bytes shape and inherited IO types, but the generic helper loops no longer need to stay in `GoCompiler`. `Bytes.ofHex`, `Bytes.toHex`, and `BytesBuffer` leaf operations now live in `hxrt`; inherited `Input` / `Output` loop helpers now live in staged source (`std/haxe/io/GoIoHelpers.hx`); RawNative/cache-coupled string helpers and numeric IO primitives remain compiler-owned because they co-own raw-native mode dispatch plus `__hx_raw` cache validity used by downstream raw-byte consumers. | `haxe.go-czm`, `haxe.go-14as.50`, `haxe.go-14as.51`, `haxe.go-14as.52`, `haxe.go-14as.54` |
| `ds` | `haxe.ds.*Map`, `List`, enum maps | 149 | Snapshot + semantic-diff dependency | Keep (for now) | Serializer and HTTP contracts rely on deterministic generated map/list shapes. | - |
| `http` | `sys.Http` request/callback/proxy contract | 542 | Semantic-diff | Keep core choreography; extract source-owned leaf helpers cautiously | Request sequencing, callback timing, and response/body normalization are still one semantic contract. `getResponseHeaderValues` and payload capture now live in a staged helper (`std/sys/GoHttpHelpers.hx`) via framework-owned `__go__`, while request lifecycle and proxy URL construction remain compiler-owned. | `haxe.go-14as.50`, `haxe.go-14as.53` |
| `stdlib_symbols` | `Std`, `Date`, `Math`, `Reflect`, crypto/xml/zip, filesystem subset | 706 | Semantic-diff | Split: keep compiler-context-sensitive core, migrate library surfaces | Keep compiler ownership for `Std`, `Reflect`, `Type`, `Xml`, and bytes/representation-sensitive crypto-zip paths where lowering context still matters. `StringTools`, `DateTools`, and direct `haxe.SysTools` helper ownership live in staged std; `haxe.io.Path` now uses the upstream Haxe stdlib implementation after the required string/array lowerings landed. | `haxe.go-7zy.12`, `haxe.go-14as.33`, `haxe.go-14as.34`, `haxe.go-14as.35`, `haxe.go-14as.36`, `haxe.go-14as.25` |
| `regex_serializer` | `EReg`, `haxe.Serializer`, `haxe.Unserializer` | 2460 | Semantic-diff | Keep | High behavior density and project metadata coupling (resolver semantics, token stream, reflection). | - |
| `net_socket` | `sys.net.Host`, `sys.net.Socket`, `sys.net.UdpSocket` | 2958 | Mixed (`semantic-diff` for TCP host/socket, snapshot for direct UDP loopback) | Keep | Deadline/select/shutdown/readiness behavior, UDP address translation, and UDP broadcast socket-option handling are target-specific and currently best enforced in one compiler-controlled path. | Keep broadcast delivery tests scoped to deterministic socket-option evidence; do not require LAN broadcast packet delivery in CI |

## Explicit Decision Records

These are the canonical per-surface decisions for shim ownership and alternatives.

| Record | Surface | Decision | Alternatives reviewed | Evidence |
| --- | --- | --- | --- | --- |
| `SDR-001` | `json` (`haxe.Json`, `haxe.format.Json*`) | Move API ownership to canonical staged std (`std/go/_std`) and keep behavior in `hxrt` (`JsonParse`/`JsonStringify`) | Compiler shim, direct lower-call rewrites, extern/runtime package | Snapshot parity + migration log (`haxe.go-7zy.10`, `haxe.go-cgk.5`) |
| `SDR-002` | `sys` (`Sys`, `sys.io.File`, `sys.io.Process`, direct `sys.io.FileInput` / `FileOutput` / `FileSeek`) | Keep thin compiler wrappers, move behavior to `hxrt` (`Sys*`, `File*`, `Process*`) | Compiler shim, direct externs, canonical `std/go/_std` | Snapshot parity + migration log (`haxe.go-7zy.11`) + direct `sys_db_io_contract` / `stdlib/sys_db_io_direct` (`haxe.go-14as.17`) |
| `SDR-003` | `io` (`haxe.io.Bytes*`, stream helpers, encoding edges) | Keep compiler ownership for representation-sensitive core and reopen algorithmic helpers for runtime/helper extraction. `Bytes.ofHex`, `Bytes.toHex`, and `BytesBuffer` leaf helpers now live in `hxrt`; inherited `Input` / `Output` loop helpers now live in staged source (`GoIoHelpers.hx`); RawNative/cache-coupled string paths and numeric IO primitives remain compiler-owned because they jointly enforce raw-native mode dispatch and `__hx_raw` cache invalidation semantics. | Compiler builtin lowering, extern/runtime package, canonical `std/go/_std`, framework-owned helper layers | Semantic-diff contracts + shim-vs-direct benchmark (`test:perf:stdlib-shims`) + post-`__go__` audit (`haxe.go-14as.50`) + IO extractions (`haxe.go-14as.51`, `haxe.go-14as.52`) + ownership lock `stdlib/bytes_raw_native_compiler_ownership` (`haxe.go-14as.54`) |
| `SDR-004` | `ds` (`haxe.ds.*Map`, `List`) | Keep compiler-owned shape generation until typed null/reflect parity is proven in a replacement path | canonical `std/go/_std`, extern-backed containers, pure runtime wrappers | Semantic-diff contracts (`ds_maps_list_contract`, map/list follow-ups) |
| `SDR-005` | `http` (`sys.Http`) | Keep compiler ownership for request/callback choreography, but move leaf helpers into staged helper islands when they only need same-package generated-shape access. `getResponseHeaderValues` and payload capture now route through `std/sys/GoHttpHelpers.hx`; proxy URL construction stays compiler-owned. | Extern-only wrappers, raw-native app code, canonical `std/go/_std`, framework-owned raw helper islands | Semantic-diff contracts (`http_request_callbacks_contract`, proxy/custom request tests) + post-`__go__` audit (`haxe.go-14as.50`) + staged helper extraction (`haxe.go-14as.53`) |
| `SDR-006` | `regex_serializer` (`EReg`, serializer/unserializer stack) | Keep compiler shims; revisit only with equivalent metadata-aware runtime path | Runtime-only package, extern-only, canonical `std/go/_std` | Serializer and regex semantic-diff suite |
| `SDR-007` | `net_socket` (`sys.net.Host`, `sys.net.Socket`, `sys.net.UdpSocket`) | Keep compiler shims for deadline/shutdown/readiness semantics | Extern-only wrappers, canonical `std/go/_std` | Socket semantic-diff contracts (loopback + advanced) plus direct UDP loopback snapshot |
| `SDR-008` | library-expressible `stdlib_symbols` surfaces (`StringTools`, `DateTools` helpers, `haxe.io.Path`, direct `haxe.SysTools`) | Default to staged std/runtime ownership; do not grow compiler-resident helper logic unless staged std/runtime is proven insufficient. `StringTools`, `DateTools`, and `haxe.SysTools` are completed staged-std migrations in this bucket; `haxe.io.Path` has graduated back to the upstream Haxe stdlib implementation. | Keep in compiler core, extern-only wrappers | Cross-target precedent (`haxe.rust` `std/StringTools.cross.hx` + runtime helper; `haxe.elixir` `std/StringTools.cross.hx`, `std/DateTools.cross.hx`) plus local migration beads (`haxe.go-14as.33`-`haxe.go-14as.36`, `haxe.go-14as.25`) |
| `SDR-009` | `haxe.Constraints` + `haxe.Rest` direct abstraction surfaces | Split ownership: staged std for `haxe.Constraints` typing/native bridge metadata, compiler lowering for native-slice `haxe.Rest` operations | Keep as compile-only debt, full compiler shim blobs, staged std abstract emission | Direct parity contracts (`haxe_constraints_contract`, `haxe_rest_contract`) plus snapshot `stdlib/haxe_constraints_rest_direct` |
| `SDR-010` | stack/main-loop `haxe.misc` surfaces (`haxe.CallStack`, `haxe.NativeStackTrace`, `haxe.EntryPoint`, `haxe.MainLoop`, `haxe.Timer`) | Split the tranche honestly: keep deterministic stack fallbacks in staged std under target-sensitive snapshot coverage, and support direct event-loop APIs through staged std wrappers over `sys.thread.EventLoop` instead of compiler-owned shims. Native Go stack capture is available only as the explicit `reflaxe_go_native_stack_trace` diagnostic capability, not as portable semantic-diff parity. | Leave as generic compile-only debt, force semantic-diff on target-timed event scheduling prematurely, grow compiler-owned stack/event-loop shims without a runtime ownership decision | Snapshot contracts `stdlib/haxe_stack_loop_target_sensitive`, `stdlib/haxe_native_stack_trace_opt_in`, and `stdlib/haxe_main_loop_runtime_direct`; staged std overrides `std/go/_std/haxe/CallStack.hx`, `std/go/_std/haxe/NativeStackTrace.hx`, `std/go/_std/haxe/EntryPoint.hx`, `std/go/_std/haxe/MainLoop.hx`, `std/go/_std/haxe/Timer.hx`; native stack spike `docs/spikes/native-stack-capture-contract.md`; cross-target precedent from `haxe.rust/std/haxe/CallStack.cross.hx` |
| `SDR-011` | legacy text surfaces (`haxe.Utf8`, `haxe.Ucs2`) | Use staged std for deprecated `haxe.Utf8` helper semantics; keep `haxe.Ucs2` as explicit target-sensitive platform exclusion under snapshot coverage | Grow compiler-owned text shims, leave both as anonymous compile-only debt | `haxe_utf8_contract`, `stdlib/haxe_utf8_basic`, `stdlib/haxe_ucs2_platform_exclusion`, staged std override `std/go/_std/haxe/Utf8.hx` |
| `SDR-012` | framework-owned raw-injection helper islands (`@:goAllowRaw` + `reflaxe.go.macros.GoInjection.__go__`) | Use this as the preferred middle layer when helper logic needs same-package generated-type access but does not need compiler-context decisions. Keep imports typed; do not use raw injection as a substitute for extern metadata. | Grow compiler-owned `GoRaw` blocks, extern-only wrappers, app-side raw `__go__` | `haxe.go-14as.48`, `haxe.go-14as.50`, sibling precedent from `reflaxe.rust` (`RustInjection.__rust__`) and `haxe.ocaml` (`__ocaml__` escape-hatch policy) |
| `SDR-013` | direct `haxe.rtti.*` support (`CType`, `Meta`, `Rtti`, `XmlParser`) | Keep public RTTI/parser logic in staged std, with a narrow compiler-owned metadata/lowering contract underneath | Move RTTI into compiler-owned stdlib authorities, invent RTTI-specific Go carrier structs, or classify the whole family unsupported | `haxe_rtti_direct_contract`, `stdlib/haxe_rtti_direct`, `haxe.go-14as.57`, `haxe.go-14as.59` |
| `SDR-014` | portable exception carrier and `sys.thread` process-lifecycle boundary | Keep the public API in staged std, require portable runtime validation failures to use `Throw`, and own worker identity/TLS state, recovery, and reporting in `hxrt`. The compiler adds a feature-gated generated-main foreground drain and, only when `sys.thread` is reachable, a cleanup scope around detached `go.Go.spawn` callbacks. Recover only explicit Haxe exception carriers; the detached scope does not join or recover native panics. | Treat every panic as a Haxe value, use raw panic for portable validation failures, retain TLS in per-instance global maps, let uncaught Haxe throws crash the Go process, join all native goroutines, or make portable workers daemon-like | `core/portable_runtime_failure_haxe_catch`, `semantic_diff/sys_thread_primitives_contract`, `stdlib/sys_thread_uncaught_exception`, `go_native/native_panic_not_haxe_catch`, `go_native/goroutine_native_panic`, `go_native/goroutine_native_shutdown`, direct `runtime/hxrt` lifecycle/race tests, and Haxe 4.3.7 interpreter probe recorded in `haxe_go-vfp.10.1` |

Review trigger for all records: revisit when an alternative path proves equal/better parity and performance under the same harness gates.

## Post-`__go__` Ownership Audit (2026-03-07)

The audit rule after restoring backend-owned `__go__` is simple:

- if a helper only stayed in `GoCompiler` because same-package generated-type access was awkward before, that is no longer a good enough reason by itself;
- prefer typed extern metadata when a real Go API exists;
- otherwise prefer framework-owned `@:goAllowRaw` + `reflaxe.go.macros.GoInjection.__go__` helper islands in `std/` or runtime code;
- keep compiler ownership only when lowering still depends on compiler context, profile policy, or representation decisions that Haxe source cannot express cleanly.

This matches sibling-target practice:

- `reflaxe.rust` already treats `__rust__` as a framework-only escape hatch and uses typed wrappers or `std/` helper layers before growing compiler-owned raw lowering;
- `haxe.ocaml` already documents `__ocaml__` as a controlled target-layer escape hatch instead of a default application coding style.

Sibling rollout note: `docs/spikes/family-raw-injection-authority-alignment.md`
captures the cross-repo handoff. It explicitly names `haxe.rust` and
`haxe.ocaml`, and tells sibling agents to compare against each compiler's own
architecture instead of copying Go-specific metadata names or helper shapes.

### Candidate List

| Candidate | Current owner | Recommended owner | Why | Follow-up |
| --- | --- | --- | --- | --- |
| `haxe.io.Bytes` algorithmic helpers (`ofHex`, `toHex`, `BytesBuffer` leaf helpers) | `hxrt` runtime helpers + thin compiler wrappers | Runtime-owned (`hxrt`) with compiler wrappers | These helpers only needed same-package generated-type access, not compiler context. The first migration moved them into `runtime/hxrt/bytes.go` so generated `haxe__io__Bytes` wrappers stay thin. | `haxe.go-14as.51` |
| `haxe.io.Bytes` RawNative/cache-coupled string helpers (`ofString`, `getString`, UTF16/raw-native conversions) | `GoCompiler` `io` shim group | Keep compiler-owned | These paths co-own both raw-native mode dispatch and `__hx_raw` cache validity. Public behavior such as RawNative `Base64.encode` after `Bytes.set(...)` depends on that cache invalidation staying coupled to the generated `haxe__io__Bytes` representation, so this is not just a misplaced helper island. | `haxe.go-14as.54` |
| `haxe.Resource` embedded payload table (`content`) | `GoCompiler` resource-table literal + source-owned std methods | Keep compiler-owned data population | The std `haxe.Resource` methods lower normally, but the actual payloads come from compiler resources (`Context.getResources()` / `__resources__()`), not reusable Haxe source. The backend must materialize `haxe__Resource_content`; otherwise direct resource calls compile but runtime sees an empty table. | `haxe.go-14as.30` |
| `haxe.io.Input` / `haxe.io.Output` helper loops (`readAll`, `readLine`, `readUntil`, `write`, `writeInput`, `readFullBytes`, `writeString`) | staged source helper + thin compiler wrappers | Keep the staged helper and only leave forwarding wrappers in `GoCompiler` | These are generic control-flow helpers around generated IO types. They now live in `std/haxe/io/GoIoHelpers.hx`, while `GoCompiler` keeps only the stable public wrapper functions and the representation-sensitive base IO types. | `haxe.go-14as.52` |
| `sys.Http` payload/proxy leaf helpers | split between `GoCompiler` and `std/sys/GoHttpHelpers.hx` | Keep core choreography in compiler; extract leaf helpers where typed imports + framework-owned `__go__` reduce raw compiler bulk | Request lifecycle, callback order, and response normalization are still one semantic contract. `getResponseHeaderValues` and payload capture were the first safe leaf extraction; proxy URL construction still stays compiler-owned because it returns native `*url.URL` transport state. | `haxe.go-14as.53` |
| `regex_serializer` (`EReg`, serializer/unserializer stack) | `GoCompiler` `regex_serializer` shim group | Keep compiler-owned | Heavy metadata coupling, resolver/reflection semantics, and token-stream behavior still make this a compiler-context surface. | Keep; revisit only with an equivalent metadata-aware runtime path |
| `net_socket` (`sys.net.Host`, `sys.net.Socket`, `sys.net.UdpSocket`) | `GoNetSocketEmitter` via the compiler-owned `net_socket` shim group | Keep compiler-owned | Deadline/select/shutdown/readiness behavior is still target-specific and tightly contract-tested, and UDP address translation now lives beside that same ownership slice. Extracting raw leaf helpers now would add churn without reducing real ownership complexity. | Keep; revisit only with a typed replacement path that preserves readiness semantics |

## Ownership Boundary (Post `haxe.go-7zy.11`)

- `runtime/hxrt/sys.go` and `runtime/hxrt/process.go` own `Sys`/`sys.io.File`/`sys.io.Process` behavior (OS args/environment/cwd, seconds-based blocking sleep, command delegation/exit, fallible file reads/writes, process launch and all three streams, PID/exit status, kill, and non-killing close/reap).
- `src/reflaxe/go/GoCompiler.hx` owns lowering, generated type shape, and translation from typed runtime status into Haxe exception/EOF/nullable stream contracts for this surface.
- `lowerSysStdlibShimDecls` must remain adapter-only: OS/process behavior belongs in `hxrt` unless a change is intentionally re-centralized and justified with parity/perf evidence.
- Portable `Sys.putEnv` explicitly discards the error retained by `hxrt.SysPutEnv` because Haxe 4.3.7 eval exposes a non-throwing `Void` contract; native Go-facing APIs may consume that error directly.
- `Sys.sleep` deliberately keeps the upstream root declaration instead of duplicating the whole class in staged std. The generated `Sys_sleep` adapter is typed Go AST, while `hxrt.SysSleep` owns Go's `time.Duration` conversion. This is the same source-adapter/runtime split used by `haxe.rust`; it does not justify compiler-owned raw statements.

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
3. Keep behavior-heavy shim groups in compiler core until an equivalent `std/go/_std` path proves equal parity under semantic-diff coverage.

## Revisit Triggers

Re-open keep decisions when one of these becomes true:

1. A canonical `std/go/_std` path reaches equal or better parity for the same fixtures.
2. Runtime package extraction can preserve source semantics and typed lowering
   policy without drift.
3. A compiler shim becomes pure forwarding with no compiler-context decisions.
