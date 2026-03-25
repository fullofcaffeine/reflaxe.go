package main

func haxe__io__FPHelper_doubleToI64(v float64) *haxe___Int64_____Int64 {
	out := haxe__io__FPHelper_littleEndianOutput()
	out.writeDouble(v)
	input := haxe__io__FPHelper_littleEndianInput(out.getBytes())
	low := input.readInt32()
	high := input.readInt32()
	x := New_haxe___Int64_____Int64(high, low)
	var this1 *haxe___Int64_____Int64
	this1 = x
	return this1
}

func haxe__io__FPHelper_floatToI32(f float64) int {
	out := haxe__io__FPHelper_littleEndianOutput()
	out.writeFloat(f)
	return haxe__io__FPHelper_littleEndianInput(out.getBytes()).readInt32()
}

func haxe__io__FPHelper_i32ToFloat(i int) float64 {
	out := haxe__io__FPHelper_littleEndianOutput()
	out.writeInt32(i)
	return haxe__io__FPHelper_littleEndianInput(out.getBytes()).readFloat()
}

func haxe__io__FPHelper_i64ToDouble(low int, high int) float64 {
	out := haxe__io__FPHelper_littleEndianOutput()
	out.writeInt32(low)
	out.writeInt32(high)
	return haxe__io__FPHelper_littleEndianInput(out.getBytes()).readDouble()
}

func haxe__io__FPHelper_littleEndianInput(bytes *haxe__io__Bytes) *haxe__io__BytesInput {
	input := New_haxe__io__BytesInput(bytes)
	input.set_bigEndian(false)
	return input
}

func haxe__io__FPHelper_littleEndianOutput() *haxe__io__BytesOutput {
	out := New_haxe__io__BytesOutput()
	out.set_bigEndian(false)
	return out
}
