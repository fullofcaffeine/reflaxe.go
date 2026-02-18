package main

import "snapshot/hxrt"

func main() {
	value := 3
	_ = value
	if value > 2 {
		hxrt.Println(hxrt.StringFromLiteral("gt"))
	} else {
		hxrt.Println(hxrt.StringFromLiteral("lte"))
	}
	if (value == 3) && true {
		hxrt.Println(hxrt.StringFromLiteral("yes"))
	} else {
		hxrt.Println(hxrt.StringFromLiteral("no"))
	}
}
