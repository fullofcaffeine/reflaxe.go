# Glossary

Plain-language definitions for terms that appear across `reflaxe.go` docs.

## Policy preset

A named bundle of policy defaults. `reflaxe_go_profile=portable|metal` is the
compatible public selector, but the selected preset does not create another
compiler engine or silently choose source semantics.

Reference: [Native policy presets and semantic boundaries](native-policy-presets.md)

## Portable

The default product path and `portable_default` policy preset. Ordinary Haxe
source keeps portable semantics while the compiler may emit direct,
performant Go where it can prove equivalence.

Reference: [Profiles](profiles.md)

## Metal compatibility preset

The supported `metal_compatibility` bundle: explicit native authority, eager
specialization, fail-fast fallback, and strict raw-boundary defaults. It is not
a second semantic product and is not required for good Go output.

Reference: [Profiles](profiles.md)

## Native boundary

An owning Haxe module explicitly declared with `@:goNative`, or native intent
expressed through typed `go.*`/extern APIs. `@:goMetal` is the compatibility
metadata alias.

Reference: [Native policy presets and semantic boundaries](native-policy-presets.md#gonative-module-boundaries)

## Contract

A behavior promise the compiler enforces, such as portable semantics or a
typed Go-native API contract. Policy presets choose defaults around a contract;
they are not contracts by themselves.

References:
- [docs/portable-canonical-contract.md](portable-canonical-contract.md)
- [docs/profile-semantics-guide.md](profile-semantics-guide.md)

## Lane

A constrained subset of code or behavior in a build. In examples, `go_native`
is a runtime adapter variant lane. Compiler policy now calls module-level
authority a native boundary and uses canonical `@:goNative`.

References:
- [docs/examples-matrix.md](examples-matrix.md)
- [docs/native-policy-presets.md](native-policy-presets.md#gonative-module-boundaries)

## go_native

An app variant used in flagship examples to run Go-first runtime adapter paths
(for example channel/select worker flows). It is not a compiler profile.

Reference: [docs/examples-matrix.md](examples-matrix.md#terms)

## Go-native

APIs or behavior tied specifically to Go, such as `go.Chan`, `go.Select`, typed
Go extern metadata, or generated goroutine/channel paths. Go-native code can be
the right choice for a Go-only deployment, but it is not portable across non-Go
Haxe targets.

References:
- [docs/profiles.md](profiles.md)
- [docs/go-concurrency-interop-guide.md](go-concurrency-interop-guide.md)

## Go-first

A design choice that prioritizes Go-specific APIs, constraints, or runtime
behavior over cross-target portability. In this repo, Go-first does not mean
"better by default"; it means the code is intentionally closer to Go.

Reference: [docs/profile-semantics-guide.md](profile-semantics-guide.md)

## Hot path

Code that runs most frequently during normal workload. Optimizing hot paths usually gives the biggest runtime gains.

Reference: [docs/examples-matrix.md](examples-matrix.md#terms)

## `hxrt`

Runtime support package copied into generated Go output (`<go_output>/hxrt`). It hosts shared runtime helpers used by generated code.

Reference: [docs/hxrt-runtime.md](hxrt-runtime.md)

## Shim

Compiler-emitted or runtime-assisted helper that bridges a Haxe API/semantic behavior to Go behavior.

References:
- [docs/stdlib-shim-rationale.md](stdlib-shim-rationale.md)
- [docs/portable-module-mapping-contract.md](portable-module-mapping-contract.md)

## Typed specialization

Lowering a generic/native facade call into a concrete typed Go path when the compiler can prove types safely.

Reference: [Native specialization](native-policy-presets.md#native-specialization)

## Fallback

A safe alternative path used when a stricter typed/native lowering cannot be applied.

Reference: [Native fallback](native-policy-presets.md#native-fallback)

## Semantic diff

Harness that compares runtime behavior between Haxe `--interp` and generated Go output.

Reference: [docs/semantic-diff-guide.md](semantic-diff-guide.md)

## Snapshot test

Harness that checks generated Go text/shape for deterministic output changes.

Reference: [docs/snapshot-policy.md](snapshot-policy.md)

## Related docs

- [Documentation map](index.md)
- [Start here](start-here.md)
- [Profiles](profiles.md)
