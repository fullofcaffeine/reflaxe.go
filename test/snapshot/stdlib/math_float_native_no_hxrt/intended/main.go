package main

import (
	"math"
	"snapshot/hxrt"
)

func main() {
	var v any = any(math.Sqrt(81.0))
	hxrt.Println(v)
	var v_1 any = any((!math.IsInf(9.0, 0) && !math.IsNaN(9.0)))
	hxrt.Println(v_1)
}
