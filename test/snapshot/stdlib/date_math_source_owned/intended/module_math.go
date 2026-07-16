package main

import (
	"math"
	"math/rand"
	"snapshot/hxrt"
)

var Math_NEGATIVE_INFINITY float64 = math.Inf(-1)

var Math_NaN float64 = math.NaN()

var Math_PI float64 = 3.141592653589793

var Math_POSITIVE_INFINITY float64 = math.Inf(1)

func Math_abs(v float64) float64 {
	return math.Abs(v)
}

func Math_acos(v float64) float64 {
	return math.Acos(v)
}

func Math_asin(v float64) float64 {
	return math.Asin(v)
}

func Math_atan(v float64) float64 {
	return math.Atan(v)
}

func Math_atan2(y float64, x float64) float64 {
	return math.Atan2(y, x)
}

func Math_ceil(v float64) int {
	return hxrt.MathCeilInt(v)
}

func Math_cos(v float64) float64 {
	return math.Cos(v)
}

func Math_exp(v float64) float64 {
	return math.Exp(v)
}

func Math_fceil(v float64) float64 {
	return math.Ceil(v)
}

func Math_ffloor(v float64) float64 {
	return math.Floor(v)
}

func Math_floor(v float64) int {
	return hxrt.MathFloorInt(v)
}

func Math_fround(v float64) float64 {
	return math.Floor((v + 0.5))
}

func Math_isFinite(f float64) bool {
	return (!math.IsInf(f, 0) && !math.IsNaN(f))
}

func Math_isNaN(f float64) bool {
	return math.IsNaN(f)
}

func Math_log(v float64) float64 {
	return math.Log(v)
}

func Math_max(a float64, b float64) float64 {
	if math.IsNaN(a) {
		return a
	}
	if math.IsNaN(b) {
		return b
	}
	var hx_if_5 float64
	if a < b {
		hx_if_5 = b
	} else {
		hx_if_5 = a
	}
	return hx_if_5
}

func Math_min(a float64, b float64) float64 {
	if math.IsNaN(a) {
		return a
	}
	if math.IsNaN(b) {
		return b
	}
	var hx_if_6 float64
	if a < b {
		hx_if_6 = a
	} else {
		hx_if_6 = b
	}
	return hx_if_6
}

func Math_pow(v float64, exp float64) float64 {
	return math.Pow(v, exp)
}

func Math_random() float64 {
	return rand.Float64()
}

func Math_round(v float64) int {
	return hxrt.MathRoundInt(v)
}

func Math_sin(v float64) float64 {
	return math.Sin(v)
}

func Math_sqrt(v float64) float64 {
	return math.Sqrt(v)
}

func Math_tan(v float64) float64 {
	return math.Tan(v)
}
