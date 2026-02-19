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
  - deterministic HTTP behavior (`requestUrl`, `customRequest`, `Http.PROXY`) without external network
  - deterministic `sys.net.Socket` loopback and advanced method parity (`bind/listen/connect/accept/read/write/close`, `setTimeout`, `waitForRead`, `setBlocking`, `setFastSend`, `select`, `shutdown`)
  - reflection compare + dynamic field semantics
  - anonymous object literal/field mutation semantics
  - crypto/xml/zip behavior parity
  - serializer/unserializer roundtrip semantics
  - serializer wire-format and sequential cursor semantics
  - serializer date/bytes token semantics
  - serializer class/enum token semantics
  - serializer extended token families (`l/b/q/M/j/x/A/B`)
  - serializer custom token + resolver semantics (`C`, `setResolver`, default/null resolver paths)
  - serializer cache/reference graph semantics (`r` parity for repeated enum/class/custom instances and cycles)
  - serializer resolver polymorphism semantics (method-shape variants and dynamic/object resolver invocation paths)
  - serializer mixed reference stress semantics (`R`/`r` interleaving and sequential cache replay)
  - EReg match/split/replace contract semantics
  - EReg edge semantics (flags, matched-group errors, global vs non-global replacement/map)
