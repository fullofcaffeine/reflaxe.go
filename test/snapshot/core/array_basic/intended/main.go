package main

import "snapshot/hxrt"

func main() {
	values := []int{10, 20, 30}
	_ = values
	hxrt.Println(values[0])
	hxrt.Println(len(values))
}
