<p align="center">
  <img src="assets/haxe.go.logo.png" alt="reflaxe.go logo" width="220" />
</p>

# reflaxe.go

[![CI Harness](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/ci-harness.yml/badge.svg)](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/ci-harness.yml)
[![Quality](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/ci-quality.yml/badge.svg)](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/ci-quality.yml)
[![Security](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/security-static-analysis.yml/badge.svg)](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/security-static-analysis.yml)
[![License: GPL-3.0](https://img.shields.io/badge/license-GPL--3.0-blue.svg)](LICENSE)

Haxe 4.3.7 -> Go compiler target built on Reflaxe.

Write Haxe, generate idiomatic-enough Go modules, and ship native binaries with profile-based control over portability vs Go-first output style.

If you want Haxe productivity with a serious Go delivery pipeline, this is that target.

## Why reflaxe.go

- One Haxe codebase, multiple Go profiles (`portable`, `gopher`, `metal`).
- Generated Go module output (`go.mod` + `main.go` + `hxrt`) with backend `go build` by default.
- Strong verification harness: snapshots, stdlib sweep, semantic diff, examples matrix, and perf checks.

## Quick Start

1. Install project dependencies:

```bash
npm install
```

2. Run core compiler snapshots:

```bash
npm test
```

3. Run the full local CI harness:

```bash
python3 test/run-ci.py
```

4. Compile and run a real example:

```bash
npm run dev:hx -- --project examples/tui_todo --profile portable --action run
```

5. Scaffold your own project:

```bash
npm run dev:new-project -- ./my_haxe_go_app
```

## First 5 Minutes

Run a real app and verify the compiler quality gates:

```bash
npm install
npm run dev:hx -- --project examples/tui_todo --profile portable --action run
python3 test/run-snapshots.py
python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_full.txt --strict --go-test
```

Then switch profiles and compare output/runtime behavior:

```bash
npm run dev:hx -- --project examples/tui_todo --profile gopher --action run
npm run dev:hx -- --project examples/tui_todo --profile metal --action run
```

## Profiles

Use `-D reflaxe_go_profile=portable|gopher|metal`.

| Profile | Best for | Contract |
| --- | --- | --- |
| `portable` (default) | Haxe-first teams | Portability-first semantics and lowest migration risk |
| `gopher` | Go-aware teams | Go-first style and safe lowering optimizations without intentional semantic drift |
| `metal` (experimental) | Low-level interop needs | `gopher` + strict boundary defaults + typed low-level interop lane |

`idiomatic` was removed and now fails fast. Use `gopher`.

Details: [docs/profiles.md](docs/profiles.md)

## Flagship Examples

- [examples/tui_todo](examples/tui_todo/README.md): complex single-codebase app compiled across all profiles.
- [examples/profile_storyboard](examples/profile_storyboard/README.md): compact profile adapter/storyboard reference.
- [examples/README.md](examples/README.md): examples overview.
- [docs/examples-matrix.md](docs/examples-matrix.md): exact compile/run/artifact matrix.

## Most Useful Commands

- New project generator: `npm run dev:new-project -- ./my_haxe_go_app`
- Compile/run wrapper: `npm run dev:hx -- --project <dir> --profile <portable|gopher|metal> --action <compile|run|build|test|vet|fmt>`
- Snapshots: `python3 test/run-snapshots.py`
- Upstream stdlib sweep (strict + go test): `python3 test/run-upstream-stdlib-sweep.py --modules-file test/upstream_std_modules_full.txt --strict --go-test`
- Semantic differential checks: `python3 test/run-semantic-diff.py`
- Examples matrix: `python3 test/run-examples.py`
- Profile perf harness: `npm run test:perf:go`
- Release visibility checks: `npm run release:status`

## Verification and Delivery

- Verification surface is built in, not optional:
  - snapshots: `test/run-snapshots.py`
  - stdlib inventory sweeps: `test/run-upstream-stdlib-sweep.py`
  - semantic differential checks: `test/run-semantic-diff.py`
  - examples matrix + generated output drift checks: `test/run-examples.py`
- CI is split into dedicated harness/quality/security/release workflows:
  - [.github/workflows/ci-harness.yml](.github/workflows/ci-harness.yml)
  - [.github/workflows/ci-quality.yml](.github/workflows/ci-quality.yml)
  - [.github/workflows/security-static-analysis.yml](.github/workflows/security-static-analysis.yml)
  - [.github/workflows/examples-artifacts.yml](.github/workflows/examples-artifacts.yml)
- Current support inventory and known tradeoffs are explicit:
  - [docs/feature-support-matrix.md](docs/feature-support-matrix.md)

## Output Model

A typical compile emits a Go module like:

```text
out/
  go.mod
  main.go
  hxrt/
    hxrt.go
```

By default, backend compile runs `go build` after codegen. Use `-D go_no_build` when you want codegen-only flows.

## Documentation

- Start here: [docs/start-here.md](docs/start-here.md)
- Feature support and coverage inventory: [docs/feature-support-matrix.md](docs/feature-support-matrix.md)
- Defines reference: [docs/defines-reference.md](docs/defines-reference.md)
- Profile admission criteria: [docs/profile-admission-criteria.md](docs/profile-admission-criteria.md)
- Snapshot policy: [docs/snapshot-policy.md](docs/snapshot-policy.md)
- Release visibility and gates: [docs/release-visibility.md](docs/release-visibility.md)
- Multi-package output evaluation: [docs/multi-package-output-evaluation.md](docs/multi-package-output-evaluation.md)
- Future compiler template/patterns: [docs/compiler-target-template.md](docs/compiler-target-template.md)
- Security policy: [SECURITY.md](SECURITY.md)
- Changelog: [CHANGELOG.md](CHANGELOG.md)

## CI and Security

- Integrated harness + release flow: [.github/workflows/ci-harness.yml](.github/workflows/ci-harness.yml)
- Quality gate workflow: [.github/workflows/ci-quality.yml](.github/workflows/ci-quality.yml)
- Static security analysis: [.github/workflows/security-static-analysis.yml](.github/workflows/security-static-analysis.yml)
- Gitleaks workflow: [.github/workflows/security-gitleaks.yml](.github/workflows/security-gitleaks.yml)
- Example artifact/release publishing: [.github/workflows/examples-artifacts.yml](.github/workflows/examples-artifacts.yml)

Local pre-commit hook setup:

```bash
npm run hooks:install
```

This enforces staged local-path leak guard, staged gitleaks scan, and staged Haxe auto-format.

## Project Status

`reflaxe.go` is currently pre-1.0 and moving quickly. The quality bar is enforced by the harnesses above, and profile/runtime parity is tracked continuously through snapshots and stdlib sweeps.

Releases: [github.com/fullofcaffeine/reflaxe.go/releases](https://github.com/fullofcaffeine/reflaxe.go/releases)

## License

GPL-3.0-only. See [LICENSE](LICENSE).
