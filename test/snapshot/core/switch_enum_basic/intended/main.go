package main

import "snapshot/hxrt"

type Flag struct {
	tag    int
	params []any
}

var Flag_Off *Flag = &Flag{tag: 0}

var Flag_On *Flag = &Flag{tag: 1}

func main() {
	current := Flag_On
	switch current.tag {
	case 0:
		hxrt.Println(any(0))
	case 1:
		hxrt.Println(any(1))
	}
	var v any = any(toInt(Flag_Off))
	hxrt.Println(v)
	var v_1 any = any(toInt(Flag_On))
	hxrt.Println(v_1)
}

func toInt(flag *Flag) int {
	var hx_switch_1 int
	switch flag.tag {
	case 0:
		hx_switch_1 = 0
	case 1:
		hx_switch_1 = 1
	}
	return hx_switch_1
}
