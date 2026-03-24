package main

import "snapshot/hxrt"

func main() {
	names := []*string{}
	hxrt.Println(len(names))
	hx_arr_1 := names
	hx_arr_1 = append(hx_arr_1, hxrt.StringFromLiteral("go"))
	names = hx_arr_1
	hx_arr_2 := names
	hx_arr_2 = append(hx_arr_2, hxrt.StringFromLiteral("haxe"))
	names = hx_arr_2
	hxrt.Println(len(names))
	hxrt.Println(names[0])
	hxrt.Println(names[1])
	nums := []int{}
	hx_arr_3 := nums
	hx_arr_3 = append(hx_arr_3, 3)
	nums = hx_arr_3
	hx_arr_4 := nums
	hx_arr_4 = append(hx_arr_4, 5)
	nums = hx_arr_4
	sum := 0
	_g := 0
	for _g < len(nums) {
		n := nums[_g]
		_g = int(int32((_g + 1)))
		sum = int(int32((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(n))))
	}
	hxrt.Println(sum)
}
