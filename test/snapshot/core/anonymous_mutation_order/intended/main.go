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
			hx_structural_old_5 := func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(hxrt.AnonymousGet(hx_structural_target_4, "count"))
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
			hx_structural_target_9 := fh_1
			hx_structural_old_10 := func(hx_value_12 any) int {
				if hx_value_12 == nil {
					var hx_zero_13 int
					return hx_zero_13
				}
				return hx_value_12.(int)
			}(hxrt.AnonymousGet(hx_structural_target_9, "count"))
			hx_structural_value_11 := int(int32((hx_structural_old_10 + 1)))
			hxrt.AnonymousSet(hx_structural_target_9, "count", hx_structural_value_11)
			return hx_structural_old_10
		}()
	}())
	hxrt.Println(v_4)
	var v_5 any = any(order)
	hxrt.Println(v_5)
	order = hxrt.StringFromLiteral("")
	var v_6 any = any(func() int {
		fh_2 := target()
		return func() int {
			hx_structural_target_14 := fh_2
			hx_structural_old_15 := func(hx_value_17 any) int {
				if hx_value_17 == nil {
					var hx_zero_18 int
					return hx_zero_18
				}
				return hx_value_17.(int)
			}(hxrt.AnonymousGet(hx_structural_target_14, "count"))
			hx_structural_value_16 := int(int32((hx_structural_old_15 + 1)))
			hxrt.AnonymousSet(hx_structural_target_14, "count", hx_structural_value_16)
			return hx_structural_value_16
		}()
	}())
	hxrt.Println(v_6)
	var v_7 any = any(order)
	hxrt.Println(v_7)
	fh_3 := target()
	func() *string {
		hx_structural_target_19 := fh_3
		hx_structural_old_20 := func(hx_value_22 any) *string {
			if hx_value_22 == nil {
				var hx_zero_23 *string
				return hx_zero_23
			}
			return hx_value_22.(*string)
		}(hxrt.AnonymousGet(hx_structural_target_19, "text"))
		hx_structural_value_21 := hxrt.StringConcatStringPtr(hx_structural_old_20, hxrt.StringFromLiteral("b"))
		hxrt.AnonymousSet(hx_structural_target_19, "text", hx_structural_value_21)
		return hx_structural_value_21
	}()
	hx_arr_24 := func(hx_value_25 any) *hxrt.Array {
		if hx_value_25 == nil {
			var hx_zero_26 *hxrt.Array
			return hx_zero_26
		}
		return hx_value_25.(*hxrt.Array)
	}(hxrt.AnonymousGet(target(), "values"))
	hx_arr_24.Push(9)
	var v_8 any = any(func(hx_value_27 any) int {
		if hx_value_27 == nil {
			var hx_zero_28 int
			return hx_zero_28
		}
		return hx_value_27.(int)
	}(hxrt.AnonymousGet(record, "count")))
	hxrt.Println(v_8)
	var v_9 any = any(func(hx_value_29 any) *string {
		if hx_value_29 == nil {
			var hx_zero_30 *string
			return hx_zero_30
		}
		return hx_value_29.(*string)
	}(hxrt.AnonymousGet(record, "text")))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringJoinAny(func(hx_value_31 any) *hxrt.Array {
		if hx_value_31 == nil {
			var hx_zero_32 *hxrt.Array
			return hx_zero_32
		}
		return hx_value_31.(*hxrt.Array)
	}(hxrt.AnonymousGet(record, "values")).Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_10)
}

var order *string = hxrt.StringFromLiteral("")

var record hxrt.Anonymous = func() hxrt.Anonymous {
	hx_obj_33 := map[string]any{}
	hx_obj_33["count"] = 2
	hx_obj_33["text"] = hxrt.StringFromLiteral("a")
	hx_obj_33["values"] = hxrt.NewArray(1)
	return hxrt.AnonymousFromMap(hx_obj_33)
}()

func right() int {
	order = hxrt.StringConcatStringPtr(order, hxrt.StringFromLiteral("R"))
	return 3
}

func target() hxrt.Anonymous {
	order = hxrt.StringConcatStringPtr(order, hxrt.StringFromLiteral("L"))
	return record
}
