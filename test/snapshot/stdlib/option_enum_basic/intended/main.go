package main

import "snapshot/hxrt"

func main() {
	var v any = any(render(haxe__ds__Option_Some(7)))
	hxrt.Println(v)
	var v_1 any = any(render(haxe__ds__Option_None))
	hxrt.Println(v_1)
}

func render(opt *haxe__ds__Option) *string {
	var hx_switch_1 *string
	switch opt.tag {
	case 0:
		_g := opt.params[0].(int)
		v := _g
		hx_switch_1 = hxrt.StringConcatAny(hxrt.StringFromLiteral("some:"), v)
	case 1:
		hx_switch_1 = hxrt.StringFromLiteral("none")
	}
	return hx_switch_1
}
