1. Exact model and reasoning setting

Model: GPT-5.6 Pro
Reasoning setting: maximum reasoning depth

I reviewed the supplied source archive as the authority. Its ZIP comment identifies commit 7d5544e738d72f2e42f88a467b43514ce75f3d6b, matching the requested pushed commit and its public commit record. I did not use an uncommitted working tree as evidence.

The review environment did not contain a Haxe executable, so I did not perform a fresh Haxe compilation. The reproduction conclusions below come from the pinned lowering code, checked-in generated Go snapshots, and the supplied strict-variant Go diagnostic.

2. Verdict

Request revision.

I approve the proposed ownership boundary:

staged haxe.Template owns iteration selection, fallback, loop state, rendering, and public errors;
the compiler may emit final-program metadata needed to select generated lowercase methods;
hxrt may invoke the resulting function value but must not discover it;
generated user methods stay lowercase and are not supplemented with exported provider methods.

I do not approve A, B, or C exactly as written.

The design I recommend is a corrected and factored form of A:

A′: a post-reachability same-package canonical-receiver switch, followed by top-level per-class own-method resolvers with superclass fallback.

A′ has no provider method on generated classes, no registry, no exported API, no unsafe, and no Template-specific compiler branch. It also avoids duplicating every inherited method in every subclass’s table.

Approval of A′ is conditional on three representation corrections:

Every generated ancestor’s __hx_this must be rebound to the most-derived receiver after construction. The current one-hop assignment is insufficient for inheritance depth greater than one.
Staged Template must remove the Iterator<Dynamic> carrier cast and invoke the resolved hasNext and next values through NativeTemplate.call.
The metadata must be reachable through a generic, method-only typed hook that future staged Reflect can use after the current monolithic Reflect_field and Reflect_hasField implementations are retired.

The natural iterator():Iterator<T> return variant is not part of .8.7.19. It requires a separately tracked nominal-to-structural expected-type adapter and must remain explicitly outside this bead’s closure claim.

3. Recommended ownership and data/control flow
3.1 Ownership split

The intended flow should be:

staged haxe.Template
    │
    ├── Reflect.field / Reflect.hasField
    │       │
    │       └── compiler-generated same-package method metadata
    │               └── exact already-bound Go method value
    │
    └── NativeTemplate.call(boundMethod, args)
            └── invocation only

That is consistent with the repository’s stated ownership contract:

Template.hx:49-68 says the three runtime operations expose representation facts while Template policy stays in staged source.
NativeTemplate.hx:3-32 defines exactly array inspection, object classification, and function invocation.
runtime/hxrt/template.go:5-75 implements only those operations.
docs/stdlib-shim-migration-log.md:1203-1230 explicitly records that the concrete-class iterable gap needs general generated-method metadata rather than restoration of a Template shim.
docs/ownership-rubric.md:55-125,180-199 permits compiler ownership for final-type-graph metadata while keeping public semantics in source.

There is no narrower compliant source-only or hxrt-only solution. Go’s reflection API does not expose these lowercase methods, and a separate package cannot directly select them. Conversely, the generated package can select them and can pass the resulting function value across the package boundary for invocation. Go method values retain their receiver when selected.

3.2 Compiler metadata record

During ordinary class lowering, retain a compiler-internal descriptor for every final emitted generated class:

GeneratedClassMethodMetadata {
    generatedGoType
    generatedParent
    hasDispatchCarrier
    ownMethods: [
        {
            haxeLookupKey
            emittedGoSelector
            emittedFunctionType
            haxeMethodKind
        }
    ]
}

Important constraints:

Record own declared or overridden methods only. Inheritance is represented by resolver fallback.
The lookup key is the Haxe-facing field name.
The selector is the actual selector chosen by method lowering. The metadata emitter should not independently repeat normalizeIdent logic and risk drifting from lowerInstanceMethodDecl.
Record only methods that survive Haxe typing/DCE and have emitted bodies. This bead should not invent a new reflection-based retention policy.
Constructors, static functions, interfaces without bodies, and Go-only synthetic methods are excluded.
The table is emitted only when the generic generated-method lookup capability is required.

compileResolvedTypes already provides the right phase boundary: project and dependency class queues are fully drained at GoCompiler.hx:359-391, while support declarations are built afterward at 399-405 and appended at 424-449.

3.3 Generated helper shape

The central helper needs two distinct operations:

Convert the physical carrier passed as Dynamic into its canonical most-derived generated receiver.
Dispatch that concrete receiver to an exact per-class resolver.

Bounded pseudocode:

func hx__generatedMethodField(obj any, key string) any {
	var receiver any

	// Physical carrier -> canonical dynamic receiver.
	switch value := obj.(type) {
	case *ConcreteIterable:
		if value == nil || value.__hx_this == nil {
			return nil
		}
		receiver = value.__hx_this

	case *ConcreteIterator:
		if value == nil || value.__hx_this == nil {
			return nil
		}
		receiver = value.__hx_this

	case *SpecializedIterator:
		if value == nil || value.__hx_this == nil {
			return nil
		}
		receiver = value.__hx_this

	default:
		return nil
	}

	// Canonical receiver -> exact class resolver.
	switch value := receiver.(type) {
	case *ConcreteIterable:
		return hx__methodField_ConcreteIterable(value, key)

	case *ConcreteIterator:
		return hx__methodField_ConcreteIterator(value, key)

	case *SpecializedIterator:
		return hx__methodField_SpecializedIterator(value, key)

	default:
		return nil
	}
}

func hx__methodField_SpecializedIterator(
	value *SpecializedIterator,
	key string,
) any {
	if value == nil {
		return nil
	}

	switch key {
	case "next":
		return value.next
	case "specializedOnly":
		return value.specializedOnly
	}

	if value.ConcreteIterator == nil {
		return nil
	}
	return hx__methodField_ConcreteIterator(
		value.ConcreteIterator,
		key,
	)
}

func hx__methodField_ConcreteIterator(
	value *ConcreteIterator,
	key string,
) any {
	if value == nil {
		return nil
	}

	switch key {
	case "hasNext":
		return value.hasNext
	case "next":
		return value.next
	}

	return nil
}

The names are illustrative. The actual names need a reserved, deterministic, collision-proof compiler namespace.

This is preferable to listing every effective method in every concrete type:

central carrier and receiver switches are O(number of emitted classes);
resolver entries are O(number of declared emitted methods);
inherited lookup is at most O(inheritance depth);
Template resolves each relevant method once before looping, so resolver-chain cost is insignificant;
subclasses that inherit everything need only a parent fallback rather than copies of the entire inherited table.
3.4 Constructor receiver-binding invariant

Current constructor lowering must be strengthened. For:

A <- B <- C

construction of C must establish:

self.B = New_B(...)

self.B.A.__hx_this = self
self.B.__hx_this = self
self.__hx_this = self

// C constructor body follows.

More generally, after the direct superclass constructor returns and before the subclass constructor body runs, every generated ancestor carrier with a __hx_this field must point to the most-derived object.

These assignments should be typed AST. They should be emitted unconditionally as a class ABI invariant, not only when Template or Reflect happens to be imported.

Compiler-owned or native ancestors that do not have the generated dispatch carrier must be excluded or terminate the path.

3.5 Current Reflect insertion point

Until the staged-Reflect migration, current compiler-owned lookup should retain its ordering:

class-token metadata
map carriers
native/direct struct field path
generated-method metadata
exported native MethodByName fallback
not found

Conceptually:

if method := hx__generatedMethodField(obj, key); method != nil {
	return method
}

method := reflect.ValueOf(obj).MethodByName(key)
if method.IsValid() {
	return method.Interface()
}

Reflect_hasField should query the same helper at the corresponding stage.

The generated helper must be a dedicated metadata emitter, analogous in ownership to GoRttiMetadataEmitter, rather than additional behavior embedded in lowerStdlibSymbolShimDecls. GoRttiMetadataEmitter.hx:7-21 is the relevant precedent: staged APIs consume a narrow table derived from final compile-context information. Its current raw implementation is not a precedent for adding new raw code.

3.6 Staged Template loop

Yes: after iterator selection, staged Template should resolve the two methods and invoke them through the existing erased invocation boundary.

Bounded source-level shape:

var hasNext:Dynamic = Reflect.field(iterator, "hasNext");
var next:Dynamic = Reflect.field(iterator, "next");

if (hasNext == null || next == null) {
	throw "Cannot iter on " + value;
}

stack.push(context);
while (NativeTemplate.call(hasNext, []) == true) {
	context = NativeTemplate.call(next, []);
	run(loop);
}
context = popStackValue();

The existing producer/direct-iterator fallback should remain in Haxe source. The important representation change is that there is no longer any:

var iterable:Iterator<Dynamic> = cast iterator;

This matters because a bound generated next():String has a Go function type corresponding to func() *string, not func() any. A Haxe cast to one fixed function type would therefore be too narrow. NativeTemplate.call is the correct existing erased invocation point.

The mainstream Template implementation remains the semantic model for lookup, fallback, iteration, and error ownership.

3.7 Future staged-Reflect hook

The planned .15.6 split should consume one exact, method-only capability, conceptually:

private extern class NativeGeneratedMethodAccess {
	static function field(
		value:Dynamic,
		key:String
	):Dynamic;
}

The exact symbol would lower directly to hx__generatedMethodField in the generated package.

Its contract should be limited to:

accept an already-lowered object and Haxe field key;
return an already-bound generated instance method or null;
perform no map lookup;
perform no data-field lookup;
invoke no getter;
perform no call;
perform no assignment;
enumerate no fields;
choose no error or fallback policy.

Current compiler-owned Reflect_field and Reflect_hasField can call this capability now. Later staged Reflect can call the same capability after those monolithic bodies are removed. That prevents .15.6 from depending on restoration of the current compiler-owned Reflect implementation.

This should be registered and documented as a generic generated-method metadata intrinsic—not as a Template intrinsic.

4. Findings ordered by severity
Critical 1 — Current receiver wiring fails at inheritance depth greater than one

The current class ABI uses embedded superclass pointers and one class-specific __hx_this interface:

lowerClassDecls, GoCompiler.hx:5035-5101, emits the embedded superclass and __hx_this.
collectDispatchMethods, 5667-5709, builds an effective inherited/overridden dispatch interface.
normal virtual reads select through target.__hx_this.<method> at 8268-8272.
upcastIfNeeded, 11208-11229, physically traverses embedded superclass pointers.

But lowerConstructorDecl, 5224-5265, updates only the immediate superclass:

self.Super.__hx_this = self
self.__hx_this = self

For A <- B <- C, construction yields:

C.B.__hx_this   = C
C.B.A.__hx_this = B

The deepest ancestor still identifies the intermediate object.

Consequences:

Upcasting C to A physically produces C.B.A.
Looking up an A-declared virtual method through that carrier sees A.__hx_this == B.
A C override is therefore lost.
Binding an inherited A method and then making a virtual call inside that method also dispatches through stale A.__hx_this, again losing C.

A transitive metadata canonicalizer could follow A -> B -> C for the immediate lookup, but it would not repair virtual calls made from inside the bound base method. Mutating receiver links only when reflection happens would also make ordinary dispatch depend on whether a value had previously been reflected. Neither is acceptable.

Therefore recursive ancestor rebinding is a mandatory representation prerequisite. It should land inside .8.7.19 or in a blocking prerequisite completed before .8.7.19 is declared green.

One limitation remains: the rebinding occurs after the superclass constructor has completed, so virtual calls made during a superclass constructor still cannot dispatch to a not-yet-bound derived receiver. That is a separate pre-existing constructor-dispatch issue and should not be silently claimed as solved here.

Critical 2 — Method discovery alone cannot make the Template fixture pass

The pinned source contains two independent failures.

First, method discovery:

Reflect_field, GoCompiler.hx:4720-4769, ends with MethodByName(key).
Reflect_hasField, 4771-4820, does the same.
the checked-in generated snapshot repeats this at test/snapshot/stdlib/haxe_template_basic/intended/main.go:627-716.

Those calls cannot expose generated lowercase methods.

Second, even after discovery succeeds:

Template.hx:478-480 casts the selected iterator to Iterator<Dynamic>.
anonymous Haxe carriers map to map[string]any in GoTypeMapper.hx:50-51,295-296.
the generated snapshot performs the concrete iterator.(map[string]any) assertion at module_haxe_template.go:1028-1034.
it then extracts hasNext and next from that map at 1036-1051.

Thus adding a generated-method helper without changing staged Template would merely move execution from the first failure to the second one.

The green change must include both:

general generated-method discovery through Reflect;
source-owned removal of the structural iterator carrier cast.
High 3 — Compile-context-generated same-package metadata is the correct narrow boundary

The compiler is the only layer that simultaneously knows:

the final reachable generated class set;
inheritance and override relationships;
Haxe lookup names;
emitted Go selector names;
exact concrete Go types;
whether a class is generated locally or external/native.

Staged Haxe source cannot enumerate those concrete Go types. hxrt, as a separate package, cannot directly select generated lowercase methods. Go reflection does not provide a package-privileged back door for unexported method discovery. Once the generated package selects a method, however, the resulting method value can safely be passed to hxrt as an ordinary function value.

This fits docs/ownership-rubric.md:95-125: final type-table and representation-sensitive primitives are acceptable compiler ownership, provided surrounding public behavior remains in source.

A per-use Template wrapper, compiler-emitted iteration loop, runtime registry, or exported provider would all own more than is necessary.

High 4 — Candidate A needs canonical receiver recovery and inheritance factoring

Candidate A as shown is sufficient for the supplied one-level fixture only because both requested keys are already members of the base dispatch interface.

It is not sufficient for general staged Reflect compatibility.

Suppose a C instance has been physically upcast to *A, and C declares a new method specializedOnly that does not exist on A. A case for *C will not match the physical *A. A case for *A can read value.__hx_this, but its static A dispatch-interface type cannot select specializedOnly.

A′ resolves that by:

recovering the canonical receiver as any;
type-switching on its actual concrete generated class;
invoking that class’s resolver.

Candidate A also risks output proportional to the sum of every class’s effective inherited method set. A′ stores each method once at its declaring/overriding class.

High 5 — Candidate B is possible but inferior; candidate C remains unjustified

B is not impossible merely because classes have already been emitted: Go permits receiver method declarations elsewhere in the same package, so a post-pass could technically add providers.

It is nevertheless inferior:

every concrete subclass needs its own provider, including subclasses that only inherit methods;
relying on a promoted base provider would bind the base receiver and would not know subclass-only keys;
the provider changes each participating class’s Go method set;
it creates collision and accidental interface-satisfaction concerns;
it couples final program-wide reflection requirements back into per-type method emission;
it provides no capability unavailable to top-level same-package resolver functions.

C has no compensating correctness benefit:

an exported provider violates the public-method-set constraint;
a callback or registry introduces initialization order, mutable global state, and package coupling;
the same generated package can already select the method directly.
High 6 — Metadata must cover emitted Haxe-facing instance methods, not just three Template names

The method key can be computed at runtime. Future staged Reflect must not depend on a fixed literal set such as iterator, hasNext, and next.

The correct initial policy is:

Include every final emitted Haxe-facing instance method with a body, for every final emitted generated class, but only when the generated-method capability is required.

Do not include only public methods. Haxe access control is a compile-time concept; dynamic lookup should follow the runtime field model. A public-only table would make private methods disappear specifically on Go and would not match the compiler’s current instance-field metadata style.

Do not let the metadata table itself keep methods that Haxe DCE removed. Existing @:keep and reflection-retention policy remains authoritative.

The detailed case decisions should be:

Case	Required behavior
Dynamic lookup key	Full surviving method inventory; no literal-key-only optimization in .8.7.19.
Inherited method	Resolved through parent resolver fallback.
Override	Most-derived resolver checks first and returns the overriding bound method.
Private method	Included if emitted as a Haxe instance method.
@:native / name metadata	Haxe name remains the lookup key; selector comes from the actual lowering descriptor. Do not independently guess aliases.
Interface-typed value	The any value retains its concrete generated pointer; concrete class metadata handles it. Interface declarations themselves contribute no body.
Extern/native class	Excluded from local generated metadata. Existing exported-native reflection remains the fallback. Cross-package lowercase native methods cannot be recovered under the stated constraints.
Generic method	Return its exact emitted bound function value; do not add a generic invocation adapter here.
Parameters	Supported by the metadata because the exact method value is returned. Argument conversion remains the caller/runtime contract.
Multiple Go results	Not generalized by this bead. Ordinary generated Haxe methods normally have the backend’s Haxe result shape; native multi-result APIs remain separate.
Static method	Excluded.
Property/getter/setter	No synthetic property aliases. Reflect.getProperty policy belongs to staged Reflect.
Function-valued data field	Existing data-field lookup must precede method metadata. Generated unexported data-field reflection is a separate gap.
dynamic Haxe method	Initial emitted method read may be included. Per-instance replacement and assignment are not solved.
Typed nil pointer	Return null/false without a Go panic.
Nonnil but unbound/incomplete object	Conservatively return null until empty-instance/deserialization binding is separately made complete.
High 7 — value.__hx_this.<method> is conditionally sufficient, not generally sufficient

For a normally constructed direct subclass and a method that exists in the static dispatch interface, the expression produces the desired overriding bound method.

It is not sufficient in these cases:

Deep inheritance: a base carrier’s __hx_this is currently stale after more than one inheritance edge.
Subclass-only method through a physical base pointer: the base dispatch interface cannot name the subclass-only selector.
Empty or deserialized object: __hx_this or embedded superclass pointers may be nil.
Inherited method body: the bound receiver may be the owner’s embedded subobject; virtual calls inside it still require every ancestor’s __hx_this to identify the most-derived object.
Superclass-constructor virtual call: the derived receiver has not yet been rebound.
Signature incompatibility: any override that lowers to a Go signature different from an ancestor’s dispatch-interface signature already prevents the derived pointer from satisfying that interface. Metadata cannot repair such ABI incompatibility.

A′ handles items 1–4 for normally constructed values through recursive rebinding, concrete receiver recovery, and leaf-to-parent resolver order. Items 5 and any broader signature-normalization issue remain separate.

Go method-value semantics themselves are suitable here: selecting value.next saves the receiver in the resulting function value.

Medium 8 — Empty-instance and deserialization paths need explicit guards

The repository has object-creation paths that do not establish the normal constructor invariant:

GoTypeReflectionEmitter.hx:401-414 returns &Type{} for empty construction.
GoRegexSerializerEmitter.hx:87-99 also starts from an empty struct.
its binder at 529-557 binds only the outer __hx_this; it does not allocate and bind the inherited carrier chain.

A method helper that blindly selects through these fields can panic.

For .8.7.19, typed nils, nil self interfaces, and missing embedded parents should return no generated method. Fully repairing Type.createEmptyInstance and unserialization should be a follow-up unless the bead deliberately expands to cover and test those creation paths.

The new helper must not reuse or expand the serializer’s existing unsafe path.

Medium 9 — NativeTemplate.call must not accidentally define malformed-iterator policy

TemplateCall returns nil for nil and non-function values, and runtime/hxrt/template_test.go:64-69 explicitly tests that behavior.

For valid generated method metadata, that is not a problem: the helper returns actual function values.

For malformed structural iterators, however, blindly doing:

while (NativeTemplate.call(hasNext, []) == true)

could silently terminate rather than producing the failure that a direct dynamic call would have produced. A missing next combined with a true hasNext could also produce bad loop behavior unless checked.

Therefore staged Template must at least resolve and validate the presence of both method values before entering the loop. The treatment of a non-null but non-callable value must be specified and covered by source-level tests. It must not be determined accidentally by TemplateCall’s nil-on-nonfunction implementation.

A generic future Reflect.isFunction representation predicate would be a better home for callability than a fourth Template-specific runtime operation. Exact malformed-iterator exception parity may also be explicitly deferred, but the limitation must be recorded.

Medium 10 — The new emitter should be fully typed AST

The required nodes already exist:

GoTypeSwitchCase, GoTypeSwitch, GoSwitch, selectors, calls, and returns are defined in GoAST.hx:34-53,77-135.
their typed printer support is present in GoASTPrinter.hx:249-334.

No new GoRaw is justified for the class switches, key switches, selectors, nil checks, fallback calls, or constructor rebinding.

docs/compiler-debt-ratchet.md:32-48,97-132 classifies raw Go AST as avoidable debt and has a zero baseline for direct unsafe touchpoints. The new boundary should add neither.

5. Explicit .8.7.19 scope decision for Iterator<T> returns
Decision: exclude it and track it separately

The stricter declaration:

public function iterator():Iterator<String> {
	return new SpecializedIterator(values);
}

fails at a different representation boundary from dynamic method discovery.

Evidence:

Iterator<T> follows the anonymous/structural type path and maps to map[string]any in GoTypeMapper.hx:50-51,295-296.
return lowering at GoCompiler.hx:5921-5955 applies only upcastIfNeeded.
upcastIfNeeded, 11208-11229, understands nominal superclass paths, not nominal-to-structural adaptation.

The failure therefore occurs while lowering the return value, before Template or Reflect receives the iterator.

The separate issue should be titled along the lines of:

Nominal generated class to structural method-carrier coercion at typed expected-type boundaries

It must cover more than return statements:

return values;
assignments and local initializers;
function arguments;
conditional/switch branch joins;
callback results;
any other site where a nominal source is consumed under an anonymous method-shaped expected type.

The adapter should:

evaluate the source expression once;
prove that the nominal type supplies the required structural methods;
produce the target’s established structural carrier or typed callbacks;
use the same central method symbol/signature inventory as the generated-method metadata;
preserve virtual dispatch through the same receiver-binding invariant;
contain no iteration, fallback, or Template error semantics.

The Rust sibling is relevant only as precedent that a target may need a compile-time callback adapter for its structural iterator ABI. Its ownership and borrowing machinery should not be copied into Go. The Elixir sibling supports the source-owned Reflect split, but its maps and callables do not address Go package visibility.

The .8.7.19 closure statement should say exactly:

.8.7.19 makes emitted instance methods on nominal generated class carriers discoverable as bound method values through a compile-context-generated same-package capability, and removes Template’s post-selection Iterator<Dynamic> map-carrier cast. It does not make nominal generated classes assignable to structural Iterator<T> or other anonymous method carriers at typed expected-type boundaries.

The strict fixture should remain a linked expected failure. Do not make its current incidental Go compiler wording a permanent diagnostic contract.

6. Required tests and generated-output invariants
6.1 Semantic red/green contracts

The minimum green suite should include:

The supplied fixture verbatim, with exact stdout:

special:a;special:b;
base:x;base:y;
Existing array and anonymous structural iterator coverage remains green.
The current fixture in test/semantic_diff/haxe_template_contract/Main.hx:10-33 already covers arrays and an anonymous iterator/hasNext/next carrier.

Direct generated-method reflection independent of Template.

Exercise:

var key = "ne" + "xt";
var fn:Void->String = cast Reflect.field(baseTypedSubclass, key);
Sys.println(fn());

Also assert Reflect.hasField and unknown-key behavior.

Three-level inheritance.

For A <- B <- C:

A declares hasNext and next;
C overrides next;
the object is observed as C, B, and A;
all reflective next lookups invoke C;
inherited hasNext resolves;
the Template loop uses C’s override.

Subclass-only method through a base physical carrier.

var a:A = new C();
var fn:Void->String = cast Reflect.field(a, "cOnly");

This is the contract that distinguishes A′ from the proposed A example.

Virtual dispatch from inside a bound base method.

A declares:

public function callNext():String {
    return next();
}

Reflectively resolve callNext from an A-typed C and verify that the internal call reaches C’s override. This proves recursive ancestor rebinding rather than only correct top-level lookup.

Computed key and normalized selector.

Include a legal Haxe method whose Go selector must be normalized, while the lookup key remains the original Haxe name.

Private emitted method.

Verify the chosen all-Haxe-instance-method policy rather than accidentally implementing public-only metadata.

Interface-typed reference with a generated concrete implementation.

The runtime Dynamic value must recover the concrete generated pointer.

Typed nil and unknown key.

Neither Reflect.field nor Reflect.hasField may trigger a Go panic.

Parameterized bound method.

Resolve a generated method that accepts an argument and verify receiver identity and state mutation. This tests that metadata returns the exact method value rather than a zero-argument Template wrapper.

Initial emitted dynamic method read, if MethDynamic bodies are included by the selected inventory policy. Reassignment remains a negative/deferred contract.
6.2 Explicit red dependency

Keep the iterator():Iterator<String> variant as an expected failure linked to the separate expected-type-adapter issue.

Its test description must state that .8.7.19 does not close it.

6.3 Generated-output invariants

Snapshots and structural checks should require:

generated methods such as iterator, hasNext, and next remain lowercase;
no exported internal provider is added;
no unexported provider method is added to generated classes;
no registry, init registration, mutable callback slot, or package-global method map is emitted;
the central helper is unexported;
helper and resolver names are deterministic and collision-proof;
the helper is absent when the generic generated-method lookup capability is not used;
the helper’s first switch covers final emitted generated physical carriers;
it canonicalizes through __hx_this;
the second switch dispatches the canonical concrete receiver;
each per-class resolver contains only that class’s declared/overridden Haxe method entries;
each resolver has at most one generated-parent fallback;
a subclass override is checked before parent fallback;
Haxe-facing names are used as string keys;
existing emitted Go selectors are used for method selection;
typed nil and missing-parent guards are present;
deep constructors rebind every generated ancestor’s __hx_this;
those bindings occur after the superclass constructor and before the subclass body;
current Reflect class-token, map, and data-field ordering is preserved;
generated-method lookup occurs before exported-native MethodByName;
the generated Template file contains no iterator.(map[string]any) conversion;
it contains no map extraction of hasNext or next after iterator selection;
it still contains the source-owned iterator-selection and context-stack behavior;
NativeTemplate remains exactly the existing three functions;
no template_support shim group reappears;
no Template loop, fallback policy, or error string is generated by the compiler;
no new GoRaw, unsafe, go:linkname, or application-side injection is introduced.

The constructor ABI correction will necessarily change generated output for affected deep inheritance programs. The method metadata helper, however, must remain absent from programs that do not require generated method discovery.

6.4 Runtime tests

runtime/hxrt/template_test.go should add a bound-method invocation test such as:

type templateCounter struct {
	value int
}

func (c *templateCounter) next() int {
	c.value++
	return c.value
}

func TestTemplateCallBoundMethod(t *testing.T) {
	counter := &templateCounter{}
	next := counter.next

	if got := TemplateCall(next, nil); got != 1 {
		t.Fatalf("first call = %#v, want 1", got)
	}
	if got := TemplateCall(next, nil); got != 2 {
		t.Fatalf("second call = %#v, want 2", got)
	}
}

Also retain:

parameterized function invocation;
void-result invocation;
nil and non-function behavior unless a separately reviewed general callability contract changes it.

There should be no runtime method-discovery test, because method discovery is deliberately not owned by hxrt.

6.5 Governance checks

Required governance updates:

a dedicated What / Why / How document for the generated-method metadata emitter;
a generic compiler-intrinsic registry entry if an exact Haxe-facing hook is introduced;
no Template-specific intrinsic or template_support restoration;
compiler-debt ratchet verification showing no raw or unsafe growth;
ownership-map and migration-log updates distinguishing method metadata from Template semantics;
negative tests equivalent to test_stdlib_migration_ledger_contract.py:700-789;
generated-output tests asserting no provider method or exported user API growth;
full semantic-diff, snapshot, Go test, vet/static-analysis, and selective-runtime footprint lanes appropriate to compiler/runtime/generated-output changes.
7. Deferred scope and follow-up recommendations

Approval of A′ must not imply support for the following:

Nominal-to-structural expected-type coercion.
Track the Iterator<T> return case and all equivalent assignment/argument/result boundaries separately.
General staged Reflect parity.
Generated data fields, Reflect.fields, class/static reflection, compare-method semantics, property policy, and complete native-field behavior are not part of .8.7.19.
Mutable Haxe dynamic-method assignment.
Reading an initially emitted dynamic method is not support for replacing it per instance.
Generated function-valued data fields.
Those have the same unexported-field visibility problem and need a separate generated-field metadata design.
Properties and accessors.
Reflect.field method metadata must not implicitly become Reflect.getProperty.
Static methods.
Global generated functions and class-token static reflection require separate metadata.
Empty-instance and unserialized inherited object binding.
Type.createEmptyInstance and deserialization need complete embedded-parent allocation and receiver binding before their method semantics can be claimed.
Virtual dispatch during superclass construction.
Recursive post-super rebinding does not make derived overrides available while the superclass constructor itself is running.
Exact malformed structural-iterator error parity.
Non-null/non-function hasNext or next values need an explicit staged-source contract or a separate generic callability capability.
Cross-package lowercase methods on extern/native classes.
They remain inaccessible without violating the package-visibility constraints.
General argument/result conversion for reflective calls.
Metadata returns exact method values. Generic coercion, variadics, native multiple results, and signature normalization belong to the eventual Reflect invocation boundary.
Every generic or structural-interface coercion.
The method table is not an implicit solution for arbitrary structural typing.
Future multi-package generated output.
A′ assumes the current generated classes and helper share a Go package. A future package-splitting architecture would need one local resolver per package and an explicitly designed aggregation boundary.

With recursive ancestor rebinding, canonical concrete-receiver recovery, per-class top-level resolver fallback, and the source-owned removal of the iterator map cast, A′ is an approvable narrow representation boundary. Without those corrections, the proposed metadata would pass only the shallow fixture while leaving required inheritance and carrier semantics unsound.
