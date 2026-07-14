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

Write a report that explains every `Dynamic` fallback:

```bash
npm run dev:goextern -- --package fmt --dynamic-report test/.test-cache/goextern-fmt-dynamic.json
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
- supported multi-return signatures -> generated tuple carrier classes
- unsupported/complex boundaries -> `Dynamic`

Generated tuple carrier pattern:

- Go can return multiple values, for example `time.Time.Zone() (name string, offset int)`.
- Haxe functions return one value, so `goextern` generates a small carrier class such as `TimeZoneResult`.
- The extern method is marked with `@:go.tupleReturn`.
- `haxe.go` lowers the call into the carrier and converts common native Go values, such as Go `string` and Go `error`, into the Haxe-facing representation.
- Tuple carriers are generated when every result value maps without `Dynamic`, including scalar values, slices/arrays, `error`, and named types from the same generated package.
- If any result value cannot be typed honestly yet, the method still returns `Dynamic`.

First-class `(T,error)` interop pattern:

- Declare extern methods with return type `go.Result<T>`.
- Add `@:go.valueError` to the method metadata.
- Compiler lowering wraps the Go multi-return call into `go.Result<T>` automatically.
- See `examples/interop_smoke` and `test/semantic_diff/go_value_error_result_contract`.

This keeps generated externs broadly usable while avoiding unsafe precision claims.
You can refine generated files manually where stronger typing is needed.

## Advanced Signature Boundary Policy

Most Go APIs fit into one of three interop paths.

Terms:

- **Ordinary typed extern metadata** means `@:go.import`, `@:go.name`, and
  optionally `@:go.receiver`. Use this when a Go function has zero or one return
  value that maps cleanly to a Haxe type.
- A **single `(T,error)`** signature means a Go function returns one value plus an `error`,
  such as `strconv.Atoi(s) (int, error)`.
- **Typed facade wrapper** means a small Go-facing Haxe extern or helper that
  presents a simpler shape to application code. Use this when the raw Go
  signature is too complex for the generator to type honestly.
- **Tuple carrier** means a generated Haxe class with one field per Go return
  value. For example, a Go `(name string, offset int)` return can become a
  `TimeZoneResult` value with `name` and `offset` fields.

| Go signature shape | What `goextern` does today | Recommended user path |
| --- | --- | --- |
| `func Name()` or `func Name(T) U` | Emits ordinary typed extern metadata when the types map cleanly. | Use the generated extern directly. |
| `func Name(T) (U, error)` | Can be represented as `go.Result<U>` when the extern is explicitly authored with `@:go.valueError`. | Prefer a typed extern or facade that returns `go.Result<U>`. See `examples/interop_smoke` and `test/semantic_diff/go_value_error_result_contract`. |
| `func Name(...) (A, B)` where both values map cleanly | Emits a generated tuple carrier such as `NameResult`, plus `@:go.tupleReturn`. | Use the generated carrier directly. Example evidence: `test/snapshot/go_native/extern_tuple_return`. |
| `func Name(...) (A, B, C...)` where every value maps cleanly | Emits a generated tuple carrier with one field per Go result value. | Use the generated carrier directly, or write a smaller facade if a domain-specific name is clearer. |
| any multi-return signature containing unsupported result types | Emits `Dynamic` for that method. | Add a typed facade wrapper. Do not pretend the generated `Dynamic` value is portable or fully typed. |
| callbacks, channels, generics, unsafe pointers, structs from another package | Usually emits `Dynamic` at the boundary. | Wrap the API behind a smaller typed facade first, then expose that facade to Haxe. |

Generated tuple carriers are intentionally simple. They are good when the Go
result names are already meaningful, such as `name` and `offset`. A hand-written
typed facade is still better when the raw Go shape needs domain-specific names,
validation, retries, or conversion into a higher-level Haxe API.

## Dynamic Fallback Report

`Dynamic` means "the generator could not honestly describe this Go value with a
specific Haxe type yet." It is usable, but it is not the ideal authoring
surface: you lose compile-time type checking at that boundary.

Use `--dynamic-report <path>` when generating externs to see exactly where this
happened:

```bash
npm run dev:goextern -- --package fmt --dynamic-report test/.test-cache/goextern-fmt-dynamic.json
```

The report is deterministic JSON:

```json
{
	"schemaVersion": 1,
	"fallbacks": [
		{
			"package": "fmt",
			"symbol": "Fprint",
			"position": "param:w",
			"goType": "io.Writer",
			"reason": "external_named_type"
		}
	]
}
```

Field meanings:

- `package`: the Go package being generated.
- `symbol`: the Go function or method that contains the fallback.
- `position`: where the fallback happened, such as `param:w` or `result:1`.
- `goType`: the original Go type at that boundary.
- `reason`: why `goextern` used `Dynamic`.

Common reason codes:

- `callback_signature`: the boundary is a Go function value, such as `func(rune) bool`.
- `external_named_type`: the type comes from another Go package that was not generated with this package.
- `unsupported_map_key`: the Go map key is not a `string`, so it does not map to `haxe.DynamicAccess<T>`.
- `struct`: the boundary is an anonymous Go struct.
- `empty_interface`: the boundary is `any` / `interface{}`.
- `non_empty_interface`: the boundary is an interface with methods.
- `channel`: the boundary is a Go channel.
- `type_parameter`: the boundary uses a Go generic type parameter.
- `unsafe_pointer`: the boundary uses `unsafe.Pointer`.

How to use this report:

1. Generate the externs and the report.
2. Find fallbacks in the APIs your app actually calls.
3. For important APIs, write a small typed facade wrapper that converts the raw
   Go shape into a simpler Haxe-facing shape.
4. Leave low-value or truly dynamic boundaries as `Dynamic`.

App code and examples must not use raw `__go__` to bypass these boundaries. Raw
injection is reserved for controlled framework/runtime layers; see
`docs/profiles.md` and `docs/stdlib-shim-rationale.md`.

The current boundary is locked by:

- `tools/goextern/main_test.go`, especially the `fmt` multi-return boundary test.
- `npm run test:goextern`
- `npm run test:goextern:fixtures`

If this policy changes, update the generator, fixtures, and this section in the
same change.

## Validation

Run generator unit/integration tests:

```bash
npm run test:goextern
```

Run deterministic fixture drift checks (used in CI):

```bash
npm run test:goextern:fixtures
```

## Fixture Pinning

`goextern` fixtures are checked in so CI can detect accidental generator drift.
Those committed fixtures are pinned to one Go release because the Go standard
library surface can change between releases. The default pinned release is Go
`1.23`, controlled by `GOEXTERN_FIXTURE_GO_VERSION` in `test/run-ci.py`.
This non-production compatibility fixture does not establish security support.
The supported build lines are governed separately by
[`toolchain-policy.json`](toolchain-policy.json).

Full CI still runs `goextern` confidence checks on other local Go versions:

- `npm run test:goextern` always runs the generator unit tests.
- If the current Go release matches the fixture pin, CI runs the committed
  fixture drift check.
- If the current Go release differs from the fixture pin, CI runs a
  current-toolchain smoke generation into `test/.test-cache` without comparing
  that output to the pinned committed fixtures.

Use `python3 test/run-goextern-fixtures.py --update` only when intentionally
refreshing the committed pinned fixtures under the pinned Go release.
