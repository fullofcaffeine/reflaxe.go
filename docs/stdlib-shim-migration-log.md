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

### 2026-03-07: inherited IO helper loop extraction (`haxe.go-14as.52`)

Implementation:

- Added staged helper source:
  - `std/haxe/io/GoIoHelpers.cross.hx`
- Moved inherited `haxe.io.Input` / `haxe.io.Output` loop bodies out of `src/reflaxe/go/GoCompiler.hx`:
  - `readAll`
  - `readFullBytes`
  - `read`
  - `readUntil`
  - `readLine`
  - `write`
  - `writeFullBytes`
  - `writeInput`
  - `writeString`
- Kept compiler ownership for:
  - public wrapper symbols (`haxe__io__input_*`, `haxe__io__output_*`)
  - base IO type-shape generation
  - numeric IO helpers and RawNative/cache-coupled bytes/string paths
- Taught `GoCompiler` to resolve and queue `haxe.io.GoIoHelpers` on demand instead of requiring global inclusion.

Validation evidence:

- Ownership snapshot:
  - `python3 test/run-snapshots.py --case stdlib/io_helper_source_owned_ownership --update`
- Focused semantic parity:
  - `python3 test/run-semantic-diff.py --case io_input_output_helpers_contract --case bytes_io_stream_contract`
- Focused upstream sweep:
  - `python3 test/run-upstream-stdlib-sweep.py --strict --go-test --module haxe.io.Input --module haxe.io.Output --module haxe.io.BytesInput --module haxe.io.BytesOutput`

Observed result:

- Generated output now routes inherited IO helper wrappers through `haxe__io__GoIoHelpers_*` symbols instead of embedding those loop bodies directly in compiler-owned declarations.
- The staged helper extraction leaves app/example boundary policy unchanged because the helper is ordinary staged std source, not app-level raw injection.

Observed result:

- The repo now states the stronger default rule explicitly: library-expressible stdlib does not belong in `GoCompiler` unless there is a concrete compiler-only reason.
- Future migration work is split into concrete beads instead of one-off local exceptions.

### 2026-03-07: post-`__go__` ownership audit for compiler shim groups (`haxe.go-14as.50`)

Implementation:

- Re-audited the largest `GoStmt.GoRaw` / `GoExpr.GoRaw` clusters in `src/reflaxe/go/GoCompiler.hx` after restoring backend-owned `__go__` lowering and scoped `@:goAllowRaw`.
- Recorded the new ownership rule in `docs/stdlib-shim-rationale.md`:
  - typed extern metadata first when a real Go API exists,
  - framework-owned `@:goAllowRaw` + `reflaxe.go.macros.GoInjection.__go__` helper islands second when same-package generated-type access is the real need,
  - compiler-owned lowering last when compiler context or representation policy still matters.
- Classified the top migration candidates:
  - `haxe.io.Bytes` algorithmic helpers
  - `haxe.io.Input` / `haxe.io.Output` helper loops
  - `sys.Http` payload/proxy leaf helpers (while keeping request/callback choreography compiler-owned)
- Explicitly kept `regex_serializer` and `net_socket` compiler-owned until a materially better parity-preserving path exists.
- Opened focused follow-up beads:
  - `haxe.go-14as.51`
  - `haxe.go-14as.52`
  - `haxe.go-14as.53`
- Recorded sibling-family alignment context already in flight:
  - `haxe.rust-oo3.19`
  - `haxe.ocaml-x0r2`

Validation evidence:

- Source audit:
  - `rg -n "GoStmt\\.GoRaw|GoExpr\\.GoRaw" src/reflaxe/go/GoCompiler.hx`
  - focused scans of `lowerIoStdlibShimDecls`, `lowerHttpStdlibShimDecls`, `lowerRegexSerializerShimDecls`, and `lowerNetSocketShimDecls`
- Repo hygiene:
  - `git diff --check`

Observed result:

- The repository now has an explicit post-`__go__` ownership policy instead of an implicit “leave it in `GoCompiler` because raw injection is awkward” bias.

### 2026-03-07: `sys.Http` leaf helper extraction (`haxe.go-14as.53`)

Implementation:

- Added staged helper module:
  - `std/sys/GoHttpHelpers.cross.hx`
- Routed `sys.Http.getResponseHeaderValues` through the staged helper instead of keeping the lookup glue in `GoCompiler`.
- Routed HTTP payload capture fan-out for `customRequest` through the staged helper instead of keeping the type-switch leaf block in `GoCompiler`.
- Kept request lifecycle, callback ordering, response normalization, and proxy URL construction compiler-owned.
- Used framework-owned `@:goAllowRaw` bridges only for the hidden generated-shape access that staged Haxe cannot express directly.

Validation evidence:

- `python3 test/run-snapshots.py --case sys/http_helper_source_owned_ownership --case sys/http_custom_request_parity --case sys/http_proxy_socket_contract --update`
- `python3 test/run-semantic-diff.py --case http_request_callbacks_contract --case http_proxy_custom_request`

Observed result:

- `sys.Http` now uses a source-owned helper island for the first safe leaf extractions without weakening the existing request/callback semantic contract.
- Header lookup kept the case-insensitive fallback behavior by normalizing keys inside the staged raw bridge instead of relying on unsupported staged `String.toLowerCase()` lowering.
- Future stdlib work can use framework-owned raw helper islands without weakening the app/example boundary rules or replacing typed extern metadata.

### 2026-03-07: explicit keep decision for RawNative `haxe.io.Bytes` helpers (`haxe.go-14as.54`)

Implementation:

- Kept the RawNative/cache-coupled `haxe.io.Bytes` string helpers in `src/reflaxe/go/GoCompiler.hx`:
  - `haxe__io__bytes_fromStringRawNativeUTF16LE`
  - `haxe__io__bytes_toStringRawNativeUTF16LE`
  - `haxe__io__Bytes_ofString`
  - `haxe__io__Bytes.getString`
  - `hxrt_haxeBytesToRaw`
- Added an ownership regression snapshot:
  - `test/snapshot/stdlib/bytes_raw_native_compiler_ownership`
- Tightened the ownership rationale in the shim strategy docs instead of treating this subset as an unreviewed leftover.

Validation evidence:

- `python3 test/run-snapshots.py --case stdlib/bytes_raw_native_compiler_ownership --update`
- `python3 test/run-semantic-diff.py --case io_encoding_contract`

Observed result:

- The retained subset is compiler-owned for a concrete reason, not historical inertia: it co-owns raw-native mode dispatch and `__hx_raw` cache invalidation on generated `haxe__io__Bytes`.
- The ownership snapshot proves both generated shape and public behavior:
  - RawNative UTF-16LE conversion helpers remain compiler-emitted in `main.go`
  - mutating `Bytes` before `Base64.encode(...)` still updates the raw-byte view seen by downstream consumers
  - RawNative `BytesOutput` / `BytesInput` roundtrip behavior remains unchanged

### 2026-03-07: first `haxe.io.Bytes` helper extraction to `hxrt` (`haxe.go-14as.51`)

Implementation:

- Added pure byte-helper functions to `runtime/hxrt/bytes.go`:
  - `BytesOfHex`
  - `BytesToHex`
  - `BytesBufferAddByte`
  - `BytesBufferAdd`
  - `BytesBufferAddSlice`
  - `BytesBufferLength`
- Reworked the `io` shim wrappers in `src/reflaxe/go/GoCompiler.hx` so:
  - `haxe__io__Bytes_ofHex`
  - `haxe__io__Bytes.toHex`
  - `haxe__io__BytesBuffer.addByte`
  - `haxe__io__BytesBuffer.add`
  - `haxe__io__BytesBuffer.addBytes`
  - `haxe__io__BytesBuffer.get_length`
  now delegate to `hxrt` instead of embedding their own loop bodies.
- Added a generated-shape regression snapshot:
  - `test/snapshot/stdlib/bytes_helper_runtime_ownership`
- Kept the RawNative/cache-coupled string path compiler-owned for now:
  - `haxe__io__Bytes_ofString`
  - `haxe__io__Bytes.getString`
  - `haxe__io__bytes_fromStringRawNativeUTF16LE`
  - `haxe__io__bytes_toStringRawNativeUTF16LE`
- Opened follow-up `haxe.go-14as.54` so the remaining compiler-owned subset stays explicit instead of quietly lumped into “Bytes already migrated”.

Validation evidence:

- `python3 test/run-snapshots.py --case stdlib/bytes_helper_runtime_ownership`
- `python3 test/run-semantic-diff.py --case bytes_hex_contract --case bytes_normalization_contract`

Observed result:

- The pure `Bytes` hex/buffer leaf operations are no longer large raw Go blocks inside `GoCompiler`.
- The remaining compiler-owned `Bytes` helpers now have one explicit reason: they still co-own cache and RawNative encoding semantics.

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

Historical note:

- The direct `haxe.Template` blocker listed below was closed on 2026-03-08 by `haxe.go-14as.38`.
- The direct `haxe.ValueException` blocker listed below was closed on 2026-03-08 by `haxe.go-14as.39`.

Implementation:

- Routed direct `haxe.Log` and `haxe.Resource` references through source-owned std inclusion instead of leaving them as missing symbols in generated output.
- Added `std/haxe/SysTools.cross.hx` with explicit `What / Why / How` HaxeDoc and moved direct quoting helper semantics into staged std instead of adding compiler-owned raw-Go helpers.
- Added direct semantic-diff coverage:
  - `test/semantic_diff/direct_haxe_helpers_contract`
  - `test/semantic_diff/direct_haxe_resource_contract`
- Added explicit compile-time blocker fixtures for the remaining direct-helper surfaces that were not honestly portable-ready yet:
  - `test/snapshot/negative/direct_haxe_template_unsupported`
  - `test/snapshot/negative/direct_haxe_value_exception_unsupported`
- Split the remaining debt into focused blocker beads:
  - `haxe.go-14as.39` (`haxe.ValueException`)

Validation evidence:

- `python3 test/run-semantic-diff.py --case direct_haxe_helpers_contract --case direct_haxe_resource_contract`
- `python3 test/run-snapshots.py --case negative/direct_haxe_value_exception_unsupported`

Observed result:

- Direct `haxe.Log`, `haxe.Resource`, and `haxe.SysTools` references no longer fall off the backend as undefined symbols.
- `haxe.ValueException` no longer failed late in `go test`; at that point it failed fast with a named blocker and regression fixture until the remaining boxing/message parity gap was closed.
- Regression coverage on existing callers:
  - `python3 test/run-semantic-diff.py --case stringtools_math --case file_read_write_contract --case process_echo_contract --case http_proxy_custom_request --case http_request_callbacks_contract`
- Snapshot coverage:
  - `python3 test/run-snapshots.py --case stdlib/stringtools_cross_std_basic --update --runtime`
  - `python3 test/run-snapshots.py --pattern '^(stdlib/(crypto_xml_zip_basic|date_path_basic|dynamic_access_basic|int64_parity|math_basic|option_enum_basic|stringtools_basic|stringtools_cross_std_basic|unicode_string_basic|xml_root_dom_basic)|sys/filesystem_basic_smoke)$' --update`

Observed result:

- `StringTools` no longer bloats `GoCompiler` with library semantics.
- Existing StringTools consumers now compile through the staged std source, and iterator helper modules are emitted as normal modules rather than shim-group fragments.

### 2026-03-08: direct `haxe.Template` support via staged std override (`haxe.go-14as.38`)

Implementation:

- Replaced the old direct `haxe.Template` hard-fail with real module inclusion.
- Added `std/haxe/Template.cross.hx` so `haxe.go` owns the portable `Template` constructor/execute contract in staged std code instead of trying to force the untouched upstream module through unsupported source-owned assumptions.
- Extended the backend/runtime just enough to support that contract cleanly:
  - direct `Template.execute(context)` now gets an explicit omitted-`macros` bridge instead of relying on a generic instance-method default-argument padding rule
  - `Reflect_getProperty`, `Reflect_isObject`, and `Reflect_callMethod` are emitted only when the new `template_support` shim group is required
  - a narrow package-level helper (`haxe__Template_anyArrayToSlice_runtime`) bridges dynamic Go slice/array iteration back to staged std code without bloating `hxrt/core.go`
- Added direct parity coverage:
  - `test/semantic_diff/haxe_template_contract`
  - `test/snapshot/stdlib/haxe_template_basic`
- Retired the old negative blocker fixture:
  - `test/snapshot/negative/direct_haxe_template_unsupported`

Validation evidence:

- `python3 test/run-semantic-diff.py --case haxe_template_contract`
- `python3 test/run-snapshots.py --case stdlib/haxe_template_basic --update`

Observed result:

- Direct `new haxe.Template(...).execute(...)` now compiles and runs under the portable contract.
- The ownership boundary is clearer: `Template` semantics now live in staged std code, while the compiler/runtime only provide the narrow helper surfaces the override genuinely needs.

### 2026-03-07: `haxe.Resource` embedded payload table wired from compiler resources (`haxe.go-14as.30`)

Implementation:

- Added runtime snapshot coverage in `test/snapshot/stdlib/haxe_resource_embedded_basic` with a real `--resource greet.txt@greet` input.
- Kept the `haxe.Resource` methods source-owned, but moved the actual `content` payload population into `GoCompiler`.
- Materialized `haxe__Resource_content` from `Context.getResources()` as a deterministic `[]map[string]any` literal, encoding every payload into the existing base64-backed `data` field so std `getString` / `getBytes` semantics stay unchanged.

Validation evidence:

- Red snapshot:
  - `python3 test/run-snapshots.py --case stdlib/haxe_resource_embedded_basic --runtime`
- Green validation:
  - `python3 test/run-snapshots.py --case stdlib/haxe_resource_embedded_basic --runtime --update`
  - `python3 test/run-semantic-diff.py --case direct_haxe_resource_contract`
  - `python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_full.txt --strict --go-test --module haxe.Resource`

Observed result:

- `haxe.Resource` no longer has the “compiles but empty at runtime” failure mode when `--resource` is used.
- Ownership is explicit: std methods stay source-owned, while backend resource extraction remains compiler-owned because the data originates in compiler resources rather than reusable Haxe source.

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

### 2026-03-06: legacy text tranche closure (`haxe.go-14as.29`)

Implementation:

- Added `std/haxe/Utf8.cross.hx` with explicit `What / Why / How` HaxeDoc and moved deprecated `haxe.Utf8` helper semantics into staged std instead of growing compiler-owned text shims.
- Kept `haxe.Ucs2` on the upstream platform-exclusion path and promoted that behavior to explicit snapshot coverage through `stdlib/haxe_ucs2_platform_exclusion`.
- Added direct parity coverage:
  - `test/semantic_diff/haxe_utf8_contract`
  - `test/snapshot/stdlib/haxe_utf8_basic`
- Added explicit exclusion coverage for the still-unresolved optional constructor shape:
  - `test/snapshot/negative/direct_haxe_utf8_size_ctor_unsupported`
- Split the optional size-constructor residue into follow-up `haxe.go-14as.42` instead of leaving the whole legacy-text tranche open.
- Updated parity promotions, inventory notes, and provenance records for the new staged std override.

Validation evidence:

- `python3 test/run-semantic-diff.py --case haxe_utf8_contract`
- `python3 test/run-snapshots.py --case stdlib/haxe_utf8_basic --runtime --update`
- `python3 test/run-snapshots.py --case stdlib/haxe_ucs2_platform_exclusion --runtime`

Observed result:

- `haxe.Utf8` no longer falls through backend gaps around source-owned constructor lowering, `String.fromCharCode`, string comparison, and callback typing.
- `haxe.Ucs2` is no longer anonymous compile-only debt; its Go behavior is an explicit platform exclusion with a named snapshot contract.
- The staged override preserves the upstream helper surface while making the current lowering gaps explicit and local to `std/`.

### 2026-03-06: stack/main-loop `haxe.misc` tranche moved from compile-only debt to explicit snapshot scope (`haxe.go-14as.28`)

Implementation:

- Added `std/haxe/CallStack.cross.hx` and `std/haxe/NativeStackTrace.cross.hx` with explicit `What / Why / How` HaxeDoc.
- Kept these stack APIs deterministic and target-sensitive on Go:
  - `CallStack.callStack()` / `exceptionStack()` return `[]`
  - `CallStack.toString()` returns `""`
  - `NativeStackTrace` exposes empty-stack fallbacks instead of pretending native capture exists
- Wired source-owned std inclusion for direct:
  - `haxe.EntryPoint`
  - `haxe.MainLoop`
  - `haxe.Timer`
  - `haxe._CallStack.CallStack_Impl_`
- Added direct snapshot coverage in `test/snapshot/stdlib/haxe_stack_loop_target_sensitive`.
- Updated parity inventory/docs so these modules are snapshot-classified instead of lingering as generic compile-only debt.

Validation evidence:

- `python3 test/run-snapshots.py --case stdlib/haxe_stack_loop_target_sensitive --update`
- `python3 test/run-snapshots.py --case stdlib/haxe_stack_loop_target_sensitive`

Observed result:

- Direct `haxe.CallStack` / `haxe.NativeStackTrace` usage now compiles with explicit deterministic fallback behavior on Go.
- Direct `haxe.EntryPoint` / `haxe.MainLoop` / `haxe.Timer` usage is now compiled through source-owned std inclusion instead of falling off the backend as missing symbols.
- The tranche is now explicitly documented as target-sensitive snapshot coverage rather than portable semantic-diff parity.

### 2026-03-08: direct exception + source-owned collection parity promoted (`haxe.go-14as.43`, `haxe.go-14as.46`)

Implementation:

- Added staged direct exception overrides:
  - `std/haxe/exceptions/PosException.cross.hx`
  - `std/haxe/exceptions/ArgumentException.cross.hx`
  - `std/haxe/exceptions/NotImplementedException.cross.hx`
- Added staged direct collection overrides:
  - `std/haxe/ds/BalancedTree.cross.hx`
  - `std/haxe/ds/GenericStack.cross.hx`
- Tightened exception-family lowering so subclass `toString()` overrides stay callable while `.message` still uses the shared hxrt carrier.
- Added direct semantic-diff coverage:
  - `test/semantic_diff/haxe_exceptions_direct_contract`
  - `test/semantic_diff/haxe_ds_source_owned_collections_contract`
- Added direct snapshot coverage:
  - `test/snapshot/stdlib/haxe_exceptions_direct`
  - `test/snapshot/stdlib/haxe_ds_source_owned_collections`

Observed result:

- Direct `haxe.exceptions.PosException`, `ArgumentException`, and `NotImplementedException` construction now preserves message and subclass `toString()` parity on Go.
- Direct `haxe.ds.BalancedTree` and `haxe.ds.GenericStack` runtime use no longer falls off compiler/source-owned ownership gaps for the covered set/get/remove/pop/toString surface.
- Iterator-typed generic object parity remains tracked separately; this closure only covers the direct collection runtime surface exercised by the new contracts.

## Open migration track

- Legacy `haxe.go-7zy.*` shim migration sequence is closed.
- Staged stdlib migration follow-ups continue under `haxe.go-cgk.*` (portable parity program).

### 2026-03-24: direct Haxe event-loop surfaces reclassified as explicit unsupported usage (`haxe.go-dt4s`)

Implementation:

- Removed the temporary source-owned inclusion route for direct:
  - `haxe.EntryPoint`
  - `haxe.MainLoop`
  - `haxe.Timer`
- Restored early compile-time failure in the source-owned std planner with an explicit ownership message:
  - direct event-loop modules are unsupported on Go until a real runtime-backed
    `sys.thread.EventLoop` / `sys.thread.Thread` contract exists
- Kept the existing negative snapshot contracts:
  - `negative/direct_haxe_entrypoint_unsupported`
  - `negative/direct_haxe_mainloop_unsupported`
  - `negative/direct_haxe_timer_unsupported`
- Reclassified the parity inventory/docs from “compile-only blocker” to “explicit unsupported surface”.

Validation evidence:

- `python3 test/run-snapshots.py --case negative/direct_haxe_entrypoint_unsupported --case negative/direct_haxe_mainloop_unsupported --case negative/direct_haxe_timer_unsupported`
- `python3 test/run-portable-stdlib-inventory.py --update`
- `python3 test/run-portable-parity-closure.py`

Observed result:

- The compiler no longer accepts these modules and then emits broken Go later.
- The repo now states one consistent thing across planner, tests, and docs:
  direct Haxe event-loop surfaces are explicitly unsupported on Go today.
- Future real support is deferred to the `sys.thread` runtime tranche (`haxe.go-14as.19`), which is where a runtime-backed event-loop contract would belong.

### 2026-03-24: direct `haxe.http` baseline promoted and RTTI blocker split (`haxe.go-14as.14`)

Implementation:

- Added staged std ownership for direct `haxe.http.HttpBase` through:
  - `std/haxe/http/HttpBase.cross.hx`
- Routed direct `haxe.http.HttpBase` usage through source-owned std planning in:
  - `src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx`
- Added direct parity evidence:
  - `test/semantic_diff/haxe_http_base_contract`
  - `test/snapshot/stdlib/haxe_http_base_direct`
- Added explicit negative contracts for target-conditional modules that should stay unsupported on Go:
  - `test/snapshot/negative/direct_haxe_httpjs_unsupported`
  - `test/snapshot/negative/direct_haxe_httpnodejs_unsupported`
- Split the old mixed HTTP/RTTI blocker so direct RTTI reflection now has its own explicit tracker:
  - `haxe.go-14as.57`

Validation evidence:

- `python3 test/run-semantic-diff.py --case haxe_http_base_contract`
- `python3 test/run-snapshots.py --case stdlib/haxe_http_base_direct`
- `python3 test/run-snapshots.py --case negative/direct_haxe_httpjs_unsupported --case negative/direct_haxe_httpnodejs_unsupported`
- `python3 test/run-portable-stdlib-inventory.py --update`
- `python3 test/run-portable-parity-closure.py`

Observed result:

- Direct `haxe.http.HttpBase` constructor/base-field/request baseline is now supported under staged std ownership instead of compiling with a missing emitted class body.
- Direct `haxe.http.HttpMethod` and `haxe.http.HttpStatus` now ride the same parity evidence instead of staying compile-only by omission.
- `haxe.http.HttpJs` and `haxe.http.HttpNodeJs` are now documented honestly as explicit unsupported target-conditional modules on Go.
- The remaining direct `haxe.rtti.*` reflection debt is tracked separately instead of being hidden inside the HTTP tranche.
