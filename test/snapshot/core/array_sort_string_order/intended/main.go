package main

import "snapshot/hxrt"

func compare(left *string, right *string) int {
	var hx_if_2 int
	if hxrt.StringCompareStringPtr(left, right) < 0 {
		hx_if_2 = -1
	} else {
		var hx_if_1 int
		if hxrt.StringEqualStringPtr(left, right) {
			hx_if_1 = 0
		} else {
			hx_if_1 = 1
		}
		hx_if_2 = hx_if_1
	}
	return hx_if_2
}

func main() {
	values := hxrt.NewArray(hxrt.StringFromLiteral("beta"), hxrt.StringFromLiteral("alpha"), hxrt.StringFromLiteral("gamma"))
	hxrt.ArraySort(values, func(hx_comparator_5 *hxrt.Function[func(*string, *string) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_5, func(hx_cmp_left_3 any, hx_cmp_right_4 any) int {
			return hxrt.FunctionCall(hx_comparator_5)(func(hx_value_6 any) *string {
				if hx_value_6 == nil {
					var hx_zero_7 *string
					return hx_zero_7
				}
				return hx_value_6.(*string)
			}(hx_cmp_left_3), func(hx_value_8 any) *string {
				if hx_value_8 == nil {
					var hx_zero_9 *string
					return hx_zero_9
				}
				return hx_value_8.(*string)
			}(hx_cmp_right_4))
		})
	}(hxrt.NamedFunction(compare, "Main.compare")))
	score := 0
	if hxrt.StringEqualAny(values.Get(0), hxrt.StringFromLiteral("alpha")) {
		score = int((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(1)))
	}
	if hxrt.StringEqualAny(values.Get(1), hxrt.StringFromLiteral("beta")) {
		score = int((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(2)))
	}
	if hxrt.StringEqualAny(values.Get(2), hxrt.StringFromLiteral("gamma")) {
		score = int((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(4)))
	}
	if hxrt.StringCompareStringPtr(hxrt.StringFromLiteral("alpha"), hxrt.StringFromLiteral("beta")) < 0 {
		score = int((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(8)))
	}
	if hxrt.StringCompareStringPtr(hxrt.StringFromLiteral("beta"), hxrt.StringFromLiteral("beta")) <= 0 {
		score = int((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(16)))
	}
	if hxrt.StringCompareStringPtr(hxrt.StringFromLiteral("gamma"), hxrt.StringFromLiteral("beta")) > 0 {
		score = int((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(32)))
	}
	if hxrt.StringCompareStringPtr(hxrt.StringFromLiteral("gamma"), hxrt.StringFromLiteral("gamma")) >= 0 {
		score = int((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(64)))
	}
	if score != 127 {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("unexpected sort or string ordering result")))
	}
}

func hx_recover_class_0(value any) *haxe__ValueException {
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

func hx_recover_class_1(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
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
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
