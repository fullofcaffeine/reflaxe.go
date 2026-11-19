package main

import "snapshot/hxrt"

func main() {
	source := hxrt.NewArray(1, 2)
	tail := hxrt.NewArray(3, 2)
	combined := func() *hxrt.Array {
		hx_concat_target_1 := source
		hx_concat_other_2 := tail
		hx_concat_result_3 := hx_concat_target_1.Copy()
		for _, hx_concat_item_4 := range hx_concat_other_2.Values() {
			hx_concat_result_3.Push(hx_concat_item_4)
		}
		return hx_concat_result_3
	}()
	var v any = any(hxrt.StringJoinAny(combined.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v)
	var v_1 any = any(func() int {
		hx_indexof_target_5 := combined
		var hx_indexof_value_6 int = 2
		hx_indexof_length_7 := hx_indexof_target_5.Len()
		hx_indexof_index_8 := 0
		for hx_indexof_index_8 < hx_indexof_length_7 {
			hx_indexof_element_9 := hxrt.IntFromNullableAny(hx_indexof_target_5.Get(hx_indexof_index_8))
			if hx_indexof_element_9 == hx_indexof_value_6 {
				return hx_indexof_index_8
			}
			hx_indexof_index_8 = (hx_indexof_index_8 + 1)
		}
		return -1
	}())
	hxrt.Println(v_1)
	var v_2 any = any(func() int {
		hx_indexof_target_10 := combined
		var hx_indexof_value_11 int = 2
		hx_indexof_index_13 := 2
		hx_indexof_length_12 := hx_indexof_target_10.Len()
		if hx_indexof_index_13 < 0 {
			hx_indexof_index_13 = (hx_indexof_length_12 + hx_indexof_index_13)
		}
		if hx_indexof_index_13 < 0 {
			hx_indexof_index_13 = 0
		}
		for hx_indexof_index_13 < hx_indexof_length_12 {
			hx_indexof_element_14 := hxrt.IntFromNullableAny(hx_indexof_target_10.Get(hx_indexof_index_13))
			if hx_indexof_element_14 == hx_indexof_value_11 {
				return hx_indexof_index_13
			}
			hx_indexof_index_13 = (hx_indexof_index_13 + 1)
		}
		return -1
	}())
	hxrt.Println(v_2)
	var v_3 any = any(func() int {
		hx_indexof_target_15 := combined
		var hx_indexof_value_16 int = 2
		hx_indexof_index_18 := -1
		hx_indexof_length_17 := hx_indexof_target_15.Len()
		if hx_indexof_index_18 < 0 {
			hx_indexof_index_18 = (hx_indexof_length_17 + hx_indexof_index_18)
		}
		if hx_indexof_index_18 < 0 {
			hx_indexof_index_18 = 0
		}
		for hx_indexof_index_18 < hx_indexof_length_17 {
			hx_indexof_element_19 := hxrt.IntFromNullableAny(hx_indexof_target_15.Get(hx_indexof_index_18))
			if hx_indexof_element_19 == hx_indexof_value_16 {
				return hx_indexof_index_18
			}
			hx_indexof_index_18 = (hx_indexof_index_18 + 1)
		}
		return -1
	}())
	hxrt.Println(v_3)
	var v_4 any = any(func() int {
		hx_indexof_target_20 := combined
		var hx_indexof_value_21 int = 2
		hx_indexof_index_23 := -8
		hx_indexof_length_22 := hx_indexof_target_20.Len()
		if hx_indexof_index_23 < 0 {
			hx_indexof_index_23 = (hx_indexof_length_22 + hx_indexof_index_23)
		}
		if hx_indexof_index_23 < 0 {
			hx_indexof_index_23 = 0
		}
		for hx_indexof_index_23 < hx_indexof_length_22 {
			hx_indexof_element_24 := hxrt.IntFromNullableAny(hx_indexof_target_20.Get(hx_indexof_index_23))
			if hx_indexof_element_24 == hx_indexof_value_21 {
				return hx_indexof_index_23
			}
			hx_indexof_index_23 = (hx_indexof_index_23 + 1)
		}
		return -1
	}())
	hxrt.Println(v_4)
	var v_5 any = any(func() int {
		hx_indexof_target_25 := combined
		var hx_indexof_value_26 int = 9
		hx_indexof_length_27 := hx_indexof_target_25.Len()
		hx_indexof_index_28 := 0
		for hx_indexof_index_28 < hx_indexof_length_27 {
			hx_indexof_element_29 := hxrt.IntFromNullableAny(hx_indexof_target_25.Get(hx_indexof_index_28))
			if hx_indexof_element_29 == hx_indexof_value_26 {
				return hx_indexof_index_28
			}
			hx_indexof_index_28 = (hx_indexof_index_28 + 1)
		}
		return -1
	}())
	hxrt.Println(v_5)
	hx_array_target_30 := source
	hx_array_index_31 := 0
	hx_array_target_30.Set(hx_array_index_31, 9)
	hx_array_target_32 := tail
	hx_array_index_33 := 0
	hx_array_target_32.Set(hx_array_index_33, 8)
	var v_6 any = any(hxrt.StringJoinAny(combined.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_6)
	rebuiltA := hxrt.StringSubstrStringPtr(hxrt.StringFromLiteral("alpha"), 0, 1, true)
	words := hxrt.NewArray(rebuiltA, hxrt.StringFromLiteral("b"))
	var v_7 any = any(func() int {
		hx_indexof_target_34 := words
		var hx_indexof_value_35 *string = hxrt.StringFromLiteral("a")
		hx_indexof_length_36 := hx_indexof_target_34.Len()
		hx_indexof_index_37 := 0
		for hx_indexof_index_37 < hx_indexof_length_36 {
			hx_indexof_element_38 := func(hx_value_39 any) *string {
				if hx_value_39 == nil {
					var hx_zero_40 *string
					return hx_zero_40
				}
				return hx_value_39.(*string)
			}(hx_indexof_target_34.Get(hx_indexof_index_37))
			if hxrt.StringEqualStringPtr(hx_indexof_element_38, hx_indexof_value_35) {
				return hx_indexof_index_37
			}
			hx_indexof_index_37 = (hx_indexof_index_37 + 1)
		}
		return -1
	}())
	hxrt.Println(v_7)
}
