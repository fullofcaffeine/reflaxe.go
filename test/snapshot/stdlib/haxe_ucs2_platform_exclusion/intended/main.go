package main

import "snapshot/hxrt"

func main() {
	safe(hxrt.StringFromLiteral("fromCharCode"), func() *string {
		str := hxrt.StringFromLiteral("A")
		var this1_1 *string
		hxrt.Throw(hxrt.StringFromLiteral("Ucs2 String not supported on this platform"))
		this1_1 = str
		this1 := hxrt.StdString(any(this1_1))
		return this1
	})
}

func safe(label *string, fn func() *string) {
	hxrt.TryCatch(func() {
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), fn()))
		hxrt.Println(v_1)
	}, func(hx_caught_1 any) {
		error := hx_caught_1
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=!")), hxrt.StdString(error)))
		hxrt.Println(v)
	})
}
