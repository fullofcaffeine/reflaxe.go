package main

import "snapshot/hxrt"

func main() {
	if !hxrt.StringEqualStringPtr(readValue(false), hxrt.StringFromLiteral("value")) || !hxrt.StringEqualStringPtr(readValue(true), hxrt.StringFromLiteral("fallback")) {
		hxrt.Throw(hxrt.StringFromLiteral("unexpected result"))
	}
}

func readValue(fail bool) *string {
	hx_try_return_1 := false
	var hx_try_value_2 *string
	hxrt.TryCatch(func() {
		if fail {
			hxrt.Throw(hxrt.StringFromLiteral("failed"))
		}
		hx_try_value_2 = hxrt.StringFromLiteral("value")
		hx_try_return_1 = true
		return
	}, func(hx_caught_3 any) {
		hx_tmp := hx_caught_3
		_ = hx_tmp
		hx_try_value_2 = hxrt.StringFromLiteral("fallback")
		hx_try_return_1 = true
		return
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	var hx_try_zero_5 *string
	return hx_try_zero_5
}
