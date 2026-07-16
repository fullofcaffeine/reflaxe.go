# Ownership Rubric

This document is the canonical rulebook for deciding where portable semantics
and explicit Go-native behavior are allowed to live in `haxe.go`.

Read this before adding:

- a new stdlib parity patch,
- a new `hxrt` helper,
- a new compiler-emitted shim,
- or a new `go.*` facade surface.

Related documents:

- `docs/native-policy-presets.md` for the source semantic boundary and policy presets
- `docs/portable-stdlib-parity-program.md` for parity goals and blocker tranches
- `docs/portable-module-mapping-contract.md` for per-module ownership mapping
- `docs/stdlib-shim-rationale.md` for current shim-by-shim decisions

## What This Document Decides

This document answers one question:

When `haxe.go` needs behavior, which layer should own it?

That question matters because the repo now has several valid implementation layers:

1. compiler lowering in `src/reflaxe/go/GoCompiler.hx`
2. dedicated compiler-owned emitter modules under `src/reflaxe/go/compiler/**`
3. staged std overrides under the canonical `std/go/_std/**` source root
4. Go runtime helpers under `runtime/hxrt/**`
5. explicit Go-native facade modules under `go.*`
6. explicit exclusions when a portable claim would be dishonest

If this choice is not made up front, parity work tends to drift into `GoCompiler.hx`, and the compiler becomes harder to reason about and easier to regress.

## Contract Invariants

These rules are fixed:

1. Portable Haxe is the default semantic contract; typed native APIs/externs
   and `@:goNative` are explicit native boundaries.
2. The compatible presets and planner/runtime/policy controls are additive
   defaults/capabilities, not replacement contracts.
3. Portable parity work lands against the `portable` contract first.
4. `go.*` is always a typed native facade, never the answer to a portable parity gap.

In practice:

- If a module is part of the portable Haxe surface, solving it with `go.*` does not count as closure.
- If a feature needs a Go-native power surface, it belongs in `go.*` and must be documented as non-portable.

## The Ownership Layers

## 1. Staged Std (`std/go/_std/**/*.hx` source)

Default owner for library-expressible behavior.

Use staged std when:

- the behavior is ordinary Haxe/library logic,
- the semantics can be expressed in Haxe source,
- and the implementation does not need compiler-only metadata tables or backend-only type graph decisions.

Good current examples:

- `std/go/_std/haxe/Template.hx`
- `std/go/_std/haxe/exceptions/PosException.hx`
- `std/go/_std/haxe/ds/BalancedTree.hx`
- `std/haxe/io/GoIoHelpers.hx`

The first three are canonical upstream overrides. `GoIoHelpers` is a
repo-authored staged-support module in its ordinary source location. Package
staging flattens only the canonical override root and renames those files to
`.cross.hx`; support, typed `hxrt` bindings, and public `go.*` facades remain
ordinary modules.

## 2. `hxrt` Runtime (`runtime/hxrt/**`)

Owner for reusable runtime behavior over already-lowered Go representations.

Use `hxrt` when:

- the behavior is runtime logic, not compile-time planning,
- multiple surfaces need the same Go-side implementation,
- and Haxe source alone would be an awkward or misleading place for the logic.

Good current examples:

- exception carrier/wrapping behavior
- JSON helpers
- file/process/sys runtime helpers
- bytes/string helper logic that operates on already-lowered representations

## 3. Compiler-Owned Intrinsics

Last-resort owner for an exact primitive that depends on compile-time information.

Use compiler-owned emitters when correctness depends on:

- compile-time metadata,
- generated type tables,
- representation-sensitive lowering,
- profile/boundary policy,
- or a backend representation fact that staged std and `hxrt` cannot express honestly.

Good current examples:

- reflection/type creation metadata derived from the final reachable type graph
- `Std.isOfType` checks against a compiler-known type token
- the exact `haxe.Rest` slice construction bridge
- exception-carrier conversion at typed catch/throw boundaries
- exact `Lambda` iterable/callback/nested-carrier/result and sort generic-erasure
  adapters whose algorithms remain in Haxe source; the private nominal Lambda
  companion only delegates `iterator()`

Important rule:

- a tested implementation is not automatically an intrinsic
- runtime behavior such as sockets, HTTP, regex, atomics, parsing, compression,
  and collection algorithms belongs in staged source or `hxrt`
- representation sensitivity justifies only the smallest exact primitive, not
  compiler ownership of the surrounding public API
- every admitted intrinsic must appear in
  [`compiler-stdlib-intrinsics.json`](compiler-stdlib-intrinsics.json)

## 4. `go.*` Native Facade

Owner for explicit Go-native APIs only.

Use `go.*` when:

- the user is intentionally asking for Go-native behavior,
- the surface is not portability-safe,
- and the right answer is a typed native facade, not raw `__go__`.

Examples:

- `go.Chan<T>`
- `go.Select`
- typed package extern facades like `go.Fmt`, `go.Time`, `go.Http`

## 5. Explicit Exclusion

Owner of last resort when a portable claim would be dishonest.

Use explicit exclusion when:

- the semantics are not portable enough to promise honestly,
- or the target/platform behavior is intentionally narrower than upstream Haxe.

Examples already in this repo:

- target-sensitive stack/UCS2 behavior under snapshot-only coverage
- surfaces still classified as compile-only with named blockers and closure targets

The rule here is simple:

- never fake parity
- either support it, or classify it explicitly

## Decision Order

When adding or fixing a parity slice, choose ownership in this order:

1. Can this live in staged std as ordinary Haxe code?
2. If not, can it live in `hxrt` as reusable runtime behavior?
3. If not, is there a real compile-time reason it must stay compiler-owned?
4. If it is Go-only, should it actually be a `go.*` facade surface instead?
5. If none of the above are honest, should it be an explicit exclusion?

That order is deliberate:

- staged std is the default,
- `hxrt` is the reusable runtime layer,
- compiler ownership is the exception,
- `go.*` is not a portability escape hatch,
- exclusion is better than counterfeit support.

## Mixed Ownership

Mixed ownership is allowed, but only when the split is explicit and documented.

Mixed ownership means:

- the public Haxe-facing behavior may live in staged std,
- while one or more low-level helpers remain in `hxrt` or compiler-owned emitters,
- and the split is recorded in `docs/portable-module-mapping-contract.md`

Good mixed examples:

- `Std`
- `Sys`
- `haxe.Json`
- staged RTTI source over a narrow generated metadata table

`haxe.io.Bytes` remains a mixed migration implementation. `Lambda` and sort
calls are mixed only at exact registered Go representation adapters; their
behavior is source-owned. Neither case permits new compiler-owned algorithms.

Bad mixed ownership looks like:

- half the behavior in `GoCompiler.hx`,
- half in a helper file,
- no mapping doc update,
- and no explanation of why the split exists

That is not mixed ownership. That is drift.

## Canonical Migration Ledger

`docs/stdlib-provenance-ledger.json` is the executable per-file decision
record. It distinguishes six ownership classes:

1. `upstream_std_override` migrates to ordinary source under
   `std/go/_std` and becomes `.cross.hx` only in a staged package.
2. `staged_support` remains an ordinary `.hx` module under its Haxe package.
3. `hxrt_binding` moves under `std/hxrt` and models a real Go runtime API.
4. `public_go_facade` remains under `std/go`, outside `_std`.
5. `obsolete` requires an explicit removal owner and evidence.
6. `intentional_boundary_fixture` requires an explicit policy-fixture owner.

Each ledger entry has one exact destination and migration Bead. Its
`compilerShimGroups` list is also exact: an empty list means the adjacent shim
audit found no directly selected compiler group, not that the audit was
skipped. Any unresolved classification must appear in the ledger's
`ambiguities` list with a follow-up Bead.

The file ledger does not approve compiler ownership. The separate
[`compiler-stdlib-intrinsics.json`](compiler-stdlib-intrinsics.json) registry
inventories exact Haxe symbols, selector paths, compiler entry points, direct
call rewrites, debt classification, evidence, review conditions, and migration
Beads. Its bidirectional test fails when a compiler stdlib symbol or entry point
is added without an explicit decision.

## What Not To Do

Do not:

1. add a semantic profile where explicit source boundaries and orthogonal
   policies express the distinction
2. solve portable parity gaps by routing users to `go.*`
3. grow behavior-heavy raw Go blobs inside `GoCompiler.hx` when staged std or `hxrt` would work
4. leave mixed ownership implicit
5. claim semantic-diff parity for surfaces the repo itself only treats as snapshot-only or compile-only
6. label a whole shim group “required” because one member consumes compiler metadata

## Required Update Sequence

When ownership changes for a portable surface, update all relevant artifacts in the same change:

1. implementation layer (`std/**`, `runtime/hxrt/**`, or compiler emitter)
2. `docs/portable-module-mapping-contract.md`
3. `docs/stdlib-shim-rationale.md`
4. `docs/portable-stdlib-parity-program.md` if blocker/phase status changes
5. provenance/governance artifacts when staged std files change
6. parity evidence:
   - snapshot
   - semantic-diff
   - sweep/inventory status

## Current High-Signal Examples

These examples represent the intended direction:

- `haxe.Template`
  - semantics in staged std
  - its remaining compiler-generated runtime reflection bridge is migration debt,
    not an approved pattern
- direct `haxe.exceptions.*` construction
  - staged std ownership for the subclass surface
  - `hxrt` carrier for runtime exception transport
- `haxe.ValueException`
  - direct constructor/catch behavior through the existing runtime exception carrier
- `BalancedTree` / `GenericStack`
  - staged std ownership, not more compiler-resident collection blobs
- `sys.Http`
  - currently mixed, with compiler-owned choreography explicitly scheduled for
    staged source plus typed `hxrt` migration

Those examples should be copied as patterns.
They should not be treated as one-off accidents.
