package main

import (
	"math"
	"snapshot/hxrt"
)

func main() {
	x := 3.8
	var v any = any(hxrt.MathFloorInt(x))
	hxrt.Println(v)
	var v_1 any = any(hxrt.MathCeilInt(x))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.MathRoundInt(x))
	hxrt.Println(v_2)
	var v_3 any = any(math.Floor((x + 0.5)))
	hxrt.Println(v_3)
	var v_4 any = any(math.Abs(x))
	hxrt.Println(v_4)
	var v_5 any = any(Math_min(x, 2.1))
	hxrt.Println(v_5)
	var v_6 any = any(Math_max(x, 2.1))
	hxrt.Println(v_6)
}
