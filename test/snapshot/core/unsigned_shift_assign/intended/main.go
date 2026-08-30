package main

import "snapshot/hxrt"

func main() {
	a := -1
	b := func() int {
		a = int(int32((uint32(hxrt.Int32Wrap(a)) >> uint(1))))
		return a
	}()
	var v any = any(hxrt.StdString(a))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(b))
	hxrt.Println(v_1)
}
