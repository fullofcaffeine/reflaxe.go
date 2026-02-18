package main

import "snapshot/hxrt"

func main() {
	total := 1
	total = (total + 4)
	flag := false
	flag = (total > 3)
	hxrt.Println(total)
	hxrt.Println(flag)
}
