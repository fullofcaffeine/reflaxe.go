# Agent Instructions

This project uses **bd** (Beads) for issue tracking. Run `bd prime` to load the current workflow and persistent project memories.

## Quick Reference

```bash
bd dolt pull          # Pull issue history from refs/dolt/data
bd ready              # Find available work
bd show <id>          # View issue details
bd update <id> --claim  # Claim work
bd close <id>         # Complete work
bd dolt push          # Push issue history to refs/dolt/data
```

Source Git history and Beads Dolt history are separate synchronization channels. See `.beads/README.md` before changing tracker storage, export, or recovery configuration.

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
- If a task starts to exceed its current thinking level, stop and say so explicitly before widening scope.
  - Ask the user for the higher thinking level before continuing at that level.
  - If the current thinking level remains correct, continue without interruption.
  - Only stop to ask when the reasoning level truly needs to change.
  - Raise to extended thinking when local tracing shows the fix is no longer a bounded implementation and has become a broader design problem.
  - Escalate to Oracle only when local investigation still leaves multiple defensible designs unresolved after that deeper pass.
- `thinking:xhigh` should get a second-pass review before closure.
  - Preferred: an Oracle checkpoint/review.
  - Acceptable fallback: an explicit written second-pass design review recorded in the bead comments.
- Oracle is a review/escalation tool for `thinking:xhigh`; it is not a substitute for implementation, tests, or CI evidence.

## High-Level Goal

- Make `haxe.go` the best way to write Go without writing raw Go, while preserving first-class portability via Haxe for codebases that need cross-target builds.
- Product rule: portable Haxe semantics are the default product path; typed `go.*`/extern APIs and `@:goNative` modules are explicit Go-native source boundaries. The public `portable|metal` selector remains for compatibility, with `metal` defined as a convenience policy preset rather than a second semantic product. Portable by default, Go-native by explicit source boundary, Go-shaped generated output whenever the compiler can prove the lowering preserves the source contract.
- Do not frame `portable` as the slow/basic mode or `metal` as the only "real Go" mode. Compiler optimization work should make portable output idiomatic and performant where semantics allow. New behavior must branch on typed authority/specialization/fallback/strictness/runtime policies, not directly on the legacy profile name; `@:goMetal` remains a compatibility alias for canonical `@:goNative`.

## Landing the Plane (Session Completion)

**When ending a work session**, you MUST complete ALL steps below. Work is NOT complete until `git push` succeeds.

**MANDATORY WORKFLOW:**

1. **File issues for remaining work** - Create issues for anything that needs follow-up
2. **Run quality gates** (if code changed) - Tests, linters, builds
3. **Update issue status** - Close finished work, update in-progress items
4. **PUSH TO REMOTE** - This is MANDATORY:
   ```bash
   bd dolt pull
   git pull --rebase
   bd dolt push
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
- Prefer staged std overrides as ordinary source under `std/go/_std` for library-expressible behavior, then thin `hxrt` helpers when Go-side runtime support is actually needed; keep compiler shims as the last resort. Package staging—not source control—owns the `.cross.hx` conversion.
- Use externs only to model real Go-native APIs. Do not use externs to smuggle Haxe stdlib behavior into the target layer when staged std or `hxrt` is the correct ownership.
- Framework-owned `__go__` is a valid middle layer now: prefer typed externs first, then narrow `reflaxe.go.macros.GoInjection.__go__` / `@:goAllowRaw` abstraction islands in `std/` or runtime helpers, before growing compiler-owned `GoRaw` emitters.
- Raw `__go__` still does not carry package imports by itself. When a snippet needs external Go packages, keep imports typed and explicit through extern metadata (`@:go.import`, `@:go.name`, `@:go.receiver`) or existing framework wrappers.
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
  - run `npm run test:examples` for compiler, runtime, staged std, profile/strictness, or example changes that can affect generated app behavior,
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
- Examples must not use raw `__go__`, `GoInjection`, or `@:goAllowRaw`; teach staged std, `hxrt`, `go.*` facades, or typed externs instead.
- Intentional boundary fixtures that exercise raw-injection enforcement must live in explicit allowlisted test locations, not ordinary examples or app-style fixtures.
- Run `python3 test/test_raw_injection_hygiene_contract.py` when changing example/test interop surfaces or raw-injection policy.

## Session Lessons (Interop + Std)

- Examples are QA contracts.
  - Every shipped example must be visible to `python3 test/run-examples.py`.
  - Every runnable example/profile lane must compile, pass `go test ./...`, run through `go run .`, and match `expected/*.stdout` unless there is an explicit documented reason it is compile-only.
  - If compiler/runtime/std/profile changes can affect examples, run `npm run test:examples`, not only `npm run test:examples:changed`.
  - When an example behavior changes intentionally, update `expected/*.stdout` and committed `generated/<profile>` trees in the same change.
  - Do not add example directories that are invisible to the examples harness.
- Higher-level rule: preserve **layered DX contracts** in canonical examples.
  - Show the default safe/productive layer (framework-owned wrappers).
  - Also show the explicit power layer (user-owned typed externs), without forcing users into it for common cases.
- Keep examples as contract docs: refactors that improve default DX must not erase the lower-level teaching path.
- For std wrapper modules, keep importable types in their own module files (one primary import target per file) to avoid hidden module/type resolution pitfalls.
- Documentation readability rule: prefer neutral, direct titles and scannable structure; avoid labels like "beginner-friendly" unless explicitly requested.
- When documenting harnesses/contracts, explain `what it is`, `why it exists`, and `how it works` in a short step-by-step flow first, then list commands.
- Anti-assumption rule: do not assume prior knowledge; define terms on first use (or link to their canonical doc) so a new reader can follow without hidden context.

<!-- BEGIN BEADS INTEGRATION v:1 profile:minimal hash:7510c1e2 -->
## Beads Issue Tracker

This project uses **bd (beads)** for issue tracking. Run `bd prime` to see full workflow context and commands.

### Quick Reference

```bash
bd dolt pull          # Pull issue history from refs/dolt/data
bd ready              # Find available work
bd show <id>          # View issue details
bd update <id> --claim  # Claim work
bd close <id>         # Complete work
bd dolt push          # Push issue history to refs/dolt/data
```

### Rules

- Use `bd` for ALL task tracking — do NOT use TodoWrite, TaskCreate, or markdown TODO lists
- Run `bd prime` for detailed command reference and session close protocol
- Use `bd remember` for persistent knowledge — do NOT use MEMORY.md files

**Architecture in one line:** active issues live in a local Dolt DB and synchronize through `refs/dolt/data`; this repository's tracked `.beads/issues.jsonl` is the canonical legacy provenance archive and must not be overwritten by auto-export. See `.beads/README.md` and https://github.com/gastownhall/beads/blob/main/docs/SYNC_CONCEPTS.md.

## Session Completion

**When ending a work session**, you MUST complete ALL steps below. Work is NOT complete until `git push` succeeds.

**MANDATORY WORKFLOW:**

1. **File issues for remaining work** - Create issues for anything that needs follow-up
2. **Run quality gates** (if code changed) - Tests, linters, builds
3. **Update issue status** - Close finished work, update in-progress items
4. **PUSH TO REMOTE** - This is MANDATORY:
   ```bash
   bd dolt pull
   git pull --rebase
   bd dolt push
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
<!-- END BEADS INTEGRATION -->
