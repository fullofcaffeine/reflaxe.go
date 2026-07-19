package main

import "snapshot/hxrt"

func main() {
	manual := position(hxrt.StringFromLiteral("manual.hx"), 7, hxrt.NewArray(hxrt.StringFromLiteral("tail"), nil))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("format.plain="), haxe__Log_formatOutput(hxrt.StringFromLiteral("value"), nil)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("format.pos="), haxe__Log_formatOutput(hxrt.StringFromLiteral("value"), manual)))
	hxrt.Println(v_1)
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, hxrt.StringFromLiteral("default"), manual)
	original := haxe__Log_trace
	haxe__Log_trace = func(value any, infos map[string]any) {
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("custom.value="), hxrt.StdString(value)))
		hxrt.Println(v)
		var hx_if_7 int
		if (infos == nil) || (func(hx_obj_1 map[string]any) *hxrt.Array {
			hx_field_2 := hx_obj_1["customParams"]
			if hx_field_2 == nil {
				var hx_zero_3 *hxrt.Array
				return hx_zero_3
			}
			return hx_field_2.(*hxrt.Array)
		}(infos) == nil) {
			hx_if_7 = -1
		} else {
			hx_if_7 = func(hx_obj_4 map[string]any) *hxrt.Array {
				hx_field_5 := hx_obj_4["customParams"]
				if hx_field_5 == nil {
					var hx_zero_6 *hxrt.Array
					return hx_zero_6
				}
				return hx_field_5.(*hxrt.Array)
			}(infos).Len()
		}
		count := hx_if_7
		var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("custom.info="), hxrt.StdString((infos != nil))), hxrt.StringFromLiteral(":")), func(hx_obj_8 map[string]any) *string {
			hx_field_9 := hx_obj_8["className"]
			if hx_field_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_field_9.(*string)
		}(infos)), hxrt.StringFromLiteral(":")), func(hx_obj_11 map[string]any) *string {
			hx_field_12 := hx_obj_11["methodName"]
			if hx_field_12 == nil {
				var hx_zero_13 *string
				return hx_zero_13
			}
			return hx_field_12.(*string)
		}(infos)), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.StringEqualStringPtr(func(hx_obj_14 map[string]any) *string {
			hx_field_15 := hx_obj_14["fileName"]
			if hx_field_15 == nil {
				var hx_zero_16 *string
				return hx_zero_16
			}
			return hx_field_15.(*string)
		}(infos), hxrt.StringFromLiteral("Main.hx")))), hxrt.StringFromLiteral(":")), count))
		hxrt.Println(v_1)
	}
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, hxrt.StringFromLiteral("rebound"), func() map[string]any {
		hx_obj_17 := map[string]any{}
		hx_obj_17["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_17["lineNumber"] = 27
		hx_obj_17["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_17["methodName"] = hxrt.StringFromLiteral("main")
		hx_obj_17["customParams"] = hxrt.NewArray(hxrt.StringFromLiteral("tail"), nil)
		return hx_obj_17
	}())
	haxe__Log_trace = original
	direct := haxe__Log_trace
	direct(hxrt.StringFromLiteral("function.value"), nil)
	haxe__Log_trace = nil
	hxrt.TryCatch(func() {
		func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0, hx_arg_1)
		}(haxe__Log_trace, hxrt.StringFromLiteral("ignored"), nil)
		hxrt.Println(any(hxrt.StringFromLiteral("null=no-throw")))
	}, func(hx_caught_18 any) {
		hx_tmp := hx_caught_18
		_ = hx_tmp
		hxrt.Println(any(hxrt.StringFromLiteral("null=throws")))
	})
	haxe__Log_trace = original
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, hxrt.StringFromLiteral("restored"), nil)
}

func position(fileName *string, lineNumber int, customParams *hxrt.Array) map[string]any {
	hx_obj_20 := map[string]any{}
	hx_obj_20["fileName"] = fileName
	hx_obj_20["lineNumber"] = lineNumber
	hx_obj_20["className"] = hxrt.StringFromLiteral("Main")
	hx_obj_20["methodName"] = hxrt.StringFromLiteral("main")
	hx_obj_20["customParams"] = customParams
	return hx_obj_20
}
