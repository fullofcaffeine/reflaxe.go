package main

import "snapshot/hxrt"

func failInt() int {
	hxrt.Throw(hxrt.StringFromLiteral("boom-int"))
	var hx_throw_zero_1 int
	return hx_throw_zero_1
}

func failString() *string {
	hxrt.Throw(hxrt.StringFromLiteral("boom-string"))
	var hx_throw_zero_2 *string
	return hx_throw_zero_2
}

func main() {
	var hx_try_3 int
	hxrt.TryCatch(func() {
		hx_try_3 = failInt()
	}, func(hx_caught_4 any) {
		hx_tmp := hx_caught_4
		_ = hx_tmp
		hx_try_3 = 7
	})
	a := hx_try_3
	_ = a
	var hx_try_6 *string
	hxrt.TryCatch(func() {
		hx_try_6 = failString()
	}, func(hx_caught_7 any) {
		hx_tmp_1 := hx_caught_7
		_ = hx_tmp_1
		hx_try_6 = hxrt.StringFromLiteral("ok")
	})
	b := hx_try_6
	_ = b
	hxrt.Println(hxrt.StdString(a))
	hxrt.Println(b)
}
