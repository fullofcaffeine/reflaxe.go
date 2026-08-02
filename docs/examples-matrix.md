# Examples Matrix

`reflaxe.go` ships canonical examples designed to show portable-first semantics,
explicit Go-native source boundaries, and policy-preset differences where they
are technically meaningful.

`metal` is not required for good Go output or Go-native APIs. Dual-preset
examples exist where comparing guarded/proven/allow with
explicit/eager/error/strict defaults produces useful policy evidence.

Executable tier, product-surface, profile, and independent expected-output
provenance are enforced by `examples/qa-manifest.json`; see the
[examples QA contract](examples-qa-contract.md). The labels below describe the
teaching purpose and do not merge portable, native, runtime, or package claims.

The word **release-bearing** means a successful lane may support only the exact
portable-beta operations it names. It is stricter than ordinary example QA.

| Example | portable | metal | Release-bearing? | Purpose |
| --- | --- | --- | --- | --- |
| `examples/portable_beta` | Yes | No | Yes, exact operation IDs only | Small real compiler/build/run path limited to beta-admitted language, data, reflection, file, and filesystem behavior. |
| `examples/profile_storyboard` | Yes | No | No | Portable-first release dashboard reference. |
| `examples/tui_todo` | Yes | No | No | Portable-first deterministic CLI app reference. |
| `examples/incident_api` | Yes | Yes | No | Runnable loopback HTTP service using Haxe stdlib sockets, JSON, config files, and file-backed state. It intentionally avoids `go.*`, Go externs, and raw `__go__`; `metal` is a preset-audit run, not a different app implementation. |
| `examples/interop_smoke` | Yes | Yes | No | Typed interop smoke reference for `@:go.import`, `@:go.name`, `@:go.receiver`, `@:go.valueError` (`(T,error)` -> `go.Result<T>`), and package APIs (`fmt`/`time`/`context`/`net/http`/`strconv`). This app is intentionally preset-neutral, so generated Go is expected to be near-identical across presets. |
| `examples/worker_pool_select` | Yes | Yes | No | Deterministic worker pool with channel fan-out plus typed `go.Select` helper flows (`recv`/`recv2`/`send`/`send2`). |
| `examples/pulseforge` | Yes | Yes | No | Flagship app scaffold proving profile matrix + explicit variant lanes (`core` via `*.hxml`, `go_native` via `*.ci.hxml`). |
| `examples/fluxproxy` | Yes | Yes | No | Flagship proxy scaffold with profile matrix + variant lanes (`core` via `*.hxml`, `go_native` via `*.ci.hxml`). |

A green non-release-bearing example still proves the behavior written in its
README and expected output. It cannot add APIs to the public support matrix.

Portable-only examples are intentional: if a second preset adds no useful
policy or benchmark evidence, we avoid synthetic duplication.

## Terms

- `go_native` (app variant): compile-time runtime-adapter lane used by an example app for Go-first execution paths (for example worker/channel/select flows). It is not a compiler profile.
- `Go-native`: APIs or runtime behavior tied specifically to Go, such as `go.Chan`, `go.Select`, or typed Go extern metadata.
- `hot path`: the frequently-executed part of an app where optimization has the biggest runtime impact.

## Native-adapter collection purity gates (legacy command names)

Dual-profile examples use two collection-purity gates:

- Hard boundary gate:
  - `examples/*/app/runtime/GoNativeRuntime.hx`
- Full-tree gate: audits `haxe.ds.*` imports across all example modules.

Commands:

```bash
# hard gate
python3 test/run-metal-example-boundary.py

# full-tree audit report
python3 test/run-metal-example-boundary.py --scope full --mode audit --report .cache/metal-example-boundary/full-scope-audit.json

# full-tree threshold gate (fails if violation count exceeds threshold)
python3 test/run-metal-example-boundary.py --scope full --mode audit --report .cache/metal-example-boundary/full-scope-audit.json --max-violations 0
```

CI:

- `ci-harness.yml` quality job runs hard gate via `npm run test:ci`.
- the same job enforces full-scope threshold mode by default:
  - `GO_METAL_COLLECTION_AUDIT_ENFORCE=1`
  - `GO_METAL_COLLECTION_AUDIT_MAX=0`
- the job also uploads the full-scope audit artifact (`metal-collection-audit`).

Allowlist rule:

- if temporary exceptions are needed, keep them file-specific, justified, and time-bounded.
- do not use broad `examples/**` exception patterns.

Policy spike:
- `docs/spikes/metal-build-collection-purity-policy.md`

## Preset performance teaching contract

Examples intentionally show two valid outcomes:

- parity cases: when code stays on portable-facing surfaces, `portable` and
  `metal` must match behavior and can be near-identical in code shape;
- specialization-delta cases: typed Go-native hot paths measure the effect of
  proven versus eager specialization without promising that one preset is
  always faster.

Use these anchors:

- parity-focused: `examples/incident_api`, `examples/interop_smoke`, `core` lanes in `examples/pulseforge` and `examples/fluxproxy`, plus portable-only references (`examples/tui_todo`, `examples/profile_storyboard`).
- policy-delta-focused: `examples/worker_pool_select`, `go_native` variants in
  `examples/pulseforge` and `examples/fluxproxy`.

Evidence sources:

- micro profile baselines: `scripts/ci/perf/go-profile-baseline.json`
- flagship app baselines: `scripts/ci/perf/app-profile-baseline.json`
- methodology and delta interpretation: `docs/benchmark-methodology-apps.md`

## Flagship app behavior docs

- Incident API service: `examples/incident_api/README.md`
- PulseForge: `examples/pulseforge/README.md`
- FluxProxy: `examples/fluxproxy/README.md`
- Planning + scope contract: `docs/flagship-apps-plan.md`
- Benchmark fairness + reproducibility rules: `docs/benchmark-methodology-apps.md`

## Pure-Go parity baselines

| Baseline | Status | Purpose |
| --- | --- | --- |
| `benchmarks/pure_go/pulseforge` | Ready | Handwritten Go parity baseline for PulseForge workload/contract and benchmark comparison against generated portable/metal outputs. Does **not** use `hxrt`. |
| `benchmarks/pure_go/fluxproxy` | Ready | Handwritten Go parity baseline for FluxProxy workload/contract and benchmark comparison against generated portable/metal outputs. Does **not** use `hxrt`. |

Run:

```bash
npm run test:pure-go:pulseforge
npm run test:pure-go:fluxproxy
npm run test:perf:pure-go:pulseforge
npm run test:perf:pure-go:fluxproxy
```

## Flagship app perf harness

Run the full app-level harness:

```bash
npm run test:perf:apps
# regenerate baseline ratios
npm run test:perf:apps:update-baseline
# optional CI parity: enforce metal hard budgets
GO_APP_PERF_ENFORCE_METAL_BUDGET=1 npm run test:perf:apps
```

This covers both flagship apps across:

- generated Haxe lanes: `portable/core`, `portable/go_native`, `metal/core`, `metal/go_native`
- handwritten Go parity lanes: `pure_go/core`, `pure_go/go_native`

Terminology:

- `pure_go`: handwritten baseline modules under `benchmarks/pure_go/*` (no `hxrt`)
- `go_native`: runtime adapter variant in the app codebase enabling Go-first lane behavior (`chan/select` worker paths) while preserving the same domain contract

Outputs:

- `.cache/perf-apps/results/current.json`
- `.cache/perf-apps/results/comparison.json`
- `.cache/perf-apps/results/summary.md`
- `.cache/perf-apps/results/raw_metrics.tsv`
- `.cache/perf-apps/results/warnings.txt`
- `.cache/perf-apps/results/hard_failures.txt`
- `scripts/ci/perf/app-profile-baseline.json`

Metrics reported per app/preset/variant lane (the artifact schema retains the
legacy `profile` field name):

- throughput (`ops/s`)
- latency (`avg`, `p95`, `p99`, milliseconds)
- allocation (`B/op`, `allocs/op`)
- memory (`max RSS`, KB)
- startup (`avg ms`, help-mode loop)
- binary size (raw + stripped)

Execution plan and acceptance gates for flagship examples: `docs/flagship-apps-plan.md`.

## Build and run matrix

```bash
python3 test/run-examples.py
```

QA contract: [docs/examples-qa-contract.md](examples-qa-contract.md).

This runs discovered profile cases per example:

- `compile.<profile>.ci.hxml` + `go test` + `go run` expected output checks
- `compile.<profile>.hxml` + `go test` + `go run` expected output checks
- generated tree drift checks (`generated/<profile>` vs `out_<profile>`) for available profiles in that example

## Generated Go + binaries

Refresh committed generated trees:

```bash
python3 test/run-examples.py --bless-generated
```

Build cross-platform binaries from committed generated trees:

```bash
bash scripts/examples/build-binaries.sh
```

Binary matrix targets:

- `linux/amd64`
- `linux/arm64`
- `darwin/arm64`
- `windows/amd64`

## CI artifact/release automation

- Workflow: `.github/workflows/examples-artifacts.yml`
- Triggers:
  - `push` to `master`: builds matrix and uploads workflow artifacts.
  - `push` tags: builds matrix, uploads workflow artifacts, and publishes release assets.
  - `workflow_dispatch`: manual artifact build/upload.
- Why this uses artifact upload/download:
  - The workflow intentionally splits `build` and `release` jobs so release publishing only runs on tag pushes and with `contents: write` permissions.
  - Job filesystems are isolated in GitHub Actions, so assets must cross jobs via `upload-artifact`/`download-artifact`.
  - Downloaded artifact layout is not guaranteed to preserve the original `dist/...` prefix, so the release job normalizes discovered files into a deterministic staging directory before `action-gh-release`.
- Release assets:
  - `examples-<tag>.tar.gz`
  - `examples-<tag>.tar.gz.sha256`
  - `manifest.json`
  - `checksums.txt`

Flagship app perf artifacts are published by `.github/workflows/ci-harness.yml` job `perf-apps` as `go-app-perf-results`.

## Related docs

- `docs/start-here.md` - first-run setup and the default command flow.
- `docs/profiles.md` - what `portable` and `metal` mean in day-to-day use.
- `docs/profile-semantics-guide.md` - behavior guarantees and contract boundaries.
- `docs/benchmark-methodology-apps.md` - how app/perf comparisons are measured and interpreted.
- `docs/flagship-apps-plan.md` - execution plan and acceptance criteria for flagship examples.
- `docs/glossary.md` - shared terms used across compiler and example docs.
