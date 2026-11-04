package main

import "snapshot/hxrt"

func main() {
	manual := position(hxrt.StringFromLiteral("manual.hx"), 7, hxrt.NewArray(hxrt.StringFromLiteral("tail"), nil))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("format.plain="), haxe__Log_formatOutput(hxrt.StringFromLiteral("value"), nil)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("format.pos="), haxe__Log_formatOutput(hxrt.StringFromLiteral("value"), manual)))
	hxrt.Println(v_1)
	func() {
		var hx_callable_1 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_2 any = hxrt.StringFromLiteral("default")
		var hx_argument_3 hxrt.Anonymous = manual
		hxrt.FunctionCall(hx_callable_1)(hx_argument_2, hx_argument_3)
	}()
	original := haxe__Log_trace
	haxe__Log_trace = hxrt.NewFunction(func(value any, infos hxrt.Anonymous) {
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("custom.value="), hxrt.StdString(value)))
		hxrt.Println(v)
		var hx_if_8 int
		if (infos == nil) || (func(hx_value_4 any) *hxrt.Array {
			if hx_value_4 == nil {
				var hx_zero_5 *hxrt.Array
				return hx_zero_5
			}
			return hx_value_4.(*hxrt.Array)
		}(hxrt.AnonymousGet(infos, "customParams")) == nil) {
			hx_if_8 = -1
		} else {
			hx_if_8 = func(hx_value_6 any) *hxrt.Array {
				if hx_value_6 == nil {
					var hx_zero_7 *hxrt.Array
					return hx_zero_7
				}
				return hx_value_6.(*hxrt.Array)
			}(hxrt.AnonymousGet(infos, "customParams")).Len()
		}
		count := hx_if_8
		var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("custom.info="), hxrt.StdString((infos != nil))), hxrt.StringFromLiteral(":")), func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
		}(hxrt.AnonymousGet(infos, "className"))), hxrt.StringFromLiteral(":")), func(hx_value_11 any) *string {
			if hx_value_11 == nil {
				var hx_zero_12 *string
				return hx_zero_12
			}
			return hx_value_11.(*string)
		}(hxrt.AnonymousGet(infos, "methodName"))), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.StringEqualStringPtr(func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(hxrt.AnonymousGet(infos, "fileName")), hxrt.StringFromLiteral("Main.hx")))), hxrt.StringFromLiteral(":")), count))
		hxrt.Println(v_1)
	})
	func() {
		var hx_callable_16 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_17 any = hxrt.StringFromLiteral("rebound")
		var hx_argument_18 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_15 := map[string]any{}
			hx_obj_15["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_15["lineNumber"] = 35
			hx_obj_15["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_15["methodName"] = hxrt.StringFromLiteral("main")
			hx_obj_15["customParams"] = hxrt.NewArray(hxrt.StringFromLiteral("tail"), nil)
			return hxrt.AnonymousFromMap(hx_obj_15)
		}()
		hxrt.FunctionCall(hx_callable_16)(hx_argument_17, hx_argument_18)
	}()
	haxe__Log_trace = original
	direct := haxe__Log_trace
	func() {
		var hx_callable_19 *hxrt.Function[func(any, hxrt.Anonymous)] = direct
		var hx_argument_20 any = hxrt.StringFromLiteral("function.value")
		var hx_argument_21 hxrt.Anonymous = nil
		hxrt.FunctionCall(hx_callable_19)(hx_argument_20, hx_argument_21)
	}()
	haxe__Log_trace = nil
	hxrt.TryCatch(func() {
		func() {
			var hx_callable_25 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
			var hx_argument_26 any = hxrt.StringFromLiteral("ignored")
			var hx_argument_27 hxrt.Anonymous = nil
			hxrt.FunctionCall(hx_callable_25)(hx_argument_26, hx_argument_27)
		}()
		hxrt.Println(any(hxrt.StringFromLiteral("null=no-throw")))
	}, func(hx_caught_22 any) {
		hx_exception_23 := haxe__Exception_caught(hx_caught_22)
		hx_payload_24 := haxe__Exception_catchValue(hx_caught_22)
		_ = hx_exception_23
		var hx_tmp any = hx_payload_24
		_ = hx_tmp
		hxrt.Println(any(hxrt.StringFromLiteral("null=throws")))
	})
	haxe__Log_trace = original
	func() {
		var hx_callable_28 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_29 any = hxrt.StringFromLiteral("restored")
		var hx_argument_30 hxrt.Anonymous = nil
		hxrt.FunctionCall(hx_callable_28)(hx_argument_29, hx_argument_30)
	}()
}

func position(fileName *string, lineNumber int, customParams *hxrt.Array) hxrt.Anonymous {
	hx_obj_31 := map[string]any{}
	hx_obj_31["fileName"] = fileName
	hx_obj_31["lineNumber"] = lineNumber
	hx_obj_31["className"] = hxrt.StringFromLiteral("Main")
	hx_obj_31["methodName"] = hxrt.StringFromLiteral("main")
	hx_obj_31["customParams"] = customParams
	return hxrt.AnonymousFromMap(hx_obj_31)
}

func hx_recover_class_0(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_1(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
