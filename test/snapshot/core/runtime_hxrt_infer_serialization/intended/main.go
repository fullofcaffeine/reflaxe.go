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

func main() {
	encoded := haxe__Serializer_run(New_SerializationSnapshotChild(1, hxrt.StringFromLiteral("ok")))
	decoded := func(hx_value_1 any) *SerializationSnapshotChild {
		if hx_value_1 == nil {
			var hx_zero_2 *SerializationSnapshotChild
			return hx_zero_2
		}
		return hx_value_1.(*SerializationSnapshotChild)
	}(haxe__Unserializer_run(encoded))
	var v any = any((decoded.__hx_this.readBase() == 1))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringEqualStringPtr(decoded.readChild(), hxrt.StringFromLiteral("ok")))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringEqualStringPtr(decoded.__hx_this.dispatch(), hxrt.StringFromLiteral("child:ok")))
	hxrt.Println(v_2)
}

type I_SerializationSnapshotBase interface {
	readBase() int
	label() *string
	dispatch() *string
}

type SerializationSnapshotBase struct {
	__hx_this I_SerializationSnapshotBase
	baseValue int
}

func New_SerializationSnapshotBase(baseValue int) *SerializationSnapshotBase {
	self := &SerializationSnapshotBase{}
	self.__hx_this = self
	self.baseValue = baseValue
	return self
}

func (self *SerializationSnapshotBase) readBase() int {
	return self.baseValue
}

func (self *SerializationSnapshotBase) label() *string {
	return hxrt.StringFromLiteral("base")
}

func (self *SerializationSnapshotBase) dispatch() *string {
	return self.__hx_this.label()
}

type I_SerializationSnapshotChild interface {
	readBase() int
	label() *string
	dispatch() *string
	readChild() *string
}

type SerializationSnapshotChild struct {
	*SerializationSnapshotBase
	__hx_this  I_SerializationSnapshotChild
	childValue *string
}

func New_SerializationSnapshotChild(baseValue int, childValue *string) *SerializationSnapshotChild {
	self := &SerializationSnapshotChild{}
	self.SerializationSnapshotBase = New_SerializationSnapshotBase(baseValue)
	self.SerializationSnapshotBase.__hx_this = self
	self.__hx_this = self
	self.childValue = childValue
	return self
}

func (self *SerializationSnapshotChild) readChild() *string {
	return self.childValue
}

func (self *SerializationSnapshotChild) label() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("child:"), self.childValue)
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := obj.(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *SerializationSnapshotBase:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *SerializationSnapshotChild:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Serializer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Unserializer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__IntMap:
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
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_method_field__Date(value, key)
	case *SerializationSnapshotBase:
		return hxrt__generated_method_field__SerializationSnapshotBase(value, key)
	case *SerializationSnapshotChild:
		return hxrt__generated_method_field__SerializationSnapshotChild(value, key)
	case *haxe__Serializer:
		return hxrt__generated_method_field__haxe__Serializer(value, key)
	case *haxe__Unserializer:
		return hxrt__generated_method_field__haxe__Unserializer(value, key)
	case *haxe__ds__IntMap:
		return hxrt__generated_method_field__haxe__ds__IntMap(value, key)
	case *haxe__ds__List:
		return hxrt__generated_method_field__haxe__ds__List(value, key)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_method_field__haxe__ds__ObjectMap(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_method_field__haxe__ds__StringMap(value, key)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_method_field__haxe__ds___List__GoListIterator(value, key)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_method_field__haxe__ds___List__GoListKeyValueIterator(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_method_field__haxe__io__Bytes(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_method_field__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Date(value *Date, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "getDate":
		return value.getDate
	case "getDay":
		return value.getDay
	case "getFullYear":
		return value.getFullYear
	case "getHours":
		return value.getHours
	case "getMinutes":
		return value.getMinutes
	case "getMonth":
		return value.getMonth
	case "getSeconds":
		return value.getSeconds
	case "getTime":
		return value.getTime
	case "getTimezoneOffset":
		return value.getTimezoneOffset
	case "getUTCDate":
		return value.getUTCDate
	case "getUTCDay":
		return value.getUTCDay
	case "getUTCFullYear":
		return value.getUTCFullYear
	case "getUTCHours":
		return value.getUTCHours
	case "getUTCMinutes":
		return value.getUTCMinutes
	case "getUTCMonth":
		return value.getUTCMonth
	case "getUTCSeconds":
		return value.getUTCSeconds
	case "localParts":
		return value.localParts
	case "toString":
		return value.toString
	case "utcParts":
		return value.utcParts
	}
	return nil
}

func hxrt__generated_method_field__SerializationSnapshotBase(value *SerializationSnapshotBase, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "dispatch":
		return value.dispatch
	case "label":
		return value.label
	case "readBase":
		return value.readBase
	}
	return nil
}

func hxrt__generated_method_field__SerializationSnapshotChild(value *SerializationSnapshotChild, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "label":
		return value.label
	case "readChild":
		return value.readChild
	}
	if value.SerializationSnapshotBase == nil {
		return nil
	}
	return hxrt__generated_method_field__SerializationSnapshotBase(value.SerializationSnapshotBase, key)
}

func hxrt__generated_method_field__haxe__Serializer(value *haxe__Serializer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "flushNulls":
		return value.flushNulls
	case "serialize":
		return value.serialize
	case "serializeArray":
		return value.serializeArray
	case "serializeBytes":
		return value.serializeBytes
	case "serializeClass":
		return value.serializeClass
	case "serializeEnum":
		return value.serializeEnum
	case "serializeException":
		return value.serializeException
	case "serializeFields":
		return value.serializeFields
	case "serializeRef":
		return value.serializeRef
	case "serializeString":
		return value.serializeString
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__Unserializer(value *haxe__Unserializer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "base64Value":
		return value.base64Value
	case "decodeBytes":
		return value.decodeBytes
	case "get":
		return value.get
	case "getResolver":
		return value.getResolver
	case "isLegacyDate":
		return value.isLegacyDate
	case "readDigits":
		return value.readDigits
	case "readFloat":
		return value.readFloat
	case "resolveClass":
		return value.resolveClass
	case "resolveEnum":
		return value.resolveEnum
	case "setResolver":
		return value.setResolver
	case "unserialize":
		return value.unserialize
	case "unserializeEnum":
		return value.unserializeEnum
	case "unserializeObject":
		return value.unserializeObject
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__IntMap(value *haxe__ds__IntMap, key string) any {
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

func hxrt__generated_method_field__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_nativeView":
		return value.__hx_nativeView
	case "blit":
		return value.blit
	case "compare":
		return value.compare
	case "fill":
		return value.fill
	case "get":
		return value.get
	case "getData":
		return value.getData
	case "getDouble":
		return value.getDouble
	case "getFloat":
		return value.getFloat
	case "getInt32":
		return value.getInt32
	case "getInt64":
		return value.getInt64
	case "getString":
		return value.getString
	case "getUInt16":
		return value.getUInt16
	case "readString":
		return value.readString
	case "set":
		return value.set
	case "setDouble":
		return value.setDouble
	case "setFloat":
		return value.setFloat
	case "setInt32":
		return value.setInt32
	case "setInt64":
		return value.setInt64
	case "setUInt16":
		return value.setUInt16
	case "sub":
		return value.sub
	case "toHex":
		return value.toHex
	case "toString":
		return value.toString
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

func hxrt__generated_method_field__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
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

func hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
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

func hxrt_typeCreateEmpty__Date() *Date {
	instance := &Date{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__SerializationSnapshotBase() *SerializationSnapshotBase {
	instance := &SerializationSnapshotBase{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__SerializationSnapshotChild() *SerializationSnapshotChild {
	instance := &SerializationSnapshotChild{}
	instance.SerializationSnapshotBase = &SerializationSnapshotBase{}
	instance.SerializationSnapshotBase.__hx_this = instance
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__StringBuf() *StringBuf {
	instance := &StringBuf{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__Serializer() *haxe__Serializer {
	instance := &haxe__Serializer{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__Unserializer() *haxe__Unserializer {
	instance := &haxe__Unserializer{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe___Int64_____Int64() *haxe___Int64_____Int64 {
	instance := &haxe___Int64_____Int64{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds__IntMap() *haxe__ds__IntMap {
	instance := &haxe__ds__IntMap{}
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

func hxrt_typeCreateEmpty__haxe__io__Bytes() *haxe__io__Bytes {
	instance := &haxe__io__Bytes{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__MapKeyValueIterator() *haxe__iterators__MapKeyValueIterator {
	instance := &haxe__iterators__MapKeyValueIterator{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__StringIterator() *haxe__iterators__StringIterator {
	instance := &haxe__iterators__StringIterator{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__StringKeyValueIterator() *haxe__iterators__StringKeyValueIterator {
	instance := &haxe__iterators__StringKeyValueIterator{}
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
	case "Date":
		return hxrt_typeCallAny(New_Date, args)
	case "Main":
		return nil, false
	case "Math":
		return nil, false
	case "Reflect":
		return nil, false
	case "SerializationSnapshotBase":
		return hxrt_typeCallAny(New_SerializationSnapshotBase, args)
	case "SerializationSnapshotChild":
		return hxrt_typeCallAny(New_SerializationSnapshotChild, args)
	case "StringBuf":
		return hxrt_typeCallAny(New_StringBuf, args)
	case "StringTools":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe.Serializer":
		return hxrt_typeCallAny(New_haxe__Serializer, args)
	case "haxe.Unserializer":
		return hxrt_typeCallAny(New_haxe__Unserializer, args)
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.ds.IntMap":
		return hxrt_typeCallAny(New_haxe__ds__IntMap, args)
	case "haxe.ds.List":
		return hxrt_typeCallAny(New_haxe__ds__List, args)
	case "haxe.ds.ObjectMap":
		return hxrt_typeCallAny(New_haxe__ds__ObjectMap, args)
	case "haxe.ds.StringMap":
		return hxrt_typeCallAny(New_haxe__ds__StringMap, args)
	case "haxe.ds._List.GoListIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListIterator, args)
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListKeyValueIterator, args)
	case "haxe.io.Bytes":
		return hxrt_typeCallAny(New_haxe__io__Bytes, args)
	case "haxe.io.FPHelper":
		return nil, false
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__MapKeyValueIterator, args)
	case "haxe.iterators.StringIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringIterator, args)
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringKeyValueIterator, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "Date":
		return hxrt_typeCreateEmpty__Date(), true
	case "SerializationSnapshotBase":
		return hxrt_typeCreateEmpty__SerializationSnapshotBase(), true
	case "SerializationSnapshotChild":
		return hxrt_typeCreateEmpty__SerializationSnapshotChild(), true
	case "StringBuf":
		return hxrt_typeCreateEmpty__StringBuf(), true
	case "haxe.Serializer":
		return hxrt_typeCreateEmpty__haxe__Serializer(), true
	case "haxe.Unserializer":
		return hxrt_typeCreateEmpty__haxe__Unserializer(), true
	case "haxe._Int64.___Int64":
		return hxrt_typeCreateEmpty__haxe___Int64_____Int64(), true
	case "haxe.ds.IntMap":
		return hxrt_typeCreateEmpty__haxe__ds__IntMap(), true
	case "haxe.ds.List":
		return hxrt_typeCreateEmpty__haxe__ds__List(), true
	case "haxe.ds.ObjectMap":
		return hxrt_typeCreateEmpty__haxe__ds__ObjectMap(), true
	case "haxe.ds.StringMap":
		return hxrt_typeCreateEmpty__haxe__ds__StringMap(), true
	case "haxe.ds._List.GoListIterator":
		return hxrt_typeCreateEmpty__haxe__ds___List__GoListIterator(), true
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt_typeCreateEmpty__haxe__ds___List__GoListKeyValueIterator(), true
	case "haxe.io.Bytes":
		return hxrt_typeCreateEmpty__haxe__io__Bytes(), true
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__MapKeyValueIterator(), true
	case "haxe.iterators.StringIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__StringIterator(), true
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__StringKeyValueIterator(), true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
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
	case "haxe.io.Encoding":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Encoding_UTF8, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Encoding_RawNative, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "UTF8":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Encoding_UTF8, true
		case "RawNative":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Encoding_RawNative, true
		default:
			return nil, false
		}
	case "haxe.io.Error":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_Blocked, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_Overflow, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_OutsideBounds, true
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__io__Error_Custom, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "Blocked":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_Blocked, true
		case "Overflow":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_Overflow, true
		case "OutsideBounds":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_OutsideBounds, true
		case "Custom":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__io__Error_Custom, args)
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
	case *Date:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Date")}
	case *SerializationSnapshotBase:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("SerializationSnapshotBase")}
	case *SerializationSnapshotChild:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("SerializationSnapshotChild")}
	case *StringBuf:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("StringBuf")}
	case *haxe__Serializer:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Serializer")}
	case *haxe__Unserializer:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Unserializer")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe__ds__IntMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.IntMap")}
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
	case *haxe__io__Bytes:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Bytes")}
	case *haxe__iterators__MapKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.MapKeyValueIterator")}
	case *haxe__iterators__StringIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.StringIterator")}
	case *haxe__iterators__StringKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.StringKeyValueIterator")}
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
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	case *haxe__io__Encoding:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.io.Encoding")}
	case *haxe__io__Error:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.io.Error")}
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
	case "Date":
		return nil
	case "Main":
		return nil
	case "Math":
		return nil
	case "Reflect":
		return nil
	case "SerializationSnapshotBase":
		return nil
	case "SerializationSnapshotChild":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("SerializationSnapshotBase")}
	case "StringBuf":
		return nil
	case "StringTools":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe.Serializer":
		return nil
	case "haxe.Unserializer":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.ds.IntMap":
		return nil
	case "haxe.ds.List":
		return nil
	case "haxe.ds.ObjectMap":
		return nil
	case "haxe.ds.StringMap":
		return nil
	case "haxe.ds._List.GoListIterator":
		return nil
	case "haxe.ds._List.GoListKeyValueIterator":
		return nil
	case "haxe.io.Bytes":
		return nil
	case "haxe.io.FPHelper":
		return nil
	case "haxe.iterators.MapKeyValueIterator":
		return nil
	case "haxe.iterators.StringIterator":
		return nil
	case "haxe.iterators.StringKeyValueIterator":
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
	case "Date":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromMilliseconds"), hxrt.StringFromLiteral("fromString"), hxrt.StringFromLiteral("fromTime"), hxrt.StringFromLiteral("now"))
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("main"))
	case "Math":
		return hxrt.NewArray(hxrt.StringFromLiteral("NEGATIVE_INFINITY"), hxrt.StringFromLiteral("NaN"), hxrt.StringFromLiteral("PI"), hxrt.StringFromLiteral("POSITIVE_INFINITY"), hxrt.StringFromLiteral("abs"), hxrt.StringFromLiteral("acos"), hxrt.StringFromLiteral("asin"), hxrt.StringFromLiteral("atan"), hxrt.StringFromLiteral("atan2"), hxrt.StringFromLiteral("ceil"), hxrt.StringFromLiteral("cos"), hxrt.StringFromLiteral("exp"), hxrt.StringFromLiteral("fceil"), hxrt.StringFromLiteral("ffloor"), hxrt.StringFromLiteral("floor"), hxrt.StringFromLiteral("fround"), hxrt.StringFromLiteral("isFinite"), hxrt.StringFromLiteral("isNaN"), hxrt.StringFromLiteral("log"), hxrt.StringFromLiteral("max"), hxrt.StringFromLiteral("min"), hxrt.StringFromLiteral("pow"), hxrt.StringFromLiteral("random"), hxrt.StringFromLiteral("round"), hxrt.StringFromLiteral("sin"), hxrt.StringFromLiteral("sqrt"), hxrt.StringFromLiteral("tan"))
	case "Reflect":
		return hxrt.NewArray(hxrt.StringFromLiteral("callMethod"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("compareMethods"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("deleteField"), hxrt.StringFromLiteral("field"), hxrt.StringFromLiteral("fields"), hxrt.StringFromLiteral("getProperty"), hxrt.StringFromLiteral("hasField"), hxrt.StringFromLiteral("isEnumValue"), hxrt.StringFromLiteral("isFunction"), hxrt.StringFromLiteral("isObject"), hxrt.StringFromLiteral("makeVarArgs"), hxrt.StringFromLiteral("setField"), hxrt.StringFromLiteral("setProperty"))
	case "SerializationSnapshotBase":
		return hxrt.NewArray()
	case "SerializationSnapshotChild":
		return hxrt.NewArray()
	case "StringBuf":
		return hxrt.NewArray()
	case "StringTools":
		return hxrt.NewArray(hxrt.StringFromLiteral("MAX_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("contains"), hxrt.StringFromLiteral("containsImpl"), hxrt.StringFromLiteral("endsWith"), hxrt.StringFromLiteral("endsWithImpl"), hxrt.StringFromLiteral("fastCodeAt"), hxrt.StringFromLiteral("hex"), hxrt.StringFromLiteral("hexDigitValue"), hxrt.StringFromLiteral("htmlEscape"), hxrt.StringFromLiteral("htmlUnescape"), hxrt.StringFromLiteral("isEof"), hxrt.StringFromLiteral("isSpace"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("lpad"), hxrt.StringFromLiteral("ltrim"), hxrt.StringFromLiteral("replace"), hxrt.StringFromLiteral("rpad"), hxrt.StringFromLiteral("rtrim"), hxrt.StringFromLiteral("startsWith"), hxrt.StringFromLiteral("startsWithImpl"), hxrt.StringFromLiteral("trim"), hxrt.StringFromLiteral("unsafeCodeAt"), hxrt.StringFromLiteral("urlDecode"), hxrt.StringFromLiteral("urlEncode"), hxrt.StringFromLiteral("utf16CodePointAt"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Serializer":
		return hxrt.NewArray(hxrt.StringFromLiteral("BASE64"), hxrt.StringFromLiteral("USE_CACHE"), hxrt.StringFromLiteral("USE_ENUM_INDEX"), hxrt.StringFromLiteral("run"))
	case "haxe.Unserializer":
		return hxrt.NewArray(hxrt.StringFromLiteral("BASE64"), hxrt.StringFromLiteral("DEFAULT_RESOLVER"), hxrt.StringFromLiteral("NULL_RESOLVER"), hxrt.StringFromLiteral("run"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.ds.IntMap":
		return hxrt.NewArray()
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("sameValue"))
	case "haxe.ds.ObjectMap":
		return hxrt.NewArray()
	case "haxe.ds.StringMap":
		return hxrt.NewArray()
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray()
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_fromNativeView"), hxrt.StringFromLiteral("alloc"), hxrt.StringFromLiteral("fastGet"), hxrt.StringFromLiteral("ofData"), hxrt.StringFromLiteral("ofHex"), hxrt.StringFromLiteral("ofString"), hxrt.StringFromLiteral("rawNativeUsesUtf16LE"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray(hxrt.StringFromLiteral("doubleToI64"), hxrt.StringFromLiteral("floatToI32"), hxrt.StringFromLiteral("i32ToFloat"), hxrt.StringFromLiteral("i64ToDouble"))
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringKeyValueIterator":
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
	case "Date":
		return hxrt.NewArray(hxrt.StringFromLiteral("getDate"), hxrt.StringFromLiteral("getDay"), hxrt.StringFromLiteral("getFullYear"), hxrt.StringFromLiteral("getHours"), hxrt.StringFromLiteral("getMinutes"), hxrt.StringFromLiteral("getMonth"), hxrt.StringFromLiteral("getSeconds"), hxrt.StringFromLiteral("getTime"), hxrt.StringFromLiteral("getTimezoneOffset"), hxrt.StringFromLiteral("getUTCDate"), hxrt.StringFromLiteral("getUTCDay"), hxrt.StringFromLiteral("getUTCFullYear"), hxrt.StringFromLiteral("getUTCHours"), hxrt.StringFromLiteral("getUTCMinutes"), hxrt.StringFromLiteral("getUTCMonth"), hxrt.StringFromLiteral("getUTCSeconds"), hxrt.StringFromLiteral("localParts"), hxrt.StringFromLiteral("ms"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("utcParts"))
	case "Main":
		return hxrt.NewArray()
	case "Math":
		return hxrt.NewArray()
	case "Reflect":
		return hxrt.NewArray()
	case "SerializationSnapshotBase":
		return hxrt.NewArray(hxrt.StringFromLiteral("baseValue"), hxrt.StringFromLiteral("dispatch"), hxrt.StringFromLiteral("label"), hxrt.StringFromLiteral("readBase"))
	case "SerializationSnapshotChild":
		return hxrt.NewArray(hxrt.StringFromLiteral("baseValue"), hxrt.StringFromLiteral("childValue"), hxrt.StringFromLiteral("dispatch"), hxrt.StringFromLiteral("label"), hxrt.StringFromLiteral("readBase"), hxrt.StringFromLiteral("readChild"))
	case "StringBuf":
		return hxrt.NewArray(hxrt.StringFromLiteral("b"))
	case "StringTools":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Serializer":
		return hxrt.NewArray(hxrt.StringFromLiteral("buf"), hxrt.StringFromLiteral("cache"), hxrt.StringFromLiteral("flushNulls"), hxrt.StringFromLiteral("scount"), hxrt.StringFromLiteral("serialize"), hxrt.StringFromLiteral("serializeArray"), hxrt.StringFromLiteral("serializeBytes"), hxrt.StringFromLiteral("serializeClass"), hxrt.StringFromLiteral("serializeEnum"), hxrt.StringFromLiteral("serializeException"), hxrt.StringFromLiteral("serializeFields"), hxrt.StringFromLiteral("serializeRef"), hxrt.StringFromLiteral("serializeString"), hxrt.StringFromLiteral("shash"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("useCache"), hxrt.StringFromLiteral("useEnumIndex"))
	case "haxe.Unserializer":
		return hxrt.NewArray(hxrt.StringFromLiteral("base64Value"), hxrt.StringFromLiteral("buf"), hxrt.StringFromLiteral("cache"), hxrt.StringFromLiteral("decodeBytes"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getResolver"), hxrt.StringFromLiteral("isLegacyDate"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("pos"), hxrt.StringFromLiteral("readDigits"), hxrt.StringFromLiteral("readFloat"), hxrt.StringFromLiteral("resolveClass"), hxrt.StringFromLiteral("resolveEnum"), hxrt.StringFromLiteral("resolver"), hxrt.StringFromLiteral("scache"), hxrt.StringFromLiteral("setResolver"), hxrt.StringFromLiteral("unserialize"), hxrt.StringFromLiteral("unserializeEnum"), hxrt.StringFromLiteral("unserializeObject"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.ds.IntMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("filter"), hxrt.StringFromLiteral("first"), hxrt.StringFromLiteral("isEmpty"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("join"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("last"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds.ObjectMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds.StringMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_nativeView"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("blit"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("fill"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getData"), hxrt.StringFromLiteral("getDouble"), hxrt.StringFromLiteral("getFloat"), hxrt.StringFromLiteral("getInt32"), hxrt.StringFromLiteral("getInt64"), hxrt.StringFromLiteral("getString"), hxrt.StringFromLiteral("getUInt16"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setDouble"), hxrt.StringFromLiteral("setFloat"), hxrt.StringFromLiteral("setInt32"), hxrt.StringFromLiteral("setInt64"), hxrt.StringFromLiteral("setUInt16"), hxrt.StringFromLiteral("sub"), hxrt.StringFromLiteral("toHex"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray()
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("next"))
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
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
	case "Date":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Math":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Reflect":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "SerializationSnapshotBase":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "SerializationSnapshotChild":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringBuf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringTools":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Serializer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Unserializer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.IntMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.List":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.ObjectMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.StringMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Bytes":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.FPHelper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.MapKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringKeyValueIterator":
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
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Encoding":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Error":
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
	case *haxe__io__Encoding:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("UTF8")
		case 1:
			return hxrt.StringFromLiteral("RawNative")
		default:
			return nil
		}
	case *haxe__io__Error:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("Blocked")
		case 1:
			return hxrt.StringFromLiteral("Overflow")
		case 2:
			return hxrt.StringFromLiteral("OutsideBounds")
		case 3:
			return hxrt.StringFromLiteral("Custom")
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
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__io__Encoding:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__io__Error:
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
	case "ValueType":
		return hxrt.NewArray(hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown"))
	case "haxe.io.Encoding":
		return hxrt.NewArray(hxrt.StringFromLiteral("UTF8"), hxrt.StringFromLiteral("RawNative"))
	case "haxe.io.Error":
		return hxrt.NewArray(hxrt.StringFromLiteral("Blocked"), hxrt.StringFromLiteral("Overflow"), hxrt.StringFromLiteral("OutsideBounds"), hxrt.StringFromLiteral("Custom"))
	default:
		return hxrt.NewArray()
	}
}

func Type_enumParameters(e any) *hxrt.Array {
	if hxrt.AnyEqualsNull(e) {
		return hxrt.NewArray()
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__io__Encoding:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__io__Error:
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
	case "ValueType":
		return hxrt.NewArray(ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown)
	case "haxe.io.Encoding":
		return hxrt.NewArray(haxe__io__Encoding_UTF8, haxe__io__Encoding_RawNative)
	case "haxe.io.Error":
		return hxrt.NewArray(haxe__io__Error_Blocked, haxe__io__Error_Overflow, haxe__io__Error_OutsideBounds)
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
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *SerializationSnapshotBase:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *SerializationSnapshotChild:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *StringBuf:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Serializer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Unserializer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__IntMap:
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
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_field_lookup__Date(value, key)
	case *SerializationSnapshotBase:
		return hxrt__generated_field_lookup__SerializationSnapshotBase(value, key)
	case *SerializationSnapshotChild:
		return hxrt__generated_field_lookup__SerializationSnapshotChild(value, key)
	case *StringBuf:
		return hxrt__generated_field_lookup__StringBuf(value, key)
	case *haxe__Serializer:
		return hxrt__generated_field_lookup__haxe__Serializer(value, key)
	case *haxe__Unserializer:
		return hxrt__generated_field_lookup__haxe__Unserializer(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe__ds__IntMap:
		return hxrt__generated_field_lookup__haxe__ds__IntMap(value, key)
	case *haxe__ds__List:
		return hxrt__generated_field_lookup__haxe__ds__List(value, key)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_field_lookup__haxe__ds__ObjectMap(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_lookup__haxe__ds__StringMap(value, key)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_field_lookup__haxe__ds___List__GoListIterator(value, key)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__ds___List__GoListKeyValueIterator(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_lookup__haxe__io__Bytes(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *SerializationSnapshotBase:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *SerializationSnapshotChild:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *StringBuf:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Serializer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Unserializer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__IntMap:
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
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_field_has__Date(value, key)
	case *SerializationSnapshotBase:
		return hxrt__generated_field_has__SerializationSnapshotBase(value, key)
	case *SerializationSnapshotChild:
		return hxrt__generated_field_has__SerializationSnapshotChild(value, key)
	case *StringBuf:
		return hxrt__generated_field_has__StringBuf(value, key)
	case *haxe__Serializer:
		return hxrt__generated_field_has__haxe__Serializer(value, key)
	case *haxe__Unserializer:
		return hxrt__generated_field_has__haxe__Unserializer(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe__ds__IntMap:
		return hxrt__generated_field_has__haxe__ds__IntMap(value, key)
	case *haxe__ds__List:
		return hxrt__generated_field_has__haxe__ds__List(value, key)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_field_has__haxe__ds__ObjectMap(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_has__haxe__ds__StringMap(value, key)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_field_has__haxe__ds___List__GoListIterator(value, key)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_field_has__haxe__ds___List__GoListKeyValueIterator(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_has__haxe__io__Bytes(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_has__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *SerializationSnapshotBase:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *SerializationSnapshotChild:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *StringBuf:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Serializer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Unserializer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__IntMap:
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
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_field_set__Date(value, key, incoming)
	case *SerializationSnapshotBase:
		return hxrt__generated_field_set__SerializationSnapshotBase(value, key, incoming)
	case *SerializationSnapshotChild:
		return hxrt__generated_field_set__SerializationSnapshotChild(value, key, incoming)
	case *StringBuf:
		return hxrt__generated_field_set__StringBuf(value, key, incoming)
	case *haxe__Serializer:
		return hxrt__generated_field_set__haxe__Serializer(value, key, incoming)
	case *haxe__Unserializer:
		return hxrt__generated_field_set__haxe__Unserializer(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe__ds__IntMap:
		return hxrt__generated_field_set__haxe__ds__IntMap(value, key, incoming)
	case *haxe__ds__List:
		return hxrt__generated_field_set__haxe__ds__List(value, key, incoming)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_field_set__haxe__ds__ObjectMap(value, key, incoming)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_set__haxe__ds__StringMap(value, key, incoming)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_field_set__haxe__ds___List__GoListIterator(value, key, incoming)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_field_set__haxe__ds___List__GoListKeyValueIterator(value, key, incoming)
	case *haxe__io__Bytes:
		return hxrt__generated_field_set__haxe__io__Bytes(value, key, incoming)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value, key, incoming)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_set__haxe__iterators__StringIterator(value, key, incoming)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := object.(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *SerializationSnapshotBase:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *SerializationSnapshotChild:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *StringBuf:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Serializer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Unserializer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__IntMap:
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
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *Date:
		return hxrt.NewArray(hxrt.StringFromLiteral("ms"))
	case *SerializationSnapshotBase:
		return hxrt.NewArray(hxrt.StringFromLiteral("baseValue"))
	case *SerializationSnapshotChild:
		return hxrt.NewArray(hxrt.StringFromLiteral("baseValue"), hxrt.StringFromLiteral("childValue"))
	case *StringBuf:
		return hxrt.NewArray(hxrt.StringFromLiteral("b"))
	case *haxe__Serializer:
		return hxrt.NewArray(hxrt.StringFromLiteral("buf"), hxrt.StringFromLiteral("cache"), hxrt.StringFromLiteral("shash"), hxrt.StringFromLiteral("scount"), hxrt.StringFromLiteral("useCache"), hxrt.StringFromLiteral("useEnumIndex"))
	case *haxe__Unserializer:
		return hxrt.NewArray(hxrt.StringFromLiteral("buf"), hxrt.StringFromLiteral("pos"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("cache"), hxrt.StringFromLiteral("scache"), hxrt.StringFromLiteral("resolver"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__IntMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__ds__List:
		return hxrt.NewArray(hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("length"))
	case *haxe__ds__ObjectMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__ds__StringMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__ds___List__GoListIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("index"))
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("index"))
	case *haxe__io__Bytes:
		return hxrt.NewArray(hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("__hx_dataExposed"))
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("keys"))
	case *haxe__iterators__StringIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Date(value *Date, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "ms":
		return value.ms
	}
	return nil
}

func hxrt__generated_field_has__Date(value *Date, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "ms":
		return true
	}
	return false
}

func hxrt__generated_field_set__Date(value *Date, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "ms":
		if incoming == nil {
			var zero float64
			value.ms = zero
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.ms = typed
			return true
		case int:
			value.ms = float64(typed)
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__SerializationSnapshotBase(value *SerializationSnapshotBase, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "baseValue":
		return value.baseValue
	}
	return nil
}

func hxrt__generated_field_has__SerializationSnapshotBase(value *SerializationSnapshotBase, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "baseValue":
		return true
	}
	return false
}

func hxrt__generated_field_set__SerializationSnapshotBase(value *SerializationSnapshotBase, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "baseValue":
		if incoming == nil {
			var zero int
			value.baseValue = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.baseValue = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__SerializationSnapshotChild(value *SerializationSnapshotChild, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "childValue":
		return value.childValue
	}
	if value.SerializationSnapshotBase == nil {
		return nil
	}
	return hxrt__generated_field_lookup__SerializationSnapshotBase(value.SerializationSnapshotBase, key)
}

func hxrt__generated_field_has__SerializationSnapshotChild(value *SerializationSnapshotChild, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "childValue":
		return true
	}
	if value.SerializationSnapshotBase == nil {
		return false
	}
	return hxrt__generated_field_has__SerializationSnapshotBase(value.SerializationSnapshotBase, key)
}

func hxrt__generated_field_set__SerializationSnapshotChild(value *SerializationSnapshotChild, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "childValue":
		if incoming == nil {
			var zero *string
			value.childValue = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.childValue = typed
			return true
		default:
			return false
		}
	}
	if value.SerializationSnapshotBase == nil {
		return false
	}
	return hxrt__generated_field_set__SerializationSnapshotBase(value.SerializationSnapshotBase, key, incoming)
}

func hxrt__generated_field_lookup__StringBuf(value *StringBuf, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "b":
		return value.b
	}
	return nil
}

func hxrt__generated_field_has__StringBuf(value *StringBuf, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		return true
	}
	return false
}

func hxrt__generated_field_set__StringBuf(value *StringBuf, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		if incoming == nil {
			var zero *string
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.b = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__Serializer(value *haxe__Serializer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "buf":
		return value.buf
	case "cache":
		return value.cache
	case "scount":
		return value.scount
	case "shash":
		return value.shash
	case "useCache":
		return value.useCache
	case "useEnumIndex":
		return value.useEnumIndex
	}
	return nil
}

func hxrt__generated_field_has__haxe__Serializer(value *haxe__Serializer, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "buf":
		return true
	case "cache":
		return true
	case "scount":
		return true
	case "shash":
		return true
	case "useCache":
		return true
	case "useEnumIndex":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Serializer(value *haxe__Serializer, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "buf":
		if incoming == nil {
			var zero *StringBuf
			value.buf = zero
			return true
		}
		switch typed := incoming.(type) {
		case *StringBuf:
			value.buf = typed
			return true
		default:
			return false
		}
	case "cache":
		if incoming == nil {
			var zero *hxrt.Array
			value.cache = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.cache = typed
			return true
		default:
			return false
		}
	case "scount":
		if incoming == nil {
			var zero int
			value.scount = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.scount = typed
			return true
		default:
			return false
		}
	case "shash":
		if incoming == nil {
			var zero *haxe__ds__StringMap
			value.shash = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds__StringMap:
			value.shash = typed
			return true
		default:
			return false
		}
	case "useCache":
		if incoming == nil {
			var zero bool
			value.useCache = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.useCache = typed
			return true
		default:
			return false
		}
	case "useEnumIndex":
		if incoming == nil {
			var zero bool
			value.useEnumIndex = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.useEnumIndex = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__Unserializer(value *haxe__Unserializer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "buf":
		return value.buf
	case "cache":
		return value.cache
	case "length":
		return value.length
	case "pos":
		return value.pos
	case "resolver":
		return value.resolver
	case "scache":
		return value.scache
	}
	return nil
}

func hxrt__generated_field_has__haxe__Unserializer(value *haxe__Unserializer, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "buf":
		return true
	case "cache":
		return true
	case "length":
		return true
	case "pos":
		return true
	case "resolver":
		return true
	case "scache":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Unserializer(value *haxe__Unserializer, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "buf":
		if incoming == nil {
			var zero *string
			value.buf = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.buf = typed
			return true
		default:
			return false
		}
	case "cache":
		if incoming == nil {
			var zero *hxrt.Array
			value.cache = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.cache = typed
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
	case "pos":
		if incoming == nil {
			var zero int
			value.pos = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.pos = typed
			return true
		default:
			return false
		}
	case "resolver":
		if incoming == nil {
			var zero any
			value.resolver = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.resolver = typed
			return true
		default:
			return false
		}
	case "scache":
		if incoming == nil {
			var zero *hxrt.Array
			value.scache = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.scache = typed
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

func hxrt__generated_field_lookup__haxe__ds__IntMap(value *haxe__ds__IntMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "h":
		return value.h
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__IntMap(value *haxe__ds__IntMap, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__IntMap(value *haxe__ds__IntMap, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		if incoming == nil {
			var zero *hxrt.IntMapCell
			value.h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.IntMapCell:
			value.h = typed
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

func hxrt__generated_field_lookup__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_dataExposed":
		return value.__hx_dataExposed
	case "__hx_raw":
		return value.__hx_raw
	case "__hx_rawValid":
		return value.__hx_rawValid
	case "b":
		return value.b
	case "length":
		return value.length
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Bytes(value *haxe__io__Bytes, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__hx_dataExposed":
		return true
	case "__hx_raw":
		return true
	case "__hx_rawValid":
		return true
	case "b":
		return true
	case "length":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Bytes(value *haxe__io__Bytes, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__hx_dataExposed":
		if incoming == nil {
			var zero bool
			value.__hx_dataExposed = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.__hx_dataExposed = typed
			return true
		default:
			return false
		}
	case "__hx_raw":
		if incoming == nil {
			var zero *hxrt.ByteView
			value.__hx_raw = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.ByteView:
			value.__hx_raw = typed
			return true
		default:
			return false
		}
	case "__hx_rawValid":
		if incoming == nil {
			var zero bool
			value.__hx_rawValid = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.__hx_rawValid = typed
			return true
		default:
			return false
		}
	case "b":
		if incoming == nil {
			var zero []int
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case []int:
			value.b = typed
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

func hxrt__generated_field_lookup__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "offset":
		return value.offset
	case "s":
		return value.s
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		if incoming == nil {
			var zero int
			value.offset = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.offset = typed
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "offset":
		return value.offset
	case "s":
		return value.s
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		if incoming == nil {
			var zero int
			value.offset = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.offset = typed
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
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
	case *haxe__io__Encoding:
		return (enumValue != nil)
	case *haxe__io__Error:
		return (enumValue != nil)
	default:
		return false
	}
}
