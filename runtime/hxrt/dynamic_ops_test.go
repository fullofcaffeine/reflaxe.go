package hxrt

import (
	"math"
	"testing"
)

// dynamicOutcome renders one helper result, or the message of the Haxe error
// that it raises, with a type tag: `I:` for Int, `F:` for Float, `S:` for
// String and `B:` for Bool.
func dynamicOutcome(run func() any) (outcome string) {
	defer func() {
		if recovered := recover(); recovered != nil {
			failure, ok := recovered.(HaxeException)
			if !ok {
				panic(recovered)
			}
			outcome = "error: " + *StdString(failure.Value)
		}
	}()
	value := run()
	switch value.(type) {
	case int:
		return "I:" + *StdString(value)
	case float64:
		return "F:" + *StdString(value)
	case *string:
		return "S:" + *StdString(value)
	case bool:
		return "B:" + *StdString(value)
	case nil:
		return "null"
	default:
		return "?:" + *StdString(value)
	}
}

// The rows are `--interp` results for the same operands.
func TestDynamicArithmeticFollowsInterp(t *testing.T) {
	text := StringFromLiteral("x")
	cases := []struct {
		name string
		run  func() any
		want string
	}{
		{"int + int", func() any { return DynamicAdd(7, 2) }, "I:9"},
		{"int + int wraps", func() any { return DynamicAdd(2147483647, 1) }, "I:-2147483648"},
		{"int + float", func() any { return DynamicAdd(7, 0.5) }, "F:7.5"},
		{"float + int", func() any { return DynamicAdd(2.5, 2) }, "F:4.5"},
		{"int + string", func() any { return DynamicAdd(7, text) }, "S:7x"},
		{"string + int", func() any { return DynamicAdd(text, 7) }, "S:x7"},
		{"null + string", func() any { return DynamicAdd(nil, text) }, "S:nullx"},
		{"string + bool", func() any { return DynamicAdd(text, false) }, "S:xfalse"},
		{"int + null", func() any { return DynamicAdd(7, nil) }, "error: Invalid operation: 7 + null"},
		{"bool + int", func() any { return DynamicAdd(true, 1) }, "error: Invalid operation: true + 1"},
		{"int - int wraps", func() any { return DynamicSub(-2147483648, 1) }, "I:2147483647"},
		{"string - int", func() any { return DynamicSub(text, 1) }, "error: Invalid operation: x - 1"},
		{"int * int wraps", func() any { return DynamicMul(65536, 65536) }, "I:0"},
		{"int * float", func() any { return DynamicMul(7, 0.5) }, "F:3.5"},
		{"int / int", func() any { return DynamicDiv(7, 2) }, "F:3.5"},
		{"null / int", func() any { return DynamicDiv(nil, 2) }, "error: Invalid operation: null / 2"},
		{"int % int", func() any { return DynamicMod(-7, 2) }, "I:-1"},
		{"min % -1", func() any { return DynamicMod(-2147483648, -1) }, "I:0"},
		{"float % int", func() any { return DynamicMod(-7.5, 2) }, "F:-1.5"},
		{"int & int", func() any { return DynamicAnd(7, 3) }, "I:3"},
		{"int | int", func() any { return DynamicOr(4, 3) }, "I:7"},
		{"int ^ int", func() any { return DynamicXor(7, 2) }, "I:5"},
		{"whole float & int", func() any { return DynamicAnd(2.0, 1) }, "error: Invalid operation: 2 & 1"},
		{"int << 33", func() any { return DynamicShl(1, 33) }, "I:2"},
		{"max << 1", func() any { return DynamicShl(2147483647, 1) }, "I:-2"},
		{"-8 >> 1", func() any { return DynamicShr(-8, 1) }, "I:-4"},
		{"-8 >>> 1", func() any { return DynamicUShr(-8, 1) }, "I:2147483644"},
		{"min >>> 0", func() any { return DynamicUShr(-2147483648, 0) }, "I:-2147483648"},
		{"-int", func() any { return DynamicNeg(7) }, "I:-7"},
		{"-min", func() any { return DynamicNeg(-2147483648) }, "I:-2147483648"},
		{"-float", func() any { return DynamicNeg(2.5) }, "F:-2.5"},
		{"-null", func() any { return DynamicNeg(nil) }, "error: Invalid operation"},
		{"~int", func() any { return DynamicComplement(7) }, "I:-8"},
		{"~float", func() any { return DynamicComplement(2.5) }, "F:-3.5"},
		{"~string", func() any { return DynamicComplement(text) }, "error: Invalid operation: -1 - x"},
		{"int64 storage", func() any { return DynamicAdd(int64(3), int32(4)) }, "I:7"},
	}
	for _, c := range cases {
		if got := dynamicOutcome(c.run); got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
	if !math.IsNaN(DynamicMod(5, 0).(float64)) {
		t.Errorf("int %% 0 is NaN")
	}
	if !math.IsInf(DynamicDiv(1, 0), 1) || !math.IsInf(DynamicDiv(-1, 0), -1) {
		t.Errorf("int / 0 is an infinity")
	}
}

func TestDynamicTruthAndOrder(t *testing.T) {
	text := StringFromLiteral("x")
	var nilText *string
	truth := []struct {
		value     any
		condition bool
		not       bool
	}{
		{true, true, false},
		{false, false, true},
		{nil, false, true},
		{nilText, false, true},
		{0, false, false},
		{1, false, false},
		{StringFromLiteral(""), false, false},
		{text, false, false},
	}
	for _, c := range truth {
		if DynamicIsTrue(c.value) != c.condition || DynamicNot(c.value) != c.not {
			t.Errorf("%v: condition %v not %v", c.value, DynamicIsTrue(c.value), DynamicNot(c.value))
		}
	}

	object := StringFromLiteral("same")
	order := []struct {
		name                   string
		left, right            any
		less, lessEqual, great bool
	}{
		{"int < int", 3, 4, true, true, false},
		{"float vs int", 2.0, 2, false, true, false},
		{"strings", StringFromLiteral("10"), StringFromLiteral("9"), true, true, false},
		{"bools", false, true, true, true, false},
		{"string vs int", text, 1, false, false, false},
		{"null vs int", nil, 1, false, false, false},
		{"null vs null", nil, nil, false, true, false},
		{"NaN", math.NaN(), 1, false, false, false},
		{"one string", object, object, false, true, false},
	}
	for _, c := range order {
		if DynamicLess(c.left, c.right) != c.less || DynamicLessEqual(c.left, c.right) != c.lessEqual || DynamicGreater(c.left, c.right) != c.great {
			t.Errorf("%s: < %v <= %v > %v", c.name, DynamicLess(c.left, c.right), DynamicLessEqual(c.left, c.right), DynamicGreater(c.left, c.right))
		}
	}
	if !DynamicGreaterEqual(nil, nil) || DynamicGreaterEqual(7, nil) {
		t.Errorf("null >= null is true and 7 >= null is false")
	}
}

func TestDynamicToFloatReadsNumbers(t *testing.T) {
	if DynamicToFloat(3) != 3 || DynamicToFloat(2.5) != 2.5 {
		t.Errorf("numbers convert")
	}
	if got := dynamicOutcome(func() any { return DynamicToFloat(nil) }); got != "error: Invalid operation: null" {
		t.Errorf("null: %q", got)
	}
}

// The rows are `--interp` results for the same targets and indexes.
func TestDynamicIndexFollowsInterp(t *testing.T) {
	array := NewArray(10, 20, 30)
	var nilArray *Array
	cases := []struct {
		name string
		run  func() any
		want string
	}{
		{"read", func() any { return DynamicIndexGet(array, 1) }, "I:20"},
		{"read past the end", func() any { return DynamicIndexGet(array, 3) }, "null"},
		{"read negative", func() any { return DynamicIndexGet(array, -1) }, "null"},
		{"write", func() any { return DynamicIndexSet(array, 0, 5) }, "I:5"},
		{"write past the end", func() any { return DynamicIndexSet(array, 4, 7) }, "I:7"},
		{"string", func() any { return DynamicIndexGet(StringFromLiteral("hello"), 0) }, "error: Unexpected value hello, expected Array"},
		{"string write", func() any { return DynamicIndexSet(StringFromLiteral("hello"), 0, 1) }, "error: Unexpected value hello, expected Array"},
		{"int", func() any { return DynamicIndexGet(5, 0) }, "error: Unexpected value 5, expected Array"},
		{"bool", func() any { return DynamicIndexGet(true, 0) }, "error: Unexpected value true, expected Array"},
		{"null", func() any { return DynamicIndexGet(nil, 0) }, "error: Null Access"},
		{"typed null", func() any { return DynamicIndexSet(nilArray, 0, 1) }, "error: Null Access"},
	}
	for _, c := range cases {
		if got := dynamicOutcome(c.run); got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
	if got := *StdString(array); got != "[5,20,30,null,7]" {
		t.Errorf("writes: got %s", got)
	}
}
