package hxrt

import "reflect"

// functionIdentity is deliberately nonzero-sized: Go permits distinct zero-sized
// allocations to share an address. A traced token survives aliases and signature
// adapters without retaining callbacks in a global registry.
type functionIdentity struct{ marker byte }

// Function carries a first-class Haxe callable and its source identity together.
// F is always a concrete Go function signature. Go cannot express a constraint
// covering every function arity, so construction validates that boundary once;
// invocation remains fully typed and does not use reflection.
//
// This primitive does not grant identity to raw native callbacks after an opaque
// roundtrip. Compiler lowering must preserve the whole value through Haxe slots.
type Function[F any] struct {
	call     F
	identity *functionIdentity
}

// functionValue is the erased identity boundary used across adapted signatures.
// The callable itself remains typed in Function; identity never depends on Go's
// executable address, captured contents, or wrapper allocation.
type functionValue interface {
	functionIdentity() *functionIdentity
}

func (value *Function[F]) functionIdentity() *functionIdentity {
	if value == nil {
		return nil
	}
	return value.identity
}

func sameFunctionIdentity(left, right functionValue) bool {
	return left.functionIdentity() == right.functionIdentity()
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
		Throw(StringFromLiteral("Cannot call null function"))
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
