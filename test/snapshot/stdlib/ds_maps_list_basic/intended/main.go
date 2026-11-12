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

func (self *Box) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Box(id int) *Box {
	return hx_init_Box(hx_alloc_Box(nil), id)
}

func hx_alloc_Box(owner any) *Box {
	self := &Box{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Box)
	}
	return self
}

func hx_init_Box(self *Box, id int) *Box {
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
	switch value := hxrt.ObjectSource(obj).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
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
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
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

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_177 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_177.__hx_this.details()
			}, func() any {
				return hx_receiver_177.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_178 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_178.__hx_this.get_message()
			}, func() any {
				return hx_receiver_178.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_179 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_179.__hx_this.get_native()
			}, func() any {
				return hx_receiver_179.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_180 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_180.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_180.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_181 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_181.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_181.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_182 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_182.__hx_this.toString()
			}, func() any {
				return hx_receiver_182.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_183 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_183.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_183.__hx_this
			}, "unwrap")
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
			hx_receiver_184 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_184.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_184.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__ds__EnumValueMap(value *haxe__ds__EnumValueMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "balance":
		return func() *hxrt.Function[func(*haxe__ds___EnumValueMap__EnumValueTreeNode, any, any, *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode] {
			hx_receiver_185 := value
			return hxrt.BoundFunction(func(left *haxe__ds___EnumValueMap__EnumValueTreeNode, key any, value any, right *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
				return hx_receiver_185.__hx_this.balance(left, key, value, right)
			}, func() any {
				return hx_receiver_185.__hx_this
			}, "balance")
		}()
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_186 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_186.__hx_this.clear()
			}, func() any {
				return hx_receiver_186.__hx_this
			}, "clear")
		}()
	case "compare":
		return func() *hxrt.Function[func(any, any) int] {
			hx_receiver_187 := value
			return hxrt.BoundFunction(func(left any, right any) int {
				return hx_receiver_187.__hx_this.compare(left, right)
			}, func() any {
				return hx_receiver_187.__hx_this
			}, "compare")
		}()
	case "compareArg":
		return func() *hxrt.Function[func(any, any) int] {
			hx_receiver_188 := value
			return hxrt.BoundFunction(func(left any, right any) int {
				return hx_receiver_188.__hx_this.compareArg(left, right)
			}, func() any {
				return hx_receiver_188.__hx_this
			}, "compareArg")
		}()
	case "compareArgs":
		return func() *hxrt.Function[func(*hxrt.Array, *hxrt.Array) int] {
			hx_receiver_189 := value
			return hxrt.BoundFunction(func(left *hxrt.Array, right *hxrt.Array) int {
				return hx_receiver_189.__hx_this.compareArgs(left, right)
			}, func() any {
				return hx_receiver_189.__hx_this
			}, "compareArgs")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__EnumValueMap] {
			hx_receiver_190 := value
			return hxrt.BoundFunction(func() *haxe__ds__EnumValueMap {
				return hx_receiver_190.__hx_this.copy()
			}, func() any {
				return hx_receiver_190.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_191 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_191.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_191.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_192 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_192.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_192.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_193 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_193.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_193.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_194 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_194.__hx_this.get(key)
			}, func() any {
				return hx_receiver_194.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_195 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_195.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_195.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_196 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_196.__hx_this.iterator()
			}, func() any {
				return hx_receiver_196.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_197 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_197.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_197.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_198 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_198.__hx_this.keys()
			}, func() any {
				return hx_receiver_198.__hx_this
			}, "keys")
		}()
	case "merge":
		return func() *hxrt.Function[func(*haxe__ds___EnumValueMap__EnumValueTreeNode, *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode] {
			hx_receiver_199 := value
			return hxrt.BoundFunction(func(left *haxe__ds___EnumValueMap__EnumValueTreeNode, right *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
				return hx_receiver_199.__hx_this.merge(left, right)
			}, func() any {
				return hx_receiver_199.__hx_this
			}, "merge")
		}()
	case "minBinding":
		return func() *hxrt.Function[func(*haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode] {
			hx_receiver_200 := value
			return hxrt.BoundFunction(func(node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
				return hx_receiver_200.__hx_this.minBinding(node)
			}, func() any {
				return hx_receiver_200.__hx_this
			}, "minBinding")
		}()
	case "remove":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_201 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_201.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_201.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_202 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_202.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_202.__hx_this
			}, "removeIMap")
		}()
	case "removeLoop":
		return func() *hxrt.Function[func(any, *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode] {
			hx_receiver_203 := value
			return hxrt.BoundFunction(func(key any, node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
				return hx_receiver_203.__hx_this.removeLoop(key, node)
			}, func() any {
				return hx_receiver_203.__hx_this
			}, "removeLoop")
		}()
	case "removeMinBinding":
		return func() *hxrt.Function[func(*haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode] {
			hx_receiver_204 := value
			return hxrt.BoundFunction(func(node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
				return hx_receiver_204.__hx_this.removeMinBinding(node)
			}, func() any {
				return hx_receiver_204.__hx_this
			}, "removeMinBinding")
		}()
	case "set":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_205 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_205.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_205.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_206 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_206.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_206.__hx_this
			}, "setIMap")
		}()
	case "setLoop":
		return func() *hxrt.Function[func(any, any, *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode] {
			hx_receiver_207 := value
			return hxrt.BoundFunction(func(key any, value any, node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
				return hx_receiver_207.__hx_this.setLoop(key, value, node)
			}, func() any {
				return hx_receiver_207.__hx_this
			}, "setLoop")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_208 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_208.__hx_this.toString()
			}, func() any {
				return hx_receiver_208.__hx_this
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
			hx_receiver_209 := value
			return hxrt.BoundFunction(func(item any) {
				hx_receiver_209.__hx_this.add(item)
			}, func() any {
				return hx_receiver_209.__hx_this
			}, "add")
		}()
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_210 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_210.__hx_this.clear()
			}, func() any {
				return hx_receiver_210.__hx_this
			}, "clear")
		}()
	case "filter":
		return func() *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List] {
			hx_receiver_211 := value
			return hxrt.BoundFunction(func(predicate *hxrt.Function[func(any) bool]) *haxe__ds__List {
				return hx_receiver_211.__hx_this.filter(predicate)
			}, func() any {
				return hx_receiver_211.__hx_this
			}, "filter")
		}()
	case "first":
		return func() *hxrt.Function[func() any] {
			hx_receiver_212 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_212.__hx_this.first()
			}, func() any {
				return hx_receiver_212.__hx_this
			}, "first")
		}()
	case "isEmpty":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_213 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_213.__hx_this.isEmpty()
			}, func() any {
				return hx_receiver_213.__hx_this
			}, "isEmpty")
		}()
	case "iterator":
		return func() *hxrt.Function[func() *haxe__ds___List__GoListIterator] {
			hx_receiver_214 := value
			return hxrt.BoundFunction(func() *haxe__ds___List__GoListIterator {
				return hx_receiver_214.__hx_this.iterator()
			}, func() any {
				return hx_receiver_214.__hx_this
			}, "iterator")
		}()
	case "join":
		return func() *hxrt.Function[func(*string) *string] {
			hx_receiver_215 := value
			return hxrt.BoundFunction(func(separator *string) *string {
				return hx_receiver_215.__hx_this.join(separator)
			}, func() any {
				return hx_receiver_215.__hx_this
			}, "join")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator] {
			hx_receiver_216 := value
			return hxrt.BoundFunction(func() *haxe__ds___List__GoListKeyValueIterator {
				return hx_receiver_216.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_216.__hx_this
			}, "keyValueIterator")
		}()
	case "last":
		return func() *hxrt.Function[func() any] {
			hx_receiver_217 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_217.__hx_this.last()
			}, func() any {
				return hx_receiver_217.__hx_this
			}, "last")
		}()
	case "map":
		return func() *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List] {
			hx_receiver_218 := value
			return hxrt.BoundFunction(func(transform *hxrt.Function[func(any) any]) *haxe__ds__List {
				return hx_receiver_218.__hx_this.map_(transform)
			}, func() any {
				return hx_receiver_218.__hx_this
			}, "map")
		}()
	case "pop":
		return func() *hxrt.Function[func() any] {
			hx_receiver_219 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_219.__hx_this.pop()
			}, func() any {
				return hx_receiver_219.__hx_this
			}, "pop")
		}()
	case "push":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_220 := value
			return hxrt.BoundFunction(func(item any) {
				hx_receiver_220.__hx_this.push(item)
			}, func() any {
				return hx_receiver_220.__hx_this
			}, "push")
		}()
	case "remove":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_221 := value
			return hxrt.BoundFunction(func(value any) bool {
				return hx_receiver_221.__hx_this.remove(value)
			}, func() any {
				return hx_receiver_221.__hx_this
			}, "remove")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_222 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_222.__hx_this.toString()
			}, func() any {
				return hx_receiver_222.__hx_this
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
			hx_receiver_223 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_223.__hx_this.clear()
			}, func() any {
				return hx_receiver_223.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__ObjectMap] {
			hx_receiver_224 := value
			return hxrt.BoundFunction(func() *haxe__ds__ObjectMap {
				return hx_receiver_224.__hx_this.copy()
			}, func() any {
				return hx_receiver_224.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_225 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_225.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_225.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_226 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_226.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_226.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_227 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_227.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_227.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_228 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_228.__hx_this.get(key)
			}, func() any {
				return hx_receiver_228.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_229 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_229.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_229.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_230 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_230.__hx_this.iterator()
			}, func() any {
				return hx_receiver_230.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_231 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_231.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_231.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_232 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_232.__hx_this.keys()
			}, func() any {
				return hx_receiver_232.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_233 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_233.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_233.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_234 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_234.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_234.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_235 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_235.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_235.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_236 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_236.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_236.__hx_this
			}, "setIMap")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_237 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_237.__hx_this.toString()
			}, func() any {
				return hx_receiver_237.__hx_this
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
			hx_receiver_238 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_238.__hx_this.clear()
			}, func() any {
				return hx_receiver_238.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__StringMap] {
			hx_receiver_239 := value
			return hxrt.BoundFunction(func() *haxe__ds__StringMap {
				return hx_receiver_239.__hx_this.copy()
			}, func() any {
				return hx_receiver_239.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_240 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_240.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_240.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_241 := value
			return hxrt.BoundFunction(func(key *string) bool {
				return hx_receiver_241.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_241.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_242 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_242.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_242.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_243 := value
			return hxrt.BoundFunction(func(key *string) any {
				return hx_receiver_243.__hx_this.get(key)
			}, func() any {
				return hx_receiver_243.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_244 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_244.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_244.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_245 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_245.__hx_this.iterator()
			}, func() any {
				return hx_receiver_245.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_246 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_246.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_246.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_247 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_247.__hx_this.keys()
			}, func() any {
				return hx_receiver_247.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_248 := value
			return hxrt.BoundFunction(func(key *string) bool {
				return hx_receiver_248.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_248.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_249 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_249.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_249.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(*string, any)] {
			hx_receiver_250 := value
			return hxrt.BoundFunction(func(key *string, value any) {
				hx_receiver_250.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_250.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_251 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_251.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_251.__hx_this
			}, "setIMap")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_252 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_252.__hx_this.toString()
			}, func() any {
				return hx_receiver_252.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds___EnumValueMap__EnumValueTreeNode(value *haxe__ds___EnumValueMap__EnumValueTreeNode, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "getHeight":
		return func() *hxrt.Function[func() int] {
			hx_receiver_253 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_253.__hx_this.getHeight()
			}, func() any {
				return hx_receiver_253.__hx_this
			}, "getHeight")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_254 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_254.__hx_this.toString()
			}, func() any {
				return hx_receiver_254.__hx_this
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
			hx_receiver_255 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_255.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_255.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() any] {
			hx_receiver_256 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_256.__hx_this.next()
			}, func() any {
				return hx_receiver_256.__hx_this
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
			hx_receiver_257 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_257.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_257.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_258 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_258.__hx_this.next()
			}, func() any {
				return hx_receiver_258.__hx_this
			}, "next")
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
			hx_receiver_259 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_259.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_259.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_260 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_260.__hx_this.next()
			}, func() any {
				return hx_receiver_260.__hx_this
			}, "next")
		}()
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

func hxrt_typeCreateEmpty__haxe__Exception() *haxe__Exception {
	instance := &haxe__Exception{}
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
	case "haxe.Exception":
		return hxrt_typeCallAny(New_haxe__Exception, args)
	case "haxe.Int64Helper":
		return nil, false
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
	case "haxe.Exception":
		return hxrt_typeCreateEmpty__haxe__Exception(), true
	case "haxe.ValueException":
		return hxrt_typeCreateEmpty__haxe__ValueException(), true
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
	case *Box:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Box")}
	case *haxe__Exception:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
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
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.StackItem")}
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
	case "haxe.Exception":
		return nil
	case "haxe.Int64Helper":
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
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("catchValue"), hxrt.StringFromLiteral("caught"), hxrt.StringFromLiteral("thrown"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
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
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
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
	case "haxe.Exception":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
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
	case "haxe.StackItem":
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
	case *haxe__StackItem:
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
	case "haxe.StackItem":
		return hxrt.NewArray(hxrt.StringFromLiteral("CFunction"), hxrt.StringFromLiteral("Module"), hxrt.StringFromLiteral("FilePos"), hxrt.StringFromLiteral("Method"), hxrt.StringFromLiteral("LocalFunction"))
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
	case *haxe__StackItem:
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
	case "haxe.StackItem":
		return hxrt.NewArray(haxe__StackItem_CFunction)
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
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
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
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
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
	switch value := hxrt.ObjectSource(object).(type) {
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
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
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
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
	switch value := hxrt.ObjectSource(object).(type) {
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
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
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
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
	switch value := hxrt.ObjectSource(object).(type) {
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
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
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
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
		return hxrt.ReadStoredField(value.id)
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

func hxrt__generated_field_lookup__haxe__ds__EnumValueMap(value *haxe__ds__EnumValueMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "root":
		return hxrt.ReadStoredField(value.root)
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

func hxrt__generated_field_lookup__haxe__ds___EnumValueMap__EnumValueTreeNode(value *haxe__ds___EnumValueMap__EnumValueTreeNode, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "height":
		return hxrt.ReadStoredField(value.height)
	case "key":
		return hxrt.ReadStoredField(value.key)
	case "left":
		return hxrt.ReadStoredField(value.left)
	case "right":
		return hxrt.ReadStoredField(value.right)
	case "value":
		return hxrt.ReadStoredField(value.value)
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
	case *haxe__StackItem:
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
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*haxe__ds___EnumValueMap__EnumValueTreeNode, any, any, *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode]:
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
		return hxrt.FunctionCall(source)(hx_recover_class_3(argument_0), argument_1, argument_2, hx_recover_class_3(argument_3))
	case *hxrt.Function[func()]:
		hxrt.FunctionCall(source)()
		return nil
	case *hxrt.Function[func(any, any) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1)
	case *hxrt.Function[func(*hxrt.Array, *hxrt.Array) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_261 any) *hxrt.Array {
			if hx_value_261 == nil {
				var hx_zero_262 *hxrt.Array
				return hx_zero_262
			}
			return hx_value_261.(*hxrt.Array)
		}(argument_0), func(hx_value_263 any) *hxrt.Array {
			if hx_value_263 == nil {
				var hx_zero_264 *hxrt.Array
				return hx_zero_264
			}
			return hx_value_263.(*hxrt.Array)
		}(argument_1))
	case *hxrt.Function[func() *haxe__ds__EnumValueMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(any) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(any) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*haxe__ds___EnumValueMap__EnumValueTreeNode, *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_3(argument_0), hx_recover_class_3(argument_1))
	case *hxrt.Function[func(*haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_3(argument_0))
	case *hxrt.Function[func(any, *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, hx_recover_class_3(argument_1))
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
	case *hxrt.Function[func(any, any, *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode]:
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
		return hxrt.FunctionCall(source)(argument_0, argument_1, hx_recover_class_3(argument_2))
	case *hxrt.Function[func(any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(argument_0)
		return nil
	case *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_3(argument_0))
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_265 any) *string {
			if hx_value_265 == nil {
				var hx_zero_266 *string
				return hx_zero_266
			}
			return hx_value_265.(*string)
		}(argument_0))
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_4(argument_0))
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_267 any) *string {
			if hx_value_267 == nil {
				var hx_zero_268 *string
				return hx_zero_268
			}
			return hx_value_267.(*string)
		}(argument_0))
	case *hxrt.Function[func(*string) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_269 any) *string {
			if hx_value_269 == nil {
				var hx_zero_270 *string
				return hx_zero_270
			}
			return hx_value_269.(*string)
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
		hxrt.FunctionCall(source)(func(hx_value_271 any) *string {
			if hx_value_271 == nil {
				var hx_zero_272 *string
				return hx_zero_272
			}
			return hx_value_271.(*string)
		}(argument_0), argument_1)
		return nil
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
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
			return func(hx_value_273 any) bool {
				if hx_value_273 == nil {
					var hx_zero_274 bool
					return hx_zero_274
				}
				return hx_value_273.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_275 any) bool {
				if hx_value_275 == nil {
					var hx_zero_276 bool
					return hx_zero_276
				}
				return hx_value_275.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_277 any) bool {
				if hx_value_277 == nil {
					var hx_zero_278 bool
					return hx_zero_278
				}
				return hx_value_277.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_279 any) bool {
				if hx_value_279 == nil {
					var hx_zero_280 bool
					return hx_zero_280
				}
				return hx_value_279.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_281 any) bool {
				if hx_value_281 == nil {
					var hx_zero_282 bool
					return hx_zero_282
				}
				return hx_value_281.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_283 any) bool {
				if hx_value_283 == nil {
					var hx_zero_284 bool
					return hx_zero_284
				}
				return hx_value_283.(bool)
			}(nil)
		})
	case *hxrt.Function[func() *haxe__ds__EnumValueMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_285 any) bool {
				if hx_value_285 == nil {
					var hx_zero_286 bool
					return hx_zero_286
				}
				return hx_value_285.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_287 any) bool {
				if hx_value_287 == nil {
					var hx_zero_288 bool
					return hx_zero_288
				}
				return hx_value_287.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_289 any) bool {
				if hx_value_289 == nil {
					var hx_zero_290 bool
					return hx_zero_290
				}
				return hx_value_289.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_291 any) bool {
				if hx_value_291 == nil {
					var hx_zero_292 bool
					return hx_zero_292
				}
				return hx_value_291.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_293 any) bool {
				if hx_value_293 == nil {
					var hx_zero_294 bool
					return hx_zero_294
				}
				return hx_value_293.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_295 any) bool {
				if hx_value_295 == nil {
					var hx_zero_296 bool
					return hx_zero_296
				}
				return hx_value_295.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_297 any) bool {
				if hx_value_297 == nil {
					var hx_zero_298 bool
					return hx_zero_298
				}
				return hx_value_297.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func(...any) any]:
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
	case *hxrt.Function[func() *haxe__ds__EnumValueMap]:
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
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() *string] {
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
			return func(hx_value_299 any) *string {
				if hx_value_299 == nil {
					var hx_zero_300 *string
					return hx_zero_300
				}
				return hx_value_299.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_301 any) *string {
				if hx_value_301 == nil {
					var hx_zero_302 *string
					return hx_zero_302
				}
				return hx_value_301.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_303 any) *string {
				if hx_value_303 == nil {
					var hx_zero_304 *string
					return hx_zero_304
				}
				return hx_value_303.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_305 any) *string {
				if hx_value_305 == nil {
					var hx_zero_306 *string
					return hx_zero_306
				}
				return hx_value_305.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_307 any) *string {
				if hx_value_307 == nil {
					var hx_zero_308 *string
					return hx_zero_308
				}
				return hx_value_307.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_309 any) *string {
				if hx_value_309 == nil {
					var hx_zero_310 *string
					return hx_zero_310
				}
				return hx_value_309.(*string)
			}(nil)
		})
	case *hxrt.Function[func() *haxe__ds__EnumValueMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_311 any) *string {
				if hx_value_311 == nil {
					var hx_zero_312 *string
					return hx_zero_312
				}
				return hx_value_311.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_313 any) *string {
				if hx_value_313 == nil {
					var hx_zero_314 *string
					return hx_zero_314
				}
				return hx_value_313.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_315 any) *string {
				if hx_value_315 == nil {
					var hx_zero_316 *string
					return hx_zero_316
				}
				return hx_value_315.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_317 any) *string {
				if hx_value_317 == nil {
					var hx_zero_318 *string
					return hx_zero_318
				}
				return hx_value_317.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_319 any) *string {
				if hx_value_319 == nil {
					var hx_zero_320 *string
					return hx_zero_320
				}
				return hx_value_319.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_321 any) *string {
				if hx_value_321 == nil {
					var hx_zero_322 *string
					return hx_zero_322
				}
				return hx_value_321.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_323 any) *string {
				if hx_value_323 == nil {
					var hx_zero_324 *string
					return hx_zero_324
				}
				return hx_value_323.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_3(value any) *hxrt.Function[func(any) bool] {
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
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_325 any) bool {
				if hx_value_325 == nil {
					var hx_zero_326 bool
					return hx_zero_326
				}
				return hx_value_325.(bool)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(any) bool]:
		return source
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_327 any) bool {
				if hx_value_327 == nil {
					var hx_zero_328 bool
					return hx_zero_328
				}
				return hx_value_327.(bool)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(*haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_329 any) bool {
				if hx_value_329 == nil {
					var hx_zero_330 bool
					return hx_zero_330
				}
				return hx_value_329.(bool)
			}(hxrt.FunctionCall(source)(hx_recover_class_3(arg_0)))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			hxrt.FunctionCall(source)(arg_0)
			return func(hx_value_331 any) bool {
				if hx_value_331 == nil {
					var hx_zero_332 bool
					return hx_zero_332
				}
				return hx_value_331.(bool)
			}(nil)
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_333 any) bool {
				if hx_value_333 == nil {
					var hx_zero_334 bool
					return hx_zero_334
				}
				return hx_value_333.(bool)
			}(hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0)))
		})
	case *hxrt.Function[func(*string) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_337 any) bool {
				if hx_value_337 == nil {
					var hx_zero_338 bool
					return hx_zero_338
				}
				return hx_value_337.(bool)
			}(hxrt.FunctionCall(source)(func(hx_value_335 any) *string {
				if hx_value_335 == nil {
					var hx_zero_336 *string
					return hx_zero_336
				}
				return hx_value_335.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_339 any) bool {
				if hx_value_339 == nil {
					var hx_zero_340 bool
					return hx_zero_340
				}
				return hx_value_339.(bool)
			}(hxrt.FunctionCall(source)(hx_recover_callable_4(arg_0)))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return hxrt.FunctionCall(source)(func(hx_value_341 any) *string {
				if hx_value_341 == nil {
					var hx_zero_342 *string
					return hx_zero_342
				}
				return hx_value_341.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_345 any) bool {
				if hx_value_345 == nil {
					var hx_zero_346 bool
					return hx_zero_346
				}
				return hx_value_345.(bool)
			}(hxrt.FunctionCall(source)(func(hx_value_343 any) *string {
				if hx_value_343 == nil {
					var hx_zero_344 *string
					return hx_zero_344
				}
				return hx_value_343.(*string)
			}(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_4(value any) *hxrt.Function[func(any) any] {
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
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) any]:
		return source
	case *hxrt.Function[func(*haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(hx_recover_class_3(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			hxrt.FunctionCall(source)(arg_0)
			return nil
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))
		})
	case *hxrt.Function[func(*string) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_347 any) *string {
				if hx_value_347 == nil {
					var hx_zero_348 *string
					return hx_zero_348
				}
				return hx_value_347.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(hx_recover_callable_4(arg_0))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_349 any) *string {
				if hx_value_349 == nil {
					var hx_zero_350 *string
					return hx_zero_350
				}
				return hx_value_349.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_351 any) *string {
				if hx_value_351 == nil {
					var hx_zero_352 *string
					return hx_zero_352
				}
				return hx_value_351.(*string)
			}(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__ds__StringMap {
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

func hx_recover_class_1(value any) *haxe__ds__ObjectMap {
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

func hx_recover_class_2(value any) *haxe__ds___List__GoListIterator {
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

func hx_recover_class_3(value any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds___EnumValueMap__EnumValueTreeNode:
		return source
	default:
		return value.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}
}

func hx_recover_class_4(value any) *haxe__ds__EnumValueMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__EnumValueMap:
		return source
	default:
		return value.(*haxe__ds__EnumValueMap)
	}
}

func hx_recover_class_5(value any) *haxe__ValueException {
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

func hx_recover_class_6(value any) *haxe__Exception {
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
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
