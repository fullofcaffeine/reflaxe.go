package main

import "snapshot/hxrt"

func main() {
	var value any = nil
	hxrt.TryCatch(func() {
		typed := hxrt.IntFromNullableAny(value)
		hxrt.Println(any(int((hxrt.Int32Wrap(typed) + hxrt.Int32Wrap(1)))))
	}, func(hx_caught_1 any) {
		hx_tmp := hx_caught_1
		_ = hx_tmp
		hxrt.Println(any(hxrt.StringFromLiteral("caught-portable-runtime-failure")))
	})
}
