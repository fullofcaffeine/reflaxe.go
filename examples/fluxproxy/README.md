# fluxproxy

Reverse-proxy style flagship app with deterministic policy contract and profile/variant benchmark lanes.

## Why this example exists

- Demonstrates a realistic service-style workload with deterministic policy assertions.
- Shows separation between profile contract (`portable` vs `metal`) and app capability variant (`core` vs `go_native`).
- Provides benchmark-ready lanes for generated-vs-handwritten Go comparisons.

## Terms used in this README

- `go_native` (app variant): compile-time runtime-adapter choice that uses Go-first execution paths (worker/channel/select dispatch). It is **not** a compiler profile.
- `hot path`: the code that runs most often during normal request flow (the highest-impact place to optimize).

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

## Portable vs metal in practice

`metal` is not required for good Go output. This app includes `metal` lanes so
you can test explicit Go-native runtime adapters under stricter checks.

| Profile | Choose this when... | What you get | What to watch for |
| --- | --- | --- | --- |
| `portable` | You want proxy/domain code to remain cross-target friendly and easy to reuse. | Stable portable behavior and the safest default for shared code, while still allowing safe Go-shaped optimizations. | Frequently-executed `go_native` paths may show remaining portable-vs-metal delta; treat that as compiler convergence signal. |
| `metal` | This proxy deployment intentionally owns Go-native APIs and you want stricter compile-time constraints. | Typed specialization in `go_native` paths and fail-fast checks when typed specialization is impossible. | You need clearer static types (avoid loose `Dynamic`/`Any` in hot paths), and generated Go can grow due to specialized helpers. |

Both profiles preserve the same proxy/policy behavior and scripted outputs. The practical difference is optimization aggressiveness in Go-native lanes.

Practical rule for this app:

- Start with `portable` for shared policy/domain modules.
- Move to `metal` when this service is Go-only and perf data points to `go_native` bottlenecks.
- Expect larger profile deltas in `go_native`; `core` often differs less.

## Variant choices in plain terms

| Variant | What changes in the app | Choose this when |
| --- | --- | --- |
| `core` | Uses simple loop-based dispatch in runtime adapters. | You want the clearest portable reference lane. |
| `go_native` | Uses worker fanout with channels/select helpers in runtime adapters. | You are testing/tuning Go-first execution paths and want that benchmark lane. |

`go_native` is a compile-time app variant (`-D fluxproxy_variant_go_native`), not a compiler profile.

`core` is effectively the portable-equivalent app path. Compiling `core` with the `metal` profile is still useful for strict-boundary verification and migration readiness, but typically yields smaller code-shape differences than `go_native`.

## Generated Go Highlights

Committed generated trees under `generated/portable` and `generated/metal` currently track the `core` lane:

- runtime adapter code: `generated/portable/module_app_runtime_coreruntime.go`
- build constants: `generated/portable/module_app_runtime_buildconfig.go`
- variant selection path: `generated/portable/module_app_runtime_runtimefactory.go`
- entry orchestration: `generated/portable/main.go`

For `go_native` generated Go inspection, compile the `*.ci.hxml` lanes and inspect `out_portable_ci` / `out_metal_ci`.

Generated diff commands:

```bash
diff -ru generated/portable generated/metal
diff -ru out_portable_ci out_metal_ci
```

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

## Related docs

- `docs/profiles.md`
- `docs/profile-semantics-guide.md`
- `docs/examples-matrix.md`
- `docs/benchmark-methodology-apps.md`
