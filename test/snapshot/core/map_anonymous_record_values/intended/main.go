package main

import "snapshot/hxrt"

func main() {
	entries := New_haxe__ds__StringMap()
	entries.__hx_this.set(hxrt.StringFromLiteral("first"), func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["name"] = hxrt.StringFromLiteral("alpha")
		hx_obj_1["count"] = 3
		return hxrt.AnonymousFromMap(hx_obj_1)
	}())
	_g := hxrt.NewArray()
	entry := func(hx_erased_iterator_2 hxrt.Anonymous) hxrt.Anonymous {
		hx_erased_iterator_has_next_3 := hxrt.AnonymousGet(hx_erased_iterator_2, "hasNext").(*hxrt.Function[func() bool])
		hx_erased_iterator_next_4 := hxrt.AnonymousGet(hx_erased_iterator_2, "next").(*hxrt.Function[func() any])
		hx_applied_iterator_5 := map[string]any{}
		hx_applied_iterator_5["hasNext"] = hx_erased_iterator_has_next_3
		hx_applied_iterator_5["next"] = hxrt.AdaptFunction(hx_erased_iterator_next_4, func() hxrt.Anonymous {
			return hxrt.AnonymousFromValue(hxrt.FunctionCall(hx_erased_iterator_next_4)())
		})
		return hxrt.AnonymousFromMap(hx_applied_iterator_5)
	}(entries.__hx_this.iterator())
	for func() bool {
		var hx_callable_9 *hxrt.Function[func() bool] = func(hx_obj_6 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_7 := hxrt.AnonymousGet(hx_obj_6, "hasNext")
			if hx_field_7 == nil {
				var hx_zero_8 *hxrt.Function[func() bool]
				return hx_zero_8
			}
			return hx_field_7.(*hxrt.Function[func() bool])
		}(entry)
		return hxrt.FunctionCall(hx_callable_9)()
	}() {
		entry_1 := func() hxrt.Anonymous {
			var hx_callable_13 *hxrt.Function[func() hxrt.Anonymous] = func(hx_obj_10 hxrt.Anonymous) *hxrt.Function[func() hxrt.Anonymous] {
				hx_field_11 := hxrt.AnonymousGet(hx_obj_10, "next")
				if hx_field_11 == nil {
					var hx_zero_12 *hxrt.Function[func() hxrt.Anonymous]
					return hx_zero_12
				}
				return hx_field_11.(*hxrt.Function[func() hxrt.Anonymous])
			}(entry)
			return hxrt.FunctionCall(hx_callable_13)()
		}()
		_g.Push(entry_1)
	}
	copied := _g
	entry_2 := hxrt.AnonymousFromValue(copied.Get(0))
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), func(hx_obj_15 hxrt.Anonymous) *string {
		hx_field_16 := hxrt.AnonymousGet(hx_obj_15, "name")
		if hx_field_16 == nil {
			var hx_zero_17 *string
			return hx_zero_17
		}
		return hx_field_16.(*string)
	}(entry_2)), hxrt.StringFromLiteral(":")), func(hx_obj_18 hxrt.Anonymous) int {
		hx_field_19 := hxrt.AnonymousGet(hx_obj_18, "count")
		if hx_field_19 == nil {
			var hx_zero_20 int
			return hx_zero_20
		}
		return hx_field_19.(int)
	}(entry_2)))
	hxrt.Println(v)
}
