package main

import "snapshot/hxrt"

func main() {
	i := 1
	_ = i
	now := func() int {
		i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2))))
		return i
	}()
	_ = now
	hxrt.Println(hxrt.StdString(now))
	hxrt.Println(hxrt.StdString(i))
	i = int(int32((hxrt.Int32Wrap(i) * hxrt.Int32Wrap(4))))
	hxrt.Println(hxrt.StdString(i))
}
