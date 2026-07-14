# `hxrt` Runtime: Why It Exists, How It Works, and What It Does

## Terms

- `hxrt`: the runtime helper package copied into each generated Go module (`<go_module>/hxrt`).
- `compiler shim`: helper declarations emitted by the compiler at build time (not copied from `runtime/hxrt` files).
- `semantic contract`: the behavior rules we keep stable for a profile, especially `portable`.

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
  - `Throw`, `TryCatch`, `UnwrapException`
  - `ExceptionCaught`, `ExceptionThrown`, `ExceptionMessage`
- JSON wrappers:
  - `JsonParse`, `JsonStringify`
- System/file/process wrappers:
  - `SysGetCwd`, `SysArgs`, `SysGetEnv`, `SysPutEnv`, `SysCommand`, `SysExit`
  - `FileSaveContent`, `FileGetContent`
  - `NewProcess`; process stdin/stdout/stderr; byte I/O; PID, blocking/non-blocking exit status, kill, and close

These helpers preserve native failures at the runtime boundary. Portable file wrappers turn read/write failures into Haxe exceptions, while process startup and non-EOF read failures remain distinct from normal EOF and child exit codes. Portable `Sys.putEnv` is the intentional exception: its compiler wrapper discards `SysPutEnv`'s returned error to match the upstream Haxe 4.3.7 eval contract, leaving the error available to typed Go-native bindings.
- Byte representation helpers:
  - `BytesFromString`, `BytesToString`, `BytesClone`

## What `hxrt` does not own

`hxrt` is not the whole Haxe stdlib implementation. Behavior-heavy stdlib surfaces still live in compiler-emitted shims where they depend on compile-time context or profile-sensitive lowering.

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
- `docs/profiles.md` - profile contracts and boundary behavior.
- `docs/portable-canonical-contract.md` - portable semantics baseline used by tests.
- `docs/glossary.md` - shared definitions used across docs.
