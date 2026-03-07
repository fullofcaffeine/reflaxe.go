package main

import "snapshot/hxrt"

func cmp(a int, b int) int {
	return int(int32((hxrt.Int32Wrap(a) - hxrt.Int32Wrap(b))))
}

func main() {
	values := []int{5, 2, 4, 1, 3}
	func(hx_sort_src_2 []int) {
		hx_sort_raw_1 := func(hx_sort_src_3 []int) []any {
			hx_sort_out_5 := make([]any, 0, len(hx_sort_src_3))
			for _, hx_sort_item_4 := range hx_sort_src_3 {
				hx_sort_out_5 = append(hx_sort_out_5, hx_sort_item_4)
			}
			return hx_sort_out_5
		}(hx_sort_src_2)
		haxe__ds__ArraySort_sort(hx_sort_raw_1, func(hx_cmp_left_6 any, hx_cmp_right_7 any) int {
			return cmp(func(hx_value_8 any) int {
				if hx_value_8 == nil {
					var hx_zero_9 int
					return hx_zero_9
				}
				return hx_value_8.(int)
			}(hx_cmp_left_6), func(hx_value_10 any) int {
				if hx_value_10 == nil {
					var hx_zero_11 int
					return hx_zero_11
				}
				return hx_value_10.(int)
			}(hx_cmp_right_7))
		})
		func(hx_sort_raw_12 []any, hx_sort_dst_13 []int) {
			for hx_sort_i_14, hx_sort_item_15 := range hx_sort_raw_12 {
				hx_sort_dst_13[hx_sort_i_14] = func(hx_value_16 any) int {
					if hx_value_16 == nil {
						var hx_zero_17 int
						return hx_zero_17
					}
					return hx_value_16.(int)
				}(hx_sort_item_15)
			}
		}(hx_sort_raw_1, hx_sort_src_2)
	}(values)
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(values[0], hxrt.StringFromLiteral(",")), values[1]), hxrt.StringFromLiteral(",")), values[2]), hxrt.StringFromLiteral(",")), values[3]), hxrt.StringFromLiteral(",")), values[4]))
}
