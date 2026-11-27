package main

import "snapshot/hxrt"

func main() {
	var d any = 1
	var v any = any(hxrt.IsInt(d))
	hxrt.Println(v)
	var v_1 any = any(func(hx_value any) bool {
		switch hx_value.(type) {
		case *string:
			return true
		case string:
			return true
		default:
			return false
		}
	}(d))
	hxrt.Println(v_1)
}
