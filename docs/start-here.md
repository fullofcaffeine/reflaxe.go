# Start Here

This is the shortest path from clone -> running app -> understanding profile choices.

## Terms

- [profile](docs/glossary.md#profile): build contract (`portable` or `metal`).
- [`hxrt`](docs/glossary.md#hxrt): runtime package copied into generated output.
- [semantic diff](docs/glossary.md#semantic-diff): behavior parity test versus Haxe `--interp`.

## What this target does

`reflaxe.go` compiles Haxe into a Go module (`go_output`) with:

- `main.go`
- per-module `module_*.go`
- copied runtime files under `hxrt/`

By default it also runs `go build` after code generation.

## First successful run

```bash
npm install
npm run hooks:install
python3 test/run-snapshots.py
python3 test/run-ci.py
npm run dev:hx -- --project examples/tui_todo --profile portable --action run
```

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
- `metal`: opt-in Go-first lanes with stricter defaults and stronger typed specialization pressure.

Detailed profile policy: [docs/profiles.md](docs/profiles.md)

## Strict policy knobs

- `-D reflaxe_go_strict_examples`: forbids raw `__go__` in repo examples/snapshots.
- `-D reflaxe_go_strict`: forbids raw `__go__` in app code.
- `-D reflaxe_go_strict_policy=auto|on|off`: app strictness policy (`auto` default: strict in `metal`, relaxed in `portable`).

## Performance and release checks

```bash
npm run test:perf:go
npm run test:perf:hxrt-selective
npm run test:perf:apps
npm run release:status
```

Release checklist: [docs/release-readiness-checklist.md](docs/release-readiness-checklist.md)

## Related docs

- Docs map: [docs/index.md](index.md)
- Glossary: [docs/glossary.md](glossary.md)
- Profiles: [docs/profiles.md](profiles.md)
- Profile semantics guide: [docs/profile-semantics-guide.md](profile-semantics-guide.md)
- Examples matrix: [docs/examples-matrix.md](examples-matrix.md)
- Semantic diff guide: [docs/semantic-diff-guide.md](semantic-diff-guide.md)
- `hxrt` runtime: [docs/hxrt-runtime.md](hxrt-runtime.md)
- Release readiness checklist: [docs/release-readiness-checklist.md](release-readiness-checklist.md)
