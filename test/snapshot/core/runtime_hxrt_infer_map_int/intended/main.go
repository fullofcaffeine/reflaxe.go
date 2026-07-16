package main

import "snapshot/hxrt"

func main() {
	values := New_haxe__ds__IntMap()
	values.set(1, hxrt.StringFromLiteral("one"))
	func(hx_value_1 any) bool {
		if hx_value_1 == nil {
			var hx_zero_2 bool
			return hx_zero_2
		}
		return hx_value_1.(bool)
	}(values.exists(1))
}
