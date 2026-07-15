# `hxrt` Runtime: Why It Exists, How It Works, and What It Does

## Terms

- `hxrt`: the runtime helper package copied into each generated Go module (`<go_module>/hxrt`).
- `compiler shim`: helper declarations emitted by the compiler at build time (not copied from `runtime/hxrt` files).
- `semantic contract`: the source behavior rules we keep stable, especially
  portable Haxe semantics and explicit typed native API contracts.

## What `hxrt` is

`hxrt` is the Go runtime support package emitted with every `reflaxe.go` build output.

- Source of truth: `runtime/hxrt/*.go`
- Generated location: `<go_output>/hxrt/*.go`
- Import path in generated code: `<go_module>/hxrt`

`hxrt` is not a separate dependency you install from the network. It is copied into the generated module so output stays self-contained and reproducible.

Selective runtime strategy and profile interaction are documented in `docs/hxrt-selective-runtime.md`.

## Why `hxrt` is needed

`hxrt` exists to bridge semantic and representation gaps between Haxe and Go in a deterministic, reusable way.

This is not a Go-only idea; many compiler targets use a runtime package. In `reflaxe.go`, the runtime is intentionally paired with compiler shims and staged stdlib migration to keep parity work incremental and verifiable.

1. Haxe semantics do not map 1:1 to native Go primitives.
   - String behavior and nullability need helper semantics (`Std.string` shape, null-safe concat/equality, rune-aware length/indexing).
   - Haxe exception flow (`throw`/`try`/`catch`) needs a controlled panic/recover boundary.
   - Some stdlib/runtime contracts require target-specific behavior (`Sys`, file/process wrappers, atomic cell behavior).

2. Centralizing helpers avoids re-emitting large behavior blocks into every generated file.
   - Compiler lowers calls to stable helpers instead of duplicating logic.
   - Generated output remains smaller and easier to inspect.

3. It supports strict boundary policy.
   - Project code does not need raw `__go__` for basic runtime glue.
   - Compiler/runtime contract is explicit and testable.

For broader shim ownership tradeoffs, see `docs/stdlib-shim-rationale.md`.

## How it works

Compilation wiring:

1. `go_module` is resolved (default `snapshot`).
2. Runtime import path is computed as `<module>/hxrt` in `src/reflaxe/go/CompilationContext.hx`.
3. Compiler emits generated Go that imports and calls `hxrt` helpers.
4. On output, backend writes:
   - `go.mod`
   - generated `.go` files
   - copied runtime directory `hxrt/` from `runtime/hxrt`
   Every destination is validated by the
   [generated-output confinement boundary](generated-output-confinement.md)
   immediately before the managed write; an existing or broken output symlink
   is rejected rather than followed.
5. Backend runs `go build` by default and fails compilation if it cannot
   launch or exits nonzero. `-D go_no_build` and `-D go_codegen_only` are
   explicit opt-outs for callers that own a separate Go build/test gate.

Key implementation points:

- Runtime copy/write: `src/reflaxe/go/GoReflaxeCompiler.hx`
- Runtime copy helper for iterator flows: `src/reflaxe/go/GoOutputIterator.hx`
- Runtime source: `runtime/hxrt/*.go`

## What `hxrt` currently does

`hxrt` currently owns helper functions in these groups:

- String/runtime conversion helpers:
  - `StringFromLiteral`, `StdString`, `StringSlice`
  - `StringConcatAny`, `StringEqualAny`
  - `StringConcatStringPtr`, `StringEqualStringPtr`
  - `StringLength`, `StringCharAt`, `StringCharCodeAt`, `StringSubstring`
- Numeric helpers:
  - `FloatMod`, `Int32Wrap`
- Atomic runtime cells:
  - `AtomicInt*` helpers
  - `AtomicObject*` helpers
- Exception bridging:
  - `Throw`, `TryCatch`, `UnwrapException`, `ReportUncaughtException`
  - `ExceptionCaught`, `ExceptionThrown`, `ExceptionMessage`
- Portable thread lifecycle:
  - `ThreadSpawn`, `ThreadSpawnWithEventLoop`, `ThreadWaitForAll`
  - `ThreadLocalNew`, `ThreadLocalGet`, `ThreadLocalSet`
  - `ThreadSpawnDetached` for lifecycle-only scoping of compiler-owned
    `go.Go.spawn` callbacks when portable thread identity is reachable
  - logical thread identity, thread-local values, message queues,
    synchronization, and event-loop state
- JSON wrappers:
  - `JsonParse`, `JsonStringify`
- System wrappers (`runtime/hxrt/sys.go`):
  - `SysGetCwd`, `SysArgs`, `SysGetEnv`, `SysPutEnv`, `SysCommand`, `SysExit`
- File capabilities (`runtime/hxrt/file.go`):
  - typed `FileReadContent`, `FileWriteContent`, `FileReadByteValues`, `FileWriteByteValues`, and `FileCopyContents`
  - opaque `FileInput` / `FileOutput` handles plus `FileOpen*`, read/write, seek/tell/eof, flush, and close operations
  - non-owning `SysStdin`, `SysStdout`, and `SysStderr` handles used by the staged file-stream classes
- Process wrappers (`runtime/hxrt/process.go`):
  - `NewProcess`; process stdin/stdout/stderr; byte I/O; PID, blocking/non-blocking exit status, kill, and close

These helpers preserve native failures at the runtime boundary. Canonical staged file wrappers turn read/write failures into Haxe exceptions and construct `haxe.io.Eof` in Haxe source, while process startup and non-EOF read failures remain distinct from normal EOF and child exit codes. Arbitrary file bytes cross the typed boundary as `Array<Int>` / `[]int`, so `hxrt` does not depend on generated `haxe.io.Bytes` internals. Portable `Sys.putEnv` is the intentional exception: its compiler wrapper discards `SysPutEnv`'s returned error to match the upstream Haxe 4.3.7 eval contract, leaving the error available to typed Go-native bindings.
- Byte representation helpers:
  - `BytesFromString`, `BytesToString`, `BytesClone`

## Exception and concurrency boundaries

Portable Haxe exceptions and native Go panics are deliberately different
failure domains:

1. `throw value` creates an `hxrt.HaxeException` carrier.
   The carrier and `Throw` live in the `core` runtime feature so even a
   core-only selective runtime can represent portable validation failures;
   catch/message helpers remain in the additive `exception` feature.
2. Generated Haxe `try`/`catch` unwraps only that carrier. A panic originating
   in a typed Go extern, the Go runtime, or malformed generated code continues
   unwinding as a native panic and cannot be mistaken for a Haxe value.
   Portable runtime validation failures, such as a failed dynamic-to-`Int`
   conversion, must call `Throw` and therefore remain Haxe-catchable; they must
   never use a raw Go panic as a shortcut.
3. An uncaught Haxe throw in `sys.thread.Thread.create` or
   `createWithEventLoop` ends that worker, writes
   `Uncaught exception <message>` to stderr, and does not crash the process.
4. Portable Haxe workers are foreground threads: a foreground thread keeps the
   generated program alive until it finishes. When the inferred runtime plan
   includes `thread`, generated `main` defers `ThreadWaitForAll`, which also
   drains portable workers created by other portable workers.
5. `go.Go.spawn` remains an explicit Go-native boundary. It is not included in
   the portable foreground count and preserves Go's normal fatal-panic and
   process-shutdown behavior. If `sys.thread` is reachable, the compiler uses
   `ThreadSpawnDetached` to initialize the runtime on the caller before launch
   and release only callback-owned logical identity and TLS state on return or
   panic unwind; it neither joins nor recovers the goroutine, and a nil callback
   retains native panic behavior.
6. Arbitrary goroutines created outside compiler-owned `Thread` and
   `go.Go.spawn` boundaries have no observable exit hook. Calling
   `Thread.current()` or `Tls` from such a foreign goroutine is therefore
   outside the automatic lifecycle-reclamation contract.

Synchronization, pool linearization, event-loop cancellation, condition
generations, channel close behavior, and the bounded `runtime.Stack` identity
fallback are specified in [`docs/concurrency-contract.md`](concurrency-contract.md).

The split is covered by `stdlib/sys_thread_uncaught_exception`,
`go_native/native_panic_not_haxe_catch`,
`go_native/goroutine_native_panic`, and
`go_native/goroutine_native_shutdown`, plus direct race-capable runtime tests
in `runtime/hxrt/exception_test.go` and `runtime/hxrt/thread_test.go`.

## What `hxrt` does not own

`hxrt` is not the whole Haxe stdlib implementation. Public library semantics
belong in upstream or canonical staged Haxe source. The remaining audited
compiler exceptions are limited to surfaces that still depend on compile-time
context or representation-sensitive lowering.

`sys.FileSystem` follows the preferred split: `std/go/_std/sys/FileSystem.hx`
owns its public API and `FileStat` construction, while typed `std/hxrt/fs`
bindings reach the selectively copied native capabilities in
`runtime/hxrt/filesystem.go`.

Examples that are currently compiler-owned (not `hxrt`-owned):

- `sys.Http`
- `sys.net.Socket` / `sys.net.Host`
- `haxe.Serializer` / `haxe.Unserializer`
- most `haxe.io` and `haxe.ds` shim declarations

## Change guidelines

Use `hxrt` when a helper is:

- target-runtime behavior (not just AST rewriting),
- reusable across many lowering sites,
- easier to verify once than duplicated per generated file.

Keep behavior in compiler shims when it depends on compile-time metadata/profile policy or large generated type-shape contracts.

When changing `hxrt`, update evidence:

- snapshots: `python3 test/run-snapshots.py`
- semantic diff: `python3 test/run-semantic-diff.py`
- full CI harness: `python3 test/run-ci.py`

And if ownership boundaries move, update:

- `docs/stdlib-shim-rationale.md`
- `docs/feature-support-matrix.md`

## Related docs

- `docs/start-here.md` - first-run commands and how generated output is produced.
- `docs/hxrt-selective-runtime.md` - runtime feature slicing policy and controls.
- `docs/stdlib-shim-rationale.md` - ownership split between runtime, compiler shims, and staged stdlib overrides.
- `docs/native-policy-presets.md` - policy presets and native-boundary behavior.
- `docs/concurrency-contract.md` - portable and Go-native concurrency invariants and evidence.
- `docs/portable-canonical-contract.md` - portable semantics baseline used by tests.
- `docs/glossary.md` - shared definitions used across docs.
