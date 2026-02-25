# fluxproxy

Reverse-proxy style flagship app with deterministic policy contract and profile/variant benchmark lanes.

## Architecture

| Layer | Files | Responsibility |
| --- | --- | --- |
| Entry + UX | `Main.hx`, `InteractiveCli.hx` | Scripted contract mode and interactive command session |
| Contract harness | `Harness.hx` | Deterministic baseline request set + policy assertions |
| Domain core | `app/core/*` | Routing, retries, rate-limit, breaker, per-route aggregation |
| Runtime adapters | `app/runtime/CoreRuntime.hx`, `app/runtime/GoNativeRuntime.hx` | Variant-specific dispatch path (`core` vs `go_native`) |
| Build selectors | `app/runtime/BuildConfig.hx`, `app/runtime/RuntimeFactory.hx` | Compile-time profile/variant constants and adapter selection |

## Compile Matrix

| Lane | Profile | Variant | HXML |
| --- | --- | --- | --- |
| portable/core | `portable` | `core` | `compile.portable.hxml` |
| metal/core | `metal` | `core` | `compile.metal.hxml` |
| portable/go_native | `portable` | `go_native` | `compile.portable.ci.hxml` |
| metal/go_native | `metal` | `go_native` | `compile.metal.ci.hxml` |

```bash
haxe compile.portable.hxml
haxe compile.metal.hxml
haxe compile.portable.ci.hxml
haxe compile.metal.ci.hxml
```

## Run

Core lanes (`out_portable`, `out_metal`) support interactive and scripted modes:

```bash
(cd out_portable && go run . --scripted)
(cd out_portable && go run . status)
(cd out_portable && go run . ingest /v1/items 45 200 status)
```

Build binary form:

```bash
(cd out_portable && go build -o fluxproxy_portable . && ./fluxproxy_portable --scripted)
(cd out_metal && go build -o fluxproxy_metal . && ./fluxproxy_metal --scripted)
```

Modes:

- scripted: deterministic contract output (`--scripted`)
- interactive: command session (`help`, `profile`, `status`, `ingest`, `reset`, `scripted`)

## Profile Behavior Matrix

| Profile | Contract intent | Boundary policy | Runtime package |
| --- | --- | --- | --- |
| `portable` | Semantic baseline lane | Strict examples boundary enabled via compile files | Generated Go + `hxrt` |
| `metal` | Go-first/perf lane under explicit opt-in | Strict examples boundary enabled; metal compiler profile selected | Generated Go + `hxrt` |

Both profiles preserve the same proxy/policy contract. Differences are code shape and optimization policy, not app behavior removal.

## Variant Behavior Matrix

| Variant | Capability id | Strategy | Notes |
| --- | --- | --- | --- |
| `core` | `loop_dispatch` | Deterministic loop-based dispatch path | Baseline lane with simple flow |
| `go_native` | `worker_chan_fanout` | Worker fan-out with channels/select helpers | Go-first dispatch lane |

`go_native` is a compile-time app variant (`-D fluxproxy_variant_go_native`), not a compiler profile.

`core` is effectively the portable-equivalent app path. Compiling `core` with the `metal` profile is still useful for strict-boundary verification and migration readiness, but typically yields smaller code-shape differences than `go_native`.

## Generated Go Highlights

Committed generated trees under `generated/portable` and `generated/metal` currently track the `core` lane:

- runtime adapter code: `generated/portable/module_app_runtime_coreruntime.go`
- build constants: `generated/portable/module_app_runtime_buildconfig.go`
- variant selection path: `generated/portable/module_app_runtime_runtimefactory.go`
- entry orchestration: `generated/portable/main.go`

For `go_native` generated Go inspection, compile the `*.ci.hxml` lanes and inspect `out_portable_ci` / `out_metal_ci`.

## Matrix Expectations

- `run.args` pins example harness execution to scripted mode.
- `*.stdout` files represent `core` variant scripted output.
- `*.ci.stdout` files represent `go_native` variant scripted output.

## Benchmarking

`pure_go` is handwritten parity baseline code in `benchmarks/pure_go/fluxproxy` and does not use `hxrt`.

```bash
(cd benchmarks/pure_go/fluxproxy && go run . --scripted)
(cd benchmarks/pure_go/fluxproxy && go test ./...)
(cd benchmarks/pure_go/fluxproxy && go test -run '^$' -bench . -benchmem)
```

Run cross-lane app harness:

```bash
npm run test:perf:apps
```

Harness artifacts:

- `.cache/perf-apps/results/current.json`
- `.cache/perf-apps/results/comparison.json`
- `.cache/perf-apps/results/summary.md`
- `.cache/perf-apps/results/raw_metrics.tsv`
- `.cache/perf-apps/results/warnings.txt`
- `.cache/perf-apps/results/hard_failures.txt`

Methodology and fairness constraints: `docs/benchmark-methodology-apps.md`.

## Known Tradeoffs

- `go_native` is compile-time selected; it is not switchable at runtime inside one binary.
- Committed generated trees are currently `core` snapshots; `go_native` codegen is validated via CI compile lanes and perf harness runs.
- App-level allocation and binary-size overhead vs pure-Go is measurable and expected while `hxrt` ownership and stdlib shims remain in active optimization work.
- `metal` is not a "fewer lines" mode. In `go_native` lanes it can emit additional specialized typed helpers (`go__concurrency_*`, `go__result_*`) to reduce dynamic paths in hot code, which may increase generated LOC while improving runtime behavior.
