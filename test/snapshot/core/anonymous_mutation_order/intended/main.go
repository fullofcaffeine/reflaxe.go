package main

import "snapshot/hxrt"

func main() {
	var v any = any(func() int {
		hx_structural_target_1 := target()
		hx_structural_value_3 := right()
		hxrt.AnonymousSet(hx_structural_target_1, "count", hx_structural_value_3)
		return hx_structural_value_3
	}())
	hxrt.Println(v)
	var v_1 any = any(order)
	hxrt.Println(v_1)
	order = hxrt.StringFromLiteral("")
	var v_2 any = any(func() int {
		fh := target()
		return func() int {
			hx_structural_target_4 := fh
			hx_structural_old_5 := func(hx_obj_7 hxrt.Anonymous) int {
				hx_field_8 := hxrt.AnonymousGet(hx_obj_7, "count")
				if hx_field_8 == nil {
					var hx_zero_9 int
					return hx_zero_9
				}
				return hx_field_8.(int)
			}(hx_structural_target_4)
			hx_structural_value_6 := int((hxrt.Int32Wrap(hx_structural_old_5) + hxrt.Int32Wrap(right())))
			hxrt.AnonymousSet(hx_structural_target_4, "count", hx_structural_value_6)
			return hx_structural_value_6
		}()
	}())
	hxrt.Println(v_2)
	var v_3 any = any(order)
	hxrt.Println(v_3)
	order = hxrt.StringFromLiteral("")
	var v_4 any = any(func() int {
		fh_1 := target()
		return func() int {
			hx_structural_target_10 := fh_1
			hx_structural_old_11 := func(hx_obj_13 hxrt.Anonymous) int {
				hx_field_14 := hxrt.AnonymousGet(hx_obj_13, "count")
				if hx_field_14 == nil {
					var hx_zero_15 int
					return hx_zero_15
				}
				return hx_field_14.(int)
			}(hx_structural_target_10)
			hx_structural_value_12 := int(int32((hx_structural_old_11 + 1)))
			hxrt.AnonymousSet(hx_structural_target_10, "count", hx_structural_value_12)
			return hx_structural_old_11
		}()
	}())
	hxrt.Println(v_4)
	var v_5 any = any(order)
	hxrt.Println(v_5)
	order = hxrt.StringFromLiteral("")
	var v_6 any = any(func() int {
		fh_2 := target()
		return func() int {
			hx_structural_target_16 := fh_2
			hx_structural_old_17 := func(hx_obj_19 hxrt.Anonymous) int {
				hx_field_20 := hxrt.AnonymousGet(hx_obj_19, "count")
				if hx_field_20 == nil {
					var hx_zero_21 int
					return hx_zero_21
				}
				return hx_field_20.(int)
			}(hx_structural_target_16)
			hx_structural_value_18 := int(int32((hx_structural_old_17 + 1)))
			hxrt.AnonymousSet(hx_structural_target_16, "count", hx_structural_value_18)
			return hx_structural_value_18
		}()
	}())
	hxrt.Println(v_6)
	var v_7 any = any(order)
	hxrt.Println(v_7)
	fh_3 := target()
	func() *string {
		hx_structural_target_22 := fh_3
		hx_structural_old_23 := func(hx_obj_25 hxrt.Anonymous) *string {
			hx_field_26 := hxrt.AnonymousGet(hx_obj_25, "text")
			if hx_field_26 == nil {
				var hx_zero_27 *string
				return hx_zero_27
			}
			return hx_field_26.(*string)
		}(hx_structural_target_22)
		hx_structural_value_24 := hxrt.StringConcatStringPtr(hx_structural_old_23, hxrt.StringFromLiteral("b"))
		hxrt.AnonymousSet(hx_structural_target_22, "text", hx_structural_value_24)
		return hx_structural_value_24
	}()
	hx_arr_28 := func(hx_obj_29 hxrt.Anonymous) *hxrt.Array {
		hx_field_30 := hxrt.AnonymousGet(hx_obj_29, "values")
		if hx_field_30 == nil {
			var hx_zero_31 *hxrt.Array
			return hx_zero_31
		}
		return hx_field_30.(*hxrt.Array)
	}(target())
	hx_arr_28.Push(9)
	var v_8 any = any(func(hx_obj_32 hxrt.Anonymous) int {
		hx_field_33 := hxrt.AnonymousGet(hx_obj_32, "count")
		if hx_field_33 == nil {
			var hx_zero_34 int
			return hx_zero_34
		}
		return hx_field_33.(int)
	}(record))
	hxrt.Println(v_8)
	var v_9 any = any(func(hx_obj_35 hxrt.Anonymous) *string {
		hx_field_36 := hxrt.AnonymousGet(hx_obj_35, "text")
		if hx_field_36 == nil {
			var hx_zero_37 *string
			return hx_zero_37
		}
		return hx_field_36.(*string)
	}(record))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringJoinAny(func(hx_obj_38 hxrt.Anonymous) *hxrt.Array {
		hx_field_39 := hxrt.AnonymousGet(hx_obj_38, "values")
		if hx_field_39 == nil {
			var hx_zero_40 *hxrt.Array
			return hx_zero_40
		}
		return hx_field_39.(*hxrt.Array)
	}(record).Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_10)
}

var order *string = hxrt.StringFromLiteral("")

var record hxrt.Anonymous = func() hxrt.Anonymous {
	hx_obj_41 := map[string]any{}
	hx_obj_41["count"] = 2
	hx_obj_41["text"] = hxrt.StringFromLiteral("a")
	hx_obj_41["values"] = hxrt.NewArray(1)
	return hxrt.AnonymousFromMap(hx_obj_41)
}()

func right() int {
	order = hxrt.StringConcatStringPtr(order, hxrt.StringFromLiteral("R"))
	return 3
}

func target() hxrt.Anonymous {
	order = hxrt.StringConcatStringPtr(order, hxrt.StringFromLiteral("L"))
	return record
}
