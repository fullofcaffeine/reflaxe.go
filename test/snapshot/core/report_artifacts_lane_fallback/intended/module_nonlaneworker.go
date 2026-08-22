package main

import "snapshot/hxrt"

func NonLaneWorker_produce() {
	nonLaneResult := NonLaneWorker_unresolvedFail(hxrt.StringFromLiteral("non-lane"))
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, (nonLaneResult == nil), func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("NonLaneWorker.hx")
		hx_obj_1["lineNumber"] = 9
		hx_obj_1["className"] = hxrt.StringFromLiteral("NonLaneWorker")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("produce")
		return hx_obj_1
	}())
}

func NonLaneWorker_unresolvedFail(message *string) *go___Result {
	return go___Go_fail(message)
}
