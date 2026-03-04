# Release Readiness Checklist

This checklist is the canonical production gate for `reflaxe.go`.
Run these checks from repo root on a clean branch before a release cut.

## Required GA gates

1. Full CI harness contract:
   - `python3 test/run-ci.py`
2. Snapshot + semantic baseline (already included in CI harness, can be run directly for debugging):
   - `python3 test/run-snapshots.py`
   - `python3 test/run-semantic-diff.py`
3. Release visibility contract:
   - `npm run release:status`
4. Performance visibility gates:
   - `npm run test:perf:go`
   - `npm run test:perf:hxrt-selective`
   - `npm run test:perf:apps`

## Reproducible command set

Use this exact command order when validating a release candidate locally:

```bash
python3 test/run-ci.py
npm run release:status
npm run test:perf:go
npm run test:perf:hxrt-selective
npm run test:perf:apps
```

Optional strict local parity flags:

```bash
GO_PERF_ENFORCE_METAL_BUDGET=1 npm run test:perf:go
GO_HXRT_SLICE_ENFORCE=1 npm run test:perf:hxrt-selective
GO_APP_PERF_ENFORCE_METAL_BUDGET=1 npm run test:perf:apps
```

## Pass criteria

- `python3 test/run-ci.py` exits `0` with no failed stages.
- `npm run release:status` exits `0` and reports release wiring as healthy.
- Perf runs complete and budgets are within expected thresholds for the current baseline policy.

## Related references

- CI stage contract source: `test/run-ci.py`
- Release automation visibility checks: `docs/release-visibility.md`
- Snapshot policy: `docs/snapshot-policy.md`
- Semantic differential guide: `docs/semantic-diff-guide.md`
