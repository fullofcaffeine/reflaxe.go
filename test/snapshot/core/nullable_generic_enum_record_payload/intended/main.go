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
		_g := hx_structural_recover(outcome.params[0])
		value := _g
		var hx_if_4 *string
		if value == nil {
			hx_if_4 = hxrt.StringFromLiteral("missing")
		} else {
			hx_if_4 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("name="), func(hx_value_2 any) *string {
				if hx_value_2 == nil {
					var hx_zero_3 *string
					return hx_zero_3
				}
				return hx_value_2.(*string)
			}(hxrt.AnonymousGet(value, "name")))
		}
		hx_switch_1 = hx_if_4
	case 1:
		_g_1 := func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(outcome.params[0])
		message := _g_1
		hx_switch_1 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("failure="), message)
	}
	return hx_switch_1
}

func describeInterface(outcome *Outcome) *string {
	var hx_switch_7 *string
	switch outcome.tag {
	case 0:
		var _g Named = func(hx_value_8 any) Named {
			if hx_value_8 == nil {
				var hx_zero_9 Named
				return hx_zero_9
			}
			return hx_value_8.(Named)
		}(hxrt.ObjectSource(outcome.params[0]))
		var value Named = _g
		var hx_if_10 *string
		if value == nil {
			hx_if_10 = hxrt.StringFromLiteral("missing")
		} else {
			hx_if_10 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("name="), value.name())
		}
		hx_switch_7 = hx_if_10
	case 1:
		_g_1 := func(hx_value_11 any) *string {
			if hx_value_11 == nil {
				var hx_zero_12 *string
				return hx_zero_12
			}
			return hx_value_11.(*string)
		}(outcome.params[0])
		message := _g_1
		hx_switch_7 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("failure="), message)
	}
	return hx_switch_7
}

func main() {
	if !hxrt.StringEqualStringPtr(describe(Outcome_Success(nil)), hxrt.StringFromLiteral("missing")) {
		hxrt.Throw(hxrt.StringFromLiteral("null payload changed"))
	}
	if !hxrt.StringEqualStringPtr(describe(Outcome_Success(func() hxrt.Anonymous {
		hx_obj_13 := map[string]any{}
		hx_obj_13["name"] = hxrt.StringFromLiteral("kept")
		return hxrt.AnonymousFromMap(hx_obj_13)
	}())), hxrt.StringFromLiteral("name=kept")) {
		hxrt.Throw(hxrt.StringFromLiteral("record payload changed"))
	}
	if !hxrt.StringEqualStringPtr(describeInterface(Outcome_Success(nil)), hxrt.StringFromLiteral("missing")) {
		hxrt.Throw(hxrt.StringFromLiteral("null interface payload changed"))
	}
	if !hxrt.StringEqualStringPtr(describeInterface(Outcome_Success(New_NamedItem())), hxrt.StringFromLiteral("name=kept")) {
		hxrt.Throw(hxrt.StringFromLiteral("interface payload changed"))
	}
}

type Named interface {
	name() *string
}

type I_NamedItem interface {
	name() *string
}

type NamedItem struct {
	__hx_this I_NamedItem
}

func (self *NamedItem) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_NamedItem() *NamedItem {
	return hx_init_NamedItem(hx_alloc_NamedItem(nil))
}

func hx_alloc_NamedItem(owner any) *NamedItem {
	self := &NamedItem{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_NamedItem)
	}
	return self
}

func hx_init_NamedItem(self *NamedItem) *NamedItem {
	return self
}

func (self *NamedItem) name() *string {
	return hxrt.StringFromLiteral("kept")
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *NamedItem:
		return func(hx_structural_source_3 *NamedItem) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_3, func(name string) hxrt.AnonymousField {
				switch name {
				case "name":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_4 := hx_structural_source_3
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_4.__hx_this.name()
						}, func() any {
							return hx_receiver_4.__hx_this
						}, "name")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}
