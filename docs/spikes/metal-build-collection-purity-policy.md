# Metal Build Collection Purity Policy Spike

Date: 2026-03-04  
Issue: `haxe.go-qzku`

## Question

Should we enforce a strict `haxe.ds.*` ban across every module compiled under `-D reflaxe_go_profile=metal`, or only in explicitly metal-native modules/lanes?

## Current baseline

- Existing CI guard checks only metal-specific example modules:
  - `examples/*/profile/MetalRuntime.hx`
  - `examples/*/app/runtime/GoNativeRuntime.hx`
- Guard currently forbids:
  - `haxe.ds.List`
  - `haxe.ds.IntMap`
- Implementation: `test/run-metal-example-boundary.py`

## Evidence: current `haxe.ds.*` usage in examples

Current imports in example app/domain code (not just metal adapters):

- `examples/tui_todo/InteractiveCli.hx` (`haxe.ds.List`)
- `examples/tui_todo/model/TodoStore.hx` (`haxe.ds.List`)
- `examples/tui_todo/model/TodoItem.hx` (`haxe.ds.List`)
- `examples/tui_todo/app/TodoApp.hx` (`haxe.ds.List`)
- `examples/tui_todo/Harness.hx` (`haxe.ds.List`)
- `examples/profile_storyboard/domain/StoryCard.hx` (`haxe.ds.List`)
- `examples/profile_storyboard/Harness.hx` (`haxe.ds.List`)
- `examples/fluxproxy/app/core/FluxPipeline.hx` (`haxe.ds.IntMap`, `haxe.ds.StringMap`)
- `examples/pulseforge/app/core/PulsePipeline.hx` (`haxe.ds.StringMap`)

## Options evaluated

## 1) Keep current policy (metal-module-only purity)

Enforce collection purity only in modules that explicitly represent metal-native boundaries (profile adapters/runtime adapters/lane modules).

Pros:
- Preserves single-codebase portable/metal examples.
- Keeps enforcement aligned with declared native boundary ownership.
- Low migration cost and low semantic risk.

Cons:
- Full-program metal purity is not guaranteed.

## 2) Enforce full-build purity immediately

Ban `haxe.ds.*` across all modules compiled in `metal`.

Pros:
- Maximum purity signal.
- Forces faster migration to metal-native structures.

Cons:
- High migration cost now (touches core/domain/example harness files).
- Pushes profile divergence early; weakens “same codebase across profiles” teaching contract.
- Risks changing semantics/perf contracts without targeted benchmarks per refactor.

## 3) Stage toward full-build purity (recommended)

Keep strict hard gate for metal-boundary modules now, plus add full-build audit mode (non-blocking) and tighten later only when migration criteria are met.

Pros:
- Preserves current profile teaching model.
- Creates measurable path to deeper purity.
- Avoids forced churn before typed-native replacements are ready and benchmarked.

Cons:
- Requires disciplined staged enforcement.

## Recommendation

Adopt option 3.

Policy:

1. Hard-fail CI for metal boundary modules only (current model, expanded rule coverage).
2. Add full-build audit mode that reports `haxe.ds.*` usage in all modules compiled for metal.
3. Move to full-build hard enforcement only after:
   - designated examples have completed migration plans,
   - semantic-diff/snapshot/perf baselines stay green,
   - docs define clear allowed replacements and exceptions.

## Migration impact estimate

If full-build hard enforcement were enabled now, at least 10 imports across 4 examples would require refactors in core/domain/harness modules.

Expected impact:

- `tui_todo` and `profile_storyboard`: collection model rewrites in domain/harness paths.
- `fluxproxy` and `pulseforge`: map/list structure rewrites in pipeline core.
- Potential divergence pressure between portable-oriented and metal-optimized example code.

Conclusion: immediate full-build hard fail is high-cost for current roadmap stage.

## Proposed CI strategy

## Stage A (now)

- Keep hard gate in `test/run-metal-example-boundary.py` for boundary modules.
- Expand forbidden set to include `haxe.ds.StringMap` in boundary modules as well.

## Stage B (next)

- Add audit mode to same script:
  - Scan full example trees for metal-target builds.
  - Emit deterministic report artifact (path + import + category).
  - Non-blocking by default; optionally blocking in nightly/experimental CI.

## Stage C (later, explicit flip)

- Enable full-build hard fail only for examples that declare “metal-pure certified”.
- Keep per-example allowlist exceptions with explicit owner + reason.

## Allowlist rules (if needed)

Allowlist entries must be explicit and temporary:

- include exact file path,
- include why replacement is not yet safe/practical,
- include removal target milestone/issue.

No glob-wide allowlists for `examples/**`.

## Follow-up implementation tasks

1. Extend metal boundary checker with rule-set config and audit mode.
2. Add CI wiring for boundary hard gate + full-build audit report publication.
3. Add per-example migration tickets (starting with `fluxproxy` and `pulseforge` core pipelines).
4. Update example/profile docs to explain:
   - what is enforced today,
   - what is audit-only,
   - what conditions trigger stricter enforcement.

## Decision summary

- Do not enforce full metal-build `haxe.ds.*` purity yet.
- Keep hard guarantees at explicit metal boundaries.
- Add deterministic full-build auditing now, then tighten in staged, benchmark-backed increments.
