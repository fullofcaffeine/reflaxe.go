<p align="center">
  <img src="assets/haxe.go.logo.png" alt="reflaxe.go logo" width="220" />
</p>

# reflaxe.go

[![CI Harness](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/ci-harness.yml/badge.svg)](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/ci-harness.yml)
[![Quality](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/ci-quality.yml/badge.svg)](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/ci-quality.yml)
[![Security](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/security-static-analysis.yml/badge.svg)](https://github.com/fullofcaffeine/reflaxe.go/actions/workflows/security-static-analysis.yml)
[![License: GPL-3.0](https://img.shields.io/badge/license-GPL--3.0-blue.svg)](LICENSE)

Haxe 4.3.7 -> Go compiler target built on Reflaxe.

Write Haxe, generate readable Go, and keep semantics visible in source:

- portable Haxe is the default product path;
- typed `go.*`/extern APIs and `@:goNative` mark explicit Go-native boundaries;
- `portable|metal` remains a compatible build selector, with `metal` defined as
  a convenience policy preset rather than a second semantic product.

`metal` is not required for good Go output. It currently bundles the `explicit`
native-authority policy, eager specialization, fail-fast fallback, and strict
raw boundaries. Each policy can also be selected independently.

## Start here

If you are new, read these first:

1. [docs/start-here.md](docs/start-here.md)
2. [docs/index.md](docs/index.md)
3. [docs/glossary.md](docs/glossary.md)

## Terms

- [policy preset](docs/glossary.md#policy-preset): compatible bundle of build-policy defaults.
- [portable](docs/glossary.md#portable): default portable Haxe semantics and policy preset.
- [metal](docs/glossary.md#metal-compatibility-preset): compatible Go-first policy preset.
- [native boundary](docs/glossary.md#native-boundary): explicit `@:goNative` module authority.
- [Go-native](docs/glossary.md#go-native): APIs or behavior tied specifically to Go.
- [`hxrt`](docs/glossary.md#hxrt): runtime support package copied into generated Go output.

## Quick start

```bash
npm install
npm run hooks:install
npm run dev:hx -- --project examples/tui_todo --profile portable --action run
```

That gets a fresh checkout to a running generated-Go example. Use `portable`
first and declare Go-specific modules with typed APIs and `@:goNative`.

## Validation commands

After the first example runs, use these commands to validate changes:

```bash
# full snapshot suite
npm test

# examples matrix
python3 test/run-examples.py

# full local CI harness
python3 test/run-ci.py
```

## Current status and caveats

Haxe.Go is a pre-1.0 beta for pinned, application-qualified portable workloads on the admitted toolchain, platform, and operation/member surface.

That sentence is deliberately narrower than "all green modules are supported."
The authoritative [machine manifest](docs/compatibility-support-manifest.json),
its [human support matrix](docs/compatibility-support-matrix.md), and the
[generated release status](docs/compatibility-release-status.md) name the exact
admitted scope. Anything unlisted is excluded; `metal`, Go-native APIs,
concurrency, networking, non-canonical platforms, and untrusted-source builds
remain outside the current release claim even when implementation evidence
exists.

Before using the compiler as a release dependency, also read:

- [docs/toolchain-policy.md](docs/toolchain-policy.md) for supported Haxe, Go,
  and Node versions and the distinction between the generated Go language floor
  and patched build toolchains.
- [docs/release-readiness-checklist.md](docs/release-readiness-checklist.md) for
  the required validation gates.
- [docs/known-gaps.md](docs/known-gaps.md) for target-sensitive APIs,
  advanced Go extern interop limits, performance warning policy, and the
  current single-package output decision.

## Which preset should I choose?

Choose one compatibility preset per build. Source APIs and `@:goNative`, not
the preset, define where semantics become Go-native.

| If most of this build wants... | Choose |
| --- | --- |
| Normal Haxe authoring, cross-target-friendly code, and safe Go output | `portable` |
| Compatibility bundle for the `explicit` authority policy, eager specialization, strict raw boundaries, and fail-fast fallback | `metal` |

### Can I mix portable and Go-native code?

Yes, with one important distinction:

- A single compiler invocation selects one preset: `portable` or `metal`.
- A codebase can mix layers. Keep shared/domain code portable, then isolate
  Go-specific code behind typed adapters, `go.*` APIs, or `@:goNative` modules.
- Under guarded native authority, unscoped native usage can be warned about or
  rejected. The compiler can also lower portable code to Go-shaped output when it
  can prove the faster lowering preserves portable semantics.
- The `metal` selector supplies stricter Go-native defaults but does not
  reclassify ordinary Haxe APIs or create another semantic backend.

So `portable` is the default for most projects, including projects that contain
small Go-specific islands. Existing `metal` builds remain supported; new code
should make native ownership explicit at API/module boundaries.

Detailed behavior and policy knobs: [docs/native-policy-presets.md](docs/native-policy-presets.md)

## Example apps

| Example | Preset support | Why run it |
| --- | --- | --- |
| [examples/tui_todo](examples/tui_todo/README.md) | portable only | Portable CLI baseline and deterministic output contract. |
| [examples/profile_storyboard](examples/profile_storyboard/README.md) | portable only | Portable dashboard/reporting baseline. |
| [examples/interop_smoke](examples/interop_smoke/README.md) | portable + metal | Typed interop patterns (`@:go.import`, receiver/name metadata). |
| [examples/worker_pool_select](examples/worker_pool_select/README.md) | portable + metal | Go channel/select usage and specialization-policy comparison. |
| [examples/pulseforge](examples/pulseforge/README.md) | portable + metal | Flagship app with `core` and `go_native` variants. |
| [examples/fluxproxy](examples/fluxproxy/README.md) | portable + metal | Flagship proxy app with `core` and `go_native` variants. |

Full matrix: [docs/examples-matrix.md](docs/examples-matrix.md)

## Most useful commands

- Snapshots: `python3 test/run-snapshots.py`
- Semantic diff: `python3 test/run-semantic-diff.py`
- Examples matrix: `python3 test/run-examples.py`
- Full CI harness: `python3 test/run-ci.py`
- Profile perf harness: `npm run test:perf:go`
- App perf harness: `npm run test:perf:apps`
- New project: `npm run dev:new-project -- ./my_haxe_go_app`
- Compile/run wrapper: `npm run dev:hx -- --project <dir> --profile <portable|metal> --action <compile|run|build|test|vet|fmt>`

## Output model

The generated Go directory is controlled by `-D go_output=<dir>` in the
selected `.hxml` file. Example projects usually write to `out_<profile>` or
`out_<profile>_ci` so portable and metal builds do not overwrite each other.

The quick-start command above uses the portable TUI todo config, so it emits:

```text
examples/tui_todo/out_portable/
  go.mod
  main.go
  module_<haxe_module>.go
  hxrt/
    *.go
```

Generated output directories are safe to delete. They are regenerated on the
next compile or run.

You can override the generated output directory through the dev wrapper:

```bash
npm run dev:hx -- --project examples/tui_todo --profile portable --out /tmp/tui_todo_go --action run
```

By default the backend runs `go build` after codegen and fails the Haxe
invocation if that build cannot launch or exits nonzero. Use `-D go_no_build`
only for explicit codegen-only workflows that own their Go build/test stage.

## Related docs

- Docs map: [docs/index.md](docs/index.md)
- Glossary: [docs/glossary.md](docs/glossary.md)
- Start guide: [docs/start-here.md](docs/start-here.md)
- Profiles and policy: [docs/profiles.md](docs/profiles.md)
- Native policy presets and semantic boundaries: [docs/native-policy-presets.md](docs/native-policy-presets.md)
- Profile semantics guide: [docs/profile-semantics-guide.md](docs/profile-semantics-guide.md)
- Portable contract: [docs/portable-canonical-contract.md](docs/portable-canonical-contract.md)
- Versioned semantics: [docs/portable-semantics-v1.md](docs/portable-semantics-v1.md)
- `hxrt` runtime: [docs/hxrt-runtime.md](docs/hxrt-runtime.md)
- Admitted compatibility scope: [docs/compatibility-support-matrix.md](docs/compatibility-support-matrix.md)
- Machine-readable compatibility authority: [docs/compatibility-support-manifest.json](docs/compatibility-support-manifest.json)
- Generated release compatibility status: [docs/compatibility-release-status.md](docs/compatibility-release-status.md)
- Implementation evidence inventory: [docs/feature-support-matrix.md](docs/feature-support-matrix.md)
- Supported toolchains: [docs/toolchain-policy.md](docs/toolchain-policy.md)
- Defines reference: [docs/defines-reference.md](docs/defines-reference.md)
- Semantic diff guide: [docs/semantic-diff-guide.md](docs/semantic-diff-guide.md)
- Release version and source identity: [docs/release-version-policy.md](docs/release-version-policy.md)
- Release readiness checklist: [docs/release-readiness-checklist.md](docs/release-readiness-checklist.md)
