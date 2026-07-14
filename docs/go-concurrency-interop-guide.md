# Go Concurrency + Interop Guide

Practical patterns for using `reflaxe.go` as a production-grade Go target
without dropping to raw `__go__`.

Quick context:

- `portable` and `metal` are compatible policy presets. See
  `docs/native-policy-presets.md`.
- `go.*` APIs (`go.Go`, `go.Chan`, `go.Select`) are Go-native facades. They are powerful, but they are not portability-safe across non-Go targets.
- `interop` means calling Go packages through typed extern metadata (`@:go.import`, `@:go.name`, `@:go.receiver`) instead of raw string injection.

`metal` is not required for good Go output. Start with `portable`; use typed
Go APIs and `@:goNative` when a module intentionally owns Go semantics.

Reference glossary: `docs/glossary.md`.

## 1) Worker pool with channels

Reference app: `examples/worker_pool_select`.

Core pattern:

- `go.Go.newChan<T>(buffer)` for buffered queues.
- `go.Go.spawn(fn)` for worker fan-out.
- `go.Select.recv/recv2` for typed receive-side branching.
- `go.Select.send/send2` for typed non-blocking send-side branching.
- `Chan.recvOr(default)` + sentinel values for deterministic loop shutdown loops.
- `Chan.trySend(value)` / `Chan.tryRecv()` for low-level non-blocking operations.

Channel lifecycle is explicit:

- `tryRecv()` returns an error result with `"empty"` when no value is ready and
  `"closed"` after a closed channel is drained.
- `recvOr(default)` returns the default for both empty/nil and drained-closed
  channels; `recv()` retains Go's zero-value-after-close behavior.
- sending (including `trySend`) after close and closing a nil/already-closed
  channel are native Go panics, not Haxe catch values.
- close is producer-owned; synchronize send and close in application code.

The full operation table is in [the concurrency contract](concurrency-contract.md).

Why this is recommended:

- Maps to real goroutine/channel/select behavior in generated Go output.
- Still runs with either compatibility preset from one codebase, so preset
  defaults can be compared without changing the API-scoped native contract.

This is an explicit Go-native lifecycle contract. `go.Go.spawn(fn)` emits a
bare goroutine: returning from `main` does not wait for it, and a panic inside
it is a normal fatal Go panic. Coordinate completion through channels or other
typed Go synchronization when work must finish before shutdown.

For portable Haxe thread semantics, use `sys.thread.Thread.create` or
`createWithEventLoop`. Those workers are foreground threads, meaning generated
`main` waits for them (including nested portable workers). An uncaught Haxe
throw is reported on stderr and ends only that worker. Native panics are never
converted into Haxe catch values in either model.

## 2) Typed interop wrappers + user externs

Reference app: `examples/interop_smoke`.

Note: this reference is intentionally preset-neutral for interop parity. For
visible specialization-policy deltas, use `examples/worker_pool_select` and the
`go_native` variants in `examples/pulseforge` / `examples/fluxproxy`.

The reference app (`examples/interop_smoke`) demonstrates both:

- framework-owned wrapper surfaces from `std/go/*`, and
- app-level extern metadata declarations for user-owned package bindings.

Shared package surfaces under both presets:

- `fmt.Println`
- `time.Now` + receiver/static `Unix`
- `context.Background`
- `net/http.StatusText`

Compiler metadata support (`@:go.import`, `@:go.name`, `@:go.receiver`) is covered in
`test/snapshot/go_native/extern_metadata_mapping`.

`(T,error)` metadata support (`@:go.valueError` -> `go.Result<T>`) is covered in:

- `test/snapshot/go_native/extern_value_error_result`
- `test/semantic_diff/go_value_error_result_contract`

## 3) Native-policy guidance

These typed low-level APIs are already contract-covered under either preset:

- `go.Chan<T>`
- `go.Slice<T>`
- `go.Map<K,V>`
- `go.Result<T>`

Boundary rule:

1. Keep semantic baseline in `portable`.
2. Declare Go-owned modules with `@:goNative` and typed APIs.
3. Select eager/error/strict policy only when the build needs those guarantees.
4. Require benchmark evidence for performance-driven lowering work.

Practical interpretation:

- Start in `portable` when cross-target compatibility and predictable Haxe semantics are primary.
- Use the `metal` compatibility preset when its bundled defaults are convenient;
  it does not itself grant a different semantic product.

## 4) Caveats (important)

- `go.Select` multi-branch helpers are deterministic and branch-priority ordered; they do not mirror Go runtime pseudo-random ready-case selection.
- `go.Go.spawn` goroutines are not joined automatically. Use an explicit channel,
  wait-group façade, or another typed Go synchronization boundary when shutdown
  must wait for them.
- Haxe `try`/`catch` catches only values raised with Haxe `throw`. A panic from a
  typed Go extern remains native and fatal unless user-owned Go code explicitly
  recovers it inside the same goroutine.
- Complex Go extern signatures may need façade wrappers until broader mapping support lands.
- `sys.thread.ElasticThreadPool.maxThreadsCount` is a core-API writable field;
  synchronize application code if it mutates that field concurrently with pool use.
- `sys.thread.Tls` lifecycle reclamation is still experimental; do not use it
  for unbounded short-lived-thread churn until `haxe_go-vfp.10.7` closes.
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
- Go users still keep direct access to native `go.*` surfaces when needed.
- semantic-diff tests can validate portable behavior independently from native optimizations.
