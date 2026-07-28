# Go Surface Contract Registry

## What it is

The Go surface contract registry is the compiler's proof-backed allowlist for
using a native Go representation behind portable Haxe source.

In simple terms: seeing `Array<Int>` is not permission to emit a Go slice. The
compiler also needs a reviewed contract that says which Haxe behavior must stay
the same, which exact type shapes are safe, what Go representation may be used,
what happens when the shape is unsafe, and which tests prove those claims.

Production admits eight proof-backed surfaces: portable Haxe `Array`, `String`,
`haxe.io.Bytes`, `StringMap`, `IntMap`, and structural `Iterator<T>`, plus the
`reflaxe.std.Option` and `reflaxe.std.Result` facades. Their exact contracts are
documented in
[Portable Collection Representation Contract](portable-collection-contract.md),
[Portable String, Bytes, and Iterator Contract](portable-string-bytes-iterator-contract.md),
and [Portable Option and Result Contract](portable-option-result-contract.md).
ObjectMap and functions remain absent until their identity evidence lands.

Enable the optional report with:

```text
-D reflaxe_go_surface_contract_report
```

The compiler writes:

- `surface_contracts.json`, validated by
  [`surface-contracts-v2.schema.json`](schemas/surface-contracts-v2.schema.json);
- `surface_contracts.md`, the same decisions in a review-friendly form.

Reports now use schema version 2 and registry version 1. Schema v2 adds the
`haxe.Iterator` surface ID and the exact anonymous-field pattern needed to
describe its structural protocol. The published
[`surface-contracts-v1.schema.json`](schemas/surface-contracts-v1.schema.json)
remains unchanged and read-only so tools that validate historical v1 reports
keep working. A report schema version describes the JSON vocabulary; the
registry version independently describes the production contract catalog.

## Why it exists

A generated Go type can look reasonable and still break Haxe behavior. Examples
include two variables no longer seeing the same mutation, `null` and empty
values becoming indistinguishable, map keys changing meaning, Unicode indexing
changing, or a callback capturing a different value.

Profiles cannot admit a surface. `portable` and `metal` are policy presets; they
do not prove that a representation preserves semantics. Source-text import
scans also cannot prove it because aliases, generics, dead-code elimination,
and compiler-introduced support are typed facts.

The registry combines two independent requirements:

1. the [typed usage ledger](typed-usage-ledger.md) proves the exact shape that
   survived typing and dead-code elimination;
2. a versioned registry entry proves that this shape has an explicit source
   contract, eligibility rule, native carrier, fallback, runtime/import
   consequences, and semantic fixture.

If either half is missing, admission fails.

## How it works

The current flow is:

```text
typed Haxe program
    -> immutable typed usage ledger
    -> recognize an exact known surface root
    -> look up its validated versioned contract
    -> match the complete nominal/function/anonymous shape
    -> apply closed eligibility rules
    -> admit or reject with a stable reason
    -> publish one immutable snapshot and optional report
```

The four decision reasons are:

- `contract_admitted`: a validated contract matched the complete shape;
- `contract_missing`: the compiler knows the surface, but production has not
  admitted it;
- `shape_mismatch`: a contract exists, but its recursive type pattern did not
  match;
- `eligibility_rejected`: the shape matched, but a closed rule such as
  "contains no `Dynamic`" or "bound map key is Go-comparable" failed.

Unknown types are ignored instead of being turned into accidental candidates.
Known nominal candidates use exact roots such as Haxe `Array`, `String`,
`haxe.io.Bytes`, and `reflaxe.std.Option`/`Result`; a namespace prefix is never
sufficient. Structural Iterator candidates must contain exactly non-optional
zero-argument `hasNext():Bool` and `next():T` fields. Other anonymous objects
are ignored. The compiler resolves only Haxe's canonical
`StdTypes.Iterator<T>` typedef to that structural evidence; a user-owned
Iterator typedef remains nominal and opaque, so it cannot widen admission by
aliasing the same underlying type. Bare generic enum declaration references do not choose a
representation: the applied typed usage carries the parameters needed for that
decision.

### What every admitted contract records

Each contract is deeply read-only and includes:

- stable surface, contract, and source-semantics versions;
- portable Haxe or shared-family source authority;
- a recursive type pattern with named generic bindings;
- typed eligibility rules;
- the native Go representation;
- native Go imports and `hxrt` requirements;
- the fallback representation, policy, imports, and runtime requirements;
- no-`hxrt` eligibility status;
- stable proof IDs and repository-relative fixture paths;
- target-local or shared-family synchronization expectations.

Validation rejects duplicate surface IDs, unknown vocabulary values, invalid
versions, malformed recursive patterns or nested lists, a shape rooted at the
wrong source surface, rules that reference an unknown binding, absolute or
traversing proof paths, unknown runtime features, and contracts without a
semantic-diff proof. Binding names use a portable identifier grammar so they
cannot corrupt encoded eligibility rules. A null catalog is malformed rather
than an alias for an explicitly supplied zero-entry catalog. Invalid catalog
data cannot become compiler authority; `create()` reports typed validation
issues instead of exposing a raw null or enum-match failure.

Repository CI validates both populated and empty reports against the published
JSON Schema with Ajv, checks compiler/schema vocabulary synchronization, and
checks fixture existence. The packaged compiler validates stable relative
references without depending on a developer's checkout path.

### Current authority boundary

`CompilationContext.surfaceContractRegistry` is assigned once, before Go
lowering, beside `typedUsageLedger`. Neither report generation nor a
compatibility profile can mutate it.

The registry is observational for representation selection today. Collections,
String, Bytes, Iterator, Option, and Result are admitted, but the planner does
not consume their carrier decisions until `.7.6`; current generated
declarations retain their portable fallbacks. `go_slice` means a shared,
slice-backed carrier rather than a naked Go slice, while portable `go_map`
admission currently covers only the fixed StringMap and IntMap key contracts.
`go_string`, `go_byte_slice`, and `go_iterator` likewise name the reviewed
semantic carriers, not raw values or a Go `range` rewrite.
In other words, the planner does not consume registry admission in this task.

Function shapes remain known but report `contract_missing`.
`haxe_go-vfp.7.11` owns the bound-method identity gap that must close before a
portable callable carrier can be admitted. `.7.6` makes optimizer and runtime
planners consume the proven decisions; it must not infer closure admission.

Portable `reflaxe.std.Result<T,E>` is not native `go.Result<T>` and is not a Go
`(T,error)` pair. The registry's `go_result` value names a future typed
two-parameter carrier that preserves `E`; it never authorizes implicit Go-error
conversion.

## Adding a surface safely

Use this order:

1. Define the portable source behavior, including identity, mutation,
   null/empty, iteration, Unicode, callback, or error rules that apply.
2. Add a semantic-diff fixture before the contract.
3. Add target-only generated-shape, Go runtime, race, or performance evidence
   where relevant.
4. Add the exact known surface ID and root if the vocabulary does not already
   contain it.
5. Add a versioned contract with the smallest eligible type pattern.
6. Name both native and fallback consequences.
7. Run the registry contract, semantic diff, snapshots, compiler-debt gate,
   examples, and relevant performance/race suites.
8. Get the required second-pass review before changing a default
   representation.

A broader pattern requires broader proof. One ASCII string fixture cannot admit
all Haxe strings, and one simple `Array<Int>` fixture cannot prove nested
mutation or callback cases.

## Sibling compiler comparison

The useful family precedent is governance, not identical implementation:

| Compiler | Current pattern | What haxe.go learns |
| --- | --- | --- |
| `haxe.rust` | Has a `SurfaceContractRegistry` for portable Option/Result and deterministic contract reports. Its first registry matches consumed module paths and records Rust representation/fallback facts. | Reuse stable IDs, versioned source semantics, explicit representation/fallback, and reports. Strengthen admission with recursive typed shapes, mandatory validation/proofs, deep immutability, and Go-specific nil/comparability/aliasing/error rules. Do not copy Rust ownership, borrow, `Send`/`Sync`, Cargo, or `no_hxrt` mechanics. |
| `haxe.ruby` | Uses Reflaxe typed usage to select deterministic Ruby `require` entries, but does not have an equivalent native-representation proof registry. | Typed reachability is useful; Ruby require selection is not semantic admission precedent. |
| `haxe.elixir.codex` | Uses target-owned typed/macro and staged-standard-library governance, without a comparable `SurfaceContractRegistry`. | Keep source-library ownership and explicit target boundaries, but do not infer Go representation policy from BEAM lowering. |
| Genes | Uses target-specific mapping and compatibility governance, without a comparable proof registry for native storage shapes. | Share documentation and compatibility discipline, not JavaScript representation decisions. |

Family synchronization is represented in each contract but enforced by the
separate `.7.8` task. Shared source contract IDs and versions may synchronize;
Go representations, imports, fallbacks, and proof details remain target-owned.

## Evidence

Run:

```bash
npm run test:surface-contract-registry
npm run test:type-usage-ledger
npm run test:auto-planner:schema
npm run test:compiler-debt
```

The focused contract runs the real Haxe validator against valid, duplicate,
unknown, unproven, malformed-pattern, malformed-nested-list, unsafe-path, and
unknown-runtime entries. It proves Option/Result admission and fallback,
typed Array/StringMap/IntMap admission and fallback, nullable String and shared
Bytes carriers, exact typed Iterator admission with Dynamic/opaque fallback,
continued function and ObjectMap identity rejection, fixed-key comparability,
explicit `go.Map` exclusion, unregistered-surface rejection, deep catalog
copying, deterministic JSON/Markdown, actual JSON Schema conformance,
compiler/schema vocabulary synchronization, generated fallback declarations,
path hygiene, and byte-identical portable/metal reports.
