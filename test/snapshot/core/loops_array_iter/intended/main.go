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
		_g = (_g + 1)
		sum = (sum + value)
	}
	hxrt.Println(sum)
}
