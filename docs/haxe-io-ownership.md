# `haxe.io` ownership contract

## What it is

The base `haxe.io` hierarchy is ordinary staged Haxe source. The canonical Go
target overrides live under `std/go/_std/haxe/io` and package as `.cross.hx`:

- `BufferInput`, `Input`, `Output`, `BytesInput`, `BytesOutput`, and
  `StringInput` own stream behavior and normal source inheritance.
- `Bytes` and `BytesBuffer` own public storage behavior, bounds checks, alias
  semantics, mutation, encoding selection, and byte algorithms.
- `Encoding`, `Eof`, and `Error` own their portable public contracts.
- `FPHelper` owns Haxe word ordering and `Int64` construction.

There is no `io` compiler-shim group, compiler-owned public IO type, or
profile-specific IO implementation.

## Why the boundary exists

Haxe exposes bytes as integer-indexed `BytesData`, while Go codecs and hashes
consume `[]byte`. Reimplementing all of `haxe.io` in the compiler merely to
bridge those representations made portable library policy hard to review and
forced unrelated IO users through a large generated declaration block.

The source/runtime split is narrower:

| Owner | Responsibility |
| --- | --- |
| Staged `haxe.io` source | Public API, validation, EOF/error behavior, stream loops, endian policy, cache invalidation, and alias observation |
| Typed `std/hxrt/io` bindings | Opaque `ByteView`, native allocation/conversion/copy calls, and IEEE-754 bit reinterpretation |
| `runtime/hxrt/bytes.go` | `[]int`/`[]byte` conversion, UTF-8/UTF-16LE conversion, overlap-safe copy, allocation, and Go float-bit primitives |
| Compiler | Normal class lowering, virtual dispatch, source inheritance, and generic representation lowering only |

These runtime operations are target capabilities, not a second public IO API.
No generated `haxe__io__Bytes` layout crosses the typed boundary.

## How the byte cache works

`Bytes` keeps its public integer data and may cache an opaque immutable
`ByteView`:

1. `Bytes.ofString` receives a native view and derives the public integer data.
2. Reads that need native bytes call `__hx_nativeView()`.
3. Every `Bytes` mutation invalidates the cached view.
4. `getData()` marks the integer data as externally exposed. Later native-view
   reads compare the cache with the live values, so mutation through a
   `BytesData` alias cannot return stale bytes.
5. Native results use `__hx_fromNativeView()` to seed both representations.

Base64 and digest APIs consume the same opaque view. This removes the previous
`Bytes -> Array<Int> -> []int -> []byte` copy chain while keeping alphabets,
padding, public construction, and digest API policy in staged Haxe.

## Inheritance and dispatch

`Input` and `Output` use the compiler's ordinary `__hx_this` virtual-dispatch
path. A user class or staged system class that extends either base embeds the
source-owned base view, so assignments and arguments typed as the base class
remain valid Go without IO-specific synthetic wrappers.

The same rule applies in both compatibility presets. `portable` is the default
product path and `metal` is a convenience policy preset; neither selector
changes `haxe.io` semantics or chooses a different hierarchy.

Nominal type references also participate in source inclusion. If manual
dead-code elimination keeps a staged IO class only in a field, argument, or
superclass type, the compiler queues that exact typed declaration rather than
falling back to a mainstream target implementation or restoring an entire
module. Compiler-owned consumers follow the same explicit rule: for example,
the current `sys.Http` carrier queues `BytesBuffer` because its narrow
framework-owned type switch cannot expose that dependency to typed traversal.

The runtime does not retain the former string, hex, or buffer-length policy
helpers. Those algorithms now live entirely in staged `Bytes` and
`BytesBuffer`; only operations that must allocate, reinterpret, or replace a Go
slice cross `std/hxrt/io`.

## Evidence

The migration is guarded by:

- the fail-closed source/ledger/registry contract in
  `test/test_stdlib_migration_ledger_contract.py`;
- semantic-diff cases for bytes, streams, endian values, EOF/errors,
  `BytesData` aliases, serialization, file/process subclasses, and crypto;
- generated-output snapshots, strict stdlib sweep, examples, compiler-debt,
  selective-runtime, performance, and Go security/race gates.

The canonical migration record is `haxe_go-vfp.8.7.11`.
