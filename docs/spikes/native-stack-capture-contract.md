# Spike: Native Stack Capture Contract

Issue: `haxe.go-14as.40`

## Decision

Do not make native stack capture part of the portable semantic baseline yet.

Keep the current default behavior for `haxe.CallStack` and `haxe.NativeStackTrace`:

- `CallStack.callStack()` returns an empty stack.
- `CallStack.exceptionStack()` returns an empty stack.
- `NativeStackTrace.callStack()` returns an empty stack carrier.
- `NativeStackTrace.toHaxe(...)` remains deterministic for the carriers it already accepts.

Accept a future Go-only implementation as an explicit target-sensitive diagnostic feature, not as semantic-diff parity.

Recommended future knob:

```text
-D reflaxe_go_native_stack_trace
```

That define would mean: generated Go may use Go runtime stack inspection to return best-effort call frames for debugging. It would not mean those frames are portable across targets or stable enough for semantic-diff comparison.

## Terms

A stack trace is a list of active function calls at one moment in time.

A stack frame is one entry in that list. It usually contains a function name, source file, and line number.

Native stack capture means asking the target runtime, here Go, for its own call stack. In Go this would likely use `runtime.Callers` and `runtime.CallersFrames`.

Portable semantics means behavior we can compare against the Haxe reference interpreter and promise across targets. Stack traces are not a clean portable semantic surface because frame names, generated helper frames, inlining, runtime wrappers, and optimization decisions differ by target.

Snapshot coverage means we assert a deterministic generated-output or runtime-smoke shape for this target. It is useful for target-sensitive behavior that should be stable on Go but cannot honestly be compared against the Haxe interpreter.

Semantic-diff coverage means we run the same Haxe program through Haxe `--interp` and through `haxe.go`, then compare observable output. This is the right gate for portable behavior, but the wrong first gate for native stack frames.

## Current State

The current Go implementation is staged std code:

- `std/haxe/CallStack.cross.hx`
- `std/haxe/NativeStackTrace.cross.hx`

Those files intentionally expose the Haxe API surface with deterministic empty-stack behavior. The current fixture is:

- `test/snapshot/stdlib/haxe_stack_loop_target_sensitive`

The inventory and roadmap classify this as target-sensitive snapshot coverage instead of semantic-diff parity.

This is honest: user code can import and call the APIs, but the compiler does not pretend Go-native stack frames already match Haxe interpreter stack behavior.

## Why Not Promote This To Portable Semantic-Diff Now?

Native stack traces are useful for debugging, but they are not stable portable semantics.

Go stack frames would include generated Go function names such as lowered module names, helper wrappers, `hxrt` calls, and runtime functions. The exact frame list can change when:

- generated code shape changes,
- helper functions move between compiler, staged std, and `hxrt`,
- Go inlining changes,
- Go version changes,
- exception or panic plumbing changes,
- snapshot fixtures run from different source locations.

If we compared those frames directly against Haxe `--interp`, we would either get noisy failures or be forced to normalize away most of the information that made native capture useful in the first place.

## Future Target Design

Keep public Haxe-facing behavior in staged std:

- `std/haxe/CallStack.cross.hx` owns the `CallStack` abstract and `StackItem` formatting surface.
- `std/haxe/NativeStackTrace.cross.hx` owns the `NativeStackTrace` API shape.

Add Go runtime support under `hxrt` only when the opt-in define is enabled:

- new runtime helper, likely `runtime/hxrt/stack.go`, captures Go frames with `runtime.Callers`.
- each runtime frame carries at least function name, file, and line.
- staged std converts runtime frames into `haxe.CallStack.StackItem` values.

Suggested Haxe mapping:

```text
Go function/file/line -> FilePos(Method(classNameOrNull, methodName), file, line, 0)
```

The first implementation may use `null` for class names and the Go function name as the method string. Better Haxe source mapping can come later after generated-symbol and source-map policy is stable.

## Boundary Rules

This feature must be explicit.

Do not silently change the default portable behavior from empty deterministic stacks to native Go stacks. That would change user-visible output for a target-sensitive surface without a clear opt-in.

The opt-in feature may be available under both profiles:

- `portable + reflaxe_go_native_stack_trace`: still portable-profile code, but with a Go-only diagnostic capability enabled.
- `metal + reflaxe_go_native_stack_trace`: same diagnostic capability, with metal's normal Go-first boundary policy.

The feature must not:

- change the meaning of `portable` or `metal`,
- require app/test/example code to use raw `__go__`,
- add compiler-owned raw Go blocks when `hxrt` can own the runtime work,
- claim semantic-diff parity for raw native frame contents.

## Harness Plan

Use snapshot/runtime tests first.

Required first tests:

- keep `stdlib/haxe_stack_loop_target_sensitive` for default empty-stack behavior.
- add a new snapshot fixture for `-D reflaxe_go_native_stack_trace` that checks stable facts, not exact full paths.

Stable facts can include:

- captured stack length is greater than zero,
- `NativeStackTrace.toHaxe(stack, 0)` returns at least one item,
- `skip` reduces or preserves length as expected,
- `CallStack.toString(...)` contains a known generated function or module marker after path normalization.

Do not assert absolute machine-local paths.

Do not add this to semantic-diff until a later design defines a stable normalization format that compares Haxe-level frames, not raw Go frames.

## Implementation Tasks

Follow-up implementation should be split into small beads:

1. Add the explicit define and runtime plan reporting.
2. Add `hxrt` native frame capture behind that define.
3. Convert captured Go frames to `StackItem` in staged std.
4. Add target-sensitive snapshot/runtime coverage.
5. Revisit whether any normalized semantic-diff subset is possible.

## Non-Goals

Do not replace the current deterministic fallback in this spike.

Do not build full Haxe source-map reconstruction in the first implementation.

Do not promise exact frame names, exact frame counts, or exact file paths as portable behavior.

Do not copy another target's stack model mechanically. `haxe.rust` currently also uses empty stack arrays for this surface, while `haxe.elixir` has a separate source-mapping and diagnostic story. Go should use Go's runtime facilities through `hxrt` when we implement this.

## Result

The current fallback is retained and documented as the correct default for now.

Native stack capture is worth implementing later as an explicit Go diagnostic capability, with snapshot/runtime coverage first and semantic-diff only if a stable normalized Haxe-level frame contract is designed.
