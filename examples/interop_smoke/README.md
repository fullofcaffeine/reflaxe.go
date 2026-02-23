# interop_smoke

Minimal typed-interop reference example for reflaxe.go.

## What It Demonstrates

- `@:go.import` package mapping (`fmt`, `time`, `context`).
- `@:go.name` symbol mapping for types and fields.
- `@:go.receiver` lowering for static receiver-style calls.
- Interface-returning package APIs (`context.Background(): context.Context`).
- One Haxe codebase compiled across `portable`, `gopher`, and `metal`.

## Compile

```bash
haxe compile.portable.hxml
haxe compile.gopher.hxml
haxe compile.metal.hxml
```

## Run

```bash
(cd out_portable && go run .)
(cd out_gopher && go run .)
(cd out_metal && go run .)
```

Expected output for every profile:

```text
1
```
