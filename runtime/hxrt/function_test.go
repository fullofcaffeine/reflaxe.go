package hxrt

import (
	"runtime"
	"testing"
)

func TestFunctionIdentityPreservesAliasesAndCapture(t *testing.T) {
	makeCounter := func() *Function[func(int) int] {
		total := 0
		return NewFunction(func(delta int) int { total += delta; return total })
	}
	first, second := makeCounter(), makeCounter()
	alias := first
	if !sameFunctionIdentity(first, alias) || sameFunctionIdentity(first, second) {
		t.Fatal("aliases must match and independent factory evaluations must differ")
	}
	if FunctionCall(first)(2) != 2 || FunctionCall(alias)(3) != 5 || FunctionCall(second)(1) != 1 {
		t.Fatal("identity transport changed mutable capture ownership")
	}
	runtime.GC()
	if FunctionCall(alias)(1) != 6 {
		t.Fatal("a live alias must retain its captures")
	}
}

func TestFunctionAdaptationPreservesIdentityAcrossSignatures(t *testing.T) {
	original := NewFunction(func(value int) int { return value + 1 })
	adapted := AdaptFunction(original, func(value int64) int64 { return int64(FunctionCall(original)(int(value))) })
	if !sameFunctionIdentity(original, adapted) || FunctionCall(adapted)(4) != 5 {
		t.Fatal("signature-only adaptation must preserve identity and invocation")
	}
	independent := NewFunction(FunctionNative(original))
	if sameFunctionIdentity(original, independent) {
		t.Fatal("a fresh native import cannot infer provenance from executable code")
	}
}

func TestFunctionNullTransportAndGuard(t *testing.T) {
	var native func(int) int
	null := NewFunction(native)
	if null != nil || FunctionNative(null) != nil || AdaptFunction(null, func(int64) int64 { return 0 }) != nil {
		t.Fatal("null must survive import, export, and signature adaptation")
	}
	var otherNull *Function[func(int64) int64]
	if !sameFunctionIdentity(null, otherNull) || sameFunctionIdentity(null, NewFunction(func(int) int { return 0 })) {
		t.Fatal("null identity must be independent of the callable signature")
	}
	caught := false
	TryCatch(func() { FunctionCall(null)(1) }, func(value any) {
		caught = true
		message, ok := value.(*string)
		if !ok || *message != "Cannot call null" {
			t.Fatalf("null failure must match the pinned Haxe interpreter: %v", value)
		}
	})
	if !caught {
		t.Fatal("null invocation must enter a Haxe catch")
	}
}

func TestFunctionNativePanicRemainsForeign(t *testing.T) {
	marker := &struct{ value int }{1}
	defer func() {
		if recover() != marker {
			t.Fatal("foreign panic must retain its original value")
		}
	}()
	callback := NewFunction(func() { panic(marker) })
	TryCatch(func() { FunctionCall(callback)() }, func(value any) {
		t.Fatal("a Haxe catch must not absorb a foreign native panic")
	})
}

func TestFunctionRejectsNonCallableNativePayload(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("a non-function payload must be rejected at construction")
		}
	}()
	NewFunction(42)
}

func TestFunctionNativeCallbackRetainsCapture(t *testing.T) {
	value := 7
	callback := FunctionNative(NewFunction(func(delta int) int { value += delta; return value }))
	runtime.GC()
	if callback(2) != 9 || callback(1) != 10 {
		t.Fatal("a retained native callback must keep the original mutable capture")
	}
}
