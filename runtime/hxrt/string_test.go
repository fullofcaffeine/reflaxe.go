package hxrt

import (
	"math"
	"math/bits"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestStringSliceCodePointsStringPtrUsesLogicalRuneBounds(t *testing.T) {
	value := "a😀bé"
	got := StringSliceCodePointsStringPtr(&value, 1, 3)
	if got == nil || *got != "😀b" {
		t.Fatalf("StringSliceCodePointsStringPtr() = %v, want %q", got, "😀b")
	}
}

func TestStringEqualStringPtrDistinguishesNullFromLiteral(t *testing.T) {
	nullLiteral := "null"
	if StringEqualStringPtr(nil, &nullLiteral) {
		t.Fatal("StringEqualStringPtr treated null as the literal string \"null\"")
	}
	if !StringEqualStringPtr(nil, nil) {
		t.Fatal("StringEqualStringPtr rejected two null strings")
	}
}

func TestStringCompareStringPtrOrdersValuesAndNull(t *testing.T) {
	alpha := "alpha"
	beta := "beta"
	if StringCompareStringPtr(&alpha, &beta) >= 0 {
		t.Fatal("alpha must sort before beta")
	}
	if StringCompareStringPtr(&beta, &alpha) <= 0 {
		t.Fatal("beta must sort after alpha")
	}
	if StringCompareStringPtr(&alpha, &alpha) != 0 {
		t.Fatal("equal string values must compare equally")
	}
	if StringCompareStringPtr(nil, &alpha) >= 0 || StringCompareStringPtr(&alpha, nil) <= 0 || StringCompareStringPtr(nil, nil) != 0 {
		t.Fatal("null string ordering must be deterministic and distinct")
	}
}

func TestStringIndexOfStringPtrUsesLogicalRuneIndexes(t *testing.T) {
	value := "a😀bé😀"
	needle := "😀"
	empty := ""
	missing := "missing"

	tests := []struct {
		name     string
		search   *string
		start    int
		hasStart bool
		want     int
	}{
		{name: "first unicode match", search: &needle, want: 1},
		{name: "match after start", search: &needle, start: 2, hasStart: true, want: 4},
		{name: "empty at start", search: &empty, start: 3, hasStart: true, want: 3},
		{name: "empty clamps past end", search: &empty, start: 99, hasStart: true, want: 5},
		{name: "negative start is relative to end", search: &needle, start: -1, hasStart: true, want: 4},
		{name: "empty clamps negative start to zero", search: &empty, start: -1, hasStart: true, want: 0},
		{name: "large negative start clamps to zero", search: &needle, start: -99, hasStart: true, want: 1},
		{name: "missing", search: &missing, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := StringIndexOfStringPtr(&value, test.search, test.start, test.hasStart); got != test.want {
				t.Fatalf("StringIndexOfStringPtr() = %d, want %d", got, test.want)
			}
		})
	}
}

func TestStringParseFloatExactAcceptsOnlyValidatedCompleteTokens(t *testing.T) {
	valid := "-125.5e-1"
	if got := StringParseFloatExact(&valid); got != -12.55 {
		t.Fatalf("StringParseFloatExact(%q) = %v, want -12.55", valid, got)
	}

	invalid := "12tail"
	if got := StringParseFloatExact(&invalid); !math.IsNaN(got) {
		t.Fatalf("StringParseFloatExact(%q) = %v, want NaN", invalid, got)
	}
	if got := StringParseFloatExact(nil); !math.IsNaN(got) {
		t.Fatalf("StringParseFloatExact(nil) = %v, want NaN", got)
	}
}

// stdStringTestEnum mirrors a generated enum carrier and its HxEnumValue method.
type stdStringTestEnum struct {
	tag    int
	params []any
}

func (self *stdStringTestEnum) HxEnumValue() (string, []any) {
	return [2]string{"None", "Some"}[self.tag], self.params
}

// stdStringTestClass mirrors a generated class without toString().
type stdStringTestClass struct{}

func (self *stdStringTestClass) HxClassName() string { return "pack.Point" }

// stdStringTestNamed mirrors a generated class whose toString() adapter is String().
type stdStringTestNamed struct {
	stdStringTestClass
}

func (self *stdStringTestNamed) String() string { return "Named(n)" }

func TestStdStringPrintsEnumsClassesAndObjectsLikeHaxe(t *testing.T) {
	none := &stdStringTestEnum{tag: 0}
	text := "hello"
	tests := []struct {
		name  string
		value any
		want  string
	}{
		{name: "enum without arguments", value: none, want: "None"},
		{name: "enum with arguments", value: &stdStringTestEnum{tag: 1, params: []any{1}}, want: "Some(1)"},
		{name: "nested enum and string", value: &stdStringTestEnum{tag: 1, params: []any{&stdStringTestEnum{tag: 1, params: []any{&text}}}}, want: "Some(Some(hello))"},
		{name: "enum with null argument", value: &stdStringTestEnum{tag: 1, params: []any{nil}}, want: "Some(null)"},
		{name: "typed nil enum", value: (*stdStringTestEnum)(nil), want: "null"},
		{name: "class without toString", value: &stdStringTestClass{}, want: "pack.Point"},
		{name: "toString wins over an embedded class name", value: &stdStringTestNamed{}, want: "Named(n)"},
		{name: "anonymous object", value: map[string]any{"b": &text, "a": 1}, want: "{a: 1, b: hello}"},
		{name: "empty anonymous object", value: map[string]any{}, want: "{}"},
		{name: "anonymous view", value: AnonymousFromMap(map[string]any{"k": none}), want: "{k: None}"},
		{name: "array of shapes", value: NewArray(none, map[string]any{"a": 1}, &stdStringTestClass{}), want: "[None,{a: 1},pack.Point]"},
		{name: "function inside a value", value: NewArray(NewFunction(func() {})), want: "[#fun]"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := *StdString(test.value); got != test.want {
				t.Fatalf("StdString() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestStdStringCallsAnonymousToStringField(t *testing.T) {
	object := map[string]any{"toString": NewFunction(func() *string { return StringFromLiteral("custom") }), "a": 1}
	if got := *StdString(object); got != "custom" {
		t.Fatalf("StdString() = %q, want custom", got)
	}
	notCallable := map[string]any{"toString": NewFunction(func(int) *string { return nil })}
	if got := *StdString(notCallable); got != "{toString: #fun}" {
		t.Fatalf("StdString() = %q, want the object fields", got)
	}
}

// The Haxe interpreter orders anonymous fields by the bit-reversed OCaml
// Hashtbl.hash of each name. The expected orders below come from `haxe --interp`.
func TestStdStringOrdersAnonymousFieldsLikeTheInterpreter(t *testing.T) {
	tests := []struct {
		names []string
		want  string
	}{
		{names: []string{"b", "a"}, want: "{a: 0, b: 0}"},
		{names: []string{"mid", "alpha", "zeta"}, want: "{zeta: 0, alpha: 0, mid: 0}"},
		{names: []string{"s", "f", "inner", "nul", "e"}, want: "{e: 0, nul: 0, inner: 0, f: 0, s: 0}"},
		{names: []string{"quoted key", "normal"}, want: "{normal: 0, quoted key: 0}"},
	}
	for _, test := range tests {
		object := map[string]any{}
		for _, name := range test.names {
			object[name] = 0
		}
		if got := *StdString(object); got != test.want {
			t.Fatalf("StdString(%v) = %q, want %q", test.names, got, test.want)
		}
	}
	// Known OCaml 5 `Hashtbl.hash` values, before bit reversal.
	for name, hash := range map[string]uint32{"a": 721651713, "zeta": 762974764, "inner": 979392339, "nul": 573017198} {
		if got, want := haxeInterpFieldOrderKey(name), bits.Reverse32(hash); got != want {
			t.Fatalf("haxeInterpFieldOrderKey(%q) = %d, want %d", name, got, want)
		}
	}
}

func TestStdStringStopsAtTheInterpreterPrintDepth(t *testing.T) {
	cyclic := NewArray()
	cyclic.Push(cyclic)
	if got, want := *StdString(cyclic), "[[[[[[<...>]]]]]]"; got != want {
		t.Fatalf("StdString(cyclic Array) = %q, want %q", got, want)
	}
	object := map[string]any{}
	object["self"] = object
	if got, want := *StdString(object), "{self: {self: {self: {self: {self: {self: <...>}}}}}}"; got != want {
		t.Fatalf("StdString(cyclic object) = %q, want %q", got, want)
	}
	deep := any(&stdStringTestEnum{tag: 0})
	for level := 0; level < 6; level++ {
		deep = &stdStringTestEnum{tag: 1, params: []any{deep}}
	}
	if got, want := *StdString(deep), "Some(Some(Some(Some(Some(Some(<...>))))))"; got != want {
		t.Fatalf("StdString(deep enum) = %q, want %q", got, want)
	}
	toString := NewArray(NewArray(NewArray(NewArray(NewArray(&stdStringTestNamed{})))))
	if got, want := *StdString(toString), "[[[[[<...>]]]]]"; got != want {
		t.Fatalf("StdString(deep toString) = %q, want %q", got, want)
	}
}

func TestSortAnonymousFieldOrderMatchesStandardSort(t *testing.T) {
	for _, count := range []int{0, 1, 2, 3, 17, 500} {
		fields := make([]anonymousFieldOrder, 0, count)
		for index := 0; index < count; index++ {
			name := "field" + strconv.Itoa(index*7919%1000)
			fields = append(fields, anonymousFieldOrder{key: haxeInterpFieldOrderKey(name) % 64, name: name})
		}
		want := append([]anonymousFieldOrder(nil), fields...)
		sort.Slice(want, func(left, right int) bool {
			return want[left].key < want[right].key || want[left].key == want[right].key && want[left].name < want[right].name
		})
		sortAnonymousFieldOrder(fields)
		for index := range want {
			if fields[index] != want[index] {
				t.Fatalf("count %d: position %d = %v, want %v", count, index, fields[index], want[index])
			}
		}
	}
}

// TestStdStringFloatMatchesInterp checks Float text against `haxe --interp`
// output for the same values (haxe_go-rlvo).
func TestStdStringFloatMatchesInterp(t *testing.T) {
	rows := []struct {
		value float64
		want  string
	}{
		{0, "0"},
		{math.Copysign(0, -1), "-0"},
		{1, "1"},
		{-1, "-1"},
		{0.5, "0.5"},
		{3.14159, "3.14159"},
		{0.1, "0.1"},
		{0.0001, "0.0001"},
		{0.000123, "0.000123"},
		{1e-5, "1e-05"},
		{1e-7, "1e-07"},
		{1.5e-7, "1.5e-07"},
		{1.23e-18, "1.23e-18"},
		{1e5, "100000"},
		{1e6, "1000000"},
		{12345678.9, "12345678.9"},
		{1e8, "100000000"},
		{-1e8, "-100000000"},
		{2147483647, "2147483647"},
		{-2147483648, "-2147483648"},
		{4294967296, "4294967296"},
		{1e9 + 0.5, "1000000000.5"},
		{123456789.123, "123456789.123"},
		{100000000000.5, "100000000000.5"},
		{999999999999, "999999999999"},
		{1e12, "1e+12"},
		{-1e12, "-1e+12"},
		{1e12 + 1, "1000000000001"},
		{2e12, "2e+12"},
		{1234567890000, "1.23456789e+12"},
		{99999999999999, "99999999999999"},
		{1e14 + 0.25, "100000000000000.25"},
		{999999999999999, "999999999999999"},
		{1e15, "1e+15"},
		{1e15 + 0.5, "1000000000000000.5"},
		{9007199254740992, "9007199254740992"},
		{1e16, "1e+16"},
		{123456789012345678, "123456789012345680"},
		{1234567890123456789, "1.23456789012345677e+18"},
		{1e21, "1e+21"},
		{-1e21, "-1e+21"},
		{1e100, "1e+100"},
		{1.5e300, "1.5e+300"},
		{math.MaxFloat64, "1.79769313486231571e+308"},
		{2.2250738585072014e-308, "2.22507385850720138e-308"},
		{5e-324, "4.94065645841e-324"},
		{1e-300, "1e-300"},
		{0.30000000000000004, "0.300000000000000044"},
		{1.0 / 3, "0.333333333333333315"},
		{-2.0 / 3, "-0.66666666666666663"},
		{math.NaN(), "nan"},
		{math.Inf(1), "infinity"},
		{math.Inf(-1), "neg_infinity"},
	}
	for _, row := range rows {
		if got := *StdString(row.value); got != row.want {
			t.Errorf("Std.string(%v) = %q, want %q", row.value, got, row.want)
		}
	}
	array := NewArray(1e8, 0.30000000000000004, 1e-7)
	if got, want := *StdString(array), "[100000000,0.300000000000000044,1e-07]"; got != want {
		t.Errorf("Std.string(Array<Float>) = %q, want %q", got, want)
	}
}

// TestStdStringFloat32KeepsItsShortestDigits covers native Go float32
// values, which Haxe does not have: they use the same layout with their own
// shortest digits, so float32(0.1) stays `0.1`.
func TestStdStringFloat32KeepsItsShortestDigits(t *testing.T) {
	rows := []struct {
		value float32
		want  string
	}{
		{0.1, "0.1"},
		{1e8, "100000000"},
		{16777217, "16777216"},
		{1.5e-7, "1.5e-07"},
		{3e38, "3e+38"},
	}
	for _, row := range rows {
		if got := *StdString(row.value); got != row.want {
			t.Errorf("Std.string(float32(%v)) = %q, want %q", row.value, got, row.want)
		}
	}
}

// haxeFloatReference is Haxe's own rule, written as plainly as possible: the
// interpreter prints `Numeric.float_repres`, which keeps the first of C's
// `%.12g` and `%.15g` that reads back as the same Float, else `%.18g`. Go's
// `g` format with a precision lays digits out as C's `%g` does.
func haxeFloatReference(value float64) string {
	switch {
	case math.IsNaN(value):
		return "nan"
	case math.IsInf(value, 1):
		return "infinity"
	case math.IsInf(value, -1):
		return "neg_infinity"
	}
	for _, precision := range []int{12, 15} {
		text := strconv.FormatFloat(value, 'g', precision, 64)
		if parsed, _ := strconv.ParseFloat(text, 64); parsed == value {
			return text
		}
	}
	return strconv.FormatFloat(value, 'g', 18, 64)
}

// TestStdStringFloatMatchesReferenceRule compares the runtime with
// haxeFloatReference on random bit patterns, random whole numbers, random
// short decimals, and the neighbours of every power of ten.
func TestStdStringFloatMatchesReferenceRule(t *testing.T) {
	random := rand.New(rand.NewSource(20261123))
	var values []float64
	for index := 0; index < 100000; index++ {
		values = append(values, math.Float64frombits(random.Uint64()))
		values = append(values, float64(random.Int63n(1<<53))*math.Copysign(1, random.Float64()-0.5))
		digits := 1 + random.Intn(17)
		decimal := strconv.FormatInt(random.Int63n(int64(math.Pow10(digits))), 10) + "e" + strconv.Itoa(random.Intn(60)-30)
		parsed, _ := strconv.ParseFloat(decimal, 64)
		values = append(values, parsed)
	}
	for exponent := -323; exponent <= 308; exponent++ {
		power := math.Pow10(exponent)
		values = append(values, power, math.Nextafter(power, 0), math.Nextafter(power, math.Inf(1)), 1.5*power, -power)
	}
	for _, value := range values {
		got, _ := scalarHaxeString(value)
		if want := haxeFloatReference(value); got != want {
			t.Fatalf("Std.string(%b) = %q, want %q", value, got, want)
		}
	}
}

// haxeFloatBenchmarkValues is a representative mix of the Floats that string
// building meets: whole numbers, short decimals, results with 16 or 17
// significant digits, and values printed in exponent form.
var haxeFloatBenchmarkValues = map[string][]any{
	"whole":    {0.0, 1.0, 42.0, -7.0, 2147483647.0, 1e8},
	"short":    {0.5, 3.14159, 0.1, 123456789.123, 12345678.9, -2.75},
	"long":     {0.30000000000000004, 1.0 / 3, 2.0 / 3, math.Pi, math.Sqrt2, 1e15 + 0.5},
	"exponent": {1e15, 1e21, 1e-7, 2.5e-5, 1.5e300, 6.02214076e23},
}

var haxeFloatBenchmarkSink string

func BenchmarkScalarHaxeStringFloat(b *testing.B) {
	for _, name := range []string{"whole", "short", "long", "exponent"} {
		values := haxeFloatBenchmarkValues[name]
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				for _, value := range values {
					haxeFloatBenchmarkSink, _ = scalarHaxeString(value)
				}
			}
		})
	}
}
