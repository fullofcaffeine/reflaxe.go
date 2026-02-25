package main

import "snapshot/hxrt"

func main() {
	a := -1
	b := int(int32(int32((uint32(hxrt.Int32Wrap(a)) >> uint(1)))))
	hxrt.Println(hxrt.StdString(b))
	hxrt.Println(hxrt.StringFromLiteral("2"))
}
