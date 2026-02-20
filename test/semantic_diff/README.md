# Semantic Differential Fixtures

These fixtures compare runtime output between:

1. Haxe reference execution (`--interp`)
2. `reflaxe.go` output (`portable` profile, then `go run .`)

Runner:

```bash
python3 test/run-semantic-diff.py
```

Goal:

- catch semantic drift where generated Go behavior diverges from the Haxe baseline
- keep a focused set of high-signal semantic fixtures:
  - nullability/string behavior
  - exceptions
  - enum/switch bindings
  - virtual dispatch
  - numeric edge behavior
  - nullable reference fields
  - selected sys/io behavior
  - deterministic `sys.io.File` save/get parity behavior
  - deterministic `sys.FileSystem` create/read/rename/delete/stat-size parity behavior
  - deterministic `haxe.ds` map/list core operation parity (`set/get/exists/remove`, `add/push/pop/first/last/length`, missing-key/empty typed-null reads)
  - deterministic `sys.io.Process` spawn/stdout behavior (cross-platform echo and `haxe --version` smoke)
  - deterministic throw-as-expression parity (`throw` in value-returning expression contexts)
  - deterministic `sys.net.Host` constructor/resolve/reverse/localhost behavior
  - deterministic `haxe.Int64` arithmetic/compare/parse/fromFloat/toInt behavior parity
  - deterministic `haxe.Int32` overflow/bitwise/shift/ucompare operator behavior parity
  - deterministic `Std.isOfType` behavior parity for typed and Dynamic class/array/enum/null checks
  - deterministic `Std.isOfType` behavior for unresolved `@:runtimeValue @:coreType` abstract targets (no hard-fail path)
  - deterministic type-value expression parity for class/enum refs (`TTypeExpr`)
  - deterministic `haxe.PosInfos` default-argument injection behavior
  - deterministic `haxe.PosInfos.customParams` missing-field/null-access behavior
  - deterministic HTTP behavior (`requestUrl`, `customRequest`, `Http.PROXY`) without external network
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
  - serializer resolver polymorphism semantics (method-shape variants and dynamic/object resolver invocation paths)
  - serializer mixed reference stress semantics (`R`/`r` interleaving and sequential cache replay)
  - EReg match/split/replace contract semantics
  - EReg edge semantics (flags, matched-group errors, global vs non-global replacement/map)
