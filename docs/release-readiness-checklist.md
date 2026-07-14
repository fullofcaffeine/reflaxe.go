# Release Readiness Checklist

This checklist is the canonical production gate for `reflaxe.go`.
Run these checks from repo root on a clean branch before a release cut.

## Required GA gates

1. Supported toolchain contract:
   - Read the [supported toolchain policy](toolchain-policy.md).
   - Confirm the candidate jobs use supported Haxe, Go, and Node lines.
   - Record the exact resolved patch versions; a floating CI selector is not
     release provenance.
2. Full CI harness contract:
   - `python3 test/run-ci.py`
3. Snapshot + semantic baseline (already included in CI harness, can be run directly for debugging):
   - `python3 test/run-snapshots.py`
   - `python3 test/run-semantic-diff.py`
4. Portable parity closure visibility:
   - `python3 test/run-portable-stdlib-inventory.py`
   - `python3 test/run-portable-parity-closure.py`
5. Ownership and stdlib governance gates:
   - `npm run test:stdlib:governance`
   - `npm run test:release-contracts`
6. Family stdlib sync gates:
   - `npm run test:family-stdlib-sync`
   - `npm run test:family-stdlib-bootstrap`
7. Release visibility contract:
   - `npm run release:status`
8. Go dynamic and static tooling gates:
   - `npm run security:go-tooling`
   - Confirm the race detector, strict checkptr, vet, and pinned Staticcheck
     reports pass on both supported Go lines.
9. Locked supply-chain provenance:
   - `npm run security:supply-chain`
   - Confirm clean npm lock installation, immutable action pins, and vendored
     Reflaxe patch reconstruction all pass.
10. Performance visibility gates:
   - `npm run test:perf:go`
   - `npm run test:perf:hxrt-selective`
   - `npm run test:perf:apps`
11. Production caveat scoreboard review:
   - Read `docs/known-gaps.md#production-hardening-scoreboard`
   - Confirm each row still has an owner, current decision, evidence, and reopen trigger.
   - Confirm the scoreboard references durable evidence links and commands, not retired tracker IDs.
   - For multi-package output, check `docs/multi-package-output-evaluation.md#measurable-production-reopen-triggers` before filing implementation work.
12. Performance budget policy review:
   - Read `docs/performance-budget-policy.md`
   - Confirm warning-only perf drift is not being treated as release-blocking without the promotion criteria in that policy.

## Reproducible command set

Use this exact command order when validating a release candidate locally:

```bash
python3 test/run-ci.py
python3 test/run-portable-stdlib-inventory.py
python3 test/run-portable-parity-closure.py
npm run test:stdlib:governance
npm run test:release-contracts
npm run test:family-stdlib-sync
npm run test:family-stdlib-bootstrap
npm run release:status
npm run security:go-tooling
npm run security:supply-chain
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

- Candidate Haxe, Go, and Node versions match
  [`toolchain-policy.json`](toolchain-policy.json), and the evidence index
  records exact resolved patch versions.
- The generated `go 1.22` directive remains a language-floor contract; it
  does not make an old compatibility fixture production-supported.
- `python3 test/run-ci.py` exits `0` with no failed stages.
- `python3 test/run-portable-stdlib-inventory.py` exits `0` and every remaining `compile-only` module carries blocker issue + target metadata.
- `python3 test/run-portable-parity-closure.py` exits `0`, reports `0 actionable blockers`, and keeps any remaining non-semantic-diff surfaces policy-locked as target-sensitive snapshots or explicit exclusions.
- `npm run test:stdlib:governance` exits `0` and confirms provenance/boundary discipline for staged std ownership.
- `npm run test:release-contracts` exits `0` and confirms ownership mapping plus release docs still match the live inventory/tracker state.
- `npm run test:family-stdlib-sync` and `npm run test:family-stdlib-bootstrap` exit `0`.
- `npm run release:status` exits `0` and reports release wiring as healthy.
- `npm run security:go-tooling` exits `0`; race detector, strict checkptr,
  vet, and pinned Staticcheck reports contain no blocking findings on every
  supported Go line.
- `npm run security:supply-chain` exits `0`; the npm lock matches
  declared dependencies, every external action matches the immutable action
  manifest, and vendored Reflaxe passes its offline patch round-trip.
- Perf runs complete and budgets are within expected thresholds for the current baseline policy.
- Performance budget policy in `docs/performance-budget-policy.md` still matches CI enforcement settings and any warning-only drift has an explicit follow-up decision.
- The Production caveat scoreboard in `docs/known-gaps.md#production-hardening-scoreboard` still matches current evidence and does not hide target-sensitive or warning-only caveats as full semantic guarantees.
- Multi-package output remains deferred unless a trigger in `docs/multi-package-output-evaluation.md#measurable-production-reopen-triggers` has repeatable evidence attached.

## Related references

- [Supply-chain policy](supply-chain-policy.md)
- [Vendored Reflaxe provenance](vendor-reflaxe-provenance.md)

- CI stage contract source: `test/run-ci.py`
- Supported toolchain policy: `docs/toolchain-policy.md`
- Release automation visibility checks: `docs/release-visibility.md`
- Snapshot policy: `docs/snapshot-policy.md`
- Semantic differential guide: `docs/semantic-diff-guide.md`
- Performance budget policy: `docs/performance-budget-policy.md`
- Multi-package output reopen triggers: `docs/multi-package-output-evaluation.md#measurable-production-reopen-triggers`
