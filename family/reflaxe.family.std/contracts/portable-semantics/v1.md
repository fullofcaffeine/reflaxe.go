# Portable Semantics Spec v1

Spec ID: `portable-semantics-v1`  
Status: Active  
Baseline: Haxe `4.3.7` portable-eligible standard library surface  
Canonical implementation target: `haxe.go` (`reflaxe.go`)

This document defines the normative portable semantics contract for the compiler family.
It is intentionally strict on cross-target risk areas and mapped to deterministic tests.

## Scope

This spec applies when code stays on the portable contract surface:

- Haxe language + portable-eligible stdlib/application code
- no target-native facade usage (`go.*`, `rust.*`, `ocaml.*`, etc.) in portable-contract modules
- no raw target injection in portable-contract modules

Portable surface membership is governed by `test/portable_allowlist.json`.

Out of scope:

- target-native APIs/facades
- ABI-level native interop details
- profile-specific optimization strategy details that do not change portable observable behavior

## Normative Rules

The rules below are mandatory for portable contract behavior.

### 1) Null + Dynamic Boxing Semantics

1. `Std.string(null)` must produce `"null"`.
2. String concatenation with null-like values (including boxed typed-nil values) must preserve portable `"null"` behavior.
3. Null equality through boxed dynamic pathways must remain portable-correct (for example `d == null` where `d` comes from a typed-nil source should evaluate `true` in portable pathways).
4. Typed-nil values must not leak target-native interface-nil quirks into portable outcomes (for example `"<nil>"` stringification drift).

Conformance fixtures:

- `test/semantic_diff/null_string_concat`
- `test/semantic_diff/non_string_null_equality_contract`
- `test/semantic_diff/typed_nil_dynamic_string_contract`
- `test/semantic_diff/nullable_struct_refs`

### 2) Exception Identity + Flow Semantics

1. `throw` / `try` / `catch` must preserve Haxe-visible behavior for typed and dynamic catches.
2. `haxe.Exception` API contract (`caught`, `thrown`, `message`) must remain stable.
3. Throw-as-expression and try/catch value-forwarding behavior must match interpreter-observed behavior.
4. Cross-module exception behavior must remain deterministic and profile-invariant for portable-surface code.

Conformance fixtures:

- `test/semantic_diff/exceptions_typed_dynamic`
- `test/semantic_diff/exception_api_contract`
- `test/semantic_diff/throw_expr_contract`
- `test/semantic_diff/try_catch_return_forwarding_contract`

### 3) String/Bytes/Encoding Semantics

1. Core `haxe.io.Bytes` behavior (indexing, conversion, hex, aliasing contracts) must match portable fixture expectations.
2. Stream behaviors (`BytesInput` / `BytesOutput` helper subset) must be deterministic and interpreter-aligned for covered APIs.
3. Encoding behavior must be explicit:
   - default mode (`reflaxe_go_raw_native_mode=interp`) is portable baseline
   - alternate mode (`utf16le`) is opt-in compatibility behavior and must not silently alter default portable expectations
4. Serializer/date/bytes token handling in covered paths must remain contract-stable.

Conformance fixtures:

- `test/semantic_diff/bytes_normalization_contract`
- `test/semantic_diff/bytes_ops_contract`
- `test/semantic_diff/bytes_of_data_contract`
- `test/semantic_diff/bytes_hex_contract`
- `test/semantic_diff/bytes_io_stream_contract`
- `test/semantic_diff/io_input_output_helpers_contract`
- `test/semantic_diff/io_input_output_edge_contract`
- `test/semantic_diff/io_encoding_contract`
- `test/semantic_diff/serializer_date_bytes_contract`

### 4) Reflection + Type Introspection Limits

Portable reflection guarantees are intentionally bounded by tested behavior.

1. `Reflect` behavior is guaranteed only for covered operations and semantics.
2. `Std.isOfType` behavior must remain stable for the covered runtime/core/abstract scenarios.
3. Type-value expression behavior (`TTypeExpr`-mapped semantics) must remain stable where explicitly covered.
4. Behavior outside the covered reflection subset is not implicitly guaranteed; it must be added via explicit fixtures before being claimed portable.

Conformance fixtures:

- `test/semantic_diff/reflect_compare`
- `test/semantic_diff/reflect_field_ops`
- `test/semantic_diff/std_is_of_type_contract`
- `test/semantic_diff/std_is_of_type_runtime_core_abstract_contract`
- `test/semantic_diff/type_expr_contract`

### 5) Numeric Edge Behavior

1. Numeric edge behavior must match the portable fixtures (overflow/shift/coercion-sensitive cases under current contract scope).
2. `haxe.Int32` and `haxe.Int64` covered operations must preserve fixture-defined outcomes.
3. Any optimization that changes numeric observable behavior for covered cases is a portable contract regression.

Conformance fixtures:

- `test/semantic_diff/numeric_edge_cases`
- `test/semantic_diff/int32_contract`
- `test/semantic_diff/int64_contract`

### 6) Array Identity and Growth

1. A portable `Array<T>` is one mutable reference object. Assignments, fields,
   parameters, returns, callbacks, `Dynamic`, and erased generic boundaries must
   observe the same length and contents after mutation.
2. `push`, `pop`, `insert`, `remove`, and `length` changes must update that shared
   identity rather than a copied Go slice header.
3. Indexed assignment at or beyond the current length must grow the Array and
   preserve null-filled gaps, including for statically non-null element types.
4. `copy()` creates a distinct Array identity. Converting between portable
   `Array<T>` and explicit `go.NativeSlice<T>` storage is a shallow copy and must
   be written explicitly at the native boundary.
5. Erased element equality retains Haxe value rules, including content equality
   for strings and identity equality for reference-shaped values.

Conformance fixtures:

- `test/semantic_diff/array_identity_contract`
- `test/snapshot/core/array_identity`
- `test/snapshot/go_native/native_slice_boundary`

## Contract Invariance Across `portable` and `metal`

If a program stays on portable surfaces, these semantics must remain equivalent when compiled with:

- `-D reflaxe_go_profile=portable`
- `-D reflaxe_go_profile=metal`

Differences are allowed only when code explicitly opts into target-native behavior outside portable contract boundaries.

## Conformance Gates

Portable semantics changes are valid only when all gates remain green:

```bash
python3 test/run-semantic-diff.py
python3 test/run-snapshots.py
python3 test/run-ci.py
```

Portable allowlist and governance gates must also stay green:

```bash
python3 test/run-portable-allowlist.py
npm run test:stdlib:governance
```

## Versioning Policy

`portable-semantics-v1` is a versioned contract artifact.

- Patch-level clarifications: wording or reference cleanup with no behavior change.
- Minor-level updates: additive guarantees with new fixtures, no break to existing guarantees.
- Major-level updates: semantic changes to existing guarantees (must create `portable-semantics-v2`).

When a normative rule changes:

1. Update this spec version or publish a new version.
2. Update fixture mappings.
3. Update `docs/portable-canonical-contract.md` and `test/README.md`.
4. Document migration impact in release notes.
