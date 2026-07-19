package main

import "snapshot/hxrt"

func LaneClean_run() {
	ok := go__result_ok___string_f613ccd0(hxrt.StringFromLiteral("clean"))
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, go__result_isOk___string_f613ccd0(ok), func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("LaneClean.hx")
		hx_obj_1["lineNumber"] = 5
		hx_obj_1["className"] = hxrt.StringFromLiteral("LaneClean")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("run")
		return hx_obj_1
	}())
}
