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

func (self *ConcreteIterable) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_ConcreteIterable(values *hxrt.Array) *ConcreteIterable {
	return hx_init_ConcreteIterable(hx_alloc_ConcreteIterable(nil), values)
}

func hx_alloc_ConcreteIterable(owner any) *ConcreteIterable {
	self := &ConcreteIterable{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_ConcreteIterable)
	}
	return self
}

func hx_init_ConcreteIterable(self *ConcreteIterable, values *hxrt.Array) *ConcreteIterable {
	self.values = values
	return self
}

func (self *ConcreteIterable) iterator() *SpecializedIterator {
	return New_SpecializedIterator(self.values)
}

func (self *ConcreteIterable) HxClassName() string {
	return "ConcreteIterable"
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

func (self *ConcreteIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_ConcreteIterator(values *hxrt.Array) *ConcreteIterator {
	return hx_init_ConcreteIterator(hx_alloc_ConcreteIterator(nil), values)
}

func hx_alloc_ConcreteIterator(owner any) *ConcreteIterator {
	self := &ConcreteIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_ConcreteIterator)
	}
	return self
}

func hx_init_ConcreteIterator(self *ConcreteIterator, values *hxrt.Array) *ConcreteIterator {
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

func (self *ConcreteIterator) HxClassName() string {
	return "ConcreteIterator"
}

func call0(obj any, key *string) *string {
	template := New_haxe__Template(hxrt.StringFromLiteral("$$invoke(dummy)"))
	return template.__hx_this.execute(func() hxrt.Anonymous {
		hx_obj_5 := map[string]any{}
		hx_obj_5["dummy"] = nil
		return hxrt.AnonymousFromMap(hx_obj_5)
	}(), func() hxrt.Anonymous {
		hx_obj_6 := map[string]any{}
		hx_obj_6["invoke"] = Reflect_field(obj, key)
		return hxrt.AnonymousFromMap(hx_obj_6)
	}())
}

func call1(obj any, key *string, value any) *string {
	template := New_haxe__Template(hxrt.StringFromLiteral("$$invoke(value)"))
	return template.__hx_this.execute(func() hxrt.Anonymous {
		hx_obj_7 := map[string]any{}
		hx_obj_7["value"] = value
		return hxrt.AnonymousFromMap(hx_obj_7)
	}(), func() hxrt.Anonymous {
		hx_obj_8 := map[string]any{}
		hx_obj_8["invoke"] = Reflect_field(obj, key)
		return hxrt.AnonymousFromMap(hx_obj_8)
	}())
}

func main() {
	template := New_haxe__Template(hxrt.StringFromLiteral("::foreach items::::__current__::;::end::"))
	var v any = any(template.__hx_this.execute(func() hxrt.Anonymous {
		hx_obj_9 := map[string]any{}
		hx_obj_9["items"] = New_ConcreteIterable(hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b")))
		return hxrt.AnonymousFromMap(hx_obj_9)
	}(), nil))
	hxrt.Println(v)
	var v_1 any = any(template.__hx_this.execute(func() hxrt.Anonymous {
		hx_obj_10 := map[string]any{}
		hx_obj_10["items"] = New_ConcreteIterator(hxrt.NewArray(hxrt.StringFromLiteral("x"), hxrt.StringFromLiteral("y")))
		return hxrt.AnonymousFromMap(hx_obj_10)
	}(), nil))
	hxrt.Println(v_1)
	leaf := New_MethodLeaf()
	middle := func(hx_value *MethodLeaf) *MethodMiddle {
		if hx_value == nil {
			return nil
		}
		return hx_value.MethodMiddle
	}(leaf)
	base := func(hx_value *MethodLeaf) *MethodBase {
		if hx_value == nil {
			return nil
		}
		return hx_value.MethodMiddle.MethodBase
	}(leaf)
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
	macroName(resolve *hxrt.Function[func(*string) any], ignored any) *string
	describe(resolve *hxrt.Function[func(*string) any], key *string) *string
	bump(resolve *hxrt.Function[func(*string) any], key *string) int
	type_(resolve *hxrt.Function[func(*string) any], ignored any) *string
	secret(resolve *hxrt.Function[func(*string) any], ignored any) *string
	retainSecret() *string
}

type MethodBase struct {
	__hx_this I_MethodBase
	total     int
}

func (self *MethodBase) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_MethodBase() *MethodBase {
	return hx_init_MethodBase(hx_alloc_MethodBase(nil))
}

func hx_alloc_MethodBase(owner any) *MethodBase {
	self := &MethodBase{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_MethodBase)
	}
	return self
}

func hx_init_MethodBase(self *MethodBase) *MethodBase {
	self.total = 0
	return self
}

func (self *MethodBase) virtualName() *string {
	return hxrt.StringFromLiteral("base")
}

func (self *MethodBase) macroName(resolve *hxrt.Function[func(*string) any], ignored any) *string {
	return self.__hx_this.virtualName()
}

func (self *MethodBase) describe(resolve *hxrt.Function[func(*string) any], key *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func() any {
		var hx_callable_11 *hxrt.Function[func(*string) any] = resolve
		var hx_argument_12 *string = key
		return hxrt.FunctionCall(hx_callable_11)(hx_argument_12)
	}()), hxrt.StringFromLiteral(":")), self.__hx_this.virtualName())
}

func (self *MethodBase) bump(resolve *hxrt.Function[func(*string) any], key *string) int {
	amount := hxrt.IntFromNullableAny(func() any {
		var hx_callable_13 *hxrt.Function[func(*string) any] = resolve
		var hx_argument_14 *string = key
		return hxrt.FunctionCall(hx_callable_13)(hx_argument_14)
	}())
	self.total = int((hxrt.Int32Wrap(self.total) + hxrt.Int32Wrap(amount)))
	return self.total
}

func (self *MethodBase) type_(resolve *hxrt.Function[func(*string) any], ignored any) *string {
	return hxrt.StringFromLiteral("keyword")
}

func (self *MethodBase) secret(resolve *hxrt.Function[func(*string) any], ignored any) *string {
	return hxrt.StringFromLiteral("secret")
}

func (self *MethodBase) retainSecret() *string {
	return self.__hx_this.secret(nil, nil)
}

func (self *MethodBase) HxClassName() string {
	return "MethodBase"
}

type I_MethodLeaf interface {
	virtualName() *string
	macroName(resolve *hxrt.Function[func(*string) any], ignored any) *string
	describe(resolve *hxrt.Function[func(*string) any], key *string) *string
	bump(resolve *hxrt.Function[func(*string) any], key *string) int
	type_(resolve *hxrt.Function[func(*string) any], ignored any) *string
	secret(resolve *hxrt.Function[func(*string) any], ignored any) *string
	retainSecret() *string
	middleOnly(resolve *hxrt.Function[func(*string) any], ignored any) *string
	leafOnly(resolve *hxrt.Function[func(*string) any], ignored any) *string
}

type MethodLeaf struct {
	*MethodMiddle
	__hx_this I_MethodLeaf
}

func (self *MethodLeaf) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_MethodLeaf() *MethodLeaf {
	return hx_init_MethodLeaf(hx_alloc_MethodLeaf(nil))
}

func hx_alloc_MethodLeaf(owner any) *MethodLeaf {
	self := &MethodLeaf{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_MethodLeaf)
	}
	self.MethodMiddle = hx_alloc_MethodMiddle(self.__hx_this)
	return self
}

func hx_init_MethodLeaf(self *MethodLeaf) *MethodLeaf {
	hx_init_MethodMiddle(self.MethodMiddle)
	return self
}

func (self *MethodLeaf) virtualName() *string {
	return hxrt.StringFromLiteral("leaf")
}

func (self *MethodLeaf) leafOnly(resolve *hxrt.Function[func(*string) any], ignored any) *string {
	return hxrt.StringFromLiteral("leaf-only")
}

func (self *MethodLeaf) HxClassName() string {
	return "MethodLeaf"
}

type I_MethodMiddle interface {
	virtualName() *string
	macroName(resolve *hxrt.Function[func(*string) any], ignored any) *string
	describe(resolve *hxrt.Function[func(*string) any], key *string) *string
	bump(resolve *hxrt.Function[func(*string) any], key *string) int
	type_(resolve *hxrt.Function[func(*string) any], ignored any) *string
	secret(resolve *hxrt.Function[func(*string) any], ignored any) *string
	retainSecret() *string
	middleOnly(resolve *hxrt.Function[func(*string) any], ignored any) *string
}

type MethodMiddle struct {
	*MethodBase
	__hx_this I_MethodMiddle
}

func (self *MethodMiddle) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_MethodMiddle() *MethodMiddle {
	return hx_init_MethodMiddle(hx_alloc_MethodMiddle(nil))
}

func hx_alloc_MethodMiddle(owner any) *MethodMiddle {
	self := &MethodMiddle{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_MethodMiddle)
	}
	self.MethodBase = hx_alloc_MethodBase(self.__hx_this)
	return self
}

func hx_init_MethodMiddle(self *MethodMiddle) *MethodMiddle {
	hx_init_MethodBase(self.MethodBase)
	return self
}

func (self *MethodMiddle) middleOnly(resolve *hxrt.Function[func(*string) any], ignored any) *string {
	return hxrt.StringFromLiteral("middle")
}

func (self *MethodMiddle) HxClassName() string {
	return "MethodMiddle"
}

type NameContract interface {
	macroName(resolve *hxrt.Function[func(*string) any], ignored any) *string
}

type I_SpecializedIterator interface {
	hasNext() bool
	next() *string
}

type SpecializedIterator struct {
	*ConcreteIterator
	__hx_this I_SpecializedIterator
}

func (self *SpecializedIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SpecializedIterator(values *hxrt.Array) *SpecializedIterator {
	return hx_init_SpecializedIterator(hx_alloc_SpecializedIterator(nil), values)
}

func hx_alloc_SpecializedIterator(owner any) *SpecializedIterator {
	self := &SpecializedIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SpecializedIterator)
	}
	self.ConcreteIterator = hx_alloc_ConcreteIterator(self.__hx_this)
	return self
}

func hx_init_SpecializedIterator(self *SpecializedIterator, values *hxrt.Array) *SpecializedIterator {
	hx_init_ConcreteIterator(self.ConcreteIterator, values)
	return self
}

func (self *SpecializedIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("special:"), self.values.Get(func() int {
		hx_post_18 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_18
	}()))
}

func (self *SpecializedIterator) HxClassName() string {
	return "SpecializedIterator"
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *ConcreteIterable:
		return func(hx_structural_source_1 *ConcreteIterable) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "values":
					return hxrt.AnonymousField{Value: hx_structural_source_1.values, Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *SpecializedIterator] {
						hx_receiver_2 := hx_structural_source_1
						return hxrt.BoundFunction(func() *SpecializedIterator {
							return hx_receiver_2.__hx_this.iterator()
						}, func() any {
							return hx_receiver_2.__hx_this
						}, "iterator")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *ConcreteIterator:
		return func(hx_structural_source_3 *ConcreteIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_3, func(name string) hxrt.AnonymousField {
				switch name {
				case "values":
					return hxrt.AnonymousField{Value: hx_structural_source_3.values, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_3.index, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_6 := hx_structural_source_3
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_6.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_6.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_7 := hx_structural_source_3
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_7.__hx_this.next()
						}, func() any {
							return hx_receiver_7.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "index":
					hx_structural_source_3.index = func(hx_value_4 any) int {
						if hx_value_4 == nil {
							var hx_zero_5 int
							return hx_zero_5
						}
						return hx_value_4.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *MethodBase:
		return func(hx_structural_source_8 *MethodBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_8, func(name string) hxrt.AnonymousField {
				switch name {
				case "total":
					return hxrt.AnonymousField{Value: hx_structural_source_8.total, Found: true}
				case "virtualName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_11 := hx_structural_source_8
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_11.__hx_this.virtualName()
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "virtualName")
					}(), Found: true}
				case "macroName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_12 := hx_structural_source_8
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_12.__hx_this.macroName(resolve, ignored)
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "macroName")
					}(), Found: true}
				case "describe":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) *string] {
						hx_receiver_13 := hx_structural_source_8
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) *string {
							return hx_receiver_13.__hx_this.describe(resolve, key)
						}, func() any {
							return hx_receiver_13.__hx_this
						}, "describe")
					}(), Found: true}
				case "bump":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) int] {
						hx_receiver_14 := hx_structural_source_8
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) int {
							return hx_receiver_14.__hx_this.bump(resolve, key)
						}, func() any {
							return hx_receiver_14.__hx_this
						}, "bump")
					}(), Found: true}
				case "type":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_15 := hx_structural_source_8
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_15.__hx_this.type_(resolve, ignored)
						}, func() any {
							return hx_receiver_15.__hx_this
						}, "type")
					}(), Found: true}
				case "secret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_16 := hx_structural_source_8
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_16.__hx_this.secret(resolve, ignored)
						}, func() any {
							return hx_receiver_16.__hx_this
						}, "secret")
					}(), Found: true}
				case "retainSecret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_17 := hx_structural_source_8
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_17.__hx_this.retainSecret()
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "retainSecret")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "total":
					hx_structural_source_8.total = func(hx_value_9 any) int {
						if hx_value_9 == nil {
							var hx_zero_10 int
							return hx_zero_10
						}
						return hx_value_9.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *MethodLeaf:
		return func(hx_structural_source_18 *MethodLeaf) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_18, func(name string) hxrt.AnonymousField {
				switch name {
				case "virtualName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_19 := hx_structural_source_18
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_19.__hx_this.virtualName()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "virtualName")
					}(), Found: true}
				case "leafOnly":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_20 := hx_structural_source_18
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_20.__hx_this.leafOnly(resolve, ignored)
						}, func() any {
							return hx_receiver_20.__hx_this
						}, "leafOnly")
					}(), Found: true}
				case "middleOnly":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_21 := hx_structural_source_18
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_21.__hx_this.middleOnly(resolve, ignored)
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "middleOnly")
					}(), Found: true}
				case "total":
					return hxrt.AnonymousField{Value: hx_structural_source_18.total, Found: true}
				case "macroName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_24 := hx_structural_source_18
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_24.__hx_this.macroName(resolve, ignored)
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "macroName")
					}(), Found: true}
				case "describe":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) *string] {
						hx_receiver_25 := hx_structural_source_18
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) *string {
							return hx_receiver_25.__hx_this.describe(resolve, key)
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "describe")
					}(), Found: true}
				case "bump":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) int] {
						hx_receiver_26 := hx_structural_source_18
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) int {
							return hx_receiver_26.__hx_this.bump(resolve, key)
						}, func() any {
							return hx_receiver_26.__hx_this
						}, "bump")
					}(), Found: true}
				case "type":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_27 := hx_structural_source_18
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_27.__hx_this.type_(resolve, ignored)
						}, func() any {
							return hx_receiver_27.__hx_this
						}, "type")
					}(), Found: true}
				case "secret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_28 := hx_structural_source_18
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_28.__hx_this.secret(resolve, ignored)
						}, func() any {
							return hx_receiver_28.__hx_this
						}, "secret")
					}(), Found: true}
				case "retainSecret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_29 := hx_structural_source_18
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_29.__hx_this.retainSecret()
						}, func() any {
							return hx_receiver_29.__hx_this
						}, "retainSecret")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "total":
					hx_structural_source_18.total = func(hx_value_22 any) int {
						if hx_value_22 == nil {
							var hx_zero_23 int
							return hx_zero_23
						}
						return hx_value_22.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *MethodMiddle:
		return func(hx_structural_source_30 *MethodMiddle) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_30, func(name string) hxrt.AnonymousField {
				switch name {
				case "middleOnly":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_31 := hx_structural_source_30
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_31.__hx_this.middleOnly(resolve, ignored)
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "middleOnly")
					}(), Found: true}
				case "total":
					return hxrt.AnonymousField{Value: hx_structural_source_30.total, Found: true}
				case "virtualName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_34 := hx_structural_source_30
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_34.__hx_this.virtualName()
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "virtualName")
					}(), Found: true}
				case "macroName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_35 := hx_structural_source_30
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_35.__hx_this.macroName(resolve, ignored)
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "macroName")
					}(), Found: true}
				case "describe":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) *string] {
						hx_receiver_36 := hx_structural_source_30
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) *string {
							return hx_receiver_36.__hx_this.describe(resolve, key)
						}, func() any {
							return hx_receiver_36.__hx_this
						}, "describe")
					}(), Found: true}
				case "bump":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) int] {
						hx_receiver_37 := hx_structural_source_30
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) int {
							return hx_receiver_37.__hx_this.bump(resolve, key)
						}, func() any {
							return hx_receiver_37.__hx_this
						}, "bump")
					}(), Found: true}
				case "type":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_38 := hx_structural_source_30
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_38.__hx_this.type_(resolve, ignored)
						}, func() any {
							return hx_receiver_38.__hx_this
						}, "type")
					}(), Found: true}
				case "secret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_39 := hx_structural_source_30
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_39.__hx_this.secret(resolve, ignored)
						}, func() any {
							return hx_receiver_39.__hx_this
						}, "secret")
					}(), Found: true}
				case "retainSecret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_40 := hx_structural_source_30
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_40.__hx_this.retainSecret()
						}, func() any {
							return hx_receiver_40.__hx_this
						}, "retainSecret")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "total":
					hx_structural_source_30.total = func(hx_value_32 any) int {
						if hx_value_32 == nil {
							var hx_zero_33 int
							return hx_zero_33
						}
						return hx_value_32.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *SpecializedIterator:
		return func(hx_structural_source_41 *SpecializedIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_41, func(name string) hxrt.AnonymousField {
				switch name {
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_42 := hx_structural_source_41
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_42.__hx_this.next()
						}, func() any {
							return hx_receiver_42.__hx_this
						}, "next")
					}(), Found: true}
				case "values":
					return hxrt.AnonymousField{Value: hx_structural_source_41.values, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_41.index, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_45 := hx_structural_source_41
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_45.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_45.__hx_this
						}, "hasNext")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "index":
					hx_structural_source_41.index = func(hx_value_43 any) int {
						if hx_value_43 == nil {
							var hx_zero_44 int
							return hx_zero_44
						}
						return hx_value_43.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_46 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_46, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_46.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_46.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_46.high = func(hx_value_47 any) int {
						if hx_value_47 == nil {
							var hx_zero_48 int
							return hx_zero_48
						}
						return hx_value_47.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_46.low = func(hx_value_49 any) int {
						if hx_value_49 == nil {
							var hx_zero_50 int
							return hx_zero_50
						}
						return hx_value_49.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Template__TokenCursor:
		return func(hx_structural_source_51 *haxe___Template__TokenCursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_51, func(name string) hxrt.AnonymousField {
				switch name {
				case "tokens":
					return hxrt.AnonymousField{Value: hx_structural_source_51.tokens, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_51.index, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "tokens":
					hx_structural_source_51.tokens = func(hx_value_52 any) *hxrt.Array {
						if hx_value_52 == nil {
							var hx_zero_53 *hxrt.Array
							return hx_zero_53
						}
						return hx_value_52.(*hxrt.Array)
					}(incoming)
					return true
				case "index":
					hx_structural_source_51.index = func(hx_value_54 any) int {
						if hx_value_54 == nil {
							var hx_zero_55 int
							return hx_zero_55
						}
						return hx_value_54.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Template__ExprCursor:
		return func(hx_structural_source_56 *haxe___Template__ExprCursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_56, func(name string) hxrt.AnonymousField {
				switch name {
				case "tokens":
					return hxrt.AnonymousField{Value: hx_structural_source_56.tokens, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_56.index, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "tokens":
					hx_structural_source_56.tokens = func(hx_value_57 any) *hxrt.Array {
						if hx_value_57 == nil {
							var hx_zero_58 *hxrt.Array
							return hx_zero_58
						}
						return hx_value_57.(*hxrt.Array)
					}(incoming)
					return true
				case "index":
					hx_structural_source_56.index = func(hx_value_59 any) int {
						if hx_value_59 == nil {
							var hx_zero_60 int
							return hx_zero_60
						}
						return hx_value_59.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Template:
		return func(hx_structural_source_61 *haxe__Template) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_61, func(name string) hxrt.AnonymousField {
				switch name {
				case "expr":
					return hxrt.AnonymousField{Value: hx_structural_source_61.expr, Found: true}
				case "context":
					return hxrt.AnonymousField{Value: hx_structural_source_61.context, Found: true}
				case "macros":
					return hxrt.AnonymousField{Value: hx_structural_source_61.macros, Found: true}
				case "stack":
					return hxrt.AnonymousField{Value: hx_structural_source_61.stack, Found: true}
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_61.output, Found: true}
				case "execute":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any) *string] {
						hx_receiver_68 := hx_structural_source_61
						return hxrt.BoundFunction(func(context any, macros any) *string {
							return hx_receiver_68.__hx_this.execute(context, macros)
						}, func() any {
							return hx_receiver_68.__hx_this
						}, "execute")
					}(), Found: true}
				case "resolve":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_69 := hx_structural_source_61
						return hxrt.BoundFunction(func(v *string) any {
							return hx_receiver_69.__hx_this.resolve(v)
						}, func() any {
							return hx_receiver_69.__hx_this
						}, "resolve")
					}(), Found: true}
				case "parseTokens":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_70 := hx_structural_source_61
						return hxrt.BoundFunction(func(data *string) *hxrt.Array {
							return hx_receiver_70.__hx_this.parseTokens(data)
						}, func() any {
							return hx_receiver_70.__hx_this
						}, "parseTokens")
					}(), Found: true}
				case "parseBlock":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
						hx_receiver_71 := hx_structural_source_61
						return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
							return hx_receiver_71.__hx_this.parseBlock(cursor)
						}, func() any {
							return hx_receiver_71.__hx_this
						}, "parseBlock")
					}(), Found: true}
				case "parse":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
						hx_receiver_72 := hx_structural_source_61
						return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
							return hx_receiver_72.__hx_this.parse(cursor)
						}, func() any {
							return hx_receiver_72.__hx_this
						}, "parse")
					}(), Found: true}
				case "parseExpr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
						hx_receiver_73 := hx_structural_source_61
						return hxrt.BoundFunction(func(data *string) *hxrt.Function[func() any] {
							return hx_receiver_73.__hx_this.parseExpr(data)
						}, func() any {
							return hx_receiver_73.__hx_this
						}, "parseExpr")
					}(), Found: true}
				case "makeConst":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
						hx_receiver_74 := hx_structural_source_61
						return hxrt.BoundFunction(func(v *string) *hxrt.Function[func() any] {
							return hx_receiver_74.__hx_this.makeConst(v)
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "makeConst")
					}(), Found: true}
				case "makePath":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_75 := hx_structural_source_61
						return hxrt.BoundFunction(func(e *hxrt.Function[func() any], cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_75.__hx_this.makePath(e, cursor)
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "makePath")
					}(), Found: true}
				case "makeExpr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_76 := hx_structural_source_61
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_76.__hx_this.makeExpr(cursor)
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "makeExpr")
					}(), Found: true}
				case "skipSpaces":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor)] {
						hx_receiver_77 := hx_structural_source_61
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) {
							hx_receiver_77.__hx_this.skipSpaces(cursor)
						}, func() any {
							return hx_receiver_77.__hx_this
						}, "skipSpaces")
					}(), Found: true}
				case "makeExpr2":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_78 := hx_structural_source_61
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_78.__hx_this.makeExpr2(cursor)
						}, func() any {
							return hx_receiver_78.__hx_this
						}, "makeExpr2")
					}(), Found: true}
				case "run":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TemplateExpr)] {
						hx_receiver_79 := hx_structural_source_61
						return hxrt.BoundFunction(func(e *haxe___Template__TemplateExpr) {
							hx_receiver_79.__hx_this.run(e)
						}, func() any {
							return hx_receiver_79.__hx_this
						}, "run")
					}(), Found: true}
				case "popStackValue":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_80 := hx_structural_source_61
						return hxrt.BoundFunction(func() any {
							return hx_receiver_80.__hx_this.popStackValue()
						}, func() any {
							return hx_receiver_80.__hx_this
						}, "popStackValue")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "expr":
					hx_structural_source_61.expr = func(hx_value_62 any) *haxe___Template__TemplateExpr {
						if hx_value_62 == nil {
							var hx_zero_63 *haxe___Template__TemplateExpr
							return hx_zero_63
						}
						return hx_value_62.(*haxe___Template__TemplateExpr)
					}(incoming)
					return true
				case "context":
					hx_structural_source_61.context = incoming
					return true
				case "macros":
					hx_structural_source_61.macros = incoming
					return true
				case "stack":
					hx_structural_source_61.stack = func(hx_value_64 any) *hxrt.Array {
						if hx_value_64 == nil {
							var hx_zero_65 *hxrt.Array
							return hx_zero_65
						}
						return hx_value_64.(*hxrt.Array)
					}(incoming)
					return true
				case "output":
					hx_structural_source_61.output = func(hx_value_66 any) *string {
						if hx_value_66 == nil {
							var hx_zero_67 *string
							return hx_zero_67
						}
						return hx_value_66.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_81 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_81, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_81.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_81.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_86 := hx_structural_source_81
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_86.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_86.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_87 := hx_structural_source_81
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_87.__hx_this.next()
						}, func() any {
							return hx_receiver_87.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_81.offset = func(hx_value_82 any) int {
						if hx_value_82 == nil {
							var hx_zero_83 int
							return hx_zero_83
						}
						return hx_value_82.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_81.s = func(hx_value_84 any) *string {
						if hx_value_84 == nil {
							var hx_zero_85 *string
							return hx_zero_85
						}
						return hx_value_84.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_88 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_88, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_88.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_88.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_93 := hx_structural_source_88
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_93.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_93.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_94 := hx_structural_source_88
						return hxrt.BoundFunction(func() int {
							return hx_receiver_94.__hx_this.next()
						}, func() any {
							return hx_receiver_94.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_88.offset = func(hx_value_89 any) int {
						if hx_value_89 == nil {
							var hx_zero_90 int
							return hx_zero_90
						}
						return hx_value_89.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_88.s = func(hx_value_91 any) *string {
						if hx_value_91 == nil {
							var hx_zero_92 *string
							return hx_zero_92
						}
						return hx_value_91.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_95 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_95, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_95.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_95.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_95.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_95.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_98 := hx_structural_source_95
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_98.__hx_this.get_message()
						}, func() any {
							return hx_receiver_98.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_99 := hx_structural_source_95
						return hxrt.BoundFunction(func() any {
							return hx_receiver_99.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_99.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_100 := hx_structural_source_95
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_100.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_100.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_101 := hx_structural_source_95
						return hxrt.BoundFunction(func() any {
							return hx_receiver_101.__hx_this.get_native()
						}, func() any {
							return hx_receiver_101.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_102 := hx_structural_source_95
						return hxrt.BoundFunction(func() any {
							return hx_receiver_102.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_102.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_103 := hx_structural_source_95
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_103.__hx_this.toString()
						}, func() any {
							return hx_receiver_103.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_104 := hx_structural_source_95
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_104.__hx_this.details()
						}, func() any {
							return hx_receiver_104.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_95.savedMessage = func(hx_value_96 any) *string {
						if hx_value_96 == nil {
							var hx_zero_97 *string
							return hx_zero_97
						}
						return hx_value_96.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_95.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_95.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_95.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *EReg:
		return func(hx_structural_source_105 *EReg) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_105, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_105.handle, Found: true}
				case "global":
					return hxrt.AnonymousField{Value: hx_structural_source_105.global, Found: true}
				case "lastSource":
					return hxrt.AnonymousField{Value: hx_structural_source_105.lastSource, Found: true}
				case "lastMatch":
					return hxrt.AnonymousField{Value: hx_structural_source_105.lastMatch, Found: true}
				case "match":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_114 := hx_structural_source_105
						return hxrt.BoundFunction(func(s *string) bool {
							return hx_receiver_114.__hx_this.match(s)
						}, func() any {
							return hx_receiver_114.__hx_this
						}, "match")
					}(), Found: true}
				case "matched":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_115 := hx_structural_source_105
						return hxrt.BoundFunction(func(n int) *string {
							return hx_receiver_115.__hx_this.matched(n)
						}, func() any {
							return hx_receiver_115.__hx_this
						}, "matched")
					}(), Found: true}
				case "matchedLeft":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_116 := hx_structural_source_105
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_116.__hx_this.matchedLeft()
						}, func() any {
							return hx_receiver_116.__hx_this
						}, "matchedLeft")
					}(), Found: true}
				case "matchedRight":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_117 := hx_structural_source_105
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_117.__hx_this.matchedRight()
						}, func() any {
							return hx_receiver_117.__hx_this
						}, "matchedRight")
					}(), Found: true}
				case "matchedPos":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_118 := hx_structural_source_105
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_118.__hx_this.matchedPos()
						}, func() any {
							return hx_receiver_118.__hx_this
						}, "matchedPos")
					}(), Found: true}
				case "matchSub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, int, any) bool] {
						hx_receiver_119 := hx_structural_source_105
						return hxrt.BoundFunction(func(s *string, pos int, len_local any) bool {
							if len_local == nil {
								len_local = -1
							}
							return hx_receiver_119.__hx_this.matchSub(s, pos, func(hx_value_120 any) int {
								if hx_value_120 == nil {
									var hx_zero_121 int
									return hx_zero_121
								}
								return hx_value_120.(int)
							}(len_local))
						}, func() any {
							return hx_receiver_119.__hx_this
						}, "matchSub")
					}(), Found: true}
				case "split":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_122 := hx_structural_source_105
						return hxrt.BoundFunction(func(s *string) *hxrt.Array {
							return hx_receiver_122.__hx_this.split(s)
						}, func() any {
							return hx_receiver_122.__hx_this
						}, "split")
					}(), Found: true}
				case "replace":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string) *string] {
						hx_receiver_123 := hx_structural_source_105
						return hxrt.BoundFunction(func(s *string, by *string) *string {
							return hx_receiver_123.__hx_this.replace(s, by)
						}, func() any {
							return hx_receiver_123.__hx_this
						}, "replace")
					}(), Found: true}
				case "map":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string] {
						hx_receiver_124 := hx_structural_source_105
						return hxrt.BoundFunction(func(s *string, f *hxrt.Function[func(*EReg) *string]) *string {
							return hx_receiver_124.__hx_this.map_(s, f)
						}, func() any {
							return hx_receiver_124.__hx_this
						}, "map")
					}(), Found: true}
				case "remember":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *hxrt.RegexMatch)] {
						hx_receiver_125 := hx_structural_source_105
						return hxrt.BoundFunction(func(source *string, match *hxrt.RegexMatch) {
							hx_receiver_125.__hx_this.remember(source, match)
						}, func() any {
							return hx_receiver_125.__hx_this
						}, "remember")
					}(), Found: true}
				case "requireMatch":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.RegexMatch] {
						hx_receiver_126 := hx_structural_source_105
						return hxrt.BoundFunction(func() *hxrt.RegexMatch {
							return hx_receiver_126.__hx_this.requireMatch()
						}, func() any {
							return hx_receiver_126.__hx_this
						}, "requireMatch")
					}(), Found: true}
				case "expandReplacement":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *hxrt.RegexMatch) *string] {
						hx_receiver_127 := hx_structural_source_105
						return hxrt.BoundFunction(func(by *string, source *string, currentMatch *hxrt.RegexMatch) *string {
							return hx_receiver_127.__hx_this.expandReplacement(by, source, currentMatch)
						}, func() any {
							return hx_receiver_127.__hx_this
						}, "expandReplacement")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_105.handle = func(hx_value_106 any) *hxrt.RegexHandle {
						if hx_value_106 == nil {
							var hx_zero_107 *hxrt.RegexHandle
							return hx_zero_107
						}
						return hx_value_106.(*hxrt.RegexHandle)
					}(incoming)
					return true
				case "global":
					hx_structural_source_105.global = func(hx_value_108 any) bool {
						if hx_value_108 == nil {
							var hx_zero_109 bool
							return hx_zero_109
						}
						return hx_value_108.(bool)
					}(incoming)
					return true
				case "lastSource":
					hx_structural_source_105.lastSource = func(hx_value_110 any) *string {
						if hx_value_110 == nil {
							var hx_zero_111 *string
							return hx_zero_111
						}
						return hx_value_110.(*string)
					}(incoming)
					return true
				case "lastMatch":
					hx_structural_source_105.lastMatch = func(hx_value_112 any) *hxrt.RegexMatch {
						if hx_value_112 == nil {
							var hx_zero_113 *hxrt.RegexMatch
							return hx_zero_113
						}
						return hx_value_112.(*hxrt.RegexMatch)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_128 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_128, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_128.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_129 := hx_structural_source_128
						return hxrt.BoundFunction(func() any {
							return hx_receiver_129.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_129.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_128.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_128.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_128.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_128.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_132 := hx_structural_source_128
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_132.__hx_this.get_message()
						}, func() any {
							return hx_receiver_132.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_133 := hx_structural_source_128
						return hxrt.BoundFunction(func() any {
							return hx_receiver_133.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_133.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_134 := hx_structural_source_128
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_134.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_134.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_135 := hx_structural_source_128
						return hxrt.BoundFunction(func() any {
							return hx_receiver_135.__hx_this.get_native()
						}, func() any {
							return hx_receiver_135.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_136 := hx_structural_source_128
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_136.__hx_this.toString()
						}, func() any {
							return hx_receiver_136.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_137 := hx_structural_source_128
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_137.__hx_this.details()
						}, func() any {
							return hx_receiver_137.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_128.savedMessage = func(hx_value_130 any) *string {
						if hx_value_130 == nil {
							var hx_zero_131 *string
							return hx_zero_131
						}
						return hx_value_130.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_128.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_128.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_128.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_138 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_138, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_138.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_138.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_138.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_138.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_138.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_147 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_147.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_147.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_148 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_148.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_148.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_149 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len_local int) {
							hx_receiver_149.__hx_this.blit(pos, src, srcpos, len_local)
						}, func() any {
							return hx_receiver_149.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_150 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, len_local int, value int) {
							hx_receiver_150.__hx_this.fill(pos, len_local, value)
						}, func() any {
							return hx_receiver_150.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_151 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, len_local int) *haxe__io__Bytes {
							return hx_receiver_151.__hx_this.sub(pos, len_local)
						}, func() any {
							return hx_receiver_151.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_152 := hx_structural_source_138
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_152.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_152.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_153 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_153.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_153.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_154 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_154.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_154.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_155 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_155.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_155.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_156 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_156.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_156.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_157 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_157.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_157.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_158 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_158.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_158.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_159 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_159.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_159.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_160 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_160.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_160.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_161 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_161.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_161.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_162 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_162.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_162.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_163 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, len_local int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_163.__hx_this.getString(pos, len_local, encoding)
						}, func() any {
							return hx_receiver_163.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_164 := hx_structural_source_138
						return hxrt.BoundFunction(func(pos int, len_local int) *string {
							return hx_receiver_164.__hx_this.readString(pos, len_local)
						}, func() any {
							return hx_receiver_164.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_165 := hx_structural_source_138
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_165.__hx_this.toString()
						}, func() any {
							return hx_receiver_165.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_166 := hx_structural_source_138
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_166.__hx_this.toHex()
						}, func() any {
							return hx_receiver_166.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_167 := hx_structural_source_138
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_167.__hx_this.getData()
						}, func() any {
							return hx_receiver_167.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_168 := hx_structural_source_138
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_168.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_168.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_138.b = func(hx_value_139 any) []int {
						if hx_value_139 == nil {
							var hx_zero_140 []int
							return hx_zero_140
						}
						return hx_value_139.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_138.__hx_raw = func(hx_value_141 any) *hxrt.ByteView {
						if hx_value_141 == nil {
							var hx_zero_142 *hxrt.ByteView
							return hx_zero_142
						}
						return hx_value_141.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_138.__hx_rawValid = func(hx_value_143 any) bool {
						if hx_value_143 == nil {
							var hx_zero_144 bool
							return hx_zero_144
						}
						return hx_value_143.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_138.__hx_dataExposed = func(hx_value_145 any) bool {
						if hx_value_145 == nil {
							var hx_zero_146 bool
							return hx_zero_146
						}
						return hx_value_145.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
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
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Template:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
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
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__Template:
		return hxrt__generated_method_field__haxe__Template(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
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
		return func() *hxrt.Function[func() *SpecializedIterator] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *SpecializedIterator {
				return hx_receiver_1.__hx_this.iterator()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "iterator")
		}()
	}
	return nil
}

func hxrt__generated_method_field__ConcreteIterator(value *ConcreteIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__EReg(value *EReg, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "expandReplacement":
		return func() *hxrt.Function[func(*string, *string, *hxrt.RegexMatch) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(by *string, source *string, currentMatch *hxrt.RegexMatch) *string {
				return hx_receiver_1.__hx_this.expandReplacement(by, source, currentMatch)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "expandReplacement")
		}()
	case "map":
		return func() *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(s *string, f *hxrt.Function[func(*EReg) *string]) *string {
				return hx_receiver_1.__hx_this.map_(s, f)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "map")
		}()
	case "match":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(s *string) bool {
				return hx_receiver_1.__hx_this.match(s)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "match")
		}()
	case "matchSub":
		return func() *hxrt.Function[func(*string, int, any) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(s *string, pos int, len_local any) bool {
				if len_local == nil {
					len_local = -1
				}
				return hx_receiver_1.__hx_this.matchSub(s, pos, func(hx_value_2 any) int {
					if hx_value_2 == nil {
						var hx_zero_3 int
						return hx_zero_3
					}
					return hx_value_2.(int)
				}(len_local))
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "matchSub")
		}()
	case "matched":
		return func() *hxrt.Function[func(int) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(n int) *string {
				return hx_receiver_1.__hx_this.matched(n)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "matched")
		}()
	case "matchedLeft":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.matchedLeft()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "matchedLeft")
		}()
	case "matchedPos":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.matchedPos()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "matchedPos")
		}()
	case "matchedRight":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.matchedRight()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "matchedRight")
		}()
	case "remember":
		return func() *hxrt.Function[func(*string, *hxrt.RegexMatch)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(source *string, match *hxrt.RegexMatch) {
				hx_receiver_1.__hx_this.remember(source, match)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "remember")
		}()
	case "replace":
		return func() *hxrt.Function[func(*string, *string) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(s *string, by *string) *string {
				return hx_receiver_1.__hx_this.replace(s, by)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "replace")
		}()
	case "requireMatch":
		return func() *hxrt.Function[func() *hxrt.RegexMatch] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *hxrt.RegexMatch {
				return hx_receiver_1.__hx_this.requireMatch()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "requireMatch")
		}()
	case "split":
		return func() *hxrt.Function[func(*string) *hxrt.Array] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(s *string) *hxrt.Array {
				return hx_receiver_1.__hx_this.split(s)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "split")
		}()
	}
	return nil
}

func hxrt__generated_method_field__MethodBase(value *MethodBase, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "bump":
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) int {
				return hx_receiver_1.__hx_this.bump(resolve, key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "bump")
		}()
	case "describe":
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) *string {
				return hx_receiver_1.__hx_this.describe(resolve, key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "describe")
		}()
	case "macroName":
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
				return hx_receiver_1.__hx_this.macroName(resolve, ignored)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "macroName")
		}()
	case "retainSecret":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.retainSecret()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "retainSecret")
		}()
	case "secret":
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
				return hx_receiver_1.__hx_this.secret(resolve, ignored)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "secret")
		}()
	case "type":
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
				return hx_receiver_1.__hx_this.type_(resolve, ignored)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "type")
		}()
	case "virtualName":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.virtualName()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "virtualName")
		}()
	}
	return nil
}

func hxrt__generated_method_field__MethodLeaf(value *MethodLeaf, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "leafOnly":
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
				return hx_receiver_1.__hx_this.leafOnly(resolve, ignored)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "leafOnly")
		}()
	case "virtualName":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.virtualName()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "virtualName")
		}()
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
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
				return hx_receiver_1.__hx_this.middleOnly(resolve, ignored)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "middleOnly")
		}()
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
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	if value.ConcreteIterator == nil {
		return nil
	}
	return hxrt__generated_method_field__ConcreteIterator(value.ConcreteIterator, key)
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.details()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.get_message()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_native()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_1.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Template(value *haxe__Template, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "execute":
		return func() *hxrt.Function[func(any, any) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(context any, macros any) *string {
				return hx_receiver_1.__hx_this.execute(context, macros)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "execute")
		}()
	case "makeConst":
		return func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(v *string) *hxrt.Function[func() any] {
				return hx_receiver_1.__hx_this.makeConst(v)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "makeConst")
		}()
	case "makeExpr":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_1.__hx_this.makeExpr(cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "makeExpr")
		}()
	case "makeExpr2":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_1.__hx_this.makeExpr2(cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "makeExpr2")
		}()
	case "makePath":
		return func() *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(e *hxrt.Function[func() any], cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_1.__hx_this.makePath(e, cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "makePath")
		}()
	case "parse":
		return func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
				return hx_receiver_1.__hx_this.parse(cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "parse")
		}()
	case "parseBlock":
		return func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
				return hx_receiver_1.__hx_this.parseBlock(cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "parseBlock")
		}()
	case "parseExpr":
		return func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(data *string) *hxrt.Function[func() any] {
				return hx_receiver_1.__hx_this.parseExpr(data)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "parseExpr")
		}()
	case "parseTokens":
		return func() *hxrt.Function[func(*string) *hxrt.Array] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(data *string) *hxrt.Array {
				return hx_receiver_1.__hx_this.parseTokens(data)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "parseTokens")
		}()
	case "popStackValue":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.popStackValue()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "popStackValue")
		}()
	case "resolve":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(v *string) any {
				return hx_receiver_1.__hx_this.resolve(v)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "resolve")
		}()
	case "run":
		return func() *hxrt.Function[func(*haxe___Template__TemplateExpr)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(e *haxe___Template__TemplateExpr) {
				hx_receiver_1.__hx_this.run(e)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "run")
		}()
	case "skipSpaces":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) {
				hx_receiver_1.__hx_this.skipSpaces(cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "skipSpaces")
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
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_nativeView":
		return func() *hxrt.Function[func() *hxrt.ByteView] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *hxrt.ByteView {
				return hx_receiver_1.__hx_this.__hx_nativeView()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "__hx_nativeView")
		}()
	case "blit":
		return func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len_local int) {
				hx_receiver_1.__hx_this.blit(pos, src, srcpos, len_local)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "blit")
		}()
	case "compare":
		return func() *hxrt.Function[func(*haxe__io__Bytes) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
				return hx_receiver_1.__hx_this.compare(other)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "compare")
		}()
	case "fill":
		return func() *hxrt.Function[func(int, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int, value int) {
				hx_receiver_1.__hx_this.fill(pos, len_local, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "fill")
		}()
	case "get":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.get(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get")
		}()
	case "getData":
		return func() *hxrt.Function[func() []int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() []int {
				return hx_receiver_1.__hx_this.getData()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getData")
		}()
	case "getDouble":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_1.__hx_this.getDouble(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getDouble")
		}()
	case "getFloat":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_1.__hx_this.getFloat(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getFloat")
		}()
	case "getInt32":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.getInt32(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getInt32")
		}()
	case "getInt64":
		return func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
				return hx_receiver_1.__hx_this.getInt64(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getInt64")
		}()
	case "getString":
		return func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_1.__hx_this.getString(pos, len_local, encoding)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getString")
		}()
	case "getUInt16":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.getUInt16(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getUInt16")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, int) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int) *string {
				return hx_receiver_1.__hx_this.readString(pos, len_local)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readString")
		}()
	case "set":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.set(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set")
		}()
	case "setDouble":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_1.__hx_this.setDouble(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setDouble")
		}()
	case "setFloat":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_1.__hx_this.setFloat(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setFloat")
		}()
	case "setInt32":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.setInt32(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setInt32")
		}()
	case "setInt64":
		return func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
				hx_receiver_1.__hx_this.setInt64(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setInt64")
		}()
	case "setUInt16":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.setUInt16(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setUInt16")
		}()
	case "sub":
		return func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int) *haxe__io__Bytes {
				return hx_receiver_1.__hx_this.sub(pos, len_local)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "sub")
		}()
	case "toHex":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toHex()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toHex")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
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
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
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
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
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
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Template:
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
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__Template:
		return hxrt__generated_field_lookup__haxe__Template(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Template:
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
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__Template:
		return hxrt__generated_field_has__haxe__Template(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Template:
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
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__Template:
		return hxrt__generated_field_set__haxe__Template(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Template:
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
		return hxrt.NewArray(hxrt.StringFromLiteral("values"), hxrt.StringFromLiteral("index"))
	case *EReg:
		return hxrt.NewArray(hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("global"), hxrt.StringFromLiteral("lastSource"), hxrt.StringFromLiteral("lastMatch"))
	case *MethodBase:
		return hxrt.NewArray(hxrt.StringFromLiteral("total"))
	case *MethodLeaf:
		return hxrt.NewArray(hxrt.StringFromLiteral("total"))
	case *MethodMiddle:
		return hxrt.NewArray(hxrt.StringFromLiteral("total"))
	case *SpecializedIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("values"), hxrt.StringFromLiteral("index"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__Template:
		return hxrt.NewArray(hxrt.StringFromLiteral("expr"), hxrt.StringFromLiteral("context"), hxrt.StringFromLiteral("macros"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("output"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe___Template__ExprCursor:
		return hxrt.NewArray(hxrt.StringFromLiteral("tokens"), hxrt.StringFromLiteral("index"))
	case *haxe___Template__TokenCursor:
		return hxrt.NewArray(hxrt.StringFromLiteral("tokens"), hxrt.StringFromLiteral("index"))
	case *haxe__io__Bytes:
		return hxrt.NewArray(hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("__hx_dataExposed"))
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
		return hxrt.ReadStoredField(value.values)
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
		return hxrt.ReadStoredField(value.index)
	case "values":
		return hxrt.ReadStoredField(value.values)
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
		return hxrt.ReadStoredField(value.global)
	case "handle":
		return hxrt.ReadStoredField(value.handle)
	case "lastMatch":
		return hxrt.ReadStoredField(value.lastMatch)
	case "lastSource":
		return hxrt.ReadStoredField(value.lastSource)
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
		return hxrt.ReadStoredField(value.total)
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

func hxrt__generated_field_lookup__haxe__Template(value *haxe__Template, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "context":
		return hxrt.ReadStoredField(value.context)
	case "expr":
		return hxrt.ReadStoredField(value.expr)
	case "macros":
		return hxrt.ReadStoredField(value.macros)
	case "output":
		return hxrt.ReadStoredField(value.output)
	case "stack":
		return hxrt.ReadStoredField(value.stack)
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

func hxrt__generated_field_lookup__haxe___Template__ExprCursor(value *haxe___Template__ExprCursor, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "index":
		return hxrt.ReadStoredField(value.index)
	case "tokens":
		return hxrt.ReadStoredField(value.tokens)
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
		return hxrt.ReadStoredField(value.index)
	case "tokens":
		return hxrt.ReadStoredField(value.tokens)
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

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(argument_0))
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *SpecializedIterator]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_d7744367(argument_0), argument_1)
	case *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_d7744367(argument_0), func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(argument_1))
	case *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_d7744367(argument_0), func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(argument_1))
	case *hxrt.Function[func(any, any) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1)
	case *hxrt.Function[func(*string) *hxrt.Array]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(argument_0))
	case *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__TokenCursor(argument_0))
	case *hxrt.Function[func(*string) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
		}(argument_0))
	case *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_c5309b43(argument_0), hx_recover_class_haxe___Template__ExprCursor(argument_1))
	case *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__ExprCursor(argument_0))
	case *hxrt.Function[func(*haxe___Template__ExprCursor)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__ExprCursor(argument_0))
		return nil
	case *hxrt.Function[func(*haxe___Template__TemplateExpr)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_11 any) *haxe___Template__TemplateExpr {
			if hx_value_11 == nil {
				var hx_zero_12 *haxe___Template__TemplateExpr
				return hx_zero_12
			}
			return hx_value_11.(*haxe___Template__TemplateExpr)
		}(argument_0))
		return nil
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(argument_0))
	case *hxrt.Function[func(int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_15 any) int {
			if hx_value_15 == nil {
				var hx_zero_16 int
				return hx_zero_16
			}
			return hx_value_15.(int)
		}(argument_0))
	case *hxrt.Function[func(*string, int, any) bool]:
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
		return hxrt.FunctionCall(source)(func(hx_value_17 any) *string {
			if hx_value_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_value_17.(*string)
		}(argument_0), func(hx_value_19 any) int {
			if hx_value_19 == nil {
				var hx_zero_20 int
				return hx_zero_20
			}
			return hx_value_19.(int)
		}(argument_1), argument_2)
	case *hxrt.Function[func(*string, *string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_21 any) *string {
			if hx_value_21 == nil {
				var hx_zero_22 *string
				return hx_zero_22
			}
			return hx_value_21.(*string)
		}(argument_0), func(hx_value_23 any) *string {
			if hx_value_23 == nil {
				var hx_zero_24 *string
				return hx_zero_24
			}
			return hx_value_23.(*string)
		}(argument_1))
	case *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_25 any) *string {
			if hx_value_25 == nil {
				var hx_zero_26 *string
				return hx_zero_26
			}
			return hx_value_25.(*string)
		}(argument_0), hx_recover_callable_97658df1(argument_1))
	case *hxrt.Function[func(*string, *hxrt.RegexMatch)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_27 any) *string {
			if hx_value_27 == nil {
				var hx_zero_28 *string
				return hx_zero_28
			}
			return hx_value_27.(*string)
		}(argument_0), func(hx_value_29 any) *hxrt.RegexMatch {
			if hx_value_29 == nil {
				var hx_zero_30 *hxrt.RegexMatch
				return hx_zero_30
			}
			return hx_value_29.(*hxrt.RegexMatch)
		}(argument_1))
		return nil
	case *hxrt.Function[func() *hxrt.RegexMatch]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string, *string, *hxrt.RegexMatch) *string]:
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
		return hxrt.FunctionCall(source)(func(hx_value_31 any) *string {
			if hx_value_31 == nil {
				var hx_zero_32 *string
				return hx_zero_32
			}
			return hx_value_31.(*string)
		}(argument_0), func(hx_value_33 any) *string {
			if hx_value_33 == nil {
				var hx_zero_34 *string
				return hx_zero_34
			}
			return hx_value_33.(*string)
		}(argument_1), func(hx_value_35 any) *hxrt.RegexMatch {
			if hx_value_35 == nil {
				var hx_zero_36 *hxrt.RegexMatch
				return hx_zero_36
			}
			return hx_value_35.(*hxrt.RegexMatch)
		}(argument_2))
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_37 any) int {
			if hx_value_37 == nil {
				var hx_zero_38 int
				return hx_zero_38
			}
			return hx_value_37.(int)
		}(argument_0))
	case *hxrt.Function[func(int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_39 any) int {
			if hx_value_39 == nil {
				var hx_zero_40 int
				return hx_zero_40
			}
			return hx_value_39.(int)
		}(argument_0), func(hx_value_41 any) int {
			if hx_value_41 == nil {
				var hx_zero_42 int
				return hx_zero_42
			}
			return hx_value_41.(int)
		}(argument_1))
		return nil
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
		hxrt.FunctionCall(source)(func(hx_value_43 any) int {
			if hx_value_43 == nil {
				var hx_zero_44 int
				return hx_zero_44
			}
			return hx_value_43.(int)
		}(argument_0), hx_recover_class_haxe__io__Bytes(argument_1), func(hx_value_45 any) int {
			if hx_value_45 == nil {
				var hx_zero_46 int
				return hx_zero_46
			}
			return hx_value_45.(int)
		}(argument_2), func(hx_value_47 any) int {
			if hx_value_47 == nil {
				var hx_zero_48 int
				return hx_zero_48
			}
			return hx_value_47.(int)
		}(argument_3))
		return nil
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
		hxrt.FunctionCall(source)(func(hx_value_49 any) int {
			if hx_value_49 == nil {
				var hx_zero_50 int
				return hx_zero_50
			}
			return hx_value_49.(int)
		}(argument_0), func(hx_value_51 any) int {
			if hx_value_51 == nil {
				var hx_zero_52 int
				return hx_zero_52
			}
			return hx_value_51.(int)
		}(argument_1), func(hx_value_53 any) int {
			if hx_value_53 == nil {
				var hx_zero_54 int
				return hx_zero_54
			}
			return hx_value_53.(int)
		}(argument_2))
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
		return hxrt.FunctionCall(source)(func(hx_value_55 any) int {
			if hx_value_55 == nil {
				var hx_zero_56 int
				return hx_zero_56
			}
			return hx_value_55.(int)
		}(argument_0), func(hx_value_57 any) int {
			if hx_value_57 == nil {
				var hx_zero_58 int
				return hx_zero_58
			}
			return hx_value_57.(int)
		}(argument_1))
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(argument_0))
	case *hxrt.Function[func(int) float64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_59 any) int {
			if hx_value_59 == nil {
				var hx_zero_60 int
				return hx_zero_60
			}
			return hx_value_59.(int)
		}(argument_0))
	case *hxrt.Function[func(int, float64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_61 any) int {
			if hx_value_61 == nil {
				var hx_zero_62 int
				return hx_zero_62
			}
			return hx_value_61.(int)
		}(argument_0), func(hx_value_63 any) float64 {
			if hx_value_63 == nil {
				var hx_zero_64 float64
				return hx_zero_64
			}
			return hx_value_63.(float64)
		}(argument_1))
		return nil
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_65 any) int {
			if hx_value_65 == nil {
				var hx_zero_66 int
				return hx_zero_66
			}
			return hx_value_65.(int)
		}(argument_0))
	case *hxrt.Function[func(int, *haxe___Int64_____Int64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_67 any) int {
			if hx_value_67 == nil {
				var hx_zero_68 int
				return hx_zero_68
			}
			return hx_value_67.(int)
		}(argument_0), hx_recover_class_haxe___Int64_____Int64(argument_1))
		return nil
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
		return hxrt.FunctionCall(source)(func(hx_value_69 any) int {
			if hx_value_69 == nil {
				var hx_zero_70 int
				return hx_zero_70
			}
			return hx_value_69.(int)
		}(argument_0), func(hx_value_71 any) int {
			if hx_value_71 == nil {
				var hx_zero_72 int
				return hx_zero_72
			}
			return hx_value_71.(int)
		}(argument_1), func(hx_value_73 any) *haxe__io__Encoding {
			if hx_value_73 == nil {
				var hx_zero_74 *haxe__io__Encoding
				return hx_zero_74
			}
			return hx_value_73.(*haxe__io__Encoding)
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
		return hxrt.FunctionCall(source)(func(hx_value_75 any) int {
			if hx_value_75 == nil {
				var hx_zero_76 int
				return hx_zero_76
			}
			return hx_value_75.(int)
		}(argument_0), func(hx_value_77 any) int {
			if hx_value_77 == nil {
				var hx_zero_78 int
				return hx_zero_78
			}
			return hx_value_77.(int)
		}(argument_1))
	case *hxrt.Function[func() []int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_97658df1(value any) *hxrt.Function[func(*EReg) *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*EReg) *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__TokenCursor(arg_0)))
		})
	case *hxrt.Function[func(*string) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__ExprCursor(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor)]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__ExprCursor(arg_0))
			return func(hx_value_19 any) *string {
				if hx_value_19 == nil {
					var hx_zero_20 *string
					return hx_zero_20
				}
				return hx_value_19.(*string)
			}(nil)
		})
	case *hxrt.Function[func(*haxe___Template__TemplateExpr)]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			hxrt.FunctionCall(source)(func(hx_value_21 any) *haxe___Template__TemplateExpr {
				if hx_value_21 == nil {
					var hx_zero_22 *haxe___Template__TemplateExpr
					return hx_zero_22
				}
				return hx_value_21.(*haxe___Template__TemplateExpr)
			}(arg_0))
			return func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(nil)
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_27 any) *string {
				if hx_value_27 == nil {
					var hx_zero_28 *string
					return hx_zero_28
				}
				return hx_value_27.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_25 any) *string {
				if hx_value_25 == nil {
					var hx_zero_26 *string
					return hx_zero_26
				}
				return hx_value_25.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return hxrt.FunctionCall(source)(func(hx_value_29 any) int {
				if hx_value_29 == nil {
					var hx_zero_30 int
					return hx_zero_30
				}
				return hx_value_29.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_33 any) *string {
				if hx_value_33 == nil {
					var hx_zero_34 *string
					return hx_zero_34
				}
				return hx_value_33.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_31 any) int {
				if hx_value_31 == nil {
					var hx_zero_32 int
					return hx_zero_32
				}
				return hx_value_31.(int)
			}(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_35 any) *string {
				if hx_value_35 == nil {
					var hx_zero_36 *string
					return hx_zero_36
				}
				return hx_value_35.(*string)
			}(hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0)))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_39 any) *string {
				if hx_value_39 == nil {
					var hx_zero_40 *string
					return hx_zero_40
				}
				return hx_value_39.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_37 any) int {
				if hx_value_37 == nil {
					var hx_zero_38 int
					return hx_zero_38
				}
				return hx_value_37.(int)
			}(arg_0)))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_43 any) *string {
				if hx_value_43 == nil {
					var hx_zero_44 *string
					return hx_zero_44
				}
				return hx_value_43.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_41 any) int {
				if hx_value_41 == nil {
					var hx_zero_42 int
					return hx_zero_42
				}
				return hx_value_41.(int)
			}(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_c5309b43(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *SpecializedIterator]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.RegexMatch]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_d7744367(value any) *hxrt.Function[func(*string) any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*string) any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string) any]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__TokenCursor(arg_0))
		})
	case *hxrt.Function[func(*string) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__ExprCursor(arg_0))
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__ExprCursor(arg_0))
			return nil
		})
	case *hxrt.Function[func(*haxe___Template__TemplateExpr)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			hxrt.FunctionCall(source)(func(hx_value_1 any) *haxe___Template__TemplateExpr {
				if hx_value_1 == nil {
					var hx_zero_2 *haxe___Template__TemplateExpr
					return hx_zero_2
				}
				return hx_value_1.(*haxe___Template__TemplateExpr)
			}(arg_0))
			return nil
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
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

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_recover_class_haxe___Int64_____Int64(value any) *haxe___Int64_____Int64 {
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

func hx_recover_class_haxe___Template__ExprCursor(value any) *haxe___Template__ExprCursor {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Template__ExprCursor:
		return source
	default:
		return value.(*haxe___Template__ExprCursor)
	}
}

func hx_recover_class_haxe___Template__TokenCursor(value any) *haxe___Template__TokenCursor {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Template__TokenCursor:
		return source
	default:
		return value.(*haxe___Template__TokenCursor)
	}
}

func hx_recover_class_haxe__io__Bytes(value any) *haxe__io__Bytes {
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

func hx_is_nominal_haxe__Exception(value any) bool {
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

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
