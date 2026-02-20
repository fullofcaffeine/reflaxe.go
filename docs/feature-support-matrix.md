# Feature Support Matrix and Unsupported Inventory

## Support contract

`reflaxe.go` treats a surface area as **supported** when all of the following are true:

1. It has snapshot or harness coverage in `test/snapshot`.
2. It runs through `python3 test/run-ci.py` in CI (`snapshot` + `stdlib sweep` + `semantic diff` + `examples`).
3. Generated Go passes `go test ./...` for the covered case(s).

Anything outside that bar is either **partial** (implemented but not fully gated) or **unsupported**.

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
| `haxe.Serializer` / `haxe.Unserializer` | `semantic-diff` | `serializer_wire_contract`, `serializer_cache_reference_contract`, `serializer_resolver_polymorphism_contract`, `serializer_reference_stress_contract` |
| `EReg` | `semantic-diff` | `ereg_behavior_contract`, `ereg_edge_contract` |
| `sys.Http` | `semantic-diff` | `http_proxy_custom_request`, `http_request_callbacks_contract` |
| `sys.net.Socket` | `semantic-diff` | `socket_loopback_contract`, `socket_advanced_contract` |
| `haxe.crypto.*` + `haxe.xml.*` + `haxe.zip.*` subset | `semantic-diff` | `crypto_xml_zip` |
| `haxe.Json` | `semantic-diff` | `json_parse_stringify_contract`, `stdlib/json_parse_stringify` |
| `haxe.io.Bytes` / `haxe.io.BytesBuffer` / `haxe.io.BytesInput` / `haxe.io.BytesOutput` (core ops + Input/Output helper subset) | `semantic-diff` | `bytes_normalization_contract`, `bytes_ops_contract`, `bytes_io_stream_contract`, `io_input_output_helpers_contract`, `io_input_output_edge_contract`, `stdlib/bytes_basic` |
| `sys.io.Process` | `semantic-diff` | `process_echo_contract`, `sys/process_echo_smoke` |
| `sys.io.File` | `semantic-diff` | `file_read_write_contract`, `sys/file_read_write_smoke` |
| `sys.FileSystem` | `semantic-diff` | `filesystem_contract`, `sys/filesystem_basic_smoke` |
| `haxe.ds.*Map` + `haxe.ds.List` (core ops subset) | `semantic-diff` | `ds_maps_list_contract`, `stdlib/ds_maps_list_basic` |
| `haxe.ds.Vector` | `semantic-diff` | `vector_contract`, `stdlib/vector_basic` |
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
| Exception subset (`throw`, typed/dynamic catch, rethrow, throw-as-expression`) | Supported | `core/haxe_exception_subset`, `core/try_catch_typed`, `core/try_catch_dynamic`, `core/try_catch_rethrow`, `throw_expr_contract` |
| `Std.isOfType` behavior | Supported | `core/std_is_of_type_basic`, `core/std_is_of_type_dynamic`, `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract` |
| Type-value expressions (`TTypeExpr`) for class/enum refs | Supported | `type_expr_contract` |
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
- `test/semantic_diff/exceptions_typed_dynamic`
- `test/semantic_diff/enum_switch_bindings`
- `test/semantic_diff/virtual_dispatch`
- `test/semantic_diff/stringtools_math`
- `test/semantic_diff/option_date_path`
- `test/semantic_diff/numeric_edge_cases`
- `test/semantic_diff/nullable_struct_refs`
- `test/semantic_diff/sys_io_roundtrip`
- `test/semantic_diff/file_read_write_contract`
- `test/semantic_diff/filesystem_contract`
- `test/semantic_diff/ds_maps_list_contract`
- `test/semantic_diff/bytes_normalization_contract`
- `test/semantic_diff/bytes_ops_contract`
- `test/semantic_diff/bytes_io_stream_contract`
- `test/semantic_diff/io_input_output_helpers_contract`
- `test/semantic_diff/io_input_output_edge_contract`
- `test/semantic_diff/host_basic_contract`
- `test/semantic_diff/int32_contract`
- `test/semantic_diff/int64_contract`
- `test/semantic_diff/posinfos_contract`
- `test/semantic_diff/posinfos_custom_params_contract`
- `test/semantic_diff/vector_contract`
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
- `test/semantic_diff/ereg_behavior_contract`
- `test/semantic_diff/ereg_edge_contract`
- `test/semantic_diff/json_parse_stringify_contract`
- `test/semantic_diff/std_is_of_type_contract`
- `test/semantic_diff/std_is_of_type_runtime_core_abstract_contract`
- `test/semantic_diff/type_expr_contract`
- `test/semantic_diff/throw_expr_contract`
- `test/semantic_diff/atomic_int_bool_contract`
- `test/semantic_diff/atomic_object_contract`

## Profile matrix

| Surface | Status | Evidence (snapshot IDs) |
| --- | --- | --- |
| `portable` dispatch path preserved | Supported | `core/portable_leaf_virtual_dispatch`, `core/portable_leaf_virtual_alias_dispatch`, `core/portable_leaf_virtual_inline_ctor_dispatch`, `core/portable_leaf_virtual_function_return_dispatch` |
| `gopher` safe devirtualization path | Supported | `core/gopher_leaf_virtual_devirtualization`, `core/gopher_leaf_virtual_alias_devirtualization`, `core/gopher_leaf_virtual_inline_ctor_devirtualization`, `core/gopher_leaf_virtual_function_return_devirtualization`, `core/gopher_non_leaf_virtual_dispatch_preserved` |
| `gopher` string helper optimizations | Supported | `core/gopher_string_ptr_helpers`, `core/gopher_string_literal_folding`, `core/portable_string_literal_no_folding` |
| Profile policy enforcement and removed aliases | Supported | `negative/profile_conflict`, `negative/profile_invalid`, `negative/profile_removed_idiomatic`, `negative/profile_removed_idiomatic_alias` |
| Strict examples/app boundary policy | Supported | `negative/strict_examples_injection`, `negative/strict_mode_injection`, `negative/metal_profile_injection` |

## Go-native abstraction matrix

| Surface | Status | Evidence (snapshot IDs) |
| --- | --- | --- |
| Channels and goroutines | Supported | `go_native/channel_basic`, `go_native/goroutine_smoke` |
| Result/Error mapping | Supported | `go_native/result_basic`, `go_native/error_result_mapping` |
| Slice/Map wrappers | Supported | `go_native/slice_map_basic` |

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

- Coverage includes `test/semantic_diff/bytes_io_stream_contract`, `test/semantic_diff/io_input_output_helpers_contract`, `test/semantic_diff/io_input_output_edge_contract`, and `test/semantic_diff/io_error_constructor_contract` for deterministic constructor bounds checks, `position`/`length`, EOF behavior, `readByte`/`readBytes`, inherited helper subset parity (`readAll`, `readFullBytes`, `read`, `readUntil`, `readLine`, `readString`, `readFloat`/`readDouble`, signed/unsigned numeric reads), output helper subset parity (`write`, `writeFullBytes`, `writeInput`, `writeString`, numeric typed writes, overflow guards), `haxe.io.Error` typed constructor matching (`Overflow`, `Custom`), and `readLine` EOF/tail/CRLF edge paths.
- Current tradeoff: parity remains focused on `BytesInput`/`BytesOutput` stream behavior and does not yet claim full cross-target `haxe.io.Input`/`haxe.io.Output` edge compatibility (for example less-common encoding variants and exhaustive constructor-path coverage for every `haxe.io.Error` throw site).

### `sys.Http` shim contract and tradeoffs

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
- `haxe.Serializer`/`haxe.Unserializer` now cover a wire-format-compatible baseline for core tokens used by fixtures (`n/t/f/z/i/d/k/p/m/v/s/y/a/o/l/b/q/M/c/w/j/C/x/A/B/g/u/h/r/R`) plus sequential `Unserializer` cursor behavior (`serializer_wire_contract`), resolver paths (`serializer_custom_resolver_contract`), resolver method-shape polymorphism (`serializer_resolver_polymorphism_contract`), cache/reference graph parity (`serializer_cache_reference_contract`), and mixed string/object reference stress (`serializer_reference_stress_contract`).
- Remaining gap: full Haxe serializer surface is still in progress (notably under-tested exotic edge combinations outside current fixtures and cross-target differences for less common resolver payload shapes).
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

| Inventory item | Current behavior | Acceptance criteria for closure |
| --- | --- | --- |
| Non-lvalue assignment targets in `lowerLValue` | Fatal: `Unsupported assignment target` | Either (a) support any newly reachable legal lvalue shape, or (b) keep as invariant and add a dedicated negative test if a reproducer becomes possible. |
| Non-`++/--` postfix unary in `lowerExpr` / `lowerExprWithPrefix` | Fatal: `Unsupported postfix unary operator` | Keep parser/typed-ast assumptions validated; if new postfix forms become reachable, add lowering + snapshots before enabling. |
| Catch-all `lowerExpr` default | Fatal: `Unsupported expression` | Continue replacing reachable typed-node gaps with explicit lowering (for example `TTypeExpr` class/enum refs are now covered via `type_expr_contract`, and `TThrow` in value positions is covered via `throw_expr_contract`) and keep dedicated coverage as each newly reachable node is supported. |
| Unsupported constant kind in `lowerConst` | Fatal: `Unsupported constant` | Add lowering + snapshot for any new constant kind encountered in real programs. |
| Unsupported `Std.isOfType` target kind | No compiler hard-fail for unresolved runtime-value abstract targets; falls back to conservative `false`/type-switch check | Keep adding explicit lowering for newly important target families and lock behavior with semantic diff coverage. |

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
- `haxe.go-6fc`: completed `haxe.ds` map/list core-ops semantic parity coverage and `List.push` parity alignment.
- `haxe.go-rlj`: completed nil-safe typed-read null semantics for `haxe.ds` map/list generic call results.
