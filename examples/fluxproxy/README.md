# fluxproxy

Reverse-proxy style flagship app scaffold.

Current scope implements a deterministic proxy contract with:

- bounded ingress queue + deterministic backpressure
- route normalization + upstream mapping
- retry accounting from status policy
- per-route aggregate digest
- one shared Haxe codebase across `portable` and `metal`
- explicit runtime variants (`core`, `go_native`)
- deterministic scripted output for CI
- interactive command mode for local demos

## Compile

```bash
haxe compile.portable.hxml
haxe compile.metal.hxml
haxe compile.portable.ci.hxml
haxe compile.metal.ci.hxml
```

## Run

```bash
(cd out_portable && go run . --scripted)
(cd out_metal && go run . --scripted)
```

## Modes

- scripted: deterministic contract output (`--scripted`)
- interactive: command session (`help`, `profile`, `status`, `ingest`, `reset`, `scripted`)

Examples:

```bash
(cd out_portable && go run . --scripted)
(cd out_portable && go run . status)
(cd out_portable && go run . ingest /v1/items 45 200 status)
```

## Variant strategy

- `core` variant:
  - deterministic loop dispatch path (`runtime.capability=loop_dispatch`)
- `go_native` variant:
  - typed worker channel fan-out path (`runtime.capability=worker_chan_fanout`)

## Matrix expectation

- `run.args` pins example harness execution to scripted mode.
- `*.stdout` files represent `core` variant scripted output.
- `*.ci.stdout` files represent `go_native` variant scripted output.

## Pure-Go parity baseline

```bash
(cd benchmarks/pure_go/fluxproxy && go run . --scripted)
(cd benchmarks/pure_go/fluxproxy && go test ./...)
(cd benchmarks/pure_go/fluxproxy && go test -run '^$' -bench . -benchmem)
```
