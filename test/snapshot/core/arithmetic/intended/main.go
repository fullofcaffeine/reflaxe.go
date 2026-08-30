package main

import "snapshot/hxrt"

func main() {
	a := 17
	b := 2.5
	input := 7
	negated := int(-int32(input))
	complemented := int(^int32(input))
	hxrt.Println(any(a))
	hxrt.Println(any(b))
	hxrt.Println(any(negated))
	hxrt.Println(any(complemented))
}
