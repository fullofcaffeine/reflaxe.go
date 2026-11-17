package main

import "snapshot/hxrt"

func main() {
	//line Main.hx:4
	math_local := New_helper__LineMath()
	var v any = any(math_local.doubleIt(21))
	hxrt.Println(v)
}
