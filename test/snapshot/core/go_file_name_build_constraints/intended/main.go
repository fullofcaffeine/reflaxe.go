package main

import "snapshot/hxrt"

func main() {
	var v any = any(unit__Test_name())
	hxrt.Println(v)
	var v_1 any = any(tools__Linux_name())
	hxrt.Println(v_1)
	var v_2 any = any(tools__Windows_amd64_name())
	hxrt.Println(v_2)
	var v_3 any = any(x__Arm64_name())
	hxrt.Println(v_3)
	var v_4 any = any(x__Js_name())
	hxrt.Println(v_4)
	var v_5 any = any(x__Unix_name())
	hxrt.Println(v_5)
}
