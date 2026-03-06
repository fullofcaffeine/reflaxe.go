# Agent Instructions

This project uses **bd** (beads) for issue tracking. Run `bd onboard` to get started.

## Quick Reference

```bash
bd ready              # Find available work
bd show <id>          # View issue details
bd update <id> --status in_progress  # Claim work
bd close <id>         # Complete work
bd sync               # Sync with git
```

## Thinking Levels (Bead Labels)

Use a `thinking:*` label on active beads so execution effort matches task risk.

- `thinking:low`
  - Mechanical edits, simple docs cleanup, straightforward renames, obvious wiring.
- `thinking:medium`
  - CI/job plumbing, runner scripts, artifact flow, bounded retry/timeout logic.
- `thinking:high`
  - Parity contracts, gate semantics, dependency graph changes, perf-policy changes, compiler/macro architecture decisions.
- `thinking:xhigh`
  - Scope-definition changes, release enforcement, provenance-sensitive implementation strategy, or any task where a wrong decision would create misleading release evidence.

Agent policy:

- When a bead has a `thinking:*` label, match reasoning depth to that label automatically.
- If a claimed bead has no `thinking:*` label, infer one immediately and add it before substantial work.
- `thinking:xhigh` should get a second-pass review before closure.
  - Preferred: an Oracle checkpoint/review.
  - Acceptable fallback: an explicit written second-pass design review recorded in the bead comments.
- Oracle is a review/escalation tool for `thinking:xhigh`; it is not a substitute for implementation, tests, or CI evidence.

## High-Level Goal

- Make `haxe.go` the best way to write Go without writing raw Go, while preserving first-class portability via Haxe for codebases that need cross-target builds.

## Landing the Plane (Session Completion)

**When ending a work session**, you MUST complete ALL steps below. Work is NOT complete until `git push` succeeds.

**MANDATORY WORKFLOW:**

1. **File issues for remaining work** - Create issues for anything that needs follow-up
2. **Run quality gates** (if code changed) - Tests, linters, builds
3. **Update issue status** - Close finished work, update in-progress items
4. **PUSH TO REMOTE** - This is MANDATORY:
   ```bash
   git pull --rebase
   bd sync
   git push
   git status  # MUST show "up to date with origin"
   ```
5. **Clean up** - Clear stashes, prune remote branches
6. **Verify** - All changes committed AND pushed
7. **Hand off** - Provide context for next session

**CRITICAL RULES:**
- Work is NOT complete until `git push` succeeds
- NEVER stop before pushing - that leaves work stranded locally
- NEVER say "ready to push when you are" - YOU must push
- If push fails, resolve and retry until it succeeds

## Compiler Guardrails

- Prefer the AST-first pipeline: builder/lowering -> transform passes -> printer/output.
- **Hard rule:** avoid `Dynamic`/`Any` whenever possible and prefer explicit typed abstractions end-to-end.
- **Hard rule:** use test-first development (TDD) for all code changes (compiler, runtime, std/shims, examples, docs-with-contracts).
- If `Dynamic`/`Any` is truly unavoidable, keep it localized behind runtime/shim boundaries and include a short justification in code/docs.
- If a change starts re-implementing Haxe stdlib library semantics as large `GoStmt.GoRaw` blocks inside `GoCompiler`, stop and reconsider ownership before proceeding.
- Prefer staged std overrides under `std/_std` for library-expressible behavior, then thin `hxrt` helpers when Go-side runtime support is actually needed; keep compiler shims as the last resort.
- Use externs only to model real Go-native APIs. Do not use externs to smuggle Haxe stdlib behavior into the target layer when staged std or `hxrt` is the correct ownership.
- If you hit one compiler-owned stdlib helper that looks misplaced, audit adjacent helpers in the same shim group and file follow-up beads instead of expanding the compiler one function at a time.
- Use sibling-target precedent before keeping stdlib behavior in `GoCompiler`: `haxe.rust` and `haxe.elixir` default library surfaces like `StringTools`/`DateTools` to staged std + small runtime helpers, and reserve compiler ownership for compile-context-sensitive behavior.
- Documentation threshold rule: do not reserve HaxeDoc only for obviously "big" constructs or artifacts. If a type/function/abstract/macro/extern override/metadata pattern is even slightly non-obvious, surprising, or easy to misuse, document it with `Why / What / How` HaxeDoc where it is declared.
- Bias toward documenting earlier rather than later, especially for abstracts, compiler helpers, runtime bindings, lowering hooks, and `std/` compatibility shims.
- For each staged std/runtime compatibility override you add, include `What / Why / How` HaxeDoc at the declaration and make the `Why` explicit about why the mainstream Haxe stdlib implementation could not be used unchanged on `haxe.go`.
- Never emit absolute machine-local paths in generated output or snapshots.
- When fixing a bug, always add or update a regression test in `test/snapshot`.

## Test-First Workflow (Mandatory)

- Start with the expected behavior as a failing contract before implementation.
- Prefer contract-first artifacts by change type:
  - snapshots for generated Go shape/output,
  - semantic-diff tests for portable semantics/behavior,
  - example `expected/*.stdout` and generated-output bless files for end-to-end UX.
- Apply red -> green -> refactor for each task slice.
- Every bug fix must include a regression test that fails before the fix.
- If behavior/output intentionally changes, update expected files in the same change and document why in commit/PR notes.
- Minimum validation before landing code:
  - targeted: `npm run test:changed` (and `npm run test:examples:changed` when examples changed),
  - full: `npm test`,
  - add `npm run test:semantic-diff` and `npm run test:stdlib-sweep:go-test` when runtime/semantics/profile code is touched,
  - run relevant perf harness (`npm run test:perf:go`, `npm run test:perf:hxrt-selective`) when optimization/runtime-slicing work changes.

## Snapshot Workflow

- Run all snapshots:
  ```bash
  npm test
  ```
- Update intended outputs intentionally:
  ```bash
  python3 test/run-snapshots.py --update
  ```
- Run upstream stdlib sweep:
  ```bash
  python3 test/run-upstream-stdlib-sweep.py --strict --go-test
  ```

## Injection Policy

- App/test/example code must not use raw `__go__` escapes.
- `__go__` usage is reserved for controlled target layers (e.g. std/runtime shims), not business logic.

## Session Lessons (Interop + Std)

- Higher-level rule: preserve **layered DX contracts** in canonical examples.
  - Show the default safe/productive layer (framework-owned wrappers).
  - Also show the explicit power layer (user-owned typed externs), without forcing users into it for common cases.
- Keep examples as contract docs: refactors that improve default DX must not erase the lower-level teaching path.
- For std wrapper modules, keep importable types in their own module files (one primary import target per file) to avoid hidden module/type resolution pitfalls.
- Documentation readability rule: prefer neutral, direct titles and scannable structure; avoid labels like "beginner-friendly" unless explicitly requested.
- When documenting harnesses/contracts, explain `what it is`, `why it exists`, and `how it works` in a short step-by-step flow first, then list commands.
- Anti-assumption rule: do not assume prior knowledge; define terms on first use (or link to their canonical doc) so a new reader can follow without hidden context.
