# Stdlib Shim Migration Log

This log tracks the end-to-end shim migration process so decisions, rollout order, and validation evidence stay auditable.

How to use this log:

1. Read `docs/stdlib-shim-rationale.md` first for current ownership decisions.
2. Use this page as the chronological execution record (what changed, when, and how it was validated).
3. Use linked test commands to reproduce evidence on current HEAD.

Terms:

- **shim**: compatibility glue between Haxe std APIs and Go output behavior.
- **semantic-diff**: runtime parity harness comparing generated Go output against Haxe `--interp`.
- **staged stdlib**: target-specific overrides under `std/_std`.

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

### 2026-02-25: staged JSON stdlib ownership (`haxe.go-cgk.5`)

Implementation:

- Added staged std overrides:
  - `std/_std/haxe/Json.cross.hx`
  - `std/_std/haxe/format/JsonParser.cross.hx`
  - `std/_std/haxe/format/JsonPrinter.cross.hx`
- Marked JSON classes as required staged stdlib classes in compiler ownership routing so override modules are compiled into output:
  - `haxe.Json`
  - `haxe.format.JsonParser`
  - `haxe.format.JsonPrinter`
- Removed compiler-call special cases that previously bypassed staged module ownership:
  - removed `haxe.Json.parse/stringify` direct lowerings
  - removed `haxe.format.JsonPrinter.print` direct lowering
  - removed `JsonParser` constructor/string-pointer special casing
- Kept runtime behavior centralized in `hxrt` through staged std wrappers.
- Updated runtime/shim ownership docs and provenance ledger for staged std files.

Validation evidence:

- `python3 test/run-snapshots.py --case stdlib/json_parse_stringify`
- `python3 test/run-semantic-diff.py --case json_parse_stringify_contract`
- `python3 test/run-upstream-stdlib-sweep.py --module haxe.Json --strict --go-test`
- `python3 test/run-ci.py --changed --skip-stdlib-sweep`

Observed result:

- JSON API ownership is now staged-stdlib-first while preserving runtime behavior and parity gates.
- Contract/runtime report and shim ownership remain deterministic after migration.

### 2026-03-06: `stdlib_symbols` anti-bloat audit against sibling targets (`haxe.go-14as.33`)

Implementation:

- Audited compiler-resident library-style helper surfaces in `src/reflaxe/go/GoCompiler.hx` instead of adding more behavior-heavy `GoStmt.GoRaw` blocks.
- Compared ownership patterns with sibling targets:
  - `haxe.rust`: `std/StringTools.cross.hx`, `std/hxrt/string/NativeString.hx`, `runtime/hxrt/src/string.rs`
  - `haxe.elixir`: `std/StringTools.cross.hx`, `std/DateTools.cross.hx`, target-gated `std/_std/**`
- Reclassified the `stdlib_symbols` strategy:
  - keep compiler ownership only for compile-context-sensitive or representation-bound surfaces (`Std`, `Reflect`, `Type`, `Xml`, bytes-sensitive crypto/zip paths)
  - migrate library-expressible helpers (`StringTools`, `DateTools` helpers, `haxe.io.Path`) to staged std and thin runtime helpers
- Opened focused migration beads:
  - `haxe.go-14as.34` (`StringTools`)
  - `haxe.go-14as.35` (`DateTools` helper formatting ownership`)
  - `haxe.go-14as.36` (`haxe.io.Path`)
- Updated `AGENTS.md` so future compiler work audits adjacent shim-group surfaces instead of expanding `GoCompiler` helper-by-helper.

Validation evidence:

- Local source audit against current HEAD:
  - `src/reflaxe/go/GoCompiler.hx` (`lowerStdlibSymbolShimDecls`)
  - sibling reference: `haxe.rust/std/StringTools.cross.hx`
  - sibling reference: `haxe.rust/std/hxrt/string/NativeString.hx`
  - sibling reference: `haxe.rust/runtime/hxrt/src/string.rs`
  - sibling reference: `haxe.elixir.codex/std/StringTools.cross.hx`
  - sibling reference: `haxe.elixir.codex/std/DateTools.cross.hx`
  - sibling reference: `haxe.elixir.codex/docs/04-api-reference/STANDARD_LIBRARY_HANDLING.md`
- Repo hygiene:
  - `git diff --check`

Observed result:

- The repo now states the stronger default rule explicitly: library-expressible stdlib does not belong in `GoCompiler` unless there is a concrete compiler-only reason.
- Future migration work is split into concrete beads instead of one-off local exceptions.

### 2026-03-06: direct `haxe.Constraints` + `haxe.Rest` abstraction closure (`haxe.go-14as.26`)

Implementation:

- Added `std/haxe/Constraints.cross.hx` so `haxe.Constraints` aliases and `IMap` bridge metadata live in staged std instead of implicit upstream assumptions.
- Updated staged std map externs:
  - `std/_std/haxe/ds/StringMap.cross.hx`
  - `std/_std/haxe/ds/IntMap.cross.hx`
  - `std/_std/haxe/ds/ObjectMap.cross.hx`
- Changed compiler DS shims in `src/reflaxe/go/GoCompiler.hx` to emit `copyIMap()` bridge methods with interface returns for concrete maps.
- Taught interface lowering to honor native field metadata for interface methods so `IMap.copy()` lowers to `copyIMap` without changing Haxe user code.
- Added minimal representation-aware lowering for direct `haxe.Rest` usage:
  - array `copy()` cloning for native-slice arrays
  - `haxe._Rest.Rest_Impl_.append`
  - `haxe._Rest.Rest_Impl_.prepend`
- Added focused parity coverage:
  - `test/semantic_diff/haxe_constraints_contract`
  - `test/semantic_diff/haxe_rest_contract`
  - `test/snapshot/stdlib/haxe_constraints_rest_direct`

Validation evidence:

- `python3 test/run-semantic-diff.py --case haxe_constraints_contract --case haxe_rest_contract`
- `python3 test/run-snapshots.py --case stdlib/haxe_constraints_rest_direct --update`
- `python3 test/run-snapshots.py --case stdlib/haxe_constraints_rest_direct --runtime`

Observed result:

- Direct `haxe.Constraints.IMap` assignment and `copy()` parity now compile and run correctly.
- Direct `haxe.Rest` `toArray()` / `append()` / `prepend()` behavior now lowers through explicit slice-cloning paths instead of falling through to missing raw-slice methods.

### 2026-03-06: `StringTools` moved from compiler shims to staged std (`haxe.go-14as.34`)

Implementation:

- Added `std/StringTools.cross.hx` and moved portable helper semantics there instead of emitting `StringTools_*` declarations from `GoCompiler`.
- Removed compiler-owned `StringTools` declarations from `lowerStdlibSymbolShimDecls`.
- Switched `StringTools` inclusion to on-demand staged-stdlib ownership so only real user-code references pull `StringTools` and its iterator helper modules.
- Removed `StringTools` from the `stdlib_symbols` shim-group classifier.
- Updated the stdlib provenance ledger and ownership map so governance points at the staged source instead of compiler helpers.

Validation evidence:

- Red/green semantic-diff:
  - `python3 test/run-semantic-diff.py --case stringtools_cross_std_contract`

### 2026-03-06: direct `haxe.Log` / `haxe.Resource` / `haxe.SysTools` support and explicit helper blockers (`haxe.go-14as.25`)

Implementation:

- Routed direct `haxe.Log` and `haxe.Resource` references through source-owned std inclusion instead of leaving them as missing symbols in generated output.
- Added `std/haxe/SysTools.cross.hx` with explicit `What / Why / How` HaxeDoc and moved direct quoting helper semantics into staged std instead of adding compiler-owned raw-Go helpers.
- Added direct semantic-diff coverage:
  - `test/semantic_diff/direct_haxe_helpers_contract`
  - `test/semantic_diff/direct_haxe_resource_contract`
- Added explicit compile-time blocker fixtures for the remaining direct-helper surfaces that are not honestly portable-ready yet:
  - `test/snapshot/negative/direct_haxe_template_unsupported`
  - `test/snapshot/negative/direct_haxe_value_exception_unsupported`
- Split the remaining debt into focused blocker beads:
  - `haxe.go-14as.38` (`haxe.Template`)
  - `haxe.go-14as.39` (`haxe.ValueException`)

Validation evidence:

- `python3 test/run-semantic-diff.py --case direct_haxe_helpers_contract --case direct_haxe_resource_contract`
- `python3 test/run-snapshots.py --case negative/direct_haxe_template_unsupported --case negative/direct_haxe_value_exception_unsupported`

Observed result:

- Direct `haxe.Log`, `haxe.Resource`, and `haxe.SysTools` references no longer fall off the backend as undefined symbols.
- `haxe.Template` and `haxe.ValueException` no longer fail late in `go test`; they now fail fast with named blockers and regression fixtures until the underlying architecture gaps are closed.
- Regression coverage on existing callers:
  - `python3 test/run-semantic-diff.py --case stringtools_math --case file_read_write_contract --case process_echo_contract --case http_proxy_custom_request --case http_request_callbacks_contract`
- Snapshot coverage:
  - `python3 test/run-snapshots.py --case stdlib/stringtools_cross_std_basic --update --runtime`
  - `python3 test/run-snapshots.py --pattern '^(stdlib/(crypto_xml_zip_basic|date_path_basic|dynamic_access_basic|int64_parity|math_basic|option_enum_basic|stringtools_basic|stringtools_cross_std_basic|unicode_string_basic|xml_root_dom_basic)|sys/filesystem_basic_smoke)$' --update`

Observed result:

- `StringTools` no longer bloats `GoCompiler` with library semantics.
- Existing StringTools consumers now compile through the staged std source, and iterator helper modules are emitted as normal modules rather than shim-group fragments.

### 2026-03-06: `DateTools` helper surface moved from compiler shims to staged std (`haxe.go-14as.35`)

Implementation:

- Added `std/DateTools.cross.hx` and moved `DateTools.format`, `getMonthDays`, `parse`, and `make` into staged std instead of keeping helper semantics in `GoCompiler`.
- Removed compiler-owned `DateTools_format` declarations from `lowerStdlibSymbolShimDecls`.
- Reused the on-demand staged-stdlib inclusion path so `DateTools` is only compiled when user code actually references it.
- Kept core `Date` representation in compiler ownership, but added the missing `getDay`, `getMinutes`, and `getSeconds` accessors required by the staged helper surface.
- Updated the stdlib provenance ledger and ownership map so governance points at the staged source instead of compiler helpers.

Validation evidence:

- Red/green semantic-diff:
  - `python3 test/run-semantic-diff.py --case datetools_cross_std_contract`
- Snapshot coverage:
  - `python3 test/run-snapshots.py --case stdlib/datetools_cross_std_basic --update --runtime`
- Existing mixed contract coverage:
  - `python3 test/run-semantic-diff.py --case stringbuf_datetools_lambda_contract`

Observed result:

- `DateTools` helper semantics no longer live in `GoCompiler`.
- Date formatting now follows the staged std implementation, while core `Date` storage and time conversion remain compiler-owned.

### 2026-03-06: `haxe.io.Path` moved from compiler shims to staged std (`haxe.go-14as.36`)

Implementation:

- Added `std/haxe/io/Path.cross.hx` and moved `Path` parsing, formatting, normalization, trailing-slash helpers, and absolute-path checks into staged std.
- Removed compiler-owned `haxe__io__Path` constructor and `haxe__io__Path_join` declarations from `lowerStdlibSymbolShimDecls`.
- Reused the on-demand staged-stdlib inclusion path so `haxe.io.Path` is only compiled when user code actually references the class.
- Kept the override narrowly expressed in target-supported string primitives and documented explicitly why upstream `haxe.io.Path` could not be reused unchanged yet.
- Updated the stdlib provenance ledger and ownership map so governance points at the staged source instead of compiler helpers.

Validation evidence:

- Red/green semantic-diff:
  - `python3 test/run-semantic-diff.py --case path_cross_std_contract`
- Existing mixed contract coverage:
  - `python3 test/run-semantic-diff.py --case option_date_path`
- Snapshot coverage:
  - `python3 test/run-snapshots.py --case stdlib/path_cross_std_basic --case stdlib/date_path_basic --update --runtime`

Observed result:

- `haxe.io.Path` no longer bloats `GoCompiler` with library semantics.
- The staged override preserves the upstream helper surface while making the current lowering gaps explicit and local to `std/`.

## Open migration track

- Legacy `haxe.go-7zy.*` shim migration sequence is closed.
- Staged stdlib migration follow-ups continue under `haxe.go-cgk.*` (portable parity program).
