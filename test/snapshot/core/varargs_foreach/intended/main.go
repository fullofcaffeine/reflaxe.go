package main

import "snapshot/hxrt"

func main() {
	var v any = any(sum([]int{1, 2, 3}))
	hxrt.Println(v)
	var v_1 any = any(sum([]int{4}))
	hxrt.Println(v_1)
}

func sum(values []int) int {
	total := 0
	var _g_current int
	var _g_args []int
	_g_current = 0
	_g_args = func(hx_value_1 any) []int {
		if hx_value_1 == nil {
			var hx_zero_2 []int
			return hx_zero_2
		}
		return hx_value_1.([]int)
	}(any(values))
	for _g_current < len(_g_args) {
		this1 := _g_args
		hx_post_3 := _g_current
		_g_current = int(int32((_g_current + 1)))
		index := hx_post_3
		value := this1[index]
		total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(value))))
	}
	return total
}
