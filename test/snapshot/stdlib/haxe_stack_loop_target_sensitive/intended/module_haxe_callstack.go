package main

import "snapshot/hxrt"

func haxe___CallStack__CallStack_Impl__asArray(this1 *hxrt.Array) *hxrt.Array {
	return this1
}

func haxe___CallStack__CallStack_Impl__callStack() *hxrt.Array {
	return hxrt.NewArray()
}

func haxe___CallStack__CallStack_Impl__copy(this1 *hxrt.Array) any {
	return this1.Copy()
}

func haxe___CallStack__CallStack_Impl__exceptionStack(fullStack bool) *hxrt.Array {
	return hxrt.NewArray()
}

func haxe___CallStack__CallStack_Impl__exceptionToString(e *hxrt.ExceptionValue) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Exception: "), hxrt.ExceptionMessage(e))
}

func haxe___CallStack__CallStack_Impl__get(this1 *hxrt.Array, index int) *haxe__StackItem {
	return func(hx_value_1 any) *haxe__StackItem {
		if hx_value_1 == nil {
			var hx_zero_2 *haxe__StackItem
			return hx_zero_2
		}
		return hx_value_1.(*haxe__StackItem)
	}(this1.Get(index))
}

func haxe___CallStack__CallStack_Impl__get_length(this1 *hxrt.Array) int {
	return this1.Len()
}

func haxe___CallStack__CallStack_Impl__itemToString(item *haxe__StackItem) *string {
	var hx_switch_3 *string
	switch item.tag {
	case 0:
		hx_switch_3 = hxrt.StringFromLiteral("a C function")
	case 1:
		_g := item.params[0].(*string)
		m := _g
		hx_switch_3 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("module "), m)
	case 2:
		_g_1 := item.params[0].(*haxe__StackItem)
		_g1 := item.params[1].(*string)
		_g2 := item.params[2].(int)
		var _g3 any = item.params[3]
		inner := _g_1
		file := _g1
		line := _g2
		var column any = _g3
		var hx_if_4 *string
		if inner == nil {
			hx_if_4 = file
		} else {
			hx_if_4 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(haxe___CallStack__CallStack_Impl__itemToString(inner), hxrt.StringFromLiteral(" (")), file)
		}
		rendered := hx_if_4
		rendered = hxrt.StringConcatStringPtr(rendered, hxrt.StringConcatAny(hxrt.StringFromLiteral(" line "), line))
		if hxrt.IntFromNullableAny(column) > 0 {
			rendered = hxrt.StringConcatStringPtr(rendered, hxrt.StringConcatAny(hxrt.StringFromLiteral(" column "), column))
		}
		if inner != nil {
			rendered = hxrt.StringConcatStringPtr(rendered, hxrt.StringFromLiteral(")"))
		}
		hx_switch_3 = rendered
	case 3:
		_g_2 := item.params[0].(*string)
		_g1_1 := item.params[1].(*string)
		classname := _g_2
		method := _g1_1
		hx_switch_3 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
			var hx_if_5 *string
			if hxrt.StringEqualStringPtr(classname, nil) {
				hx_if_5 = hxrt.StringFromLiteral("<unknown>")
			} else {
				hx_if_5 = classname
			}
			return hx_if_5
		}(), hxrt.StringFromLiteral(".")), method)
	case 4:
		var _g_3 any = item.params[0]
		var v any = _g_3
		hx_switch_3 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("local function #"), hxrt.StdString(v))
	}
	return hx_switch_3
}

var haxe___CallStack__CallStack_Impl__length int

func haxe___CallStack__CallStack_Impl__subtract(this1 *hxrt.Array, stack any) any {
	return this1
}

func haxe___CallStack__CallStack_Impl__toString(stack any) *string {
	out := hxrt.StringFromLiteral("")
	_g := 0
	_g1 := func(hx_value_6 any) *hxrt.Array {
		if hx_value_6 == nil {
			var hx_zero_7 *hxrt.Array
			return hx_zero_7
		}
		return hx_value_6.(*hxrt.Array)
	}(stack)
	for _g < _g1.Len() {
		item := func(hx_value_8 any) *haxe__StackItem {
			if hx_value_8 == nil {
				var hx_zero_9 *haxe__StackItem
				return hx_zero_9
			}
			return hx_value_8.(*haxe__StackItem)
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
