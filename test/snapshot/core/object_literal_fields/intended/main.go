package main

import "snapshot/hxrt"

func main() {
	user := makeUser(hxrt.StringFromLiteral("marcelo"), 10)
	_ = user
	hxrt.Println(user["name"].(*string))
	user["score"] = (user["score"].(int) + 5)
	hxrt.Println(user["score"].(int))
	var nested_inner_flag bool
	_ = nested_inner_flag
	var nested_inner_count int
	_ = nested_inner_count
	nested_inner_flag = true
	nested_inner_count = 2
	hxrt.Println(nested_inner_flag)
	nested_inner_count = (nested_inner_count + 3)
	hxrt.Println(nested_inner_count)
}

func makeUser(name *string, score int) map[string]any {
	hx_obj_1 := map[string]any{}
	hx_obj_1["name"] = name
	hx_obj_1["score"] = score
	return hx_obj_1
}
