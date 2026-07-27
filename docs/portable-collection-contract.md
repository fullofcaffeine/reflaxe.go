# Portable Collection Representation Contract

## What it is

This contract says when ordinary portable Haxe collections may use Go-native
storage without changing what the Haxe program means.

The admitted surfaces are:

- `Array<T>` with a fully typed element shape;
- `haxe.ds.StringMap<V>` with a fully typed value shape;
- `haxe.ds.IntMap<V>` with a fully typed value shape.

“Fully typed” means the recursive shape contains neither `Dynamic`, an
unresolved type, nor an alias whose underlying storage is absent from the
ledger. A nested `Array<Array<Int>>` is eligible; `Array<Dynamic>` is not.
Named type parameters, typedefs, and user abstracts currently fall back because
the ledger cannot prove their concrete underlying carrier. Core primitive and
`Null<T>` abstracts are closed, explicitly recognized cases.

Admission records future planner authority. Until the registry-consuming
planner task lands, generated code continues to use the existing semantic
fallbacks.

## Why the carrier is not a raw Go collection

A portable Haxe `Array<T>` is one shared mutable object. If two variables point
to it and one appends an item, both variables observe the new length. A raw Go
slice value contains a copied header, so append can change one variable without
changing the other. Haxe sparse writes can also create null holes, including in
an `Array<Int>`, while a raw `[]int` has no “missing” state.

For that reason, `go_slice` means a shared, slice-backed carrier that preserves:

- shared identity and alias-visible length/content changes;
- sparse holes and null elements;
- null array versus allocated empty array;
- nested collection identity;
- iteration and callback-visible mutation.

It does not mean that the compiler may replace `Array<T>` with a naked `[]T`.
The `go.NativeSlice<T>` and `go.Slice` APIs remain the explicit raw Go slice
boundary and use copying adapters when they cross into portable arrays.

Portable maps have a similar wrapper requirement. A map carrier must preserve
shared mutation, distinguish a missing entry from a present null value through
`exists()`, retain shallow-copy behavior, expose every current entry to
iteration, and preserve nested value identity. It must not expose Go's
randomized map iteration as a new source-level promise.

## Key equality and comparability

The admitted map keys come from typed Haxe surfaces:

| Portable surface | Proven key | Equality authority | Status |
| --- | --- | --- | --- |
| `StringMap<V>` | `String` | exact string equality; Go-comparable | admitted for typed `V` |
| `IntMap<V>` | `Int` | exact integer equality; Go-comparable | admitted for typed `V` |
| `ObjectMap<K,V>` | object identity | Haxe object identity, not ordinary Go value equality | not admitted |
| `EnumValueMap<K,V>` | enum structure | Haxe structural enum comparison | not in this registry yet |

`StringMap` and `IntMap` encode their key types in their nominal type, so their
ledger shapes carry only `V`. The typed
`surface_has_fixed_go_comparable_map_key` rule proves that fixed key. This is
different from checking whether an arbitrary value happens to fit in a Go
interface.

The explicit target-native `go.Map<K,V>` facade is deliberately outside this
portable registry. Its compatibility fallback currently calls `Std.string`
for keys. That behavior is useful only as a target API fallback; the key
coercion is not semantic evidence that two portable Haxe keys have the same
equality or hash contract. No portable admission may rely on that coercion.

## Eligibility and fallback matrix

| Observed typed shape | Decision | Selected carrier | Fallback |
| --- | --- | --- | --- |
| `Array<Int>` | admitted | shared `go_slice` semantic carrier | `hxrt_array` |
| `Array<Array<Int>>` | admitted | recursive shared `go_slice` carriers | `hxrt_array` |
| `Array<Dynamic>` | rejected | none | `hxrt_array` |
| `Array<T>` | rejected: unproven carrier | none | `hxrt_array` |
| `Array<Hidden>` where `typedef Hidden = Dynamic` | rejected: unproven carrier | none | `hxrt_array` |
| `StringMap<Int>` | admitted | fixed-string-key `go_map` semantic carrier | `hxrt_map` + `map_string` |
| `StringMap<Array<Int>>` | admitted | fixed-string-key `go_map` with recursive value carrier | `hxrt_map` + `map_string` |
| `IntMap<String>` | admitted | fixed-int-key `go_map` semantic carrier | `hxrt_map` + `map_int` |
| `StringMap<Dynamic>` | rejected | none | `hxrt_map` + `map_string` |
| `ObjectMap<MyClass,String>` | rejected: `contract_missing` | none | existing staged ObjectMap behavior |
| `go.Map<Array<Int>,Int>` | ignored by portable registry | explicit target-native policy only | target API policy |

Rejected Array/StringMap/IntMap shapes retain the named fallback and its runtime
cost in the surface report. ObjectMap has no admitted contract, so its report
uses `contract_missing`; existing lowering still retains the staged ObjectMap
implementation. Bare generic declaration shapes have no applied type arguments
and reject as `shape_mismatch`; they never authorize a representation.

## Sibling compiler precedent

`haxe.rust` makes the same important product distinction: ordinary Haxe
`Array<T>` keeps shared mutable Haxe semantics, while `rust.Vec<T>` is the
explicit native collection boundary. Its representation plan does not treat a
source Array as permission to emit a raw Vec. `haxe.go` follows that semantic
boundary while using Go-specific key comparability, nil/empty, and slice-header
rules.

The Ruby, Elixir, and Genes sibling targets provide useful typed-reachability,
staged-stdlib, and compatibility-governance precedent, but none currently
supplies a portable native-collection registry contract that can replace this
Go-specific proof.

## How it is proved

The evidence is split by responsibility:

1. `test/semantic_diff/portable_collections_contract` compares interpreter and
   Go behavior for aliasing, mutation, sparse/null and empty states, nested
   values, shallow map copies, iteration, callbacks, and `Dynamic` fallbacks.
2. `test/semantic_diff/array_identity_contract` supplies the deeper Array
   identity matrix across fields, parameters, returns, generics, callbacks, and
   erased values.
3. `test/semantic_diff/ds_maps_list_contract` supplies StringMap, IntMap, and
   ObjectMap equality/copy/iteration evidence. ObjectMap evidence explains why
   it is rejected from native map admission.
4. `test/fixtures/surface_contract_registry` proves deterministic admitted and
   rejected decisions, exact runtime fallbacks, fixed-key comparability, schema
   validity, real typed typedef/abstract hidden-`Dynamic` rejection, profile
   independence, and the current Array generated fallback shape.
   StringMap/IntMap do not claim generated-shape proof until a dedicated
   asserted shape fixture exists.

Run:

```bash
npm run test:surface-contract-registry
python3 test/run-semantic-diff.py --case portable_collections_contract
python3 test/run-semantic-diff.py --case array_identity_contract
python3 test/run-semantic-diff.py --case ds_maps_list_contract
```
