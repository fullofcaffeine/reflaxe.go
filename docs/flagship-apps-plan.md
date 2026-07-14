# Flagship Apps Plan: PulseForge + FluxProxy

Terms used in this plan:

- **Compatibility selector**: `portable` or `metal`, each expanding to a policy
  preset documented in `docs/profiles.md`.
- **Variant**: app runtime adapter path (`core` or `go_native`) inside one shared
  app codebase. This is a source/API choice, not a compiler profile.
- **core variant**: lowest-risk implementation path that avoids target-specific `go.*` API dependence where possible.
- **go_native variant**: Go-first runtime adapter path that intentionally uses typed `go.*` surfaces.

Reference glossary: `docs/glossary.md`.

## Decision

Build both apps, with strict sequencing:

1. `PulseForge` first (flagship "wow" demo + primary benchmark).
2. `FluxProxy` second (network-heavy focused benchmark).

Both must be usable and runnable, not design-only artifacts.

## Usability Contract (non-negotiable)

Each app must ship with:

- `app/` Haxe source that compiles with both compatibility selectors.
- `compile.portable.hxml`, `compile.metal.hxml`.
- `compile.portable.ci.hxml`, `compile.metal.ci.hxml`.
- deterministic scripted mode used by CI (`expected/<profile>.stdout`).
- interactive mode for local demo.
- generated Go trees committed under `generated/<profile>/`.
- profile run outputs committed under `out_<profile>/` (from harness runs).
- app README with preset/variant behavior matrix and benchmark commands.

This keeps both examples visible in `python3 test/run-examples.py` and prevents
preset or source-boundary drift. The `<profile>` path token is retained by the
existing harness as a compatibility name for the selector dimension.

## Variant Strategy (per app)

Use one shared codebase per app, with explicit runtime variants instead of separate forks:

- `core` variant:
  - must run under both presets (`portable`, `metal`)
  - preset differences may affect diagnostics, specialization attempts, and
    code shape, not application features or portable semantics
- `go_native` variant:
  - still one codebase, but enables Go-first lanes through typed adapters
  - declares its Go-specific contract through typed `go.*` APIs and
    `@:goNative`, under either preset

Implementation mechanism:

- Keep shared domain logic in `app/core/*`.
- Keep variant adapters in `app/runtime/*` and mark Go-owning modules
  `@:goNative`.
- Select variant through explicit compile define (for example `-D app_variant=core|go_native`).

This gives one app across two compatibility presets and two explicit source
variants without fragmenting into unrelated implementations.

## How Presets Are Showcased

Use one shared app core and explicit runtime variants:

- Keep business logic on portable source surfaces.
- Isolate Go-native behavior behind a small typed runtime adapter layer.
- Document behavior contracts explicitly:
  - `portable`: guarded/proven/allow defaults;
  - `metal`: explicit/eager/error/strict compatibility defaults;
  - `go_native`: the source/API boundary that actually opts into Go-specific
    behavior.

For each app README, include:

- one table showing policy behavior per preset.
- one table showing variant behavior (`core` vs `go_native`) per preset.
- one section showing generated Go differences (code-shape highlights).
- one section showing binary/runtime benchmark deltas by preset and variant.

## App Scope

### PulseForge (flagship)

Real-time observability stream processor with:

- ingest endpoints (HTTP; optional TCP/UDP follow-up).
- pipeline stages: `parse -> enrich -> aggregate -> alert`.
- goroutine/channel/select-based worker graph.
- bounded queues and backpressure policy.
- live status output (TUI or rich CLI dashboard mode).
- WAL + replay mode for deterministic restart demonstration.

Primary compiler surface exercised:

- concurrency (`go`/`chan`/`select`)
- interfaces + dispatch
- generic containers (`go.Slice<T>`, `go.Map<K,V>`, `go.Result<T>`)
- `context` cancellation/timeouts
- typed interop externs
- error flow and recovery boundaries

### FluxProxy (secondary)

High-throughput reverse proxy / gateway with:

- upstream pool and request routing.
- concurrency limits + per-route rate limits.
- timeout/retry/circuit-breaker style policies.
- live stats endpoint and structured logs.

Primary compiler surface exercised:

- `net/http` interop-heavy path
- context propagation
- sync/atomic counters
- error wrapping + classification

## Benchmark Protocol (Haxe->Go vs Pure Go)

Benchmark both apps with the same workload generator and data:

- variants:
  - pure Go baseline
  - Haxe->Go `portable`
  - Haxe->Go `metal`
- metrics:
  - throughput (events or requests/sec)
  - p95/p99 latency
  - allocations/op and bytes/op
  - CPU and RSS
  - startup time
  - binary size (raw + stripped)
- constraints:
  - same hardware class and Go toolchain
  - same workload profile and duration
  - warm-up + steady-state windows
  - deterministic scripted mode for CI repeatability

## Implementation Order

1. PulseForge MVP (usable across profiles + scripted mode + examples harness integration).
2. PulseForge pure-Go parity baseline + benchmark harness entry.
3. FluxProxy MVP (same usability contract + profile matrix).
4. FluxProxy pure-Go parity baseline + benchmark harness entry.
5. CI benchmark job/docs updates after both are stable.

## Repo Integration Targets

- `examples/pulseforge`
- `examples/fluxproxy`
- pure-Go parity implementations:
  - `benchmarks/pure_go/pulseforge`
  - `benchmarks/pure_go/fluxproxy`
- benchmark orchestration:
  - `scripts/ci/perf-apps.sh` (`npm run test:perf:apps`)
  - baseline ratios: `scripts/ci/perf/app-profile-baseline.json` (`npm run test:perf:apps:update-baseline`)
  - artifacts under `.cache/perf-apps/results/` (`current.json`, `comparison.json`, `summary.md`, `raw_metrics.tsv`, `warnings.txt`, `hard_failures.txt`)

## Acceptance Gates

Before calling either app complete:

- examples matrix passes for all declared presets.
- generated-tree drift checks pass.
- documented preset/variant behavior table exists.
- benchmark run produces machine-readable results.
- README links to benchmark summary and known tradeoffs.

Benchmark methodology and fairness rules are defined in `docs/benchmark-methodology-apps.md`.
