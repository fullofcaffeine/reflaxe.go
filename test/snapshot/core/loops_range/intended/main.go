package main

import "snapshot/hxrt"

func main() {
	sum := 0
	_ = sum
	sum = sum
	sum = (sum + 1)
	sum = (sum + 2)
	sum = (sum + 3)
	hxrt.Println(sum)
}
