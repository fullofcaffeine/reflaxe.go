# worker_pool_select

Deterministic worker-pool and select-style channel example across `portable`, `gopher`, and `metal`.

## What it demonstrates

- `go.Go.spawn` worker fan-out from one Haxe codebase.
- `go.Chan<T>` buffered queues and deterministic work collection.
- Select-backed non-blocking channel operations (`trySend`, `recvOr`) that map to `select { ... default: }` on Go output.
- Same source compiled to all profiles with stable output.

## Compile

```bash
haxe compile.portable.hxml
haxe compile.gopher.hxml
haxe compile.metal.hxml
```

## Run

```bash
(cd out_portable && go run .)
(cd out_gopher && go run .)
(cd out_metal && go run .)
```

Expected output:

```text
worker.count=4
select.trySend=true,false
select.recvOr=5,99
```
