# Concurrency Contract

This document defines the current concurrency boundaries for generated Go. It
covers portable `sys.thread`, explicit Go-native `go.*` APIs, and the runtime
mechanisms that keep those two products distinct.

## Product boundary

`sys.thread` is the portable Haxe concurrency surface. Its workers are
foreground workers: generated `main` waits for them, including nested workers.
An uncaught Haxe throw ends only that worker after a stable stderr report.

`go.Go.spawn`, `go.Chan<T>`, and `go.Select` are explicit Go-native APIs. A
spawned goroutine is not joined when `main` returns, and a native panic is not a
Haxe catch value. These semantics are API-scoped; choosing the `metal`
compatibility preset is neither necessary nor sufficient to opt into them.

## Portable runtime invariants

| Surface | Contract |
| --- | --- |
| `FixedThreadPool.run` versus `shutdown` | One pool mutex is the linearization point. Returning from `run` means its task was queued before every shutdown sentinel; otherwise `run` throws `ThreadPoolException`. Every accepted task executes exactly once, and a Haxe-throwing task installs a replacement worker before its failure is reported. |
| `ElasticThreadPool.run` versus `shutdown` | One pool mutex owns admission, pending/live counts, timeout retirement, task completion, and failed-worker replacement. Each accepted task has one queue item and one counted wake token. Shutdown rejects later tasks, drains accepted tasks, then consumes one exit token per live worker. |
| `Condition.signal` | Targets the oldest currently unsignaled waiter. Repeated signals do not accumulate credits for future waiters. |
| `Condition.broadcast` | Closes only the per-waiter wake channels registered in that broadcast generation. A later waiter cannot consume an earlier broadcast. |
| Event-loop ownership | A per-thread ownership lock protects loop install, lookup, and removal. Callback execution does not hold that lock. |
| Repeating-event cancellation | Queued cancellation removes the event without a tombstone. Only an executing callback can hold a temporary cancellation marker, which is removed when progress decides whether to reschedule. Unknown IDs are bounded no-ops. |
| Timed event-loop waits | An insertion, cancellation, or timeout ends one wait generation. The loop then recomputes the full schedule, so an earlier timer is not delayed to the old deadline and cancelling the last timer cannot strand the worker. |
| Portable thread identity state | Spawned portable workers register before executing and remove both their logical state and goroutine mapping on normal/Haxe-throw completion. Native panics remain fatal. |

The elastic pool's public `maxThreadsCount` remains a plain writable field
because that is Haxe's core API shape. Concurrently mutating that field while
calling pool methods requires synchronization by the application. Normal
`run`, `shutdown`, timeout, and read-only pool-property concurrency is owned by
the implementation.

## Goroutine identity dependency

Go does not expose a supported goroutine-local identifier. Re-entrant Haxe
`Mutex`, `Condition`, `Tls`, and `Thread.current()` nevertheless require stable
caller identity, so `hxrt` uses a deliberately bounded fallback:

1. Capture only the first 64 bytes from `runtime.Stack` for the current
   goroutine.
2. Parse only `goroutine <positive decimal id> `.
3. Reuse fixed-size buffers through `sync.Pool`; no complete stack is retained.
4. Fail with a native runtime panic if the header format is invalid. Identity
   zero is reserved for “unowned,” so silently returning zero would corrupt
   mutex ownership.
5. Remove mappings for every supported portable worker lifecycle.

The direct regression test drains 10,000 portable workers in batches and
requires identity maps to return to their original size after every batch. On
2026-07-14, an Apple M2 Pro / Darwin arm64 benchmark measured
`BenchmarkCurrentGoroutineID` at 1.72–1.77 microseconds per operation with
`0 B/op` and `0 allocs/op`. This is characterization evidence, not a
hardware-dependent release threshold.

Bare goroutines, including those created by `go.Go.spawn`, are not portable
foreground threads. `go.Go.spawn` deliberately owns only native non-joined
shutdown and panic behavior; the runtime does not promise portable lifecycle
discovery for arbitrary user-created Go goroutines.

`sys.thread.Tls` is not part of the admitted portable concurrency surface yet.
Its get/set/clear behavior and synchronization are tested, but the current
per-instance ID map does not prove value reclamation when a thread exits, and a
`go.Go.spawn` callback that asks for `Thread.current()` has no explicit detached
identity cleanup transition. [`haxe_go-vfp.10.7`](compatibility-support-matrix.md#known-blockers)
owns that narrower lifecycle design; the admitted `Thread`, synchronization,
event-loop, and pool contracts do not depend on claiming it complete.

## Go-native channel lifecycle

`go.Chan<T>` preserves native Go ownership and panic rules. Buffered values are
received before closed state becomes visible.

| Operation | Open, ready | Open, not ready / nil | Closed and drained |
| --- | --- | --- | --- |
| `send(value)` | Sends, blocking if required | Blocks; a nil channel blocks forever | Native `send on closed channel` panic |
| `trySend(value)` | `true` | `false` | Native `send on closed channel` panic |
| `recv()` | Returns a value | Blocks; a nil channel blocks forever | Returns `T`'s Go zero value |
| `tryRecv()` | Successful `go.Result<T>` | Error result with `"empty"` | Error result with `"closed"` |
| `recvOr(default)` | Returns the immediately ready value | Returns `default` | Returns `default` |
| `close()` | Closes producer side | Closing nil panics | Closing again panics |

Channel close is producer-owned. Applications must synchronize send and close;
a send/close race is not made safe by `Chan<T>`. Native channel panics continue
unwinding as Go panics and are not accepted by Haxe `try`/`catch`.

`go.Select` multi-branch helpers remain deterministic priority polling, not
Go's pseudo-random selection among simultaneously ready native cases. That
separate limitation remains part of the Go-native capability roadmap.

## Evidence and no-retry policy

The release tooling runs these contracts without flaky-test retries:

- direct `hxrt` normal and race tests, including condition generations,
  cancellation boundedness, event-loop ownership and timed-wakeup
  recalculation, and identity churn;
- generated fixed/elastic pool tests with 10,000 concurrent submissions,
  repeated shutdown at `GOMAXPROCS=1,2,8`, and replacement after a task's Haxe
  throw;
- generated generic and specialized channel tests for buffered close, nil
  channels, closed/empty comma-ok behavior, send-after-close, and double-close;
- snapshot and runtime contracts in `stdlib/sys_thread_runtime_direct`,
  `go_native/channel_try_recv`, and `core/ast_select_stmt_printer`.

Run the release-facing race/static-analysis lane with:

```bash
npm run security:go-tooling
```

See also [the Go concurrency and interop guide](go-concurrency-interop-guide.md),
[the `hxrt` runtime guide](hxrt-runtime.md), and
[the compatibility support matrix](compatibility-support-matrix.md). The
`thinking:xhigh` closure verdict and adjudicated findings are recorded in the
[written second-pass review](reviews/concurrency-audit-vfp-10.2.md).
