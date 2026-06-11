# Profile Semantics Guide (`portable` vs `metal`)

This guide explains the practical and semantic differences between profiles.

Short version: portable is the default product path; metal is an explicit Go-native authoring contract.

Portable by default, Go-native by opt-in, metal-like generated Go whenever the compiler can prove the lowering preserves portable Haxe semantics.

Use `portable` as the normal way to write Haxe that becomes readable,
idiomatic Go. Use `metal` only when you intentionally want Go-native authoring
surfaces, stricter boundaries, or fail-fast native-lane checks.

## Terms

- [portable](/docs/glossary.md#portable-profile): portability-first profile contract.
- [metal](/docs/glossary.md#metal-profile): Go-first profile contract.
- [fallback](/docs/glossary.md#fallback): safe path used when strict typed lowering cannot apply.
- [semantic diff](/docs/glossary.md#semantic-diff): runtime behavior parity test against Haxe `--interp`.

## Architecture model (current)

The profile system is explicit and layered:

1. semantic contract: `portable|metal`
2. boundary policy: strict mode + portable native-import policy
3. runtime policy: full vs selective hxrt copy
4. planner policy: `off|auto|auto_strict`
5. lane scope: `@:goMetal` modules

`GoBuildContextResolver.resolve()` computes these axes once, and `GoReflaxeCompiler` uses the resolved context to drive compile behavior and report emission.

## Quick answer

- Choose `portable` when you want shared, cross-target-friendly Haxe behavior.
- Choose `metal` when you intentionally want explicit Go-native APIs and stricter native-lane checks.
- Do not choose `metal` just because you want good Go output; the portable optimizer should generate Go-shaped fast paths whenever it can do so without changing Haxe semantics.

## What usually stays the same

If code remains on portable surfaces (Haxe stdlib/app-level APIs, no target-native shortcuts), runtime behavior should stay equivalent across profiles.

Examples:

- `Std.string(null)` should still behave as portable `"null"` semantics.
- portable null/equality paths should not silently flip just because profile changed.

When this does not hold, treat it as a regression or an explicitly documented exception.

## What usually changes

1. **Boundary strictness**
   - `metal` defaults to stricter app-side injection policy.
2. **Typed specialization pressure**
   - `metal` is where typed native lowering is expected first.
3. **Portability posture**
   - `portable` optimizes for cross-target stability.
   - `metal` accepts lower portability in native-first paths.

## Real-world examples

### Example A: portable code in metal build

Portable-style code can compile under `metal` and keep behavior parity.
This is useful for “metal-readiness” checks before fully adopting native-first lanes.

### Example B: native-first code in portable build

Native facades (`go.*`) may still compile in portable depending on policy, but they are outside the portable compatibility contract by design.
Use:

- `reflaxe_go_portable_native_policy=error` to fail fast in CI
- `reflaxe_go_portable_native_policy=warn` for local migration visibility

### Example C: typed lowering fallback

When typed specialization is not possible, the compiler can use a fallback path.

- in `metal`: fallback is a hard error by default unless explicitly allowed.
- in `portable`: fallback is allowed by default, with report visibility.

## `@:goMetal` lanes inside portable

`@:goMetal` lets you harden selected modules while keeping the whole build portable.

Under `-D reflaxe_go_auto=auto_strict`, lane modules fail if typed specialization for go-native collections/concurrency/result paths cannot be applied.

This supports incremental migration:

1. keep shared code portable
2. mark hot/native modules as lane modules
3. enforce stricter rules only where needed

## Why explicit profiles (not inferred semantics)

The repo keeps explicit profile selection because semantic intent must stay visible in CI and code review.

If semantics were inferred from usage:

- small dependency changes could silently alter behavior,
- teams would lose reviewable contract intent,
- debugging profile-related behavior drift would become harder.

Inference is still used for additive planning (runtime feature selection, optimizer plans), not for hidden semantic profile switching.

## What profiles do not do

- They do not infer semantic contract from imports or optimizer outcomes.
- They do not collapse runtime slicing into contract selection.
- They do not replace lane metadata (`@:goMetal`) with implicit module inference.

## Portable to metal migration checklist

1. Keep semantic-diff green in portable.
2. Enable metal in one app lane.
3. Resolve strict-boundary violations using typed wrappers/facades.
4. Benchmark before/after.
5. Keep portable shared modules untouched unless evidence justifies native-only tradeoffs.

## Portable to metal admission criteria (pass/fail)

### Pass criteria

1. Lane and boundary checks are green:
   - `python3 test/run-semantic-diff.py --suite lanes`
   - `python3 test/run-snapshots.py`
2. Strict policy is enforced as intended for the promoted lane.
3. Contract reports show deterministic fallback diagnostics.
4. Perf evidence exists for promoted modules.

### Fail criteria

1. Lane semantic-diff or boundary checks fail.
2. Promotion depends on raw app-side `__go__` injection.
3. Fallback diagnostics are missing or unstable.
4. Profile switch is made without benchmark evidence or explicit tradeoff documentation.

## Metal back to portable checklist

1. Remove target-native-only usage in shared modules.
2. Re-run snapshots + semantic diff in portable.
3. Re-check examples/support matrix expectations.
4. Document any remaining native-only islands.

## Command set for profile-sensitive changes

```bash
python3 test/run-snapshots.py
python3 test/run-semantic-diff.py
python3 test/run-semantic-diff.py --suite lanes
python3 test/run-ci.py
```

## Related docs

- Docs map: [docs/index.md](index.md)
- Glossary: [docs/glossary.md](glossary.md)
- Profiles reference: [docs/profiles.md](profiles.md)
- Portable contract: [docs/portable-canonical-contract.md](portable-canonical-contract.md)
- Versioned semantics spec: [docs/portable-semantics-v1.md](portable-semantics-v1.md)
- Semantic diff guide: [docs/semantic-diff-guide.md](semantic-diff-guide.md)
- Examples matrix: [docs/examples-matrix.md](examples-matrix.md)
