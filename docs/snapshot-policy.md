# Snapshot Policy

## Canonical output rule

Snapshot goldens track the final **post-Reflaxe optimized AST** Go output.

This means snapshots intentionally preserve simplifications such as:

- constant folding (`7 + 5 * 2` -> `17`)
- boolean simplification (`(value == 3) && true` -> `value == 3`)
- normalized temporary variable naming

## Why this policy exists

- Keeps snapshot diffs aligned with what users actually compile and ship.
- Avoids churn from internal pre-optimization tree differences.
- Separates concerns: behavior is validated by build/runtime checks, shape by snapshots.

## Sentinel test

`test/snapshot/core/optimized_ast_policy` is the explicit policy sentinel.

## Snapshot Realignment Notes

### 2026-02-24 broad realignment (`haxe.go-cme`)

Snapshot fixtures were intentionally refreshed after coordinated runtime/profile/codegen updates:

- `hxrt` split output (`hxrt/*.go`) replaced monolithic fixture shapes.
- generic-call result assertions were expanded to preserve typed semantics after monomorphization/generic erasure boundaries.
- primitive-vs-`null` comparisons are now normalized at lowering time to valid constant boolean outcomes.

Policy reminder:

- this type of broad fixture update is expected when generated Go shape changes intentionally,
- semantic parity is validated in parallel via `python3 test/run-semantic-diff.py`,
- example UX parity is validated via `python3 test/run-examples.py`.
