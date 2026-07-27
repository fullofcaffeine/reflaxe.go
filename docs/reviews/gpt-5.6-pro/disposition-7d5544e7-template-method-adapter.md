# Template method-adapter review disposition

## Outcome

The review verdict is **Request revision**. Haxe.Go accepts the corrected A′
architecture as the design direction for Bead `haxe_go-vfp.8.7.19`, subject to
red-first local proof. It does not accept candidates A, B, or C exactly as they
appeared in the prompt.

A′ keeps Template semantics in staged `haxe.Template`, lets the compiler emit
selective final-program method metadata, and leaves `hxrt` responsible only for
invoking an already-resolved function value. Generated Haxe methods remain
lowercase Go methods. The compiler emits no provider method, exported adapter,
mutable registry, reflection workaround, or Template-specific behavior shim.

The response was supplied manually. Its model and reasoning claims are
reviewer-declared rather than Caf-receipt-backed, so the response is advisory
evidence rather than implementation authority.

## Accepted design conditions

| Review condition | Local disposition | Owner |
| --- | --- | --- |
| Recover the canonical most-derived receiver, then dispatch through per-class own-method resolvers with one superclass fallback. | Accepted as the metadata shape to prove. It avoids inherited-table duplication and supports subclass-only methods through a physically upcast carrier. | `haxe_go-vfp.8.7.19` |
| Rebind every generated ancestor's `__hx_this` after the direct superclass constructor returns. | Accepted as a blocking class-ABI correction. Current source still updates only the immediate superclass, so a three-level red contract is required before implementation. | `haxe_go-vfp.8.7.19` or a blocking child if the slice must land independently |
| Remove Template's `Iterator<Dynamic>` cast and invoke resolved `hasNext`/`next` method values through `NativeTemplate.call`. | Accepted. Current staged source still contains the cast, so metadata alone cannot make the fixture green. Lookup, fallback, loop state, and public errors remain in Haxe source. | `haxe_go-vfp.8.7.19` |
| Expose one generic method-only typed hook usable by current compiler-owned Reflect and future staged Reflect. | Accepted. The hook returns an already-bound generated method or null and owns no field, property, call, enumeration, or error policy. | `haxe_go-vfp.8.7.19`, consumed later by `haxe_go-vfp.8.7.15.6` |
| Emit metadata only when the generic capability is reachable, using typed Go AST and the selector already chosen by normal lowering. | Accepted. No new `GoRaw`, `unsafe`, runtime discovery, or separately normalized selector is allowed. | `haxe_go-vfp.8.7.19` |
| Keep the `iterator():Iterator<T>` expected-type adapter outside this bead. | Accepted as a scope boundary, but the work is no longer pending: `haxe_go-vfp.8.3.3` landed the shared typed adapter after the reviewed commit. `.8.7.19` must not claim that work as its own. | Closed `haxe_go-vfp.8.3.3` and its follow-up slices |

## Finding adjudication

- Critical findings 1 and 2 are accepted. Local source confirms both one-hop
  receiver rebinding and the remaining staged Template cast.
- High findings 3 through 7 are accepted as design constraints. A′ is narrower
  than provider or registry designs and preserves the future staged-Reflect
  ownership boundary.
- Medium finding 8 is accepted as a required nil/missing-parent guard for this
  feature. Fully repairing empty-instance and deserialization construction is
  deferred and must not be implied by closure.
- Medium finding 9 is accepted with a source-owned contract: Template must at
  least reject missing `hasNext` or `next` before entering the loop. Exact
  non-null/non-callable parity remains deferred unless a red source contract
  proves it belongs in this bead.
- Medium finding 10 is accepted without qualification: the new declarations,
  switches, selectors, guards, and constructor assignments use typed Go AST.

## Required first red contracts

Implementation starts with failures that distinguish A′ from the shallow
candidate A:

1. the supplied concrete Template iterable and direct-iterator output;
2. direct computed-key `Reflect.field` and `Reflect.hasField` method lookup;
3. three-level inheritance observed through leaf, middle, and base carriers;
4. a subclass-only method recovered through a physically upcast base carrier;
5. virtual dispatch from inside a reflectively bound base method;
6. typed nil, unknown-key, parameterized method, normalized-selector, and private
   emitted-method behavior;
7. generated-output contracts excluding providers, registries, exported method
   growth, raw code, and unconditional metadata;
8. a bound-method invocation test for `NativeTemplate.call` with no runtime
   method-discovery responsibility.

The review does not close the bead. Green Haxe/Go evidence, compiler-debt gates,
documentation, and the required `thinking:xhigh` closure pass remain mandatory.
