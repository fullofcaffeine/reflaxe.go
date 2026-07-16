# Portable Module Mapping Contract (Tier1 Seed)

This document is the module-level companion to `docs/ownership-rubric.md`.

Use the rubric document to decide which layer should own behavior.
Use this document to record the actual ownership split for concrete portable modules.

This document defines ownership mapping for Tier1 portable modules:

- Haxe-source implementation
- runtime binding (`hxrt`)
- compiler intrinsic/shim
- compiler compatibility migration debt
- mixed ownership (explicitly split)

It is the canonical Tier1 module-mapping seed for family extraction work.

Contract inputs:

- `test/portable_allowlist.json` (Tier1 module set)
- `test/portable_conformance_tier1.json` (Tier1 module->conformance cases)
- `docs/portable-semantics-v1.md` (semantic rules)

## Ownership Class Definitions

1. `haxe_source`
   - Behavior lives in Haxe std sources (`std/go/_std` overrides or the upstream std implementation).
2. `runtime_binding`
   - Haxe surface delegates behavior to runtime package functions in `runtime/hxrt/*.go`.
3. `compiler_intrinsic`
   - An exact operation is emitted directly because it needs closed compiler
     metadata or a backend representation fact. Admission is defined by
     `docs/compiler-stdlib-intrinsics.json`.
4. `mixed`
   - Surface spans more than one class above; split is explicit and test-gated.
5. `compiler_migration`
   - The current implementation is compiler-emitted for compatibility, but its
     behavior belongs in staged source or `hxrt` and has a concrete migration
     bead. This is current location, not approved architecture.

## Tier1 Mapping Table

| Module | Ownership class | Primary implementation location | Runtime dependency | Tier1 conformance cases |
| --- | --- | --- | --- | --- |
| `Math` | `compiler_migration` | Current compatibility implementation: `lowerStdlibSymbolShimDecls`; target owner: staged source plus narrow runtime math capabilities | Indirect via core helpers where needed | `numeric_edge_cases`, `stringtools_math`; migration `haxe_go-vfp.8.7.15` |
| `Std` | `mixed` (exact `compiler_intrinsic` primitives + `compiler_migration` + runtime helpers) | `Std.isOfType` and representation conversion are individually registered; the broad `stdlib_symbols` block is migration debt | `runtime/hxrt/string.go`, `runtime/hxrt/exception.go`, core helpers | `exception_api_contract`, `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract`, `typed_nil_dynamic_string_contract`; migration `haxe_go-vfp.8.7.15` |
| `DateTools` | `haxe_source` | `std/go/_std/DateTools.hx` | None beyond core `Date` and string primitives already owned elsewhere | `stringbuf_datetools_lambda_contract`, `datetools_cross_std_contract` |
| `StringTools` | `haxe_source` | `std/go/_std/StringTools.hx` | None beyond core string/runtime primitives used by normal lowering | `stringtools_math`, `stringtools_cross_std_contract` |
| `Lambda` | `mixed` (`haxe_source` + exact representation intrinsics) | `std/go/_std/Lambda.hx` owns all 19 public algorithms: traversal, comparison, early exit, allocation, filtering, mapping, folding, lookup, flattening, and concatenation. Its private representation-only `LambdaGoIterableCarrier` companion delegates `iterator()` without owning an algorithm. | `GoLambdaIterableLowering.hx` plus `lowerLambdaSourceCallAdapter` wrap only Go-invariant iterable, callback, constrained nested-iterable, and erased-result shapes. Already-erased nested carriers fail with explicit diagnostics instead of reaching `go test`. | `lambda_full_api_contract`, `lambda_generic_iterable_count_empty_contract`, `lambda_iter_array_list_contract`, `lambda_iter_generic_iterable_contract`, `lambda_list_contract`, `negative/lambda_flatten_erased_nested`, `negative/lambda_flat_map_erased_result`; source migration `haxe_go-vfp.8.7.17`, full API/carrier closure `haxe_go-vfp.8.7.18` |
| `haxe.Template` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/Template.hx` owns parsing, lookup/fallback, iteration, macro argument construction, errors, and rendering | `std/hxrt/template/NativeTemplate.hx` maps three dynamic representation operations to footprint-explicit `runtime/hxrt/template.go`; no Template helper is compiler-generated | `haxe_template_contract`, `stdlib/haxe_template_basic`, direct `template_test.go`; compiler bridge retirement `haxe_go-vfp.8.7.16` |
| `Sys` | `mixed` (`haxe_source` + `runtime_binding`, with one compile-time unsupported diagnostic) | `std/go/_std/Sys.hx` with typed bindings in `std/hxrt/sys` and `std/hxrt/fs` | `runtime/hxrt/sys.go` provides root process capabilities, `runtime/hxrt/file.go` provides non-owning standard-stream handles, and footprint-explicit build-tagged `runtime/hxrt/terminal*.go` provides terminal state plus one-byte input. Staged Haxe retains `getChar` EOF/echo semantics. The complete root surface is specified in [Portable root `Sys` contract](portable-sys-contract.md). Portable `putEnv` discards the retained native error for Haxe 4.3.7 eval parity; `setTimeLocale` reports `false`; `cpuTime` fails at Haxe compile time. | `root_sys_contract`, `root_sys_portable_contract`, `sys_io_roundtrip`, `file_error_semantics_contract`, `sys_sleep_contract`, `sys/root_sys_portable`, `sys/sys_get_char_terminal`, `test_sys_get_char_terminal.py`, `core/runtime_hxrt_infer_sys`, `negative/sys_cpu_time_unsupported` |
| `haxe.Utf8` | `haxe_source` | `std/go/_std/haxe/Utf8.hx` | None beyond `haxe.io.Bytes`, `UnicodeString`, and shared Go string runtime helpers already owned elsewhere | `haxe_utf8_contract`, `stdlib/haxe_utf8_basic` |
| `haxe.Json` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/Json.hx` | `runtime/hxrt/json.go` | `json_parse_stringify_contract` |
| `haxe.crypto.Base64`, `Md5`, `Sha1`, `Sha224`, `Sha256` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/crypto` owns public APIs, Base64 alphabets and padding, and `Bytes` conversion | `std/hxrt/crypto/NativeCrypto.hx` maps strings and integer byte arrays to footprint-explicit `runtime/hxrt/crypto.go`; no generated `haxe.io.Bytes` layout crosses the boundary | `crypto_source_owned`, `crypto_xml_zip`, `stdlib/crypto_xml_zip_basic`, direct `crypto_test.go`; migration `haxe_go-vfp.8.7.15.1` |
| root `Xml`, `haxe.xml.Parser`, `haxe.xml.Printer` | `haxe_source` | `std/go/_std/Xml.hx` owns the DOM with narrow Go-representation adaptations; unchanged upstream Haxe 4.3.7 source owns the parser state machine; `std/go/_std/haxe/xml/Printer.hx` preserves upstream formatting without an incidental regex dependency | No native XML runtime or compiler declarations; ordinary staged `StringMap`, `StringTools`, and `StringBuf` dependencies only | `xml_source_owned`, `root_xml_contract`, `crypto_xml_zip`, `stdlib/xml_root_dom_basic`, `stdlib/crypto_xml_zip_basic`; migration `haxe_go-vfp.8.7.15.2` |
| `haxe.ds.EnumValueMap` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/ds/EnumValueMap.hx` owns recursive comparison, AVL balancing, iteration, copying, and the public API | `std/hxrt/collections/NativeEnumValue.hx` and `runtime/hxrt/enum_value.go` provide only the generated-enum carrier predicate | `ds_maps_list_contract`; source migration `haxe_go-vfp.8.7.10` |
| `haxe.ds.IntMap` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/ds/IntMap.hx` owns the public API, iteration, copying, rendering, and clearing | Typed `std/hxrt/collections` bindings over selectively copied `runtime/hxrt/map_int.go` provide native storage and deterministic key snapshots | `ds_maps_list_contract`, `core/runtime_hxrt_infer_map_int`; source migration `haxe_go-vfp.8.7.10` |
| `haxe.ds.ObjectMap` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/ds/ObjectMap.hx` owns the public API, iteration, copying, rendering, and clearing | Typed `std/hxrt/collections` bindings over selectively copied `runtime/hxrt/map_object.go` provide retained reference-identity storage and deterministic key snapshots | `ds_maps_list_contract`, `core/runtime_hxrt_infer_map_object`; source migration `haxe_go-vfp.8.7.10` |
| `haxe.ds.StringMap` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/ds/StringMap.hx` owns the public API, iteration, copying, rendering, and clearing | Typed `std/hxrt/collections` bindings over selectively copied `runtime/hxrt/map_string.go` provide native storage and deterministic key snapshots | `ds_maps_list_contract`, `core/runtime_hxrt_infer_map_string`; source migration `haxe_go-vfp.8.7.10` |
| `haxe.ds.List` | `haxe_source` | `std/go/_std/haxe/ds/List.hx`; Lambda and serializer integration now use only its public iterator/API instead of reading its private carrier | None beyond ordinary array/string primitives | `ds_maps_list_contract`, `list_std_contract`, `lambda_list_contract`; source migration `haxe_go-vfp.8.7.10`, adapter cleanup `haxe_go-vfp.8.7.17` |
| `haxe.io.Bytes` | `mixed` (`compiler_migration` + runtime helper calls) | Current compatibility implementation: `lowerIoStdlibShimDecls`; target owner: staged source over typed raw-byte storage | `runtime/hxrt/bytes.go`, `runtime/hxrt/string.go` | Existing bytes/IO contracts; migration `haxe_go-vfp.8.7.11` |
| `haxe.io.Path` | `haxe_source` | upstream Haxe stdlib `haxe/io/Path.hx` | Core string and array helpers lowered by the target (`lastIndexOf`, `split`, `Array.join`, `String.fromCharCode`) | `option_date_path`, `path_cross_std_contract` |
| `sys.FileSystem` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/sys/FileSystem.hx` with typed carriers in `std/hxrt/fs` | `runtime/hxrt/filesystem.go`, `runtime/hxrt/string.go`, `runtime/hxrt/exception.go` | `filesystem_contract`, `sys/filesystem_basic_smoke` |
| `sys.io.File` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/sys/io/File.hx`, `FileInput.hx`, `FileOutput.hx`, and `FileSeek.hx`, with typed carriers in `std/hxrt/fs` | selectively copied `runtime/hxrt/file.go`; native OS failures remain Haxe exceptions, while byte conversion, bounds/EOF semantics, seek mapping, and public stream construction stay in Haxe source | `file_read_write_contract`, `file_error_semantics_contract`, `semantic_diff/sys_db_io_contract`, `sys/file_error_semantics`, `stdlib/sys_db_io_direct` |
| `sys.io.Process` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/sys/io/Process.hx` with typed opaque carriers in `std/hxrt/process` | selectively copied `runtime/hxrt/process.go`; native spawn, pipes, waits, signals, and close stay in Go, while stream construction, bounds/EOF translation, detached rejection, `Null<Int>` exit status, and public lifecycle policy stay in Haxe source | `process_echo_contract`, `process_error_semantics_contract`, `sys/process_error_semantics`, `core/runtime_hxrt_infer_process` |
| `sys.net.Socket` | `compiler_migration` | Current compatibility implementation: `GoNetSocketEmitter`; target owner: staged API over typed runtime socket handles | Uses core runtime helpers and the shared `haxe.io.Input` helper layer today | Existing TCP contracts; migration `haxe_go-vfp.8.7.14` |
| `sys.net.UdpSocket` | `compiler_migration` | Current compatibility implementation: `GoNetSocketEmitter`; target owner: staged API over typed runtime UDP handles | Uses core runtime helpers today | `stdlib/sys_net_udp_socket_direct`; migration `haxe_go-vfp.8.7.14` |

## Additional Mixed-Ownership Rows

These modules are outside the Tier1 seed table above, but they are important enough to record explicitly because
their ownership split is easy to misunderstand.

| Module family | Ownership class | Public implementation location | Backend-owned support beneath it | Evidence |
| --- | --- | --- | --- | --- |
| `haxe.io` misc direct surfaces (`BufferInput`, `BytesData`, `Encoding`, `Eof`, `Error`, `FPHelper`, `Mime`, `Scheme`, `StringInput`) | `mixed` with `compiler_migration` | upstream `std/haxe/io/**` plus `std/go/_std/haxe/io/FPHelper.hx` | current migration-debt base IO/encoding/error/input hierarchy plus the Bytes carrier; target ownership is staged source over a narrow typed representation boundary | `semantic_diff/haxe_io_misc_contract`, `stdlib/haxe_io_misc_direct`, `haxe_go-vfp.8.7.11` |
| `haxe.io` typed arrays (`ArrayBufferView`, `UInt8Array`, `UInt16Array`, `UInt32Array`, `Int32Array`, `Float32Array`, `Float64Array`) | `mixed` with `compiler_migration` | `std/go/_std/haxe/io/*.hx` | current migration-debt `haxe.io.Bytes` / `ArrayBufferViewImpl` carrier plus source-owned abstract behavior; float arrays already reuse staged `haxe.io.FPHelper` | `semantic_diff/haxe_io_typed_arrays_contract`, `stdlib/haxe_io_typed_arrays_direct`, `haxe_go-vfp.8.7.11` |
| `sys.db` direct surfaces (`Connection`, `ResultSet`, `Mysql`, `Sqlite`) | `mixed` | upstream `std/sys/db/**` interfaces and platform stubs | no fake DB runtime; Go keeps the upstream platform contract where `Mysql.connect` / `Sqlite.open` remain explicit unsupported runtime stubs instead of inventing target-owned behavior | `semantic_diff/sys_db_io_contract`, `stdlib/sys_db_io_direct` |
| `sys.io` direct handle surfaces (`FileInput`, `FileOutput`, `FileSeek`) | `mixed` (`FileSeek` is fully source-owned) | canonical staged modules under `std/go/_std/sys/io`; no File-specific compiler declarations or branches remain | typed opaque handles and native operations in `std/hxrt/fs` + `runtime/hxrt/file.go` beneath the source-owned public Haxe stream API | `semantic_diff/sys_db_io_contract`, `stdlib/sys_db_io_direct` |
| `sys.ssl` direct surfaces (`Certificate`, `Digest`, `DigestAlgorithm`, `Key`, `Socket`) | `mixed` | `std/go/_std/sys/ssl/*.hx` for the public API; `DigestAlgorithm` is fully source-owned | `runtime/hxrt/ssl.go` for certificate parsing, key parsing, digest/sign/verify, and TLS socket dial/listen/handshake/SNI selection helpers beneath the public wrappers | `stdlib/sys_ssl_leaf_direct`, `stdlib/sys_ssl_socket_direct`, `stdlib/sys_ssl_socket_sni_direct`, `semantic_diff/sys_net_address_ssl_digest_algorithm_contract`, `stdlib/sys_net_address_ssl_digest_algorithm_direct`; policy spike: `docs/spikes/ssl-udp-semantic-diff-spike.md` |
| `sys.thread` direct surfaces (`Condition`, `Deque`, `EventLoop`, `ElasticThreadPool`, `FixedThreadPool`, `IThreadPool`, `Lock`, `Mutex`, `NoEventLoopException`, `Semaphore`, `Thread`, `ThreadPoolException`, `Tls`) | `mixed` | `std/go/_std/sys/thread/*.hx` for the public API and pool policies; the compiler adds a feature-gated generated-main foreground drain and a lifecycle-only scope for detached `go.Go.spawn` callbacks; target-only worker companions remain staged support | `runtime/hxrt/thread.go` for blocking primitives, logical thread identity, lifecycle-owned TLS values, message queues, event-loop handles, nested foreground lifecycle, detached identity cleanup, and carrier-only uncaught-thread reporting beneath the staged wrappers | `semantic_diff/sys_thread_primitives_contract`, `semantic_diff/sys_thread_runtime_contract`, `stdlib/sys_thread_primitives_direct`, `stdlib/sys_thread_runtime_direct`, `stdlib/sys_thread_uncaught_exception`, `go_native/goroutine_native_panic` |
| `haxe.EntryPoint` / `haxe.MainLoop` / `haxe.Timer` direct event-loop surfaces | `mixed` | `std/go/_std/haxe/EntryPoint.hx`, `std/go/_std/haxe/MainLoop.hx`, `std/go/_std/haxe/Timer.hx` for the public API | `runtime/hxrt/thread.go` through `sys.thread.EventLoop` for main-thread callback queues, worker promises, repeating timers, and monotonic timer stamps | `stdlib/haxe_main_loop_runtime_direct`; policy spike: `docs/spikes/event-loop-semantic-diff-spike.md` |
| `haxe.rtti.*` (`CType`, `Meta`, `Rtti`, `XmlParser`) | `mixed` | `std/go/_std/haxe/rtti/*.hx` | class-token `__meta__` / `__rtti` lookup contract plus anonymous-record array-field mutation lowering in the compiler | `semantic_diff/haxe_rtti_direct_contract`, `stdlib/haxe_rtti_direct` |
| `haxe.ds.ArraySort` and `haxe.ds.ListSort` | `mixed` (upstream `haxe_source` + exact representation intrinsics) | Haxe 4.3.7 upstream stable merge-sort algorithms remain authoritative | `lowerDsSortHelperCall` only boxes/copies typed slices, adapts comparators, and restores typed linked-node heads across erased Go generic signatures | `haxe_ds_sort_helpers_contract`, `stdlib/haxe_ds_sort_helpers_direct`; adapter admission `haxe_go-vfp.8.7.17` |

## Notes on Staged Source Selection

Source-checkout builds put staged portable overrides on the initial classpath
before any macro is typed through:

- `haxe_libraries/reflaxe.go.hxml`

`CompilerBootstrap` does not own override precedence; installed packages use
the corresponding flattened `src/**/*.cross.hx` artifacts.

The canonical override inventory and its compiler-shim splits are locked by
`docs/stdlib-provenance-ledger.json`; Tier1 rows above record the corresponding
public ownership view.

## Transition Notes (Post-`__go__` Audit)

- `haxe.io.Bytes`
  - The current `mixed` classification describes implementation location today;
    it is not an approval of the compiler-owned API.
  - The first post-`__go__` extraction already moved pure hex and `BytesBuffer` leaf helpers into `runtime/hxrt/bytes.go`, leaving thin compiler wrappers in place.
  - The remaining compiler-emitted RawNative/cache-coupled string path (`ofString`, `getString`, UTF16/raw-native conversion helpers) is migration debt under `haxe_go-vfp.8.7.11`. Its `__hx_raw` cache validity and encoding-tag behavior are acceptance constraints for the replacement, not reasons to keep public library behavior in the compiler.
  - The ownership lock is `stdlib/bytes_raw_native_compiler_ownership`, which proves RawNative `Bytes.set(...)` still needs to invalidate the cached raw-byte view seen by downstream consumers such as Base64.
  - Closed evidence: `haxe.go-14as.51`, `haxe.go-14as.54`
- `haxe.io.Input` / `haxe.io.Output`
  - These surfaces are not listed as separate Tier1 rows here, but their inherited helper loops no longer live as raw loop bodies in `GoCompiler`.
  - `readAll`, `readLine`, `readUntil`, `readFullBytes`, `write`, `writeFullBytes`, `writeInput`, and `writeString` now route through `std/haxe/io/GoIoHelpers.hx`; the remaining compiler wrappers and base IO types are migration debt under `haxe_go-vfp.8.7.11`.
  - Closed evidence: `haxe.go-14as.52`
- `haxe.io` misc direct tranche
  - `haxe.io.FPHelper` is now the model staged-std slice for this family: public bit-conversion behavior lives in `std/go/_std/haxe/io/FPHelper.hx` on top of the existing little-endian `BytesInput` / `BytesOutput` contract.
  - `haxe.io.Mime` and `haxe.io.Scheme` remain plain upstream source-owned string abstracts.
  - `haxe.io.StringInput`, `haxe.io.BufferInput`, `haxe.io.Encoding`, `haxe.io.Eof`, and `haxe.io.Error` currently ride on the compiler-emitted compatibility hierarchy. `haxe_go-vfp.8.7.11` must move their public behavior to staged source and isolate any true representation primitive.
  - Baseline evidence: `haxe.go-14as.15`; migration owner: `haxe_go-vfp.8.7.11`
- `sys.Http`
  - Tier1 mapping records current compiler ownership as migration debt. A single
    request/callback contract requires coherent regression evidence, not compiler
    ownership.
  - `getResponseHeaderValues` and payload capture already live in `std/sys/GoHttpHelpers.hx`. The remaining request sequencing and proxy URL construction are migration debt under `haxe_go-vfp.8.7.12`, with the existing parity suite defining the replacement contract.
  - Baseline evidence: `haxe.go-14as.53`; migration owner: `haxe_go-vfp.8.7.12`

## Governance Rule

Any ownership change for a Tier1 module must update all of:

1. this mapping document,
2. `docs/ownership-rubric.md` when the rule itself changes,
3. `test/portable_conformance_tier1.json`,
4. `docs/stdlib-provenance-ledger.json` (when staged source files are added/changed),
5. relevant conformance fixtures in `test/semantic_diff`.
