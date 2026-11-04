package main

import "snapshot/hxrt"

func main() {
	func() {
		var hx_callable_2 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_3 any = hxrt.StringFromLiteral("reports")
		var hx_argument_4 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_1 := map[string]any{}
			hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_1["lineNumber"] = 3
			hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
			return hxrt.AnonymousFromMap(hx_obj_1)
		}()
		hxrt.FunctionCall(hx_callable_2)(hx_argument_3, hx_argument_4)
	}()
}
