package main

import "snapshot/hxrt"

func main() {
	i := 1
	_ = i
	now := func() int {
		i = int(int32((int32(i) + int32(2))))
		return i
	}()
	_ = now
	hxrt.Println(hxrt.StdString(now))
	hxrt.Println(hxrt.StdString(i))
	i = int(int32((int32(i) * int32(4))))
	hxrt.Println(hxrt.StdString(i))
}
