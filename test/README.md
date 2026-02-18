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

Force stdlib sweep on a shard:

```bash
python3 test/run-ci.py --chunk 0/4 --force-stdlib-sweep
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

## Strict examples mode

Snapshots compile with `-D reflaxe_go_strict_examples` so app/test code cannot rely on raw `__go__` escape hatches.

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
