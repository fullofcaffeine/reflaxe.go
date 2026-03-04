# Spike: Explicit `auto` Mode vs Two-Profile Model

Issue: `haxe.go-pgl`

## Decision

Keep `portable` + `metal` as the only semantic contracts.

Do **not** add `auto` as a third semantic profile.

If we add `auto` later, implement it as an explicit planner/preset mode that is additive-only and deterministic, with semantics still anchored to an explicit contract (`portable` or `metal`).

## Why

`portable` vs `metal` is a semantic boundary, not just a performance preset.

Inference can detect feature usage but cannot infer user intent safely. Making semantics implicit creates hidden flips and brittle CI behavior.

Current repo architecture already follows the right foundation:

- explicit profile resolver and conflict diagnostics (`src/reflaxe/go/ProfileResolver.hx`)
- canonical portable contract (`docs/portable-canonical-contract.md`)
- orthogonal runtime slicing controls (`reflaxe_go_hxrt_*`) that are additive and deterministic

## Option Comparison

| Option | Benefits | Risks | Verdict |
| --- | --- | --- | --- |
| Two profiles only (`portable`,`metal`) | Clear contracts, stable CI matrix, explicit review boundaries | Slightly more upfront user choice | Keep as canonical |
| Add `auto` as semantic profile | Fewer flags for newcomers | Hidden semantic flips, unclear ownership of strictness/runtime/interop policy | Reject |
| Add explicit `auto` planner mode (non-semantic) | Better UX while preserving contracts, additive inference, deterministic reporting | Requires lock/report plumbing and docs discipline | Accept only as future experimental mode |

## Contract for a Future `auto` Planner (if implemented)

`auto` may infer only additive dimensions:

- runtime feature slices (`hxrt` modules)
- optional helper packages/shims
- non-semantic build knobs

`auto` must **not** infer/flip:

- semantic contract (`portable` vs `metal`)
- strictness policy (error vs warning)
- injection boundary policy
- interop permission model

## CI/Test Implications

If a future planner mode is implemented, require:

- deterministic planner report artifact per build
- lockable plan input/output (plan lock file or equivalent)
- no semantic test matrix expansion caused by planner nondeterminism
- explicit CI checks that planner output is stable for unchanged inputs

Required gates (future):

- snapshot + semantic-diff stay profile-contract based
- planner-specific determinism tests compare emitted plan reports
- examples/perf jobs can consume planner reports but must still declare semantic profile contract

## Implemented safety gates (current repo)

The CI harness now includes a dedicated auto-planner schema stage on full runs:

- `npm run test:auto-planner:schema`
  - Validates deterministic report contracts from snapshot artifacts:
    - `profile_contract.json` schema v7 + required planner/native-scan keys
    - `optimizer_plan.json` schema v5 + pass-selection source/reason keys + capability-level auto-lowering summary keys (`autoLoweringCapabilities`)
    - `hxrt_plan.json` schema v1 + runtime-plan core keys
  - Verifies planner vs compatibility fallback sources:
    - planner case: `goAstPassSelectionSource=planner`
    - legacy granular case: `goAstPassSelectionSource=legacy_granular_bundle`

CI wiring:

- `python3 test/run-ci.py` runs auto-planner schema stage by default on full runs.
- `--skip-auto-planner-schema` / `--force-auto-planner-schema` control this stage explicitly.

Related semantic/perf commands for planner safety:

- `npm run test:semantic-diff`
- `npm run test:semantic-diff:optimizer-matrix`
- `npm run test:semantic-diff:lanes`
- `npm run test:perf:go`

## Recommended Future Shape (Low Priority)

1. Keep `reflaxe_go_profile=portable|metal` unchanged.
2. Introduce explicit non-semantic opt-in define (example: `reflaxe_go_auto_plan`).
3. Emit planner report JSON (selected features + reasons + deterministic key ordering).
4. Optionally allow lockfile compare mode in CI.

## Result

This spike recommends **no immediate compiler behavior change**.

Two-profile model remains canonical. Any future `auto` support should be implemented as a non-semantic planning layer, explicitly opted in, experimental, and deterministic.
