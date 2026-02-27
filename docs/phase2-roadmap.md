# Phase 2 Roadmap (Canonical)

## Status

- This file is the active Phase-2 roadmap.
- `prd.md` is retained as Phase-1 historical context.

## North Star

1. Keep portable Haxe correctness and harness discipline as the baseline.
2. Make Go-target `go.*` APIs genuinely Go-native.
3. Provide typed interop so Go ecosystem usage does not depend on raw `__go__`.
4. Make `metal` a real typed-performance lane with explicit budgets.
5. Add selective `hxrt` runtime slicing so metal can minimize runtime overhead without dropping portability contracts.
6. Reach full portable-eligible Haxe stdlib parity in `portable`, with explicit portable-vs-native facade boundaries and deterministic parity reporting.

## Profile Contract

- `portable`: semantics-first output, lowest migration risk.
- `metal`: `portable` + typed interop/performance lane (strict defaults).

Selective runtime slicing is tracked separately and does not replace profile contracts.
See `docs/hxrt-selective-runtime.md`.

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

- Implement unified `go.Go` / `go.Chan` authority in `src/go/*` with target-conditional lowering hooks (Go-native on `go` builds, deterministic simulation for non-Go harness paths).
- Allow `std/go/*` override classes through compiler project-class filtering.
- Add deterministic go-native concurrency contracts (no sleep-based race tests).
- Latest follow-up: typed `go.Chan<T>` `recv`/`recvOr` call results now route through generic assertion bridging in `portable`, so typed channel reads compile without forcing `Dynamic` callsites.
- Latest follow-up: `go.Chan<T>.tryRecv():go.Result<T>` now lowers through select-backed concurrency shims across `portable`/`metal`, with dedicated go_native snapshot coverage.
- Latest follow-up: typed `go.Select` helpers (`recv`, `recv2`, `send`, `send2`) now provide a deterministic Haxe-level select API with explicit branch-priority semantics.

### M3 - Typed Interop Foundation

Depends on: M2

- Implement typed interop metadata and import resolution (`@:go.import`, symbol mapping).
- Add smoke examples for `fmt`, `time`, and `context`.
- Latest follow-up: extern calls typed as Haxe `String` now normalize return values through `hxrt.StdString` (covers static/instance/receiver interop forms without `Dynamic` callsite workarounds).

### M3.5 - Interop Autopilot

Depends on: M3

- Add deterministic `tools/goextern` generation flow + fixture checks in CI.

### M4 - Metal Typed-Performance Lane

Depends on: M3, M3.5

- Start with monomorphization-first typed lane for hot generic surfaces.
  - Prototype landed: `go.Chan<T>` metal call-site specialization with typed channel shims per concrete element type.
  - Extended prototype landed: `go.Slice<T>` + `go.Map<K,V>` metal call-site specialization with typed collection shims per concrete type set.
- `go.Result<T>` metal lowering now includes typed call-site shims with internal `(T, error)` helpers; continue iterating toward broader direct idiom emission.
- Perf harness now covers channel/map/generic microbench lanes with metal budget enforcement.

### M4.5 - Selective `hxrt` Runtime Slicing

Depends on: M4

- Split runtime into feature files (`core` + optional slices).
- Infer required runtime features from compilation inputs and shim usage.
- Support selective runtime copy with define-based overrides.
- Keep full runtime copy fallback for compatibility/debugging.

### M5 - Output Ergonomics

Depends on: M3

- Add file-per-module output within a single Go package.
- Add optional `//line` directives and harness support for multi-file snapshots.

### M6 - Stdlib Completion + Onboarding

Depends on: M2 (with task-level dependencies on M3+ lanes where needed)

- Promote high-value compile-only surfaces to semantic-diff contracts.
  - Latest promotion: `haxe.Exception` (`caught`/`thrown`/`message`) via `test/semantic_diff/exception_api_contract`.
- Continue unsupported-expression inventory reduction.
  - Latest reduction: explicit `TIdent` lowering for untyped identifiers with snapshot coverage in `test/snapshot/core/untyped_ident_nil`.
- Publish concrete onboarding docs and showcase examples.
  - Latest docs/examples pass: added `examples/worker_pool_select`, expanded `examples/interop_smoke` with `net/http`, and published `docs/go-concurrency-interop-guide.md` + `docs/known-gaps.md`.

### M7 - Portable Stdlib 100% Parity Program

Depends on: M6 (and cross-cuts M2-M4.5)

- Adopt explicit stdlib layering contract: portable canonical surface + `go.*` native facade (non-portable).
- Close module-by-module parity gaps until all portable-eligible upstream Haxe stdlib modules are supported in portable mode.
- Add deterministic inventory/provenance artifacts and CI boundary gates for stdlib synchronization.
- Add portable contract native-import policy with migration warnings and CI/release hard errors.
- Consolidate duplicate native facade authorities (`src/go/*` vs `std/go/*`) under one canonical path.
- Track family-shared portable stdlib extraction path once Go parity program stabilizes.

Execution tracker: `haxe.go-cgk` (`docs/portable-stdlib-parity-program.md`).

Approach-C closure follow-up tracker: `haxe.go-qhv` (lane enforcement depth + lane semantic gates).

## Execution Rules

- Harness is the source of truth: snapshots, stdlib sweep, semantic diff, examples, perf.
- Every milestone must land with deterministic tests and docs updates.
- Profile semantics must be explicit; no implicit drift between profiles.
