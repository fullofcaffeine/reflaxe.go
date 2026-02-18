package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(sum([]int{1, 2, 3}))
	hxrt.Println(sum([]int{4}))
}

func sum(values []int) int {
	total := 0
	_ = total
	hx_tmp_args := values
	hx_tmp_current := 0
	for hx_tmp_current < len(hx_tmp_args) {
		self := hx_tmp_args
		_ = self
		hx_post_1 := hx_tmp_current
		hx_tmp_current = (hx_tmp_current + 1)
		index := hx_post_1
		_ = index
		value := self[index]
		_ = value
		total = (total + value)
	}
	return total
}
