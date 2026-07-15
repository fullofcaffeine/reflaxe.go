package main

import "snapshot/hxrt"

func main() {
	//line Main.hx:4
	math := New_helper__LineMath()
	var v any = any(math.doubleIt(21))
	hxrt.Println(v)
}
