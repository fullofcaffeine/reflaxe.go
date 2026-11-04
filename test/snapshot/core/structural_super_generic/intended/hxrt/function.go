package hxrt

import "reflect"

// functionIdentity is deliberately nonzero-sized: Go permits distinct zero-sized
// allocations to share an address. A traced token survives aliases and signature
// adapters without retaining callbacks in a global registry.
type functionIdentity struct {
	marker   byte
	name     string
	receiver func() any
}

// Function carries a first-class Haxe callable and its source identity together.
// F is always a concrete Go function signature. Go cannot express a constraint
// covering every function arity, so construction validates that boundary once;
// invocation remains fully typed and does not use reflection.
//
// This primitive does not grant identity to raw native callbacks after an opaque
// roundtrip. Compiler lowering must preserve the whole value through Haxe slots.
type Function[F any] struct {
	call          F
	identity      *functionIdentity
	defaultMethod bool
}

// functionValue is the erased identity boundary used across adapted signatures.
// The callable itself remains typed in Function; identity never depends on Go's
// executable address, captured contents, or wrapper allocation.
type functionValue interface {
	functionIdentity() *functionIdentity
	functionPayload() any
}

// Erasure is confined to reflection; ordinary generated calls use FunctionCall.
func (value *Function[F]) functionPayload() any {
	if value == nil {
		return nil
	}
	return value.call
}

func (value *Function[F]) functionIdentity() *functionIdentity {
	if value == nil {
		return nil
	}
	return value.identity
}

func sameFunctionIdentity(left, right functionValue) bool {
	l, r := left.functionIdentity(), right.functionIdentity()
	if l == r {
		return true
	}
	if l == nil || r == nil || l.name == "" || l.name != r.name {
		return false
	}
	if l.receiver == nil || r.receiver == nil {
		return l.receiver == nil && r.receiver == nil
	}
	// Generated receiver accessors return a nonzero-sized object pointer. Resolve
	// it now: superclass construction can rebind the logical receiver after an
	// earlier method extraction. No native memory layout is inspected.
	return l.receiver() == r.receiver()
}

// equalFunctionIdentity implements ordinary Haxe equality. Bound extraction is
// a fresh closure even when Reflect.compareMethods recognizes the same method.
// Named static declarations retain declaration identity; adapters retain tokens.
func equalFunctionIdentity(left, right functionValue) bool {
	l, r := left.functionIdentity(), right.functionIdentity()
	return l == r || (l != nil && r != nil && l.receiver == nil && r.receiver == nil && sameFunctionIdentity(left, right))
}

// String keeps the runtime carrier opaque to ordinary Haxe string conversion.
func (value *Function[F]) String() string { return "#fun" }

// NamedFunction identifies a finite compiler-owned declaration, independently
// of signature adapters or repeated reads. The key is emitted from source
// declaration identity; it is never derived from executable code.
func NamedFunction[F any](call F, key string) *Function[F] {
	value := NewFunction(call)
	if value != nil {
		value.identity.name = key
	}
	return value
}

// BoundFunction combines one virtual method slot with a live logical receiver.
// Both the callable and receiver accessor must retain the physical receiver.
// Generated accessors return comparable object pointers; arbitrary native value
// receivers require an explicit boundary instead of guessing their identity.
func BoundFunction[F any](call F, receiver func() any, slot string) *Function[F] {
	value := NamedFunction(call, slot)
	if value != nil {
		value.identity.receiver = receiver
	}
	return value
}

// FunctionReceiver normalizes generated object and interface views. The narrow
// method is compiler-owned and remains outside Haxe-visible member metadata.
func FunctionReceiver(value any) any {
	if receiver, ok := value.(interface{ HxFunctionReceiver() any }); ok {
		return receiver.HxFunctionReceiver()
	}
	return value
}

func functionPayloadIsNil[F any](call F) bool {
	value := reflect.ValueOf(call)
	if !value.IsValid() || value.Kind() != reflect.Func {
		panic("hxrt.Function requires a typed Go function")
	}
	return value.IsNil()
}

// NewFunction records one fresh semantic function creation or native import.
// Nil native callbacks remain null. Aliases must copy this value, not re-import
// its native payload, because raw Go functions do not carry portable identity.
func NewFunction[F any](call F) *Function[F] {
	if functionPayloadIsNil(call) {
		return nil
	}
	return &Function[F]{call: call, identity: &functionIdentity{}}
}

// AdaptFunction changes invocation shape without changing source identity.
// The generated adapter must close over the already-evaluated source value;
// it must never re-evaluate the callback expression on each invocation.
func AdaptFunction[F, G any](source *Function[F], call G) *Function[G] {
	if source == nil {
		return nil
	}
	if functionPayloadIsNil(call) {
		panic("hxrt.Function adapter requires a non-null callable")
	}
	return &Function[G]{call: call, identity: source.identity}
}

// FunctionCall guards portable invocation before a nil Go callback can panic.
// Only this source-level null failure becomes a Haxe exception; foreign native
// panics still follow the existing exception boundary unchanged.
func FunctionCall[F any](value *Function[F]) F {
	if value == nil {
		Throw(StringFromLiteral("Cannot call null"))
	}
	return value.call
}

// FunctionNative exposes an exact Go callback signature at an explicit native
// boundary. The returned closure retains its captures through normal Go GC.
func FunctionNative[F any](value *Function[F]) F {
	if value == nil {
		var zero F
		return zero
	}
	return value.call
}

// DefaultDynamicFunction marks the initial implementation in a mutable method
// slot. Haxe binds that default freshly on extraction; assigned function values
// instead retain their stored identity. Only compiler initialization uses this.
func DefaultDynamicFunction[F any](value *Function[F], receiver func() any, slot string) *Function[F] {
	if value == nil {
		return nil
	}
	value.identity.name = slot
	value.identity.receiver = receiver
	value.defaultMethod = true
	return value
}

// ReadDynamicFunction binds an untouched default once, or returns the current
// replacement as stored. The extracted result is never a template: saving or
// assigning it elsewhere cannot cause later reads to bind again.
func ReadDynamicFunction[F any](value *Function[F]) *Function[F] {
	if value == nil || !value.defaultMethod {
		return value
	}
	return BoundFunction(value.call, value.identity.receiver, value.identity.name)
}

func (value *Function[F]) readStoredFunction() any { return ReadDynamicFunction(value) }

// ReadStoredField is the erased generated-field boundary used by reflection.
// Ordinary fields pass through; only an unextracted dynamic default is rebound.
func ReadStoredField(value any) any {
	if function, ok := value.(interface{ readStoredFunction() any }); ok {
		return function.readStoredFunction()
	}
	return value
}
