package main

import "snapshot/hxrt"

func haxe__Log_formatOutput(v any, infos map[string]any) *string {
	str := hxrt.StdString(v)
	if infos == nil {
		return str
	}
	pstr := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func(hx_obj_21 map[string]any) *string {
		hx_field_22 := hx_obj_21["fileName"]
		if hx_field_22 == nil {
			var hx_zero_23 *string
			return hx_zero_23
		}
		return hx_field_22.(*string)
	}(infos), hxrt.StringFromLiteral(":")), func(hx_obj_24 map[string]any) int {
		hx_field_25 := hx_obj_24["lineNumber"]
		if hx_field_25 == nil {
			var hx_zero_26 int
			return hx_zero_26
		}
		return hx_field_25.(int)
	}(infos))
	if func(hx_obj_30 map[string]any) *hxrt.Array {
		hx_field_31 := hx_obj_30["customParams"]
		if hx_field_31 == nil {
			var hx_zero_32 *hxrt.Array
			return hx_zero_32
		}
		return hx_field_31.(*hxrt.Array)
	}(infos) != nil {
		_g := 0
		_g1 := func(hx_obj_27 map[string]any) *hxrt.Array {
			hx_field_28 := hx_obj_27["customParams"]
			if hx_field_28 == nil {
				var hx_zero_29 *hxrt.Array
				return hx_zero_29
			}
			return hx_field_28.(*hxrt.Array)
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
