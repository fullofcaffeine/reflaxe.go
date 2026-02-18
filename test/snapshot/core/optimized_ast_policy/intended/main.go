package main

import "snapshot/hxrt"

func main() {
	x := 17
	_ = x
	value := 3
	_ = value
	if value == 3 {
		hxrt.Println(x)
	} else {
		hxrt.Println(0)
	}
}
