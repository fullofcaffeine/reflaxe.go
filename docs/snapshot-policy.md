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
