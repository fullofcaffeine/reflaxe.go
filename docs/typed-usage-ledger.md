# Typed Usage Ledger

## What it is

The typed usage ledger is the compiler's deterministic record of which Haxe
types and members survived typing and dead-code elimination. Collection is
always enabled because later compiler planning needs the typed snapshot. To
write the optional human/machine-readable report, compile with:

```text
-D reflaxe_go_type_usage_report
```

The compiler then writes `type_usage.json` and `type_usage.md` beside the
generated Go files.

The report answers four practical questions:

1. Which project declarations retained typed usage after dead-code elimination?
2. Which generic/function/anonymous shapes, fields, constructors, and calls did
   those declarations use?
3. Which typed Go package imports did that usage select?
4. Which `hxrt` capabilities did the completed lowering require?

It is analysis evidence, not a new whole-program intermediate representation.
The normal pipeline remains:

```text
Haxe typed AST
    -> Reflaxe typed usage map
    -> haxe.go typed usage snapshot
    -> Go lowering and runtime-capability inference
    -> optional deterministic report
```

## Why it exists

Source-text scans cannot reliably decide runtime or representation policy.
Imports can be aliased, types can arrive through generic signatures, dead code
can disappear, and the compiler can introduce typed support that is not visible
as a source import.

Reflaxe already exposes a `TypeUsageTracker` after Haxe has typed the program.
Before this integration, haxe.go left `trackUsedTypes` disabled and discarded
that evidence. The ledger now gives later registry and planner work one typed,
auditable input without serializing arbitrary Haxe macro objects or using
`Dynamic` payloads.

## How it works

### Collection

`GoReflaxeCompiler.setupModule` builds Reflaxe's `TypeUsageMap` for each
project-owned class, enum, typedef, or abstract, at the point where Reflaxe's
own `trackUsedTypes` option would build it. `CompilerInit` leaves that option
off because it builds a map for every module type, including standard-library
and compiler types whose maps the ledger never reads. During each project
class, enum, typedef, or abstract callback, `GoTypeUsageLedger` consumes the
current `TypeUsageMap`.

The map comes from `GoTypeUsageTracker`. It follows the rules of Reflaxe's
`TypeUsageTracker` and returns the same map, but it reads each enum, typedef,
and abstract declaration only once per compilation instead of once for every
expression that names it. Reflaxe needs a declaration's id for every type it
records, and reading an abstract rebuilds its whole declaration, so this was
most of the tracker's time in large programs. Classes are still read every
time, because type parameter classes print alike: every local `pick<V>` owns a
`pick.V`. `GoTypeUsageTrackerContract` compares the two trackers' maps on
fixture and standard-library modules, using the unchanged vendored Reflaxe
tracker as the reference.

The ledger walks every field body of a project class too. The tracker decodes
each body once and the compiler hands the ledger those fields and bodies
(`GoTypeUsageTracker.classFields`), so a body is not built twice. Between the
tracker's read in `setupModule` and the class callback, Reflaxe only asks
which classes and fields to generate, so the ledger sees the same values it
would read itself. `GoTypeUsageTrackerContract` also checks that a report
built from the shared fields equals one whose ledger decodes them itself.

The ledger records the seven Reflaxe usage levels with stable names:

- `expression`
- `variable_type`
- `static_access`
- `constructed`
- `function_declaration`
- `variable_declaration`
- `extended_from`

It also walks the same typed module expressions for closed shape and operation
facts. `GoTypeShape` preserves nominal type parameters, function parameters and
returns, sorted anonymous fields, dynamic inner types, and deterministic
recursion markers. These are typed AST observations; the collector does not
reread source text or retain macro objects.

A large program names the same declarations many thousands of times, so the
collector avoids reading a declaration more often than it must. It reads a
class once for both its path and its `@:go.import` metadata, keeps the path and
import of each enum, typedef, and abstract by type path, names a field by
printing its reference instead of decoding it, and builds the JSON of each
shape without children (`Int`, `String`, a type parameter) once. Classes are
not kept, for the same reason as in the tracker. `GoTypeUsageLedgerContract`
and `GoTypeShapeContract` compare these answers with the plain reads.

### Locations and paths

Locations use `Module:line`, such as `Main:6`. Reports therefore remain useful
without exposing an absolute checkout path. Modules, usage entries, native
imports, and capability reasons are deduplicated and sorted before rendering.

Only project-owned declarations with retained usage appear as owners. Two
declarations in one `.hx` module retain separate `Module.Type` identities.
Standard-library, `go.*`, `hxrt`, and compiler types can still appear as typed
targets. This keeps the report focused while preserving the dependencies that
later planners need.

### Generics and dead code

The ledger reports Reflaxe's actual typed facts. A used generic surface such as
`UsedBox<Int>` is one nominal shape whose parameter is `Int`; the relationship
is not flattened into unrelated facts. Function and anonymous shapes likewise
retain their ordered arguments/return and sorted fields. The compiler converts
these observations into a small closed algebra instead of serializing macro
`Type` values.

Nested uses such as `UsedBox<UsedBox<Int>>` remain nested. A repeated nominal
type is not itself recursion because the ledger never expands its declaration;
recursion guards apply only while structurally expanding function, anonymous,
dynamic, lazy, or unresolved types.

An otherwise reachable generic referenced only by a compile-time-false branch
has no retained evidence after Haxe dead-code elimination. The
`core/type_usage_ledger` fixture proves this with `UsedBox<Int>` and
`DeadBox<String>`, two declarations from one source module, function and
anonymous shapes, a typed `go.Fmt` import, and an `hxrt` metadata token.

Native imports expose both `metadataImportPath` and `resolvedImportPath`.
Ordinary Go paths such as `fmt` are identical in both fields. The special
`hxrt` metadata token resolves to the configured module path such as
`snapshot/hxrt`, matching generated Go.

### Planner boundary

One deeply read-only snapshot is stored as a final `CompilationContext` field
before Go lowering. Its lists clone their source arrays and expose no mutable
alias. The optional post-lowering report uses a separate snapshot to append
runtime capability consequences; it never mutates planner authority. The
[Go surface contract registry](surface-contract-registry.md) now consumes that
context snapshot and publishes a second immutable admission snapshot. The
[portable surface planner](surface-planner.md) combines those decisions with
the resolved build context once, then supplies the same selected carrier,
imports, and runtime requirements to every downstream consumer.

The existing portable-native source scanner remains available only for
transitional contract diagnostics through
`reflaxe_go_portable_native_scan_mode=scanner|hybrid`. The ledger records this
as `transitional_contract_diagnostics_only`; source scanning is not planning
authority.

## Sibling compiler comparison

The integration follows established Reflaxe practice while making the evidence
more explicit:

| Compiler | Current use of Reflaxe typed usage |
| --- | --- |
| `haxe.rust` | Enables `trackUsedTypes`; `TypeUsageAnalyzer` reduces the map to normalized module paths used by runtime-feature selection. |
| `haxe.ruby` | Enables `trackUsedTypes`; `RequireRegistry` consumes type metadata to select deterministic Ruby `require` entries. |
| `haxe.elixir.codex` | Does not currently enable or consume Reflaxe `TypeUsageTracker`; framework discovery uses its own typed/macro mechanisms. |
| `haxe.go` | Runs the tracker for project-owned module types; preserves types, member/call locations, native imports, and runtime consequences; admits exact portable shapes through a contract registry; and feeds one immutable surface plan to optimizer, import, and runtime consumers. |

The useful common pattern is compiler-observed typed reachability. The exact
consumer differs by target: Rust selects runtime features, Ruby selects
requires, and Go needs both registry admission and selective `hxrt` planning.

## Evidence

Run:

```bash
npm run test:type-usage-ledger
python3 test/run-snapshots.py --case core/type_usage_ledger --runtime
npm run test:auto-planner:schema
```

`npm run test:type-usage-ledger` also runs `GoTypeUsageTrackerContract`,
`GoTypeShapeContract`, and `GoTypeUsageLedgerContract` (see Collection above).

The contract checks deterministic bytes across repeated compilations, deep
read-only publication, preserved generic/function/anonymous relationships,
DCE-eliminated generic evidence, same-module owner identity, metadata and
resolved native imports, member/call evidence, capability counts, and absence
of absolute paths.
