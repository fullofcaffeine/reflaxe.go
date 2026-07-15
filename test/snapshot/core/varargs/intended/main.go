package main

import "snapshot/hxrt"

func main() {
	var v any = any(sum([]int{1, 2, 3}))
	hxrt.Println(v)
	var v_1 any = any(sum([]int{4}))
	hxrt.Println(v_1)
}

func sum(values []int) int {
	total := 0
	i := 0
	for i < len(values) {
		total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(values[i]))))
		i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))
	}
	return total
}
