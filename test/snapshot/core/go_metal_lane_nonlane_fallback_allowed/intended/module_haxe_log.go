package main

import "snapshot/hxrt"

func haxe__Log_formatOutput(v any, infos map[string]any) *string {
	str := hxrt.StdString(v)
	if infos == nil {
		return str
	}
	pstr := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func(hx_obj_9 map[string]any) *string {
		hx_field_10 := hx_obj_9["fileName"]
		if hx_field_10 == nil {
			var hx_zero_11 *string
			return hx_zero_11
		}
		return hx_field_10.(*string)
	}(infos), hxrt.StringFromLiteral(":")), func(hx_obj_12 map[string]any) int {
		hx_field_13 := hx_obj_12["lineNumber"]
		if hx_field_13 == nil {
			var hx_zero_14 int
			return hx_zero_14
		}
		return hx_field_13.(int)
	}(infos))
	if func(hx_obj_18 map[string]any) *hxrt.Array {
		hx_field_19 := hx_obj_18["customParams"]
		if hx_field_19 == nil {
			var hx_zero_20 *hxrt.Array
			return hx_zero_20
		}
		return hx_field_19.(*hxrt.Array)
	}(infos) != nil {
		_g := 0
		_g1 := func(hx_obj_15 map[string]any) *hxrt.Array {
			hx_field_16 := hx_obj_15["customParams"]
			if hx_field_16 == nil {
				var hx_zero_17 *hxrt.Array
				return hx_zero_17
			}
			return hx_field_16.(*hxrt.Array)
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
