package main

import "snapshot/hxrt"

func id(value int) int {
	return value
}

func main() {
	i := 1
	before := id(func() int {
		hx_post_1 := i
		i = int(int32((i + 1)))
		return hx_post_1
	}())
	var v any = any(hxrt.StdString(before))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(i))
	hxrt.Println(v_1)
}
