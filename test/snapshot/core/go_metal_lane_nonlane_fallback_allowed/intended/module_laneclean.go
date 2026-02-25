package main

import "snapshot/hxrt"

func LaneClean_run() {
	ok := go___Go_ok(hxrt.StringFromLiteral("clean"))
	hxrt.Println(func(hx_value_1 any) bool {
		if hx_value_1 == nil {
			var hx_zero_2 bool
			return hx_zero_2
		}
		return hx_value_1.(bool)
	}(ok.isOk()))
}
