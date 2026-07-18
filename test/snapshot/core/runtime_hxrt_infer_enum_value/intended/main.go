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

type CollectionFeatureKey struct {
	tag    int
	params []any
}

func CollectionFeatureKey_Entry(value int) *CollectionFeatureKey {
	enumValue := &CollectionFeatureKey{tag: 0}
	enumValue.params = []any{value}
	return enumValue
}

func main() {
	values := New_haxe__ds__EnumValueMap()
	values.__hx_this.set(CollectionFeatureKey_Entry(1), hxrt.StringFromLiteral("one"))
	func(hx_value_1 any) bool {
		if hx_value_1 == nil {
			var hx_zero_2 bool
			return hx_zero_2
		}
		return hx_value_1.(bool)
	}(values.__hx_this.exists(CollectionFeatureKey_Entry(1)))
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := obj.(type) {
	case *haxe__ds__EnumValueMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
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
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return hxrt__generated_method_field__haxe__ds___EnumValueMap__EnumValueTreeNode(value, key)
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

type Type struct {
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
	case "Main":
		return nil, false
	case "Reflect":
		return nil, false
	case "haxe.ds.EnumValueMap":
		return hxrt_typeCallAny(New_haxe__ds__EnumValueMap, args)
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
		return hxrt_typeCallAny(New_haxe__ds___EnumValueMap__EnumValueTreeNode, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "haxe.ds.EnumValueMap":
		return &haxe__ds__EnumValueMap{}, true
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
		return &haxe__ds___EnumValueMap__EnumValueTreeNode{}, true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "CollectionFeatureKey":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(CollectionFeatureKey_Entry, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "Entry":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(CollectionFeatureKey_Entry, args)
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
	case *haxe__ds__EnumValueMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.EnumValueMap")}
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._EnumValueMap.EnumValueTreeNode")}
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
	case *CollectionFeatureKey:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("CollectionFeatureKey")}
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
	case "Main":
		return nil
	case "Reflect":
		return nil
	case "haxe.ds.EnumValueMap":
		return nil
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
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
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("main"))
	case "Reflect":
		return hxrt.NewArray(hxrt.StringFromLiteral("callMethod"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("compareMethods"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("deleteField"), hxrt.StringFromLiteral("field"), hxrt.StringFromLiteral("fields"), hxrt.StringFromLiteral("getProperty"), hxrt.StringFromLiteral("hasField"), hxrt.StringFromLiteral("isEnumValue"), hxrt.StringFromLiteral("isFunction"), hxrt.StringFromLiteral("isObject"), hxrt.StringFromLiteral("makeVarArgs"), hxrt.StringFromLiteral("setField"), hxrt.StringFromLiteral("setProperty"))
	case "haxe.ds.EnumValueMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("isEnumValue"), hxrt.StringFromLiteral("keysLoop"), hxrt.StringFromLiteral("valuesLoop"))
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
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
	case "Main":
		return hxrt.NewArray()
	case "Reflect":
		return hxrt.NewArray()
	case "haxe.ds.EnumValueMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("balance"), hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("compareArg"), hxrt.StringFromLiteral("compareArgs"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("merge"), hxrt.StringFromLiteral("minBinding"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("removeLoop"), hxrt.StringFromLiteral("removeMinBinding"), hxrt.StringFromLiteral("root"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("setLoop"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
		return hxrt.NewArray(hxrt.StringFromLiteral("getHeight"), hxrt.StringFromLiteral("height"), hxrt.StringFromLiteral("key"), hxrt.StringFromLiteral("left"), hxrt.StringFromLiteral("right"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("value"))
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
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Reflect":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.EnumValueMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._EnumValueMap.EnumValueTreeNode":
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
	case "CollectionFeatureKey":
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
	case *CollectionFeatureKey:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("Entry")
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
	case *CollectionFeatureKey:
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
	case "CollectionFeatureKey":
		return hxrt.NewArray(hxrt.StringFromLiteral("Entry"))
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
	case *CollectionFeatureKey:
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
	case "CollectionFeatureKey":
		return hxrt.NewArray()
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
	case *haxe__ds__EnumValueMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__ds__EnumValueMap:
		return hxrt__generated_field_lookup__haxe__ds__EnumValueMap(value, key)
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return hxrt__generated_field_lookup__haxe__ds___EnumValueMap__EnumValueTreeNode(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *haxe__ds__EnumValueMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe__ds__EnumValueMap:
		return hxrt__generated_field_has__haxe__ds__EnumValueMap(value, key)
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return hxrt__generated_field_has__haxe__ds___EnumValueMap__EnumValueTreeNode(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *haxe__ds__EnumValueMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe__ds__EnumValueMap:
		return hxrt__generated_field_set__haxe__ds__EnumValueMap(value, key, incoming)
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return hxrt__generated_field_set__haxe__ds___EnumValueMap__EnumValueTreeNode(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := object.(type) {
	case *haxe__ds__EnumValueMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *haxe__ds__EnumValueMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("root"))
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return hxrt.NewArray(hxrt.StringFromLiteral("height"), hxrt.StringFromLiteral("key"), hxrt.StringFromLiteral("left"), hxrt.StringFromLiteral("right"), hxrt.StringFromLiteral("value"))
	default:
		return nil
	}
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
	case *CollectionFeatureKey:
		return (enumValue != nil)
	default:
		return false
	}
}
