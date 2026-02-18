package main

import "snapshot/hxrt"

func main() {
	values := []int{2, 4, 6}
	_ = values
	sum := 0
	_ = sum
	hx_tmp := 0
	_ = hx_tmp
	for hx_tmp < len(values) {
		value := values[hx_tmp]
		_ = value
		hx_tmp = (hx_tmp + 1)
		sum = (sum + value)
	}
	hxrt.Println(sum)
}
