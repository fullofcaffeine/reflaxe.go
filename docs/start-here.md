# Start Here

## What this target does

`reflaxe.go` compiles Haxe to a generated Go module (`go_output`) with `main.go` plus per-module `module_*.go` files, copies runtime support under `out/hxrt`, and runs `go build` by default (unless disabled).

## First successful run

1. Install toolchain deps:

```bash
npm install
```

Install repo pre-commit hooks (recommended):

```bash
npm run hooks:install
```

2. Run snapshots:

```bash
python3 test/run-snapshots.py
```

3. Run CI entrypoint:

```bash
python3 test/run-ci.py
```

4. Run examples matrix:

```bash
python3 test/run-examples.py
```

5. Run perf harness + release visibility checks:

```bash
npm run test:perf:go
# optional local CI parity: hard-fail if metal exceeds budget
GO_PERF_ENFORCE_METAL_BUDGET=1 npm run test:perf:go
npm run test:perf:hxrt-selective
# optional local CI parity: hard-fail if selective runtime exceeds budget
GO_HXRT_SLICE_ENFORCE=1 npm run test:perf:hxrt-selective
npm run test:perf:apps
# optional local CI parity: hard-fail if app metal ratios exceed budget
GO_APP_PERF_ENFORCE_METAL_BUDGET=1 npm run test:perf:apps
# baseline management for app harness
npm run test:perf:apps:update-baseline
npm run release:status
```

Perf harness lanes now include hello/array/atomic plus channel/map/generic/string/virtual/select microbench cases and TUI profile spread.
Flagship app harness output is written under `.cache/perf-apps/results/` (`current.json`, `comparison.json`, `summary.md`, `raw_metrics.tsv`) and uses `scripts/ci/perf/app-profile-baseline.json`.

## Scaffold a new project

```bash
npm run dev:new-project -- ./my_haxe_go_app
cd ./my_haxe_go_app
npm install
npm run setup
npm run hx:run
```

## Task and package manager model

- Generated project dependencies/build/test use the Go toolchain directly (`go mod`, `go run`, `go test`, `go build`).
- Compiler/dev orchestration uses `npm` scripts so workflow stays consistent with `haxe.rust` and `haxe.elixir.codex`.
- Direct hxml compiles run backend `go build` by default; use `-D go_no_build` for codegen-only flows.

## Why stdlib ownership is hybrid

`reflaxe.go` intentionally uses a hybrid stdlib model:

- runtime helpers in `hxrt`
- compiler-owned shims for metadata/profile-sensitive behavior
- staged `std/_std` migration surfaces

This pattern is not unique to Go, but Go's exception/string/dynamic semantics and profile policy model make it especially practical here. Details: `docs/stdlib-shim-rationale.md`.

Portable target objective is full portable-eligible Haxe stdlib parity (excluding target-specific namespaces, with explicit module-by-module tracking and closure plan):
`docs/portable-stdlib-parity-program.md`.

Quick compile+go action from this repo:

```bash
npm run dev:hx -- --project examples/tui_todo --profile portable --action run
```

The wrapper resolves `compile*.hxml`, runs Haxe, resolves `-D go_output=...`, then executes the selected Go action.

Reference examples for this phase:

- `examples/worker_pool_select` (worker pool + select-style channel ops)
- `examples/interop_smoke` (`fmt`/`time`/`context`/`net/http` typed interop)
- `examples/pulseforge` (flagship observability app with profile + variant behavior matrices)
- `examples/fluxproxy` (flagship proxy app with profile + variant behavior matrices)
- `docs/examples-matrix.md` (benchmark harness commands and artifact paths)

## Profile selection

Set via:

```bash
-D reflaxe_go_profile=portable|metal
```

- `portable` (default): choose this for portability and lowest migration risk.
- `metal` (experimental): choose this when portable abstractions are not enough and you need typed low-level interop with strict boundaries.
- `@:goMetal`: optional lane metadata for portable builds when you want module-by-module metal-clean enforcement.

Compatibility note:

- `idiomatic` is removed and intentionally fails fast; use `portable` instead.
- `gopher` is removed and intentionally fails fast; use `portable` instead.

## Strict policy knobs

- `-D reflaxe_go_strict_examples`: forbids raw `__go__` in repo examples/snapshots.
- `-D reflaxe_go_strict`: forbids raw `__go__` in app project sources.
- `metal` enables strict mode by default for app-side injection boundaries.
- `-D reflaxe_go_metal_allow_fallback`: allows typed-specialization fallback in metal builds (instead of hard error) and disables metal strict-by-default boundary behavior.
- `-D reflaxe_go_portable_native_policy=warn|error|off`: controls `go.*` diagnostics in portable builds (`warn` default, `error` recommended in CI/release).
- `-D reflaxe_go_portable_native_allow=<csv>`: optional module-prefix allowlist for sanctioned native adapter modules in portable builds.
- `-D reflaxe_go_line_directives`: opt-in `//line` source mapping directives in generated user functions.

## Contract/runtime report knobs

- `-D reflaxe_go_contract_report`: writes `profile_contract.json` + `profile_contract.md` to output root.
- `-D reflaxe_go_runtime_plan_report`: writes `hxrt_plan.json` + `hxrt_plan.md` to output root.

## GitHub CI harness

- `.github/workflows/ci-harness.yml`: integrated quality + security gates (`test:ci`, gitleaks, dependency audit) plus semantic-release on `master`.
- `ci-harness.yml` includes flagship app benchmark job `perf-apps` (PR/push/manual/weekly schedule) with `go-app-perf-results` artifact upload.
- `.github/workflows/security-static-analysis.yml`: dependency review/codeql and scheduled security analysis.
- `.github/workflows/examples-artifacts.yml`: builds and uploads examples binary artifacts on `master`, and on tag pushes publishes release assets via a separate release job.

## Related docs

- `docs/hxrt-runtime.md`
- `docs/hxrt-selective-runtime.md`
- `docs/portable-canonical-contract.md`
- `docs/profiles.md`
- `docs/profile-semantics-guide.md`
- `docs/profile-auto-spike.md`
- `docs/go-concurrency-interop-guide.md`
- `docs/flagship-apps-plan.md`
- `docs/feature-support-matrix.md`
- `docs/examples-matrix.md`
- `docs/benchmark-methodology-apps.md`
- `docs/known-gaps.md`
- `docs/release-visibility.md`
- `docs/stdlib-shim-rationale.md`
- `docs/portable-stdlib-parity-program.md`
- `docs/defines-reference.md`
- `docs/profile-admission-criteria.md`
- `docs/compiler-target-template.md`
- `docs/snapshot-policy.md`
- `SECURITY.md`
