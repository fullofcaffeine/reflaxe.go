package main

import "snapshot/hxrt"

func add(a int, b int) int {
	return int(int32((hxrt.Int32Wrap(a) + hxrt.Int32Wrap(b))))
}

func main() {
	local := func(v int) int {
		return int(int32((hxrt.Int32Wrap(v) + hxrt.Int32Wrap(1))))
	}
	var v any = any(add(1, 2))
	hxrt.Println(v)
	var v_1 any = any(add(5, 2))
	hxrt.Println(v_1)
	var v_2 any = any(add(5, 6))
	hxrt.Println(v_2)
	var v_3 any = any(local(10))
	hxrt.Println(v_3)
	var v_4 any = any(local(20))
	hxrt.Println(v_4)
}
