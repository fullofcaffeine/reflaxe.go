package main

import "snapshot/hxrt"

func NonLaneFallback_run() {
	nonLaneResult := NonLaneFallback_unresolvedFail(hxrt.StringFromLiteral("non-lane"))
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, (nonLaneResult == nil), func() map[string]any {
		hx_obj_3 := map[string]any{}
		hx_obj_3["fileName"] = hxrt.StringFromLiteral("NonLaneFallback.hx")
		hx_obj_3["lineNumber"] = 9
		hx_obj_3["className"] = hxrt.StringFromLiteral("NonLaneFallback")
		hx_obj_3["methodName"] = hxrt.StringFromLiteral("run")
		return hx_obj_3
	}())
}

func NonLaneFallback_unresolvedFail(message *string) *go___Result {
	return go___Go_fail(message)
}
