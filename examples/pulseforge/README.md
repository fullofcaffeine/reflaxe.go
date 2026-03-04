# pulseforge

Flagship observability-stream pipeline demo with one Haxe codebase compiled across profile and runtime-variant lanes.

## Why this example exists

- Demonstrates a production-shaped app architecture, not just a toy program.
- Shows profile contract (`portable` vs `metal`) and app variant (`core` vs `go_native`) as separate axes.
- Provides benchmark-ready lanes against handwritten Go baselines.

## Terms used in this README

- `go_native` (app variant): compile-time runtime-adapter choice that uses Go-first execution paths (channel/select worker fanout). It is **not** a compiler profile.
- `hot path`: the code that runs most often during normal workload (the part where performance changes matter most).

## Architecture

| Layer | Files | Responsibility |
| --- | --- | --- |
| Entry + UX | `Main.hx`, `InteractiveCli.hx` | Scripted contract mode and interactive command session |
| Contract harness | `Harness.hx` | Deterministic baseline workload + assertable output contract |
| Domain core | `app/core/*` | Parse, enrich, aggregate, alert pipeline and report rendering |
| Runtime adapters | `app/runtime/CoreRuntime.hx`, `app/runtime/GoNativeRuntime.hx` | Variant-specific execution strategy (`core` vs `go_native`) |
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
(cd out_portable && go run . ingest edge 8 iad status)
```

Build binary form:

```bash
(cd out_portable && go build -o pulseforge_portable . && ./pulseforge_portable --scripted)
(cd out_metal && go build -o pulseforge_metal . && ./pulseforge_metal --scripted)
```

Modes:

- scripted: deterministic contract output (`--scripted`)
- interactive: command session (`help`, `profile`, `status`, `ingest`, `reset`, `scripted`)

## Portable vs metal in practice

| Profile | Choose this when... | What you get | What to watch for |
| --- | --- | --- | --- |
| `portable` | You want this codebase to stay cross-target friendly and easy to share with other Haxe targets. | Stable portable behavior and the lowest migration risk for shared app/domain code. | In frequently-executed `go_native` paths, generated Go can rely more on generic/runtime helper paths, so peak Go performance may be lower than `metal`. |
| `metal` | This deployment is Go-first and you want stricter compile-time checks plus stronger typed lowering in hot paths. | More aggressive typed specialization in `go_native` paths (`go.Chan`/`go.Select` style flows) and fail-fast checks for unsupported typed specialization cases. | You may need more explicit typing (avoid loose `Dynamic`/`Any` paths), and generated code can be larger because of specialized helpers. |

Both profiles keep the same app behavior contract and scripted outputs. The main difference is how aggressively the compiler optimizes Go-native paths.

Practical rule for this app:

- Start with `portable` for shared domain logic.
- Use `metal` when this service is Go-only and benchmark data shows the `go_native` lane is a bottleneck.
- Expect the biggest profile differences in `go_native`, not `core`.

## Variant choices in plain terms

| Variant | What changes in the app | Choose this when |
| --- | --- | --- |
| `core` | Uses simple loop-based processing for parse/enrich stages. | You want the most straightforward, portable reference behavior. |
| `go_native` | Uses worker fanout with channels/select helpers in runtime adapters. | You are testing/tuning Go-first execution paths and want to benchmark that lane. |

`go_native` is a compile-time app variant (`-D pulseforge_variant_go_native`), not a compiler profile.

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

`pure_go` is handwritten parity baseline code in `benchmarks/pure_go/pulseforge` and does not use `hxrt`.

```bash
(cd benchmarks/pure_go/pulseforge && go run . --scripted)
(cd benchmarks/pure_go/pulseforge && go test ./...)
(cd benchmarks/pure_go/pulseforge && go test -run '^$' -bench . -benchmem)
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

## Related docs

- `docs/profiles.md`
- `docs/profile-semantics-guide.md`
- `docs/examples-matrix.md`
- `docs/benchmark-methodology-apps.md`
