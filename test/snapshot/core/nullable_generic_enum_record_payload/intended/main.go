package main

import "snapshot/hxrt"

type Outcome struct {
	tag    int
	params []any
}

func Outcome_Success(value any) *Outcome {
	enumValue := &Outcome{tag: 0}
	enumValue.params = []any{value}
	return enumValue
}

func Outcome_Failure(message *string) *Outcome {
	enumValue := &Outcome{tag: 1}
	enumValue.params = []any{message}
	return enumValue
}

func describe(outcome *Outcome) *string {
	var hx_switch_1 *string
	switch outcome.tag {
	case 0:
		_g := func(hx_value_2 any) map[string]any {
			if hx_value_2 == nil {
				var hx_zero_3 map[string]any
				return hx_zero_3
			}
			return hx_value_2.(map[string]any)
		}(outcome.params[0])
		value := _g
		var hx_if_7 *string
		if value == nil {
			hx_if_7 = hxrt.StringFromLiteral("missing")
		} else {
			hx_if_7 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("name="), func(hx_obj_4 map[string]any) *string {
				hx_field_5 := hx_obj_4["name"]
				if hx_field_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_field_5.(*string)
			}(value))
		}
		hx_switch_1 = hx_if_7
	case 1:
		_g_1 := func(hx_value_8 any) *string {
			if hx_value_8 == nil {
				var hx_zero_9 *string
				return hx_zero_9
			}
			return hx_value_8.(*string)
		}(outcome.params[0])
		message := _g_1
		hx_switch_1 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("failure="), message)
	}
	return hx_switch_1
}

func main() {
	if !hxrt.StringEqualStringPtr(describe(Outcome_Success(nil)), hxrt.StringFromLiteral("missing")) {
		hxrt.Throw(hxrt.StringFromLiteral("null payload changed"))
	}
	if !hxrt.StringEqualStringPtr(describe(Outcome_Success(func() map[string]any {
		hx_obj_10 := map[string]any{}
		hx_obj_10["name"] = hxrt.StringFromLiteral("kept")
		return hx_obj_10
	}())), hxrt.StringFromLiteral("name=kept")) {
		hxrt.Throw(hxrt.StringFromLiteral("record payload changed"))
	}
}
