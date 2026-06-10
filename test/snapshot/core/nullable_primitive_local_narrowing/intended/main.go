package main

import "snapshot/hxrt"

func boolBranch(value any) *string {
	if value != nil {
		narrowed := value.(bool)
		var hx_if_1 *string
		if narrowed {
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
		narrowed := value.(float64)
		return (narrowed / 2.0)
	}
	return -1.0
}

func intBranch(value any) int {
	if value != nil {
		narrowed := value.(int)
		return int(int32((hxrt.Int32Wrap(narrowed) + hxrt.Int32Wrap(1))))
	}
	return -1
}

func main() {
	hxrt.Println(intBranch(4))
	hxrt.Println(intBranch(nil))
	hxrt.Println(floatBranch(5.0))
	hxrt.Println(floatBranch(nil))
	hxrt.Println(boolBranch(true))
	hxrt.Println(boolBranch(nil))
}
