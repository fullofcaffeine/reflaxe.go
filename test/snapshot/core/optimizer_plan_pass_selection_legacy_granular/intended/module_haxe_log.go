package main

import "snapshot/hxrt"

func haxe__Log_formatOutput(v any, infos map[string]any) *string {
	str := hxrt.StdString(v)
	if infos == nil {
		return str
	}
	pstr := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func(hx_obj_2 map[string]any) *string {
		hx_field_3 := hx_obj_2["fileName"]
		if hx_field_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_field_3.(*string)
	}(infos), hxrt.StringFromLiteral(":")), func(hx_obj_5 map[string]any) int {
		hx_field_6 := hx_obj_5["lineNumber"]
		if hx_field_6 == nil {
			var hx_zero_7 int
			return hx_zero_7
		}
		return hx_field_6.(int)
	}(infos))
	if func(hx_obj_11 map[string]any) *hxrt.Array {
		hx_field_12 := hx_obj_11["customParams"]
		if hx_field_12 == nil {
			var hx_zero_13 *hxrt.Array
			return hx_zero_13
		}
		return hx_field_12.(*hxrt.Array)
	}(infos) != nil {
		_g := 0
		_g1 := func(hx_obj_8 map[string]any) *hxrt.Array {
			hx_field_9 := hx_obj_8["customParams"]
			if hx_field_9 == nil {
				var hx_zero_10 *hxrt.Array
				return hx_zero_10
			}
			return hx_field_9.(*hxrt.Array)
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
