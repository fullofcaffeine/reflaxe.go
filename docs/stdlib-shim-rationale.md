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

This document records which compiler-core shims migrated and which exact
compiler primitives have enough compile-context evidence to remain. The
behavior-heavy migration ledger is closed: no portable compiler-stdlib
migration-debt exception remains.
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

Former behavior-heavy compiler groups were compatibility implementations, not
approved architecture, and have moved to their proper source/runtime owners.
The remaining declaration emitters are individually registered metadata or
representation capabilities; explicit `go.*` emitters are separately classified
native APIs. A shared selector may dispatch those registered capabilities, but
it owns no library behavior and is not a second stdlib layer.

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
| Compiler-lowered builtins/intrinsics only | Minimal generated wrappers for hot paths | Only viable when behavior depends on compile-time facts or an exact generated representation; ordinary token, regex, exception, and conversion policy belongs in source/runtime owners. |
| Externs + external Go runtime package | Clean boundary and reuse potential | Externs are type-only and `ignoreExterns: true` is currently required for deterministic emission in `src/reflaxe/go/CompilerInit.hx`. |
| Raw `__go__` in Haxe std/app code | Minimal indirection for target-native calls | Still the wrong default for app/examples, but now valid in framework-owned low-level abstraction islands via `@:goAllowRaw` + `reflaxe.go.macros.GoInjection.__go__`. Keep imports typed with extern metadata; do not use raw injection as a substitute for `@:go.import`/`@:go.name`. |
| Canonical target `_std` (`std/go/_std`) | Authoritative home for Haxe-visible behavior | Completed migrations show that typed runtime handles and narrow representation bridges preserve this ownership without compiler-emitted library algorithms. |

## Decision Matrix

`Compiler LOC` values below are historical shim-function spans measured on 2026-02-19 unless a row explicitly reports its current post-migration span and debt-ratchet count.

| Shim group | Primary surfaces | Compiler LOC | Highest CI tier | Decision | Reason | Follow-up |
| --- | --- | ---: | --- | --- | --- | --- |
| `json` | `haxe.Json`, `haxe.format.JsonParser/JsonPrinter` | 38 | Snapshot | Migrated (canonical staged std + runtime-owned behavior) | Staged std source under `std/go/_std/**` owns JSON API surfaces; behavior delegates to `hxrt.JsonParse`/`hxrt.JsonStringify`. | `haxe.go-7zy.10`, `haxe.go-cgk.5` |
| `sys` | root `Sys` | 0 (compiler ownership retired) | Semantic-diff + real PTY | Migrated (canonical staged std + typed runtime capabilities) | `std/go/_std/Sys.hx` owns the public API, map construction, fallbacks, aliases, stream wrapping, `getChar` EOF construction, and echo policy. Typed `std/hxrt` bindings expose capabilities in `runtime/hxrt/sys.go`, `file.go`, build-tagged `terminal*.go`, and the baseline print slice. Only the honest compile-time `cpuTime` rejection remains in the compiler. | `haxe.go-7zy.11`, `haxe_go-vfp.8.7.1`, `haxe_go-vfp.8.7.2`, `haxe_go-vfp.8.7.3`, `haxe_go-vfp.8.7.6` |
| `process` | `sys.io.Process` | 0 (compiler ownership retired) | Semantic-diff | Migrated (canonical staged std + typed runtime capabilities) | `std/go/_std/sys/io/Process.hx` owns the public API, stream adapters, bounds/EOF translation, nullable status, detached rejection, and lifecycle policy. Typed `std/hxrt/process` bindings expose opaque handles and native capabilities in selectively copied `runtime/hxrt/process.go`; no Process declaration emitter or branch remains in `GoCompiler`. | `haxe.go-7zy.11`, `haxe_go-vfp.8.7.7` |
| `file_io` | `sys.io.File`, `FileInput`, `FileOutput`, `FileSeek` | 0 (compiler ownership retired) | Semantic-diff | Migrated (canonical staged std + typed runtime capabilities) | `std/go/_std/sys/io` owns the public Haxe API, byte conversion, bounds/EOF behavior, and seek mapping. Typed `std/hxrt/fs` bindings expose opaque handles and native operations in selectively copied `runtime/hxrt/file.go`. No File declarations, handle maps, seek mappers, methods, or File-specific subclass branches remain in `GoCompiler`. | `haxe.go-14as.17`, `haxe_go-vfp.8.7.5` |
| `atomic` | `haxe.atomic.*` operations | 0 (compiler ownership retired) | Semantic-diff + race | Migrated (mainstream/staged Haxe + typed runtime capabilities) | Mainstream Haxe continues to own `AtomicBool`; canonical target overrides implement only the body-less `AtomicInt` and `AtomicObject` core types over opaque typed handles. Native ordering, storage, and reference comparison remain in `hxrt`; no atomic compiler group or direct lowering remains. | `haxe_go-vfp.8.7.9` |
| retired `io` | `haxe.io.Bytes`, buffers, input/output base wiring | 0 (compiler group retired) | Semantic-diff + snapshot + strict sweep + performance | Migrated (canonical staged std + typed representation capabilities) | Staged source owns all public types, validation, loops, encoding, EOF/error, endian, alias, and cache policy. Typed `std/hxrt/io` exposes only opaque byte views, native conversion/copy/UTF operations, and scalar IEEE-754 reinterpretation. | `haxe_go-vfp.8.7.11`; [ownership contract](haxe-io-ownership.md) |
| `ds` | `haxe.ds.*Map`, `List`, enum maps, complete `Lambda` API, sort helpers | 0 declaration shims; exact call adapters only | Snapshot + semantic-diff | Migrated (canonical/upstream Haxe + typed storage and representation capabilities) | Ordinary Haxe owns every public collection API and algorithm. Typed `hxrt` handles retain only native storage facts. Exact registered compiler adapters wrap Go's invariant iterable, callback, nested-carrier, array, comparator, and linked-node shapes without implementing traversal or sorting. `LambdaGoIterableCarrier` is a private representation-only staged companion. | `haxe_go-vfp.8.7.10`, `haxe_go-vfp.8.7.17`, `haxe_go-vfp.8.7.18` |
| retired `http` | `sys.Http` request/callback/proxy contract | 0 (compiler group retired) | Semantic-diff + snapshot + direct runtime | Migrated (canonical staged std + typed runtime capabilities) | `std/go/_std/sys/Http.hx` owns request selection, payload/header policy, callback order, response normalization, and status/error handling. Typed `std/hxrt/http` handles expose footprint-explicit Go URL/transport work in `runtime/hxrt/http.go`; no generated public HTTP type or helper island remains. | `haxe_go-vfp.8.7.12` |
| `filesystem` | `sys.FileSystem` | 0 (compiler group retired) | Semantic-diff | Migrated (canonical staged std + typed runtime capabilities) | `std/go/_std/sys/FileSystem.hx` owns the complete Haxe API and constructs `sys.FileStat`; `std/hxrt/fs` provides typed bindings to native operations in selectively copied `runtime/hxrt/filesystem.go`. No compiler filesystem declarations or imports remain. | `haxe_go-vfp.8.7.4` |
| `crypto` | `haxe.crypto.Base64`, `Md5`, `Sha1`, `Sha224`, `Sha256` | 0 compiler declarations | Semantic-diff + snapshot + direct runtime | Migrated (canonical staged std + typed runtime capabilities) | Staged Haxe owns public APIs and Base64 alphabets/padding. Footprint-explicit `runtime/hxrt/crypto.go` owns only native codec and digest execution over strings and opaque cached byte views. | `haxe_go-vfp.8.7.15.1`, byte-view sharing `haxe_go-vfp.8.7.11` |
| `xml` | root `Xml`, `haxe.xml.Parser`, `haxe.xml.Printer` | 0 compiler declarations | Semantic-diff + snapshot | Migrated (canonical/upstream staged Haxe) | `std/go/_std/Xml.hx` owns DOM storage, validation, mutation, parent links, and structural iteration; unchanged upstream source owns strict/non-strict parsing and structured errors; staged Printer preserves upstream formatting without an incidental `EReg` dependency. No native XML parser or compiler helper remains. | `haxe_go-vfp.8.7.15.2` |
| `zip` | `haxe.zip.Compress`, `haxe.zip.Uncompress` | 0 compiler declarations | Semantic-diff + snapshot + direct runtime/race | Migrated (canonical staged std + typed runtime capabilities) | Staged Haxe owns levels, buffer defaults, `Bytes` conversion, offsets/results, flush/lifecycle policy, one-shot helpers, and raw-DEFLATE selection. Footprint-explicit `runtime/hxrt/zip.go` retains live codecs behind opaque typed handles and returns bounded typed steps. Go supports exact `NO` / `SYNC` / `FINISH`; `FULL` / `BLOCK` fail explicitly instead of being silently downgraded. | `haxe_go-vfp.8.7.15.3`, streaming `haxe_go-vfp.8.7.21`; [contract](haxe-zip-streaming.md) |
| `date` | root `Date` | 0 compiler declarations | Semantic-diff + snapshot + direct runtime | Migrated (canonical staged std + typed runtime capabilities) | Staged Haxe owns the epoch-millisecond carrier and complete public API. Footprint-explicit `runtime/hxrt/date.go` owns only host clock, timezone, parsing, formatting, and calendar conversion over scalars and `DateParts`; staged Serializer consumes the public `Date.getTime()` contract without a compiler representation probe. | `haxe_go-vfp.8.7.15.4` |
| `math` | root `Math` | 0 compiler declarations | Semantic-diff + snapshot + direct runtime | Migrated (canonical staged std + typed native capabilities) | Staged Haxe owns Haxe rounding, finiteness, NaN propagation, and operand-order signed-zero behavior. Float operations bind directly to Go `math` / `math/rand`; footprint-explicit `runtime/hxrt/math.go` owns only three Int-returning signature adapters. | `haxe_go-vfp.8.7.15.4` |
| `unicode_string` | root `UnicodeString` | 0 compiler declarations | Semantic-diff + snapshot + direct runtime | Migrated (canonical staged std + typed representation capabilities) | Staged Haxe owns code-point bounds, slicing, searching, comparison, iteration, operators, and UTF-8 validation. Typed `GoStringRuntime` calls expose only rune length, code-point lookup, and already-normalized slicing over pointer-backed Go strings. | `haxe_go-vfp.8.7.15.5` |
| staged `Std` / `haxe.Log` | complete root `Std` API plus `haxe.Log.formatOutput` / mutable `trace` | 0 behavior-specific direct call rewrites | Semantic-diff + snapshot runtime + strict sweep + performance | Migrated (canonical staged std + typed representation/native capabilities) | Staged Haxe owns parsing, overflow, aliases, downcast policy, random bounds, trace formatting, position/custom parameters, rebinding, and `Sys.println`. `Std.string` and `Std.isOfType` remain exact registered representation intrinsics; narrow typed bindings perform native float conversion, truncation, and random generation. General typed class lowering—not a Log profile branch—owns mutable static dynamic functions and catchable null calls. | `haxe_go-vfp.8.7.22` |
| retired `stdlib_symbols` | former nominal `Std` carrier, `haxe.ds.Option`, and unrelated classifier/planner selectors | 0 | Snapshot + semantic-diff + strict sweep | Retired | `Option` follows ordinary source enum lowering. The empty `Std` carrier and false `BalancedTree`, `Path`, Template, and SSL selectors are gone. Type, Reflect, generated-method metadata, type tests, string/exception representations, and Rest remain separately registered exact capabilities. The later `Std`/`Log` migration retired both residual behavior-specific direct calls without recreating a compatibility group. | retirement `haxe_go-vfp.8.7.15.7`; source closure `haxe_go-vfp.8.7.22` |
| retired `regex_serializer` | `EReg`, serialization token/cache/traversal policy | 0 behavior-heavy compiler declarations | Semantic-diff + snapshot + direct runtime | Migrated (canonical staged std + typed runtime capabilities) | Staged Haxe owns regex state/policy and the full serialization algorithms. `regex.go` owns RE2 execution; `serialization.go` owns only bounded float parsing. Existing Reflect metadata owns typed generated field/method access, and Type metadata owns class/enum lookup plus safe constructor-free allocation. | `haxe_go-vfp.8.7.13`; typed accessor completion `haxe_go-vfp.10.5.1` |
| retired `serialization_source_bridge` | private/inherited fields, `hxSerialize` / `hxUnserialize`, and structural resolver callbacks | 0 serializer-specific compiler declarations; 0 unsafe runtime accesses | Semantic-diff + generated-output + direct runtime + compiler-debt ratchet | Retired into shared typed metadata | Serializer and Unserializer call staged Reflect. The generic Reflect field/method adapters provide same-package access; Type empty-instance helpers allocate inherited carriers and repair virtual dispatch. No duplicate serializer registry or bridge remains. | `haxe_go-vfp.10.5.1` |
| `net_socket` | `sys.net.Host`, `sys.net.Socket`, `sys.net.UdpSocket`, TLS socket composition | 0 compiler declarations (group retired) | Semantic-diff + snapshot runtime + direct race/cross-build | Migrated (canonical staged std + typed runtime capabilities) | Staged source owns public objects, stream/error policy, address construction, select identity, TLS configuration, and accepted SSL identity. Footprint-explicit `socket.go`, build-tagged broadcast adapters, and `socket_ssl.go` own only DNS/OS transport, deadline/readiness, socket options, and TLS resources over typed handles. | `haxe_go-vfp.8.7.14` |
| `template_support` | runtime representation beneath `haxe.Template` | 0 Template-specific compiler helpers (group retired) | Semantic-diff + generated-output + direct runtime | Migrated (staged Haxe + generic generated metadata + typed `hxrt`) | Staged `haxe.Template` owns parsing, lookup, concrete and structural iteration, macros, errors, and rendering. The generic `Reflect.field` / `Reflect.hasField` path can consume selective same-package generated-method metadata; this is not a Template helper. Footprint-explicit `runtime/hxrt/template.go` still owns only dynamic array inspection, object classification, and invocation through `std/hxrt/template/NativeTemplate.hx`. | bridge retirement `haxe_go-vfp.8.7.16`; concrete iteration `haxe_go-vfp.8.7.19` |

## Explicit Decision Records

These are the canonical per-surface decisions for shim ownership and alternatives.

| Record | Surface | Decision | Alternatives reviewed | Evidence |
| --- | --- | --- | --- | --- |
| `SDR-001` | `json` (`haxe.Json`, `haxe.format.Json*`) | Move API ownership to canonical staged std (`std/go/_std`) and keep behavior in `hxrt` (`JsonParse`/`JsonStringify`) | Compiler shim, direct lower-call rewrites, extern/runtime package | Snapshot parity + migration log (`haxe.go-7zy.10`, `haxe.go-cgk.5`) |
| `SDR-002` | `process` (`sys.io.Process`) | Keep the complete public API and stream policy in canonical staged Haxe; cross only typed opaque handles and native spawn/pipe/wait/signal/close capabilities through `std/hxrt/process` and `runtime/hxrt/process.go`. Compiler ownership is retired. | Keep the former combined Sys/Process group, keep the isolated Process emitter, expose the public stdlib as externs, use raw injection in the override, or move Haxe bounds/EOF/null policy into `hxrt` | Process semantic-diff contracts, direct runtime tests, selective `core/runtime_hxrt_infer_process`, sibling staged-source precedent, and the permanent 134-site `GoRaw` reduction under `haxe_go-vfp.8.7.7` |
| `SDR-003` | retired `io` group (`haxe.io.Bytes*`, stream helpers, encoding edges) | Keep the complete public hierarchy and algorithms in canonical staged Haxe. Cross only opaque `ByteView`, conversion/copy/UTF, and scalar IEEE-754 capabilities through typed `std/hxrt/io`; retain no compiler IO primitive or profile branch. | Preserve the whole group, move the blob unchanged into `hxrt`, canonical staged source plus opaque typed storage | Bytes/IO/RawNative/alias/endian/EOF semantic-diff, file/process subclass evidence, generated snapshots, strict sweep, crypto byte-view contract, and performance gates under `haxe_go-vfp.8.7.11`. |
| `SDR-004` | collections (`haxe.ds.*Map`, `List`, `Lambda`, stable sort helpers) | Keep public APIs and algorithms in staged or upstream Haxe. Use narrow typed native storage only for actual runtime facts, and admit exact call adapters or representation-only staged carriers only where erased Haxe generics are not assignable under Go. | Preserve generated map/list classes, extern-only containers, move algorithms into `hxrt`, retain compiler-owned loops, staged source plus typed storage/representation capabilities | Map/list, complete Lambda/Iterable carrier, serializer, sort, selective-runtime, and compiler-debt contracts; migrations `haxe_go-vfp.8.7.10`, `haxe_go-vfp.8.7.17`, and `haxe_go-vfp.8.7.18` |
| `SDR-005` | retired `http` (`sys.Http`) | Keep Haxe-visible request choreography in canonical staged source and native URL/HTTP/socket resources behind typed opaque `std/hxrt/http` handles. The compiler group and intermediate raw helper island are retired. | Keep choreography in compiler, move the generated declaration blob into `hxrt`, retain raw helper islands, expose public std as externs, staged source plus typed transport handles | HTTP callback/proxy/custom-request semantic contracts, four generated/runtime snapshots, `core/runtime_hxrt_infer_http`, direct transport/timeout/truncated-body/socket-cleanup tests, and compiler-debt reduction under `haxe_go-vfp.8.7.12` |
| `SDR-006` | retired `regex_serializer` and `serialization_source_bridge` | Keep the complete public algorithms in canonical staged Haxe. Cross only typed RE2 execution and bounded float parsing into feature-sliced `hxrt`; reuse staged Reflect plus shared generated field/method metadata for private access and hooks, and Type metadata for class/enum lookup and constructor-free allocation. | Keep the mixed compiler emitter, retain `reflect.NewAt`/`unsafe.Pointer`, generate serializer-specific duplicate metadata tables, preserve a separate invocation bridge, or export every generated member | Regex/serializer semantic-diff contracts, selective-runtime snapshots, direct runtime tests, intrinsic registry, compiler-debt reduction under `haxe_go-vfp.8.7.13`, and the three-level inherited typed-accessor contract under `haxe_go-vfp.10.5.1` |
| `SDR-007` | retired `net_socket` group (`sys.net.Host`, `sys.net.Socket`, `sys.net.UdpSocket`, `sys.ssl.Socket` composition) | Keep the complete public APIs and Haxe-facing policy in canonical staged source. Cross only typed opaque socket/certificate/key/SNI handles and concrete result carriers into footprint-explicit `socket.go`, `ssl.go`, and `socket_ssl.go`; compiler ownership is retired. | Keep compiler classes, retain raw injection or `Dynamic` native handles, extern-only wrappers, move public objects/exceptions into `hxrt`, or use staged source plus typed handles | Host/TCP semantic-diff, UDP/TLS/SNI runtime snapshots, selective-runtime positive/negative evidence, direct close/timeout/readiness/race tests, sibling staged-source precedent, and the permanent compiler-debt reduction under `haxe_go-vfp.8.7.14` |
| `SDR-008` | retired `stdlib_symbols` library surfaces | Retire the group completely. Emit upstream `haxe.ds.Option` through the ordinary enum pipeline, remove the empty `Std` nominal carrier and false selectors, and keep every legitimate metadata/representation capability independently named. Its two residual behavior calls were explicitly tracked until `SDR-021` moved them to staged source. | Preserve the mixed blob, invent a special Option carrier, keep an empty compatibility group, stage an incomplete `Std`, or attach serialization to Type/Reflect | Sibling-target staged-source precedent; option snapshot + portable parity; bidirectional registry/provenance/debt contracts; retirement `haxe_go-vfp.8.7.15.7`; source closure `haxe_go-vfp.8.7.22` |
| `SDR-009` | `haxe.Constraints` + `haxe.Rest` direct abstraction surfaces | Split ownership: staged std for `haxe.Constraints` typing/native bridge metadata, compiler lowering for native-slice `haxe.Rest` operations | Keep as compile-only debt, full compiler shim blobs, staged std abstract emission | Direct parity contracts (`haxe_constraints_contract`, `haxe_rest_contract`) plus snapshot `stdlib/haxe_constraints_rest_direct` |
| `SDR-010` | stack/main-loop `haxe.misc` surfaces (`haxe.CallStack`, `haxe.NativeStackTrace`, `haxe.EntryPoint`, `haxe.MainLoop`, `haxe.Timer`) | Split the tranche honestly: keep deterministic stack fallbacks in staged std under target-sensitive snapshot coverage, and support direct event-loop APIs through staged std wrappers over `sys.thread.EventLoop` instead of compiler-owned shims. Native Go stack capture is available only as the explicit `reflaxe_go_native_stack_trace` diagnostic capability, not as portable semantic-diff parity. | Leave as generic compile-only debt, force semantic-diff on target-timed event scheduling prematurely, grow compiler-owned stack/event-loop shims without a runtime ownership decision | Snapshot contracts `stdlib/haxe_stack_loop_target_sensitive`, `stdlib/haxe_native_stack_trace_opt_in`, and `stdlib/haxe_main_loop_runtime_direct`; staged std overrides `std/go/_std/haxe/CallStack.hx`, `std/go/_std/haxe/NativeStackTrace.hx`, `std/go/_std/haxe/EntryPoint.hx`, `std/go/_std/haxe/MainLoop.hx`, `std/go/_std/haxe/Timer.hx`; native stack spike `docs/spikes/native-stack-capture-contract.md`; cross-target precedent from `haxe.rust/std/haxe/CallStack.cross.hx` |
| `SDR-011` | legacy text surfaces (`haxe.Utf8`, `haxe.Ucs2`) | Use staged std for deprecated `haxe.Utf8` helper semantics; keep `haxe.Ucs2` as explicit target-sensitive platform exclusion under snapshot coverage | Grow compiler-owned text shims, leave both as anonymous compile-only debt | `haxe_utf8_contract`, `stdlib/haxe_utf8_basic`, `stdlib/haxe_ucs2_platform_exclusion`, staged std override `std/go/_std/haxe/Utf8.hx` |
| `SDR-012` | framework-owned raw-injection helper islands (`@:goAllowRaw` + `reflaxe.go.macros.GoInjection.__go__`) | Use this as the preferred middle layer when helper logic needs same-package generated-type access but does not need compiler-context decisions. Keep imports typed; do not use raw injection as a substitute for extern metadata. | Grow compiler-owned `GoRaw` blocks, extern-only wrappers, app-side raw `__go__` | `haxe.go-14as.48`, `haxe.go-14as.50`, sibling precedent from `reflaxe.rust` (`RustInjection.__rust__`) and `haxe.ocaml` (`__ocaml__` escape-hatch policy) |
| `SDR-013` | direct `haxe.rtti.*` support (`CType`, `Meta`, `Rtti`, `XmlParser`) | Keep public RTTI/parser logic in staged std, with a narrow compiler-owned metadata/lowering contract underneath | Move RTTI into compiler-owned stdlib authorities, invent RTTI-specific Go carrier structs, or classify the whole family unsupported | `haxe_rtti_direct_contract`, `stdlib/haxe_rtti_direct`, `haxe.go-14as.57`, `haxe.go-14as.59` |
| `SDR-014` | portable exception carrier and `sys.thread` process-lifecycle boundary | Keep the public API in staged std, require portable runtime validation failures to use `Throw`, and own worker identity/TLS state, recovery, and reporting in `hxrt`. The compiler adds a feature-gated generated-main foreground drain and, only when `sys.thread` is reachable, a cleanup scope around detached `go.Go.spawn` callbacks. Recover only explicit Haxe exception carriers; the detached scope does not join or recover native panics. | Treat every panic as a Haxe value, use raw panic for portable validation failures, retain TLS in per-instance global maps, let uncaught Haxe throws crash the Go process, join all native goroutines, or make portable workers daemon-like | `core/portable_runtime_failure_haxe_catch`, `semantic_diff/sys_thread_primitives_contract`, `stdlib/sys_thread_uncaught_exception`, `go_native/native_panic_not_haxe_catch`, `go_native/goroutine_native_panic`, `go_native/goroutine_native_shutdown`, direct `runtime/hxrt` lifecycle/race tests, and Haxe 4.3.7 interpreter probe recorded in `haxe_go-vfp.10.1` |
| `SDR-015` | `sys.FileSystem` | Keep the complete public API and `FileStat` construction in canonical staged Haxe; expose only typed native filesystem capabilities through `std/hxrt/fs` and `runtime/hxrt/filesystem.go`. Compiler ownership is retired, and runtime slicing excludes filesystem support from unrelated `sys.*` programs. | Keep the multi-function `GoCompiler` shim group, duplicate the upstream `FileStat` typedef, put the helpers in the broad `sys.go` runtime slice, or use raw `__go__` in the override | Complete Haxe 4.3.7 surface, snapshot + semantic-diff filesystem contracts, direct runtime tests, selective-runtime evidence, and compiler-debt reduction under `haxe_go-vfp.8.7.4` |
| `SDR-016` | `sys.io.File`, `FileInput`, `FileOutput`, `FileSeek` | Keep the complete public API and stream semantics in canonical staged Haxe; cross only typed opaque handles and native file capabilities through `std/hxrt/fs` and `runtime/hxrt/file.go`. Retain exact public base-I/O fields at macro initialization solely so indirectly staged subclasses can type-check after DCE. | Keep File declarations and algorithms in the former combined `lowerSysStdlibShimDecls`, expose generated `Bytes` internals to `hxrt`, put file support in broad `sys.go`, or use raw injection in the override | Binary/text/copy/write/append/update/seek/tell/bounds/EOF semantic-diff coverage, direct snapshots/runtime tests, selective-runtime evidence, and permanent `GoRaw` reductions under `haxe_go-vfp.8.7.5` |
| `SDR-017` | root `Sys` | Keep the complete supported public API in canonical `std/go/_std/Sys.hx`; use narrow typed console, process-state, environment, path, clock, command, and standard-file capabilities. Inline one-step source methods for Go-shaped direct calls, while first-class references materialize the same source-owned contract. | Retain generated root declarations in the combined shim, use externs as the public stdlib API, put map/alias/fallback behavior in `hxrt`, or substitute wall time for `cpuTime` | Root semantic-diff and negative contracts, direct runtime tests, separate Sys-only/Process-only selective snapshots, sibling `haxe.rust` staged-source precedent, and compiler-debt reduction under `haxe_go-vfp.8.7.6` |
| `SDR-018` | `Sys.getChar` terminal behavior | Keep EOF construction and requested echo in staged `Sys`; expose only serialized terminal-state control and one-byte native input through typed `NativeTerminal` and a footprint-explicit, build-tagged `terminal` runtime slice. Confine the POSIX `unsafe.Pointer` to one typed ioctl helper with exact debt ceilings. | Reintroduce a compiler shim, use ordinary line-buffered `FileInput.readByte`, add raw Haxe injection, depend on a current `x/term` that raises the generated Go floor, or pin the older compatible release with a known advisory | Real-PTY no-newline/echo/restoration behavior executed by a `checkptr=2`-instrumented binary, redirected EOF, Linux/macOS/Windows implementation cross-builds, an unsupported-FreeBSD cross-build, selective snapshot evidence, supported-toolchain race/checkptr gates, and the exact unsafe ratchet closed under `haxe_go-vfp.10.5` |
| `SDR-019` | portable compiler stdlib intrinsic registry | Inventory exact Haxe symbols, compiler entry points, selectors, direct call rewrites, dependencies, evidence, review conditions, and migration beads in a fail-closed machine-readable registry. Classify behavior-heavy groups as avoidable migration debt; distinguish narrow admitted intrinsics from explicit `go.*` native emitters. | Keep the file ledger as an indirect audit, treat every named shim as required, or ban all compiler primitives without recognizing type-metadata/representation needs | `docs/compiler-stdlib-intrinsics.json`, `test/test_compiler_stdlib_intrinsic_registry.py`, and the split compiler-debt exceptions under `haxe_go-vfp.8.7.8` |
| `SDR-020` | lowercase generated-method discovery through `Reflect.field` / `Reflect.hasField` | Retain one selective, method-only compiler metadata plan over the final reachable generated class graph. Emit a canonical-receiver switch and per-class own-method resolvers with one generated-superclass fallback; return only an already-bound method value. Keep lookup policy in `Reflect`, iteration policy in staged `haxe.Template`, and invocation only in the existing typed runtime boundary. | Export or duplicate generated methods, use a provider interface, add a global registry/map, use unsafe reflection in `hxrt`, restore a Template-specific compiler shim, or introduce a full semantic program IR for this bounded closed-world fact | `haxe_template_concrete_iterable_contract`, `stdlib/haxe_template_generated_method_lookup`, direct bound-method runtime coverage, the intrinsic registry, and the written xhigh second pass under `haxe_go-vfp.8.7.19` |
| `SDR-021` | complete `Std` and `haxe.Log` source ownership | Keep all Haxe-visible parsing, overflow, downcast, formatting, position, custom-parameter, trace-rebinding, and output policy in canonical staged source. Retain only `Std.string` / `Std.isOfType` as exact representation intrinsics and use typed native bindings for float conversion, truncation, and random generation. Extend general typed Go class/cast lowering for mutable static dynamic functions, catchable null calls, concrete downcasts, and virtual-carrier runtime tests; do not add a Log/Std-specific semantic product or universal IR. | Preserve direct `Std.parseInt` / `haxe.Log.trace` rewrites, move library policy into `hxrt`, publish a partial `Std`, treat nil Go-function panics as Haxe catches globally, add profile-specific lowering, or introduce a whole-program IR for four bounded representation seams | `std_complete_api_contract`, `stdlib/std_log_source_owned`, prior Std type-test contracts, provenance/intrinsic/debt gates, sibling staged-source precedent, and `haxe_go-vfp.8.7.22` |
| `SDR-022` | stdlib ownership closeout and capability dispatch | Retire the obsolete `lowerStdlibShimDecls` migration-debt identity. Keep one fail-closed `lowerRegisteredCompilerCapabilityDecls` selector for the three approved portable metadata/representation groups and three explicit `go.*` native groups, while measuring every selected declaration emitter independently. Remove the now-empty portable stdlib migration-debt exception. | Keep reporting a completed migration as avoidable debt, inline the same six checks at the generation call site, or misclassify structural selection as another approved intrinsic | Exact bidirectional and structural dispatcher/registry contract, removal of one plumbing function from the named compiler-shim metric while all nine declaration emitters remain ratcheted, zero `migration_required` registry entries, unchanged generated snapshots, and parent closeout review under `haxe_go-vfp.8.7` |

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
| `haxe.io.Bytes` algorithms and RawNative/cache-coupled conversion | Canonical staged `Bytes` over typed `std/hxrt/io` | Keep this split | Source owns algorithms, bounds, aliases, invalidation, and encoding policy; runtime owns only opaque native byte and UTF/copy capabilities. No compiler wrapper remains. | Completed by `haxe_go-vfp.8.7.11` |
| `haxe.Resource` embedded payload table (`content`) | `GoCompiler` resource-table literal + source-owned std methods | Keep compiler-owned data population | The std `haxe.Resource` methods lower normally, but the actual payloads come from compiler resources (`Context.getResources()` / `__resources__()`), not reusable Haxe source. The backend must materialize `haxe__Resource_content`; otherwise direct resource calls compile but runtime sees an empty table. | `haxe.go-14as.30` |
| `haxe.io.Input` / `haxe.io.Output` helper loops (`readAll`, `readLine`, `readUntil`, `write`, `writeInput`, `readFullBytes`, `writeString`) | Canonical staged base classes | Keep this split | Ordinary source inheritance and `__hx_this` dispatch now carry all algorithms; `GoIoHelpers` and generated IO-specific wrappers are retired. | Completed by `haxe_go-vfp.8.7.11` |
| retired `sys.Http` compiler/helper split | Canonical `std/go/_std/sys/Http.hx` over typed `std/hxrt/http` and `runtime/hxrt/http.go` | Keep this split | Source owns request/callback/header/status policy; the runtime owns only Go URL parsing, transport resources, response bytes/headers, proxy configuration, and typed socket consumption. | Completed by `haxe_go-vfp.8.7.12` |
| retired `regex_serializer` (`EReg`, serializer/unserializer stack) | Canonical staged source over typed `std/hxrt/regex`, float-only `std/hxrt/serialization`, staged Reflect with shared generated metadata, and Type metadata | Keep this split | Regex execution and float parsing are runtime facts; token streams, match state, resolver policy, traversal, caches, and field policy are source behavior. Shared typed Reflect/Type metadata supplies generated-program facts without unsafe access or a serializer-specific registry. | Completed by `haxe_go-vfp.8.7.13`; unsafe/bridge retirement `haxe_go-vfp.10.5.1` |
| `net_socket` (`sys.net.Host`, `sys.net.Socket`, `sys.net.UdpSocket`) | Canonical staged source over typed `std/hxrt/net` and `std/hxrt/ssl` bindings; the former compiler group and `GoNetSocketEmitter` are retired | Keep this split; widen only typed runtime capabilities when an OS resource is genuinely required | Target-sensitive OS behavior belongs in `hxrt`; public objects, exceptions, byte copying, select identity, and TLS configuration remain ordinary Haxe source. | Completed by `haxe_go-vfp.8.7.14` |

## Ownership Boundary (Post `haxe_go-vfp.8.7.7`)

- `runtime/hxrt/sys.go`, `runtime/hxrt/file.go`, and `runtime/hxrt/process.go` own distinct native capabilities: root process state, file/standard-stream handles, and child processes respectively.
- `std/go/_std/sys/io` owns the File and Process public APIs and stream semantics, and `std/go/_std/Sys.hx` owns root `Sys`. None of these families has compiler-generated declarations or semantic branches.
- `std/hxrt/process` exposes only typed opaque process/pipe handles plus native capability calls. OS child-process behavior belongs in `runtime/hxrt/process.go`; Haxe-visible bounds, EOF, nullable status, detached rejection, and lifecycle behavior remain reviewable staged source.
- Staged `Sys.putEnv` explicitly selects the non-throwing `SysSetEnvironment` capability because Haxe 4.3.7 eval exposes `Void`; `hxrt.SysPutEnv` retains its native error for Go-native callers.
- Simple staged root methods inline to typed capabilities so direct output remains Go-shaped. First-class function references materialize source-owned functions instead of restoring compiler special cases.
- `Sys.cpuTime` is the explicit compile-context exception: it is rejected at Haxe compile time because no Go-standard-library implementation can satisfy the CPU-time contract. See [Portable root `Sys` contract](portable-sys-contract.md).

## Historical Measured Tradeoff: Shim vs Simpler Path

The following baseline measured `haxe.crypto.Base64.encode` before its compiler
shim was retired by `haxe_go-vfp.8.7.15.1`.

The same harness now measures the replacement staged-Haxe + typed-runtime path:

```bash
npm run test:perf:stdlib-shims
```

Artifacts:

- `.cache/perf-stdlib-shim-review/report.json`
- `.cache/perf-stdlib-shim-review/report.md`

Those artifacts describe the current staged boundary. The fixed table below is
the historical compiler-shim baseline retained for comparison.

Measured at `2026-02-24T01:22:42Z` on `darwin/arm64` (`Apple M2 Pro`):

| Path | ns/op | B/op | allocs/op | Code-shape LOC (call path) |
| --- | ---: | ---: | ---: | ---: |
| Generated shim (`haxe__crypto__Base64_encode` + bytes conversion helpers) | 64.05 | 112 | 3 | 33 |
| Direct Go (`base64.StdEncoding.EncodeToString`) | 57.09 | 96 | 2 | 3 |
| Delta | +12.19% | +16 | +1 | +30 |

Interpretation:

- overhead is primarily representation conversion (`[]int` <-> `[]byte`) rather than base64 algorithm cost
- this remains a regression baseline for the staged Base64 API over the typed
  `NativeCrypto` boundary; it no longer describes an approved or active compiler
  implementation

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
