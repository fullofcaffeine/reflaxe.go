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

## CI artifact/release automation

- Workflow: `.github/workflows/examples-artifacts.yml`
- Triggers:
  - `push` to `master`: builds matrix and uploads workflow artifacts.
  - `push` tags: builds matrix, uploads workflow artifacts, and publishes release assets.
  - `workflow_dispatch`: manual artifact build/upload.
- Why this uses artifact upload/download:
  - The workflow intentionally splits `build` and `release` jobs so release publishing only runs on tag pushes and with `contents: write` permissions.
  - Job filesystems are isolated in GitHub Actions, so assets must cross jobs via `upload-artifact`/`download-artifact`.
  - Downloaded artifact layout is not guaranteed to preserve the original `dist/...` prefix, so the release job normalizes discovered files into a deterministic staging directory before `action-gh-release`.
- Release assets:
  - `examples-<tag>.tar.gz`
  - `examples-<tag>.tar.gz.sha256`
  - `manifest.json`
  - `checksums.txt`
