# Multi-Package Output Re-evaluation

## Date
2026-03-04

## Scope
Issue `haxe.go-l3qt.4` evaluates whether output should move from the current single-package Go model to multi-package emission for production GA.

## Current State
- The backend emits one Go package (`main`) split across multiple files with deterministic mangling.
- This model is already covered by snapshots and semantic/runtime harness gates.
- Runtime + shim ownership stays deterministic because package-level import graph planning is not required yet.

## Findings
- A direct Haxe-package-to-Go-package split introduces import-cycle risk across inheritance, enum wiring, and helper/shim dependencies.
- Existing single-package output already satisfies current production-readiness goals (portable contract, metal contract, runtime/report artifacts, and CI reproducibility).
- There is no proven production blocker in current workloads that requires multi-package emission today.
- Safe multi-package support needs new compiler stages (package graph planner + cycle strategy) and dedicated graph-focused tests.

## GA Decision
Decision: defer multi-package output for production GA.

This is a non-blocking defer, not a rejection of multi-package output.

## Boundary conditions that re-open this work
1. A real deployment requires package-scoped ownership that cannot be represented in single-package output.
2. We introduce an import-graph planner with deterministic package assignment and cycle diagnostics.
3. We define a cycle-breaking strategy (interface extraction, adapter package, or dependency inversion) and validate it with regression tests.
4. We add importer/graph tests that assert deterministic package graphs and no-cycle guarantees.
5. We keep a single-package compatibility mode so existing users are not forced into package-graph migration.

## If/when implementation starts
Use staged delivery:
1. Planner-only phase: build package graph + diagnostics, no output behavior change.
2. Dual-output phase: optional multi-package emission behind a define with importer/graph snapshot tests.
3. Graduation phase: production criteria based on deterministic graph/test evidence.
