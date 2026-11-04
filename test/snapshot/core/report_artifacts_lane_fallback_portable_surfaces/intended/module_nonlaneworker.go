package main

import "snapshot/hxrt"

func NonLaneWorker_produce() {
	func() {
		var hx_callable_2 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_3 any = PortableSurfaceDigest_compute(13)
		var hx_argument_4 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_1 := map[string]any{}
			hx_obj_1["fileName"] = hxrt.StringFromLiteral("NonLaneWorker.hx")
			hx_obj_1["lineNumber"] = 7
			hx_obj_1["className"] = hxrt.StringFromLiteral("NonLaneWorker")
			hx_obj_1["methodName"] = hxrt.StringFromLiteral("produce")
			return hxrt.AnonymousFromMap(hx_obj_1)
		}()
		hxrt.FunctionCall(hx_callable_2)(hx_argument_3, hx_argument_4)
	}()
	nonLaneResult := NonLaneWorker_unresolvedFail(hxrt.StringFromLiteral("non-lane"))
	func() {
		var hx_callable_6 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_7 any = (nonLaneResult == nil)
		var hx_argument_8 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_5 := map[string]any{}
			hx_obj_5["fileName"] = hxrt.StringFromLiteral("NonLaneWorker.hx")
			hx_obj_5["lineNumber"] = 10
			hx_obj_5["className"] = hxrt.StringFromLiteral("NonLaneWorker")
			hx_obj_5["methodName"] = hxrt.StringFromLiteral("produce")
			return hxrt.AnonymousFromMap(hx_obj_5)
		}()
		hxrt.FunctionCall(hx_callable_6)(hx_argument_7, hx_argument_8)
	}()
}

func NonLaneWorker_unresolvedFail(message *string) *go___Result {
	return go___Go_fail(message)
}
