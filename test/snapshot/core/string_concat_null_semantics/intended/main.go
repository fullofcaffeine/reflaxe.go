package main

import "snapshot/hxrt"

func main() {
	var left *string = nil
	right := hxrt.StringFromLiteral("value")
	a := hxrt.StringConcatStringPtr(left, hxrt.StringFromLiteral("x"))
	b := hxrt.StringConcatStringPtr(right, nil)
	c := hxrt.StringFromLiteral("p12")
	hxrt.Println(a)
	hxrt.Println(b)
	hxrt.Println(c)
}
