# Semantic Diff Guide

This guide explains what the semantic-diff harness is, why it exists, how one
case runs, when Haxe's interpreter is a valid reference, and when to choose a
different test surface.

## Terms

- `portable contract`: the baseline behavior contract for cross-target-compatible code.
  - reference: `docs/portable-canonical-contract.md`
- `reference execution`: the stdout and exit result produced by running the same
  source program with Haxe `--interp`.
- `semantic diff case`: a fixture where we compare runtime stdout between Haxe `--interp` and generated Go.
- `lane suite`: semantic-diff fixtures for native-boundary guarantees (canonical
  `@:goNative`, with existing `@:goMetal` compatibility fixtures) under
  `reflaxe_go_auto=auto_strict`. The source behavior must still be portable; the
  annotation alone does not make the interpreter a valid oracle for target-only
  operations.
- `go_defines.txt`: optional case-local define list used to toggle compiler behavior for a semantic-diff fixture.

## What Semantic Diff Is

Semantic diff compares **program behavior**, not just generated code shape.

For each case, the harness runs the same source program twice:

1. Haxe reference execution (`--interp`)
2. `reflaxe.go` compilation (`portable` contract)
3. generated Go runtime (`go run .`)
4. exact stdout comparison (`reference` vs `generated-go`)

There is no hand-written expected stdout file. The interpreter result is captured
at test time and becomes the reference. If either execution fails or their stdout
differs exactly, the case fails.

## Why It Exists

Snapshot tests answer: "did generated Go text change?"

Semantic diff answers: "did runtime meaning change?"

You need both:

- Snapshot catches codegen shape drift.
- Semantic diff catches behavior drift even when code still compiles and looks reasonable.

Typical regressions semantic diff catches:

- null/stringification differences (`null` vs `<nil>`)
- exception/catch behavior drift
- equality/boxing edge cases
- serializer/wire-format runtime differences
- stdlib behavior mismatches across contracts

## When the Haxe interpreter is a valid reference

The interpreter is a valid oracle only when the **same source program** has a
defined portable Haxe meaning. A good case can be copied from `test/semantic_diff`
and run with `--interp` without replacing its types, APIs, or algorithm.

Target-native behavior does not meet that rule. Code exercising `go.*`,
target-only operations inside `@:goNative` modules, typed extern calls, raw Go
injection, or Go runtime details must not use semantic-diff as its primary
behavior test. Haxe's interpreter cannot define the intended semantics of those
APIs.

Do not solve an invalid reference by adding an interpreter fallback to production
code. A portable substitute created only for the reference run would test two
different programs and could hide a Go-target regression. Move the behavior to a
target-only contract instead.

Metadata-only lane cases may remain here when their executable source is otherwise
portable. In that situation semantic-diff proves that boundary analysis or
specialization did not change the portable behavior; it does not prove a
target-native API.

## How one case runs

Entrypoint: `test/run-semantic-diff.py`

Per case pipeline:

1. Clean case output dir (`<case>/out`).
2. Run the reference execution:
   - `haxe -cp <case> -cp src -main Main --interp`
3. Compile Go output:
   - Haxe macros bootstrapping `reflaxe.go`
   - `-D reflaxe_go_profile=portable`
   - optional per-case defines from `go_defines.txt`
4. Run generated Go checks:
   - `go test ./...`
   - `go run .`
5. Compare generated-Go stdout with the captured reference stdout exactly.
6. Record the case result and any failure stage under `test/.test-cache/`.

The harness reports stage-level failures (`reference`, `compile`, `go test`, `runtime`, `diff`).

The core and lane suites use separate locks because every run recreates each
case's `out/` directory. This prevents concurrent processes from deleting or
rewriting one another's generated Go.

## Case Layout

Core suite root: `test/semantic_diff/`

Lane suite root: `test/semantic_diff_lanes/`

Typical case:

```text
test/semantic_diff/<case_id>/
  Main.hx
  go_defines.txt   # optional; one define per line
```

`go_defines.txt` lets you pin optimizer/profile capability toggles per case without changing harness code.

## Choosing the right harness

Choose evidence based on the contract being proved:

| Contract | Test surface | Why |
| --- | --- | --- |
| Portable Haxe runtime meaning | Semantic diff | Haxe `--interp` is a meaningful oracle for the same source. |
| Go-native API behavior (`go.*`, target-only `@:goNative`, typed extern calls) | Go-only runtime snapshot with `expected.stdout`, or a runnable example | The intended behavior exists only on the Go target. |
| Generated Go representation or optimization shape | Snapshot `intended/*.go` | Runtime output cannot prove which lowering was selected. |
| Planner eligibility, fallback reason, or specialization count | Optimizer/report contract | Reports prove the compiler decision independently of behavior. |
| Isolated `hxrt` behavior | Go unit test | The contract is a Go runtime helper, not a Haxe cross-target program. |
| Performance or footprint | Performance harness | Semantic equivalence does not establish a budget. |

For a feature with both portable and native paths, use more than one surface:
semantic-diff for the portable source contract, a Go-only test for native behavior,
and an optimizer/report or generated-shape snapshot for path selection. Do not
make the native API interpreter-compatible merely to collapse those proofs into
one fixture.

Examples of the target-only path live under `test/snapshot/go_native/`:
`channel_try_recv`, `select_helpers`, `extern_value_error_result`, and
`native_boundary_collections_strict`. The snapshot runner compiles their generated
Go, runs `go test ./...`, and, with `--runtime`, compares program output with the
committed `expected.stdout` file.

## What semantic diff does not prove

Semantic diff does not prove:

- that generated Go is idiomatic or has a particular shape;
- that `go.*`, target-only `@:goNative`, or typed extern APIs work correctly;
- that a particular optimizer path was selected;
- performance, allocation count, binary footprint, race freedom, or scheduling;
- behavior not made observable through successful execution and deterministic stdout.

Use the dedicated harness from the table above for each of those claims.

## Commands You Will Use

Run core suite:

```bash
python3 test/run-semantic-diff.py
```

Run lane suite:

```bash
python3 test/run-semantic-diff.py --suite lanes
```

List cases:

```bash
python3 test/run-semantic-diff.py --list
```

Run changed cases only:

```bash
python3 test/run-semantic-diff.py --changed
```

Run optimizer semantic matrix:

```bash
npm run test:semantic-diff:optimizer-matrix
```

Fail fast if another semantic-diff run is active:

```bash
python3 test/run-semantic-diff.py --lock-timeout 0
```

## Locking And Stability

The harness uses suite-scoped lock files under `test/.test-cache/`:

- `run-semantic-diff-core.lock`
- `run-semantic-diff-lanes.lock`

This prevents overlapping runs from stepping on the same case `out/` directories.

Use `--lock-timeout` to control wait behavior.

## Reading Failures Quickly

- `reference`: Haxe `--interp` failed. The case itself is invalid.
- `compile`: Go generation failed. Codegen/typing/defines issue.
- `go test`: generated module failed to compile/test.
- `runtime`: generated binary crashed/errored.
- `diff`: both ran, but outputs differ (semantic regression).

Start debugging from the first failing stage.

## Authoring Checklist For New Cases

1. Confirm the same source program has meaningful portable behavior under
   `--interp`; reject the case if it relies on a target-specific API.
2. Keep the case deterministic (no clocks, network, or randomness unless controlled).
3. Print only contract-relevant outputs.
4. Keep the case minimal; one behavior family per fixture.
5. If compiler toggles matter, encode them in `go_defines.txt`.
6. Add comments in `Main.hx` only when behavior is non-obvious.
7. Run the case directly before broad suite runs.

## Scope Rule

Semantic diff is for **portable contract behavior** and lane contracts whose
same executable source remains portable.
It is not a benchmark harness and not a generated-code-style harness.

Use:

- semantic diff for behavior
- Go-only runtime snapshots or examples for target-native behavior
- snapshots for generated shape
- perf harness for performance budgets

## Related docs

- `docs/start-here.md` - first-run setup and default test flow.
- `docs/native-policy-presets.md` - source boundary, compatibility presets, and policy axes.
- `docs/profile-semantics-guide.md` - behavior guarantees and migration guidance.
- `docs/snapshot-policy.md` - generated Go shape contracts.
- `docs/glossary.md` - shared vocabulary used across compiler docs.
