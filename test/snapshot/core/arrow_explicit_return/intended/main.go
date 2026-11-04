package main

import "snapshot/hxrt"

func main() {
	rows := hxrt.NewArray(func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["index"] = 0
		hx_obj_1["key"] = 2
		return hx_obj_1
	}(), func() map[string]any {
		hx_obj_2 := map[string]any{}
		hx_obj_2["index"] = 1
		hx_obj_2["key"] = 1
		return hx_obj_2
	}(), func() map[string]any {
		hx_obj_3 := map[string]any{}
		hx_obj_3["index"] = 2
		hx_obj_3["key"] = 1
		return hx_obj_3
	}())
	hxrt.ArraySort(rows, func(hx_comparator_27 *hxrt.Function[func(map[string]any, map[string]any) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_27, func(hx_cmp_left_25 any, hx_cmp_right_26 any) int {
			return hxrt.FunctionCall(hx_comparator_27)(func(hx_value_28 any) map[string]any {
				if hx_value_28 == nil {
					var hx_zero_29 map[string]any
					return hx_zero_29
				}
				return hx_value_28.(map[string]any)
			}(hx_cmp_left_25), func(hx_value_30 any) map[string]any {
				if hx_value_30 == nil {
					var hx_zero_31 map[string]any
					return hx_zero_31
				}
				return hx_value_30.(map[string]any)
			}(hx_cmp_right_26))
		})
	}(hxrt.NewFunction(func(left map[string]any, right map[string]any) int {
		compared := int((hxrt.Int32Wrap(func(hx_obj_4 map[string]any) int {
			hx_field_5 := hx_obj_4["key"]
			if hx_field_5 == nil {
				var hx_zero_6 int
				return hx_zero_6
			}
			return hx_field_5.(int)
		}(left)) - hxrt.Int32Wrap(func(hx_obj_7 map[string]any) int {
			hx_field_8 := hx_obj_7["key"]
			if hx_field_8 == nil {
				var hx_zero_9 int
				return hx_zero_9
			}
			return hx_field_8.(int)
		}(right))))
		var hx_if_24 int
		if compared != 0 {
			hx_if_24 = compared
		} else {
			var hx_if_23 int
			if func(hx_obj_10 map[string]any) int {
				hx_field_11 := hx_obj_10["index"]
				if hx_field_11 == nil {
					var hx_zero_12 int
					return hx_zero_12
				}
				return hx_field_11.(int)
			}(left) < func(hx_obj_13 map[string]any) int {
				hx_field_14 := hx_obj_13["index"]
				if hx_field_14 == nil {
					var hx_zero_15 int
					return hx_zero_15
				}
				return hx_field_14.(int)
			}(right) {
				hx_if_23 = -1
			} else {
				var hx_if_22 int
				if func(hx_obj_16 map[string]any) int {
					hx_field_17 := hx_obj_16["index"]
					if hx_field_17 == nil {
						var hx_zero_18 int
						return hx_zero_18
					}
					return hx_field_17.(int)
				}(left) == func(hx_obj_19 map[string]any) int {
					hx_field_20 := hx_obj_19["index"]
					if hx_field_20 == nil {
						var hx_zero_21 int
						return hx_zero_21
					}
					return hx_field_20.(int)
				}(right) {
					hx_if_22 = 0
				} else {
					hx_if_22 = 1
				}
				hx_if_23 = hx_if_22
			}
			hx_if_24 = hx_if_23
		}
		return hx_if_24
	})))
	var v any = any(hxrt.StringJoinAny(func() *hxrt.Array {
		_g := hxrt.NewArray()
		_g1 := 0
		for _g1 < rows.Len() {
			row := func(hx_value_32 any) map[string]any {
				if hx_value_32 == nil {
					var hx_zero_33 map[string]any
					return hx_zero_33
				}
				return hx_value_32.(map[string]any)
			}(rows.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func(hx_obj_35 map[string]any) int {
				hx_field_36 := hx_obj_35["index"]
				if hx_field_36 == nil {
					var hx_zero_37 int
					return hx_zero_37
				}
				return hx_field_36.(int)
			}(row))
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v)
}
