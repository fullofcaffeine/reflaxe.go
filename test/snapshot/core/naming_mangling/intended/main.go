package main

import "snapshot/hxrt"

func Keywords_map_() int {
	return 30
}

func Keywords_range_() int {
	return 40
}

func main() {
	hxrt.Println(a_b__Util_value())
	hxrt.Println(a__b__Util_value())
	hxrt.Println(Keywords_map_())
	hxrt.Println(Keywords_range_())
}

func a__b__Util_value() int {
	return 2
}

func a_b__Util_value() int {
	return 1
}
