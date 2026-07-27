# Portable Option and Result Contract

## What it is

`reflaxe.std.Option<T>` and `reflaxe.std.Result<T, E>` are portable source
contracts that a Reflaxe target may represent efficiently without changing what
the Haxe program means.

The first contract is deliberately small:

- `Option<T>` has `Some(value)` and `None`;
- `Result<T, E>` has `Ok(value)` and `Err(error)`;
- generic helpers, nesting, and independently typed errors retain those cases;
- `Some(null)` is different from `None`;
- `Result<T, E>` keeps `E`; it does not silently replace it with Go's `error`.

The future standalone `reflaxe.std` package is the intended owner of the Haxe
module definitions. `reflaxe.go` does not publish those definitions today.
Tests supply the two tiny enums locally, just as `haxe.rust` does, so compiler
admission does not make a false package-availability claim.

## Why it exists

Go has no built-in `Option<T>`, and its conventional `(T, error)` result has one
specific error interface. Those facts do not make a nullable value equivalent
to portable Option, or a Go error pair equivalent to portable Result:

- nullable storage cannot distinguish `Some(null)` from `None`;
- converting an arbitrary `E` to `error` can lose its type, identity, or data;
- a generated wrapper can preserve the cases while a later typed optimizer
  chooses a more Go-shaped carrier.

The registry therefore admits the portable source contract independently from
native interop. This lets optimization improve generated Go without turning a
portable API into a native API.

## How it works

The compiler observes exact typed shapes and applies these rules:

| Surface | Eligible shape | Selected contract carrier | Fallback |
| --- | --- | --- | --- |
| `Option<T>` | `T` contains neither `Dynamic` nor an unresolved shape | `go_option` | `portable_option` |
| `Result<T, E>` | both `T` and `E` contain neither `Dynamic` nor an unresolved shape | `go_result` | `portable_result` |

Named type parameters are eligible. They retain their identity for later
monomorphization or typed generic lowering. Nested types are checked
recursively. A `Dynamic` or unresolved child rejects native admission and keeps
the portable tagged-enum fallback.

`go_option` means a typed Go-side presence carrier that still has two cases.
`go_result` means a typed two-parameter success/error carrier that preserves
both `T` and `E`. It is not native `go.Result<T>` and it is not `(T, error)`.

Both admitted carriers and both fallbacks declare:

- no Go package imports of their own;
- no direct `hxrt` feature requirement;
- eligibility for an otherwise `hxrt`-free program.

That statement is surface-local. Printing, strings, reflection, exceptions, or
other code in the same program may still require `hxrt`.

## Current generated shape

This admission task records semantic and representation authority before
lowering. The optimizer/runtime planner does not consume these entries until
`haxe_go-vfp.7.6`.

For now, supplied Option/Result declarations are retained and emitted through
the existing portable enum shape:

```go
type reflaxe__std__Result struct {
    tag    int
    params []any
}
```

That shape is the reported fallback, not the final typed carrier. The registry
report may say that an applied shape is admitted for `go_result`; it does not
claim that the current planner has already emitted that representation.
`test/fixtures/surface_contract_registry` locks this transition state so `.7.6`
can change it intentionally.

The retention rule matches only the exact `reflaxe.std.Option` and
`reflaxe.std.Result` enum identities that Haxe actually typed. It does not treat
the `reflaxe.*` namespace as compiler or representation authority.

## Native Go boundaries

There are three different contracts:

| Contract | Error type | Meaning |
| --- | --- | --- |
| `reflaxe.std.Result<T, E>` | caller-chosen `E` | portable tagged success/error value |
| `go.Result<T>` | current backend-local `go.Error` wrapper | explicit Go-target facade |
| `(T, error)` | Go `error` interface | native Go multi-result convention |

No implicit conversion exists in either direction between portable Result and
`go.Result`. Negative compile fixtures enforce that boundary.

This task intentionally does not ship a generic adapter. A truthful adapter
would need an explicit `E -> error` mapping and an explicit reverse mapping, and
native error identity must survive the Go bridge. The current `go.Result`
implementation can reduce native errors to messages; `haxe_go-vfp.9.2` owns the
first-class error/multi-result work needed before such an adapter can make a
stronger promise.

An application can write a domain-specific conversion today, but the lossy part
must be visible in its function signature and implementation. A profile or
compiler heuristic must never perform it silently.

Option likewise does not implicitly convert to `Null<T>` because that would
collapse `Some(null)` and `None`.

## Family and sibling status

`haxe.rust` supplies the precedent: fixture- or dependency-supplied
`reflaxe.std.Option/Result` are admitted per exact module and lower to native
Rust `Option<T>` / `Result<T,E>`. Go adopts the source distinction and proof
discipline, but not Rust's representation.

The in-repository `reflaxe.family.std` directory is currently a target-local
bootstrap mirror. Go and Rust use the same bootstrap version for payloads that
are not yet content-identical. These new registry entries therefore report
`target_local` family synchronization with no shared contract ID. They must not
claim cross-repository identity until the immutable-core work
(`haxe_go-tvjt`) and registry synchronization work (`haxe_go-vfp.7.8`) land.

## Evidence

Run:

```bash
npm run test:surface-contract-registry
python3 test/run-semantic-diff.py \
  --case portable_option_result_contract \
  --case portable_option_result_fallback_contract
python3 test/run-snapshots.py \
  --case negative/portable_result_not_go_result \
  --case negative/go_result_not_portable_result
```

The evidence covers success, typed errors, null payloads, nesting, generic
helpers, eligible and fallback shapes, exact generated fallback declarations,
profile-independent reports, and the native boundary.
