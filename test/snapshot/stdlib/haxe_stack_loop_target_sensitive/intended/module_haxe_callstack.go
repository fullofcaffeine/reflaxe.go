package main

import "snapshot/hxrt"

func haxe___CallStack__CallStack_Impl__asArray(this1 []*haxe__StackItem) []*haxe__StackItem {
	return this1
}

func haxe___CallStack__CallStack_Impl__callStack() []*haxe__StackItem {
	return []*haxe__StackItem{}
}

func haxe___CallStack__CallStack_Impl__copy(this1 []*haxe__StackItem) any {
	return func(src []*haxe__StackItem) []*haxe__StackItem {
		out := append([]*haxe__StackItem{}, src...)
		return out
	}(this1)
}

func haxe___CallStack__CallStack_Impl__exceptionStack(_fullStack bool) []*haxe__StackItem {
	return []*haxe__StackItem{}
}

func haxe___CallStack__CallStack_Impl__exceptionToString(e *hxrt.ExceptionValue) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Exception: "), hxrt.StdString(e))
}

func haxe___CallStack__CallStack_Impl__get(this1 []*haxe__StackItem, index int) *haxe__StackItem {
	return this1[index]
}

func haxe___CallStack__CallStack_Impl__get_length(this1 []*haxe__StackItem) int {
	return len(this1)
}

var haxe___CallStack__CallStack_Impl__length int

func haxe___CallStack__CallStack_Impl__subtract(this1 []*haxe__StackItem, _stack any) any {
	return this1
}

func haxe___CallStack__CallStack_Impl__toString(_stack any) *string {
	return hxrt.StringFromLiteral("")
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

func haxe__StackItem_FilePos(s *haxe__StackItem, file *string, line int, column int) *haxe__StackItem {
	enumValue := &haxe__StackItem{tag: 2}
	enumValue.params = []any{s, file, line, column}
	return enumValue
}

func haxe__StackItem_Method(classname *string, method *string) *haxe__StackItem {
	enumValue := &haxe__StackItem{tag: 3}
	enumValue.params = []any{classname, method}
	return enumValue
}

func haxe__StackItem_LocalFunction(v int) *haxe__StackItem {
	enumValue := &haxe__StackItem{tag: 4}
	enumValue.params = []any{v}
	return enumValue
}
