package main

import "snapshot/hxrt"

func main() {
	source := hxrt.NewArray(1, 2, 3)
	native := func(hx_lambda_raw_1 []any) []int {
		hx_lambda_out_2 := make([]int, 0, len(hx_lambda_raw_1))
		for _, hx_lambda_item_3 := range hx_lambda_raw_1 {
			hx_lambda_out_2 = append(hx_lambda_out_2, func(hx_value_4 any) int {
				if hx_value_4 == nil {
					var hx_zero_5 int
					return hx_zero_5
				}
				return hx_value_4.(int)
			}(hx_lambda_item_3))
		}
		return hx_lambda_out_2
	}(source.Values())
	native[0] = 7
	portable := hxrt.ArrayFromValues(func(hx_sort_src_6 []int) []any {
		hx_sort_out_8 := make([]any, 0, len(hx_sort_src_6))
		for _, hx_sort_item_7 := range hx_sort_src_6 {
			hx_sort_out_8 = append(hx_sort_out_8, hx_sort_item_7)
		}
		return hx_sort_out_8
	}(native))
	portable.Push(9)
	dynamicSource := hxrt.NewArray(hxrt.StringFromLiteral("source"))
	dynamicNative := dynamicSource.ValuesCopy()
	dynamicNative[0] = hxrt.StringFromLiteral("native")
	dynamicPortable := hxrt.ArrayFromValues(dynamicNative)
	hx_array_target_10 := dynamicPortable
	hx_array_index_11 := 0
	hx_array_target_10.Set(hx_array_index_11, hxrt.StringFromLiteral("portable"))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("source="), hxrt.StringJoinAny(source.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native="), render(native)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("portable="), hxrt.StringJoinAny(portable.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic.source="), hxrt.StdString(dynamicSource.Get(0))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic.native="), hxrt.StdString(dynamicNative[0])))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic.portable="), hxrt.StdString(dynamicPortable.Get(0))))
	hxrt.Println(v_5)
}

func render(values []int) *string {
	out := hxrt.StringFromLiteral("")
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_12 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_12
		if index > 0 {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatStringPtr(out, hxrt.StdString(values[index]))
	}
	return out
}
