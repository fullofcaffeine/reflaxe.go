# Compiler Debt Baseline and Ratchet

## What it is

The compiler debt ratchet is a deterministic inventory and no-growth gate for
six target-owned implementation signals:

- `GoStmt.GoRaw` and `GoExpr.GoRaw` construction;
- Haxe `Dynamic` and `Any` types;
- Go `reflect` and `unsafe` package imports and selectors; and
- named compiler shim entry points.

The checked-in authority is `test/compiler_debt_policy.json`. The executable
gate is `test/run-compiler-debt-ratchet.py`.

Counts are directional evidence, not correctness scores. A required reflective
runtime boundary can be correct and well-tested while an avoidable raw compiler
string remains valuable migration evidence. Neither count proves semantic
parity; snapshots, semantic-diff tests, runtime tests, and Go tooling gates keep
that responsibility.

## Why it exists

The repository deliberately supports dynamic Haxe semantics and explicit
Go-native APIs, so a global ban on `Dynamic`, `Any`, or reflection would be
false precision. At the same time, unexplained growth makes typed ownership and
AST-first lowering harder to achieve.

The policy therefore classifies current findings in two ways:

- `required` means the current boundary is justified by a language, upstream
  API, runtime, or compile-context contract. It remains counted and cannot grow
  silently.
- `avoidable` means the behavior may be valid but its present representation is
  debt. All raw Go AST strings are in this class because typed AST nodes expose
  more structure to transforms, import analysis, and validation.

Every exception records an owner and explicit What / Why / How rationale. A
baseline entry references one of those exceptions and sets a ceiling for its
exact file, function context, ownership class, capability, profile, and surface.

The target type/operator migration and its validated legacy-string boundary are
documented in [`typed-go-ir.md`](typed-go-ir.md).

## How it works

Run:

```bash
npm run test:compiler-debt
```

The scanner ignores comments and quoted strings for source-token matching,
resolves ordinary and aliased Go package imports, then reports only repository-
relative paths. It covers:

1. compiler and target API Haxe source under `src/reflaxe/go` and `src/go`;
2. staged standard-library and runtime-facing Haxe source under `std`;
3. checked-in Go runtime code under `runtime/hxrt`; and
4. committed example output under `examples/*/generated/portable` and
   `examples/*/generated/metal`.

The last two generated lanes are measured separately. `portable` is still the
default semantic product and `metal` is still a compatibility convenience
preset for explicit Go-native policy. Separate report rows make generated
output changes visible; they do not establish a second semantic product or
grant profile-wide native authority. Native authority remains module/API
scoped.

The command writes deterministic current-state reports to:

- `.cache/compiler-debt/report.json`
- `.cache/compiler-debt/report.md`

Both reports aggregate counts by file, owner, capability, profile, surface, and
required/avoidable classification. Runtime code and copied generated-runtime
code have distinct surfaces, so selective `hxrt` changes remain visible.

## Ratchet semantics

For each current report row, CI requires a matching policy entry.

- A lower count passes. Regenerate the baseline in the same migration so the
  reduction becomes the new ceiling.
- A higher count fails with the current and permitted counts.
- A finding in a new file, function context, owner, capability, profile, or
  surface fails as unexplained debt.
- Direct `unsafe` package touchpoints currently have a zero baseline. The first use fails
  even though no zero-count file row exists, because every observed finding must
  match a checked-in entry.

This gate runs from `npm test`, `npm run test:changed`, and the release-contract
suite used by `test/run-ci.py`.

## Updating a reviewed boundary

Prefer removing or typing the finding. When a justified boundary really changes:

1. update the relevant exception's owner and What / Why / How rationale;
2. add or update behavior, snapshot, runtime, and tooling evidence appropriate
   to the affected surface;
3. regenerate current ceilings:

   ```bash
   python3 test/run-compiler-debt-ratchet.py --update-baseline
   ```

4. inspect `git diff -- test/compiler_debt_policy.json` and confirm no unrelated
   file, profile, or ownership bucket grew; and
5. run `npm run test:compiler-debt` again without the update flag.

`--update-baseline` is a mechanical writer, not an approval mechanism. A larger
number is accepted only when its rationale and evidence survive ordinary code
review. Oracle review is reserved for unresolved `thinking:xhigh` architecture
choices; maintaining this `thinking:high` measurement contract does not by
itself require Oracle.
