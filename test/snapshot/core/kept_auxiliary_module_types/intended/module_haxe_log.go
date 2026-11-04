package main

import "snapshot/hxrt"

func haxe__Log_formatOutput(v any, infos hxrt.Anonymous) *string {
	str := hxrt.StdString(v)
	if infos == nil {
		return str
	}
	pstr := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(hxrt.AnonymousGet(infos, "fileName")), hxrt.StringFromLiteral(":")), func(hx_value_3 any) int {
		if hx_value_3 == nil {
			var hx_zero_4 int
			return hx_zero_4
		}
		return hx_value_3.(int)
	}(hxrt.AnonymousGet(infos, "lineNumber")))
	if func(hx_value_7 any) *hxrt.Array {
		if hx_value_7 == nil {
			var hx_zero_8 *hxrt.Array
			return hx_zero_8
		}
		return hx_value_7.(*hxrt.Array)
	}(hxrt.AnonymousGet(infos, "customParams")) != nil {
		_g := 0
		_g1 := func(hx_value_5 any) *hxrt.Array {
			if hx_value_5 == nil {
				var hx_zero_6 *hxrt.Array
				return hx_zero_6
			}
			return hx_value_5.(*hxrt.Array)
		}(hxrt.AnonymousGet(infos, "customParams"))
		for _g < _g1.Len() {
			var v_1 any = _g1.Get(_g)
			_g = int(int32((_g + 1)))
			str = hxrt.StringConcatStringPtr(str, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(", "), hxrt.StdString(v_1)))
		}
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(pstr, hxrt.StringFromLiteral(": ")), str)
}

var haxe__Log_trace *hxrt.Function[func(any, hxrt.Anonymous)] = hxrt.NewFunction(func(v any, infos hxrt.Anonymous) {
	str := haxe__Log_formatOutput(v, infos)
	hxrt.Println(any(str))
})
