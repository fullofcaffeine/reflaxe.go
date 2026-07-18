package hxrt

import "testing"

type reflectRuntimeFixture struct {
	Name  string
	Count int
}

func (fixture *reflectRuntimeFixture) Add(value int) int {
	return fixture.Count + value
}

func reflectTestName(value string) *string {
	return StringFromLiteral(value)
}

func TestReflectRuntimeFieldAndMethodLookup(t *testing.T) {
	fixture := &reflectRuntimeFixture{Name: "Ada", Count: 4}
	field := ReflectLookupField(fixture, reflectTestName("Name"))
	if field == nil || !field.Found || field.Value != "Ada" {
		t.Fatalf("unexpected field lookup: %#v", field)
	}
	missing := ReflectLookupField(fixture, reflectTestName("Missing"))
	if missing == nil || missing.Found {
		t.Fatalf("unexpected missing lookup: %#v", missing)
	}
	method := ReflectLookupMethod(fixture, reflectTestName("Add"))
	if method == nil || !method.Found {
		t.Fatalf("unexpected method lookup: %#v", method)
	}
	if got := ReflectCallMethod(method.Value, []any{3}); got != 7 {
		t.Fatalf("unexpected method result: %#v", got)
	}
}

func TestReflectRuntimeAnonymousObjectOperations(t *testing.T) {
	object := map[string]any{"name": "Ada", "count": 2}
	ReflectSetField(object, reflectTestName("name"), "Bea")
	if got := object["name"]; got != "Bea" {
		t.Fatalf("set field = %#v", got)
	}
	if !ReflectDeleteField(object, reflectTestName("count")) {
		t.Fatal("expected count deletion")
	}
	if _, found := object["count"]; found {
		t.Fatal("count remained after deletion")
	}
	copied, ok := ReflectCopy(object).(map[string]any)
	if !ok {
		t.Fatalf("copy type = %T", ReflectCopy(object))
	}
	copied["name"] = "Cleo"
	if object["name"] != "Bea" {
		t.Fatalf("copy mutated source: %#v", object)
	}
	fields := ReflectFields(object)
	if len(fields) != 1 || *fields[0] != "name" {
		t.Fatalf("fields = %#v", fields)
	}
}

func TestReflectRuntimeClassificationComparisonAndVarArgs(t *testing.T) {
	if ReflectCompare(2, 3) >= 0 || ReflectCompare(reflectTestName("z"), reflectTestName("a")) <= 0 {
		t.Fatal("unexpected comparison ordering")
	}
	callback := func(arguments *Array) any {
		total := 0
		for _, argument := range arguments.Values() {
			total += argument.(int)
		}
		return total
	}
	variadic, ok := ReflectMakeVarArgs(callback).(func(...any) any)
	if !ok {
		t.Fatalf("variadic type = %T", ReflectMakeVarArgs(callback))
	}
	if got := variadic(1, 2, 3, 4); got != 10 {
		t.Fatalf("variadic result = %#v", got)
	}
	if !ReflectIsFunction(callback) || !ReflectIsObject(map[string]any{}) || !ReflectIsObject(reflectTestName("text")) || ReflectIsObject("text") {
		t.Fatal("unexpected runtime classification")
	}
}
