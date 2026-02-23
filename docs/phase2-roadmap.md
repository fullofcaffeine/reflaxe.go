# Phase 2 Roadmap (Canonical)

## Status

- This file is the active Phase-2 roadmap.
- `prd.md` is retained as Phase-1 historical context.

## North Star

1. Keep portable Haxe correctness and harness discipline as the baseline.
2. Make Go-target `go.*` APIs genuinely Go-native.
3. Provide typed interop so Go ecosystem usage does not depend on raw `__go__`.
4. Make `metal` a real typed-performance lane with explicit budgets.

## Profile Contract

- `portable`: semantics-first output, lowest migration risk.
- `gopher`: portable behavior plus safe Go-first optimizations.
- `metal`: `gopher` + typed interop/performance lane (strict defaults).

`idiomatic` remains removed and must not be reintroduced.

## Milestones And Dependencies

### M0 - Baseline Alignment

- Fix open correctness paper-cuts before feature expansion.
- Publish canonical Phase-2 roadmap and deprecate stale roadmap language.

### M1 - Go AST Native Constructs

Depends on: M0

- Add first-class AST/printer/transformer support for:
  - `go` statements
  - `defer` statements
  - channel send statements (`ch <- v`)
  - receive expressions (`<-ch`)
  - `select` statements
  - `range` loops

### M1.5 - Compiler Structure Hardening

Depends on: M1

- Refactor `src/reflaxe/go/GoCompiler.hx` into subsystem modules before major feature growth.

### M2 - Real Go Concurrency

Depends on: M1, M1.5

- Implement Go-target overrides for `std/go/Go.hx` and `std/go/Chan.hx` with real goroutines/channels.
- Allow `std/go/*` override classes through compiler project-class filtering.
- Add deterministic go-native concurrency contracts (no sleep-based race tests).

### M3 - Typed Interop Foundation

Depends on: M2

- Implement typed interop metadata and import resolution (`@:go.import`, symbol mapping).
- Add smoke examples for `fmt`, `time`, and `context`.

### M3.5 - Interop Autopilot

Depends on: M3

- Add deterministic `tools/goextern` generation flow + fixture checks in CI.

### M4 - Metal Typed-Performance Lane

Depends on: M3, M3.5

- Start with monomorphization-first typed lane for hot generic surfaces.
- Add `go.Result<T>` lowering strategy for metal.
- Expand perf harness and enforce metal budgets.

### M5 - Output Ergonomics

Depends on: M3

- Add file-per-module output within a single Go package.
- Add optional `//line` directives and harness support for multi-file snapshots.

### M6 - Stdlib Completion + Onboarding

Depends on: M2 (with task-level dependencies on M3+ lanes where needed)

- Promote high-value compile-only surfaces to semantic-diff contracts.
- Continue unsupported-expression inventory reduction.
- Publish concrete onboarding docs and showcase examples.

## Execution Rules

- Harness is the source of truth: snapshots, stdlib sweep, semantic diff, examples, perf.
- Every milestone must land with deterministic tests and docs updates.
- Profile semantics must be explicit; no implicit drift between profiles.
