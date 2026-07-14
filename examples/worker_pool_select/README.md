# worker_pool_select

Deterministic worker-pool and select-style channel example across the
`portable` and `metal` compatibility presets.

## Why this example exists

- Demonstrates Go concurrency abstractions in a compact deterministic workload.
- Shows one explicit Go-native API contract under both policy presets.
- Serves as a high-signal reference for typed `go.Select` helper usage.

## What it demonstrates

- `go.Go.spawn` worker fan-out from one Haxe codebase.
- `go.Chan<T>` buffered queues and deterministic work collection.
- Typed `go.Select` helpers (`recv`, `recv2`, `send`, `send2`) layered on select-backed non-blocking channel operations.
- Same source compiled with both presets and stable output.

## Portable vs metal diff in this app

- `portable`: guarded/proven/allow defaults. The `go.*` calls themselves state
  that this example is Go-native.
- `metal`: explicit/eager/error/strict compatibility defaults.
- Both presets preserve the same typed Go API contract and output expectations.

`metal` is not required for good Go output or for channels/select. This example
keeps a metal preset run to test the compatibility bundle against the same
source.

## When to choose each preset here

- Choose `portable` as the default preset and isolate this Go-specific adapter
  from portable domain code.
- Choose `metal` when its eager/error/strict bundle is convenient, or select
  those policies individually under portable.

## Tradeoffs shown by this example

- Preset outputs can stay behavior-equivalent while generated helper shape differs.
- `metal` is not guaranteed to be fewer lines; it may emit extra typed helpers for hot-path stability/perf.

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

Expected output:

```text
worker.count=4
select.trySend=true,false
select.recvOr=5,99
select.recv2=right:right
select.send2=a values=11,-1
```

## Generated Go diff inspection

```bash
diff -ru generated/portable generated/metal
```

High-signal files:

- `generated/portable/main.go`
- `generated/metal/main.go`
- `generated/portable/module_go_select.go`
- `generated/metal/module_go_select.go`

## Related docs

- [`docs/profiles.md`](../../docs/profiles.md)
- [`docs/native-policy-presets.md`](../../docs/native-policy-presets.md)
- [`docs/profile-semantics-guide.md`](../../docs/profile-semantics-guide.md)
- [`docs/go-concurrency-interop-guide.md`](../../docs/go-concurrency-interop-guide.md)
