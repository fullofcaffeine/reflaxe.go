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

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *ConcreteIterable:
		return func(hx_structural_source_376 *ConcreteIterable) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_376, func(name string) hxrt.AnonymousField {
				switch name {
				case "values":
					return hxrt.AnonymousField{Value: hx_structural_source_376.values, Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *SpecializedIterator] {
						hx_receiver_377 := hx_structural_source_376
						return hxrt.BoundFunction(func() *SpecializedIterator {
							return hx_receiver_377.__hx_this.iterator()
						}, func() any {
							return hx_receiver_377.__hx_this
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
		return func(hx_structural_source_378 *ConcreteIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_378, func(name string) hxrt.AnonymousField {
				switch name {
				case "values":
					return hxrt.AnonymousField{Value: hx_structural_source_378.values, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_378.index, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_381 := hx_structural_source_378
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_381.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_381.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_382 := hx_structural_source_378
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_382.__hx_this.next()
						}, func() any {
							return hx_receiver_382.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "index":
					hx_structural_source_378.index = func(hx_value_379 any) int {
						if hx_value_379 == nil {
							var hx_zero_380 int
							return hx_zero_380
						}
						return hx_value_379.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *MethodBase:
		return func(hx_structural_source_383 *MethodBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_383, func(name string) hxrt.AnonymousField {
				switch name {
				case "total":
					return hxrt.AnonymousField{Value: hx_structural_source_383.total, Found: true}
				case "virtualName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_386 := hx_structural_source_383
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_386.__hx_this.virtualName()
						}, func() any {
							return hx_receiver_386.__hx_this
						}, "virtualName")
					}(), Found: true}
				case "macroName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_387 := hx_structural_source_383
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_387.__hx_this.macroName(resolve, ignored)
						}, func() any {
							return hx_receiver_387.__hx_this
						}, "macroName")
					}(), Found: true}
				case "describe":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) *string] {
						hx_receiver_388 := hx_structural_source_383
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) *string {
							return hx_receiver_388.__hx_this.describe(resolve, key)
						}, func() any {
							return hx_receiver_388.__hx_this
						}, "describe")
					}(), Found: true}
				case "bump":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) int] {
						hx_receiver_389 := hx_structural_source_383
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) int {
							return hx_receiver_389.__hx_this.bump(resolve, key)
						}, func() any {
							return hx_receiver_389.__hx_this
						}, "bump")
					}(), Found: true}
				case "type":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_390 := hx_structural_source_383
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_390.__hx_this.type_(resolve, ignored)
						}, func() any {
							return hx_receiver_390.__hx_this
						}, "type")
					}(), Found: true}
				case "secret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_391 := hx_structural_source_383
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_391.__hx_this.secret(resolve, ignored)
						}, func() any {
							return hx_receiver_391.__hx_this
						}, "secret")
					}(), Found: true}
				case "retainSecret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_392 := hx_structural_source_383
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_392.__hx_this.retainSecret()
						}, func() any {
							return hx_receiver_392.__hx_this
						}, "retainSecret")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "total":
					hx_structural_source_383.total = func(hx_value_384 any) int {
						if hx_value_384 == nil {
							var hx_zero_385 int
							return hx_zero_385
						}
						return hx_value_384.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *MethodLeaf:
		return func(hx_structural_source_393 *MethodLeaf) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_393, func(name string) hxrt.AnonymousField {
				switch name {
				case "virtualName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_394 := hx_structural_source_393
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_394.__hx_this.virtualName()
						}, func() any {
							return hx_receiver_394.__hx_this
						}, "virtualName")
					}(), Found: true}
				case "leafOnly":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_395 := hx_structural_source_393
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_395.__hx_this.leafOnly(resolve, ignored)
						}, func() any {
							return hx_receiver_395.__hx_this
						}, "leafOnly")
					}(), Found: true}
				case "middleOnly":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_396 := hx_structural_source_393
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_396.__hx_this.middleOnly(resolve, ignored)
						}, func() any {
							return hx_receiver_396.__hx_this
						}, "middleOnly")
					}(), Found: true}
				case "total":
					return hxrt.AnonymousField{Value: hx_structural_source_393.total, Found: true}
				case "macroName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_399 := hx_structural_source_393
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_399.__hx_this.macroName(resolve, ignored)
						}, func() any {
							return hx_receiver_399.__hx_this
						}, "macroName")
					}(), Found: true}
				case "describe":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) *string] {
						hx_receiver_400 := hx_structural_source_393
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) *string {
							return hx_receiver_400.__hx_this.describe(resolve, key)
						}, func() any {
							return hx_receiver_400.__hx_this
						}, "describe")
					}(), Found: true}
				case "bump":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) int] {
						hx_receiver_401 := hx_structural_source_393
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) int {
							return hx_receiver_401.__hx_this.bump(resolve, key)
						}, func() any {
							return hx_receiver_401.__hx_this
						}, "bump")
					}(), Found: true}
				case "type":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_402 := hx_structural_source_393
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_402.__hx_this.type_(resolve, ignored)
						}, func() any {
							return hx_receiver_402.__hx_this
						}, "type")
					}(), Found: true}
				case "secret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_403 := hx_structural_source_393
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_403.__hx_this.secret(resolve, ignored)
						}, func() any {
							return hx_receiver_403.__hx_this
						}, "secret")
					}(), Found: true}
				case "retainSecret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_404 := hx_structural_source_393
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_404.__hx_this.retainSecret()
						}, func() any {
							return hx_receiver_404.__hx_this
						}, "retainSecret")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "total":
					hx_structural_source_393.total = func(hx_value_397 any) int {
						if hx_value_397 == nil {
							var hx_zero_398 int
							return hx_zero_398
						}
						return hx_value_397.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *MethodMiddle:
		return func(hx_structural_source_405 *MethodMiddle) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_405, func(name string) hxrt.AnonymousField {
				switch name {
				case "middleOnly":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_406 := hx_structural_source_405
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_406.__hx_this.middleOnly(resolve, ignored)
						}, func() any {
							return hx_receiver_406.__hx_this
						}, "middleOnly")
					}(), Found: true}
				case "total":
					return hxrt.AnonymousField{Value: hx_structural_source_405.total, Found: true}
				case "virtualName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_409 := hx_structural_source_405
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_409.__hx_this.virtualName()
						}, func() any {
							return hx_receiver_409.__hx_this
						}, "virtualName")
					}(), Found: true}
				case "macroName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_410 := hx_structural_source_405
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_410.__hx_this.macroName(resolve, ignored)
						}, func() any {
							return hx_receiver_410.__hx_this
						}, "macroName")
					}(), Found: true}
				case "describe":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) *string] {
						hx_receiver_411 := hx_structural_source_405
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) *string {
							return hx_receiver_411.__hx_this.describe(resolve, key)
						}, func() any {
							return hx_receiver_411.__hx_this
						}, "describe")
					}(), Found: true}
				case "bump":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) int] {
						hx_receiver_412 := hx_structural_source_405
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) int {
							return hx_receiver_412.__hx_this.bump(resolve, key)
						}, func() any {
							return hx_receiver_412.__hx_this
						}, "bump")
					}(), Found: true}
				case "type":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_413 := hx_structural_source_405
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_413.__hx_this.type_(resolve, ignored)
						}, func() any {
							return hx_receiver_413.__hx_this
						}, "type")
					}(), Found: true}
				case "secret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
						hx_receiver_414 := hx_structural_source_405
						return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
							return hx_receiver_414.__hx_this.secret(resolve, ignored)
						}, func() any {
							return hx_receiver_414.__hx_this
						}, "secret")
					}(), Found: true}
				case "retainSecret":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_415 := hx_structural_source_405
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_415.__hx_this.retainSecret()
						}, func() any {
							return hx_receiver_415.__hx_this
						}, "retainSecret")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "total":
					hx_structural_source_405.total = func(hx_value_407 any) int {
						if hx_value_407 == nil {
							var hx_zero_408 int
							return hx_zero_408
						}
						return hx_value_407.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *SpecializedIterator:
		return func(hx_structural_source_416 *SpecializedIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_416, func(name string) hxrt.AnonymousField {
				switch name {
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_417 := hx_structural_source_416
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_417.__hx_this.next()
						}, func() any {
							return hx_receiver_417.__hx_this
						}, "next")
					}(), Found: true}
				case "values":
					return hxrt.AnonymousField{Value: hx_structural_source_416.values, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_416.index, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_420 := hx_structural_source_416
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_420.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_420.__hx_this
						}, "hasNext")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "index":
					hx_structural_source_416.index = func(hx_value_418 any) int {
						if hx_value_418 == nil {
							var hx_zero_419 int
							return hx_zero_419
						}
						return hx_value_418.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_421 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_421, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_421.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_421.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_421.high = func(hx_value_422 any) int {
						if hx_value_422 == nil {
							var hx_zero_423 int
							return hx_zero_423
						}
						return hx_value_422.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_421.low = func(hx_value_424 any) int {
						if hx_value_424 == nil {
							var hx_zero_425 int
							return hx_zero_425
						}
						return hx_value_424.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Template__TokenCursor:
		return func(hx_structural_source_426 *haxe___Template__TokenCursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_426, func(name string) hxrt.AnonymousField {
				switch name {
				case "tokens":
					return hxrt.AnonymousField{Value: hx_structural_source_426.tokens, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_426.index, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "tokens":
					hx_structural_source_426.tokens = func(hx_value_427 any) *hxrt.Array {
						if hx_value_427 == nil {
							var hx_zero_428 *hxrt.Array
							return hx_zero_428
						}
						return hx_value_427.(*hxrt.Array)
					}(incoming)
					return true
				case "index":
					hx_structural_source_426.index = func(hx_value_429 any) int {
						if hx_value_429 == nil {
							var hx_zero_430 int
							return hx_zero_430
						}
						return hx_value_429.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Template__ExprCursor:
		return func(hx_structural_source_431 *haxe___Template__ExprCursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_431, func(name string) hxrt.AnonymousField {
				switch name {
				case "tokens":
					return hxrt.AnonymousField{Value: hx_structural_source_431.tokens, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_431.index, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "tokens":
					hx_structural_source_431.tokens = func(hx_value_432 any) *hxrt.Array {
						if hx_value_432 == nil {
							var hx_zero_433 *hxrt.Array
							return hx_zero_433
						}
						return hx_value_432.(*hxrt.Array)
					}(incoming)
					return true
				case "index":
					hx_structural_source_431.index = func(hx_value_434 any) int {
						if hx_value_434 == nil {
							var hx_zero_435 int
							return hx_zero_435
						}
						return hx_value_434.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Template:
		return func(hx_structural_source_436 *haxe__Template) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_436, func(name string) hxrt.AnonymousField {
				switch name {
				case "expr":
					return hxrt.AnonymousField{Value: hx_structural_source_436.expr, Found: true}
				case "context":
					return hxrt.AnonymousField{Value: hx_structural_source_436.context, Found: true}
				case "macros":
					return hxrt.AnonymousField{Value: hx_structural_source_436.macros, Found: true}
				case "stack":
					return hxrt.AnonymousField{Value: hx_structural_source_436.stack, Found: true}
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_436.output, Found: true}
				case "execute":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any) *string] {
						hx_receiver_443 := hx_structural_source_436
						return hxrt.BoundFunction(func(context any, macros any) *string {
							return hx_receiver_443.__hx_this.execute(context, macros)
						}, func() any {
							return hx_receiver_443.__hx_this
						}, "execute")
					}(), Found: true}
				case "resolve":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_444 := hx_structural_source_436
						return hxrt.BoundFunction(func(v *string) any {
							return hx_receiver_444.__hx_this.resolve(v)
						}, func() any {
							return hx_receiver_444.__hx_this
						}, "resolve")
					}(), Found: true}
				case "parseTokens":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_445 := hx_structural_source_436
						return hxrt.BoundFunction(func(data *string) *hxrt.Array {
							return hx_receiver_445.__hx_this.parseTokens(data)
						}, func() any {
							return hx_receiver_445.__hx_this
						}, "parseTokens")
					}(), Found: true}
				case "parseBlock":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
						hx_receiver_446 := hx_structural_source_436
						return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
							return hx_receiver_446.__hx_this.parseBlock(cursor)
						}, func() any {
							return hx_receiver_446.__hx_this
						}, "parseBlock")
					}(), Found: true}
				case "parse":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
						hx_receiver_447 := hx_structural_source_436
						return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
							return hx_receiver_447.__hx_this.parse(cursor)
						}, func() any {
							return hx_receiver_447.__hx_this
						}, "parse")
					}(), Found: true}
				case "parseExpr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
						hx_receiver_448 := hx_structural_source_436
						return hxrt.BoundFunction(func(data *string) *hxrt.Function[func() any] {
							return hx_receiver_448.__hx_this.parseExpr(data)
						}, func() any {
							return hx_receiver_448.__hx_this
						}, "parseExpr")
					}(), Found: true}
				case "makeConst":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
						hx_receiver_449 := hx_structural_source_436
						return hxrt.BoundFunction(func(v *string) *hxrt.Function[func() any] {
							return hx_receiver_449.__hx_this.makeConst(v)
						}, func() any {
							return hx_receiver_449.__hx_this
						}, "makeConst")
					}(), Found: true}
				case "makePath":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_450 := hx_structural_source_436
						return hxrt.BoundFunction(func(e *hxrt.Function[func() any], cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_450.__hx_this.makePath(e, cursor)
						}, func() any {
							return hx_receiver_450.__hx_this
						}, "makePath")
					}(), Found: true}
				case "makeExpr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_451 := hx_structural_source_436
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_451.__hx_this.makeExpr(cursor)
						}, func() any {
							return hx_receiver_451.__hx_this
						}, "makeExpr")
					}(), Found: true}
				case "skipSpaces":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor)] {
						hx_receiver_452 := hx_structural_source_436
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) {
							hx_receiver_452.__hx_this.skipSpaces(cursor)
						}, func() any {
							return hx_receiver_452.__hx_this
						}, "skipSpaces")
					}(), Found: true}
				case "makeExpr2":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_453 := hx_structural_source_436
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_453.__hx_this.makeExpr2(cursor)
						}, func() any {
							return hx_receiver_453.__hx_this
						}, "makeExpr2")
					}(), Found: true}
				case "run":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TemplateExpr)] {
						hx_receiver_454 := hx_structural_source_436
						return hxrt.BoundFunction(func(e *haxe___Template__TemplateExpr) {
							hx_receiver_454.__hx_this.run(e)
						}, func() any {
							return hx_receiver_454.__hx_this
						}, "run")
					}(), Found: true}
				case "popStackValue":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_455 := hx_structural_source_436
						return hxrt.BoundFunction(func() any {
							return hx_receiver_455.__hx_this.popStackValue()
						}, func() any {
							return hx_receiver_455.__hx_this
						}, "popStackValue")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "expr":
					hx_structural_source_436.expr = func(hx_value_437 any) *haxe___Template__TemplateExpr {
						if hx_value_437 == nil {
							var hx_zero_438 *haxe___Template__TemplateExpr
							return hx_zero_438
						}
						return hx_value_437.(*haxe___Template__TemplateExpr)
					}(incoming)
					return true
				case "context":
					hx_structural_source_436.context = incoming
					return true
				case "macros":
					hx_structural_source_436.macros = incoming
					return true
				case "stack":
					hx_structural_source_436.stack = func(hx_value_439 any) *hxrt.Array {
						if hx_value_439 == nil {
							var hx_zero_440 *hxrt.Array
							return hx_zero_440
						}
						return hx_value_439.(*hxrt.Array)
					}(incoming)
					return true
				case "output":
					hx_structural_source_436.output = func(hx_value_441 any) *string {
						if hx_value_441 == nil {
							var hx_zero_442 *string
							return hx_zero_442
						}
						return hx_value_441.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_456 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_456, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_456.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_456.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_461 := hx_structural_source_456
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_461.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_461.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_462 := hx_structural_source_456
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_462.__hx_this.next()
						}, func() any {
							return hx_receiver_462.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_456.offset = func(hx_value_457 any) int {
						if hx_value_457 == nil {
							var hx_zero_458 int
							return hx_zero_458
						}
						return hx_value_457.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_456.s = func(hx_value_459 any) *string {
						if hx_value_459 == nil {
							var hx_zero_460 *string
							return hx_zero_460
						}
						return hx_value_459.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_463 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_463, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_463.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_463.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_468 := hx_structural_source_463
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_468.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_468.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_469 := hx_structural_source_463
						return hxrt.BoundFunction(func() int {
							return hx_receiver_469.__hx_this.next()
						}, func() any {
							return hx_receiver_469.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_463.offset = func(hx_value_464 any) int {
						if hx_value_464 == nil {
							var hx_zero_465 int
							return hx_zero_465
						}
						return hx_value_464.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_463.s = func(hx_value_466 any) *string {
						if hx_value_466 == nil {
							var hx_zero_467 *string
							return hx_zero_467
						}
						return hx_value_466.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_470 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_470, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_470.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_470.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_470.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_470.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_473 := hx_structural_source_470
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_473.__hx_this.get_message()
						}, func() any {
							return hx_receiver_473.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_474 := hx_structural_source_470
						return hxrt.BoundFunction(func() any {
							return hx_receiver_474.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_474.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_475 := hx_structural_source_470
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_475.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_475.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_476 := hx_structural_source_470
						return hxrt.BoundFunction(func() any {
							return hx_receiver_476.__hx_this.get_native()
						}, func() any {
							return hx_receiver_476.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_477 := hx_structural_source_470
						return hxrt.BoundFunction(func() any {
							return hx_receiver_477.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_477.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_478 := hx_structural_source_470
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_478.__hx_this.toString()
						}, func() any {
							return hx_receiver_478.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_479 := hx_structural_source_470
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_479.__hx_this.details()
						}, func() any {
							return hx_receiver_479.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_470.savedMessage = func(hx_value_471 any) *string {
						if hx_value_471 == nil {
							var hx_zero_472 *string
							return hx_zero_472
						}
						return hx_value_471.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_470.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_470.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_470.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *EReg:
		return func(hx_structural_source_480 *EReg) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_480, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_480.handle, Found: true}
				case "global":
					return hxrt.AnonymousField{Value: hx_structural_source_480.global, Found: true}
				case "lastSource":
					return hxrt.AnonymousField{Value: hx_structural_source_480.lastSource, Found: true}
				case "lastMatch":
					return hxrt.AnonymousField{Value: hx_structural_source_480.lastMatch, Found: true}
				case "match":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_489 := hx_structural_source_480
						return hxrt.BoundFunction(func(s *string) bool {
							return hx_receiver_489.__hx_this.match(s)
						}, func() any {
							return hx_receiver_489.__hx_this
						}, "match")
					}(), Found: true}
				case "matched":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_490 := hx_structural_source_480
						return hxrt.BoundFunction(func(n int) *string {
							return hx_receiver_490.__hx_this.matched(n)
						}, func() any {
							return hx_receiver_490.__hx_this
						}, "matched")
					}(), Found: true}
				case "matchedLeft":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_491 := hx_structural_source_480
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_491.__hx_this.matchedLeft()
						}, func() any {
							return hx_receiver_491.__hx_this
						}, "matchedLeft")
					}(), Found: true}
				case "matchedRight":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_492 := hx_structural_source_480
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_492.__hx_this.matchedRight()
						}, func() any {
							return hx_receiver_492.__hx_this
						}, "matchedRight")
					}(), Found: true}
				case "matchedPos":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_493 := hx_structural_source_480
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_493.__hx_this.matchedPos()
						}, func() any {
							return hx_receiver_493.__hx_this
						}, "matchedPos")
					}(), Found: true}
				case "matchSub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, int, any) bool] {
						hx_receiver_494 := hx_structural_source_480
						return hxrt.BoundFunction(func(s *string, pos int, len any) bool {
							if len == nil {
								len = -1
							}
							return hx_receiver_494.__hx_this.matchSub(s, pos, func(hx_value_495 any) int {
								if hx_value_495 == nil {
									var hx_zero_496 int
									return hx_zero_496
								}
								return hx_value_495.(int)
							}(any(len)))
						}, func() any {
							return hx_receiver_494.__hx_this
						}, "matchSub")
					}(), Found: true}
				case "split":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_497 := hx_structural_source_480
						return hxrt.BoundFunction(func(s *string) *hxrt.Array {
							return hx_receiver_497.__hx_this.split(s)
						}, func() any {
							return hx_receiver_497.__hx_this
						}, "split")
					}(), Found: true}
				case "replace":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string) *string] {
						hx_receiver_498 := hx_structural_source_480
						return hxrt.BoundFunction(func(s *string, by *string) *string {
							return hx_receiver_498.__hx_this.replace(s, by)
						}, func() any {
							return hx_receiver_498.__hx_this
						}, "replace")
					}(), Found: true}
				case "map":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string] {
						hx_receiver_499 := hx_structural_source_480
						return hxrt.BoundFunction(func(s *string, f *hxrt.Function[func(*EReg) *string]) *string {
							return hx_receiver_499.__hx_this.map_(s, f)
						}, func() any {
							return hx_receiver_499.__hx_this
						}, "map")
					}(), Found: true}
				case "remember":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *hxrt.RegexMatch)] {
						hx_receiver_500 := hx_structural_source_480
						return hxrt.BoundFunction(func(source *string, match *hxrt.RegexMatch) {
							hx_receiver_500.__hx_this.remember(source, match)
						}, func() any {
							return hx_receiver_500.__hx_this
						}, "remember")
					}(), Found: true}
				case "requireMatch":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.RegexMatch] {
						hx_receiver_501 := hx_structural_source_480
						return hxrt.BoundFunction(func() *hxrt.RegexMatch {
							return hx_receiver_501.__hx_this.requireMatch()
						}, func() any {
							return hx_receiver_501.__hx_this
						}, "requireMatch")
					}(), Found: true}
				case "expandReplacement":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *hxrt.RegexMatch) *string] {
						hx_receiver_502 := hx_structural_source_480
						return hxrt.BoundFunction(func(by *string, source *string, currentMatch *hxrt.RegexMatch) *string {
							return hx_receiver_502.__hx_this.expandReplacement(by, source, currentMatch)
						}, func() any {
							return hx_receiver_502.__hx_this
						}, "expandReplacement")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_480.handle = func(hx_value_481 any) *hxrt.RegexHandle {
						if hx_value_481 == nil {
							var hx_zero_482 *hxrt.RegexHandle
							return hx_zero_482
						}
						return hx_value_481.(*hxrt.RegexHandle)
					}(incoming)
					return true
				case "global":
					hx_structural_source_480.global = func(hx_value_483 any) bool {
						if hx_value_483 == nil {
							var hx_zero_484 bool
							return hx_zero_484
						}
						return hx_value_483.(bool)
					}(incoming)
					return true
				case "lastSource":
					hx_structural_source_480.lastSource = func(hx_value_485 any) *string {
						if hx_value_485 == nil {
							var hx_zero_486 *string
							return hx_zero_486
						}
						return hx_value_485.(*string)
					}(incoming)
					return true
				case "lastMatch":
					hx_structural_source_480.lastMatch = func(hx_value_487 any) *hxrt.RegexMatch {
						if hx_value_487 == nil {
							var hx_zero_488 *hxrt.RegexMatch
							return hx_zero_488
						}
						return hx_value_487.(*hxrt.RegexMatch)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_503 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_503, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_503.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_504 := hx_structural_source_503
						return hxrt.BoundFunction(func() any {
							return hx_receiver_504.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_504.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_503.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_503.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_503.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_503.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_507 := hx_structural_source_503
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_507.__hx_this.get_message()
						}, func() any {
							return hx_receiver_507.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_508 := hx_structural_source_503
						return hxrt.BoundFunction(func() any {
							return hx_receiver_508.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_508.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_509 := hx_structural_source_503
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_509.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_509.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_510 := hx_structural_source_503
						return hxrt.BoundFunction(func() any {
							return hx_receiver_510.__hx_this.get_native()
						}, func() any {
							return hx_receiver_510.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_511 := hx_structural_source_503
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_511.__hx_this.toString()
						}, func() any {
							return hx_receiver_511.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_512 := hx_structural_source_503
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_512.__hx_this.details()
						}, func() any {
							return hx_receiver_512.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_503.savedMessage = func(hx_value_505 any) *string {
						if hx_value_505 == nil {
							var hx_zero_506 *string
							return hx_zero_506
						}
						return hx_value_505.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_503.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_503.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_503.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_513 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_513, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_513.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_513.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_513.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_513.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_513.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_522 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_522.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_522.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_523 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_523.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_523.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_524 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
							hx_receiver_524.__hx_this.blit(pos, src, srcpos, len)
						}, func() any {
							return hx_receiver_524.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_525 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, len int, value int) {
							hx_receiver_525.__hx_this.fill(pos, len, value)
						}, func() any {
							return hx_receiver_525.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_526 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
							return hx_receiver_526.__hx_this.sub(pos, len)
						}, func() any {
							return hx_receiver_526.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_527 := hx_structural_source_513
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_527.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_527.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_528 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_528.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_528.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_529 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_529.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_529.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_530 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_530.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_530.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_531 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_531.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_531.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_532 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_532.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_532.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_533 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_533.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_533.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_534 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_534.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_534.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_535 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_535.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_535.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_536 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_536.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_536.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_537 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_537.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_537.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_538 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_538.__hx_this.getString(pos, len, encoding)
						}, func() any {
							return hx_receiver_538.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_539 := hx_structural_source_513
						return hxrt.BoundFunction(func(pos int, len int) *string {
							return hx_receiver_539.__hx_this.readString(pos, len)
						}, func() any {
							return hx_receiver_539.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_540 := hx_structural_source_513
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_540.__hx_this.toString()
						}, func() any {
							return hx_receiver_540.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_541 := hx_structural_source_513
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_541.__hx_this.toHex()
						}, func() any {
							return hx_receiver_541.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_542 := hx_structural_source_513
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_542.__hx_this.getData()
						}, func() any {
							return hx_receiver_542.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_543 := hx_structural_source_513
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_543.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_543.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_513.b = func(hx_value_514 any) []int {
						if hx_value_514 == nil {
							var hx_zero_515 []int
							return hx_zero_515
						}
						return hx_value_514.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_513.__hx_raw = func(hx_value_516 any) *hxrt.ByteView {
						if hx_value_516 == nil {
							var hx_zero_517 *hxrt.ByteView
							return hx_zero_517
						}
						return hx_value_516.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_513.__hx_rawValid = func(hx_value_518 any) bool {
						if hx_value_518 == nil {
							var hx_zero_519 bool
							return hx_zero_519
						}
						return hx_value_518.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_513.__hx_dataExposed = func(hx_value_520 any) bool {
						if hx_value_520 == nil {
							var hx_zero_521 bool
							return hx_zero_521
						}
						return hx_value_520.(bool)
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
			hx_receiver_544 := value
			return hxrt.BoundFunction(func() *SpecializedIterator {
				return hx_receiver_544.__hx_this.iterator()
			}, func() any {
				return hx_receiver_544.__hx_this
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
			hx_receiver_545 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_545.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_545.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_546 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_546.__hx_this.next()
			}, func() any {
				return hx_receiver_546.__hx_this
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
			hx_receiver_547 := value
			return hxrt.BoundFunction(func(by *string, source *string, currentMatch *hxrt.RegexMatch) *string {
				return hx_receiver_547.__hx_this.expandReplacement(by, source, currentMatch)
			}, func() any {
				return hx_receiver_547.__hx_this
			}, "expandReplacement")
		}()
	case "map":
		return func() *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string] {
			hx_receiver_548 := value
			return hxrt.BoundFunction(func(s *string, f *hxrt.Function[func(*EReg) *string]) *string {
				return hx_receiver_548.__hx_this.map_(s, f)
			}, func() any {
				return hx_receiver_548.__hx_this
			}, "map")
		}()
	case "match":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_549 := value
			return hxrt.BoundFunction(func(s *string) bool {
				return hx_receiver_549.__hx_this.match(s)
			}, func() any {
				return hx_receiver_549.__hx_this
			}, "match")
		}()
	case "matchSub":
		return func() *hxrt.Function[func(*string, int, any) bool] {
			hx_receiver_550 := value
			return hxrt.BoundFunction(func(s *string, pos int, len any) bool {
				if len == nil {
					len = -1
				}
				return hx_receiver_550.__hx_this.matchSub(s, pos, func(hx_value_551 any) int {
					if hx_value_551 == nil {
						var hx_zero_552 int
						return hx_zero_552
					}
					return hx_value_551.(int)
				}(any(len)))
			}, func() any {
				return hx_receiver_550.__hx_this
			}, "matchSub")
		}()
	case "matched":
		return func() *hxrt.Function[func(int) *string] {
			hx_receiver_553 := value
			return hxrt.BoundFunction(func(n int) *string {
				return hx_receiver_553.__hx_this.matched(n)
			}, func() any {
				return hx_receiver_553.__hx_this
			}, "matched")
		}()
	case "matchedLeft":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_554 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_554.__hx_this.matchedLeft()
			}, func() any {
				return hx_receiver_554.__hx_this
			}, "matchedLeft")
		}()
	case "matchedPos":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_555 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_555.__hx_this.matchedPos()
			}, func() any {
				return hx_receiver_555.__hx_this
			}, "matchedPos")
		}()
	case "matchedRight":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_556 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_556.__hx_this.matchedRight()
			}, func() any {
				return hx_receiver_556.__hx_this
			}, "matchedRight")
		}()
	case "remember":
		return func() *hxrt.Function[func(*string, *hxrt.RegexMatch)] {
			hx_receiver_557 := value
			return hxrt.BoundFunction(func(source *string, match *hxrt.RegexMatch) {
				hx_receiver_557.__hx_this.remember(source, match)
			}, func() any {
				return hx_receiver_557.__hx_this
			}, "remember")
		}()
	case "replace":
		return func() *hxrt.Function[func(*string, *string) *string] {
			hx_receiver_558 := value
			return hxrt.BoundFunction(func(s *string, by *string) *string {
				return hx_receiver_558.__hx_this.replace(s, by)
			}, func() any {
				return hx_receiver_558.__hx_this
			}, "replace")
		}()
	case "requireMatch":
		return func() *hxrt.Function[func() *hxrt.RegexMatch] {
			hx_receiver_559 := value
			return hxrt.BoundFunction(func() *hxrt.RegexMatch {
				return hx_receiver_559.__hx_this.requireMatch()
			}, func() any {
				return hx_receiver_559.__hx_this
			}, "requireMatch")
		}()
	case "split":
		return func() *hxrt.Function[func(*string) *hxrt.Array] {
			hx_receiver_560 := value
			return hxrt.BoundFunction(func(s *string) *hxrt.Array {
				return hx_receiver_560.__hx_this.split(s)
			}, func() any {
				return hx_receiver_560.__hx_this
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
			hx_receiver_561 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) int {
				return hx_receiver_561.__hx_this.bump(resolve, key)
			}, func() any {
				return hx_receiver_561.__hx_this
			}, "bump")
		}()
	case "describe":
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) *string] {
			hx_receiver_562 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], key *string) *string {
				return hx_receiver_562.__hx_this.describe(resolve, key)
			}, func() any {
				return hx_receiver_562.__hx_this
			}, "describe")
		}()
	case "macroName":
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
			hx_receiver_563 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
				return hx_receiver_563.__hx_this.macroName(resolve, ignored)
			}, func() any {
				return hx_receiver_563.__hx_this
			}, "macroName")
		}()
	case "retainSecret":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_564 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_564.__hx_this.retainSecret()
			}, func() any {
				return hx_receiver_564.__hx_this
			}, "retainSecret")
		}()
	case "secret":
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
			hx_receiver_565 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
				return hx_receiver_565.__hx_this.secret(resolve, ignored)
			}, func() any {
				return hx_receiver_565.__hx_this
			}, "secret")
		}()
	case "type":
		return func() *hxrt.Function[func(*hxrt.Function[func(*string) any], any) *string] {
			hx_receiver_566 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
				return hx_receiver_566.__hx_this.type_(resolve, ignored)
			}, func() any {
				return hx_receiver_566.__hx_this
			}, "type")
		}()
	case "virtualName":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_567 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_567.__hx_this.virtualName()
			}, func() any {
				return hx_receiver_567.__hx_this
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
			hx_receiver_568 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
				return hx_receiver_568.__hx_this.leafOnly(resolve, ignored)
			}, func() any {
				return hx_receiver_568.__hx_this
			}, "leafOnly")
		}()
	case "virtualName":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_569 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_569.__hx_this.virtualName()
			}, func() any {
				return hx_receiver_569.__hx_this
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
			hx_receiver_570 := value
			return hxrt.BoundFunction(func(resolve *hxrt.Function[func(*string) any], ignored any) *string {
				return hx_receiver_570.__hx_this.middleOnly(resolve, ignored)
			}, func() any {
				return hx_receiver_570.__hx_this
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
			hx_receiver_571 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_571.__hx_this.next()
			}, func() any {
				return hx_receiver_571.__hx_this
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
			hx_receiver_572 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_572.__hx_this.details()
			}, func() any {
				return hx_receiver_572.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_573 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_573.__hx_this.get_message()
			}, func() any {
				return hx_receiver_573.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_574 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_574.__hx_this.get_native()
			}, func() any {
				return hx_receiver_574.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_575 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_575.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_575.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_576 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_576.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_576.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_577 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_577.__hx_this.toString()
			}, func() any {
				return hx_receiver_577.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_578 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_578.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_578.__hx_this
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
			hx_receiver_579 := value
			return hxrt.BoundFunction(func(context any, macros any) *string {
				return hx_receiver_579.__hx_this.execute(context, macros)
			}, func() any {
				return hx_receiver_579.__hx_this
			}, "execute")
		}()
	case "makeConst":
		return func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
			hx_receiver_580 := value
			return hxrt.BoundFunction(func(v *string) *hxrt.Function[func() any] {
				return hx_receiver_580.__hx_this.makeConst(v)
			}, func() any {
				return hx_receiver_580.__hx_this
			}, "makeConst")
		}()
	case "makeExpr":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_581 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_581.__hx_this.makeExpr(cursor)
			}, func() any {
				return hx_receiver_581.__hx_this
			}, "makeExpr")
		}()
	case "makeExpr2":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_582 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_582.__hx_this.makeExpr2(cursor)
			}, func() any {
				return hx_receiver_582.__hx_this
			}, "makeExpr2")
		}()
	case "makePath":
		return func() *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_583 := value
			return hxrt.BoundFunction(func(e *hxrt.Function[func() any], cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_583.__hx_this.makePath(e, cursor)
			}, func() any {
				return hx_receiver_583.__hx_this
			}, "makePath")
		}()
	case "parse":
		return func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
			hx_receiver_584 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
				return hx_receiver_584.__hx_this.parse(cursor)
			}, func() any {
				return hx_receiver_584.__hx_this
			}, "parse")
		}()
	case "parseBlock":
		return func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
			hx_receiver_585 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
				return hx_receiver_585.__hx_this.parseBlock(cursor)
			}, func() any {
				return hx_receiver_585.__hx_this
			}, "parseBlock")
		}()
	case "parseExpr":
		return func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
			hx_receiver_586 := value
			return hxrt.BoundFunction(func(data *string) *hxrt.Function[func() any] {
				return hx_receiver_586.__hx_this.parseExpr(data)
			}, func() any {
				return hx_receiver_586.__hx_this
			}, "parseExpr")
		}()
	case "parseTokens":
		return func() *hxrt.Function[func(*string) *hxrt.Array] {
			hx_receiver_587 := value
			return hxrt.BoundFunction(func(data *string) *hxrt.Array {
				return hx_receiver_587.__hx_this.parseTokens(data)
			}, func() any {
				return hx_receiver_587.__hx_this
			}, "parseTokens")
		}()
	case "popStackValue":
		return func() *hxrt.Function[func() any] {
			hx_receiver_588 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_588.__hx_this.popStackValue()
			}, func() any {
				return hx_receiver_588.__hx_this
			}, "popStackValue")
		}()
	case "resolve":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_589 := value
			return hxrt.BoundFunction(func(v *string) any {
				return hx_receiver_589.__hx_this.resolve(v)
			}, func() any {
				return hx_receiver_589.__hx_this
			}, "resolve")
		}()
	case "run":
		return func() *hxrt.Function[func(*haxe___Template__TemplateExpr)] {
			hx_receiver_590 := value
			return hxrt.BoundFunction(func(e *haxe___Template__TemplateExpr) {
				hx_receiver_590.__hx_this.run(e)
			}, func() any {
				return hx_receiver_590.__hx_this
			}, "run")
		}()
	case "skipSpaces":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor)] {
			hx_receiver_591 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) {
				hx_receiver_591.__hx_this.skipSpaces(cursor)
			}, func() any {
				return hx_receiver_591.__hx_this
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
			hx_receiver_592 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_592.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_592.__hx_this
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
			hx_receiver_593 := value
			return hxrt.BoundFunction(func() *hxrt.ByteView {
				return hx_receiver_593.__hx_this.__hx_nativeView()
			}, func() any {
				return hx_receiver_593.__hx_this
			}, "__hx_nativeView")
		}()
	case "blit":
		return func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
			hx_receiver_594 := value
			return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
				hx_receiver_594.__hx_this.blit(pos, src, srcpos, len)
			}, func() any {
				return hx_receiver_594.__hx_this
			}, "blit")
		}()
	case "compare":
		return func() *hxrt.Function[func(*haxe__io__Bytes) int] {
			hx_receiver_595 := value
			return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
				return hx_receiver_595.__hx_this.compare(other)
			}, func() any {
				return hx_receiver_595.__hx_this
			}, "compare")
		}()
	case "fill":
		return func() *hxrt.Function[func(int, int, int)] {
			hx_receiver_596 := value
			return hxrt.BoundFunction(func(pos int, len int, value int) {
				hx_receiver_596.__hx_this.fill(pos, len, value)
			}, func() any {
				return hx_receiver_596.__hx_this
			}, "fill")
		}()
	case "get":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_597 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_597.__hx_this.get(pos)
			}, func() any {
				return hx_receiver_597.__hx_this
			}, "get")
		}()
	case "getData":
		return func() *hxrt.Function[func() []int] {
			hx_receiver_598 := value
			return hxrt.BoundFunction(func() []int {
				return hx_receiver_598.__hx_this.getData()
			}, func() any {
				return hx_receiver_598.__hx_this
			}, "getData")
		}()
	case "getDouble":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_599 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_599.__hx_this.getDouble(pos)
			}, func() any {
				return hx_receiver_599.__hx_this
			}, "getDouble")
		}()
	case "getFloat":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_600 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_600.__hx_this.getFloat(pos)
			}, func() any {
				return hx_receiver_600.__hx_this
			}, "getFloat")
		}()
	case "getInt32":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_601 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_601.__hx_this.getInt32(pos)
			}, func() any {
				return hx_receiver_601.__hx_this
			}, "getInt32")
		}()
	case "getInt64":
		return func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
			hx_receiver_602 := value
			return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
				return hx_receiver_602.__hx_this.getInt64(pos)
			}, func() any {
				return hx_receiver_602.__hx_this
			}, "getInt64")
		}()
	case "getString":
		return func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
			hx_receiver_603 := value
			return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_603.__hx_this.getString(pos, len, encoding)
			}, func() any {
				return hx_receiver_603.__hx_this
			}, "getString")
		}()
	case "getUInt16":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_604 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_604.__hx_this.getUInt16(pos)
			}, func() any {
				return hx_receiver_604.__hx_this
			}, "getUInt16")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, int) *string] {
			hx_receiver_605 := value
			return hxrt.BoundFunction(func(pos int, len int) *string {
				return hx_receiver_605.__hx_this.readString(pos, len)
			}, func() any {
				return hx_receiver_605.__hx_this
			}, "readString")
		}()
	case "set":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_606 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_606.__hx_this.set(pos, value)
			}, func() any {
				return hx_receiver_606.__hx_this
			}, "set")
		}()
	case "setDouble":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_607 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_607.__hx_this.setDouble(pos, value)
			}, func() any {
				return hx_receiver_607.__hx_this
			}, "setDouble")
		}()
	case "setFloat":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_608 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_608.__hx_this.setFloat(pos, value)
			}, func() any {
				return hx_receiver_608.__hx_this
			}, "setFloat")
		}()
	case "setInt32":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_609 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_609.__hx_this.setInt32(pos, value)
			}, func() any {
				return hx_receiver_609.__hx_this
			}, "setInt32")
		}()
	case "setInt64":
		return func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
			hx_receiver_610 := value
			return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
				hx_receiver_610.__hx_this.setInt64(pos, value)
			}, func() any {
				return hx_receiver_610.__hx_this
			}, "setInt64")
		}()
	case "setUInt16":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_611 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_611.__hx_this.setUInt16(pos, value)
			}, func() any {
				return hx_receiver_611.__hx_this
			}, "setUInt16")
		}()
	case "sub":
		return func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
			hx_receiver_612 := value
			return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
				return hx_receiver_612.__hx_this.sub(pos, len)
			}, func() any {
				return hx_receiver_612.__hx_this
			}, "sub")
		}()
	case "toHex":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_613 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_613.__hx_this.toHex()
			}, func() any {
				return hx_receiver_613.__hx_this
			}, "toHex")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_614 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_614.__hx_this.toString()
			}, func() any {
				return hx_receiver_614.__hx_this
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
			hx_receiver_615 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_615.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_615.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_616 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_616.__hx_this.next()
			}, func() any {
				return hx_receiver_616.__hx_this
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
			hx_receiver_617 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_617.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_617.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_618 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_618.__hx_this.next()
			}, func() any {
				return hx_receiver_618.__hx_this
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
		return hxrt.FunctionCall(source)(func(hx_value_619 any) *string {
			if hx_value_619 == nil {
				var hx_zero_620 *string
				return hx_zero_620
			}
			return hx_value_619.(*string)
		}(any(argument_0)))
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
		return hxrt.FunctionCall(source)(hx_recover_callable_1(argument_0), argument_1)
	case *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_1(argument_0), func(hx_value_621 any) *string {
			if hx_value_621 == nil {
				var hx_zero_622 *string
				return hx_zero_622
			}
			return hx_value_621.(*string)
		}(any(argument_1)))
	case *hxrt.Function[func(*hxrt.Function[func(*string) any], *string) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_1(argument_0), func(hx_value_623 any) *string {
			if hx_value_623 == nil {
				var hx_zero_624 *string
				return hx_zero_624
			}
			return hx_value_623.(*string)
		}(any(argument_1)))
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
		return hxrt.FunctionCall(source)(func(hx_value_625 any) *string {
			if hx_value_625 == nil {
				var hx_zero_626 *string
				return hx_zero_626
			}
			return hx_value_625.(*string)
		}(any(argument_0)))
	case *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_2(argument_0))
	case *hxrt.Function[func(*string) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_627 any) *string {
			if hx_value_627 == nil {
				var hx_zero_628 *string
				return hx_zero_628
			}
			return hx_value_627.(*string)
		}(any(argument_0)))
	case *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_0(argument_0), hx_recover_class_3(argument_1))
	case *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_3(argument_0))
	case *hxrt.Function[func(*haxe___Template__ExprCursor)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_3(argument_0))
		return nil
	case *hxrt.Function[func(*haxe___Template__TemplateExpr)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_629 any) *haxe___Template__TemplateExpr {
			if hx_value_629 == nil {
				var hx_zero_630 *haxe___Template__TemplateExpr
				return hx_zero_630
			}
			return hx_value_629.(*haxe___Template__TemplateExpr)
		}(any(argument_0)))
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
		return hxrt.FunctionCall(source)(func(hx_value_631 any) *string {
			if hx_value_631 == nil {
				var hx_zero_632 *string
				return hx_zero_632
			}
			return hx_value_631.(*string)
		}(any(argument_0)))
	case *hxrt.Function[func(int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_633 any) int {
			if hx_value_633 == nil {
				var hx_zero_634 int
				return hx_zero_634
			}
			return hx_value_633.(int)
		}(any(argument_0)))
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
		return hxrt.FunctionCall(source)(func(hx_value_635 any) *string {
			if hx_value_635 == nil {
				var hx_zero_636 *string
				return hx_zero_636
			}
			return hx_value_635.(*string)
		}(any(argument_0)), func(hx_value_637 any) int {
			if hx_value_637 == nil {
				var hx_zero_638 int
				return hx_zero_638
			}
			return hx_value_637.(int)
		}(any(argument_1)), argument_2)
	case *hxrt.Function[func(*string, *string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_639 any) *string {
			if hx_value_639 == nil {
				var hx_zero_640 *string
				return hx_zero_640
			}
			return hx_value_639.(*string)
		}(any(argument_0)), func(hx_value_641 any) *string {
			if hx_value_641 == nil {
				var hx_zero_642 *string
				return hx_zero_642
			}
			return hx_value_641.(*string)
		}(any(argument_1)))
	case *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_643 any) *string {
			if hx_value_643 == nil {
				var hx_zero_644 *string
				return hx_zero_644
			}
			return hx_value_643.(*string)
		}(any(argument_0)), hx_recover_callable_2(argument_1))
	case *hxrt.Function[func(*string, *hxrt.RegexMatch)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_645 any) *string {
			if hx_value_645 == nil {
				var hx_zero_646 *string
				return hx_zero_646
			}
			return hx_value_645.(*string)
		}(any(argument_0)), func(hx_value_647 any) *hxrt.RegexMatch {
			if hx_value_647 == nil {
				var hx_zero_648 *hxrt.RegexMatch
				return hx_zero_648
			}
			return hx_value_647.(*hxrt.RegexMatch)
		}(any(argument_1)))
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
		return hxrt.FunctionCall(source)(func(hx_value_649 any) *string {
			if hx_value_649 == nil {
				var hx_zero_650 *string
				return hx_zero_650
			}
			return hx_value_649.(*string)
		}(any(argument_0)), func(hx_value_651 any) *string {
			if hx_value_651 == nil {
				var hx_zero_652 *string
				return hx_zero_652
			}
			return hx_value_651.(*string)
		}(any(argument_1)), func(hx_value_653 any) *hxrt.RegexMatch {
			if hx_value_653 == nil {
				var hx_zero_654 *hxrt.RegexMatch
				return hx_zero_654
			}
			return hx_value_653.(*hxrt.RegexMatch)
		}(any(argument_2)))
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_655 any) int {
			if hx_value_655 == nil {
				var hx_zero_656 int
				return hx_zero_656
			}
			return hx_value_655.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_657 any) int {
			if hx_value_657 == nil {
				var hx_zero_658 int
				return hx_zero_658
			}
			return hx_value_657.(int)
		}(any(argument_0)), func(hx_value_659 any) int {
			if hx_value_659 == nil {
				var hx_zero_660 int
				return hx_zero_660
			}
			return hx_value_659.(int)
		}(any(argument_1)))
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
		hxrt.FunctionCall(source)(func(hx_value_661 any) int {
			if hx_value_661 == nil {
				var hx_zero_662 int
				return hx_zero_662
			}
			return hx_value_661.(int)
		}(any(argument_0)), hx_recover_class_4(argument_1), func(hx_value_663 any) int {
			if hx_value_663 == nil {
				var hx_zero_664 int
				return hx_zero_664
			}
			return hx_value_663.(int)
		}(any(argument_2)), func(hx_value_665 any) int {
			if hx_value_665 == nil {
				var hx_zero_666 int
				return hx_zero_666
			}
			return hx_value_665.(int)
		}(any(argument_3)))
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
		hxrt.FunctionCall(source)(func(hx_value_667 any) int {
			if hx_value_667 == nil {
				var hx_zero_668 int
				return hx_zero_668
			}
			return hx_value_667.(int)
		}(any(argument_0)), func(hx_value_669 any) int {
			if hx_value_669 == nil {
				var hx_zero_670 int
				return hx_zero_670
			}
			return hx_value_669.(int)
		}(any(argument_1)), func(hx_value_671 any) int {
			if hx_value_671 == nil {
				var hx_zero_672 int
				return hx_zero_672
			}
			return hx_value_671.(int)
		}(any(argument_2)))
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
		return hxrt.FunctionCall(source)(func(hx_value_673 any) int {
			if hx_value_673 == nil {
				var hx_zero_674 int
				return hx_zero_674
			}
			return hx_value_673.(int)
		}(any(argument_0)), func(hx_value_675 any) int {
			if hx_value_675 == nil {
				var hx_zero_676 int
				return hx_zero_676
			}
			return hx_value_675.(int)
		}(any(argument_1)))
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_4(argument_0))
	case *hxrt.Function[func(int) float64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_677 any) int {
			if hx_value_677 == nil {
				var hx_zero_678 int
				return hx_zero_678
			}
			return hx_value_677.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(int, float64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_679 any) int {
			if hx_value_679 == nil {
				var hx_zero_680 int
				return hx_zero_680
			}
			return hx_value_679.(int)
		}(any(argument_0)), func(hx_value_681 any) float64 {
			if hx_value_681 == nil {
				var hx_zero_682 float64
				return hx_zero_682
			}
			return hx_value_681.(float64)
		}(any(argument_1)))
		return nil
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_683 any) int {
			if hx_value_683 == nil {
				var hx_zero_684 int
				return hx_zero_684
			}
			return hx_value_683.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(int, *haxe___Int64_____Int64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_685 any) int {
			if hx_value_685 == nil {
				var hx_zero_686 int
				return hx_zero_686
			}
			return hx_value_685.(int)
		}(any(argument_0)), hx_recover_class_5(argument_1))
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
		return hxrt.FunctionCall(source)(func(hx_value_687 any) int {
			if hx_value_687 == nil {
				var hx_zero_688 int
				return hx_zero_688
			}
			return hx_value_687.(int)
		}(any(argument_0)), func(hx_value_689 any) int {
			if hx_value_689 == nil {
				var hx_zero_690 int
				return hx_zero_690
			}
			return hx_value_689.(int)
		}(any(argument_1)), func(hx_value_691 any) *haxe__io__Encoding {
			if hx_value_691 == nil {
				var hx_zero_692 *haxe__io__Encoding
				return hx_zero_692
			}
			return hx_value_691.(*haxe__io__Encoding)
		}(any(argument_2)))
	case *hxrt.Function[func(int, int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_693 any) int {
			if hx_value_693 == nil {
				var hx_zero_694 int
				return hx_zero_694
			}
			return hx_value_693.(int)
		}(any(argument_0)), func(hx_value_695 any) int {
			if hx_value_695 == nil {
				var hx_zero_696 int
				return hx_zero_696
			}
			return hx_value_695.(int)
		}(any(argument_1)))
	case *hxrt.Function[func() []int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() any] {
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

func hx_recover_callable_1(value any) *hxrt.Function[func(*string) any] {
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
			return hxrt.FunctionCall(source)(hx_recover_class_2(arg_0))
		})
	case *hxrt.Function[func(*string) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(hx_recover_class_3(arg_0))
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			hxrt.FunctionCall(source)(hx_recover_class_3(arg_0))
			return nil
		})
	case *hxrt.Function[func(*haxe___Template__TemplateExpr)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			hxrt.FunctionCall(source)(func(hx_value_697 any) *haxe___Template__TemplateExpr {
				if hx_value_697 == nil {
					var hx_zero_698 *haxe___Template__TemplateExpr
					return hx_zero_698
				}
				return hx_value_697.(*haxe___Template__TemplateExpr)
			}(any(arg_0)))
			return nil
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(func(hx_value_699 any) int {
				if hx_value_699 == nil {
					var hx_zero_700 int
					return hx_zero_700
				}
				return hx_value_699.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(func(hx_value_701 any) int {
				if hx_value_701 == nil {
					var hx_zero_702 int
					return hx_zero_702
				}
				return hx_value_701.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(hx_recover_class_4(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(func(hx_value_703 any) int {
				if hx_value_703 == nil {
					var hx_zero_704 int
					return hx_zero_704
				}
				return hx_value_703.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) any {
			return hxrt.FunctionCall(source)(func(hx_value_705 any) int {
				if hx_value_705 == nil {
					var hx_zero_706 int
					return hx_zero_706
				}
				return hx_value_705.(int)
			}(any(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func(*EReg) *string] {
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
			return func(hx_value_709 any) *string {
				if hx_value_709 == nil {
					var hx_zero_710 *string
					return hx_zero_710
				}
				return hx_value_709.(*string)
			}(any(hxrt.FunctionCall(source)(func(hx_value_707 any) *string {
				if hx_value_707 == nil {
					var hx_zero_708 *string
					return hx_zero_708
				}
				return hx_value_707.(*string)
			}(any(arg_0)))))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_711 any) *string {
				if hx_value_711 == nil {
					var hx_zero_712 *string
					return hx_zero_712
				}
				return hx_value_711.(*string)
			}(any(hxrt.FunctionCall(source)(arg_0)))
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_715 any) *string {
				if hx_value_715 == nil {
					var hx_zero_716 *string
					return hx_zero_716
				}
				return hx_value_715.(*string)
			}(any(hxrt.FunctionCall(source)(func(hx_value_713 any) *string {
				if hx_value_713 == nil {
					var hx_zero_714 *string
					return hx_zero_714
				}
				return hx_value_713.(*string)
			}(any(arg_0)))))
		})
	case *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_717 any) *string {
				if hx_value_717 == nil {
					var hx_zero_718 *string
					return hx_zero_718
				}
				return hx_value_717.(*string)
			}(any(hxrt.FunctionCall(source)(hx_recover_class_2(arg_0))))
		})
	case *hxrt.Function[func(*string) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_721 any) *string {
				if hx_value_721 == nil {
					var hx_zero_722 *string
					return hx_zero_722
				}
				return hx_value_721.(*string)
			}(any(hxrt.FunctionCall(source)(func(hx_value_719 any) *string {
				if hx_value_719 == nil {
					var hx_zero_720 *string
					return hx_zero_720
				}
				return hx_value_719.(*string)
			}(any(arg_0)))))
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_723 any) *string {
				if hx_value_723 == nil {
					var hx_zero_724 *string
					return hx_zero_724
				}
				return hx_value_723.(*string)
			}(any(hxrt.FunctionCall(source)(hx_recover_class_3(arg_0))))
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor)]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			hxrt.FunctionCall(source)(hx_recover_class_3(arg_0))
			return func(hx_value_725 any) *string {
				if hx_value_725 == nil {
					var hx_zero_726 *string
					return hx_zero_726
				}
				return hx_value_725.(*string)
			}(any(nil))
		})
	case *hxrt.Function[func(*haxe___Template__TemplateExpr)]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			hxrt.FunctionCall(source)(func(hx_value_727 any) *haxe___Template__TemplateExpr {
				if hx_value_727 == nil {
					var hx_zero_728 *haxe___Template__TemplateExpr
					return hx_zero_728
				}
				return hx_value_727.(*haxe___Template__TemplateExpr)
			}(any(arg_0)))
			return func(hx_value_729 any) *string {
				if hx_value_729 == nil {
					var hx_zero_730 *string
					return hx_zero_730
				}
				return hx_value_729.(*string)
			}(any(nil))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_733 any) *string {
				if hx_value_733 == nil {
					var hx_zero_734 *string
					return hx_zero_734
				}
				return hx_value_733.(*string)
			}(any(hxrt.FunctionCall(source)(func(hx_value_731 any) *string {
				if hx_value_731 == nil {
					var hx_zero_732 *string
					return hx_zero_732
				}
				return hx_value_731.(*string)
			}(any(arg_0)))))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return hxrt.FunctionCall(source)(func(hx_value_735 any) int {
				if hx_value_735 == nil {
					var hx_zero_736 int
					return hx_zero_736
				}
				return hx_value_735.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_739 any) *string {
				if hx_value_739 == nil {
					var hx_zero_740 *string
					return hx_zero_740
				}
				return hx_value_739.(*string)
			}(any(hxrt.FunctionCall(source)(func(hx_value_737 any) int {
				if hx_value_737 == nil {
					var hx_zero_738 int
					return hx_zero_738
				}
				return hx_value_737.(int)
			}(any(arg_0)))))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_741 any) *string {
				if hx_value_741 == nil {
					var hx_zero_742 *string
					return hx_zero_742
				}
				return hx_value_741.(*string)
			}(any(hxrt.FunctionCall(source)(hx_recover_class_4(arg_0))))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_745 any) *string {
				if hx_value_745 == nil {
					var hx_zero_746 *string
					return hx_zero_746
				}
				return hx_value_745.(*string)
			}(any(hxrt.FunctionCall(source)(func(hx_value_743 any) int {
				if hx_value_743 == nil {
					var hx_zero_744 int
					return hx_zero_744
				}
				return hx_value_743.(int)
			}(any(arg_0)))))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_749 any) *string {
				if hx_value_749 == nil {
					var hx_zero_750 *string
					return hx_zero_750
				}
				return hx_value_749.(*string)
			}(any(hxrt.FunctionCall(source)(func(hx_value_747 any) int {
				if hx_value_747 == nil {
					var hx_zero_748 int
					return hx_zero_748
				}
				return hx_value_747.(int)
			}(any(arg_0)))))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__ValueException {
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

func hx_recover_class_1(value any) *haxe__Exception {
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

func hx_recover_class_2(value any) *haxe___Template__TokenCursor {
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

func hx_recover_class_3(value any) *haxe___Template__ExprCursor {
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

func hx_recover_class_4(value any) *haxe__io__Bytes {
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

func hx_recover_class_5(value any) *haxe___Int64_____Int64 {
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
