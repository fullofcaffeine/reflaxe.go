package main

import "snapshot/hxrt"

func main() {
	add := func(a int, b int) int {
		return int(int32((int32(a) + int32(b))))
	}
	_ = add
	mul := func(v int) int {
		return int(int32((int32(v) * int32(3))))
	}
	_ = mul
	hxrt.Println(twice(5))
	hxrt.Println(add(2, 7))
	hxrt.Println(mul(4))
}

func twice(value int) int {
	return int(int32((int32(value) * int32(2))))
}
