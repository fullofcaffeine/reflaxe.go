# Flagship App Benchmark Methodology

This document defines fairness and reproducibility rules for app-level benchmarking of:

- generated Haxe->Go outputs (`portable` / `metal`, `core` / `go_native`)
- handwritten Go parity baselines (`pure_go`, `core` / `go_native`)

Harness entrypoint: `scripts/ci/perf-apps.sh` (`npm run test:perf:apps`).

## Terms used in this doc

- **Profile** (`portable`, `metal`): compiler contract selected with `-D reflaxe_go_profile=...`. See `docs/profiles.md`.
- **Variant** (`core`, `go_native`): app runtime adapter choice inside one app codebase. This is not a compiler profile. See `docs/examples-matrix.md`.
- **Lane**: one profile+variant combination in benchmark output (for example `portable/core`).
- **pure_go**: handwritten Go reference implementation used as a baseline for ratio comparisons.

Reference glossary: `docs/glossary.md`.

## Scope

The methodology applies to flagship apps:

- `examples/pulseforge`
- `examples/fluxproxy`

And parity modules:

- `benchmarks/pure_go/pulseforge`
- `benchmarks/pure_go/fluxproxy`

## Lane Matrix

Each run captures these lanes for both apps:

| Kind | Profile | Variant |
| --- | --- | --- |
| `haxe` | `portable` | `core` |
| `haxe` | `portable` | `go_native` |
| `haxe` | `metal` | `core` |
| `haxe` | `metal` | `go_native` |
| `pure_go` | `pure` | `core` |
| `pure_go` | `pure` | `go_native` |

`pure_go` is handwritten baseline code and does not use `hxrt`.
`hxrt` is the shared runtime helper package generated projects import. See `docs/hxrt-runtime.md`.

## Workload Shape

Workloads are deterministic and contract-driven:

- PulseForge input shape: `Harness.baselineFrames()` in `examples/pulseforge/Harness.hx`
- FluxProxy input shape: `Harness.baselineRequests()` in `examples/fluxproxy/Harness.hx`
- Pure-Go parity mirrors the same workload and contract keys (`runScripted(...)` in each pure-Go baseline module)

All throughput/latency measurements are based on these fixed scripted workloads.

## Warmup and Steady-State Protocol

Recommended local reproducibility protocol:

1. Warmup run (discard output):
   - `npm run test:perf:apps`
2. Steady-state runs (record output):
   - run `npm run test:perf:apps` at least 3 times
   - compare `summary.md`/`comparison.json`; use median or majority trend for conclusions
3. Baseline refresh (only when intentionally re-baselining):
   - `npm run test:perf:apps:update-baseline`

Why: first run includes additional cache/build effects (Go/Haxe/module caches). Steady-state runs reduce cache/JIT/toolchain noise in comparisons.

## Toolchain Pinning

Use pinned toolchains when comparing runs:

- Haxe: `4.3.7` (CI pinned)
- Go: CI pinned in `.github/workflows/ci-harness.yml` (`GO_VERSION`)

Record versions in every analysis:

```bash
haxe --version
go version
```

The harness writes toolchain versions into `.cache/perf-apps/results/current.json`.

## Metrics and Interpretation

Per-lane metrics:

- throughput: `ops/s` from scripted contract counters
- latency: `avg`, `p95`, `p99` (ms) from repeated scripted runs
- allocation: `B/op`, `allocs/op` from `go test -benchmem`
- memory: max RSS (`KB`) from `/usr/bin/time`
- startup: avg ms from repeated `help` command runs
- size: raw + stripped binary bytes

Derived comparison uses `haxe_vs_pure` ratios (same app + variant):

- throughput ratio: higher is better
- latency/alloc/rss/startup/size ratios: lower is better

The harness also derives `portable_vs_metal` deltas per app+variant:

- throughput delta (`portable/metal`): higher is better
- latency/alloc/rss/startup/size deltas (`portable/metal`): lower is better

Reading tip:

- Use `haxe_vs_pure` to answer "how close are generated outputs to handwritten Go?"
- Use `portable_vs_metal` to answer "how much Go-native convergence headroom
  remains?" `metal` is not required for good Go output; the delta tells us
  where portable codegen can still become more Go-shaped without changing Haxe
  semantics.

## Baselines and Budgets

Baseline file:

- `scripts/ci/perf/app-profile-baseline.json`

Update mode:

- `npm run test:perf:apps:update-baseline`

Compare mode:

- `npm run test:perf:apps`

Default warning budgets (configurable via env):

- throughput drop: `GO_APP_PERF_THROUGHPUT_WARN_PCT` (default `12`)
- latency rise: `GO_APP_PERF_LATENCY_WARN_PCT` (default `12`)
- alloc rise: `GO_APP_PERF_ALLOC_WARN_PCT` (default `15`)
- memory rise: `GO_APP_PERF_MEMORY_WARN_PCT` (default `12`)
- startup rise: `GO_APP_PERF_STARTUP_WARN_PCT` (default `15`)
- size rise: `GO_APP_PERF_SIZE_WARN_PCT` (default `8`)

Optional metal hard-gate:

- enable with `GO_APP_PERF_ENFORCE_METAL_BUDGET=1`
- hard thresholds:
  - `GO_APP_PERF_METAL_THROUGHPUT_FAIL_PCT`
  - `GO_APP_PERF_METAL_LATENCY_FAIL_PCT`
  - `GO_APP_PERF_METAL_ALLOC_FAIL_PCT`
  - `GO_APP_PERF_METAL_MEMORY_FAIL_PCT`
  - `GO_APP_PERF_METAL_STARTUP_FAIL_PCT`
  - `GO_APP_PERF_METAL_SIZE_FAIL_PCT`

Optional portable-vs-metal delta hard-gate:

- enable with `GO_APP_PERF_ENFORCE_DELTA_BUDGET=1`
- selectors/thresholds:
  - `GO_APP_PERF_DELTA_CASES` (comma-separated `app:variant`, or `all`)
  - `GO_APP_PERF_DELTA_WARN_PCT`
  - `GO_APP_PERF_DELTA_FAIL_PCT`

## CI Cadence

CI stage: `.github/workflows/ci-harness.yml` job `perf-apps`.

Cadence:

- pull requests: runs comparison mode and uploads artifacts
- pushes to `master`: runs comparison mode and uploads artifacts
- manual runs: `workflow_dispatch`
- scheduled run: weekly cron in `ci-harness.yml`

Artifacts published by the job:

- `.cache/perf-apps/results/current.json`
- `.cache/perf-apps/results/comparison.json`
- `.cache/perf-apps/results/summary.md`
- `.cache/perf-apps/results/raw_metrics.tsv`
- `.cache/perf-apps/results/warnings.txt`
- `.cache/perf-apps/results/hard_failures.txt`

`comparison.json` includes independent counts for metal hard-fail and delta hard-fail candidates.

## Fairness Constraints

For valid run-to-run comparisons:

- Use the same machine/hardware class for baseline and compare runs.
- Keep background load low (avoid concurrent heavy builds/tests).
- Use the same harness run parameters (`GO_APP_PERF_*` env vars).
- Do not compare reduced-loop local smoke runs against production baselines.
- Interpret warning/hard-fail output as regression signals, not absolute truth; confirm with repeated steady-state runs.

## Artifacts Contract

Harness outputs:

- `.cache/perf-apps/results/current.json`
- `.cache/perf-apps/results/comparison.json`
- `.cache/perf-apps/results/summary.md`
- `.cache/perf-apps/results/raw_metrics.tsv`
- `.cache/perf-apps/results/warnings.txt`
- `.cache/perf-apps/results/hard_failures.txt`
