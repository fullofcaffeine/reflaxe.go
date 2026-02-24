# Flagship Apps Plan: PulseForge + FluxProxy

## Decision

Build both apps, with strict sequencing:

1. `PulseForge` first (flagship "wow" demo + primary benchmark).
2. `FluxProxy` second (network-heavy focused benchmark).

Both must be usable and runnable, not design-only artifacts.

## Usability Contract (non-negotiable)

Each app must ship with:

- `app/` Haxe source that compiles in `portable`, `gopher`, and `metal`.
- `compile.portable.hxml`, `compile.gopher.hxml`, `compile.metal.hxml`.
- `compile.portable.ci.hxml`, `compile.gopher.ci.hxml`, `compile.metal.ci.hxml`.
- deterministic scripted mode used by CI (`expected/<profile>.stdout`).
- interactive mode for local demo.
- generated Go trees committed under `generated/<profile>/`.
- profile run outputs committed under `out_<profile>/` (from harness runs).
- app README with profile behavior matrix and benchmark commands.

This keeps both examples visible in `python3 test/run-examples.py` and prevents profile drift.

## Variant Strategy (per app)

Use one shared codebase per app, with explicit runtime variants instead of separate forks:

- `core` variant:
  - must run on all profiles (`portable`, `gopher`, `metal`)
  - profile differences should be mostly code-shape/perf, not feature removals
- `go_native` variant:
  - still one codebase, but enables Go-first lanes through typed adapters
  - can expose additional capability in `gopher`/`metal` (for example richer concurrency or interop paths)

Implementation mechanism:

- Keep shared domain logic in `app/core/*`.
- Keep profile/variant adapters in `app/runtime/*`.
- Select variant through explicit compile define (for example `-D app_variant=core|go_native`).

This gives "portable vs metal versions" in practice without fragmenting into multiple unrelated apps.

## How Profiles Are Showcased

Use one shared app core and profile runtime adapters:

- Keep business logic profile-agnostic where possible.
- Isolate profile-sensitive behavior behind a small runtime adapter layer.
- Document behavior contracts explicitly:
  - `portable`: semantics-first fallback paths.
  - `gopher`: same contract, safer Go-first optimizations.
  - `metal`: typed low-level interop/perf lane.

For each app README, include:

- one table showing feature behavior per profile.
- one table showing variant behavior (`core` vs `go_native`) per profile.
- one section showing generated Go differences (code-shape highlights).
- one section showing binary/runtime benchmark deltas by profile.

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
  - Haxe->Go `gopher`
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
  - extend `scripts/ci/perf-go-profiles.sh` or add `scripts/ci/perf-apps.sh`

## Acceptance Gates

Before calling either app complete:

- examples matrix passes for all profiles.
- generated-tree drift checks pass.
- documented profile behavior table exists.
- benchmark run produces machine-readable results.
- README links to benchmark summary and known tradeoffs.
