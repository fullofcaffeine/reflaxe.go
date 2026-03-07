package main

import "snapshot/hxrt"

func main() {
	safe(hxrt.StringFromLiteral("fromCharCode"), func() *string {
		str := hxrt.StringFromLiteral("A")
		var this1_1 *string
		hxrt.Throw(hxrt.StringFromLiteral("Ucs2 String not supported on this platform"))
		var hx_throw_zero_1 *string
		return hx_throw_zero_1
		this1_1 = str
		this1 := hxrt.StdString(any(this1_1))
		return this1
	})
}

func safe(label *string, fn func() *string) {
	hxrt.TryCatch(func() {
		hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), fn()))
	}, func(hx_caught_2 any) {
		error := hx_caught_2
		hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=!")), hxrt.StdString(error)))
	})
}
