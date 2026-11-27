package hxrt

import (
	"math"
	"testing"
)

// TestIsIntFollowsInterp checks Std.isOfType(value, Int) against `--interp`
// (haxe_go-jxxd): a Float counts as an Int when it is integral and within the
// 32-bit range, as on `--interp`, JS and Python.
func TestIsIntFollowsInterp(t *testing.T) {
	for _, test := range []struct {
		name  string
		value any
		want  bool
	}{
		{name: "int", value: 4, want: true},
		{name: "negative int", value: -7, want: true},
		{name: "int32", value: int32(3), want: true},
		{name: "uint8", value: uint8(3), want: true},
		{name: "4.0", value: 4.0, want: true},
		{name: "0.0", value: 0.0, want: true},
		{name: "-0.0", value: math.Copysign(0, -1), want: true},
		{name: "2147483647.0", value: 2147483647.0, want: true},
		{name: "-2147483648.0", value: -2147483648.0, want: true},
		{name: "float32 5", value: float32(5), want: true},
		{name: "2147483648.0", value: 2147483648.0, want: false},
		{name: "-2147483649.0", value: -2147483649.0, want: false},
		{name: "4294967296.0", value: 4294967296.0, want: false},
		{name: "1e20", value: 1e20, want: false},
		{name: "0.5", value: 0.5, want: false},
		{name: "-1.5", value: -1.5, want: false},
		{name: "float32 5.5", value: float32(5.5), want: false},
		{name: "NaN", value: math.NaN(), want: false},
		{name: "+Inf", value: math.Inf(1), want: false},
		{name: "-Inf", value: math.Inf(-1), want: false},
		{name: "null", value: nil, want: false},
		{name: "String", value: StringFromLiteral("4"), want: false},
		{name: "Bool", value: true, want: false},
	} {
		if got := IsInt(test.value); got != test.want {
			t.Errorf("IsInt(%s) = %v, want %v", test.name, got, test.want)
		}
		if value, ok := test.value.(float64); ok {
			if got := FloatIsInt(value); got != test.want {
				t.Errorf("FloatIsInt(%s) = %v, want %v", test.name, got, test.want)
			}
		}
	}
}
