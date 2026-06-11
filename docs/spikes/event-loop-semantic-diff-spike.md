# Event-loop Semantic-diff Spike

This spike records why direct `haxe.EntryPoint`, `haxe.MainLoop`, and `haxe.Timer`
use stays under snapshot/runtime evidence for now instead of semantic-diff evidence.

## Terms

- `semantic-diff`: the harness that runs the same Haxe program through Haxe
  `--interp` and generated Go, then compares stdout. See
  `/docs/semantic-diff-guide.md`.
- snapshot/runtime evidence: a test that proves generated Go compiles and runs a
  stable Go-side contract, without claiming byte-for-byte interpreter parity.
- event loop: a queue that runs callbacks later, often after a thread wakeup or
  timer fires.
- target-sensitive: behavior that depends on the runtime, operating system,
  scheduler, wall clock, stack frames, network, or TLS implementation.

## What was tested

A narrow `haxe.EntryPoint` probe queued two callbacks and then called
`EntryPoint.run()`:

```haxe
import haxe.EntryPoint;

class Main {
	static function main() {
		var log = new Array<String>();
		log.push("start");
		EntryPoint.runInMainThread(function() log.push("first"));
		EntryPoint.runInMainThread(function() log.push("second"));
		log.push("queued");
		EntryPoint.run();
		log.push("done");
		Sys.println(log.join("|"));
	}
}
```

Under Haxe `--interp`, the output was:

```text
start|queued|done
```

That means Haxe `--interp` does not execute `EntryPoint.runInMainThread`
callbacks in this direct probe. The generated Go implementation does execute
queued callbacks through `sys.thread.EventLoop` / `runtime/hxrt/thread.go`.

## Decision

Do not add a semantic-diff fixture for direct `haxe.EntryPoint`,
`haxe.MainLoop`, or `haxe.Timer` yet.

The current honest evidence remains:

- `test/snapshot/stdlib/haxe_main_loop_runtime_direct`

That fixture proves the generated Go event-loop bridge compiles and runs, but it
is not an interpreter-vs-Go parity claim.

## Why this is the right boundary

A semantic-diff fixture is useful only when the Haxe interpreter is a meaningful
reference for the behavior being compared. For event-loop APIs, the useful Go
behavior is callback scheduling through the Go-backed runtime event loop. The
interpreter does not expose the same observable callback drain behavior for the
simple direct probe above.

Adding a semantic-diff test anyway would create one of two bad outcomes:

- It would fail for a real oracle mismatch rather than a compiler bug.
- It would avoid the callback behavior and only test superficial API access,
  which would not protect users from event-loop regressions.

## Reopen trigger

Reopen this decision only when the harness can compare event-loop behavior with a
normalization layer that is explicit about what is being compared.

A credible future harness would need to:

1. run the target until callbacks reach a known quiescent point,
2. compare logical facts such as callback order or "callback eventually ran",
3. avoid comparing wall-clock timing,
4. use bounded timeouts so CI cannot hang forever,
5. document which parts of `EntryPoint`, `MainLoop`, and `Timer` are covered.

Until then, snapshot/runtime evidence is the safer production contract.

## Related docs

- `/docs/known-gaps.md#target-sensitive-parity-policy`
- `/docs/portable-module-mapping-contract.md`
- `/docs/semantic-diff-guide.md`
