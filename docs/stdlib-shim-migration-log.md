# Stdlib Shim Migration Log

This log tracks the end-to-end shim migration process so decisions, rollout order, and validation evidence stay auditable.

## Process Template

For each shim surface:

1. Scope and classify the surface in `docs/stdlib-shim-rationale.md`.
2. Pick target ownership (`compiler core`, `runtime-lowered`, or `std/_std`).
3. Implement the migration in a minimal step.
4. Validate with focused harnesses first, then strict sweeps.
5. Record results and open/close follow-up beads.

## Timeline

### 2026-02-19: baseline architecture review (`haxe.go-7zy.7`)

- Published a shim decision matrix with keep/migrate choices.
- Added reproducible benchmark harness:
  - `scripts/ci/perf-stdlib-shim-review.sh`
  - `npm run test:perf:stdlib-shims`
- Opened migration/perf follow-up beads:
  - `haxe.go-7zy.10` (`haxe.Json` migration)
  - `haxe.go-7zy.11` (`Sys/File/Process` migration)
  - `haxe.go-7zy.12` (`stdlib_symbols` bytes conversion optimization)

### 2026-02-19: JSON shim extraction (`haxe.go-7zy.10`)

Implementation:

- Removed compiler-emitted JSON shim declarations from `src/reflaxe/go/GoCompiler.hx`.
- Removed JSON shim-group activation from stdlib shim routing (`requiredStdlibShimGroups` no longer tracks `json`).
- Kept JSON behavior via direct runtime lowering:
  - `haxe.Json.parse` -> `hxrt.JsonParse`
  - `haxe.Json.stringify` -> `hxrt.JsonStringify`
  - `haxe.format.JsonPrinter.print` -> `hxrt.JsonStringify`
  - `haxe.format.JsonParser.doParse` -> `hxrt.JsonParse`
- Lowered `new haxe.format.JsonParser(source)` to the source pointer representation (`*string`) to avoid synthetic parser struct emission.

Validation evidence:

- `python3 test/run-snapshots.py --case stdlib/json_parse_stringify`
- `python3 test/run-upstream-stdlib-sweep.py --module haxe.Json --strict --go-test`
- `python3 test/run-upstream-stdlib-sweep.py --strict`
- `python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_full.txt --strict`

Observed result:

- Snapshot no longer emits `haxe__Json`/`haxe__format__JsonParser` declarations.
- Strict stdlib sweeps remain green (`55/55` and `175/175`).

## Open migration track

- `haxe.go-7zy.11`: move `Sys`/`sys.io.File`/`sys.io.Process` shims out of compiler core.
- `haxe.go-7zy.12`: reduce `stdlib_symbols` conversion overhead while preserving parity.
