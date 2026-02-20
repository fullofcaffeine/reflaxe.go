package main

import "snapshot/hxrt"

func main() {
	factor := 3
	_ = factor
	mul := func(v int) int {
		return int(int32((hxrt.Int32Wrap(v) * hxrt.Int32Wrap(factor))))
	}
	_ = mul
	factor = 4
	hxrt.Println(mul(2))
}
