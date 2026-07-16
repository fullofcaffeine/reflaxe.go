package hxrt

import "math"

// MathFloorInt returns Haxe Math.floor through Go's native int conversion.
//
// What: Produces the largest integer not greater than value.
// Why: Go math.Floor returns float64 while Haxe Math.floor returns Int.
// How: Apply math.Floor first, then convert the result to Go int.
func MathFloorInt(value float64) int {
	return int(math.Floor(value))
}

// MathCeilInt returns Haxe Math.ceil through Go's native int conversion.
//
// What: Produces the smallest integer not less than value.
// Why: Go math.Ceil returns float64 while Haxe Math.ceil returns Int.
// How: Apply math.Ceil first, then convert the result to Go int.
func MathCeilInt(value float64) int {
	return int(math.Ceil(value))
}

// MathRoundInt implements the Haxe 4.3.7 ties-up rounding rule.
//
// What: Rounds to the nearest Int with 0.5 moving toward positive infinity.
// Why: Go math.Round uses half-away-from-zero, which disagrees for negative ties.
// How: Compute floor(value + 0.5), then convert the result to Go int.
func MathRoundInt(value float64) int {
	return int(math.Floor(value + 0.5))
}
