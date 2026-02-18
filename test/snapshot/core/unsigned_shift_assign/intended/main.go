package main

import "snapshot/hxrt"

func main() {
	a := -1
	_ = a
	b := func() int {
		a = int((uint32(a) >> uint(1)))
		return a
	}()
	_ = b
	hxrt.Println(hxrt.StdString(a))
	hxrt.Println(hxrt.StdString(b))
}
