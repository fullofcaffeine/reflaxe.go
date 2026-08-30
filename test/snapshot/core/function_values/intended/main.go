package main

import "snapshot/hxrt"

func main() {
	add := func(a int, b int) int {
		return int((hxrt.Int32Wrap(a) + hxrt.Int32Wrap(b)))
	}
	mul := func(v int) int {
		return int((hxrt.Int32Wrap(v) * hxrt.Int32Wrap(3)))
	}
	var v any = any(twice(5))
	hxrt.Println(v)
	var v_1 any = any(add(2, 7))
	hxrt.Println(v_1)
	var v_2 any = any(mul(4))
	hxrt.Println(v_2)
}

func twice(value int) int {
	return int((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(2)))
}
