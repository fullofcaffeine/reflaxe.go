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
	custom := func(hx_obj_2 map[string]any) *hxrt.Array {
		hx_field_3 := hx_obj_2["customParams"]
		if hx_field_3 == nil {
			var hx_zero_4 *hxrt.Array
			return hx_zero_4
		}
		return hx_field_3.(*hxrt.Array)
	}(pos)
	count := 0
	if custom != nil {
		count = custom.Len()
	}
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("count="), count)))
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("line="), func(hx_obj_5 map[string]any) int {
		hx_field_6 := hx_obj_5["lineNumber"]
		if hx_field_6 == nil {
			var hx_zero_7 int
			return hx_zero_7
		}
		return hx_field_6.(int)
	}(pos)))
	hxrt.Println(v)
}
