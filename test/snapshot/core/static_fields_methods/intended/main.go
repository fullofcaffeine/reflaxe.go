package main

import "snapshot/hxrt"

func main() {
	var v any = any(MathBox_mul(4))
	hxrt.Println(v)
	MathBox_factor = 5
	var v_1 any = any(MathBox_mul(4))
	hxrt.Println(v_1)
}

var MathBox_factor int = 3

func MathBox_mul(value int) int {
	return int(int32((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(MathBox_factor))))
}
