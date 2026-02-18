package main

import "snapshot/hxrt"

func main() {
	values := []int{}
	_ = values
	values = append(values, 4)
	values = append(values, 9)
	if len(values) > 0 {
		values = values[:(len(values) - 1)]
	}
	hxrt.Println(len(values))
	hxrt.Println(values[0])
}
