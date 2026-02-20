package main

import "snapshot/hxrt"

func main() {
	values := []int{2, 4, 6}
	_ = values
	sum := 0
	_ = sum
	_g := 0
	_ = _g
	for _g < len(values) {
		value := values[_g]
		_ = value
		_g = int(int32((_g + 1)))
		sum = int(int32((int32(sum) + int32(value))))
	}
	hxrt.Println(sum)
}
