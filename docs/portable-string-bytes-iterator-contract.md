# Portable String, Bytes, and Iterator Contract

## Practical outcome

The registry now recognizes three more ordinary portable Haxe surfaces:

- `String` selects `go_string`, meaning the compiler's nullable pointer-backed
  Go string carrier;
- `haxe.io.Bytes` selects `go_byte_slice`, meaning the staged shared data/view
  carrier, not a naked `[]byte`;
- `Iterator<T>` selects `go_iterator` only for the exact `hasNext`/`next` protocol
  and a recursively proven `T`.

These proof-backed decisions now feed the shared surface planner. String and
Bytes route their already-implemented carriers only after admission; String
fast paths also require the selected String carrier. Iterator remains on its
registered `hxrt_iterator` fallback until `.7.7` introduces a distinct
statically typed Go carrier and an exact per-shape lowering gate.

Functions and closures are deliberately not admitted. Their capture behavior is
well covered, but portable method identity still has an incorrect edge case.
That narrower fix is tracked by `haxe_go-vfp.7.11`.

## String: a semantic carrier, not raw bytes

The admitted `go_string` value is a nullable pointer-backed Go string. The
pointer keeps Haxe `null` distinct from both `""` and the literal `"null"`.
Runtime operations preserve value equality, null-aware concatenation and
stringification, and haxe.go's established Unicode-scalar behavior for length,
character lookup, code-point lookup, split, substring, and iteration.

“Unicode scalar” means one decoded code point, which Go calls a rune. It is not a
UTF-8 byte offset. For example, the emoji `😀` occupies four UTF-8 bytes but is
one character under this source contract.

There are two easily confused encoding boundaries:

1. ordinary haxe.go `String` operations use the scalar/rune contract above;
2. `-D reflaxe_go_raw_native_mode=utf16le` changes how
   `haxe.io.Encoding.RawNative` converts a string to and from bytes.

The UTF-16LE switch is an explicit Bytes encoding policy. It never authorizes
String indexing by UTF-8 bytes or UTF-16 code units. This separation is tested
by the portable String/Bytes semantic fixture, UnicodeString coverage, and the
RawNative UTF-16LE generated-output snapshot.

The native and fallback paths both report the `string` runtime capability today,
so the contract is marked no-`hxrt` ineligible.

## Bytes: one object with two coherent views

Portable `Bytes` owns a shared data/view carrier:

- the authoritative public `BytesData` storage is an aliased `[]int`, matching
  Haxe's integer-valued byte access;
- an opaque `ByteView` may cache a native `[]byte` for Go runtime operations;
- public or object-side mutation invalidates or revalidates that cache before
  it is reused.

This preserves the behavior users can observe:

- assigning the Bytes object keeps object identity;
- `getData()` and `ofData()` preserve mutation in both directions;
- values are masked to `0...255`;
- `sub()` returns independent copied storage;
- overlapping `blit()` behaves like a safe overlapping copy;
- null and allocated-empty Bytes stay distinct;
- UTF-8 and RawNative conversion use their explicit encoding policies.

Therefore `go_byte_slice` does not mean “replace `Bytes` with `[]byte`.” It
names this reviewed hybrid carrier. Both selected and fallback paths report the
`bytes` runtime capability and are no-`hxrt` ineligible today.
The `bytes_normalization_contract` semantic fixture specifically proves that
integer writes normalize into the `0...255` range; the broader portable
String/Bytes fixture proves aliasing, copying, overlap, null/empty, and encoding
behavior.

## Iterator: shared progress, not a collection rewrite

Haxe `Iterator<T>` is structural rather than nominal. The typed ledger resolves
the one canonical `StdTypes.Iterator<T>` typedef to its anonymous protocol;
user typedefs remain opaque. The registry recognizes that anonymous type only
when it has exactly these two non-optional fields:

```haxe
hasNext: Void -> Bool
next: Void -> T
```

An object with extra fields, parameters, an optional method, or a non-Boolean
`hasNext` is just an anonymous object to this registry.

For a recursively proven `T`, `go_iterator` names a future statically typed Go
carrier whose operations share one cursor or state owner. The semantic contract
requires that carrier to preserve:

- one evaluation of the iterator source;
- source order;
- repeated `hasNext()` calls without advancing;
- live mutation that the source iterator exposes;
- one cursor and exhaustion state shared by aliases.

It does not mean Go `range`, a copied collection, or a snapshot. Haxe leaves
`next()` after exhaustion unspecified, so the contract does not invent a new
result there.

Today, even eligible Iterator shapes retain `hxrt_iterator`: the established
structural value is an erased `map[string]any` containing `hasNext` and `next`
closures. A stored `next` closure may have a concrete Go return type, but the
surrounding erased map is still the fallback representation. Calling that map
`go_iterator` would make reports claim a native carrier that does not exist.

`Iterator<Dynamic>`, unresolved element shapes, named generic parameters, and
opaque typedef or abstract element storage also retain `hxrt_iterator` and
report the `core` runtime capability. A user-defined typedef of `Iterator<T>`
remains nominal and opaque to admission.

## Why closures remain unadmitted

The semantic fixture proves mutable captures, callback reuse, independent
closure state, loop capture, bound-receiver mutation, same-receiver method
comparison, null comparison, and Dynamic callback fallback.

It also found a real boundary: the current Go implementation of
`Reflect.compareMethods` compares function code pointers. The same method on two
different receiver objects can therefore compare equal, while portable Haxe
requires different receivers to compare unequal. Admitting `haxe.Function`
from simple capture tests would hide that mismatch.

The registry continues to recognize function shapes and reports
`contract_missing`. `haxe_go-vfp.7.11` owns a stable callable-identity carrier
and guarded null invocation before function admission can be reconsidered.

## What sibling compilers teach us

Sibling targets are useful evidence about semantic risks, but they do not choose
Go representations:

| Compiler | Useful precedent | What is still Go-specific |
| --- | --- | --- |
| `haxe.rust` | Distinguishes nullable/owned strings, shared Bytes identity, shared iterator state, and callable identity/capture cells. Its closure, iterator, Unicode, and Bytes fixtures informed this test matrix. | Rust ownership and representation plans do not prove Go pointer, slice, map, or function behavior. |
| `haxe.elixir` | Keeps portable behavior in staged source and uses small target runtime boundaries. | BEAM binaries, processes, and function identity are not Go carrier authority. |
| `haxe.ruby` | Uses typed reachability and portable behavior tests without a comparable native-representation registry. | Ruby object/string/closure behavior cannot admit a Go carrier. |
| Genes | Supplies compatibility and documentation discipline for another dynamic target. | JavaScript string, array, iterator, and closure shapes do not establish Go semantics. |

The family lesson is to preserve source semantics and make target boundaries
explicit. Each compiler still needs target-local representation proof.

## Evidence

Run the focused contracts with:

```bash
npm run test:surface-contract-registry
python3 test/run-semantic-diff.py --case portable_string_bytes_contract
python3 test/run-semantic-diff.py --case bytes_normalization_contract
python3 test/run-semantic-diff.py --case portable_iterator_closure_contract
python3 test/run-semantic-diff.py --case unicode_string_source_owned
python3 test/run-semantic-diff.py --case structural_iterator_assignment_contract
```

Generated-shape evidence also lives in:

- `test/snapshot/core/string_concat_null_semantics`;
- `test/snapshot/stdlib/bytes_basic`;
- `test/snapshot/core/raw_native_utf16_mode`;
- `test/snapshot/core/structural_iterator_assignment`.
