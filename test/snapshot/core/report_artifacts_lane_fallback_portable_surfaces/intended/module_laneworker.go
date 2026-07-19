package main

import "snapshot/hxrt"

func LaneWorker_produce() {
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, PortableSurfaceDigest_compute(11), func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("LaneWorker.hx")
		hx_obj_1["lineNumber"] = 8
		hx_obj_1["className"] = hxrt.StringFromLiteral("LaneWorker")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("produce")
		return hx_obj_1
	}())
	laneResult := LaneWorker_unresolvedFail(hxrt.StringFromLiteral("lane"))
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, (laneResult == nil), func() map[string]any {
		hx_obj_2 := map[string]any{}
		hx_obj_2["fileName"] = hxrt.StringFromLiteral("LaneWorker.hx")
		hx_obj_2["lineNumber"] = 11
		hx_obj_2["className"] = hxrt.StringFromLiteral("LaneWorker")
		hx_obj_2["methodName"] = hxrt.StringFromLiteral("produce")
		return hx_obj_2
	}())
}

func LaneWorker_unresolvedFail(message *string) *go___Result {
	return go___Go_fail(message)
}
