package main

import "snapshot/hxrt"

func main() {
	values := []int{}
	values = append(values, 4)
	values = append(values, 9)
	if len(values) > 0 {
		values = values[:(len(values) - 1)]
	}
	pushLen := func() int {
		values = append(values, 12)
		return len(values)
	}()
	removed := func() int {
		hx_len_6 := len(values)
		if hx_len_6 == 0 {
			var hx_zero_8 int
			return hx_zero_8
		}
		hx_value_7 := values[(hx_len_6 - 1)]
		values = values[:(hx_len_6 - 1)]
		return hx_value_7
	}()
	hxrt.Println(len(values))
	hxrt.Println(values[0])
	hxrt.Println(pushLen)
	hxrt.Println(removed)
}
