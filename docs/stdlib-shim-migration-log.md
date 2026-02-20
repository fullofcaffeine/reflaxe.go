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

### 2026-02-19: Sys/File/Process extraction from compiler core (`haxe.go-7zy.11`)

Implementation:

- Removed behavior-heavy `sys` imports (`bufio`, `os`, `os/exec`) from compiler shim import wiring in `src/reflaxe/go/GoCompiler.hx`.
- Reworked `lowerSysStdlibShimDecls` to forwarding wrappers only:
  - `Sys_getCwd` -> `hxrt.SysGetCwd`
  - `Sys_args` -> `hxrt.SysArgs`
  - `sys__io__File_saveContent` -> `hxrt.FileSaveContent`
  - `sys__io__File_getContent` -> `hxrt.FileGetContent`
  - `New_sys__io__Process` -> `hxrt.NewProcess`
  - `sys__io__ProcessOutput.readLine` -> `hxrt.ProcessOutput.ReadLine`
  - `sys__io__Process.close` -> `hxrt.Process.Close`
- Added runtime-owned behavior to `runtime/hxrt/hxrt.go`:
  - `SysGetCwd`, `SysArgs`, `FileSaveContent`, `FileGetContent`
  - `ProcessOutput` and `Process` runtime types
  - `NewProcess`, `Stdout`, `ReadLine`, `Close`
- Preserved generated Haxe type-shape parity by keeping compiler-side wrapper structs that now delegate all behavior.

Validation evidence:

- Focused parity:
  - `python3 test/run-snapshots.py --case sys/file_read_write_smoke --case sys/process_echo_smoke`
  - `python3 test/run-upstream-stdlib-sweep.py --module Sys --module sys.io.File --module sys.io.Process --strict --go-test`
- Full regression:
  - `npm run test:ci`

Observed result:

- `lowerSysStdlibShimDecls` no longer carries behavior-heavy file/process logic.
- Local CI remains fully green after migration:
  - snapshots: `94/94`
  - strict stdlib sweep: `55/55`
  - semantic diff: `27/27`
  - examples: `6/6`

### 2026-02-19: `stdlib_symbols` bytes-conversion optimization (`haxe.go-7zy.12`)

Implementation:

- Added an internal raw-byte cache to generated `haxe__io__Bytes` in `src/reflaxe/go/GoCompiler.hx`:
  - `__hx_raw []byte`
  - `__hx_rawValid bool`
- Updated `haxe__io__Bytes_ofString` to initialize both int-backed (`b`) and raw-byte representations.
- Invalidated cache in mutating `haxe__io__Bytes.set`.
- Updated conversion helpers:
  - `hxrt_haxeBytesToRaw` now reuses cached raw bytes when valid.
  - `hxrt_rawToHaxeBytes` now seeds cache on construction.

Validation evidence:

- Perf harness:
  - `npm run test:perf:stdlib-shims`
  - comparative 3-run sample using prior commit (`8b18b3f`) vs optimized commit:
    - baseline shim `ns/op`: `135.0`, `179.7`, `79.58` (median `135.0`)
    - optimized shim `ns/op`: `107.8`, `70.25`, `74.75` (median `74.75`)
    - median delta: `-44.63%` shim `ns/op`
- Semantic parity:
  - `python3 test/run-semantic-diff.py --case crypto_xml_zip`
- Full regression harness:
  - `npm run test:ci`

Observed result:

- `stdlib_symbols` bytes conversion path keeps parity while improving measured shim-path performance versus baseline sample on the same machine.
- Snapshot/example goldens were refreshed for impacted stdlib/sys surfaces.

### 2026-02-20: IO helper surface gating + edge coverage (`haxe.go-czm`)

Implementation:

- Kept `haxe.io` ownership in compiler core but split emission policy:
  - core bytes stream declarations are always emitted when `io` shims are required.
  - inherited `haxe.io.Input`/`haxe.io.Output` helper surface declarations are emitted only when helper methods are actually referenced by lowered code.
- Added usage tracking in lowering:
  - `noteIoHelperFieldUsage` now marks helper-surface requirement when `haxe.io.Input`/`BytesInput` helper reads or `haxe.io.Output`/`BytesOutput` helper writes are accessed.
- Added selective trimming path in `lowerIoStdlibShimDecls`:
  - `trimIoShimToCoreSurface` removes helper-only interfaces/functions/method wrappers when not needed.
  - helper mode still preserves full parity subset introduced by `haxe.go-vxe`.
- Added semantic edge fixture:
  - `test/semantic_diff/io_input_output_edge_contract` for `readLine` EOF/tail/CRLF behavior.

Validation evidence:

- `python3 test/run-snapshots.py --update --timeout 180`
- `python3 test/run-semantic-diff.py --timeout 180` (`51/51`, includes new edge fixture)

Observed result:

- Generated footprint for non-IO-heavy fixtures dropped without behavior regressions:
  - `stdlib/math_basic`: `1488` -> `971` lines (`-517`, `-34.7%`)
  - `stdlib/stringtools_basic`: `1484` -> `967` lines (`-517`, `-34.8%`)
- IO-smoke fixture also shrank after eliminating always-on helper declarations:
  - `stdlib/io_type_smoke`: `901` -> `381` lines (`-520`, `-57.7%`)
- Snapshot refresh delta for this optimization pass: `6218` deleted lines, `95` inserted lines.

## Open migration track

- No open shim migration beads remain in the `haxe.go-7zy.10`/`.11`/`.12` sequence.
