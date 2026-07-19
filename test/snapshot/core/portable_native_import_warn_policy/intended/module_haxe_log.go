package main

import "snapshot/hxrt"

func haxe__Log_formatOutput(v any, infos map[string]any) *string {
	str := hxrt.StdString(v)
	if infos == nil {
		return str
	}
	pstr := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func(hx_obj_7 map[string]any) *string {
		hx_field_8 := hx_obj_7["fileName"]
		if hx_field_8 == nil {
			var hx_zero_9 *string
			return hx_zero_9
		}
		return hx_field_8.(*string)
	}(infos), hxrt.StringFromLiteral(":")), func(hx_obj_10 map[string]any) int {
		hx_field_11 := hx_obj_10["lineNumber"]
		if hx_field_11 == nil {
			var hx_zero_12 int
			return hx_zero_12
		}
		return hx_field_11.(int)
	}(infos))
	if func(hx_obj_16 map[string]any) *hxrt.Array {
		hx_field_17 := hx_obj_16["customParams"]
		if hx_field_17 == nil {
			var hx_zero_18 *hxrt.Array
			return hx_zero_18
		}
		return hx_field_17.(*hxrt.Array)
	}(infos) != nil {
		_g := 0
		_g1 := func(hx_obj_13 map[string]any) *hxrt.Array {
			hx_field_14 := hx_obj_13["customParams"]
			if hx_field_14 == nil {
				var hx_zero_15 *hxrt.Array
				return hx_zero_15
			}
			return hx_field_14.(*hxrt.Array)
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
