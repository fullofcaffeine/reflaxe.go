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
	_ = current
	switch current.tag {
	case 0:
		hxrt.Println(0)
	case 1:
		hxrt.Println(1)
	}
	hxrt.Println(toInt(Flag_Off))
	hxrt.Println(toInt(Flag_On))
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
