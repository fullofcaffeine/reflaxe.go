# Start Here

## What this target does

`reflaxe.go` compiles Haxe to a generated Go module (`go_output`), copies runtime support under `out/hxrt`, and runs `go build` by default (unless disabled).

## First successful run

1. Install toolchain deps:

```bash
npm install
```

Install repo pre-commit hooks (recommended):

```bash
npm run hooks:install
```

2. Run snapshots:

```bash
python3 test/run-snapshots.py
```

3. Run CI entrypoint:

```bash
python3 test/run-ci.py
```

4. Run examples matrix:

```bash
python3 test/run-examples.py
```

## Scaffold a new project

```bash
npm run dev:new-project -- ./my_haxe_go_app
cd ./my_haxe_go_app
npm install
npm run setup
npm run hx:run
```

## Task and package manager model

- Generated project dependencies/build/test use the Go toolchain directly (`go mod`, `go run`, `go test`, `go build`).
- Compiler/dev orchestration uses `npm` scripts so workflow stays consistent with `haxe.rust` and `haxe.elixir.codex`.
- Direct hxml compiles run backend `go build` by default; use `-D go_no_build` for codegen-only flows.

Quick compile+go action from this repo:

```bash
npm run dev:hx -- --project examples/tui_todo --profile portable --action run
```

The wrapper resolves `compile*.hxml`, runs Haxe, resolves `-D go_output=...`, then executes the selected Go action.

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
- `docs/feature-support-matrix.md`
- `docs/examples-matrix.md`
- `docs/defines-reference.md`
- `docs/profile-admission-criteria.md`
- `docs/compiler-target-template.md`
- `docs/snapshot-policy.md`
- `SECURITY.md`
