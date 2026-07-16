package main

import "snapshot/hxrt"

func main() {
	m := New_haxe__ds__IntMap()
	m.set(1, hxrt.StringFromLiteral("one"))
	m.set(2, hxrt.StringFromLiteral("two"))
	one := func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(m.get(1))
	hxrt.Println(any(one))
	var v any = any(func(hx_value_3 any) bool {
		if hx_value_3 == nil {
			var hx_zero_4 bool
			return hx_zero_4
		}
		return hx_value_3.(bool)
	}(m.exists(3)))
	hxrt.Println(v)
}
