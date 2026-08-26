package main

import "snapshot/hxrt"

func describe(value int) *string {
	var hx_switch_1 *string
	switch value {
	case 0:
		return hxrt.StringFromLiteral("zero")
	case 1:
		hx_switch_1 = hxrt.StringFromLiteral("one")
	default:
		if value < 0 {
			return hxrt.StringFromLiteral("negative")
		}
		hx_switch_1 = hxrt.StringFromLiteral("many")
	}
	selected := hx_switch_1
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("selected:"), selected)
}

func main() {
	var v any = any(describe(0))
	hxrt.Println(v)
	var v_1 any = any(describe(1))
	hxrt.Println(v_1)
	var v_2 any = any(describe(-1))
	hxrt.Println(v_2)
	var v_3 any = any(describe(2))
	hxrt.Println(v_3)
}
