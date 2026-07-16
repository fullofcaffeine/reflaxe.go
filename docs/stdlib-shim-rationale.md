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

This document records which compiler-core shims have migrated, which are still
explicit migration debt, and which exact compiler primitives have enough
compile-context evidence to remain.
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

Use the first applicable ownership layer that preserves parity:

- staged `std/go/_std` for Haxe-visible library behavior,
- runtime (`hxrt`) for reusable target-runtime behavior,
- compiler intrinsics only for exact compile-time metadata, policy, or
  representation primitives that neither source layer can express.

Existing behavior-heavy compiler groups are compatibility implementations, not
approved architecture. They remain measured and tested while their linked
migration beads move them to the proper source/runtime owner.

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
| Canonical target `_std` (`std/go/_std`) | Authoritative home for Haxe-visible behavior | Remaining families need typed runtime handles and representation bridges migrated incrementally; their size is execution scope, not a reason for compiler ownership. |

## Decision Matrix

`Compiler LOC` values below are historical shim-function spans measured on 2026-02-19 unless a row explicitly reports its current post-migration span and debt-ratchet count.

| Shim group | Primary surfaces | Compiler LOC | Highest CI tier | Decision | Reason | Follow-up |
| --- | --- | ---: | --- | --- | --- | --- |
| `json` | `haxe.Json`, `haxe.format.JsonParser/JsonPrinter` | 38 | Snapshot | Migrated (canonical staged std + runtime-owned behavior) | Staged std source under `std/go/_std/**` owns JSON API surfaces; behavior delegates to `hxrt.JsonParse`/`hxrt.JsonStringify`. | `haxe.go-7zy.10`, `haxe.go-cgk.5` |
| `sys` | root `Sys` | 0 (compiler ownership retired) | Semantic-diff + real PTY | Migrated (canonical staged std + typed runtime capabilities) | `std/go/_std/Sys.hx` owns the public API, map construction, fallbacks, aliases, stream wrapping, `getChar` EOF construction, and echo policy. Typed `std/hxrt` bindings expose capabilities in `runtime/hxrt/sys.go`, `file.go`, build-tagged `terminal*.go`, and the baseline print slice. Only the honest compile-time `cpuTime` rejection remains in the compiler. | `haxe.go-7zy.11`, `haxe_go-vfp.8.7.1`, `haxe_go-vfp.8.7.2`, `haxe_go-vfp.8.7.3`, `haxe_go-vfp.8.7.6` |
| `process` | `sys.io.Process` | 0 (compiler ownership retired) | Semantic-diff | Migrated (canonical staged std + typed runtime capabilities) | `std/go/_std/sys/io/Process.hx` owns the public API, stream adapters, bounds/EOF translation, nullable status, detached rejection, and lifecycle policy. Typed `std/hxrt/process` bindings expose opaque handles and native capabilities in selectively copied `runtime/hxrt/process.go`; no Process declaration emitter or branch remains in `GoCompiler`. | `haxe.go-7zy.11`, `haxe_go-vfp.8.7.7` |
| `file_io` | `sys.io.File`, `FileInput`, `FileOutput`, `FileSeek` | 0 (compiler ownership retired) | Semantic-diff | Migrated (canonical staged std + typed runtime capabilities) | `std/go/_std/sys/io` owns the public Haxe API, byte conversion, bounds/EOF behavior, and seek mapping. Typed `std/hxrt/fs` bindings expose opaque handles and native operations in selectively copied `runtime/hxrt/file.go`. No File declarations, handle maps, seek mappers, methods, or File-specific subclass branches remain in `GoCompiler`. | `haxe.go-14as.17`, `haxe_go-vfp.8.7.5` |
| `atomic` | `haxe.atomic.*` operations | 0 (compiler ownership retired) | Semantic-diff + race | Migrated (mainstream/staged Haxe + typed runtime capabilities) | Mainstream Haxe continues to own `AtomicBool`; canonical target overrides implement only the body-less `AtomicInt` and `AtomicObject` core types over opaque typed handles. Native ordering, storage, and reference comparison remain in `hxrt`; no atomic compiler group or direct lowering remains. | `haxe_go-vfp.8.7.9` |
| `io` | `haxe.io.Bytes`, buffers, input/output base wiring | 108 | Snapshot + semantic-diff dependency | Migration required; retain only separately proven representation primitives | Raw byte storage and cache invalidation may need a typed representation boundary, but public types, validation, loops, encoding policy, EOF, and numeric IO are staged-source behavior. | `haxe_go-vfp.8.7.11` |
| `ds` | `haxe.ds.*Map`, `List`, enum maps | 0 (compiler group retired) | Snapshot + semantic-diff | Migrated (canonical staged std + typed runtime storage capabilities) | Ordinary Haxe owns every public collection API and algorithm. Typed `hxrt` handles retain only integer/string hash storage, object identity, deterministic snapshots, and generated-enum recognition. Direct `Lambda`/sort lowering is separate follow-up scope. | `haxe_go-vfp.8.7.10`; direct algorithms `haxe_go-vfp.8.7.17` |
| `http` | `sys.Http` request/callback/proxy contract | 542 | Semantic-diff | Migration required | Transport resources belong in `hxrt`; request sequencing and callback policy belong in staged Haxe. Existing staged leaf helpers prove the direction. | `haxe_go-vfp.8.7.12` |
| `filesystem` | `sys.FileSystem` | 0 (compiler group retired) | Semantic-diff | Migrated (canonical staged std + typed runtime capabilities) | `std/go/_std/sys/FileSystem.hx` owns the complete Haxe API and constructs `sys.FileStat`; `std/hxrt/fs` provides typed bindings to native operations in selectively copied `runtime/hxrt/filesystem.go`. No compiler filesystem declarations or imports remain. | `haxe_go-vfp.8.7.4` |
| `stdlib_symbols` | `Std`, `Date`, `Math`, `Reflect`, crypto/xml/zip | 706 | Semantic-diff | Migration required except exact registered metadata/representation primitives | Date, Math, DOM, codecs, hashes, parsing, printing, and compression are source/runtime behavior. Only the separately registered type metadata, type-test, string representation, exception carrier, and Rest primitives are admitted. | `haxe_go-vfp.8.7.15` |
| `regex_serializer` | `EReg`, `haxe.Serializer`, `haxe.Unserializer` | 2460 | Semantic-diff | Migration required; split generated metadata from runtime algorithms | Reachable type metadata may remain generated data, but regex matching, token streams, resolver policy, traversal, and identity behavior do not belong in a compiler emitter. | `haxe_go-vfp.8.7.13` |
| `net_socket` | `sys.net.Host`, `sys.net.Socket`, `sys.net.UdpSocket` | 2958 | Mixed (`semantic-diff` for TCP host/socket, snapshot for direct UDP loopback) | Migration required | Deadline, readiness, address translation, shutdown, broadcast, and TLS composition are typed runtime-handle behavior beneath staged public APIs. | `haxe_go-vfp.8.7.14` |
| `template_support` | runtime `Reflect` helpers beneath `haxe.Template` | 3 helpers | Semantic-diff | Migration required | Property lookup, object checks, and method invocation use runtime values and no closed compiler metadata. | `haxe_go-vfp.8.7.16` |

## Explicit Decision Records

These are the canonical per-surface decisions for shim ownership and alternatives.

| Record | Surface | Decision | Alternatives reviewed | Evidence |
| --- | --- | --- | --- | --- |
| `SDR-001` | `json` (`haxe.Json`, `haxe.format.Json*`) | Move API ownership to canonical staged std (`std/go/_std`) and keep behavior in `hxrt` (`JsonParse`/`JsonStringify`) | Compiler shim, direct lower-call rewrites, extern/runtime package | Snapshot parity + migration log (`haxe.go-7zy.10`, `haxe.go-cgk.5`) |
| `SDR-002` | `process` (`sys.io.Process`) | Keep the complete public API and stream policy in canonical staged Haxe; cross only typed opaque handles and native spawn/pipe/wait/signal/close capabilities through `std/hxrt/process` and `runtime/hxrt/process.go`. Compiler ownership is retired. | Keep the former combined Sys/Process group, keep the isolated Process emitter, expose the public stdlib as externs, use raw injection in the override, or move Haxe bounds/EOF/null policy into `hxrt` | Process semantic-diff contracts, direct runtime tests, selective `core/runtime_hxrt_infer_process`, sibling staged-source precedent, and the permanent 134-site `GoRaw` reduction under `haxe_go-vfp.8.7.7` |
| `SDR-003` | `io` (`haxe.io.Bytes*`, stream helpers, encoding edges) | Treat the broad compiler group as migration debt. Move public behavior to staged Haxe and raw storage to typed runtime capabilities; admit a residual representation primitive only after it is separated and individually justified. | Preserve the whole group, move the blob unchanged into `hxrt`, canonical staged source plus opaque typed storage | Existing semantic-diff, RawNative cache, file/process subclass, and performance evidence becomes the regression baseline for `haxe_go-vfp.8.7.11`. |
| `SDR-004` | `ds` (`haxe.ds.*Map`, `List`) | Keep public APIs, iteration, copies, ordering policy, enum comparison, and list algorithms in staged Haxe. Use narrow typed native storage only for hash representation, object identity, and generated-enum recognition; the compiler `ds` group is retired. | Preserve generated map/list classes, extern-only containers, move the former compiler blob into `hxrt`, staged source plus typed storage capabilities | `ds_maps_list_contract`, `map_keyvalue_iterator_contract`, `serializer_extended_tokens_contract`, four selective collection-runtime snapshots; source migration `haxe_go-vfp.8.7.10`, separate direct-algorithm follow-up `haxe_go-vfp.8.7.17` |
| `SDR-005` | `http` (`sys.Http`) | Move Haxe-visible request choreography to staged source and native HTTP resources to typed `hxrt`; existing source-owned leaf helpers are an intermediate step, not the final split. | Keep choreography in compiler, raw app code, staged source plus typed transport handles | HTTP callback/proxy/custom-request contracts; `haxe_go-vfp.8.7.12` |
| `SDR-006` | `regex_serializer` (`EReg`, serializer/unserializer stack) | Split exact generated metadata from runtime behavior, then move public regex and serialization algorithms to staged Haxe plus typed `hxrt`. | Keep the mixed compiler emitter, runtime-only blob, staged algorithms plus generated metadata data/hooks | Serializer and regex semantic-diff suite; `haxe_go-vfp.8.7.13` |
| `SDR-007` | `net_socket` (`sys.net.Host`, `sys.net.Socket`, `sys.net.UdpSocket`) | Move public APIs to staged source and OS socket lifecycle/readiness/deadline behavior to typed runtime handles. | Keep compiler classes, extern-only wrappers, staged source plus typed handles | Socket semantic-diff contracts, direct UDP loopback snapshot, and `haxe_go-vfp.8.7.14` |
| `SDR-008` | `stdlib_symbols` library surfaces | Continue completed staged migrations and retire the remaining monolithic Date/Math/Reflect/Xml/crypto/zip block. Only exact metadata and representation primitives in the intrinsic registry may remain. | Preserve the mixed blob, extern-only wrappers, staged source plus narrow runtime helpers | Sibling-target precedent, existing local migrations, and `haxe_go-vfp.8.7.15` |
| `SDR-009` | `haxe.Constraints` + `haxe.Rest` direct abstraction surfaces | Split ownership: staged std for `haxe.Constraints` typing/native bridge metadata, compiler lowering for native-slice `haxe.Rest` operations | Keep as compile-only debt, full compiler shim blobs, staged std abstract emission | Direct parity contracts (`haxe_constraints_contract`, `haxe_rest_contract`) plus snapshot `stdlib/haxe_constraints_rest_direct` |
| `SDR-010` | stack/main-loop `haxe.misc` surfaces (`haxe.CallStack`, `haxe.NativeStackTrace`, `haxe.EntryPoint`, `haxe.MainLoop`, `haxe.Timer`) | Split the tranche honestly: keep deterministic stack fallbacks in staged std under target-sensitive snapshot coverage, and support direct event-loop APIs through staged std wrappers over `sys.thread.EventLoop` instead of compiler-owned shims. Native Go stack capture is available only as the explicit `reflaxe_go_native_stack_trace` diagnostic capability, not as portable semantic-diff parity. | Leave as generic compile-only debt, force semantic-diff on target-timed event scheduling prematurely, grow compiler-owned stack/event-loop shims without a runtime ownership decision | Snapshot contracts `stdlib/haxe_stack_loop_target_sensitive`, `stdlib/haxe_native_stack_trace_opt_in`, and `stdlib/haxe_main_loop_runtime_direct`; staged std overrides `std/go/_std/haxe/CallStack.hx`, `std/go/_std/haxe/NativeStackTrace.hx`, `std/go/_std/haxe/EntryPoint.hx`, `std/go/_std/haxe/MainLoop.hx`, `std/go/_std/haxe/Timer.hx`; native stack spike `docs/spikes/native-stack-capture-contract.md`; cross-target precedent from `haxe.rust/std/haxe/CallStack.cross.hx` |
| `SDR-011` | legacy text surfaces (`haxe.Utf8`, `haxe.Ucs2`) | Use staged std for deprecated `haxe.Utf8` helper semantics; keep `haxe.Ucs2` as explicit target-sensitive platform exclusion under snapshot coverage | Grow compiler-owned text shims, leave both as anonymous compile-only debt | `haxe_utf8_contract`, `stdlib/haxe_utf8_basic`, `stdlib/haxe_ucs2_platform_exclusion`, staged std override `std/go/_std/haxe/Utf8.hx` |
| `SDR-012` | framework-owned raw-injection helper islands (`@:goAllowRaw` + `reflaxe.go.macros.GoInjection.__go__`) | Use this as the preferred middle layer when helper logic needs same-package generated-type access but does not need compiler-context decisions. Keep imports typed; do not use raw injection as a substitute for extern metadata. | Grow compiler-owned `GoRaw` blocks, extern-only wrappers, app-side raw `__go__` | `haxe.go-14as.48`, `haxe.go-14as.50`, sibling precedent from `reflaxe.rust` (`RustInjection.__rust__`) and `haxe.ocaml` (`__ocaml__` escape-hatch policy) |
| `SDR-013` | direct `haxe.rtti.*` support (`CType`, `Meta`, `Rtti`, `XmlParser`) | Keep public RTTI/parser logic in staged std, with a narrow compiler-owned metadata/lowering contract underneath | Move RTTI into compiler-owned stdlib authorities, invent RTTI-specific Go carrier structs, or classify the whole family unsupported | `haxe_rtti_direct_contract`, `stdlib/haxe_rtti_direct`, `haxe.go-14as.57`, `haxe.go-14as.59` |
| `SDR-014` | portable exception carrier and `sys.thread` process-lifecycle boundary | Keep the public API in staged std, require portable runtime validation failures to use `Throw`, and own worker identity/TLS state, recovery, and reporting in `hxrt`. The compiler adds a feature-gated generated-main foreground drain and, only when `sys.thread` is reachable, a cleanup scope around detached `go.Go.spawn` callbacks. Recover only explicit Haxe exception carriers; the detached scope does not join or recover native panics. | Treat every panic as a Haxe value, use raw panic for portable validation failures, retain TLS in per-instance global maps, let uncaught Haxe throws crash the Go process, join all native goroutines, or make portable workers daemon-like | `core/portable_runtime_failure_haxe_catch`, `semantic_diff/sys_thread_primitives_contract`, `stdlib/sys_thread_uncaught_exception`, `go_native/native_panic_not_haxe_catch`, `go_native/goroutine_native_panic`, `go_native/goroutine_native_shutdown`, direct `runtime/hxrt` lifecycle/race tests, and Haxe 4.3.7 interpreter probe recorded in `haxe_go-vfp.10.1` |
| `SDR-015` | `sys.FileSystem` | Keep the complete public API and `FileStat` construction in canonical staged Haxe; expose only typed native filesystem capabilities through `std/hxrt/fs` and `runtime/hxrt/filesystem.go`. Compiler ownership is retired, and runtime slicing excludes filesystem support from unrelated `sys.*` programs. | Keep the multi-function `GoCompiler` shim group, duplicate the upstream `FileStat` typedef, put the helpers in the broad `sys.go` runtime slice, or use raw `__go__` in the override | Complete Haxe 4.3.7 surface, snapshot + semantic-diff filesystem contracts, direct runtime tests, selective-runtime evidence, and compiler-debt reduction under `haxe_go-vfp.8.7.4` |
| `SDR-016` | `sys.io.File`, `FileInput`, `FileOutput`, `FileSeek` | Keep the complete public API and stream semantics in canonical staged Haxe; cross only typed opaque handles and native file capabilities through `std/hxrt/fs` and `runtime/hxrt/file.go`. Retain exact public base-I/O fields at macro initialization solely so indirectly staged subclasses can type-check after DCE. | Keep File declarations and algorithms in the former combined `lowerSysStdlibShimDecls`, expose generated `Bytes` internals to `hxrt`, put file support in broad `sys.go`, or use raw injection in the override | Binary/text/copy/write/append/update/seek/tell/bounds/EOF semantic-diff coverage, direct snapshots/runtime tests, selective-runtime evidence, and permanent `GoRaw` reductions under `haxe_go-vfp.8.7.5` |
| `SDR-017` | root `Sys` | Keep the complete supported public API in canonical `std/go/_std/Sys.hx`; use narrow typed console, process-state, environment, path, clock, command, and standard-file capabilities. Inline one-step source methods for Go-shaped direct calls, while first-class references materialize the same source-owned contract. | Retain generated root declarations in the combined shim, use externs as the public stdlib API, put map/alias/fallback behavior in `hxrt`, or substitute wall time for `cpuTime` | Root semantic-diff and negative contracts, direct runtime tests, separate Sys-only/Process-only selective snapshots, sibling `haxe.rust` staged-source precedent, and compiler-debt reduction under `haxe_go-vfp.8.7.6` |
| `SDR-018` | `Sys.getChar` terminal behavior | Keep EOF construction and requested echo in staged `Sys`; expose only serialized terminal-state control and one-byte native input through typed `NativeTerminal` and a footprint-explicit, build-tagged `terminal` runtime slice. Confine the POSIX `unsafe.Pointer` to one typed ioctl helper with exact debt ceilings. | Reintroduce a compiler shim, use ordinary line-buffered `FileInput.readByte`, add raw Haxe injection, depend on a current `x/term` that raises the generated Go floor, or pin the older advisory-bearing dependency | Real-PTY no-newline/echo/restoration behavior, redirected EOF, Linux/macOS/Windows implementation cross-builds, an unsupported-FreeBSD cross-build, selective snapshot evidence, race/checkptr gates, and the reviewed unsafe ratchet under `haxe_go-vfp.8.7.3` |
| `SDR-019` | portable compiler stdlib intrinsic registry | Inventory exact Haxe symbols, compiler entry points, selectors, direct call rewrites, dependencies, evidence, review conditions, and migration beads in a fail-closed machine-readable registry. Classify behavior-heavy groups as avoidable migration debt; distinguish narrow admitted intrinsics from explicit `go.*` native emitters. | Keep the file ledger as an indirect audit, treat every named shim as required, or ban all compiler primitives without recognizing type-metadata/representation needs | `docs/compiler-stdlib-intrinsics.json`, `test/test_compiler_stdlib_intrinsic_registry.py`, and the split compiler-debt exceptions under `haxe_go-vfp.8.7.8` |

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
| `haxe.io.Bytes` RawNative/cache-coupled string helpers (`ofString`, `getString`, UTF16/raw-native conversions) | `GoCompiler` `io` shim group | Move public policy to staged source; retain only a typed storage/cache primitive if still necessary | Cache invalidation is real representation evidence, but it justifies a narrow opaque byte-store capability rather than compiler ownership of the complete Bytes API. | `haxe_go-vfp.8.7.11` |
| `haxe.Resource` embedded payload table (`content`) | `GoCompiler` resource-table literal + source-owned std methods | Keep compiler-owned data population | The std `haxe.Resource` methods lower normally, but the actual payloads come from compiler resources (`Context.getResources()` / `__resources__()`), not reusable Haxe source. The backend must materialize `haxe__Resource_content`; otherwise direct resource calls compile but runtime sees an empty table. | `haxe.go-14as.30` |
| `haxe.io.Input` / `haxe.io.Output` helper loops (`readAll`, `readLine`, `readUntil`, `write`, `writeInput`, `readFullBytes`, `writeString`) | staged source helper + thin compiler wrappers | Remove the remaining compiler wrappers with the base IO migration | The algorithms already live in `std/haxe/io/GoIoHelpers.hx`; keeping generated public forwarding bodies is transitional ownership, not a compile-context requirement. | `haxe_go-vfp.8.7.11` |
| `sys.Http` payload/proxy leaf helpers | split between `GoCompiler` and `std/sys/GoHttpHelpers.hx` | Move all Haxe-visible choreography to staged source and native URL/transport state to typed `hxrt` | One semantic contract is a testing reason to migrate coherently, not a compiler-ownership reason. | `haxe_go-vfp.8.7.12` |
| `regex_serializer` (`EReg`, serializer/unserializer stack) | `GoCompiler` `regex_serializer` shim group | Split and migrate | Only final reachable-type metadata may remain compiler-generated; regex and token-stream algorithms move to source/runtime owners. | `haxe_go-vfp.8.7.13` |
| `net_socket` (`sys.net.Host`, `sys.net.Socket`, `sys.net.UdpSocket`) | `GoNetSocketEmitter` via the compiler-owned `net_socket` shim group | Move to staged APIs over typed runtime handles | Target-sensitive OS behavior belongs in `hxrt`; deterministic parity tests guard the migration rather than justify compiler classes. | `haxe_go-vfp.8.7.14` |

## Ownership Boundary (Post `haxe_go-vfp.8.7.7`)

- `runtime/hxrt/sys.go`, `runtime/hxrt/file.go`, and `runtime/hxrt/process.go` own distinct native capabilities: root process state, file/standard-stream handles, and child processes respectively.
- `std/go/_std/sys/io` owns the File and Process public APIs and stream semantics, and `std/go/_std/Sys.hx` owns root `Sys`. None of these families has compiler-generated declarations or semantic branches.
- `std/hxrt/process` exposes only typed opaque process/pipe handles plus native capability calls. OS child-process behavior belongs in `runtime/hxrt/process.go`; Haxe-visible bounds, EOF, nullable status, detached rejection, and lifecycle behavior remain reviewable staged source.
- Staged `Sys.putEnv` explicitly selects the non-throwing `SysSetEnvironment` capability because Haxe 4.3.7 eval exposes `Void`; `hxrt.SysPutEnv` retains its native error for Go-native callers.
- Simple staged root methods inline to typed capabilities so direct output remains Go-shaped. First-class function references materialize source-owned functions instead of restoring compiler special cases.
- `Sys.cpuTime` is the explicit compile-context exception: it is rejected at Haxe compile time because no Go-standard-library implementation can satisfy the CPU-time contract. See [Portable root `Sys` contract](portable-sys-contract.md).

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
- this is a regression baseline for moving public Base64 behavior to staged
  source over a narrower byte-conversion capability; it does not approve the
  current `stdlib_symbols` group

## Migration Sequence

1. Move `json` out of compiler core first (`haxe.go-7zy.10`) and then promote staged std ownership (`haxe.go-cgk.5`) because it is the thinnest shim and lowest risk.
2. Move `sys` wrappers second (`haxe.go-7zy.11`, completed 2026-02-19) once snapshot parity remains stable.
3. Migrate the remaining behavior-heavy families through `haxe_go-vfp.8.7.9`
   to `.17`. Existing parity suites are the preservation oracle during that
   work, not a reason to keep compiler ownership.

## Migration Closure Triggers

Remove a migration-debt entry when one of these becomes true:

1. A canonical `std/go/_std` path reaches equal or better parity for the same fixtures.
2. Runtime package extraction can preserve source semantics and typed lowering
   policy without drift.
3. A compiler shim becomes pure forwarding with no compiler-context decisions.
