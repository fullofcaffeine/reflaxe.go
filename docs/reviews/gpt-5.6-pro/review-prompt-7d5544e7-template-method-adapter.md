# Independent pre-refactor review: generated method access for `haxe.Template`

Review Haxe.Go source commit
`7d5544e738d72f2e42f88a467b43514ce75f3d6b` at maximum reasoning depth. Treat
that pushed Git commit as the source authority. This is a pre-refactor design
checkpoint for Bead `haxe_go-vfp.8.7.19`; no implementation has been selected
yet, and an uncommitted working tree must not be treated as evidence.

## Provenance requirement

This prompt is intended for a genuine GPT-5.6 Pro independent reviewer. Name
the exact model and reasoning setting in the response. If GPT-5.6 Pro is not
actually available, stop and report the mismatch. Do not silently substitute
an older model, Oracle configuration, or local written review while claiming
GPT-5.6 provenance.

## Decision to review

`haxe.Template` owns its parsing, lookup, iteration, rendering, and error
semantics in staged Haxe source under `std/go/_std`. Its Go runtime binding owns
only three erased-representation facts: array inspection, object
classification, and invoking an already-resolved function value.

Template foreach works for arrays and anonymous structural iterators, but not
for generated Haxe class instances. Generated Go methods intentionally keep
their lowercase Haxe-shaped names, while Go reflection exposes only exported
methods. After method discovery is fixed, Template's current
`cast iterator : Iterator<Dynamic>` would still lower a concrete pointer through
the anonymous `map[string]any` iterator carrier and fail.

Adjudicate the narrow representation boundary that should make generated
method values discoverable and callable without moving Template behavior into
the compiler or exporting user methods as public Go API.

## Non-negotiable ownership and product constraints

- The mainstream Haxe stdlib remains the semantic model. Iteration policy and
  public Template errors stay in staged `haxe.Template` source.
- Compiler ownership is acceptable only for compile-context-sensitive
  generated metadata or a typed representation adapter. It must not become a
  behavior-heavy Template or Reflect reimplementation.
- `hxrt` may invoke an already-resolved function or inspect native carrier
  facts. Because it is a separate Go package, it cannot use reflection to
  access generated lowercase methods.
- Existing generated methods such as `iterator`, `hasNext`, and `next` must
  remain lowercase. Do not rename or export them, and do not add an exported
  internal method that broadens every generated class's public Go method set.
- Do not use `unsafe`, `go:linkname`, application-side raw injection, or a
  Template-specific compiler shim.
- Prefer typed Go AST declarations and expressions. Any residual raw boundary
  must be narrow, justified with Why / What / How documentation, and covered by
  the compiler-debt ratchet.
- Preserve inheritance and virtual dispatch. Looking up `next` on a generated
  subclass must return the overriding bound method, while inherited `hasNext`
  still resolves.
- Keep the solution compatible with the planned staged-Reflect split in
  `haxe_go-vfp.8.7.15.6`; do not make that migration depend on restoring the
  current monolithic compiler-owned Reflect implementation.

## Required source evidence

Read these files at the pinned Haxe.Go commit:

- `std/go/_std/haxe/Template.hx`;
- `std/hxrt/template/NativeTemplate.hx`;
- `runtime/hxrt/template.go` and `runtime/hxrt/template_test.go`;
- `src/reflaxe/go/GoCompiler.hx`, especially `lowerStdlibSymbolShimDecls`,
  `lowerClassDecls`, `lowerConstructorDecl`, `collectDispatchMethods`, and the
  dynamic field/call lowering paths;
- `src/reflaxe/go/ast/GoAST.hx` and `GoASTPrinter.hx` for typed switch support;
- `src/reflaxe/go/compiler/emit/GoRttiMetadataEmitter.hx` as precedent for
  compiler-generated metadata that is separate from library semantics;
- `docs/stdlib-shim-migration-log.md` around the completed Template bridge
  removal;
- `test/semantic_diff/haxe_template_contract/Main.hx` and
  `test/snapshot/stdlib/haxe_template_basic`;
- the compiler-debt and stdlib-ownership contracts that cover compiler, staged
  std, runtime, and generated-output changes.

Use sibling targets only as representation-specific precedent:

- Haxe.Rust commit `ef5c57c90fb694b00b24da94135fe1c53dcfc94f`, especially
  `coerceNominalIteratorToStructural` in `RustCompiler.hx`, demonstrates a
  compile-time callback adapter for its distinct structural iterator ABI;
- Reflaxe.Elixir commit `7f54623cd55ffe3a412fde918b63b56b1c732ecd`
  keeps Reflect behavior in staged std, but Elixir's callable and map
  representations do not solve Go's package visibility problem.

Do not copy Rust ownership/borrowing machinery or Elixir raw injection into Go.

## Reproduction contract

The following semantic-diff fixture passes on Haxe 4.3.7 `--interp` and is the
intended first red contract:

```haxe
class ConcreteIterator {
	public final values:Array<String>;
	public var index:Int;

	public function new(values:Array<String>) {
		this.values = values;
		this.index = 0;
	}

	public function hasNext():Bool {
		return index < values.length;
	}

	public function next():String {
		return "base:" + values[index++];
	}
}

class SpecializedIterator extends ConcreteIterator {
	public function new(values:Array<String>) {
		super(values);
	}

	override public function next():String {
		return "special:" + values[index++];
	}
}

class ConcreteIterable {
	final values:Array<String>;

	public function new(values:Array<String>) {
		this.values = values;
	}

	public function iterator():SpecializedIterator {
		return new SpecializedIterator(values);
	}
}

class Main {
	static function main() {
		var template = new haxe.Template("::foreach items::::__current__::;::end::");
		Sys.println(template.execute({items: new ConcreteIterable(["a", "b"])}));
		Sys.println(template.execute({items: new ConcreteIterator(["x", "y"])}));
	}
}
```

Reference stdout is:

```text
special:a;special:b;
base:x;base:y;
```

At the pinned commit, generated Go instead reaches `Reflect_field` /
`Reflect_hasField`, whose final method path is `MethodByName(key)`. The lowercase
methods are invisible, both iterator probes fail, and execution panics through
Template's `Cannot iter` path. The generated loop also contains an independent
`iterator.(map[string]any)` conversion before reading `hasNext` and `next`.

A natural stricter variant changes the iterable declaration to:

```haxe
public function iterator():Iterator<String> {
	return new SpecializedIterator(values);
}
```

Haxe accepts that structural return, but current generated Go fails to build:

```text
cannot use New_SpecializedIterator(self.values) (value of type
*SpecializedIterator) as map[string]any value in return statement
```

Decide explicitly whether this nominal-to-structural return adaptation belongs
in `.8.7.19` or must be a separately tracked dependency. Do not let the narrow
concrete-return fixture imply that the structural-return gap is solved.

## Candidate metadata shapes

These are discussion candidates, not approved designs.

### A. Selective same-package generated type switch

After all project classes are known, a dedicated compiler emitter could produce
one unexported same-package helper when generated-method discovery is required:

```go
func hxrt__generatedMethodField(obj any, key string) any {
	switch value := obj.(type) {
	case *ConcreteIterable:
		switch key {
		case "iterator":
			return value.__hx_this.iterator
		}
	case *SpecializedIterator:
		switch key {
		case "hasNext":
			return value.__hx_this.hasNext
		case "next":
			return value.__hx_this.next
		}
	}
	return nil
}
```

The real form should use typed `GoTypeSwitch`, `GoSwitch`, selectors, and
returns. Haxe-facing names are lookup keys; existing generated Go selectors are
the returned bound methods. Enumerating the subclass separately and selecting
through `__hx_this` preserves overriding. The helper can remain absent from
programs that never require generated-method discovery.

This avoids changes to user method sets and avoids cross-package access, but it
centralizes a potentially large class/method table in generated support code.

### B. Per-class unexported provider

Generated classes could instead implement an unexported local capability:

```go
type hxrt__GeneratedMethodProvider interface {
	__hx_methodField(string) any
}
```

Each concrete class would emit its own string switch and return bound methods;
the generic helper would only type-assert this provider. This localizes metadata
and avoids one large type switch, but it adds an internal method to every
participating concrete type, complicates selective emission and compiler phase
ordering, and risks universal output growth if lookup requirements are learned
after early classes have already been lowered.

### C. Cross-package registry or exported provider

`hxrt` could receive a generated callback/registry or require an exported
provider method. That crosses the package boundary, but introduces global
mutable registration or a new exported method on generated classes. Treat this
as disfavored unless there is a concrete correctness requirement that A and B
cannot satisfy.

### Rejected directions

- exporting or renaming `iterator`, `hasNext`, or `next`;
- unsafe access to unexported methods;
- a Template-only compiler dispatcher;
- per-application wrappers or raw injection;
- moving foreach state, fallback policy, or errors into `hxrt`.

## Questions to adjudicate

1. Is compile-context-generated method metadata the correct ownership boundary,
   or is there a narrower source/runtime design that satisfies every constraint?
2. Choose A, B, a justified hybrid, or a better concrete alternative. Explain
   the Go visibility, compiler-phase, output-size, and future staged-Reflect
   consequences.
3. Should metadata include all generated instance methods, only public Haxe
   methods, or a statically provable subset? Explain behavior for dynamic lookup
   keys, inherited methods, overrides, `@:native`/name metadata, nil receivers,
   and mutable Haxe dynamic methods.
4. Is returning `value.__hx_this.<method>` sufficient to preserve bound receiver
   identity and virtual dispatch for base-typed values whose concrete receiver
   is a subclass? Identify counterexamples.
5. How should staged `Template` consume the result? In particular, should it
   resolve `hasNext` and `next` with `Reflect.field` and invoke them through the
   existing narrow `NativeTemplate.call`, thereby removing the
   `Iterator<Dynamic>` carrier cast while retaining iteration policy in Haxe?
6. What narrow typed hook should staged `Reflect` eventually use in `.15.6` so
   the generated metadata remains available after the current compiler-owned
   `Reflect_field` and `Reflect_hasField` bodies are retired?
7. Does the natural `Iterator<T>` return failure belong in this bead? If yes,
   specify how a nominal-to-structural adapter composes with method metadata
   without duplicating semantics. If no, define the dependency and wording that
   prevents false closure.
8. Identify any semantic or representation cases the candidates miss: generic
   methods, interfaces, extern/native classes, embedded superclasses, method
   values with parameters or multiple results, static methods, properties,
   private methods, or function-valued fields.
9. Define the minimum red/green contracts: semantic diff, generated-output
   visibility/dispatch shape, direct dynamic method access if warranted,
   negative ownership checks, compiler-debt ratchets, and runtime tests.
10. State intentionally deferred scope. Approval must not imply general
    `Reflect` parity, mutable dynamic-method assignment, static reflection, or
    every structural-interface coercion unless the proposed implementation and
    tests genuinely cover them.

## Required response format

1. Exact model and reasoning setting.
2. Verdict: approve one design, request revision, or block pending evidence.
3. Recommended ownership and data/control flow, with bounded pseudocode.
4. Findings ordered by severity and tied to repository evidence.
5. Explicit `.8.7.19` scope decision for the `Iterator<T>` return variant.
6. Required tests and generated-output invariants.
7. Deferred scope and follow-up issue recommendations.

Do not implement code. This review is the checkpoint that must be adjudicated
before the broad representation refactor begins.
