package main

import "snapshot/hxrt"

func NonLaneWorker_produce() {
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, PortableSurfaceDigest_compute(13), func() map[string]any {
		hx_obj_4 := map[string]any{}
		hx_obj_4["fileName"] = hxrt.StringFromLiteral("NonLaneWorker.hx")
		hx_obj_4["lineNumber"] = 7
		hx_obj_4["className"] = hxrt.StringFromLiteral("NonLaneWorker")
		hx_obj_4["methodName"] = hxrt.StringFromLiteral("produce")
		return hx_obj_4
	}())
	nonLaneResult := NonLaneWorker_unresolvedFail(hxrt.StringFromLiteral("non-lane"))
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, (nonLaneResult == nil), func() map[string]any {
		hx_obj_5 := map[string]any{}
		hx_obj_5["fileName"] = hxrt.StringFromLiteral("NonLaneWorker.hx")
		hx_obj_5["lineNumber"] = 10
		hx_obj_5["className"] = hxrt.StringFromLiteral("NonLaneWorker")
		hx_obj_5["methodName"] = hxrt.StringFromLiteral("produce")
		return hx_obj_5
	}())
}

func NonLaneWorker_unresolvedFail(message *string) *go___Result {
	return go___Go_fail(message)
}
