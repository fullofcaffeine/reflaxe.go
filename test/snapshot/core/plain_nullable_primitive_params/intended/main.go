package main

import "snapshot/hxrt"

func boolBranch(value any) *string {
	if value != nil {
		var hx_if_1 *string
		if value.(bool) {
			hx_if_1 = hxrt.StringFromLiteral("true")
		} else {
			hx_if_1 = hxrt.StringFromLiteral("false")
		}
		return hx_if_1
	}
	return hxrt.StringFromLiteral("missing")
}

func floatBranch(value any) float64 {
	if value != nil {
		return (value.(float64) / 2.0)
	}
	return -1.0
}

func intBranch(value any) int {
	if value != nil {
		return int(int32((hxrt.Int32Wrap(value.(int)) + hxrt.Int32Wrap(1))))
	}
	return -1
}

func main() {
	var v any = any(intBranch(4))
	hxrt.Println(v)
	var v_1 any = any(intBranch(nil))
	hxrt.Println(v_1)
	var v_2 any = any(floatBranch(5.0))
	hxrt.Println(v_2)
	var v_3 any = any(floatBranch(nil))
	hxrt.Println(v_3)
	var v_4 any = any(boolBranch(true))
	hxrt.Println(v_4)
	var v_5 any = any(boolBranch(nil))
	hxrt.Println(v_5)
}
