package main

import "snapshot/hxrt"

func main() {
	sum := 0
	_ = sum
	sum = sum
	sum = int(int32((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(1))))
	sum = int(int32((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(2))))
	sum = int(int32((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(3))))
	hxrt.Println(sum)
}
