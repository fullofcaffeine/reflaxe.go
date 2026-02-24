# Go Concurrency + Interop Guide

Practical patterns for using `reflaxe.go` as a production-grade Go-target lane without dropping to raw `__go__`.

## 1) Worker pool with channels

Reference app: `examples/worker_pool_select`.

Core pattern:

- `go.Go.newChan<T>(buffer)` for buffered queues.
- `go.Go.spawn(fn)` for worker fan-out.
- `Chan.recvOr(default)` + sentinel values for deterministic loop shutdown.
- `Chan.trySend(value)` for non-blocking enqueue/backpressure checks.

Why this is recommended:

- Maps to real goroutine/channel/select behavior in generated Go output.
- Still runs in `portable`/`gopher`/`metal` profile matrix from one codebase.

## 2) Typed interop externs

Reference app: `examples/interop_smoke`.

Interop metadata:

- `@:go.import("pkg/path")`
- `@:go.name("SymbolName")`
- `@:go.receiver` for receiver-style static wrappers

Example surfaces in the reference app:

- `fmt.Println`
- `time.Now` + receiver/static `Unix`
- `context.Background`
- `net/http.StatusText`

## 3) Metal subset guidance

Use `metal` when you need typed low-level lanes that are already contract-covered:

- `go.Chan<T>`
- `go.Slice<T>`
- `go.Map<K,V>`
- `go.Result<T>`

Promotion rule:

1. Keep semantic baseline in `portable`/`gopher`.
2. Move hot paths to `metal` only with benchmark evidence.
3. Keep strict boundary enforcement enabled (`reflaxe_go_strict`).

## 4) Caveats (important)

- Non-blocking channel operations are currently exposed via `trySend`/`recvOr` helpers; there is no finalized general-purpose Haxe `select` expression API yet.
- Complex Go extern signatures may need façade wrappers until broader mapping lanes land.
- For current limitations and planning guidance, use `docs/known-gaps.md`.
