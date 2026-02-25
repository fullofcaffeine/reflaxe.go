package main

import "snapshot/hxrt"

func main() {
	a := -1
	b := func() int {
		a = int(int32(int32((uint32(hxrt.Int32Wrap(a)) >> uint(1)))))
		return a
	}()
	hxrt.Println(hxrt.StdString(a))
	hxrt.Println(hxrt.StdString(b))
}
