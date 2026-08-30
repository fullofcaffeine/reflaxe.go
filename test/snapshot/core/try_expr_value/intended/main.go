package main

import "snapshot/hxrt"

func main() {
	var hx_try_1 int
	hxrt.TryCatch(func() {
		hx_try_1 = risky(0)
	}, func(hx_caught_2 any) {
		e := hx_caught_2
		_ = e
		hx_try_1 = 11
	})
	a := hx_try_1
	var hx_try_4 int
	hxrt.TryCatch(func() {
		hx_try_4 = risky(4)
	}, func(hx_caught_5 any) {
		e_1 := hx_caught_5
		_ = e_1
		hx_try_4 = 11
	})
	b := hx_try_4
	var v any = any(hxrt.StdString(a))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(b))
	hxrt.Println(v_1)
}

func risky(v int) int {
	if v == 0 {
		hxrt.Throw(hxrt.StringFromLiteral("bad"))
	}
	return int((hxrt.Int32Wrap(v) + hxrt.Int32Wrap(1)))
}
