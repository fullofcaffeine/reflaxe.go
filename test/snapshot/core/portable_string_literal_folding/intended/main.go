package main

import "snapshot/hxrt"

func main() {
	a := hxrt.StringFromLiteral("portable")
	b := hxrt.StringFromLiteral("nullx")
	c := true
	d := true
	hxrt.Println(a)
	hxrt.Println(b)
	hxrt.Println(c)
	hxrt.Println(d)
}
