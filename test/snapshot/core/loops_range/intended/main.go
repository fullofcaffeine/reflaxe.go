package main

import "snapshot/hxrt"

func main() {
	sum := 0
	_ = sum
	sum = sum
	sum = int(int32((int32(sum) + int32(1))))
	sum = int(int32((int32(sum) + int32(2))))
	sum = int(int32((int32(sum) + int32(3))))
	hxrt.Println(sum)
}
