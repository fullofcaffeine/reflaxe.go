package main

import "snapshot/hxrt"

func main() {
	a := -1
	_ = a
	b := func() int {
		a = int(int32(int32((uint32(int32(a)) >> uint(1)))))
		return a
	}()
	_ = b
	hxrt.Println(hxrt.StdString(a))
	hxrt.Println(hxrt.StdString(b))
}
