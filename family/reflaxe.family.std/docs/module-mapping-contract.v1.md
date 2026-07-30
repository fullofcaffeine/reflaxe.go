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
| `Date` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/Date.hx` owns the epoch-millisecond carrier, constructors, all local/UTC accessors, and public API policy | `std/hxrt/date` maps scalar components and `DateParts` to footprint-explicit `runtime/hxrt/date.go` for host clock, timezone, parsing, formatting, and calendar conversion; no generated `Date` layout crosses the boundary | `date_source_owned`, `option_date_path`, `datetools_cross_std_contract`, `stdlib/date_math_source_owned`, direct runtime date tests; migration `haxe_go-vfp.8.7.15.4` |
| `Math` | `mixed` (`haxe_source` + typed native bindings) | `std/go/_std/Math.hx` owns Haxe rounding, finiteness, NaN propagation, and operand-order signed-zero policy | Typed `std/hxrt/math` externs call Go `math` and `math/rand` directly; only the three Int-returning rounders use footprint-explicit `runtime/hxrt/math.go` because Go's signatures return `float64` | `math_source_owned`, `numeric_edge_cases`, `stringtools_math`, `stdlib/date_math_source_owned`, `stdlib/math_float_native_no_hxrt`, direct runtime math tests; migration `haxe_go-vfp.8.7.15.4` |
| `UnicodeString` | `mixed` (`haxe_source` + representation binding) | `std/go/_std/UnicodeString.hx` owns all code-point bounds, slicing, searching, comparison, iteration, operators, and UTF-8 validation | Typed `std/hxrt/string/GoStringRuntime.hx` exposes only rune length, code-point lookup, and already-normalized slicing over pointer-backed Go strings; no Haxe library policy crosses the boundary | `unicode_string_source_owned`, `stdlib/unicode_string_basic`, direct `runtime/hxrt` string tests; migration `haxe_go-vfp.8.7.15.5` |
| `Std` | `mixed` (`haxe_source` + exact representation intrinsics + typed native bindings) | `std/go/_std/Std.hx` owns the complete Haxe 4.3.7 API: parseInt/parseFloat scanning and overflow policy, aliases, downcast behavior, truncation, and random bounds. `Std.string` and `Std.isOfType` remain individually registered because they consume erased Go representation or compiler-known type tokens. | Typed `std/hxrt/math` and `std/hxrt/string` bindings expose only native truncation, random generation, and exact float conversion; no parsing algorithm or trace policy lives in `hxrt` | `std_complete_api_contract`, `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract`, `typed_nil_dynamic_string_contract`, core Std snapshots; migration `haxe_go-vfp.8.7.22` |
| `haxe.Log` | `haxe_source` | `std/go/_std/haxe/Log.hx` owns complete `formatOutput` and mutable `trace` behavior, including `PosInfos`, custom parameters, direct function values, rebinding, null rebinding, and restoration | Source-owned `Sys.println`; general typed class lowering represents static dynamic methods as mutable Go function variables and converts null invocation into a Haxe exception, with no Log-specific compiler or runtime helper | `direct_haxe_helpers_contract`, `stdlib/std_log_source_owned`; migration `haxe_go-vfp.8.7.22` |
| `haxe.ds.Option` | `haxe_source` | Upstream Haxe 4.3.7 `haxe/ds/Option.hx`, queued and emitted through the same ordinary enum pipeline as project enums | No dedicated compiler carrier, raw constructor, or runtime helper | `option_date_path`, `stdlib/option_enum_basic`; residual-group retirement `haxe_go-vfp.8.7.15.7` |
| `DateTools` | `haxe_source` | `std/go/_std/DateTools.hx` | Staged `Date` and `Math` plus string primitives already owned elsewhere | `stringbuf_datetools_lambda_contract`, `datetools_cross_std_contract` |
| `StringTools` | `haxe_source` | `std/go/_std/StringTools.hx` | None beyond core string/runtime primitives used by normal lowering | `stringtools_math`, `stringtools_cross_std_contract` |
| `Lambda` / structural iterables | `mixed` (`haxe_source` + exact representation intrinsics) | `std/go/_std/Lambda.hx` owns all 19 public algorithms: traversal, comparison, early exit, allocation, filtering, mapping, folding, lookup, flattening, and concatenation. Its private representation-only `LambdaGoIterableCarrier` companion delegates `iterator()` without owning an algorithm. | `GoLambdaIterableLowering.hx` plus `lowerLambdaSourceCallAdapter` wrap only Go-invariant iterable, callback, constrained nested-iterable, and erased-result shapes. The same representation owner adapts concrete generated iterators to structural `Iterator<T>` maps and preserves ordered inline setup before typed concrete tails. Direct `ArrayIterator` sources retain a separate live-slice cursor specialization. Declarations, assignments, returns, ordinary calls, and constructors reuse that expected-type path; generic constructors select the emitted erased iterator ABI. The bridge adds no traversal policy, reflection, or runtime helper. Already-erased nested carriers fail with explicit diagnostics instead of reaching `go test`. | `lambda_full_api_contract`, `lambda_generic_iterable_count_empty_contract`, `lambda_iter_array_list_contract`, `lambda_iter_generic_iterable_contract`, `lambda_list_contract`, `structural_iterator_assignment_contract`, `inline_structural_iterator_effect_contract`, `inline_concrete_structural_iterator_effect_contract`, `structural_iterator_constructor_argument_contract`, `core/structural_iterator_assignment`, `core/inline_structural_iterator_effect`, `core/inline_concrete_structural_iterator_effect`, `core/structural_iterator_constructor_argument`, `negative/structural_iterator_element_mismatch`, `negative/inline_concrete_structural_iterator_element_mismatch`, `negative/structural_iterator_constructor_element_mismatch`, `negative/lambda_flatten_erased_nested`, `negative/lambda_flat_map_erased_result`; source migration `haxe_go-vfp.8.7.17`, full API/carrier closure `haxe_go-vfp.8.7.18`, structural assignment `haxe_go-vfp.8.3.3`, array inline effects `haxe_go-vfp.8.3.4`, constructor arguments `haxe_go-vfp.8.3.5`, concrete inline tails `haxe_go-vfp.8.3.6` |
| `haxe.Template` | `mixed` (`haxe_source` + exact generated metadata + `runtime_binding`) | `std/go/_std/haxe/Template.hx` owns parsing, lookup/fallback, iteration, macro argument construction, errors, and rendering | Selective compiler metadata lets the existing generic `Reflect.field` / `Reflect.hasField` path return already-bound lowercase generated methods; `std/hxrt/template/NativeTemplate.hx` still maps exactly three representation operations to footprint-explicit `runtime/hxrt/template.go`, and no Template-specific compiler helper exists | `haxe_template_contract`, `haxe_template_concrete_iterable_contract`, `stdlib/haxe_template_basic`, `stdlib/haxe_template_generated_method_lookup`, direct `template_test.go`; compiler bridge retirement `haxe_go-vfp.8.7.16`, method metadata `haxe_go-vfp.8.7.19` |
| `Sys` | `mixed` (`haxe_source` + `runtime_binding`, with one compile-time unsupported diagnostic) | `std/go/_std/Sys.hx` with typed bindings in `std/hxrt/sys` and `std/hxrt/fs` | `runtime/hxrt/sys.go` provides root process capabilities, `runtime/hxrt/file.go` provides non-owning standard-stream handles, and footprint-explicit build-tagged `runtime/hxrt/terminal*.go` provides terminal state plus one-byte input. Staged Haxe retains `getChar` EOF/echo semantics. The complete root surface is specified in [Portable root `Sys` contract](portable-sys-contract.md). Portable `putEnv` discards the retained native error for Haxe 4.3.7 eval parity; `setTimeLocale` reports `false`; `cpuTime` fails at Haxe compile time. | `root_sys_contract`, `root_sys_portable_contract`, `sys_io_roundtrip`, `file_error_semantics_contract`, `sys_sleep_contract`, `sys/root_sys_portable`, `sys/sys_get_char_terminal`, `test_sys_get_char_terminal.py`, `core/runtime_hxrt_infer_sys`, `negative/sys_cpu_time_unsupported` |
| `haxe.Utf8` | `haxe_source` | `std/go/_std/haxe/Utf8.hx` | None beyond `haxe.io.Bytes`, `UnicodeString`, and shared Go string runtime helpers already owned elsewhere | `haxe_utf8_contract`, `stdlib/haxe_utf8_basic` |
| `haxe.Json` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/Json.hx` | `runtime/hxrt/json.go` | `json_parse_stringify_contract` |
| `haxe.crypto.Base64`, `Md5`, `Sha1`, `Sha224`, `Sha256` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/crypto` owns public APIs, Base64 alphabets/padding, and public `Bytes` construction | `std/hxrt/crypto/NativeCrypto.hx` maps strings and opaque cached `ByteView` handles to footprint-explicit `runtime/hxrt/crypto.go`; no generated `haxe.io.Bytes` layout crosses the boundary and repeat integer/native-byte copies are avoided | `crypto_source_owned`, `crypto_xml_zip`, `stdlib/crypto_xml_zip_basic`, direct `crypto_test.go`; crypto migration `haxe_go-vfp.8.7.15.1`, shared byte view `haxe_go-vfp.8.7.11` |
| root `Xml`, `haxe.xml.Parser`, `haxe.xml.Printer` | `haxe_source` | `std/go/_std/Xml.hx` owns the DOM with only a nullable empty-read adaptation; its child mutation uses the upstream `Array.remove` / `Array.insert` calls through general typed slice lowering, and its four throwing accessors plus child iterator retain upstream inline forms; unchanged upstream Haxe 4.3.7 source owns the parser state machine; `std/go/_std/haxe/xml/Printer.hx` preserves upstream formatting without an incidental regex dependency | No native XML runtime or compiler declarations; ordinary staged `StringMap`, `StringTools`, and `StringBuf` dependencies only | `xml_source_owned`, `root_xml_contract`, `crypto_xml_zip`, `array_remove_insert_contract`, `core/array_remove_insert`, `inline_structural_iterator_effect_contract`, `inline_throw_accessor_result_type_contract`, `core/inline_throw_accessor_result_type`, `stdlib/xml_root_dom_basic`, `stdlib/crypto_xml_zip_basic`; migration `haxe_go-vfp.8.7.15.2`, iterator cleanup `haxe_go-vfp.8.3.3`, inline iterator restoration `haxe_go-vfp.8.3.4`, inline accessor restoration `haxe_go-vfp.8.3.1`, array mutation restoration `haxe_go-vfp.8.3.2` |
| `haxe.zip.Compress`, `haxe.zip.Uncompress` | `mixed` (`haxe_source` + `runtime_binding`) | Staged overrides own levels, optional defaults, `Bytes` conversion, offsets, bounded result records, flush/lifecycle policy, static helpers, and raw-DEFLATE selection; the rest of `haxe.zip` remains ordinary upstream source | `std/hxrt/zip` maps opaque typed deflate/inflate handles plus one typed step carrier to footprint-explicit `runtime/hxrt/zip.go`; the live codecs pause across partial buffers and no generated `haxe.io.Bytes` layout crosses the boundary | `zip_source_owned`, `zip_streaming_contract`, `crypto_xml_zip`, `stdlib/crypto_xml_zip_basic`, `stdlib/zip_streaming_policy`, direct runtime/race tests; source migration `haxe_go-vfp.8.7.15.3`, streaming completion `haxe_go-vfp.8.7.21`; [streaming contract](haxe-zip-streaming.md) |
| `haxe.ds.EnumValueMap` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/ds/EnumValueMap.hx` owns recursive comparison, AVL balancing, iteration, copying, and the public API | `std/hxrt/collections/NativeEnumValue.hx` and `runtime/hxrt/enum_value.go` provide only the generated-enum carrier predicate | `ds_maps_list_contract`; source migration `haxe_go-vfp.8.7.10` |
| `haxe.ds.IntMap` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/ds/IntMap.hx` owns the public API, iteration, copying, rendering, and clearing | Typed `std/hxrt/collections` bindings over selectively copied `runtime/hxrt/map_int.go` provide native storage and deterministic key snapshots | `ds_maps_list_contract`, `core/runtime_hxrt_infer_map_int`; source migration `haxe_go-vfp.8.7.10` |
| `haxe.ds.ObjectMap` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/ds/ObjectMap.hx` owns the public API, iteration, copying, rendering, and clearing | Typed `std/hxrt/collections` bindings over selectively copied `runtime/hxrt/map_object.go` provide retained reference-identity storage and deterministic key snapshots | `ds_maps_list_contract`, `core/runtime_hxrt_infer_map_object`; source migration `haxe_go-vfp.8.7.10` |
| `haxe.ds.StringMap` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/haxe/ds/StringMap.hx` owns the public API, iteration, copying, rendering, and clearing | Typed `std/hxrt/collections` bindings over selectively copied `runtime/hxrt/map_string.go` provide native storage and deterministic key snapshots | `ds_maps_list_contract`, `core/runtime_hxrt_infer_map_string`; source migration `haxe_go-vfp.8.7.10` |
| `haxe.ds.List` | `haxe_source` | `std/go/_std/haxe/ds/List.hx`; Lambda and serializer integration now use only its public iterator/API instead of reading its private carrier | None beyond ordinary array/string primitives | `ds_maps_list_contract`, `list_std_contract`, `lambda_list_contract`; source migration `haxe_go-vfp.8.7.10`, adapter cleanup `haxe_go-vfp.8.7.17` |
| `haxe.io.Bytes` | `mixed` (`haxe_source` + typed representation capabilities) | `std/go/_std/haxe/io/Bytes.hx` owns the public buffer API, bounds and mutation rules, hex and string policy, `BytesData` alias observation, RawNative selection, and opaque-view cache invalidation | `std/hxrt/io` exposes only opaque `ByteView`, allocation/conversion/copy, and UTF codec capabilities backed by `runtime/hxrt/bytes.go`; generated `Bytes` layout and public policy never cross the boundary | `bytes_hex_contract`, `bytes_io_stream_contract`, `bytes_normalization_contract`, `bytes_of_data_contract`, `bytes_ops_contract`, `io_encoding_contract`, generated snapshots, strict sweep, and direct runtime/perf gates; migration `haxe_go-vfp.8.7.11`; [ownership contract](haxe-io-ownership.md) |
| base `haxe.io` streams and support (`BytesBuffer`, `Input`, `Output`, encoding, EOF, error, FP helpers) | `mixed` (`haxe_source` + typed representation capabilities) | Canonical overrides under `std/go/_std/haxe/io` own every public type, stream loop, bounds/error rule, endian operation, alias contract, and IEEE-754 word-order policy | The same narrow `std/hxrt/io` boundary supplies byte capabilities plus scalar IEEE-754 reinterpretation; no `io` compiler group or generated public IO authority remains | bytes/IO/serializer/sys-IO semantic-diff, generated snapshots, strict sweep, direct runtime/perf gates; migration `haxe_go-vfp.8.7.11`; [ownership contract](haxe-io-ownership.md) |
| `haxe.io.Path` | `haxe_source` | upstream Haxe stdlib `haxe/io/Path.hx` | Core string and array helpers lowered by the target (`lastIndexOf`, `split`, `Array.join`, `String.fromCharCode`) | `option_date_path`, `path_cross_std_contract` |
| `sys.FileSystem` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/sys/FileSystem.hx` with typed carriers in `std/hxrt/fs` | `runtime/hxrt/filesystem.go`, `runtime/hxrt/string.go`, `runtime/hxrt/exception.go` | `filesystem_contract`, `sys/filesystem_basic_smoke` |
| `sys.io.File` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/sys/io/File.hx`, `FileInput.hx`, `FileOutput.hx`, and `FileSeek.hx`, with typed carriers in `std/hxrt/fs` | selectively copied `runtime/hxrt/file.go`; native OS failures remain Haxe exceptions, while byte conversion, bounds/EOF semantics, seek mapping, and public stream construction stay in Haxe source | `file_read_write_contract`, `file_error_semantics_contract`, `semantic_diff/sys_db_io_contract`, `sys/file_error_semantics`, `stdlib/sys_db_io_direct` |
| `sys.io.Process` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/sys/io/Process.hx` with typed opaque carriers in `std/hxrt/process` | selectively copied `runtime/hxrt/process.go`; native spawn, pipes, waits, signals, and close stay in Go, while stream construction, bounds/EOF translation, detached rejection, `Null<Int>` exit status, and public lifecycle policy stay in Haxe source | `process_echo_contract`, `process_error_semantics_contract`, `sys/process_error_semantics`, `core/runtime_hxrt_infer_process` |
| `sys.net.Host` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/sys/net/Host.hx` owns the public API and address conversions | Typed DNS/address helpers in `std/hxrt/net` over footprint-explicit `runtime/hxrt/socket.go`; no compiler-owned host declarations remain | `host_basic_contract`, `sys/host_basic_smoke`; source migration `haxe_go-vfp.8.7.14` |
| `sys.net.Socket` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/sys/net/Socket.hx` and `std/sys/net/_SocketIO.hx` own the public API, bind/listen sequencing, stream adapters, Haxe exceptions, and select object identity | Typed opaque handles and concrete result carriers in `std/hxrt/net` over footprint-explicit `runtime/hxrt/socket.go` plus build-tagged listener and readiness adapters; typed connection capabilities own transactional installation, TCP/TLS shutdown, and socket options; no `net_socket` compiler emitter or authority remains | `socket_advanced_contract`, `socket_loopback_contract`, `socket_server_lifecycle_contract`, `socket_readiness_contract`, `sys/socket_input_service_surface`, `sys/socket_server_lifecycle_contract`, `sys/socket_readiness_contract`, `stdlib/sys_ssl_socket_direct`, `core/runtime_hxrt_infer_socket`, direct runtime race/backlog/saturation/OOB/TLS-control/convergence tests; source migration `haxe_go-vfp.8.7.14`; server lifecycle `haxe_go-vfp.10.9.5`; real readiness `haxe_go-vfp.10.9.6`; TLS controls `haxe_go-vfp.10.9.7`; convergence review `haxe_go-vfp.10.9.8` |
| `sys.net.UdpSocket` | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/sys/net/UdpSocket.hx` owns the public API, peer-address construction, and Haxe exception translation | The same typed `std/hxrt/net` handle boundary over footprint-explicit `runtime/hxrt/socket.go` owns datagram transport; build-tagged `socket_broadcast_*.go` helpers adapt the native socket-option descriptor on POSIX and Windows | `stdlib/sys_net_udp_socket_direct`, direct runtime race tests, `test_socket_runtime_cross_build.py`; policy spike: `docs/spikes/ssl-udp-semantic-diff-spike.md`; source migration `haxe_go-vfp.8.7.14` |

## Additional Mixed-Ownership Rows

These modules are outside the Tier1 seed table above, but they are important enough to record explicitly because
their ownership split is easy to misunderstand.

| Module family | Ownership class | Public implementation location | Backend-owned support beneath it | Evidence |
| --- | --- | --- | --- | --- |
| `haxe.io` misc direct surfaces (`BufferInput`, `BytesData`, `Encoding`, `Eof`, `Error`, `FPHelper`, `Mime`, `Scheme`, `StringInput`) | `haxe_source`, with typed runtime capabilities for `Bytes` and `FPHelper` | canonical overrides under `std/go/_std/haxe/io` plus unchanged upstream source for simple abstracts | opaque `ByteView`/UTF/copy and IEEE-754 scalar capabilities only; no compiler public type or algorithm | `semantic_diff/haxe_io_misc_contract`, `stdlib/haxe_io_misc_direct`, `haxe_go-vfp.8.7.11` |
| `haxe.io` typed arrays (`ArrayBufferView`, `UInt8Array`, `UInt16Array`, `UInt32Array`, `Int32Array`, `Float32Array`, `Float64Array`) | `mixed` (`haxe_source` + ordinary representation lowering) | `std/go/_std/haxe/io/*.hx` owns typed-array behavior over the now-staged `Bytes` hierarchy | ordinary abstract/type representation mapping plus the shared typed byte/float capabilities; no IO compiler shim | `semantic_diff/haxe_io_typed_arrays_contract`, `stdlib/haxe_io_typed_arrays_direct`, `haxe_go-vfp.8.7.11` |
| `sys.db` direct surfaces (`Connection`, `ResultSet`, `Mysql`, `Sqlite`) | `mixed` | upstream `std/sys/db/**` interfaces and platform stubs | no fake DB runtime; Go keeps the upstream platform contract where `Mysql.connect` / `Sqlite.open` remain explicit unsupported runtime stubs instead of inventing target-owned behavior | `semantic_diff/sys_db_io_contract`, `stdlib/sys_db_io_direct` |
| `sys.io` direct handle surfaces (`FileInput`, `FileOutput`, `FileSeek`) | `mixed` (`FileSeek` is fully source-owned) | canonical staged modules under `std/go/_std/sys/io`; no File-specific compiler declarations or branches remain | typed opaque handles and native operations in `std/hxrt/fs` + `runtime/hxrt/file.go` beneath the source-owned public Haxe stream API | `semantic_diff/sys_db_io_contract`, `stdlib/sys_db_io_direct` |
| `sys.ssl` direct surfaces (`Certificate`, `Digest`, `DigestAlgorithm`, `Key`, `Socket`) | `mixed` (`haxe_source` + `runtime_binding`) | `std/go/_std/sys/ssl/*.hx` owns the public API and composes `sys.ssl.Socket` over the shared source-owned `sys.net.Socket`; `DigestAlgorithm` is fully source-owned | Typed carriers in `std/hxrt/ssl` plus `std/hxrt/net/SocketEndpoint.hx`; footprint-explicit `runtime/hxrt/ssl.go` owns certificate/key/digest primitives while `runtime/hxrt/socket_ssl.go` owns TLS handshake, peer-certificate, and SNI composition over the shared socket handle. The endpoint keeps the numeric route separate from the logical verification identity; shared typed connection capabilities send TLS close-notify and reach the wrapped TCP option without exposing native handles. | `stdlib/sys_ssl_leaf_direct`, `stdlib/sys_ssl_socket_direct`, `stdlib/sys_ssl_socket_sni_direct`, `core/runtime_hxrt_infer_socket_ssl`, direct runtime TLS-control/race/convergence tests, `semantic_diff/sys_net_address_ssl_digest_algorithm_contract`, `stdlib/sys_net_address_ssl_digest_algorithm_direct`; policy spike: `docs/spikes/ssl-udp-semantic-diff-spike.md`; socket source migration `haxe_go-vfp.8.7.14`; logical-host repair `haxe_go-vfp.10.9.3`; TLS controls `haxe_go-vfp.10.9.7`; convergence review `haxe_go-vfp.10.9.8` |
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
  - The `mixed` classification now means canonical staged source over a narrow
    typed runtime capability, not compiler migration debt.
  - Staged `Bytes` owns hex, bounds, mutation, RawNative selection, alias
    observation, and opaque-view cache invalidation. `std/hxrt/io` owns only
    native allocation/conversion/copy/UTF operations.
  - `stdlib/bytes_raw_native_compiler_ownership` now proves the replacement
    source-owned shape and cache invalidation seen by Base64/digests.
  - Closed evidence: `haxe.go-14as.51`, `haxe.go-14as.54`,
    `haxe_go-vfp.8.7.11`
- `haxe.io.Input` / `haxe.io.Output`
  - `readAll`, `readLine`, `readUntil`, `readFullBytes`, `write`,
    `writeFullBytes`, `writeInput`, and `writeString` live directly in canonical
    staged base classes and dispatch through normal source inheritance.
  - `GoIoHelpers`, IO-specific synthetic wrappers, and the broad compiler group
    are retired.
  - Closed evidence: `haxe.go-14as.52`, `haxe_go-vfp.8.7.11`
- `haxe.io` misc direct tranche
  - `haxe.io.FPHelper` owns public bit-conversion and word-order behavior over
    typed scalar `NativeFloatBits` capabilities, avoiding recursive stream calls.
  - `haxe.io.Mime` and `haxe.io.Scheme` remain plain upstream source-owned string abstracts.
  - `StringInput`, `BufferInput`, `Encoding`, `Eof`, and `Error` are canonical
    staged source with no compiler-owned carrier.
  - Baseline evidence: `haxe.go-14as.15`; closure:
    `haxe_go-vfp.8.7.11`
- `sys.Http`
  - Canonical `std/go/_std/sys/Http.hx` owns the Haxe 4.3.7 API, request
    selection, data-URL behavior, payload/header assembly, callback order,
    response maps, and status/error policy.
  - Typed `std/hxrt/http` bindings cross only opaque request/response handles,
    scalar header access, `ByteView`, and the existing typed `SocketHandle` into
    footprint-explicit `runtime/hxrt/http.go`. No generated `sys.Http` or
    `haxe.io.Bytes` layout crosses that boundary.
  - The public `portable|metal` selector does not choose HTTP semantics. Both
    compatibility presets use the same staged API; native transport is an
    API-scoped capability selected by typed reachability.
  - Baseline evidence: `haxe.go-14as.53`; completed migration:
    `haxe_go-vfp.8.7.12`; lifecycle/partial-I/O audit:
    `haxe_go-vfp.10.4`

## Governance Rule

Any ownership change for a Tier1 module must update all of:

1. this mapping document,
2. `docs/ownership-rubric.md` when the rule itself changes,
3. `test/portable_conformance_tier1.json`,
4. `docs/stdlib-provenance-ledger.json` (when staged source files are added/changed),
5. relevant conformance fixtures in `test/semantic_diff`.
