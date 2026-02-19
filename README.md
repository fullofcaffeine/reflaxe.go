# reflaxe.go

Haxe 4.3.7 -> Go target.

This backend prioritizes portable Haxe semantics first, with an opt-in Go-first profile surface.

## Start Here

- Onboarding: `docs/start-here.md`
- CI quality gate: `.github/workflows/ci-quality.yml`
- CI static security gate: `.github/workflows/security-static-analysis.yml`
- Feature support matrix: `docs/feature-support-matrix.md`
- Profiles: `docs/profiles.md`
- Defines reference: `docs/defines-reference.md`
- Examples matrix: `docs/examples-matrix.md`
- Future target template: `docs/compiler-target-template.md`
- Security policy: `SECURITY.md`

## Profiles

Use `-D reflaxe_go_profile=portable|gopher|metal`.

- `portable` (default): Haxe-first portability and predictable semantics.
- `gopher`: Go-first authoring/output style without changing core semantic guarantees (includes safe literal string-op folding, typed String helper lowering, and safe leaf dispatch devirtualization for `self`, inline constructor targets, tracked local aliases, and known leaf-returning call targets).
- `metal` (experimental): gopher+typed low-level interop lane and strict app boundary defaults.

`idiomatic` has been removed and now fails fast.

## Useful Commands

- Scaffold a starter project: `npm run dev:new-project -- ./my_haxe_go_app`
- Run snapshots: `python3 test/run-snapshots.py`
- Run CI surface: `python3 test/run-ci.py`
- Run semantic differential checks (interp vs Go portable): `python3 test/run-semantic-diff.py`
- Run examples profile matrix: `python3 test/run-examples.py`
- Re-run previous failures: `python3 test/run-snapshots.py --failed`
- Bless changed snapshot files only: `python3 test/run-snapshots.py --bless`
- Bless generated example Go trees: `python3 test/run-examples.py --bless-generated`
- Install repo pre-commit hook: `npm run hooks:install`
- Run repository gitleaks scan: `npm run security:gitleaks`
- Run dependency vulnerability audit: `npm run security:deps`

## Strictness

- `-D reflaxe_go_strict_examples`: repo examples/snapshots may not use raw `__go__`.
- `-D reflaxe_go_strict`: user project strict app-boundary policy.
- `metal` enables strict mode defaults for app-side injections.

## Local Security Gate

The repo-managed pre-commit hook (`npm run hooks:install`) enforces:

- staged local path leak guard (`scripts/lint/local_path_guard_staged.sh`)
- staged secret scan with `gitleaks` (`scripts/security/run-gitleaks.sh --staged`)
- staged Haxe auto-format (`haxelib run formatter`)

## Examples

- `examples/profile_storyboard`: compact profile-adapter reference.
- `examples/tui_todo`: complex app reference (single Haxe codebase -> `portable|gopher|metal`).

See `docs/examples-matrix.md` for commands, output layout, and artifact matrix.
