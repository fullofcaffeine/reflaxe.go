# Snapshot Harness

## Run all snapshots

```bash
python3 test/run-snapshots.py
```

## List snapshots

```bash
python3 test/run-snapshots.py --list
```

## Run one snapshot

```bash
python3 test/run-snapshots.py --case core/hello_trace
```

## Update intended outputs

```bash
python3 test/run-snapshots.py --update
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

Run one module:

```bash
python3 test/run-upstream-stdlib-sweep.py --module haxe.Json
```

## Strict examples mode

Snapshots compile with `-D reflaxe_go_strict_examples` so app/test code cannot rely on raw `__go__` escape hatches.

## Troubleshooting

- Use `KEEP_ARTIFACTS=1` to keep generated `out/` folders on failures.
- Use `--failed` to rerun only the previous failing set.
- Use `--changed` to focus only on touched snapshot cases.
