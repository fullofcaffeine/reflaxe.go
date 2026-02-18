# profile_storyboard

Compact profile-idiom example showing one shared domain compiled to `portable`, `gopher`, and `metal`.

## Why this exists

- Keeps profile differences easy to inspect in a small codebase.
- Demonstrates runtime-adapter pattern (`profile/RuntimeFactory.hx`) for future compilers.
- Shows additive profile behavior while preserving baseline semantics.

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
