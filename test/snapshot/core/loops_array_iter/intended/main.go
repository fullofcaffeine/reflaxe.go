package main

import "snapshot/hxrt"

func main() {
	values := []int{2, 4, 6}
	sum := 0
	hx_tmp := 0
	for hx_tmp < len(values) {
		value := values[hx_tmp]
		hx_tmp = (hx_tmp + 1)
		sum = (sum + value)
	}
	hxrt.Println(sum)
}
