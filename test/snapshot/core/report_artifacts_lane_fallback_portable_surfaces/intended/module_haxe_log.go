package main

import "snapshot/hxrt"

func haxe__Log_formatOutput(v any, infos map[string]any) *string {
	str := hxrt.StdString(v)
	if infos == nil {
		return str
	}
	pstr := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func(hx_obj_1 map[string]any) *string {
		hx_field_2 := hx_obj_1["fileName"]
		if hx_field_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_field_2.(*string)
	}(infos), hxrt.StringFromLiteral(":")), func(hx_obj_4 map[string]any) int {
		hx_field_5 := hx_obj_4["lineNumber"]
		if hx_field_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_field_5.(int)
	}(infos))
	if func(hx_obj_10 map[string]any) *hxrt.Array {
		hx_field_11 := hx_obj_10["customParams"]
		if hx_field_11 == nil {
			var hx_zero_12 *hxrt.Array
			return hx_zero_12
		}
		return hx_field_11.(*hxrt.Array)
	}(infos) != nil {
		_g := 0
		_g1 := func(hx_obj_7 map[string]any) *hxrt.Array {
			hx_field_8 := hx_obj_7["customParams"]
			if hx_field_8 == nil {
				var hx_zero_9 *hxrt.Array
				return hx_zero_9
			}
			return hx_field_8.(*hxrt.Array)
		}(infos)
		for _g < _g1.Len() {
			var v_1 any = _g1.Get(_g)
			_g = int(int32((_g + 1)))
			str = hxrt.StringConcatStringPtr(str, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(", "), hxrt.StdString(v_1)))
		}
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(pstr, hxrt.StringFromLiteral(": ")), str)
}

var haxe__Log_trace func(any, map[string]any) = func(v any, infos map[string]any) {
	str := haxe__Log_formatOutput(v, infos)
	hxrt.Println(any(str))
}
