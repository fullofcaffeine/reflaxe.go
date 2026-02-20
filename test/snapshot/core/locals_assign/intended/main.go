package main

import "snapshot/hxrt"

func main() {
	total := 1
	_ = total
	total = int(int32((int32(total) + int32(4))))
	flag := false
	_ = flag
	flag = (total > 3)
	hxrt.Println(total)
	hxrt.Println(flag)
}
