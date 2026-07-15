package main

import "snapshot/hxrt"

func id(value int) int {
	return value
}

func main() {
	i := 1
	now := id(func() int {
		i = int(int32((i + 1)))
		return i
	}())
	var v any = any(hxrt.StdString(now))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(i))
	hxrt.Println(v_1)
}
