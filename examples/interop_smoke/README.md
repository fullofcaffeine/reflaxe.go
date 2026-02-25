# interop_smoke

Minimal typed-interop reference example for reflaxe.go.

## What It Demonstrates

- Framework-owned interop wrappers from `std/go/*` (`go.Fmt`, `go.Time`, `go.ContextPkg`, `go.Http`).
- User-level extern metadata in app code (`UserGoTime`, `UserGoContextPkg`, `UserGoHttp`) using:
  - `@:go.import("pkg/path")`
  - `@:go.name("SymbolName")`
  - `@:go.receiver`
- One Haxe codebase compiled across `portable` and `metal`.

Compiler metadata coverage is additionally locked by snapshot fixtures
(`test/snapshot/go_native/extern_metadata_mapping`).

## Profile Note

This example is expected to generate near-identical Go for `portable` and `metal`
because it only uses profile-safe interop calls.
The profile matrix here validates contract consistency, not profile-divergent code shape.

## Compile

```bash
haxe compile.portable.hxml
haxe compile.metal.hxml
```

## Run

```bash
(cd out_portable && go run .)
(cd out_metal && go run .)
```

Expected output for every profile:

```text
1
```
