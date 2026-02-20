package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(sum([]int{1, 2, 3}))
	hxrt.Println(sum([]int{4}))
}

func sum(values []int) int {
	total := 0
	_ = total
	var _g_current int
	_ = _g_current
	var _g_args []int
	_ = _g_args
	_g_current = 0
	_g_args = values
	for _g_current < len(_g_args) {
		this1 := _g_args
		_ = this1
		hx_post_1 := _g_current
		_g_current = int(int32((_g_current + 1)))
		index := hx_post_1
		_ = index
		value := this1[index]
		_ = value
		total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(value))))
	}
	return total
}
