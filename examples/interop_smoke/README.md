# interop_smoke

Minimal typed-interop reference example for reflaxe.go.

## What It Demonstrates

- `@:go.import` package mapping (`fmt`, `time`, `context`, `net/http`).
- `@:go.name` symbol mapping for types and fields.
- `@:go.receiver` lowering for static receiver-style calls.
- Interface-returning package APIs (`context.Background(): context.Context`).
- `net/http` static API interop (`http.StatusText(200)`).
- One Haxe codebase compiled across `portable` and `metal`.

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
