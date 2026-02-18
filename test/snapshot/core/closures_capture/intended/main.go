package main

import "snapshot/hxrt"

func main() {
	factor := 3
	_ = factor
	mul := func(v int) int {
		return (v * factor)
	}
	_ = mul
	factor = 4
	hxrt.Println(mul(2))
}
