package main

import "snapshot/hxrt"

func main() {
	add := func(a int, b int) int {
		return int(int32((hxrt.Int32Wrap(a) + hxrt.Int32Wrap(b))))
	}
	_ = add
	mul := func(v int) int {
		return int(int32((hxrt.Int32Wrap(v) * hxrt.Int32Wrap(3))))
	}
	_ = mul
	hxrt.Println(twice(5))
	hxrt.Println(add(2, 7))
	hxrt.Println(mul(4))
}

func twice(value int) int {
	return int(int32((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(2))))
}
