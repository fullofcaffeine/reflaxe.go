# Multi-Package Output Re-evaluation

## Date
2026-02-18

## Scope
Issue `haxe.go-xd9.1.5` asks whether output should move from the current single `main` package model to multi-package emission.

## Current State
- The backend currently emits a single Go package with deterministic global mangling.
- Milestone 9 stdlib shims are now implemented and covered by snapshots/sweep checks.
- Runtime and stdlib shims are intentionally centralized to keep behavior deterministic.

## Findings
- A naive Haxe package -> Go package split will introduce import cycles quickly (class inheritance, enum usage, and stdlib helper dependencies).
- Existing naming/mangling already solves collisions in the single-package model with low complexity.
- The current pipeline does not yet have an import-graph planner, cycle breaker, or package-level topological staging.
- Multi-package output would require a dedicated planning phase and additional regression coverage before it is safe.

## Decision
Keep single-package output through Milestone 10.

## Entry Criteria For Re-opening Multi-Package Work
1. Add an explicit package graph planner pass in the compiler pipeline.
2. Define cycle-breaking rules (interface extraction, adapter package, or dependency inversion strategy).
3. Add snapshot coverage that asserts import graph determinism and no-cycle guarantees.
4. Keep a compatibility mode that preserves current single-package output.

## Follow-up
Track implementation as a future phase item after Milestone 10 stabilization.
