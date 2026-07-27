package main

import (
	"reflect"
	"snapshot/hxrt"
)

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type EKey struct {
	tag    int
	params []any
}

var EKey_A *EKey = &EKey{tag: 0}

func EKey_B(v int) *EKey {
	enumValue := &EKey{tag: 1}
	enumValue.params = []any{v}
	return enumValue
}

type I_Box interface {
}

type Box struct {
	__hx_this I_Box
	id        int
}

func New_Box(id int) *Box {
	self := &Box{}
	self.__hx_this = self
	self.id = id
	return self
}

func main() {
	sm := New_haxe__ds__StringMap()
	sm.__hx_this.set(hxrt.StringFromLiteral("a"), 1)
	av := hxrt.IntFromNullableAny(func(hx_value_1 any) any {
		if hx_value_1 == nil {
			return nil
		}
		return hx_value_1.(int)
	}(sm.__hx_this.get(hxrt.StringFromLiteral("a"))))
	hxrt.Println(any(av))
	om := New_haxe__ds__ObjectMap()
	box := New_Box(7)
	om.__hx_this.set(box, hxrt.StringFromLiteral("box"))
	ov := func(hx_value_2 any) *string {
		if hx_value_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_value_2.(*string)
	}(om.__hx_this.get(box))
	hxrt.Println(any(ov))
	em := New_haxe__ds__EnumValueMap()
	em.__hx_this.set(EKey_A, hxrt.StringFromLiteral("enum"))
	ev := func(hx_value_4 any) *string {
		if hx_value_4 == nil {
			var hx_zero_5 *string
			return hx_zero_5
		}
		return hx_value_4.(*string)
	}(em.__hx_this.get(EKey_A))
	hxrt.Println(any(ev))
	list := New_haxe__ds__List()
	list.__hx_this.add(4)
	list.__hx_this.add(5)
	var v any = any(list.length)
	hxrt.Println(v)
	var v_1 any = any(func(hx_value_6 any) any {
		if hx_value_6 == nil {
			return nil
		}
		return hx_value_6.(int)
	}(list.__hx_this.first()))
	hxrt.Println(v_1)
	var v_2 any = any(func(hx_value_7 any) any {
		if hx_value_7 == nil {
			return nil
		}
		return hx_value_7.(int)
	}(list.__hx_this.last()))
	hxrt.Println(v_2)
	var v_3 any = any(func(hx_value_8 any) any {
		if hx_value_8 == nil {
			return nil
		}
		return hx_value_8.(int)
	}(list.__hx_this.pop()))
	hxrt.Println(v_3)
	var v_4 any = any(list.length)
	hxrt.Println(v_4)
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := obj.(type) {
	case *haxe__ds__EnumValueMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__List:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__ObjectMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__ds__EnumValueMap:
		return hxrt__generated_method_field__haxe__ds__EnumValueMap(value, key)
	case *haxe__ds__List:
		return hxrt__generated_method_field__haxe__ds__List(value, key)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_method_field__haxe__ds__ObjectMap(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_method_field__haxe__ds__StringMap(value, key)
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return hxrt__generated_method_field__haxe__ds___EnumValueMap__EnumValueTreeNode(value, key)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_method_field__haxe__ds___List__GoListIterator(value, key)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_method_field__haxe__ds___List__GoListKeyValueIterator(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__haxe__ds__EnumValueMap(value *haxe__ds__EnumValueMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "balance":
		return value.balance
	case "clear":
		return value.clear
	case "compare":
		return value.compare
	case "compareArg":
		return value.compareArg
	case "compareArgs":
		return value.compareArgs
	case "copy":
		return value.copy
	case "copyIMap":
		return value.copyIMap
	case "exists":
		return value.exists
	case "existsIMap":
		return value.existsIMap
	case "get":
		return value.get
	case "getIMap":
		return value.getIMap
	case "iterator":
		return value.iterator
	case "keyValueIterator":
		return value.keyValueIterator
	case "keys":
		return value.keys
	case "merge":
		return value.merge
	case "minBinding":
		return value.minBinding
	case "remove":
		return value.remove
	case "removeIMap":
		return value.removeIMap
	case "removeLoop":
		return value.removeLoop
	case "removeMinBinding":
		return value.removeMinBinding
	case "set":
		return value.set
	case "setIMap":
		return value.setIMap
	case "setLoop":
		return value.setLoop
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__List(value *haxe__ds__List, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return value.add
	case "clear":
		return value.clear
	case "filter":
		return value.filter
	case "first":
		return value.first
	case "isEmpty":
		return value.isEmpty
	case "iterator":
		return value.iterator
	case "join":
		return value.join
	case "keyValueIterator":
		return value.keyValueIterator
	case "last":
		return value.last
	case "map":
		return value.map_
	case "pop":
		return value.pop
	case "push":
		return value.push
	case "remove":
		return value.remove
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__ObjectMap(value *haxe__ds__ObjectMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return value.clear
	case "copy":
		return value.copy
	case "copyIMap":
		return value.copyIMap
	case "exists":
		return value.exists
	case "existsIMap":
		return value.existsIMap
	case "get":
		return value.get
	case "getIMap":
		return value.getIMap
	case "iterator":
		return value.iterator
	case "keyValueIterator":
		return value.keyValueIterator
	case "keys":
		return value.keys
	case "remove":
		return value.remove
	case "removeIMap":
		return value.removeIMap
	case "set":
		return value.set
	case "setIMap":
		return value.setIMap
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return value.clear
	case "copy":
		return value.copy
	case "copyIMap":
		return value.copyIMap
	case "exists":
		return value.exists
	case "existsIMap":
		return value.existsIMap
	case "get":
		return value.get
	case "getIMap":
		return value.getIMap
	case "iterator":
		return value.iterator
	case "keyValueIterator":
		return value.keyValueIterator
	case "keys":
		return value.keys
	case "remove":
		return value.remove
	case "removeIMap":
		return value.removeIMap
	case "set":
		return value.set
	case "setIMap":
		return value.setIMap
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds___EnumValueMap__EnumValueTreeNode(value *haxe__ds___EnumValueMap__EnumValueTreeNode, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "getHeight":
		return value.getHeight
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds___List__GoListIterator(value *haxe__ds___List__GoListIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return value.hasNext
	case "next":
		return value.next
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds___List__GoListKeyValueIterator(value *haxe__ds___List__GoListKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return value.hasNext
	case "next":
		return value.next
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return value.hasNext
	case "next":
		return value.next
	}
	return nil
}

type Type struct {
}

func hxrt_typeCreateEmpty__Box() *Box {
	instance := &Box{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe___Int64_____Int64() *haxe___Int64_____Int64 {
	instance := &haxe___Int64_____Int64{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	instance := &haxe__ds__EnumValueMap{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds__List() *haxe__ds__List {
	instance := &haxe__ds__List{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	instance := &haxe__ds__ObjectMap{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds__StringMap() *haxe__ds__StringMap {
	instance := &haxe__ds__StringMap{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds___EnumValueMap__EnumValueTreeNode() *haxe__ds___EnumValueMap__EnumValueTreeNode {
	instance := &haxe__ds___EnumValueMap__EnumValueTreeNode{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds___List__GoListIterator() *haxe__ds___List__GoListIterator {
	instance := &haxe__ds___List__GoListIterator{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds___List__GoListKeyValueIterator() *haxe__ds___List__GoListKeyValueIterator {
	instance := &haxe__ds___List__GoListKeyValueIterator{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__MapKeyValueIterator() *haxe__iterators__MapKeyValueIterator {
	instance := &haxe__iterators__MapKeyValueIterator{}
	instance.__hx_this = instance
	return instance
}

type ValueType struct {
	tag    int
	params []any
}

var ValueType_TNull *ValueType = &ValueType{tag: 0, params: []any{}}

var ValueType_TInt *ValueType = &ValueType{tag: 1, params: []any{}}

var ValueType_TFloat *ValueType = &ValueType{tag: 2, params: []any{}}

var ValueType_TBool *ValueType = &ValueType{tag: 3, params: []any{}}

var ValueType_TObject *ValueType = &ValueType{tag: 4, params: []any{}}

var ValueType_TFunction *ValueType = &ValueType{tag: 5, params: []any{}}

var ValueType_TUnknown *ValueType = &ValueType{tag: 8, params: []any{}}

func ValueType_TClass(c any) *ValueType {
	return &ValueType{tag: 6, params: []any{c}}
}

func ValueType_TEnum(e any) *ValueType {
	return &ValueType{tag: 7, params: []any{e}}
}

func hxrt_typeCallAny(callable any, args []any) (any, bool) {
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	if callable == nil {
		return nil, false
	}
	fn := reflect.ValueOf(callable)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, false
	}
	fnType := fn.Type()
	if fnType.NumIn() != len(args) {
		return nil, false
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		paramType := fnType.In(i)
		arg := args[i]
		if arg == nil {
			in[i] = reflect.Zero(paramType)
			continue
		}
		v := reflect.ValueOf(arg)
		if v.IsValid() && v.Type().AssignableTo(paramType) {
			in[i] = v
			continue
		}
		if v.IsValid() && v.Type().ConvertibleTo(paramType) {
			in[i] = v.Convert(paramType)
			continue
		}
		if paramType.Kind() == reflect.Interface && v.IsValid() {
			in[i] = v
			continue
		}
		return nil, false
	}
	out := fn.Call(in)
	if len(out) == 0 {
		return nil, true
	}
	first := out[0]
	if !first.IsValid() {
		return nil, true
	}
	result = first.Interface()
	ok = true
	return result, ok
}

func hxrt_typeArrayValues(value *hxrt.Array) []any {
	if value == nil {
		return []any{}
	}
	return value.Values()
}

func hxrt_typeResolvedClassName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeClassValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeClassValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeResolvedEnumName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeEnumValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeEnumValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeCreateClassInstance(className string, args []any) (any, bool) {
	switch className {
	case "Box":
		return hxrt_typeCallAny(New_Box, args)
	case "Main":
		return nil, false
	case "Reflect":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.ds.EnumValueMap":
		return hxrt_typeCallAny(New_haxe__ds__EnumValueMap, args)
	case "haxe.ds.List":
		return hxrt_typeCallAny(New_haxe__ds__List, args)
	case "haxe.ds.ObjectMap":
		return hxrt_typeCallAny(New_haxe__ds__ObjectMap, args)
	case "haxe.ds.StringMap":
		return hxrt_typeCallAny(New_haxe__ds__StringMap, args)
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
		return hxrt_typeCallAny(New_haxe__ds___EnumValueMap__EnumValueTreeNode, args)
	case "haxe.ds._List.GoListIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListIterator, args)
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListKeyValueIterator, args)
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__MapKeyValueIterator, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "Box":
		return hxrt_typeCreateEmpty__Box(), true
	case "haxe._Int64.___Int64":
		return hxrt_typeCreateEmpty__haxe___Int64_____Int64(), true
	case "haxe.ds.EnumValueMap":
		return hxrt_typeCreateEmpty__haxe__ds__EnumValueMap(), true
	case "haxe.ds.List":
		return hxrt_typeCreateEmpty__haxe__ds__List(), true
	case "haxe.ds.ObjectMap":
		return hxrt_typeCreateEmpty__haxe__ds__ObjectMap(), true
	case "haxe.ds.StringMap":
		return hxrt_typeCreateEmpty__haxe__ds__StringMap(), true
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
		return hxrt_typeCreateEmpty__haxe__ds___EnumValueMap__EnumValueTreeNode(), true
	case "haxe.ds._List.GoListIterator":
		return hxrt_typeCreateEmpty__haxe__ds___List__GoListIterator(), true
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt_typeCreateEmpty__haxe__ds___List__GoListKeyValueIterator(), true
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__MapKeyValueIterator(), true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "EKey":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return EKey_A, true
			case 1:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(EKey_B, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "A":
			if len(args) != 0 {
				return nil, false
			}
			return EKey_A, true
		case "B":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(EKey_B, args)
		default:
			return nil, false
		}
	case "ValueType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TNull, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TInt, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFloat, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TBool, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TObject, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFunction, true
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TClass, args)
			case 7:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TEnum, args)
			case 8:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TUnknown, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TNull":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TNull, true
		case "TInt":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TInt, true
		case "TFloat":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFloat, true
		case "TBool":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TBool, true
		case "TObject":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TObject, true
		case "TFunction":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFunction, true
		case "TClass":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TClass, args)
		case "TEnum":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TEnum, args)
		case "TUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TUnknown, true
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeClassValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeClassValue:
		copyValue := value
		return &copyValue
	case *hxrt.Array:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")}
	case *Box:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Box")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe__ds__EnumValueMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.EnumValueMap")}
	case *haxe__ds__List:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.List")}
	case *haxe__ds__ObjectMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.ObjectMap")}
	case *haxe__ds__StringMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.StringMap")}
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._EnumValueMap.EnumValueTreeNode")}
	case *haxe__ds___List__GoListIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._List.GoListIterator")}
	case *haxe__ds___List__GoListKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._List.GoListKeyValueIterator")}
	case *haxe__iterators__MapKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.MapKeyValueIterator")}
	default:
		return nil
	}
}

func Type_getEnum(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeEnumValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeEnumValue:
		copyValue := value
		return &copyValue
	case *EKey:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("EKey")}
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	default:
		return nil
	}
}

func Type_getSuperClass(c any) any {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	switch className {
	case "Box":
		return nil
	case "Main":
		return nil
	case "Reflect":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.ds.EnumValueMap":
		return nil
	case "haxe.ds.List":
		return nil
	case "haxe.ds.ObjectMap":
		return nil
	case "haxe.ds.StringMap":
		return nil
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
		return nil
	case "haxe.ds._List.GoListIterator":
		return nil
	case "haxe.ds._List.GoListKeyValueIterator":
		return nil
	case "haxe.iterators.MapKeyValueIterator":
		return nil
	default:
		return nil
	}
}

func Type_getClassName(c any) *string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(className)
}

func Type_getClassFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Box":
		return hxrt.NewArray()
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("main"))
	case "Reflect":
		return hxrt.NewArray(hxrt.StringFromLiteral("callMethod"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("compareMethods"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("deleteField"), hxrt.StringFromLiteral("field"), hxrt.StringFromLiteral("fields"), hxrt.StringFromLiteral("getProperty"), hxrt.StringFromLiteral("hasField"), hxrt.StringFromLiteral("isEnumValue"), hxrt.StringFromLiteral("isFunction"), hxrt.StringFromLiteral("isObject"), hxrt.StringFromLiteral("makeVarArgs"), hxrt.StringFromLiteral("setField"), hxrt.StringFromLiteral("setProperty"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.ds.EnumValueMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("isEnumValue"), hxrt.StringFromLiteral("keysLoop"), hxrt.StringFromLiteral("valuesLoop"))
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("sameValue"))
	case "haxe.ds.ObjectMap":
		return hxrt.NewArray()
	case "haxe.ds.StringMap":
		return hxrt.NewArray()
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
		return hxrt.NewArray()
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray()
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray()
	default:
		return hxrt.NewArray()
	}
}

func Type_getInstanceFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Box":
		return hxrt.NewArray(hxrt.StringFromLiteral("id"))
	case "Main":
		return hxrt.NewArray()
	case "Reflect":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.ds.EnumValueMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("balance"), hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("compareArg"), hxrt.StringFromLiteral("compareArgs"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("merge"), hxrt.StringFromLiteral("minBinding"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("removeLoop"), hxrt.StringFromLiteral("removeMinBinding"), hxrt.StringFromLiteral("root"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("setLoop"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("filter"), hxrt.StringFromLiteral("first"), hxrt.StringFromLiteral("isEmpty"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("join"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("last"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds.ObjectMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds.StringMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
		return hxrt.NewArray(hxrt.StringFromLiteral("getHeight"), hxrt.StringFromLiteral("height"), hxrt.StringFromLiteral("key"), hxrt.StringFromLiteral("left"), hxrt.StringFromLiteral("right"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("value"))
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("next"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getEnumName(e any) *string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(enumName)
}

func Type_resolveClass(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "Box":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Reflect":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.EnumValueMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.List":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.ObjectMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.StringMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.MapKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_resolveEnum(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "EKey":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_createInstance(cl any, args *hxrt.Array) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, hxrt_typeArrayValues(args))
	if !ok {
		return nil
	}
	return instance
}

func Type_createEmptyInstance(cl any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassEmptyInstance(className)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEnum(e any, constr *string, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_enumConstructor(e any) *string {
	if hxrt.AnyEqualsNull(e) {
		return nil
	}
	switch value := e.(type) {
	case *EKey:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("A")
		case 1:
			return hxrt.StringFromLiteral("B")
		default:
			return nil
		}
	case *ValueType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TNull")
		case 1:
			return hxrt.StringFromLiteral("TInt")
		case 2:
			return hxrt.StringFromLiteral("TFloat")
		case 3:
			return hxrt.StringFromLiteral("TBool")
		case 4:
			return hxrt.StringFromLiteral("TObject")
		case 5:
			return hxrt.StringFromLiteral("TFunction")
		case 6:
			return hxrt.StringFromLiteral("TClass")
		case 7:
			return hxrt.StringFromLiteral("TEnum")
		case 8:
			return hxrt.StringFromLiteral("TUnknown")
		default:
			return nil
		}
	default:
		return nil
	}
}

func Type_enumIndex(e any) int {
	if hxrt.AnyEqualsNull(e) {
		return -1
	}
	switch value := e.(type) {
	case *EKey:
		if value == nil {
			return -1
		}
		return value.tag
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "EKey":
		return hxrt.NewArray(hxrt.StringFromLiteral("A"), hxrt.StringFromLiteral("B"))
	case "ValueType":
		return hxrt.NewArray(hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown"))
	default:
		return hxrt.NewArray()
	}
}

func Type_enumParameters(e any) *hxrt.Array {
	if hxrt.AnyEqualsNull(e) {
		return hxrt.NewArray()
	}
	switch value := e.(type) {
	case *EKey:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *ValueType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	default:
		return hxrt.NewArray()
	}
}

func Type_allEnums(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "EKey":
		return hxrt.NewArray(EKey_A)
	case "ValueType":
		return hxrt.NewArray(ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown)
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) *ValueType {
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
	}
	switch v.(type) {
	case *hxrt__TypeClassValue, hxrt__TypeClassValue, *hxrt__TypeEnumValue, hxrt__TypeEnumValue:
		return ValueType_TObject
	}
	if enumValue := Type_getEnum(v); enumValue != nil {
		return ValueType_TEnum(enumValue)
	}
	if classValue := Type_getClass(v); classValue != nil {
		return ValueType_TClass(classValue)
	}
	switch v.(type) {
	case bool:
		return ValueType_TBool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return ValueType_TInt
	case float32, float64:
		return ValueType_TFloat
	case string, *string:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")})
	case *hxrt.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	}
	ref := reflect.ValueOf(v)
	if !ref.IsValid() {
		return ValueType_TNull
	}
	switch ref.Kind() {
	case reflect.Func:
		return ValueType_TFunction
	case reflect.Slice, reflect.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:
		return ValueType_TObject
	default:
		return ValueType_TUnknown
	}
}

func Type_enumEq(a any, b any) bool {
	return reflect.DeepEqual(a, b)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__EnumValueMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__List:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__ObjectMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Box:
		return hxrt__generated_field_lookup__Box(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe__ds__EnumValueMap:
		return hxrt__generated_field_lookup__haxe__ds__EnumValueMap(value, key)
	case *haxe__ds__List:
		return hxrt__generated_field_lookup__haxe__ds__List(value, key)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_field_lookup__haxe__ds__ObjectMap(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_lookup__haxe__ds__StringMap(value, key)
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return hxrt__generated_field_lookup__haxe__ds___EnumValueMap__EnumValueTreeNode(value, key)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_field_lookup__haxe__ds___List__GoListIterator(value, key)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__ds___List__GoListKeyValueIterator(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__EnumValueMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__List:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__ObjectMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Box:
		return hxrt__generated_field_has__Box(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe__ds__EnumValueMap:
		return hxrt__generated_field_has__haxe__ds__EnumValueMap(value, key)
	case *haxe__ds__List:
		return hxrt__generated_field_has__haxe__ds__List(value, key)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_field_has__haxe__ds__ObjectMap(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_has__haxe__ds__StringMap(value, key)
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return hxrt__generated_field_has__haxe__ds___EnumValueMap__EnumValueTreeNode(value, key)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_field_has__haxe__ds___List__GoListIterator(value, key)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_field_has__haxe__ds___List__GoListKeyValueIterator(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__EnumValueMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__List:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__ObjectMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Box:
		return hxrt__generated_field_set__Box(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe__ds__EnumValueMap:
		return hxrt__generated_field_set__haxe__ds__EnumValueMap(value, key, incoming)
	case *haxe__ds__List:
		return hxrt__generated_field_set__haxe__ds__List(value, key, incoming)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_field_set__haxe__ds__ObjectMap(value, key, incoming)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_set__haxe__ds__StringMap(value, key, incoming)
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return hxrt__generated_field_set__haxe__ds___EnumValueMap__EnumValueTreeNode(value, key, incoming)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_field_set__haxe__ds___List__GoListIterator(value, key, incoming)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_field_set__haxe__ds___List__GoListKeyValueIterator(value, key, incoming)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := object.(type) {
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__EnumValueMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__List:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__ObjectMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *Box:
		return hxrt.NewArray(hxrt.StringFromLiteral("id"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__EnumValueMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("root"))
	case *haxe__ds__List:
		return hxrt.NewArray(hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("length"))
	case *haxe__ds__ObjectMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__ds__StringMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return hxrt.NewArray(hxrt.StringFromLiteral("left"), hxrt.StringFromLiteral("right"), hxrt.StringFromLiteral("key"), hxrt.StringFromLiteral("value"), hxrt.StringFromLiteral("height"))
	case *haxe__ds___List__GoListIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("index"))
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("index"))
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("keys"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Box(value *Box, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "id":
		return value.id
	}
	return nil
}

func hxrt__generated_field_has__Box(value *Box, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "id":
		return true
	}
	return false
}

func hxrt__generated_field_set__Box(value *Box, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "id":
		if incoming == nil {
			var zero int
			value.id = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.id = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return value.high
	case "low":
		return value.low
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__EnumValueMap(value *haxe__ds__EnumValueMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "root":
		return value.root
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__EnumValueMap(value *haxe__ds__EnumValueMap, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "root":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__EnumValueMap(value *haxe__ds__EnumValueMap, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "root":
		if incoming == nil {
			var zero *haxe__ds___EnumValueMap__EnumValueTreeNode
			value.root = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds___EnumValueMap__EnumValueTreeNode:
			value.root = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__List(value *haxe__ds__List, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "items":
		return value.items
	case "length":
		return value.length
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__List(value *haxe__ds__List, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "items":
		return true
	case "length":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__List(value *haxe__ds__List, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "items":
		if incoming == nil {
			var zero *hxrt.Array
			value.items = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.items = typed
			return true
		default:
			return false
		}
	case "length":
		if incoming == nil {
			var zero int
			value.length = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.length = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__ObjectMap(value *haxe__ds__ObjectMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "h":
		return value.h
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__ObjectMap(value *haxe__ds__ObjectMap, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__ObjectMap(value *haxe__ds__ObjectMap, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		if incoming == nil {
			var zero *hxrt.ObjectMapCell
			value.h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.ObjectMapCell:
			value.h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "h":
		return value.h
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__StringMap(value *haxe__ds__StringMap, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		if incoming == nil {
			var zero *hxrt.StringMapCell
			value.h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.StringMapCell:
			value.h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds___EnumValueMap__EnumValueTreeNode(value *haxe__ds___EnumValueMap__EnumValueTreeNode, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "height":
		return value.height
	case "key":
		return value.key
	case "left":
		return value.left
	case "right":
		return value.right
	case "value":
		return value.value
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds___EnumValueMap__EnumValueTreeNode(value *haxe__ds___EnumValueMap__EnumValueTreeNode, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "height":
		return true
	case "key":
		return true
	case "left":
		return true
	case "right":
		return true
	case "value":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds___EnumValueMap__EnumValueTreeNode(value *haxe__ds___EnumValueMap__EnumValueTreeNode, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "height":
		if incoming == nil {
			var zero int
			value.height = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.height = typed
			return true
		default:
			return false
		}
	case "key":
		if incoming == nil {
			var zero any
			value.key = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.key = typed
			return true
		default:
			return false
		}
	case "left":
		if incoming == nil {
			var zero *haxe__ds___EnumValueMap__EnumValueTreeNode
			value.left = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds___EnumValueMap__EnumValueTreeNode:
			value.left = typed
			return true
		default:
			return false
		}
	case "right":
		if incoming == nil {
			var zero *haxe__ds___EnumValueMap__EnumValueTreeNode
			value.right = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds___EnumValueMap__EnumValueTreeNode:
			value.right = typed
			return true
		default:
			return false
		}
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds___List__GoListIterator(value *haxe__ds___List__GoListIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "index":
		return value.index
	case "items":
		return value.items
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds___List__GoListIterator(value *haxe__ds___List__GoListIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		return true
	case "items":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds___List__GoListIterator(value *haxe__ds___List__GoListIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		if incoming == nil {
			var zero int
			value.index = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.index = typed
			return true
		default:
			return false
		}
	case "items":
		if incoming == nil {
			var zero *hxrt.Array
			value.items = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.items = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds___List__GoListKeyValueIterator(value *haxe__ds___List__GoListKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "index":
		return value.index
	case "items":
		return value.items
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds___List__GoListKeyValueIterator(value *haxe__ds___List__GoListKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		return true
	case "items":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds___List__GoListKeyValueIterator(value *haxe__ds___List__GoListKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		if incoming == nil {
			var zero int
			value.index = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.index = typed
			return true
		default:
			return false
		}
	case "items":
		if incoming == nil {
			var zero *hxrt.Array
			value.items = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.items = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "keys":
		return value.keys
	case "map":
		return value.map_
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		return true
	case "map":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		if incoming == nil {
			var zero map[string]any
			value.keys = zero
			return true
		}
		switch typed := incoming.(type) {
		case map[string]any:
			value.keys = typed
			return true
		default:
			return false
		}
	case "map":
		if incoming == nil {
			var zero haxe__IMap
			value.map_ = zero
			return true
		}
		switch typed := incoming.(type) {
		case haxe__IMap:
			value.map_ = typed
			return true
		default:
			return false
		}
	}
	return false
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	switch enumValue := value.(type) {
	case *EKey:
		return (enumValue != nil)
	default:
		return false
	}
}
