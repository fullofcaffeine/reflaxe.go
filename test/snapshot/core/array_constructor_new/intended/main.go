package main

import "snapshot/hxrt"

func main() {
	names := hxrt.NewArray()
	var v any = any(names.Len())
	hxrt.Println(v)
	names.Push(hxrt.StringFromLiteral("go"))
	names.Push(hxrt.StringFromLiteral("haxe"))
	var v_1 any = any(names.Len())
	hxrt.Println(v_1)
	hxrt.Println(any(names.Get(0)))
	hxrt.Println(any(names.Get(1)))
	nums := hxrt.NewArray()
	nums.Push(3)
	nums.Push(5)
	sum := 0
	_g := 0
	for _g < nums.Len() {
		n := hxrt.IntFromNullableAny(nums.Get(_g))
		_g = int(int32((_g + 1)))
		sum = int((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(n)))
	}
	hxrt.Println(any(sum))
}
