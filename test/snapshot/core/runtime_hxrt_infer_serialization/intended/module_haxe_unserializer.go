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
		hx_obj_160 := map[string]any{}
		hx_obj_160["resolveClass"] = func(name *string) any {
			return Type_resolveClass(name)
		}
		hx_obj_160["resolveEnum"] = func(name *string) any {
			return Type_resolveEnum(name)
		}
		current = hx_obj_160
		haxe__Unserializer_DEFAULT_RESOLVER = current
	}
	self.resolver = current
	return self
}

func (self *haxe__Unserializer) setResolver(value any) {
	var hx_if_161 any
	if hxrt.AnyEqualsNull(value) {
		hx_if_161 = haxe__Unserializer_NULL_RESOLVER
	} else {
		hx_if_161 = value
	}
	self.resolver = hx_if_161
}

func (self *haxe__Unserializer) getResolver() any {
	return self.resolver
}

func (self *haxe__Unserializer) get(index int) int {
	s := self.buf
	var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
	var hx_if_162 int
	if c == nil {
		hx_if_162 = -1
	} else {
		hx_if_162 = c.(int)
	}
	return hx_if_162
}

func (self *haxe__Unserializer) readDigits() int {
	value := 0
	negative := false
	start := self.pos
	for true {
		index := self.pos
		s := self.buf
		var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
		var hx_if_163 int
		if c == nil {
			hx_if_163 = -1
		} else {
			hx_if_163 = c.(int)
		}
		code := hx_if_163
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
	var hx_if_164 int
	if negative {
		hx_if_164 = int(int32(-int32(value)))
	} else {
		hx_if_164 = value
	}
	return hx_if_164
}

func (self *haxe__Unserializer) readFloat() float64 {
	start := self.pos
	for true {
		index := self.pos
		s := self.buf
		var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
		var hx_if_165 int
		if c == nil {
			hx_if_165 = -1
		} else {
			hx_if_165 = c.(int)
		}
		code := hx_if_165
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
			var hx_if_166 int
			if c == nil {
				hx_if_166 = -1
			} else {
				hx_if_166 = c.(int)
			}
			return hx_if_166
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
		Reflect_setField(target, hxrt.StdString(func(hx_value_167 any) *string {
			if hx_value_167 == nil {
				var hx_zero_168 *string
				return hx_zero_168
			}
			return hx_value_167.(*string)
		}(key)), self.__hx_this.unserialize())
	}
	self.pos = int(int32((self.pos + 1)))
}

func (self *haxe__Unserializer) resolveClass(name *string) any {
	var method any = Reflect_field(self.resolver, hxrt.StringFromLiteral("resolveClass"))
	var hx_if_169 any
	if hxrt.AnyEqualsNull(method) {
		hx_if_169 = nil
	} else {
		hx_if_169 = Reflect_callMethod(self.resolver, method, hxrt.NewArray(name))
	}
	return hx_if_169
}

func (self *haxe__Unserializer) resolveEnum(name *string) any {
	var method any = Reflect_field(self.resolver, hxrt.StringFromLiteral("resolveEnum"))
	var hx_if_170 any
	if hxrt.AnyEqualsNull(method) {
		hx_if_170 = nil
	} else {
		hx_if_170 = Reflect_callMethod(self.resolver, method, hxrt.NewArray(name))
	}
	return hx_if_170
}

func (self *haxe__Unserializer) unserializeEnum(declaration any, tag *string) any {
	if func() int {
		hx_post_171 := self.pos
		self.pos = int(int32((self.pos + 1)))
		index := hx_post_171
		s := self.buf
		var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
		var hx_if_172 int
		if c == nil {
			hx_if_172 = -1
		} else {
			hx_if_172 = c.(int)
		}
		return hx_if_172
	}() != 58 {
		hxrt.Throw(hxrt.StringFromLiteral("Invalid enum format"))
	}
	count := self.__hx_this.readDigits()
	if count == 0 {
		return Type_createEnum(declaration, tag, hxrt.NewArray())
	}
	arguments := hxrt.NewArray()
	for func() int {
		hx_post_173 := count
		count = int(int32((count - 1)))
		return hx_post_173
	}() > 0 {
		arguments.Push(self.__hx_this.unserialize())
	}
	return Type_createEnum(declaration, tag, arguments)
}

func (self *haxe__Unserializer) decodeBytes(encoded *string) *haxe__io__Bytes {
	rest := int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(encoded)) & hxrt.Int32Wrap(3))))
	size := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(encoded)) >> uint(2))))) * hxrt.Int32Wrap(3))))) + hxrt.Int32Wrap(func() int {
		var hx_if_175 int
		if rest >= 2 {
			hx_if_175 = int(int32((hxrt.Int32Wrap(rest) - hxrt.Int32Wrap(1))))
		} else {
			hx_if_175 = 0
		}
		return hx_if_175
	}()))))
	bytes := haxe__io__Bytes_alloc(size)
	index := 0
	output := 0
	complete := int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(encoded)) - hxrt.Int32Wrap(rest))))
	for index < complete {
		first := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_176 := index
			index = int(int32((index + 1)))
			return hx_post_176
		}())))
		second := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_177 := index
			index = int(int32((index + 1)))
			return hx_post_177
		}())))
		third := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_178 := index
			index = int(int32((index + 1)))
			return hx_post_178
		}())))
		fourth := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_179 := index
			index = int(int32((index + 1)))
			return hx_post_179
		}())))
		hx_post_180 := output
		output = int(int32((output + 1)))
		pos := hx_post_180
		bytes.b[pos] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(first) << uint(2))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second) >> uint(4))))))))) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
		hx_post_181 := output
		output = int(int32((output + 1)))
		pos_1 := hx_post_181
		bytes.b[pos_1] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second) << uint(4))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(third) >> uint(2))))))))) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
		hx_post_182 := output
		output = int(int32((output + 1)))
		pos_2 := hx_post_182
		bytes.b[pos_2] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(third) << uint(6))))) | hxrt.Int32Wrap(fourth))))) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	if rest >= 2 {
		first_1 := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_183 := index
			index = int(int32((index + 1)))
			return hx_post_183
		}())))
		second_1 := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
			hx_post_184 := index
			index = int(int32((index + 1)))
			return hx_post_184
		}())))
		hx_post_185 := output
		output = int(int32((output + 1)))
		pos_3 := hx_post_185
		bytes.b[pos_3] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(first_1) << uint(2))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second_1) >> uint(4))))))))) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
		if rest == 3 {
			third_1 := self.__hx_this.base64Value(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(encoded, func() int {
				hx_post_186 := index
				index = int(int32((index + 1)))
				return hx_post_186
			}())))
			hx_post_187 := output
			output = int(int32((output + 1)))
			pos_4 := hx_post_187
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
		hx_post_188 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_188
		if hxrt.StringCharCodeAtAnyStringPtr(haxe__Unserializer_BASE64, index) == code {
			return index
		}
	}
	return -1
}

func (self *haxe__Unserializer) unserialize() any {
	hx_post_189 := self.pos
	self.pos = int(int32((self.pos + 1)))
	index := hx_post_189
	s := self.buf
	var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
	var hx_if_190 int
	if c == nil {
		hx_if_190 = -1
	} else {
		hx_if_190 = c.(int)
	}
	_g := hx_if_190
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
		hx_arr_191 := self.cache
		hx_arr_191.Push(object)
		var hook any = Reflect_field(object, hxrt.StringFromLiteral("hxUnserialize"))
		if hxrt.AnyEqualsNull(hook) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Class "), name_2), hxrt.StringFromLiteral(" has no hxUnserialize hook")))
		}
		Reflect_callMethod(object, hook, hxrt.NewArray(self))
		if func() int {
			hx_post_192 := self.pos
			self.pos = int(int32((self.pos + 1)))
			index_1 := hx_post_192
			s_1 := self.buf
			var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(s_1, index_1)
			var hx_if_193 int
			if c_1 == nil {
				hx_if_193 = -1
			} else {
				hx_if_193 = c_1.(int)
			}
			return hx_if_193
		}() != 103 {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid custom data"))
		}
		return object
	case 77:
		map_ := New_haxe__ds__ObjectMap()
		hx_arr_194 := self.cache
		hx_arr_194.Push(map_)
		for func() int {
			index_2 := self.pos
			s_2 := self.buf
			var c_2 any = hxrt.StringCharCodeAtAnyStringPtr(s_2, index_2)
			var hx_if_195 int
			if c_2 == nil {
				hx_if_195 = -1
			} else {
				hx_if_195 = c_2.(int)
			}
			return hx_if_195
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
		hx_arr_196 := self.cache
		hx_arr_196.Push(array)
	hx_loop_198:
		for true {
			index_4 := self.pos
			s_3 := self.buf
			var c_3 any = hxrt.StringCharCodeAtAnyStringPtr(s_3, index_4)
			var hx_if_197 int
			if c_3 == nil {
				hx_if_197 = -1
			} else {
				hx_if_197 = c_3.(int)
			}
			code := hx_if_197
			if code == 104 {
				self.pos = int(int32((self.pos + 1)))
				break hx_loop_198
			}
			if code == 117 {
				self.pos = int(int32((self.pos + 1)))
				count := self.__hx_this.readDigits()
				hx_array_target_199 := array
				hx_array_index_200 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(array.Len()) + hxrt.Int32Wrap(count))))) - hxrt.Int32Wrap(1))))
				hx_array_target_199.Set(hx_array_index_200, nil)
			} else {
				array.Push(self.__hx_this.unserialize())
			}
		}
		return array
	case 98:
		map__1 := New_haxe__ds__StringMap()
		hx_arr_202 := self.cache
		hx_arr_202.Push(map__1)
		for func() int {
			index_5 := self.pos
			s_4 := self.buf
			var c_4 any = hxrt.StringCharCodeAtAnyStringPtr(s_4, index_5)
			var hx_if_203 int
			if c_4 == nil {
				hx_if_203 = -1
			} else {
				hx_if_203 = c_4.(int)
			}
			return hx_if_203
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
		hx_arr_204 := self.cache
		hx_arr_204.Push(object_1)
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
		tag := func(hx_value_205 any) *string {
			if hx_value_205 == nil {
				var hx_zero_206 *string
				return hx_zero_206
			}
			return hx_value_205.(*string)
		}(Type_getEnumConstructs(declaration_4).Get(index_6))
		if hxrt.StringEqualStringPtr(tag, nil) {
			hxrt.Throw(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unknown enum index "), name_4), hxrt.StringFromLiteral("@")), index_6))
		}
		var value any = self.__hx_this.unserializeEnum(declaration_4, tag)
		hx_arr_207 := self.cache
		hx_arr_207.Push(value)
		return value
	case 107:
		return Math_NaN
	case 108:
		list := New_haxe__ds__List()
		hx_arr_208 := self.cache
		hx_arr_208.Push(list)
		for func() int {
			index_7 := self.pos
			s_5 := self.buf
			var c_5 any = hxrt.StringCharCodeAtAnyStringPtr(s_5, index_7)
			var hx_if_209 int
			if c_5 == nil {
				hx_if_209 = -1
			} else {
				hx_if_209 = c_5.(int)
			}
			return hx_if_209
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
		hx_obj_210 := map[string]any{}
		var object_2 any = hx_obj_210
		hx_arr_211 := self.cache
		hx_arr_211.Push(object_2)
		self.__hx_this.unserializeObject(object_2)
		return object_2
	case 112:
		return Math_POSITIVE_INFINITY
	case 113:
		map__2 := New_haxe__ds__IntMap()
		hx_arr_212 := self.cache
		hx_arr_212.Push(map__2)
		hx_post_213 := self.pos
		self.pos = int(int32((self.pos + 1)))
		index_8 := hx_post_213
		s_6 := self.buf
		var c_6 any = hxrt.StringCharCodeAtAnyStringPtr(s_6, index_8)
		var hx_if_214 int
		if c_6 == nil {
			hx_if_214 = -1
		} else {
			hx_if_214 = c_6.(int)
		}
		code_1 := hx_if_214
		for code_1 == 58 {
			map__2.__hx_this.set(self.__hx_this.readDigits(), self.__hx_this.unserialize())
			hx_post_215 := self.pos
			self.pos = int(int32((self.pos + 1)))
			index_9 := hx_post_215
			s_7 := self.buf
			var c_7 any = hxrt.StringCharCodeAtAnyStringPtr(s_7, index_9)
			var hx_if_216 int
			if c_7 == nil {
				hx_if_216 = -1
			} else {
				hx_if_216 = c_7.(int)
			}
			code_1 = hx_if_216
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
			hx_post_217 := self.pos
			self.pos = int(int32((self.pos + 1)))
			index_11 := hx_post_217
			s_8 := self.buf
			var c_8 any = hxrt.StringCharCodeAtAnyStringPtr(s_8, index_11)
			var hx_if_218 int
			if c_8 == nil {
				hx_if_218 = -1
			} else {
				hx_if_218 = c_8.(int)
			}
			return hx_if_218
		}() != 58) || (int(int32((hxrt.Int32Wrap(self.length) - hxrt.Int32Wrap(self.pos)))) < bytesLength) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid bytes length"))
		}
		bytes := self.__hx_this.decodeBytes(hxrt.StringSubstrStringPtr(self.buf, self.pos, bytesLength, true))
		self.pos = int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(bytesLength))))
		hx_arr_219 := self.cache
		hx_arr_219.Push(bytes)
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
		hx_arr_220 := self.cache
		hx_arr_220.Push(date)
		return date
	case 119:
		name_5 := hxrt.StdString(self.__hx_this.unserialize())
		var declaration_5 any = self.__hx_this.resolveEnum(name_5)
		if hxrt.AnyEqualsNull(declaration_5) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Enum not found "), name_5))
		}
		var value_1 any = self.__hx_this.unserializeEnum(declaration_5, hxrt.StdString(self.__hx_this.unserialize()))
		hx_arr_221 := self.cache
		hx_arr_221.Push(value_1)
		return value_1
	case 120:
		hxrt.Throw(self.__hx_this.unserialize())
	case 121:
		stringLength := self.__hx_this.readDigits()
		if (func() int {
			hx_post_222 := self.pos
			self.pos = int(int32((self.pos + 1)))
			index_12 := hx_post_222
			s_9 := self.buf
			var c_9 any = hxrt.StringCharCodeAtAnyStringPtr(s_9, index_12)
			var hx_if_223 int
			if c_9 == nil {
				hx_if_223 = -1
			} else {
				hx_if_223 = c_9.(int)
			}
			return hx_if_223
		}() != 58) || (int(int32((hxrt.Int32Wrap(self.length) - hxrt.Int32Wrap(self.pos)))) < stringLength) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid string length"))
		}
		value_2 := StringTools_urlDecode(hxrt.StringSubstrStringPtr(self.buf, self.pos, stringLength, true))
		self.pos = int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(stringLength))))
		hx_arr_224 := self.scache
		hx_arr_224.Push(value_2)
		return value_2
	case 122:
		return 0
	default:
	}
	self.pos = int(int32((self.pos - 1)))
	hxrt.Throw(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid char "), hxrt.StringCharAtStringPtr(self.buf, self.pos)), hxrt.StringFromLiteral(" at position ")), self.pos))
	var hx_throw_zero_225 any
	return hx_throw_zero_225
}

func (self *haxe__Unserializer) isLegacyDate() bool {
	index := self.pos
	s := self.buf
	var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
	var hx_if_226 int
	if c == nil {
		hx_if_226 = -1
	} else {
		hx_if_226 = c.(int)
	}
	code := hx_if_226
	if (code < 48) || (code > 57) {
		return false
	}
	index_1 := int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(1))))
	s_1 := self.buf
	var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(s_1, index_1)
	var hx_if_227 int
	if c_1 == nil {
		hx_if_227 = -1
	} else {
		hx_if_227 = c_1.(int)
	}
	code_1 := hx_if_227
	if (code_1 < 48) || (code_1 > 57) {
		return false
	}
	index_2 := int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(2))))
	s_2 := self.buf
	var c_2 any = hxrt.StringCharCodeAtAnyStringPtr(s_2, index_2)
	var hx_if_228 int
	if c_2 == nil {
		hx_if_228 = -1
	} else {
		hx_if_228 = c_2.(int)
	}
	code_2 := hx_if_228
	if (code_2 < 48) || (code_2 > 57) {
		return false
	}
	index_3 := int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(3))))
	s_3 := self.buf
	var c_3 any = hxrt.StringCharCodeAtAnyStringPtr(s_3, index_3)
	var hx_if_229 int
	if c_3 == nil {
		hx_if_229 = -1
	} else {
		hx_if_229 = c_3.(int)
	}
	code_3 := hx_if_229
	if (code_3 < 48) || (code_3 > 57) {
		return false
	}
	return (func() int {
		index_4 := int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(4))))
		s_4 := self.buf
		var c_4 any = hxrt.StringCharCodeAtAnyStringPtr(s_4, index_4)
		var hx_if_230 int
		if c_4 == nil {
			hx_if_230 = -1
		} else {
			hx_if_230 = c_4.(int)
		}
		return hx_if_230
	}() == 45)
}

var haxe__Unserializer_BASE64 *string = hxrt.StringFromLiteral("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789%:")

var haxe__Unserializer_DEFAULT_RESOLVER map[string]any = func() map[string]any {
	hx_obj_231 := map[string]any{}
	hx_obj_231["resolveClass"] = func(name *string) any {
		return Type_resolveClass(name)
	}
	hx_obj_231["resolveEnum"] = func(name *string) any {
		return Type_resolveEnum(name)
	}
	return hx_obj_231
}()

var haxe__Unserializer_NULL_RESOLVER map[string]any = func() map[string]any {
	hx_obj_232 := map[string]any{}
	hx_obj_232["resolveClass"] = func(_name *string) any {
		return nil
	}
	hx_obj_232["resolveEnum"] = func(_name *string) any {
		return nil
	}
	return hx_obj_232
}()

func haxe__Unserializer_run(value *string) any {
	return New_haxe__Unserializer(value).__hx_this.unserialize()
}
