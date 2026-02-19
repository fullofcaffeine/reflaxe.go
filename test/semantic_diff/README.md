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
  - reflection compare + dynamic field semantics
  - anonymous object literal/field mutation semantics
  - crypto/xml/zip behavior parity
  - serializer/unserializer roundtrip semantics
  - EReg match/split/replace contract semantics
