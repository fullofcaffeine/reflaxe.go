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
		_g_1 := _g.params[0].(*hxrt.Array)
		_ = _g_1
	case 1:
		_g_2 := _g.params[0].(*string)
		message := _g_2
		hxrt.Println(any(message))
	}
	_g_3 := selectItem(ItemOutcome_ItemSuccess(hxrt.NewArray(func() map[string]any {
		hx_obj_2 := map[string]any{}
		hx_obj_2["id"] = hxrt.StringFromLiteral("one")
		return hx_obj_2
	}())))
	switch _g_3.tag {
	case 0:
		_g_4 := _g_3.params[0].(*hxrt.Array)
		items := _g_4
		var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("selected-item:"), func(hx_obj_5 map[string]any) *string {
			hx_field_6 := hx_obj_5["id"]
			if hx_field_6 == nil {
				var hx_zero_7 *string
				return hx_zero_7
			}
			return hx_field_6.(*string)
		}(func(hx_value_3 any) map[string]any {
			if hx_value_3 == nil {
				var hx_zero_4 map[string]any
				return hx_zero_4
			}
			return hx_value_3.(map[string]any)
		}(items.Get(0)))))
		hxrt.Println(v_4)
	case 1:
		_g_5 := _g_3.params[0].(*string)
		_ = _g_5
	}
}

func selectItem(outcome *ItemOutcome) *ItemOutcome {
	var hx_switch_8 map[string]any
	switch outcome.tag {
	case 0:
		_g := outcome.params[0].(*hxrt.Array)
		items := _g
		hx_switch_8 = func(hx_value_9 any) map[string]any {
			if hx_value_9 == nil {
				var hx_zero_10 map[string]any
				return hx_zero_10
			}
			return hx_value_9.(map[string]any)
		}(items.Get(0))
	case 1:
		_g_1 := outcome.params[0].(*string)
		message := _g_1
		return ItemOutcome_ItemFailure(message)
	}
	selected := hx_switch_8
	return ItemOutcome_ItemSuccess(hxrt.NewArray(selected))
}
