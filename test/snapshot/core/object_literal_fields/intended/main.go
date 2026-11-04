package main

import "snapshot/hxrt"

func main() {
	user := makeUser(hxrt.StringFromLiteral("marcelo"), 10)
	var v any = any(func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(hxrt.AnonymousGet(user, "name")))
	hxrt.Println(v)
	func() int {
		hx_structural_target_3 := user
		hx_structural_value_5 := int((hxrt.Int32Wrap(func(hx_value_6 any) int {
			if hx_value_6 == nil {
				var hx_zero_7 int
				return hx_zero_7
			}
			return hx_value_6.(int)
		}(hxrt.AnonymousGet(user, "score"))) + hxrt.Int32Wrap(5)))
		hxrt.AnonymousSet(hx_structural_target_3, "score", hx_structural_value_5)
		return hx_structural_value_5
	}()
	var v_1 any = any(func(hx_value_8 any) int {
		if hx_value_8 == nil {
			var hx_zero_9 int
			return hx_zero_9
		}
		return hx_value_8.(int)
	}(hxrt.AnonymousGet(user, "score")))
	hxrt.Println(v_1)
	var nested_inner_flag bool
	var nested_inner_count int
	nested_inner_flag = true
	nested_inner_count = 2
	var v_2 any = any(nested_inner_flag)
	hxrt.Println(v_2)
	nested_inner_count = int((hxrt.Int32Wrap(nested_inner_count) + hxrt.Int32Wrap(3)))
	var v_3 any = any(nested_inner_count)
	hxrt.Println(v_3)
}

func makeUser(name *string, score int) hxrt.Anonymous {
	hx_obj_10 := map[string]any{}
	hx_obj_10["name"] = name
	hx_obj_10["score"] = score
	return hxrt.AnonymousFromMap(hx_obj_10)
}
