package main

import "snapshot/hxrt"

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

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
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
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Template__ExprCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Template__TokenCursor:
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
		return hxrt__generated_field_lookup__ConcreteIterable(value, key)
	case *ConcreteIterator:
		return hxrt__generated_field_lookup__ConcreteIterator(value, key)
	case *EReg:
		return hxrt__generated_field_lookup__EReg(value, key)
	case *MethodBase:
		return hxrt__generated_field_lookup__MethodBase(value, key)
	case *MethodLeaf:
		return hxrt__generated_field_lookup__MethodLeaf(value, key)
	case *MethodMiddle:
		return hxrt__generated_field_lookup__MethodMiddle(value, key)
	case *SpecializedIterator:
		return hxrt__generated_field_lookup__SpecializedIterator(value, key)
	case *haxe__Template:
		return hxrt__generated_field_lookup__haxe__Template(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe___Template__ExprCursor:
		return hxrt__generated_field_lookup__haxe___Template__ExprCursor(value, key)
	case *haxe___Template__TokenCursor:
		return hxrt__generated_field_lookup__haxe___Template__TokenCursor(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_lookup__haxe__io__Bytes(value, key)
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
	case *ConcreteIterable:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *ConcreteIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *MethodBase:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *MethodLeaf:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *MethodMiddle:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *SpecializedIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Template:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Template__ExprCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Template__TokenCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
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
	case *ConcreteIterable:
		return hxrt__generated_field_has__ConcreteIterable(value, key)
	case *ConcreteIterator:
		return hxrt__generated_field_has__ConcreteIterator(value, key)
	case *EReg:
		return hxrt__generated_field_has__EReg(value, key)
	case *MethodBase:
		return hxrt__generated_field_has__MethodBase(value, key)
	case *MethodLeaf:
		return hxrt__generated_field_has__MethodLeaf(value, key)
	case *MethodMiddle:
		return hxrt__generated_field_has__MethodMiddle(value, key)
	case *SpecializedIterator:
		return hxrt__generated_field_has__SpecializedIterator(value, key)
	case *haxe__Template:
		return hxrt__generated_field_has__haxe__Template(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe___Template__ExprCursor:
		return hxrt__generated_field_has__haxe___Template__ExprCursor(value, key)
	case *haxe___Template__TokenCursor:
		return hxrt__generated_field_has__haxe___Template__TokenCursor(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_has__haxe__io__Bytes(value, key)
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
	case *ConcreteIterable:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *ConcreteIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *MethodBase:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *MethodLeaf:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *MethodMiddle:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *SpecializedIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Template:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Template__ExprCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Template__TokenCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
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
	case *ConcreteIterable:
		return hxrt__generated_field_set__ConcreteIterable(value, key, incoming)
	case *ConcreteIterator:
		return hxrt__generated_field_set__ConcreteIterator(value, key, incoming)
	case *EReg:
		return hxrt__generated_field_set__EReg(value, key, incoming)
	case *MethodBase:
		return hxrt__generated_field_set__MethodBase(value, key, incoming)
	case *MethodLeaf:
		return hxrt__generated_field_set__MethodLeaf(value, key, incoming)
	case *MethodMiddle:
		return hxrt__generated_field_set__MethodMiddle(value, key, incoming)
	case *SpecializedIterator:
		return hxrt__generated_field_set__SpecializedIterator(value, key, incoming)
	case *haxe__Template:
		return hxrt__generated_field_set__haxe__Template(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe___Template__ExprCursor:
		return hxrt__generated_field_set__haxe___Template__ExprCursor(value, key, incoming)
	case *haxe___Template__TokenCursor:
		return hxrt__generated_field_set__haxe___Template__TokenCursor(value, key, incoming)
	case *haxe__io__Bytes:
		return hxrt__generated_field_set__haxe__io__Bytes(value, key, incoming)
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
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Template__ExprCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Template__TokenCursor:
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
	switch receiver.(type) {
	case *ConcreteIterable:
		return hxrt.NewArray(hxrt.StringFromLiteral("values"))
	case *ConcreteIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("values"))
	case *EReg:
		return hxrt.NewArray(hxrt.StringFromLiteral("global"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("lastMatch"), hxrt.StringFromLiteral("lastSource"))
	case *MethodBase:
		return hxrt.NewArray(hxrt.StringFromLiteral("total"))
	case *MethodLeaf:
		return hxrt.NewArray(hxrt.StringFromLiteral("total"))
	case *MethodMiddle:
		return hxrt.NewArray(hxrt.StringFromLiteral("total"))
	case *SpecializedIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("values"))
	case *haxe__Template:
		return hxrt.NewArray(hxrt.StringFromLiteral("context"), hxrt.StringFromLiteral("expr"), hxrt.StringFromLiteral("macros"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("stack"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe___Template__ExprCursor:
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("tokens"))
	case *haxe___Template__TokenCursor:
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("tokens"))
	case *haxe__io__Bytes:
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("length"))
	case *haxe__iterators__StringIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__ConcreteIterable(value *ConcreteIterable, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "values":
		return value.values
	}
	return nil
}

func hxrt__generated_field_has__ConcreteIterable(value *ConcreteIterable, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "values":
		return true
	}
	return false
}

func hxrt__generated_field_set__ConcreteIterable(value *ConcreteIterable, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "values":
		if incoming == nil {
			var zero *hxrt.Array
			value.values = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.values = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__ConcreteIterator(value *ConcreteIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "index":
		return value.index
	case "values":
		return value.values
	}
	return nil
}

func hxrt__generated_field_has__ConcreteIterator(value *ConcreteIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		return true
	case "values":
		return true
	}
	return false
}

func hxrt__generated_field_set__ConcreteIterator(value *ConcreteIterator, key string, incoming any) bool {
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
	case "values":
		if incoming == nil {
			var zero *hxrt.Array
			value.values = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.values = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__EReg(value *EReg, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "global":
		return value.global
	case "handle":
		return value.handle
	case "lastMatch":
		return value.lastMatch
	case "lastSource":
		return value.lastSource
	}
	return nil
}

func hxrt__generated_field_has__EReg(value *EReg, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "global":
		return true
	case "handle":
		return true
	case "lastMatch":
		return true
	case "lastSource":
		return true
	}
	return false
}

func hxrt__generated_field_set__EReg(value *EReg, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "global":
		if incoming == nil {
			var zero bool
			value.global = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.global = typed
			return true
		default:
			return false
		}
	case "handle":
		if incoming == nil {
			var zero *hxrt.RegexHandle
			value.handle = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.RegexHandle:
			value.handle = typed
			return true
		default:
			return false
		}
	case "lastMatch":
		if incoming == nil {
			var zero *hxrt.RegexMatch
			value.lastMatch = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.RegexMatch:
			value.lastMatch = typed
			return true
		default:
			return false
		}
	case "lastSource":
		if incoming == nil {
			var zero *string
			value.lastSource = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.lastSource = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__MethodBase(value *MethodBase, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "total":
		return value.total
	}
	return nil
}

func hxrt__generated_field_has__MethodBase(value *MethodBase, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "total":
		return true
	}
	return false
}

func hxrt__generated_field_set__MethodBase(value *MethodBase, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "total":
		if incoming == nil {
			var zero int
			value.total = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.total = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__MethodLeaf(value *MethodLeaf, key string) any {
	if value == nil {
		return nil
	}
	if value.MethodMiddle == nil {
		return nil
	}
	return hxrt__generated_field_lookup__MethodMiddle(value.MethodMiddle, key)
}

func hxrt__generated_field_has__MethodLeaf(value *MethodLeaf, key string) bool {
	if value == nil {
		return false
	}
	if value.MethodMiddle == nil {
		return false
	}
	return hxrt__generated_field_has__MethodMiddle(value.MethodMiddle, key)
}

func hxrt__generated_field_set__MethodLeaf(value *MethodLeaf, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.MethodMiddle == nil {
		return false
	}
	return hxrt__generated_field_set__MethodMiddle(value.MethodMiddle, key, incoming)
}

func hxrt__generated_field_lookup__MethodMiddle(value *MethodMiddle, key string) any {
	if value == nil {
		return nil
	}
	if value.MethodBase == nil {
		return nil
	}
	return hxrt__generated_field_lookup__MethodBase(value.MethodBase, key)
}

func hxrt__generated_field_has__MethodMiddle(value *MethodMiddle, key string) bool {
	if value == nil {
		return false
	}
	if value.MethodBase == nil {
		return false
	}
	return hxrt__generated_field_has__MethodBase(value.MethodBase, key)
}

func hxrt__generated_field_set__MethodMiddle(value *MethodMiddle, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.MethodBase == nil {
		return false
	}
	return hxrt__generated_field_set__MethodBase(value.MethodBase, key, incoming)
}

func hxrt__generated_field_lookup__SpecializedIterator(value *SpecializedIterator, key string) any {
	if value == nil {
		return nil
	}
	if value.ConcreteIterator == nil {
		return nil
	}
	return hxrt__generated_field_lookup__ConcreteIterator(value.ConcreteIterator, key)
}

func hxrt__generated_field_has__SpecializedIterator(value *SpecializedIterator, key string) bool {
	if value == nil {
		return false
	}
	if value.ConcreteIterator == nil {
		return false
	}
	return hxrt__generated_field_has__ConcreteIterator(value.ConcreteIterator, key)
}

func hxrt__generated_field_set__SpecializedIterator(value *SpecializedIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.ConcreteIterator == nil {
		return false
	}
	return hxrt__generated_field_set__ConcreteIterator(value.ConcreteIterator, key, incoming)
}

func hxrt__generated_field_lookup__haxe__Template(value *haxe__Template, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "context":
		return value.context
	case "expr":
		return value.expr
	case "macros":
		return value.macros
	case "output":
		return value.output
	case "stack":
		return value.stack
	}
	return nil
}

func hxrt__generated_field_has__haxe__Template(value *haxe__Template, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "context":
		return true
	case "expr":
		return true
	case "macros":
		return true
	case "output":
		return true
	case "stack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Template(value *haxe__Template, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "context":
		if incoming == nil {
			var zero any
			value.context = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.context = typed
			return true
		default:
			return false
		}
	case "expr":
		if incoming == nil {
			var zero *haxe___Template__TemplateExpr
			value.expr = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe___Template__TemplateExpr:
			value.expr = typed
			return true
		default:
			return false
		}
	case "macros":
		if incoming == nil {
			var zero any
			value.macros = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.macros = typed
			return true
		default:
			return false
		}
	case "output":
		if incoming == nil {
			var zero *string
			value.output = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.output = typed
			return true
		default:
			return false
		}
	case "stack":
		if incoming == nil {
			var zero *hxrt.Array
			value.stack = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.stack = typed
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

func hxrt__generated_field_lookup__haxe___Template__ExprCursor(value *haxe___Template__ExprCursor, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "index":
		return value.index
	case "tokens":
		return value.tokens
	}
	return nil
}

func hxrt__generated_field_has__haxe___Template__ExprCursor(value *haxe___Template__ExprCursor, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		return true
	case "tokens":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Template__ExprCursor(value *haxe___Template__ExprCursor, key string, incoming any) bool {
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
	case "tokens":
		if incoming == nil {
			var zero *hxrt.Array
			value.tokens = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.tokens = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe___Template__TokenCursor(value *haxe___Template__TokenCursor, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "index":
		return value.index
	case "tokens":
		return value.tokens
	}
	return nil
}

func hxrt__generated_field_has__haxe___Template__TokenCursor(value *haxe___Template__TokenCursor, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		return true
	case "tokens":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Template__TokenCursor(value *haxe___Template__TokenCursor, key string, incoming any) bool {
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
	case "tokens":
		if incoming == nil {
			var zero *hxrt.Array
			value.tokens = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.tokens = typed
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
	case *haxe___Template__TemplateExpr:
		return (enumValue != nil)
	case *haxe__io__Encoding:
		return (enumValue != nil)
	case *haxe__io__Error:
		return (enumValue != nil)
	default:
		return false
	}
}
