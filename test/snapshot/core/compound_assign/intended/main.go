package main

import "snapshot/hxrt"

func main() {
	i := 1
	now := func() int {
		i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2))))
		return i
	}()
	hxrt.Println(hxrt.StdString(now))
	hxrt.Println(hxrt.StdString(i))
	i = int(int32((hxrt.Int32Wrap(i) * hxrt.Int32Wrap(4))))
	hxrt.Println(hxrt.StdString(i))
}
