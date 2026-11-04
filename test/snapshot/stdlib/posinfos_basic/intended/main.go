package main

import "snapshot/hxrt"

func main() {
	probe(hxrt.StringFromLiteral("main"), func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_1["lineNumber"] = 12
		hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
		return hxrt.AnonymousFromMap(hx_obj_1)
	}())
	fn := hxrt.NewFunction(func() {
		probe(hxrt.StringFromLiteral("closure"), func() hxrt.Anonymous {
			hx_obj_2 := map[string]any{}
			hx_obj_2["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_2["lineNumber"] = 14
			hx_obj_2["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_2["methodName"] = hxrt.StringFromLiteral("main")
			return hxrt.AnonymousFromMap(hx_obj_2)
		}())
	})
	func() {
		var hx_callable_3 *hxrt.Function[func()] = fn
		hxrt.FunctionCall(hx_callable_3)()
	}()
}

func probe(label *string, pos hxrt.Anonymous) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(".class=")), func(hx_value_4 any) *string {
		if hx_value_4 == nil {
			var hx_zero_5 *string
			return hx_zero_5
		}
		return hx_value_4.(*string)
	}(hxrt.AnonymousGet(pos, "className"))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(".method=")), func(hx_value_6 any) *string {
		if hx_value_6 == nil {
			var hx_zero_7 *string
			return hx_zero_7
		}
		return hx_value_6.(*string)
	}(hxrt.AnonymousGet(pos, "methodName"))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(".line=")), func(hx_value_8 any) int {
		if hx_value_8 == nil {
			var hx_zero_9 int
			return hx_zero_9
		}
		return hx_value_8.(int)
	}(hxrt.AnonymousGet(pos, "lineNumber"))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(".fileNonEmpty=")), hxrt.StdString(!hxrt.StringEqualStringPtr(func(hx_value_10 any) *string {
		if hx_value_10 == nil {
			var hx_zero_11 *string
			return hx_zero_11
		}
		return hx_value_10.(*string)
	}(hxrt.AnonymousGet(pos, "fileName")), hxrt.StringFromLiteral("")))))
	hxrt.Println(v_3)
}
