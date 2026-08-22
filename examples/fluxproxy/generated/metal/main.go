package main

import "examples_fluxproxy_metal/hxrt"

func hasArg(flag *string) bool {
	_g := 0
	_g1 := hxrt.ArrayFromValues(func(hx_sort_src_1 []*string) []any {
		hx_sort_out_3 := make([]any, 0, len(hx_sort_src_1))
		for _, hx_sort_item_2 := range hx_sort_src_1 {
			hx_sort_out_3 = append(hx_sort_out_3, hx_sort_item_2)
		}
		return hx_sort_out_3
	}(hxrt.SysArgs()))
	for _g < _g1.Len() {
		arg := func(hx_value_4 any) *string {
			if hx_value_4 == nil {
				var hx_zero_5 *string
				return hx_zero_5
			}
			return hx_value_4.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(arg, flag) {
			return true
		}
	}
	return false
}

func main() {
	var runtime app__runtime__FluxRuntime = app__runtime__RuntimeFactory_create()
	if hasArg(hxrt.StringFromLiteral("--scripted")) {
		var v any = any(Harness_run(runtime))
		hxrt.Println(v)
	} else {
		InteractiveCli_run(runtime)
	}
}
