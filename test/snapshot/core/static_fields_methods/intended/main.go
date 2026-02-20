package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(MathBox_mul(4))
	MathBox_factor = 5
	hxrt.Println(MathBox_mul(4))
}

var MathBox_factor int = 3

func MathBox_mul(value int) int {
	return int(int32((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(MathBox_factor))))
}
