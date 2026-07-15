#!/usr/bin/env python3

from __future__ import annotations

import argparse
import json
from pathlib import Path
from typing import Any

ROOT = Path(__file__).resolve().parent.parent
CACHE_ROOT = ROOT / "test" / ".test-cache"
FULL_MODULES_FILE = ROOT / "test" / "upstream_std_modules_full.txt"
STRICT_SWEEP_MODULES_FILE = ROOT / "test" / "upstream_std_modules.txt"
INVENTORY_FILE = ROOT / "test" / "portable_stdlib_inventory.json"
PROMOTIONS_FILE = ROOT / "test" / "portable_parity_promotions.json"
SUMMARY_JSON = CACHE_ROOT / "portable_stdlib_inventory_summary.json"
SUMMARY_MD = CACHE_ROOT / "portable_stdlib_inventory_summary.md"

ALLOWED_STATUS = {"unsupported", "compile-only", "snapshot", "semantic-diff"}
ALLOWED_OWNER = {"unassigned", "mixed", "compiler_shim", "runtime_hxrt", "staged_std"}

SEMANTIC_DIFF_EXPLICIT = {
    "Array",
    "Class",
    "Date",
    "DateTools",
    "EReg",
    "Enum",
    "EnumValue",
    "IntIterator",
    "Lambda",
    "Math",
    "Reflect",
    "String",
    "StringBuf",
    "Std",
    "StringTools",
    "Type",
    "haxe.Exception",
    "haxe.Int32",
    "haxe.Int64",
    "haxe.Int64Helper",
    "haxe.Json",
    "haxe.PosInfos",
    "haxe.Serializer",
    "haxe.Unserializer",
    "haxe.atomic.AtomicBool",
    "haxe.atomic.AtomicInt",
    "haxe.atomic.AtomicObject",
    "haxe.ds.EnumValueMap",
    "haxe.ds.IntMap",
    "haxe.ds.List",
    "haxe.ds.Map",
    "haxe.ds.Option",
    "haxe.ds.ObjectMap",
    "haxe.ds.ReadOnlyArray",
    "haxe.ds.StringMap",
    "haxe.ds.Vector",
    "haxe.iterators.HashMapKeyValueIterator",
    "haxe.iterators.MapKeyValueIterator",
    "haxe.format.JsonParser",
    "haxe.format.JsonPrinter",
    "haxe.io.Bytes",
    "haxe.io.BytesBuffer",
    "haxe.io.BytesInput",
    "haxe.io.BytesOutput",
    "haxe.io.Input",
    "haxe.io.Output",
    "haxe.io.Path",
    "sys.FileSystem",
    "sys.Http",
    "sys.io.File",
    "sys.io.Process",
    "sys.net.Host",
    "sys.net.Socket",
}

SEMANTIC_DIFF_PREFIXES = (
    "haxe.crypto.",
    "haxe.xml.",
    "haxe.zip.",
)

OWNER_OVERRIDES = {
	"Sys": "mixed",
    "haxe.http.HttpBase": "staged_std",
    "haxe.EntryPoint": "mixed",
    "haxe.CallStack": "staged_std",
    "haxe.Json": "runtime_hxrt",
    "haxe.MainLoop": "mixed",
    "haxe.format.JsonParser": "runtime_hxrt",
    "haxe.format.JsonPrinter": "runtime_hxrt",
    "haxe.atomic.AtomicBool": "runtime_hxrt",
    "haxe.atomic.AtomicInt": "runtime_hxrt",
    "haxe.atomic.AtomicObject": "runtime_hxrt",
    "sys.io.File": "mixed",
    "sys.io.FileInput": "mixed",
    "sys.io.FileOutput": "mixed",
    "sys.io.FileSeek": "staged_std",
    "sys.io.Process": "runtime_hxrt",
    "sys.FileSystem": "mixed",
    "sys.db.Connection": "staged_std",
    "sys.db.Mysql": "staged_std",
    "sys.db.ResultSet": "staged_std",
    "sys.db.Sqlite": "staged_std",
    "sys.net.Address": "staged_std",
    "sys.net.UdpSocket": "compiler_shim",
    "sys.ssl.Certificate": "mixed",
    "sys.ssl.Digest": "mixed",
    "sys.ssl.DigestAlgorithm": "staged_std",
    "sys.ssl.Key": "mixed",
    "sys.ssl.Socket": "mixed",
    "DateTools": "staged_std",
    "StringTools": "staged_std",
    "haxe.SysTools": "staged_std",
    "haxe.io.Path": "staged_std",
    "haxe.NativeStackTrace": "staged_std",
    "haxe.Template": "staged_std",
    "haxe.Timer": "mixed",
    "haxe.Utf8": "staged_std",
    "haxe.io.BufferInput": "compiler_shim",
    "haxe.io.BytesData": "mixed",
    "haxe.io.Encoding": "compiler_shim",
    "haxe.io.Eof": "compiler_shim",
    "haxe.io.Error": "compiler_shim",
    "haxe.io.FPHelper": "staged_std",
    "haxe.io.ArrayBufferView": "mixed",
    "haxe.io.UInt8Array": "mixed",
    "haxe.io.UInt16Array": "mixed",
    "haxe.io.UInt32Array": "mixed",
    "haxe.io.Int32Array": "mixed",
    "haxe.io.Float32Array": "mixed",
    "haxe.io.Float64Array": "mixed",
    "haxe.io.Mime": "staged_std",
    "haxe.io.Scheme": "staged_std",
    "haxe.io.StringInput": "compiler_shim",
    "haxe.rtti.CType": "mixed",
    "haxe.rtti.Meta": "mixed",
    "haxe.rtti.Rtti": "mixed",
    "haxe.rtti.XmlParser": "mixed",
    "haxe.ds.WeakMap": "staged_std",
    "EReg": "compiler_shim",
    "haxe.Serializer": "compiler_shim",
    "haxe.Unserializer": "compiler_shim",
    "sys.Http": "compiler_shim",
    "sys.net.Host": "compiler_shim",
    "sys.net.Socket": "compiler_shim",
    "haxe.io.Bytes": "compiler_shim",
    "haxe.io.BytesBuffer": "compiler_shim",
    "haxe.io.BytesInput": "compiler_shim",
    "haxe.io.BytesOutput": "compiler_shim",
    "haxe.io.Input": "compiler_shim",
    "haxe.io.Output": "compiler_shim",
}

UNSUPPORTED_EXPLICIT = {
    "haxe.http.HttpJs",
    "haxe.http.HttpNodeJs",
}

MODULE_NOTES_OVERRIDES = {
	"Sys": (
		"The supported Haxe 4.3.7 root API is canonical staged source in "
		"std/go/_std/Sys.hx over typed std/hxrt bindings. Public map construction, "
		"fallbacks, aliases, and standard-stream wrappers stay in Haxe; native process "
		"capabilities live in runtime/hxrt/sys.go and standard handles in file.go. "
		"Evidence: root_sys_contract, root_sys_portable_contract, sys_command_contract, "
		"sys_sleep_contract, sys/root_sys_portable, core/runtime_hxrt_infer_sys, and "
		"negative/sys_cpu_time_unsupported."
	),
	"haxe.CallStack": (
		"Covered by target-sensitive snapshot contracts in stdlib/haxe_stack_loop_target_sensitive. "
		"Go uses a staged std deterministic empty-stack fallback by default; opt-in native Go stack "
		"diagnostics are covered by stdlib/haxe_native_stack_trace_opt_in and do not claim portable "
		"semantic-diff stack parity."
	),
    "haxe.EntryPoint": (
        "Direct haxe.EntryPoint usage is covered by snapshot/runtime smoke contract "
        "stdlib/haxe_main_loop_runtime_direct. The staged override connects the public "
        "EntryPoint API to the runtime-backed sys.thread.EventLoop main-thread loop."
    ),
    "haxe.http.HttpBase": (
        "Direct haxe.http.HttpBase constructor/base-field/request baseline now has semantic-diff coverage "
        "through the staged override in std/go/_std/haxe/http/HttpBase.hx. Evidence: "
        "semantic_diff/haxe_http_base_contract and stdlib/haxe_http_base_direct."
    ),
    "haxe.http.HttpJs": (
        "Direct haxe.http.HttpJs usage is explicitly unsupported on Go. "
        "This is a JS-only target-conditional std module and the Haxe frontend already rejects it "
        "on non-JS targets. Evidence: negative/direct_haxe_httpjs_unsupported."
    ),
    "haxe.http.HttpNodeJs": (
        "Direct haxe.http.HttpNodeJs usage is explicitly unsupported on Go. "
        "This is a Node-only target-conditional std module and the Haxe frontend already rejects it "
        "on non-Node targets. Evidence: negative/direct_haxe_httpnodejs_unsupported."
    ),
    "haxe.http.HttpMethod": (
        "Direct haxe.http.HttpMethod abstract usage now has semantic-diff coverage through "
        "haxe_http_base_contract and stdlib/haxe_http_base_direct."
    ),
    "haxe.http.HttpStatus": (
        "Direct haxe.http.HttpStatus abstract usage now has semantic-diff coverage through "
        "haxe_http_base_contract and stdlib/haxe_http_base_direct."
    ),
    "haxe.MainLoop": (
        "Direct haxe.MainLoop usage is covered by snapshot/runtime smoke contract "
        "stdlib/haxe_main_loop_runtime_direct. The staged override keeps the Haxe-facing "
        "MainEvent facade while scheduling callbacks through haxe.EntryPoint and sys.thread.EventLoop."
    ),
	"haxe.NativeStackTrace": (
		"Covered by target-sensitive snapshot contracts in stdlib/haxe_stack_loop_target_sensitive. "
		"Go exposes a deterministic empty-stack fallback by default; opt-in native Go stack diagnostics "
		"are covered by stdlib/haxe_native_stack_trace_opt_in and do not claim portable semantic-diff stack parity."
	),
    "haxe.Ucs2": (
        "Covered by target-sensitive snapshot contract stdlib/haxe_ucs2_platform_exclusion. "
        "Go keeps the upstream platform-exclusion behavior and does not claim native UCS2 strings."
    ),
    "haxe.Utf8": (
        "Covered by staged-std semantic-diff contract haxe_utf8_contract plus snapshot "
        "stdlib/haxe_utf8_basic. Go preserves the deprecated buffer constructor with and without "
        "the optional size hint plus the helper subset through std/go/_std/haxe/Utf8.hx. The size "
        "hint is ignored because it is only a deprecated capacity hint and has no visible buffer "
        "semantics."
    ),
    "haxe.Timer": (
        "Direct haxe.Timer usage is covered by snapshot/runtime smoke contract "
        "stdlib/haxe_main_loop_runtime_direct. The staged override registers timers on "
        "Thread.current().events and uses the hxrt monotonic thread clock for Timer.stamp()."
    ),
    "Xml": (
        "Root Xml DOM subset is covered by root_xml_contract and stdlib/xml_root_dom_basic, "
        "including parsed CDATA node-type preservation."
    ),
    "haxe.Template": (
        "Direct haxe.Template constructor/execute usage now has semantic-diff coverage through "
        "a staged std override in std/go/_std/haxe/Template.hx. Evidence: "
        "semantic_diff/haxe_template_contract and stdlib/haxe_template_basic."
    ),
    "haxe.io.BufferInput": (
        "Direct `haxe.io.BufferInput` constructor and buffered-read baseline now have semantic-diff "
        "coverage through `semantic_diff/haxe_io_misc_contract` and snapshot coverage in "
        "`stdlib/haxe_io_misc_direct`. Ownership stays compiler-owned with the base IO hierarchy "
        "because this backend still owns the representation-sensitive `Input` / `BytesInput` / "
        "`BufferInput` type shapes and inherited helper wiring."
    ),
    "haxe.io.BytesData": (
        "Direct `haxe.io.BytesData` alias semantics now have semantic-diff coverage through "
        "`semantic_diff/haxe_io_misc_contract` and snapshot coverage in `stdlib/haxe_io_misc_direct`, "
        "including `Bytes.getData()` / `Bytes.ofData()` alias mutation behavior. Ownership stays mixed "
        "because the public alias is source-level, while the actual behavior still rides on the "
        "compiler-owned `haxe.io.Bytes` carrier."
    ),
    "haxe.io.Encoding": (
        "Direct `haxe.io.Encoding` constructor and pattern-match parity now have semantic-diff coverage "
        "through `semantic_diff/haxe_io_misc_contract` and snapshot coverage in `stdlib/haxe_io_misc_direct`. "
        "Ownership stays compiler-owned with the base IO hierarchy because the backend still emits the "
        "encoding tag values used by `Bytes`, `Input`, and `Output` shims."
    ),
    "haxe.io.Eof": (
        "Direct `haxe.io.Eof` construction and string parity now have semantic-diff coverage through "
        "`semantic_diff/haxe_io_misc_contract` and snapshot coverage in `stdlib/haxe_io_misc_direct`. "
        "Ownership stays compiler-owned with the base IO hierarchy because the backend still emits the "
        "`Eof` carrier used by IO shims and exception matching."
    ),
    "haxe.io.Error": (
        "Direct `haxe.io.Error` constructor and pattern-match parity now have semantic-diff coverage "
        "through `semantic_diff/haxe_io_misc_contract` and snapshot coverage in `stdlib/haxe_io_misc_direct`. "
        "Ownership stays compiler-owned with the base IO hierarchy because the backend still emits the "
        "error-tag carrier used by the public IO shims."
    ),
    "haxe.io.FPHelper": (
        "Direct `haxe.io.FPHelper` bit-conversion helpers now have semantic-diff coverage through "
        "`semantic_diff/haxe_io_misc_contract` and snapshot coverage in `stdlib/haxe_io_misc_direct`. "
        "The public API now lives in the staged override `std/go/_std/haxe/io/FPHelper.hx`, expressed on "
        "top of the existing little-endian `BytesInput` / `BytesOutput` contract instead of more raw Go."
    ),
    "haxe.io.ArrayBufferView": (
        "Direct `haxe.io.ArrayBufferView` usage now has semantic-diff coverage through "
        "`semantic_diff/haxe_io_typed_arrays_contract` and snapshot coverage in "
        "`stdlib/haxe_io_typed_arrays_direct`. Ownership stays mixed: the public typed-array API now lives "
        "in staged overrides under `std/go/_std/haxe/io/*.hx`, while the underlying carrier still rides on "
        "the compiler-owned `haxe.io.Bytes` / `ArrayBufferViewImpl` representation."
    ),
    "haxe.io.UInt8Array": (
        "Direct `haxe.io.UInt8Array` usage now has semantic-diff coverage through "
        "`semantic_diff/haxe_io_typed_arrays_contract` and snapshot coverage in "
        "`stdlib/haxe_io_typed_arrays_direct`. Ownership stays mixed: staged std owns the public typed-array "
        "API, while the actual storage and byte normalization still ride on the compiler-owned "
        "`haxe.io.Bytes` carrier."
    ),
    "haxe.io.UInt16Array": (
        "Direct `haxe.io.UInt16Array` usage now has semantic-diff coverage through "
        "`semantic_diff/haxe_io_typed_arrays_contract` and snapshot coverage in "
        "`stdlib/haxe_io_typed_arrays_direct`. Ownership stays mixed: staged std owns the public typed-array "
        "API, while the actual storage and byte normalization still ride on the compiler-owned "
        "`haxe.io.Bytes` carrier."
    ),
    "haxe.io.UInt32Array": (
        "Direct `haxe.io.UInt32Array` usage now has semantic-diff coverage through "
        "`semantic_diff/haxe_io_typed_arrays_contract` and snapshot coverage in "
        "`stdlib/haxe_io_typed_arrays_direct`. Ownership stays mixed: staged std owns the public typed-array "
        "API, while the actual storage and byte normalization still ride on the compiler-owned "
        "`haxe.io.Bytes` carrier."
    ),
    "haxe.io.Int32Array": (
        "Direct `haxe.io.Int32Array` usage now has semantic-diff coverage through "
        "`semantic_diff/haxe_io_typed_arrays_contract` and snapshot coverage in "
        "`stdlib/haxe_io_typed_arrays_direct`. Ownership stays mixed: staged std owns the public typed-array "
        "API, while the actual storage and byte normalization still ride on the compiler-owned "
        "`haxe.io.Bytes` carrier."
    ),
    "haxe.io.Float32Array": (
        "Direct `haxe.io.Float32Array` usage now has semantic-diff coverage through "
        "`semantic_diff/haxe_io_typed_arrays_contract` and snapshot coverage in "
        "`stdlib/haxe_io_typed_arrays_direct`. Ownership stays mixed: staged std owns the public typed-array "
        "API, the actual storage still rides on the compiler-owned `haxe.io.Bytes` carrier, and the float "
        "bit conversions are expressed through staged `haxe.io.FPHelper` helpers instead of new raw compiler "
        "bytes logic."
    ),
    "haxe.io.Float64Array": (
        "Direct `haxe.io.Float64Array` usage now has semantic-diff coverage through "
        "`semantic_diff/haxe_io_typed_arrays_contract` and snapshot coverage in "
        "`stdlib/haxe_io_typed_arrays_direct`. Ownership stays mixed: staged std owns the public typed-array "
        "API, the actual storage still rides on the compiler-owned `haxe.io.Bytes` carrier, and the float "
        "bit conversions are expressed through staged `haxe.io.FPHelper` helpers instead of new raw compiler "
        "bytes logic."
    ),
    "haxe.io.Mime": (
        "Direct `haxe.io.Mime` abstract usage now has semantic-diff coverage through "
        "`semantic_diff/haxe_io_misc_contract` and snapshot coverage in `stdlib/haxe_io_misc_direct`. "
        "Ownership stays source-owned because this surface is just the upstream string-backed abstract "
        "running on the normal string lowering contract."
    ),
    "haxe.io.Path": (
        "Direct `haxe.io.Path` usage now has semantic-diff coverage through "
        "`semantic_diff/path_cross_std_contract` and snapshot coverage in `stdlib/path_cross_std_basic`. "
        "Ownership has graduated back to the upstream Haxe stdlib implementation; Go only owns the reusable "
        "string and array lowerings that upstream `Path.hx` depends on."
    ),
    "haxe.io.Scheme": (
        "Direct `haxe.io.Scheme` abstract usage now has semantic-diff coverage through "
        "`semantic_diff/haxe_io_misc_contract` and snapshot coverage in `stdlib/haxe_io_misc_direct`. "
        "Ownership stays source-owned because this surface is just the upstream string-backed abstract "
        "running on the normal string lowering contract."
    ),
    "haxe.io.StringInput": (
        "Direct `haxe.io.StringInput` constructor and inherited-read baseline now have semantic-diff "
        "coverage through `semantic_diff/haxe_io_misc_contract` and snapshot coverage in "
        "`stdlib/haxe_io_misc_direct`. Ownership stays compiler-owned with the base IO hierarchy because "
        "this backend still owns the representation-sensitive `BytesInput` / `StringInput` type shapes "
        "and inherited helper wiring."
    ),
    "haxe.ValueException": (
        "Direct haxe.ValueException constructor/message/value parity now has semantic-diff coverage. "
        "Evidence: semantic_diff/haxe_value_exception_contract and stdlib/haxe_value_exception_basic."
    ),
    "haxe.ds.WeakMap": (
        "Direct WeakMap usage now has semantic-diff coverage for the upstream platform behavior: "
        "on Go, `new haxe.ds.WeakMap()` preserves the generic Haxe stdlib contract and throws "
        "`haxe.exceptions.NotImplementedException` instead of pretending to expose real weak references. "
        "Evidence: semantic_diff/haxe_ds_weakmap_contract and stdlib/haxe_ds_weakmap_platform."
    ),
    "haxe.rtti.CType": (
        "Direct RTTI typedefs, enums, `TypeApi`, and `CTypeTools` now have semantic-diff coverage through "
        "`semantic_diff/haxe_rtti_direct_contract` and snapshot coverage in `stdlib/haxe_rtti_direct`. "
        "Ownership stays mixed: the public Haxe-facing API lives in staged std overrides under `std/haxe/rtti/**`, "
        "while the backend still owns the narrow class-token `__meta__` / `__rtti` contract and anonymous-record "
        "array-field mutation lowering that those overrides rely on."
    ),
    "haxe.rtti.Meta": (
        "Direct `haxe.rtti.Meta` access now has semantic-diff coverage through `semantic_diff/haxe_rtti_direct_contract` "
        "and snapshot coverage in `stdlib/haxe_rtti_direct`. The staged override keeps the public API source-owned "
        "while routing metadata lookup through the backend-owned class-token `__meta__` contract."
    ),
    "haxe.rtti.Rtti": (
        "Direct `haxe.rtti.Rtti` access now has semantic-diff coverage through `semantic_diff/haxe_rtti_direct_contract` "
        "and snapshot coverage in `stdlib/haxe_rtti_direct`. The staged override keeps the public API source-owned "
        "while routing RTTI lookup through the backend-owned class-token `__rtti` contract."
    ),
    "haxe.rtti.XmlParser": (
        "Direct `haxe.rtti.XmlParser` parsing now has semantic-diff coverage through `semantic_diff/haxe_rtti_direct_contract` "
        "and snapshot coverage in `stdlib/haxe_rtti_direct`. Ownership stays mixed: parser logic is staged std, while the "
        "backend still owns the anonymous-record array-field mutation fix that makes RTTI record merging lower honestly on Go."
    ),
    "sys.db.Connection": (
        "Direct `sys.db.Connection` interface usage now has semantic-diff coverage through "
        "`semantic_diff/sys_db_io_contract` and snapshot coverage in `stdlib/sys_db_io_direct`. "
        "Ownership is explicit instead of vague: the public DB contract stays in upstream Haxe std source, while Go-specific "
        "runtime binding only begins at the file-handle side of the same tranche."
    ),
    "sys.db.Mysql": (
        "Direct `sys.db.Mysql.connect` usage now has semantic-diff coverage through `semantic_diff/sys_db_io_contract` "
        "and snapshot coverage in `stdlib/sys_db_io_direct`, preserving the upstream Go-platform contract that this API "
        "throws `haxe.exceptions.NotImplementedException` instead of pretending to expose a fake database runtime."
    ),
    "sys.db.ResultSet": (
        "Direct `sys.db.ResultSet` interface usage now has semantic-diff coverage through "
        "`semantic_diff/sys_db_io_contract` and snapshot coverage in `stdlib/sys_db_io_direct`, including length/field-name/"
        "cursor access through user-defined result-set implementations."
    ),
    "sys.db.Sqlite": (
        "Direct `sys.db.Sqlite.open` usage now has semantic-diff coverage through `semantic_diff/sys_db_io_contract` "
        "and snapshot coverage in `stdlib/sys_db_io_direct`, preserving the upstream Go-platform contract that this API "
        "throws `Not implemented for this platform` instead of claiming a nonexistent SQLite binding."
    ),
    "sys.net.Address": (
        "Direct `sys.net.Address` usage now has semantic-diff coverage through "
        "`semantic_diff/sys_net_address_ssl_digest_algorithm_contract` and snapshot coverage in "
        "`stdlib/sys_net_address_ssl_digest_algorithm_direct`. Ownership stays source-owned because this "
        "surface is just the upstream `{host, port}` carrier and helper methods, expressed without growing "
        "the compiler-owned socket runtime."
    ),
    "sys.net.UdpSocket": (
        "Direct `sys.net.UdpSocket` usage now has deterministic snapshot/runtime coverage through "
        "`stdlib/sys_net_udp_socket_direct`, covering loopback `bind` / `host` / `sendTo` / `readFrom` / `setBroadcast` "
        "plus peer address round-tripping. Ownership stays compiler-owned inside the `net_socket` shim "
        "group because UDP deadline/blocking/address translation and broadcast socket-option behavior are still target-sensitive. "
        "The evidence covers OS socket-option installation without requiring LAN broadcast packet delivery in CI."
    ),
    "sys.ssl.Certificate": (
        "Direct `sys.ssl.Certificate` leaf usage now has snapshot runtime coverage through "
        "`stdlib/sys_ssl_leaf_direct`. Ownership stays mixed: the public Haxe-facing wrapper lives in "
        "`std/go/_std/sys/ssl/Certificate.hx`, while certificate parsing and trust-root handling live in "
        "`runtime/hxrt/ssl.go`."
    ),
    "sys.ssl.Digest": (
        "Direct `sys.ssl.Digest` leaf usage now has snapshot runtime coverage through "
        "`stdlib/sys_ssl_leaf_direct`, including deterministic SHA-256/SHA-512 digests plus sign/verify "
        "with a parsed private key. Ownership stays mixed: the public Haxe API lives in "
        "`std/go/_std/sys/ssl/Digest.hx`, while the actual cryptographic work lives in `runtime/hxrt/ssl.go`."
    ),
    "sys.io.File": (
        "The complete Haxe 4.3.7 `sys.io.File` API is canonical staged source in `std/go/_std/sys/io/File.hx`, "
        "with typed `std/hxrt/fs` bindings to native capabilities in selectively copied `runtime/hxrt/file.go`. "
        "Semantic coverage includes binary save/read/copy, write/append/update modes, seek/tell, bounds, and EOF "
        "through `file_read_write_contract`, `file_error_semantics_contract`, and `sys_db_io_contract`; snapshot "
        "shape coverage lives in `sys/file_read_write_smoke`, `sys/file_error_semantics`, and `stdlib/sys_db_io_direct`."
    ),
    "sys.io.FileInput": (
        "Direct `sys.io.FileInput` usage now has semantic-diff coverage through `semantic_diff/sys_db_io_contract` "
        "and snapshot coverage in `stdlib/sys_db_io_direct`. The public stream implementation, bounds checks, EOF translation, "
        "and seek-origin mapping live in `std/go/_std/sys/io/FileInput.hx`; only its opaque OS handle and native operations live "
        "in `runtime/hxrt/file.go`."
    ),
    "sys.io.FileOutput": (
        "Direct `sys.io.FileOutput` usage now has semantic-diff coverage through `semantic_diff/sys_db_io_contract` "
        "and snapshot coverage in `stdlib/sys_db_io_direct`. The public stream implementation, bounds checks, byte conversion, "
        "and seek-origin mapping live in `std/go/_std/sys/io/FileOutput.hx`; only its opaque OS handle and native operations live "
        "in `runtime/hxrt/file.go`."
    ),
    "sys.io.FileSeek": (
        "Direct `sys.io.FileSeek` enum usage now has semantic-diff coverage through `semantic_diff/sys_db_io_contract` "
        "and snapshot coverage in `stdlib/sys_db_io_direct`. The enum is canonical staged source, and the staged stream methods "
        "select native seek origins explicitly without a compiler-synthesized carrier or mapper."
    ),
    "sys.ssl.DigestAlgorithm": (
        "Direct `sys.ssl.DigestAlgorithm` usage now has semantic-diff coverage through "
        "`semantic_diff/sys_net_address_ssl_digest_algorithm_contract` and snapshot coverage in "
        "`stdlib/sys_net_address_ssl_digest_algorithm_direct`. Ownership stays source-owned because this surface "
        "is the upstream string-backed algorithm enum, independent of the target-owned TLS and digest machinery underneath."
    ),
    "sys.ssl.Key": (
        "Direct `sys.ssl.Key` leaf usage now has snapshot runtime coverage through `stdlib/sys_ssl_leaf_direct`. "
        "Ownership stays mixed: the public Haxe constructors live in `std/go/_std/sys/ssl/Key.hx`, while PEM/DER "
        "parsing and native key storage live in `runtime/hxrt/ssl.go`. Omitted optional pass arguments still depend "
        "on the broader source-owned static optional-argument lowering follow-up."
    ),
    "sys.ssl.Socket": (
        "Direct `sys.ssl.Socket` usage now has snapshot runtime coverage through `stdlib/sys_ssl_socket_direct` "
        "and SNI selection coverage through `stdlib/sys_ssl_socket_sni_direct`, "
        "covering staged public socket configuration on top of the compiler-owned `sys.net.Socket` carrier plus "
        "runtime TLS dial/listen/handshake/peer-certificate/SNI behavior in `runtime/hxrt/ssl.go`. Ownership stays mixed: "
        "the public Haxe API lives in `std/go/_std/sys/ssl/Socket.hx`, while the underlying connection/listener TLS "
        "behavior remains runtime-owned."
    ),
}

PROMOTION_LEVEL_KEYS = ("snapshot", "semantic_diff")

BLOCKER_FAMILY_SPECS = (
    {
        "issue": "haxe.go-14as.25",
        "family": "haxe_misc_symbols",
        "closure_target": "2026-03-20",
        "modules": {
        },
    },
    {
        "issue": "haxe.go-14as.26",
        "family": "haxe_misc_abstractions",
        "closure_target": "2026-03-27",
        "modules": {
            "haxe.Constraints",
            "haxe.Rest",
        },
    },
    {
        "issue": "haxe.go-14as.27",
        "family": "haxe_misc_enum_helpers",
        "closure_target": "2026-04-03",
        "modules": {
            "haxe.EnumFlags",
            "haxe.EnumTools",
        },
    },
    {
        "issue": "haxe.go-14as.29",
        "family": "haxe_misc_legacy_text",
        "closure_target": "2026-04-17",
        "modules": {
            "haxe.Ucs2",
            "haxe.Utf8",
        },
    },
    {
        "issue": "haxe.go-14as.46",
        "family": "haxe_ds_source_owned_collections",
        "closure_target": "2026-04-14",
        "modules": {
            "haxe.ds.BalancedTree",
            "haxe.ds.GenericStack",
        },
    },
    {
        "issue": "haxe.go-14as.47",
        "family": "haxe_ds_weakmap",
        "closure_target": "2026-04-14",
        "modules": {
            "haxe.ds.WeakMap",
        },
    },
    {
        "issue": "haxe.go-14as.43",
        "family": "haxe_exceptions_direct",
        "closure_target": "2026-04-14",
        "modules": {
            "haxe.exceptions.ArgumentException",
            "haxe.exceptions.NotImplementedException",
            "haxe.exceptions.PosException",
        },
    },
    {
        "issue": "haxe.go-14as.57",
        "family": "haxe_rtti_direct",
        "closure_target": "2026-04-28",
        "modules": {
            "haxe.rtti.CType",
            "haxe.rtti.Meta",
            "haxe.rtti.Rtti",
            "haxe.rtti.XmlParser",
        },
    },
    {
        "issue": "haxe.go-14as.15",
        "family": "haxe_io_misc",
        "closure_target": "2026-04-30",
        "modules": {
            "haxe.io.BufferInput",
            "haxe.io.BytesData",
            "haxe.io.Encoding",
            "haxe.io.Eof",
            "haxe.io.Error",
            "haxe.io.FPHelper",
            "haxe.io.Mime",
            "haxe.io.Scheme",
            "haxe.io.StringInput",
        },
    },
    {
        "issue": "haxe.go-14as.16",
        "family": "haxe_io_typed_arrays",
        "closure_target": "2026-05-07",
        "modules": {
            "haxe.io.ArrayBufferView",
            "haxe.io.Float32Array",
            "haxe.io.Float64Array",
            "haxe.io.Int32Array",
            "haxe.io.UInt16Array",
            "haxe.io.UInt32Array",
            "haxe.io.UInt8Array",
        },
    },
    {
        "issue": "haxe.go-14as.18",
        "family": "sys_net_ssl",
        "closure_target": "2026-05-21",
        "modules": {
            "sys.net.Address",
            "sys.ssl.Certificate",
            "sys.ssl.Digest",
            "sys.ssl.DigestAlgorithm",
            "sys.ssl.Key",
        },
    },
    {
        "issue": "haxe.go-14as.19",
        "family": "sys_thread",
        "closure_target": "2026-05-31",
        "modules": {
            "sys.thread.ElasticThreadPool",
            "sys.thread.EventLoop",
            "sys.thread.FixedThreadPool",
            "sys.thread.Thread",
        },
    },
)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Validate and generate portable stdlib inventory (portable-eligible Haxe 4.3.7 modules)."
    )
    parser.add_argument(
        "--update",
        action="store_true",
        help="Write generated inventory to test/portable_stdlib_inventory.json.",
    )
    return parser.parse_args()


def load_modules(path: Path) -> list[str]:
    modules: list[str] = []
    for raw in path.read_text(encoding="utf-8").splitlines():
        line = raw.strip()
        if not line or line.startswith("#"):
            continue
        modules.append(line)
    return modules


def load_promotions(path: Path, valid_modules: set[str], auto_semantic_modules: set[str]) -> dict[str, set[str]]:
    try:
        raw = json.loads(path.read_text(encoding="utf-8"))
    except FileNotFoundError as exc:
        raise SystemExit(
            "portable parity promotions file missing. "
            "Expected: test/portable_parity_promotions.json"
        ) from exc
    except json.JSONDecodeError as exc:
        raise SystemExit(f"{path.name}: invalid JSON ({exc})") from exc

    if not isinstance(raw, dict):
        raise SystemExit(f"{path.name}: root must be an object")
    if raw.get("schema_version") != 1:
        raise SystemExit(f"{path.name}: schema_version must be 1")

    levels = raw.get("levels")
    if not isinstance(levels, dict):
        raise SystemExit(f"{path.name}: levels must be an object")

    unknown_keys = sorted(set(levels.keys()) - set(PROMOTION_LEVEL_KEYS))
    if unknown_keys:
        raise SystemExit(
            f"{path.name}: unknown levels keys: {', '.join(unknown_keys)} "
            f"(allowed: {', '.join(PROMOTION_LEVEL_KEYS)})"
        )

    normalized: dict[str, set[str]] = {}
    for key in PROMOTION_LEVEL_KEYS:
        raw_modules = levels.get(key, [])
        if not isinstance(raw_modules, list):
            raise SystemExit(f"{path.name}: levels.{key} must be an array")
        if raw_modules != sorted(raw_modules):
            raise SystemExit(f"{path.name}: levels.{key} must be sorted")

        seen: set[str] = set()
        modules: set[str] = set()
        for module in raw_modules:
            if not isinstance(module, str) or not module.strip():
                raise SystemExit(f"{path.name}: levels.{key} contains invalid module entry: {module!r}")
            if module in seen:
                raise SystemExit(f"{path.name}: levels.{key} has duplicate module `{module}`")
            seen.add(module)
            if module not in valid_modules:
                raise SystemExit(f"{path.name}: levels.{key} module `{module}` not found in full module list")
            modules.add(module)
        normalized[key] = modules

    overlap = normalized["snapshot"] & normalized["semantic_diff"]
    if overlap:
        raise SystemExit(
            f"{path.name}: modules cannot be in both snapshot and semantic_diff levels: "
            + ", ".join(sorted(overlap))
        )

    redundant_snapshot = normalized["snapshot"] & auto_semantic_modules
    if redundant_snapshot:
        raise SystemExit(
            f"{path.name}: snapshot promotions cannot include modules already auto-promoted to semantic-diff: "
            + ", ".join(sorted(redundant_snapshot))
        )

    return normalized


def is_semantic_diff_module(module: str) -> bool:
    if module in SEMANTIC_DIFF_EXPLICIT:
        return True
    for prefix in SEMANTIC_DIFF_PREFIXES:
        if module.startswith(prefix):
            return True
    return False


def select_owner(module: str, status: str, in_strict_sweep: bool) -> str:
    if module in OWNER_OVERRIDES:
        return OWNER_OVERRIDES[module]
    if module.startswith(SEMANTIC_DIFF_PREFIXES):
        return "compiler_shim"
    if status == "compile-only" or in_strict_sweep:
        return "mixed"
    return "unassigned"


def blocker_plan(module: str) -> dict[str, str] | None:
    for spec in BLOCKER_FAMILY_SPECS:
        if module in spec["modules"]:
            return {
                "issue": str(spec["issue"]),
                "family": str(spec["family"]),
                "closure_target": str(spec["closure_target"]),
            }
    return None


def module_notes(module: str, status: str, in_strict_sweep: bool) -> str:
    if status == "semantic-diff":
        base = (
            "Covered by semantic-diff/runtime contracts in test/semantic_diff and "
            "documented in docs/feature-support-matrix.md."
        )
        override = MODULE_NOTES_OVERRIDES.get(module)
        if override:
            return base + " " + override
        return base
    if status == "snapshot":
        base = "Covered by snapshot-level deterministic generated-code/runtime smoke contracts."
        override = MODULE_NOTES_OVERRIDES.get(module)
        if override:
            return base + " " + override
        return base
    if status == "compile-only":
        blocker = blocker_plan(module)
        if blocker is None:
            raise SystemExit(f"compile-only module missing blocker plan: {module}")
        if in_strict_sweep:
            base = (
                "Covered by strict upstream stdlib sweep compile/go-test checks "
                "(test/upstream_std_modules.txt)."
            )
        else:
            base = (
                "Covered by full portable-eligible upstream stdlib sweep compile checks "
                "(test/upstream_std_modules_full.txt); runtime parity contracts are not yet promoted."
            )
        return (
            base
            + " "
            + f"Tracked by {blocker['issue']} ({blocker['family']}); closure target {blocker['closure_target']}."
        )
    if status == "unsupported":
        override = MODULE_NOTES_OVERRIDES.get(module)
        if override:
            return override
        return "Portable-eligible module is explicitly unsupported on Go; the repo does not claim parity for it."
    return "Portable-eligible module inventoried; parity promotion is pending."


def module_evidence(status: str, in_full_sweep: bool, in_strict_sweep: bool) -> list[str]:
    evidence: list[str] = []
    if in_full_sweep:
        evidence.append("upstream_sweep:full_compile")
    if in_strict_sweep:
        evidence.append("upstream_sweep:strict_go_test")
    if status == "snapshot":
        evidence.append("snapshot")
    if status == "semantic-diff":
        evidence.append("semantic_diff")
    return evidence


def build_inventory(
    full_modules: list[str], strict_sweep_modules: set[str], promotions: dict[str, set[str]]
) -> dict[str, Any]:
    snapshot_promotions = promotions["snapshot"]
    semantic_promotions = promotions["semantic_diff"]
    modules_payload: list[dict[str, Any]] = []
    for module in sorted(full_modules):
        in_full_sweep = True
        in_strict_sweep = module in strict_sweep_modules
        status = "compile-only"
        if module in UNSUPPORTED_EXPLICIT:
            status = "unsupported"
        elif is_semantic_diff_module(module) or module in semantic_promotions:
            status = "semantic-diff"
        elif module in snapshot_promotions:
            status = "snapshot"

        owner = select_owner(module, status, in_strict_sweep)
        entry = {
            "module": module,
            "portable_eligible": True,
            "status": status,
            "owner": owner,
            "in_full_sweep": in_full_sweep,
            "in_strict_sweep": in_strict_sweep,
            "coverage_evidence": module_evidence(status, in_full_sweep, in_strict_sweep),
            "notes": module_notes(module, status, in_strict_sweep),
        }
        blocker = blocker_plan(module) if status == "compile-only" else None
        if blocker is not None:
            entry["blocker_issue"] = blocker["issue"]
            entry["blocker_family"] = blocker["family"]
            entry["closure_target"] = blocker["closure_target"]
        modules_payload.append(entry)

    return {
        "schema_version": 1,
        "baseline": {
            "haxe_version": "4.3.7",
            "module_source": "test/upstream_std_modules_full.txt",
            "portable_scope": "portable-eligible modules only; target-specific namespaces excluded",
            "excluded_prefix_examples": ["cpp.*", "java.*", "cs.*", "hl.*", "lua.*", "php.*", "python.*", "js.*"],
        },
        "generated_by": "test/run-portable-stdlib-inventory.py",
        "modules": modules_payload,
    }


def validate_inventory_schema(inventory: dict[str, Any], full_modules: list[str]) -> None:
    if inventory.get("schema_version") != 1:
        raise SystemExit("portable_stdlib_inventory.json: schema_version must be 1")

    modules = inventory.get("modules")
    if not isinstance(modules, list):
        raise SystemExit("portable_stdlib_inventory.json: modules must be an array")

    seen: set[str] = set()
    ordered_modules: list[str] = []
    for entry in modules:
        if not isinstance(entry, dict):
            raise SystemExit("portable_stdlib_inventory.json: module entries must be objects")

        module = entry.get("module")
        status = entry.get("status")
        owner = entry.get("owner")
        portable_eligible = entry.get("portable_eligible")

        if not isinstance(module, str) or not module:
            raise SystemExit("portable_stdlib_inventory.json: module must be a non-empty string")
        if module in seen:
            raise SystemExit(f"portable_stdlib_inventory.json: duplicate module entry: {module}")
        seen.add(module)
        ordered_modules.append(module)

        if status not in ALLOWED_STATUS:
            raise SystemExit(f"portable_stdlib_inventory.json: invalid status for {module}: {status!r}")
        if owner not in ALLOWED_OWNER:
            raise SystemExit(f"portable_stdlib_inventory.json: invalid owner for {module}: {owner!r}")
        if portable_eligible is not True:
            raise SystemExit(f"portable_stdlib_inventory.json: portable_eligible must be true for {module}")
        if status == "compile-only":
            blocker_issue = entry.get("blocker_issue")
            blocker_family = entry.get("blocker_family")
            closure_target = entry.get("closure_target")
            if not isinstance(blocker_issue, str) or not blocker_issue.strip():
                raise SystemExit(f"portable_stdlib_inventory.json: compile-only module missing blocker_issue: {module}")
            if not isinstance(blocker_family, str) or not blocker_family.strip():
                raise SystemExit(f"portable_stdlib_inventory.json: compile-only module missing blocker_family: {module}")
            if not isinstance(closure_target, str) or not closure_target.strip():
                raise SystemExit(f"portable_stdlib_inventory.json: compile-only module missing closure_target: {module}")

    expected = sorted(full_modules)
    if sorted(seen) != expected:
        missing = sorted(set(expected) - set(seen))
        extra = sorted(set(seen) - set(expected))
        details: list[str] = []
        if missing:
            details.append(f"missing={missing[:10]}")
        if extra:
            details.append(f"extra={extra[:10]}")
        raise SystemExit(
            "portable_stdlib_inventory.json: module set must match test/upstream_std_modules_full.txt "
            + "; ".join(details)
        )

    if ordered_modules != sorted(ordered_modules):
        raise SystemExit("portable_stdlib_inventory.json: modules must be sorted by module name")


def build_summary(inventory: dict[str, Any]) -> dict[str, Any]:
    modules = inventory["modules"]
    status_counts: dict[str, int] = {status: 0 for status in sorted(ALLOWED_STATUS)}
    owner_counts: dict[str, int] = {owner: 0 for owner in sorted(ALLOWED_OWNER)}
    full_sweep_count = 0
    strict_sweep_count = 0
    for entry in modules:
        status_counts[entry["status"]] += 1
        owner_counts[entry["owner"]] += 1
        if entry.get("in_full_sweep"):
            full_sweep_count += 1
        if entry["in_strict_sweep"]:
            strict_sweep_count += 1

    return {
        "schema_version": 1,
        "total_modules": len(modules),
        "full_sweep_modules": full_sweep_count,
        "strict_sweep_modules": strict_sweep_count,
        "status_counts": status_counts,
        "owner_counts": owner_counts,
    }


def write_summary(summary: dict[str, Any]) -> None:
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    SUMMARY_JSON.write_text(json.dumps(summary, indent=2, sort_keys=True) + "\n", encoding="utf-8")

    lines = [
        "# Portable Stdlib Inventory Summary",
        "",
        f"- total_modules: `{summary['total_modules']}`",
        f"- full_sweep_modules: `{summary['full_sweep_modules']}`",
        f"- strict_sweep_modules: `{summary['strict_sweep_modules']}`",
        "",
        "## Status counts",
    ]
    for status, count in summary["status_counts"].items():
        lines.append(f"- `{status}`: `{count}`")

    lines.append("")
    lines.append("## Owner counts")
    for owner, count in summary["owner_counts"].items():
        lines.append(f"- `{owner}`: `{count}`")

    lines.append("")
    lines.append("Artifacts:")
    lines.append(f"- `{SUMMARY_JSON.relative_to(ROOT)}`")
    lines.append(f"- `{INVENTORY_FILE.relative_to(ROOT)}`")
    SUMMARY_MD.write_text("\n".join(lines) + "\n", encoding="utf-8")


def main() -> int:
    args = parse_args()

    full_modules = load_modules(FULL_MODULES_FILE)
    strict_sweep_modules = set(load_modules(STRICT_SWEEP_MODULES_FILE))
    full_module_set = set(full_modules)
    auto_semantic_modules = {module for module in full_modules if is_semantic_diff_module(module)}
    promotions = load_promotions(PROMOTIONS_FILE, full_module_set, auto_semantic_modules)
    generated = build_inventory(full_modules, strict_sweep_modules, promotions)

    if args.update:
        INVENTORY_FILE.write_text(json.dumps(generated, indent=2, sort_keys=True) + "\n", encoding="utf-8")
        print(f"[PASS] wrote {INVENTORY_FILE.relative_to(ROOT)}")

    if not INVENTORY_FILE.exists():
        raise SystemExit(
            "portable stdlib inventory file missing. Run: python3 test/run-portable-stdlib-inventory.py --update"
        )

    existing = json.loads(INVENTORY_FILE.read_text(encoding="utf-8"))
    validate_inventory_schema(existing, full_modules)

    if existing != generated:
        raise SystemExit(
            "portable stdlib inventory drift detected. "
            "Run: python3 test/run-portable-stdlib-inventory.py --update"
        )

    summary = build_summary(existing)
    write_summary(summary)
    print("[PASS] portable stdlib inventory validated")
    print(f"[PASS] summary: {SUMMARY_JSON.relative_to(ROOT)}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
