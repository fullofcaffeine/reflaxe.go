# Portable Module Mapping Contract (Tier1 Seed)

This document is the module-level companion to `docs/ownership-rubric.md`.

Use the rubric document to decide which layer should own behavior.
Use this document to record the actual ownership split for concrete portable modules.

This document defines ownership mapping for Tier1 portable modules:

- Haxe-source implementation
- runtime binding (`hxrt`)
- compiler intrinsic/shim
- mixed ownership (explicitly split)

It is the canonical Tier1 module-mapping seed for family extraction work.

Contract inputs:

- `test/portable_allowlist.json` (Tier1 module set)
- `test/portable_conformance_tier1.json` (Tier1 module->conformance cases)
- `docs/portable-semantics-v1.md` (semantic rules)

## Ownership Class Definitions

1. `haxe_source`
   - Behavior lives in Haxe std sources (`std/_std` overrides or upstream std implementation).
2. `runtime_binding`
   - Haxe surface delegates behavior to runtime package functions in `runtime/hxrt/*.go`.
3. `compiler_intrinsic`
   - Behavior is emitted directly by compiler lowering/shim generation.
4. `mixed`
   - Surface spans more than one class above; split is explicit and test-gated.

## Tier1 Mapping Table

| Module | Ownership class | Primary implementation location | Runtime dependency | Tier1 conformance cases |
| --- | --- | --- | --- | --- |
| `Math` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerStdlibSymbolShimDecls`) | Indirect via core helpers where needed | `numeric_edge_cases`, `stringtools_math` |
| `Std` | `mixed` (`compiler_intrinsic` + runtime helper calls) | `src/reflaxe/go/GoCompiler.hx` (`lowerStdlibSymbolShimDecls`, core lowering paths) | `runtime/hxrt/string.go`, `runtime/hxrt/exception.go`, core helpers | `exception_api_contract`, `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract`, `typed_nil_dynamic_string_contract` |
| `DateTools` | `haxe_source` | `std/DateTools.cross.hx` | None beyond core `Date` and string primitives already owned elsewhere | `stringbuf_datetools_lambda_contract`, `datetools_cross_std_contract` |
| `StringTools` | `haxe_source` | `std/StringTools.cross.hx` | None beyond core string/runtime primitives used by normal lowering | `stringtools_math`, `stringtools_cross_std_contract` |
| `Sys` | `mixed` (`compiler wrapper` + `runtime_binding`) | `src/reflaxe/go/GoCompiler.hx` (`lowerSysStdlibShimDecls`) | `runtime/hxrt/sys.go` | `sys_io_roundtrip` |
| `haxe.Utf8` | `haxe_source` | `std/haxe/Utf8.cross.hx` | None beyond `haxe.io.Bytes`, `UnicodeString`, and shared Go string runtime helpers already owned elsewhere | `haxe_utf8_contract`, `stdlib/haxe_utf8_basic` |
| `haxe.Json` | `mixed` (`haxe_source` + `runtime_binding`) | `std/_std/haxe/Json.cross.hx` | `runtime/hxrt/json.go` | `json_parse_stringify_contract` |
| `haxe.ds.EnumValueMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.ds.IntMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.ds.ObjectMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.ds.StringMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.io.Bytes` | `mixed` (`compiler_intrinsic` + runtime helper calls) | `src/reflaxe/go/GoCompiler.hx` (`lowerIoStdlibShimDecls`) | `runtime/hxrt/bytes.go`, `runtime/hxrt/string.go` | `bytes_hex_contract`, `bytes_io_stream_contract`, `bytes_normalization_contract`, `bytes_of_data_contract`, `bytes_ops_contract`, `io_encoding_contract` |
| `haxe.io.Path` | `haxe_source` | `std/haxe/io/Path.cross.hx` | None beyond core string primitives already lowered by the target | `option_date_path`, `path_cross_std_contract` |
| `sys.FileSystem` | `mixed` (`compiler_intrinsic` + runtime helper calls) | `src/reflaxe/go/GoCompiler.hx` (`lowerFileSystemShimDecls`) | `runtime/hxrt/string.go`, `runtime/hxrt/exception.go` | `filesystem_contract` |
| `sys.io.File` | `runtime_binding` | `src/reflaxe/go/GoCompiler.hx` (`lowerSysStdlibShimDecls` forwarding wrappers) | `runtime/hxrt/sys.go` (`FileSaveContent`, `FileGetContent`, `FileGetBytes`, `FileSaveBytes`, `FileCopy`, `OpenFile*Output`, `OpenFileInput`) | `file_read_write_contract`, `semantic_diff/sys_db_io_contract`, `stdlib/sys_db_io_direct` |
| `sys.io.Process` | `runtime_binding` | `src/reflaxe/go/GoCompiler.hx` (`lowerSysStdlibShimDecls` forwarding wrappers) | `runtime/hxrt/process.go` (`NewProcess`, `ProcessOutput`) | `process_echo_contract` |
| `sys.net.Socket` | `compiler_intrinsic` | `src/reflaxe/go/compiler/emit/GoNetSocketEmitter.hx` (wired from `src/reflaxe/go/GoCompiler.hx`) | Uses core runtime helpers (`Throw`, string conversion) where needed | `socket_advanced_contract`, `socket_loopback_contract` |
| `sys.net.UdpSocket` | `compiler_intrinsic` | `src/reflaxe/go/compiler/emit/GoNetSocketEmitter.hx` (wired from `src/reflaxe/go/GoCompiler.hx`) | Uses core runtime helpers (`Throw`, string conversion) where needed | `stdlib/sys_net_udp_socket_direct` |

## Additional Mixed-Ownership Rows

These modules are outside the Tier1 seed table above, but they are important enough to record explicitly because
their ownership split is easy to misunderstand.

| Module family | Ownership class | Public implementation location | Backend-owned support beneath it | Evidence |
| --- | --- | --- | --- | --- |
| `haxe.io` misc direct surfaces (`BufferInput`, `BytesData`, `Encoding`, `Eof`, `Error`, `FPHelper`, `Mime`, `Scheme`, `StringInput`) | `mixed` | upstream `std/haxe/io/**` plus `std/haxe/io/FPHelper.cross.hx` | compiler-owned base IO/encoding/error/input hierarchy for `BufferInput` / `StringInput` / `Encoding` / `Eof` / `Error`, plus the `haxe.io.Bytes` carrier beneath `BytesData` | `semantic_diff/haxe_io_misc_contract`, `stdlib/haxe_io_misc_direct` |
| `haxe.io` typed arrays (`ArrayBufferView`, `UInt8Array`, `UInt16Array`, `UInt32Array`, `Int32Array`, `Float32Array`, `Float64Array`) | `mixed` | `std/haxe/io/*.cross.hx` | compiler-owned `haxe.io.Bytes` / `ArrayBufferViewImpl` carrier plus source-owned abstract static-method/default-arg routing in the compiler; float arrays reuse staged `haxe.io.FPHelper` instead of adding more compiler-owned bytes logic | `semantic_diff/haxe_io_typed_arrays_contract`, `stdlib/haxe_io_typed_arrays_direct` |
| `sys.db` direct surfaces (`Connection`, `ResultSet`, `Mysql`, `Sqlite`) | `mixed` | upstream `std/sys/db/**` interfaces and platform stubs | no fake DB runtime; Go keeps the upstream platform contract where `Mysql.connect` / `Sqlite.open` remain explicit unsupported runtime stubs instead of inventing target-owned behavior | `semantic_diff/sys_db_io_contract`, `stdlib/sys_db_io_direct` |
| `sys.io` direct handle surfaces (`FileInput`, `FileOutput`, `FileSeek`) | `mixed` | compiler-owned `lowerSysStdlibShimDecls` type/wrapper layer | `runtime/hxrt/sys.go` file-handle runtime (`OpenFileInput`, `OpenFile*Output`, seek/tell/eof/write/read`) beneath the public Haxe file-handle API | `semantic_diff/sys_db_io_contract`, `stdlib/sys_db_io_direct` |
| `sys.ssl` direct surfaces (`Certificate`, `Digest`, `DigestAlgorithm`, `Key`, `Socket`) | `mixed` | `std/sys/ssl/*.cross.hx` for the public API; `DigestAlgorithm` is fully source-owned | `runtime/hxrt/ssl.go` for certificate parsing, key parsing, digest/sign/verify, and TLS socket dial/listen/handshake/SNI selection helpers beneath the public wrappers | `stdlib/sys_ssl_leaf_direct`, `stdlib/sys_ssl_socket_direct`, `stdlib/sys_ssl_socket_sni_direct`, `semantic_diff/sys_net_address_ssl_digest_algorithm_contract`, `stdlib/sys_net_address_ssl_digest_algorithm_direct` |
| `sys.thread` direct surfaces (`Condition`, `Deque`, `EventLoop`, `ElasticThreadPool`, `FixedThreadPool`, `IThreadPool`, `Lock`, `Mutex`, `NoEventLoopException`, `Semaphore`, `Thread`, `ThreadPoolException`, `Tls`) | `mixed` | `std/sys/thread/*.cross.hx` for the public API, queue/storage helpers, and pool policies | `runtime/hxrt/thread.go` for blocking primitives, logical thread identity, message queues, and event-loop handles beneath the staged wrappers | `semantic_diff/sys_thread_primitives_contract`, `semantic_diff/sys_thread_runtime_contract`, `stdlib/sys_thread_primitives_direct`, `stdlib/sys_thread_runtime_direct` |
| `haxe.rtti.*` (`CType`, `Meta`, `Rtti`, `XmlParser`) | `mixed` | `std/haxe/rtti/*.cross.hx` | class-token `__meta__` / `__rtti` lookup contract plus anonymous-record array-field mutation lowering in the compiler | `semantic_diff/haxe_rtti_direct_contract`, `stdlib/haxe_rtti_direct` |
| `Lambda` and generic `Iterable<T>` call sites | `mixed` | source-owned `Lambda` plus staged iterator overrides under `std/_std/haxe/iterators/*.cross.hx` | `src/reflaxe/go/compiler/GoLambdaIterableLowering.hx` owns the small representation bridge for arrays, `haxe.ds.List`, and unknown manual-iterator carriers; `GoCompiler.hx` keeps only the direct optimized call lowering around that bridge | `lambda_generic_iterable_count_empty_contract`, `lambda_iter_generic_iterable_contract`, `stringbuf_datetools_lambda_contract` |

## Notes on Staged Source Injection

Staged portable overrides are injected first for Go builds by:

- `src/reflaxe/go/CompilerBootstrap.hx`

Current staged Tier1 coverage includes the JSON family and `StringTools`, with additional migrations gated by semantic-diff and Tier1 conformance coverage.

## Transition Notes (Post-`__go__` Audit)

- `haxe.io.Bytes`
  - The current `mixed` classification is still correct for parity today.
  - The first post-`__go__` extraction already moved pure hex and `BytesBuffer` leaf helpers into `runtime/hxrt/bytes.go`, leaving thin compiler wrappers in place.
  - The remaining compiler-owned subset is the RawNative/cache-coupled string path (`ofString`, `getString`, UTF16/raw-native conversion helpers) because it still co-owns `__hx_raw` cache validity and encoding-tag behavior.
  - The ownership lock is `stdlib/bytes_raw_native_compiler_ownership`, which proves RawNative `Bytes.set(...)` still needs to invalidate the cached raw-byte view seen by downstream consumers such as Base64.
  - Tracking: `haxe.go-14as.51`, `haxe.go-14as.54`
- `haxe.io.Input` / `haxe.io.Output`
  - These surfaces are not listed as separate Tier1 rows here, but their inherited helper loops no longer live as raw loop bodies in `GoCompiler`.
  - `readAll`, `readLine`, `readUntil`, `readFullBytes`, `write`, `writeFullBytes`, `writeInput`, and `writeString` now route through `std/haxe/io/GoIoHelpers.cross.hx`, with `GoCompiler` keeping only the public wrapper functions and the representation-sensitive base IO types.
  - Tracking: `haxe.go-14as.52`
- `haxe.io` misc direct tranche
  - `haxe.io.FPHelper` is now the model staged-std slice for this family: public bit-conversion behavior lives in `std/haxe/io/FPHelper.cross.hx` on top of the existing little-endian `BytesInput` / `BytesOutput` contract.
  - `haxe.io.Mime` and `haxe.io.Scheme` remain plain upstream source-owned string abstracts.
  - `haxe.io.StringInput`, `haxe.io.BufferInput`, `haxe.io.Encoding`, `haxe.io.Eof`, and `haxe.io.Error` stay compiler-owned with the base IO hierarchy because their type shapes and inherited helper wiring are still representation-sensitive on Go.
  - Tracking: `haxe.go-14as.15`
- `sys.Http`
  - Tier1 mapping still treats the surface as compiler-owned because request/callback choreography remains one semantic contract.
  - The audit narrowed extraction to leaf payload/proxy helpers only; `getResponseHeaderValues` and payload capture now live in `std/sys/GoHttpHelpers.cross.hx`, while core request sequencing and proxy URL construction stay in compiler scope unless parity evidence proves otherwise.
  - Tracking: `haxe.go-14as.53`

## Governance Rule

Any ownership change for a Tier1 module must update all of:

1. this mapping document,
2. `docs/ownership-rubric.md` when the rule itself changes,
3. `test/portable_conformance_tier1.json`,
4. `docs/stdlib-provenance-ledger.json` (when staged source files are added/changed),
5. relevant conformance fixtures in `test/semantic_diff`.
