package main

import "snapshot/hxrt"

func main() {
	rows := hxrt.NewArray(func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["index"] = 0
		hx_obj_1["key"] = 2
		return hxrt.AnonymousFromMap(hx_obj_1)
	}(), func() hxrt.Anonymous {
		hx_obj_2 := map[string]any{}
		hx_obj_2["index"] = 1
		hx_obj_2["key"] = 1
		return hxrt.AnonymousFromMap(hx_obj_2)
	}(), func() hxrt.Anonymous {
		hx_obj_3 := map[string]any{}
		hx_obj_3["index"] = 2
		hx_obj_3["key"] = 1
		return hxrt.AnonymousFromMap(hx_obj_3)
	}())
	hxrt.ArraySort(rows, func(hx_comparator_21 *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_21, func(hx_cmp_left_19 any, hx_cmp_right_20 any) int {
			return hxrt.FunctionCall(hx_comparator_21)(hx_structural_recover(hx_cmp_left_19), hx_structural_recover(hx_cmp_right_20))
		})
	}(hxrt.NewFunction(func(left hxrt.Anonymous, right hxrt.Anonymous) int {
		compared := int((hxrt.Int32Wrap(func(hx_value_4 any) int {
			if hx_value_4 == nil {
				var hx_zero_5 int
				return hx_zero_5
			}
			return hx_value_4.(int)
		}(hxrt.AnonymousGet(left, "key"))) - hxrt.Int32Wrap(func(hx_value_6 any) int {
			if hx_value_6 == nil {
				var hx_zero_7 int
				return hx_zero_7
			}
			return hx_value_6.(int)
		}(hxrt.AnonymousGet(right, "key")))))
		var hx_if_18 int
		if compared != 0 {
			hx_if_18 = compared
		} else {
			var hx_if_17 int
			if func(hx_value_8 any) int {
				if hx_value_8 == nil {
					var hx_zero_9 int
					return hx_zero_9
				}
				return hx_value_8.(int)
			}(hxrt.AnonymousGet(left, "index")) < func(hx_value_10 any) int {
				if hx_value_10 == nil {
					var hx_zero_11 int
					return hx_zero_11
				}
				return hx_value_10.(int)
			}(hxrt.AnonymousGet(right, "index")) {
				hx_if_17 = -1
			} else {
				var hx_if_16 int
				if func(hx_value_12 any) int {
					if hx_value_12 == nil {
						var hx_zero_13 int
						return hx_zero_13
					}
					return hx_value_12.(int)
				}(hxrt.AnonymousGet(left, "index")) == func(hx_value_14 any) int {
					if hx_value_14 == nil {
						var hx_zero_15 int
						return hx_zero_15
					}
					return hx_value_14.(int)
				}(hxrt.AnonymousGet(right, "index")) {
					hx_if_16 = 0
				} else {
					hx_if_16 = 1
				}
				hx_if_17 = hx_if_16
			}
			hx_if_18 = hx_if_17
		}
		return hx_if_18
	})))
	var v any = any(hxrt.StringJoinAny(func() *hxrt.Array {
		_g := hxrt.NewArray()
		_g1 := 0
		for _g1 < rows.Len() {
			row := hx_structural_recover(rows.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func(hx_value_23 any) int {
				if hx_value_23 == nil {
					var hx_zero_24 int
					return hx_zero_24
				}
				return hx_value_23.(int)
			}(hxrt.AnonymousGet(row, "index")))
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_1 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_1.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_1.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_1.high = func(hx_value_2 any) int {
						if hx_value_2 == nil {
							var hx_zero_3 int
							return hx_zero_3
						}
						return hx_value_2.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_1.low = func(hx_value_4 any) int {
						if hx_value_4 == nil {
							var hx_zero_5 int
							return hx_zero_5
						}
						return hx_value_4.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}
