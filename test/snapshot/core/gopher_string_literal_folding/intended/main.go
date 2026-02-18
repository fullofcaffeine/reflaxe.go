package main

import "snapshot/hxrt"

func main() {
	a := hxrt.StringFromLiteral("gopher")
	_ = a
	b := hxrt.StringFromLiteral("nullx")
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
