package main

import "snapshot/hxrt"

func comparator() *hxrt.Function[func(int, int) int] {
	factories = int(int32((factories + 1)))
	return hxrt.NewFunction(func(left int, right int) int {
		return int((hxrt.Int32Wrap(left) - hxrt.Int32Wrap(right)))
	})
}

var factories int = 0

func main() {
	ordinary := hxrt.NewArray(3, 1, 2)
	hxrt.ArraySort(ordinary, func(hx_comparator_3 *hxrt.Function[func(int, int) int]) *hxrt.Function[func(any, any) int] {
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
	}(comparator()))
	var v any = any(hxrt.StringJoinAny(ordinary.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v)
	var v_1 any = any(factories)
	hxrt.Println(v_1)
	staged := hxrt.NewArray(6, 4, 5)
	haxe__ds__ArraySort_sort(staged, func(hx_comparator_10 *hxrt.Function[func(int, int) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_10, func(hx_cmp_left_8 any, hx_cmp_right_9 any) int {
			return hxrt.FunctionCall(hx_comparator_10)(func(hx_value_11 any) int {
				if hx_value_11 == nil {
					var hx_zero_12 int
					return hx_zero_12
				}
				return hx_value_11.(int)
			}(hx_cmp_left_8), func(hx_value_13 any) int {
				if hx_value_13 == nil {
					var hx_zero_14 int
					return hx_zero_14
				}
				return hx_value_13.(int)
			}(hx_cmp_right_9))
		})
	}(comparator()))
	var v_2 any = any(hxrt.StringJoinAny(staged.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_2)
	var v_3 any = any(factories)
	hxrt.Println(v_3)
	empty := hxrt.NewArray()
	hxrt.ArraySort(empty, func(hx_comparator_17 *hxrt.Function[func(int, int) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_17, func(hx_cmp_left_15 any, hx_cmp_right_16 any) int {
			return hxrt.FunctionCall(hx_comparator_17)(func(hx_value_18 any) int {
				if hx_value_18 == nil {
					var hx_zero_19 int
					return hx_zero_19
				}
				return hx_value_18.(int)
			}(hx_cmp_left_15), func(hx_value_20 any) int {
				if hx_value_20 == nil {
					var hx_zero_21 int
					return hx_zero_21
				}
				return hx_value_20.(int)
			}(hx_cmp_right_16))
		})
	}(nil))
	hxrt.Println(any(hxrt.StringFromLiteral("empty")))
	hxrt.TryCatch(func() {
		hxrt.ArraySort(hxrt.NewArray(2, 1), func(hx_comparator_27 *hxrt.Function[func(int, int) int]) *hxrt.Function[func(any, any) int] {
			return hxrt.AdaptFunction(hx_comparator_27, func(hx_cmp_left_25 any, hx_cmp_right_26 any) int {
				return hxrt.FunctionCall(hx_comparator_27)(func(hx_value_28 any) int {
					if hx_value_28 == nil {
						var hx_zero_29 int
						return hx_zero_29
					}
					return hx_value_28.(int)
				}(hx_cmp_left_25), func(hx_value_30 any) int {
					if hx_value_30 == nil {
						var hx_zero_31 int
						return hx_zero_31
					}
					return hx_value_30.(int)
				}(hx_cmp_right_26))
			})
		}(nil))
	}, func(hx_caught_22 any) {
		hx_exception_23 := haxe__Exception_caught(hx_caught_22)
		hx_payload_24 := haxe__Exception_catchValue(hx_caught_22)
		_ = hx_payload_24
		var error *haxe__Exception = hx_recover_class_0(hx_exception_23)
		var v_4 any = any(error.__hx_this.get_message())
		hxrt.Println(v_4)
	})
}

func hx_recover_class_0(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_1(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
