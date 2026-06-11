# Glossary

Plain-language definitions for terms that appear across `reflaxe.go` docs.

## Profile

A named build contract (`portable` or `metal`) that controls semantics and policy defaults.

Reference: [docs/profiles.md](profiles.md)

## Portable profile

Default profile and normal product path for Haxe code that should keep portable
semantics while still generating readable, performant Go where the compiler can
prove the lowering is safe.

Reference: [docs/profiles.md](profiles.md#matrix)

## Metal profile

Opt-in profile for explicit Go-native authoring, stricter defaults, and
fail-fast native-lane checks. `metal` is not required for good Go output; use it
when you deliberately want Go-specific APIs or constraints.

Reference: [docs/profiles.md](profiles.md#matrix)

## Contract

A behavior promise the compiler enforces (for example profile rules, portable semantics, lane policy).

References:
- [docs/portable-canonical-contract.md](portable-canonical-contract.md)
- [docs/profile-semantics-guide.md](profile-semantics-guide.md)

## Lane

A constrained subset of code or behavior in a build. In examples, `go_native` is a runtime adapter variant lane. In compiler policy, `@:goMetal` marks metal-clean enforcement islands.

References:
- [docs/examples-matrix.md](examples-matrix.md)
- [docs/profiles.md](profiles.md#goMetal-lanes-portable-builds)

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

Reference: [docs/profiles.md](profiles.md#portable-convergence-optimizer-controls)

## Fallback

A safe alternative path used when a stricter typed/native lowering cannot be applied.

Reference: [docs/profiles.md](profiles.md#portable-convergence-optimizer-controls)

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
