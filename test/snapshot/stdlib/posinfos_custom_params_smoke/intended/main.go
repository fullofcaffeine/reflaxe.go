package main

import "snapshot/hxrt"

func main() {
	read(func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_1["lineNumber"] = 15
		hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_1
	}())
}

func read(pos map[string]any) {
	custom := func(hx_obj_2 map[string]any) []any {
		hx_field_3 := hx_obj_2["customParams"]
		if hx_field_3 == nil {
			var hx_zero_4 []any
			return hx_zero_4
		}
		return hx_field_3.([]any)
	}(pos)
	_ = custom
	count := 0
	_ = count
	if !hxrt.StringEqualAny(custom, nil) {
		count = len(custom)
	}
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("count="), count))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("line="), func(hx_obj_5 map[string]any) int {
		hx_field_6 := hx_obj_5["lineNumber"]
		if hx_field_6 == nil {
			var hx_zero_7 int
			return hx_zero_7
		}
		return hx_field_6.(int)
	}(pos)))
}
