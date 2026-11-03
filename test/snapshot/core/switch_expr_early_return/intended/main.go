package main

import "snapshot/hxrt"

type ItemOutcome struct {
	tag    int
	params []any
}

func ItemOutcome_ItemSuccess(items *hxrt.Array) *ItemOutcome {
	enumValue := &ItemOutcome{tag: 0}
	enumValue.params = []any{items}
	return enumValue
}

func ItemOutcome_ItemFailure(message *string) *ItemOutcome {
	enumValue := &ItemOutcome{tag: 1}
	enumValue.params = []any{message}
	return enumValue
}

func describe(value int) *string {
	var hx_switch_1 *string
	switch value {
	case 0:
		return hxrt.StringFromLiteral("zero")
	case 1:
		hx_switch_1 = hxrt.StringFromLiteral("one")
	default:
		if value < 0 {
			return hxrt.StringFromLiteral("negative")
		}
		hx_switch_1 = hxrt.StringFromLiteral("many")
	}
	selected := hx_switch_1
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("selected:"), selected)
}

func main() {
	var v any = any(describe(0))
	hxrt.Println(v)
	var v_1 any = any(describe(1))
	hxrt.Println(v_1)
	var v_2 any = any(describe(-1))
	hxrt.Println(v_2)
	var v_3 any = any(describe(2))
	hxrt.Println(v_3)
	_g := selectItem(ItemOutcome_ItemFailure(hxrt.StringFromLiteral("zero-item")))
	switch _g.tag {
	case 0:
		_g_1 := func(hx_value_2 any) *hxrt.Array {
			if hx_value_2 == nil {
				var hx_zero_3 *hxrt.Array
				return hx_zero_3
			}
			return hx_value_2.(*hxrt.Array)
		}(_g.params[0])
		_ = _g_1
	case 1:
		_g_2 := func(hx_value_4 any) *string {
			if hx_value_4 == nil {
				var hx_zero_5 *string
				return hx_zero_5
			}
			return hx_value_4.(*string)
		}(_g.params[0])
		message := _g_2
		hxrt.Println(any(message))
	}
	_g_3 := selectItem(ItemOutcome_ItemSuccess(hxrt.NewArray(func() map[string]any {
		hx_obj_6 := map[string]any{}
		hx_obj_6["id"] = hxrt.StringFromLiteral("one")
		return hx_obj_6
	}())))
	switch _g_3.tag {
	case 0:
		_g_4 := func(hx_value_7 any) *hxrt.Array {
			if hx_value_7 == nil {
				var hx_zero_8 *hxrt.Array
				return hx_zero_8
			}
			return hx_value_7.(*hxrt.Array)
		}(_g_3.params[0])
		items := _g_4
		var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("selected-item:"), func(hx_obj_11 map[string]any) *string {
			hx_field_12 := hx_obj_11["id"]
			if hx_field_12 == nil {
				var hx_zero_13 *string
				return hx_zero_13
			}
			return hx_field_12.(*string)
		}(func(hx_value_9 any) map[string]any {
			if hx_value_9 == nil {
				var hx_zero_10 map[string]any
				return hx_zero_10
			}
			return hx_value_9.(map[string]any)
		}(items.Get(0)))))
		hxrt.Println(v_4)
	case 1:
		_g_5 := func(hx_value_14 any) *string {
			if hx_value_14 == nil {
				var hx_zero_15 *string
				return hx_zero_15
			}
			return hx_value_14.(*string)
		}(_g_3.params[0])
		_ = _g_5
	}
}

func selectItem(outcome *ItemOutcome) *ItemOutcome {
	var hx_switch_16 map[string]any
	switch outcome.tag {
	case 0:
		_g := func(hx_value_17 any) *hxrt.Array {
			if hx_value_17 == nil {
				var hx_zero_18 *hxrt.Array
				return hx_zero_18
			}
			return hx_value_17.(*hxrt.Array)
		}(outcome.params[0])
		items := _g
		hx_switch_16 = func(hx_value_19 any) map[string]any {
			if hx_value_19 == nil {
				var hx_zero_20 map[string]any
				return hx_zero_20
			}
			return hx_value_19.(map[string]any)
		}(items.Get(0))
	case 1:
		_g_1 := func(hx_value_21 any) *string {
			if hx_value_21 == nil {
				var hx_zero_22 *string
				return hx_zero_22
			}
			return hx_value_21.(*string)
		}(outcome.params[0])
		message := _g_1
		return ItemOutcome_ItemFailure(message)
	}
	selected := hx_switch_16
	return ItemOutcome_ItemSuccess(hxrt.NewArray(selected))
}
