package main

import "snapshot/hxrt"

func main() {
	a := hxrt.StringFromLiteral("portable")
	b := hxrt.StringFromLiteral("nullx")
	c := true
	d := true
	hxrt.Println(any(a))
	hxrt.Println(any(b))
	hxrt.Println(any(c))
	hxrt.Println(any(d))
}
