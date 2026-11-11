package main

import "snapshot/hxrt"

func main() {
	var v any = any(selected())
	hxrt.Println(v)
	func() any {
		hx_structural_target_1 := record
		hx_structural_value_3 := false
		hxrt.AnonymousSet(hx_structural_target_1, "enabled", hx_structural_value_3)
		return hx_structural_value_3
	}()
	var v_1 any = any(selected())
	hxrt.Println(v_1)
	hx_obj_4 := map[string]any{}
	record = hxrt.AnonymousFromMap(hx_obj_4)
	var v_2 any = any(selected())
	hxrt.Println(v_2)
	visits := 0
	checks := 0
	hx_obj_5 := map[string]any{}
	hx_obj_5["enabled"] = true
	record = hxrt.AnonymousFromMap(hx_obj_5)
	for func() bool {
		checks = int(int32((checks + 1)))
		return func(hx_value_7 any) bool {
			if hx_value_7 == nil {
				var hx_zero_8 bool
				return hx_zero_8
			}
			return hx_value_7.(bool)
		}(func(hx_value_6 any) any {
			if hx_value_6 == nil {
				return nil
			}
			return hx_value_6.(bool)
		}(hxrt.AnonymousGet(record, "enabled")))
	}() {
		visits = int(int32((visits + 1)))
		func() any {
			hx_structural_target_9 := record
			hx_structural_value_11 := false
			hxrt.AnonymousSet(hx_structural_target_9, "enabled", hx_structural_value_11)
			return hx_structural_value_11
		}()
	}
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatAny(visits, hxrt.StringFromLiteral(":")), checks)))
	hx_obj_12 := map[string]any{}
	record = hxrt.AnonymousFromMap(hx_obj_12)
	hx_do_first_13 := true
	for hx_do_first_13 || func() bool {
		checks = int(int32((checks + 1)))
		return func(hx_value_15 any) bool {
			if hx_value_15 == nil {
				var hx_zero_16 bool
				return hx_zero_16
			}
			return hx_value_15.(bool)
		}(func(hx_value_14 any) any {
			if hx_value_14 == nil {
				return nil
			}
			return hx_value_14.(bool)
		}(hxrt.AnonymousGet(record, "enabled")))
	}() {
		hx_do_first_13 = false
		visits = int(int32((visits + 1)))
	}
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatAny(visits, hxrt.StringFromLiteral(":")), checks)))
	hx_obj_17 := map[string]any{}
	hx_obj_17["enabled"] = true
	record = hxrt.AnonymousFromMap(hx_obj_17)
	var hx_if_21 any
	checks = int(int32((checks + 1)))
	if func(hx_value_19 any) bool {
		if hx_value_19 == nil {
			var hx_zero_20 bool
			return hx_zero_20
		}
		return hx_value_19.(bool)
	}(func(hx_value_18 any) any {
		if hx_value_18 == nil {
			return nil
		}
		return hx_value_18.(bool)
	}(hxrt.AnonymousGet(record, "enabled"))) {
		hx_if_21 = hxrt.StringFromLiteral("yes")
	} else {
		hx_if_21 = hxrt.StringFromLiteral("no")
	}
	var v_3 any = hx_if_21
	hxrt.Println(v_3)
	hxrt.Println(any(checks))
}

var record hxrt.Anonymous = func() hxrt.Anonymous {
	hx_obj_22 := map[string]any{}
	hx_obj_22["enabled"] = true
	return hxrt.AnonymousFromMap(hx_obj_22)
}()

func selected() *string {
	if func(hx_value_24 any) bool {
		if hx_value_24 == nil {
			var hx_zero_25 bool
			return hx_zero_25
		}
		return hx_value_24.(bool)
	}(func(hx_value_23 any) any {
		if hx_value_23 == nil {
			return nil
		}
		return hx_value_23.(bool)
	}(hxrt.AnonymousGet(record, "enabled"))) {
		return hxrt.StringFromLiteral("on")
	}
	return hxrt.StringFromLiteral("off")
}
