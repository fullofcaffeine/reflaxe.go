package hxrt

import (
	"math"
	"testing"
)

// flags renders four comparison results as "t"/"f" so a row reads like the
// Haxe expectation for `<`, `<=`, `>` and `>=`.
func flags(results ...bool) string {
	out := ""
	for _, result := range results {
		if result {
			out += "t"
		} else {
			out += "f"
		}
	}
	return out
}

func TestNullableIntWithIntFollowsHaxeNullRules(t *testing.T) {
	cases := []struct {
		name  string
		left  any
		right int
		want  string
	}{
		{"null", nil, 1, "ffff"},
		{"null zero", nil, 0, "ffff"},
		{"null negative", nil, -1, "ffff"},
		{"less", 1, 2, "ttff"},
		{"equal", 2, 2, "ftft"},
		{"greater", 3, 2, "fftt"},
		{"int64 storage", int64(5), 4, "fftt"},
	}
	for _, c := range cases {
		got := flags(NullableIntLess(c.left, c.right), NullableIntLessEqual(c.left, c.right),
			NullableIntGreater(c.left, c.right), NullableIntGreaterEqual(c.left, c.right))
		if got != c.want {
			t.Errorf("%s: Null<Int> %v op %d = %s, want %s", c.name, c.left, c.right, got, c.want)
		}
	}
}

func TestIntWithNullableIntFollowsHaxeNullRules(t *testing.T) {
	cases := []struct {
		name  string
		left  int
		right any
		want  string
	}{
		{"null", 1, nil, "ffff"},
		{"null zero", 0, nil, "ffff"},
		{"less", 1, 2, "ttff"},
		{"equal", 2, 2, "ftft"},
		{"greater", 3, int32(2), "fftt"},
	}
	for _, c := range cases {
		got := flags(IntLessNullable(c.left, c.right), IntLessEqualNullable(c.left, c.right),
			IntGreaterNullable(c.left, c.right), IntGreaterEqualNullable(c.left, c.right))
		if got != c.want {
			t.Errorf("%s: %d op Null<Int> %v = %s, want %s", c.name, c.left, c.right, got, c.want)
		}
	}
}

func TestNullableIntsTreatTwoNullsAsEqual(t *testing.T) {
	cases := []struct {
		name        string
		left, right any
		want        string
	}{
		{"both null", nil, nil, "ftft"},
		{"left null", nil, 0, "ffff"},
		{"right null", 0, nil, "ffff"},
		{"less", 1, 2, "ttff"},
		{"equal", 2, 2, "ftft"},
		{"greater", 2, 1, "fftt"},
	}
	for _, c := range cases {
		got := flags(NullableIntsLess(c.left, c.right), NullableIntsLessEqual(c.left, c.right),
			NullableIntsGreater(c.left, c.right), NullableIntsGreaterEqual(c.left, c.right))
		if got != c.want {
			t.Errorf("%s: Null<Int> %v op Null<Int> %v = %s, want %s", c.name, c.left, c.right, got, c.want)
		}
	}
}

func TestNullableFloatsTreatTwoNullsAsEqualAndNaNAsUnordered(t *testing.T) {
	cases := []struct {
		name        string
		left, right any
		want        string
	}{
		{"both null", nil, nil, "ftft"},
		{"left null", nil, 1.5, "ffff"},
		{"right null", 2, nil, "ffff"},
		{"less", 1.5, 2.5, "ttff"},
		{"int and float equal", 2, 2.0, "ftft"},
		{"float and int greater", 2.5, 2, "fftt"},
		{"nan", math.NaN(), math.NaN(), "ffff"},
		{"nan and null", math.NaN(), nil, "ffff"},
	}
	for _, c := range cases {
		got := flags(NullableFloatsLess(c.left, c.right), NullableFloatsLessEqual(c.left, c.right),
			NullableFloatsGreater(c.left, c.right), NullableFloatsGreaterEqual(c.left, c.right))
		if got != c.want {
			t.Errorf("%s: %v op %v = %s, want %s", c.name, c.left, c.right, got, c.want)
		}
	}
}
