package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(sum([]int{1, 2, 3}))
	hxrt.Println(sum([]int{4}))
}

func sum(values []int) int {
	total := 0
	i := 0
	for i < len(values) {
		total = (total + values[i])
		i = (i + 1)
	}
	return total
}
