package main

import "snapshot/hxrt"

func main() {
	var left *string = nil
	_ = left
	right := hxrt.StringFromLiteral("value")
	_ = right
	a := hxrt.StringConcatAny(left, hxrt.StringFromLiteral("x"))
	_ = a
	b := hxrt.StringConcatAny(right, nil)
	_ = b
	c := hxrt.StringConcatAny(hxrt.StringFromLiteral("p"), 12)
	_ = c
	hxrt.Println(a)
	hxrt.Println(b)
	hxrt.Println(c)
}
