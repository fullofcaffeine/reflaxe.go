# pulseforge

Flagship observability-stream pipeline demo.

Current scope implements the full scripted contract pipeline:

- ingest with bounded queue + deterministic backpressure policy
- parse stage (typed ingress frame normalization)
- enrich stage (severity + weighted value)
- aggregate stage (per-source rollups)
- alert stage (weighted-threshold alerts)
- one shared Haxe codebase across `portable` and `metal`
- explicit runtime variants (`core`, `go_native`)
- deterministic scripted output for CI
- interactive command mode for local demos

## Compile

Core variant (default compile files):

```bash
haxe compile.portable.hxml
haxe compile.metal.hxml
```

Go-native variant (CI compile files):

```bash
haxe compile.portable.ci.hxml
haxe compile.metal.ci.hxml
```

## Run

```bash
(cd out_portable && go run .)
(cd out_metal && go run .)
```

Or build binaries:

```bash
(cd out_portable && go build -o pulseforge_portable . && ./pulseforge_portable)
(cd out_metal && go build -o pulseforge_metal . && ./pulseforge_metal)
```

## Modes

- scripted: deterministic contract output (`--scripted`)
- interactive: command session (`help`, `profile`, `status`, `ingest`, `reset`, `scripted`)

Examples:

```bash
(cd out_portable && go run . --scripted)
(cd out_portable && go run . status)
(cd out_portable && go run . ingest edge 8 iad status)
```

## Variant strategy

- `core` variant:
  - deterministic loop-based runtime path (`runtime.capability=core_loop`)
- `go_native` variant:
  - typed channel/select path (`runtime.capability=chan_fanout_select`)

Current profile behavior differs by runtime adapter and generated code shape while keeping the same domain contract and deterministic outputs.

- `core`:
  - loop-based stage adapters
  - capability id: `core_loop`
- `go_native`:
  - channel/select worker adapters with fan-out + fan-in flow
  - capability id: `chan_fanout_select`

## Matrix expectation

- `run.args` pins example harness execution to scripted mode.
- `*.stdout` files represent `core` variant scripted output.
- `*.ci.stdout` files represent `go_native` variant scripted output.

## Pure-Go parity baseline

```bash
(cd benchmarks/pure_go/pulseforge && go run . --scripted)
(cd benchmarks/pure_go/pulseforge && go test ./...)
(cd benchmarks/pure_go/pulseforge && go test -run '^$' -bench . -benchmem)
```
