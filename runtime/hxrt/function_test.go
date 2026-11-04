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
	if !ReflectCompareMethods(null, nil) || !ReflectCompareMethods(nil, null) || !HaxeEqual(null, nil) {
		t.Fatal("typed and erased null functions must compare equally in both orders")
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

func TestFunctionReflectionAcrossSignatures(t *testing.T) {
	original := NewFunction(func(value int) int { return value + 1 })
	adapted := AdaptFunction(original, func(value int64) int64 { return int64(FunctionCall(original)(int(value))) })
	if !ReflectCompareMethods(original, adapted) || !HaxeEqual(original, adapted) {
		t.Fatal("reflection and erased equality must preserve adapted identity")
	}
	if !ReflectIsFunction(original) || ReflectIsObject(original) || len(ReflectFields(original)) != 0 {
		t.Fatal("the runtime carrier must remain a function with hidden implementation fields")
	}
	if ReflectCallMethod(original, []any{4}) != 5 {
		t.Fatal("reflective invocation must reach the typed payload")
	}
}

func TestFunctionNamedAndReboundMethodIdentity(t *testing.T) {
	first := NamedFunction(func() int { return 1 }, "Module.first")
	secondRead := NamedFunction(func() int { return 1 }, "Module.first")
	other := NamedFunction(func() int { return 1 }, "Module.other")
	if !sameFunctionIdentity(first, secondRead) || sameFunctionIdentity(first, other) {
		t.Fatal("declaration identity must survive reads without equating different declarations")
	}
	base, child := &struct{ value int }{1}, &struct{ value int }{2}
	logical := base
	before := BoundFunction(func() int { return logical.value }, func() any { return logical }, "read")
	logical = child
	after := BoundFunction(func() int { return child.value }, func() any { return child }, "read")
	different := BoundFunction(func() int { return base.value }, func() any { return base }, "read")
	if !sameFunctionIdentity(before, after) || sameFunctionIdentity(before, different) || FunctionCall(before)() != 2 {
		t.Fatal("method identity and invocation must follow constructor rebinding")
	}
}

func TestFunctionVarArgsCreationIdentity(t *testing.T) {
	callback := NewFunction(func(arguments *Array) any { return len(arguments.Values()) })
	values := make([]any, 2)
	for index := range values {
		values[index] = ReflectMakeVarArgs(callback)
	}
	if !ReflectCompareMethods(values[0], values[0]) || ReflectCompareMethods(values[0], values[1]) {
		t.Fatal("each varargs factory evaluation must create one distinct function")
	}
	if ReflectCallMethod(values[0], []any{1, 2}) != 2 {
		t.Fatal("varargs identity must preserve reflective invocation")
	}
}

// Ordinary equality keeps each bound extraction distinct, while method comparison
// deliberately identifies the same receiver and method slot.
func TestFunctionOrdinaryEqualityAndObservation(t *testing.T) {
	receiver := &struct{ value int }{4}
	first := BoundFunction(func() int { return receiver.value }, func() any { return receiver }, "read")
	second := BoundFunction(func() int { return receiver.value }, func() any { return receiver }, "read")
	if HaxeEqual(first, second) || !HaxeEqual(first, first) || !ReflectCompareMethods(first, second) {
		t.Fatal("ordinary equality must distinguish repeated bound method extractions")
	}
	named := NamedFunction(func() int { return 1 }, "Module.named")
	again := NamedFunction(func() int { return 1 }, "Module.named")
	if !HaxeEqual(named, again) {
		t.Fatal("static declaration equality must survive repeated reads")
	}
	if got := *StdString(first); got != "#fun" {
		t.Fatalf("function string = %q", got)
	}
	if got := *JsonStringify(first, nil); got != `"<fun>"` {
		t.Fatalf("function JSON = %q", got)
	}
	if got := *JsonStringify(map[string]any{"call": first}, nil); got != `{}` {
		t.Fatalf("function field JSON = %q", got)
	}
	if TemplateIsObject(first) {
		t.Fatal("function carrier is not a template object")
	}
	if got := TemplateCall(first, nil); got != 4 {
		t.Fatalf("template callback = %v", got)
	}
}

func TestDynamicMethodDefaultExtraction(t *testing.T) {
	receiver := &struct{ value int }{value: 3}
	slot := DefaultDynamicFunction(NewFunction(func() int { return receiver.value }), func() any { return receiver }, "next")
	first, second := ReadDynamicFunction(slot), ReadDynamicFunction(slot)
	if HaxeEqual(first, second) || !ReflectCompareMethods(first, second) {
		t.Fatal("default method extraction identity")
	}
	if ReadDynamicFunction(first) != first {
		t.Fatal("saved extraction remained a default template")
	}
	replacement := NewFunction(func() int { return 9 })
	slot = replacement
	if ReadDynamicFunction(slot) != replacement {
		t.Fatal("replacement identity changed")
	}
	if FunctionCall(first)() != 3 || FunctionCall(ReadDynamicFunction(slot))() != 9 {
		t.Fatal("saved binding followed replacement")
	}
}

// Projection is a compiler representation change; it must not erase the source
// result or default policy observed through Dynamic and explicit reflection.
func TestFunctionProjectionRetainsNormalizedSource(t *testing.T) {
	original := NewFunction(func(value int) int { return value + 2 })
	normalized := AdaptFunction(original, func(value any) int {
		if value == nil {
			return FunctionCall(original)(7)
		}
		return FunctionCall(original)(value.(int))
	})
	discarded := ProjectFunction(normalized, func(value any) { FunctionCall(normalized)(value) })
	if FunctionOrigin(discarded) != normalized {
		t.Fatal("result discard replaced the normalized callable source")
	}
	if !HaxeEqual(original, discarded) || !ReflectCompareMethods(original, discarded) {
		t.Fatal("projection changed source function identity")
	}
	if got := ReflectCallMethod(discarded, []any{nil}); got != 9 {
		t.Fatalf("reflection lost source default/result: %v", got)
	}
	if got := ReflectCallMethod(discarded, []any{3}); got != 5 {
		t.Fatalf("reflection ignored an explicit source argument: %v", got)
	}
	current := discarded
	for i := 0; i < 1000; i++ {
		// Generated recovery must build its next call from this source, never from
		// the previous projection. This asserts the runtime provenance contract.
		root := FunctionOrigin(current).(*Function[func(any) int])
		current = ProjectFunction(root, func(value any) { FunctionCall(root)(value) })
		if FunctionOrigin(current) != normalized {
			t.Fatal("projection accumulated origins")
		}
	}
	if ReflectCallMethod(current, []any{nil}) != 9 {
		t.Fatal("roundtrip lost source invocation")
	}
	var absent *Function[func(int) int]
	if ProjectFunction(absent, func(int) {}) != nil || FunctionOrigin(absent) != nil {
		t.Fatal("projection materialized a null callable")
	}
}
