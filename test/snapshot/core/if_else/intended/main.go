package main

import "snapshot/hxrt"

func main() {
	value := 3
	if value > 2 {
		hxrt.Println(any(hxrt.StringFromLiteral("gt")))
	} else {
		hxrt.Println(any(hxrt.StringFromLiteral("lte")))
	}
	if value == 3 {
		hxrt.Println(any(hxrt.StringFromLiteral("yes")))
	} else {
		hxrt.Println(any(hxrt.StringFromLiteral("no")))
	}
}
