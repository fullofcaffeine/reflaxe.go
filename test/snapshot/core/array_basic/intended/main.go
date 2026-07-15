package main

import "snapshot/hxrt"

func main() {
	var values_2 int
	_ = values_2
	var values_1 int
	_ = values_1
	var values_0 int
	values_0 = 10
	values_1 = 20
	values_2 = 30
	hxrt.Println(any(values_0))
	var v any = any(3)
	hxrt.Println(v)
}
