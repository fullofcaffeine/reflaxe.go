package hxrt

import (
	"math"
	"testing"
)

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

func TestJsonParseBoxesStringsForTypedDynamicCasts(t *testing.T) {
	decoded := JsonParse(StringFromLiteral(`{"name":"alpha","nested":{"name":"beta"}}`))
	object, ok := decoded.(map[string]any)
	if !ok {
		t.Fatalf("decoded root type = %T, want map[string]any", decoded)
	}

	name, ok := object["name"].(*string)
	if !ok || name == nil || *name != "alpha" {
		t.Fatalf("decoded name = %#v (%T), want *string(alpha)", object["name"], object["name"])
	}
	nested, ok := object["nested"].(map[string]any)
	if !ok {
		t.Fatalf("decoded nested type = %T, want map[string]any", object["nested"])
	}
	nestedName, ok := nested["name"].(*string)
	if !ok || nestedName == nil || *nestedName != "beta" {
		t.Fatalf("decoded nested name = %#v (%T), want *string(beta)", nested["name"], nested["name"])
	}
}

func TestJsonStringifyTraversesPortableArrayCarriers(t *testing.T) {
	value := map[string]any{
		"items": NewArray(1, NewArray(2, 3)),
	}

	got := JsonStringify(value, nil)
	if got == nil || *got != `{"items":[1,[2,3]]}` {
		t.Fatalf("JsonStringify() = %v, want nested JSON arrays", StdString(got))
	}
}

func TestJsonStringifyRejectsCyclicPortableArray(t *testing.T) {
	value := NewArray(nil)
	value.Set(0, value)

	got := JsonStringify(value, nil)
	if got == nil || *got != "null" {
		t.Fatalf("JsonStringify(cycle) = %v, want null", StdString(got))
	}
}

func TestJsonStringifyAllowsRepeatedArrayOutsideActivePath(t *testing.T) {
	shared := NewArray(1)
	value := NewArray(shared, shared)

	got := JsonStringify(value, nil)
	if got == nil || *got != `[[1],[1]]` {
		t.Fatalf("JsonStringify(repeated alias) = %v, want [[1],[1]]", StdString(got))
	}
}

func TestJsonStringifyHonorsPresentIndentation(t *testing.T) {
	value := map[string]any{"items": NewArray(1, 2)}

	for _, test := range []struct {
		name  string
		space *string
		want  string
	}{
		{name: "omitted", space: nil, want: `{"items":[1,2]}`},
		{name: "empty", space: StringFromLiteral(""), want: `{
"items": [
1,
2
]
}`},
		{name: "two spaces", space: StringFromLiteral("  "), want: `{
  "items": [
    1,
    2
  ]
}`},
		{name: "tab", space: StringFromLiteral("\t"), want: "{\n\t\"items\": [\n\t\t1,\n\t\t2\n\t]\n}"},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := JsonStringify(value, test.space)
			if got == nil {
				t.Fatal("JsonStringify() = nil")
			}
			if *got != test.want {
				t.Fatalf("JsonStringify() = %q, want %q", *got, test.want)
			}
		})
	}
}

// TestJsonStringifyWritesFloatsAsHaxe checks haxe.Json against `--interp`:
// Haxe's JsonPrinter writes a finite Float as Std.string does and any other
// Float as null, without dropping the rest of the document (haxe_go-rlvo).
func TestJsonStringifyWritesFloatsAsHaxe(t *testing.T) {
	values := NewArray(1e8, 2147483647.0, 0.30000000000000004, 1e15, 1e21, 1e-7, math.Copysign(0, -1), 1.5, math.NaN(), math.Inf(1))
	if got, want := *JsonStringify(values, nil), "[100000000,2147483647,0.300000000000000044,1e+15,1e+21,1e-07,-0,1.5,null,null]"; got != want {
		t.Fatalf("JsonStringify(floats) = %s, want %s", got, want)
	}
	if got, want := *JsonStringify(map[string]any{"big": 1e8, "count": 3}, nil), `{"big":100000000,"count":3}`; got != want {
		t.Fatalf("JsonStringify(object) = %s, want %s", got, want)
	}
}
