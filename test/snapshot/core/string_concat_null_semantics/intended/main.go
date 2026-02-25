package main

import "snapshot/hxrt"

func main() {
	var left *string = nil
	_ = left
	right := hxrt.StringFromLiteral("value")
	_ = right
	a := hxrt.StringConcatStringPtr(left, hxrt.StringFromLiteral("x"))
	_ = a
	b := hxrt.StringConcatStringPtr(right, nil)
	_ = b
	c := hxrt.StringFromLiteral("p12")
	_ = c
	hxrt.Println(a)
	hxrt.Println(b)
	hxrt.Println(c)
}
