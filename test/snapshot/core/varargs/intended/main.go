package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(sum([]int{1, 2, 3}))
	hxrt.Println(sum([]int{4}))
}

func sum(values []int) int {
	total := 0
	_ = total
	i := 0
	_ = i
	for i < len(values) {
		total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(values[i]))))
		i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))
	}
	return total
}
