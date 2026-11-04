package main

import "snapshot/hxrt"

func main() {
	read(func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_1["lineNumber"] = 15
		hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
		return hxrt.AnonymousFromMap(hx_obj_1)
	}())
}

func read(pos hxrt.Anonymous) {
	custom := func(hx_value_2 any) *hxrt.Array {
		if hx_value_2 == nil {
			var hx_zero_3 *hxrt.Array
			return hx_zero_3
		}
		return hx_value_2.(*hxrt.Array)
	}(hxrt.AnonymousGet(pos, "customParams"))
	count := 0
	if custom != nil {
		count = custom.Len()
	}
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("count="), count)))
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("line="), func(hx_value_4 any) int {
		if hx_value_4 == nil {
			var hx_zero_5 int
			return hx_zero_5
		}
		return hx_value_4.(int)
	}(hxrt.AnonymousGet(pos, "lineNumber"))))
	hxrt.Println(v)
}
