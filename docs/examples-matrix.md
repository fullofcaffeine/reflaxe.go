# Examples Matrix

`reflaxe.go` ships two canonical examples designed to show one Haxe codebase compiled into all supported profiles.

| Example | portable | gopher | metal | Purpose |
| --- | --- | --- | --- | --- |
| `examples/profile_storyboard` | Yes | Yes | Yes | Compact profile-idiom reference with profile runtime adapters. |
| `examples/tui_todo` | Yes | Yes | Yes | Complex app reference with deterministic pseudo-TUI flow and additive profile capabilities. |

## Build and run matrix

```bash
python3 test/run-examples.py
```

This runs:

- `compile.<profile>.ci.hxml` + `go test` + `go run` expected output checks
- `compile.<profile>.hxml` + `go test` + `go run` expected output checks
- generated tree drift checks (`generated/<profile>` vs `out_<profile>`)

## Generated Go + binaries

Refresh committed generated trees:

```bash
python3 test/run-examples.py --bless-generated
```

Build cross-platform binaries from committed generated trees:

```bash
bash scripts/examples/build-binaries.sh
```

Binary matrix targets:

- `linux/amd64`
- `linux/arm64`
- `darwin/arm64`
- `windows/amd64`
