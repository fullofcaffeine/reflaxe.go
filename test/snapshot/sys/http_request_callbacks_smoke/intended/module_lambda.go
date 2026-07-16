package main

func Lambda_exists(it map[string]any, predicate func(any) bool) bool {
	found := false
	value := func(hx_obj_1 map[string]any) func() map[string]any {
		hx_field_2 := hx_obj_1["iterator"]
		if hx_field_2 == nil {
			var hx_zero_3 func() map[string]any
			return hx_zero_3
		}
		return hx_field_2.(func() map[string]any)
	}(it)()
	for func(hx_obj_4 map[string]any) func() bool {
		hx_field_5 := hx_obj_4["hasNext"]
		if hx_field_5 == nil {
			var hx_zero_6 func() bool
			return hx_zero_6
		}
		return hx_field_5.(func() bool)
	}(value)() {
		var value_1 any = func(hx_obj_7 map[string]any) func() any {
			hx_field_8 := hx_obj_7["next"]
			if hx_field_8 == nil {
				var hx_zero_9 func() any
				return hx_zero_9
			}
			return hx_field_8.(func() any)
		}(value)()
		if predicate(value_1) {
			found = true
			break
		}
	}
	return found
}
