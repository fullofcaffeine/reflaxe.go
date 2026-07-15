package main

import "snapshot/hxrt"

func main() {
	a := 17
	b := 2.5
	hxrt.Println(any(a))
	hxrt.Println(any(b))
}
