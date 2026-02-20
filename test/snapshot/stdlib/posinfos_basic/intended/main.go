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
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral(".class=")), pos["className"].(*string)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral(".method=")), pos["methodName"].(*string)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral(".line=")), pos["lineNumber"].(int)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral(".fileNonEmpty=")), hxrt.StdString(!hxrt.StringEqualAny(pos["fileName"].(*string), hxrt.StringFromLiteral("")))))
}
