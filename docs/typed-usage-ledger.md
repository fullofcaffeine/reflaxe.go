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

`CompilerInit` enables Reflaxe's `trackUsedTypes` option. During each project
class, enum, typedef, or abstract callback, `GoTypeUsageLedger` consumes the
current `TypeUsageMap`.

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
contract registry and runtime planner can consume the context snapshot in their
dedicated follow-up Beads without this change silently altering representation
admission or runtime selection.

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
| `haxe.go` | Enables the tracker and preserves types, member/call locations, native imports, and runtime consequences in a typed snapshot and optional report. Planner consumption remains an explicit later step. |

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

The contract checks deterministic bytes across repeated compilations, deep
read-only publication, preserved generic/function/anonymous relationships,
DCE-eliminated generic evidence, same-module owner identity, metadata and
resolved native imports, member/call evidence, capability counts, and absence
of absolute paths.
