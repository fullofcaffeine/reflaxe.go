# Start Here

## What this target does

`reflaxe.go` compiles Haxe to a generated Go module (`go_output`) and copies runtime support under `out/hxrt`.

## First successful run

1. Install toolchain deps:

```bash
npm install
```

2. Run snapshots:

```bash
python3 test/run-snapshots.py
```

3. Run CI entrypoint:

```bash
python3 test/run-ci.py
```

## Profile selection

Set via:

```bash
-D reflaxe_go_profile=portable|gopher|metal
```

- `portable` (default): choose this for portability and lowest migration risk.
- `gopher`: choose this for Go-first APIs/output style while keeping core semantics stable.
- `metal` (experimental): choose this when gopher abstractions are not enough and you need typed low-level interop with strict boundaries.

Compatibility note:

- `idiomatic` is removed and intentionally fails fast; use `gopher` instead.

## Strict policy knobs

- `-D reflaxe_go_strict_examples`: forbids raw `__go__` in repo examples/snapshots.
- `-D reflaxe_go_strict`: forbids raw `__go__` in app project sources.
- `metal` enables strict mode by default for app-side injection boundaries.

## Related docs

- `docs/profiles.md`
- `docs/defines-reference.md`
- `docs/profile-admission-criteria.md`
- `docs/compiler-target-template.md`
