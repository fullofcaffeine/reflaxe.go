package main

import (
	"math"
	"snapshot/hxrt"
)

type I_haxe__Serializer interface {
	toString() *string
	serializeString(value *string)
	serializeRef(value any) bool
	serializeFields(value any)
	serializeArray(value *hxrt.Array)
	flushNulls(count int)
	serializeBytes(value *haxe__io__Bytes)
	serializeClass(value any, declaration any)
	serializeEnum(value any, declaration any)
	serialize(value any)
	serializeException(value any)
}

type haxe__Serializer struct {
	__hx_this    I_haxe__Serializer
	buf          *StringBuf
	cache        *hxrt.Array
	shash        *haxe__ds__StringMap
	scount       int
	useCache     bool
	useEnumIndex bool
}

func (self *haxe__Serializer) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__Serializer() *haxe__Serializer {
	return hx_init_haxe__Serializer(hx_alloc_haxe__Serializer(nil))
}

func hx_alloc_haxe__Serializer(owner any) *haxe__Serializer {
	self := &haxe__Serializer{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__Serializer)
	}
	return self
}

func hx_init_haxe__Serializer(self *haxe__Serializer) *haxe__Serializer {
	self.buf = New_StringBuf()
	self.cache = hxrt.NewArray()
	self.useCache = haxe__Serializer_USE_CACHE
	self.useEnumIndex = haxe__Serializer_USE_ENUM_INDEX
	self.shash = New_haxe__ds__StringMap()
	self.scount = 0
	return self
}

func (self *haxe__Serializer) toString() *string {
	return self.buf.b
}

func (self *haxe__Serializer) serializeString(value *string) {
	var known any = func(hx_value_1 any) any {
		if hx_value_1 == nil {
			return nil
		}
		return hx_value_1.(int)
	}(self.shash.__hx_this.get(value))
	if known != nil {
		_this := self.buf
		_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("R"))
		_this_1 := self.buf
		_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StdString(known.(int)))
		return
	}
	self.shash.__hx_this.set(value, func() int {
		hx_post_2 := self.scount
		self.scount = int(int32((self.scount + 1)))
		return hx_post_2
	}())
	_this_2 := self.buf
	_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StringFromLiteral("y"))
	encoded := StringTools_urlEncode(value)
	_this_3 := self.buf
	x := hxrt.StringLengthStringPtr(encoded)
	_this_3.b = hxrt.StringConcatStringPtr(_this_3.b, hxrt.StdString(x))
	_this_4 := self.buf
	_this_4.b = hxrt.StringConcatStringPtr(_this_4.b, hxrt.StringFromLiteral(":"))
	_this_5 := self.buf
	_this_5.b = hxrt.StringConcatStringPtr(_this_5.b, hxrt.StdString(encoded))
}

func (self *haxe__Serializer) serializeRef(value any) bool {
	_g := 0
	_g1 := self.cache.Len()
	for _g < _g1 {
		hx_post_3 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_3
		if hxrt.HaxeEqual(self.cache.Get(index), value) {
			_this := self.buf
			_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("r"))
			_this_1 := self.buf
			_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StdString(index))
			return true
		}
	}
	hx_arr_4 := self.cache
	hx_arr_4.Push(value)
	return false
}

func (self *haxe__Serializer) serializeFields(value any) {
	_g := 0
	_g1 := Reflect_fields(value)
	for _g < _g1.Len() {
		field := func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		self.__hx_this.serializeString(field)
		self.__hx_this.serialize(Reflect_field(value, field))
	}
	_this := self.buf
	_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("g"))
}

func (self *haxe__Serializer) serializeArray(value *hxrt.Array) {
	nullCount := 0
	_this := self.buf
	_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("a"))
	_g := 0
	for _g < value.Len() {
		var item any = value.Get(_g)
		_g = int(int32((_g + 1)))
		if hxrt.AnyEqualsNull(item) {
			nullCount = int(int32((nullCount + 1)))
			continue
		}
		if nullCount == 1 {
			_this_1 := self.buf
			_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StringFromLiteral("n"))
		} else {
			if nullCount > 1 {
				_this_2 := self.buf
				_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StringFromLiteral("u"))
				_this_3 := self.buf
				_this_3.b = hxrt.StringConcatStringPtr(_this_3.b, hxrt.StdString(nullCount))
			}
		}
		nullCount = 0
		self.__hx_this.serialize(item)
	}
	if nullCount == 1 {
		_this_4 := self.buf
		_this_4.b = hxrt.StringConcatStringPtr(_this_4.b, hxrt.StringFromLiteral("n"))
	} else {
		if nullCount > 1 {
			_this_5 := self.buf
			_this_5.b = hxrt.StringConcatStringPtr(_this_5.b, hxrt.StringFromLiteral("u"))
			_this_6 := self.buf
			_this_6.b = hxrt.StringConcatStringPtr(_this_6.b, hxrt.StdString(nullCount))
		}
	}
	_this_7 := self.buf
	_this_7.b = hxrt.StringConcatStringPtr(_this_7.b, hxrt.StringFromLiteral("h"))
}

func (self *haxe__Serializer) flushNulls(count int) {
	if count == 1 {
		_this := self.buf
		_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("n"))
	} else {
		if count > 1 {
			_this_1 := self.buf
			_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StringFromLiteral("u"))
			_this_2 := self.buf
			_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StdString(count))
		}
	}
}

func (self *haxe__Serializer) serializeBytes(value *haxe__io__Bytes) {
	_this := self.buf
	_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("s"))
	_this_1 := self.buf
	v := (float64(int((hxrt.Int32Wrap(value.length) * hxrt.Int32Wrap(8)))) / float64(6))
	x := hxrt.MathCeilInt(v)
	_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StdString(x))
	_this_2 := self.buf
	_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StringFromLiteral(":"))
	index := 0
	max := int((hxrt.Int32Wrap(value.length) - hxrt.Int32Wrap(2)))
	for index < max {
		hx_post_7 := index
		index = int(int32((index + 1)))
		pos := hx_post_7
		first := value.b[pos]
		hx_post_8 := index
		index = int(int32((index + 1)))
		pos_1 := hx_post_8
		second := value.b[pos_1]
		hx_post_9 := index
		index = int(int32((index + 1)))
		pos_2 := hx_post_9
		third := value.b[pos_2]
		_this_3 := self.buf
		c := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int((hxrt.Int32Wrap(first) >> uint(2)))))
		_this_3.b = hxrt.StringConcatStringPtr(_this_3.b, hxrt.StringFromCharCode(c))
		_this_4 := self.buf
		c_1 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(first) << uint(4)))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(second) >> uint(4))))))) & hxrt.Int32Wrap(63)))))
		_this_4.b = hxrt.StringConcatStringPtr(_this_4.b, hxrt.StringFromCharCode(c_1))
		_this_5 := self.buf
		c_2 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(second) << uint(2)))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(third) >> uint(6))))))) & hxrt.Int32Wrap(63)))))
		_this_5.b = hxrt.StringConcatStringPtr(_this_5.b, hxrt.StringFromCharCode(c_2))
		_this_6 := self.buf
		c_3 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int((hxrt.Int32Wrap(third) & hxrt.Int32Wrap(63)))))
		_this_6.b = hxrt.StringConcatStringPtr(_this_6.b, hxrt.StringFromCharCode(c_3))
	}
	if index == max {
		hx_post_10 := index
		index = int(int32((index + 1)))
		pos_3 := hx_post_10
		first_1 := value.b[pos_3]
		hx_post_11 := index
		index = int(int32((index + 1)))
		pos_4 := hx_post_11
		second_1 := value.b[pos_4]
		_this_7 := self.buf
		c_4 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int((hxrt.Int32Wrap(first_1) >> uint(2)))))
		_this_7.b = hxrt.StringConcatStringPtr(_this_7.b, hxrt.StringFromCharCode(c_4))
		_this_8 := self.buf
		c_5 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(first_1) << uint(4)))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(second_1) >> uint(4))))))) & hxrt.Int32Wrap(63)))))
		_this_8.b = hxrt.StringConcatStringPtr(_this_8.b, hxrt.StringFromCharCode(c_5))
		_this_9 := self.buf
		c_6 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(second_1) << uint(2)))) & hxrt.Int32Wrap(63)))))
		_this_9.b = hxrt.StringConcatStringPtr(_this_9.b, hxrt.StringFromCharCode(c_6))
	} else {
		if index == int((hxrt.Int32Wrap(max) + hxrt.Int32Wrap(1))) {
			hx_post_12 := index
			index = int(int32((index + 1)))
			pos_5 := hx_post_12
			first_2 := value.b[pos_5]
			_this_10 := self.buf
			c_7 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int((hxrt.Int32Wrap(first_2) >> uint(2)))))
			_this_10.b = hxrt.StringConcatStringPtr(_this_10.b, hxrt.StringFromCharCode(c_7))
			_this_11 := self.buf
			c_8 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(first_2) << uint(4)))) & hxrt.Int32Wrap(63)))))
			_this_11.b = hxrt.StringConcatStringPtr(_this_11.b, hxrt.StringFromCharCode(c_8))
		}
	}
}

func (self *haxe__Serializer) serializeClass(value any, declaration any) {
	className := Type_getClassName(declaration)
	if hxrt.StringEqualStringPtr(className, hxrt.StringFromLiteral("String")) {
		self.__hx_this.serializeString(hxrt.StdString(value))
		return
	}
	if self.useCache && self.__hx_this.serializeRef(value) {
		return
	}
	switch *hxrt.StdString(className) {
	case *hxrt.StdString(hxrt.StringFromLiteral("Array")):
		self.__hx_this.serializeArray(func(hx_value_15 any) *hxrt.Array {
			if hx_value_15 == nil {
				var hx_zero_16 *hxrt.Array
				return hx_zero_16
			}
			return hx_value_15.(*hxrt.Array)
		}(any(func(hx_value_13 any) *hxrt.Array {
			if hx_value_13 == nil {
				var hx_zero_14 *hxrt.Array
				return hx_zero_14
			}
			return hx_value_13.(*hxrt.Array)
		}(value))))
	case *hxrt.StdString(hxrt.StringFromLiteral("Date")):
		date := hx_recover_class_5(hx_recover_class_5(value))
		_this := self.buf
		_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("v"))
		_this_1 := self.buf
		x := date.ms
		_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StdString(x))
	case *hxrt.StdString(hxrt.StringFromLiteral("haxe.ds.IntMap")):
		_this_2 := self.buf
		_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StringFromLiteral("q"))
		map_ := hx_recover_class_4(hx_recover_class_4(value))
		key := map_.__hx_this.keys()
		for func() bool {
			var hx_callable_17 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(key, "hasNext"))
			return hxrt.FunctionCall(hx_callable_17)()
		}() {
			key_1 := func() int {
				var hx_callable_18 *hxrt.Function[func() int] = hx_recover_callable_3(hxrt.AnonymousGet(key, "next"))
				return hxrt.FunctionCall(hx_callable_18)()
			}()
			_this_3 := self.buf
			_this_3.b = hxrt.StringConcatStringPtr(_this_3.b, hxrt.StringFromLiteral(":"))
			_this_4 := self.buf
			_this_4.b = hxrt.StringConcatStringPtr(_this_4.b, hxrt.StdString(key_1))
			self.__hx_this.serialize(map_.__hx_this.get(key_1))
		}
		_this_5 := self.buf
		_this_5.b = hxrt.StringConcatStringPtr(_this_5.b, hxrt.StringFromLiteral("h"))
	case *hxrt.StdString(hxrt.StringFromLiteral("haxe.ds.List")):
		_this_6 := self.buf
		_this_6.b = hxrt.StringConcatStringPtr(_this_6.b, hxrt.StringFromLiteral("l"))
		list := hx_recover_class_6(hx_recover_class_6(value))
		item := hx_recover_class_3(list.__hx_this.iterator())
		for func(hx_value_19 any) bool {
			if hx_value_19 == nil {
				var hx_zero_20 bool
				return hx_zero_20
			}
			return hx_value_19.(bool)
		}(item.__hx_this.hasNext()) {
			var item_1 any = item.__hx_this.next()
			self.__hx_this.serialize(item_1)
		}
		_this_7 := self.buf
		_this_7.b = hxrt.StringConcatStringPtr(_this_7.b, hxrt.StringFromLiteral("h"))
	case *hxrt.StdString(hxrt.StringFromLiteral("haxe.ds.ObjectMap")):
		_this_8 := self.buf
		_this_8.b = hxrt.StringConcatStringPtr(_this_8.b, hxrt.StringFromLiteral("M"))
		map__1 := hx_recover_class_2(hx_recover_class_2(value))
		key_2 := map__1.__hx_this.keys()
		for func() bool {
			var hx_callable_21 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(key_2, "hasNext"))
			return hxrt.FunctionCall(hx_callable_21)()
		}() {
			var key_3 any = func() any {
				var hx_callable_22 *hxrt.Function[func() any] = hx_recover_callable_2(hxrt.AnonymousGet(key_2, "next"))
				return hxrt.FunctionCall(hx_callable_22)()
			}()
			self.__hx_this.serialize(key_3)
			self.__hx_this.serialize(map__1.__hx_this.get(key_3))
		}
		_this_9 := self.buf
		_this_9.b = hxrt.StringConcatStringPtr(_this_9.b, hxrt.StringFromLiteral("h"))
	case *hxrt.StdString(hxrt.StringFromLiteral("haxe.ds.StringMap")):
		_this_10 := self.buf
		_this_10.b = hxrt.StringConcatStringPtr(_this_10.b, hxrt.StringFromLiteral("b"))
		map__2 := hx_recover_class_1(hx_recover_class_1(value))
		key_4 := map__2.__hx_this.keys()
		for func() bool {
			var hx_callable_23 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(key_4, "hasNext"))
			return hxrt.FunctionCall(hx_callable_23)()
		}() {
			key_5 := func() *string {
				var hx_callable_24 *hxrt.Function[func() *string] = hx_recover_callable_1(hxrt.AnonymousGet(key_4, "next"))
				return hxrt.FunctionCall(hx_callable_24)()
			}()
			self.__hx_this.serializeString(key_5)
			self.__hx_this.serialize(map__2.__hx_this.get(key_5))
		}
		_this_11 := self.buf
		_this_11.b = hxrt.StringConcatStringPtr(_this_11.b, hxrt.StringFromLiteral("h"))
	case *hxrt.StdString(hxrt.StringFromLiteral("haxe.io.Bytes")):
		self.__hx_this.serializeBytes(hx_recover_class_7(hx_recover_class_7(value)))
	default:
		if self.useCache {
			hx_arr_25 := self.cache
			hx_arr_25.Pop()
		}
		var hook any = Reflect_field(value, hxrt.StringFromLiteral("hxSerialize"))
		if !hxrt.AnyEqualsNull(hook) {
			_this_12 := self.buf
			_this_12.b = hxrt.StringConcatStringPtr(_this_12.b, hxrt.StringFromLiteral("C"))
			self.__hx_this.serializeString(className)
			if self.useCache {
				hx_arr_26 := self.cache
				hx_arr_26.Push(value)
			}
			Reflect_callMethod(value, hook, hxrt.NewArray(self))
			_this_13 := self.buf
			_this_13.b = hxrt.StringConcatStringPtr(_this_13.b, hxrt.StringFromLiteral("g"))
		} else {
			_this_14 := self.buf
			_this_14.b = hxrt.StringConcatStringPtr(_this_14.b, hxrt.StringFromLiteral("c"))
			self.__hx_this.serializeString(className)
			if self.useCache {
				hx_arr_27 := self.cache
				hx_arr_27.Push(value)
			}
			self.__hx_this.serializeFields(value)
		}
	}
}

func (self *haxe__Serializer) serializeEnum(value any, declaration any) {
	if self.useCache {
		if self.__hx_this.serializeRef(value) {
			return
		}
		hx_arr_28 := self.cache
		hx_arr_28.Pop()
	}
	_this := self.buf
	var hx_if_29 *string
	if self.useEnumIndex {
		hx_if_29 = hxrt.StringFromLiteral("j")
	} else {
		hx_if_29 = hxrt.StringFromLiteral("w")
	}
	x := hx_if_29
	_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StdString(x))
	self.__hx_this.serializeString(Type_getEnumName(declaration))
	if self.useEnumIndex {
		_this_1 := self.buf
		_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StringFromLiteral(":"))
		_this_2 := self.buf
		x_1 := Type_enumIndex(value)
		_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StdString(x_1))
	} else {
		self.__hx_this.serializeString(Type_enumConstructor(value))
	}
	_this_3 := self.buf
	_this_3.b = hxrt.StringConcatStringPtr(_this_3.b, hxrt.StringFromLiteral(":"))
	parameters := Type_enumParameters(value)
	_this_4 := self.buf
	x_2 := parameters.Len()
	_this_4.b = hxrt.StringConcatStringPtr(_this_4.b, hxrt.StdString(x_2))
	_g := 0
	for _g < parameters.Len() {
		var parameter any = parameters.Get(_g)
		_g = int(int32((_g + 1)))
		self.__hx_this.serialize(parameter)
	}
	if self.useCache {
		hx_arr_30 := self.cache
		hx_arr_30.Push(value)
	}
}

func (self *haxe__Serializer) serialize(value any) {
	if hx_is_nominal_0(value) {
		if self.useCache && self.__hx_this.serializeRef(value) {
			return
		}
		self.__hx_this.serializeBytes(hx_recover_class_7(hx_recover_class_7(value)))
		return
	}
	_g := Type_typeof(value)
	switch _g.tag {
	case 0:
		_this := self.buf
		_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("n"))
	case 1:
		integer := hxrt.IntFromNullableAny(value)
		if integer == 0 {
			_this_1 := self.buf
			_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StringFromLiteral("z"))
		} else {
			_this_2 := self.buf
			_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StringFromLiteral("i"))
			_this_3 := self.buf
			_this_3.b = hxrt.StringConcatStringPtr(_this_3.b, hxrt.StdString(integer))
		}
	case 2:
		number := func(hx_value_31 any) float64 {
			if hx_value_31 == nil {
				var hx_zero_32 float64
				return hx_zero_32
			}
			return hx_value_31.(float64)
		}(value)
		if math.IsNaN(number) {
			_this_4 := self.buf
			_this_4.b = hxrt.StringConcatStringPtr(_this_4.b, hxrt.StringFromLiteral("k"))
		} else {
			if !(!math.IsInf(number, 0) && !math.IsNaN(number)) {
				_this_5 := self.buf
				var hx_if_33 *string
				if number < 0 {
					hx_if_33 = hxrt.StringFromLiteral("m")
				} else {
					hx_if_33 = hxrt.StringFromLiteral("p")
				}
				_this_5.b = hxrt.StringConcatStringPtr(_this_5.b, hx_if_33)
			} else {
				_this_6 := self.buf
				_this_6.b = hxrt.StringConcatStringPtr(_this_6.b, hxrt.StringFromLiteral("d"))
				_this_7 := self.buf
				_this_7.b = hxrt.StringConcatStringPtr(_this_7.b, hxrt.StdString(number))
			}
		}
	case 3:
		boolean := func(hx_value_34 any) bool {
			if hx_value_34 == nil {
				var hx_zero_35 bool
				return hx_zero_35
			}
			return hx_value_34.(bool)
		}(value)
		_this_8 := self.buf
		var hx_if_36 *string
		if boolean {
			hx_if_36 = hxrt.StringFromLiteral("t")
		} else {
			hx_if_36 = hxrt.StringFromLiteral("f")
		}
		_this_8.b = hxrt.StringConcatStringPtr(_this_8.b, hx_if_36)
	case 4:
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *hxrt__TypeClassValue:
				return true
			default:
				return false
			}
		}(value) {
			_this_9 := self.buf
			_this_9.b = hxrt.StringConcatStringPtr(_this_9.b, hxrt.StringFromLiteral("A"))
			self.__hx_this.serializeString(Type_getClassName(value))
		} else {
			if func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt__TypeEnumValue:
					return true
				default:
					return false
				}
			}(value) {
				_this_10 := self.buf
				_this_10.b = hxrt.StringConcatStringPtr(_this_10.b, hxrt.StringFromLiteral("B"))
				self.__hx_this.serializeString(Type_getEnumName(value))
			} else {
				if self.useCache && self.__hx_this.serializeRef(value) {
					return
				}
				_this_11 := self.buf
				_this_11.b = hxrt.StringConcatStringPtr(_this_11.b, hxrt.StringFromLiteral("o"))
				self.__hx_this.serializeFields(value)
			}
		}
	case 5:
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Cannot serialize function")))
	case 6:
		var _g_1 any = _g.params[0]
		var declaration any = _g_1
		self.__hx_this.serializeClass(value, declaration)
	case 7:
		var _g_2 any = _g.params[0]
		var declaration_1 any = _g_2
		self.__hx_this.serializeEnum(value, declaration_1)
	case 8:
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot serialize "), hxrt.StdString(value))))
	}
}

func (self *haxe__Serializer) serializeException(value any) {
	_this := self.buf
	_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("x"))
	self.__hx_this.serialize(value)
}

func (self *haxe__Serializer) String() string {
	return *self.__hx_this.toString()
}

var haxe__Serializer_BASE64 *string = hxrt.StringFromLiteral("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789%:")

var haxe__Serializer_USE_CACHE bool = false

var haxe__Serializer_USE_ENUM_INDEX bool = false

func haxe__Serializer_run(value any) *string {
	serializer := New_haxe__Serializer()
	serializer.__hx_this.serialize(value)
	return serializer.__hx_this.toString()
}
