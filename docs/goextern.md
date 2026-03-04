# goextern (Go -> Haxe extern generator)

`tools/goextern` is a deterministic generator that inspects a Go package using
`go/packages` + `go/types` and emits Haxe externs compatible with `reflaxe.go`.

## Why

- Reduces manual extern authoring for Go interop surfaces.
- Keeps extern generation reproducible (sorted symbols, stable file layout).
- Establishes the foundation for fixture-based drift checks in CI.

## Quick Usage

Generate externs for a Go package:

```bash
npm run dev:goextern -- --package fmt
```

Custom output root and Haxe package prefix:

```bash
npm run dev:goextern -- --package context --out gen/goextern --haxe-package goextern
```

Print generated files without writing to disk:

```bash
npm run dev:goextern -- --package time --stdout
```

## Output Layout

By default, output is written under:

```text
gen/goextern/<go import path>/
```

Examples:

- `fmt` -> `gen/goextern/fmt/*.hx`
- `context` -> `gen/goextern/context/*.hx`
- `github.com/foo/bar` -> `gen/goextern/github.com/foo/bar/*.hx`

## What It Emits

- One extern class per exported named type in the package.
- One package static extern class (`<PkgName>Pkg`) for exported package-level functions.
- Metadata for compiler lowering:
  - `@:go.import("<path>")`
  - `@:go.name("<symbol>")`
  - `@:go.valueError` (when the extern returns Go `(T,error)` and the Haxe return type is `go.Result<T>`)

## Determinism Contract

- Exported symbols are processed in stable lexical order.
- Methods are sorted with stable keys (name + signature).
- File output order is stable.
- Stale `.hx` files in the destination package directory are removed on write.

## Implementation Note

The tool imports `golang.org/x/tools/go/packages` and `go/types`.
For deterministic/offline bootstrap, this repository currently vendors a minimal
`go/packages` compatibility implementation under:

`tools/goextern/third_party/golang.org/x/tools/go/packages`

This keeps the call surface stable while M3.5 lands fixture and CI drift gates.

## Current Type-Mapping Scope

The generator intentionally starts conservative:

- `bool` -> `Bool`
- integer scalars -> `Int`
- floating scalars -> `Float`
- `string` -> `String`
- slices/arrays -> `Array<T>`
- map with string keys -> `haxe.DynamicAccess<T>`
- multi-return signatures -> `Dynamic`
- unsupported/complex boundaries -> `Dynamic`

First-class `(T,error)` interop pattern:

- Declare extern methods with return type `go.Result<T>`.
- Add `@:go.valueError` to the method metadata.
- Compiler lowering wraps the Go multi-return call into `go.Result<T>` automatically.
- See `examples/interop_smoke` and `test/semantic_diff/go_value_error_result_contract`.

This keeps generated externs broadly usable while avoiding unsafe precision claims.
You can refine generated files manually where stronger typing is needed.

## Validation

Run generator unit/integration tests:

```bash
npm run test:goextern
```

Run deterministic fixture drift checks (used in CI):

```bash
npm run test:goextern:fixtures
```
