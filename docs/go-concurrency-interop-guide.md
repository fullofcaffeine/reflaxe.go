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
- Still runs in `portable`/`metal` profile matrix from one codebase.

## 2) Typed interop wrappers + user externs

Reference app: `examples/interop_smoke`.

Note: this reference is intentionally profile-neutral for interop parity. For visible
portable-vs-metal value, use `examples/worker_pool_select` and the `go_native`
lanes in `examples/pulseforge` / `examples/fluxproxy`.

The reference app (`examples/interop_smoke`) demonstrates both:

- framework-owned wrapper surfaces from `std/go/*`, and
- app-level extern metadata declarations for user-owned package bindings.

Shared package surfaces in both lanes:

- `fmt.Println`
- `time.Now` + receiver/static `Unix`
- `context.Background`
- `net/http.StatusText`

Compiler metadata support (`@:go.import`, `@:go.name`, `@:go.receiver`) is covered in
`test/snapshot/go_native/extern_metadata_mapping`.

`(T,error)` metadata support (`@:go.valueError` -> `go.Result<T>`) is covered in:

- `test/snapshot/go_native/extern_value_error_result`
- `test/semantic_diff/go_value_error_result_contract`

## 3) Metal subset guidance

Use `metal` when you need typed low-level lanes that are already contract-covered:

- `go.Chan<T>`
- `go.Slice<T>`
- `go.Map<K,V>`
- `go.Result<T>`

Promotion rule:

1. Keep semantic baseline in `portable`.
2. Move hot paths to `metal` only with benchmark evidence.
3. Keep strict boundary enforcement enabled (`reflaxe_go_strict`).

## 4) Caveats (important)

- `go.Select` multi-branch helpers are deterministic and branch-priority ordered; they do not mirror Go runtime pseudo-random ready-case selection.
- Complex Go extern signatures may need façade wrappers until broader mapping lanes land.
- For current limitations and planning guidance, use `docs/known-gaps.md`.

## 5) Planned portable channel facade

Goal: support channel-style concurrency in portable code without exposing `go.*` as the portable contract surface.

Do not make `go.Chan` itself portable.
`go.Chan<T>` remains a Go-native facade API.

Instead, introduce a portable abstraction layer (planned), for example:

- `hx.concurrent.Channel<T>`

Design rule:

1. Define semantics explicitly first (blocking behavior, close behavior, receive-after-close behavior, and select guarantees).
2. Implement target-specific backends for the same contract:
   - Go backend: lower to real goroutines/channels/select.
   - non-Go backends: runtime/scheduler implementations that match the declared contract.
3. Keep `go.*` available as native power APIs outside the portable contract.

Why this shape:

- portable users get a stable cross-target API.
- Go users still keep direct access to native `go.*` lanes when needed.
- semantic-diff tests can validate portable behavior independently from native optimizations.
