package main

import "snapshot/hxrt"

func haxe__Log_formatOutput(v any, infos map[string]any) *string {
	str := hxrt.StdString(v)
	if infos == nil {
		return str
	}
	pstr := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func(hx_obj_34 map[string]any) *string {
		hx_field_35 := hx_obj_34["fileName"]
		if hx_field_35 == nil {
			var hx_zero_36 *string
			return hx_zero_36
		}
		return hx_field_35.(*string)
	}(infos), hxrt.StringFromLiteral(":")), func(hx_obj_37 map[string]any) int {
		hx_field_38 := hx_obj_37["lineNumber"]
		if hx_field_38 == nil {
			var hx_zero_39 int
			return hx_zero_39
		}
		return hx_field_38.(int)
	}(infos))
	if func(hx_obj_43 map[string]any) *hxrt.Array {
		hx_field_44 := hx_obj_43["customParams"]
		if hx_field_44 == nil {
			var hx_zero_45 *hxrt.Array
			return hx_zero_45
		}
		return hx_field_44.(*hxrt.Array)
	}(infos) != nil {
		_g := 0
		_g1 := func(hx_obj_40 map[string]any) *hxrt.Array {
			hx_field_41 := hx_obj_40["customParams"]
			if hx_field_41 == nil {
				var hx_zero_42 *hxrt.Array
				return hx_zero_42
			}
			return hx_field_41.(*hxrt.Array)
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
