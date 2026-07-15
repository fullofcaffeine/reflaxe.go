package main

import "snapshot/hxrt"

func main() {
	i := 1
	now := func() int {
		i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2))))
		return i
	}()
	var v any = any(hxrt.StdString(now))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(i))
	hxrt.Println(v_1)
	i = int(int32((hxrt.Int32Wrap(i) * hxrt.Int32Wrap(4))))
	var v_2 any = any(hxrt.StdString(i))
	hxrt.Println(v_2)
}
