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

## Strict examples mode

Snapshots compile with `-D reflaxe_go_strict_examples` so app/test code cannot rely on raw `__go__` escape hatches.

## Troubleshooting

- Use `KEEP_ARTIFACTS=1` to keep generated `out/` folders on failures.
- Use `--failed` to rerun only the previous failing set.
