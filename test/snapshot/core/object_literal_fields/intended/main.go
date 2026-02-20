package main

import "snapshot/hxrt"

func main() {
	user := makeUser(hxrt.StringFromLiteral("marcelo"), 10)
	_ = user
	hxrt.Println(func(hx_obj_1 map[string]any) *string {
		hx_field_2 := hx_obj_1["name"]
		if hx_field_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_field_2.(*string)
	}(user))
	user["score"] = (func(hx_obj_4 map[string]any) int {
		hx_field_5 := hx_obj_4["score"]
		if hx_field_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_field_5.(int)
	}(user) + 5)
	hxrt.Println(func(hx_obj_7 map[string]any) int {
		hx_field_8 := hx_obj_7["score"]
		if hx_field_8 == nil {
			var hx_zero_9 int
			return hx_zero_9
		}
		return hx_field_8.(int)
	}(user))
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
	hx_obj_10 := map[string]any{}
	hx_obj_10["name"] = name
	hx_obj_10["score"] = score
	return hx_obj_10
}
