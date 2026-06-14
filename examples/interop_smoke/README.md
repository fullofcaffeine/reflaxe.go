# interop_smoke

Minimal typed-interop reference example for reflaxe.go.

## Why this example exists

- Demonstrates typed interop metadata in real Haxe code.
- Keeps the scenario intentionally small so interop annotations are easy to audit.
- Proves that portable-surface interop wrappers can stay profile-consistent.

## What it demonstrates

- Framework-owned interop wrappers from `std/go/*` (`go.Fmt`, `go.Time`, `go.ContextPkg`, `go.Http`).
- User-level extern metadata in app code (`UserGoTime`, `UserGoContextPkg`, `UserGoHttp`) using:
  - `@:go.import("pkg/path")`
  - `@:go.name("SymbolName")`
  - `@:go.receiver`
- One Haxe codebase compiled across `portable` and `metal`.

Compiler metadata coverage is additionally locked by snapshot fixtures
(`test/snapshot/go_native/extern_metadata_mapping`).

## Portable vs metal diff in this app

This example is expected to generate near-identical Go for `portable` and `metal`
because it only uses profile-safe interop calls.
The profile matrix here validates contract consistency, not profile-divergent code shape.

## When to choose each profile here

- Choose `portable` when this interop adapter is meant to remain cross-target friendly at the Haxe contract level.
- Choose `metal` when this adapter is part of a Go-native lane and you want strict metal boundary defaults around surrounding modules.

## Tradeoffs shown by this example

- Near-identical generated Go across profiles is expected for portable-surface interop.
- This example is not the right place to inspect large profile code-shape divergence.
- Use `examples/worker_pool_select` and `go_native` lanes in `examples/pulseforge` / `examples/fluxproxy` for visible profile divergence.

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

Expected output for every profile:

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
- [`docs/profile-semantics-guide.md`](../../docs/profile-semantics-guide.md)
- [`docs/go-concurrency-interop-guide.md`](../../docs/go-concurrency-interop-guide.md)
