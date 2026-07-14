# interop_smoke

Minimal typed-interop reference example for reflaxe.go.

## Why this example exists

- Demonstrates typed interop metadata in real Haxe code.
- Keeps the scenario intentionally small so interop annotations are easy to audit.
- Proves that typed Go interop wrappers keep one API contract across presets.

## What it demonstrates

- Framework-owned interop wrappers from `std/go/*` (`go.Fmt`, `go.Time`, `go.ContextPkg`, `go.Http`).
- User-level extern metadata in app code (`UserGoTime`, `UserGoContextPkg`, `UserGoHttp`) using:
  - `@:go.import("pkg/path")`
  - `@:go.name("SymbolName")`
  - `@:go.receiver`
- One Haxe codebase compiled across `portable` and `metal`.

Compiler metadata coverage is additionally locked by snapshot fixtures
(`test/snapshot/go_native/extern_metadata_mapping`).

## Portable vs metal preset diff in this app

This example is expected to generate near-identical Go for `portable` and `metal`
because the typed APIs already carry the Go-native contract.
The matrix validates API consistency across policy defaults, not divergent source semantics.

## When to choose each preset here

- Choose `portable` as the default preset; isolate this target-specific adapter
  behind a portable application interface if cross-target reuse matters.
- Choose `metal` when its strict/eager/fail-fast bundle is convenient, or select
  the relevant axes independently.

## Tradeoffs shown by this example

- Near-identical generated Go across presets is expected for the same typed API.
- This example is not the right place to inspect large policy-driven code-shape
  divergence.
- Use `examples/worker_pool_select` and `go_native` variants in
  `examples/pulseforge` / `examples/fluxproxy` for visible specialization
  policy deltas.

## Compile

```bash
haxe compile.portable.hxml
haxe compile.metal.hxml
```

## Run

```bash
(cd out_portable && go run .)
(cd out_metal && go run .)
```

Expected output for every preset:

```text
1
```

## Generated Go diff inspection

```bash
diff -ru generated/portable generated/metal
```

You should mostly see module-name/path differences (`go.mod`, package path strings), not major structural divergence.

## Related docs

- [`docs/profiles.md`](../../docs/profiles.md)
- [`docs/native-policy-presets.md`](../../docs/native-policy-presets.md)
- [`docs/profile-semantics-guide.md`](../../docs/profile-semantics-guide.md)
- [`docs/go-concurrency-interop-guide.md`](../../docs/go-concurrency-interop-guide.md)
