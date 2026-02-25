# Profile Semantics Guide (`portable` vs `metal`)

This document is the canonical profile guide for `reflaxe.go` and the template reference for sibling compilers in the same family.

## Why profiles exist

Profiles are semantic contracts, not just optimization toggles.

- `portable` defines the compatibility baseline.
- `metal` defines the explicit low-level Go lane.

Profiles should be explicit in source control and CI (`-D reflaxe_go_profile=...`) so semantic intent is visible and reviewable.

## Quick definitions

| Profile | Primary goal | Portability expectation | Interop expectation | Codegen tendency |
| --- | --- | --- | --- | --- |
| `portable` (default) | Stable Haxe semantics and lowest migration risk | Highest (within documented support matrix) | Use Haxe stdlib/application surfaces first | More `hxrt`/shim-mediated behavior for semantic stability |
| `metal` (experimental) | Go-native control with strict boundaries | Lower by design when using Go-native surfaces | Typed Go-facing APIs and stricter escape-hatch policy | More typed specialization and direct Go-shaped output in supported lanes |

## What changes and what does not

### Non-negotiable

- Portable semantics are the baseline contract.
- Runtime feature inference/slicing is orthogonal and must not silently change profile semantics.
- Removed profiles (`gopher`, `idiomatic`) are hard errors, not aliases.

### What can differ between profiles

- Default boundary policy (`metal` enables strict app boundary checks by default).
- Amount of typed specialization in generated Go for supported metal lanes.
- How much generated code routes through generic runtime helpers vs typed direct paths.

### What should not silently differ

For programs written against portable surfaces (Haxe stdlib + app code, no target-only APIs), behavior should remain equivalent when compiled with either profile, modulo documented gaps.

If this rule regresses, it is a compiler bug or a contract deviation that must be documented and tested.

### Portable null/string/dynamic semantics to keep in mind

These are easy to break if a compiler silently switches to native-first behavior in loose-typing paths:

1. `Std.string(null)` should behave as portable contract `"null"` stringification.
2. String concatenation with null-like dynamic values should keep portable `"null"` semantics.
3. Boxed typed-nil values in `Dynamic`/`any` pathways should still satisfy portable null expectations (`d == null` style behavior).
4. Dispatch and inheritance behavior should remain portable-correct even with optimization passes (no silent override bypass).

In practice, this is why “explicit profile contract + tests” is safer than “implicit inferred profile.”

### How this compares in `metal`

| Case | `portable` | `metal` (when code stays on portable surfaces) | `metal` (when using native-first surfaces) |
| --- | --- | --- | --- |
| `Std.string(null)` | `"null"` | `"null"` (same contract) | May differ if you bypass portable pathway with native formatting APIs |
| `"" + dynamicNull` | `"null"` | `"null"` (same contract) | May differ if concatenation goes through native-only path |
| `d == null` where `d` is boxed typed-nil | `true` | `true` (same contract) | May differ if code depends on raw target-native boxed-nil behavior |

Rule of thumb:

- If your code remains on portable APIs, null semantics should remain portable even under `metal`.
- Differences appear when you intentionally opt into target-native behavior outside the portable contract.

### How you opt in (or out)

There is no implicit profile switching during compilation.
`reflaxe_go_profile` is explicit and fixed per build.

Opt into native-first behavior:

1. Compile with `-D reflaxe_go_profile=metal`.
2. Use target-native surfaces (for example `go.*` APIs or explicit native interop wrappers).

Stay in portable semantics:

1. Keep code on portable Haxe stdlib/application APIs.
2. Avoid target-native-only surfaces in shared/core modules.
3. You can still compile with `metal` for boundary/perf validation without changing semantics if those modules remain portable-surface-only.

## Choosing a profile

### Choose `portable` when

1. You want the best chance to keep code cross-target.
2. You need predictable Haxe-first semantics.
3. You plan to share logic with other Haxe target outputs.
4. You are starting a new codebase and optimizing later.

### Choose `metal` when

1. You need explicit Go-native interop lanes now.
2. You are optimizing known hot paths with benchmark evidence.
3. You accept a stricter boundary model and lower cross-target portability for those paths.

## Practical default

Start in `portable`, then promote targeted hotspots or interop-heavy modules to metal-oriented patterns once you have benchmark data.

## Why we did not choose a single metal-first surface

We intentionally kept explicit `portable` and `metal` contracts instead of a single inferred “metal-first” mode.

Key reasons:

1. Intent cannot be inferred reliably from usage.
2. Small code changes could silently flip inferred semantics.
3. CI/review needs explicit semantic mode selection in diffs.
4. Cross-target compatibility is easier to preserve when portable is a named contract.

Semantic-flip examples we want to avoid:

- A dependency starts using target-native surfaces and an inferred global mode begins treating nearby loosely-typed value paths (`Dynamic`, stringification, `null` handling) as native-first instead of portable-contract-first.
  In practice, this can change observable behavior:
  ```haxe
  var n:Node = null;
  var d:Dynamic = n;
  Sys.println(Std.string(d)); // portable contract: "null"
  Sys.println("" + d);        // portable contract: "null"
  Sys.println(d == null);     // portable contract: true (null stays null when boxed as Dynamic)
  ```
  A native-first Go pathway can instead produce interface-nil behavior like `"<nil>"` stringification and `d == null` mismatches.
- A refactor that looks “type-only” (for example, replacing a generic container path with a target-specific fast path) quietly changes dispatch/runtime-helper behavior for the same public API.
- A minor dependency update changes inferred feature sets and produces different exception/stringification/equality behavior without an explicit profile change in version control.

In short: runtime feature inference is useful, but semantic profile inference is too risky as a default model.

## Future `auto` direction

`auto` is planned as an explicit additive planner, not a hidden semantic profile.

Expected shape:

- User opts in explicitly.
- Compiler infers runtime/feature selection and emits a deterministic report.
- Compiler does not silently relax boundary or semantic contracts.
- If code crosses restricted boundaries, `auto` should error or require explicit user opt-in flags.

Design spike and rationale are tracked in `docs/profile-auto-spike.md`.

## Compiling portable-oriented code with `metal`

This is a supported workflow.

- If the code uses only portable surfaces, it should compile and usually behave the same.
- Generated Go may look similar when no metal-only specialization is triggered.
- You still get metal boundary defaults (strict app-side injection policy).

This is useful for “audit mode” (checking metal readiness without rewriting the codebase).

## Compiling metal-oriented code with `portable`

This depends on what the code uses.

- If code stays on portable surfaces, it should compile in `portable`.
- If code relies on target-native metal-only surfaces/patterns, portability is intentionally reduced and portable compilation may fail or require adaptation.

Treat metal-only constructs as deliberate opt-in technical debt against cross-target compatibility.

## Interop and other Haxe targets

If you need interoperability with other Haxe targets, keep shared logic in portable surfaces.

- Shared/core logic: portable.
- Target adapter layer (Go interop, target-native APIs): metal-capable boundary modules.

This split preserves cross-target leverage while allowing Go-native power where needed.

## Authoring guidance when starting a new Go project

1. Build domain logic in portable Haxe first.
2. Keep interop behind explicit adapter modules.
3. Turn on strict modes early (`reflaxe_go_strict`, `reflaxe_go_strict_examples`) to avoid raw injection drift.
4. Add snapshots + semantic-diff coverage before profile-sensitive refactors.
5. Move to metal-oriented APIs only with measured perf or interoperability requirements.

## Portable to metal migration checklist

1. Confirm semantic-diff suite is green in `portable`.
2. Switch profile to `metal` in CI for one target/app lane.
3. Fix strict-boundary violations using typed facades, not raw app-side injection.
4. Benchmark before/after (`npm run test:perf:go`, `npm run test:perf:apps`).
5. Keep portable-compatible modules untouched unless data proves metal-only specialization is needed.

## Metal back to portable checklist

1. Remove/replace metal-only API usage in shared modules.
2. Re-run snapshots and semantic-diff in portable profile.
3. Re-check examples matrix and support matrix expectations.
4. Document any remaining target-native islands explicitly.

## Generated code expectations by profile

- `portable`: readability is acceptable but may include more runtime helper calls to preserve semantics.
- `metal`: aim for hand-written-Go-like shape in typed lanes, while still preserving correctness contracts.

If metal output degrades readability without measurable benefit, treat it as a codegen quality issue and add snapshot/perf evidence.

## Contract guardrails in this repo

Use these gates when profile behavior changes:

- `python3 test/run-snapshots.py`
- `python3 test/run-semantic-diff.py`
- `python3 test/run-ci.py`
- `python3 test/run-examples.py`

Canonical portable semantics reference: `docs/portable-canonical-contract.md`.

## Family template requirements

Any sibling compiler reusing this model should document the same sections:

1. Profile definitions and contract boundaries.
2. Semantic guarantees vs codegen tendencies.
3. Profile choice decision guide.
4. Portable<->metal migration rules.
5. Cross-target interoperability strategy.
6. Explicit test gates that enforce the contract.
