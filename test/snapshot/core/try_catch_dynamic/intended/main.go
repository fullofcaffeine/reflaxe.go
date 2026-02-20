package main

import "snapshot/hxrt"

func main() {
	hxrt.TryCatch(func() {
		hxrt.Println(mayFail(false))
		hxrt.Println(mayFail(true))
	}, func(hx_caught_1 any) {
		e := hx_caught_1
		_ = e
		hxrt.Println(e)
	})
	hxrt.Println(9)
}

func mayFail(flag bool) int {
	if flag {
		hxrt.Throw(hxrt.StringFromLiteral("boom"))
		var hx_throw_zero_3 int
		return hx_throw_zero_3
	}
	return 7
}
