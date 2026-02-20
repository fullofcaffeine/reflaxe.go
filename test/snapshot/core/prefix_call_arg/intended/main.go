package main

import "snapshot/hxrt"

func id(value int) int {
	return value
}

func main() {
	i := 1
	_ = i
	now := id(func() int {
		i = int(int32((i + 1)))
		return i
	}())
	_ = now
	hxrt.Println(hxrt.StdString(now))
	hxrt.Println(hxrt.StdString(i))
}
