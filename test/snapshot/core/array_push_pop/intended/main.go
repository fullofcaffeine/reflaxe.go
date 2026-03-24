package main

import "snapshot/hxrt"

func main() {
	values := []int{}
	hx_arr_1 := values
	hx_arr_1 = append(hx_arr_1, 4)
	values = hx_arr_1
	hx_arr_2 := values
	hx_arr_2 = append(hx_arr_2, 9)
	values = hx_arr_2
	hx_arr_3 := values
	if len(hx_arr_3) > 0 {
		hx_arr_3 = hx_arr_3[:(len(hx_arr_3) - 1)]
	}
	values = hx_arr_3
	hxrt.Println(len(values))
	hxrt.Println(values[0])
}
