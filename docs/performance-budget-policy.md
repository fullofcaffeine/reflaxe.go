# Performance Budget Policy

This page explains how `reflaxe.go` treats benchmark output during production
readiness checks.

Performance numbers are useful, but they are noisy. A single benchmark run can
move because of machine load, cache warmth, Go version, Haxe version, or a small
OS scheduling difference. The policy below keeps real regressions visible
without pretending that every warning is automatically a release blocker.

## Terms

- **baseline**: the checked-in JSON file used as the comparison point for a
  harness run.
- **soft warning**: a visible `::warning::` annotation and artifact entry. It
  asks a human to inspect the change, but it does not fail CI by itself.
- **hard gate**: a budget check that fails CI when the configured threshold is
  exceeded.
- **pure-Go baseline**: handwritten Go code that runs the same workload as a
  generated Haxe-to-Go example. It helps answer "how close are we to ordinary
  Go?"
- **portable-vs-metal delta**: the ratio between default `portable` output and
  opt-in `metal` output. It helps answer "how much performance headroom remains
  between portable code and Go-first code?"

## Current Release Policy

| Area | Default release decision | Why |
| --- | --- | --- |
| Portable regressions | Portable regressions stay warning-only by default. | `portable` is the semantic baseline. A single noisy run should not block a release unless repeated evidence shows user-visible drift. |
| Metal microbenches | Metal microbench regressions are release-blocking in CI. | `metal` is the Go-first performance lane, so regressions there are higher signal. |
| Flagship app metal metrics | Flagship app metal regressions stay warning-only by default. | App-level benchmarks include more moving parts and should be promoted to hard gates only after repeated stable drift. |
| Portable-vs-metal deltas | Delta regressions stay warning-only by default. | Delta gates are useful trend signals, but we should not make them release-blocking until the selected cases are stable across CI runs. |
| HXRT selective runtime | HXRT selective runtime drift is release-blocking in CI. | Runtime slicing is expected to preserve size/perf wins once enabled, so drift is a concrete release risk. |

Do not update perf baselines just to make warnings disappear. Update a baseline
only when the code change is intentional, the run was collected under the
documented methodology, and the commit explains why the new numbers are the
right comparison point.

## Harnesses

| Harness | Command | Baseline | Current hard gates |
| --- | --- | --- | --- |
| Go profile microbench | `npm run test:perf:go` | `scripts/ci/perf/go-profile-baseline.json` | `GO_PERF_ENFORCE_METAL_BUDGET=1` in CI |
| Flagship app perf | `npm run test:perf:apps` | `scripts/ci/perf/app-profile-baseline.json` | none by default; emits warning and hard-candidate artifacts |
| HXRT selective runtime | `npm run test:perf:hxrt-selective` | `scripts/ci/perf/hxrt-selective-baseline.json` | `GO_HXRT_SLICE_ENFORCE=1` in CI |

Important knobs:

- `GO_PERF_ENFORCE_METAL_BUDGET`: fail Go profile microbench CI when metal
  metrics exceed the configured hard budget.
- `GO_PERF_ENFORCE_DELTA_BUDGET`: fail Go profile microbench CI when selected
  portable-vs-metal deltas exceed the configured hard budget.
- `GO_APP_PERF_ENFORCE_METAL_BUDGET`: fail flagship app CI when metal app
  metrics exceed the configured hard budget.
- `GO_APP_PERF_ENFORCE_DELTA_BUDGET`: fail flagship app CI when selected
  portable-vs-metal deltas exceed the configured hard budget.
- `GO_HXRT_SLICE_ENFORCE`: fail HXRT selective runtime CI on configured
  source/binary drift.

## When To Promote A Warning To A Hard Gate

Promote a warning-only budget to a hard gate only when all of these are true:

1. The same warning repeats across stable CI runs.
2. The affected metric maps to user-visible behavior, such as startup time,
   binary size, throughput, latency, or allocation pressure.
3. The selected case has low enough variance that a hard gate will not be flaky.
4. The new threshold and selected cases are documented in this file and in the
   harness configuration.

This keeps the project honest: warnings stay useful signals, hard gates stay
trustworthy, and baseline updates remain explicit engineering decisions.
