package main

import "snapshot/hxrt"

func add(a int, b int) int {
	return int(int32((int32(a) + int32(b))))
}

func main() {
	local := func(v int) int {
		return int(int32((int32(v) + int32(1))))
	}
	_ = local
	hxrt.Println(add(1, 2))
	hxrt.Println(add(5, 2))
	hxrt.Println(add(5, 6))
	hxrt.Println(local(10))
	hxrt.Println(local(20))
}
