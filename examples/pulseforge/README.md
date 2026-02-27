# pulseforge

Flagship observability-stream pipeline demo with one Haxe codebase compiled across profile and runtime-variant lanes.

## Why this example exists

- Demonstrates a production-shaped app architecture, not just a toy program.
- Shows profile contract (`portable` vs `metal`) and app variant (`core` vs `go_native`) as separate axes.
- Provides benchmark-ready lanes against handwritten Go baselines.

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

## Profile Behavior Matrix

| Profile | Contract intent | Boundary policy | Runtime package |
| --- | --- | --- | --- |
| `portable` | Semantic baseline lane | Strict examples boundary enabled via compile files | Generated Go + `hxrt` |
| `metal` | Go-first/perf lane under explicit opt-in | Strict examples boundary enabled; metal compiler profile selected | Generated Go + `hxrt` |

Both profiles keep the same domain contract and workload semantics. Differences are code shape and optimization policy, not app feature removal.

## When to choose each profile here

- Choose `portable` when PulseForge core logic must stay aligned with cross-target portable semantics.
- Choose `metal` when Go-native optimization and strict boundary policy are required for this deployment.

## Tradeoffs shown by this example

- `portable` and `metal` can preserve the same app contract while using different codegen strategies.
- `metal` may emit more typed helper code in hot lanes; more LOC does not imply lower quality.
- Profile choice and variant choice are independent decisions.

## Variant Behavior Matrix

| Variant | Capability id | Strategy | Notes |
| --- | --- | --- | --- |
| `core` | `core_loop` | Deterministic loop-based parse/enrich stages | Lowest-risk baseline behavior |
| `go_native` | `chan_fanout_select` | Channel fan-out/fan-in plus `go.Select` helpers | Go-first execution lane |

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
