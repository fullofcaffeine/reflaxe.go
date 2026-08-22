package main

import "snapshot/hxrt"

type I_haxe__Unserializer interface {
	setResolver(value any)
	getResolver() any
	get(index int) int
	readDigits() int
	readFloat() float64
	unserializeObject(target any)
	resolveClass(name *string) any
	resolveEnum(name *string) any
	unserializeEnum(declaration any, tag *string) any
	decodeBytes(encoded *string) *haxe__io__Bytes
	base64Value(code int) int
	unserialize() any
	isLegacyDate() bool
}

type haxe__Unserializer struct {
	__hx_this I_haxe__Unserializer
	buf       *string
	pos       int
	length    int
	cache     *hxrt.Array
	scache    *hxrt.Array
	resolver  any
}

func New_haxe__Unserializer(buf *string) *haxe__Unserializer {
	self := &haxe__Unserializer{}
	self.__hx_this = self
	self.buf = buf
	self.length = hxrt.StringLengthStringPtr(buf)
	self.pos = 0
	self.scache = hxrt.NewArray()
	self.cache = hxrt.NewArray()
	current := haxe__Unserializer_DEFAULT_RESOLVER
	if current == nil {
		hx_obj_1 := map[string]any{}
		hx_obj_1["resolveClass"] = func(name *string) any {
			return Type_resolveClass(name)
		}
		hx_obj_1["resolveEnum"] = func(name *string) any {
			return Type_resolveEnum(name)
		}
		current = hx_obj_1
		haxe__Unserializer_DEFAULT_RESOLVER = current
	}
	self.resolver = current
	return self
}

func (self *haxe__Unserializer) setResolver(value any) {
	var hx_if_2 any
	if hxrt.AnyEqualsNull(value) {
		hx_if_2 = haxe__Unserializer_NULL_RESOLVER
	} else {
		hx_if_2 = value
	}
	self.resolver = hx_if_2
}

func (self *haxe__Unserializer) getResolver() any {
	return self.resolver
}

func (self *haxe__Unserializer) get(index int) int {
	s := self.buf
	var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
	var hx_if_3 int
	if c == nil {
		hx_if_3 = -1
	} else {
		hx_if_3 = c.(int)
	}
	return hx_if_3
}

func (self *haxe__Unserializer) readDigits() int {
	value := 0
	negative := false
	start := self.pos
	for true {
		index := self.pos
		s := self.buf
		var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
		var hx_if_4 int
		if c == nil {
			hx_if_4 = -1
		} else {
			hx_if_4 = c.(int)
		}
		code := hx_if_4
		if code == -1 {
			break
		}
		if code == 45 {
			if self.pos != start {
				break
			}
			negative = true
			self.pos = int(int32((self.pos + 1)))
			continue
		}
		if (code < 48) || (code > 57) {
			break
		}
		value = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(10))))) + hxrt.Int32Wrap(code))))) - hxrt.Int32Wrap(48))))
		self.pos = int(int32((self.pos + 1)))
	}
	var hx_if_5 int
	if negative {
		hx_if_5 = int(int32(-int32(value)))
	} else {
		hx_if_5 = value
	}
	return hx_if_5
}

func (self *haxe__Unserializer) readFloat() float64 {
	start := self.pos
	for true {
		index := self.pos
		s := self.buf
		var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
		var hx_if_6 int
		if c == nil {
			hx_if_6 = -1
		} else {
			hx_if_6 = c.(int)
		}
		code := hx_if_6
		if code == -1 {
			break
		}
		if (((code >= 43) && (code < 58)) || (code == 101)) || (code == 69) {
			self.pos = int(int32((self.pos + 1)))
		} else {
			break
		}
	}
	return hxrt.SerializationParseFloat(hxrt.StringSubstrStringPtr(self.buf, start, int(int32((hxrt.Int32Wrap(self.pos) - hxrt.Int32Wrap(start)))), true))
}

func (self *haxe__Unserializer) unserializeObject(target any) {
	for true {
		if self.pos >= self.length {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid object"))
		}
		if func() int {
			index := self.pos
			s := self.buf
			var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
			var hx_if_7 int
			if c == nil {
				hx_if_7 = -1
			} else {
				hx_if_7 = c.(int)
			}
			return hx_if_7
		}() == 103 {
			break
		}
		var key any = self.__hx_this.unserialize()
		if !func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			case string:
				return true
			default:
				return false
			}
		}(any(key)) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid object key"))
		}
		Reflect_setField(target, hxrt.StdString(func(hx_value_8 any) *string {
			if hx_value_8 == nil {
				var hx_zero_9 *string
				return hx_zero_9
			}
			return hx_value_8.(*string)
		}(key)), self.__hx_this.unserialize())
	}
	self.pos = int(int32((self.pos + 1)))
}

func (self *haxe__Unserializer) resolveClass(name *string) any {
	var method any = Reflect_field(self.resolver, hxrt.StringFromLiteral("resolveClass"))
	var hx_if_10 any
	if hxrt.AnyEqualsNull(method) {
		hx_if_10 = nil
	} else {
		hx_if_10 = Reflect_callMethod(self.resolver, method, hxrt.NewArray(name))
	}
	return hx_if_10
}

func (self *haxe__Unserializer) resolveEnum(name *string) any {
	var method any = Reflect_field(self.resolver, hxrt.StringFromLiteral("resolveEnum"))
	var hx_if_11 any
	if hxrt.AnyEqualsNull(method) {
		hx_if_11 = nil
	} else {
		hx_if_11 = Reflect_callMethod(self.resolver, method, hxrt.NewArray(name))
	}
	return hx_if_11
}

func (self *haxe__Unserializer) unserializeEnum(declaration any, tag *string) any {
	if func() int {
		hx_post_12 := self.pos
		self.pos = int(int32((self.pos + 1)))
		index := hx_post_12
		s := self.buf
		var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
		var hx_if_13 int
		if c == nil {
			hx_if_13 = -1
		} else {
			hx_if_13 = c.(int)
		}
		return hx_if_13
	}() != 58 {
		hxrt.Throw(hxrt.StringFromLiteral("Invalid enum format"))
	}
	count := self.__hx_this.readDigits()
	if count == 0 {
		return Type_createEnum(declaration, tag, hxrt.NewArray())
	}
	arguments := hxrt.NewArray()
	for func() int {
		hx_post_14 := count
		count = int(int32((count - 1)))
		return hx_post_14
	}() > 0 {
		arguments.Push(self.__hx_this.unserialize())
	}
	return Type_createEnum(declaration, tag, arguments)
}

func (self *haxe__Unserializer) decodeBytes(encoded *string) *haxe__io__Bytes {
	rest := int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(encoded)) & hxrt.Int32Wrap(3))))
	size := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(encoded)) >> uint(2))))) * hxrt.Int32Wrap(3))))) + hxrt.Int32Wrap(func() int {
		var hx_if_16 int
		if rest >= 2 {
			hx_if_16 = int(int32((hxrt.Int32Wrap(rest) - hxrt.Int32Wrap(1))))
		} else {
			hx_if_16 = 0
		}
		return hx_if_16
	}()))))
	bytes := haxe__io__Bytes_alloc(size)
	index := 0
	output := 0
	complete := int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(encoded)) - hxrt.Int32Wrap(rest))))
	for index < complete {
		first := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_17 := index
			index = int(int32((index + 1)))
			return hx_post_17
		}())))
		second := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_18 := index
			index = int(int32((index + 1)))
			return hx_post_18
		}())))
		third := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_19 := index
			index = int(int32((index + 1)))
			return hx_post_19
		}())))
		fourth := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_20 := index
			index = int(int32((index + 1)))
			return hx_post_20
		}())))
		hx_post_21 := output
		output = int(int32((output + 1)))
		pos := hx_post_21
		bytes.b[pos] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(first) << uint(2))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second) >> uint(4))))))))) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
		hx_post_22 := output
		output = int(int32((output + 1)))
		pos_1 := hx_post_22
		bytes.b[pos_1] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second) << uint(4))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(third) >> uint(2))))))))) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
		hx_post_23 := output
		output = int(int32((output + 1)))
		pos_2 := hx_post_23
		bytes.b[pos_2] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(third) << uint(6))))) | hxrt.Int32Wrap(fourth))))) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	if rest >= 2 {
		first_1 := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_24 := index
			index = int(int32((index + 1)))
			return hx_post_24
		}())))
		second_1 := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_25 := index
			index = int(int32((index + 1)))
			return hx_post_25
		}())))
		hx_post_26 := output
		output = int(int32((output + 1)))
		pos_3 := hx_post_26
		bytes.b[pos_3] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(first_1) << uint(2))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second_1) >> uint(4))))))))) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
		if rest == 3 {
			third_1 := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
				hx_post_27 := index
				index = int(int32((index + 1)))
				return hx_post_27
			}())))
			hx_post_28 := output
			output = int(int32((output + 1)))
			pos_4 := hx_post_28
			bytes.b[pos_4] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second_1) << uint(4))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(third_1) >> uint(2))))))))) & hxrt.Int32Wrap(255))))
			bytes.__hx_rawValid = false
		}
	}
	return bytes
}

func (self *haxe__Unserializer) base64Value(code int) int {
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(haxe__Unserializer_BASE64)
	for _g < _g1 {
		hx_post_29 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_29
		if hxrt.StringCharCodeAtAnyStringPtr(haxe__Unserializer_BASE64, index) == code {
			return index
		}
	}
	return -1
}

func (self *haxe__Unserializer) unserialize() any {
	hx_post_30 := self.pos
	self.pos = int(int32((self.pos + 1)))
	index := hx_post_30
	s := self.buf
	var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
	var hx_if_31 int
	if c == nil {
		hx_if_31 = -1
	} else {
		hx_if_31 = c.(int)
	}
	_g := hx_if_31
	switch _g {
	case 65:
		name := hxrt.StdString(self.__hx_this.unserialize())
		var declaration any = self.__hx_this.resolveClass(name)
		if hxrt.AnyEqualsNull(declaration) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Class not found "), name))
		}
		return declaration
	case 66:
		name_1 := hxrt.StdString(self.__hx_this.unserialize())
		var declaration_1 any = self.__hx_this.resolveEnum(name_1)
		if hxrt.AnyEqualsNull(declaration_1) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Enum not found "), name_1))
		}
		return declaration_1
	case 67:
		name_2 := hxrt.StdString(self.__hx_this.unserialize())
		var declaration_2 any = self.__hx_this.resolveClass(name_2)
		if hxrt.AnyEqualsNull(declaration_2) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Class not found "), name_2))
		}
		var object any = Type_createEmptyInstance(declaration_2)
		hx_arr_32 := self.cache
		hx_arr_32.Push(object)
		var hook any = Reflect_field(object, hxrt.StringFromLiteral("hxUnserialize"))
		if hxrt.AnyEqualsNull(hook) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Class "), name_2), hxrt.StringFromLiteral(" has no hxUnserialize hook")))
		}
		Reflect_callMethod(object, hook, hxrt.NewArray(self))
		if func() int {
			hx_post_33 := self.pos
			self.pos = int(int32((self.pos + 1)))
			index_1 := hx_post_33
			s_1 := self.buf
			var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(s_1, index_1)
			var hx_if_34 int
			if c_1 == nil {
				hx_if_34 = -1
			} else {
				hx_if_34 = c_1.(int)
			}
			return hx_if_34
		}() != 103 {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid custom data"))
		}
		return object
	case 77:
		map_ := New_haxe__ds__ObjectMap()
		hx_arr_35 := self.cache
		hx_arr_35.Push(map_)
		for func() int {
			index_2 := self.pos
			s_2 := self.buf
			var c_2 any = hxrt.StringCharCodeAtAnyStringPtr(s_2, index_2)
			var hx_if_36 int
			if c_2 == nil {
				hx_if_36 = -1
			} else {
				hx_if_36 = c_2.(int)
			}
			return hx_if_36
		}() != 104 {
			map_.__hx_this.set(self.__hx_this.unserialize(), self.__hx_this.unserialize())
		}
		self.pos = int(int32((self.pos + 1)))
		return map_
	case 82:
		index_3 := self.__hx_this.readDigits()
		if (index_3 < 0) || (index_3 >= self.scache.Len()) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid string reference"))
		}
		return self.scache.Get(index_3)
	case 97:
		array := hxrt.NewArray()
		hx_arr_37 := self.cache
		hx_arr_37.Push(array)
	hx_loop_39:
		for true {
			index_4 := self.pos
			s_3 := self.buf
			var c_3 any = hxrt.StringCharCodeAtAnyStringPtr(s_3, index_4)
			var hx_if_38 int
			if c_3 == nil {
				hx_if_38 = -1
			} else {
				hx_if_38 = c_3.(int)
			}
			code := hx_if_38
			if code == 104 {
				self.pos = int(int32((self.pos + 1)))
				break hx_loop_39
			}
			if code == 117 {
				self.pos = int(int32((self.pos + 1)))
				count := self.__hx_this.readDigits()
				hx_array_target_40 := array
				hx_array_index_41 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(array.Len()) + hxrt.Int32Wrap(count))))) - hxrt.Int32Wrap(1))))
				hx_array_target_40.Set(hx_array_index_41, nil)
			} else {
				array.Push(self.__hx_this.unserialize())
			}
		}
		return array
	case 98:
		map__1 := New_haxe__ds__StringMap()
		hx_arr_43 := self.cache
		hx_arr_43.Push(map__1)
		for func() int {
			index_5 := self.pos
			s_4 := self.buf
			var c_4 any = hxrt.StringCharCodeAtAnyStringPtr(s_4, index_5)
			var hx_if_44 int
			if c_4 == nil {
				hx_if_44 = -1
			} else {
				hx_if_44 = c_4.(int)
			}
			return hx_if_44
		}() != 104 {
			key := hxrt.StdString(self.__hx_this.unserialize())
			map__1.__hx_this.set(key, self.__hx_this.unserialize())
		}
		self.pos = int(int32((self.pos + 1)))
		return map__1
	case 99:
		name_3 := hxrt.StdString(self.__hx_this.unserialize())
		var declaration_3 any = self.__hx_this.resolveClass(name_3)
		if hxrt.AnyEqualsNull(declaration_3) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Class not found "), name_3))
		}
		var object_1 any = Type_createEmptyInstance(declaration_3)
		hx_arr_45 := self.cache
		hx_arr_45.Push(object_1)
		self.__hx_this.unserializeObject(object_1)
		return object_1
	case 100:
		return self.__hx_this.readFloat()
	case 102:
		return false
	case 105:
		return self.__hx_this.readDigits()
	case 106:
		name_4 := hxrt.StdString(self.__hx_this.unserialize())
		var declaration_4 any = self.__hx_this.resolveEnum(name_4)
		if hxrt.AnyEqualsNull(declaration_4) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Enum not found "), name_4))
		}
		self.pos = int(int32((self.pos + 1)))
		index_6 := self.__hx_this.readDigits()
		tag := func(hx_value_46 any) *string {
			if hx_value_46 == nil {
				var hx_zero_47 *string
				return hx_zero_47
			}
			return hx_value_46.(*string)
		}(Type_getEnumConstructs(declaration_4).Get(index_6))
		if hxrt.StringEqualStringPtr(tag, nil) {
			hxrt.Throw(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unknown enum index "), name_4), hxrt.StringFromLiteral("@")), index_6))
		}
		var value any = self.__hx_this.unserializeEnum(declaration_4, tag)
		hx_arr_48 := self.cache
		hx_arr_48.Push(value)
		return value
	case 107:
		return Math_NaN
	case 108:
		list := New_haxe__ds__List()
		hx_arr_49 := self.cache
		hx_arr_49.Push(list)
		for func() int {
			index_7 := self.pos
			s_5 := self.buf
			var c_5 any = hxrt.StringCharCodeAtAnyStringPtr(s_5, index_7)
			var hx_if_50 int
			if c_5 == nil {
				hx_if_50 = -1
			} else {
				hx_if_50 = c_5.(int)
			}
			return hx_if_50
		}() != 104 {
			list.__hx_this.add(self.__hx_this.unserialize())
		}
		self.pos = int(int32((self.pos + 1)))
		return list
	case 109:
		return Math_NEGATIVE_INFINITY
	case 110:
		return nil
	case 111:
		hx_obj_51 := map[string]any{}
		var object_2 any = hx_obj_51
		hx_arr_52 := self.cache
		hx_arr_52.Push(object_2)
		self.__hx_this.unserializeObject(object_2)
		return object_2
	case 112:
		return Math_POSITIVE_INFINITY
	case 113:
		map__2 := New_haxe__ds__IntMap()
		hx_arr_53 := self.cache
		hx_arr_53.Push(map__2)
		hx_post_54 := self.pos
		self.pos = int(int32((self.pos + 1)))
		index_8 := hx_post_54
		s_6 := self.buf
		var c_6 any = hxrt.StringCharCodeAtAnyStringPtr(s_6, index_8)
		var hx_if_55 int
		if c_6 == nil {
			hx_if_55 = -1
		} else {
			hx_if_55 = c_6.(int)
		}
		code_1 := hx_if_55
		for code_1 == 58 {
			map__2.__hx_this.set(self.__hx_this.readDigits(), self.__hx_this.unserialize())
			hx_post_56 := self.pos
			self.pos = int(int32((self.pos + 1)))
			index_9 := hx_post_56
			s_7 := self.buf
			var c_7 any = hxrt.StringCharCodeAtAnyStringPtr(s_7, index_9)
			var hx_if_57 int
			if c_7 == nil {
				hx_if_57 = -1
			} else {
				hx_if_57 = c_7.(int)
			}
			code_1 = hx_if_57
		}
		if code_1 != 104 {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid IntMap format"))
		}
		return map__2
	case 114:
		index_10 := self.__hx_this.readDigits()
		if (index_10 < 0) || (index_10 >= self.cache.Len()) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid reference"))
		}
		return self.cache.Get(index_10)
	case 115:
		bytesLength := self.__hx_this.readDigits()
		if (func() int {
			hx_post_58 := self.pos
			self.pos = int(int32((self.pos + 1)))
			index_11 := hx_post_58
			s_8 := self.buf
			var c_8 any = hxrt.StringCharCodeAtAnyStringPtr(s_8, index_11)
			var hx_if_59 int
			if c_8 == nil {
				hx_if_59 = -1
			} else {
				hx_if_59 = c_8.(int)
			}
			return hx_if_59
		}() != 58) || (int(int32((hxrt.Int32Wrap(self.length) - hxrt.Int32Wrap(self.pos)))) < bytesLength) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid bytes length"))
		}
		bytes := self.__hx_this.decodeBytes(hxrt.StringSubstrStringPtr(self.buf, self.pos, bytesLength, true))
		self.pos = int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(bytesLength))))
		hx_arr_60 := self.cache
		hx_arr_60.Push(bytes)
		return bytes
	case 116:
		return true
	case 118:
		var date *Date
		if self.__hx_this.isLegacyDate() {
			date = Date_fromString(hxrt.StringSubstrStringPtr(self.buf, self.pos, 19, true))
			self.pos = int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(19))))
		} else {
			date = Date_fromTime(self.__hx_this.readFloat())
		}
		hx_arr_61 := self.cache
		hx_arr_61.Push(date)
		return date
	case 119:
		name_5 := hxrt.StdString(self.__hx_this.unserialize())
		var declaration_5 any = self.__hx_this.resolveEnum(name_5)
		if hxrt.AnyEqualsNull(declaration_5) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Enum not found "), name_5))
		}
		var value_1 any = self.__hx_this.unserializeEnum(declaration_5, hxrt.StdString(self.__hx_this.unserialize()))
		hx_arr_62 := self.cache
		hx_arr_62.Push(value_1)
		return value_1
	case 120:
		hxrt.Throw(self.__hx_this.unserialize())
	case 121:
		stringLength := self.__hx_this.readDigits()
		if (func() int {
			hx_post_63 := self.pos
			self.pos = int(int32((self.pos + 1)))
			index_12 := hx_post_63
			s_9 := self.buf
			var c_9 any = hxrt.StringCharCodeAtAnyStringPtr(s_9, index_12)
			var hx_if_64 int
			if c_9 == nil {
				hx_if_64 = -1
			} else {
				hx_if_64 = c_9.(int)
			}
			return hx_if_64
		}() != 58) || (int(int32((hxrt.Int32Wrap(self.length) - hxrt.Int32Wrap(self.pos)))) < stringLength) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid string length"))
		}
		value_2 := StringTools_urlDecode(hxrt.StringSubstrStringPtr(self.buf, self.pos, stringLength, true))
		self.pos = int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(stringLength))))
		hx_arr_65 := self.scache
		hx_arr_65.Push(value_2)
		return value_2
	case 122:
		return 0
	default:
	}
	self.pos = int(int32((self.pos - 1)))
	hxrt.Throw(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid char "), hxrt.StringCharAtStringPtr(self.buf, self.pos)), hxrt.StringFromLiteral(" at position ")), self.pos))
	var hx_throw_zero_66 any
	return hx_throw_zero_66
}

func (self *haxe__Unserializer) isLegacyDate() bool {
	index := self.pos
	s := self.buf
	var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
	var hx_if_67 int
	if c == nil {
		hx_if_67 = -1
	} else {
		hx_if_67 = c.(int)
	}
	code := hx_if_67
	if (code < 48) || (code > 57) {
		return false
	}
	index_1 := int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(1))))
	s_1 := self.buf
	var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(s_1, index_1)
	var hx_if_68 int
	if c_1 == nil {
		hx_if_68 = -1
	} else {
		hx_if_68 = c_1.(int)
	}
	code_1 := hx_if_68
	if (code_1 < 48) || (code_1 > 57) {
		return false
	}
	index_2 := int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(2))))
	s_2 := self.buf
	var c_2 any = hxrt.StringCharCodeAtAnyStringPtr(s_2, index_2)
	var hx_if_69 int
	if c_2 == nil {
		hx_if_69 = -1
	} else {
		hx_if_69 = c_2.(int)
	}
	code_2 := hx_if_69
	if (code_2 < 48) || (code_2 > 57) {
		return false
	}
	index_3 := int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(3))))
	s_3 := self.buf
	var c_3 any = hxrt.StringCharCodeAtAnyStringPtr(s_3, index_3)
	var hx_if_70 int
	if c_3 == nil {
		hx_if_70 = -1
	} else {
		hx_if_70 = c_3.(int)
	}
	code_3 := hx_if_70
	if (code_3 < 48) || (code_3 > 57) {
		return false
	}
	return (func() int {
		index_4 := int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(4))))
		s_4 := self.buf
		var c_4 any = hxrt.StringCharCodeAtAnyStringPtr(s_4, index_4)
		var hx_if_71 int
		if c_4 == nil {
			hx_if_71 = -1
		} else {
			hx_if_71 = c_4.(int)
		}
		return hx_if_71
	}() == 45)
}

var haxe__Unserializer_BASE64 *string = hxrt.StringFromLiteral("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789%:")

var haxe__Unserializer_DEFAULT_RESOLVER map[string]any = func() map[string]any {
	hx_obj_72 := map[string]any{}
	hx_obj_72["resolveClass"] = func(name *string) any {
		return Type_resolveClass(name)
	}
	hx_obj_72["resolveEnum"] = func(name *string) any {
		return Type_resolveEnum(name)
	}
	return hx_obj_72
}()

var haxe__Unserializer_NULL_RESOLVER map[string]any = func() map[string]any {
	hx_obj_73 := map[string]any{}
	hx_obj_73["resolveClass"] = func(_name *string) any {
		return nil
	}
	hx_obj_73["resolveEnum"] = func(_name *string) any {
		return nil
	}
	return hx_obj_73
}()

func haxe__Unserializer_run(value *string) any {
	return New_haxe__Unserializer(value).__hx_this.unserialize()
}
