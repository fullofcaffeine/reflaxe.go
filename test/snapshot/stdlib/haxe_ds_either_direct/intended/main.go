package main

import "snapshot/hxrt"

func main() {
	left := haxe__ds__Either_Left(hxrt.StringFromLiteral("go"))
	right := haxe__ds__Either_Right(7)
	hxrt.Println(func() any {
		var hx_switch_1 any
		switch left.tag {
		case 0:
			_g := left.params[0].(*string)
			value := _g
			hx_switch_1 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("left="), value)
		case 1:
			_g_1 := left.params[0].(int)
			value_1 := _g_1
			hx_switch_1 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("right="), hxrt.StdString(value_1))
		}
		return hx_switch_1
	}())
	hxrt.Println(func() any {
		var hx_switch_2 any
		switch right.tag {
		case 0:
			_g_2 := right.params[0].(*string)
			value_2 := _g_2
			hx_switch_2 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("left="), value_2)
		case 1:
			_g_3 := right.params[0].(int)
			value_3 := _g_3
			hx_switch_2 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("right="), hxrt.StdString(value_3))
		}
		return hx_switch_2
	}())
}
