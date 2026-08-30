package main

import "snapshot/hxrt"

func main() {
	values := hxrt.NewArray(2, 4, 6)
	sum := 0
	_g := 0
	for _g < values.Len() {
		value := hxrt.IntFromNullableAny(values.Get(_g))
		_g = int(int32((_g + 1)))
		sum = int((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(value)))
	}
	hxrt.Println(any(sum))
}
