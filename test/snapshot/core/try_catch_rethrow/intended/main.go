package main

import "snapshot/hxrt"

func main() {
	hxrt.TryCatch(func() {
		hxrt.TryCatch(func() {
			hxrt.Throw(true)
		}, func(hx_caught_3 any) {
			switch hx_typed_4 := hx_caught_3.(type) {
			case int:
				i := hx_typed_4
				_ = i
				hxrt.Println(i)
			default:
				hxrt.Throw(hx_caught_3)
			}
		})
	}, func(hx_caught_1 any) {
		e := hx_caught_1
		_ = e
		hxrt.Println(hxrt.StringFromLiteral("outer"))
	})
}
