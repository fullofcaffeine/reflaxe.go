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
| `metal` | Go-native control with strict boundaries | Lower by design when using Go-native surfaces | Typed Go-facing APIs and stricter escape-hatch policy | More typed specialization and direct Go-shaped output in supported lanes |

## What differs today (practical reality)

If your code stays fully on portable surfaces, outputs can look similar across profiles. That is intentional.

Current concrete differences:

1. Boundary defaults:
   `metal` enables stricter app-side injection policy by default.
2. Native-surface lane:
   `metal` is the intended lane for explicit Go-native APIs and interop patterns.
3. Typed specialization focus:
   `metal` is where typed low-level specialization work is prioritized (`go.Chan<T>`, `go.Slice<T>`, `go.Map<K,V>`, `go.Result<T>` lanes).
4. Portability posture:
   `portable` is the cross-target baseline contract; `metal` accepts lower portability when you opt into native-first surfaces.

## Concrete examples (where you can see the difference)

### Example A: raw native injection boundary

```haxe
class Main {
  static function main() {
    untyped __go__("fmt.Println(\"hi\")");
  }
}
```

- `portable`: may compile when strict mode is not enabled (still discouraged for maintainability).
- `metal`: fails by default because metal enables strict app-boundary policy.

This is a real profile difference today.

### Example B: portable semantics path

```haxe
class Main {
  static function main() {
    var n:Dynamic = null;
    Sys.println(Std.string(n));
  }
}
```

- `portable`: prints `"null"`.
- `metal` (portable surfaces only): also prints `"null"`.

This is an intentional “same semantics” case.

### Example C: metal-intended native lane

```haxe
import go.Go;

class Main {
  static function main() {
    var ch = Go.newChan<Int>();
    Go.spawn(() -> ch.send(1));
    Sys.println(ch.recv());
  }
}
```

- `metal`: intended lane for this style, with typed specialization work prioritized here.
- `portable`: can compile on Go target, but this is target-native code and not part of cross-target portable expectations.

### Example D: portable build with a metal lane island

```haxe
@:goMetal
class Adapter {
  static function boot() {
    // raw __go__ is forbidden in this lane when contract=portable
    // with reflaxe_go_auto=auto_strict, non-monomorphizable
    // go.Result/go.Chan/go.Slice/go.Map calls are also forbidden
  }
}
```

`@:goMetal` lets you mark incremental migration islands inside a portable build. Those modules get metal-clean enforcement rules even when global contract stays portable.
Typed fallback fail-fast in lane modules is enabled under `-D reflaxe_go_auto=auto_strict`; under `off|auto` those fallback paths are allowed and reported.

Typed-lowering eligibility is centralized in `src/reflaxe/go/compiler/GoMetalTypeEligibility.hx`:
- concrete non-`any` Go type is required,
- nullable primitive dynamic-path representations are excluded for semantic safety,
- `go.Map` key specialization additionally requires a comparable Go key type.

Lane enforcement snapshots:
- `test/snapshot/negative/go_metal_lane_injection`
- `test/snapshot/negative/go_metal_lane_fallback_result`
- `test/snapshot/negative/go_metal_lane_fallback_chan`
- `test/snapshot/negative/go_metal_lane_fallback_slice`
- `test/snapshot/negative/go_metal_lane_fallback_map`
- `test/snapshot/negative/go_metal_lane_fallback_map_noncomparable_key`
- `test/snapshot/core/go_metal_lane_nonlane_fallback_allowed`
- `test/snapshot/core/go_metal_lane_fallback_allowed_off`

## What changes and what does not

### Non-negotiable

- Portable semantics are the baseline contract.
- Runtime feature inference/slicing is orthogonal and must not silently change profile semantics.

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
3. If typed specialization cannot be monomorphized, use `-D reflaxe_go_metal_allow_fallback` to allow fallback (otherwise metal hard-errors by default).
4. If you need raw `__go__`, prefer typed framework facades; explicit `-D reflaxe_go_strict` still forbids raw injection.

Stay in portable semantics:

1. Keep code on portable Haxe stdlib/application APIs.
2. Avoid target-native-only surfaces in shared/core modules.
3. Use `-D reflaxe_go_portable_native_policy=error` in CI/release to hard-fail accidental `go.*` usage in portable modules (`warn` is the local default).
4. If you want deterministic import/using scanner detection (instead of typed traversal), set `-D reflaxe_go_portable_native_scan_mode=scanner` (`hybrid` unions both sources).
4. You can still compile with `metal` for boundary/perf validation without changing semantics if those modules remain portable-surface-only.

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

## Portable convergence optimizer controls

These controls are additive and do not redefine profile semantics:

- `-D reflaxe_go_auto=off|auto|auto_strict` (explicit planner mode, default `off`)
- in `portable`, `auto|auto_strict` currently attempt typed `go.Slice` / `go.Map` / `go.Result` lowerings (outcomes recorded in contract reports)
- `-D reflaxe_go_opt=portable_fast|none` (default `portable_fast`)
- `-D reflaxe_go_opt_go_concurrency_fastpath=...` (typed portable concurrency fastpath capability)
- `-D reflaxe_go_optimizer_plan_report` (emits deterministic optimizer plan artifacts)

Use them to tune portable performance convergence without switching semantic contract.

Optimizer-plan concurrency counters are source-aware: typed fastpath hits/fallbacks count user/app lowering sites and intentionally exclude framework-internal (`std/`, `src/go/`, `src/reflaxe/`) emission paths.

## `auto` direction

`auto` is an explicit additive planner, not a hidden semantic profile.

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

## Portable to metal admission criteria (pass/fail)

### Pass criteria

1. Contract gates are green in CI:
   - `python3 test/run-semantic-diff.py --suite lanes`
   - `python3 test/run-metal-example-boundary.py`
   - `python3 test/run-snapshots.py --case core/report_artifacts_lane_fallback`
2. Strict boundary policy stays enabled for app code (`reflaxe_go_strict_policy=auto|on`; no raw app-side `__go__`).
3. `profile_contract.json` shows deterministic fallback diagnostics (`metalFallbackViolations`, lane/non-lane counts).
4. Perf evidence exists for promoted modules (`npm run test:perf:go` and/or `npm run test:perf:apps`) with documented rationale.

### Fail criteria

1. Lane semantic-diff or metal boundary gates fail.
2. Fallback diagnostics are missing, unstable, or only visible through ad-hoc local commands.
3. Promotion relies on raw app-side `__go__` instead of typed facades.
4. Profile switch is made without benchmark evidence or without documenting accepted portability tradeoffs.

## Metal back to portable checklist

1. Remove/replace metal-only API usage in shared modules.
2. Re-run snapshots and semantic-diff in portable profile.
3. Re-check examples matrix and support matrix expectations.
4. Document any remaining target-native islands explicitly.

## Generated code expectations by profile

- `portable`: readability is acceptable but may include more runtime helper calls to preserve semantics.
- `metal`: aim for hand-written-Go-like shape in typed lanes, while still preserving correctness contracts.

If metal output degrades readability without measurable benefit, treat it as a codegen quality issue and add snapshot/perf evidence.

## Build report artifacts

Use these optional flags when auditing effective contract/runtime behavior:

- `-D reflaxe_go_contract_report` emits `profile_contract.json` and `profile_contract.md`.
- `-D reflaxe_go_runtime_plan_report` emits `hxrt_plan.json` and `hxrt_plan.md`.

`profile_contract.json` (schema v7) carries centralized analyzer diagnostics (`contractDiagnosticCount`, `contractDiagnostics`), portable native scan summary fields (`portableNativeImportScanMode`, `portableNativeImportHitCount`, `portableNativeImportHits`, `portableNativeImportTypedHitCount`, `portableNativeImportTypedHits`, `portableNativeImportScannerHitCount`, `portableNativeImportScannerHits`), deterministic lowering-decision ledger fields (`loweringDecisionCount`, `loweringDecisionAttemptCount`, `loweringDecisionSuccessCount`, `loweringDecisionFallbackCount`, `loweringDecisions`), includes `autoLoweringMode`, and preserves module/lane attribution (`module`, `inMetalLane`) plus lane summary fields:

- `metalFallbackLaneViolationCount`
- `metalFallbackNonLaneViolationCount`
- `metalFallbackViolationsByModule`

Reference snapshots:
- `test/snapshot/core/report_artifacts_basic`
- `test/snapshot/core/report_artifacts_lane_fallback`

## Contract guardrails in this repo

Use these gates when profile behavior changes:

- `python3 test/run-snapshots.py`
- `python3 test/run-semantic-diff.py`
- `python3 test/run-semantic-diff.py --suite lanes`
- `npm run test:semantic-diff:lanes`
- `npm run test:auto-planner:schema`
- `npm run test:perf:go`
- `python3 test/run-ci.py`
- `python3 test/run-ci.py --force-semantic-diff-lanes`
- `python3 test/run-examples.py`

CI auto-planner schema stage controls:

- `python3 test/run-ci.py --skip-auto-planner-schema`
- `python3 test/run-ci.py --force-auto-planner-schema`

Canonical portable semantics references:
- `docs/portable-canonical-contract.md`
- `docs/portable-semantics-v1.md`

## Family template requirements

Any sibling compiler reusing this model should document the same sections:

1. Profile definitions and contract boundaries.
2. Semantic guarantees vs codegen tendencies.
3. Profile choice decision guide.
4. Portable<->metal migration rules.
5. Cross-target interoperability strategy.
6. Explicit test gates that enforce the contract.
