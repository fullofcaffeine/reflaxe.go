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
- Existing single-package output already satisfies current production-readiness
  goals (portable semantics, explicit native boundaries, policy presets,
  runtime/report artifacts, and CI reproducibility).
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

## Measurable production reopen triggers

Single-package output stays the default until a real project produces repeatable
evidence that the single-package model is the bottleneck. A **trigger** is a
measured signal that should create or unblock a Bead for multi-package output.
It is not a hunch that multi-package output would look nicer.

| Trigger | What to measure | Reopen threshold | Trigger evidence |
| --- | --- | --- | --- |
| Generated file size | Size of the largest generated `.go` file and total generated Go size for a production workload. | A generated file becomes too large for normal review/tooling, or repeated builds produce a single file that is clearly the source of formatter, editor, or compiler pain. | Attach generated-size numbers, the workload name, and the command that produced them. |
| Go compile time | Wall-clock time for `go test` / `go build` on generated output, measured on the same machine or CI runner. | Single-package output becomes a repeatable compile-time bottleneck after unrelated runtime/test costs are excluded. | Attach at least three runs, median time, command, Go version, and whether the same workload improves with a prototype package split. |
| Package-private boundary needs | Need for Go package boundaries to use package-private APIs, generated helper isolation, or separate review/ownership areas. | A production integration cannot be represented honestly through public generated symbols or typed extern facades. | Attach the required boundary, why public wrappers are not enough, and the affected generated modules. |
| Go tooling limit | A concrete Go tool, editor, linter, formatter, or build system limit caused by single-package output. | The tool fails, times out, or becomes unusably slow only because output is one package. | Attach the tool name/version, failing command, error/log excerpt, and a minimized reproduction if possible. |

CI records the first two trigger categories automatically for the examples and
flagship apps that run through `test/run-examples.py`:

- `.cache/generated-output-telemetry/examples.json`
- `.cache/generated-output-telemetry/examples.md`

The artifact records largest generated Go file size, total generated Go size,
and `go test ./...` elapsed time for each generated example output tree. Use
that artifact as the starting evidence before reopening multi-package output.
If it points to a real bottleneck, confirm with repeated runs on the same runner
class before changing output architecture.

If a trigger fires, implementation still starts with the planner-only phase
below. Do not skip directly to changing output shape; package graphs need their
own deterministic tests first.

## If/when implementation starts
Use staged delivery:
1. Planner-only phase: build package graph + diagnostics, no output behavior change.
2. Dual-output phase: optional multi-package emission behind a define with importer/graph snapshot tests.
3. Graduation phase: production criteria based on deterministic graph/test evidence.
