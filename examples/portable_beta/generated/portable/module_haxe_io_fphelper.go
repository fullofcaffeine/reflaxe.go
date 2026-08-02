package main

import "examples_portable_beta/hxrt"

func haxe__io__FPHelper_doubleToI64(v float64) *haxe___Int64_____Int64 {
	high := hxrt.Float64HighWord(v)
	low := hxrt.Float64LowWord(v)
	x := New_haxe___Int64_____Int64(high, low)
	var this1 *haxe___Int64_____Int64
	this1 = x
	return this1
}

func haxe__io__FPHelper_floatToI32(f float64) int {
	return hxrt.Float32Bits(f)
}

func haxe__io__FPHelper_i32ToFloat(i int) float64 {
	return hxrt.Float32FromBits(i)
}

func haxe__io__FPHelper_i64ToDouble(low int, high int) float64 {
	return hxrt.Float64FromWords(low, high)
}
