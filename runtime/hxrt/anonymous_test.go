package hxrt

import "testing"

// These contracts exercise storage identity and live access independently of
// compiler lowering. Generated access plans remain responsible for member policy.
func TestAnonymousMapRetainsStorage(t *testing.T) {
	raw := map[string]any{"value": 1, "present": nil}
	value := AnonymousFromMap(raw)
	if !HaxeEqual(value, raw) || !HaxeEqual(raw, value) {
		t.Fatal("map admission changed identity")
	}
	if got := AnonymousLookup(value, "present"); !got.Found || got.Value != nil {
		t.Fatal("lost present null field")
	}
	if AnonymousLookup(value, "missing").Found {
		t.Fatal("invented missing field")
	}
	AnonymousSet(value, "value", 2)
	if raw["value"] != 2 {
		t.Fatal("write did not reach original map")
	}
	raw["value"] = 3
	if AnonymousGet(value, "value") != 3 {
		t.Fatal("read used a snapshot")
	}
	if AnonymousFromMap(nil) != nil {
		t.Fatal("nil map became an object")
	}
}

type anonymousTestSource struct {
	value   int
	logical *anonymousTestSource
}

func (source *anonymousTestSource) HxFunctionReceiver() any {
	if source.logical != nil {
		return source.logical
	}
	return source
}

func TestAnonymousViewRetainsSourceAndLiveAccess(t *testing.T) {
	source := &anonymousTestSource{value: 1}
	reads := 0
	view := NewAnonymousView(source, func(name string) AnonymousField {
		reads++
		return AnonymousField{Value: source.value, Found: name == "value"}
	}, func(name string, value any) bool {
		if name != "value" {
			return false
		}
		source.value = value.(int)
		return true
	})
	if reads != 0 {
		t.Fatal("view creation read a member")
	}
	if !HaxeEqual(view, source) || !HaxeEqual(source, view) {
		t.Fatal("view identity is not source identity")
	}
	source.value = 4
	if AnonymousGet(view, "value") != 4 {
		t.Fatal("view lost live source")
	}
	AnonymousSet(view, "value", 5)
	if source.value != 5 {
		t.Fatal("view write did not reach source")
	}
	derived := &anonymousTestSource{value: 9}
	source.logical = derived
	if !HaxeEqual(view, derived) || ObjectSource(view) != derived {
		t.Fatal("logical receiver was frozen")
	}
	var absent *anonymousTestSource
	if NewAnonymousView(absent, nil, nil) != nil {
		t.Fatal("typed nil created a view")
	}
}

func TestAnonymousMapObservationDoesNotExposeCarrier(t *testing.T) {
	raw := map[string]any{"name": StringFromLiteral("sample")}
	value := AnonymousFromMap(raw)
	if *StdString(value) != *StdString(raw) {
		t.Fatal("string exposes carrier")
	}
	field := StringFromLiteral("name")
	if !ReflectLookupField(value, field).Found {
		t.Fatal("reflection lost source map")
	}
	copy := ReflectCopy(value)
	if HaxeEqual(copy, value) {
		t.Fatal("copy retained source identity")
	}
	if !ReflectSetField(copy, field, StringFromLiteral("copy")) {
		t.Fatal("copy is not writable")
	}
	if *raw["name"].(*string) != "sample" {
		t.Fatal("copy mutated source")
	}
}

func TestAnonymousJSONNormalizesBeforeCycleDetection(t *testing.T) {
	raw := map[string]any{"value": 3}
	value := AnonymousFromMap(raw)
	if got := *JsonStringify(value, nil); got != `{"value":3}` {
		t.Fatalf("carrier JSON: %s", got)
	}
	raw["self"] = value
	if got := *JsonStringify(value, nil); got != "null" {
		t.Fatalf("cycle JSON: %s", got)
	}
}
