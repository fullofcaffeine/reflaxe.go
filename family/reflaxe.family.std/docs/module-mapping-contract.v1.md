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
| `StringTools` | `mixed` (`compiler_intrinsic` + runtime helper calls) | `src/reflaxe/go/GoCompiler.hx` (`lowerStdlibSymbolShimDecls`) | `runtime/hxrt/string.go` helpers (`StdString`, literal conversions) | `stringtools_math` |
| `Sys` | `mixed` (`compiler wrapper` + `runtime_binding`) | `src/reflaxe/go/GoCompiler.hx` (`lowerSysStdlibShimDecls`) | `runtime/hxrt/sys.go` | `sys_io_roundtrip` |
| `haxe.Json` | `mixed` (`haxe_source` + `runtime_binding`) | `std/_std/haxe/Json.cross.hx` | `runtime/hxrt/json.go` | `json_parse_stringify_contract` |
| `haxe.ds.EnumValueMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.ds.IntMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.ds.ObjectMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.ds.StringMap` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerDsStdlibShimDecls`) | Uses core runtime helpers for dynamic/null pathways | `ds_maps_list_contract` |
| `haxe.io.Bytes` | `mixed` (`compiler_intrinsic` + runtime helper calls) | `src/reflaxe/go/GoCompiler.hx` (`lowerIoStdlibShimDecls`) | `runtime/hxrt/bytes.go`, `runtime/hxrt/string.go` | `bytes_hex_contract`, `bytes_io_stream_contract`, `bytes_normalization_contract`, `bytes_of_data_contract`, `bytes_ops_contract`, `io_encoding_contract` |
| `sys.FileSystem` | `mixed` (`compiler_intrinsic` + runtime helper calls) | `src/reflaxe/go/GoCompiler.hx` (`lowerFileSystemShimDecls`) | `runtime/hxrt/string.go`, `runtime/hxrt/exception.go` | `filesystem_contract` |
| `sys.io.File` | `runtime_binding` | `src/reflaxe/go/GoCompiler.hx` (`lowerSysStdlibShimDecls` forwarding wrappers) | `runtime/hxrt/sys.go` (`FileSaveContent`, `FileGetContent`) | `file_read_write_contract` |
| `sys.io.Process` | `runtime_binding` | `src/reflaxe/go/GoCompiler.hx` (`lowerSysStdlibShimDecls` forwarding wrappers) | `runtime/hxrt/process.go` (`NewProcess`, `ProcessOutput`) | `process_echo_contract` |
| `sys.net.Socket` | `compiler_intrinsic` | `src/reflaxe/go/GoCompiler.hx` (`lowerNetSocketShimDecls`) | Uses core runtime helpers (`Throw`, string conversion) where needed | `socket_advanced_contract`, `socket_loopback_contract` |

## Notes on Staged Source Injection

Staged portable overrides are injected first for Go builds by:

- `src/reflaxe/go/CompilerBootstrap.hx`

Current staged Tier1 coverage is intentionally narrow (JSON family) while preserving deterministic parity via semantic-diff and Tier1 conformance gating.

## Governance Rule

Any ownership change for a Tier1 module must update all of:

1. this mapping document,
2. `test/portable_conformance_tier1.json`,
3. `docs/stdlib-provenance-ledger.json` (when staged source files are added/changed),
4. relevant conformance fixtures in `test/semantic_diff`.
