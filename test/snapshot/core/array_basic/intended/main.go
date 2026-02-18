package main

import "snapshot/hxrt"

func main() {
	values := []int{10, 20, 30}
	hxrt.Println(values[0])
	hxrt.Println(len(values))
}
