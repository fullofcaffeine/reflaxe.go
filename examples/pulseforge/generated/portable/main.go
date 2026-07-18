package main

import "examples_pulseforge_portable/hxrt"

func hasArg(flag *string) bool {
	_g := 0
	_g1 := hxrt.ArrayFromValues(func(hx_sort_src_21 []*string) []any {
		hx_sort_out_23 := make([]any, 0, len(hx_sort_src_21))
		for _, hx_sort_item_22 := range hx_sort_src_21 {
			hx_sort_out_23 = append(hx_sort_out_23, hx_sort_item_22)
		}
		return hx_sort_out_23
	}(hxrt.SysArgs()))
	for _g < _g1.Len() {
		arg := func(hx_value_24 any) *string {
			if hx_value_24 == nil {
				var hx_zero_25 *string
				return hx_zero_25
			}
			return hx_value_24.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(arg, flag) {
			return true
		}
	}
	return false
}

func main() {
	var runtime app__runtime__PulseRuntime = app__runtime__RuntimeFactory_create()
	if hasArg(hxrt.StringFromLiteral("--scripted")) {
		var v any = any(Harness_run(runtime))
		hxrt.Println(v)
	} else {
		InteractiveCli_run(runtime)
	}
}
