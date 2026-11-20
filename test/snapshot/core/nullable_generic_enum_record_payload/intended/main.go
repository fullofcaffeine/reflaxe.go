package main

import "snapshot/hxrt"

type Outcome struct {
	tag    int
	params []any
}

func (self *Outcome) HxEnumValue() (string, []any) {
	return [2]string{"Success", "Failure"}[self.tag], self.params
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
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("null payload changed")))
	}
	if !hxrt.StringEqualStringPtr(describe(Outcome_Success(func() hxrt.Anonymous {
		hx_obj_13 := map[string]any{}
		hx_obj_13["name"] = hxrt.StringFromLiteral("kept")
		return hxrt.AnonymousFromMap(hx_obj_13)
	}())), hxrt.StringFromLiteral("name=kept")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("record payload changed")))
	}
	if !hxrt.StringEqualStringPtr(describeInterface(Outcome_Success(nil)), hxrt.StringFromLiteral("missing")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("null interface payload changed")))
	}
	if !hxrt.StringEqualStringPtr(describeInterface(Outcome_Success(New_NamedItem())), hxrt.StringFromLiteral("name=kept")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("interface payload changed")))
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

func (self *NamedItem) HxClassName() string {
	return "NamedItem"
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *NamedItem:
		return func(hx_structural_source_1 *NamedItem) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "name":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_2 := hx_structural_source_1
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_2.__hx_this.name()
						}, func() any {
							return hx_receiver_2.__hx_this
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
	case *haxe__Exception:
		return func(hx_structural_source_3 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_3, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_3.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_3.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_3.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_3.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_6 := hx_structural_source_3
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_6.__hx_this.get_message()
						}, func() any {
							return hx_receiver_6.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_7 := hx_structural_source_3
						return hxrt.BoundFunction(func() any {
							return hx_receiver_7.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_7.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_8 := hx_structural_source_3
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_8.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_8.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_9 := hx_structural_source_3
						return hxrt.BoundFunction(func() any {
							return hx_receiver_9.__hx_this.get_native()
						}, func() any {
							return hx_receiver_9.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_10 := hx_structural_source_3
						return hxrt.BoundFunction(func() any {
							return hx_receiver_10.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_10.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_11 := hx_structural_source_3
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_11.__hx_this.toString()
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_12 := hx_structural_source_3
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_12.__hx_this.details()
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_3.savedMessage = func(hx_value_4 any) *string {
						if hx_value_4 == nil {
							var hx_zero_5 *string
							return hx_zero_5
						}
						return hx_value_4.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_3.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_3.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_3.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_13 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_13, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_13.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_14 := hx_structural_source_13
						return hxrt.BoundFunction(func() any {
							return hx_receiver_14.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_14.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_13.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_13.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_13.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_13.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_17 := hx_structural_source_13
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_17.__hx_this.get_message()
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_18 := hx_structural_source_13
						return hxrt.BoundFunction(func() any {
							return hx_receiver_18.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_19 := hx_structural_source_13
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_19.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_20 := hx_structural_source_13
						return hxrt.BoundFunction(func() any {
							return hx_receiver_20.__hx_this.get_native()
						}, func() any {
							return hx_receiver_20.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_21 := hx_structural_source_13
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_21.__hx_this.toString()
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_22 := hx_structural_source_13
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_22.__hx_this.details()
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_13.savedMessage = func(hx_value_15 any) *string {
						if hx_value_15 == nil {
							var hx_zero_16 *string
							return hx_zero_16
						}
						return hx_value_15.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_13.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_13.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_13.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
