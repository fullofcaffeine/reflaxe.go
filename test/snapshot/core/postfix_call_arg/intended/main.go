package main

import "snapshot/hxrt"

func id(value int) int {
	return value
}

func main() {
	i := 1
	_ = i
	before := id(func() int {
		hx_post_1 := i
		i = int(int32((i + 1)))
		return hx_post_1
	}())
	_ = before
	hxrt.Println(hxrt.StdString(before))
	hxrt.Println(hxrt.StdString(i))
}
