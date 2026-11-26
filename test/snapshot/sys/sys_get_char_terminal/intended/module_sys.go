package main

import "snapshot/hxrt"

func Sys_args() *hxrt.Array {
	return hxrt.ArrayFromValues(func(hx_sort_src_1 []*string) []any {
		hx_sort_out_3 := make([]any, 0, len(hx_sort_src_1))
		for _, hx_sort_item_2 := range hx_sort_src_1 {
			hx_sort_out_3 = append(hx_sort_out_3, hx_sort_item_2)
		}
		return hx_sort_out_3
	}(hxrt.SysArgs()))
}

func Sys_getChar(echo bool) int {
	value := hxrt.SysReadCharValue()
	if value < 0 {
		hxrt.Throw(haxe__Exception_thrown(New_haxe__io__Eof()))
	}
	if echo {
		New_sys__io__FileOutput(hxrt.SysStdout()).__hx_this.writeByte(value)
	}
	return value
}
