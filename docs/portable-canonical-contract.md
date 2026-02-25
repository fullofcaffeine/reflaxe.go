# Portable Canonical Contract

This document is the source-of-truth portable contract for the Haxe native compiler family (`haxe.go`, then `haxe.rust`, then `haxe.elixir.codex`).

## Purpose

`portable` is the cross-compiler compatibility lane.

If a program uses only Haxe stdlib/application code (no target-only APIs like `go.*`), then it should compile and run with equivalent semantics across:

- `haxe.go` (`portable`)
- `haxe.rust` (`portable`)
- `haxe.elixir.codex` (portable authoring style / equivalent lane)

## Contract Boundaries

In scope for canonical portability:

- Haxe language semantics (core expressions/control-flow/exceptions)
- Stdlib and `sys.*` behavior that each target claims as supported
- Null and string behavior along portable rules
- Virtual dispatch and inheritance semantics

Out of scope for canonical portability:

- Target-native APIs (`go.*`, Rust-only or Elixir-only native facades)
- Target-native performance lanes (`metal`)
- Target-native interop behavior and ABI details

## Profile Model

`haxe.go` has two profiles:

- `portable`: canonical semantics baseline
- `metal`: opt-in low-level/performance lane

Selective `hxrt` runtime slicing is orthogonal to profiles. It optimizes packaging/runtime footprint and must not change portable semantics.

## Portable Semantic Rules (current)

These are explicitly locked in tests:

- Dynamic dispatch must remain correct across modules.
- Leaf-call devirtualization is allowed only when globally proven safe.
- `Std.string(null)` and `"" + null` produce `"null"`.
- Typed-nil values boxed into `Dynamic`/`any` still behave as null for portable stringification/equality expectations.
- Removed profile selectors (`gopher`, `idiomatic`) fail fast with clear diagnostics.

### Portable null semantics (practical contract)

For portable pathways, null-facing behavior is Haxe-oriented and should not drift to raw target-native nil quirks.

Expected behavior examples:

```haxe
var n:Node = null;
var d:Dynamic = n;
Sys.println(Std.string(d)); // "null"
Sys.println("" + d);        // "null"
Sys.println(d == null);     // true (boxed typed-nil still compares as null in portable contract)
```

This specifically guards against Go interface-nil traps leaking into portable semantics (for example `"<nil>"` stringification or false negative null-equality on boxed typed-nil values).

## Canonical Test Gates in `haxe.go`

Portable contract changes require green results in:

- `python3 test/run-snapshots.py`
- `python3 test/run-semantic-diff.py`
- `python3 test/run-ci.py`

Key semantic guards include:

- `test/semantic_diff/virtual_dispatch_cross_module`
- `test/semantic_diff/typed_nil_dynamic_string_contract`
- `test/snapshot/core/portable_leaf_virtual_devirtualization`
- `test/snapshot/core/portable_non_leaf_virtual_dispatch_preserved`
- `test/snapshot/negative/profile_removed_gopher`
- `test/snapshot/negative/profile_removed_gopher_alias`

## Adoption Pattern for Other Compilers

When aligning `haxe.rust` and `haxe.elixir.codex`:

1. Mirror this contract document into each repo.
2. Port the same semantic-diff/snapshot cases first.
3. Mark each deviation explicitly in docs until resolved.
4. Keep target-native lanes (`metal`-equivalent) separate from portable behavior.

Do not weaken portable behavior to fit a target-native optimization. Target-native lanes should absorb those tradeoffs instead.
