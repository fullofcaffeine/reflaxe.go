package hxrt

import "testing"

func TestMathIntRoundingUsesHaxeRules(t *testing.T) {
	tests := []struct {
		value float64
		floor int
		ceil  int
		round int
	}{
		{value: 3.8, floor: 3, ceil: 4, round: 4},
		{value: -3.8, floor: -4, ceil: -3, round: -4},
		{value: 0.5, floor: 0, ceil: 1, round: 1},
		{value: -0.5, floor: -1, ceil: 0, round: 0},
		{value: -1.5, floor: -2, ceil: -1, round: -1},
	}
	for _, test := range tests {
		if got := MathFloorInt(test.value); got != test.floor {
			t.Errorf("MathFloorInt(%v) = %d, want %d", test.value, got, test.floor)
		}
		if got := MathCeilInt(test.value); got != test.ceil {
			t.Errorf("MathCeilInt(%v) = %d, want %d", test.value, got, test.ceil)
		}
		if got := MathRoundInt(test.value); got != test.round {
			t.Errorf("MathRoundInt(%v) = %d, want %d", test.value, got, test.round)
		}
	}
}
