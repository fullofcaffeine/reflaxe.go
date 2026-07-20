# Start Here

This is the shortest path from clone -> running app -> understanding policy
presets and explicit native boundaries.

## Terms

- [policy preset](glossary.md#policy-preset): compatible defaults selected by
  `portable` or `metal`.
- [native boundary](glossary.md#native-boundary): module-level Go authority
  declared with `@:goNative`.
- [`hxrt`](glossary.md#hxrt): runtime package copied into generated output.
- [semantic diff](glossary.md#semantic-diff): behavior parity test versus Haxe `--interp`.

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
- Real Haxe stdlib service: `examples/incident_api`
- Preset-neutral interop: `examples/interop_smoke`
- Clear specialization-policy comparison: `examples/worker_pool_select`
- Full app benchmark lanes: `examples/pulseforge`, `examples/fluxproxy`

Run all examples:

```bash
python3 test/run-examples.py
```

## Preset choice (practical)

Use:

```bash
-D reflaxe_go_profile=portable|metal
```

- `portable`: default policy bundle for portable Haxe source.
- `metal`: supported compatibility bundle for the `explicit` native-authority
  policy, eager specialization, strict raw boundaries, and fail-fast fallback.

`metal` is not required for good Go output. Start in `portable`; the compiler
should still emit Go-shaped fast paths whenever it can prove they preserve Haxe
semantics. Use typed `go.*`/extern APIs or `@:goNative` when source itself is
Go-native.

Detailed contract: [Native policy presets and semantic boundaries](native-policy-presets.md)

## Current release scope

Haxe.Go's current claim is a pre-1.0 beta for the exact portable workload,
toolchain, platform, operation/member, and trust boundaries listed in the
[generated compatibility matrix](compatibility-support-matrix.md). The
[machine manifest](compatibility-support-manifest.json) is authoritative, and
the [generated release status](compatibility-release-status.md) supplies the
wording used in release notes. Module-level evidence and the `metal`
compatibility preset do not widen that admission.

Verify the generated artifacts before release work:

```bash
npm run compatibility:verify
```

## Strict policy knobs

- `-D reflaxe_go_strict_examples`: forbids raw `__go__` in repo examples/snapshots.
- `-D reflaxe_go_strict`: forbids raw `__go__` in app code.
- `-D reflaxe_go_strict_policy=auto|on|off`: app strictness policy (`auto`
  follows the selected compatibility preset).

## Where is the stdlib?

Short answer: most of it is **not** copied into this repo under `std/`.

Beginner mental model:

1. You still start from the upstream Haxe stdlib.
2. `reflaxe.go` adds only the target-specific parts it needs:
   - canonical override source in `std/go/_std`
   - target support and typed runtime bindings in ordinary `std` modules
   - Go facade modules in `std/go`
   - runtime helper files in `runtime/hxrt`
   - compiler-emitted shims in `src/reflaxe/go/GoCompiler.hx`

Those are source-checkout paths. In an installed Haxelib package, the package
runner flattens only canonical `std/go/_std/**/*.hx` overrides into
`src/**/*.cross.hx`; support, bindings, and facades remain ordinary `.hx`
modules. The [canonical `_std` migration closeout](canonical-std-migration-closeout.md)
explains the mapping and its isolated install proof.

Why this is done:

- it avoids maintaining a full stdlib fork in this repository,
- it lets parity work land incrementally with tests,
- and it keeps target-specific behavior isolated.

Deep dive and ownership rule: [docs/ownership-rubric.md](ownership-rubric.md).
Current shim-by-shim decisions: [docs/stdlib-shim-rationale.md](stdlib-shim-rationale.md).

## Performance and release checks

```bash
npm run test:perf:go
npm run test:perf:hxrt-selective
npm run test:perf:apps
npm run security:go-tooling
npm run security:supply-chain
npm run release:policy
npm run release:license-policy
npm run compatibility:verify
npm run release:status
```

Release version and tested-source policy:
[docs/release-version-policy.md](release-version-policy.md)
Public compatibility and SemVer boundary:
[docs/public-contract.md](public-contract.md)
Licensing and generated-output policy: [LICENSING.md](../LICENSING.md)
Release checklist: [docs/release-readiness-checklist.md](release-readiness-checklist.md)
Performance budget policy: [docs/performance-budget-policy.md](performance-budget-policy.md)
Go tooling gate policy: [docs/go-tooling-gates.md](go-tooling-gates.md)
Supply-chain policy: [docs/supply-chain-policy.md](supply-chain-policy.md)

## Related docs

- Docs map: [docs/index.md](index.md)
- Glossary: [docs/glossary.md](glossary.md)
- Profiles: [docs/profiles.md](profiles.md)
- Native policy presets: [docs/native-policy-presets.md](native-policy-presets.md)
- Profile semantics guide: [docs/profile-semantics-guide.md](profile-semantics-guide.md)
- Examples matrix: [docs/examples-matrix.md](examples-matrix.md)
- Semantic diff guide: [docs/semantic-diff-guide.md](semantic-diff-guide.md)
- Compatibility and support matrix: [docs/compatibility-support-matrix.md](compatibility-support-matrix.md)
- Machine compatibility manifest: [docs/compatibility-support-manifest.json](compatibility-support-manifest.json)
- Generated compatibility release status: [docs/compatibility-release-status.md](compatibility-release-status.md)
- `hxrt` runtime: [docs/hxrt-runtime.md](hxrt-runtime.md)
- Go tooling release gates: [docs/go-tooling-gates.md](go-tooling-gates.md)
- Supply-chain policy: [docs/supply-chain-policy.md](supply-chain-policy.md)
- Vendored Reflaxe provenance: [docs/vendor-reflaxe-provenance.md](vendor-reflaxe-provenance.md)
- Public contract and SemVer boundary: [docs/public-contract.md](public-contract.md)
- Release version and source identity: [docs/release-version-policy.md](release-version-policy.md)
- Release readiness checklist: [docs/release-readiness-checklist.md](release-readiness-checklist.md)
