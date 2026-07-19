package main

import "snapshot/hxrt"

func haxe__Log_formatOutput(v any, infos map[string]any) *string {
	str := hxrt.StdString(v)
	if infos == nil {
		return str
	}
	pstr := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func(hx_obj_11 map[string]any) *string {
		hx_field_12 := hx_obj_11["fileName"]
		if hx_field_12 == nil {
			var hx_zero_13 *string
			return hx_zero_13
		}
		return hx_field_12.(*string)
	}(infos), hxrt.StringFromLiteral(":")), func(hx_obj_14 map[string]any) int {
		hx_field_15 := hx_obj_14["lineNumber"]
		if hx_field_15 == nil {
			var hx_zero_16 int
			return hx_zero_16
		}
		return hx_field_15.(int)
	}(infos))
	if func(hx_obj_20 map[string]any) *hxrt.Array {
		hx_field_21 := hx_obj_20["customParams"]
		if hx_field_21 == nil {
			var hx_zero_22 *hxrt.Array
			return hx_zero_22
		}
		return hx_field_21.(*hxrt.Array)
	}(infos) != nil {
		_g := 0
		_g1 := func(hx_obj_17 map[string]any) *hxrt.Array {
			hx_field_18 := hx_obj_17["customParams"]
			if hx_field_18 == nil {
				var hx_zero_19 *hxrt.Array
				return hx_zero_19
			}
			return hx_field_18.(*hxrt.Array)
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
