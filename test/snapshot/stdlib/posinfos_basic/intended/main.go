package main

import "snapshot/hxrt"

func main() {
	probe(hxrt.StringFromLiteral("main"), func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_1["lineNumber"] = 12
		hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_1
	}())
	fn := func() {
		probe(hxrt.StringFromLiteral("closure"), func() map[string]any {
			hx_obj_2 := map[string]any{}
			hx_obj_2["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_2["lineNumber"] = 14
			hx_obj_2["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_2["methodName"] = hxrt.StringFromLiteral("main")
			return hx_obj_2
		}())
	}
	_ = fn
	fn()
}

func probe(label *string, pos map[string]any) {
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral(".class=")), func(hx_obj_3 map[string]any) *string {
		hx_field_4 := hx_obj_3["className"]
		if hx_field_4 == nil {
			var hx_zero_5 *string
			return hx_zero_5
		}
		return hx_field_4.(*string)
	}(pos)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral(".method=")), func(hx_obj_6 map[string]any) *string {
		hx_field_7 := hx_obj_6["methodName"]
		if hx_field_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_field_7.(*string)
	}(pos)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral(".line=")), func(hx_obj_9 map[string]any) int {
		hx_field_10 := hx_obj_9["lineNumber"]
		if hx_field_10 == nil {
			var hx_zero_11 int
			return hx_zero_11
		}
		return hx_field_10.(int)
	}(pos)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral(".fileNonEmpty=")), hxrt.StdString(!hxrt.StringEqualAny(func(hx_obj_12 map[string]any) *string {
		hx_field_13 := hx_obj_12["fileName"]
		if hx_field_13 == nil {
			var hx_zero_14 *string
			return hx_zero_14
		}
		return hx_field_13.(*string)
	}(pos), hxrt.StringFromLiteral("")))))
}
