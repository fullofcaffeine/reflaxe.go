package main

import "snapshot/hxrt"

func Keywords_map_() int {
	return 30
}

func Keywords_range_() int {
	return 40
}

func main() {
	var v any = any(a_b__Util_value())
	hxrt.Println(v)
	var v_1 any = any(a__b__Util_value())
	hxrt.Println(v_1)
	var v_2 any = any(Keywords_map_())
	hxrt.Println(v_2)
	var v_3 any = any(Keywords_range_())
	hxrt.Println(v_3)
}
