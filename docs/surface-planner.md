# Portable Surface Planner

## Practical outcome

The compiler now makes one shared, typed decision about each recognized
portable Haxe surface before it lowers code to Go. The optimizer, import
collector, and selective `hxrt` runtime packager consume that same immutable
plan.

This closes an important authority gap: a profile name, optimization define, or
generated-text scan cannot independently decide that ordinary portable Haxe
code may use a native carrier. A carrier is considered only after the
[surface contract registry](surface-contract-registry.md) admits the exact
typed shape.

The compatibility presets still work, but they do not change these semantic
decisions. `portable` with eager native specialization and `metal` with proven
specialization produce the same portable-surface plan.

## The simple model

Think of the registry as the safety certificate and the planner as the routing
table:

```text
typed use in the program
    -> registry: is this exact shape proved safe?
    -> planner: which implemented carrier or fallback applies?
    -> optimizer, imports, and runtime packaging consume one answer
```

For example, the registry can admit an `Array<Int>` shape because its semantic
contract is proved. The planner still selects `hxrt_array` today because the
separate default-promotion work has not activated the new Array carrier. This
is intentional: proof that a representation is eligible and permission to make
it the default are different release decisions.

## Current selections

| Surface | Admitted carrier | Planner behavior now |
| --- | --- | --- |
| `String` | `go_string` | Routes the existing proved carrier and gates String fast paths on registry admission. |
| `haxe.io.Bytes` | `go_byte_slice` | Routes the existing proved shared data/view carrier. |
| structural `Iterator<T>` | future statically typed `go_iterator` | Keeps the established `map[string]any` closure carrier as `hxrt_iterator`; `.7.7` must introduce a distinct typed Go carrier before activation. |
| `Array<T>` | `go_slice` | Keeps `hxrt_array` until the dedicated promotion and rollback gates land. |
| `StringMap<V>` / `IntMap<V>` | `go_map` | Keep their registered map fallbacks until dedicated carrier promotion. |
| portable `Option<T>` / `Result<T,E>` | `go_option` / `go_result` | Keep portable enum emission until dedicated carrier promotion. |
| `Function` / `ObjectMap` | no admitted contract | Preserve existing lowering; the planner cannot invent a carrier or fallback contract. |

Rejected registered shapes always select the fallback recorded by their
contract. A recognized surface without a registered fallback selects
`existing`, which means the compiler preserves the established lowering
instead of guessing.

Explicit target-native APIs such as `go.Slice`, `go.Map`, and `go.Result` are
outside this registry. Their source API is already an explicit Go-native
boundary, so they continue to follow their own typed contracts.

## One authority, three consumers

`CompilationContext.surfacePlan` is assigned once from the resolved
`GoBuildContext` and the immutable registry snapshot. It is not recomputed
during lowering.

The consumers use it as follows:

1. The optimizer checks the selected carrier before applying a governed
   portable specialization. String fast paths therefore cannot bypass a
   rejected String decision.
2. The import collector adds only the imports selected by the plan. Its normal
   structural import filter still removes aliases that the generated module
   does not actually reference.
3. The runtime planner adds the selected fallback or carrier requirements to
   selective `hxrt` packaging.

An exact typed Go AST walk also replaces the old final-output scan that looked
for Array support text in generated Go. It recognizes only the closed
`hxrt.Array`, `hxrt.NewArray`, and `hxrt.ArrayFromValues` symbols after lowering
and before printing. This keeps Array support for staged stdlib and explicit
`go.Slice` adapters without copying `array.go` for a specialized source shape,
such as `Rest<T>`, that does not materialize the portable carrier.

## Reading the reports

`optimizer_plan.json` schema v7 and `hxrt_plan.json` schema v4 contain the same
surface-plan authority:

- `surfacePlanAuthority`: identifies the typed build-context plus registry
  authority;
- `surfacePlanDecisionCount`: number of per-use decisions;
- `requiredSurfaceImports`: deduplicated selected import consequences;
- `requiredSurfaceRuntimeFeatures`: deduplicated selected `hxrt` consequences;
- `surfacePlans`: the complete decisions.

Schema v4 also groups the exact copied files and inclusion reasons under each
typed runtime capability. The runtime copier and the report consume the same
immutable manifest, so the report is evidence of the generated file set rather
than a second prediction.

Each `surfacePlans` entry records:

- `usedType`, module, location, and usage level;
- the surface ID and contract version;
- eligibility outcome, stable reason, and explanatory detail;
- `native`, `fallback`, or `existing` selection and its reason;
- selected representation and fallback explanation;
- selected imports and runtime requirements;
- the registry-owned no-`hxrt` status and whether the representation actually
  selected for this use is no-`hxrt` eligible.

The report is deliberately explicit. It lets CI and reviewers answer both
“was this shape proved safe?” and “what did this compiler version actually
choose?” without inferring either answer from generated text.

## Define precedence

Native policy defines may affect explicit native APIs, but they do not overrule
portable semantic proof. The planner records the resolved specialization value
and source for provenance, then derives every portable selection only from the
registry decision and the compiler's closed carrier-activation table.

The paired surface-registry fixture compiles:

- portable preset plus an explicit eager specialization policy;
- metal preset plus an explicit proven specialization policy.

It requires byte-equivalent surface selections, imports, and runtime
requirements. This proves both profile independence and explicit-define
precedence without making either policy an admission authority.

## Promotion boundary

Adding a representation to the activation table is a default behavior change,
not a documentation-only update. The follow-up promotion work must provide the
generated carrier or exact per-shape lowering gate, semantics and runtime
evidence, paired rollback coverage, and performance evidence before switching
Iterator, Array, maps, Option, or Result away from their current fallbacks.

For Iterator specifically, the existing structural value is a
`map[string]any` containing `hasNext` and `next` closures. A concrete Go result
type on one stored closure does not make that erased map the future
`go_iterator`. `.7.7` must define and prove a distinct statically typed carrier;
until then both admitted and rejected Iterator shapes remain auditable planner
fallbacks with the registered `core` consequence.

Runtime slicing consumes this plan directly. It does not reconstruct surface
choices from profile names, defines, generated text, or empty import lists.

## Evidence

Run:

```bash
npm run test:surface-contract-registry
npm run test:auto-planner:schema
npm run test:compiler-debt
```

The focused contract exercises real Haxe typing, admitted and rejected shapes,
fallback consequences, profile/define independence, report schemas, source
authority checks, and the typed Go AST runtime-reachability replacement.
