package main

import "snapshot/hxrt"

func main() {
	x := 17
	value := 3
	if value == 3 {
		hxrt.Println(x)
	} else {
		hxrt.Println(0)
	}
}
