# reflaxe.go

Haxe 4.3.7 -> Go target.

This backend prioritizes portable Haxe semantics first, with an opt-in Go-first profile surface.

## Start Here

- Onboarding: `docs/start-here.md`
- Profiles: `docs/profiles.md`
- Defines reference: `docs/defines-reference.md`
- Future target template: `docs/compiler-target-template.md`

## Profiles

Use `-D reflaxe_go_profile=portable|gopher|metal`.

- `portable` (default): Haxe-first portability and predictable semantics.
- `gopher`: Go-first authoring/output style without changing core semantic guarantees (includes safe literal string-op folding, typed String helper lowering, and safe leaf dispatch devirtualization for `self`/constructor-local calls).
- `metal` (experimental): gopher+typed low-level interop lane and strict app boundary defaults.

`idiomatic` has been removed and now fails fast.

## Useful Commands

- Run snapshots: `python3 test/run-snapshots.py`
- Run CI surface: `python3 test/run-ci.py`
- Re-run previous failures: `python3 test/run-snapshots.py --failed`
- Bless changed snapshot files only: `python3 test/run-snapshots.py --bless`

## Strictness

- `-D reflaxe_go_strict_examples`: repo examples/snapshots may not use raw `__go__`.
- `-D reflaxe_go_strict`: user project strict app-boundary policy.
- `metal` enables strict mode defaults for app-side injections.
