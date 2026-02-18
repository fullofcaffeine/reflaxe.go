package main

import "snapshot/hxrt"

func main() {
	a := 17
	_ = a
	b := 2.5
	_ = b
	hxrt.Println(a)
	hxrt.Println(b)
}
