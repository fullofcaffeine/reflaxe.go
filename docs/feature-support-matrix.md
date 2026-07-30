# Feature Support Matrix and Unsupported Inventory

This page is an implementation and evidence inventory. It is not the release
support authority. The authoritative release scope is the generated
[compatibility and support matrix](compatibility-support-matrix.md), backed by
the [machine manifest](compatibility-support-manifest.json). A `Supported` row
here proves the covered case; it does not admit every member, error path,
platform, or trust model.

How to read this page:

1. Start with the generated compatibility matrix to learn what the current
   release claim admits.
2. Read the evidence contract below to understand what `Supported` means in
   this implementation inventory.
3. Use the coverage tiers table and linked fixtures to inspect exact evidence.

Key terms:

- `snapshot`: checks generated Go code shape/text (`docs/snapshot-policy.md`).
- `semantic-diff`: compares runtime behavior of generated Go vs Haxe `--interp` (`docs/semantic-diff-guide.md`).
- `portable` / `metal`: compatible policy presets, documented in
  `docs/native-policy-presets.md`.

## Evidence support contract

`reflaxe.go` treats a surface area as **supported** when all of the following are true:

1. It has snapshot or harness coverage in `test/snapshot`.
2. It runs through `python3 test/run-ci.py` in CI (`snapshot` + `stdlib sweep` + `semantic diff` + `examples`).
3. Generated Go passes `go test ./...` for the covered case(s).

Anything outside that bar is either **partial** (implemented but not fully gated) or **unsupported**.

This label is intentionally evidence-scoped. Release admission requires a
separate explicit operation/member entry in
`compatibility-support-manifest.json`; unknown or unlisted surfaces are
excluded by default.

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
| `haxe.Http` / `sys.Http` | `semantic-diff` + direct race | `http_proxy_custom_request`, `http_request_callbacks_contract`, `http_multipart_streaming_contract`, `http_custom_request_lifecycle_contract`, `runtime/hxrt/http_test.go` |
| Direct `haxe.http.HttpBase` baseline plus direct `haxe.http.HttpMethod` / `haxe.http.HttpStatus` use | `semantic-diff` | `haxe_http_base_contract`, `stdlib/haxe_http_base_direct` |
| `sys.net.Socket` | `semantic-diff` + `snapshot` + direct race/cross-build | `socket_loopback_contract`, `socket_advanced_contract`, `sys/socket_input_service_surface`, `core/runtime_hxrt_infer_socket`, partial-I/O/peer-close/deadline cases in `runtime/hxrt/socket_test.go`, `test_socket_runtime_cross_build.py` |
| `haxe.crypto.Base64`, `Md5`, `Sha1`, `Sha224`, `Sha256` | `semantic-diff` + `snapshot` | `crypto_source_owned`, `crypto_xml_zip`, `stdlib/crypto_xml_zip_basic`, direct runtime crypto tests |
| root `Xml`, `haxe.xml.Parser`, `haxe.xml.Printer` | `semantic-diff` + `snapshot` | `xml_source_owned`, `root_xml_contract`, `crypto_xml_zip`, `stdlib/xml_root_dom_basic`, `stdlib/crypto_xml_zip_basic` |
| `haxe.zip.Compress`, `haxe.zip.Uncompress`, and `haxe.zip.Tools` one-shot and progressive partial-buffer paths | `semantic-diff` + snapshot/runtime + direct runtime/race | `zip_source_owned`, `zip_streaming_contract`, `crypto_xml_zip`, `stdlib/crypto_xml_zip_basic`, `stdlib/zip_streaming_policy`, direct runtime zip tests; exact `NO` / `SYNC` / `FINISH`, explicit unsupported errors for `FULL` / `BLOCK`; [streaming contract](haxe-zip-streaming.md) |
| root `Date` complete Haxe 4.3.7 API | `semantic-diff` + `snapshot` + direct runtime | `date_source_owned`, `option_date_path`, `datetools_cross_std_contract`, `stdlib/date_math_source_owned`, direct runtime date/timezone tests |
| root `Math` complete Haxe 4.3.7 API | `semantic-diff` + `snapshot` + direct runtime | `math_source_owned`, `numeric_edge_cases`, `stringtools_math`, `stdlib/date_math_source_owned`, `stdlib/math_float_native_no_hxrt`, direct runtime rounding tests |
| root `UnicodeString` complete Haxe 4.3.7 API | `semantic-diff` + `snapshot` + direct runtime | `unicode_string_source_owned`, `stdlib/unicode_string_basic`, direct runtime rune-slice tests |
| `haxe.Json` | `semantic-diff` | `json_parse_stringify_contract`, `stdlib/json_parse_stringify` |
| `haxe.io.Bytes` / `haxe.io.BytesBuffer` / `haxe.io.BytesInput` / `haxe.io.BytesOutput` (core ops + Input/Output helper subset) | `semantic-diff` | `bytes_normalization_contract`, `bytes_ops_contract`, `bytes_of_data_contract`, `bytes_hex_contract`, `bytes_io_stream_contract`, `io_input_output_helpers_contract`, `io_input_output_edge_contract`, `stdlib/bytes_basic` |
| `haxe.io` typed arrays (`ArrayBufferView`, `UInt8Array`, `UInt16Array`, `UInt32Array`, `Int32Array`, `Float32Array`, `Float64Array`) | `semantic-diff` | `haxe_io_typed_arrays_contract`, `stdlib/haxe_io_typed_arrays_direct` |
| `sys.io.Process` | `semantic-diff` + `snapshot` | `process_echo_contract`, `process_error_semantics_contract`, `sys/process_echo_smoke`, `sys/process_error_semantics` |
| `Sys.command` / `Sys.exit` wrapper delegation | `semantic-diff` + `snapshot` | `sys_command_contract`, `sys/sys_command_exit_wrapper` |
| `sys.io.File` | `semantic-diff` + `snapshot` | `file_read_write_contract`, `file_error_semantics_contract`, `sys/file_read_write_smoke`, `sys/file_error_semantics` |
| `sys.FileSystem` | `semantic-diff` | `filesystem_contract`, `sys/filesystem_basic_smoke` |
| `sys.net.Address` | `semantic-diff` | `sys_net_address_ssl_digest_algorithm_contract`, `stdlib/sys_net_address_ssl_digest_algorithm_direct` |
| `sys.net.UdpSocket` | `snapshot` + direct race/cross-build | `stdlib/sys_net_udp_socket_direct`, `runtime/hxrt/socket_test.go`, `test_socket_runtime_cross_build.py` |
| `Xml` (root DOM subset: document/element creation, attributes, child iteration, parse/print baseline) | `semantic-diff` | `root_xml_contract`, `stdlib/xml_root_dom_basic` |
| `haxe.ds.*Map` + `haxe.ds.List` (core ops subset) | `semantic-diff` | `ds_maps_list_contract`, `stdlib/ds_maps_list_basic` |
| `haxe.ds.WeakMap` (upstream platform contract: constructor throws `haxe.exceptions.NotImplementedException` on this target) | `semantic-diff` | `haxe_ds_weakmap_contract`, `stdlib/haxe_ds_weakmap_platform` |
| root `Reflect` complete Haxe 4.3.7 API | `semantic-diff` + `snapshot` + direct runtime | `reflect_compare`, `reflect_field_ops`, `reflect_extended_contract`, `stdlib/dynamic_access_basic`, direct runtime Reflect tests |
| `Date` + `haxe.ds.Option` + `haxe.io.Path` subset | `semantic-diff` | `option_date_path`, `path_cross_std_contract`, `stdlib/date_path_basic`, `stdlib/path_cross_std_basic` |
| Complete Haxe 4.3.7 `haxe.Log` (`formatOutput`, position/custom-parameter formatting, compiler-injected trace arguments, mutable rebinding, direct function values, catchable null invocation, restoration, and `Sys.println`) plus direct `haxe.Resource` / `haxe.SysTools` helpers | `semantic-diff` + snapshot runtime | `direct_haxe_helpers_contract`, `stdlib/std_log_source_owned`, `direct_haxe_resource_contract`, `stdlib/haxe_resource_embedded_basic` |
| `haxe.Utf8` deprecated helper subset (buffer ctor with or without optional size hint, `addChar`, `toString`, `iter`, `charCodeAt`, `validate`, byte-length `length`, byte-compare `compare`, UTF-8 character-position `sub`, `encode`, `decode`) | `semantic-diff` | `haxe_utf8_contract`, `stdlib/haxe_utf8_basic` |
| `haxe.Ucs2` platform exclusion | `snapshot` | `stdlib/haxe_ucs2_platform_exclusion` |
| `haxe.http.HttpJs` / `haxe.http.HttpNodeJs` target-conditional exclusion on Go | `snapshot` | `negative/direct_haxe_httpjs_unsupported`, `negative/direct_haxe_httpnodejs_unsupported` |
| `Array` + `IntIterator` core subset (`push`/`pop` statement-form, length/index access, array iteration, `0...N` range iteration) | `semantic-diff` | `array_string_intiterator_contract` |
| `String` core subset (`length`, `charAt`, `charCodeAt` with null on out-of-range, `substring`, `fromCharCode`, `toLowerCase`, `toUpperCase`) | `semantic-diff` + `snapshot` | `array_string_intiterator_contract`, `string_charcodeat_bounds_contract`, `sys/socket_input_service_surface` |
| `StringBuf` + `DateTools` (`add`/`addChar`/`addSub`, `DateTools.delta` with duration helpers, staged-std `DateTools.format` strftime subset plus `getMonthDays`/`parse`/`make`) | `semantic-diff` | `stringbuf_datetools_lambda_contract`, `datetools_cross_std_contract` |
| Complete Haxe 4.3.7 `Lambda` API: `array`, `list`, `map`, `mapi`, `flatten`, `flatMap`, `has`, `exists`, `foreach`, `iter`, `filter`, `fold`, `foldi`, `count`, `empty`, `indexOf`, `find`, `findIndex`, and `concat`. Every input-taking helper accepts `Array<T>`, `haxe.ds.List<T>`, and a concrete manual `Iterable<T>` carrier; callbacks and generic/nullable results are restored to their typed Haxe contract. Existing function-value coverage includes `map` and `fold`. | `semantic-diff` | `lambda_full_api_contract`, `lambda_list_contract`, `lambda_generic_iterable_count_empty_contract`, `lambda_iter_array_list_contract`, `lambda_iter_generic_iterable_contract` |
| `Lambda.flatten` / `Lambda.flatMap` nested-carrier boundary: nested arrays, staged lists, and concrete manual iterables are supported, as are callbacks returning concrete arrays or lists. A value already erased to structural `Iterable<Iterable<T>>`, or a callback already erased to `T -> Iterable<U>`, is rejected during Haxe compilation because Go cannot recover the missing nominal carrier after erasure. | `semantic-diff` + negative snapshot | `lambda_full_api_contract`, `negative/lambda_flatten_erased_nested`, `negative/lambda_flat_map_erased_result` |
| Core `Class`/`Enum`/`EnumValue` type-value subset | `semantic-diff` | `type_expr_contract` |
| `Type` reflection subset (`getClass`, `getSuperClass`, `getClassFields`, `getInstanceFields`, `resolveClass`, `createInstance`, `createEmptyInstance`, `getEnum`, `getEnumConstructs`, `resolveEnum`, `allEnums`, enum constructor/index/parameters/equality) | `semantic-diff` | `type_reflection_contract`, `type_reflection_extended_contract` |
| `haxe.ds.Vector` | `semantic-diff` | `vector_contract`, `stdlib/vector_basic` |
| `haxe.ds.ReadOnlyArray` (length/index/read-only view subset) | `semantic-diff` | `readonly_array_contract` |
| `sys.net.Host` | `semantic-diff` | `host_basic_contract`, `sys/host_basic_smoke` |
| `sys.ssl.Certificate` / `sys.ssl.Digest` / `sys.ssl.Key` | `snapshot` | `stdlib/sys_ssl_leaf_direct` |
| `sys.ssl.Socket` | `snapshot runtime` + direct race | `stdlib/sys_ssl_socket_direct`, `stdlib/sys_ssl_socket_sni_direct`, `core/runtime_hxrt_infer_socket_ssl`, stalled-handshake timeout coverage in `runtime/hxrt/socket_test.go` |
| `sys.ssl.DigestAlgorithm` | `semantic-diff` | `sys_net_address_ssl_digest_algorithm_contract`, `stdlib/sys_net_address_ssl_digest_algorithm_direct` |
| `haxe.PosInfos` | `semantic-diff` | `posinfos_contract`, `stdlib/posinfos_basic` |
| `haxe.Int32` | `semantic-diff` | `int32_contract` |
| `haxe.Int64` / `haxe.Int64Helper` | `semantic-diff` | `int64_contract`, `stdlib/int64_parity` |
| Complete Haxe 4.3.7 `Std` API, including parsing prefixes/signs/whitespace/overflow, float-prefix parsing, truncation, downcast aliases, type aliases, and random bounds | `semantic-diff` + `snapshot` | `std_complete_api_contract`, `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract`, `core/std_is_of_type_basic`, `core/std_is_of_type_dynamic`, `sys/socket_input_service_surface` |
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
| Inheritance and override dispatch | Supported for normally constructed generated hierarchies, including deep upcasts | `core/inheritance_override_dispatch`, `core/inheritance_ctor_chain_upcast`, `core/inheritance_return_upcast`, `core/inheritance_self_dispatch_wiring`, `core/deep_inheritance_dispatch_rebinding` |
| Dynamic lookup and mutation of emitted generated fields/methods through `Reflect` | Supported through selective typed same-package metadata; generated members remain lowercase in Go output | `reflect_extended_contract`, `haxe_template_concrete_iterable_contract`, `stdlib/dynamic_access_basic`, `stdlib/haxe_template_generated_method_lookup` |
| Interface dispatch | Supported | `core/interface_dispatch_basic` |
| Super calls | Supported | `core/super_calls` |
| Enums and switch pattern bindings | Supported | `core/enum_constructors`, `core/switch_enum_basic`, `core/enum_switch_bindings` |
| Anonymous object literals and structural field mutation | Supported | `core/object_literal_fields` |
| Exception subset (`throw`, typed/dynamic catch, rethrow, throw-as-expression, return-forwarding in statement-form try/catch, `haxe.Exception` API mapping, Haxe-catchable portable runtime validation failures, and carrier-only separation from native Go panics) | Supported | `core/haxe_exception_subset`, `core/try_catch_typed`, `core/try_catch_dynamic`, `core/try_catch_rethrow`, `core/try_catch_return_forwarding`, `core/portable_runtime_failure_haxe_catch`, `throw_expr_contract`, `try_catch_return_forwarding_contract`, `exception_api_contract`, `go_native/native_panic_not_haxe_catch` |
| `Std.isOfType` behavior | Supported | `core/std_is_of_type_basic`, `core/std_is_of_type_dynamic`, `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract` |
| Type-value expressions (`TTypeExpr`) for class/enum refs | Supported | `type_expr_contract` |
| `Type` reflection subset (`getClass`, `getSuperClass`, `getClassFields`, `getInstanceFields`, `resolveClass`, `createInstance`, `createEmptyInstance`, `getEnum`, `getEnumConstructs`, `resolveEnum`, `createEnum`, `createEnumIndex`, `allEnums`, `enumConstructor`, `enumIndex`, `enumParameters`, `enumEq`) | Supported | `type_reflection_contract`, `type_reflection_extended_contract` |
| Unsigned right shift behavior | Supported | `core/unsigned_shift`, `core/unsigned_shift_assign` |
| Naming/mangling and deterministic code shape | Supported | `core/naming_mangling`, `core/optimized_ast_policy` |
| HXML define/include resolution | Supported | `core/nested_hxml_define_detection`, `core/nested_hxml_long_define_detection`, `core/nested_hxml_quoted_define_detection`, `core/nested_hxml_root_relative_include_detection` |

### Semantic diff fixture coverage

- `test/semantic_diff/date_source_owned`
- `test/semantic_diff/math_source_owned`
- `test/semantic_diff/crypto_xml_zip`
- `test/semantic_diff/zip_source_owned`
- `test/semantic_diff/zip_streaming_contract`
- `test/semantic_diff/http_proxy_custom_request`
- `test/semantic_diff/http_request_callbacks_contract`
- `test/semantic_diff/socket_loopback_contract`
- `test/semantic_diff/socket_advanced_contract`
- `test/semantic_diff/null_string_concat`
- `test/semantic_diff/non_string_null_equality_contract`
- `test/semantic_diff/exceptions_typed_dynamic`
- `test/semantic_diff/exception_api_contract`
- `test/semantic_diff/haxe_http_base_contract`
- `test/semantic_diff/enum_switch_bindings`
- `test/semantic_diff/virtual_dispatch`
- `test/semantic_diff/deep_inheritance_dispatch_rebinding`
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
- `test/semantic_diff/haxe_template_concrete_iterable_contract`
- `test/semantic_diff/throw_expr_contract`
- `test/semantic_diff/try_catch_return_forwarding_contract`
- `test/semantic_diff/atomic_int_bool_contract`
- `test/semantic_diff/atomic_object_contract`
- `test/semantic_diff/go_result_contract`
- `test/semantic_diff/go_chan_nonblocking_contract`
- `test/semantic_diff/go_select_contract`

## Policy preset and native-boundary matrix

| Surface | Status | Evidence (snapshot IDs) |
| --- | --- | --- |
| `portable` safe devirtualization path | Supported | `core/portable_leaf_virtual_devirtualization`, `core/portable_leaf_virtual_alias_devirtualization`, `core/portable_leaf_virtual_inline_ctor_devirtualization`, `core/portable_leaf_virtual_function_return_devirtualization`, `core/portable_non_leaf_virtual_dispatch_preserved` |
| `portable` string helper optimizations | Supported | `core/portable_string_ptr_helpers`, `core/portable_string_literal_folding` |
| Compatibility selector enforcement | Supported | `negative/profile_conflict`, `negative/profile_invalid` |
| Native policy axes, precedence, and invalid/conflicting inputs | Supported | `core/report_artifacts_native_policy_overrides`, `core/report_artifacts_metal_proven_override`, `core/report_artifacts_lane_fallback_portable_surfaces`, `negative/native_authority_invalid`, `negative/native_specialization_invalid`, `negative/native_fallback_invalid`, `negative/native_fallback_conflict`, `negative/native_authority_guarded_metal`, `negative/native_fallback_error_portable` |
| Canonical `@:goNative` boundary plus `@:goMetal` compatibility | Supported | `core/native_boundary_guarded_authority`, `negative/go_metal_lane_injection`, existing `go_metal_lane_*` fixtures |
| Strict examples/app boundary policy + guarded native-usage policy modes | Supported | `negative/strict_examples_injection`, `negative/strict_mode_injection`, `negative/metal_profile_injection`, `negative/go_metal_lane_injection`, `negative/portable_native_import_error`, `core/portable_native_import_warn_policy` |
| RawNative encoding policy define (`reflaxe_go_raw_native_mode`) | Supported | `core/raw_native_utf16_mode`, `negative/raw_native_mode_invalid` |

## Go-native abstraction matrix

| Surface | Status | Evidence (snapshot IDs) |
| --- | --- | --- |
| Channels and goroutines | Supported (real goroutine/channel/select lowering; closed/empty receive is comma-ok aware; send-after-close and double-close preserve native panics; `go.Go.spawn` retains native fatal-panic and non-joined shutdown behavior while a feature-gated scope releases lazily-created portable thread identity/TLS state; typed deterministic `go.Select` helpers are available; concrete typed shims are selected by eager specialization or the proven concurrency fastpath, not by a separate semantic backend) | `go_native/channel_basic`, `go_native/channel_try_recv`, generated `go-channel-runtime` tooling gate, `go_native/channel_select_handshake`, `go_native/channel_metal_monomorph`, `core/native_boundary_guarded_authority`, `core/report_artifacts_metal_proven_override`, `go_native/goroutine_smoke`, `go_native/goroutine_native_panic`, `go_native/goroutine_native_shutdown`, `go_native/select_helpers`, `go_native/select_metal_monomorph` |
| Extern metadata mapping | Supported (`@:go.import`/`@:go.name`/`@:go.receiver`, extern `String` return normalization via `hxrt.StdString`, `@:go.valueError` mapping for `(T,error)` extern calls to `go.Result<T>`, and `@:go.tupleReturn` mapping for generated multi-return carrier classes) | `go_native/extern_metadata_mapping`, `go_native/extern_value_error_result`, `go_native/extern_tuple_return` |
| Result/Error mapping | Supported (eager or proven specialization adds typed `go.Result<T>` shim lowering with internal `(T,error)` helper emission) | `go_native/result_basic`, `go_native/error_result_mapping`, `go_native/result_metal_monomorph`, `core/report_artifacts_native_policy_overrides` |
| Slice/Map wrappers | Supported (eager or proven specialization adds typed shims for concrete `go.Slice<T>` and `go.Map<K,V>` call-sites) | `go_native/slice_map_basic`, `go_native/slice_map_metal_monomorph` |

## Stdlib matrix

Shim strategy and alternatives are documented in:

- `docs/stdlib-shim-rationale.md`
- `docs/stdlib-shim-migration-log.md`

### Snapshot-level behavioral coverage

- `stdlib/bytes_basic`
- `stdlib/crypto_xml_zip_basic`
- `stdlib/zip_streaming_policy`
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

- `haxe.Template` constructor/execute, nested lookup, array, structural-Dynamic, and concrete generated-class iteration, stack fallback, and macro invocation have direct parity coverage through `test/semantic_diff/haxe_template_contract`, `test/semantic_diff/haxe_template_concrete_iterable_contract`, `test/snapshot/stdlib/haxe_template_basic`, and `test/snapshot/stdlib/haxe_template_generated_method_lookup`. Template behavior is staged Haxe; three runtime representation operations cross a typed, footprint-explicit `hxrt` binding. Generic selective generated-method metadata supplies only already-bound lowercase methods to `Reflect.field` / `Reflect.hasField`.
- `haxe.ValueException` direct constructor/message/value parity now has semantic-diff coverage through `test/semantic_diff/haxe_value_exception_contract` and snapshot coverage in `test/snapshot/stdlib/haxe_value_exception_basic`.

### `Reflect` ownership and lookup contract

- `std/go/_std/Reflect.hx` owns the complete public Haxe 4.3.7 API, including lookup precedence, property access, copying, deletion, comparison, and varargs composition. This source contract is shared by the compatible `portable` and `metal` policy presets; no reflection behavior branches on the legacy profile name.
- `runtime/hxrt/reflect.go` owns only ordinary Go representation operations: map/exported-field and exported-method inspection, safe reflected calls and assignment, comparison, copying, deletion, object/function classification, and varargs adaptation. It is copied only when the reflection runtime feature is inferred.
- The compiler owns only closed-world facts that cannot be recovered from a separate Go package: class-token RTTI, generated lowercase fields and bound methods, and exact enum carriers. Typed semantic plans feed typed Go AST emitters for these adapters; there is no unsafe access, runtime type registry, or behavior-heavy `Reflect_*` compiler shim.
- Dynamic field lookup order is fixed and covered as: class-token RTTI, native map/exported field, generated Haxe field, generated Haxe method, then exported native Go method. `Type` metadata remains a separate exact intrinsic group and does not select ordinary `Reflect` runtime behavior.
- “Complete API” describes the 15 public Haxe 4.3.7 entrypoints, not an open-world promise for every possible Go carrier. The release manifest limits admission to the named generated-object, anonymous-object, native exported-member, comparison, and function cases with direct evidence; arbitrary external package-private members and untested function-identity edge cases remain excluded.

### `haxe.Json` runtime-lowered contract

- `haxe.Json.parse`/`haxe.Json.stringify`, `haxe.format.JsonPrinter.print`, and `haxe.format.JsonParser.doParse` now lower directly to `hxrt.JsonParse`/`hxrt.JsonStringify`.
- Compiler-emitted JSON shim declarations were removed from `GoCompiler` (no `haxe__Json`/`haxe__format__JsonParser` synthetic declarations in generated output).
- Runtime parity evidence: `test/semantic_diff/json_parse_stringify_contract`.

### `Sys` / `sys.io.File` / `sys.io.Process` ownership contract

- Native behavior is split by capability across `runtime/hxrt/sys.go`, `runtime/hxrt/file.go`, `runtime/hxrt/process.go`, and build-tagged `runtime/hxrt/terminal*.go`:
  - `hxrt.SysGetCwd`, `hxrt.SysArgs`, `hxrt.SysGetEnv`, typed environment entries/mutation, `hxrt.SysSleep`, cwd/path helpers, `hxrt.SysCommand`, `hxrt.SysExit`
  - typed `FileRead*` / `FileWrite*` / `FileOpen*`, opaque file handles, seek/tell/eof/flush/close capabilities, and non-owning standard-stream handles
  - `hxrt.NewProcess`; process stdin/stdout/stderr; byte I/O; PID, blocking/non-blocking exit status, kill, and close
  - footprint-explicit terminal character mode, one-byte input, and state restoration for `Sys.getChar`
- Canonical staged source in `std/go/_std/sys/io` owns `File`, `FileInput`, `FileOutput`, and `FileSeek`. It performs byte conversion, bounds checks, EOF construction, seek-origin selection, and stream construction; typed bindings in `std/hxrt/fs` expose only real native capabilities.
- Canonical staged `std/go/_std/Sys.hx` owns root `Sys`, including environment-map construction, aliases, fallbacks, standard-stream construction, `getChar` EOF construction, and echo policy. Canonical staged `std/go/_std/sys/io/Process.hx` likewise owns the public child-process API and stream semantics. Narrow typed bindings in `std/hxrt/sys`, `std/hxrt/fs`, and `std/hxrt/process` cross into genuine native capabilities; neither surface retains compiler-generated declarations.
- Portable `File.getContent` and `File.saveContent` propagate missing-path, permission, directory, and other OS failures through Haxe exceptions; an error is never converted to empty content or apparent success.
- Process startup and pipe failures throw instead of returning partial objects. Normal EOF is represented by `haxe.io.Eof`, non-EOF read failures remain errors, nonzero child exits remain ordinary exit codes, and `close()` releases/reaps without implicitly killing the child.
- Portable `Sys.putEnv` deliberately discards the runtime error to match Haxe 4.3.7 eval's `Void`, non-throwing contract. `hxrt.SysPutEnv` still returns the native `os.Setenv`/`os.Unsetenv` error so Go-native facades can preserve it.
- Portable `Sys.sleep(seconds)` blocks through an inline staged call to runtime-owned `hxrt.SysSleep`; non-positive and NaN durations return immediately. Its bounded timing contract avoids exact scheduler assertions.
- The root surface is explicit staged source rather than symbol-dependent compiler output: simple one-step methods inline to typed bindings, first-class references materialize source-owned functions, `setTimeLocale` reports `false`, deprecated `executablePath` aliases `programPath`, and `Sys.cpuTime` fails during Haxe compilation because Go has no portable process CPU clock. See [Portable root `Sys` contract](portable-sys-contract.md).
- Standard-stream carriers are non-owning: `close()` cannot close the process descriptor, and stdout/stderr flush avoids `os.File.Sync` failures on pipes and terminals. Ordinary staged `sys.io.File` carriers remain owning and sync-on-flush.
- On admitted `linux-amd64`, `Sys.getChar` reads immediately from a terminal with canonical input and host echo temporarily disabled, restores terminal state on success or failure, preserves redirected byte-stream EOF, and performs requested echo exactly once in staged Haxe. The terminal slice is absent from output that does not use `getChar`; `test/test_sys_get_char_terminal.py` executes the real PTY transition through a `checkptr=2`-instrumented binary and also covers redirected input plus Linux/macOS/Windows/unsupported-host cross-builds.
- No Process-, root-Sys-, or File-specific declaration emitter remains in `GoCompiler`. File evidence covers text and arbitrary binary content, copy, write/append/update modes, seek/tell, bounds, EOF, and OS failures through `sys/file_read_write_smoke`, `test/semantic_diff/file_read_write_contract`, `sys/file_error_semantics`, `test/semantic_diff/file_error_semantics_contract`, `test/semantic_diff/sys_db_io_contract`, and `stdlib/sys_db_io_direct`. Root/process evidence remains in `sys/process_echo_smoke`, `test/semantic_diff/process_echo_contract`, `sys/process_error_semantics`, `test/semantic_diff/process_error_semantics_contract`, `test/semantic_diff/root_sys_contract`, `test/semantic_diff/sys_command_contract`, `sys/sys_command_exit_wrapper`, `test/semantic_diff/sys_sleep_contract`, `sys/sys_sleep_portable`, `test/semantic_diff/root_sys_portable_contract`, `sys/root_sys_portable`, `core/runtime_hxrt_infer_sys`, `core/runtime_hxrt_infer_process`, and `negative/sys_cpu_time_unsupported`.

### `sys.thread` failure and shutdown contract

- `sys.thread.Thread.create` and `createWithEventLoop` create portable foreground
  workers. Generated `main` waits for all such workers and their nested portable
  workers before returning.
- An uncaught Haxe throw writes `Uncaught exception <message>` to stderr and ends
  only that worker, matching the observed Haxe 4.3.7 interpreter contract.
- A foreign Go panic is never accepted by Haxe `try`/`catch` or by the portable
  worker reporter; it remains fatal native behavior.
- `go.Go.spawn` is outside the foreground count and keeps normal Go shutdown
  semantics. When `sys.thread` is reachable, its compiler-owned callback scope
  releases only logical identity and TLS state on return or native-panic unwind;
  it does not join or recover. Evidence: `stdlib/sys_thread_uncaught_exception`,
  `go_native/native_panic_not_haxe_catch`, `go_native/goroutine_native_panic`,
  `go_native/goroutine_native_shutdown`, and the direct `runtime/hxrt` tests.
- Fixed and elastic pool admission is linearized with shutdown: every `run`
  either returns after durable exactly-once ownership or throws deterministic
  rejection. Generated race evidence covers 10,000 submissions at
  `GOMAXPROCS=1,2,8`, plus worker replacement after a task throws, with no
  retries.
- Condition wakeups are per waiter rather than global credits, event-loop
  ownership is synchronized, timed waits recompute after insert/cancel wakeups,
  and repeating-event cancellation state returns to baseline after 100,000
  cancellations. `sys.thread.Tls` stores values in the owning `ThreadState`, and
  portable/detached lifecycle churn must return identities and TLS payloads to
  baseline. Arbitrary foreign goroutines still have no automatic detach hook. See
  [`docs/concurrency-contract.md`](concurrency-contract.md).

### `sys.FileSystem` source/runtime contract and tradeoffs

- Canonical staged source in `std/go/_std/sys/FileSystem.hx` owns the complete Haxe 4.3.7 API: `exists`, `rename`, `stat`, `fullPath`, `absolutePath`, `isDirectory`, `createDirectory`, `deleteFile`, `deleteDirectory`, and `readDirectory`.
- Typed bindings under `std/hxrt/fs` delegate only native filesystem capabilities to the selectively copied `runtime/hxrt/filesystem.go`; `GoCompiler` has no filesystem shim group, synthetic `sys__FileSystem` declaration, or filesystem imports.
- Coverage includes `sys/filesystem_basic_smoke`, `test/semantic_diff/filesystem_contract`, and direct `runtime/hxrt` tests for deterministic create/read/rename/delete/stat-size behavior, canonical existing paths, and absolute paths that need not exist.
- Current tradeoff: `stat` maps metadata not exposed portably by Go's `os.FileInfo` (`uid/gid/dev/ino/rdev`, and platform-specific time distinctions) to stable fallback values behind the typed carrier. The staged Haxe source still constructs the unchanged upstream `sys.FileStat` record.

### `sys.io.File*` source/runtime contract and tradeoffs

- Canonical staged source in `std/go/_std/sys/io` owns the complete Haxe 4.3.7 `File` static API, both stream subclasses, and `FileSeek`; the mainstream extern declarations cannot supply a target implementation unchanged.
- Typed `std/hxrt/fs` bindings cross only opaque handles, strings, integers, booleans, and `Array<Int>` byte values. `runtime/hxrt/file.go` owns OS resources and native failures without depending on generated `haxe.io.Bytes` internals.
- The compiler initialization retains exact public `haxe.io.Input`, `Output`, and `Bytes` fields used by staged subclasses. This is type-only late-staging support, not a File behavior shim; no File-specific lowering remains.
- Current tradeoff: bulk byte operations copy through `Array<Int>` at the typed runtime boundary. That keeps ownership and representation honest; a future optimization must preserve the same source contract and prove its benefit with the runtime/perf gates before introducing a narrower representation bridge.

### `haxe.ds.*Map` / `haxe.ds.List` shim contract and tradeoffs

- Coverage includes `stdlib/ds_maps_list_basic` and `test/semantic_diff/ds_maps_list_contract` for deterministic `set/get/exists/remove` behavior across `StringMap`/`IntMap`/`ObjectMap`/`EnumValueMap`, plus `List` `add`/`push`/`pop`/`first`/`last`/`length`.
- `List.push` now prepends to match Haxe semantics (with `pop` removing the list head).
- Missing-key map reads and empty `List` reads in typed call sites now lower through nil-safe typed assertions, returning typed zero values (`null` for reference-like types) instead of panicking.

### `haxe.io.BytesInput` / `haxe.io.BytesOutput` source/runtime contract and tradeoffs

- Coverage includes `test/semantic_diff/bytes_io_stream_contract`, `test/semantic_diff/bytes_of_data_contract`, `test/semantic_diff/bytes_hex_contract`, `test/semantic_diff/io_input_output_helpers_contract`, `test/semantic_diff/io_input_output_edge_contract`, `test/semantic_diff/io_error_constructor_contract`, `test/semantic_diff/io_encoding_contract`, and `test/semantic_diff/haxe_io_misc_contract` for deterministic constructor bounds checks, `position`/`length`, EOF behavior, `readByte`/`readBytes`, inherited helper subset parity (`readAll`, `readFullBytes`, `read`, `readUntil`, `readLine`, `readString`, `readFloat`/`readDouble`, signed/unsigned numeric reads), output helper subset parity (`write`, `writeFullBytes`, `writeInput`, `writeString`, numeric typed writes, overflow guards), direct `StringInput` / `BufferInput` constructor+read behavior, `haxe.io.Error` typed constructor matching (`Blocked`, `Overflow`, `OutsideBounds`, `Custom`), `haxe.io.Encoding` constructor parity (`UTF8`, `RawNative`), `Bytes.getString` bounds behavior, `Bytes.getData`/`Bytes.ofData` alias semantics, direct `Mime` / `Scheme` abstract usage, `FPHelper` bit conversions, `Bytes.toHex`/`Bytes.ofHex` behavior, and `readLine` EOF/tail/CRLF edge paths. Snapshot evidence for the direct tranche lives in `test/snapshot/stdlib/haxe_io_misc_direct`.
- Canonical staged modules under `std/go/_std/haxe/io` own the public hierarchy,
  stream algorithms, validation, endian behavior, EOF/error policy, alias
  observation, and RawNative cache invalidation. Typed `std/hxrt/io` bindings
  cross only opaque byte views, native conversion/copy/UTF operations, and
  scalar IEEE-754 reinterpretation. No compiler `io` group remains.
- Current tradeoff: parity remains focused on `BytesInput`/`BytesOutput` stream behavior with interpreter-compatible semantics by default (`reflaxe_go_raw_native_mode=interp`, where `UTF8` and `RawNative` both map to UTF-8 conversion). For projects that need Java/C#-style RawNative byte layout, `reflaxe_go_raw_native_mode=utf16le` provides an explicit opt-in UTF-16LE path; full target-by-target RawNative equivalence is still not claimed outside these documented modes.

### `haxe.io` typed-array contract and tradeoffs

- Coverage includes `test/semantic_diff/haxe_io_typed_arrays_contract` and `test/snapshot/stdlib/haxe_io_typed_arrays_direct` for direct `ArrayBufferView`, `UInt8Array`, `UInt16Array`, `UInt32Array`, `Int32Array`, `Float32Array`, and `Float64Array` usage, including `fromBytes`, `fromArray`, direct construction, indexing, `sub`, `subarray`, aliasing through `Bytes`, and bounds errors.
- Public typed-array behavior now lives in ordinary staged source under `std/go/_std/haxe/io/*.hx`, so the Haxe-facing API stays library-owned instead of growing more raw compiler-owned bytes logic.
- Current tradeoff: typed arrays use ordinary generated abstract representation
  mapping over staged `Bytes`; native byte and float-bit facts cross the shared
  typed `ByteView` / `NativeFloatBits` capabilities. There is no IO compiler shim.

### `haxe.Http` / `sys.Http` shim contract and tradeoffs

- `haxe.Http` is a `typedef` alias of `sys.Http` on `sys` targets, so the same semantic-diff fixtures now serve as the portable contract for both entry points.
- `sys.Http` is canonical staged source in `std/go/_std/sys/Http.hx`. It owns synchronous request selection for `http`/`https`, deterministic `data:` handling, payload/header assembly, callback order, public response maps, and status/error policy.
- Go URL parsing, proxy setup, response resources, and optional typed socket consumption live behind opaque `std/hxrt/http` handles in footprint-explicit `runtime/hxrt/http.go`; no generated Haxe object layout crosses the runtime boundary and no compiler HTTP group remains.
- Covered behaviors: `setHeader`/`addHeader`, `setParameter`/`addParameter`, `setPostData`/`setPostBytes`, `fileTransfer`/`fileTransfert`, `customRequest` (including optional socket transport injection), proxy URL wiring (`Http.PROXY`), `getResponseHeaderValues`, dynamic callbacks (`onData`, `onBytes`, `onError`, `onStatus`), `responseData`/`responseBytes`, and `requestUrl`.
- Semantic diff now also locks callback/status/header/error parity for local deterministic HTTP servers (`http_request_callbacks_contract`), including 4xx `onError` formatting (`Http Error #<status>`).
- Multipart uploads pull bounded chunks from the caller's `Input`; partial reads
  are retried, the full file is not staged in memory, and early EOF or a source
  error aborts the exchange while preserving the source error
  (`http_multipart_streaming_contract` and direct runtime tests).
- Direct `customRequest` proves only the final state of small complete responses:
  it writes and closes the supplied `Output` on success, does not also publish
  `responseData` or fire `onData`/`onBytes`, and leaves the output open after an
  HTTP-status error (`http_custom_request_lifecycle_contract`). Native
  execution currently buffers the complete response first, so status and body
  writes are not streamed and partial bytes are lost when body reading later
  fails.
- Direct runtime tests cover the current GET/POST/query/body representation,
  timeout cancellation, truncated-body status retention, proxy formatting,
  early-upload abort, idle transport cleanup, and typed custom-socket closure.
  They do not establish Haxe parity for ordered repeated parameters/headers,
  query/form separation, per-request multipart boundaries, generated-Haxe
  callback context, or cancellation of a blocked upload source.
- HTTP is release-excluded under `haxe_go-vfp.10.8`. The required typed design
  is a source-driven upload sink plus a bounded native response-exchange handle;
  see the generated compatibility matrix and the
  [independent disposition](reviews/network-admission-oracle-disposition-vfp-10.4.md).

### `sys.net.Socket` staged-source contract and tradeoffs

- Deterministic loopback fixtures exercise `bind`/`listen`/`connect`/`accept`/`read`/`write`/`close` plus advanced methods such as `setTimeout`, `waitForRead`, `setBlocking`, `setFastSend`, `select`, and `shutdown` (`socket_loopback_contract`, `socket_advanced_contract`). Fixture execution is evidence for those cases, not blanket API parity.
- `sys.net.Socket.input` now satisfies the generated `haxe.io.Input` stream contract for service-style code paths (`readByte`, `readBytes`, `readAll`, typed numeric/string helper forwarding, and endian control). The focused snapshot is `sys/socket_input_service_surface`.
- `select` preserves source object identity and tests read readiness, but its
  write set currently reports connection presence rather than operating-system
  writability, and its exceptional set is not a real exceptional-condition
  probe. Advanced readiness remains release-excluded.
- The public API, stream wrappers, Haxe exceptions, address construction, and select object identity are canonical staged Haxe. A typed opaque `SocketHandle` and concrete result carriers cross into footprint-explicit `runtime/hxrt/socket.go`; the former `net_socket` compiler group and `GoNetSocketEmitter` are gone.
- Direct runtime tests cover TCP/UDP round trips, partial-write progress,
  accept/datagram deadlines, peer close after a partial read, explicit
  timeout/readiness state, idempotent concurrent close, close-unblocks-read,
  and closed-handle `waitForRead` under the Go race detector.
- Current tradeoffs: `setBlocking` uses deadline behavior rather than true
  nonblocking file descriptors; `bind` starts listening; `listen` does not
  apply backlog; eager DNS is outside socket timeouts. The admitted operation
  list is only the Linux/amd64 blocking IPv4 TCP client core. All other socket
  members are owned by `haxe_go-vfp.10.9`.

### `sys.net.UdpSocket` direct baseline and tradeoffs

- Direct UDP behavior has deterministic loopback snapshot/runtime coverage for
  `bind`, `host`, `sendTo`, `readFrom`, `setBroadcast`, and `Address`
  round-tripping (`stdlib/sys_net_udp_socket_direct`). A valid zero-byte
  datagram currently loses its sender identity, so UDP remains release-excluded.
- `haxe_go-vfp.8.7.14` moved the public API to canonical staged source over the shared typed TCP/UDP handle. No compiler-emitted UdpSocket implementation or raw-injection path remains.
- `setBroadcast(true)` maps to Go's operating-system socket option path (`SO_BROADCAST`) on the underlying UDP connection. Build-tagged POSIX and Windows helpers preserve each platform's native descriptor type, and a cross-build regression test keeps both compiling. The portable evidence checks that the option is installed and that normal UDP behavior still works; it does not require sending packets to a LAN broadcast address, because CI machines and developer laptops can block that at the network-policy level.

### `sys.ssl.Socket` staged TLS composition and tradeoffs

- Public verification/CA/hostname/certificate/SNI configuration lives in staged Haxe over the source-owned `sys.net.Socket` and its shared typed handle.
- `runtime/hxrt/socket_ssl.go` owns only native TLS client/listener installation, handshake, peer-certificate access, and synchronized SNI selection. Certificate/key/digest primitives remain in `ssl.go`; SSL leaf users therefore do not select network transport.
- Runnable snapshots prove covered TLS loopback I/O, peer certificate fields,
  accepted `sys.ssl.Socket` runtime identity, default-certificate selection,
  and callback-driven SNI selection. They do not prove the ordinary default
  path: the original logical hostname is currently lost when the resolved
  numeric address is passed to TLS unless callers explicitly set a hostname.
  Selective snapshots prove only the `ssl` leaf versus
  `socket + ssl + socket_ssl` transport split.
- A direct race test also proves `Socket.setTimeout` bounds a TLS client when a
  peer accepts TCP but never completes its handshake. TCP and TLS now share
  the same snapshotted dial policy.
- TLS, HTTPS, inherited shutdown/fast-send controls, DNS identity, and runtime
  behavior outside Linux/amd64 remain release-excluded under
  `haxe_go-vfp.10.9`.

### `EReg` + `haxe.Serializer` contract and tradeoffs

- `std/go/_std/EReg.hx`, `std/go/_std/haxe/Serializer.hx`, and `Unserializer.hx` are now the canonical public implementations. Match state, group validation, global policy, token selection/parsing, caches, recursive traversal, resolver policy, and custom-hook sequencing are ordinary staged Haxe behavior; the retired `regex_serializer` compiler group no longer emits them.
- Typed `std/hxrt/regex` bindings select `runtime/hxrt/regex.go`, which owns only compiled RE2 resources, native matching/quoting, and conversion from Go UTF-8 byte indexes to Haxe code-point indexes. Staged `EReg` expands Haxe replacement templates itself so `$1x` cannot be misread as Go's named-capture syntax. `core/runtime_hxrt_infer_regex` proves this slice does not copy serialization support.
- Typed `std/hxrt/serialization` bindings select `runtime/hxrt/serialization.go`, which now owns only bounded host float parsing. Serialization adds no serializer-specific reflection or unsafe boundary; its staged `Reflect` calls intentionally select the existing safe `runtime/hxrt/reflect.go` fallback for dynamic objects and calls.
- Field enumeration, private/inherited field access, assignment, custom hooks, and structural resolver calls reuse staged `Reflect`. Generated members use existing typed same-package field/method metadata; erased call invocation continues through the shared safe `Reflect.callMethod` boundary. Class/enum name resolution and construction reuse Type metadata, whose empty-instance helpers initialize all embedded superclass carriers and bind virtual dispatch to the final child. There is no serializer-specific bridge or metadata table.
- `serializer_typed_accessor_contract` proves private fields across three inheritance levels, `@:transient` parity, constructor-free virtual dispatch, custom hook round-trips, and integral wire tokens assigned to declared `Float` fields without `unsafe`.
- `core/runtime_hxrt_infer_serialization` proves the serialization slice, shared Reflect helper, and equality dependency are copied without `regex.go`. The serialization-specific selective-runtime perf case budgets their source and binary footprint. Runtime feature selection follows typed reachability, not the legacy `portable|metal` preset.
- `EReg` parity now covers: `g/i/m/s/u` option handling, global vs non-global `replace`/`map`, `matched`/`matchedPos`/`matchedLeft`/`matchedRight` error semantics, and group/null behavior via semantic diff fixtures (`ereg_behavior_contract`, `ereg_edge_contract`).
- `haxe.Serializer`/`haxe.Unserializer` now cover a wire-format-compatible baseline for core tokens used by fixtures (`n/t/f/z/i/d/k/p/m/v/s/y/a/o/l/b/q/M/c/w/j/C/x/A/B/g/u/h/r/R`) plus sequential `Unserializer` cursor behavior (`serializer_wire_contract`), resolver paths (`serializer_custom_resolver_contract`), resolver method-shape polymorphism (`serializer_resolver_polymorphism_contract`), cache/reference graph parity (`serializer_cache_reference_contract`), global serializer default flag behavior (`Serializer.USE_CACHE`/`Serializer.USE_ENUM_INDEX`) with `serializeException` interaction (`serializer_global_flags_contract`), and mixed string/object reference stress (`serializer_reference_stress_contract`).
- `serializer_boundary_matrix_contract` characterizes the remaining parent
  boundary directly: an interface-typed field containing a generic class
  round-trips with method dispatch intact, 64 deterministic generic values
  satisfy the same round-trip property, a cached generic self-cycle preserves
  identity, and malformed token, truncated string, and unknown-class inputs
  remain catchable errors.
- Current scope note: serializer/unserializer coverage is broad, but the fixtures still focus on deterministic, portable payload shapes rather than claiming every obscure cross-target edge combination.
- Completed follow-up evidence:
  - `haxe.go-7zy.10` (migrate `haxe.Json` shim out of compiler core, completed 2026-02-19)
  - `haxe.go-7zy.11` (move native `Sys`/file/process behavior out of compiler core, completed 2026-02-19)
  - `haxe_go-vfp.8.7.5` (retire File API/type ownership from the remaining Sys/Process shim, completed 2026-07-15)
  - `haxe_go-vfp.8.7.6` (retire root `Sys` API/type ownership from the remaining Process shim, completed 2026-07-15)
  - `haxe_go-vfp.8.7.7` (retire the final Process compiler emitter in favor of canonical staged source and typed runtime capabilities, completed 2026-07-15)
  - `haxe.go-7zy.12` (reduce `stdlib_symbols` bytes-conversion overhead, completed 2026-02-19)
  - `haxe.go-re8` (support resolver-returned type-value markers for class/enum name extraction + serialization, completed 2026-02-20)
  - `haxe_go-vfp.8.7.13` (retire the mixed regex/serializer emitter in favor of staged algorithms and exact typed boundaries, completed 2026-07-18)

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
sys.net.UdpSocket
```

### Full runtime-eligible inventory sweep

Source list: `test/upstream_std_modules_full.txt` (175 modules).

Current generated sweep contract:

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

These are explicit fatal guards in `src/reflaxe/go/GoCompiler.hx`.
The remaining inventory is invariant-only: parser/front-end rejection or node-family closure proof, not active supported-language holes.
As of **2026-03-24**, the hard-fail inventory count remains **4**, with named fixture strategies per path.

| Inventory item | Current behavior | Fixture strategy (named) | Invariant guard strategy | Owner |
| --- | --- | --- | --- | --- |
| Non-lvalue assignment targets in `lowerLValue` | Fatal: `Unsupported assignment target` | `negative/non_lvalue_assignment_invariant` locks Haxe front-end rejection (`Invalid assign`) so backend fatal remains an explicit invariant unless a typed reproducer becomes reachable. | Closed as an invariant: if a legal lvalue shape becomes reachable later, the same change must add lowering plus a focused regression fixture. | `haxe.go-14as.56` |
| Non-`++/--` postfix unary in `lowerExpr` / `lowerExprWithPrefix` | Fatal: `Unsupported postfix unary operator` | `negative/postfix_non_inc_dec_invariant` locks parser-level rejection (`Postfix ! is not supported`) so only `++/--` postfix forms can reach lowering today. | Closed as an invariant: parser/typed-AST assumptions stay release-checked, and any new reachable postfix form must land with lowering plus snapshots. | `haxe.go-14as.56` |
| Catch-all `lowerExpr` default | Fatal: `Unsupported expression` | Node-family closure map: `semantic-diff/type_expr_contract`, `semantic-diff/throw_expr_contract`, `core/untyped_ident_nil`, and `core/const_kinds_contract`; new reachable node families must add a dedicated contract fixture in the same change. | Closed as an invariant: known reachable node families are covered, and future node-family expansion must add explicit lowering plus semantic or snapshot evidence. | `haxe.go-14as.56` |
| Unsupported `Std.isOfType` target kind | No compiler hard-fail for unresolved runtime-value abstract targets; falls back to conservative `false`/type-switch check | `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract`, `core/std_is_of_type_basic`, `core/std_is_of_type_dynamic`, `core/type_switch_no_binding_std_is_of_type` lock current fallback semantics. | Closed as an invariant: current fallback semantics are contract-tested, and future target-family expansion must add explicit lowering plus semantic-diff coverage. | `haxe.go-14as.56` |

## Known non-semantic-diff stdlib surfaces

The broader probe list:

```bash
python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_gap_probe.txt --strict --go-test
```

reports:

- `53 passed / 0 expected policy / 0 failed / 0 unexpected present`

There are currently no active expected-policy rules in the full inventory, and no modules remain in `compile-only` status.

The portable parity closure summary is still useful after compile-only debt reaches zero. It tracks the smaller set of modules that are not semantic-diff surfaces yet:

- `snapshot`: implemented enough for generated Go/runtime smoke evidence, but not promoted to interpreter-vs-Go semantic-diff evidence.
- `unsupported`: explicit target-conditional exclusions, not hidden missing work.
- `closure_policy`: why the module is still non-semantic-diff, such as `target_sensitive_snapshot` or `explicit_exclusion`.
- `actionable`: whether the remaining item is real follow-up work (`true`) or an intentionally locked policy (`false`).

Generate the live list with:

```bash
python3 test/run-portable-parity-closure.py --list-blockers
```

## Completed Tracking History

All entries below are closed provenance. They explain how this matrix reached its
current state; they are not the active task queue. Use the generated inventory
and closure artifacts above for the current live status.

- `haxe.go-d5u`: published and maintained this matrix/inventory baseline.
- `haxe.go-61w`: reduced the compiler hard-fail unsupported expression surface.
- `haxe.go-19u`: expanded stdlib parity from the documented probe gap list.
- `haxe.go-ab2`: added the semantic differential regression harness.
- `haxe.go-3d4`: reduced unsupported expression surface by lowering `TTypeExpr` class/enum value nodes.
- `haxe.go-8zt`: lowered `TThrow` in expression positions and locked it with semantic diff coverage.
- `haxe.go-888`: promoted `sys.FileSystem` with deterministic snapshot + semantic parity contracts.
- `haxe.go-uz4.10`: enabled typed `go.Chan<T>` recv/recvOr assertions in `portable`.
- `haxe.go-6fc`: completed `haxe.ds` map/list core-ops semantic parity coverage and `List.push` parity alignment.
- `haxe.go-rlj`: completed nil-safe typed-read null semantics for `haxe.ds` map/list generic call results.
- `haxe.go-aiy`: completed `haxe.io.Encoding` constructor parity and `Bytes.getString` coverage (`io_encoding_contract`).
- `haxe.go-dq2`: evaluated and guarded RawNative compatibility policy with explicit mode controls.
- `haxe.go-rcv`: added the `haxe.io.Bytes.ofData` shim and locked `getData` alias semantics (`bytes_of_data_contract`).
- `haxe.go-nmg`: added `haxe.io.Bytes.toHex` / `haxe.io.Bytes.ofHex` shim parity (`bytes_hex_contract`).
- `haxe.go-9v6`: promoted `haxe.ds.ReadOnlyArray` from compile-only to semantic-diff coverage (`readonly_array_contract`).
- `haxe.go-8hs`: added serializer global default flag semantic coverage (`serializer_global_flags_contract`).
- `haxe.go-14as.6`: promoted iterator-family parity (`haxe.iterators.ArrayIterator`, `ArrayKeyValueIterator`, `DynamicAccessIterator`, `DynamicAccessKeyValueIterator`, `RestIterator`, `RestKeyValueIterator`, `StringIterator`, `StringIteratorUnicode`, `StringKeyValueIterator`, `StringKeyValueIteratorUnicode`) via `iterators_family_contract` and staged std overrides.
- `haxe.go-14as.7`: added portable-vs-metal invariance coverage for iterator/list/map/string portable surfaces (`portable_surfaces_metal_invariance_contract`, `portable_surfaces_lane_invariance_contract`) and snapshot fallback-report attribution lock (`core/report_artifacts_lane_fallback_portable_surfaces`).
- `haxe.go-14as.8`: closed unsupported-expression inventory with explicit invariant fixture strategy mapping (`negative/non_lvalue_assignment_invariant`, `negative/postfix_non_inc_dec_invariant`, plus node-family closure fixtures for `lowerExpr` and `Std.isOfType` fallback behavior).
- `haxe.go-14as.9`: added deterministic hxrt feature reason provenance to runtime reports (`core/report_artifacts_runtime_reason_provenance`, `test/run-auto-planner-schema.py`) without changing runtime semantics.
- `haxe.go-14as.10`: final closure sync task; inventory/closure artifacts now require explicit blocker metadata for remaining compile-only modules, and release readiness docs include portable parity + family sync gates.
- `haxe.go-14as.11`: root/core tranche triage closure. Direct semantic-diff contracts promoted `Any`, `StdTypes`, and `sys.FileStat`; remaining root blockers were split into dedicated tasks for `Sys`, `Xml`, and `UnicodeString`.
- `haxe.go-14as.20`: closed root `Sys` surface split. Added direct semantic-diff coverage in `root_sys_contract` and wired `Sys.getEnv`, `Sys.putEnv`, `Sys.environment`, and `Sys.systemName` through new hxrt/compiler shims.
- `haxe.go-14as.21`: closed root `Xml` surface split. Added direct semantic-diff coverage in `root_xml_contract` plus snapshot coverage in `stdlib/xml_root_dom_basic`.
- `haxe.go-14as.22`: closed root `UnicodeString` surface split. Added direct semantic-diff coverage in `root_unicode_string_contract` plus snapshot coverage in `stdlib/unicode_string_basic`, and wired the `_UnicodeString__UnicodeString_Impl__*` lowering surface through generated stdlib-symbol shims.
- `haxe.go-14as.24`: closed the remaining `Xml.parse()` parsed-CDATA node-type follow-up, so root `Xml` now preserves `CData` instead of collapsing it to `PCData`.
- `haxe.go-14as.12`: closed generic `haxe.misc` triage by promoting `haxe.Http` from existing semantic-diff fixtures and splitting the remaining modules into `haxe.go-14as.25` to `haxe.go-14as.29`.
- `haxe.go-14as.26`: closed `haxe.Constraints` + `haxe.Rest` abstraction lowering. Added direct semantic-diff coverage in `haxe_constraints_contract` and `haxe_rest_contract`, snapshot coverage in `stdlib/haxe_constraints_rest_direct`, staged std ownership for the `haxe.Constraints` bridge contract, and direct native lowering for `haxe.Rest` copy/append/prepend semantics.
- `haxe.go-14as.27`: closed `haxe.EnumFlags` + `haxe.EnumTools` enum-helper parity. Added direct semantic-diff coverage in `haxe_enum_helpers_contract` plus snapshot coverage in `stdlib/haxe_enum_helpers_direct`, fixed `haxe.EnumFlags` abstract lowering to preserve its `Int` backing, and required the stdlib impl class `haxe._EnumFlags.EnumFlags_Impl_` so upstream helper code can compile without a target-owned override.
- `haxe.go-14as.28`: closed the stack-fallback half of the old stack/main-loop triage. `haxe.CallStack` and `haxe.NativeStackTrace` remain under explicit target-sensitive snapshot coverage through `stdlib/haxe_stack_loop_target_sensitive`.
- `haxe.go-14as.40`: kept native Go stack capture out of the portable semantic baseline and documented the opt-in diagnostic design in `docs/spikes/native-stack-capture-contract.md`. `haxe.go-14as.76` implemented the `-D reflaxe_go_native_stack_trace` path with target-sensitive snapshot/runtime coverage in `stdlib/haxe_native_stack_trace_opt_in`.
- `haxe.go-14as.29`: closed the legacy text tranche. `haxe.Utf8` now uses staged std ownership with semantic-diff coverage (`haxe_utf8_contract`) and snapshot coverage (`stdlib/haxe_utf8_basic`), while `haxe.Ucs2` stays under explicit target-sensitive snapshot coverage (`stdlib/haxe_ucs2_platform_exclusion`).
- `haxe.go-14as.43`: closed direct `haxe.exceptions` subclass construction parity. `haxe.exceptions.PosException`, `haxe.exceptions.ArgumentException`, and `haxe.exceptions.NotImplementedException` now have semantic-diff coverage in `haxe_exceptions_direct_contract` plus snapshot coverage in `stdlib/haxe_exceptions_direct`.
- `haxe.go-14as.46`: closed the direct `haxe.ds.BalancedTree` / `haxe.ds.GenericStack` runtime tranche. The staged overrides now hold direct set/get/remove/toString/pop behavior under semantic-diff coverage in `haxe_ds_source_owned_collections_contract` plus snapshot coverage in `stdlib/haxe_ds_source_owned_collections`; broader iterator parity stays tracked separately.
- `haxe.go-14as.47`: closed the direct `haxe.ds.WeakMap` stance blocker. Go now preserves the upstream Haxe platform contract by allowing direct construction to compile and then throw `haxe.exceptions.NotImplementedException` at runtime; evidence lives in `haxe_ds_weakmap_contract` and `stdlib/haxe_ds_weakmap_platform`.
- `haxe.go-14as.14`: closed the direct `haxe.http` half of the old combined HTTP/RTTI blocker. `haxe.http.HttpBase` now has staged-stdlib baseline support through `std/go/_std/haxe/http/HttpBase.hx`, with parity evidence in `haxe_http_base_contract` and `stdlib/haxe_http_base_direct`. `haxe.http.HttpJs` and `haxe.http.HttpNodeJs` are now explicit unsupported target-conditional modules on Go, with negative coverage in `negative/direct_haxe_httpjs_unsupported` and `negative/direct_haxe_httpnodejs_unsupported`.
- `haxe.go-14as.57`: closed the direct `haxe.rtti.*` reflection tranche. `haxe.rtti.Meta`, `haxe.rtti.Rtti`, `haxe.rtti.CType`, and `haxe.rtti.XmlParser` now have staged-stdlib baseline support with mixed ownership: public parser/type logic lives in `std/go/_std/haxe/rtti/*.hx`, while the backend owns the narrow class-token `__meta__` / `__rtti` contract plus anonymous-record array-field mutation lowering underneath. Evidence: `haxe_rtti_direct_contract` and `stdlib/haxe_rtti_direct`.
- `haxe.go-14as.17`: closed the original direct `sys.db` + `sys.io` parity tranche through `sys_db_io_contract` and `stdlib/sys_db_io_direct`. Its interim compiler-owned File wrapper decision was superseded by `haxe_go-vfp.8.7.5`: `sys.io.File`, `FileInput`, `FileOutput`, and `FileSeek` are now canonical staged std over typed `runtime/hxrt/file.go` capabilities, while `sys.db` remains upstream/public source-owned.
- `haxe.go-14as.19` is now closed end-to-end. The first wave (`Condition`, `Deque`, `IThreadPool`, `Lock`, `Mutex`, `NoEventLoopException`, `Semaphore`, `ThreadPoolException`, `Tls`) stays covered by `sys_thread_primitives_contract` and `stdlib/sys_thread_primitives_direct`, and the second wave (`Thread`, `EventLoop`, `ElasticThreadPool`, `FixedThreadPool`) is now covered by `sys_thread_runtime_contract` and `stdlib/sys_thread_runtime_direct`.
- `haxe.go-14as.69`: direct `haxe.EntryPoint` / `haxe.MainLoop` / `haxe.Timer` usage now has staged-stdlib baseline support over the runtime-backed `sys.thread.EventLoop` contract. Evidence: `stdlib/haxe_main_loop_runtime_direct`. This stays snapshot/runtime coverage for now because asynchronous timing is target-sensitive rather than a clean semantic-diff comparison.
- No compile-only portable blocker families remain in the generated inventory. Use `test/portable_stdlib_inventory.json` plus `test/.test-cache/portable_parity_closure_summary.md` as the authoritative current list for snapshot-only and explicitly unsupported surfaces.
