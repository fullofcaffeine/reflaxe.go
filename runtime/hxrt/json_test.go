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
	if nested.Len() != 2 || nested.Get(0) != 2 || nested.Get(1) != 3 {
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

// TestJsonParseDecodesIntegralNumbersAsHaxeInt checks haxe.Json against
// `--interp` (haxe_go-2bge). Haxe's JsonParser returns an Int for a number
// written without a `.` whose value is integral and fits in 32 bits, and a
// Float for every other number. A typed `{count:Int}` read asserts a Go int.
func TestJsonParseDecodesIntegralNumbersAsHaxeInt(t *testing.T) {
	for _, test := range []struct {
		text string
		want any
	}{
		{text: "4", want: 4},
		{text: "0", want: 0},
		{text: "-5", want: -5},
		{text: "-0", want: 0},
		{text: "1e3", want: 1000},
		{text: "1E2", want: 100},
		{text: "1e+2", want: 100},
		{text: "10e-1", want: 1},
		{text: "1e-400", want: 0},
		{text: "2147483647", want: 2147483647},
		{text: "-2147483648", want: -2147483648},
		{text: "25e-1", want: 2.5},
		{text: "1.5", want: 1.5},
		{text: "4.0", want: 4.0},
		{text: "0.0", want: 0.0},
		{text: "-0.0", want: math.Copysign(0, -1)},
		{text: "2147483647.0", want: 2147483647.0},
		{text: "2147483648", want: 2147483648.0},
		{text: "-2147483649", want: -2147483649.0},
		{text: "4294967296", want: 4294967296.0},
		{text: "9007199254740993", want: 9007199254740992.0},
		{text: "12345678901234567890", want: 12345678901234567890.0},
		{text: "1e20", want: 1e20},
		{text: "1e400", want: math.Inf(1)},
		{text: "-1e400", want: math.Inf(-1)},
	} {
		t.Run(test.text, func(t *testing.T) {
			got := JsonParse(StringFromLiteral(test.text))
			if got != test.want {
				t.Fatalf("JsonParse(%s) = %#v (%T), want %#v (%T)", test.text, got, got, test.want, test.want)
			}
			if want, ok := test.want.(float64); ok && math.Signbit(want) != math.Signbit(got.(float64)) {
				t.Fatalf("JsonParse(%s) sign = %v, want %v", test.text, math.Signbit(got.(float64)), math.Signbit(want))
			}
		})
	}

	decoded, ok := JsonParse(StringFromLiteral(`{"count":4,"items":[1,2.5,{"n":-0}]}`)).(map[string]any)
	if !ok {
		t.Fatal("JsonParse(object) is not a map")
	}
	if decoded["count"] != 4 {
		t.Fatalf("count = %#v (%T), want int 4", decoded["count"], decoded["count"])
	}
	items := decoded["items"].(*Array)
	if items.Get(0) != 1 || items.Get(1) != 2.5 || items.Get(2).(map[string]any)["n"] != 0 {
		t.Fatalf("items = %#v, want [1 2.5 {n:0}]", items.Values())
	}
}

// TestJsonParseRejectsWhatUnmarshalRejects keeps the "null on invalid
// document" result while numbers are decoded from their text.
func TestJsonParseRejectsWhatUnmarshalRejects(t *testing.T) {
	for _, text := range []string{"", "   ", "1 2", "[1] x", "{\"a\":}", "01", "1.", ".5", "+1", "-", "1e", "NaN"} {
		if got := JsonParse(StringFromLiteral(text)); got != nil {
			t.Fatalf("JsonParse(%q) = %#v, want nil", text, got)
		}
	}
	if got := JsonParse(StringFromLiteral(" 7 \n")); got != 7 {
		t.Fatalf("JsonParse(padded) = %#v, want 7", got)
	}
}

// TestJsonStringifyRoundTripsParsedNumbers checks that a decoded Int prints
// as itself and a decoded Float keeps the text Haxe's JsonPrinter writes.
func TestJsonStringifyRoundTripsParsedNumbers(t *testing.T) {
	source := `[4,4.5,[1,2.0,-0,1e3,2147483648],{"e":-0.0},{"f":-7}]`
	if got, want := *JsonStringify(JsonParse(StringFromLiteral(source)), nil), `[4,4.5,[1,2,0,1000,2147483648],{"e":-0},{"f":-7}]`; got != want {
		t.Fatalf("JsonStringify(JsonParse(...)) = %s, want %s", got, want)
	}
}
