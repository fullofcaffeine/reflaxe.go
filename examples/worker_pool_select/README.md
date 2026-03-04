# worker_pool_select

Deterministic worker-pool and select-style channel example across `portable` and `metal`.

## Why this example exists

- Demonstrates Go concurrency abstractions in a compact deterministic workload.
- Shows how one Haxe codebase can target both profile contracts.
- Serves as a high-signal reference for typed `go.Select` helper usage.

## What it demonstrates

- `go.Go.spawn` worker fan-out from one Haxe codebase.
- `go.Chan<T>` buffered queues and deterministic work collection.
- Typed `go.Select` helpers (`recv`, `recv2`, `send`, `send2`) layered on select-backed non-blocking channel operations.
- Same source compiled to all profiles with stable output.

## Portable vs metal diff in this app

- `portable`: semantic baseline and portability-first defaults.
- `metal`: explicit Go-first/perf lane; typed specialization is prioritized here.
- Both profiles preserve the same workload contract and output expectations.

## When to choose each profile here

- Choose `portable` when this worker pattern must remain portable with shared domain code.
- Choose `metal` when this worker path is a Go hot path and you want stricter metal policy plus stronger typed specialization pressure.

## Tradeoffs shown by this example

- Profile outputs can stay behavior-equivalent while generated helper shape differs.
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

- `docs/profiles.md`
- `docs/profile-semantics-guide.md`
- `docs/go-concurrency-interop-guide.md`
