# Family Raw-Injection Authority Alignment

This note is a handoff for sibling compiler work in `haxe.rust` and
`haxe.ocaml`. It records the `haxe.go` raw-injection authority model so those
repos can compare against it without blindly copying Go-specific mechanics.

## Terms

- **Raw injection** means embedding target-language code directly from Haxe
  source, such as Go snippets through `__go__`.
- **Framework-owned** means code owned by the compiler/runtime/standard-library
  compatibility layer, not normal application code.
- **Typed extern metadata** means target binding metadata that names a real
  target package, symbol, or receiver, such as Go's `@:go.import`,
  `@:go.name`, and `@:go.receiver`.
- **Allow tag** means metadata that grants a narrow framework module permission
  to use raw injection under strict boundary checks. In `haxe.go`, that tag is
  `@:goAllowRaw`.
- **Compiler-owned raw emitter** means generated target code emitted directly by
  compiler internals, for example `GoRaw` blocks in `GoCompiler`.

## Go Decision

`haxe.go` treats raw target injection as a controlled framework authority, not
as the default user-facing interop story.

The intended order is:

1. Use typed extern metadata when calling a real Go API.
2. Use a narrow framework-owned helper island when staged std or runtime code
   needs raw generated-shape access that Haxe source cannot express directly.
3. Keep compiler-owned raw emitters only when correctness depends on compiler
   context, profile policy, or backend representation details.

In Go, the helper-island mechanism is:

- `@:goAllowRaw` on the module/type that owns the low-level bridge.
- `reflaxe.go.macros.GoInjection.__go__(...)` or direct framework-owned
  `untyped __go__(...)` inside that bridge.
- Typed extern metadata for imports and symbol binding. Raw `__go__` snippets do
  not add Go package imports by themselves.

`std/sys/GoHttpHelpers.hx` remains an example of an acceptable framework helper
island. The former `std/haxe/io/GoIoHelpers.hx` island was retired by
`haxe_go-vfp.8.7.11` when the complete base IO hierarchy moved to canonical
staged source. App code and examples still must not teach raw `__go__` as normal
business logic.

## Sibling Review Instructions

### `haxe.rust`

The local agent should compare this model against `haxe.rust`'s own compiler,
runtime, and staged-stdlib architecture.

Do not copy Go names or mechanics directly. Instead, check whether Rust already
has equivalent concepts, such as:

- a framework-owned raw Rust escape hatch,
- typed wrappers for real Rust APIs,
- staged std/runtime helpers before compiler-owned raw emission,
- strict boundary or profile rules that decide where raw target code is allowed.

If `haxe.rust` adopts or changes an allow-tag mechanism, its docs should explain:

- what the tag allows,
- why the tag is framework-owned rather than app-owned,
- how typed imports/wrappers remain the default interop path,
- which tests prove strict/profile boundaries still reject app-level raw escapes.

### `haxe.ocaml`

The local agent should compare this model against `haxe.ocaml`'s local
architecture and existing `__ocaml__` policy.

Do not assume Go's `@:goAllowRaw` maps one-to-one to an OCaml metadata tag.
Instead, decide what fits OCaml's own staged std, runtime, profile, and boundary
enforcement design.

If `haxe.ocaml` adopts or changes an allow-tag mechanism, its docs should
explain:

- the intended framework-owned scope,
- why normal app code should prefer typed or staged abstractions,
- how raw OCaml snippets interact with imports/modules,
- which boundary tests prevent the allow tag from becoming a general escape
  hatch.

## What Must Stay Aligned

- Raw injection is not the normal application interop surface.
- Typed imports/wrappers stay the preferred way to bind real target APIs.
- Scoped allow tags are for narrow framework-owned abstraction islands.
- Compiler-owned raw emitters should shrink when staged std or runtime helpers
  can own the behavior honestly.
- Sibling repos must compare against their own architecture before adopting this
  pattern.
- Any adopted allow-tag mechanism must be documented with its rationale and
  regression-tested against strict/profile boundaries.

## What Must Not Be Copied Blindly

- Go metadata names such as `@:goAllowRaw`, `@:go.import`, or `@:go.receiver`.
- Go's exact helper shape using `reflaxe.go.macros.GoInjection.__go__`.
- Go-specific ownership decisions for `sys.Http`, `haxe.io`, sockets, or bytes.

Those are examples of the rule, not the family rule itself.

## Local Go References

- `docs/profiles.md` documents strict raw `__go__` policy and `@:goAllowRaw`.
- `docs/defines-reference.md` defines the `@:goAllowRaw` metadata contract.
- `docs/stdlib-shim-rationale.md` records the post-`__go__` ownership rule and
  current shim decisions.
- `docs/stdlib-shim-migration-log.md` records the `haxe.io` and `sys.Http`
  extractions that motivated this alignment handoff.
- `src/reflaxe/go/macros/GoInjection.hx` documents the Go macro shim.
- `src/reflaxe/go/analyze/GoRawInjectionAuthorityAnalyzer.hx` centralizes Go's
  allow-tag detection.
