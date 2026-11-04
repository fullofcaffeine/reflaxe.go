package main

import "snapshot/hxrt"

func LaneWorker_produce() {
	laneResult := LaneWorker_unresolvedFail(hxrt.StringFromLiteral("lane"))
	func() {
		var hx_callable_2 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_3 any = (laneResult == nil)
		var hx_argument_4 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_1 := map[string]any{}
			hx_obj_1["fileName"] = hxrt.StringFromLiteral("LaneWorker.hx")
			hx_obj_1["lineNumber"] = 10
			hx_obj_1["className"] = hxrt.StringFromLiteral("LaneWorker")
			hx_obj_1["methodName"] = hxrt.StringFromLiteral("produce")
			return hxrt.AnonymousFromMap(hx_obj_1)
		}()
		hxrt.FunctionCall(hx_callable_2)(hx_argument_3, hx_argument_4)
	}()
}

func LaneWorker_unresolvedFail(message *string) *go___Result {
	return go___Go_fail(message)
}
