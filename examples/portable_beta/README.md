# Portable beta

This is the smallest example that carries a portable-beta release claim. It is
deliberately narrower than the larger application examples.

## What it proves

The program runs ordinary Haxe through the custom Go backend, formats and tests
the generated Go module, and executes it. Its checks cover named, currently
admitted operations from these groups:

- control flow, classes, enums, closures, and typed exception messages;
- core arrays, strings, and `StringTools`;
- JSON, bytes, `StringMap`, and paths;
- reflection used to inspect parsed JSON;
- file reads/writes and an application-controlled temporary-file lifecycle.

The exact operation IDs live with this example in
[`examples/qa-manifest.json`](../qa-manifest.json). The harness resolves every
ID against
[`docs/compatibility-support-source.json`](../../docs/compatibility-support-source.json)
and refuses to run a release claim when an operation is missing or excluded.

## Why the program is quiet

Successful execution writes nothing to standard output. Each behavior has a
small, manually authored invariant in `Main.hx`; a mismatch throws and makes
`go run` fail. This avoids pretending that unrelated console APIs are part of
the narrow claim. The examples harness reports the visible `PASS` result.

## Run it

```bash
python3 test/run-examples.py --example portable_beta --profile portable
```

That one command performs the real Haxe → Go path, `gofmt`, `go test ./...`,
`go run .`, the empty-output comparison, and generated-tree drift checking.

This example does not claim Go-native APIs, the `metal` preset, HTTP, DNS,
public networking, another operating system, or any compatibility operation
not listed in the QA manifest.
