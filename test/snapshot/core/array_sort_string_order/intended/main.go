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
	hxrt.ArraySort(values, func(hx_cmp_left_3 any, hx_cmp_right_4 any) int {
		return compare(func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(hx_cmp_left_3), func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(hx_cmp_right_4))
	})
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
		hxrt.Throw(hxrt.StringFromLiteral("unexpected sort or string ordering result"))
	}
}
