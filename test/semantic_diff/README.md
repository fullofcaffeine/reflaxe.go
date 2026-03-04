# Semantic Differential Fixtures

These fixtures compare runtime output between:

1. Haxe reference execution (`--interp`)
2. `reflaxe.go` output (`portable` profile, then `go run .`)

Runner:

```bash
python3 test/run-semantic-diff.py
```

Optional case-local Go defines:

- Add `go_defines.txt` inside a semantic case directory.
- Each non-empty, non-comment line is appended as `-D <value>` to the Go compile invocation only.
- Use this for targeted parity checks (for example optimizer/fastpath toggle coverage) without changing global harness defaults.

Goal:

- catch semantic drift where generated Go behavior diverges from the Haxe baseline
- keep a focused set of high-signal semantic fixtures:
  - nullability/string behavior
  - deterministic non-string null-equality behavior (`== null` / `!= null` on typed references)
  - deterministic `go.Chan<T>` non-blocking behavior (`trySend`, `recvOr`, `tryRecv`)
  - deterministic `go.Result` success/failure semantics (`isOk`/`isErr`/`unwrap`/`error`, `Go.ok`/`Go.fail`)
  - deterministic portable+auto fallback parity for ineligible typed `go.Slice`/`go.Map`/`go.Result` specializations
  - exceptions
  - enum/switch bindings
  - virtual dispatch
  - numeric edge behavior
  - nullable reference fields
  - selected sys/io behavior
  - deterministic `sys.io.File` save/get parity behavior
  - deterministic `sys.FileSystem` create/read/rename/delete/stat-size parity behavior
  - deterministic `haxe.ds` map/list core operation parity (`set/get/exists/remove`, `add/push/pop/first/last/length`, missing-key/empty typed-null reads)
  - deterministic `haxe.ds.ReadOnlyArray` length/index/read-only-view parity subset
  - deterministic `sys.io.Process` spawn/stdout behavior (cross-platform echo and `haxe --version` smoke)
  - deterministic throw-as-expression parity (`throw` in value-returning expression contexts)
  - deterministic try/catch return-forwarding parity for statement-form `try` branches that `return` from enclosing functions
  - deterministic `sys.net.Host` constructor/resolve/reverse/localhost behavior
  - deterministic `haxe.io.Bytes`/`haxe.io.BytesBuffer` semantics (`set`/`addByte` byte normalization, plus `blit`/`fill`/`sub`/`compare` behavior)
  - deterministic `haxe.io.Bytes.getData` / `haxe.io.Bytes.ofData` alias semantics
  - deterministic `haxe.io.Bytes.toHex` / `haxe.io.Bytes.ofHex` semantics (including odd-digit error path)
  - deterministic `haxe.io.BytesInput`/`haxe.io.BytesOutput` stream + inherited helper subset semantics (`position`/`length`, EOF, bounds, read/write contracts, `readLine`/`readAll`/`readUntil`/`readFullBytes`, typed numeric read/write helpers, `writeInput`, and `readLine` EOF/tail/CRLF edge behavior)
  - deterministic `haxe.io.Encoding` constructor and `haxe.io.Bytes` string conversion semantics (`UTF8`/`RawNative`, `Bytes.ofString`, `Bytes.getString`, and IO read/write string paths)
  - deterministic `haxe.Int64` arithmetic/compare/parse/fromFloat/toInt behavior parity
  - deterministic `haxe.Int32` overflow/bitwise/shift/ucompare operator behavior parity
  - deterministic `Std.isOfType` behavior parity for typed and Dynamic class/array/enum/null checks
  - deterministic `Std.isOfType` behavior for unresolved `@:runtimeValue @:coreType` abstract targets (no hard-fail path)
  - deterministic type-value expression parity for class/enum refs (`TTypeExpr`)
  - deterministic `Type` reflection subset parity (`getClass`, `getSuperClass`, `getClassFields`, `getInstanceFields`, `resolveClass`, `createInstance`, `createEmptyInstance`, `getEnum`, `getEnumConstructs`, `resolveEnum`, `allEnums`, enum constructor/index/parameters/equality)
  - deterministic core `Array`/`String`/`IntIterator` subset parity (`push`/`pop` statement-form, length/index/iteration, `0...N` range iteration, `String.length`/`charAt`/`charCodeAt` null-on-out-of-range/`substring`/`fromCharCode`)
  - deterministic `StringBuf`/`DateTools`/`Lambda` subset parity (`add`/`addChar`/`addSub`, `DateTools.delta` with duration helpers and `%Y-%m-%d %H:%M:%S` formatting, `Lambda.filter`/`map`/`fold`/`has`/`exists`/`count`/`empty` over `Array<T>` and `haxe.ds.List<T>`)
  - deterministic `haxe.PosInfos` default-argument injection behavior
  - deterministic `haxe.PosInfos.customParams` missing-field/null-access behavior
  - deterministic HTTP behavior (`requestUrl`, `customRequest`, `customRequest` socket transport, callback/status/error flows, `Http.PROXY`) without external network
  - deterministic `sys.net.Socket` loopback and advanced method parity (`bind/listen/connect/accept/read/write/close`, `setTimeout`, `waitForRead`, `setBlocking`, `setFastSend`, `select`, `shutdown`)
  - reflection compare + dynamic field semantics
  - anonymous object literal/field mutation semantics
  - deterministic `haxe.Json` parse/stringify + `JsonParser`/`JsonPrinter` behavior
  - crypto/xml/zip behavior parity
  - serializer/unserializer roundtrip semantics
  - serializer wire-format and sequential cursor semantics
  - serializer date/bytes token semantics
  - serializer class/enum token semantics
  - serializer extended token families (`l/b/q/M/j/x/A/B`)
  - serializer custom token + resolver semantics (`C`, `setResolver`, default/null resolver paths)
  - serializer resolver type-value semantics (`Class<T>`/`Enum<T>` returns from resolver methods)
  - serializer cache/reference graph semantics (`r` parity for repeated enum/class/custom instances and cycles)
  - serializer global default flag semantics (`Serializer.USE_CACHE`, `Serializer.USE_ENUM_INDEX`) and `serializeException` interaction
  - serializer resolver polymorphism semantics (method-shape variants and dynamic/object resolver invocation paths)
  - serializer mixed reference stress semantics (`R`/`r` interleaving and sequential cache replay)
  - EReg match/split/replace contract semantics
  - EReg edge semantics (flags, matched-group errors, global vs non-global replacement/map)
