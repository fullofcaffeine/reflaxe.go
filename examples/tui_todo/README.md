# tui_todo

Canonical complex example for reflaxe.go profile comparisons.

## What it demonstrates

- Shared todo app/domain code compiled to `portable`, `gopher`, and `metal`.
- Profile runtime adapters chosen at compile-time via `profile/RuntimeFactory.hx`.
- Equivalent baseline semantics across profiles.
- Additive profile capabilities:
  - `gopher`: batch add helper.
  - `metal`: batch add + diagnostics.
- User-driven command session mode for local demo runs.
- Deterministic scripted mode for CI (`--scripted`).

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

This starts command-session mode with commands like:

- `help`
- `add <priority> <title_token>`
- `toggle <id>`
- `tag <id> <tag_token>`
- `batch <priority> <title1_token> <title2_token>` (gopher/metal)
- `list`
- `summary`
- `diag`

Token note: use `_` where you want spaces (for example `Write_profile_docs`).

Scripted deterministic output mode (used by harness):

```bash
(cd out_portable && go run . --scripted)
```

Command-session mode examples:

```bash
(cd out_portable && go run . help)
(cd out_portable && go run . add 2 Write_profile_docs tag 1 docs list)
(cd out_gopher && go run . batch 3 Ship_generated_go_sync Add_binary_matrix list)
```
