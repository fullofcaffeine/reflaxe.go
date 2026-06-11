# Start Here

This is the shortest path from clone -> running app -> understanding profile choices.

## Terms

- [profile](/docs/glossary.md#profile): build contract (`portable` or `metal`).
- [`hxrt`](/docs/glossary.md#hxrt): runtime package copied into generated output.
- [semantic diff](/docs/glossary.md#semantic-diff): behavior parity test versus Haxe `--interp`.

## What this target does

`reflaxe.go` compiles Haxe into a Go module (`go_output`) with:

- `main.go`
- per-module `module_*.go`
- copied runtime files under `hxrt/`

By default it also runs `go build` after code generation.

## First successful run

Use this path when you want to prove the local checkout works and see generated
Go run a real example:

```bash
npm install
npm run hooks:install
npm run dev:hx -- --project examples/tui_todo --profile portable --action run
```

That command runs the portable TUI todo example. `portable` is the default
starting point because it keeps the code closest to normal Haxe semantics and
cross-target-friendly APIs.

## Generated output and local artifacts

The first run creates generated Go under:

```text
examples/tui_todo/out_portable/
```

That path comes from `examples/tui_todo/compile.portable.hxml`:

```text
-D go_output=out_portable
```

The generated directory contains `go.mod`, generated `.go` files, and the
copied `hxrt/` runtime package. It is safe to delete because it is regenerated
the next time you run the compile/run wrapper.

If you want generated output somewhere else, pass `--out`:

```bash
npm run dev:hx -- --project examples/tui_todo --profile portable --out /tmp/tui_todo_go --action run
```

## Validation after first run

After the first example works, use these checks depending on what you changed:

```bash
# generated-output contracts
python3 test/run-snapshots.py

# example apps
python3 test/run-examples.py

# full local CI harness
python3 test/run-ci.py
```

For normal development, start with the narrowest command that covers your
change. Use the full CI harness before release cuts or broad compiler/runtime
changes.

## Choose your first example

- Portable baseline: `examples/tui_todo`
- Profile-neutral interop: `examples/interop_smoke`
- Clear portable vs metal comparison: `examples/worker_pool_select`
- Full app benchmark lanes: `examples/pulseforge`, `examples/fluxproxy`

Run all examples:

```bash
python3 test/run-examples.py
```

## Profile choice (practical)

Use:

```bash
-D reflaxe_go_profile=portable|metal
```

- `portable`: best default for shared/cross-target-friendly code.
- `metal`: opt-in Go-native authoring contract with stricter defaults and
  fail-fast native-lane checks.

`metal` is not required for good Go output. Start in `portable`; the compiler
should still emit Go-shaped fast paths whenever it can prove they preserve Haxe
semantics.

Detailed profile policy: [/docs/profiles.md](/docs/profiles.md)

## Strict policy knobs

- `-D reflaxe_go_strict_examples`: forbids raw `__go__` in repo examples/snapshots.
- `-D reflaxe_go_strict`: forbids raw `__go__` in app code.
- `-D reflaxe_go_strict_policy=auto|on|off`: app strictness policy (`auto` default: strict in `metal`, relaxed in `portable`).

## Where is the stdlib?

Short answer: most of it is **not** copied into this repo under `std/`.

Beginner mental model:

1. You still start from the upstream Haxe stdlib.
2. `reflaxe.go` adds only the target-specific parts it needs:
   - staged overrides in `std/_std`
   - Go facade modules in `std/go`
   - runtime helper files in `runtime/hxrt`
   - compiler-emitted shims in `src/reflaxe/go/GoCompiler.hx`

Why this is done:

- it avoids maintaining a full stdlib fork in this repository,
- it lets parity work land incrementally with tests,
- and it keeps target-specific behavior isolated.

Deep dive and ownership rule: [/docs/ownership-rubric.md](/docs/ownership-rubric.md).
Current shim-by-shim decisions: [/docs/stdlib-shim-rationale.md](/docs/stdlib-shim-rationale.md).

## Performance and release checks

```bash
npm run test:perf:go
npm run test:perf:hxrt-selective
npm run test:perf:apps
npm run release:status
```

Release checklist: [/docs/release-readiness-checklist.md](/docs/release-readiness-checklist.md)
Performance budget policy: [/docs/performance-budget-policy.md](/docs/performance-budget-policy.md)

## Related docs

- Docs map: [docs/index.md](index.md)
- Glossary: [docs/glossary.md](glossary.md)
- Profiles: [docs/profiles.md](profiles.md)
- Profile semantics guide: [docs/profile-semantics-guide.md](profile-semantics-guide.md)
- Examples matrix: [docs/examples-matrix.md](examples-matrix.md)
- Semantic diff guide: [docs/semantic-diff-guide.md](semantic-diff-guide.md)
- `hxrt` runtime: [docs/hxrt-runtime.md](hxrt-runtime.md)
- Release readiness checklist: [docs/release-readiness-checklist.md](release-readiness-checklist.md)
