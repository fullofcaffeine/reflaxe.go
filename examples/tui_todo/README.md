# tui_todo

Canonical complex example for reflaxe.go profile comparisons.

## What it demonstrates

- Shared todo app/domain code compiled to `portable`, `gopher`, and `metal`.
- Profile runtime adapters chosen at compile-time via `profile/RuntimeFactory.hx`.
- Equivalent baseline semantics across profiles.
- Additive profile capabilities:
  - `gopher`: batch add helper.
  - `metal`: batch add + diagnostics.

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
