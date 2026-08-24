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
		hx_indexof_value_6 := 2
		var hx_indexof_start_input_7 any = 0
		hx_indexof_start_8 := hxrt.IntFromNullableAny(hx_indexof_start_input_7)
		hx_indexof_length_9 := hx_indexof_target_5.Len()
		if hx_indexof_start_8 < 0 {
			hx_indexof_start_8 = (hx_indexof_length_9 + hx_indexof_start_8)
		}
		if hx_indexof_start_8 < 0 {
			hx_indexof_start_8 = 0
		}
		hx_indexof_index_10 := hx_indexof_start_8
		for hx_indexof_index_10 < hx_indexof_length_9 {
			hx_indexof_element_11 := hxrt.IntFromNullableAny(hx_indexof_target_5.Get(hx_indexof_index_10))
			if hx_indexof_element_11 == hx_indexof_value_6 {
				return hx_indexof_index_10
			}
			hx_indexof_index_10 = (hx_indexof_index_10 + 1)
		}
		return -1
	}())
	hxrt.Println(v_1)
	var v_2 any = any(func() int {
		hx_indexof_target_12 := combined
		hx_indexof_value_13 := 2
		var hx_indexof_start_input_14 any = 2
		hx_indexof_start_15 := hxrt.IntFromNullableAny(hx_indexof_start_input_14)
		hx_indexof_length_16 := hx_indexof_target_12.Len()
		if hx_indexof_start_15 < 0 {
			hx_indexof_start_15 = (hx_indexof_length_16 + hx_indexof_start_15)
		}
		if hx_indexof_start_15 < 0 {
			hx_indexof_start_15 = 0
		}
		hx_indexof_index_17 := hx_indexof_start_15
		for hx_indexof_index_17 < hx_indexof_length_16 {
			hx_indexof_element_18 := hxrt.IntFromNullableAny(hx_indexof_target_12.Get(hx_indexof_index_17))
			if hx_indexof_element_18 == hx_indexof_value_13 {
				return hx_indexof_index_17
			}
			hx_indexof_index_17 = (hx_indexof_index_17 + 1)
		}
		return -1
	}())
	hxrt.Println(v_2)
	var v_3 any = any(func() int {
		hx_indexof_target_19 := combined
		hx_indexof_value_20 := 2
		var hx_indexof_start_input_21 any = -1
		hx_indexof_start_22 := hxrt.IntFromNullableAny(hx_indexof_start_input_21)
		hx_indexof_length_23 := hx_indexof_target_19.Len()
		if hx_indexof_start_22 < 0 {
			hx_indexof_start_22 = (hx_indexof_length_23 + hx_indexof_start_22)
		}
		if hx_indexof_start_22 < 0 {
			hx_indexof_start_22 = 0
		}
		hx_indexof_index_24 := hx_indexof_start_22
		for hx_indexof_index_24 < hx_indexof_length_23 {
			hx_indexof_element_25 := hxrt.IntFromNullableAny(hx_indexof_target_19.Get(hx_indexof_index_24))
			if hx_indexof_element_25 == hx_indexof_value_20 {
				return hx_indexof_index_24
			}
			hx_indexof_index_24 = (hx_indexof_index_24 + 1)
		}
		return -1
	}())
	hxrt.Println(v_3)
	var v_4 any = any(func() int {
		hx_indexof_target_26 := combined
		hx_indexof_value_27 := 2
		var hx_indexof_start_input_28 any = -8
		hx_indexof_start_29 := hxrt.IntFromNullableAny(hx_indexof_start_input_28)
		hx_indexof_length_30 := hx_indexof_target_26.Len()
		if hx_indexof_start_29 < 0 {
			hx_indexof_start_29 = (hx_indexof_length_30 + hx_indexof_start_29)
		}
		if hx_indexof_start_29 < 0 {
			hx_indexof_start_29 = 0
		}
		hx_indexof_index_31 := hx_indexof_start_29
		for hx_indexof_index_31 < hx_indexof_length_30 {
			hx_indexof_element_32 := hxrt.IntFromNullableAny(hx_indexof_target_26.Get(hx_indexof_index_31))
			if hx_indexof_element_32 == hx_indexof_value_27 {
				return hx_indexof_index_31
			}
			hx_indexof_index_31 = (hx_indexof_index_31 + 1)
		}
		return -1
	}())
	hxrt.Println(v_4)
	var v_5 any = any(func() int {
		hx_indexof_target_33 := combined
		hx_indexof_value_34 := 9
		var hx_indexof_start_input_35 any = 0
		hx_indexof_start_36 := hxrt.IntFromNullableAny(hx_indexof_start_input_35)
		hx_indexof_length_37 := hx_indexof_target_33.Len()
		if hx_indexof_start_36 < 0 {
			hx_indexof_start_36 = (hx_indexof_length_37 + hx_indexof_start_36)
		}
		if hx_indexof_start_36 < 0 {
			hx_indexof_start_36 = 0
		}
		hx_indexof_index_38 := hx_indexof_start_36
		for hx_indexof_index_38 < hx_indexof_length_37 {
			hx_indexof_element_39 := hxrt.IntFromNullableAny(hx_indexof_target_33.Get(hx_indexof_index_38))
			if hx_indexof_element_39 == hx_indexof_value_34 {
				return hx_indexof_index_38
			}
			hx_indexof_index_38 = (hx_indexof_index_38 + 1)
		}
		return -1
	}())
	hxrt.Println(v_5)
	hx_array_target_40 := source
	hx_array_index_41 := 0
	hx_array_target_40.Set(hx_array_index_41, 9)
	hx_array_target_42 := tail
	hx_array_index_43 := 0
	hx_array_target_42.Set(hx_array_index_43, 8)
	var v_6 any = any(hxrt.StringJoinAny(combined.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_6)
	rebuiltA := hxrt.StringSubstrStringPtr(hxrt.StringFromLiteral("alpha"), 0, 1, true)
	words := hxrt.NewArray(rebuiltA, hxrt.StringFromLiteral("b"))
	var v_7 any = any(func() int {
		hx_indexof_target_44 := words
		hx_indexof_value_45 := hxrt.StringFromLiteral("a")
		var hx_indexof_start_input_46 any = 0
		hx_indexof_start_47 := hxrt.IntFromNullableAny(hx_indexof_start_input_46)
		hx_indexof_length_48 := hx_indexof_target_44.Len()
		if hx_indexof_start_47 < 0 {
			hx_indexof_start_47 = (hx_indexof_length_48 + hx_indexof_start_47)
		}
		if hx_indexof_start_47 < 0 {
			hx_indexof_start_47 = 0
		}
		hx_indexof_index_49 := hx_indexof_start_47
		for hx_indexof_index_49 < hx_indexof_length_48 {
			hx_indexof_element_50 := func(hx_value_51 any) *string {
				if hx_value_51 == nil {
					var hx_zero_52 *string
					return hx_zero_52
				}
				return hx_value_51.(*string)
			}(hx_indexof_target_44.Get(hx_indexof_index_49))
			if hxrt.StringEqualStringPtr(hx_indexof_element_50, hx_indexof_value_45) {
				return hx_indexof_index_49
			}
			hx_indexof_index_49 = (hx_indexof_index_49 + 1)
		}
		return -1
	}())
	hxrt.Println(v_7)
}
