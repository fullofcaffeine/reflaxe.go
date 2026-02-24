# Go Concurrency + Interop Guide

Practical patterns for using `reflaxe.go` as a production-grade Go-target lane without dropping to raw `__go__`.

## 1) Worker pool with channels

Reference app: `examples/worker_pool_select`.

Core pattern:

- `go.Go.newChan<T>(buffer)` for buffered queues.
- `go.Go.spawn(fn)` for worker fan-out.
- `go.Select.recv/recv2` for typed receive-side branching.
- `go.Select.send/send2` for typed non-blocking send-side branching.
- `Chan.recvOr(default)` + sentinel values for deterministic loop shutdown loops.
- `Chan.trySend(value)` / `Chan.tryRecv()` for low-level non-blocking operations.

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

- `go.Select` multi-branch helpers are deterministic and branch-priority ordered; they do not mirror Go runtime pseudo-random ready-case selection.
- Complex Go extern signatures may need façade wrappers until broader mapping lanes land.
- For current limitations and planning guidance, use `docs/known-gaps.md`.
