package main

import "snapshot/hxrt"

func main() {
	v := 1
	_ = v
	switch v {
	case 0:
		hxrt.Println(0)
	case 1:
		hxrt.Println(1)
	default:
		hxrt.Println(9)
	}
	hxrt.Println(pick(0))
	hxrt.Println(pick(2))
	hxrt.Println(pick(7))
}

func pick(v int) int {
	var hx_switch_1 int
	switch v {
	case 0, 1:
		hx_switch_1 = 10
	case 2:
		hx_switch_1 = 20
	default:
		hx_switch_1 = 30
	}
	return hx_switch_1
}
