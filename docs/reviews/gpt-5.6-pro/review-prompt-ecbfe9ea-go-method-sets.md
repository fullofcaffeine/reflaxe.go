# Independent pre-refactor review: Go extern method-set satisfaction

Review Haxe.Go source commit
`ecbfe9ea61e086d3afd069e1017deb711604bf0e` at maximum reasoning depth. Treat
that pushed Git commit as the source authority. This is a pre-refactor design
checkpoint for Bead `haxe_go-vfp.8.4.1`; no implementation has been selected,
and an uncommitted working tree must not be treated as evidence.

Repository: `https://github.com/fullofcaffeine/reflaxe.go`

## Provenance requirement

This prompt is intended for a genuine GPT-5.6 Pro independent reviewer. Name
the exact model route and reasoning setting in the response. If GPT-5.6 Pro is
not actually selected, stop and report the mismatch. Do not silently substitute
GPT-5.5 Pro, an older model, or a local written review while claiming GPT-5.6
provenance.

## Decision to review

Go accepts a concrete pointer such as `*bytes.Buffer` wherever its method set
satisfies `io.Writer`. Haxe normally requires a nominal relationship, while
`goextern` currently emits cross-package named types and non-empty interfaces
as `Dynamic`. We need an ordinary typed Haxe call such as this to compile with
no user-written cast and lower to the unchanged native Go value:

```haxe
function encode(writer:goextern.io.Writer):Void;

function useBuffer(buffer:goextern.bytes.Buffer):Void {
	encode(buffer);
}
```

The source-level convenience must be backed by a real proof. A type with a
missing method, a mismatched parameter or result, the wrong package identity,
or only a pointer-receiver method when a Go value is required must fail with an
actionable Haxe.Go diagnostic before `go test` becomes the type checker.

Adjudicate the smallest architecture that lets generated and deliberately
hand-authored typed externs express this relationship without adding portable
Haxe structural subtyping.

## Current facts at the pinned commit

1. `tools/goextern/main.go` loads one Go package per invocation. It maps a named
   type from another package to `Dynamic` with reason
   `external_named_type`, and maps a non-empty anonymous interface to `Dynamic`
   with reason `non_empty_interface`.
2. `collectMethods` enumerates both `types.NewMethodSet(T)` and
   `types.NewMethodSet(*T)`, merges them into one Haxe method list, and discards
   whether each method came from the value or pointer set. This is convenient
   for calls but cannot prove Go assignability.
3. Go's `go/types` package is already available in `goextern`. It can supply
   exact package-qualified signatures, embedded-interface completion, and
   `types.Implements` / `types.Satisfies` checks. Reuse that authority instead
   of recreating the Go language rules by string matching.
4. The Haxe compiler maps an imported extern class to a Go pointer such as
   `*bytes.Buffer`, and an imported extern interface to a Go interface such as
   `io.Writer`. Once Haxe accepts the call, ordinary argument lowering already
   leaves an unrelated imported extern value unchanged.
5. Haxe does not accept `Buffer` where nominal `Writer` is expected unless the
   binding surface supplies an `implements` relationship, a structural Haxe
   shape, or an implicit abstract conversion. The backend cannot repair a call
   that the Haxe typer rejected before macro lowering.
6. `src/reflaxe/go/ast/GoType.hx` structurally models named types, pointers,
   functions, and interface method signatures, but it does not model a named
   concrete type's value/pointer method sets or a proof that one named type
   satisfies another named interface.
7. `src/reflaxe/go/compiler/GoTypeMapper.hx` follows ordinary extern class and
   interface metadata, and has an extern-backed abstract path. That abstract
   path currently prepends `*` even when its carrier is an interface, so it is
   not already a safe implicit-interface solution.
8. `docs/goextern.md` explicitly describes `Dynamic` as an honest fallback and
   records `fmt.Fprint`'s `io.Writer` parameter as
   `external_named_type`. Any change must update the deterministic fallback
   report and fixtures rather than silently hiding unsupported boundaries.

## Non-negotiable product and ownership constraints

- Limit this feature to explicit Go-native authority: imported typed externs,
  generated bindings, and `@:goNative` source boundaries. Do not change
  portable Haxe assignability or make ordinary Haxe interfaces structural.
- Application and example code must not need `cast`, `untyped`, raw `__go__`,
  `Dynamic`, reflection, or a handwritten Go adapter for a proven method set.
- Do not keep a compiler name list for `bytes.Buffer`, `io.Writer`, or standard
  packages. Package path and type identity must be structural inputs.
- Preserve the distinction between method sets of `T` and `*T`, including
  promoted methods. Do not claim that a value implements an interface merely
  because its pointer does.
- Compare exact Go parameter and result types, variadic shape, package identity,
  and method names. Respect Go's exported/unexported method rules and embedded
  interfaces. Unsupported generic or ambiguous shapes must fail closed or stay
  an explicitly reported `Dynamic` fallback.
- Use `go/types` as the source authority in `goextern`. If the Haxe compiler
  consumes generated proof metadata, keep that metadata typed, deterministic,
  versioned, and validated before emitting the unchanged native value.
- Prefer the AST-first compiler path. Do not add `GoRaw` or an `hxrt` runtime
  conversion; Go interface assignment needs no runtime adapter.
- Generated bindings must remain deterministic across runs, safe when packages
  are generated in different orders, and clear about whether a dependency
  package or interface catalog is required.
- Hand-authored externs need a documented, narrow path. Do not imply that a
  user-written false metadata claim has been verified against installed Go
  source if only generated signature metadata was compared.

## Candidate Haxe typing boundaries

These are discussion candidates, not approved designs. Reject or combine them
as the evidence requires.

### A. Multi-package generation plus nominal `implements`

Extend `goextern` so one deterministic generation unit knows the selected
concrete packages and named interfaces. Use `go/types` to emit
`extern class Buffer implements goextern.io.Writer` only when the mapped Go
pointer really implements the interface. This gives Haxe a normal nominal
relationship and requires little backend adaptation.

Questions: Is requiring all participating packages in one generation unit an
acceptable first-class contract? How are interface method signatures with Go
multi-results represented identically on the concrete class and Haxe
interface? How is output kept order-independent when packages are regenerated
separately?

### B. Generated proof metadata plus an implicit target-interface abstract

Represent a named Go interface with a Haxe abstract that permits an implicit
generic conversion only inside the Go-native binding layer. Attach or reference
versioned method-set facts generated by `go/types`; after Haxe typing, the
backend validates source method set against target interface and erases the
conversion so the original native value is emitted unchanged.

Questions: Can this be made fail-closed without making every Haxe value
convertible to every Go interface on non-Go targets? Should validation live in
a typing macro, the Reflaxe compiler, or both? What exact typed metadata shape
avoids parsing an ad-hoc Go-signature mini-language?

### C. Metadata-backed structural Haxe interface typedef

Generate a method-only anonymous Haxe shape so the Haxe typer performs duck
typing, then retain imported Go interface identity through metadata on the
typedef/abstract rather than lowering the anonymous shape to
`map[string]any`. Use `go/types` metadata to validate facts that Haxe cannot
express exactly, especially multi-results and pointer receiver provenance.

Questions: Does following typedefs make native package identity too fragile?
Can concrete and interface method declarations share canonical multi-result
types without broad `Dynamic`? Would this accidentally leak structural
assignability into portable Haxe surfaces?

## Required first-slice contracts

Recommend exact fixtures and diagnostic ownership for all of these:

1. Positive generated or hand-authored `*bytes.Buffer -> io.Writer` call with no
   cast. Generated Go must pass the native pointer directly and compile/run.
2. Positive value-receiver method: both `T` and `*T` satisfy the interface.
3. Negative pointer-only method: `T` fails while `*T` succeeds.
4. Negative missing method.
5. Negative same method name with a parameter mismatch.
6. Negative same unqualified type name from a different package.
7. Embedded interface completion.
8. Exported and unexported method behavior across package boundaries.
9. Deterministic `goextern` output and fallback-report drift showing that a
   supported interface parameter is no longer `external_named_type` while
   unsupported shapes remain visible.
10. A compiler snapshot proving imports use structural package paths and the
    interface conversion adds no cast, wrapper, reflection, or runtime helper.

The positive runtime fixture may receive its concrete value through a typed
extern function or a small test-only Go package; it must not use application
raw injection. Negative fixtures must stop during Haxe/goextern validation, not
merely preserve a known `go test` failure.

## Required source evidence

Read these files at the pinned commit:

- `tools/goextern/main.go` and `tools/goextern/main_test.go`;
- `test/run-goextern-fixtures.py` and the committed `fmt`, `context`, and `sync`
  fixtures under `test/fixtures/goextern`;
- `src/reflaxe/go/ast/GoType.hx`;
- `src/reflaxe/go/compiler/GoTypeMapper.hx`;
- `src/reflaxe/go/GoCompiler.hx`, especially ordinary call-argument lowering,
  `upcastIfNeeded`, extern package/type mapping, and metadata validation;
- `docs/goextern.md`, `docs/typed-go-ir.md`, `docs/known-gaps.md`, and the public
  compatibility/support matrix;
- `test/snapshot/go_native/extern_metadata_mapping`,
  `extern_tuple_return`, and the negative-fixture conventions in
  `test/run-snapshots.py`.

Do not use go2hx implementation details as authority: it is a separate project
and is not part of this repository. Go's language specification and `go/types`
behavior are the relevant external semantics.

## Questions the review must answer

1. Which candidate boundary, or which smaller alternative, should own the
   first production slice? Give the end-to-end data flow from `go/types` through
   generated Haxe typing to typed Go IR and final unchanged Go expression.
2. What is the minimal versioned structural model for named type identity,
   receiver kind, method signatures, embedded interfaces, and proof results?
3. Where must invalid bindings fail, and which diagnostics should be produced
   by `goextern`, a Haxe typing macro, or the backend?
4. How should separately generated packages share interface identity without
   generation order becoming semantic state?
5. What scope must be explicitly deferred so this bead does not falsely claim
   general Go structural typing, generics, type assertions, or full go2hx-like
   binding generation?

## Required response shape

Start with one verdict: `APPROVE A`, `APPROVE B`, `APPROVE C`,
`APPROVE HYBRID`, or `REJECT ALL / NEED MORE EVIDENCE`.

Then provide:

1. exact model route and reasoning setting;
2. the selected ownership/data-flow design;
3. invariants and fail-closed rules;
4. the first red-to-green implementation sequence;
5. findings with severity, file evidence, violated invariant, and regression
   strategy;
6. explicitly deferred scope;
7. a final `PROCEED`, `PROCEED WITH CONDITIONS`, or `DO NOT PROCEED` decision.

Keep Oracle output advisory. The repository contracts, local source tracing,
tests, and recorded Beads disposition remain implementation authority.
