package main

import "snapshot/hxrt"

func cmp(a int, b int) int {
	return int((hxrt.Int32Wrap(a) - hxrt.Int32Wrap(b)))
}

func main() {
	values := hxrt.NewArray(5, 2, 4, 1, 3)
	haxe__ds__ArraySort_sort(values, func(hx_comparator_3 *hxrt.Function[func(int, int) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_3, func(hx_cmp_left_1 any, hx_cmp_right_2 any) int {
			return hxrt.FunctionCall(hx_comparator_3)(func(hx_value_4 any) int {
				if hx_value_4 == nil {
					var hx_zero_5 int
					return hx_zero_5
				}
				return hx_value_4.(int)
			}(hx_cmp_left_1), func(hx_value_6 any) int {
				if hx_value_6 == nil {
					var hx_zero_7 int
					return hx_zero_7
				}
				return hx_value_6.(int)
			}(hx_cmp_right_2))
		})
	}(hxrt.NamedFunction(cmp, "Main.cmp")))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(values.Get(0), hxrt.StringFromLiteral(",")), values.Get(1)), hxrt.StringFromLiteral(",")), values.Get(2)), hxrt.StringFromLiteral(",")), values.Get(3)), hxrt.StringFromLiteral(",")), values.Get(4))))
}
