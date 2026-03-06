# Feature Support Matrix and Unsupported Inventory

How to read this page:

1. Start with the support contract below to understand what "Supported" means in this repo.
2. Use the coverage tiers table to understand the strength of evidence behind each surface.
3. Use linked fixtures to inspect exact expected behavior.

Key terms:

- `snapshot`: checks generated Go code shape/text (`docs/snapshot-policy.md`).
- `semantic-diff`: compares runtime behavior of generated Go vs Haxe `--interp` (`docs/semantic-diff-guide.md`).
- `portable` / `metal`: compiler profiles (contracts), documented in `docs/profiles.md`.

## Support contract

`reflaxe.go` treats a surface area as **supported** when all of the following are true:

1. It has snapshot or harness coverage in `test/snapshot`.
2. It runs through `python3 test/run-ci.py` in CI (`snapshot` + `stdlib sweep` + `semantic diff` + `examples`).
3. Generated Go passes `go test ./...` for the covered case(s).

Anything outside that bar is either **partial** (implemented but not fully gated) or **unsupported**.

Portable semantic rulebook for high-risk behavior classes is versioned in `docs/portable-semantics-v1.md`.
Tier1 ownership mapping for portable modules is tracked in `docs/portable-module-mapping-contract.md`.

Portable-eligible upstream module status is tracked in machine-readable form at:

- `test/portable_stdlib_inventory.json` (validated by `python3 test/run-portable-stdlib-inventory.py`)
- summary artifact: `test/.test-cache/portable_stdlib_inventory_summary.json`
- `test/portable_allowlist.json` (validated by `python3 test/run-portable-allowlist.py`)
- closure summary: `test/.test-cache/portable_parity_closure_summary.json` (validated by `python3 test/run-portable-parity-closure.py`)

## Coverage tiers

Coverage is tracked in explicit tiers; a surface can appear in multiple tiers, and the highest tier is its effective guarantee.

| Tier | Guarantee | Primary evidence |
| --- | --- | --- |
| `compile-only` | Symbol/type availability and codegen viability (`go test` on generated probe), without runtime parity guarantees on its own | `test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_full.txt --strict --go-test` |
| `snapshot` | Deterministic generated Go shape and targeted runtime smoke behavior | `test/run-snapshots.py`, fixtures in `test/snapshot` |
| `semantic-diff` | Runtime output parity against Haxe `--interp` for deterministic behavior contracts | `test/run-semantic-diff.py`, fixtures in `test/semantic_diff` |

`compile-only` by itself is a **partial** guarantee. Surfaces become fully supported only when they satisfy the support contract above.

### Representative tier map

| Surface | Highest tier | Evidence |
| --- | --- | --- |
| `haxe.Serializer` / `haxe.Unserializer` | `semantic-diff` | `serializer_wire_contract`, `serializer_cache_reference_contract`, `serializer_global_flags_contract`, `serializer_resolver_polymorphism_contract`, `serializer_reference_stress_contract` |
| `haxe.Exception` (`caught`/`thrown`/`message`) | `semantic-diff` | `exception_api_contract`, `exceptions_typed_dynamic` |
| `EReg` | `semantic-diff` | `ereg_behavior_contract`, `ereg_edge_contract` |
| `haxe.Http` / `sys.Http` | `semantic-diff` | `http_proxy_custom_request`, `http_request_callbacks_contract` |
| `sys.net.Socket` | `semantic-diff` | `socket_loopback_contract`, `socket_advanced_contract` |
| `haxe.crypto.*` + `haxe.xml.*` + `haxe.zip.*` subset | `semantic-diff` | `crypto_xml_zip` |
| `haxe.Json` | `semantic-diff` | `json_parse_stringify_contract`, `stdlib/json_parse_stringify` |
| `haxe.io.Bytes` / `haxe.io.BytesBuffer` / `haxe.io.BytesInput` / `haxe.io.BytesOutput` (core ops + Input/Output helper subset) | `semantic-diff` | `bytes_normalization_contract`, `bytes_ops_contract`, `bytes_of_data_contract`, `bytes_hex_contract`, `bytes_io_stream_contract`, `io_input_output_helpers_contract`, `io_input_output_edge_contract`, `stdlib/bytes_basic` |
| `sys.io.Process` | `semantic-diff` | `process_echo_contract`, `sys/process_echo_smoke` |
| `sys.io.File` | `semantic-diff` | `file_read_write_contract`, `sys/file_read_write_smoke` |
| `sys.FileSystem` | `semantic-diff` | `filesystem_contract`, `sys/filesystem_basic_smoke` |
| `Xml` (root DOM subset: document/element creation, attributes, child iteration, parse/print baseline) | `semantic-diff` | `root_xml_contract`, `stdlib/xml_root_dom_basic` |
| `haxe.ds.*Map` + `haxe.ds.List` (core ops subset) | `semantic-diff` | `ds_maps_list_contract`, `stdlib/ds_maps_list_basic` |
| `Reflect` (compare + dynamic field ops subset) | `semantic-diff` | `reflect_compare`, `reflect_field_ops` |
| `Date` + `haxe.ds.Option` + `haxe.io.Path` subset | `semantic-diff` | `option_date_path`, `path_cross_std_contract`, `stdlib/date_path_basic`, `stdlib/path_cross_std_basic` |
| `haxe.Log` + `haxe.Resource` + `haxe.SysTools` direct helper subset | `semantic-diff` | `direct_haxe_helpers_contract`, `direct_haxe_resource_contract` |
| `Array` + `IntIterator` core subset (`push`/`pop` statement-form, length/index access, array iteration, `0...N` range iteration) | `semantic-diff` | `array_string_intiterator_contract` |
| `String` core subset (`length`, `charAt`, `charCodeAt` with null on out-of-range, `substring`, `fromCharCode`) | `semantic-diff` | `array_string_intiterator_contract`, `string_charcodeat_bounds_contract` |
| `StringBuf` + `DateTools` + `Lambda` core subset (`add`/`addChar`/`addSub`, `DateTools.delta` with duration helpers, staged-std `DateTools.format` strftime subset plus `getMonthDays`/`parse`/`make`, `Lambda.filter`/`map`/`fold`/`has`/`exists`/`iter` over `Array<T>`/`haxe.ds.List<T>`, generic-`Iterable<T>` support for `Lambda.count`/`Lambda.empty`/`Lambda.exists`/`Lambda.has`/`Lambda.filter`/`Lambda.map`/`Lambda.fold`/`Lambda.iter` including function-value `Lambda.map` call-sites) | `semantic-diff` | `stringbuf_datetools_lambda_contract`, `datetools_cross_std_contract`, `lambda_list_contract`, `lambda_generic_iterable_count_empty_contract`, `lambda_iter_array_list_contract`, `lambda_iter_generic_iterable_contract` |
| Core `Class`/`Enum`/`EnumValue` type-value subset | `semantic-diff` | `type_expr_contract` |
| `Type` reflection subset (`getClass`, `getSuperClass`, `getClassFields`, `getInstanceFields`, `resolveClass`, `createInstance`, `createEmptyInstance`, `getEnum`, `getEnumConstructs`, `resolveEnum`, `allEnums`, enum constructor/index/parameters/equality) | `semantic-diff` | `type_reflection_contract`, `type_reflection_extended_contract` |
| `haxe.ds.Vector` | `semantic-diff` | `vector_contract`, `stdlib/vector_basic` |
| `haxe.ds.ReadOnlyArray` (length/index/read-only view subset) | `semantic-diff` | `readonly_array_contract` |
| `sys.net.Host` | `semantic-diff` | `host_basic_contract`, `sys/host_basic_smoke` |
| `haxe.PosInfos` | `semantic-diff` | `posinfos_contract`, `stdlib/posinfos_basic` |
| `haxe.Int32` | `semantic-diff` | `int32_contract` |
| `haxe.Int64` / `haxe.Int64Helper` | `semantic-diff` | `int64_contract`, `stdlib/int64_parity` |
| `Std.isOfType` | `semantic-diff` | `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract`, `core/std_is_of_type_basic`, `core/std_is_of_type_dynamic` |
| `haxe.atomic.AtomicInt` / `haxe.atomic.AtomicBool` | `semantic-diff` | `atomic_int_bool_contract`, `stdlib/atomic_int_bool_basic` |
| `haxe.atomic.AtomicObject` | `semantic-diff` | `atomic_object_contract`, `stdlib/atomic_object_basic` |

## Language/Core matrix

| Surface | Status | Evidence (snapshot IDs) |
| --- | --- | --- |
| Arithmetic, comparisons, boolean flow | Supported | `core/arithmetic`, `core/if_else` |
| Locals, assignment, compound assignment | Supported | `core/locals_assign`, `core/compound_assign`, `core/compound_assign_string` |
| Arrays and indexing | Supported | `core/array_basic`, `core/array_push_pop`, `core/array_constructor_new` |
| Loops (`for`, `while`, `do-while`, break/continue) | Supported | `core/loops_array_iter`, `core/loops_range`, `core/do_while_semantics`, `core/loop_break_continue` |
| Expression-form control flow in value positions (`if`, `switch`, `try`) | Supported | `core/if_expr_call_arg`, `core/switch_expr_call_arg`, `core/try_expr_value` |
| Functions, function values, closures | Supported | `core/function_values`, `core/closures_capture` |
| Default arguments and varargs | Supported | `core/default_args`, `core/varargs`, `core/varargs_foreach` |
| Prefix/postfix call argument order | Supported | `core/prefix_call_arg`, `core/postfix_call_arg` |
| Classes, instance/static fields/methods | Supported | `core/class_fields_methods`, `core/static_fields_methods` |
| Inheritance and override dispatch | Supported | `core/inheritance_override_dispatch`, `core/inheritance_ctor_chain_upcast`, `core/inheritance_return_upcast`, `core/inheritance_self_dispatch_wiring` |
| Interface dispatch | Supported | `core/interface_dispatch_basic` |
| Super calls | Supported | `core/super_calls` |
| Enums and switch pattern bindings | Supported | `core/enum_constructors`, `core/switch_enum_basic`, `core/enum_switch_bindings` |
| Anonymous object literals and structural field mutation | Supported | `core/object_literal_fields` |
| Exception subset (`throw`, typed/dynamic catch, rethrow, throw-as-expression, return-forwarding in statement-form try/catch, `haxe.Exception` API mapping) | Supported | `core/haxe_exception_subset`, `core/try_catch_typed`, `core/try_catch_dynamic`, `core/try_catch_rethrow`, `core/try_catch_return_forwarding`, `throw_expr_contract`, `try_catch_return_forwarding_contract`, `exception_api_contract` |
| `Std.isOfType` behavior | Supported | `core/std_is_of_type_basic`, `core/std_is_of_type_dynamic`, `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract` |
| Type-value expressions (`TTypeExpr`) for class/enum refs | Supported | `type_expr_contract` |
| `Type` reflection subset (`getClass`, `getSuperClass`, `getClassFields`, `getInstanceFields`, `resolveClass`, `createInstance`, `createEmptyInstance`, `getEnum`, `getEnumConstructs`, `resolveEnum`, `createEnum`, `createEnumIndex`, `allEnums`, `enumConstructor`, `enumIndex`, `enumParameters`, `enumEq`) | Supported | `type_reflection_contract`, `type_reflection_extended_contract` |
| Unsigned right shift behavior | Supported | `core/unsigned_shift`, `core/unsigned_shift_assign` |
| Naming/mangling and deterministic code shape | Supported | `core/naming_mangling`, `core/optimized_ast_policy` |
| HXML define/include resolution | Supported | `core/nested_hxml_define_detection`, `core/nested_hxml_long_define_detection`, `core/nested_hxml_quoted_define_detection`, `core/nested_hxml_root_relative_include_detection` |

### Semantic diff fixture coverage

- `test/semantic_diff/crypto_xml_zip`
- `test/semantic_diff/http_proxy_custom_request`
- `test/semantic_diff/http_request_callbacks_contract`
- `test/semantic_diff/socket_loopback_contract`
- `test/semantic_diff/socket_advanced_contract`
- `test/semantic_diff/null_string_concat`
- `test/semantic_diff/non_string_null_equality_contract`
- `test/semantic_diff/exceptions_typed_dynamic`
- `test/semantic_diff/exception_api_contract`
- `test/semantic_diff/enum_switch_bindings`
- `test/semantic_diff/virtual_dispatch`
- `test/semantic_diff/stringtools_math`
- `test/semantic_diff/string_charcodeat_bounds_contract`
- `test/semantic_diff/stringbuf_datetools_lambda_contract`
- `test/semantic_diff/lambda_list_contract`
- `test/semantic_diff/lambda_generic_iterable_count_empty_contract`
- `test/semantic_diff/option_date_path`
- `test/semantic_diff/array_string_intiterator_contract`
- `test/semantic_diff/numeric_edge_cases`
- `test/semantic_diff/nullable_struct_refs`
- `test/semantic_diff/sys_io_roundtrip`
- `test/semantic_diff/file_read_write_contract`
- `test/semantic_diff/filesystem_contract`
- `test/semantic_diff/ds_maps_list_contract`
- `test/semantic_diff/bytes_normalization_contract`
- `test/semantic_diff/bytes_ops_contract`
- `test/semantic_diff/bytes_of_data_contract`
- `test/semantic_diff/bytes_hex_contract`
- `test/semantic_diff/bytes_io_stream_contract`
- `test/semantic_diff/io_input_output_helpers_contract`
- `test/semantic_diff/io_input_output_edge_contract`
- `test/semantic_diff/host_basic_contract`
- `test/semantic_diff/int32_contract`
- `test/semantic_diff/int64_contract`
- `test/semantic_diff/posinfos_contract`
- `test/semantic_diff/posinfos_custom_params_contract`
- `test/semantic_diff/vector_contract`
- `test/semantic_diff/readonly_array_contract`
- `test/semantic_diff/process_echo_contract`
- `test/semantic_diff/reflect_compare`
- `test/semantic_diff/reflect_field_ops`
- `test/semantic_diff/anonymous_object_literals`
- `test/semantic_diff/serializer_unserializer_roundtrip`
- `test/semantic_diff/serializer_wire_contract`
- `test/semantic_diff/serializer_date_bytes_contract`
- `test/semantic_diff/serializer_class_enum_contract`
- `test/semantic_diff/serializer_extended_tokens_contract`
- `test/semantic_diff/serializer_custom_resolver_contract`
- `test/semantic_diff/serializer_cache_reference_contract`
- `test/semantic_diff/serializer_resolver_polymorphism_contract`
- `test/semantic_diff/serializer_resolver_type_value_contract`
- `test/semantic_diff/serializer_reference_stress_contract`
- `test/semantic_diff/serializer_global_flags_contract`
- `test/semantic_diff/ereg_behavior_contract`
- `test/semantic_diff/ereg_edge_contract`
- `test/semantic_diff/json_parse_stringify_contract`
- `test/semantic_diff/std_is_of_type_contract`
- `test/semantic_diff/std_is_of_type_runtime_core_abstract_contract`
- `test/semantic_diff/type_expr_contract`
- `test/semantic_diff/type_reflection_contract`
- `test/semantic_diff/type_reflection_extended_contract`
- `test/semantic_diff/throw_expr_contract`
- `test/semantic_diff/try_catch_return_forwarding_contract`
- `test/semantic_diff/atomic_int_bool_contract`
- `test/semantic_diff/atomic_object_contract`
- `test/semantic_diff/go_result_contract`
- `test/semantic_diff/go_chan_nonblocking_contract`
- `test/semantic_diff/go_select_contract`

## Profile matrix

| Surface | Status | Evidence (snapshot IDs) |
| --- | --- | --- |
| `portable` safe devirtualization path | Supported | `core/portable_leaf_virtual_devirtualization`, `core/portable_leaf_virtual_alias_devirtualization`, `core/portable_leaf_virtual_inline_ctor_devirtualization`, `core/portable_leaf_virtual_function_return_devirtualization`, `core/portable_non_leaf_virtual_dispatch_preserved` |
| `portable` string helper optimizations | Supported | `core/portable_string_ptr_helpers`, `core/portable_string_literal_folding` |
| Profile policy enforcement | Supported | `negative/profile_conflict`, `negative/profile_invalid` |
| Strict examples/app boundary policy + portable native-import policy modes | Supported | `negative/strict_examples_injection`, `negative/strict_mode_injection`, `negative/metal_profile_injection`, `negative/go_metal_lane_injection`, `negative/portable_native_import_error`, `core/portable_native_import_warn_policy` |
| RawNative encoding policy define (`reflaxe_go_raw_native_mode`) | Supported | `core/raw_native_utf16_mode`, `negative/raw_native_mode_invalid` |

## Go-native abstraction matrix

| Surface | Status | Evidence (snapshot IDs) |
| --- | --- | --- |
| Channels and goroutines | Supported (real goroutine/channel/select lowering; non-metal applies typed recv/recvOr/tryRecv assertion bridging for `go.Chan<T>` reads; typed deterministic `go.Select` helper API is available for receive/send branching; `metal` adds concrete typed shim lanes, including specialized `go.Select` helper routing) | `go_native/channel_basic`, `go_native/channel_try_recv`, `go_native/channel_select_handshake`, `go_native/channel_metal_monomorph`, `go_native/goroutine_smoke`, `go_native/select_helpers`, `go_native/select_metal_monomorph` |
| Extern metadata mapping | Supported (`@:go.import`/`@:go.name`/`@:go.receiver`, extern `String` return normalization via `hxrt.StdString`, and `@:go.valueError` mapping for `(T,error)` extern calls to `go.Result<T>`) | `go_native/extern_metadata_mapping`, `go_native/extern_value_error_result` |
| Result/Error mapping | Supported (`metal` adds typed `go.Result<T>` shim lowering with internal `(T,error)` helper emission) | `go_native/result_basic`, `go_native/error_result_mapping`, `go_native/result_metal_monomorph` |
| Slice/Map wrappers | Supported (`metal` adds typed shim specialization for concrete `go.Slice<T>` and `go.Map<K,V>` call-sites) | `go_native/slice_map_basic`, `go_native/slice_map_metal_monomorph` |

## Stdlib matrix

Shim strategy and alternatives are documented in:

- `docs/stdlib-shim-rationale.md`
- `docs/stdlib-shim-migration-log.md`

### Snapshot-level behavioral coverage

- `stdlib/bytes_basic`
- `stdlib/crypto_xml_zip_basic`
- `stdlib/date_path_basic`
- `stdlib/ds_maps_list_basic`
- `stdlib/intmap_basic`
- `stdlib/int64_parity`
- `stdlib/atomic_int_bool_basic`
- `stdlib/atomic_object_basic`
- `stdlib/io_type_smoke`
- `stdlib/json_parse_stringify`
- `stdlib/math_basic`
- `stdlib/option_enum_basic`
- `stdlib/posinfos_basic`
- `stdlib/posinfos_custom_params_smoke`
- `stdlib/stringtools_basic`
- `stdlib/vector_basic`
- `sys/file_read_write_smoke`
- `sys/filesystem_basic_smoke`
- `sys/http_custom_request_parity`
- `sys/http_proxy_socket_contract`
- `sys/http_request_callbacks_smoke`
- `sys/process_echo_smoke`
- `sys/host_basic_smoke`

### Explicit direct-helper exclusions

- `haxe.Template` direct usage is intentionally blocked until source-owned std inclusion can emit the module-local enum support it needs. Evidence: `negative/direct_haxe_template_unsupported`.
- `haxe.ValueException` direct usage is intentionally blocked until string-payload message parity is restored across `Any`/string boxing. Evidence: `negative/direct_haxe_value_exception_unsupported`.

### `haxe.Json` runtime-lowered contract

- `haxe.Json.parse`/`haxe.Json.stringify`, `haxe.format.JsonPrinter.print`, and `haxe.format.JsonParser.doParse` now lower directly to `hxrt.JsonParse`/`hxrt.JsonStringify`.
- Compiler-emitted JSON shim declarations were removed from `GoCompiler` (no `haxe__Json`/`haxe__format__JsonParser` synthetic declarations in generated output).
- Runtime parity evidence: `test/semantic_diff/json_parse_stringify_contract`.

### `Sys` / `sys.io.File` / `sys.io.Process` ownership contract

- Runtime behavior now lives in `runtime/hxrt/hxrt.go`:
  - `hxrt.SysGetCwd`, `hxrt.SysArgs`
  - `hxrt.FileSaveContent`, `hxrt.FileGetContent`
  - `hxrt.NewProcess`, `Process.Stdout`, `ProcessOutput.ReadLine`, `Process.Close`
- Compiler-generated `sys` declarations remain as thin wrappers to preserve Haxe type shape and call signatures.
- `lowerSysStdlibShimDecls` is forwarding-only for this surface; behavior changes must be implemented in runtime and verified by `sys/file_read_write_smoke`, `test/semantic_diff/file_read_write_contract`, `sys/process_echo_smoke`, and `test/semantic_diff/process_echo_contract`.

### `sys.FileSystem` shim contract and tradeoffs

- `sys.FileSystem` static calls now lower to compiler-emitted wrappers in `lowerFileSystemShimDecls`:
  - `sys__FileSystem_exists`, `rename`, `stat`, `fullPath`, `isDirectory`, `createDirectory`, `deleteFile`, `deleteDirectory`, `readDirectory`.
- Coverage now includes `sys/filesystem_basic_smoke` and `test/semantic_diff/filesystem_contract` (deterministic create/read/rename/delete/stat-size behavior).
- Current tradeoff: `stat` currently maps non-portable fields (`uid/gid/dev/ino/nlink/rdev`) to stable fallback values when unavailable from portable Go APIs.

### `haxe.ds.*Map` / `haxe.ds.List` shim contract and tradeoffs

- Coverage includes `stdlib/ds_maps_list_basic` and `test/semantic_diff/ds_maps_list_contract` for deterministic `set/get/exists/remove` behavior across `StringMap`/`IntMap`/`ObjectMap`/`EnumValueMap`, plus `List` `add`/`push`/`pop`/`first`/`last`/`length`.
- `List.push` now prepends to match Haxe semantics (with `pop` removing the list head).
- Missing-key map reads and empty `List` reads in typed call sites now lower through nil-safe typed assertions, returning typed zero values (`null` for reference-like types) instead of panicking.

### `haxe.io.BytesInput` / `haxe.io.BytesOutput` shim contract and tradeoffs

- Coverage includes `test/semantic_diff/bytes_io_stream_contract`, `test/semantic_diff/bytes_of_data_contract`, `test/semantic_diff/bytes_hex_contract`, `test/semantic_diff/io_input_output_helpers_contract`, `test/semantic_diff/io_input_output_edge_contract`, `test/semantic_diff/io_error_constructor_contract`, and `test/semantic_diff/io_encoding_contract` for deterministic constructor bounds checks, `position`/`length`, EOF behavior, `readByte`/`readBytes`, inherited helper subset parity (`readAll`, `readFullBytes`, `read`, `readUntil`, `readLine`, `readString`, `readFloat`/`readDouble`, signed/unsigned numeric reads), output helper subset parity (`write`, `writeFullBytes`, `writeInput`, `writeString`, numeric typed writes, overflow guards), `haxe.io.Error` typed constructor matching (`Blocked`, `Overflow`, `OutsideBounds`, `Custom`), `haxe.io.Encoding` constructor parity (`UTF8`, `RawNative`), `Bytes.getString` bounds behavior, `Bytes.getData`/`Bytes.ofData` alias semantics, `Bytes.toHex`/`Bytes.ofHex` behavior, and `readLine` EOF/tail/CRLF edge paths.
- Current tradeoff: parity remains focused on `BytesInput`/`BytesOutput` stream behavior with interpreter-compatible semantics by default (`reflaxe_go_raw_native_mode=interp`, where `UTF8` and `RawNative` both map to UTF-8 conversion). For projects that need Java/C#-style RawNative byte layout, `reflaxe_go_raw_native_mode=utf16le` provides an explicit opt-in UTF-16LE path; full target-by-target RawNative equivalence is still not claimed outside these documented modes.

### `haxe.Http` / `sys.Http` shim contract and tradeoffs

- `haxe.Http` is a `typedef` alias of `sys.Http` on `sys` targets, so the same semantic-diff fixtures now serve as the portable contract for both entry points.
- `sys.Http` now includes synchronous request semantics for `http`/`https` and deterministic `data:` handling used by tests.
- Covered behaviors: `setHeader`/`addHeader`, `setParameter`/`addParameter`, `setPostData`/`setPostBytes`, `fileTransfer`/`fileTransfert`, `customRequest` (including optional socket transport injection), proxy URL wiring (`Http.PROXY`), `getResponseHeaderValues`, dynamic callbacks (`onData`, `onBytes`, `onError`, `onStatus`), `responseData`/`responseBytes`, and `requestUrl`.
- Semantic diff now also locks callback/status/header/error parity for local deterministic HTTP servers (`http_request_callbacks_contract`), including 4xx `onError` formatting (`Http Error #<status>`).
- Current tradeoff: execution remains synchronous, and `customRequest` socket injection currently maps into Go `http.Transport` dialing semantics rather than the exact byte-level write/read loop used by upstream `sys.Http`.

### `sys.net.Socket` shim contract and tradeoffs

- Socket parity now covers deterministic loopback contracts (`bind`/`listen`/`connect`/`accept`/`read`/`write`/`close`) plus advanced methods: `setTimeout`, `waitForRead`, `setBlocking`, `setFastSend`, `select`, and `shutdown` (`socket_loopback_contract`, `socket_advanced_contract`).
- `select` now returns readiness-filtered arrays for read/write groups under timeout control, and `waitForRead` delegates to this readiness path.
- Current tradeoff: `setBlocking` is implemented through deadline behavior rather than true OS-level non-blocking file descriptor mode.

### `EReg` + `haxe.Serializer` contract and tradeoffs

- `EReg` parity now covers: `g/i/m/s/u` option handling, global vs non-global `replace`/`map`, `matched`/`matchedPos`/`matchedLeft`/`matchedRight` error semantics, and group/null behavior via semantic diff fixtures (`ereg_behavior_contract`, `ereg_edge_contract`).
- `haxe.Serializer`/`haxe.Unserializer` now cover a wire-format-compatible baseline for core tokens used by fixtures (`n/t/f/z/i/d/k/p/m/v/s/y/a/o/l/b/q/M/c/w/j/C/x/A/B/g/u/h/r/R`) plus sequential `Unserializer` cursor behavior (`serializer_wire_contract`), resolver paths (`serializer_custom_resolver_contract`), resolver method-shape polymorphism (`serializer_resolver_polymorphism_contract`), cache/reference graph parity (`serializer_cache_reference_contract`), global serializer default flag behavior (`Serializer.USE_CACHE`/`Serializer.USE_ENUM_INDEX`) with `serializeException` interaction (`serializer_global_flags_contract`), and mixed string/object reference stress (`serializer_reference_stress_contract`).
- Remaining gap: full Haxe serializer surface is still in progress (mainly exotic/uncommon resolver payload shapes and less-common cross-target edge combinations outside current fixtures).
- Active follow-up tracking:
  - `haxe.go-7zy.10` (migrate `haxe.Json` shim out of compiler core, completed 2026-02-19)
  - `haxe.go-7zy.11` (migrate `Sys`/`sys.io.File`/`sys.io.Process` shim path out of compiler core, completed 2026-02-19)
  - `haxe.go-7zy.12` (reduce `stdlib_symbols` bytes-conversion overhead, completed 2026-02-19)
  - `haxe.go-re8` (support resolver-returned type-value markers for class/enum name extraction + serialization)

### Upstream module sweep (strict CI-gated)

Source list: `test/upstream_std_modules.txt`
This sweep validates module symbol/type availability and target compatibility (`haxe` compile + `go test ./...`) for each listed module.

```text
haxe.CallStack
haxe.Exception
haxe.Int32
haxe.Int64
haxe.Http
haxe.Json
haxe.PosInfos
haxe.Serializer
haxe.Unserializer
haxe.crypto.Base64
haxe.crypto.Md5
haxe.crypto.Sha1
haxe.crypto.Sha224
haxe.crypto.Sha256
haxe.ds.BalancedTree
haxe.ds.EnumValueMap
haxe.ds.IntMap
haxe.ds.List
haxe.ds.Map
haxe.ds.ObjectMap
haxe.ds.Option
haxe.ds.ReadOnlyArray
haxe.ds.StringMap
haxe.ds.Vector
haxe.io.Bytes
haxe.io.BytesBuffer
haxe.io.BytesData
haxe.io.BytesInput
haxe.io.BytesOutput
haxe.io.Eof
haxe.io.Error
haxe.io.Input
haxe.io.Output
haxe.io.Path
haxe.io.StringInput
haxe.xml.Access
haxe.xml.Parser
haxe.xml.Printer
haxe.zip.Compress
haxe.zip.Uncompress
Date
EReg
Math
Reflect
Std
StringTools
Type
Xml
Sys
sys.FileSystem
sys.Http
sys.io.File
sys.io.Process
sys.net.Host
sys.net.Socket
```

### Full runtime-eligible inventory sweep

Source list: `test/upstream_std_modules_full.txt` (175 modules).

As of **2026-02-19**:

- Compile-only strict sweep:
  - `python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_full.txt --strict`
  - Result: `175 passed / 0 expected policy / 0 failed / 0 unexpected present`
- Compile + generated Go validation:
  - `python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_full.txt --strict --go-test`
  - Result: `175 passed / 0 expected policy / 0 failed / 0 unexpected present`

Policy sources:

- `test/upstream_std_expected_missing.json` (currently empty)
- `test/upstream_std_expected_unavailable.json` (currently empty)

## Unsupported expression inventory (compiler hard-fail paths)

These are explicit fatal guards in `src/reflaxe/go/GoCompiler.hx` that represent unsupported paths.
As of **2026-03-06**, the hard-fail inventory count remains **4**, with invariant fixture strategies now explicitly named per path.

| Inventory item | Current behavior | Fixture strategy (named) | Acceptance criteria for closure | Owner |
| --- | --- | --- | --- | --- |
| Non-lvalue assignment targets in `lowerLValue` | Fatal: `Unsupported assignment target` | `negative/non_lvalue_assignment_invariant` locks Haxe front-end rejection (`Invalid assign`) so backend fatal remains an explicit invariant unless a typed reproducer becomes reachable. | Either (a) support any newly reachable legal lvalue shape, or (b) keep as invariant and add a dedicated negative test if a reproducer becomes possible. | `haxe.go-14as.8` |
| Non-`++/--` postfix unary in `lowerExpr` / `lowerExprWithPrefix` | Fatal: `Unsupported postfix unary operator` | `negative/postfix_non_inc_dec_invariant` locks parser-level rejection (`Postfix ! is not supported`) so only `++/--` postfix forms can reach lowering today. | Keep parser/typed-ast assumptions validated; if new postfix forms become reachable, add lowering + snapshots before enabling. | `haxe.go-14as.8` |
| Catch-all `lowerExpr` default | Fatal: `Unsupported expression` | Node-family closure map: `semantic-diff/type_expr_contract`, `semantic-diff/throw_expr_contract`, `core/untyped_ident_nil`, and `core/const_kinds_contract`; new reachable node families must add a dedicated contract fixture in the same change. | Continue replacing reachable typed-node gaps with explicit lowering and keep dedicated coverage as each newly reachable node is supported. | `haxe.go-14as.8` |
| Unsupported `Std.isOfType` target kind | No compiler hard-fail for unresolved runtime-value abstract targets; falls back to conservative `false`/type-switch check | `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract`, `core/std_is_of_type_basic`, `core/std_is_of_type_dynamic`, `core/type_switch_no_binding_std_is_of_type` lock current fallback semantics. | Keep adding explicit lowering for newly important target families and lock behavior with semantic diff coverage. | `haxe.go-14as.8` |

## Known stdlib parity gaps (probe inventory)

As of **2026-02-19**, the broader probe list:

```bash
python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_gap_probe.txt --strict --go-test
```

reports:

- `53 passed / 0 expected policy / 0 failed / 0 unexpected present`

There are currently no active expected-policy rules in the full inventory.

## Tracking

- `haxe.go-d5u`: publish and maintain this matrix/inventory.
- `haxe.go-61w`: reduce compiler hard-fail unsupported expression surface.
- `haxe.go-19u`: expand stdlib parity from the documented probe gap list.
- `haxe.go-ab2`: add semantic differential regression harness.
- `haxe.go-3d4`: reduce unsupported expression surface by lowering `TTypeExpr` class/enum value nodes.
- `haxe.go-8zt`: lower `TThrow` in expression positions and lock with semantic diff coverage.
- `haxe.go-888`: promote `sys.FileSystem` with deterministic snapshot + semantic parity contracts.
- `haxe.go-uz4.10`: enable typed `go.Chan<T>` recv/recvOr assertions in `portable`.
- `haxe.go-6fc`: completed `haxe.ds` map/list core-ops semantic parity coverage and `List.push` parity alignment.
- `haxe.go-rlj`: completed nil-safe typed-read null semantics for `haxe.ds` map/list generic call results.
- `haxe.go-aiy`: add `haxe.io.Encoding` constructor parity and `Bytes.getString` coverage (`io_encoding_contract`).
- `haxe.go-dq2`: evaluate and guard RawNative compatibility policy with explicit mode controls.
- `haxe.go-rcv`: add `haxe.io.Bytes.ofData` shim and lock `getData` alias semantics (`bytes_of_data_contract`).
- `haxe.go-nmg`: add `haxe.io.Bytes.toHex` / `haxe.io.Bytes.ofHex` shim parity (`bytes_hex_contract`).
- `haxe.go-9v6`: promote `haxe.ds.ReadOnlyArray` from compile-only to semantic-diff coverage (`readonly_array_contract`).
- `haxe.go-8hs`: add serializer global default flag semantic coverage (`serializer_global_flags_contract`).
- `haxe.go-14as.6`: promote iterator-family parity (`haxe.iterators.ArrayIterator`, `ArrayKeyValueIterator`, `DynamicAccessIterator`, `DynamicAccessKeyValueIterator`, `RestIterator`, `RestKeyValueIterator`, `StringIterator`, `StringIteratorUnicode`, `StringKeyValueIterator`, `StringKeyValueIteratorUnicode`) via `iterators_family_contract` and staged std overrides.
- `haxe.go-14as.7`: add portable-vs-metal invariance coverage for iterator/list/map/string portable surfaces (`portable_surfaces_metal_invariance_contract`, `portable_surfaces_lane_invariance_contract`) and snapshot fallback-report attribution lock (`core/report_artifacts_lane_fallback_portable_surfaces`).
- `haxe.go-14as.8`: close unsupported-expression inventory with explicit invariant fixture strategy mapping (`negative/non_lvalue_assignment_invariant`, `negative/postfix_non_inc_dec_invariant`, plus node-family closure fixtures for `lowerExpr` and `Std.isOfType` fallback behavior).
- `haxe.go-14as.9`: add deterministic hxrt feature reason provenance to runtime reports (`core/report_artifacts_runtime_reason_provenance`, `test/run-auto-planner-schema.py`) without changing runtime semantics.
- `haxe.go-14as.10`: final closure sync task; inventory/closure artifacts now require explicit blocker metadata for remaining compile-only modules, and release readiness docs include portable parity + family sync gates.
- `haxe.go-14as.11`: root/core tranche triage closure. Direct semantic-diff contracts promoted `Any`, `StdTypes`, and `sys.FileStat`; remaining root blockers were split into dedicated tasks for `Sys`, `Xml`, and `UnicodeString`.
- `haxe.go-14as.20`: closed root `Sys` surface split. Added direct semantic-diff coverage in `root_sys_contract` and wired `Sys.getEnv`, `Sys.putEnv`, `Sys.environment`, and `Sys.systemName` through new hxrt/compiler shims.
- `haxe.go-14as.21`: closed root `Xml` surface split. Added direct semantic-diff coverage in `root_xml_contract` plus snapshot coverage in `stdlib/xml_root_dom_basic`.
- `haxe.go-14as.22`: closed root `UnicodeString` surface split. Added direct semantic-diff coverage in `root_unicode_string_contract` plus snapshot coverage in `stdlib/unicode_string_basic`, and wired the `_UnicodeString__UnicodeString_Impl__*` lowering surface through generated stdlib-symbol shims.
- `haxe.go-14as.24`: closed the remaining `Xml.parse()` parsed-CDATA node-type follow-up, so root `Xml` now preserves `CData` instead of collapsing it to `PCData`.
- `haxe.go-14as.12`: closed generic `haxe.misc` triage by promoting `haxe.Http` from existing semantic-diff fixtures and splitting the remaining modules into `haxe.go-14as.25` to `haxe.go-14as.29`.
- `haxe.go-14as.13` to `haxe.go-14as.19` and `haxe.go-14as.25` to `haxe.go-14as.29`: dated blocker families for the remaining compile-only portable tranches (`haxe_ds_exceptions`, `haxe_http_rtti`, `haxe_io_misc`, `haxe_io_typed_arrays`, `sys_db_io`, `sys_net_ssl`, `sys_thread`, `haxe_misc_symbols`, `haxe_misc_abstractions`, `haxe_misc_enum_helpers`, `haxe_misc_stack_loop`, `haxe_misc_legacy_text`).
