# Feature Support Matrix and Unsupported Inventory

## Support contract

`reflaxe.go` treats a surface area as **supported** when all of the following are true:

1. It has snapshot or harness coverage in `test/snapshot`.
2. It runs through `python3 test/run-ci.py` in CI (`snapshot` + `stdlib sweep` + `semantic diff` + `examples`).
3. Generated Go passes `go test ./...` for the covered case(s).

Anything outside that bar is either **partial** (implemented but not fully gated) or **unsupported**.

## Language/Core matrix

| Surface | Status | Evidence (snapshot IDs) |
| --- | --- | --- |
| Arithmetic, comparisons, boolean flow | Supported | `core/arithmetic`, `core/if_else` |
| Locals, assignment, compound assignment | Supported | `core/locals_assign`, `core/compound_assign`, `core/compound_assign_string` |
| Arrays and indexing | Supported | `core/array_basic`, `core/array_push_pop` |
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
| Exception subset (`throw`, typed/dynamic catch, rethrow) | Supported | `core/haxe_exception_subset`, `core/try_catch_typed`, `core/try_catch_dynamic`, `core/try_catch_rethrow` |
| `Std.isOfType` behavior | Supported | `core/std_is_of_type_basic`, `core/std_is_of_type_dynamic` |
| Unsigned right shift behavior | Supported | `core/unsigned_shift`, `core/unsigned_shift_assign` |
| Naming/mangling and deterministic code shape | Supported | `core/naming_mangling`, `core/optimized_ast_policy` |
| HXML define/include resolution | Supported | `core/nested_hxml_define_detection`, `core/nested_hxml_long_define_detection`, `core/nested_hxml_quoted_define_detection`, `core/nested_hxml_root_relative_include_detection` |

### Semantic diff fixture coverage

- `test/semantic_diff/null_string_concat`
- `test/semantic_diff/exceptions_typed_dynamic`
- `test/semantic_diff/enum_switch_bindings`
- `test/semantic_diff/virtual_dispatch`
- `test/semantic_diff/stringtools_math`
- `test/semantic_diff/option_date_path`
- `test/semantic_diff/numeric_edge_cases`
- `test/semantic_diff/nullable_struct_refs`
- `test/semantic_diff/sys_io_roundtrip`
- `test/semantic_diff/reflect_compare`
- `test/semantic_diff/reflect_field_ops`
- `test/semantic_diff/anonymous_object_literals`

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

### Snapshot-level behavioral coverage

- `stdlib/bytes_basic`
- `stdlib/crypto_xml_zip_basic`
- `stdlib/date_path_basic`
- `stdlib/ds_maps_list_basic`
- `stdlib/intmap_basic`
- `stdlib/io_type_smoke`
- `stdlib/json_parse_stringify`
- `stdlib/math_basic`
- `stdlib/option_enum_basic`
- `stdlib/stringtools_basic`
- `sys/file_read_write_smoke`
- `sys/process_echo_smoke`

### Upstream module sweep (strict CI-gated)

Source list: `test/upstream_std_modules.txt`
This sweep validates module symbol/type availability and target compatibility (`haxe` compile + `go test ./...`) for each listed module.

```text
haxe.CallStack
haxe.Exception
haxe.Int32
haxe.Int64
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
  - Result: `174 passed / 1 expected policy / 0 failed / 0 unexpected present`

Policy sources:

- `test/upstream_std_expected_missing.json` (currently empty)
- `test/upstream_std_expected_unavailable.json` (currently one `go_test`-stage rule for `haxe.Http`)

## Unsupported expression inventory (compiler hard-fail paths)

These are explicit fatal guards in `src/reflaxe/go/GoCompiler.hx` that represent unsupported paths.

| Inventory item | Current behavior | Acceptance criteria for closure |
| --- | --- | --- |
| Non-lvalue assignment targets in `lowerLValue` | Fatal: `Unsupported assignment target` | Either (a) support any newly reachable legal lvalue shape, or (b) keep as invariant and add a dedicated negative test if a reproducer becomes possible. |
| Non-`++/--` postfix unary in `lowerExpr` / `lowerExprWithPrefix` | Fatal: `Unsupported postfix unary operator` | Keep parser/typed-ast assumptions validated; if new postfix forms become reachable, add lowering + snapshots before enabling. |
| Catch-all `lowerExpr` default | Fatal: `Unsupported expression` | Replace with exhaustive typed-expression handling or narrower typed-node diagnostics; add negative regression coverage for each unsupported reachable node. |
| Unsupported constant kind in `lowerConst` | Fatal: `Unsupported constant` | Add lowering + snapshot for any new constant kind encountered in real programs. |
| Unsupported `Std.isOfType` target kind | Fatal with target-type diagnostic | Add explicit lowering path and dedicated snapshot for each new accepted target type. |

## Known stdlib parity gaps (probe inventory)

As of **2026-02-19**, the broader probe list:

```bash
python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_gap_probe.txt --strict --go-test
```

reports:

- `53 passed / 0 expected policy / 0 failed / 0 unexpected present`

The only active expected-policy rule in the full inventory is:

- `haxe.Http` (`go_test` stage): generated probe output references an unresolved `sys__Http` symbol.

## Tracking

- `haxe.go-d5u`: publish and maintain this matrix/inventory.
- `haxe.go-61w`: reduce compiler hard-fail unsupported expression surface.
- `haxe.go-19u`: expand stdlib parity from the documented probe gap list.
- `haxe.go-ab2`: add semantic differential regression harness.
