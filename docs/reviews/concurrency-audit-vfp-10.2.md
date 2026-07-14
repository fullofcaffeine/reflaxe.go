# `haxe_go-vfp.10.2` Written Second-Pass Review

Date: 2026-07-14

Verdict: **approve the bounded concurrency closure with explicit exclusions**.
Portable synchronization, foreground threads, event loops, and fixed/elastic
pools have enough deterministic semantic and supported-Go tooling evidence for
the named pre-1.0 beta surface. `sys.thread.Tls` lifecycle reclamation and the
complete Go-native concurrency API remain outside that admission.

## Review provenance

This is the explicit written second-pass fallback permitted for a
`thinking:xhigh` bead. It is a local engineering review of the implementation,
generated output, tests, and compatibility claims; it is not an Oracle result
and does not claim independent-model provenance.

The requested GPT-5.6 Pro Oracle path was unavailable because the Codex account
had reached its usage limit, with a reported reset of July 19, 2026 at 12:58 PM.
The installed fallback was not used because it mapped the requested model name
to older `gpt-5-pro`, which would make the provenance misleading. Local tracing
resolved the concurrency designs to one defensible implementation in each
case, so another deep model pass is not required to land this bounded audit.

## Review method

The pass traced every mutable field to its owner and linearization point, then
checked the generated Go rather than relying only on Haxe source shape. It also
compared relevant behavior with Haxe 4.3.7's `EventLoop` and
`FixedThreadPool`, exercised the runtime under Go's race detector, and reviewed
the compatibility source rather than inferring admission from implementation.

The following invariants were reviewed:

| Area | Required invariant | Verdict |
| --- | --- | --- |
| Fixed pool | `run` either queues before shutdown sentinels or rejects; a throwing task cannot strand later accepted work | Pass |
| Elastic pool | One mutex owns admission, shutdown, pending/live counts, timeout retirement, and worker replacement | Pass |
| Conditions | Signals target existing waiters only; broadcasts cannot leak credits to a later generation | Pass |
| Event loop | Ownership is race-free; cancellation state is bounded; insert/cancel wakes force deadline recomputation | Pass |
| Channels | Empty and drained-closed receives are distinct; native send/close panics remain native | Pass for `go.Chan<T>`; broader native select remains excluded |
| Identity | The `runtime.Stack` dependency is fixed-size, strict, zero-allocation in steady state, and portable-worker mappings return to baseline | Pass for admitted portable workers |
| Product boundary | Concurrency semantics are selected by `sys.thread` versus explicit `go.*` APIs, not by the `portable`/`metal` preset name | Pass |

## Findings and disposition

### C-01: a timed event-loop wait retained its stale deadline

Severity: P0 for the admitted event-loop surface.

After a cancellation or earlier timer insertion signaled the condition, the
wait loop checked only whether work was ready at that instant and went back to
sleep until the old timer. Cancelling the last event could therefore strand a
worker, and an earlier timer could run late.

Disposition: fixed. A wait now ends after one scheduler-state transition and
lets `ThreadEventLoopLoop` recompute the schedule. Deterministic tests replace
the condition locker with a notification wrapper, proving that the tested
goroutine has entered `sync.Cond.Wait` before cancellation/insertion occurs.

### C-02: a fixed-pool worker was not replaced after a Haxe throw

Severity: P0 for the accepted-work contract.

The staged worker caught only the shutdown sentinel. With one worker, a task
that threw could terminate that worker and leave later accepted tasks queued
forever. Haxe 4.3.7's implementation starts a replacement before preserving
the original throw.

Disposition: fixed without shared worker-generation state. The old worker calls
`Thread.create(loop)` before rethrowing the original Haxe value. A generated
one-worker regression fails on the old code and now proves that both fixed and
elastic pools run the following accepted task exactly once.

### C-03: TLS values and detached identities lack an exit reclamation contract

Severity: P1, non-blocking only because it is excluded from release admission.

`sys.thread.Tls` stores values in a per-instance `IntMap` keyed by logical
thread ID. Ordinary get/set/clear is synchronized, but completed-thread value
reclamation is not proved. A `go.Go.spawn` callback that requests
`Thread.current()` likewise has no explicit detach transition.

Disposition: deferred honestly to `haxe_go-vfp.10.7`. The compatibility matrix
splits `sys.thread.Tls` from the admitted primitive group and marks it
experimental. The follow-up requires lifecycle-owned storage, no generation
contamination, race evidence, and preservation of native non-joined/panic
semantics.

## Evidence reviewed

- `npm test`: 248/248 snapshots, including runtime execution where configured.
- `npm run test:semantic-diff`: 129/129 interpreter-versus-Go contracts; the
  changed thread contract was also rerun after the fixed-worker repair.
- `npm run test:stdlib-sweep:go-test`: 55/55 strict upstream modules.
- `npm run test:examples`: 12/12 compile, `go test`, run, and stdout contracts
  after regenerating portable/metal example trees.
- `npm run security:go-tooling`: seven scopes times race, checkptr, vet, and
  Staticcheck `SA*`; all 28 gates passed on Go 1.25.6 with retries disabled.
- Generated pool stress: 10,000 submissions, 64 submitters, eight concurrent
  shutdown callers, `GOMAXPROCS=1,2,8`, exact accepted/executed correspondence,
  plus Haxe-throw worker replacement.
- Direct event-loop/condition/identity tests: 100,000 cancellations, real
  signal/broadcast generations, deterministic timed-wakeup tests, concurrent
  ownership reads, and 10,000 portable-worker lifecycle drains.
- `BenchmarkCurrentGoroutineID`: 1.72–1.77 microseconds/op, `0 B/op`,
  `0 allocs/op` on Apple M2 Pro / Darwin arm64. This is characterization, not a
  release threshold.

## Residual boundaries

- `sys.thread.Tls` remains experimental under `haxe_go-vfp.10.7`.
- Concurrent mutation of `ElasticThreadPool.maxThreadsCount` requires caller
  synchronization because the writable core-API field cannot provide an
  intercepting setter without changing the public contract.
- Arbitrary foreign goroutines are not discoverable portable threads.
- Directional channels, range syntax, and true native multi-case select remain
  governed by `haxe_go-vfp.9.1`.

No unresolved finding blocks closure of `haxe_go-vfp.10.2` under these stated
boundaries. This verdict does not admit stable 1.x, networking, untrusted-input,
release-provenance, or the excluded TLS/native-concurrency surfaces.
