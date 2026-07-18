package main

import "examples_tui_todo_portable/hxrt"

func hasArg(flag *string) bool {
	_g := 0
	_g1 := hxrt.ArrayFromValues(func(hx_sort_src_57 []*string) []any {
		hx_sort_out_59 := make([]any, 0, len(hx_sort_src_57))
		for _, hx_sort_item_58 := range hx_sort_src_57 {
			hx_sort_out_59 = append(hx_sort_out_59, hx_sort_item_58)
		}
		return hx_sort_out_59
	}(hxrt.SysArgs()))
	for _g < _g1.Len() {
		arg := func(hx_value_60 any) *string {
			if hx_value_60 == nil {
				var hx_zero_61 *string
				return hx_zero_61
			}
			return hx_value_60.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(arg, flag) {
			return true
		}
	}
	return false
}

func main() {
	var runtime profile__TodoRuntime = profile__RuntimeFactory_create()
	if hasArg(hxrt.StringFromLiteral("--scripted")) {
		var v any = any(Harness_run(runtime))
		hxrt.Println(v)
	} else {
		InteractiveCli_run(runtime)
	}
}
