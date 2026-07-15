package main

import "snapshot/hxrt"

func main() {
	hxrt.TryCatch(func() {
		var v any = any(mayFail(false))
		hxrt.Println(v)
		var v_1 any = any(mayFail(true))
		hxrt.Println(v_1)
	}, func(hx_caught_1 any) {
		e := hx_caught_1
		hxrt.Println(e)
	})
	hxrt.Println(any(9))
}

func mayFail(flag bool) int {
	if flag {
		hxrt.Throw(hxrt.StringFromLiteral("boom"))
		var hx_throw_zero_3 int
		return hx_throw_zero_3
	}
	return 7
}
