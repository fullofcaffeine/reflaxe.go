package main

import "snapshot/hxrt"

func id(value int) int {
	return value
}

func main() {
	cond := true
	_ = cond
	hxrt.Println(hxrt.StdString(id(func() int {
		var hx_if_1 int
		if cond {
			hx_if_1 = 7
		} else {
			hx_if_1 = 9
		}
		return hx_if_1
	}())))
	cond = false
	hxrt.Println(hxrt.StdString(id(func() int {
		var hx_if_2 int
		if cond {
			hx_if_2 = 7
		} else {
			hx_if_2 = 9
		}
		return hx_if_2
	}())))
}
