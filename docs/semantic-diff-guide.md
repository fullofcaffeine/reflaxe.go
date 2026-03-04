# Semantic Diff Guide

This guide explains what semantic diff is, why we use it, and how to add/diagnose cases.

## Key Terms

- `portable contract`: the baseline behavior contract for cross-target-compatible code.
  - reference: `docs/portable-canonical-contract.md`
- `semantic diff case`: a fixture where we compare runtime stdout between Haxe `--interp` and generated Go.
- `lane suite`: semantic-diff fixtures for lane-specific guarantees (for example `@:goMetal` lane behavior under `reflaxe_go_auto=auto_strict`).
- `go_defines.txt`: optional case-local define list used to toggle compiler behavior for a semantic-diff fixture.

## What Semantic Diff Is

Semantic diff compares **program behavior**, not just generated code shape.

For each case, the harness runs:

1. Haxe reference execution (`--interp`)
2. `reflaxe.go` compilation (`portable` contract)
3. generated Go runtime (`go run .`)
4. exact stdout comparison (`reference` vs `generated-go`)

If stdout differs, the case fails.

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

## How It Works Internally

Entrypoint: `test/run-semantic-diff.py`

Per case pipeline:

1. Clean case output dir (`<case>/out`).
2. Run Haxe reference:
   - `haxe -cp <case> -cp src -main Main --interp`
3. Compile Go output:
   - Haxe macros bootstrapping `reflaxe.go`
   - `-D reflaxe_go_profile=portable`
   - optional per-case defines from `go_defines.txt`
4. Run generated Go checks:
   - `go test ./...`
   - `go run .`
5. Compare stdout exactly.

The harness reports stage-level failures (`reference`, `compile`, `go test`, `runtime`, `diff`).

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

## Portable+Auto Contract Pattern

When adding planner-driven lowering behavior, keep two paired semantic-diff cases:

1. typed-success case (eligible types; specialization should apply)
2. fallback case (ineligible types; fallback path should preserve semantics)

Current reference pair:

- `test/semantic_diff/go_auto_collections_result_typed_contract`
- `test/semantic_diff/go_auto_collections_result_fallback_contract`

Both cases pin:

```text
reflaxe_go_auto=auto
```

via `go_defines.txt`.

Why this pairing exists:

- semantic diff proves runtime parity vs `--interp` for both execution paths.
- optimizer report contracts prove planner-path intent:
  - typed case records typed lowerings with zero typed fallbacks.
  - fallback case records typed fallbacks with reason counts.

The optimizer side is validated by:

```bash
npm run test:auto-planner:schema
```

Use this same pattern whenever a new auto-lowering capability is introduced.

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

1. Keep case deterministic (no clocks/network/randomness unless controlled).
2. Print only contract-relevant outputs.
3. Keep case minimal; one behavior family per fixture.
4. If toggles matter, encode them in `go_defines.txt`.
5. Add comments in `Main.hx` only when behavior is non-obvious.
6. Run the case directly before broad suite runs.

## Scope Rule

Semantic diff is for **portable contract behavior** and explicit lane contracts.
It is not a benchmark harness and not a generated-code-style harness.

Use:

- semantic diff for behavior
- snapshots for generated shape
- perf harness for performance budgets
