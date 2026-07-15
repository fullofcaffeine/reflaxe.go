package main

import "snapshot/hxrt"

func main() {
	names := []*string{}
	var v any = any(len(names))
	hxrt.Println(v)
	names = append(names, hxrt.StringFromLiteral("go"))
	names = append(names, hxrt.StringFromLiteral("haxe"))
	var v_1 any = any(len(names))
	hxrt.Println(v_1)
	hxrt.Println(any(names[0]))
	hxrt.Println(any(names[1]))
	nums := []int{}
	nums = append(nums, 3)
	nums = append(nums, 5)
	sum := 0
	_g := 0
	for _g < len(nums) {
		n := nums[_g]
		_g = int(int32((_g + 1)))
		sum = int(int32((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(n))))
	}
	hxrt.Println(any(sum))
}
