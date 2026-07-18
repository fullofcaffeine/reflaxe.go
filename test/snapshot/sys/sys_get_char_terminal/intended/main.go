package main

import "snapshot/hxrt"

func main() {
	arguments := hxrt.ArrayFromValues(func(hx_sort_src_1 []*string) []any {
		hx_sort_out_3 := make([]any, 0, len(hx_sort_src_1))
		for _, hx_sort_item_2 := range hx_sort_src_1 {
			hx_sort_out_3 = append(hx_sort_out_3, hx_sort_item_2)
		}
		return hx_sort_out_3
	}(hxrt.SysArgs()))
	echo := ((arguments.Len() > 0) && hxrt.StringEqualAny(arguments.Get(0), hxrt.StringFromLiteral("echo")))
	hxrt.Print(any(hxrt.StringFromLiteral("ready|")))
	hxrt.TryCatch(func() {
		value := Sys_getChar(echo)
		hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("|"), value), hxrt.StringFromLiteral("|"))))
	}, func(hx_caught_4 any) {
		switch hx_typed_5 := hx_caught_4.(type) {
		case *haxe__io__Eof:
			hx_tmp := hx_typed_5
			_ = hx_tmp
			hxrt.Println(any(hxrt.StringFromLiteral("eof|")))
		default:
			hxrt.Throw(hx_caught_4)
		}
	})
}
