# Ownership Rubric

This document is the canonical rulebook for deciding where portable- and metal-contract behavior is allowed to live in `haxe.go`.

Read this before adding:

- a new stdlib parity patch,
- a new `hxrt` helper,
- a new compiler-emitted shim,
- or a new `go.*` facade surface.

Related documents:

- `docs/profiles.md` for semantic contracts (`portable`, `metal`)
- `docs/portable-stdlib-parity-program.md` for parity goals and blocker tranches
- `docs/portable-module-mapping-contract.md` for per-module ownership mapping
- `docs/stdlib-shim-rationale.md` for current shim-by-shim decisions

## What This Document Decides

This document answers one question:

When `haxe.go` needs behavior, which layer should own it?

That question matters because the repo now has several valid implementation layers:

1. compiler lowering in `src/reflaxe/go/GoCompiler.hx`
2. dedicated compiler-owned emitter modules under `src/reflaxe/go/compiler/**`
3. staged std overrides under `std/**` and `std/_std/**`
4. Go runtime helpers under `runtime/hxrt/**`
5. explicit Go-native facade modules under `go.*`
6. explicit exclusions when a portable claim would be dishonest

If this choice is not made up front, parity work tends to drift into `GoCompiler.hx`, and the compiler becomes harder to reason about and easier to regress.

## Contract Invariants

These rules are fixed:

1. `portable` and `metal` are the only semantic contracts.
2. Planner/runtime/lane controls are additive capabilities, not replacement contracts.
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

- `std/haxe/Template.cross.hx`
- `std/haxe/exceptions/PosException.cross.hx`
- `std/haxe/ds/BalancedTree.cross.hx`
- `std/haxe/io/GoIoHelpers.cross.hx`

The paths above are pre-migration source locations. The canonical source
contract keeps upstream overrides as ordinary `.hx` files under
`std/go/_std`; package staging alone flattens that root and renames those files
to `.cross.hx`. Repo-authored support, typed `hxrt` bindings, and public
`go.*` facades do not become override artifacts.

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

## 3. Compiler-Owned Emitters

Owner for irreducible backend-specific behavior that depends on compile-time information.

Use compiler-owned emitters when correctness depends on:

- compile-time metadata,
- generated type tables,
- representation-sensitive lowering,
- profile/boundary policy,
- or backend-specific orchestration that staged std and `hxrt` cannot express honestly.

Good current examples:

- reflection/type creation metadata
- serializer/unserializer metadata-driven emission
- socket/readiness/deadline behavior
- representation-coupled bytes/string paths

Important rule:

- compiler-owned does **not** mean “leave it in the `GoCompiler.hx` monolith”
- compiler-owned means it may stay in compiler space, but should live in dedicated emitters/planners when the surface is non-trivial

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
- `haxe.io.Bytes`
- `Lambda` / generic `Iterable<T>` calls, where public stdlib behavior stays source-owned and the compiler owns only the representation bridge for Go arrays, lists, and manual iterators

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

## What Not To Do

Do not:

1. add a third semantic profile
2. solve portable parity gaps by routing users to `go.*`
3. grow behavior-heavy raw Go blobs inside `GoCompiler.hx` when staged std or `hxrt` would work
4. leave mixed ownership implicit
5. claim semantic-diff parity for surfaces the repo itself only treats as snapshot-only or compile-only

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
  - narrow helper support in compiler/runtime only where necessary
- direct `haxe.exceptions.*` construction
  - staged std ownership for the subclass surface
  - `hxrt` carrier for runtime exception transport
- `haxe.ValueException`
  - direct constructor/catch behavior through the existing runtime exception carrier
- `BalancedTree` / `GenericStack`
  - staged std ownership, not more compiler-resident collection blobs
- `sys.Http`
  - mixed ownership, with core choreography still compiler-owned and leaf helpers extracted cautiously

Those examples should be copied as patterns.
They should not be treated as one-off accidents.
