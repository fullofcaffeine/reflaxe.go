# Portable Module Mapping Contract (Tier1 Seed)

This document defines ownership mapping for Tier1 portable modules:

- Haxe-source implementation
- runtime binding (`hxrt`)
- compiler intrinsic/shim
- mixed ownership (explicitly split)

It is the canonical Tier1 mapping seed for family extraction work.

Contract inputs:

- `test/portable_allowlist.json` (Tier1 module set)
- `test/portable_conformance_tier1.json` (Tier1 module->conformance cases)
- `docs/portable-semantics-v1.md` (semantic rules)

## Ownership Class Definitions

1. `haxe_source`
   - Behavior lives in Haxe std sources (`std/_std` overrides or upstream std implementation).
2. `runtime_binding`
   - Haxe surface delegates behavior to runtime package functions in `runtime/hxrt/*.go`.
3. `compiler_intrinsic`
   - Behavior is emitted directly by compiler lowering/shim generation.
4. `mixed`
   - Surface spans more than one class above; split is explicit and test-gated.

## Tier1 Mapping Table

| Module | Ownership class | Primary implementation location | Runtime dependency | Tier1 conformance cases |
| --- | --- | --- | --- | --- |
| `Math` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerStdlibSymbolShimDecls`) | Indirect via core helpers where needed | `numeric_edge_cases`, `stringtools_math` |
| `Std` | `mixed` (`compiler_intrinsic` + runtime helper calls) | `src/reflaxe/go/GoCompiler.hx` (`lowerStdlibSymbolShimDecls`, core lowering paths) | `runtime/hxrt/string.go`, `runtime/hxrt/exception.go`, core helpers | `exception_api_contract`, `std_is_of_type_contract`, `std_is_of_type_runtime_core_abstract_contract`, `typed_nil_dynamic_string_contract` |
| `DateTools` | `haxe_source` | `std/DateTools.cross.hx` | None beyond core `Date` and string primitives already owned elsewhere | `stringbuf_datetools_lambda_contract`, `datetools_cross_std_contract` |
| `StringTools` | `haxe_source` | `std/StringTools.cross.hx` | None beyond core string/runtime primitives used by normal lowering | `stringtools_math`, `stringtools_cross_std_contract` |
| `Sys` | `mixed` (`compiler wrapper` + `runtime_binding`) | `src/reflaxe/go/GoCompiler.hx` (`lowerSysStdlibShimDecls`) | `runtime/hxrt/sys.go` | `sys_io_roundtrip` |
| `haxe.Utf8` | `haxe_source` | `std/haxe/Utf8.cross.hx` | None beyond `haxe.io.Bytes`, `UnicodeString`, and shared Go string runtime helpers already owned elsewhere | `haxe_utf8_contract`, `stdlib/haxe_utf8_basic` |
| `haxe.Json` | `mixed` (`haxe_source` + `runtime_binding`) | `std/_std/haxe/Json.cross.hx` | `runtime/hxrt/json.go` | `json_parse_stringify_contract` |
| `haxe.ds.EnumValueMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.ds.IntMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.ds.ObjectMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.ds.StringMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.io.Bytes` | `mixed` (`compiler_intrinsic` + runtime helper calls) | `src/reflaxe/go/GoCompiler.hx` (`lowerIoStdlibShimDecls`) | `runtime/hxrt/bytes.go`, `runtime/hxrt/string.go` | `bytes_hex_contract`, `bytes_io_stream_contract`, `bytes_normalization_contract`, `bytes_of_data_contract`, `bytes_ops_contract`, `io_encoding_contract` |
| `haxe.io.Path` | `haxe_source` | `std/haxe/io/Path.cross.hx` | None beyond core string primitives already lowered by the target | `option_date_path`, `path_cross_std_contract` |
| `sys.FileSystem` | `mixed` (`compiler_intrinsic` + runtime helper calls) | `src/reflaxe/go/GoCompiler.hx` (`lowerFileSystemShimDecls`) | `runtime/hxrt/string.go`, `runtime/hxrt/exception.go` | `filesystem_contract` |
| `sys.io.File` | `runtime_binding` | `src/reflaxe/go/GoCompiler.hx` (`lowerSysStdlibShimDecls` forwarding wrappers) | `runtime/hxrt/sys.go` (`FileSaveContent`, `FileGetContent`) | `file_read_write_contract` |
| `sys.io.Process` | `runtime_binding` | `src/reflaxe/go/GoCompiler.hx` (`lowerSysStdlibShimDecls` forwarding wrappers) | `runtime/hxrt/process.go` (`NewProcess`, `ProcessOutput`) | `process_echo_contract` |
| `sys.net.Socket` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerNetSocketShimDecls`) | Uses core runtime helpers (`Throw`, string conversion) where needed | `socket_advanced_contract`, `socket_loopback_contract` |

## Notes on Staged Source Injection

Staged portable overrides are injected first for Go builds by:

- `src/reflaxe/go/CompilerBootstrap.hx`

Current staged Tier1 coverage includes the JSON family and `StringTools`, with additional migrations gated by semantic-diff and Tier1 conformance coverage.

## Transition Notes (Post-`__go__` Audit)

- `haxe.io.Bytes`
  - The current `mixed` classification is still correct for parity today.
  - The first post-`__go__` extraction already moved pure hex and `BytesBuffer` leaf helpers into `runtime/hxrt/bytes.go`, leaving thin compiler wrappers in place.
  - The remaining compiler-owned subset is the RawNative/cache-coupled string path (`ofString`, `getString`, UTF16/raw-native conversion helpers) because it still co-owns `__hx_raw` cache validity and encoding-tag behavior.
  - Tracking: `haxe.go-14as.51`, `haxe.go-14as.54`
- `haxe.io.Input` / `haxe.io.Output`
  - These surfaces are not listed as separate Tier1 rows here, but their inherited helper loops no longer live as raw loop bodies in `GoCompiler`.
  - `readAll`, `readLine`, `readUntil`, `readFullBytes`, `write`, `writeFullBytes`, `writeInput`, and `writeString` now route through `std/haxe/io/GoIoHelpers.cross.hx`, with `GoCompiler` keeping only the public wrapper functions and the representation-sensitive base IO types.
  - Tracking: `haxe.go-14as.52`
- `sys.Http`
  - Tier1 mapping still treats the surface as compiler-owned because request/callback choreography remains one semantic contract.
  - The audit narrowed extraction to leaf payload/proxy helpers only; `getResponseHeaderValues` and payload capture now live in `std/sys/GoHttpHelpers.cross.hx`, while core request sequencing and proxy URL construction stay in compiler scope unless parity evidence proves otherwise.
  - Tracking: `haxe.go-14as.53`

## Governance Rule

Any ownership change for a Tier1 module must update all of:

1. this mapping document,
2. `test/portable_conformance_tier1.json`,
3. `docs/stdlib-provenance-ledger.json` (when staged source files are added/changed),
4. relevant conformance fixtures in `test/semantic_diff`.
