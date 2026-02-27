# Examples

This repository ships six first-class multi-profile examples.

| Example | What it teaches | Profile visibility |
| --- | --- | --- |
| `profile_storyboard` | Compact profile walkthrough with explicit runtime adapters | High (`portable` vs `metal` differences are easy to inspect) |
| `tui_todo` | Complex single-codebase app with deterministic scripted contract and interactive mode | Medium (same contract, additive metal capabilities) |
| `interop_smoke` | Typed Go interop metadata (`@:go.import`, `@:go.name`, `@:go.receiver`) | Low by design (intentionally near-identical output) |
| `worker_pool_select` | Go concurrency surfaces (`go.Chan`, `go.Select`, worker fan-out) | Medium-high (typed helper shape differences) |
| `pulseforge` | Flagship observability app with `core` and `go_native` variants | Medium-high (especially in `go_native` lanes) |
| `fluxproxy` | Flagship reverse-proxy app with `core` and `go_native` variants | Medium-high (especially in `go_native` lanes) |

## Why two profiles exist in these examples

- `portable` is the semantic portability contract.
- `metal` is the explicit Go-native/perf lane with stricter boundary defaults.
- One Haxe codebase can target both. If your code stays on portable surfaces, outputs may be similar; that is expected and desirable for compatibility.

## Start here by goal

- Understand profile differences first: `examples/profile_storyboard`.
- Learn baseline app portability: `examples/tui_todo`.
- Learn typed interop metadata: `examples/interop_smoke`.
- Learn Go concurrency surfaces: `examples/worker_pool_select`.
- Evaluate larger app architecture + perf lanes: `examples/pulseforge`, `examples/fluxproxy`.

## Quick commands

Compile/run all example-profile cases and validate expected stdout:

```bash
python3 test/run-examples.py
```

Refresh committed generated trees from fresh outputs:

```bash
python3 test/run-examples.py --bless-generated
```

Build binary artifacts from committed generated trees:

```bash
bash scripts/examples/build-binaries.sh
```

Inspect generated Go differences for a specific example:

```bash
diff -ru examples/profile_storyboard/generated/portable examples/profile_storyboard/generated/metal
```

## Canonical references

- Profile contract matrix: `docs/profiles.md`
- Deep semantics and migration guidance: `docs/profile-semantics-guide.md`
- Example/perf matrix: `docs/examples-matrix.md`
