package main

func Lambda_exists(it map[string]any, f func(any) bool) bool {
	x := func(hx_obj_1 map[string]any) func() map[string]any {
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
	}(x)() {
		var x_1 any = func(hx_obj_7 map[string]any) func() any {
			hx_field_8 := hx_obj_7["next"]
			if hx_field_8 == nil {
				var hx_zero_9 func() any
				return hx_zero_9
			}
			return hx_field_8.(func() any)
		}(x)()
		if f(x_1) {
			return true
		}
	}
	return false
}
