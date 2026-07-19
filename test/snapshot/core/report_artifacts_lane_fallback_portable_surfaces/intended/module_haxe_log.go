package main

import "snapshot/hxrt"

func haxe__Log_formatOutput(v any, infos map[string]any) *string {
	str := hxrt.StdString(v)
	if infos == nil {
		return str
	}
	pstr := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func(hx_obj_29 map[string]any) *string {
		hx_field_30 := hx_obj_29["fileName"]
		if hx_field_30 == nil {
			var hx_zero_31 *string
			return hx_zero_31
		}
		return hx_field_30.(*string)
	}(infos), hxrt.StringFromLiteral(":")), func(hx_obj_32 map[string]any) int {
		hx_field_33 := hx_obj_32["lineNumber"]
		if hx_field_33 == nil {
			var hx_zero_34 int
			return hx_zero_34
		}
		return hx_field_33.(int)
	}(infos))
	if func(hx_obj_38 map[string]any) *hxrt.Array {
		hx_field_39 := hx_obj_38["customParams"]
		if hx_field_39 == nil {
			var hx_zero_40 *hxrt.Array
			return hx_zero_40
		}
		return hx_field_39.(*hxrt.Array)
	}(infos) != nil {
		_g := 0
		_g1 := func(hx_obj_35 map[string]any) *hxrt.Array {
			hx_field_36 := hx_obj_35["customParams"]
			if hx_field_36 == nil {
				var hx_zero_37 *hxrt.Array
				return hx_zero_37
			}
			return hx_field_36.(*hxrt.Array)
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
