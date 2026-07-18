# Generated Class Dispatch ABI

## What it is

Generated Haxe classes use two related Go representations:

- a concrete struct, which embeds the concrete struct for its generated
  superclass; and
- a class-specific Go interface stored in `__hx_this`, which identifies the
  most-derived receiver for Haxe virtual method calls.

The embedded structs preserve Haxe's nominal upcast path. The dispatch carrier
preserves Haxe overriding after that physical upcast. This is a compiler ABI
rule, not an API that application code should access.

## Why it exists

Go embedding alone does not reproduce Haxe virtual dispatch. Given `Base <-
Middle <- Leaf`, a Haxe value typed as `Base` is represented by the embedded
`Base` pointer inside `Leaf.Middle`. A call to a Haxe virtual method therefore
uses that carrier's `__hx_this` interface to recover `Leaf` before selecting the
method.

Every generated ancestor has its own carrier. Rebinding only the direct
superclass would leave `Leaf.Middle.Base.__hx_this` pointing to `Middle`, losing
a `Leaf` override through a `Base` value and from inside an inherited base
method.

## How construction binds the receiver

After the direct superclass constructor returns, a generated subclass
constructor rewires every generated ancestor from deepest to nearest, then
wires its own carrier. For `Base <- Middle <- Leaf`, the relevant shape is:

```go
self.Middle = New_Middle()
self.Middle.Base.__hx_this = self
self.Middle.__hx_this = self
self.__hx_this = self
```

These are ordinary typed `GoExpr.GoSelector` and `GoStmt.GoAssign` nodes. The
compiler does not use `GoRaw`, reflection, `unsafe`, or profile-specific
branches for this invariant.

The complete constructor order is:

1. allocate the concrete receiver;
2. run the direct superclass constructor;
3. rebind generated ancestor carriers, deepest first;
4. bind the concrete receiver's own carrier;
5. initialize Haxe `dynamic function` instance fields; and
6. run the subclass constructor body.

The deepest-first order is deterministic. A one-level hierarchy retains its
existing shape: direct superclass carrier, then concrete receiver carrier.

## Ownership boundary

The traversal follows only `projectSuperClass`, the same path that owns emitted
Go struct embedding. Native and extern ancestors do not enter that generated
path. A compiler-owned ancestor also terminates carrier rebinding because its
synthetic layout is authoritative and must not be assumed to contain
`__hx_this`.

This rule is profile-neutral. `portable` and the compatibility `metal` preset
share the same generated class representation; explicit Go-native source
boundaries do not change ordinary generated Haxe class semantics.

## Dynamic generated-method lookup

`Reflect.field` and `Reflect.hasField` sometimes receive a class instance only
as `Dynamic`. Go's `reflect` package cannot discover the lowercase methods that
haxe.go deliberately emits for ordinary Haxe methods. Exporting or duplicating
those methods would change the public Go shape, while moving discovery into
`hxrt` would cross a Go package-visibility boundary.

When either lookup API is reachable, the compiler therefore emits one selective,
method-only metadata capability in the generated program package:

1. a type switch maps the physical embedded carrier to its canonical
   `__hx_this` receiver;
2. a second type switch selects the resolver for that exact concrete receiver;
3. each class resolver lists only its own emitted methods by Haxe field name;
4. a miss follows the one direct generated-superclass embedding, with a nil
   guard; and
5. the result is an already-bound Go function value or `nil`.

The selector stored for each entry is the selector chosen by ordinary method
lowering. For example, Haxe method `type` remains the lookup key `"type"` even
when its Go selector is normalized to `type_`. Constructors, static functions,
interfaces without bodies, mutable Haxe `dynamic function` fields, and synthetic
Go adapters are not method-metadata entries.

The central helper and class resolvers use typed Go AST only. They are absent
when `Reflect.field` and `Reflect.hasField` are unreachable. They do not own
field/property lookup, method invocation, iteration, errors, or retention policy;
current `Reflect` checks them after native/data fields and before exported native
Go methods. `hxrt.TemplateCall` invokes only an already-resolved function value.

## Deliberate limits

This invariant applies to normally constructed generated classes after each
direct superclass constructor has returned. It does not claim that a virtual
call made *during* a superclass constructor can already reach a not-yet-bound
leaf receiver.

Constructor-free allocation and decoding are also outside this constructor
rule. Any runtime API that creates objects without normal constructors must
independently establish the carrier invariants promised by that API.

## Contract evidence

- `test/semantic_diff/deep_inheritance_dispatch_rebinding` compares a virtual
  call from the leaf constructor body, leaf/middle/base-typed calls, and a bound
  inherited base method against Haxe `--interp`.
- `test/snapshot/core/deep_inheritance_dispatch_rebinding` fixes the generated
  constructor order—including rebinding before the leaf body—and runs the
  generated Go behavior.
- `test/snapshot/core/inheritance_self_dispatch_wiring` protects the unchanged
  one-level shape.
- `test/snapshot/stdlib/haxe_io_misc_direct` proves the same rule on the staged
  `Input <- BytesInput <- StringInput` hierarchy.
- `test/semantic_diff/haxe_template_concrete_iterable_contract` proves computed
  method lookup, deep canonical receiver recovery, inherited/private/normalized
  selectors, bound arguments and mutation, interface values, nil/misses, and
  concrete `haxe.Template` iteration against Haxe `--interp`.
- `test/snapshot/stdlib/haxe_template_generated_method_lookup` fixes the typed
  helper shape, superclass fallback and guards, Reflect ordering, unchanged
  lowercase methods, and the source-owned Template loop.
