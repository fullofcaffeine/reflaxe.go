package main

import "snapshot/hxrt"

func main() {
	m := New_haxe__ds__IntMap()
	m.__hx_this.set(1, hxrt.StringFromLiteral("one"))
	m.__hx_this.set(2, hxrt.StringFromLiteral("two"))
	one := func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(m.__hx_this.get(1))
	hxrt.Println(any(one))
	var v any = any(func(hx_value_3 any) bool {
		if hx_value_3 == nil {
			var hx_zero_4 bool
			return hx_zero_4
		}
		return hx_value_3.(bool)
	}(m.__hx_this.exists(3)))
	hxrt.Println(v)
}
