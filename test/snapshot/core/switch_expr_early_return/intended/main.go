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
	_g_3 := selectItem(ItemOutcome_ItemSuccess(hxrt.NewArray(func() hxrt.Anonymous {
		hx_obj_6 := map[string]any{}
		hx_obj_6["id"] = hxrt.StringFromLiteral("one")
		return hxrt.AnonymousFromMap(hx_obj_6)
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
		var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("selected-item:"), func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
		}(hxrt.AnonymousGet(hx_structural_recover(items.Get(0)), "id"))))
		hxrt.Println(v_4)
	case 1:
		_g_5 := func(hx_value_11 any) *string {
			if hx_value_11 == nil {
				var hx_zero_12 *string
				return hx_zero_12
			}
			return hx_value_11.(*string)
		}(_g_3.params[0])
		_ = _g_5
	}
}

func selectItem(outcome *ItemOutcome) *ItemOutcome {
	var hx_switch_13 hxrt.Anonymous
	switch outcome.tag {
	case 0:
		_g := func(hx_value_14 any) *hxrt.Array {
			if hx_value_14 == nil {
				var hx_zero_15 *hxrt.Array
				return hx_zero_15
			}
			return hx_value_14.(*hxrt.Array)
		}(outcome.params[0])
		items := _g
		hx_switch_13 = hx_structural_recover(hx_structural_recover(hx_structural_recover(items.Get(0))))
	case 1:
		_g_1 := func(hx_value_16 any) *string {
			if hx_value_16 == nil {
				var hx_zero_17 *string
				return hx_zero_17
			}
			return hx_value_16.(*string)
		}(outcome.params[0])
		message := _g_1
		return ItemOutcome_ItemFailure(message)
	}
	selected := hx_switch_13
	return ItemOutcome_ItemSuccess(hxrt.NewArray(selected))
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_1 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_1.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_1.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_1.high = func(hx_value_2 any) int {
						if hx_value_2 == nil {
							var hx_zero_3 int
							return hx_zero_3
						}
						return hx_value_2.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_1.low = func(hx_value_4 any) int {
						if hx_value_4 == nil {
							var hx_zero_5 int
							return hx_zero_5
						}
						return hx_value_4.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}
