package main

import "snapshot/hxrt"

func consume(v any) {
}

func main() {
	consume(func() any {
		x := 0
		hxrt.Println(hxrt.StdString(x))
		x_1 := 1
		hxrt.Println(hxrt.StdString(x_1))
		x_2 := 2
		hxrt.Println(hxrt.StdString(x_2))
		return nil
	}())
	hxrt.Println(hxrt.StringFromLiteral("ok"))
}
