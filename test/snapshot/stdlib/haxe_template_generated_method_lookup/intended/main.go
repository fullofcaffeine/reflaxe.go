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

type I_ConcreteIterable interface {
	iterator() *SpecializedIterator
}

type ConcreteIterable struct {
	__hx_this I_ConcreteIterable
	values    *hxrt.Array
}

func New_ConcreteIterable(values *hxrt.Array) *ConcreteIterable {
	self := &ConcreteIterable{}
	self.__hx_this = self
	self.values = values
	return self
}

func (self *ConcreteIterable) iterator() *SpecializedIterator {
	return New_SpecializedIterator(self.values)
}

type I_ConcreteIterator interface {
	hasNext() bool
	next() *string
}

type ConcreteIterator struct {
	__hx_this I_ConcreteIterator
	values    *hxrt.Array
	index     int
}

func New_ConcreteIterator(values *hxrt.Array) *ConcreteIterator {
	self := &ConcreteIterator{}
	self.__hx_this = self
	self.values = values
	self.index = 0
	return self
}

func (self *ConcreteIterator) hasNext() bool {
	return (self.index < self.values.Len())
}

func (self *ConcreteIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("base:"), self.values.Get(func() int {
		hx_post_4 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_4
	}()))
}

func call0(obj any, key *string) *string {
	template := New_haxe__Template(hxrt.StringFromLiteral("$$invoke(dummy)"))
	return template.__hx_this.execute(func() map[string]any {
		hx_obj_5 := map[string]any{}
		hx_obj_5["dummy"] = nil
		return hx_obj_5
	}(), func() map[string]any {
		hx_obj_6 := map[string]any{}
		hx_obj_6["invoke"] = Reflect_field(obj, key)
		return hx_obj_6
	}())
}

func call1(obj any, key *string, value any) *string {
	template := New_haxe__Template(hxrt.StringFromLiteral("$$invoke(value)"))
	return template.__hx_this.execute(func() map[string]any {
		hx_obj_7 := map[string]any{}
		hx_obj_7["value"] = value
		return hx_obj_7
	}(), func() map[string]any {
		hx_obj_8 := map[string]any{}
		hx_obj_8["invoke"] = Reflect_field(obj, key)
		return hx_obj_8
	}())
}

func main() {
	template := New_haxe__Template(hxrt.StringFromLiteral("::foreach items::::__current__::;::end::"))
	var v any = any(template.__hx_this.execute(func() map[string]any {
		hx_obj_9 := map[string]any{}
		hx_obj_9["items"] = New_ConcreteIterable(hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b")))
		return hx_obj_9
	}(), nil))
	hxrt.Println(v)
	var v_1 any = any(template.__hx_this.execute(func() map[string]any {
		hx_obj_10 := map[string]any{}
		hx_obj_10["items"] = New_ConcreteIterator(hxrt.NewArray(hxrt.StringFromLiteral("x"), hxrt.StringFromLiteral("y")))
		return hx_obj_10
	}(), nil))
	hxrt.Println(v_1)
	leaf := New_MethodLeaf()
	middle := leaf.MethodMiddle
	base := leaf.MethodMiddle.MethodBase
	var named NameContract = leaf
	var v_2 any = any(Reflect_hasField(leaf, hxrt.StringFromLiteral("leafOnly")))
	hxrt.Println(v_2)
	var v_3 any = any(!hxrt.AnyEqualsNull(Reflect_field(leaf, hxrt.StringFromLiteral("leafOnly"))))
	hxrt.Println(v_3)
	var v_4 any = any(call0(leaf, hxrt.StringFromLiteral("leafOnly")))
	hxrt.Println(v_4)
	var v_5 any = any(Reflect_hasField(base, hxrt.StringFromLiteral("leafOnly")))
	hxrt.Println(v_5)
	var v_6 any = any(!hxrt.AnyEqualsNull(Reflect_field(base, hxrt.StringFromLiteral("leafOnly"))))
	hxrt.Println(v_6)
	var v_7 any = any(call0(base, hxrt.StringFromLiteral("leafOnly")))
	hxrt.Println(v_7)
	var v_8 any = any(call0(middle, hxrt.StringFromLiteral("middleOnly")))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(call0(leaf, hxrt.StringFromLiteral("macroName")), hxrt.StringFromLiteral(":")), call0(middle, hxrt.StringFromLiteral("macroName"))), hxrt.StringFromLiteral(":")), call0(base, hxrt.StringFromLiteral("macroName"))), hxrt.StringFromLiteral(":")), call0(named, hxrt.StringFromLiteral("macroName"))))
	hxrt.Println(v_9)
	var v_10 any = any(call1(base, hxrt.StringFromLiteral("describe"), hxrt.StringFromLiteral("bound")))
	hxrt.Println(v_10)
	var v_11 any = any(call0(base, hxrt.StringFromLiteral("type")))
	hxrt.Println(v_11)
	var v_12 any = any(leaf.retainSecret())
	hxrt.Println(v_12)
	var v_13 any = any(call0(base, hxrt.StringFromLiteral("secret")))
	hxrt.Println(v_13)
	var v_14 any = any(call1(base, hxrt.StringFromLiteral("bump"), 2))
	hxrt.Println(v_14)
	var v_15 any = any(call1(base, hxrt.StringFromLiteral("bump"), 3))
	hxrt.Println(v_15)
	var v_16 any = any(Reflect_hasField(base, hxrt.StringFromLiteral("missing")))
	hxrt.Println(v_16)
	var v_17 any = any(hxrt.StdString(Reflect_field(base, hxrt.StringFromLiteral("missing"))))
	hxrt.Println(v_17)
	var absent *MethodBase = nil
	var v_18 any = any(Reflect_hasField(absent, hxrt.StringFromLiteral("virtualName")))
	hxrt.Println(v_18)
	var v_19 any = any(hxrt.StdString(Reflect_field(absent, hxrt.StringFromLiteral("virtualName"))))
	hxrt.Println(v_19)
}

type I_MethodBase interface {
	virtualName() *string
	macroName(resolve func(*string) any, ignored any) *string
	describe(resolve func(*string) any, key *string) *string
	bump(resolve func(*string) any, key *string) int
	type_(resolve func(*string) any, ignored any) *string
	secret(resolve func(*string) any, ignored any) *string
	retainSecret() *string
}

type MethodBase struct {
	__hx_this I_MethodBase
	total     int
}

func New_MethodBase() *MethodBase {
	self := &MethodBase{}
	self.__hx_this = self
	self.total = 0
	return self
}

func (self *MethodBase) virtualName() *string {
	return hxrt.StringFromLiteral("base")
}

func (self *MethodBase) macroName(resolve func(*string) any, ignored any) *string {
	return self.__hx_this.virtualName()
}

func (self *MethodBase) describe(resolve func(*string) any, key *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(resolve(key)), hxrt.StringFromLiteral(":")), self.__hx_this.virtualName())
}

func (self *MethodBase) bump(resolve func(*string) any, key *string) int {
	amount := hxrt.IntFromNullableAny(resolve(key))
	self.total = int(int32((hxrt.Int32Wrap(self.total) + hxrt.Int32Wrap(amount))))
	return self.total
}

func (self *MethodBase) type_(resolve func(*string) any, ignored any) *string {
	return hxrt.StringFromLiteral("keyword")
}

func (self *MethodBase) secret(resolve func(*string) any, ignored any) *string {
	return hxrt.StringFromLiteral("secret")
}

func (self *MethodBase) retainSecret() *string {
	return self.__hx_this.secret(nil, nil)
}

type I_MethodLeaf interface {
	virtualName() *string
	macroName(resolve func(*string) any, ignored any) *string
	describe(resolve func(*string) any, key *string) *string
	bump(resolve func(*string) any, key *string) int
	type_(resolve func(*string) any, ignored any) *string
	secret(resolve func(*string) any, ignored any) *string
	retainSecret() *string
	middleOnly(resolve func(*string) any, ignored any) *string
	leafOnly(resolve func(*string) any, ignored any) *string
}

type MethodLeaf struct {
	*MethodMiddle
	__hx_this I_MethodLeaf
}

func New_MethodLeaf() *MethodLeaf {
	self := &MethodLeaf{}
	self.MethodMiddle = New_MethodMiddle()
	self.MethodMiddle.MethodBase.__hx_this = self
	self.MethodMiddle.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *MethodLeaf) virtualName() *string {
	return hxrt.StringFromLiteral("leaf")
}

func (self *MethodLeaf) leafOnly(resolve func(*string) any, ignored any) *string {
	return hxrt.StringFromLiteral("leaf-only")
}

type I_MethodMiddle interface {
	virtualName() *string
	macroName(resolve func(*string) any, ignored any) *string
	describe(resolve func(*string) any, key *string) *string
	bump(resolve func(*string) any, key *string) int
	type_(resolve func(*string) any, ignored any) *string
	secret(resolve func(*string) any, ignored any) *string
	retainSecret() *string
	middleOnly(resolve func(*string) any, ignored any) *string
}

type MethodMiddle struct {
	*MethodBase
	__hx_this I_MethodMiddle
}

func New_MethodMiddle() *MethodMiddle {
	self := &MethodMiddle{}
	self.MethodBase = New_MethodBase()
	self.MethodBase.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *MethodMiddle) middleOnly(resolve func(*string) any, ignored any) *string {
	return hxrt.StringFromLiteral("middle")
}

type NameContract interface {
	macroName(resolve func(*string) any, ignored any) *string
}

type I_SpecializedIterator interface {
	hasNext() bool
	next() *string
}

type SpecializedIterator struct {
	*ConcreteIterator
	__hx_this I_SpecializedIterator
}

func New_SpecializedIterator(values *hxrt.Array) *SpecializedIterator {
	self := &SpecializedIterator{}
	self.ConcreteIterator = New_ConcreteIterator(values)
	self.ConcreteIterator.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *SpecializedIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("special:"), self.values.Get(func() int {
		hx_post_14 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_14
	}()))
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := obj.(type) {
	case *ConcreteIterable:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *ConcreteIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *MethodBase:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *MethodLeaf:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *MethodMiddle:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *SpecializedIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Template:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
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
	case *ConcreteIterable:
		return hxrt__generated_method_field__ConcreteIterable(value, key)
	case *ConcreteIterator:
		return hxrt__generated_method_field__ConcreteIterator(value, key)
	case *EReg:
		return hxrt__generated_method_field__EReg(value, key)
	case *MethodBase:
		return hxrt__generated_method_field__MethodBase(value, key)
	case *MethodLeaf:
		return hxrt__generated_method_field__MethodLeaf(value, key)
	case *MethodMiddle:
		return hxrt__generated_method_field__MethodMiddle(value, key)
	case *SpecializedIterator:
		return hxrt__generated_method_field__SpecializedIterator(value, key)
	case *haxe__Template:
		return hxrt__generated_method_field__haxe__Template(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_method_field__haxe__io__Bytes(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_method_field__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__ConcreteIterable(value *ConcreteIterable, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "iterator":
		return value.iterator
	}
	return nil
}

func hxrt__generated_method_field__ConcreteIterator(value *ConcreteIterator, key string) any {
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

func hxrt__generated_method_field__EReg(value *EReg, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "expandReplacement":
		return value.expandReplacement
	case "map":
		return value.map_
	case "match":
		return value.match
	case "matchSub":
		return value.matchSub
	case "matched":
		return value.matched
	case "matchedLeft":
		return value.matchedLeft
	case "matchedPos":
		return value.matchedPos
	case "matchedRight":
		return value.matchedRight
	case "remember":
		return value.remember
	case "replace":
		return value.replace
	case "requireMatch":
		return value.requireMatch
	case "split":
		return value.split
	}
	return nil
}

func hxrt__generated_method_field__MethodBase(value *MethodBase, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "bump":
		return value.bump
	case "describe":
		return value.describe
	case "macroName":
		return value.macroName
	case "retainSecret":
		return value.retainSecret
	case "secret":
		return value.secret
	case "type":
		return value.type_
	case "virtualName":
		return value.virtualName
	}
	return nil
}

func hxrt__generated_method_field__MethodLeaf(value *MethodLeaf, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "leafOnly":
		return value.leafOnly
	case "virtualName":
		return value.virtualName
	}
	if value.MethodMiddle == nil {
		return nil
	}
	return hxrt__generated_method_field__MethodMiddle(value.MethodMiddle, key)
}

func hxrt__generated_method_field__MethodMiddle(value *MethodMiddle, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "middleOnly":
		return value.middleOnly
	}
	if value.MethodBase == nil {
		return nil
	}
	return hxrt__generated_method_field__MethodBase(value.MethodBase, key)
}

func hxrt__generated_method_field__SpecializedIterator(value *SpecializedIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "next":
		return value.next
	}
	if value.ConcreteIterator == nil {
		return nil
	}
	return hxrt__generated_method_field__ConcreteIterator(value.ConcreteIterator, key)
}

func hxrt__generated_method_field__haxe__Template(value *haxe__Template, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "execute":
		return value.execute
	case "makeConst":
		return value.makeConst
	case "makeExpr":
		return value.makeExpr
	case "makeExpr2":
		return value.makeExpr2
	case "makePath":
		return value.makePath
	case "parse":
		return value.parse
	case "parseBlock":
		return value.parseBlock
	case "parseExpr":
		return value.parseExpr
	case "parseTokens":
		return value.parseTokens
	case "popStackValue":
		return value.popStackValue
	case "resolve":
		return value.resolve
	case "run":
		return value.run
	case "skipSpaces":
		return value.skipSpaces
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

type Std struct {
}

type Type struct {
}

type Reflect struct {
}

func Reflect_compare(a any, b any) int {
	toFloat := func(value any) (float64, bool) {
		switch v := value.(type) {
		case int:
			return float64(v), true
		case int8:
			return float64(v), true
		case int16:
			return float64(v), true
		case int32:
			return float64(v), true
		case int64:
			return float64(v), true
		case uint:
			return float64(v), true
		case uint8:
			return float64(v), true
		case uint16:
			return float64(v), true
		case uint32:
			return float64(v), true
		case uint64:
			return float64(v), true
		case float32:
			return float64(v), true
		case float64:
			return v, true
		default:
			return 0, false
		}
	}
	if af, ok := toFloat(a); ok {
		if bf, okB := toFloat(b); okB {
			if af < bf {
				return -1
			}
			if af > bf {
				return 1
			}
			return 0
		}
	}
	aStr := *hxrt.StdString(a)
	bStr := *hxrt.StdString(b)
	if aStr < bStr {
		return -1
	}
	if aStr > bStr {
		return 1
	}
	return 0
}

func Reflect_compareMethods(a any, b any) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	if !av.IsValid() || !bv.IsValid() {
		return !av.IsValid() && !bv.IsValid()
	}
	if av.Kind() == reflect.Func && bv.Kind() == reflect.Func {
		if av.IsNil() || bv.IsNil() {
			return av.IsNil() && bv.IsNil()
		}
		return av.Pointer() == bv.Pointer()
	}
	return reflect.DeepEqual(a, b)
}

func Reflect_field(obj any, field *string) any {
	if obj == nil {
		return nil
	}
	key := *hxrt.StdString(field)
	if metadataValue, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return metadataValue
	}
	switch value := obj.(type) {
	case map[string]any:
		return value[key]
	case map[any]any:
		return value[key]
	case *map[string]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	case *map[any]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return nil
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if fieldValue := rv.FieldByName(key); fieldValue.IsValid() && fieldValue.CanInterface() {
			return fieldValue.Interface()
		}
	}
	generatedMethod := hxrt__generated_method_field(obj, key)
	if generatedMethod != nil {
		return generatedMethod
	}
	method := reflect.ValueOf(obj).MethodByName(key)
	if method.IsValid() {
		return method.Interface()
	}
	return nil
}

func Reflect_hasField(obj any, field *string) bool {
	if obj == nil {
		return false
	}
	key := *hxrt.StdString(field)
	if _, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return true
	}
	switch value := obj.(type) {
	case map[string]any:
		_, ok := value[key]
		return ok
	case map[any]any:
		_, ok := value[key]
		return ok
	case *map[string]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	case *map[any]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if rv.FieldByName(key).IsValid() {
			return true
		}
	}
	generatedMethod := hxrt__generated_method_field(obj, key)
	if generatedMethod != nil {
		return true
	}
	return reflect.ValueOf(obj).MethodByName(key).IsValid()
}

func Reflect_setField(obj any, field *string, value any) {
	if obj == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	key := *hxrt.StdString(field)
	switch target := obj.(type) {
	case map[string]any:
		target[key] = value
		return
	case map[any]any:
		target[key] = value
		return
	case *map[string]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	case *map[any]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() || rv.Kind() != reflect.Pointer {
		return
	}
	if rv.IsNil() {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	fieldValue := rv.FieldByName(key)
	if !fieldValue.IsValid() || !fieldValue.CanSet() {
		return
	}
	if value == nil {
		fieldValue.Set(reflect.Zero(fieldValue.Type()))
		return
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(fieldValue.Type()) {
		fieldValue.Set(incoming)
		return
	}
	if incoming.Type().ConvertibleTo(fieldValue.Type()) {
		fieldValue.Set(incoming.Convert(fieldValue.Type()))
		return
	}
	if fieldValue.Kind() == reflect.Interface {
		fieldValue.Set(incoming)
	}
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
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
	case "ConcreteIterable":
		return hxrt_typeCallAny(New_ConcreteIterable, args)
	case "ConcreteIterator":
		return hxrt_typeCallAny(New_ConcreteIterator, args)
	case "EReg":
		return hxrt_typeCallAny(New_EReg, args)
	case "Main":
		return nil, false
	case "MethodBase":
		return hxrt_typeCallAny(New_MethodBase, args)
	case "MethodLeaf":
		return hxrt_typeCallAny(New_MethodLeaf, args)
	case "MethodMiddle":
		return hxrt_typeCallAny(New_MethodMiddle, args)
	case "SpecializedIterator":
		return hxrt_typeCallAny(New_SpecializedIterator, args)
	case "StringBuf":
		return nil, false
	case "StringTools":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe.Template":
		return hxrt_typeCallAny(New_haxe__Template, args)
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe._Template.ExprCursor":
		return hxrt_typeCallAny(New_haxe___Template__ExprCursor, args)
	case "haxe._Template.TokenCursor":
		return hxrt_typeCallAny(New_haxe___Template__TokenCursor, args)
	case "haxe.io.Bytes":
		return hxrt_typeCallAny(New_haxe__io__Bytes, args)
	case "haxe.io.FPHelper":
		return nil, false
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
	case "ConcreteIterable":
		return &ConcreteIterable{}, true
	case "ConcreteIterator":
		return &ConcreteIterator{}, true
	case "EReg":
		return &EReg{}, true
	case "MethodBase":
		return &MethodBase{}, true
	case "MethodLeaf":
		return &MethodLeaf{}, true
	case "MethodMiddle":
		return &MethodMiddle{}, true
	case "SpecializedIterator":
		return &SpecializedIterator{}, true
	case "haxe.Template":
		return &haxe__Template{}, true
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	case "haxe._Template.ExprCursor":
		return &haxe___Template__ExprCursor{}, true
	case "haxe._Template.TokenCursor":
		return &haxe___Template__TokenCursor{}, true
	case "haxe.io.Bytes":
		return &haxe__io__Bytes{}, true
	case "haxe.iterators.StringIterator":
		return &haxe__iterators__StringIterator{}, true
	case "haxe.iterators.StringKeyValueIterator":
		return &haxe__iterators__StringKeyValueIterator{}, true
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
	case "haxe._Template.TemplateExpr":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpVar, args)
			case 1:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpExpr, args)
			case 2:
				if len(args) != 3 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpIf, args)
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpStr, args)
			case 4:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpBlock, args)
			case 5:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpForeach, args)
			case 6:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpMacro, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "OpVar":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpVar, args)
		case "OpExpr":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpExpr, args)
		case "OpIf":
			if len(args) != 3 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpIf, args)
		case "OpStr":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpStr, args)
		case "OpBlock":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpBlock, args)
		case "OpForeach":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpForeach, args)
		case "OpMacro":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpMacro, args)
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
	case *ConcreteIterable:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("ConcreteIterable")}
	case *ConcreteIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("ConcreteIterator")}
	case *EReg:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("EReg")}
	case *MethodBase:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("MethodBase")}
	case *MethodLeaf:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("MethodLeaf")}
	case *MethodMiddle:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("MethodMiddle")}
	case *SpecializedIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("SpecializedIterator")}
	case *haxe__Template:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Template")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe___Template__ExprCursor:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Template.ExprCursor")}
	case *haxe___Template__TokenCursor:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Template.TokenCursor")}
	case *haxe__io__Bytes:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Bytes")}
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
	case *haxe___Template__TemplateExpr:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe._Template.TemplateExpr")}
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
	case "ConcreteIterable":
		return nil
	case "ConcreteIterator":
		return nil
	case "EReg":
		return nil
	case "Main":
		return nil
	case "MethodBase":
		return nil
	case "MethodLeaf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("MethodMiddle")}
	case "MethodMiddle":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("MethodBase")}
	case "SpecializedIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("ConcreteIterator")}
	case "StringBuf":
		return nil
	case "StringTools":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe.Template":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe._Template.ExprCursor":
		return nil
	case "haxe._Template.TokenCursor":
		return nil
	case "haxe.io.Bytes":
		return nil
	case "haxe.io.FPHelper":
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
	case "ConcreteIterable":
		return hxrt.NewArray()
	case "ConcreteIterator":
		return hxrt.NewArray()
	case "EReg":
		return hxrt.NewArray(hxrt.StringFromLiteral("escape"))
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("call0"), hxrt.StringFromLiteral("call1"), hxrt.StringFromLiteral("main"))
	case "MethodBase":
		return hxrt.NewArray()
	case "MethodLeaf":
		return hxrt.NewArray()
	case "MethodMiddle":
		return hxrt.NewArray()
	case "SpecializedIterator":
		return hxrt.NewArray()
	case "StringBuf":
		return hxrt.NewArray()
	case "StringTools":
		return hxrt.NewArray(hxrt.StringFromLiteral("MAX_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("contains"), hxrt.StringFromLiteral("containsImpl"), hxrt.StringFromLiteral("endsWith"), hxrt.StringFromLiteral("endsWithImpl"), hxrt.StringFromLiteral("fastCodeAt"), hxrt.StringFromLiteral("hex"), hxrt.StringFromLiteral("hexDigitValue"), hxrt.StringFromLiteral("htmlEscape"), hxrt.StringFromLiteral("htmlUnescape"), hxrt.StringFromLiteral("isEof"), hxrt.StringFromLiteral("isSpace"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("lpad"), hxrt.StringFromLiteral("ltrim"), hxrt.StringFromLiteral("replace"), hxrt.StringFromLiteral("rpad"), hxrt.StringFromLiteral("rtrim"), hxrt.StringFromLiteral("startsWith"), hxrt.StringFromLiteral("startsWithImpl"), hxrt.StringFromLiteral("trim"), hxrt.StringFromLiteral("unsafeCodeAt"), hxrt.StringFromLiteral("urlDecode"), hxrt.StringFromLiteral("urlEncode"), hxrt.StringFromLiteral("utf16CodePointAt"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Template":
		return hxrt.NewArray(hxrt.StringFromLiteral("addValues"), hxrt.StringFromLiteral("compareValues"), hxrt.StringFromLiteral("divideValues"), hxrt.StringFromLiteral("expr_float"), hxrt.StringFromLiteral("expr_int"), hxrt.StringFromLiteral("expr_splitter"), hxrt.StringFromLiteral("expr_trim"), hxrt.StringFromLiteral("globals"), hxrt.StringFromLiteral("isSpaceOnly"), hxrt.StringFromLiteral("joinDynamicArgs"), hxrt.StringFromLiteral("kwdEnd"), hxrt.StringFromLiteral("multiplyValues"), hxrt.StringFromLiteral("parseFloatLiteral"), hxrt.StringFromLiteral("parseIntLiteral"), hxrt.StringFromLiteral("peekExprToken"), hxrt.StringFromLiteral("peekToken"), hxrt.StringFromLiteral("popExprToken"), hxrt.StringFromLiteral("popToken"), hxrt.StringFromLiteral("splitter"), hxrt.StringFromLiteral("subtractValues"), hxrt.StringFromLiteral("trimExprToken"), hxrt.StringFromLiteral("valueAsBool"), hxrt.StringFromLiteral("valueAsFloat"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe._Template.ExprCursor":
		return hxrt.NewArray()
	case "haxe._Template.TokenCursor":
		return hxrt.NewArray()
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_fromNativeView"), hxrt.StringFromLiteral("alloc"), hxrt.StringFromLiteral("fastGet"), hxrt.StringFromLiteral("ofData"), hxrt.StringFromLiteral("ofHex"), hxrt.StringFromLiteral("ofString"), hxrt.StringFromLiteral("rawNativeUsesUtf16LE"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray(hxrt.StringFromLiteral("doubleToI64"), hxrt.StringFromLiteral("floatToI32"), hxrt.StringFromLiteral("i32ToFloat"), hxrt.StringFromLiteral("i64ToDouble"))
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
	case "ConcreteIterable":
		return hxrt.NewArray(hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("values"))
	case "ConcreteIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("values"))
	case "EReg":
		return hxrt.NewArray(hxrt.StringFromLiteral("expandReplacement"), hxrt.StringFromLiteral("global"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("lastMatch"), hxrt.StringFromLiteral("lastSource"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("match"), hxrt.StringFromLiteral("matchSub"), hxrt.StringFromLiteral("matched"), hxrt.StringFromLiteral("matchedLeft"), hxrt.StringFromLiteral("matchedPos"), hxrt.StringFromLiteral("matchedRight"), hxrt.StringFromLiteral("remember"), hxrt.StringFromLiteral("replace"), hxrt.StringFromLiteral("requireMatch"), hxrt.StringFromLiteral("split"))
	case "Main":
		return hxrt.NewArray()
	case "MethodBase":
		return hxrt.NewArray(hxrt.StringFromLiteral("bump"), hxrt.StringFromLiteral("describe"), hxrt.StringFromLiteral("macroName"), hxrt.StringFromLiteral("retainSecret"), hxrt.StringFromLiteral("secret"), hxrt.StringFromLiteral("total"), hxrt.StringFromLiteral("type"), hxrt.StringFromLiteral("virtualName"))
	case "MethodLeaf":
		return hxrt.NewArray(hxrt.StringFromLiteral("bump"), hxrt.StringFromLiteral("describe"), hxrt.StringFromLiteral("leafOnly"), hxrt.StringFromLiteral("macroName"), hxrt.StringFromLiteral("middleOnly"), hxrt.StringFromLiteral("retainSecret"), hxrt.StringFromLiteral("secret"), hxrt.StringFromLiteral("total"), hxrt.StringFromLiteral("type"), hxrt.StringFromLiteral("virtualName"))
	case "MethodMiddle":
		return hxrt.NewArray(hxrt.StringFromLiteral("bump"), hxrt.StringFromLiteral("describe"), hxrt.StringFromLiteral("macroName"), hxrt.StringFromLiteral("middleOnly"), hxrt.StringFromLiteral("retainSecret"), hxrt.StringFromLiteral("secret"), hxrt.StringFromLiteral("total"), hxrt.StringFromLiteral("type"), hxrt.StringFromLiteral("virtualName"))
	case "SpecializedIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("values"))
	case "StringBuf":
		return hxrt.NewArray()
	case "StringTools":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Template":
		return hxrt.NewArray(hxrt.StringFromLiteral("context"), hxrt.StringFromLiteral("execute"), hxrt.StringFromLiteral("expr"), hxrt.StringFromLiteral("macros"), hxrt.StringFromLiteral("makeConst"), hxrt.StringFromLiteral("makeExpr"), hxrt.StringFromLiteral("makeExpr2"), hxrt.StringFromLiteral("makePath"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("parse"), hxrt.StringFromLiteral("parseBlock"), hxrt.StringFromLiteral("parseExpr"), hxrt.StringFromLiteral("parseTokens"), hxrt.StringFromLiteral("popStackValue"), hxrt.StringFromLiteral("resolve"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("skipSpaces"), hxrt.StringFromLiteral("stack"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe._Template.ExprCursor":
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("tokens"))
	case "haxe._Template.TokenCursor":
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("tokens"))
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_nativeView"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("blit"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("fill"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getData"), hxrt.StringFromLiteral("getDouble"), hxrt.StringFromLiteral("getFloat"), hxrt.StringFromLiteral("getInt32"), hxrt.StringFromLiteral("getInt64"), hxrt.StringFromLiteral("getString"), hxrt.StringFromLiteral("getUInt16"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setDouble"), hxrt.StringFromLiteral("setFloat"), hxrt.StringFromLiteral("setInt32"), hxrt.StringFromLiteral("setInt64"), hxrt.StringFromLiteral("setUInt16"), hxrt.StringFromLiteral("sub"), hxrt.StringFromLiteral("toHex"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray()
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
	case "ConcreteIterable":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "ConcreteIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "EReg":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "MethodBase":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "MethodLeaf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "MethodMiddle":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "SpecializedIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringBuf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringTools":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Template":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Template.ExprCursor":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Template.TokenCursor":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Bytes":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.FPHelper":
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
	case "haxe._Template.TemplateExpr":
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
	case *haxe___Template__TemplateExpr:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("OpVar")
		case 1:
			return hxrt.StringFromLiteral("OpExpr")
		case 2:
			return hxrt.StringFromLiteral("OpIf")
		case 3:
			return hxrt.StringFromLiteral("OpStr")
		case 4:
			return hxrt.StringFromLiteral("OpBlock")
		case 5:
			return hxrt.StringFromLiteral("OpForeach")
		case 6:
			return hxrt.StringFromLiteral("OpMacro")
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
	case *haxe___Template__TemplateExpr:
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
	case "haxe._Template.TemplateExpr":
		return hxrt.NewArray(hxrt.StringFromLiteral("OpVar"), hxrt.StringFromLiteral("OpExpr"), hxrt.StringFromLiteral("OpIf"), hxrt.StringFromLiteral("OpStr"), hxrt.StringFromLiteral("OpBlock"), hxrt.StringFromLiteral("OpForeach"), hxrt.StringFromLiteral("OpMacro"))
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
	case *haxe___Template__TemplateExpr:
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
	case "haxe._Template.TemplateExpr":
		return hxrt.NewArray()
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
	className, ok := hxrt_typeResolvedClassName(value)
	if !ok {
		return nil, false
	}
	switch className {
	default:
		return nil, false
	}
}
