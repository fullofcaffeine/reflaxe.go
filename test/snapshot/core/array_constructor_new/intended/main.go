package main

import "snapshot/hxrt"

func main() {
	names := []*string{}
	_ = names
	hxrt.Println(len(names))
	names = append(names, hxrt.StringFromLiteral("go"))
	names = append(names, hxrt.StringFromLiteral("haxe"))
	hxrt.Println(len(names))
	hxrt.Println(names[0])
	hxrt.Println(names[1])
	nums := []int{}
	_ = nums
	nums = append(nums, 3)
	nums = append(nums, 5)
	sum := 0
	_ = sum
	_g := 0
	_ = _g
	for _g < len(nums) {
		n := nums[_g]
		_ = n
		_g = (_g + 1)
		sum = (sum + n)
	}
	hxrt.Println(sum)
}
