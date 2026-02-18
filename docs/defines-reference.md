# Defines Reference (`-D ...`)

## Core output

- `go_output=<dir>`
  - Required signal for Go generation.
- `reflaxe.dont_output_metadata_id`
  - Recommended for deterministic snapshots.

## Profiles

- `reflaxe_go_profile=portable|gopher|metal`
  - Main profile selector.
- `reflaxe_go_portable`
  - Alias selector for portable.
- `reflaxe_go_gopher`
  - Alias selector for gopher.
- `reflaxe_go_metal`
  - Alias selector for metal.

Removed:

- `reflaxe_go_profile=idiomatic` -> compile error, use `gopher`.
- `reflaxe_go_idiomatic` -> compile error, use `reflaxe_go_profile=gopher`.

## Strictness

- `reflaxe_go_strict`
  - Enforce strict no raw `__go__` policy in app project sources.
- `reflaxe_go_strict_examples`
  - Enforce strict no raw `__go__` policy for repo examples/snapshots.

## Pass registry

- `go_granular_pass_registry`
  - Use granular pass bundle.
- `reflaxe_go_test_registry_case=<duplicate|missing_dep|cycle>`
  - Test-only define used by negative snapshot cases for registry validation.
