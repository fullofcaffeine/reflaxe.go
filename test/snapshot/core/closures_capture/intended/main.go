package main

import "snapshot/hxrt"

func main() {
	factor := 3
	mul := func(v int) int {
		return (v * factor)
	}
	factor = 4
	hxrt.Println(mul(2))
}
