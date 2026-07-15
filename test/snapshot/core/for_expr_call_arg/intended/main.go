package main

import "snapshot/hxrt"

func consume(v any) {
}

func main() {
	consume(func() any {
		x := 0
		var v any = any(hxrt.StdString(x))
		hxrt.Println(v)
		x_1 := 1
		var v_1 any = any(hxrt.StdString(x_1))
		hxrt.Println(v_1)
		x_2 := 2
		var v_2 any = any(hxrt.StdString(x_2))
		hxrt.Println(v_2)
		return nil
	}())
	hxrt.Println(any(hxrt.StringFromLiteral("ok")))
}
