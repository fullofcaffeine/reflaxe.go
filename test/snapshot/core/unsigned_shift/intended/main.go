package main

import "snapshot/hxrt"

func main() {
	a := -1
	b := int(int32(int32((uint32(hxrt.Int32Wrap(a)) >> uint(1)))))
	var v any = any(hxrt.StdString(b))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.StringFromLiteral("2")))
}
