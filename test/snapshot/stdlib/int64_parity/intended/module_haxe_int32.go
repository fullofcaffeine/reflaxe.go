package main

import "snapshot/hxrt"

func haxe___Int32__Int32_Impl__ucompare(a int, b int) int {
	if a < 0 {
		var hx_if_46 int
		if b < 0 {
			hx_if_46 = int(int32((hxrt.Int32Wrap(int(int32(^int32(b)))) - hxrt.Int32Wrap(int(int32(^int32(a)))))))
		} else {
			hx_if_46 = 1
		}
		return hx_if_46
	}
	var hx_if_47 int
	if b < 0 {
		hx_if_47 = -1
	} else {
		hx_if_47 = int(int32((hxrt.Int32Wrap(a) - hxrt.Int32Wrap(b))))
	}
	return hx_if_47
}
