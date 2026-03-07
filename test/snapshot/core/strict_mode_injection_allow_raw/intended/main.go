package main

import "snapshot/hxrt"

func main() {
	base := 1
	_ = base
	direct := func(hx_value_1 any) int {
		if hx_value_1 == nil {
			var hx_zero_2 int
			return hx_zero_2
		}
		return hx_value_1.(int)
	}(2)
	interpolated := func(hx_value_3 any) int {
		if hx_value_3 == nil {
			var hx_zero_4 int
			return hx_zero_4
		}
		return hx_value_3.(int)
	}(base + 2)
	hxrt.Println(direct)
	hxrt.Println(interpolated)
}
