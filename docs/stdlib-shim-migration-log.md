# Stdlib Shim Migration Log

This log tracks the end-to-end shim migration process so decisions, rollout order, and validation evidence stay auditable.

How to use this log:

1. Read `docs/stdlib-shim-rationale.md` first for current ownership decisions.
2. Use this page as the chronological execution record (what changed, when, and how it was validated).
3. Use linked test commands to reproduce evidence on current HEAD.

Terms:

- **shim**: compatibility glue between Haxe std APIs and Go output behavior.
- **semantic-diff**: runtime parity harness comparing generated Go output against Haxe `--interp`.
- **staged stdlib**: target-specific override source under `std/go/_std`,
  converted to package-only `.cross.hx` artifacts during staging.

## Process Template

For each shim surface:

1. Scope and classify the surface in `docs/stdlib-shim-rationale.md`.
2. Pick target ownership (`compiler core`, `runtime-lowered`, canonical staged
   std, or ordinary target support).
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
  - `Sys_command` -> `hxrt.SysCommand`
  - `Sys_exit` -> `hxrt.SysExit`
  - `sys__io__File_saveContent` -> `hxrt.FileSaveContent`
  - `sys__io__File_getContent` -> `hxrt.FileGetContent`
  - `New_sys__io__Process` -> `hxrt.NewProcess`
  - `sys__io__ProcessOutput.readLine` -> `hxrt.ProcessOutput.ReadLine`
  - `sys__io__Process.close` -> `hxrt.Process.Close`
- Added runtime-owned behavior to `runtime/hxrt/hxrt.go`:
  - `SysGetCwd`, `SysArgs`, `FileSaveContent`, `FileGetContent`
  - `SysCommand`, `SysExit`
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

### 2026-05-11: `Sys.command` / `Sys.exit` wrapper delegation (`haxe.go-14as.49`)

Implementation:

- Added runtime-owned `hxrt.SysCommand` that runs a child process with inherited stdin/stdout/stderr and returns the child exit code.
- Added runtime-owned `hxrt.SysExit` so emitted CLI wrappers can terminate with the delegated command status.
- Added thin compiler forwarding declarations for generated `Sys_command` and `Sys_exit`.
- Extended the snapshot runtime harness with `expected.exit` support so nonzero CLI wrapper exits are verified against a built binary rather than hidden behind `go run` exit wrapping.

Validation evidence:

- `python3 test/run-semantic-diff.py --case sys_command_contract --timeout 120`
- `python3 test/run-snapshots.py --case sys/sys_command_exit_wrapper --runtime --timeout 120`

Observed result:

- Emitted Go can use the canonical `Sys.command(...)` + `Sys.exit(...)` wrapper shape without undefined symbols, while preserving child stdout/stderr inheritance and propagated exit code.

### 2026-07-14: portable `Sys.sleep` runtime delegation (`haxe_go-vfp.8.7.1`)

Implementation:

- Added the missing typed `Sys_sleep(Float)` generated adapter and delegated it to runtime-owned `hxrt.SysSleep`.
- Kept the mainstream Haxe 4.3.7 root `Sys` declaration instead of duplicating the compiler-owned carrier surface in a whole-class staged override.
- Converted Haxe seconds to Go `time.Duration` in `runtime/hxrt/sys.go`; non-positive and NaN values return immediately.
- Added a direct runtime unit contract plus snapshot and semantic-diff timing contracts with broad lower/upper bounds.
- Compared the adjacent root `Sys` surface with the upstream declaration. The remaining missing APIs are isolated in `haxe_go-vfp.8.7.2` rather than being folded into this regression fix.

Ownership decision:

- This follows the sibling `haxe.rust` shape: the source-facing `Sys.sleep` adapter is thin and target runtime code owns the platform duration conversion and blocking primitive.
- `lowerSysStdlibShimDecls` remains adapter-only and gains no behavior-heavy `GoRaw`.

Validation evidence:

- `GO111MODULE=off go test ./runtime/hxrt`
- `python3 test/run-snapshots.py --case sys/sys_sleep_portable --runtime`
- `python3 test/run-semantic-diff.py --case sys_sleep_contract`
- `npm test` (`251/251` snapshots)
- `npm run test:semantic-diff` (`130/130` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` modules)
- `npm run test:examples` (`12/12` lanes)

### 2026-07-14: complete portable root `Sys` surface (`haxe_go-vfp.8.7.2`)

Implementation:

- Added typed generated adapters for first-class and direct `Sys.print`,
  `setTimeLocale`, `setCwd`, `time`, `programPath`, deprecated
  `executablePath`, `getChar`, `stdin`, `stdout`, and `stderr` usage.
- Added runtime-owned wall-clock, cwd, executable-path, byte-input, and
  standard-stream helpers. Standard streams are non-owning and avoid invalid
  `Sync` calls; ordinary file handles remain owning and sync on flush.
- Added an explicit Haxe compile-time error for `Sys.cpuTime`. Go's standard
  library has no portable process CPU clock, and wall-clock substitution would
  falsify the public contract.
- Kept `lowerSysStdlibShimDecls` adapter-only. Its raw statements translate
  typed error/EOF status and attach existing Haxe IO carriers; they do not own
  OS behavior. A typed `GoMultiAssign` AST statement now represents hxrt's
  `(value, error)` / `(value, eof, error)` results, so this work adds no
  compiler `GoRaw` debt.

Sibling and review decision:

- `haxe.rust` confirms the semantic split: typed runtime helpers own OS
  behavior, `setTimeLocale` reports false, standard streams are explicit, and
  CPU time is rejected rather than faked. Go retains generated adapters instead
  of copying Rust's whole staged root class because Go's file/process carriers
  already share that generated boundary.
- The local `haxe.elixir` tree has no production root `Sys` override, and the
  `haxe.ruby` placeholder is not parity evidence.
- This bounded `thinking:high` choice had one honest design after local tracing,
  so no Oracle escalation was needed. Global `metal` selector deprecation
  remains the independent `thinking:xhigh` decision in `haxe_go-vfp.6.6`.

Validation evidence:

- `GO111MODULE=off go test ./runtime/hxrt`
- `python3 test/run-snapshots.py --case sys/root_sys_portable --case negative/sys_cpu_time_unsupported`
- `python3 test/run-snapshots.py --case core/ast_multi_assign_stmt_printer`
- `python3 test/run-semantic-diff.py --case root_sys_portable_contract`
- `npm test` (`254/254` snapshots; compiler debt unchanged)
- `npm run test:semantic-diff` (`131/131` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` modules)
- `npm run test:examples` (`12/12` lanes)

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
- Added the sibling handoff note in
  `docs/spikes/family-raw-injection-authority-alignment.md`, which preserves the
  Go reasoning while requiring `haxe.rust` and `haxe.ocaml` agents to compare
  against their own local architecture before adopting any allow-tag mechanism.

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
- Date formatting moved to staged std in this tranche. The core `Date` storage and time conversion emitter remained temporarily and is now explicit migration debt under `haxe_go-vfp.8.7.15` rather than an approved compiler boundary.

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

- `haxe.io.Path` later graduated back to the upstream Haxe stdlib implementation once the reusable string/array lowerings landed (`haxe.go-14as.37`).

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

### 2026-06-10: `haxe.Utf8` optional size constructor restored (`haxe.go-14as.42`)

Implementation:

- Updated `std/haxe/Utf8.cross.hx` so `new haxe.Utf8(size)` is accepted again through a typed default `Int` constructor argument.
- Removed the old negative fixture for `direct_haxe_utf8_size_ctor_unsupported`.
- Extended `haxe_utf8_contract` and `stdlib/haxe_utf8_basic` to cover both constructor forms.

Observed result:

- The deprecated size hint remains a capacity hint only and is intentionally ignored.
- Generated Go keeps the constructor typed instead of widening the ignored argument to `any`.

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
- The `haxe.go-cgk.*` planning work is historical context now, not the active execution tracker.
- Current stdlib parity status is governed by `docs/portable-stdlib-parity-program.md`,
  `test/portable_stdlib_inventory.json`, and the generated parity closure summaries
  under `test/.test-cache/`.

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
- At that point, direct Haxe event-loop surfaces were explicitly unsupported instead
  of compiling through a broken temporary route.
- That support was deferred to the `sys.thread` runtime tranche (`haxe.go-14as.19`),
  which is where a runtime-backed event-loop contract belonged.

### 2026-06-10: direct Haxe event-loop surfaces bridged to `sys.thread.EventLoop` (`haxe.go-14as.69`)

Implementation:

- Added staged std overrides:
  - `std/haxe/EntryPoint.cross.hx`
  - `std/haxe/MainLoop.cross.hx`
  - `std/haxe/Timer.cross.hx`
- Routed direct `haxe.EntryPoint`, `haxe.MainLoop`, `haxe.MainEvent`, and `haxe.Timer`
  usage through `GoSourceOwnedStdlibPlanner`.
- Added `hxrt.ThreadNowSeconds()` plus a typed `NativeThread.nowSeconds()` bridge so
  `Timer.stamp()` uses the same monotonic clock as the event-loop runtime.
- Removed the obsolete negative fixtures for direct event-loop usage.

Validation evidence:

- `python3 test/run-snapshots.py --case stdlib/haxe_main_loop_runtime_direct --runtime`
- `python3 test/run-portable-stdlib-inventory.py --update`

Observed result:

- Direct `haxe.EntryPoint.run()`, `haxe.MainLoop.add(...)`, and `haxe.Timer.delay(...)`
  now compile and run through the runtime-backed main event loop.
- Coverage is intentionally snapshot/runtime smoke rather than semantic-diff because
  event-loop timing is target-sensitive.
- The old `haxe.go-dt4s` unsupported classification is superseded by this bridge.

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

### 2026-07-15: `sys.FileSystem` moved from compiler shims to canonical staged std (`haxe_go-vfp.8.7.4`)

Implementation:

- Added `std/go/_std/sys/FileSystem.hx` as the source authority for all ten Haxe 4.3.7 methods, including the previously missing `absolutePath`.
- Reused the unchanged upstream `sys.FileStat` typedef and constructed its anonymous record in Haxe source.
- Added typed `std/hxrt/fs/NativeFileSystem.hx` and `FileSystemStat.hx` bindings over native capabilities in selectively copied `runtime/hxrt/filesystem.go`.
- Removed `lowerFileSystemShimDecls`, filesystem shim classification/dependencies/imports, the synthetic `sys__FileSystem` declaration, and the corresponding compiler-debt allowances.

Validation evidence:

- `python3 test/run-snapshots.py --case sys/filesystem_basic_smoke --runtime`
- `python3 test/run-semantic-diff.py --case filesystem_contract`
- `python3 test/test_stdlib_migration_ledger_contract.py`
- `npm run test:compiler-debt`

Observed result:

- The Haxe stdlib surface is now reviewable source under the canonical target `_std` tree, while Go-native I/O stays behind a typed runtime boundary.
- `absolutePath` accepts paths that do not exist, `fullPath` resolves existing paths and symlinks, and `isDirectory` preserves Haxe 4.3.7 interpreter behavior by returning `false` for a missing path despite the upstream API documentation describing an exception.
- Filesystem behavior no longer adds library algorithms, raw statements, support imports, or synthetic declarations to `GoCompiler`.

### 2026-07-15: `sys.io.File*` moved from compiler shims to canonical staged std (`haxe_go-vfp.8.7.5`)

Implementation:

- Added canonical `std/go/_std/sys/io/File.hx`, `FileInput.hx`, `FileOutput.hx`, and `FileSeek.hx` implementations of the Haxe 4.3.7 target extern surface.
- Added typed `std/hxrt/fs` handle/capability bindings over a dedicated, selectively copied `runtime/hxrt/file.go` slice. Arbitrary bytes cross as `Array<Int>` / `[]int`, not generated `Bytes` internals.
- Added type-only retention for exact public `Input`, `Output`, and `Bytes` fields needed when compiler-owned root `Sys` discovers staged stream subclasses after Haxe dead-code elimination.
- Removed File structs, handle maps, seek mapping, static APIs, stream methods, File ownership/classification, and File-specific subclass branches from `GoCompiler`; root `Sys` standard streams now construct the staged wrappers.
- Lowered the permanent debt ceilings from 339 to 138 `GoRaw` sites in `lowerSysStdlibShimDecls` and from 93 to 46 in generic I/O subclass synthesis.

Validation evidence:

- `python3 test/run-semantic-diff.py --case file_read_write_contract`
- `python3 test/run-semantic-diff.py --case file_error_semantics_contract`
- `python3 test/run-semantic-diff.py --case sys_db_io_contract`
- `python3 test/run-snapshots.py --case sys/root_sys_portable --runtime`
- `python3 test/run-snapshots.py --case sys/process_echo_smoke --runtime`
- `python3 test/run-snapshots.py --case stdlib/sys_db_io_direct --runtime`
- `python3 test/run-snapshots.py --case core/runtime_hxrt_infer_file --runtime`
- `python3 test/test_stdlib_migration_ledger_contract.py`
- `npm run test:compiler-debt`
- `npm test` (`261/261` snapshots)
- `npm run test:semantic-diff` (`131/131` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` modules)
- `npm run test:examples` (`12/12` lanes)
- `npm run security:go-tooling`
- `npm run test:perf:hxrt-selective`

Observed result:

- Text and arbitrary binary content, copy, write/append/update modes, seek/tell, bounds, EOF, native errors, owning ordinary handles, and non-owning standard streams are preserved through reviewable Haxe source plus typed native capabilities.
- Direct File use selects `file.go` without the broad root `sys.go` / `process.go` slices; the still-monolithic Sys/Process shim group selects the same staged wrappers and file runtime explicitly.
- `GoCompiler` is no longer a semantic or representation owner for `sys.io.File`, `FileInput`, `FileOutput`, or `FileSeek`.

### 2026-07-15: root `Sys` moved to canonical staged std (`haxe_go-vfp.8.7.6`)

Implementation:

- Added `std/go/_std/Sys.hx` as the authority for the supported Haxe 4.3.7 root surface. Environment-map construction, locale fallback, deprecated aliases, `getChar`, and standard-stream wrapping are ordinary typed Haxe source.
- Added narrow typed `std/hxrt/sys` bindings. Display keeps its unavoidable upstream `Dynamic` contract in `NativeConsole`; process state, environment, clocks, commands, cwd, and program paths remain typed in `NativeSys`; standard handles reuse `hxrt.fs.NativeFile`.
- Made one-step source methods inline so direct calls keep compact Go-shaped output while first-class references still materialize the source-owned API.
- Removed every root `Sys` struct/function and print/println semantic branch from `GoCompiler`. Renamed the remaining child-process group to `lowerProcessStdlibShimDecls` and reduced its raw debt ceiling from the former combined 138 sites to 134 Process-only sites.
- Fixed selective inference for surviving typed extern calls after source inlining. Sys-only output copies `sys.go` without `file.go`/`process.go`; Process-only output copies `process.go` without `sys.go`/`file.go`.
- Retained nominal types for DCE-reached static-only staged classes through the generic type mapper, so `var probe:Sys = null` emits an empty source-owned carrier without restoring a compiler `Sys` declaration or constructor.
- Preserved explicit compile-time rejection for `Sys.cpuTime`; Go's standard library still has no portable process CPU clock.

Validation evidence:

- `test/semantic_diff/root_sys_contract`
- `test/semantic_diff/root_sys_portable_contract`
- `test/semantic_diff/sys_command_contract`
- `test/semantic_diff/sys_sleep_contract`
- `test/snapshot/sys/root_sys_portable`
- `test/snapshot/core/runtime_hxrt_infer_sys`
- `test/snapshot/core/runtime_hxrt_infer_process`
- `test/snapshot/negative/sys_cpu_time_unsupported`
- `runtime/hxrt/sys_test.go`
- `test/test_stdlib_migration_ledger_contract.py`
- `npm run test:compiler-debt`
- `npm test` (`263/263` snapshots)
- `npm run test:semantic-diff` (`131/131` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` modules)
- `npm run test:examples` (`12/12` lanes)
- `npm run test:perf:hxrt-selective`
- `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)

Observed result:

- Root `Sys` now follows the same source-ownership rule as File/FileSystem and the sibling `haxe.rust` target: Haxe library behavior stays in Haxe source, while only genuine OS capabilities cross typed runtime bindings.
- Compiler shims are no longer the default implementation mechanism for root stdlib behavior. The remaining `sys.io.Process` adapters are isolated and explicitly tracked by `haxe_go-vfp.8.7.7`.

### 2026-07-15: `sys.io.Process` moved to canonical staged std (`haxe_go-vfp.8.7.7`)

Implementation:

- Added `std/go/_std/sys/io/Process.hx` as the authority for the complete Haxe 4.3.7 surface. Its ordinary Haxe source owns public stream construction, bounds and EOF translation, detached rejection, nullable exit status, and closed-state/lifecycle policy.
- Added five narrow `std/hxrt/process` modules: opaque process/input/output handles, a typed `{Available, Code}` exit-status carrier, and the `NativeProcess` capability bridge. The native boundary contains no `Dynamic`, `Any`, or raw injection.
- Extended `runtime/hxrt/process.go` with typed Haxe-shaped spawn, pipe, byte-transfer, PID, wait/poll, kill, and close functions while retaining the lower-level native API for direct runtime callers.
- Routed `sys.io.Process` through the source-owned planner and inferred selective `process.go` retention from the surviving typed `hxrt.process.NativeProcess` authority.
- Removed the Process classifier group, dependency branch, carrier/stream declarations, public API bodies, and inherited-helper synthesis from `GoCompiler`. No Process-specific compiler emitter remains.
- Preserved `Null<Int>` method results through the compiler's generic nil-capable primitive result mapping. This is language-level return storage, not a Process shim; the typed runtime boundary still uses an explicit status carrier.
- Removed the Process compiler-debt allowances. The measured compiler totals fell from 4,943 to 4,809 `GoRaw` sites and from 17 to 16 shim emitters.

Validation evidence:

- `runtime/hxrt/process_test.go`
- `test/semantic_diff/process_echo_contract`
- `test/semantic_diff/process_error_semantics_contract`
- `test/snapshot/sys/process_echo_smoke`
- `test/snapshot/sys/process_error_semantics`
- `test/snapshot/core/runtime_hxrt_infer_process`
- `test/test_stdlib_migration_ledger_contract.py`
- `npm run test:compiler-debt`
- `npm test` (`263/263` snapshots)
- `npm run test:semantic-diff` (`131/131` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` modules)
- `npm run test:examples` (`12/12` lanes)
- `npm run test:perf:hxrt-selective`
- `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)

Observed result:

- The last compiler-owned member of the former Sys/File/Process shim family is now reviewable Haxe stdlib source over a typed native capability boundary.
- Process-only selective output keeps `process.go` without `sys.go` or `file.go`; root Sys and direct File likewise remain independent.
- Startup failure, stdin/stdout/stderr behavior, normal EOF, nonblocking `null`, nonzero exits, kill, detached rejection, large output, and close-without-kill semantics remain covered by existing runtime and semantic contracts.

### 2026-07-15: `Sys.getChar` gained source-owned terminal semantics (`haxe_go-vfp.8.7.3`)

Implementation:

- Kept `haxe.io.Eof` construction and requested echo in canonical `std/go/_std/Sys.hx`; no compiler declaration, raw emitter, or Haxe injection was added.
- Added typed `std/hxrt/sys/NativeTerminal.hx` over a dedicated `terminal` runtime feature. Build-tagged Linux, Darwin, Windows, and unsupported-host files own only terminal state and the one-byte native read.
- Made the terminal feature footprint-explicit even in ordinary full-copy mode. Programs that do not use `Sys.getChar` do not acquire the POSIX boundary; disabling inference explicitly retains the all-files escape.
- Confined the required POSIX `unsafe.Pointer` to `terminalIoctlTermios`. The debt policy permits exactly one import and one selector, explains why the frozen `syscall` API lacks a safe wrapper, and records why neither a Go-floor-raising current `x/term` nor an advisory-bearing compatible pin was accepted.
- Serialized terminal transitions and restored the original mode on every return. Redirected input retains byte-stream EOF behavior; character-device hosts without an implementation fail explicitly.

Validation evidence:

- `test/test_sys_get_char_terminal.py` (real PTY no-newline input, echo off/on, state restoration, redirected EOF, Linux/macOS/Windows implementation cross-builds, and an unsupported-FreeBSD cross-build)
- `test/test_sys_get_char_terminal_contract.py`
- `test/snapshot/sys/sys_get_char_terminal`
- `test/snapshot/sys/root_sys_portable`
- `test/snapshot/core/runtime_hxrt_infer_sys`
- `runtime/hxrt/sys_test.go`
- `npm run test:compiler-debt`
- `npm run security:go-tooling`

Observed result:

- The admitted `linux-amd64` root `Sys.getChar` contract now reads a terminal byte immediately, suppresses host echo, restores terminal state, and emits requested echo exactly once from staged Haxe.
- Terminal control remains a narrow typed native capability, not a second Haxe stdlib in `GoCompiler` or `hxrt`.

### 2026-07-15: compiler stdlib ownership became fail-closed (`haxe_go-vfp.8.7.8`)

Implementation:

- Added `docs/compiler-stdlib-intrinsics.json` as the exact registry for portable
  compiler stdlib groups, generated authorities, class/enum and source-planner
  selectors, group dependencies, direct call rewrites, special data/diagnostic
  primitives, evidence, review conditions, and migration beads.
- Flattened `GoStdlibShimClassifier` into a machine-auditable fully qualified
  symbol table without changing selection behavior.
- Split the former generic required compiler-shim debt exception into three
  truthful categories: avoidable portable stdlib migration debt, exact required
  compiler stdlib intrinsics, and explicit required `go.*` native emitters.
- Registered the eight behavior-heavy declaration families and direct
  collection algorithms as unfinished work under `haxe_go-vfp.8.7.9` through
  `.17`. The registry does not imply that the parent migration is complete.
- Admitted only exact current compile-context or representation primitives:
  generated type/RTTI metadata, `Std.isOfType`, string representation
  construction/conversion, exception carrier access, `haxe.Rest` slice
  construction, compile-resource data population, and the honest
  `Sys.cpuTime` unsupported diagnostic.
- Wired the bidirectional gate into normal, changed, stdlib-governance, CI, and
  release-contract paths.

Validation evidence:

- `npm run test:stdlib:intrinsics`
- `npm run test:stdlib:governance`
- `npm run test:compiler-debt` (`16` named entry points: `10` portable
  migration-debt contexts, `1` exact portable intrinsic context, and `5`
  explicit native `go.*` contexts)
- `npm test` (`264` snapshots)
- `npm run test:semantic-diff` (`131` cases)
- `npm run test:stdlib-sweep:go-test` (`55` strict modules)
- `npm run test:examples` (`12` example/profile lanes)
- `npm run test:release-contracts`

### 2026-07-15: core `haxe.ds` maps and `List` moved to canonical staged std (`haxe_go-vfp.8.7.10`)

Implementation:

- Replaced compiler-generated `IntMap`, `StringMap`, `ObjectMap`, `EnumValueMap`,
  and `List` behavior with ordinary Haxe source under `std/go/_std/haxe/ds`.
  Public mutation, lookup, copy, iteration, ordering, filtering, mapping, joining,
  and tree-balancing algorithms are now reviewable as Haxe library code.
- Added narrow typed `std/hxrt/collections` bindings over four runtime features.
  The Go boundary owns only native hash storage, object-key identity, deterministic
  key snapshots, and recognition of generated enum carriers. Generic erased values
  use `Dynamic` only at those documented bindings and are cast back immediately.
- Added exact `IMap` bridge methods because Go requires identical interface method
  signatures, while Haxe permits concrete key types and covariant `copy()` results.
  The compiler now preserves interface-to-concrete casts used by the `Map` abstract,
  honors interface selector metadata, and adapts typed callbacks to shared erased Go
  method signatures without moving collection behavior back into the compiler.
- Routed serializer and Lambda integration through staged collection APIs and typed
  snapshots. Removed the complete compiler `ds` declaration group, its classifier,
  dependency wiring, ownership exception, and reflection-based collection probes.
- Added selective-runtime inference and isolated snapshots for `enum_value`,
  `map_int`, `map_object`, and `map_string`. Ordinary full-runtime output includes
  all four files; inferred output includes only features proven by typed use.
- Kept the temporary array-backed private `List` carrier while the representation
  adapters tracked by `haxe_go-vfp.8.7.17` still exist. That deferred work does not
  retain public list behavior or a compiler-owned `List` declaration.

Validation evidence:

- `npm test` (`270/270` snapshots)
- `npm run test:semantic-diff` (`131/131` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance` (`119` tracked sources: `36` typed runtime
  bindings, `73` upstream overrides, `5` staged support modules, and `5` public
  Go facades)
- `npm run test:compiler-debt` (`4,583` `GoRaw` sites and `14` compiler shim
  entry points; generated example copies of reviewed reflection boundaries are
  recorded separately from their source owners)
- `npm run test:perf:hxrt-selective`
- `npm run security:go-tooling` (all `28` race/checkptr/vet/staticcheck gates)
- `npm run test:release-contracts`

Observed result:

- Haxe stdlib collection semantics now live in staged Haxe, matching the project
  ownership rule and sibling-target precedent. The runtime is a typed capability
  layer rather than a second implementation of the public library.
- Null-valued entries, stable repeated iteration, insertion order, object identity,
  recursive enum-key comparison, independent copies, serializer round trips, and
  the complete supported `List` API remain covered by behavioral contracts.
- `GoCompiler` no longer owns the five collection declarations or their algorithms.
  The remaining collection-lowering work is explicitly limited to the separately
  tracked adapter cleanup and does not weaken this ownership boundary.

### 2026-07-15: `Lambda` algorithms moved to canonical staged std (`haxe_go-vfp.8.7.17`)

Implementation:

- Added canonical `std/go/_std/Lambda.hx` with the public collection algorithms as
  ordinary Haxe source. `@:dce` keeps unused helpers out of generated Go while
  preserving the upstream `Lambda` API surface.
- Removed the compiler-owned loops for `count`, `empty`, `exists`, `has`, `iter`,
  `filter`, `map`, and `fold`. The direct-call lowering now only converts concrete
  arrays, lists, and manual iterators to the erased structural carrier used by the
  staged module, adapts typed callbacks at that same boundary, and restores typed
  array and scalar results.
- Kept the `ArraySort` and `ListSort` bridges as exact representation intrinsics.
  They box erased generic elements, preserve the source comparator contract, and
  copy results back to the original carrier; sorting and linked-list merge behavior
  remain in upstream Haxe source.
- Registered each retained bridge separately, updated provenance and package
  inventories, and removed the old compiler-debt allowance for 95 raw algorithm
  sites. No new runtime feature or compiler-owned collection behavior was added.
- Added behavioral coverage for optional `count` predicates, first-class `Lambda`
  function values over both generic iterables and arrays, generic `fold`, and
  single-linked `ListSort` input. Preserved the upstream-inline `mapi`, `flatten`,
  and `flatMap` helpers so staging the module does not create new erased Go entry
  points; array `mapi` now has explicit semantic coverage.

Validation evidence:

- `npm test` (`270/270` snapshots)
- `npm run test:semantic-diff` (`131/131` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance` (`120` tracked sources: `36` typed runtime
  bindings, `74` upstream overrides, `5` staged support modules, and `5` public
  Go facades)
- `npm run test:compiler-debt` (`4,488` `GoRaw` sites and `14` compiler shim entry
  points)
- `npm run test:perf:go`
- `npm run test:perf:hxrt-selective`
- `npm run test:perf:stdlib-shims`
- `npm run security:go-tooling`
- `npm run test:release-contracts`

Observed result:

- Public `Lambda`, array-sort, and list-sort algorithms are source-owned. The
  compiler now knows only enough about Go's erased generic representations to make
  those source algorithms callable without changing their behavior.
- The retained adapters are closed, individually governed, and tested as
  representation boundaries rather than a general escape for compiler-side
  stdlib implementations.
- Complete carrier coverage for the other public `Lambda` helpers, including the
  existing nested-`Iterable` gap in `flatten` and `flatMap`, is intentionally
  deferred to `haxe_go-vfp.8.7.18`; this closure admits no adapter for them.

### 2026-07-16: complete staged `Lambda` API and carrier closure (`haxe_go-vfp.8.7.18`)

Implementation:

- Completed source-owned entrypoints for `array`, `list`, `mapi`, `flatten`,
  `flatMap`, `foreach`, `foldi`, `indexOf`, `find`, `findIndex`, and `concat`.
  Together with the prior tranche, all 19 public `Lambda` algorithms now execute
  from `std/go/_std/Lambda.hx`; no traversal, comparison, early-exit, allocation,
  or collection policy moved into the compiler or runtime.
- Extended the exact call adapters across arrays, staged lists, and concrete
  manual `Iterable<T>` classes, including typed indexed callbacks, nullable
  results, and mixed-carrier concatenation. Nested arrays, lists, and concrete
  iterable classes work for `flatten`; `flatMap` callbacks may return concrete
  arrays or lists.
- Added a private, representation-only `LambdaGoIterableCarrier` companion for
  the Go method interface created by constrained nested iterables. Its factory is
  retained only when `flatten` or `flatMap` is reachable, so ordinary `Lambda`
  calls do not add generated declarations or reflection-visible types.
- Added deterministic compile-time diagnostics for nested sources or callback
  results that have already been erased to structural `Iterable`. At that point
  the concrete carrier authority is no longer recoverable, so the compiler stops
  before emitting an incidental Go type error.
- Rebuilt the iterator bridge with typed Go AST assignments and function literals,
  reducing its reviewed `GoRaw` debt from four sites to two. Registered all 19
  exact intrinsics and documented the complete method/carrier matrix.

Validation evidence:

- `npm test` (`272/272` snapshots)
- `npm run test:semantic-diff` (`132/132` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:changed`
- `npm run test:compiler-debt` (`4,486` `GoRaw` sites and `14` compiler shim
  entry points)
- `npm run test:perf:go`
- `npm run test:perf:hxrt-selective`
- `npm run test:perf:stdlib-shims`
- `npm run security:go-tooling` (all `28` race/checkptr/vet/staticcheck gates)
- `npm run test:release-contracts`

Observed result:

- The mainstream Haxe stdlib remains the semantic model, with the Go override
  differing only where Go's invariant and method-interface representations make
  the upstream generic shape unassignable.
- The complete public `Lambda` API is staged source. The compiler boundary is a
  closed set of registered representation adapters, not an alternate collection
  implementation, and no new `hxrt` feature or public helper module was added.
- Concrete carrier authority produces valid Go; genuinely erased nested authority
  produces a stable Haxe diagnostic. Neither path leaks an incidental Go compiler
  error to the user.

### 2026-07-16: retire the compiler-owned `haxe.Template` runtime bridge (`haxe_go-vfp.8.7.16`)

Implementation:

- Removed the `template_support` compiler dispatcher branch, reflection import,
  four generated helper declarations, planner requirement, intrinsic-registry
  entry, and compiler-debt allowances. `GoCompiler` no longer emits Template
  array conversion or `Reflect.getProperty` / `isObject` / `callMethod` helpers.
- Added the typed `std/hxrt/template/NativeTemplate.hx` binding over
  `runtime/hxrt/template.go`. The runtime owns only three dynamic representation
  operations: exposing Go slices/arrays, classifying map/struct carriers, and
  invoking an already-resolved function with runtime arguments.
- Kept parsing, field lookup and stack fallback, iteration, macro argument
  construction, errors, and rendering in `std/go/_std/haxe/Template.hx`. The
  staged override no longer uses raw `__go__`; property reads use the existing
  `Reflect.field` contract while its broader source/runtime migration remains
  tracked by `haxe_go-vfp.8.7.15`.
- Added a footprint-explicit `template` runtime feature. Full and selective
  runtime plans copy `template.go` only when `haxe.Template` or its typed binding
  is used, so unrelated generated programs do not acquire reflection support.
- Expanded Template parity coverage across nested properties, record arrays,
  primitive stack fallback, structural dynamic iterators, and macros. Added
  direct runtime tests for slice/array conversion, object classification, and
  invocation.
- Filed `haxe_go-vfp.8.7.19` for the separately discovered concrete-class
  iterable gap: Go reflection cannot discover generated lowercase methods, and
  solving that requires a general generated-method metadata/adapter design rather
  than restoring a Template compiler shim.

Validation evidence:

- `npm run test:changed`
- `npm test` (`272/272` snapshots)
- `npm run test:semantic-diff` (`132/132` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance` (`121` tracked sources; `37` typed `hxrt`
  bindings)
- `npm run test:stdlib-inventory`
- `npm run test:compiler-debt` (`4,427` `GoRaw` sites, `597` bounded Go
  reflection sites, and `13` compiler shim entry points)
- direct `runtime/hxrt` Go tests and raw-injection hygiene
- `npm run test:perf:hxrt-selective`
- `npm run test:perf:stdlib-shims`
- `npm run test:perf:go` (four warning-only startup signals; no hard failures)
- `npm run security:go-tooling` (all `28` race/checkptr/vet/staticcheck gates)
- `npm run test:release-contracts`

Observed result:

- `haxe.Template` is now a source-owned public API over a narrow typed runtime
  capability. No Template algorithm, lookup rule, iteration loop, or macro policy
  originates in the compiler or `hxrt`.
- Compiler debt falls permanently by one shim entry point and 59 raw sites. The
  remaining dynamic and reflection counts are declared at the real runtime
  boundary instead of being hidden inside compiler-emitted library functions.
- Template-only reflection support is absent from unrelated output, while direct
  Template use remains portable and matches the staged Haxe reference contract.

### 2026-07-16: move `haxe.crypto` APIs to staged std and typed runtime capabilities (`haxe_go-vfp.8.7.15.1`)

Implementation:

- Tried the unchanged Haxe 4.3.7 crypto sources first. Their shared `BaseCode`
  helper creates an empty array and fills it with indexed writes; generated Go
  currently turns those writes into fixed-length slice assignments and panics.
  The general array-growth fix is tracked separately by `haxe_go-vfp.8.7.20`
  instead of being hidden behind a crypto-specific compiler exception.
- Added canonical staged overrides for `Base64`, `Md5`, `Sha1`, `Sha224`, and
  `Sha256`. Haxe source owns the public APIs, default arguments, Base64 alphabet
  and padding rules, and conversion to and from `haxe.io.Bytes`.
- Added the typed `std/hxrt/crypto/NativeCrypto.hx` boundary over
  `runtime/hxrt/crypto.go`. The Go runtime owns only native Base64 and digest
  execution over strings and integer byte arrays; generated `Bytes` fields do
  not cross the boundary.
- Removed all five crypto declarations, their imports, and their classifier
  routes from the monolithic `stdlib_symbols` compiler group. Crypto use now
  selects ordinary staged source, and a footprint-explicit `crypto` feature
  copies the native runtime file only when one of those APIs is reachable.
- Updated the provenance ledger, intrinsic registry, inventory, package-layout
  status, and fail-closed ownership tests. Governance now covers 127 staged
  sources, including 38 typed `hxrt` bindings and 79 upstream overrides.
- Modernized the stdlib-boundary performance harness to benchmark the active
  staged Base64 call path. Its byte-conversion cost is explicit follow-up input
  for the broader typed `haxe.io.Bytes` migration under `haxe_go-vfp.8.7.11`,
  not a reason to restore compiler-owned crypto algorithms.

Validation evidence:

- `npm run test:changed`
- `npm test` (`272/272` snapshots)
- `npm run test:semantic-diff` (`133/133` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance` and `npm run test:stdlib-inventory`
- `npm run test:compiler-debt` (`4,401` `GoRaw` sites and `13` compiler shim
  entry points)
- direct `runtime/hxrt` Go tests and raw-injection hygiene
- `npm run test:perf:hxrt-selective`
- `npm run test:perf:stdlib-shims`
- `npm run test:perf:go` (five warning-only startup signals; no hard failures)
- `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)

Observed result:

- Crypto library behavior is source-owned, while native algorithms remain
  behind a narrow typed capability. No crypto API body, default, padding rule,
  or generated `Bytes` layout originates in `GoCompiler` or `hxrt`.
- Compiler debt falls permanently by 26 raw sites. Snapshot regeneration removes
  7,602 lines, chiefly because unrelated `stdlib_symbols` users no longer receive
  the old crypto block; the two `incident_api` profiles each lose another 117
  unused generated lines.
- Runtime slicing includes `crypto.go` only for programs that use the staged
  crypto APIs. XML/zip-only output stays on its existing migration path without
  importing crypto packages or carrying crypto declarations.

### 2026-07-16: move XML DOM, parser, and printer behavior to staged source (`haxe_go-vfp.8.7.15.2`)

Implementation:

- Added a red ownership contract and the `xml_source_owned` semantic-diff
  fixture before removing the compiler implementation. The fixture covers DOM
  construction and mutation, parent movement, attributes and iterators,
  compact/pretty printing, node validation, comments, CDATA, directives,
  doctypes, entities, strict/loose parsing, and structured parser positions.
- Tried the unchanged Haxe 4.3.7 sources first. `haxe.xml.Parser` now works
  unchanged and remains upstream-owned. Root `Xml` needs a narrow staged
  override because inline throwing accessors, concrete generic iterators,
  `Array.remove` / `Array.insert`, and empty indexed reads expose general Go
  lowering gaps. The override preserves the upstream API and algorithms with
  source-level equivalents while those gaps are tracked separately by
  `haxe_go-vfp.8.3.1`, `haxe_go-vfp.8.3.2`, and `haxe_go-vfp.8.3.3`.
- Added a narrow staged `haxe.xml.Printer` override. It retains the upstream
  printing algorithm but removes comment line breaks with ordinary
  `StringTools.replace` calls, so XML use does not accidentally select the
  still-migrating `EReg` and serializer compiler group.
- Corrected `StringTools.htmlEscape` so its omitted optional `quotes` argument
  follows Haxe null/false behavior instead of asserting a null Go interface as
  `bool`. This surfaced when ordinary upstream PCData printing began executing.
- Removed the complete XML DOM, parser, and printer declaration block, the
  `encoding/xml` import, and all XML classifier and registry selections from
  `GoCompiler`. No native XML parser or XML-specific `hxrt` capability was
  needed; parsing and printing are ordinary staged Haxe behavior.
- Regenerated source provenance, compatibility and module inventories, package
  counts, compiler-debt ceilings, snapshots, and committed example output. The
  governed staged surface now contains 129 sources, including 81 upstream
  overrides.

Validation evidence:

- `npm run test:changed`
- `npm test` (`272/272` snapshots)
- `npm run test:semantic-diff` (`134/134` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- `npm run test:compiler-debt` (`4,133` `GoRaw` sites and `13` compiler shim
  entry points)
- direct `runtime/hxrt` Go tests and raw-injection hygiene
- `npm run test:perf:hxrt-selective`
- `npm run test:perf:stdlib-shims`
- `npm run test:perf:go` (two warning-only startup signals; no hard failures)
- `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)
- `npm run test:release-contracts`

Observed result:

- XML semantics now originate in Haxe source. Strictness, entity decoding,
  detailed source positions, DOM ordering, mutation rules, and formatting no
  longer depend on Go's `encoding/xml` behavior or a compiler declaration blob.
- The 526-line compiler XML block is gone, and the `stdlib_symbols` allowance
  falls permanently from 738 to 470 raw sites. The repository-wide `GoRaw`
  ceiling falls by 268 sites, from 4,401 to 4,133.
- Snapshot regeneration removes 15,291 tracked lines because unrelated
  `stdlib_symbols` consumers stop carrying the old XML block. XML consumers now
  receive ordinary `module_xml.go`, `module_haxe_xml_parser.go`, and
  `module_haxe_xml_printer.go` files only when those sources are reachable.

### 2026-07-16: move `haxe.zip` compression to staged std and typed runtime capabilities (`haxe_go-vfp.8.7.15.3`)

Implementation:

- Added red ownership, semantic-diff, snapshot/runtime, and direct Go runtime
  contracts before changing the compiler. The contracts cover levels `-1`, `0`,
  `1`, `6`, and `9`; positive buffer-size hints; empty and binary payloads;
  invalid levels and streams; whole-buffer instance calls; and the raw-DEFLATE
  path used by `haxe.zip.Tools`.
- Tried the mainstream Haxe 4.3.7 surfaces first. Its generic `Compress` is an
  intentional `NotImplementedException` stub, while generic `Uncompress.run`
  uses the Haxe inflater but cannot provide the target instance/raw-DEFLATE
  capability needed by `Tools`. The Go target therefore needs two narrow staged
  overrides rather than compiler declarations.
- Added canonical staged `Compress` and `Uncompress` overrides. Haxe owns level
  validation, the 64 KiB default, positive buffer-size policy, `haxe.io.Bytes`
  conversion, whole-buffer result records, and negative-window raw-DEFLATE
  selection. The established one-shot target contract retains no native state,
  so `setFlushMode` and `close` are intentionally no-ops in this slice.
- Added the typed `std/hxrt/zip/NativeZip.hx` boundary over
  `runtime/hxrt/zip.go`. Only integer byte arrays, integers, and one raw-stream
  Boolean cross the package boundary. Go owns zlib/raw-DEFLATE execution and
  native errors; no generated `haxe.io.Bytes` layout, reflection, unsafe code,
  or raw source injection is involved.
- Removed the two compiler structs, both compiler-owned run functions,
  `bytes` / `compress/zlib` / `io` imports, classifier routes, intrinsic-registry
  ownership, and the special default-argument exception. Zip calls now select
  ordinary Haxe modules and a separate footprint-explicit `zip` runtime feature.
- Corrected the inventory ownership of the adjacent `haxe.zip` modules:
  `Compress` / `Uncompress` are staged-source plus runtime bindings, while
  `Entry`, `FlushMode`, `Huffman`, `InflateImpl`, `Reader`, `Tools`, and `Writer`
  remain unchanged upstream Haxe source.
- Progressive multi-call source/destination buffers, native codec handles, and
  stateful flush/close lifecycle semantics are explicitly deferred to
  `haxe_go-vfp.8.7.21`; this migration does not imply that broader streaming
  contract is complete.

Validation evidence:

- `npm run test:changed`
- `npm test` (`272/272` snapshots)
- `npm run test:semantic-diff` (`135/135` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- `npm run test:compiler-debt` (`4,118` `GoRaw` sites and `13` compiler shim
  entry points)
- direct `runtime/hxrt` Go tests and raw-injection hygiene
- `npm run test:perf:hxrt-selective`
- `npm run test:perf:stdlib-shims`
- `npm run test:perf:go` (four warning-only startup budget signals, zero
  enforced hard failures, and one explicitly non-enforced hard-gate dry-run
  candidate)
- `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)
- `npm run test:release-contracts`

Observed result:

- Public zip behavior now originates in Haxe source, and native compression is
  isolated behind a typed, representation-neutral runtime capability. Invalid
  streams still cross the ordinary Haxe exception carrier on Go.
- The `stdlib_symbols` raw allowance falls permanently from 470 to 455 sites,
  and repository-wide `GoRaw` debt falls by 15 sites, from 4,133 to 4,118.
- Snapshot regeneration removes 1,942 tracked lines from unrelated
  `stdlib_symbols` consumers. The actual zip consumer gains ordinary
  `module_haxe_zip_*` files and selectively copied `hxrt/zip.go`; unrelated
  runtime slices do not acquire compression code.
- Package governance covers 132 sources (83 upstream overrides and 39 typed
  bindings), with 314 manifest entries and 315 archive members.

### 2026-07-16: move root `Date` and `Math` to staged std ownership (`haxe_go-vfp.8.7.15.4`)

Implementation:

- Added complete canonical `Date` and `Math` overrides under `std/go/_std`.
  The mainstream Haxe 4.3.7 declarations are target-supplied extern contracts,
  so they cannot run unchanged; the Go overrides now provide those target
  implementations as ordinary documented Haxe source instead of compiler
  declarations.
- `Date` owns a portable epoch-millisecond field, constructors, all local and
  UTC accessors, timezone offset, parsing, formatting, and wall-clock policy.
  Typed `std/hxrt/date` bindings cross only numbers, strings, and a scalar
  `DateParts` carrier into footprint-explicit `runtime/hxrt/date.go`; generated
  `Date` layout and Go `time.Time` never cross that boundary.
- The Date runtime uses Go's millisecond time APIs rather than converting the
  entire timestamp through nanoseconds. A red year-2500 regression demonstrated
  that the previous nanosecond approach wrapped to 1915; the new direct and
  semantic contracts cover dates on both sides of the roughly 1678–2262
  `UnixNano` range while preserving fractional `Date.fromTime` values in Haxe.
- `Math` owns constants, Haxe ties-up rounding policy, finiteness and NaN rules,
  and the reference runtime's asymmetric operand-order behavior for equal
  signed zeros. Float operations bind directly through typed Go `math` and
  `math/rand` externs. Only `floor`, `ceil`, and `round` use
  `runtime/hxrt/math.go`, because Go's native functions return `float64` while
  the Haxe 4.3.7 API returns `Int` and the general `Std.int` migration is not
  complete.
- Removed Date and Math declarations, `math` / `time` imports, classifier
  routes, and intrinsic-registry ownership from the monolithic
  `stdlib_symbols` compiler group. Date and the three integer Math adapters now
  use separate footprint features, while programs that use neither no longer
  carry the old declarations.
- Kept the serializer's temporary Date representation bridge narrow: it probes
  only the staged millisecond field and no longer assumes a compiler-emitted
  `time.Time`. Removing the broader reflective serializer representation bridge
  remains owned by `haxe_go-vfp.8.7.13`.
- Updated source provenance, compatibility and module inventories, package
  counts, debt ceilings, examples, and fail-closed ownership/footprint tests.
  Governance now covers 139 sources: 85 upstream/staged overrides, 44 typed
  `hxrt` bindings, five public Go facades, and five staged support files.

Validation evidence:

- `npm run test:changed`
- `npm test` (`274/274` snapshots)
- `npm run test:semantic-diff` (`137/137` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- `npm run test:compiler-debt` (`4,095` `GoRaw` sites, `441` raw sites in the
  remaining `stdlib_symbols` group, and `13` compiler shim entry points)
- direct `runtime/hxrt` Go tests and raw-injection hygiene
- `npm run test:perf:hxrt-selective`
- `npm run test:perf:stdlib-shims`
- `npm run test:perf:go` (four warning-only startup budget signals and zero
  enforced hard failures)
- `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)
- `npm run test:release-contracts`

Observed result:

- Root Date and Math behavior now originates in Haxe source. Native time and
  integer-conversion work is isolated behind typed, representation-neutral
  capabilities, with no `Dynamic`, `Any`, unsafe conversion, raw injection, or
  legacy profile branch added.
- The `stdlib_symbols` raw allowance falls permanently from 455 to 441 sites,
  and repository-wide `GoRaw` debt falls by 23 sites, from 4,118 to 4,095.
- Snapshot regeneration removes 14,408 lines and adds 8,320, including two new
  focused ownership/footprint contracts. Actual Date
  and integer-rounding consumers receive ordinary source modules plus their
  narrow runtime files; unrelated `stdlib_symbols` consumers lose the old
  Date/Math block.
- The canonical package now contains 85 staged overrides, 323 manifest entries,
  and 324 archive members. This slice does not claim closure of the remaining
  Unicode, Reflect/Type, Std/Option, logging, or serializer migration work.

### 2026-07-16: move root `UnicodeString` algorithms to staged std ownership (`haxe_go-vfp.8.7.15.5`)

Implementation:

- Added red semantic-diff, snapshot/runtime, ownership, and direct runtime
  contracts before changing the compiler. The cases cover astral code points,
  negative and out-of-range positions, reversed and omitted ranges, negative
  substring lengths, overlapping and empty searches, both iterator forms,
  comparisons, mixed concatenation, compound assignment, valid one- through
  four-byte UTF-8, malformed continuation/overlong/truncated/surrogate/range
  sequences, and the required `RawNative` error.
- Tried the mainstream Haxe 4.3.7 source unchanged first. It cannot run
  unchanged on this target: the selected UTF-16 branch assumes native code-unit
  indexing even though pointer-backed Go strings are already rune-indexed, its
  constructor parameter shadows Go's built-in `string` type, and its
  declaration-only abstract operators lower without usable string types.
- Added canonical `std/go/_std/UnicodeString.hx`. Ordinary Haxe now owns every
  bounds, slicing, search, comparison, iterator, operator, and UTF-8 validation
  rule. Relational comparison walks typed code points in Haxe, and compound
  assignment intentionally reuses the typed `+` operators so the returned value
  is assigned instead of becoming a discarded Go call.
- Expanded the existing typed `GoStringRuntime` boundary with only rune length
  and already-normalized code-point slicing. The new Go helper performs the
  representation conversion and no Haxe range policy; existing typed
  `charCodeAt` remains the only lookup primitive.
- Removed all eight UnicodeString compiler declarations, both classifier
  routes, intrinsic-registry ownership, and the final static default-argument
  exception. Direct use now plans the staged module plus its two staged Unicode
  iterator classes on demand.
- Updated provenance, inventory, compatibility, package-layout, rationale, and
  fail-closed ownership records. Governance now covers 140 sources: 86
  upstream/staged overrides, 44 typed `hxrt` bindings, five public Go facades,
  and five staged support files.

Validation evidence:

- `npm run test:changed`
- `npm test` (`274/274` snapshots)
- `npm run test:semantic-diff` (`138/138` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- `npm run test:compiler-debt` (`3,879` `GoRaw` sites, `225` raw sites in the
  remaining `stdlib_symbols` group, and `13` compiler shim entry points)
- direct `runtime/hxrt` Go tests and raw-injection hygiene
- `npm run test:perf:hxrt-selective`
- `npm run test:perf:stdlib-shims`
- `npm run test:perf:go` (seven warning-only startup signals, zero enforced hard
  failures, and one explicitly non-enforced hard-gate dry-run candidate)
- `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)
- `npm run test:release-contracts`

Observed result:

- Root UnicodeString behavior now originates in documented Haxe source. No
  algorithm, default, bounds rule, validation state machine, `Dynamic`, `Any`,
  raw injection, unsafe conversion, or legacy profile branch was added to a
  compiler or runtime shim.
- The `stdlib_symbols` raw allowance falls permanently from 441 to 225 sites,
  and repository-wide `GoRaw` debt falls by 216 sites, from 4,095 to 3,879.
- Snapshot regeneration changes 254 generated files, removing 7,941 lines and
  adding 2,849. Only the direct UnicodeString and `haxe.Utf8` consumers gain
  ordinary Unicode source modules; 25 unrelated `main.go` files lose the old
  compiler block, while shared string-runtime snapshots gain the seven-line
  representation primitive.
- The canonical package now contains 86 staged overrides, 324 manifest entries,
  and 325 archive members. This slice does not claim closure of the remaining
  Reflect/Type, Std/Option, logging, or serializer migration work.

### 2026-07-16: make concrete iterators structurally assignable (`haxe_go-vfp.8.3.3`)

Implementation:

- Added red semantic, positive snapshot/runtime, and negative compile contracts
  before changing lowering. They cover direct `ArrayIterator<Int>` assignment,
  indexed array mutation after iterator creation, direct ordinary-call
  arguments, a user-defined generic iterator, inherited virtual dispatch
  through a base-typed value, `MapKeyValueIterator`, and a mismatched element
  type that Haxe must still reject statically.
- Extended the existing `GoLambdaIterableLowering` representation owner instead
  of adding a second iterator carrier. It recognizes only the closed anonymous
  `hasNext():Bool` / `next():T` shape, evaluates a concrete source once, and
  exposes typed closures through the existing structural map. Generated method
  calls use `__hx_this` so overrides remain virtual, and erased generic `next`
  results are asserted back to the Haxe-proven target element type.
- Added a pre-lowering path for direct and safely inlined array iterators. It
  captures the live typed Go slice and one cursor, so indexed mutations remain
  visible; an earlier red semantic run proved that copying elements to `[]any`
  would incorrectly freeze the original slots.
- Restored root `Xml` child iteration to the upstream
  `children.iterator()` / `ret.iterator()` expressions and removed its private
  hand-written structural closure. `Xml.iterator()` remains non-inline because
  preserving validation effects across a multi-prefix inline coercion is the
  separate `haxe_go-vfp.8.3.4` contract.
- Centralized the one syntax-only empty-map fragment needed until typed Go map
  composite literals land under `haxe_go-vfp.8.3`. All iterator keys, closures,
  selectors, calls, and type recovery remain typed AST nodes.

Validation evidence:

- `npm run test:changed`
- snapshot coverage for all `276` cases; the aggregate replay passed `275` and
  the sole `stdlib/int64_parity` host-load timeout passed immediately with
  runtime checking under a wider timeout. The earlier full replay had also
  passed that unchanged fixture.
- `npm run test:semantic-diff` (`139/139` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- `npm run test:compiler-debt` (`3,878` `GoRaw` sites and `13` compiler shim
  entry points)
- raw-injection hygiene, terminal-input behavior, and
  `npm run test:release-contracts`
- `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)

Observed result:

- Concrete generated iterator classes now satisfy Haxe's structural
  `Iterator<T>` contract at declarations, assignments, returns, branches, and
  ordinary call arguments without reflection, unsafe conversion, a runtime
  helper, exported method wrappers, or stdlib class-name dispatch.
- Array-backed iterators retain their live typed slice, user generic iterators
  recover erased results, and base-typed subclass iterators retain virtual
  behavior. Mismatched element types still fail in Haxe before Go generation.
- Root `Xml` no longer owns an iterator protocol shim, and repository-wide
  `GoRaw` debt falls permanently by one site, from 3,879 to 3,878. This slice
  intentionally does not claim effect-preserving restoration of the upstream
  inline modifier; that remains owned by `haxe_go-vfp.8.3.4`.

### 2026-07-16: preserve inline iterator effects and restore upstream `Xml.iterator()` (`haxe_go-vfp.8.3.4`)

Implementation:

- Added a red semantic-diff and snapshot/runtime contract before changing
  lowering. An inline method performs one observable effect, returns
  `Array.iterator()`, and is consumed through both a declaration and an ordinary
  call argument. The contract also mutates the source array after iterator
  creation to retain the live-slice guarantee from `haxe_go-vfp.8.3.3`.
- Changed the existing typed native-array adapter from an expression-only result
  to an ordered prefix plus expression. It recursively separates an inline
  block's final array iterator from preceding setup, lowers every retained setup
  expression through ordinary typed statement lowering, and folds only trailing
  aliases whose evaluation position is unchanged.
- Expected-type declaration, assignment, return, and branch contexts emit the
  prefix directly. Expression-only call arguments materialize the same prefix in
  an immediately invoked function, so source order and exactly-once effects are
  preserved without retaining the erased iterator constructor.
- Restored the upstream `inline` modifier on root `Xml.iterator()`. Its node-type
  validation now remains ordered before traversal, while inline `for` consumers
  such as `haxe.xml.Printer` lower to direct typed slice loops instead of
  structural map lookups.
- Kept the change inside `GoLambdaIterableLowering` plus expected-type
  orchestration. It adds no runtime helper, reflection, unsafe conversion, raw
  fragment, profile branch, XML special case, or library traversal algorithm.

Validation evidence:

- `npm run test:changed`
- `npm test` (`277/277` snapshots)
- `npm run test:semantic-diff` (`140/140` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- `npm run test:compiler-debt` (`3,878` `GoRaw` sites and `13` compiler shim
  entry points)
- raw-injection hygiene, `npm run test:release-contracts`, and
  `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)
- `npm run test:perf:go` (three warning-only startup signals and zero enforced
  hard failures) and `npm run test:perf:stdlib-shims`

Observed result:

- Inline iterator setup now runs once and in order in both statement and call
  argument contexts. Generated Go contains one typed live-array cursor and no
  `ArrayIterator` type, constructor, copied `[]any` storage, reflection, or
  unsafe path.
- Root `Xml.iterator()` again matches the upstream inline API. Generated XML
  printer loops validate the node once before walking `value.children` directly,
  eliminating the prior structural map method lookups without moving XML policy
  into the compiler.
- Compiler debt stays flat at 3,878 raw sites; this slice changes orchestration,
  not syntax debt. Structural coercion for constructor parameters remains the
  separate `haxe_go-vfp.8.3.5` path, while effectful inline blocks ending in a
  non-array concrete iterator remain `haxe_go-vfp.8.3.6`.

### 2026-07-16: coerce structural iterators in constructor arguments (`haxe_go-vfp.8.3.5`)

Implementation:

- Added red semantic-diff, positive snapshot/runtime, and negative compile
  contracts before changing constructor lowering. They cover direct
  `ArrayIterator<Int>` input with observable argument ordering and later array
  mutation, a user-defined generic iterator, inherited virtual dispatch through
  a base-typed value, and a mismatched element type that Haxe must reject.
- Routed explicit and default `TNew` arguments through the same expected-type
  coercion used by declarations, assignments, returns, and ordinary calls.
  Ordered setup is materialized as one expression for the Go constructor call,
  so each source argument still evaluates once and from left to right.
- Resolve the expected parameter from the constructor signature that generated
  Go actually emits. Generic Haxe constructors erase their structural
  `next():T` closure to `next():any`; adapting against the applied source type
  initially produced a more specific closure that looked valid to Haxe but
  failed at runtime inside the erased Go map. Reading the declared constructor
  type keeps the adapter and generated constructor on the same ABI.
- Reused `GoLambdaIterableLowering` without adding a constructor-specific
  carrier, runtime helper, reflection, unsafe conversion, raw fragment,
  profile branch, or stdlib class-name dispatch.

Validation evidence:

- `npm run test:changed`
- `npm test` (`279/279` snapshots)
- `npm run test:semantic-diff` (`141/141` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- `npm run test:compiler-debt` (`3,878` `GoRaw` sites and `13` compiler shim
  entry points)
- raw-injection hygiene, `npm run test:release-contracts`, and
  `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)

Observed result:

- Constructor parameters now accept matching direct array, user-generic, and
  inherited concrete iterators through the same structural `Iterator<T>` map
  used everywhere else. Array cursors remain live after indexed mutation,
  generic closures match the erased constructor ABI, and subclass overrides
  remain virtual.
- Side effects in earlier arguments, iterator setup, and later arguments retain
  source order and run exactly once. Mismatched element types still fail during
  Haxe compilation rather than escaping into generated Go.
- Compiler debt stays flat at 3,878 raw sites because this is an expected-type
  orchestration fix. Effectful inline blocks ending in a non-array concrete
  iterator remain explicitly deferred to `haxe_go-vfp.8.3.6`.

### 2026-07-16: retain concrete iterator authority through effectful inline blocks (`haxe_go-vfp.8.3.6`)

Implementation:

- Added red semantic-diff, positive snapshot/runtime, and negative compile
  contracts before changing lowering. The positive case records an effect before
  constructing a user generic iterator, then repeats the shape with a subclass
  stored in a base-typed local and passed directly to an ordinary call. Before
  the fix, generated Go tried to pass both class pointers where the existing
  structural `map[string]any` carrier was required.
- Extracted the concrete iterator validation from carrier construction into one
  shared typed plan. It still requires the exact anonymous
  `hasNext():Bool` / `next():T` target and a non-extern concrete class with
  matching zero-argument methods; ordinary direct values and recovered inline
  tails now make the same decision.
- Added a narrow nested-block plan that retains every setup expression in source
  order while preserving the final expression's concrete Haxe type. The setup
  lowers once, then the terminal value flows through the existing class-agnostic
  structural adapter, including erased generic result recovery and `__hx_this`
  virtual dispatch.
- Wired the plan into both expected-type statement contexts and source-aware
  ordinary call arguments. Each entry point tries the native Array cursor first,
  so Array keeps its distinct live-slice representation instead of falling into
  the general concrete-class path.
- Added no runtime helper, reflection, unsafe conversion, raw fragment, second
  carrier, profile branch, or stdlib class-name dispatch.

Validation evidence:

- `npm run test:changed`
- `npm test` (`281/281` snapshots)
- `npm run test:semantic-diff` (`142/142` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- `npm run test:compiler-debt` (`3,878` `GoRaw` sites and `13` compiler shim
  entry points)
- raw-injection hygiene, `npm run test:release-contracts`, and
  `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)

Observed result:

- Effectful inline methods can now return user concrete generic iterators to
  matching `Iterator<T>` declarations and ordinary call parameters. Effects run
  once before construction, erased generic results recover the Haxe-proven
  element type, and base-typed subclass tails retain override dispatch.
- Mismatched element types still fail in Haxe before Go generation. Existing
  direct concrete, Array live-cursor, constructor, and inline XML paths remain
  green through the aggregate regression suites.
- Compiler debt stays flat at 3,878 raw sites because this change preserves typed
  tail authority and reuses the existing carrier rather than adding syntax or a
  representation owner.

### 2026-07-16: preserve inline throw result types (`haxe_go-vfp.8.3.1`)

Implementation:

- Added a red semantic-diff, snapshot/runtime, and static ownership contract
  before changing lowering. They cover inline accessors returning `String`,
  `Int`, `Bool`, nullable references, and generic values inside comparisons,
  interpolation, and returns, including both successful reads and caught throws.
  A nested function literal declared before an outer continuation also pins that
  suppression cannot leak across generated function scopes.
  Before the fix, the generated fallback could inherit the surrounding
  expression's Go type instead of the accessor result type.
- Split throw handling into two typed cases. A terminal throw used as a value
  now receives its immediate expected storage type, while a guard throw with
  later generated code no longer emits a synthetic return that would escape the
  enclosing function or immediately invoked expression.
- Applied continuation-aware fallback suppression to ordinary statement blocks,
  expression blocks, and wrapped block-tail throws. Terminal function throws
  still receive the Go zero return required for static typing, and nullable
  values use their storage representation rather than their unwrapped value
  type. Each generated function saves and resets suppression state so an outer
  guard continuation cannot remove a nested closure's required return.
- Restored the exact upstream Haxe 4.3.7 inline modifiers on root `Xml`'s four
  throwing accessors. Parser and printer call sites now inline the same node-type
  validation and field access without an XML-specific compiler route.
- Refreshed the affected snapshots and committed example trees only after an
  aggregate audit. Across the existing snapshot suite, the substantive general
  change is removal of unreachable zero returns after guard throws; XML adds the
  expected inline validation/field access, and remaining differences are
  deterministic temporary renumbering.

Validation evidence:

- `npm run test:changed`
- `npm test` (`282/282` snapshots)
- `npm run test:semantic-diff` (`143/143` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- `npm run test:compiler-debt` (`3,878` `GoRaw` sites and `13` compiler shim
  entry points)
- raw-injection hygiene, `npm run test:release-contracts`, and
  `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)

Observed result:

- Inline accessor throws now retain their immediate `String`, `Int`, `Bool`,
  nullable-reference, or erased generic Go result contract even when the caller
  produces a different outer type. Successful and caught-throw behavior matches
  the Haxe reference in all five cases.
- Guard throws followed by a value tail emit no dead return; terminal throwing
  value branches retain correctly typed `bool` or `any` zero returns solely for
  Go's static checker. Root `Xml` accessors can therefore match upstream source
  without compiler-owned library policy.
- Compiler debt remains flat at 3,878 raw sites. This slice does not change the
  exception runtime representation or claim the remaining typed-Go-IR
  control-flow work tracked by the other `haxe_go-vfp.8.3` children.

### 2026-07-16: lower portable Array remove and insert structurally (`haxe_go-vfp.8.3.2`)

Implementation:

- Added red semantic-diff, selective snapshot/runtime, and static ownership
  contracts before changing lowering. They cover first-match and missing
  removal, duplicate values, string contents, object identity, nullable Int and
  String values, erased generic values, every insertion position rule,
  left-to-right argument evaluation, and single evaluation plus write-back for
  anonymous-object fields.
- Lowered `Array.remove` and `Array.insert` entirely through typed Go AST nodes.
  Removal ranges to the first Haxe-equal element, shifts with `copy`, clears the
  released slot, shrinks the slice, and returns `Bool`. Insertion evaluates both
  arguments once, applies Haxe negative and oversized position rules, grows with
  `append`, shifts the suffix, and writes the final slice back to its mutation
  site. No `GoRaw`, profile branch, library-name dispatch, reflection, or unsafe
  operation owns either algorithm.
- Made array-like storage preserve nullable primitive elements as `any`, matching
  existing value storage instead of emitting impossible Go literals such as
  `[]int{nil}`. Typed comparable elements still use native equality and strings
  use their value comparator; only interface-backed or non-comparable carriers
  use the selectively inferred `hxrt.HaxeEqual` capability.
- Moved the existing reference-identity reflection island out of the atomic
  object file into the new equality runtime slice. `AtomicObject` now depends on
  that shared slice, and Array lowering requests it only when an interface-backed
  comparison is emitted. The debt baseline moved the same 130 copied reflection
  findings to their new owner while total reflection stayed exactly flat at 597.
- Restored root `Xml.removeChild()` and `insertChild()` to their upstream
  `children.remove(...)` / `children.insert(...)` source. The staged override now
  retains only its documented nullable empty `firstChild()` adaptation.
- Correct string equality so native nil remains distinct from the literal
  `"null"`. That regression exposed a previously masked `Sys.getEnv` bridge bug;
  explicit `Null<String>` extern results now bypass non-nullable Go string
  normalization and preserve an absent environment value as nil.

Validation evidence:

- `npm run test:changed`
- `npm test` (`283/283` snapshots)
- `npm run test:semantic-diff` (`144/144` cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- `npm run test:compiler-debt` (`3,878` `GoRaw` sites, `597` reflection sites,
  and `13` compiler shim entry points)
- raw-injection hygiene, `npm run test:release-contracts`, and
  `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates)
- `npm run test:perf:hxrt-selective` and `npm run test:perf:go` (no hard
  regression; the Go lane reported warning-only startup signals)

Observed result:

- Ordinary staged Haxe source can use `Array.remove` and `Array.insert` without
  rebuilding arrays or growing compiler-owned stdlib behavior. Generated Go is
  typed slice code, preserves nullable and generic equality, evaluates mutation
  inputs once, and writes changed slice headers back through anonymous fields.
- Root XML child mutation now matches upstream Haxe source and retains parent
  ownership semantics. Missing environment values also remain native nil after
  the corrected string equality stopped masking the nullable extern boundary.
- This slice deliberately does not claim full portable Array identity. Concrete
  arrays crossing erased generic function boundaries remain
  `haxe_go-vfp.8.3.7`; shared identity across length-changing aliases is the
  xhigh representation decision in `haxe_go-vfp.8.3.8`; sparse indexed growth
  remains `haxe_go-vfp.8.7.20`.

### 2026-07-17: preserve portable Array identity and sparse growth (`haxe_go-vfp.8.3.7`, `haxe_go-vfp.8.3.8`, `haxe_go-vfp.8.7.20`)

Implementation:

- Added failing contracts first for aliases crossing locals, fields, parameters,
  returns, callbacks, erased `Dynamic`, generic functions, and
  `ReadOnlyArray`. The same contracts cover sparse indexed writes, nil holes,
  length-changing mutations, detached copies, compound and unary indexed
  assignments, and once-only receiver/index/value evaluation.
- Changed portable Haxe `Array<T>` from a copied Go slice header to one shared
  `*hxrt.Array` identity carrier. The carrier owns erased element storage and the
  representation primitives needed by compiler lowering: length, indexed read
  and write, sparse growth, push/pop, insertion/removal, copy, and
  an explicit values view for controlled runtime bridges. Portable and metal use
  the same representation decision; no behavior branches on the legacy profile.
- Added `go.NativeSlice<T>` as the explicit real-Go-slice boundary and migrated
  slice externs plus `goextern` generation to it. Array-to-slice and
  slice-to-Array conversions copy deliberately, so a native API cannot silently
  acquire or sever Haxe alias identity. Fixed-size/vector, rest-argument, and
  bytes carriers keep their existing specialized representations.
- Routed array literals, indexing, iteration, mutation, generic boundaries,
  structural iterators, Lambda, serializers, JSON, reflection, regex, templates,
  sockets, and staged stdlib helpers through the shared carrier. Full-suite
  testing also exposed and fixed a string-method receiver loaded from erased
  Array storage and added the staged SSL native-certificate binding needed by
  alternative-name access.
- Added a structural `GoMakeSlice(elementType, length, capacity)` typed-IR node
  and used typed range statements for slice boxing, copy-back, regex split,
  type-reflection, and serializer paths. This preserves useful capacity hints
  while lowering compiler raw-syntax debt from 3,878 to 3,864 sites.
- Separated portable semantic-oracle cases from Go-target-only integration
  cases. Target-only `go.*` behavior now lives in Go-native snapshot/runtime
  fixtures instead of adding interpreter fallbacks to production types. The
  repository rule and semantic-diff guide now explain what the harness compares,
  why the interpreter is the oracle, how each phase runs, and when a case belongs
  outside it.

Validation evidence:

- `npm run test:changed`
- `npm test` (`287/287` snapshots)
- `npm run test:semantic-diff` (`135/135` portable oracle cases)
- `npm run test:stdlib-sweep:go-test` (`55/55` strict modules)
- `npm run test:examples` (`12/12` example/profile lanes)
- `npm run test:stdlib:governance` (141 tracked staged sources),
  `npm run test:stdlib-inventory`, and `npm run compatibility:verify`
- `npm run test:compiler-debt` (`3,864` `GoRaw`, `215` Haxe `Dynamic`, `5`
  Haxe `Any`, `2` Go unsafe, `675` reflection, and `13` compiler shim sites)
- `npm run test:release-contracts`, raw-injection hygiene, semantic-harness and
  documentation contracts, `hxrt` unit tests, and `goextern` unit tests
- `npm run security:go-tooling` (all 28 race/checkptr/vet/staticcheck gates;
  Staticcheck cache redirected to writable temporary storage in the sandbox)
- `npm run test:perf:hxrt-selective` (12–15 fewer copied files and 27.68–32.42%
  smaller binaries across the four cases) and `npm run test:perf:go` (five
  warning-only signals, zero metal or portable-vs-metal hard failures)

Observed result:

- Every ordinary Haxe Array alias now observes indexed and length-changing
  mutations through one identity, including erased and generic crossings.
  Sparse writes grow to `index + 1` with nil holes, later pushes append after the
  grown length, and `copy()` remains detached as Haxe requires.
- Explicit native-slice interop remains Go-shaped without weakening portable
  semantics. Generated adapters make the copy boundary visible, and ordinary
  app/test/example code does not need raw injection or compiler-owned stdlib
  shims to use either side.
- JSON cycle detection now recognizes shared Array carriers alongside maps and
  slices; its bounded reflection island is documented and ratcheted in both
  checked-in runtime and committed generated examples.
- Go extern interface satisfaction from structural method sets (for example, a
  buffer accepted where a writer interface is expected) is intentionally not
  claimed here. That independent interop design remains tracked as
  `haxe_go-vfp.8.4.1`.

### 2026-07-17: retire compiler-owned sockets (`haxe_go-vfp.8.7.14`)

Implementation:

- Added failing ownership and runtime contracts before implementation. They
  required canonical `Host`, `Socket`, and `UdpSocket` source, ordinary staged
  Input/Output support, typed runtime bindings, no `net_socket` classifier or
  authority, explicit timeout/readiness state, TCP/UDP round trips, and safe
  concurrent cleanup.
- Replaced `GoNetSocketEmitter` and its complete-type authorities with
  `std/go/_std/sys/net/{Host,Socket,UdpSocket}.hx` plus
  `std/sys/net/_SocketIO.hx`. Public object identity, byte bounds and copying,
  Haxe EOF/blocked translation, address construction, and select mapping now
  live in ordinary Haxe source.
- Added `std/hxrt/net` opaque handles and concrete address, byte-progress,
  accept, datagram, and readiness carriers over footprint-explicit
  `runtime/hxrt/socket.go`. The handle synchronizes replace/close/deadline and
  read/write state; close is idempotent, interrupts blocked reads, and a closed
  handle cannot leave `waitForRead` spinning forever.
- Reworked `sys.ssl.Certificate`, `Key`, `Digest`, and `Socket` to cross typed
  certificate/key/SNI/socket handles rather than raw injection or `Dynamic`.
  TLS transport composition moved from `ssl.go` into footprint-explicit
  `socket_ssl.go`, so SSL leaf APIs do not acquire networking code.
- Preserved the inherited `sys.net.Socket` return signature for TLS `accept`
  while returning the embedded base view of a real `sys.ssl.Socket`. This keeps
  Go method sets valid and matches established Haxe target behavior: dynamic
  type checks still see the accepted connection as SSL.
- Added selective-runtime cases for socket-only and TLS-socket programs. The
  existing SSL-leaf case stays free of `socket.go` and `socket_ssl.go`, and
  unrelated full-copy output excludes both new capability files unless typed
  use or explicit feature selection requires them.
- The xhigh second pass found two concurrency/portability defects before
  landing. Lazy UDP initialization could install competing ephemeral sockets,
  and the broadcast option used a POSIX descriptor type on Windows. A direct
  concurrency regression now requires one shared connection, while
  build-tagged POSIX/Windows option helpers and a permanent cross-build gate
  preserve both runtime behavior and platform compilation.
- Updated the intrinsic registry, provenance ledger, compiler-debt policy,
  portable inventory, compatibility source, feature matrix, and runtime docs.
  The broader network release blocker remains open: retiring compiler
  ownership does not claim cross-platform cancellation or hostile-peer closure.

Validation evidence:

- `npm run test:changed` and `npm test`
- `npm run test:semantic-diff` and `npm run test:stdlib-sweep:go-test`
- `npm run test:examples`
- `npm run test:stdlib:governance`, `npm run test:stdlib-inventory`, and
  `npm run compatibility:verify`
- targeted TLS and SNI snapshots with runtime execution, plus UDP and socket
  stream snapshots
- direct `runtime/hxrt` tests normally and under `go test -race`
- `python3 test/test_socket_runtime_cross_build.py` (Linux and Windows)
- `npm run test:compiler-debt`, raw-injection hygiene, release contracts, and
  the selective-runtime performance harness

Observed result:

- Portable Haxe networking remains the product semantics in both compatibility
  presets; no behavior branches on `portable|metal`. OS networking and TLS are
  explicit typed capabilities beneath source-owned APIs, not a second semantic
  product or a compiler-generated stdlib.
- TCP, UDP, deadlines, readiness, address translation, shutdown, broadcast,
  TLS handshake, peer certificates, SNI selection, accepted SSL identity, and
  concurrent cleanup have deterministic local evidence. Cross-platform network
  admission remains governed separately by `haxe_go-vfp.10.4`.

### 2026-07-18: retire the mixed regex/serializer compiler emitter (`haxe_go-vfp.8.7.13`)

Implementation:

- Added a failing ownership contract first for canonical staged `EReg`,
  `haxe.Serializer`, and `haxe.Unserializer`, typed runtime bindings, exact
  provenance, feature slicing, and complete removal of the behavior-heavy
  `regex_serializer` group and emitter.
- Moved match state, capture validation, `matchSub`, split/map traversal, Haxe
  replacement-template expansion, global replacement policy, every serialization
  token, cache ordering, recursive traversal, collection handling, resolver
  policy, and custom-hook sequencing into ordinary Haxe source under
  `std/go/_std`.
- Added typed `std/hxrt/regex` bindings over footprint-explicit
  `runtime/hxrt/regex.go`. The runtime owns only compiled RE2 resources,
  matching/quoting, and exact conversion from Go UTF-8 byte indexes
  to Haxe code-point offsets.
- Added typed `std/hxrt/serialization` bindings over footprint-explicit
  `runtime/hxrt/serialization.go`. The runtime owns deterministic erased field
  snapshots, decoded-field assignment, constructor-bypassed hidden-self repair,
  and bounded float parsing. One centralized, ratcheted
  `reflect.NewAt`/`unsafe.Pointer` lift reaches package-private generated fields;
  direct tests and checkptr guard it.
- Reused the existing approved Type metadata emitter for reachable class/enum
  names, resolution, empty construction, enum construction, and field lists.
  No serializer-specific metadata table or duplicate construction registry was
  added.
- Retained one exact compiler representation primitive:
  `lowerSerializationSourceBridgeShimDecls` emits same-package interface
  assertions for `hxSerialize`, `hxUnserialize`, `resolveClass`, and
  `resolveEnum`. Its machine-readable intrinsic decision forbids tokens,
  traversal, caches, regex, reflection, unsafe access, metadata tables, and
  constructors, and requires removal when generated method visibility gains a
  source-visible typed representation.
- Added regex-only and serialization-only selective-runtime snapshots. Each
  proves its own runtime file is copied and the other capability remains absent;
  neither behavior branches on the legacy `portable|metal` compatibility
  preset.
- Updated the intrinsic registry, provenance ledger, compiler-debt policy,
  portable inventory, compatibility artifacts, feature matrix, runtime docs,
  and ownership rationale. The old 2,491-line emitter and its complete-type
  authority are deleted.

Validation evidence:

- all 13 focused regex/serialization semantic-diff contracts
- the two selective-runtime snapshots and direct `runtime/hxrt` unit tests
- stdlib governance, inventory, compatibility, compiler-debt, and raw-injection
  contracts
- full snapshots, semantic diff, strict upstream stdlib Go sweep, and examples
- Go vet/staticcheck/race/checkptr security tooling
- Go profile, selective-runtime, and staged-stdlib performance gates
- an explicit written xhigh second pass that challenged host regex replacement
  and traversal semantics before accepting the boundary

Observed result:

- Regex and serialization are portable-by-default staged library surfaces in
  both compatibility presets. Native RE2 and erased field access are explicit
  typed runtime boundaries, not profile-wide native semantics.
- The compiler no longer owns public regex or serialization behavior. The only
  residual primitive adapts exact package-private method representation, while
  the pre-existing Type metadata authority remains the single source of
  reachable class/enum facts.
- Compiler raw-emission debt falls by more than 1,700 sites despite the exact
  bridge, and the new Dynamic findings are confined to the public serialization
  API and its erased boundaries. The reviewed unsafe ceiling is one import and
  one selector in a single runtime file.
- The second pass found and fixed two host-policy leaks before closure: Go's
  named-capture parsing of `$1x`, and its omission of empty matches adjacent to
  a previous match. Staged `EReg` now owns replacement expansion and the distinct
  zero-width progress rules for replace, map, and split; `hxrt` exposes no bulk
  traversal or replacement API.
- Follow-up `haxe_go-vfp.10.5.1` owns evaluation of typed generated field/method
  accessors that could remove the remaining unsafe lift and exact same-package
  bridge without returning serialization algorithms to compiler ownership.

### 2026-07-18: move the base `haxe.io` hierarchy to staged source (`haxe_go-vfp.8.7.11`)

Implementation:

- Added a fail-closed ownership contract before implementation. It requires
  canonical staged definitions for all eleven formerly generated public IO
  types, exact provenance/planner routing, and complete removal of the compiler
  `io` group, authorities, synthetic subclass wrappers, helper island, registry
  entry, and debt allowances.
- Moved `BufferInput`, `Bytes`, `BytesBuffer`, `BytesInput`, `BytesOutput`,
  `Encoding`, `Eof`, `Error`, `Input`, `Output`, and `StringInput` to canonical
  overrides under `std/go/_std/haxe/io`. Public validation, byte algorithms,
  stream loops, EOF/error behavior, endian policy, RawNative selection, aliases,
  and cache invalidation are now ordinary Haxe source.
- Replaced `GoIoHelpers` and compiler-generated IO forwarding methods with
  ordinary source inheritance and the existing `__hx_this` virtual-dispatch
  path. Typed source-backed std superclasses are now queued even when manual DCE
  initially selected only a user subclass, preserving normal base-class upcasts.
- Added a generic generated `String() string` adapter for source classes that
  implement typed `toString():String`. It delegates through `__hx_this`, allowing
  erased `Std.string` calls to observe source policy without an Eof-specific
  compiler rule.
- Added typed `std/hxrt/io` bindings for an opaque immutable `ByteView`, native
  allocation/conversion/copy/UTF capabilities, and scalar IEEE-754 bit
  reinterpretation. `runtime/hxrt/bytes.go` owns only those target facts; no
  generated `Bytes` layout or public IO policy crosses the package boundary.
- Staged `Bytes` retains `BytesData` alias semantics. Mutations invalidate the
  opaque cache, and a view requested after `getData()` validates against live
  integer values so external alias mutation cannot return stale native bytes.
- Routed staged Base64 and digest implementations through the same opaque view.
  This removes the previous `Bytes -> Array<Int> -> []int -> []byte` copy chain
  while leaving alphabets, padding, public construction, and API policy in Haxe.
- Retired the broad IO intrinsic-registry entry and every IO-specific compiler
  debt allowance. Neither `portable` nor `metal` selects different IO semantics;
  `metal` remains only a compatibility policy preset.

Validation evidence:

- red-to-green source/ledger/registry, exact runtime-surface, canonical-package,
  and crypto byte-view ownership contracts
- 294 generated-output snapshots and 135 portable semantic-diff cases
- 55 strict upstream stdlib modules compiled and checked with `go test`
- all 12 runnable example/profile lanes, including both portable and metal
  where the example declares both compatibility presets
- stdlib governance over 180 tracked sources, inventory, compatibility,
  release/archive, raw-injection, and compiler-debt contracts
- all 28 Go vet/staticcheck/race/checkptr gates across `hxrt` and representative
  generated portable/metal scopes
- selective-runtime, staged-stdlib boundary, and general Go profile performance
  harnesses
- an explicit written xhigh second pass covering byte aliasing, cache validity,
  endian/IEEE word ordering, inheritance/DCE, compiler-owned HTTP interaction,
  package selection, runtime slicing, and compatibility-preset invariance

Observed result:

- Public `haxe.io` policy is source-owned in both compatibility presets. The
  compiler retains no IO shim group or profile branch, and compiler debt falls
  from `go_raw=1598` / `compiler_shim=12` to `go_raw=1087` /
  `compiler_shim=11`.
- The opaque byte view removes the former double conversion for codecs and
  digests. The staged Base64 boundary measures 78.23 ns/op, 112 B/op, and 3
  allocations/op versus direct Go at 62.90 ns/op, 96 B/op, and 2 allocations/op;
  selective runtime output saves 12-15 files and 65-87% of runtime source.
- The xhigh pass found and removed five unreachable legacy runtime helpers, so
  string, hex, and buffer-length policy no longer ships in `hxrt`. It also made
  the isolated canonical source/package fixture carry the exact staged IO
  closure rather than silently selecting mainstream target implementations.
- Type-only staged declarations and staged superclasses are queued from their
  already-typed authority. At this IO closeout, the compiler-owned `sys.Http`
  carrier kept its concrete layout while its hidden `BytesBuffer` dependency
  became explicit; the carrier was subsequently retired by
  `haxe_go-vfp.8.7.12` without restoring an IO-specific compiler hierarchy.
- The general performance lane reports warning-only timing signals and no
  enforced hard failure. Those signals remain governed by the existing
  performance policy rather than being reclassified as IO closure evidence.

### 2026-07-18: move `sys.Http` request policy to staged source (`haxe_go-vfp.8.7.12`)

Implementation:

- Added a failing ownership contract first. It requires canonical staged
  `sys.Http`, typed opaque request/response bindings, selective `hxrt/http.go`
  packaging, and complete removal of the compiler HTTP group,
  `GoHttpHelpers`, registry authority, and generic debt allowances.
- Moved Haxe-visible request selection, parameters, headers, data URLs,
  multipart marker compatibility, proxy/custom-request choreography, response
  normalization, callback order, and status/error policy to
  `std/go/_std/sys/Http.hx`.
- Added typed `hxrt.http.HttpRequestHandle` and `HttpResponseHandle` boundaries.
  The native runtime receives only strings, scalars, an opaque `ByteView`, and
  an optional typed `SocketHandle`; no generated `sys.Http`, `haxe.io.Bytes`,
  map, callback, or stream layout crosses the package boundary.
- Added a one-use Go HTTP transport that owns URL parsing, native form/query
  encoding, proxy setup, response-body closure, idle-connection cleanup,
  timeout enforcement, deterministic header iteration, and custom socket
  consumption. It returns status and headers even when body reading fails so
  staged callback ordering remains observable.
- Added footprint-explicit runtime slicing for `http.go` with declared string,
  byte, and socket dependencies. HTTP code is absent unless typed `sys.Http` or
  `hxrt.http` usage requires it.
- Generalized instance dynamic-method lowering into per-instance function
  fields initialized by constructors. This is ordinary AST lowering rather
  than an HTTP special case and allows `haxe.http.HttpBase` callbacks to be
  replaced per request without incorrectly treating them as interface methods.
- Retired the compiler-emitted HTTP declaration block and its source helper.
  Neither `portable` nor `metal` selects different HTTP semantics; the legacy
  selector remains a convenience policy preset rather than a semantic branch.
- Preserved the established deterministic upload size-marker contract. Real
  streaming of the caller-supplied `Input`, including partial I/O,
  cancellation, and cleanup, remains explicitly owned by
  `haxe_go-vfp.10.4` and is not claimed as part of this migration.

Validation evidence:

- red-to-green migration-ledger and ownership assertions for source authority,
  typed runtime boundaries, selective packaging, and retired compiler debt
- focused portable semantic-diff contracts for request callbacks and
  proxy/custom-request behavior
- six focused generated-output snapshots covering selective runtime inference,
  dynamic-method RTTI shape, callbacks, proxy/socket handling, custom requests,
  and source ownership
- deterministic local Go transport tests for query/form/header/body behavior,
  multi-value response headers, truncated-body status preservation, bounded
  timeout, idle-connection cleanup, custom method/body, proxy formatting, and
  typed socket closure
- 295 generated-output snapshots and 135 portable semantic-diff cases
- 55 strict upstream stdlib modules compiled and checked with `go test`,
  including direct `haxe.Http` and `sys.Http`
- all 12 runnable example/profile lanes with unchanged portable/metal behavior
- stdlib governance over 183 tracked sources, exact 382-member package and
  381-entry manifest ratchets, inventory, compatibility, raw-injection, and
  release contracts
- all 28 race/checkptr/vet/staticcheck gates on supported Go 1.25.12, plus a
  clean reachable-vulnerability audit on that exact patched toolchain; the
  stale host Go 1.25.6 run failed closed as policy requires
- selective-runtime and general Go profile performance harnesses; the latter
  reported only its documented warning-only timing signals
- supply-chain verification for vendor provenance, lockfile, pinned actions,
  and dependency-update coverage

Observed result:

- `sys.Http` is an ordinary staged class in both compatibility presets. The
  compiler no longer owns HTTP library semantics or emits a synthetic public
  carrier.
- Native networking is isolated behind small typed capabilities, while the
  public Haxe API and callback contract remain readable and testable in source.
- The general dynamic-method representation now matches Haxe's assignable
  per-instance callback semantics and is regression-covered outside the HTTP
  fixture itself.
- Removing the 265 raw HTTP emitter statements and its shim group lowers the
  compiler-debt ratchet from `go_raw=1087` / `compiler_shim=11` to
  `go_raw=822` / `compiler_shim=10` without adding a `Dynamic` or `Any`
  transport boundary.
- The narrow `@:allow(sys.Http)` access to the existing typed socket handle
  avoids adding an HTTP-only method to the reflected `sys.net.Socket` surface.

### 2026-07-18: support concrete generated Template iterables with selective method metadata (`haxe_go-vfp.8.7.19`)

Implementation:

- Added a failing portable contract first. Haxe Eval rendered concrete custom
  `Iterable` / `Iterator` classes correctly, while generated Go could not find
  their lowercase `iterator`, `hasNext`, and `next` methods through reflection.
- Added a small closed-world metadata plan after the reachable class queue is
  complete. It records each concrete class, its canonical receiver, its direct
  generated superclass, and the exact Haxe lookup key and Go selector already
  chosen by ordinary method lowering.
- Added a dedicated typed-Go-AST emitter. One central switch recovers
  `__hx_this` from physical superclass carriers, a second selects the exact
  concrete resolver, and each per-class resolver lists only its own emitted
  methods before one nil-guarded direct-superclass fallback.
- Kept generated methods lowercase and returned already-bound function values.
  The implementation adds no exported duplicates, provider interface, global
  registry/map, `unsafe`, raw Go block, runtime discovery, or Template-specific
  compiler helper.
- Inserted the generic lookup after existing native/data-field discovery and
  before Go's exported `MethodByName` fallback in `Reflect.field` and
  `Reflect.hasField`. Metadata is absent unless either API is reachable.
- Changed staged `haxe.Template` to resolve `hasNext` and `next` through the
  existing Reflect contract and invoke them through `NativeTemplate.call`. The
  invalid `Iterator<Dynamic>` representation cast is gone; all iteration,
  fallback, macro, error, and rendering policy remains ordinary Haxe source.
- Kept `NativeTemplate` at exactly three typed runtime operations. A direct Go
  test proves that invoking an already-bound method preserves its receiver and
  mutation; `hxrt` neither discovers nor indexes generated methods.
- Documented the architectural boundary in `docs/typed-go-ir.md`: this feature
  uses a validated semantic metadata plan feeding typed Go AST. It does not
  justify copying `haxe.c`'s full ownership/lifetime/control-flow IR into the Go
  target.

Validation evidence:

- `npm test`: 297/297 snapshots, including exact selectors, inheritance
  fallback, nil guards, Reflect ordering, selective absence, and source-owned
  Template iteration.
- `npm run test:semantic-diff`: 137/137 portable parity contracts, including
  concrete Template iteration and computed generated-method lookup.
- `npm run test:stdlib-sweep:go-test`: 55/55 upstream stdlib modules; and
  `npm run test:examples`: 12/12 runnable examples.
- Stdlib governance, portable inventory, compatibility verification, raw
  injection hygiene, and the full 28-lane Go security/tooling matrix passed.
- The compiler-debt ratchet passed at `go_raw=822`, `haxe_dynamic=264`,
  `haxe_any=5`, `go_unsafe=4`, `go_reflection=695`, and
  `compiler_shim=11`.
- Selective-runtime performance passed; the full Go performance report had
  nine documented warning-only signals and zero hard failures.

Observed result:

- Concrete generated iterables now follow the same portable Template contract as
  arrays and structural iterators without changing the public generated Go API.
- The capability is generic to dynamic generated-method lookup rather than owned
  by Template, and it is emitted only when reachable Reflect use requires it.
- The semantic decision is a narrow immutable plan, while all target output is
  ordinary typed Go AST. A full second program IR remains intentionally deferred
  until repeated cross-cutting representation, effect, failure, or dispatch
  evidence demonstrates that the target AST plus feature plans is insufficient.
- A broad `Dynamic` Reflect consumer still requires conservative metadata for
  every reachable generated method; this adds 645 lines to each committed
  `incident_api` profile output. Follow-up `haxe_go-k8w2` owns deterministic
  footprint reporting and proof-based demand narrowing. It must retain the
  generic path whenever typed provenance cannot exclude generated class values.

### 2026-07-18: retire the residual `stdlib_symbols` compatibility group (`haxe_go-vfp.8.7.15.7`)

Implementation:

- Started with a failing governance contract covering the compiler dispatcher,
  classifier, source-owned planner, intrinsic registry, debt policy, provenance
  ledger, and generated `Option` ownership.
- Deleted the empty compiler-emitted `Std` carrier and the hand-written
  `haxe.ds.Option` tag/parameter carrier. Reachable upstream `Option` now enters
  the ordinary source-owned enum queue and is emitted in
  `module_haxe_ds_option.go` by the same typed enum pipeline as other Haxe
  enums.
- Removed false selectors for `Std`, `haxe.ds.BalancedTree`, `haxe.io.Path`,
  `haxe.Template`, and the staged SSL family. These selectors existed only to
  pull in the compatibility carrier and did not own behavior for those
  modules.
- Promoted the already-approved serialization invocation adapter to the
  independently named `serialization_source_bridge` group. It is now selected
  only when staged `haxe.Serializer` or `haxe.Unserializer` is reachable,
  instead of being emitted unconditionally or hidden under an unrelated group.
- Kept `Std.parseInt` and `haxe.Log.trace` honest as explicit migration-required
  direct calls. `haxe_go-vfp.8.7.22` owns complete source-level `Std` and `Log`
  semantics, including parsing edges, `PosInfos` formatting, and dynamic trace
  rebinding; this slice does not claim that the remaining calls are approved
  representation intrinsics.
- Removed the retired group from the intrinsic inventory, compiler-debt policy,
  debt capability map, and bidirectional stdlib provenance audit. Historical
  migration-log references remain historical evidence rather than current
  selectors.

Design review:

- Rejected an empty compatibility group because it would preserve a second
  ownership bucket with no semantic authority.
- Rejected a new special `Option` carrier because ordinary enum lowering already
  supplies the required typed shape and keeps source provenance visible.
- Rejected a partial staged `Std`: publishing an incomplete public API would
  move the declaration without completing Haxe 4.3.7 behavior. The focused
  `Std`/`Log` follow-up is the honest boundary.
- Rejected attaching serialization to Type or Reflect. The bridge invokes
  package-private generated methods; it neither owns metadata nor public
  reflection policy.
- The final second pass corrected the registry owner to the actual staged extern
  `haxe.GoSerializationBridge` and its five exact members. Serializer and
  Unserializer remain reachability selectors; the registry does not falsely
  claim that `hxSerialize` or `hxUnserialize` are members of those public types.
- The local second pass converged on one design, so no Oracle escalation was
  needed. This is a selector/ownership cleanup within the existing AST-first
  pipeline and provides no evidence for a universal program IR.

Validation evidence:

- the new red-to-green residual-group contract and all 27 stdlib migration
  ledger contracts
- all six compiler stdlib intrinsic-registry contracts
- all 297 generated-output snapshots, with 11 expected `main.go` reductions
  and one new ordinary `module_haxe_ds_option.go` artifact
- all 138 portable semantic-diff contracts, including `option_date_path` and the
  serialization custom-hook/resolver family
- all 55 strict upstream stdlib modules compiled and checked with `go test`,
  including direct `Std`, `haxe.ds.Option`, Template dependencies, and SSL-adjacent
  modules
- all 12 runnable portable/metal example lanes
- compiler-debt ratchet at `go_raw=554`, `haxe_dynamic=320`, `haxe_any=5`,
  `go_unsafe=4`, `go_reflection=736`, and `compiler_shim=10`

Observed result:

- `stdlib_symbols` is no longer a production selector, registry group,
  provenance dependency, or debt allowance.
- SSL, Template, Path, and collection programs no longer receive an unrelated
  empty `Std` plus `Option` carrier.
- `Option` construction and matching retain portable behavior while gaining an
  ordinary source-owned generated module.
- Serialization retains its exact same-package capability without becoming a
  second semantic product or profile-specific branch.

### 2026-07-18: finish `Std` and `haxe.Log` source ownership (`haxe_go-vfp.8.7.22`)

Implementation:

- Added failing contracts first for the complete Haxe 4.3.7 `Std` surface and
  for `haxe.Log.formatOutput` / mutable `trace`. The red evidence covered null,
  whitespace, sign, hexadecimal, prefix stopping, Int32 overflow, float
  exponents, aliases, downcasts, random bounds, compiler-injected `PosInfos`,
  custom parameters, direct function values, rebinding, and null rebinding.
- Added canonical staged `Std` and `haxe.Log` overrides. Integer and float-token
  scanning, overflow rules, aliases, downcast decisions, random-bound policy,
  trace formatting, and trace delegation through `Sys.println` are now ordinary
  Haxe source in both compatibility presets.
- Retained exactly two `Std` representation intrinsics: `Std.string` for erased
  target values and `Std.isOfType` for typed runtime tokens. Removed the direct
  `Std.parseInt -> hxrt.StdParseInt` and `haxe.Log.trace -> hxrt.Println`
  rewrites and removed their migration registry entries.
- Kept native work behind narrow typed capabilities: exact IEEE-754 token
  conversion, finite float-to-int conversion, and Go random generation. The Go
  runtime no longer owns integer parsing policy.
- Lowered static Haxe `dynamic function` methods as mutable typed Go variables.
  Calls through generated dynamic-method fields now evaluate the callee and
  arguments once, reject null through `hxrt.Throw`, and remain catchable by Haxe
  `try/catch`; typed native extern calls retain their direct path.
- Added closed, typed semantic plans for forwarded `Class<T>` / `Enum<T>` tokens
  and concrete subclass recovery through generated virtual receivers. These
  plans feed the existing typed Go AST and use no raw-Go emitter, reflection
  registry, `unsafe`, `Dynamic` transport boundary, or profile-name branch.
- Updated provenance, portability inventory, compatibility, feature mapping,
  package/archive ratchets, compiler-debt ownership, and committed snapshot and
  example trees for the new source authority.

Design review:

- Rejected a partial `Std` override because replacing the target declaration
  without its complete public API would create a misleading ownership claim.
- Rejected keeping parsing or trace formatting in `hxrt`: both are portable
  source policy, while only exact host representation conversions justify typed
  native bindings.
- Rejected treating `metal` as a separate semantic implementation. Portable
  Haxe behavior is the product default; the public selector remains only a
  compatibility policy preset, and explicit native modules/APIs remain the Go
  boundary.
- This work does not justify a universal compiler IR. The new decisions are
  small immutable semantic plans feeding the established builder/lowering,
  transform, and typed Go AST pipeline. A second whole-program IR would add
  ownership and synchronization cost without solving an observed repeated
  cross-cutting limitation.
- The `thinking:high` local design pass converged on one bounded typed approach.
  No unresolved competing design remained, so an Oracle review was not needed;
  Oracle remains appropriate if a future `thinking:xhigh` scope or repeated
  representation/effect/failure evidence reopens the architecture decision.

Validation evidence:

- 298/298 generated-output snapshots, including the focused source-owned
  `haxe.Log` runtime contract
- 139/139 portable semantic-diff contracts, including the new complete `Std`
  API case and existing type/reflection/serialization contracts
- 55/55 strict upstream stdlib modules compiled and checked with `go test`
- 12/12 runnable example/profile lanes after intentional generated-tree refresh
- all 28 Go race, checkptr, vet, and staticcheck lanes
- compiler-debt, stdlib governance, portability inventory, compatibility,
  canonical-layout, package, and intrinsic-registry contracts
- selective-runtime performance passed; the full Go profile harness reported
  only its documented warning-only signals and no enforced hard failure

Observed result:

- `Std` and `haxe.Log` are readable, testable staged library owners rather than
  special compiler call sites, with identical portable semantics under the
  compatibility presets.
- Mutable trace sinks now preserve formatting, position/custom-parameter data,
  direct function use, rebinding, restoration, and catchable null-call failure.
- Complete `Std` parsing and downcast behavior is covered without broadening the
  target runtime or introducing a universal IR.

### 2026-07-19: complete progressive `haxe.zip` codec semantics (`haxe_go-vfp.8.7.21`)

Implementation:

- Added the portable contract first and confirmed the previous whole-buffer
  implementation failed when a complete compressed stream did not fit in a
  five-byte destination. The cross-target case now covers repeated source and
  destination fragments, aggregate `read` / `write` accounting, completion,
  zlib round trips, and negative-window raw DEFLATE.
- Replaced per-call compression with an opaque typed deflate handle retaining a
  live Go zlib writer and pending output. Haxe continues to own source/destination
  positions, `Bytes` conversion and blits, and the public anonymous result.
- Added an opaque typed inflate handle whose fragment feeder implements Go's
  exact byte-reader path. A live inflater pauses at temporary input boundaries,
  resumes on the next `execute`, and stops without consuming bytes after the end
  of one stream. Partial input is not treated as terminal EOF and earlier input
  is not replayed.
- Applied destination-sized backpressure to the live inflater, so highly
  compressed input cannot expand into an unbounded hidden output buffer. Its
  pending decode state stays within the current or still-active destination
  allowance.
- Added a typed `ZipCodecStep` carrier containing only integer byte values,
  consumed input, and completion. No generated `haxe.io.Bytes` layout,
  `Dynamic` native state, raw injection, reflection, or `unsafe` crosses the
  staged-source/runtime boundary.
- Implemented exact `NO`, `SYNC`, and `FINISH` behavior. Go's public compressor
  does not expose zlib's FULL dictionary reset or BLOCK boundary stop, so both
  modes fail at `setFlushMode` with explicit target-capability errors instead of
  silently behaving like `SYNC`.
- Made close idempotent, use-after-close deterministic, invalid positions use
  `haxe.io.Error.OutsideBounds`, zero-capacity destinations return zero progress,
  and incomplete native inflater shutdown releases its paused decoder.
- Preserved source-owned `Compress.run`, `Uncompress.run`, and `haxe.zip.Tools`
  behavior, including the raw-DEFLATE ZIP entry path and footprint-explicit
  selection of `runtime/hxrt/zip.go`.

Design review:

- A first bounded prototype accumulated and replayed compressed input. It was
  semantically correct but rejected because many tiny fragments would cause
  quadratic decode work and did not retain the actual native inflater state.
- A terminal EOF per fragment was rejected because Go's inflater makes read
  errors sticky and cannot resume afterward. The final feeder blocks only the
  codec's private reader between fragments or output allowances and
  acknowledges consumption to the public call, retaining linear decode work,
  bounded pending expansion, and exact trailing-byte accounting.
- Adding cgo zlib or a third-party compression backend was rejected: both would
  add distribution/dependency cost solely to expose FULL/BLOCK controls that Go
  otherwise does not promise. Explicit limitation is the honest portable-target
  policy.
- This is a library/runtime state machine feeding the existing typed Go AST; it
  presents no evidence for a second whole-program IR or a profile-specific
  semantic product. The `thinking:high` local pass converged on one design, so
  no Oracle escalation was needed.

Validation evidence:

- `zip_streaming_contract` and the existing zip semantic-diff contracts
- `stdlib/zip_streaming_policy` plus the existing crypto/XML/zip snapshot runtime
- direct `runtime/hxrt` zip tests, including tiny fragments, bounded output,
  raw DEFLATE, trailing-byte preservation, unsupported flushes, idempotent close,
  and use-after-close, under the Go race detector
- strict upstream stdlib, examples, selective-runtime, compiler-debt, security,
  package-governance, inventory, and compatibility gates

Observed result:

- Instance codecs now preserve progressive Haxe behavior across repeated
  partial buffers while one-shot helpers remain simple source-owned APIs.
- The runtime boundary remains typed and representation-neutral, with actual
  persistent codecs rather than a compiler shim, whole-buffer restart, or
  universal IR layer.
- Unsupported zlib controls are visible and documented; no compatibility mode
  silently changes their semantics.
