package main

import (
	"fmt"
	"snapshot/hxrt"
)

func main() {
	value := hxrt.StringFromLiteral("\x00A\x01\a\b\f\n\r\t\v\x1f\x7f\"\\é🙂")
	bytes := haxe__io__Bytes_ofString(value, nil)
	fmt.Println(bytes.length)
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_1
		fmt.Println(bytes.b[index])
	}
}
