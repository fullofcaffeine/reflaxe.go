package main

import "snapshot/hxrt"

func main() {
	i := 1
	_ = i
	now := func() int {
		i = (i + 2)
		return i
	}()
	_ = now
	hxrt.Println(hxrt.StdString(now))
	hxrt.Println(hxrt.StdString(i))
	i = (i * 4)
	hxrt.Println(hxrt.StdString(i))
}
