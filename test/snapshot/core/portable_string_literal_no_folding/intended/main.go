package main

import "snapshot/hxrt"

func main() {
	a := hxrt.StringConcatAny(hxrt.StringFromLiteral("go"), hxrt.StringFromLiteral("pher"))
	_ = a
	b := hxrt.StringConcatAny(nil, hxrt.StringFromLiteral("x"))
	_ = b
	c := true
	_ = c
	d := true
	_ = d
	hxrt.Println(a)
	hxrt.Println(b)
	hxrt.Println(c)
	hxrt.Println(d)
}
