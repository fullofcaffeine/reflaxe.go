package main

import "snapshot/hxrt"

func _GuardedTry__GuardedTry_Fields__readValue(skip bool, fail bool) *string {
	if skip {
		return hxrt.StringFromLiteral("")
	}
	hx_try_return_1 := false
	var hx_try_value_2 *string
	hxrt.TryCatch(func() {
		if fail {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("failed")))
		}
		hx_try_value_2 = hxrt.StringFromLiteral("value")
		hx_try_return_1 = true
		return
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_exception_4
		var hx_tmp any = hx_payload_5
		_ = hx_tmp
		hx_try_value_2 = hxrt.StringFromLiteral("fallback")
		hx_try_return_1 = true
		return
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	var hx_try_zero_6 *string
	return hx_try_zero_6
}
