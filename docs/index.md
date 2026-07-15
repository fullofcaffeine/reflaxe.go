# Documentation Map

This page helps you find the right docs quickly.

## Terms

- [portable](glossary.md#portable): default portable policy preset and semantic baseline.
- [metal](glossary.md#metal-compatibility-preset): supported compatibility policy preset.
- [native boundary](glossary.md#native-boundary): explicit Go-native module authority.
- [semantic diff](glossary.md#semantic-diff): behavior comparison between Haxe `--interp` and generated Go.
- [`hxrt`](glossary.md#hxrt): runtime support package copied into generated output.

## If you are new

1. [README.md](../README.md)
2. [docs/start-here.md](start-here.md)
3. [docs/profiles.md](profiles.md)
4. [docs/examples-matrix.md](examples-matrix.md)

## Core usage docs

- [Start here](start-here.md)
- [Profiles](profiles.md)
- [Native policy presets and semantic boundaries](native-policy-presets.md)
- [Profile semantics guide](profile-semantics-guide.md)
- [Defines reference](defines-reference.md)
- [Examples matrix](examples-matrix.md)
- [Go concurrency + interop guide](go-concurrency-interop-guide.md)
- [Concurrency contract](concurrency-contract.md)

## Testing and quality docs

- [Compatibility and support matrix](compatibility-support-matrix.md)
- [Machine-readable compatibility manifest](compatibility-support-manifest.json)
- [Generated compatibility release status](compatibility-release-status.md)
- [Semantic diff guide](semantic-diff-guide.md)
- [Snapshot policy](snapshot-policy.md)
- [Examples QA contract](examples-qa-contract.md)
- [Thinking levels](thinking-levels.md)
- [Implementation evidence inventory](feature-support-matrix.md)
- [Supported toolchain policy](toolchain-policy.md)
- [Release version and source-identity policy](release-version-policy.md)
- [Licensing and generated-output policy](../LICENSING.md)
- [Release readiness checklist](release-readiness-checklist.md)
- [Security dependency audit](security-dependency-audit.md)
- [Go tooling release gates](go-tooling-gates.md)
- [Supply-chain policy](supply-chain-policy.md)
- [Generated-output confinement](generated-output-confinement.md)
- [Vendored Reflaxe provenance](vendor-reflaxe-provenance.md)
- [Performance budget policy](performance-budget-policy.md)
- [Compiler debt baseline and ratchet](compiler-debt-ratchet.md)
- [Perf warning triage](perf-warning-triage.md)
- [Release visibility](release-visibility.md)

## Runtime and stdlib docs

- [Ownership rubric](ownership-rubric.md)
- [`hxrt` runtime](hxrt-runtime.md)
- [Portable root `Sys` contract](portable-sys-contract.md)
- [Concurrency contract](concurrency-contract.md)
- [Selective `hxrt` runtime](hxrt-selective-runtime.md)
- [Stdlib shim rationale](stdlib-shim-rationale.md)
- [Cross overrides and hardening](cross-overrides-and-hardening.md)
- [Canonical `_std` migration closeout](canonical-std-migration-closeout.md)
- [Family raw-injection authority alignment](spikes/family-raw-injection-authority-alignment.md)
- [Portable stdlib parity program](portable-stdlib-parity-program.md)
- [Portable module mapping contract](portable-module-mapping-contract.md)

## Contract/spec docs

- [Portable canonical contract](portable-canonical-contract.md)
- [Portable semantics v1](portable-semantics-v1.md)
- [Native policy presets and semantic boundaries](native-policy-presets.md)
- [Profile admission criteria](profile-admission-criteria.md)

## Deep technical references

- [Multi-package output evaluation](multi-package-output-evaluation.md)
- [Compiler target template](compiler-target-template.md)
- [Go extern generator](goextern.md)
- [Known gaps](known-gaps.md)

## Related docs

- [Glossary](glossary.md)
- [README](../README.md)
- [Start here](start-here.md)
