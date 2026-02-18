package main

import "snapshot/hxrt"

func add(a int, b int) int {
	return (a + b)
}

func main() {
	local := func(v int) int {
		return (v + 1)
	}
	_ = local
	hxrt.Println(add(1, 2))
	hxrt.Println(add(5, 2))
	hxrt.Println(add(5, 6))
	hxrt.Println(local(10))
	hxrt.Println(local(20))
}
