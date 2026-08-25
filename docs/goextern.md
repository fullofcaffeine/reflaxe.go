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

Resolve a package from another Go module without changing that module:

```bash
npm run dev:goextern -- --package example.com/app/api --dir ../app
```

Pass one exact Go import path. Relative and wildcard package patterns are not
accepted because the import path is part of the generated ownership identity.

Print generated files without writing to disk:

```bash
npm run dev:goextern -- --package time --stdout
```

Write a report that explains every `Dynamic` fallback:

```bash
npm run dev:goextern -- --package fmt --dynamic-report test/.test-cache/goextern-fmt-dynamic.json
```

## Output Layout

By default, the Haxe package `goextern` is written under:

```text
gen/goextern/<sanitized Go import path>/
```

Examples:

- `fmt` -> `gen/goextern/fmt/*.hx`
- `context` -> `gen/goextern/context/*.hx`
- `github.com/foo/bar` -> `gen/goextern/github_com/foo/bar/*.hx`

The output path matches the declared Haxe package. The `@:go.import` metadata
keeps the original Go import path without changes. Add `gen` to the Haxe class
path when you use the default output and package prefix.

## What It Emits

- One extern class per supported exported named type in the root package.
- One Haxe `typedef` per supported exported Go alias.
- Externs for supported named dependency types that the root API reaches.
- Exported fields on named structs when their types preserve the exact Go ABI.
- One package static extern class (`<PkgName>Pkg`) for exported package-level functions.
- Metadata for compiler lowering:
  - `@:go.import("<path>")`
  - `@:go.package("<name>")`
  - `@:go.name("<symbol>")`
  - `@:go.struct` (on concrete Go structs with zero-value construction)
  - `@:go.valueError` (when the extern returns Go `(T,error)` and the Haxe return type is `go.Result<T>`)

## Determinism Contract

- The generator uses the exact Go import path and declaration name as identity.
- The root emits all exported types and package functions.
- A dependency emits only reachable exported types and their supported methods.
- A dependency package function requires a separate root invocation.
- Recursive named types terminate and use stable references.
- Symbols, methods, files, diagnostics, and manifests use stable sort keys.
- The same invocation produces the same bytes when its Go package graph is unchanged.

Each root owns its files through `.goextern/roots/<root-key>.json`. Two roots can
own the same file only when its bytes are equal. The generator does not change
an unowned file or a modified owned file. It removes a stale file only after
the final owner releases that file.

Run one generator process at a time for an output tree. The ownership manifest
does not coordinate concurrent writers.

A successful write reports its precision:

```text
generated 17 files across 4 packages; precision=exact; fallbacks=0; out=gen/goextern
```

The value is `precision=partial` when the fallback report is not empty.

The generator stops before a write for these graph and ownership errors:

- `package_load_failed`: Go did not load the selected root package.
- `haxe_package_collision`: two Go paths map to the same Haxe package.
- `output_path_collision`: two declarations map to the same output path.
- `go_import_alias_required`: two Go packages require the same Go qualifier.
- `owned_output_conflict`: two roots plan different bytes for one owned path.
- `owned_output_modified`: a person or another tool changed an owned file.
- `unowned_output_conflict`: an unowned file already uses the planned path.
- `invalid_ownership_manifest`: an existing ownership record is malformed.
- `unsafe_output_path`: a path escapes the root or crosses a symbolic link.

## Implementation Note

The tool imports `golang.org/x/tools/go/packages` and `go/types`.
For deterministic/offline bootstrap, this repository currently vendors a minimal
`go/packages` compatibility implementation under:

`tools/goextern/third_party/golang.org/x/tools/go/packages`

The compatibility layer asks the selected Go toolchain for read-only export data.
It honors the caller's module directory, so local module packages and their typed
dependencies can be inspected without editing `go.mod` or creating `go.sum`.

## Current Type-Mapping Scope

The generator intentionally starts conservative:

- `bool` -> `Bool`
- integer scalars -> `Int`
- floating scalars -> `Float`
- `string` -> `String`
- slices -> `go.NativeSlice<T>`
- fixed-size Go arrays -> `Dynamic` (`fixed_array`) until their length-bearing ABI has a typed facade
- map with string keys -> `haxe.DynamicAccess<T>`
- exported fields on named structs -> writable Haxe fields with exact `@:go.name` selectors
- supported multi-return signatures -> generated tuple carrier classes
- supported external named types -> fully qualified Haxe types and reachable dependency externs
- unsupported/complex boundaries -> `Dynamic`

Generated named structs have a zero-argument constructor. `haxe.go` lowers it
to an addressed Go composite literal such as `&image.Point{}`. This creates the
ordinary Go zero value; it does not call or invent a package constructor.

The generator emits an exported, non-embedded field only when its complete Go
ABI has an exact Haxe type. The exact set includes `bool`, `int`, `float64`,
`string`, safe native slices, named interfaces, and pointers to supported named
values. These named types can belong to the root or a dependency package.

Unexported fields stay private to Go. The generator omits unsupported fields
and records each omission in the fallback report. Examples include embedded
fields, maps, width-changing scalars, and pointers without a named carrier.

Generated tuple carrier pattern:

- Go can return multiple values, for example `time.Time.Zone() (name string, offset int)`.
- Haxe functions return one value, so `goextern` generates a small carrier class such as `TimeZoneResult`.
- The extern method is marked with `@:go.tupleReturn`.
- `haxe.go` lowers the call into the carrier and converts common native Go values, such as Go `string` and Go `error`, into the Haxe-facing representation.
- Tuple carriers are generated when every result value maps without `Dynamic`.
- Exact results include scalars, slices, `error`, and supported named dependency types.
- If any result value cannot be typed honestly yet, the method still returns `Dynamic`.

First-class `(T,error)` interop pattern:

- Declare extern methods with return type `go.Result<T>`.
- Add `@:go.valueError` to the method metadata.
- Compiler lowering wraps the Go multi-return call into `go.Result<T>` automatically.
- See `examples/interop_smoke` and `test/semantic_diff/go_value_error_result_contract`.

Generated files are not an edit surface. Put stronger user-owned externs or
typed adapters in a separate Haxe path. The ownership manifest rejects a
different unowned file at a generated path.

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
| supported named types from another package | Emits a fully qualified Haxe type and the reachable dependency declaration. | Use the generated extern directly. Generate a separate root for dependency package functions. |
| callbacks, channels, generics, unsafe pointers, or unsupported container shapes | Emits `Dynamic` or omits an unsupported field. | Use a precise user extern or a narrow typed adapter for that boundary. |

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
			"position": "param:a",
			"goType": "[]any",
			"reason": "empty_interface"
		}
	]
}
```

Field meanings:

- `package`: the Go package being generated.
- `symbol`: the Go function, method, or struct field that contains the fallback.
- `position`: where the fallback happened, such as `param:w`, `result:1`, or `field:Created`.
- `goType`: the original Go type at that boundary.
- `reason`: why `goextern` could not preserve an exact typed boundary. For
  callable signatures this usually means `Dynamic`; unsupported struct fields
  are omitted instead.

Common reason codes:

- `callback_signature`: the boundary is a Go function value, such as `func(rune) bool`.
- `generic_named_type`: the named type has type parameters or type arguments.
- `unexported_named_type`: the Go package does not export the named type.
- `alias_target_unsupported`: the exported alias target does not have an exact mapping.
- `unsupported_map_key`: the Go map key is not a `string`, so it does not map to `haxe.DynamicAccess<T>`.
- `struct`: the boundary is an anonymous Go struct.
- `empty_interface`: the boundary is `any` / `interface{}`.
- `non_empty_interface`: the boundary is an interface with methods.
- `channel`: the boundary is a Go channel.
- `type_parameter`: the boundary uses a Go generic type parameter.
- `unsafe_pointer`: the boundary uses `unsafe.Pointer`.
- `embedded_field`: the exported field is embedded and is not emitted as a direct writable selector.
- `scalar_field_abi`: the scalar width or signedness does not match the Haxe carrier.
- `pointer_field_abi`: the pointer does not have a matching named Haxe carrier.
- `slice_element_abi`: the native slice element would change representation.
- `map_field_abi`: no exact native map field carrier is available yet.
- `named_value_field_abi`: a named Go value would be pointer-backed as a Haxe extern class.

How to use this report:

1. Generate the externs and the report.
2. Find fallbacks in the APIs your app actually calls.
3. For important APIs, write a precise user extern or a narrow typed adapter.
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

- `npm run test:goextern` runs the generator unit tests and the cross-package
  Haxe-to-Go runtime tracer.
- If the current Go release matches the fixture pin, CI runs the committed
  fixture drift check.
- If the current Go release differs from the fixture pin, CI runs a
  current-toolchain smoke generation into `test/.test-cache` without comparing
  that output to the pinned committed fixtures.

Use `python3 test/run-goextern-fixtures.py --update` only when intentionally
refreshing the committed pinned fixtures under the pinned Go release.
