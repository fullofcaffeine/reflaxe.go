package main

import "snapshot/hxrt"

func main() {
	factor := 3
	mul := func(v int) int {
		return int((hxrt.Int32Wrap(v) * hxrt.Int32Wrap(factor)))
	}
	factor = 4
	var v any = any(mul(2))
	hxrt.Println(v)
}
