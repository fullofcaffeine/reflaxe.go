package main

import "snapshot/hxrt"

func main() {
	a := -1
	_ = a
	b := int(int32(int32((uint32(hxrt.Int32Wrap(a)) >> uint(1)))))
	_ = b
	hxrt.Println(hxrt.StdString(b))
	hxrt.Println(hxrt.StringFromLiteral("2"))
}
