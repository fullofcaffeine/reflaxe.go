package main

import "snapshot/hxrt"

func main() {
	names := haxe__Resource_listNames()
	var v any = any(names.Len())
	hxrt.Println(v)
	var hx_if_1 any
	if names.Len() > 0 {
		hx_if_1 = names.Get(0)
	} else {
		hx_if_1 = hxrt.StringFromLiteral("<none>")
	}
	var v_1 any = hx_if_1
	hxrt.Println(v_1)
	var v_2 any = any(StringTools_replace(haxe__Resource_getString(hxrt.StringFromLiteral("greet")), hxrt.StringFromLiteral("\n"), hxrt.StringFromLiteral("\\n")))
	hxrt.Println(v_2)
	bytes := haxe__Resource_getBytes(hxrt.StringFromLiteral("greet"))
	var hx_if_2 any
	if bytes == nil {
		hx_if_2 = hxrt.StringFromLiteral("null")
	} else {
		hx_if_2 = bytes.length
	}
	var v_3 any = hx_if_2
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringEqualStringPtr(haxe__Resource_getString(hxrt.StringFromLiteral("missing")), nil))
	hxrt.Println(v_4)
	var v_5 any = any((haxe__Resource_getBytes(hxrt.StringFromLiteral("missing")) == nil))
	hxrt.Println(v_5)
}
