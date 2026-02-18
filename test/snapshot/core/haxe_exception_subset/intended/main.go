package main

import "snapshot/hxrt"

func fail() {
	hxrt.Throw(hxrt.StringFromLiteral("boom"))
}

func main() {
	hxrt.TryCatch(func() {
		fail()
	}, func(hx_caught_1 any) {
		e := hxrt.ExceptionCaught(hx_caught_1)
		_ = e
		hxrt.Println(hxrt.ExceptionMessage(e))
	})
}
