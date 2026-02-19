# Snapshot Harness

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

Skip semantic diff stage:

```bash
python3 test/run-ci.py --skip-semantic-diff
```

Force semantic diff on a shard:

```bash
python3 test/run-ci.py --chunk 0/4 --force-semantic-diff
```

Force examples stage on a focused run:

```bash
python3 test/run-ci.py --changed --force-examples
```

## List snapshots

```bash
python3 test/run-snapshots.py --list
```

## Run one snapshot

```bash
python3 test/run-snapshots.py --case core/hello_trace
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

Result categories:

- `PASS`: module compiled (and optionally passed `go test`)
- `EXPECTED_MISSING`: compile-time `Type not found` matched policy for current Haxe version
- `FAIL`: non-policy failure
- `UNEXPECTED_PRESENT`: module compiled even though policy says expected-missing (policy drift to investigate)

Strict mode exits non-zero when any module fails:

```bash
python3 test/run-upstream-stdlib-sweep.py --strict
```

Strict mode + generated Go build checks:

```bash
python3 test/run-upstream-stdlib-sweep.py --strict --go-test
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

## Semantic differential harness

Compare runtime behavior between Haxe reference execution (`--interp`) and `reflaxe.go` generated output (`portable` profile):

```bash
python3 test/run-semantic-diff.py
```

List cases:

```bash
python3 test/run-semantic-diff.py --list
```

Run only changed semantic cases:

```bash
python3 test/run-semantic-diff.py --changed
```

## Examples matrix

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
Harness compile steps also force `-D go_no_build`, then run explicit `go test`/`go run` checks to keep stage ownership deterministic.

## Profile contract checks

Negative snapshot cases validate profile policy:

- conflicts (`portable` + `metal`) fail
- invalid profile values fail
- removed `idiomatic` value and alias fail fast

Supported profile selector values are:

- `portable`
- `gopher`
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
- Snapshot runs are process-locked to avoid concurrent `out/` races; tune with `--lock-timeout <seconds>` (or `0` for fail-fast).
