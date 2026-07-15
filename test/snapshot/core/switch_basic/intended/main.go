package main

import "snapshot/hxrt"

func main() {
	v := 1
	switch v {
	case 0:
		hxrt.Println(any(0))
	case 1:
		hxrt.Println(any(1))
	default:
		hxrt.Println(any(9))
	}
	var v_1 any = any(pick(0))
	hxrt.Println(v_1)
	var v_2 any = any(pick(2))
	hxrt.Println(v_2)
	var v_3 any = any(pick(7))
	hxrt.Println(v_3)
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
