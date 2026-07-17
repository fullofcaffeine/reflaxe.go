# Typed Go IR Contract

## What it is

The typed Go intermediate representation (IR) is the compiler-owned syntax
model between Haxe lowering and generated Go text. It now represents target
types, package names, import paths, and unary/binary operators structurally:

```text
Haxe TypedExpr
    -> Go builders and lowering
    -> typed Go AST
    -> transform passes
    -> structural import-use analysis
    -> Go printer
    -> gofmt / go test
```

The central types are:

- `GoType` and `GoBuiltinType` for Go target types;
- `GoPackageName` and `GoImportPath` for package/import identity;
- `GoBinaryOperator` and `GoUnaryOperator` for expression operators; and
- `GoAST` nodes whose relevant fields use those types instead of `String`,
  including typed slice allocation through `GoMakeSlice`.

This is an internal compiler contract. It models printable Go syntax; it is not
a replacement for Haxe's type system and does not decide whether a Haxe value
is portable or admitted at a native boundary.

## Why it exists

A rendered string such as `map[string][]*pkg.Value` is readable, but compiler
passes cannot safely answer basic questions about it:

- Is it a map, a slice, or a named type?
- Is its map key legal?
- Which package qualifier does it use?
- Is a variadic parameter in the final position?
- Does a function return one value or several?
- Is `=>` an accidentally emitted operator?

Previously those questions were deferred to generated-source compilation or
answered by regular expressions over target text. That made validation late,
imports fragile, and transforms harder to make exhaustive.

The typed boundary makes those facts traversable before printing. It also
supports the compiler-debt direction in
[`compiler-debt-ratchet.md`](compiler-debt-ratchet.md): raw Go fragments remain
measured migration debt, while ordinary type and operator syntax no longer
needs to be an analysis blind spot.

## How it works

### Type shapes

`GoType` has no unchecked raw-type constructor. It covers the target forms
required by the compiler:

| Go form | Structural builder |
| --- | --- |
| `int`, `string`, `error`, `any`, and other builtins | `GoType.builtin(...)` |
| `Local` | `GoType.named("Local")` |
| `atomic.Pointer` | `GoType.qualified(packageName, "Pointer")` |
| `*T` | `GoType.pointer(type)` |
| `[]T` | `GoType.slice(type)` |
| `[3]T` | `GoType.array(3, type)` |
| `map[K]V` | `GoType.map(key, value)` |
| `chan T`, `<-chan T`, `chan<- T` | `GoType.channel(direction, type)` |
| `func(T, ...U) (V, error)` | `GoType.functionType(params, results)` |
| `interface{ Apply(T) (V, error) }` | `GoType.interfaceType(methods)` |
| `atomic.Pointer[int]` | `GoType.generic(base, arguments)` |
| `(V, error)` inside a function signature | `GoType.multiResult(results)` |

`multiResult` represents the parenthesized result portion of a function type;
it is not a standalone Haxe value type. `variadic` represents parameter shape
and is admitted only as the final parameter by `functionType`.

### Validation before printing

Construction rejects malformed or structurally invalid input, including:

- invalid normalized package/type identifiers and Go keywords;
- empty, quoted, whitespace-containing, absolute, or traversal-like imports;
- negative array lengths;
- slices, maps, and functions used as map keys;
- non-final variadic function parameters;
- multi-result values nested where one ordinary type is required;
- duplicate interface method names;
- generic instantiation over a non-named base; and
- tokens outside the closed unary and binary operator sets.

These checks are target-syntax checks. Representation eligibility and portable
semantics remain separate compiler policies.

### Imports are structural where syntax is structural

Candidate imports still enter the compiler as module/runtime paths. They are
validated into `GoImportPath` before AST transforms. Import filtering then
derives the package qualifier and asks each `GoType` tree whether it uses that
qualifier.

For example, `atomic.Pointer[int]` retains `sync/atomic` because the generic
base contains the structural qualifier `atomic`; no regular expression over
the rendered type is needed. Expressions and statements that still use
`GoRaw` retain their measured text-scanning fallback until their own migration
beads replace those fragments with structural nodes.

### The printer owns punctuation

AST nodes carry meaning, not assembled target syntax. `GoASTPrinter` owns:

- package and import quoting;
- pointer, slice, map, channel, function, interface, and generic punctuation;
- result-list parentheses;
- operator tokens;
- slice-allocation syntax such as `make([]T, length, capacity)`; and
- ordinary declaration/expression layout.

Generated files continue through `gofmt` and `go test`, so adopting structural
nodes does not create a second formatting contract.

## Migration boundary

The compiler has many existing lowering helpers that still compute a target
type spelling while they also perform Haxe-semantic decisions. Rewriting all
of those decisions in one change would combine syntax migration with
representation-policy changes.

During the incremental migration, `GoType` therefore provides one validating
`@:from String` parser:

1. existing lowering computes its current spelling;
2. the value crosses a typed `GoAST` field;
3. the parser produces structural nodes or fails immediately; and
4. transforms and import analysis see only `GoType` after that boundary.

This is a compatibility bridge, not a raw escape hatch. New AST construction
should use direct builders. Existing string-producing helpers should migrate
when their ownership area is touched, without changing portable/native
admission semantics in the same mechanical step.

The first direct migrations include the operator mapper, function-result AST
construction, and the complete typed-shape fixture. Remaining expression and
statement structure is tracked by the dependent typed-IR beads rather than
being hidden inside this contract.

Portable `Array.remove` and `Array.insert` are also lowered through ordinary
typed statement/expression nodes: range, condition, assignment, slice, call,
and return. Typed comparable elements stay on native Go equality, strings use
their established value comparator, and only erased, nullable-primitive, or
non-comparable element shapes cross the narrow `hxrt.HaxeEqual` runtime
boundary. No library name or raw statement block owns these algorithms.

Collection representation remains a separate semantic decision. Root Haxe
`Array<T>` lowers to the shared `*hxrt.Array` carrier because a copied Go slice
header cannot preserve length-changing aliases or sparse null-filled growth.
Explicit `go.NativeSlice<T>`, `haxe.Rest`, fixed `Vector` storage, and the
temporary `BytesData` representation retain raw Go slice shapes where their
contracts require them. `NativeSlice.fromArray` and `toArray` are explicit
shallow-copy boundaries; typed externs must not declare native `[]T` values as
portable `Array<T>` merely because both can be indexed.

## Profile and native-boundary relationship

The typed Go IR is profile-neutral:

- `portable` remains the default semantic product path;
- `metal` remains a compatibility convenience preset; and
- explicit Go-native authority remains module/API scoped.

A `GoType.channel(...)` node says how a Go channel type prints. It does not, by
itself, grant source code permission to use Go-native channel semantics. The
profile/boundary analyzers make that decision before or during lowering, and
both admitted portable optimizations and explicit native APIs share the same
typed target syntax afterward.

## Sibling-target precedent

This direction matches the strongest parts of the sibling compiler designs,
while remaining target-specific:

- `haxe.rust` models Rust types, paths, generics, lifetimes, and const arguments
  structurally. Its type/path work is the closest precedent for keeping target
  syntax separate from Haxe type semantics.
- `haxe.elixir` uses closed `EBinaryOp` and `EUnaryOp` enums inside its typed
  AST, demonstrating the transform and exhaustiveness benefit of typed
  operators even on a dynamic target.
- `haxe.ruby` structurally distinguishes method-parameter and call-argument
  kinds, but currently keeps ordinary unary/binary operators as strings. That
  is a useful partial precedent, not a reason for Go to keep strings where
  static types and import traversal require more structure.

The shared family principle is structural syntax where compiler passes need to
reason about meaning. The concrete nodes differ because Go, Rust, Elixir, and
Ruby have different target grammars.

## Extending the IR

When adding a target type or operator form:

1. add a failing positive or negative AST contract first;
2. extend the closed structural model and its validation;
3. update rendering and structural import traversal exhaustively;
4. do not add an unchecked raw-type/operator variant;
5. preserve the distinction between Go syntax and Haxe semantic admission; and
6. run the targeted snapshot, the full snapshot suite, and the compiler-debt
   ratchet.

The primary contracts are:

- `core/ast_typed_type_operator_printer` for valid shape and printer coverage;
- `negative/ast_invalid_go_type` for malformed type syntax;
- `negative/ast_invalid_go_type_combination` for structural invalidity;
- `negative/ast_invalid_go_operator` for the closed operator set;
- `negative/ast_invalid_go_import_path` for whitespace/import validation; and
- `negative/ast_invalid_go_import_path_character` for Go-tooling portability.
