package main

import "snapshot/hxrt"

func cmp(a int, b int) int {
	return int((hxrt.Int32Wrap(a) - hxrt.Int32Wrap(b)))
}

func main() {
	values := hxrt.NewArray(5, 2, 4, 1, 3)
	haxe__ds__ArraySort_sort(values, func(hx_cmp_left_1 any, hx_cmp_right_2 any) int {
		return cmp(func(hx_value_3 any) int {
			if hx_value_3 == nil {
				var hx_zero_4 int
				return hx_zero_4
			}
			return hx_value_3.(int)
		}(hx_cmp_left_1), func(hx_value_5 any) int {
			if hx_value_5 == nil {
				var hx_zero_6 int
				return hx_zero_6
			}
			return hx_value_5.(int)
		}(hx_cmp_right_2))
	})
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(values.Get(0), hxrt.StringFromLiteral(",")), values.Get(1)), hxrt.StringFromLiteral(",")), values.Get(2)), hxrt.StringFromLiteral(",")), values.Get(3)), hxrt.StringFromLiteral(",")), values.Get(4))))
}
