package main

import "snapshot/hxrt"

func main() {
	var dynamicNil any = nil
	_ = dynamicNil
	if dynamicNil == nil {
		hxrt.Println(hxrt.StringFromLiteral("ident:nil"))
	} else {
		hxrt.Println(hxrt.StringFromLiteral("ident:non_nil"))
	}
}
