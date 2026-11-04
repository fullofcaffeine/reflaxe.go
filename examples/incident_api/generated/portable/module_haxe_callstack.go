package main

import "examples_incident_api_portable/hxrt"

func haxe___CallStack__CallStack_Impl__asArray(this1 *hxrt.Array) *hxrt.Array {
	return this1
}

func haxe___CallStack__CallStack_Impl__callStack() *hxrt.Array {
	return hxrt.NewArray()
}

func haxe___CallStack__CallStack_Impl__copy(this1 *hxrt.Array) any {
	return this1.Copy()
}

func haxe___CallStack__CallStack_Impl__equalItems(item1 *haxe__StackItem, item2 *haxe__StackItem) bool {
	var hx_if_33 bool
	if item1 == nil {
		var hx_if_1 bool
		if item2 == nil {
			hx_if_1 = true
		} else {
			hx_if_1 = false
		}
		hx_if_33 = hx_if_1
	} else {
		var hx_switch_2 bool
		switch item1.tag {
		case 0:
			var hx_if_4 bool
			if item2 == nil {
				hx_if_4 = false
			} else {
				var hx_if_3 bool
				if item2.tag == 0 {
					hx_if_3 = true
				} else {
					hx_if_3 = false
				}
				hx_if_4 = hx_if_3
			}
			hx_switch_2 = hx_if_4
		case 1:
			_g := func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(item1.params[0])
			var hx_if_10 bool
			if item2 == nil {
				hx_if_10 = false
			} else {
				var hx_if_9 bool
				if item2.tag == 1 {
					_g1 := func(hx_value_7 any) *string {
						if hx_value_7 == nil {
							var hx_zero_8 *string
							return hx_zero_8
						}
						return hx_value_7.(*string)
					}(item2.params[0])
					m2 := _g1
					m1 := _g
					hx_if_9 = hxrt.StringEqualStringPtr(m1, m2)
				} else {
					hx_if_9 = false
				}
				hx_if_10 = hx_if_9
			}
			hx_switch_2 = hx_if_10
		case 2:
			_g_1 := func(hx_value_11 any) *haxe__StackItem {
				if hx_value_11 == nil {
					var hx_zero_12 *haxe__StackItem
					return hx_zero_12
				}
				return hx_value_11.(*haxe__StackItem)
			}(item1.params[0])
			_g1_1 := func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(item1.params[1])
			_g2 := item1.params[2].(int)
			var _g3 any = item1.params[3]
			var hx_if_20 bool
			if item2 == nil {
				hx_if_20 = false
			} else {
				var hx_if_19 bool
				if item2.tag == 2 {
					_g4 := func(hx_value_15 any) *haxe__StackItem {
						if hx_value_15 == nil {
							var hx_zero_16 *haxe__StackItem
							return hx_zero_16
						}
						return hx_value_15.(*haxe__StackItem)
					}(item2.params[0])
					_g5 := func(hx_value_17 any) *string {
						if hx_value_17 == nil {
							var hx_zero_18 *string
							return hx_zero_18
						}
						return hx_value_17.(*string)
					}(item2.params[1])
					_g6 := item2.params[2].(int)
					var _g7 any = item2.params[3]
					item2_1 := _g4
					file2 := _g5
					line2 := _g6
					var col2 any = _g7
					var col1 any = _g3
					line1 := _g2
					file1 := _g1_1
					item1_1 := _g_1
					hx_if_19 = (((hxrt.StringEqualStringPtr(file1, file2) && (line1 == line2)) && (col1 == col2)) && haxe___CallStack__CallStack_Impl__equalItems(item1_1, item2_1))
				} else {
					hx_if_19 = false
				}
				hx_if_20 = hx_if_19
			}
			hx_switch_2 = hx_if_20
		case 3:
			_g_2 := func(hx_value_21 any) *string {
				if hx_value_21 == nil {
					var hx_zero_22 *string
					return hx_zero_22
				}
				return hx_value_21.(*string)
			}(item1.params[0])
			_g1_2 := func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(item1.params[1])
			var hx_if_30 bool
			if item2 == nil {
				hx_if_30 = false
			} else {
				var hx_if_29 bool
				if item2.tag == 3 {
					_g2_1 := func(hx_value_25 any) *string {
						if hx_value_25 == nil {
							var hx_zero_26 *string
							return hx_zero_26
						}
						return hx_value_25.(*string)
					}(item2.params[0])
					_g3_1 := func(hx_value_27 any) *string {
						if hx_value_27 == nil {
							var hx_zero_28 *string
							return hx_zero_28
						}
						return hx_value_27.(*string)
					}(item2.params[1])
					class2 := _g2_1
					method2 := _g3_1
					method1 := _g1_2
					class1 := _g_2
					hx_if_29 = (hxrt.StringEqualStringPtr(class1, class2) && hxrt.StringEqualStringPtr(method1, method2))
				} else {
					hx_if_29 = false
				}
				hx_if_30 = hx_if_29
			}
			hx_switch_2 = hx_if_30
		case 4:
			var _g_3 any = item1.params[0]
			var hx_if_32 bool
			if item2 == nil {
				hx_if_32 = false
			} else {
				var hx_if_31 bool
				if item2.tag == 4 {
					var _g1_3 any = item2.params[0]
					var v2 any = _g1_3
					var v1 any = _g_3
					hx_if_31 = (v1 == v2)
				} else {
					hx_if_31 = false
				}
				hx_if_32 = hx_if_31
			}
			hx_switch_2 = hx_if_32
		}
		hx_if_33 = hx_switch_2
	}
	return hx_if_33
}

func haxe___CallStack__CallStack_Impl__exceptionStack(fullStack bool) *hxrt.Array {
	return hxrt.NewArray()
}

func haxe___CallStack__CallStack_Impl__exceptionToString(e *haxe__Exception) *string {
	if e.__hx_this.get_previous() == nil {
		return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Exception: "), e.__hx_this.toString()), func() *string {
			var tmp any = e.__hx_this.get_stack()
			var hx_if_34 *string
			if hxrt.AnyEqualsNull(tmp) {
				hx_if_34 = hxrt.StringFromLiteral("null")
			} else {
				hx_if_34 = haxe___CallStack__CallStack_Impl__toString(tmp)
			}
			return hx_if_34
		}())
	}
	result := hxrt.StringFromLiteral("")
	e_1 := e
	var prev *haxe__Exception = nil
	for e_1 != nil {
		if prev == nil {
			result = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Exception: "), e_1.__hx_this.get_message()), func() *string {
				var tmp_1 any = e_1.__hx_this.get_stack()
				var hx_if_35 *string
				if hxrt.AnyEqualsNull(tmp_1) {
					hx_if_35 = hxrt.StringFromLiteral("null")
				} else {
					hx_if_35 = haxe___CallStack__CallStack_Impl__toString(tmp_1)
				}
				return hx_if_35
			}()), result)
		} else {
			var prevStack any = haxe___CallStack__CallStack_Impl__subtract(func(hx_value_36 any) *hxrt.Array {
				if hx_value_36 == nil {
					var hx_zero_37 *hxrt.Array
					return hx_zero_37
				}
				return hx_value_36.(*hxrt.Array)
			}(any(e_1.__hx_this.get_stack())), prev.__hx_this.get_stack())
			result = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Exception: "), e_1.__hx_this.get_message()), func() *string {
				var hx_if_38 *string
				if hxrt.AnyEqualsNull(prevStack) {
					hx_if_38 = hxrt.StringFromLiteral("null")
				} else {
					hx_if_38 = haxe___CallStack__CallStack_Impl__toString(prevStack)
				}
				return hx_if_38
			}()), hxrt.StringFromLiteral("\n\nNext ")), result)
		}
		prev = e_1
		e_1 = e_1.__hx_this.get_previous()
	}
	return result
}

func haxe___CallStack__CallStack_Impl__get(this1 *hxrt.Array, index int) *haxe__StackItem {
	return func(hx_value_39 any) *haxe__StackItem {
		if hx_value_39 == nil {
			var hx_zero_40 *haxe__StackItem
			return hx_zero_40
		}
		return hx_value_39.(*haxe__StackItem)
	}(this1.Get(index))
}

func haxe___CallStack__CallStack_Impl__get_length(this1 *hxrt.Array) int {
	return this1.Len()
}

func haxe___CallStack__CallStack_Impl__itemToString(item *haxe__StackItem) *string {
	var hx_switch_41 *string
	switch item.tag {
	case 0:
		hx_switch_41 = hxrt.StringFromLiteral("a C function")
	case 1:
		_g := func(hx_value_42 any) *string {
			if hx_value_42 == nil {
				var hx_zero_43 *string
				return hx_zero_43
			}
			return hx_value_42.(*string)
		}(item.params[0])
		m := _g
		hx_switch_41 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("module "), m)
	case 2:
		_g_1 := func(hx_value_44 any) *haxe__StackItem {
			if hx_value_44 == nil {
				var hx_zero_45 *haxe__StackItem
				return hx_zero_45
			}
			return hx_value_44.(*haxe__StackItem)
		}(item.params[0])
		_g1 := func(hx_value_46 any) *string {
			if hx_value_46 == nil {
				var hx_zero_47 *string
				return hx_zero_47
			}
			return hx_value_46.(*string)
		}(item.params[1])
		_g2 := item.params[2].(int)
		var _g3 any = item.params[3]
		inner := _g_1
		file := _g1
		line := _g2
		var column any = _g3
		var hx_if_48 *string
		if inner == nil {
			hx_if_48 = file
		} else {
			hx_if_48 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(haxe___CallStack__CallStack_Impl__itemToString(inner), hxrt.StringFromLiteral(" (")), file)
		}
		rendered := hx_if_48
		rendered = hxrt.StringConcatStringPtr(rendered, hxrt.StringConcatAny(hxrt.StringFromLiteral(" line "), line))
		if hxrt.IntFromNullableAny(column) > 0 {
			rendered = hxrt.StringConcatStringPtr(rendered, hxrt.StringConcatAny(hxrt.StringFromLiteral(" column "), column))
		}
		if inner != nil {
			rendered = hxrt.StringConcatStringPtr(rendered, hxrt.StringFromLiteral(")"))
		}
		hx_switch_41 = rendered
	case 3:
		_g_2 := func(hx_value_49 any) *string {
			if hx_value_49 == nil {
				var hx_zero_50 *string
				return hx_zero_50
			}
			return hx_value_49.(*string)
		}(item.params[0])
		_g1_1 := func(hx_value_51 any) *string {
			if hx_value_51 == nil {
				var hx_zero_52 *string
				return hx_zero_52
			}
			return hx_value_51.(*string)
		}(item.params[1])
		classname := _g_2
		method := _g1_1
		hx_switch_41 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
			var hx_if_53 *string
			if hxrt.StringEqualStringPtr(classname, nil) {
				hx_if_53 = hxrt.StringFromLiteral("<unknown>")
			} else {
				hx_if_53 = classname
			}
			return hx_if_53
		}(), hxrt.StringFromLiteral(".")), method)
	case 4:
		var _g_3 any = item.params[0]
		var v any = _g_3
		hx_switch_41 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("local function #"), hxrt.StdString(v))
	}
	return hx_switch_41
}

var haxe___CallStack__CallStack_Impl__length int

func haxe___CallStack__CallStack_Impl__subtract(this1 *hxrt.Array, stack any) any {
	startIndex := -1
	i := -1
	for func() int {
		i = int(int32((i + 1)))
		return i
	}() < this1.Len() {
		_g := 0
		_g1 := func(hx_value_54 any) *hxrt.Array {
			if hx_value_54 == nil {
				var hx_zero_55 *hxrt.Array
				return hx_zero_55
			}
			return hx_value_54.(*hxrt.Array)
		}(stack).Len()
		for _g < _g1 {
			hx_post_56 := _g
			_g = int(int32((_g + 1)))
			j := hx_post_56
			if haxe___CallStack__CallStack_Impl__equalItems(func(hx_value_57 any) *haxe__StackItem {
				if hx_value_57 == nil {
					var hx_zero_58 *haxe__StackItem
					return hx_zero_58
				}
				return hx_value_57.(*haxe__StackItem)
			}(this1.Get(i)), func(hx_value_61 any) *haxe__StackItem {
				if hx_value_61 == nil {
					var hx_zero_62 *haxe__StackItem
					return hx_zero_62
				}
				return hx_value_61.(*haxe__StackItem)
			}(func(hx_value_59 any) *hxrt.Array {
				if hx_value_59 == nil {
					var hx_zero_60 *hxrt.Array
					return hx_zero_60
				}
				return hx_value_59.(*hxrt.Array)
			}(stack).Get(j))) {
				if startIndex < 0 {
					startIndex = i
				}
				i = int(int32((i + 1)))
				if i >= this1.Len() {
					break
				}
			} else {
				startIndex = -1
			}
		}
		if startIndex >= 0 {
			break
		}
	}
	var hx_if_63 any
	if startIndex >= 0 {
		hx_if_63 = this1.SliceOptional(0, startIndex)
	} else {
		hx_if_63 = this1
	}
	return hx_if_63
}

func haxe___CallStack__CallStack_Impl__toString(stack any) *string {
	out := hxrt.StringFromLiteral("")
	_g := 0
	_g1 := func(hx_value_66 any) *hxrt.Array {
		if hx_value_66 == nil {
			var hx_zero_67 *hxrt.Array
			return hx_zero_67
		}
		return hx_value_66.(*hxrt.Array)
	}(any(func(hx_value_64 any) *hxrt.Array {
		if hx_value_64 == nil {
			var hx_zero_65 *hxrt.Array
			return hx_zero_65
		}
		return hx_value_64.(*hxrt.Array)
	}(stack)))
	for _g < _g1.Len() {
		item := func(hx_value_68 any) *haxe__StackItem {
			if hx_value_68 == nil {
				var hx_zero_69 *haxe__StackItem
				return hx_zero_69
			}
			return hx_value_68.(*haxe__StackItem)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("\nCalled from "), haxe___CallStack__CallStack_Impl__itemToString(item)))
	}
	return out
}

type haxe__StackItem struct {
	tag    int
	params []any
}

var haxe__StackItem_CFunction *haxe__StackItem = &haxe__StackItem{tag: 0}

func haxe__StackItem_Module(m *string) *haxe__StackItem {
	enumValue := &haxe__StackItem{tag: 1}
	enumValue.params = []any{m}
	return enumValue
}

func haxe__StackItem_FilePos(s *haxe__StackItem, file *string, line int, column any) *haxe__StackItem {
	enumValue := &haxe__StackItem{tag: 2}
	enumValue.params = []any{s, file, line, column}
	return enumValue
}

func haxe__StackItem_Method(classname *string, method *string) *haxe__StackItem {
	enumValue := &haxe__StackItem{tag: 3}
	enumValue.params = []any{classname, method}
	return enumValue
}

func haxe__StackItem_LocalFunction(v any) *haxe__StackItem {
	enumValue := &haxe__StackItem{tag: 4}
	enumValue.params = []any{v}
	return enumValue
}
