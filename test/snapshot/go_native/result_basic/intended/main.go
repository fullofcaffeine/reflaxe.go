package main

import "snapshot/hxrt"

func main() {
	ok := go___Result_ok(7)
	var v any = any(func(hx_value_1 any) bool {
		if hx_value_1 == nil {
			var hx_zero_2 bool
			return hx_zero_2
		}
		return hx_value_1.(bool)
	}(ok.isOk()))
	hxrt.Println(v)
	var v_1 any = any(func(hx_value_3 any) int {
		if hx_value_3 == nil {
			var hx_zero_4 int
			return hx_zero_4
		}
		return hx_value_3.(int)
	}(ok.unwrap()))
	hxrt.Println(v_1)
	err := go___Result_failure(hxrt.StringFromLiteral("boom"))
	var v_2 any = any(func(hx_value_5 any) bool {
		if hx_value_5 == nil {
			var hx_zero_6 bool
			return hx_zero_6
		}
		return hx_value_5.(bool)
	}(err.isErr()))
	hxrt.Println(v_2)
	var v_3 any = any(func(hx_value_7 any) *string {
		if hx_value_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_value_7.(*string)
	}(err.error()))
	hxrt.Println(v_3)
}
