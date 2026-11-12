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
	decoded := hx_recover_class_0(haxe__Unserializer_run(encoded))
	var v any = any((decoded.readBase() == 1))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringEqualStringPtr(decoded.readChild(), hxrt.StringFromLiteral("ok")))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringEqualStringPtr(decoded.dispatch(), hxrt.StringFromLiteral("child:ok")))
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

func (self *SerializationSnapshotBase) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SerializationSnapshotBase(baseValue int) *SerializationSnapshotBase {
	return hx_init_SerializationSnapshotBase(hx_alloc_SerializationSnapshotBase(nil), baseValue)
}

func hx_alloc_SerializationSnapshotBase(owner any) *SerializationSnapshotBase {
	self := &SerializationSnapshotBase{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SerializationSnapshotBase)
	}
	return self
}

func hx_init_SerializationSnapshotBase(self *SerializationSnapshotBase, baseValue int) *SerializationSnapshotBase {
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

func (self *SerializationSnapshotChild) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SerializationSnapshotChild(baseValue int, childValue *string) *SerializationSnapshotChild {
	return hx_init_SerializationSnapshotChild(hx_alloc_SerializationSnapshotChild(nil), baseValue, childValue)
}

func hx_alloc_SerializationSnapshotChild(owner any) *SerializationSnapshotChild {
	self := &SerializationSnapshotChild{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SerializationSnapshotChild)
	}
	self.SerializationSnapshotBase = hx_alloc_SerializationSnapshotBase(self.__hx_this)
	return self
}

func hx_init_SerializationSnapshotChild(self *SerializationSnapshotChild, baseValue int, childValue *string) *SerializationSnapshotChild {
	hx_init_SerializationSnapshotBase(self.SerializationSnapshotBase, baseValue)
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
	switch value := hxrt.ObjectSource(obj).(type) {
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
	case *haxe__Exception:
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
	case *haxe__ValueException:
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
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__Serializer:
		return hxrt__generated_method_field__haxe__Serializer(value, key)
	case *haxe__Unserializer:
		return hxrt__generated_method_field__haxe__Unserializer(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
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
		return func() *hxrt.Function[func() int] {
			hx_receiver_398 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_398.__hx_this.getDate()
			}, func() any {
				return hx_receiver_398.__hx_this
			}, "getDate")
		}()
	case "getDay":
		return func() *hxrt.Function[func() int] {
			hx_receiver_399 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_399.__hx_this.getDay()
			}, func() any {
				return hx_receiver_399.__hx_this
			}, "getDay")
		}()
	case "getFullYear":
		return func() *hxrt.Function[func() int] {
			hx_receiver_400 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_400.__hx_this.getFullYear()
			}, func() any {
				return hx_receiver_400.__hx_this
			}, "getFullYear")
		}()
	case "getHours":
		return func() *hxrt.Function[func() int] {
			hx_receiver_401 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_401.__hx_this.getHours()
			}, func() any {
				return hx_receiver_401.__hx_this
			}, "getHours")
		}()
	case "getMinutes":
		return func() *hxrt.Function[func() int] {
			hx_receiver_402 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_402.__hx_this.getMinutes()
			}, func() any {
				return hx_receiver_402.__hx_this
			}, "getMinutes")
		}()
	case "getMonth":
		return func() *hxrt.Function[func() int] {
			hx_receiver_403 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_403.__hx_this.getMonth()
			}, func() any {
				return hx_receiver_403.__hx_this
			}, "getMonth")
		}()
	case "getSeconds":
		return func() *hxrt.Function[func() int] {
			hx_receiver_404 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_404.__hx_this.getSeconds()
			}, func() any {
				return hx_receiver_404.__hx_this
			}, "getSeconds")
		}()
	case "getTime":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_405 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_405.__hx_this.getTime()
			}, func() any {
				return hx_receiver_405.__hx_this
			}, "getTime")
		}()
	case "getTimezoneOffset":
		return func() *hxrt.Function[func() int] {
			hx_receiver_406 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_406.__hx_this.getTimezoneOffset()
			}, func() any {
				return hx_receiver_406.__hx_this
			}, "getTimezoneOffset")
		}()
	case "getUTCDate":
		return func() *hxrt.Function[func() int] {
			hx_receiver_407 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_407.__hx_this.getUTCDate()
			}, func() any {
				return hx_receiver_407.__hx_this
			}, "getUTCDate")
		}()
	case "getUTCDay":
		return func() *hxrt.Function[func() int] {
			hx_receiver_408 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_408.__hx_this.getUTCDay()
			}, func() any {
				return hx_receiver_408.__hx_this
			}, "getUTCDay")
		}()
	case "getUTCFullYear":
		return func() *hxrt.Function[func() int] {
			hx_receiver_409 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_409.__hx_this.getUTCFullYear()
			}, func() any {
				return hx_receiver_409.__hx_this
			}, "getUTCFullYear")
		}()
	case "getUTCHours":
		return func() *hxrt.Function[func() int] {
			hx_receiver_410 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_410.__hx_this.getUTCHours()
			}, func() any {
				return hx_receiver_410.__hx_this
			}, "getUTCHours")
		}()
	case "getUTCMinutes":
		return func() *hxrt.Function[func() int] {
			hx_receiver_411 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_411.__hx_this.getUTCMinutes()
			}, func() any {
				return hx_receiver_411.__hx_this
			}, "getUTCMinutes")
		}()
	case "getUTCMonth":
		return func() *hxrt.Function[func() int] {
			hx_receiver_412 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_412.__hx_this.getUTCMonth()
			}, func() any {
				return hx_receiver_412.__hx_this
			}, "getUTCMonth")
		}()
	case "getUTCSeconds":
		return func() *hxrt.Function[func() int] {
			hx_receiver_413 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_413.__hx_this.getUTCSeconds()
			}, func() any {
				return hx_receiver_413.__hx_this
			}, "getUTCSeconds")
		}()
	case "localParts":
		return func() *hxrt.Function[func() *hxrt.DateParts] {
			hx_receiver_414 := value
			return hxrt.BoundFunction(func() *hxrt.DateParts {
				return hx_receiver_414.__hx_this.localParts()
			}, func() any {
				return hx_receiver_414.__hx_this
			}, "localParts")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_415 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_415.__hx_this.toString()
			}, func() any {
				return hx_receiver_415.__hx_this
			}, "toString")
		}()
	case "utcParts":
		return func() *hxrt.Function[func() *hxrt.DateParts] {
			hx_receiver_416 := value
			return hxrt.BoundFunction(func() *hxrt.DateParts {
				return hx_receiver_416.__hx_this.utcParts()
			}, func() any {
				return hx_receiver_416.__hx_this
			}, "utcParts")
		}()
	}
	return nil
}

func hxrt__generated_method_field__SerializationSnapshotBase(value *SerializationSnapshotBase, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "dispatch":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_417 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_417.__hx_this.dispatch()
			}, func() any {
				return hx_receiver_417.__hx_this
			}, "dispatch")
		}()
	case "label":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_418 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_418.__hx_this.label()
			}, func() any {
				return hx_receiver_418.__hx_this
			}, "label")
		}()
	case "readBase":
		return func() *hxrt.Function[func() int] {
			hx_receiver_419 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_419.__hx_this.readBase()
			}, func() any {
				return hx_receiver_419.__hx_this
			}, "readBase")
		}()
	}
	return nil
}

func hxrt__generated_method_field__SerializationSnapshotChild(value *SerializationSnapshotChild, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "label":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_420 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_420.__hx_this.label()
			}, func() any {
				return hx_receiver_420.__hx_this
			}, "label")
		}()
	case "readChild":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_421 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_421.__hx_this.readChild()
			}, func() any {
				return hx_receiver_421.__hx_this
			}, "readChild")
		}()
	}
	if value.SerializationSnapshotBase == nil {
		return nil
	}
	return hxrt__generated_method_field__SerializationSnapshotBase(value.SerializationSnapshotBase, key)
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_422 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_422.__hx_this.details()
			}, func() any {
				return hx_receiver_422.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_423 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_423.__hx_this.get_message()
			}, func() any {
				return hx_receiver_423.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_424 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_424.__hx_this.get_native()
			}, func() any {
				return hx_receiver_424.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_425 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_425.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_425.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_426 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_426.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_426.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_427 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_427.__hx_this.toString()
			}, func() any {
				return hx_receiver_427.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_428 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_428.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_428.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Serializer(value *haxe__Serializer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "flushNulls":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_429 := value
			return hxrt.BoundFunction(func(count int) {
				hx_receiver_429.__hx_this.flushNulls(count)
			}, func() any {
				return hx_receiver_429.__hx_this
			}, "flushNulls")
		}()
	case "serialize":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_430 := value
			return hxrt.BoundFunction(func(value any) {
				hx_receiver_430.__hx_this.serialize(value)
			}, func() any {
				return hx_receiver_430.__hx_this
			}, "serialize")
		}()
	case "serializeArray":
		return func() *hxrt.Function[func(*hxrt.Array)] {
			hx_receiver_431 := value
			return hxrt.BoundFunction(func(value *hxrt.Array) {
				hx_receiver_431.__hx_this.serializeArray(value)
			}, func() any {
				return hx_receiver_431.__hx_this
			}, "serializeArray")
		}()
	case "serializeBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes)] {
			hx_receiver_432 := value
			return hxrt.BoundFunction(func(value *haxe__io__Bytes) {
				hx_receiver_432.__hx_this.serializeBytes(value)
			}, func() any {
				return hx_receiver_432.__hx_this
			}, "serializeBytes")
		}()
	case "serializeClass":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_433 := value
			return hxrt.BoundFunction(func(value any, declaration any) {
				hx_receiver_433.__hx_this.serializeClass(value, declaration)
			}, func() any {
				return hx_receiver_433.__hx_this
			}, "serializeClass")
		}()
	case "serializeEnum":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_434 := value
			return hxrt.BoundFunction(func(value any, declaration any) {
				hx_receiver_434.__hx_this.serializeEnum(value, declaration)
			}, func() any {
				return hx_receiver_434.__hx_this
			}, "serializeEnum")
		}()
	case "serializeException":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_435 := value
			return hxrt.BoundFunction(func(value any) {
				hx_receiver_435.__hx_this.serializeException(value)
			}, func() any {
				return hx_receiver_435.__hx_this
			}, "serializeException")
		}()
	case "serializeFields":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_436 := value
			return hxrt.BoundFunction(func(value any) {
				hx_receiver_436.__hx_this.serializeFields(value)
			}, func() any {
				return hx_receiver_436.__hx_this
			}, "serializeFields")
		}()
	case "serializeRef":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_437 := value
			return hxrt.BoundFunction(func(value any) bool {
				return hx_receiver_437.__hx_this.serializeRef(value)
			}, func() any {
				return hx_receiver_437.__hx_this
			}, "serializeRef")
		}()
	case "serializeString":
		return func() *hxrt.Function[func(*string)] {
			hx_receiver_438 := value
			return hxrt.BoundFunction(func(value *string) {
				hx_receiver_438.__hx_this.serializeString(value)
			}, func() any {
				return hx_receiver_438.__hx_this
			}, "serializeString")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_439 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_439.__hx_this.toString()
			}, func() any {
				return hx_receiver_439.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Unserializer(value *haxe__Unserializer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "base64Value":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_440 := value
			return hxrt.BoundFunction(func(code int) int {
				return hx_receiver_440.__hx_this.base64Value(code)
			}, func() any {
				return hx_receiver_440.__hx_this
			}, "base64Value")
		}()
	case "decodeBytes":
		return func() *hxrt.Function[func(*string) *haxe__io__Bytes] {
			hx_receiver_441 := value
			return hxrt.BoundFunction(func(encoded *string) *haxe__io__Bytes {
				return hx_receiver_441.__hx_this.decodeBytes(encoded)
			}, func() any {
				return hx_receiver_441.__hx_this
			}, "decodeBytes")
		}()
	case "get":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_442 := value
			return hxrt.BoundFunction(func(index int) int {
				return hx_receiver_442.__hx_this.get(index)
			}, func() any {
				return hx_receiver_442.__hx_this
			}, "get")
		}()
	case "getResolver":
		return func() *hxrt.Function[func() any] {
			hx_receiver_443 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_443.__hx_this.getResolver()
			}, func() any {
				return hx_receiver_443.__hx_this
			}, "getResolver")
		}()
	case "isLegacyDate":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_444 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_444.__hx_this.isLegacyDate()
			}, func() any {
				return hx_receiver_444.__hx_this
			}, "isLegacyDate")
		}()
	case "readDigits":
		return func() *hxrt.Function[func() int] {
			hx_receiver_445 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_445.__hx_this.readDigits()
			}, func() any {
				return hx_receiver_445.__hx_this
			}, "readDigits")
		}()
	case "readFloat":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_446 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_446.__hx_this.readFloat()
			}, func() any {
				return hx_receiver_446.__hx_this
			}, "readFloat")
		}()
	case "resolveClass":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_447 := value
			return hxrt.BoundFunction(func(name *string) any {
				return hx_receiver_447.__hx_this.resolveClass(name)
			}, func() any {
				return hx_receiver_447.__hx_this
			}, "resolveClass")
		}()
	case "resolveEnum":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_448 := value
			return hxrt.BoundFunction(func(name *string) any {
				return hx_receiver_448.__hx_this.resolveEnum(name)
			}, func() any {
				return hx_receiver_448.__hx_this
			}, "resolveEnum")
		}()
	case "setResolver":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_449 := value
			return hxrt.BoundFunction(func(value any) {
				hx_receiver_449.__hx_this.setResolver(value)
			}, func() any {
				return hx_receiver_449.__hx_this
			}, "setResolver")
		}()
	case "unserialize":
		return func() *hxrt.Function[func() any] {
			hx_receiver_450 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_450.__hx_this.unserialize()
			}, func() any {
				return hx_receiver_450.__hx_this
			}, "unserialize")
		}()
	case "unserializeEnum":
		return func() *hxrt.Function[func(any, *string) any] {
			hx_receiver_451 := value
			return hxrt.BoundFunction(func(declaration any, tag *string) any {
				return hx_receiver_451.__hx_this.unserializeEnum(declaration, tag)
			}, func() any {
				return hx_receiver_451.__hx_this
			}, "unserializeEnum")
		}()
	case "unserializeObject":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_452 := value
			return hxrt.BoundFunction(func(target any) {
				hx_receiver_452.__hx_this.unserializeObject(target)
			}, func() any {
				return hx_receiver_452.__hx_this
			}, "unserializeObject")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_453 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_453.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_453.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__ds__IntMap(value *haxe__ds__IntMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_454 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_454.__hx_this.clear()
			}, func() any {
				return hx_receiver_454.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__IntMap] {
			hx_receiver_455 := value
			return hxrt.BoundFunction(func() *haxe__ds__IntMap {
				return hx_receiver_455.__hx_this.copy()
			}, func() any {
				return hx_receiver_455.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_456 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_456.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_456.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(int) bool] {
			hx_receiver_457 := value
			return hxrt.BoundFunction(func(key int) bool {
				return hx_receiver_457.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_457.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_458 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_458.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_458.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(int) any] {
			hx_receiver_459 := value
			return hxrt.BoundFunction(func(key int) any {
				return hx_receiver_459.__hx_this.get(key)
			}, func() any {
				return hx_receiver_459.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_460 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_460.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_460.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_461 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_461.__hx_this.iterator()
			}, func() any {
				return hx_receiver_461.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_462 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_462.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_462.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_463 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_463.__hx_this.keys()
			}, func() any {
				return hx_receiver_463.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(int) bool] {
			hx_receiver_464 := value
			return hxrt.BoundFunction(func(key int) bool {
				return hx_receiver_464.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_464.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_465 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_465.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_465.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(int, any)] {
			hx_receiver_466 := value
			return hxrt.BoundFunction(func(key int, value any) {
				hx_receiver_466.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_466.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_467 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_467.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_467.__hx_this
			}, "setIMap")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_468 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_468.__hx_this.toString()
			}, func() any {
				return hx_receiver_468.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__List(value *haxe__ds__List, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_469 := value
			return hxrt.BoundFunction(func(item any) {
				hx_receiver_469.__hx_this.add(item)
			}, func() any {
				return hx_receiver_469.__hx_this
			}, "add")
		}()
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_470 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_470.__hx_this.clear()
			}, func() any {
				return hx_receiver_470.__hx_this
			}, "clear")
		}()
	case "filter":
		return func() *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List] {
			hx_receiver_471 := value
			return hxrt.BoundFunction(func(predicate *hxrt.Function[func(any) bool]) *haxe__ds__List {
				return hx_receiver_471.__hx_this.filter(predicate)
			}, func() any {
				return hx_receiver_471.__hx_this
			}, "filter")
		}()
	case "first":
		return func() *hxrt.Function[func() any] {
			hx_receiver_472 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_472.__hx_this.first()
			}, func() any {
				return hx_receiver_472.__hx_this
			}, "first")
		}()
	case "isEmpty":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_473 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_473.__hx_this.isEmpty()
			}, func() any {
				return hx_receiver_473.__hx_this
			}, "isEmpty")
		}()
	case "iterator":
		return func() *hxrt.Function[func() *haxe__ds___List__GoListIterator] {
			hx_receiver_474 := value
			return hxrt.BoundFunction(func() *haxe__ds___List__GoListIterator {
				return hx_receiver_474.__hx_this.iterator()
			}, func() any {
				return hx_receiver_474.__hx_this
			}, "iterator")
		}()
	case "join":
		return func() *hxrt.Function[func(*string) *string] {
			hx_receiver_475 := value
			return hxrt.BoundFunction(func(separator *string) *string {
				return hx_receiver_475.__hx_this.join(separator)
			}, func() any {
				return hx_receiver_475.__hx_this
			}, "join")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator] {
			hx_receiver_476 := value
			return hxrt.BoundFunction(func() *haxe__ds___List__GoListKeyValueIterator {
				return hx_receiver_476.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_476.__hx_this
			}, "keyValueIterator")
		}()
	case "last":
		return func() *hxrt.Function[func() any] {
			hx_receiver_477 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_477.__hx_this.last()
			}, func() any {
				return hx_receiver_477.__hx_this
			}, "last")
		}()
	case "map":
		return func() *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List] {
			hx_receiver_478 := value
			return hxrt.BoundFunction(func(transform *hxrt.Function[func(any) any]) *haxe__ds__List {
				return hx_receiver_478.__hx_this.map_(transform)
			}, func() any {
				return hx_receiver_478.__hx_this
			}, "map")
		}()
	case "pop":
		return func() *hxrt.Function[func() any] {
			hx_receiver_479 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_479.__hx_this.pop()
			}, func() any {
				return hx_receiver_479.__hx_this
			}, "pop")
		}()
	case "push":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_480 := value
			return hxrt.BoundFunction(func(item any) {
				hx_receiver_480.__hx_this.push(item)
			}, func() any {
				return hx_receiver_480.__hx_this
			}, "push")
		}()
	case "remove":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_481 := value
			return hxrt.BoundFunction(func(value any) bool {
				return hx_receiver_481.__hx_this.remove(value)
			}, func() any {
				return hx_receiver_481.__hx_this
			}, "remove")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_482 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_482.__hx_this.toString()
			}, func() any {
				return hx_receiver_482.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__ObjectMap(value *haxe__ds__ObjectMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_483 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_483.__hx_this.clear()
			}, func() any {
				return hx_receiver_483.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__ObjectMap] {
			hx_receiver_484 := value
			return hxrt.BoundFunction(func() *haxe__ds__ObjectMap {
				return hx_receiver_484.__hx_this.copy()
			}, func() any {
				return hx_receiver_484.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_485 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_485.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_485.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_486 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_486.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_486.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_487 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_487.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_487.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_488 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_488.__hx_this.get(key)
			}, func() any {
				return hx_receiver_488.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_489 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_489.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_489.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_490 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_490.__hx_this.iterator()
			}, func() any {
				return hx_receiver_490.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_491 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_491.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_491.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_492 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_492.__hx_this.keys()
			}, func() any {
				return hx_receiver_492.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_493 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_493.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_493.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_494 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_494.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_494.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_495 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_495.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_495.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_496 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_496.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_496.__hx_this
			}, "setIMap")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_497 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_497.__hx_this.toString()
			}, func() any {
				return hx_receiver_497.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_498 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_498.__hx_this.clear()
			}, func() any {
				return hx_receiver_498.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__StringMap] {
			hx_receiver_499 := value
			return hxrt.BoundFunction(func() *haxe__ds__StringMap {
				return hx_receiver_499.__hx_this.copy()
			}, func() any {
				return hx_receiver_499.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_500 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_500.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_500.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_501 := value
			return hxrt.BoundFunction(func(key *string) bool {
				return hx_receiver_501.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_501.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_502 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_502.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_502.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_503 := value
			return hxrt.BoundFunction(func(key *string) any {
				return hx_receiver_503.__hx_this.get(key)
			}, func() any {
				return hx_receiver_503.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_504 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_504.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_504.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_505 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_505.__hx_this.iterator()
			}, func() any {
				return hx_receiver_505.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_506 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_506.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_506.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_507 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_507.__hx_this.keys()
			}, func() any {
				return hx_receiver_507.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_508 := value
			return hxrt.BoundFunction(func(key *string) bool {
				return hx_receiver_508.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_508.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_509 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_509.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_509.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(*string, any)] {
			hx_receiver_510 := value
			return hxrt.BoundFunction(func(key *string, value any) {
				hx_receiver_510.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_510.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_511 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_511.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_511.__hx_this
			}, "setIMap")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_512 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_512.__hx_this.toString()
			}, func() any {
				return hx_receiver_512.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds___List__GoListIterator(value *haxe__ds___List__GoListIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_513 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_513.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_513.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() any] {
			hx_receiver_514 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_514.__hx_this.next()
			}, func() any {
				return hx_receiver_514.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds___List__GoListKeyValueIterator(value *haxe__ds___List__GoListKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_515 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_515.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_515.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_516 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_516.__hx_this.next()
			}, func() any {
				return hx_receiver_516.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_nativeView":
		return func() *hxrt.Function[func() *hxrt.ByteView] {
			hx_receiver_517 := value
			return hxrt.BoundFunction(func() *hxrt.ByteView {
				return hx_receiver_517.__hx_this.__hx_nativeView()
			}, func() any {
				return hx_receiver_517.__hx_this
			}, "__hx_nativeView")
		}()
	case "blit":
		return func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
			hx_receiver_518 := value
			return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
				hx_receiver_518.__hx_this.blit(pos, src, srcpos, len)
			}, func() any {
				return hx_receiver_518.__hx_this
			}, "blit")
		}()
	case "compare":
		return func() *hxrt.Function[func(*haxe__io__Bytes) int] {
			hx_receiver_519 := value
			return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
				return hx_receiver_519.__hx_this.compare(other)
			}, func() any {
				return hx_receiver_519.__hx_this
			}, "compare")
		}()
	case "fill":
		return func() *hxrt.Function[func(int, int, int)] {
			hx_receiver_520 := value
			return hxrt.BoundFunction(func(pos int, len int, value int) {
				hx_receiver_520.__hx_this.fill(pos, len, value)
			}, func() any {
				return hx_receiver_520.__hx_this
			}, "fill")
		}()
	case "get":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_521 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_521.__hx_this.get(pos)
			}, func() any {
				return hx_receiver_521.__hx_this
			}, "get")
		}()
	case "getData":
		return func() *hxrt.Function[func() []int] {
			hx_receiver_522 := value
			return hxrt.BoundFunction(func() []int {
				return hx_receiver_522.__hx_this.getData()
			}, func() any {
				return hx_receiver_522.__hx_this
			}, "getData")
		}()
	case "getDouble":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_523 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_523.__hx_this.getDouble(pos)
			}, func() any {
				return hx_receiver_523.__hx_this
			}, "getDouble")
		}()
	case "getFloat":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_524 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_524.__hx_this.getFloat(pos)
			}, func() any {
				return hx_receiver_524.__hx_this
			}, "getFloat")
		}()
	case "getInt32":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_525 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_525.__hx_this.getInt32(pos)
			}, func() any {
				return hx_receiver_525.__hx_this
			}, "getInt32")
		}()
	case "getInt64":
		return func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
			hx_receiver_526 := value
			return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
				return hx_receiver_526.__hx_this.getInt64(pos)
			}, func() any {
				return hx_receiver_526.__hx_this
			}, "getInt64")
		}()
	case "getString":
		return func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
			hx_receiver_527 := value
			return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_527.__hx_this.getString(pos, len, encoding)
			}, func() any {
				return hx_receiver_527.__hx_this
			}, "getString")
		}()
	case "getUInt16":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_528 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_528.__hx_this.getUInt16(pos)
			}, func() any {
				return hx_receiver_528.__hx_this
			}, "getUInt16")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, int) *string] {
			hx_receiver_529 := value
			return hxrt.BoundFunction(func(pos int, len int) *string {
				return hx_receiver_529.__hx_this.readString(pos, len)
			}, func() any {
				return hx_receiver_529.__hx_this
			}, "readString")
		}()
	case "set":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_530 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_530.__hx_this.set(pos, value)
			}, func() any {
				return hx_receiver_530.__hx_this
			}, "set")
		}()
	case "setDouble":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_531 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_531.__hx_this.setDouble(pos, value)
			}, func() any {
				return hx_receiver_531.__hx_this
			}, "setDouble")
		}()
	case "setFloat":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_532 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_532.__hx_this.setFloat(pos, value)
			}, func() any {
				return hx_receiver_532.__hx_this
			}, "setFloat")
		}()
	case "setInt32":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_533 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_533.__hx_this.setInt32(pos, value)
			}, func() any {
				return hx_receiver_533.__hx_this
			}, "setInt32")
		}()
	case "setInt64":
		return func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
			hx_receiver_534 := value
			return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
				hx_receiver_534.__hx_this.setInt64(pos, value)
			}, func() any {
				return hx_receiver_534.__hx_this
			}, "setInt64")
		}()
	case "setUInt16":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_535 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_535.__hx_this.setUInt16(pos, value)
			}, func() any {
				return hx_receiver_535.__hx_this
			}, "setUInt16")
		}()
	case "sub":
		return func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
			hx_receiver_536 := value
			return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
				return hx_receiver_536.__hx_this.sub(pos, len)
			}, func() any {
				return hx_receiver_536.__hx_this
			}, "sub")
		}()
	case "toHex":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_537 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_537.__hx_this.toHex()
			}, func() any {
				return hx_receiver_537.__hx_this
			}, "toHex")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_538 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_538.__hx_this.toString()
			}, func() any {
				return hx_receiver_538.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_539 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_539.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_539.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_540 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_540.__hx_this.next()
			}, func() any {
				return hx_receiver_540.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_541 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_541.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_541.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_542 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_542.__hx_this.next()
			}, func() any {
				return hx_receiver_542.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_543 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_543.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_543.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_544 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_544.__hx_this.next()
			}, func() any {
				return hx_receiver_544.__hx_this
			}, "next")
		}()
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

func hxrt_typeCreateEmpty__haxe__Exception() *haxe__Exception {
	instance := &haxe__Exception{}
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

func hxrt_typeCreateEmpty__haxe__ValueException() *haxe__ValueException {
	instance := &haxe__ValueException{}
	instance.haxe__Exception = &haxe__Exception{}
	instance.haxe__Exception.__hx_this = instance
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
	case "haxe.Exception":
		return hxrt_typeCallAny(New_haxe__Exception, args)
	case "haxe.Int64Helper":
		return nil, false
	case "haxe.Serializer":
		return hxrt_typeCallAny(New_haxe__Serializer, args)
	case "haxe.Unserializer":
		return hxrt_typeCallAny(New_haxe__Unserializer, args)
	case "haxe.ValueException":
		return hxrt_typeCallAny(New_haxe__ValueException, args)
	case "haxe._CallStack.CallStack_Impl_":
		return nil, false
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
	case "haxe.Exception":
		return hxrt_typeCreateEmpty__haxe__Exception(), true
	case "haxe.Serializer":
		return hxrt_typeCreateEmpty__haxe__Serializer(), true
	case "haxe.Unserializer":
		return hxrt_typeCreateEmpty__haxe__Unserializer(), true
	case "haxe.ValueException":
		return hxrt_typeCreateEmpty__haxe__ValueException(), true
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
	case "haxe.StackItem":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__StackItem_CFunction, true
			case 1:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_Module, args)
			case 2:
				if len(args) != 4 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_FilePos, args)
			case 3:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_Method, args)
			case 4:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_LocalFunction, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "CFunction":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__StackItem_CFunction, true
		case "Module":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_Module, args)
		case "FilePos":
			if len(args) != 4 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_FilePos, args)
		case "Method":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_Method, args)
		case "LocalFunction":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_LocalFunction, args)
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
	o = hxrt.ObjectSource(o)
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
	case *haxe__Exception:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
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
	case *haxe__ValueException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ValueException")}
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
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.StackItem")}
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
	case "haxe.Exception":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe.Serializer":
		return nil
	case "haxe.Unserializer":
		return nil
	case "haxe.ValueException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "haxe._CallStack.CallStack_Impl_":
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
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("catchValue"), hxrt.StringFromLiteral("caught"), hxrt.StringFromLiteral("thrown"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Serializer":
		return hxrt.NewArray(hxrt.StringFromLiteral("BASE64"), hxrt.StringFromLiteral("USE_CACHE"), hxrt.StringFromLiteral("USE_ENUM_INDEX"), hxrt.StringFromLiteral("run"))
	case "haxe.Unserializer":
		return hxrt.NewArray(hxrt.StringFromLiteral("BASE64"), hxrt.StringFromLiteral("DEFAULT_RESOLVER"), hxrt.StringFromLiteral("NULL_RESOLVER"), hxrt.StringFromLiteral("run"))
	case "haxe.ValueException":
		return hxrt.NewArray()
	case "haxe._CallStack.CallStack_Impl_":
		return hxrt.NewArray(hxrt.StringFromLiteral("asArray"), hxrt.StringFromLiteral("callStack"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("equalItems"), hxrt.StringFromLiteral("exceptionStack"), hxrt.StringFromLiteral("exceptionToString"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("get_length"), hxrt.StringFromLiteral("itemToString"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("subtract"), hxrt.StringFromLiteral("toString"))
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
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Serializer":
		return hxrt.NewArray(hxrt.StringFromLiteral("buf"), hxrt.StringFromLiteral("cache"), hxrt.StringFromLiteral("flushNulls"), hxrt.StringFromLiteral("scount"), hxrt.StringFromLiteral("serialize"), hxrt.StringFromLiteral("serializeArray"), hxrt.StringFromLiteral("serializeBytes"), hxrt.StringFromLiteral("serializeClass"), hxrt.StringFromLiteral("serializeEnum"), hxrt.StringFromLiteral("serializeException"), hxrt.StringFromLiteral("serializeFields"), hxrt.StringFromLiteral("serializeRef"), hxrt.StringFromLiteral("serializeString"), hxrt.StringFromLiteral("shash"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("useCache"), hxrt.StringFromLiteral("useEnumIndex"))
	case "haxe.Unserializer":
		return hxrt.NewArray(hxrt.StringFromLiteral("base64Value"), hxrt.StringFromLiteral("buf"), hxrt.StringFromLiteral("cache"), hxrt.StringFromLiteral("decodeBytes"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getResolver"), hxrt.StringFromLiteral("isLegacyDate"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("pos"), hxrt.StringFromLiteral("readDigits"), hxrt.StringFromLiteral("readFloat"), hxrt.StringFromLiteral("resolveClass"), hxrt.StringFromLiteral("resolveEnum"), hxrt.StringFromLiteral("resolver"), hxrt.StringFromLiteral("scache"), hxrt.StringFromLiteral("setResolver"), hxrt.StringFromLiteral("unserialize"), hxrt.StringFromLiteral("unserializeEnum"), hxrt.StringFromLiteral("unserializeObject"))
	case "haxe.ValueException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"), hxrt.StringFromLiteral("value"))
	case "haxe._CallStack.CallStack_Impl_":
		return hxrt.NewArray()
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
	case "haxe.Exception":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Serializer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Unserializer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ValueException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._CallStack.CallStack_Impl_":
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
	case "haxe.StackItem":
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
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("CFunction")
		case 1:
			return hxrt.StringFromLiteral("Module")
		case 2:
			return hxrt.StringFromLiteral("FilePos")
		case 3:
			return hxrt.StringFromLiteral("Method")
		case 4:
			return hxrt.StringFromLiteral("LocalFunction")
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
	case *haxe__StackItem:
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
	case "haxe.StackItem":
		return hxrt.NewArray(hxrt.StringFromLiteral("CFunction"), hxrt.StringFromLiteral("Module"), hxrt.StringFromLiteral("FilePos"), hxrt.StringFromLiteral("Method"), hxrt.StringFromLiteral("LocalFunction"))
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
	case *haxe__StackItem:
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
	case "haxe.StackItem":
		return hxrt.NewArray(haxe__StackItem_CFunction)
	case "haxe.io.Encoding":
		return hxrt.NewArray(haxe__io__Encoding_UTF8, haxe__io__Encoding_RawNative)
	case "haxe.io.Error":
		return hxrt.NewArray(haxe__io__Error_Blocked, haxe__io__Error_Overflow, haxe__io__Error_OutsideBounds)
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) *ValueType {
	v = hxrt.ObjectSource(v)
	if hxrt.ReflectIsFunction(v) {
		return ValueType_TFunction
	}
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *haxe__Exception:
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
	case *haxe__ValueException:
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
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__Serializer:
		return hxrt__generated_field_lookup__haxe__Serializer(value, key)
	case *haxe__Unserializer:
		return hxrt__generated_field_lookup__haxe__Unserializer(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *haxe__Exception:
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
	case *haxe__ValueException:
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
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__Serializer:
		return hxrt__generated_field_has__haxe__Serializer(value, key)
	case *haxe__Unserializer:
		return hxrt__generated_field_has__haxe__Unserializer(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *haxe__Exception:
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
	case *haxe__ValueException:
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
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__Serializer:
		return hxrt__generated_field_set__haxe__Serializer(value, key, incoming)
	case *haxe__Unserializer:
		return hxrt__generated_field_set__haxe__Unserializer(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *haxe__Exception:
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
	case *haxe__ValueException:
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
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__Serializer:
		return hxrt.NewArray(hxrt.StringFromLiteral("buf"), hxrt.StringFromLiteral("cache"), hxrt.StringFromLiteral("shash"), hxrt.StringFromLiteral("scount"), hxrt.StringFromLiteral("useCache"), hxrt.StringFromLiteral("useEnumIndex"))
	case *haxe__Unserializer:
		return hxrt.NewArray(hxrt.StringFromLiteral("buf"), hxrt.StringFromLiteral("pos"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("cache"), hxrt.StringFromLiteral("scache"), hxrt.StringFromLiteral("resolver"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
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
		return hxrt.ReadStoredField(value.ms)
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
		return hxrt.ReadStoredField(value.baseValue)
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
		return hxrt.ReadStoredField(value.childValue)
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
		return hxrt.ReadStoredField(value.b)
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

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
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
		return hxrt.ReadStoredField(value.buf)
	case "cache":
		return hxrt.ReadStoredField(value.cache)
	case "scount":
		return hxrt.ReadStoredField(value.scount)
	case "shash":
		return hxrt.ReadStoredField(value.shash)
	case "useCache":
		return hxrt.ReadStoredField(value.useCache)
	case "useEnumIndex":
		return hxrt.ReadStoredField(value.useEnumIndex)
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
		return hxrt.ReadStoredField(value.buf)
	case "cache":
		return hxrt.ReadStoredField(value.cache)
	case "length":
		return hxrt.ReadStoredField(value.length)
	case "pos":
		return hxrt.ReadStoredField(value.pos)
	case "resolver":
		return hxrt.ReadStoredField(value.resolver)
	case "scache":
		return hxrt.ReadStoredField(value.scache)
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

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
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
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
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
		return hxrt.ReadStoredField(value.h)
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
		return hxrt.ReadStoredField(value.items)
	case "length":
		return hxrt.ReadStoredField(value.length)
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
		return hxrt.ReadStoredField(value.h)
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
		return hxrt.ReadStoredField(value.h)
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
		return hxrt.ReadStoredField(value.index)
	case "items":
		return hxrt.ReadStoredField(value.items)
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
		return hxrt.ReadStoredField(value.index)
	case "items":
		return hxrt.ReadStoredField(value.items)
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
		return hxrt.ReadStoredField(value.__hx_dataExposed)
	case "__hx_raw":
		return hxrt.ReadStoredField(value.__hx_raw)
	case "__hx_rawValid":
		return hxrt.ReadStoredField(value.__hx_rawValid)
	case "b":
		return hxrt.ReadStoredField(value.b)
	case "length":
		return hxrt.ReadStoredField(value.length)
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
		return hxrt.ReadStoredField(value.keys)
	case "map":
		return hxrt.ReadStoredField(value.map_)
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
			var zero hxrt.Anonymous
			value.keys = zero
			return true
		}
		switch typed := incoming.(type) {
		case hxrt.Anonymous:
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
		return hxrt.ReadStoredField(value.offset)
	case "s":
		return hxrt.ReadStoredField(value.s)
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
		return hxrt.ReadStoredField(value.offset)
	case "s":
		return hxrt.ReadStoredField(value.s)
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
	case *haxe__StackItem:
		return (enumValue != nil)
	case *haxe__io__Encoding:
		return (enumValue != nil)
	case *haxe__io__Error:
		return (enumValue != nil)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_545 any) *string {
			if hx_value_545 == nil {
				var hx_zero_546 *string
				return hx_zero_546
			}
			return hx_value_545.(*string)
		}(argument_0))
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() float64]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *hxrt.DateParts]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_547 any) int {
			if hx_value_547 == nil {
				var hx_zero_548 int
				return hx_zero_548
			}
			return hx_value_547.(int)
		}(argument_0))
		return nil
	case *hxrt.Function[func(any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(argument_0)
		return nil
	case *hxrt.Function[func(*hxrt.Array)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_549 any) *hxrt.Array {
			if hx_value_549 == nil {
				var hx_zero_550 *hxrt.Array
				return hx_zero_550
			}
			return hx_value_549.(*hxrt.Array)
		}(argument_0))
		return nil
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_7(argument_0))
		return nil
	case *hxrt.Function[func(any, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(argument_0, argument_1)
		return nil
	case *hxrt.Function[func(any) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*string)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_551 any) *string {
			if hx_value_551 == nil {
				var hx_zero_552 *string
				return hx_zero_552
			}
			return hx_value_551.(*string)
		}(argument_0))
		return nil
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_553 any) int {
			if hx_value_553 == nil {
				var hx_zero_554 int
				return hx_zero_554
			}
			return hx_value_553.(int)
		}(argument_0))
	case *hxrt.Function[func(*string) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_555 any) *string {
			if hx_value_555 == nil {
				var hx_zero_556 *string
				return hx_zero_556
			}
			return hx_value_555.(*string)
		}(argument_0))
	case *hxrt.Function[func(any, *string) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, func(hx_value_557 any) *string {
			if hx_value_557 == nil {
				var hx_zero_558 *string
				return hx_zero_558
			}
			return hx_value_557.(*string)
		}(argument_1))
	case *hxrt.Function[func()]:
		hxrt.FunctionCall(source)()
		return nil
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_559 any) int {
			if hx_value_559 == nil {
				var hx_zero_560 int
				return hx_zero_560
			}
			return hx_value_559.(int)
		}(argument_0))
	case *hxrt.Function[func(int) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_561 any) int {
			if hx_value_561 == nil {
				var hx_zero_562 int
				return hx_zero_562
			}
			return hx_value_561.(int)
		}(argument_0))
	case *hxrt.Function[func(any) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(int, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_563 any) int {
			if hx_value_563 == nil {
				var hx_zero_564 int
				return hx_zero_564
			}
			return hx_value_563.(int)
		}(argument_0), argument_1)
		return nil
	case *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_4(argument_0))
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_565 any) *string {
			if hx_value_565 == nil {
				var hx_zero_566 *string
				return hx_zero_566
			}
			return hx_value_565.(*string)
		}(argument_0))
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_5(argument_0))
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_567 any) *string {
			if hx_value_567 == nil {
				var hx_zero_568 *string
				return hx_zero_568
			}
			return hx_value_567.(*string)
		}(argument_0))
	case *hxrt.Function[func(*string, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_569 any) *string {
			if hx_value_569 == nil {
				var hx_zero_570 *string
				return hx_zero_570
			}
			return hx_value_569.(*string)
		}(argument_0), argument_1)
		return nil
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int, *haxe__io__Bytes, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		var argument_3 any = nil
		if 3 < len(arguments) {
			argument_3 = arguments[3]
		}
		hxrt.FunctionCall(source)(func(hx_value_571 any) int {
			if hx_value_571 == nil {
				var hx_zero_572 int
				return hx_zero_572
			}
			return hx_value_571.(int)
		}(argument_0), hx_recover_class_7(argument_1), func(hx_value_573 any) int {
			if hx_value_573 == nil {
				var hx_zero_574 int
				return hx_zero_574
			}
			return hx_value_573.(int)
		}(argument_2), func(hx_value_575 any) int {
			if hx_value_575 == nil {
				var hx_zero_576 int
				return hx_zero_576
			}
			return hx_value_575.(int)
		}(argument_3))
		return nil
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_7(argument_0))
	case *hxrt.Function[func(int, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		hxrt.FunctionCall(source)(func(hx_value_577 any) int {
			if hx_value_577 == nil {
				var hx_zero_578 int
				return hx_zero_578
			}
			return hx_value_577.(int)
		}(argument_0), func(hx_value_579 any) int {
			if hx_value_579 == nil {
				var hx_zero_580 int
				return hx_zero_580
			}
			return hx_value_579.(int)
		}(argument_1), func(hx_value_581 any) int {
			if hx_value_581 == nil {
				var hx_zero_582 int
				return hx_zero_582
			}
			return hx_value_581.(int)
		}(argument_2))
		return nil
	case *hxrt.Function[func() []int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) float64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_583 any) int {
			if hx_value_583 == nil {
				var hx_zero_584 int
				return hx_zero_584
			}
			return hx_value_583.(int)
		}(argument_0))
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_585 any) int {
			if hx_value_585 == nil {
				var hx_zero_586 int
				return hx_zero_586
			}
			return hx_value_585.(int)
		}(argument_0))
	case *hxrt.Function[func(int, int, *haxe__io__Encoding) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(func(hx_value_587 any) int {
			if hx_value_587 == nil {
				var hx_zero_588 int
				return hx_zero_588
			}
			return hx_value_587.(int)
		}(argument_0), func(hx_value_589 any) int {
			if hx_value_589 == nil {
				var hx_zero_590 int
				return hx_zero_590
			}
			return hx_value_589.(int)
		}(argument_1), func(hx_value_591 any) *haxe__io__Encoding {
			if hx_value_591 == nil {
				var hx_zero_592 *haxe__io__Encoding
				return hx_zero_592
			}
			return hx_value_591.(*haxe__io__Encoding)
		}(argument_2))
	case *hxrt.Function[func(int, int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_593 any) int {
			if hx_value_593 == nil {
				var hx_zero_594 int
				return hx_zero_594
			}
			return hx_value_593.(int)
		}(argument_0), func(hx_value_595 any) int {
			if hx_value_595 == nil {
				var hx_zero_596 int
				return hx_zero_596
			}
			return hx_value_595.(int)
		}(argument_1))
	case *hxrt.Function[func(int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_597 any) int {
			if hx_value_597 == nil {
				var hx_zero_598 int
				return hx_zero_598
			}
			return hx_value_597.(int)
		}(argument_0), func(hx_value_599 any) int {
			if hx_value_599 == nil {
				var hx_zero_600 int
				return hx_zero_600
			}
			return hx_value_599.(int)
		}(argument_1))
		return nil
	case *hxrt.Function[func(int, float64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_601 any) int {
			if hx_value_601 == nil {
				var hx_zero_602 int
				return hx_zero_602
			}
			return hx_value_601.(int)
		}(argument_0), func(hx_value_603 any) float64 {
			if hx_value_603 == nil {
				var hx_zero_604 float64
				return hx_zero_604
			}
			return hx_value_603.(float64)
		}(argument_1))
		return nil
	case *hxrt.Function[func(int, *haxe___Int64_____Int64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_605 any) int {
			if hx_value_605 == nil {
				var hx_zero_606 int
				return hx_zero_606
			}
			return hx_value_605.(int)
		}(argument_0), hx_recover_class_10(argument_1))
		return nil
	case *hxrt.Function[func(int, int) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_607 any) int {
			if hx_value_607 == nil {
				var hx_zero_608 int
				return hx_zero_608
			}
			return hx_value_607.(int)
		}(argument_0), func(hx_value_609 any) int {
			if hx_value_609 == nil {
				var hx_zero_610 int
				return hx_zero_610
			}
			return hx_value_609.(int)
		}(argument_1))
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_611 any) bool {
				if hx_value_611 == nil {
					var hx_zero_612 bool
					return hx_zero_612
				}
				return hx_value_611.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_613 any) bool {
				if hx_value_613 == nil {
					var hx_zero_614 bool
					return hx_zero_614
				}
				return hx_value_613.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_615 any) bool {
				if hx_value_615 == nil {
					var hx_zero_616 bool
					return hx_zero_616
				}
				return hx_value_615.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_617 any) bool {
				if hx_value_617 == nil {
					var hx_zero_618 bool
					return hx_zero_618
				}
				return hx_value_617.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_619 any) bool {
				if hx_value_619 == nil {
					var hx_zero_620 bool
					return hx_zero_620
				}
				return hx_value_619.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_621 any) bool {
				if hx_value_621 == nil {
					var hx_zero_622 bool
					return hx_zero_622
				}
				return hx_value_621.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.DateParts]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_623 any) bool {
				if hx_value_623 == nil {
					var hx_zero_624 bool
					return hx_zero_624
				}
				return hx_value_623.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_625 any) bool {
				if hx_value_625 == nil {
					var hx_zero_626 bool
					return hx_zero_626
				}
				return hx_value_625.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_627 any) bool {
				if hx_value_627 == nil {
					var hx_zero_628 bool
					return hx_zero_628
				}
				return hx_value_627.(bool)
			}(nil)
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_629 any) bool {
				if hx_value_629 == nil {
					var hx_zero_630 bool
					return hx_zero_630
				}
				return hx_value_629.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_631 any) bool {
				if hx_value_631 == nil {
					var hx_zero_632 bool
					return hx_zero_632
				}
				return hx_value_631.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_633 any) bool {
				if hx_value_633 == nil {
					var hx_zero_634 bool
					return hx_zero_634
				}
				return hx_value_633.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_635 any) bool {
				if hx_value_635 == nil {
					var hx_zero_636 bool
					return hx_zero_636
				}
				return hx_value_635.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_637 any) bool {
				if hx_value_637 == nil {
					var hx_zero_638 bool
					return hx_zero_638
				}
				return hx_value_637.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_639 any) bool {
				if hx_value_639 == nil {
					var hx_zero_640 bool
					return hx_zero_640
				}
				return hx_value_639.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_641 any) bool {
				if hx_value_641 == nil {
					var hx_zero_642 bool
					return hx_zero_642
				}
				return hx_value_641.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_643 any) bool {
				if hx_value_643 == nil {
					var hx_zero_644 bool
					return hx_zero_644
				}
				return hx_value_643.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_645 any) *string {
				if hx_value_645 == nil {
					var hx_zero_646 *string
					return hx_zero_646
				}
				return hx_value_645.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_647 any) *string {
				if hx_value_647 == nil {
					var hx_zero_648 *string
					return hx_zero_648
				}
				return hx_value_647.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_649 any) *string {
				if hx_value_649 == nil {
					var hx_zero_650 *string
					return hx_zero_650
				}
				return hx_value_649.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_651 any) *string {
				if hx_value_651 == nil {
					var hx_zero_652 *string
					return hx_zero_652
				}
				return hx_value_651.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_653 any) *string {
				if hx_value_653 == nil {
					var hx_zero_654 *string
					return hx_zero_654
				}
				return hx_value_653.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_655 any) *string {
				if hx_value_655 == nil {
					var hx_zero_656 *string
					return hx_zero_656
				}
				return hx_value_655.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.DateParts]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_657 any) *string {
				if hx_value_657 == nil {
					var hx_zero_658 *string
					return hx_zero_658
				}
				return hx_value_657.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_659 any) *string {
				if hx_value_659 == nil {
					var hx_zero_660 *string
					return hx_zero_660
				}
				return hx_value_659.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_661 any) *string {
				if hx_value_661 == nil {
					var hx_zero_662 *string
					return hx_zero_662
				}
				return hx_value_661.(*string)
			}(nil)
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_663 any) *string {
				if hx_value_663 == nil {
					var hx_zero_664 *string
					return hx_zero_664
				}
				return hx_value_663.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_665 any) *string {
				if hx_value_665 == nil {
					var hx_zero_666 *string
					return hx_zero_666
				}
				return hx_value_665.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_667 any) *string {
				if hx_value_667 == nil {
					var hx_zero_668 *string
					return hx_zero_668
				}
				return hx_value_667.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_669 any) *string {
				if hx_value_669 == nil {
					var hx_zero_670 *string
					return hx_zero_670
				}
				return hx_value_669.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_671 any) *string {
				if hx_value_671 == nil {
					var hx_zero_672 *string
					return hx_zero_672
				}
				return hx_value_671.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_673 any) *string {
				if hx_value_673 == nil {
					var hx_zero_674 *string
					return hx_zero_674
				}
				return hx_value_673.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_675 any) *string {
				if hx_value_675 == nil {
					var hx_zero_676 *string
					return hx_zero_676
				}
				return hx_value_675.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_677 any) *string {
				if hx_value_677 == nil {
					var hx_zero_678 *string
					return hx_zero_678
				}
				return hx_value_677.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.DateParts]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_3(value any) *hxrt.Function[func() int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_679 any) int {
				if hx_value_679 == nil {
					var hx_zero_680 int
					return hx_zero_680
				}
				return hx_value_679.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_681 any) int {
				if hx_value_681 == nil {
					var hx_zero_682 int
					return hx_zero_682
				}
				return hx_value_681.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_683 any) int {
				if hx_value_683 == nil {
					var hx_zero_684 int
					return hx_zero_684
				}
				return hx_value_683.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_685 any) int {
				if hx_value_685 == nil {
					var hx_zero_686 int
					return hx_zero_686
				}
				return hx_value_685.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_687 any) int {
				if hx_value_687 == nil {
					var hx_zero_688 int
					return hx_zero_688
				}
				return hx_value_687.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_689 any) int {
				if hx_value_689 == nil {
					var hx_zero_690 int
					return hx_zero_690
				}
				return hx_value_689.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.DateParts]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_691 any) int {
				if hx_value_691 == nil {
					var hx_zero_692 int
					return hx_zero_692
				}
				return hx_value_691.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_693 any) int {
				if hx_value_693 == nil {
					var hx_zero_694 int
					return hx_zero_694
				}
				return hx_value_693.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() int {
			hxrt.FunctionCall(source)()
			return func(hx_value_695 any) int {
				if hx_value_695 == nil {
					var hx_zero_696 int
					return hx_zero_696
				}
				return hx_value_695.(int)
			}(nil)
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_697 any) int {
				if hx_value_697 == nil {
					var hx_zero_698 int
					return hx_zero_698
				}
				return hx_value_697.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_699 any) int {
				if hx_value_699 == nil {
					var hx_zero_700 int
					return hx_zero_700
				}
				return hx_value_699.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_701 any) int {
				if hx_value_701 == nil {
					var hx_zero_702 int
					return hx_zero_702
				}
				return hx_value_701.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_703 any) int {
				if hx_value_703 == nil {
					var hx_zero_704 int
					return hx_zero_704
				}
				return hx_value_703.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_705 any) int {
				if hx_value_705 == nil {
					var hx_zero_706 int
					return hx_zero_706
				}
				return hx_value_705.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_707 any) int {
				if hx_value_707 == nil {
					var hx_zero_708 int
					return hx_zero_708
				}
				return hx_value_707.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_709 any) int {
				if hx_value_709 == nil {
					var hx_zero_710 int
					return hx_zero_710
				}
				return hx_value_709.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_711 any) int {
				if hx_value_711 == nil {
					var hx_zero_712 int
					return hx_zero_712
				}
				return hx_value_711.(int)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_4(value any) *hxrt.Function[func(any) bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any) bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_715 any) bool {
				if hx_value_715 == nil {
					var hx_zero_716 bool
					return hx_zero_716
				}
				return hx_value_715.(bool)
			}(hxrt.FunctionCall(source)(func(hx_value_713 any) *string {
				if hx_value_713 == nil {
					var hx_zero_714 *string
					return hx_zero_714
				}
				return hx_value_713.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_717 any) bool {
				if hx_value_717 == nil {
					var hx_zero_718 bool
					return hx_zero_718
				}
				return hx_value_717.(bool)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			hxrt.FunctionCall(source)(func(hx_value_719 any) int {
				if hx_value_719 == nil {
					var hx_zero_720 int
					return hx_zero_720
				}
				return hx_value_719.(int)
			}(arg_0))
			return func(hx_value_721 any) bool {
				if hx_value_721 == nil {
					var hx_zero_722 bool
					return hx_zero_722
				}
				return hx_value_721.(bool)
			}(nil)
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			hxrt.FunctionCall(source)(arg_0)
			return func(hx_value_723 any) bool {
				if hx_value_723 == nil {
					var hx_zero_724 bool
					return hx_zero_724
				}
				return hx_value_723.(bool)
			}(nil)
		})
	case *hxrt.Function[func(*hxrt.Array)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			hxrt.FunctionCall(source)(func(hx_value_725 any) *hxrt.Array {
				if hx_value_725 == nil {
					var hx_zero_726 *hxrt.Array
					return hx_zero_726
				}
				return hx_value_725.(*hxrt.Array)
			}(arg_0))
			return func(hx_value_727 any) bool {
				if hx_value_727 == nil {
					var hx_zero_728 bool
					return hx_zero_728
				}
				return hx_value_727.(bool)
			}(nil)
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
			return func(hx_value_729 any) bool {
				if hx_value_729 == nil {
					var hx_zero_730 bool
					return hx_zero_730
				}
				return hx_value_729.(bool)
			}(nil)
		})
	case *hxrt.Function[func(any) bool]:
		return source
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			hxrt.FunctionCall(source)(func(hx_value_731 any) *string {
				if hx_value_731 == nil {
					var hx_zero_732 *string
					return hx_zero_732
				}
				return hx_value_731.(*string)
			}(arg_0))
			return func(hx_value_733 any) bool {
				if hx_value_733 == nil {
					var hx_zero_734 bool
					return hx_zero_734
				}
				return hx_value_733.(bool)
			}(nil)
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_737 any) bool {
				if hx_value_737 == nil {
					var hx_zero_738 bool
					return hx_zero_738
				}
				return hx_value_737.(bool)
			}(hxrt.FunctionCall(source)(func(hx_value_735 any) int {
				if hx_value_735 == nil {
					var hx_zero_736 int
					return hx_zero_736
				}
				return hx_value_735.(int)
			}(arg_0)))
		})
	case *hxrt.Function[func(*string) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_741 any) bool {
				if hx_value_741 == nil {
					var hx_zero_742 bool
					return hx_zero_742
				}
				return hx_value_741.(bool)
			}(hxrt.FunctionCall(source)(func(hx_value_739 any) *string {
				if hx_value_739 == nil {
					var hx_zero_740 *string
					return hx_zero_740
				}
				return hx_value_739.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(int) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return hxrt.FunctionCall(source)(func(hx_value_743 any) int {
				if hx_value_743 == nil {
					var hx_zero_744 int
					return hx_zero_744
				}
				return hx_value_743.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_747 any) bool {
				if hx_value_747 == nil {
					var hx_zero_748 bool
					return hx_zero_748
				}
				return hx_value_747.(bool)
			}(hxrt.FunctionCall(source)(func(hx_value_745 any) int {
				if hx_value_745 == nil {
					var hx_zero_746 int
					return hx_zero_746
				}
				return hx_value_745.(int)
			}(arg_0)))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_749 any) bool {
				if hx_value_749 == nil {
					var hx_zero_750 bool
					return hx_zero_750
				}
				return hx_value_749.(bool)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_751 any) bool {
				if hx_value_751 == nil {
					var hx_zero_752 bool
					return hx_zero_752
				}
				return hx_value_751.(bool)
			}(hxrt.FunctionCall(source)(hx_recover_callable_4(arg_0)))
		})
	case *hxrt.Function[func(*string) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_755 any) bool {
				if hx_value_755 == nil {
					var hx_zero_756 bool
					return hx_zero_756
				}
				return hx_value_755.(bool)
			}(hxrt.FunctionCall(source)(func(hx_value_753 any) *string {
				if hx_value_753 == nil {
					var hx_zero_754 *string
					return hx_zero_754
				}
				return hx_value_753.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_757 any) bool {
				if hx_value_757 == nil {
					var hx_zero_758 bool
					return hx_zero_758
				}
				return hx_value_757.(bool)
			}(hxrt.FunctionCall(source)(hx_recover_callable_5(arg_0)))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return hxrt.FunctionCall(source)(func(hx_value_759 any) *string {
				if hx_value_759 == nil {
					var hx_zero_760 *string
					return hx_zero_760
				}
				return hx_value_759.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_761 any) bool {
				if hx_value_761 == nil {
					var hx_zero_762 bool
					return hx_zero_762
				}
				return hx_value_761.(bool)
			}(hxrt.FunctionCall(source)(hx_recover_class_7(arg_0)))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_765 any) bool {
				if hx_value_765 == nil {
					var hx_zero_766 bool
					return hx_zero_766
				}
				return hx_value_765.(bool)
			}(hxrt.FunctionCall(source)(func(hx_value_763 any) int {
				if hx_value_763 == nil {
					var hx_zero_764 int
					return hx_zero_764
				}
				return hx_value_763.(int)
			}(arg_0)))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_769 any) bool {
				if hx_value_769 == nil {
					var hx_zero_770 bool
					return hx_zero_770
				}
				return hx_value_769.(bool)
			}(hxrt.FunctionCall(source)(func(hx_value_767 any) int {
				if hx_value_767 == nil {
					var hx_zero_768 int
					return hx_zero_768
				}
				return hx_value_767.(int)
			}(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_5(value any) *hxrt.Function[func(any) any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any) any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_771 any) *string {
				if hx_value_771 == nil {
					var hx_zero_772 *string
					return hx_zero_772
				}
				return hx_value_771.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			hxrt.FunctionCall(source)(func(hx_value_773 any) int {
				if hx_value_773 == nil {
					var hx_zero_774 int
					return hx_zero_774
				}
				return hx_value_773.(int)
			}(arg_0))
			return nil
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			hxrt.FunctionCall(source)(arg_0)
			return nil
		})
	case *hxrt.Function[func(*hxrt.Array)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			hxrt.FunctionCall(source)(func(hx_value_775 any) *hxrt.Array {
				if hx_value_775 == nil {
					var hx_zero_776 *hxrt.Array
					return hx_zero_776
				}
				return hx_value_775.(*hxrt.Array)
			}(arg_0))
			return nil
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
			return nil
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			hxrt.FunctionCall(source)(func(hx_value_777 any) *string {
				if hx_value_777 == nil {
					var hx_zero_778 *string
					return hx_zero_778
				}
				return hx_value_777.(*string)
			}(arg_0))
			return nil
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_779 any) int {
				if hx_value_779 == nil {
					var hx_zero_780 int
					return hx_zero_780
				}
				return hx_value_779.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_781 any) *string {
				if hx_value_781 == nil {
					var hx_zero_782 *string
					return hx_zero_782
				}
				return hx_value_781.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(int) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_783 any) int {
				if hx_value_783 == nil {
					var hx_zero_784 int
					return hx_zero_784
				}
				return hx_value_783.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_785 any) int {
				if hx_value_785 == nil {
					var hx_zero_786 int
					return hx_zero_786
				}
				return hx_value_785.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(any) any]:
		return source
	case *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(hx_recover_callable_4(arg_0))
		})
	case *hxrt.Function[func(*string) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_787 any) *string {
				if hx_value_787 == nil {
					var hx_zero_788 *string
					return hx_zero_788
				}
				return hx_value_787.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(hx_recover_callable_5(arg_0))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_789 any) *string {
				if hx_value_789 == nil {
					var hx_zero_790 *string
					return hx_zero_790
				}
				return hx_value_789.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_791 any) int {
				if hx_value_791 == nil {
					var hx_zero_792 int
					return hx_zero_792
				}
				return hx_value_791.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_793 any) int {
				if hx_value_793 == nil {
					var hx_zero_794 int
					return hx_zero_794
				}
				return hx_value_793.(int)
			}(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *SerializationSnapshotChild {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *SerializationSnapshotChild:
		return source
	default:
		return value.(*SerializationSnapshotChild)
	}
}

func hx_recover_class_1(value any) *haxe__ds__StringMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return source
	default:
		return value.(*haxe__ds__StringMap)
	}
}

func hx_recover_class_2(value any) *haxe__ds__ObjectMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__ObjectMap:
		return source
	default:
		return value.(*haxe__ds__ObjectMap)
	}
}

func hx_recover_class_3(value any) *haxe__ds___List__GoListIterator {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds___List__GoListIterator:
		return source
	default:
		return value.(*haxe__ds___List__GoListIterator)
	}
}

func hx_recover_class_4(value any) *haxe__ds__IntMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__IntMap:
		return source
	default:
		return value.(*haxe__ds__IntMap)
	}
}

func hx_recover_class_5(value any) *Date {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Date:
		return source
	default:
		return value.(*Date)
	}
}

func hx_recover_class_6(value any) *haxe__ds__List {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__List:
		return source
	default:
		return value.(*haxe__ds__List)
	}
}

func hx_recover_class_7(value any) *haxe__io__Bytes {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return source
	default:
		return value.(*haxe__io__Bytes)
	}
}

func hx_recover_class_8(value any) *haxe__ValueException {
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

func hx_recover_class_9(value any) *haxe__Exception {
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

func hx_recover_class_10(value any) *haxe___Int64_____Int64 {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return source
	default:
		return value.(*haxe___Int64_____Int64)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
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
