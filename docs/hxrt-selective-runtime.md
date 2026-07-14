# Selective `hxrt` Runtime Plan

## Decision Summary

Selective `hxrt` is a runtime-packaging policy. It is orthogonal to source
semantics and to the compatible `portable|metal` policy presets.

- Portable semantics remain the default product contract.
- Typed APIs/externs and `@:goNative` declare Go-native source boundaries.
- Selective `hxrt` minimizes copied runtime support.

These concerns should evolve together without collapsing into one toggle.

## Why runtime slicing is separate

`metal + selective hxrt` reduces emitted runtime footprint, but it does not answer:

- what semantic compatibility baseline we guarantee,
- what strict-boundary behavior is enforced by default,
- what optimization/interoperability behaviors are opt-in vs guaranteed.

Portable Haxe remains the semantic baseline. Typed native APIs and modules
remain the interoperability boundary. `metal` remains a supported compatibility
preset for stricter/eager defaults.

## Goals

1. Keep Haxe source portability guarantees explicit (`portable` baseline).
2. Make generated output as close as possible to handwritten Go wherever the
   selected source semantics permit it.
3. Trim runtime footprint via deterministic feature inference + override controls.

## Implementation Tracks

1. Runtime slice split:
   - Split `runtime/hxrt/hxrt.go` into `runtime/hxrt/*.go` feature groups.
   - Keep package/API stable.
2. Feature inference:
   - Infer required runtime features from used module/type surfaces and compiler shim groups.
   - Keep ordering deterministic and dependency-complete.
3. Selective runtime emit:
   - Copy only required runtime files when selective mode is enabled.
   - Preserve full-copy behavior for compatibility and fallback.

## Define Matrix

- `reflaxe_go_hxrt_default_features`
  - Force full runtime copy (compat mode).
  - Takes precedence over selective runtime flags.
- `reflaxe_go_hxrt_features=core,json,sys,ssl,...`
  - Enables selective runtime mode and adds manual feature list.
  - Use empty value (`-D reflaxe_go_hxrt_features=`) to enable selective mode with inferred-only features.
- `reflaxe_go_hxrt_no_feature_infer`
  - Enables selective runtime mode and disables inference (use core + manual only).

## Rollout Policy

Phase 1:
- Keep full runtime copy as default.
- Add selective mode behind defines + tests (`reflaxe_go_hxrt_features` and/or `reflaxe_go_hxrt_no_feature_infer`).

Phase 2:
- Promote selective mode based on coverage/perf evidence, independently of the
  selected compatibility preset.
- Keep explicit full-copy fallback for debugging and migrations.

## Perf/Size Harness

Run selective-vs-full runtime footprint metrics:

```bash
bash scripts/ci/perf-hxrt-selective.sh
```

Optional hard budget enforcement (used in CI):

```bash
GO_HXRT_SLICE_ENFORCE=1 bash scripts/ci/perf-hxrt-selective.sh
```

Regenerate baseline:

```bash
bash scripts/ci/perf-hxrt-selective.sh --update-baseline
```

Artifacts:

- `.cache/perf-hxrt-selective/results/current.json`
- `.cache/perf-hxrt-selective/results/comparison.json`
- `.cache/perf-hxrt-selective/results/summary.md`
- `scripts/ci/perf/hxrt-selective-baseline.json`

Interpretation:

- `file_delta`: selective runtime files minus full runtime files (should be `<= 0`).
- `source_delta_pct`: selective runtime source bytes change vs full runtime (should be `<= 0` in normal cases).
- `binary_delta_pct`: selective binary bytes change vs full runtime (expected near 0; a small positive drift can happen by case/toolchain).
- `drift` columns in `summary.md`: delta vs `scripts/ci/perf/hxrt-selective-baseline.json` to track trend over time.

## Related Beads

- Epic: `haxe.go-e73`
- Canonical sequencing is tracked in bead dependencies under that epic.
