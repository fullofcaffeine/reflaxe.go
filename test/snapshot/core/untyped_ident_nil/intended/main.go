package main

import "snapshot/hxrt"

func main() {
	var dynamicNil any = nil
	if hxrt.AnyEqualsNull(dynamicNil) {
		hxrt.Println(any(hxrt.StringFromLiteral("ident:nil")))
	} else {
		hxrt.Println(any(hxrt.StringFromLiteral("ident:non_nil")))
	}
}
