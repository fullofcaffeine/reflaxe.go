package main

import "snapshot/hxrt"

func haxe___Int32__Int32_Impl__ucompare(a int, b int) int {
	if a < 0 {
		var hx_if_1 int
		if b < 0 {
			hx_if_1 = int((hxrt.Int32Wrap(int(^int32(b))) - hxrt.Int32Wrap(int(^int32(a)))))
		} else {
			hx_if_1 = 1
		}
		return hx_if_1
	}
	var hx_if_2 int
	if b < 0 {
		hx_if_2 = -1
	} else {
		hx_if_2 = int((hxrt.Int32Wrap(a) - hxrt.Int32Wrap(b)))
	}
	return hx_if_2
}
