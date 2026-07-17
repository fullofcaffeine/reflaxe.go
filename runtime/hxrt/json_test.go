package hxrt

import "testing"

func TestJsonParseBuildsPortableArrayCarriers(t *testing.T) {
	decoded := JsonParse(StringFromLiteral(`{"items":[1,[2,3]]}`))
	object, ok := decoded.(map[string]any)
	if !ok {
		t.Fatalf("decoded root type = %T, want map[string]any", decoded)
	}

	items, ok := object["items"].(*Array)
	if !ok {
		t.Fatalf("decoded items type = %T, want *Array", object["items"])
	}
	if items.Len() != 2 {
		t.Fatalf("decoded items length = %d, want 2", items.Len())
	}
	nested, ok := items.Get(1).(*Array)
	if !ok {
		t.Fatalf("decoded nested type = %T, want *Array", items.Get(1))
	}
	if nested.Len() != 2 || nested.Get(0) != float64(2) || nested.Get(1) != float64(3) {
		t.Fatalf("decoded nested values = %#v, want [2 3]", nested.Values())
	}
}

func TestJsonStringifyTraversesPortableArrayCarriers(t *testing.T) {
	value := map[string]any{
		"items": NewArray(1, NewArray(2, 3)),
	}

	got := JsonStringify(value)
	if got == nil || *got != `{"items":[1,[2,3]]}` {
		t.Fatalf("JsonStringify() = %v, want nested JSON arrays", StdString(got))
	}
}

func TestJsonStringifyRejectsCyclicPortableArray(t *testing.T) {
	value := NewArray(nil)
	value.Set(0, value)

	got := JsonStringify(value)
	if got == nil || *got != "null" {
		t.Fatalf("JsonStringify(cycle) = %v, want null", StdString(got))
	}
}

func TestJsonStringifyAllowsRepeatedArrayOutsideActivePath(t *testing.T) {
	shared := NewArray(1)
	value := NewArray(shared, shared)

	got := JsonStringify(value)
	if got == nil || *got != `[[1],[1]]` {
		t.Fatalf("JsonStringify(repeated alias) = %v, want [[1],[1]]", StdString(got))
	}
}
