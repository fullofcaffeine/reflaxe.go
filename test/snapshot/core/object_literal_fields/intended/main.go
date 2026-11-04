package main

import "snapshot/hxrt"

func main() {
	user := makeUser(hxrt.StringFromLiteral("marcelo"), 10)
	var v any = any(func(hx_obj_1 hxrt.Anonymous) *string {
		hx_field_2 := hxrt.AnonymousGet(hx_obj_1, "name")
		if hx_field_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_field_2.(*string)
	}(user))
	hxrt.Println(v)
	func() int {
		hx_structural_target_4 := user
		hx_structural_value_6 := int((hxrt.Int32Wrap(func(hx_obj_7 hxrt.Anonymous) int {
			hx_field_8 := hxrt.AnonymousGet(hx_obj_7, "score")
			if hx_field_8 == nil {
				var hx_zero_9 int
				return hx_zero_9
			}
			return hx_field_8.(int)
		}(user)) + hxrt.Int32Wrap(5)))
		hxrt.AnonymousSet(hx_structural_target_4, "score", hx_structural_value_6)
		return hx_structural_value_6
	}()
	var v_1 any = any(func(hx_obj_10 hxrt.Anonymous) int {
		hx_field_11 := hxrt.AnonymousGet(hx_obj_10, "score")
		if hx_field_11 == nil {
			var hx_zero_12 int
			return hx_zero_12
		}
		return hx_field_11.(int)
	}(user))
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
	hx_obj_13 := map[string]any{}
	hx_obj_13["name"] = name
	hx_obj_13["score"] = score
	return hxrt.AnonymousFromMap(hx_obj_13)
}
