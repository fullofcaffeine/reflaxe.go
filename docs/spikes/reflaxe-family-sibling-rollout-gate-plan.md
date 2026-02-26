# `reflaxe.family.std` Sibling Rollout Gate Plan

Owner issue: `haxe.go-cgk.18`  
Scope: ordered adoption plan for `haxe.rust`, `hxhx/reflaxe.ocaml`, and `haxe.elixir.codex` after `haxe.go` dual-run stabilization.

## Goal

Adopt the shared portable contract package across sibling compilers without breaking each backend's documented semantics.

This plan defines:

1. rollout order,
2. per-repo prerequisites,
3. known semantic deltas,
4. mandatory gates for the first adoption PR in each repo.

## Rollout Order

1. `haxe.rust`
2. `hxhx` / `reflaxe.ocaml`
3. `haxe.elixir.codex`

Rationale:

- `haxe.rust` already has the closest architecture match for explicit `portable|metal`, strict boundaries, and runtime feature planning.
- `hxhx/reflaxe.ocaml` is also contract-aligned but has a distinct compiler architecture and runtime module selection strategy that should be validated after rust.
- `haxe.elixir.codex` needs the largest contract formalization step (portable vs elixir-first currently documented as style, not strict contract mode), so it should adopt last.

## Global Prerequisites Before Any Sibling PR

1. `haxe.go` dual-run pin remains green:
   - `npm run test:family-stdlib-sync`
   - `npm run test:family-stdlib-bootstrap`
2. `haxe.go` portable parity closure artifacts are stable and reviewed:
   - `test/.test-cache/portable_parity_closure_summary.json`
   - `test/.test-cache/family_std_dual_run_report.json`
3. Family package contract artifacts are versioned and pinned:
   - `family/reflaxe.family.std/VERSION`
   - `family/family_std_pin.json`
4. No open blocker in `docs/spikes/reflaxe-family-stdlib-execution-checklist.md` that invalidates extraction sequencing.

## Per-Repo Adoption Matrix

### `haxe.rust` (Wave 1)

Prerequisites:

1. Keep profile contract explicit (`portable|metal`) and preserve existing `metal` fallback controls.
2. Keep existing runtime feature inference/reporting operational while wiring shared portable artifacts.

Known semantic deltas to track:

1. Metal defaults and representation choices (for example string/nullability defaults) are not identical to portable and must remain documented.
2. `rust_no_hxrt` and no-runtime lanes remain rust-specific and out of shared portable contract scope.

Required first adoption PR gates:

1. Shared Tier1 portable conformance passes with rust backend.
2. Existing rust profile contract and metal restrictions tests stay green.
3. Runtime plan report still emits deterministic artifacts with pinned family version.

### `hxhx` / `reflaxe.ocaml` (Wave 2)

Prerequisites:

1. Preserve explicit `portable|metal` profile contract and metal verifier behavior.
2. Maintain runtime module-copy flow while introducing shared portable allowlist/conformance assets.

Known semantic deltas to track:

1. Runtime module selection currently relies on token-level inference and may require explicit allow/force controls.
2. OCaml-specific runtime representation choices (for example null/string/env semantics) must remain documented as target-specific deltas.

Required first adoption PR gates:

1. Shared Tier1 portable conformance passes with OCaml backend.
2. Existing profile verifier tests and fixture harness remain green.
3. Runtime selection report (or equivalent deterministic artifact) is available in CI.

### `haxe.elixir.codex` (Wave 3)

Prerequisites:

1. Introduce formal profile contract selection (`portable|metal` or equivalent explicit contract names) instead of style-only guidance.
2. Keep strict boundary enforcement available for app/examples and define fallback policy for migration.

Known semantic deltas to track:

1. BEAM-oriented interop/idiomatic transforms may not map 1:1 to portable contract defaults.
2. Existing feature-flag surface is broad; profile contract presets must avoid semantic ambiguity.

Required first adoption PR gates:

1. Shared Tier1 portable conformance passes with Elixir backend.
2. Portable semantic subset gate exists and is CI-enforced.
3. Profile contract docs clearly separate portable semantics vs target-native interop surfaces.

## Adoption PR Checklist (Each Sibling Repo)

1. Add/refresh family pin metadata (version + source + migration mode).
2. Add deterministic sync/verify command for local canonical artifacts vs pinned family assets.
3. Wire CI gates:
   - portable allowlist gate,
   - Tier1 conformance gate,
   - governance/provenance gate (or equivalent),
   - profile/contract regression tests.
4. Document known semantic deltas that remain intentionally target-specific.
5. Publish generated reports/summary artifacts for reviewer inspection.

## Rollback Rule

If first adoption PR in a sibling repo fails any mandatory gate after merge, revert to local canonical artifacts and reopen the sibling rollout task with blocker details before continuing to the next wave.
