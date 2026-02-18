package main

import "snapshot/hxrt"

func main() {
	a := (7 + (5 * 2))
	_ = a
	b := ((8.0 / 4.0) + 0.5)
	_ = b
	hxrt.Println(a)
	hxrt.Println(b)
}
