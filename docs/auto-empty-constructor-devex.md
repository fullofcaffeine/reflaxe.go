# Auto Empty Constructor DevEx (PulseForge Spike)

## Context

PulseForge runtime classes (`CoreRuntime`, `GoNativeRuntime`) required explicit empty `new()` constructors to satisfy typed `new` calls, which added repetitive boilerplate.

## Decision

Use a compiler-level, opt-in define to auto-inject empty constructors for implementations of selected interfaces:

- Define: `-D reflaxe_go_auto_empty_ctor_interfaces=<csv>`
- Example usage: `-D reflaxe_go_auto_empty_ctor_interfaces=app.runtime.PulseRuntime`
- Implementation:
  - `src/reflaxe/go/macros/AutoEmptyConstructor.hx`
  - wired from `src/reflaxe/go/CompilerInit.hx`

This keeps application code clean and avoids per-project `@:autoBuild(...)` annotations.

## Alternatives Considered

1. Per-interface `@:autoBuild` in app code.
- Rejected: still requires users to remember/build-wire macro annotations.

2. Shared abstract/base runtime class with explicit constructor.
- Rejected: couples app architecture to constructor workaround and does not generalize.

3. Keep explicit empty constructors everywhere.
- Rejected: repetitive boilerplate with no semantic value.

## Guardrails

- Opt-in only; no global constructor injection by default.
- Interface paths are validated as fully-qualified names.
- Existing user-defined constructors are never overwritten.
