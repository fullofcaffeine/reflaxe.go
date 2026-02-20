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
		total = int(int32((int32(total) + int32(values[i]))))
		i = int(int32((int32(i) + int32(1))))
	}
	return total
}
