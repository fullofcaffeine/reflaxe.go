package main

import "snapshot/hxrt"

func main() {
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, hxrt.StringFromLiteral("reports"), func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_1["lineNumber"] = 3
		hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_1
	}())
}
