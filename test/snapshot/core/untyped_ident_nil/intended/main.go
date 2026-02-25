package main

import "snapshot/hxrt"

func main() {
	var dynamicNil any = nil
	if hxrt.AnyEqualsNull(dynamicNil) {
		hxrt.Println(hxrt.StringFromLiteral("ident:nil"))
	} else {
		hxrt.Println(hxrt.StringFromLiteral("ident:non_nil"))
	}
}
