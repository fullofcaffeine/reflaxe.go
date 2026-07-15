# Profile Semantics and Migration Guide

`portable` and `metal` are accepted compatibility selectors for policy presets.
They are not separate compiler engines, and `metal` is not a second semantic
product.

Accepted compatibility input does not imply release admission. See the
[generated compatibility matrix](compatibility-support-matrix.md): the current
pre-1.0 beta scope admits the named portable surface and records `metal` as
`compatibility-only`.

Read the full normative contract in [Native policy presets and semantic
boundaries](native-policy-presets.md).

## Terms

- [policy preset](glossary.md#policy-preset): compatible defaults selected by
  the `portable|metal` build selector.
- [Go-native](glossary.md#go-native): source that intentionally adopts Go API
  or runtime semantics.
- [native boundary](glossary.md#native-boundary): module-level Go-native
  authority declared with `@:goNative` or a typed Go API.

The product rule is: portable Haxe semantics are the default product path;
typed `go.*`/extern APIs and `@:goNative` modules are explicit Go-native source
boundaries. The `metal` selector is defined as a convenience policy preset
rather than a second semantic product. Portable by default, Go-native by
explicit source boundary.

## The practical rule

Start portable. Let the source API state when code becomes Go-native.

- Portable source surfaces preserve Haxe behavior under either preset.
- `go.*`, typed Go externs, and `@:goNative` explicitly choose Go-native
  semantics for their boundary.
- Optimized or Go-shaped output does not by itself make source Go-native.

This distinction keeps output quality independent from portability. Portable is
not the slow/basic mode, and metal is not the only real-Go mode.

## What the selectors currently change

Only defaults:

| Policy | `portable_default` | `metal_compatibility` |
| --- | --- | --- |
| Native authority | guarded | explicit |
| Native specialization | proven | eager |
| Native fallback | allow | error |
| Strict raw-boundary policy when set to `auto` | off | on |

Explicit policy defines override these defaults. Runtime slicing, planner mode,
optimizer selection, and diagnostic capabilities are independent.

## What must stay invariant

If a module uses only portable surfaces, changing the selector alone must not
change its observable Haxe behavior.

Examples include:

- null and equality behavior;
- Haxe exception and portable-thread lifecycle behavior;
- staged stdlib results;
- portable collection and string behavior;
- portable module APIs and cross-target intent.

Native representation is still allowed when a capability proof shows the same
behavior. A profile-shaped implementation branch is not proof.

## Mixed codebases

One project can combine portable domain code with explicit Go adapters.

```text
domain/                 portable Haxe semantics
application/            portable orchestration
platform/go/            typed externs and @:goNative adapters
```

The directory names are only organization. The compiler recognizes typed API
usage and `@:goNative`, not a magic folder convention.

Recommended CI configuration:

```bash
-D reflaxe_go_profile=portable
-D reflaxe_go_native_authority=guarded
-D reflaxe_go_portable_native_policy=error
```

Approved adapter modules use `@:goNative`. Accidental `go.*` usage elsewhere
then fails early.

## Migrating an existing metal build

There is no requirement to migrate today; `metal` remains supported. To make a
build's intent more explicit without changing behavior:

1. Keep `reflaxe_go_profile=metal` while adding contract reports.
2. Identify modules that actually consume `go.*` or typed Go externs.
3. Mark module-owned native boundaries with `@:goNative`.
4. Confirm raw `__go__` is absent from app and boundary code.
5. Reproduce the preset through explicit axes in a portable build:

   ```bash
   -D reflaxe_go_profile=portable
   -D reflaxe_go_native_authority=explicit
   -D reflaxe_go_native_specialization=eager
   -D reflaxe_go_native_fallback=error
   -D reflaxe_go_strict_policy=on
   ```

6. Compare snapshots, semantic-diff output, runtime output, reports, and
   performance.
7. If desired, move from global `explicit` authority to `guarded` plus
   `@:goNative` modules.

That exercise proves whether the global preset has any remaining value for the
application. It is not a deprecation requirement.

## Migrating source metadata

Existing source:

```haxe
@:goMetal
class NativeAdapter {}
```

Canonical source:

```haxe
@:goNative
class NativeAdapter {}
```

Both compile with identical boundary behavior. The old spelling remains a
silent compatibility alias; the rename can happen on the application's normal
schedule.

## Fallback strategy

Fallback means a typed native representation was attempted but could not be
proven safe.

- Use `allow` while measuring capability gaps or when the semantics-safe
  representation is acceptable.
- Use `error` when every user-owned attempted native representation must be
  concrete.
- Use `auto_strict` when explicit native modules must also reject unresolved
  fallback.

Fallback events are not automatically semantic failures. Reports intentionally
record framework-template attempts even when the error policy excludes
framework internals.

## Why the semantic boundary is explicit

Inferring a native contract from generated output would be unstable: optimizer
changes could silently change source meaning. The compiler instead reads stable
source evidence—APIs, extern metadata, and `@:goNative`—then chooses a lowering
that preserves that evidence.

This differs from inferring a global profile. A build-wide policy can change
which diagnostics or representations are attempted, but it cannot silently
reclassify ordinary Haxe source as Go-native.

## Sibling-target interpretation

- Reflaxe.Elixir demonstrates one pipeline with authoring intent carried by
  typed APIs and metadata.
- haxe.ruby keeps one pipeline and treats its profiles as real semantic
  guardrails; it explicitly avoids inventing a metal/performance profile.
- haxe.rust retains a meaningful metal contract because ownership, nullability,
  string representation, borrowing, and no-runtime behavior create genuine
  semantic differences.

Go's current profile-only effects decompose cleanly into policy axes, so Go
follows the first two precedents unless future evidence demonstrates an
irreducible semantic distinction.

## Review boundary

The compatibility-preserving refactor can land with the repository's xhigh
written second-pass fallback and full test evidence. Any decision to deprecate
or remove `metal` is broader: it needs a genuine independent deep review,
commit-pinned evidence, a SemVer migration plan, and separate approval.

## Validation

```bash
npm run test:changed
npm test
npm run test:semantic-diff
npm run test:examples
npm run test:stdlib-sweep:go-test
```

For optimization or runtime changes, also run the relevant performance and
security gates documented in the repository instructions.

## Related docs

- [Profiles reference](profiles.md)
- [Defines reference](defines-reference.md)
- [Portable canonical contract](portable-canonical-contract.md)
- [Semantic diff guide](semantic-diff-guide.md)
- [Examples matrix](examples-matrix.md)
