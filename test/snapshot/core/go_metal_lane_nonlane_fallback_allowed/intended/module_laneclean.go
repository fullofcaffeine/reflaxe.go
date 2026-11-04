package main

import "snapshot/hxrt"

func LaneClean_run() {
	ok := go__result_ok___string_f613ccd0(hxrt.StringFromLiteral("clean"))
	func() {
		var hx_callable_2 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_3 any = go__result_isOk___string_f613ccd0(ok)
		var hx_argument_4 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_1 := map[string]any{}
			hx_obj_1["fileName"] = hxrt.StringFromLiteral("LaneClean.hx")
			hx_obj_1["lineNumber"] = 5
			hx_obj_1["className"] = hxrt.StringFromLiteral("LaneClean")
			hx_obj_1["methodName"] = hxrt.StringFromLiteral("run")
			return hxrt.AnonymousFromMap(hx_obj_1)
		}()
		hxrt.FunctionCall(hx_callable_2)(hx_argument_3, hx_argument_4)
	}()
}
