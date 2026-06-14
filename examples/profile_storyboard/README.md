# profile_storyboard

Portable-first release-dashboard renderer for a small planning board.

## What this app does

- Builds a deterministic set of story cards.
- Computes readiness, lane counts, open-load, ETA, and risk summary.
- Renders a compact “command center” text dashboard.

## Profile support

- portable: Yes
- metal: No

This example is intentionally portable-only. The previous metal lane in this app changed adapter strings/thresholds for demonstration, not because metal generated materially better behavior or meaningful Go-hot-path benefits.

For real portable-vs-metal differences, use:

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
(cd out_portable && go run .)
(cd out_portable_ci && go run .)
```

## Expected scripted contract

Portable outputs are validated by:

- `expected/portable.stdout`
- `expected/portable.ci.stdout`

## Related docs

- [`docs/examples-matrix.md`](../../docs/examples-matrix.md)
- [`docs/profiles.md`](../../docs/profiles.md)
- [`docs/profile-semantics-guide.md`](../../docs/profile-semantics-guide.md)
