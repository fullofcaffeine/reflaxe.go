package main

import "snapshot/hxrt"

func main() {
	entries := New_haxe__ds__StringMap()
	entries.__hx_this.set(hxrt.StringFromLiteral("first"), func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["name"] = hxrt.StringFromLiteral("alpha")
		hx_obj_1["count"] = 3
		return hx_obj_1
	}())
	_g := hxrt.NewArray()
	entry := func(hx_erased_iterator_2 map[string]any) map[string]any {
		hx_erased_iterator_has_next_3 := hx_erased_iterator_2["hasNext"].(func() bool)
		hx_erased_iterator_next_4 := hx_erased_iterator_2["next"].(func() any)
		hx_applied_iterator_5 := map[string]any{}
		hx_applied_iterator_5["hasNext"] = func() bool {
			return hx_erased_iterator_has_next_3()
		}
		hx_applied_iterator_5["next"] = func() map[string]any {
			return func(hx_value_6 any) map[string]any {
				if hx_value_6 == nil {
					var hx_zero_7 map[string]any
					return hx_zero_7
				}
				return hx_value_6.(map[string]any)
			}(hx_erased_iterator_next_4())
		}
		return hx_applied_iterator_5
	}(entries.__hx_this.iterator())
	for func(hx_obj_8 map[string]any) func() bool {
		hx_field_9 := hx_obj_8["hasNext"]
		if hx_field_9 == nil {
			var hx_zero_10 func() bool
			return hx_zero_10
		}
		return hx_field_9.(func() bool)
	}(entry)() {
		entry_1 := func(hx_obj_11 map[string]any) func() map[string]any {
			hx_field_12 := hx_obj_11["next"]
			if hx_field_12 == nil {
				var hx_zero_13 func() map[string]any
				return hx_zero_13
			}
			return hx_field_12.(func() map[string]any)
		}(entry)()
		_g.Push(entry_1)
	}
	copied := _g
	entry_2 := func(hx_value_15 any) map[string]any {
		if hx_value_15 == nil {
			var hx_zero_16 map[string]any
			return hx_zero_16
		}
		return hx_value_15.(map[string]any)
	}(copied.Get(0))
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), func(hx_obj_17 map[string]any) *string {
		hx_field_18 := hx_obj_17["name"]
		if hx_field_18 == nil {
			var hx_zero_19 *string
			return hx_zero_19
		}
		return hx_field_18.(*string)
	}(entry_2)), hxrt.StringFromLiteral(":")), func(hx_obj_20 map[string]any) int {
		hx_field_21 := hx_obj_20["count"]
		if hx_field_21 == nil {
			var hx_zero_22 int
			return hx_zero_22
		}
		return hx_field_21.(int)
	}(entry_2)))
	hxrt.Println(v)
}
