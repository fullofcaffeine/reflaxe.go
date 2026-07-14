# tui_todo

Portable-first todo CLI reference with deterministic scripted output and interactive command mode.

## What this app does

- Manages a small todo list (add/toggle/tag/batch/list/summary/diag).
- Persists state in `.tui_todo_state.txt` in the current directory.
- Provides a deterministic `--scripted` flow used by harness/CI.

## Profile support

- portable: Yes
- metal: No

This example is intentionally portable-only. Its behavior is stdlib-heavy and
there is no meaningful second-preset lesson; a metal lane previously created
synthetic runtime differences instead of exercising real policy value.

For concrete native-policy examples, use:

- `examples/worker_pool_select`
- `examples/pulseforge` (`go_native` lanes)
- `examples/fluxproxy` (`go_native` lanes)

## Compile

```bash
haxe compile.portable.hxml
haxe compile.portable.ci.hxml
```

## Run

```bash
(cd out_portable && go run . --scripted)
(cd out_portable && go run . help)
(cd out_portable && go run . add 2 Write_profile_docs tag 1 docs list)
```

## Commands

- `reset`
- `help`
- `add <priority> <title_token>`
- `toggle <id>`
- `tag <id> <tag_token>`
- `batch <priority> <title1_token> <title2_token>`
- `list`
- `summary`
- `diag`

Token note: use `_` for spaces (for example `Write_profile_docs`).

## Expected scripted contract

Portable scripted output is validated by:

- `expected/portable.stdout`
- `expected/portable.ci.stdout`

## Related docs

- [`docs/examples-matrix.md`](../../docs/examples-matrix.md)
- [`docs/profiles.md`](../../docs/profiles.md)
- [`docs/profile-semantics-guide.md`](../../docs/profile-semantics-guide.md)
