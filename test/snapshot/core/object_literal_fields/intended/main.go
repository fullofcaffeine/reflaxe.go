package main

import "snapshot/hxrt"

func main() {
	user := makeUser(hxrt.StringFromLiteral("marcelo"), 10)
	var v any = any(func(hx_obj_1 map[string]any) *string {
		hx_field_2 := hx_obj_1["name"]
		if hx_field_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_field_2.(*string)
	}(user))
	hxrt.Println(v)
	user["score"] = int(int32((hxrt.Int32Wrap(func(hx_obj_4 map[string]any) int {
		hx_field_5 := hx_obj_4["score"]
		if hx_field_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_field_5.(int)
	}(user)) + hxrt.Int32Wrap(5))))
	var v_1 any = any(func(hx_obj_7 map[string]any) int {
		hx_field_8 := hx_obj_7["score"]
		if hx_field_8 == nil {
			var hx_zero_9 int
			return hx_zero_9
		}
		return hx_field_8.(int)
	}(user))
	hxrt.Println(v_1)
	var nested_inner_flag bool
	var nested_inner_count int
	nested_inner_flag = true
	nested_inner_count = 2
	var v_2 any = any(nested_inner_flag)
	hxrt.Println(v_2)
	nested_inner_count = int(int32((hxrt.Int32Wrap(nested_inner_count) + hxrt.Int32Wrap(3))))
	var v_3 any = any(nested_inner_count)
	hxrt.Println(v_3)
}

func makeUser(name *string, score int) map[string]any {
	hx_obj_10 := map[string]any{}
	hx_obj_10["name"] = name
	hx_obj_10["score"] = score
	return hx_obj_10
}
