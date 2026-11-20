package main

import "snapshot/hxrt"

func Edited_run(first any, second any, callback any) int {
	square := hx_recover_class_Square(first)
	circle := hx_recover_class_Circle(second)
	total := int((hxrt.Int32Wrap(square.size) + hxrt.Int32Wrap(circle.size)))
	return total
}
