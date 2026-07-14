# Examples

This repository ships canonical app examples for portable semantics, explicit
Go-native boundaries, and the compatible portable/metal policy presets.

| Example | What it does | Profile support | Portable vs metal diff |
| --- | --- | --- | --- |
| `profile_storyboard` | Renders a release-planning command-center dashboard from deterministic card data. | portable only | N/A (portable reference only; no meaningful metal delta). |
| `tui_todo` | Interactive + scripted todo CLI with local state persistence and deterministic contract output. | portable only | N/A (portable reference only; no meaningful metal delta). |
| `incident_api` | Runnable loopback HTTP incident service using Haxe stdlib sockets, JSON, config, and file-backed state. | portable + metal | Same user behavior; metal exercises stricter/eager compatibility defaults, not another semantic engine. |
| `interop_smoke` | Typed Go interop smoke (`@:go.import`, `@:go.name`, `@:go.receiver`, `@:go.valueError`). | portable + metal | Intentionally near-identical output; validates typed API behavior across preset defaults. |
| `worker_pool_select` | Deterministic worker pool using `go.Chan` + typed `go.Select` helpers. | portable + metal | Same Go-native API contract; metal emphasizes eager specialization policy. |
| `pulseforge` | Flagship observability pipeline with `core` and `go_native` runtime variants. | portable + metal | Largest practical delta appears in `go_native` lanes (typed channel/select specialization pressure). |
| `fluxproxy` | Flagship reverse-proxy policy pipeline with `core` and `go_native` runtime variants. | portable + metal | Largest practical delta appears in `go_native` lanes (worker fanout + typed channel/select paths). |

## Why not force every app to have both presets

Dual-preset examples are useful only when they exercise meaningful policy or
report differences. Portable-only examples are used when the second preset
would add synthetic, non-instructive duplication.

## Start here by goal

- Portable-first app reference: `examples/tui_todo`, `examples/profile_storyboard`.
- Real stdlib service reference: `examples/incident_api`.
- Typed interop surface: `examples/interop_smoke`.
- Concurrency/select specialization-policy comparison:
  `examples/worker_pool_select`.
- Large app + perf lanes: `examples/pulseforge`, `examples/fluxproxy`.

## Quick commands

Run example harness:

```bash
python3 test/run-examples.py
```

Refresh committed generated trees:

```bash
python3 test/run-examples.py --bless-generated
```

Build cross-platform binaries from committed generated trees:

```bash
bash scripts/examples/build-binaries.sh
```

## Canonical references

- Compatibility preset matrix: [`docs/profiles.md`](../docs/profiles.md)
- Native policy contract: [`docs/native-policy-presets.md`](../docs/native-policy-presets.md)
- Profile semantics guide: [`docs/profile-semantics-guide.md`](../docs/profile-semantics-guide.md)
- Example/perf matrix: [`docs/examples-matrix.md`](../docs/examples-matrix.md)
