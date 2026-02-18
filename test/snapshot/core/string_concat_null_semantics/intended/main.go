package main

import "snapshot/hxrt"

func main() {
	var left *string = nil
	right := hxrt.StringFromLiteral("value")
	a := hxrt.StringConcatAny(left, hxrt.StringFromLiteral("x"))
	b := hxrt.StringConcatAny(right, nil)
	c := hxrt.StringConcatAny(hxrt.StringFromLiteral("p"), 12)
	hxrt.Println(a)
	hxrt.Println(b)
	hxrt.Println(c)
}
