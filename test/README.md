# Snapshot Harness

## Canonical `_std` layout contract

What it is: a source/package layout audit plus a behavior test for Reflaxe
standard-library override selection.

Why it exists: a source checkout can appear healthy while late macro classpath
mutation hides a broken installed package. The contract separately models
ordinary source overrides under `std/go/_std/**/*.hx` and package-generated
`src/**/*.cross.hx`, then compiles and runs both against an upstream shadow.

How it works:

1. `scripts/ci/canonical_stdlib_layout_check.py` requires the live source tree
   to keep the canonical layout.
2. `test/canonical_std_layout_status.json` declares that source layout
   `required-green`; legacy support roots, source `.cross.hx`, or classpath
   regressions fail immediately.
3. `test/test_canonical_std_layout_contract.py` rejects adversarial source and
   package layouts, then proves override, support, runtime-binding, and public
   facade resolution through source and staged-package runtime output.
4. Ordinary support/facade modules must remain `.hx` in a package; only
   canonical `std/go/_std` overrides map to `.cross.hx`.

Run the integrated contract:

```bash
npm run test:canonical-std-layout
```

Inspect the live source audit directly:

```bash
python3 scripts/ci/canonical_stdlib_layout_check.py --source-root .
```

## Run all snapshots

```bash
python3 test/run-snapshots.py
```

## CI entrypoint

```bash
python3 test/run-ci.py
```

Run one CI shard (skip stdlib sweep by default for chunked runs):

```bash
python3 test/run-ci.py --chunk 0/4
```

Tune snapshot lock wait in CI wrapper:

```bash
python3 test/run-ci.py --snapshot-lock-timeout 0
```

Force stdlib sweep on a shard:

```bash
python3 test/run-ci.py --chunk 0/4 --force-stdlib-sweep
```

Force full portable-eligible stdlib sweep on a focused run:

```bash
python3 test/run-ci.py --changed --force-stdlib-full-sweep
```

Skip full portable-eligible stdlib sweep stage:

```bash
python3 test/run-ci.py --skip-stdlib-full-sweep
```

Include `go test` in full portable-eligible stdlib sweep stage:

```bash
python3 test/run-ci.py --stdlib-full-go-test
```

Skip semantic diff stage:

```bash
python3 test/run-ci.py --skip-semantic-diff
```

Skip semantic-diff optimizer matrix stage:

```bash
python3 test/run-ci.py --skip-semantic-diff-optimizer-matrix
```

Skip lane semantic diff stage:

```bash
python3 test/run-ci.py --skip-semantic-diff-lanes
```

Skip stdlib governance stage:

```bash
python3 test/run-ci.py --skip-stdlib-governance
```

Skip optimizer matrix stage:

```bash
python3 test/run-ci.py --skip-optimizer-matrix
```

Skip auto planner report-schema stage:

```bash
python3 test/run-ci.py --skip-auto-planner-schema
```

Skip portable allowlist stage:

```bash
python3 test/run-ci.py --skip-portable-allowlist
```

Skip portable conformance stage:

```bash
python3 test/run-ci.py --skip-portable-conformance
```

Skip portable parity closure summary stage:

```bash
python3 test/run-ci.py --skip-portable-parity-closure
```

Skip family std sync/verify stage:

```bash
python3 test/run-ci.py --skip-family-stdlib-bootstrap
```

Force semantic diff on a shard:

```bash
python3 test/run-ci.py --chunk 0/4 --force-semantic-diff
```

Force semantic-diff optimizer matrix on a focused run:

```bash
python3 test/run-ci.py --changed --force-semantic-diff-optimizer-matrix
```

Force lane semantic diff on a shard:

```bash
python3 test/run-ci.py --chunk 0/4 --force-semantic-diff-lanes
```

Force portable conformance on a focused run:

```bash
python3 test/run-ci.py --changed --force-portable-conformance
```

Force examples stage on a focused run:

```bash
python3 test/run-ci.py --changed --force-examples
```

Skip metal-only example boundary import policy stage:

```bash
python3 test/run-ci.py --skip-metal-example-boundary
```

Force optimizer matrix stage on a focused run:

```bash
python3 test/run-ci.py --changed --force-optimizer-matrix
```

Force auto planner report-schema stage on a focused run:

```bash
python3 test/run-ci.py --changed --force-auto-planner-schema
```

## List snapshots

```bash
python3 test/run-snapshots.py --list
```

## Run one snapshot

```bash
python3 test/run-snapshots.py --case core/hello_trace
```

## Check metal-only example import boundary directly

```bash
python3 test/run-metal-example-boundary.py
npm run test:examples:metal-boundary
```

## Run with parallel workers

```bash
python3 test/run-snapshots.py --jobs 4
```

## Update intended outputs

```bash
python3 test/run-snapshots.py --update
```

## Bless only changed intended files

```bash
python3 test/run-snapshots.py --bless
```

## Re-run previous failures

```bash
python3 test/run-snapshots.py --failed
```

## Run only changed snapshot cases

```bash
python3 test/run-snapshots.py --changed
```

## Run a deterministic CI shard

```bash
python3 test/run-snapshots.py --chunk 0/4
```

## Run optimizer-plan matrix snapshots

Runs the orthogonal optimizer axis matrix (string fastpath preset on/off and concurrency fastpath default/override combinations):

```bash
npm run test:optimizer:matrix
```

## Run optimizer semantic-diff matrix

Runs semantic parity checks across the same optimizer toggle matrix:

```bash
npm run test:semantic-diff:optimizer-matrix
```

## Run auto planner report schema gate

Validates deterministic planner/runtime/contract artifact schemas and required keys, including optimizer counter invariants for portable+auto typed-lowering vs fallback snapshot contracts (`goCollections*` / `goResult*`) and capability-level `autoLoweringCapabilities` invariants:

```bash
npm run test:auto-planner:schema
```

## Upstream stdlib sweep

Run curated upstream stdlib module compile checks:

```bash
python3 test/run-upstream-stdlib-sweep.py
```

Module list source:

```text
test/upstream_std_modules.txt
```

Expected-missing policy source (version-aware):

```text
test/upstream_std_expected_missing.json
```

Expected-unavailable policy source (version-aware):

```text
test/upstream_std_expected_unavailable.json
```

Policy rules can optionally set `"stage": "compile" | "go_test" | "any"` to scope expected failures.

Result categories:

- `PASS`: module compiled (and optionally passed `go test`)
- `EXPECTED_POLICY`: compile-time failure matched active expected-missing/expected-unavailable policy for current Haxe version
- `FAIL`: non-policy failure
- `UNEXPECTED_PRESENT`: module compiled even though policy says expected failure (policy drift to investigate)

Strict mode exits non-zero when any module fails:

```bash
python3 test/run-upstream-stdlib-sweep.py --strict
```

Strict mode + generated Go build checks:

```bash
python3 test/run-upstream-stdlib-sweep.py --strict --go-test
```

Run full runtime-eligible inventory:

```bash
python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_full.txt --strict
```

Run full inventory with generated Go compile checks:

```bash
python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_full.txt --strict --go-test
```

Run one module:

```bash
python3 test/run-upstream-stdlib-sweep.py --module haxe.Json
```

Run the broader parity-gap probe inventory list:

```bash
python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_gap_probe.txt --go-test
```

Disable expected-missing policy classification (raw failures only):

```bash
python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_gap_probe.txt --go-test --no-expected-missing-policy
```

Disable expected-unavailable policy classification (raw failures only):

```bash
python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_gap_probe.txt --go-test --no-expected-unavailable-policy
```

Equivalent npm commands:

```bash
npm run test:stdlib-sweep:full
npm run test:stdlib-sweep:full:go-test
```

## Portable stdlib inventory ledger

Validate the machine-readable portable-eligible stdlib inventory:

```bash
python3 test/run-portable-stdlib-inventory.py
```

Regenerate inventory after policy/module changes:

```bash
python3 test/run-portable-stdlib-inventory.py --update
```

Promotion overrides (`compile-only -> snapshot -> semantic-diff`) are declared in:

```text
test/portable_parity_promotions.json
```

Primary artifacts:

```text
test/portable_stdlib_inventory.json
test/portable_parity_promotions.json
test/.test-cache/portable_stdlib_inventory_summary.json
test/.test-cache/portable_stdlib_inventory_summary.md
```

## Portable allowlist (tiered contract set)

Validate the tiered portable contract allowlist:

```bash
python3 test/run-portable-allowlist.py
```

Equivalent npm command:

```bash
npm run test:portable-allowlist
```

Primary artifacts:

```text
test/portable_allowlist.json
test/.test-cache/portable_allowlist_summary.json
test/.test-cache/portable_allowlist_summary.md
```

## Portable Tier1 conformance suite

Run the dedicated Tier1 portable conformance seed (mapped from `portable_allowlist` modules to deterministic semantic-diff cases):

```bash
python3 test/run-portable-conformance.py
npm run test:portable-conformance
```

Ownership mapping reference:

- `docs/portable-module-mapping-contract.md`

List module-to-case mapping:

```bash
python3 test/run-portable-conformance.py --list
```

Run only selected Tier1 module checks:

```bash
python3 test/run-portable-conformance.py --module haxe.Json --module sys.net.Socket
```

Primary artifacts:

```text
test/portable_conformance_tier1.json
test/.test-cache/portable_conformance_tier1_summary.json
test/.test-cache/portable_conformance_tier1_summary.md
```

## Portable parity closure summary

Generate full-module parity closure summary (remaining non-semantic-diff surfaces + promotion/policy queue):

```bash
python3 test/run-portable-parity-closure.py
npm run test:portable-parity-closure
```

List non-semantic-diff modules with next promotion or policy step:

```bash
python3 test/run-portable-parity-closure.py --list-blockers
```

`--list-blockers` prints the automated promotion/policy queue with module-level `next_step`, `closure_policy`, and `actionable` guidance.
`actionable: true` means there is still hidden parity work.
`actionable: false` means the surface is intentionally policy-locked, such as target-sensitive snapshot evidence or an explicit target-conditional exclusion.

Primary artifacts:

```text
test/.test-cache/portable_parity_closure_summary.json
test/.test-cache/portable_parity_closure_summary.md
```

## Family std sync and bootstrap validation

Validate canonical<->family sync state plus bootstrap skeleton schemas/manifest:

```bash
python3 tools/family_std_sync.py --mode verify
npm run test:family-stdlib-sync
npm run test:family-stdlib-bootstrap
```

Sync commands:

```bash
npm run family:stdlib:export
npm run family:stdlib:import
```

Primary artifacts:

```text
family/family_std_pin.json
family/reflaxe.family.std/MANIFEST.v1.txt
family/reflaxe.family.std/provenance/stdlib-provenance-ledger.schema.json
family/reflaxe.family.std/provenance/upstream-boundary-policy.v1.json
test/.test-cache/family_std_dual_run_report.json
test/.test-cache/family_std_dual_run_report.md
```

## Stdlib governance guards (boundary + provenance + migration ownership)

Validate the upstream boundary plus the per-file provenance, ownership,
destination, and compiler-shim audit:

```bash
npm run test:stdlib:governance
```

Run checks individually:

```bash
npm run test:stdlib:boundary
npm run test:stdlib:provenance
npm run test:stdlib:migration-ledger
```

Primary policy artifacts:

```text
docs/stdlib-provenance-ledger.json
scripts/ci/upstream-stdlib-boundary-check.js
scripts/ci/stdlib-provenance-ledger-check.js
test/test_stdlib_migration_ledger_contract.py
```

## Compiler debt baseline and ratchet

What it is: a deterministic inventory of target-owned raw Go AST construction,
Haxe `Dynamic` / `Any`, Go `reflect` / `unsafe` imports and selectors, and named compiler shim
boundaries.

Why it exists: some dynamic or reflective behavior is required by Haxe
semantics, while raw string emission is migration debt. A single global count
would hide ownership transfers and profile-specific generated-output growth.

How it works:

1. `test/run-compiler-debt-ratchet.py` scans source, runtime, and committed
   portable/metal example output without recording absolute paths.
2. `test/compiler_debt_policy.json` classifies every current location as
   `required` or `avoidable`, links it to a Why / What / How exception, and
   sets a per-file/per-context ceiling.
3. Reductions pass. A new location or an increase fails until the underlying
   debt is removed or the reviewed policy and baseline are intentionally
   updated.
4. Current reports are written to `.cache/compiler-debt/report.json` and
   `.cache/compiler-debt/report.md`.

Run the gate:

```bash
npm run test:compiler-debt
```

After an intentional, reviewed boundary change, regenerate ceilings and
inspect the full diff:

```bash
python3 test/run-compiler-debt-ratchet.py --update-baseline
git diff -- test/compiler_debt_policy.json
```

## Semantic differential harness

Compare runtime behavior between Haxe reference execution (`--interp`) and `reflaxe.go` generated output (`portable` profile):

Portable semantic contract reference:

- `docs/portable-semantics-v1.md`
- Option/Result admission and native-boundary contract:
  `docs/portable-option-result-contract.md`
- guide: `docs/semantic-diff-guide.md`

```bash
python3 test/run-semantic-diff.py
```

Run the lane-focused suite (`@:goMetal` lane-clean fixtures):

```bash
python3 test/run-semantic-diff.py --suite lanes
npm run test:semantic-diff:lanes
```

List cases:

```bash
python3 test/run-semantic-diff.py --list
```

Run only changed semantic cases:

```bash
python3 test/run-semantic-diff.py --changed
```

Run semantic diff with fail-fast lock behavior:

```bash
python3 test/run-semantic-diff.py --lock-timeout 0
```

## Go policy-preset perf harness

Collect soft-budget benchmark ratios for the `portable|metal` compatibility
presets versus pure-Go microcases (`hello`, `array`, `atomic`, `channel`, `map`,
`generic`, `string`, `string_instance`, `virtual`, `select`) plus the historical
`examples/tui_todo` preset spread:

```bash
bash scripts/ci/perf-go-profiles.sh
```

Enforce metal-compatibility-preset budget regressions as hard failures (the
portable-default preset remains warning-only):

```bash
GO_PERF_ENFORCE_METAL_BUDGET=1 bash scripts/ci/perf-go-profiles.sh
```

Enable portable-vs-metal startup delta hard budget checks (selected cases only):

```bash
GO_PERF_ENFORCE_DELTA_BUDGET=1 bash scripts/ci/perf-go-profiles.sh
```

Tune hard-fail budgets if needed:

```bash
GO_PERF_ENFORCE_METAL_BUDGET=1 GO_PERF_METAL_RUNTIME_FAIL_PCT=90 bash scripts/ci/perf-go-profiles.sh
```

Tune delta budget cases and thresholds if needed:

```bash
GO_PERF_ENFORCE_DELTA_BUDGET=1 GO_PERF_DELTA_CASES=string,string_instance,select,channel GO_PERF_DELTA_WARN_PCT=12 GO_PERF_DELTA_FAIL_PCT=20 bash scripts/ci/perf-go-profiles.sh
```

Disable portable concurrency fastpath for A/B perf checks:

```bash
GO_PERF_PORTABLE_CONCURRENCY_FASTPATH=0 bash scripts/ci/perf-go-profiles.sh
```

Change the base `hxrt` runtime slice used by microbench builds:

```bash
GO_PERF_HXRT_FEATURES=core,string,print bash scripts/ci/perf-go-profiles.sh
```

`GO_PERF_HXRT_FEATURES` enables selective runtime copying for the microbench
cases. The default keeps the common print/string helpers in the measured
output, while the compiler can still infer extra case-specific files such as
atomic or concurrency helpers. This prevents unrelated runtime subsystems from
making every microbench binary look slower or larger.

Tune atomic workload/loop stability if needed:

```bash
GO_PERF_ATOMIC_WORK=400000 GO_PERF_ATOMIC_ITERS=80 bash scripts/ci/perf-go-profiles.sh
```

Tune select workload/loop stability if needed:

```bash
GO_PERF_SELECT_WORK=80000 GO_PERF_SELECT_ITERS=80 bash scripts/ci/perf-go-profiles.sh
```

Tune string and virtual workloads if needed:

```bash
GO_PERF_STRING_WORK=20000 GO_PERF_STRING_ITERS=70 GO_PERF_STRING_INSTANCE_WORK=10000 GO_PERF_STRING_INSTANCE_ITERS=50 GO_PERF_VIRTUAL_WORK=150000 GO_PERF_VIRTUAL_ITERS=90 bash scripts/ci/perf-go-profiles.sh
```

Regenerate baseline:

```bash
bash scripts/ci/perf-go-profiles.sh --update-baseline
```

Baseline source:

```text
scripts/ci/perf/go-profile-baseline.json
```

Result artifacts:

```text
.cache/perf-go/results/current.json
.cache/perf-go/results/comparison.json
.cache/perf-go/results/summary.md
.cache/perf-go/results/warnings.txt
.cache/perf-go/results/hard_failures.txt
```

Delta interpretation:

- `current.json` -> `derived.portableMetalDeltaRatios.<case>.startupRatio` reports portable/metal ratio per microcase (`1.000` means equal startup overhead ratio).
- `comparison.json` -> `deltaWarningCount` / `deltaHardFailureCount` and `deltaCases` summarize selected-case drift checks.
- `summary.md` includes a dedicated `Portable-vs-metal Delta` table plus `Delta Hard-Fail Candidates`.

## Flagship app perf harness

Collect app-level profile ratios for `pulseforge` and `fluxproxy` (portable/metal vs pure-go, variants `core` and `go_native`):

```bash
bash scripts/ci/perf-apps.sh
```

Enforce metal hard budgets:

```bash
GO_APP_PERF_ENFORCE_METAL_BUDGET=1 bash scripts/ci/perf-apps.sh
```

Enable portable-vs-metal delta hard budget checks:

```bash
GO_APP_PERF_ENFORCE_DELTA_BUDGET=1 bash scripts/ci/perf-apps.sh
```

Tune delta cases and thresholds:

```bash
GO_APP_PERF_ENFORCE_DELTA_BUDGET=1 GO_APP_PERF_DELTA_CASES=pulseforge:go_native,fluxproxy:go_native GO_APP_PERF_DELTA_WARN_PCT=12 GO_APP_PERF_DELTA_FAIL_PCT=20 bash scripts/ci/perf-apps.sh
```

Regenerate app baseline:

```bash
bash scripts/ci/perf-apps.sh --update-baseline
```

App baseline source:

```text
scripts/ci/perf/app-profile-baseline.json
```

App result artifacts:

```text
.cache/perf-apps/results/current.json
.cache/perf-apps/results/comparison.json
.cache/perf-apps/results/summary.md
.cache/perf-apps/results/warnings.txt
.cache/perf-apps/results/hard_failures.txt
```

App delta interpretation:

- `current.json` -> `derived.portableVsMetal` reports portable/metal deltas per `app+variant`.
- `comparison.json` -> `deltaCases`, `deltaWarningCount`, and `deltaHardFailureCount` expose selected-case budget outcomes.
- `summary.md` includes `Portable-vs-Metal Deltas` and separate `Metal Hard-Fail` vs `Delta Hard-Fail` sections.

## HXRT selective runtime perf/size harness

Collect selective-vs-full runtime footprint metrics (`runtime file count`, `runtime source bytes`, `binary bytes`) for representative portable+metal cases:

```bash
bash scripts/ci/perf-hxrt-selective.sh
```

Enforce selective runtime budget regressions as hard failures:

```bash
GO_HXRT_SLICE_ENFORCE=1 bash scripts/ci/perf-hxrt-selective.sh
```

Tune hard-fail budgets if needed:

```bash
GO_HXRT_SLICE_ENFORCE=1 GO_HXRT_SLICE_MAX_SOURCE_PCT=2 GO_HXRT_SLICE_MAX_BINARY_PCT=5 bash scripts/ci/perf-hxrt-selective.sh
```

Regenerate baseline:

```bash
bash scripts/ci/perf-hxrt-selective.sh --update-baseline
```

Baseline source:

```text
scripts/ci/perf/hxrt-selective-baseline.json
```

Result artifacts:

```text
.cache/perf-hxrt-selective/results/current.json
.cache/perf-hxrt-selective/results/comparison.json
.cache/perf-hxrt-selective/results/summary.md
```

`summary.md` includes drift columns (source/binary delta drift) relative to `scripts/ci/perf/hxrt-selective-baseline.json`.

## Examples matrix

Examples are QA contracts, not loose demos. Full policy:
`docs/examples-qa-contract.md`.

Run all example/profile cases:

```bash
python3 test/run-examples.py
```

Run only changed examples:

```bash
python3 test/run-examples.py --changed
```

Refresh committed generated example outputs:

```bash
python3 test/run-examples.py --bless-generated
```

Validate generated trees only (after compiling examples):

```bash
python3 scripts/examples/sync-generated.py
```

## Strict examples mode

Snapshots compile with `-D reflaxe_go_strict_examples` so app/test code cannot rely on raw `__go__` escape hatches.
Harness compile steps normally force `-D go_no_build`, then run explicit
`go test`/`go run` checks to keep stage ownership deterministic. A negative
fixture may include a `backend-build` marker when its contract specifically
tests the backend-owned build phase; such a fixture must expect compilation to
fail and must use a deterministic local command.

## Profile contract checks

Negative snapshot cases validate profile policy:

- conflicts (`portable` + `metal`) fail
- invalid profile values fail
- removed `gopher` value and alias fail fast
- removed `idiomatic` value and alias fail fast

Supported profile selector values are:

- `portable`
- `metal`

## Snapshot shape policy

Snapshots are canonicalized against the **post-Reflaxe optimized AST** output, not the raw pre-optimization typed tree.

- Behavioral correctness is guarded by `go test` and optional `expected.stdout` runtime checks.
- Code-shape stability is guarded by snapshot diffs.
- Optimization-shape sentinel: `core/optimized_ast_policy` (constant folding + boolean simplification).

## Troubleshooting

- Use `KEEP_ARTIFACTS=1` to keep generated `out/` folders on failures.
- Use `--failed` to rerun only the previous failing set.
- Use `--changed` to focus only on touched snapshot cases.
- Use `--bless` to update only changed files and print a gofmt/naming/runtime checklist.
- Snapshot diff failures print deterministic file-set summaries and unified diff excerpts for modified files (helpful for multi-file output cases).
- Snapshot runs are process-locked to avoid concurrent `out/` races; tune with `--lock-timeout <seconds>` (or `0` for fail-fast).
