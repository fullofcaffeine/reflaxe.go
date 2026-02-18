package main

import "snapshot/hxrt"

func main() {
	add := func(a int, b int) int {
		return (a + b)
	}
	_ = add
	mul := func(v int) int {
		return (v * 3)
	}
	_ = mul
	hxrt.Println(twice(5))
	hxrt.Println(add(2, 7))
	hxrt.Println(mul(4))
}

func twice(value int) int {
	return (value * 2)
}
