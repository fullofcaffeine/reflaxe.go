package main

import "snapshot/hxrt"

func main() {
	total := 1
	total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(4)))
	flag := false
	flag = (total > 3)
	hxrt.Println(any(total))
	hxrt.Println(any(flag))
}
