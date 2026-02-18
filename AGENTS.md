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
- Keep `Dynamic` usage minimal and localized. If unavoidable, contain it behind runtime/shim boundaries.
- Never emit absolute machine-local paths in generated output or snapshots.
- When fixing a bug, always add or update a regression test in `test/snapshot`.

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
