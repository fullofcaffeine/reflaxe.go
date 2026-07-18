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

func New_haxe__Serializer() *haxe__Serializer {
	self := &haxe__Serializer{}
	self.__hx_this = self
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
	var known any = func(hx_value_159 any) any {
		if hx_value_159 == nil {
			return nil
		}
		return hx_value_159.(int)
	}(self.shash.get(value))
	if known != nil {
		_this := self.buf
		_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("R"))
		_this_1 := self.buf
		_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StdString(known.(int)))
		return
	}
	self.shash.set(value, func() int {
		hx_post_160 := self.scount
		self.scount = int(int32((self.scount + 1)))
		return hx_post_160
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
		hx_post_161 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_161
		if hxrt.HaxeEqual(self.cache.Get(index), value) {
			_this := self.buf
			_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("r"))
			_this_1 := self.buf
			_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StdString(index))
			return true
		}
	}
	hx_arr_162 := self.cache
	hx_arr_162.Push(value)
	return false
}

func (self *haxe__Serializer) serializeFields(value any) {
	fields := hxrt.SerializationFields(value)
	_g := 0
	_g1 := len(fields)
	for _g < _g1 {
		hx_post_163 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_163
		field := fields[index]
		self.serializeString(field.Name)
		self.serialize(field.Value)
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
		self.serialize(item)
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
	v := (float64(int(int32((hxrt.Int32Wrap(value.length) * hxrt.Int32Wrap(8))))) / float64(6))
	x := hxrt.MathCeilInt(v)
	_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StdString(x))
	_this_2 := self.buf
	_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StringFromLiteral(":"))
	index := 0
	max := int(int32((hxrt.Int32Wrap(value.length) - hxrt.Int32Wrap(2))))
	for index < max {
		hx_post_164 := index
		index = int(int32((index + 1)))
		pos := hx_post_164
		first := value.b[pos]
		hx_post_165 := index
		index = int(int32((index + 1)))
		pos_1 := hx_post_165
		second := value.b[pos_1]
		hx_post_166 := index
		index = int(int32((index + 1)))
		pos_2 := hx_post_166
		third := value.b[pos_2]
		_this_3 := self.buf
		c := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int(int32((hxrt.Int32Wrap(first) >> uint(2))))))
		_this_3.b = hxrt.StringConcatStringPtr(_this_3.b, hxrt.StringFromCharCode(c))
		_this_4 := self.buf
		c_1 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(first) << uint(4))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second) >> uint(4))))))))) & hxrt.Int32Wrap(63))))))
		_this_4.b = hxrt.StringConcatStringPtr(_this_4.b, hxrt.StringFromCharCode(c_1))
		_this_5 := self.buf
		c_2 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second) << uint(2))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(third) >> uint(6))))))))) & hxrt.Int32Wrap(63))))))
		_this_5.b = hxrt.StringConcatStringPtr(_this_5.b, hxrt.StringFromCharCode(c_2))
		_this_6 := self.buf
		c_3 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int(int32((hxrt.Int32Wrap(third) & hxrt.Int32Wrap(63))))))
		_this_6.b = hxrt.StringConcatStringPtr(_this_6.b, hxrt.StringFromCharCode(c_3))
	}
	if index == max {
		hx_post_167 := index
		index = int(int32((index + 1)))
		pos_3 := hx_post_167
		first_1 := value.b[pos_3]
		hx_post_168 := index
		index = int(int32((index + 1)))
		pos_4 := hx_post_168
		second_1 := value.b[pos_4]
		_this_7 := self.buf
		c_4 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int(int32((hxrt.Int32Wrap(first_1) >> uint(2))))))
		_this_7.b = hxrt.StringConcatStringPtr(_this_7.b, hxrt.StringFromCharCode(c_4))
		_this_8 := self.buf
		c_5 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(first_1) << uint(4))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second_1) >> uint(4))))))))) & hxrt.Int32Wrap(63))))))
		_this_8.b = hxrt.StringConcatStringPtr(_this_8.b, hxrt.StringFromCharCode(c_5))
		_this_9 := self.buf
		c_6 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(second_1) << uint(2))))) & hxrt.Int32Wrap(63))))))
		_this_9.b = hxrt.StringConcatStringPtr(_this_9.b, hxrt.StringFromCharCode(c_6))
	} else {
		if index == int(int32((hxrt.Int32Wrap(max) + hxrt.Int32Wrap(1)))) {
			hx_post_169 := index
			index = int(int32((index + 1)))
			pos_5 := hx_post_169
			first_2 := value.b[pos_5]
			_this_10 := self.buf
			c_7 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int(int32((hxrt.Int32Wrap(first_2) >> uint(2))))))
			_this_10.b = hxrt.StringConcatStringPtr(_this_10.b, hxrt.StringFromCharCode(c_7))
			_this_11 := self.buf
			c_8 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(haxe__Serializer_BASE64, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(first_2) << uint(4))))) & hxrt.Int32Wrap(63))))))
			_this_11.b = hxrt.StringConcatStringPtr(_this_11.b, hxrt.StringFromCharCode(c_8))
		}
	}
}

func (self *haxe__Serializer) serializeClass(value any, declaration any) {
	className := Type_getClassName(declaration)
	if hxrt.StringEqualStringPtr(className, hxrt.StringFromLiteral("String")) {
		self.serializeString(hxrt.StdString(value))
		return
	}
	if self.useCache && self.serializeRef(value) {
		return
	}
	switch *hxrt.StdString(className) {
	case *hxrt.StdString(hxrt.StringFromLiteral("Array")):
		self.serializeArray(func(hx_value_170 any) *hxrt.Array {
			if hx_value_170 == nil {
				var hx_zero_171 *hxrt.Array
				return hx_zero_171
			}
			return hx_value_170.(*hxrt.Array)
		}(value))
	case *hxrt.StdString(hxrt.StringFromLiteral("Date")):
		date := func(hx_value_172 any) *Date {
			if hx_value_172 == nil {
				var hx_zero_173 *Date
				return hx_zero_173
			}
			return hx_value_172.(*Date)
		}(value)
		_this := self.buf
		_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("v"))
		_this_1 := self.buf
		x := date.ms
		_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StdString(x))
	case *hxrt.StdString(hxrt.StringFromLiteral("haxe.ds.IntMap")):
		_this_2 := self.buf
		_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StringFromLiteral("q"))
		map_ := func(hx_value_174 any) *haxe__ds__IntMap {
			if hx_value_174 == nil {
				var hx_zero_175 *haxe__ds__IntMap
				return hx_zero_175
			}
			return hx_value_174.(*haxe__ds__IntMap)
		}(value)
		key := func(hx_value_176 any) map[string]any {
			if hx_value_176 == nil {
				var hx_zero_177 map[string]any
				return hx_zero_177
			}
			return hx_value_176.(map[string]any)
		}(map_.keys())
		for func(hx_obj_178 map[string]any) func() bool {
			hx_field_179 := hx_obj_178["hasNext"]
			if hx_field_179 == nil {
				var hx_zero_180 func() bool
				return hx_zero_180
			}
			return hx_field_179.(func() bool)
		}(key)() {
			key_1 := func(hx_obj_181 map[string]any) func() int {
				hx_field_182 := hx_obj_181["next"]
				if hx_field_182 == nil {
					var hx_zero_183 func() int
					return hx_zero_183
				}
				return hx_field_182.(func() int)
			}(key)()
			_this_3 := self.buf
			_this_3.b = hxrt.StringConcatStringPtr(_this_3.b, hxrt.StringFromLiteral(":"))
			_this_4 := self.buf
			_this_4.b = hxrt.StringConcatStringPtr(_this_4.b, hxrt.StdString(key_1))
			self.serialize(map_.get(key_1))
		}
		_this_5 := self.buf
		_this_5.b = hxrt.StringConcatStringPtr(_this_5.b, hxrt.StringFromLiteral("h"))
	case *hxrt.StdString(hxrt.StringFromLiteral("haxe.ds.List")):
		_this_6 := self.buf
		_this_6.b = hxrt.StringConcatStringPtr(_this_6.b, hxrt.StringFromLiteral("l"))
		list := func(hx_value_184 any) *haxe__ds__List {
			if hx_value_184 == nil {
				var hx_zero_185 *haxe__ds__List
				return hx_zero_185
			}
			return hx_value_184.(*haxe__ds__List)
		}(value)
		item := func(hx_value_186 any) *haxe__ds___List__GoListIterator {
			if hx_value_186 == nil {
				var hx_zero_187 *haxe__ds___List__GoListIterator
				return hx_zero_187
			}
			return hx_value_186.(*haxe__ds___List__GoListIterator)
		}(list.iterator())
		for func(hx_value_188 any) bool {
			if hx_value_188 == nil {
				var hx_zero_189 bool
				return hx_zero_189
			}
			return hx_value_188.(bool)
		}(item.hasNext()) {
			var item_1 any = item.next()
			self.serialize(item_1)
		}
		_this_7 := self.buf
		_this_7.b = hxrt.StringConcatStringPtr(_this_7.b, hxrt.StringFromLiteral("h"))
	case *hxrt.StdString(hxrt.StringFromLiteral("haxe.ds.ObjectMap")):
		_this_8 := self.buf
		_this_8.b = hxrt.StringConcatStringPtr(_this_8.b, hxrt.StringFromLiteral("M"))
		map__1 := func(hx_value_190 any) *haxe__ds__ObjectMap {
			if hx_value_190 == nil {
				var hx_zero_191 *haxe__ds__ObjectMap
				return hx_zero_191
			}
			return hx_value_190.(*haxe__ds__ObjectMap)
		}(value)
		key_2 := func(hx_value_192 any) map[string]any {
			if hx_value_192 == nil {
				var hx_zero_193 map[string]any
				return hx_zero_193
			}
			return hx_value_192.(map[string]any)
		}(map__1.keys())
		for func(hx_obj_194 map[string]any) func() bool {
			hx_field_195 := hx_obj_194["hasNext"]
			if hx_field_195 == nil {
				var hx_zero_196 func() bool
				return hx_zero_196
			}
			return hx_field_195.(func() bool)
		}(key_2)() {
			var key_3 any = func(hx_obj_197 map[string]any) func() any {
				hx_field_198 := hx_obj_197["next"]
				if hx_field_198 == nil {
					var hx_zero_199 func() any
					return hx_zero_199
				}
				return hx_field_198.(func() any)
			}(key_2)()
			self.serialize(key_3)
			self.serialize(map__1.get(key_3))
		}
		_this_9 := self.buf
		_this_9.b = hxrt.StringConcatStringPtr(_this_9.b, hxrt.StringFromLiteral("h"))
	case *hxrt.StdString(hxrt.StringFromLiteral("haxe.ds.StringMap")):
		_this_10 := self.buf
		_this_10.b = hxrt.StringConcatStringPtr(_this_10.b, hxrt.StringFromLiteral("b"))
		map__2 := func(hx_value_200 any) *haxe__ds__StringMap {
			if hx_value_200 == nil {
				var hx_zero_201 *haxe__ds__StringMap
				return hx_zero_201
			}
			return hx_value_200.(*haxe__ds__StringMap)
		}(value)
		key_4 := func(hx_value_202 any) map[string]any {
			if hx_value_202 == nil {
				var hx_zero_203 map[string]any
				return hx_zero_203
			}
			return hx_value_202.(map[string]any)
		}(map__2.keys())
		for func(hx_obj_204 map[string]any) func() bool {
			hx_field_205 := hx_obj_204["hasNext"]
			if hx_field_205 == nil {
				var hx_zero_206 func() bool
				return hx_zero_206
			}
			return hx_field_205.(func() bool)
		}(key_4)() {
			key_5 := func(hx_obj_207 map[string]any) func() *string {
				hx_field_208 := hx_obj_207["next"]
				if hx_field_208 == nil {
					var hx_zero_209 func() *string
					return hx_zero_209
				}
				return hx_field_208.(func() *string)
			}(key_4)()
			self.serializeString(key_5)
			self.serialize(map__2.get(key_5))
		}
		_this_11 := self.buf
		_this_11.b = hxrt.StringConcatStringPtr(_this_11.b, hxrt.StringFromLiteral("h"))
	case *hxrt.StdString(hxrt.StringFromLiteral("haxe.io.Bytes")):
		self.serializeBytes(func(hx_value_210 any) *haxe__io__Bytes {
			if hx_value_210 == nil {
				var hx_zero_211 *haxe__io__Bytes
				return hx_zero_211
			}
			return hx_value_210.(*haxe__io__Bytes)
		}(value))
	default:
		if self.useCache {
			hx_arr_212 := self.cache
			hx_arr_212.Pop()
		}
		if haxe__GoSerializationBridge_hasSerializeHook(value) {
			_this_12 := self.buf
			_this_12.b = hxrt.StringConcatStringPtr(_this_12.b, hxrt.StringFromLiteral("C"))
			self.serializeString(className)
			if self.useCache {
				hx_arr_213 := self.cache
				hx_arr_213.Push(value)
			}
			haxe__GoSerializationBridge_callSerializeHook(value, self)
			_this_13 := self.buf
			_this_13.b = hxrt.StringConcatStringPtr(_this_13.b, hxrt.StringFromLiteral("g"))
		} else {
			_this_14 := self.buf
			_this_14.b = hxrt.StringConcatStringPtr(_this_14.b, hxrt.StringFromLiteral("c"))
			self.serializeString(className)
			if self.useCache {
				hx_arr_214 := self.cache
				hx_arr_214.Push(value)
			}
			self.serializeFields(value)
		}
	}
}

func (self *haxe__Serializer) serializeEnum(value any, declaration any) {
	if self.useCache {
		if self.serializeRef(value) {
			return
		}
		hx_arr_215 := self.cache
		hx_arr_215.Pop()
	}
	_this := self.buf
	var hx_if_216 *string
	if self.useEnumIndex {
		hx_if_216 = hxrt.StringFromLiteral("j")
	} else {
		hx_if_216 = hxrt.StringFromLiteral("w")
	}
	x := hx_if_216
	_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StdString(x))
	self.serializeString(Type_getEnumName(declaration))
	if self.useEnumIndex {
		_this_1 := self.buf
		_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StringFromLiteral(":"))
		_this_2 := self.buf
		x_1 := Type_enumIndex(value)
		_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StdString(x_1))
	} else {
		self.serializeString(Type_enumConstructor(value))
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
		self.serialize(parameter)
	}
	if self.useCache {
		hx_arr_217 := self.cache
		hx_arr_217.Push(value)
	}
}

func (self *haxe__Serializer) serialize(value any) {
	if func(hx_value any) bool {
		switch hx_value.(type) {
		case *haxe__io__Bytes:
			return true
		default:
			return false
		}
	}(any(value)) {
		if self.useCache && self.serializeRef(value) {
			return
		}
		self.serializeBytes(func(hx_value_218 any) *haxe__io__Bytes {
			if hx_value_218 == nil {
				var hx_zero_219 *haxe__io__Bytes
				return hx_zero_219
			}
			return hx_value_218.(*haxe__io__Bytes)
		}(value))
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
		number := func(hx_value_220 any) float64 {
			if hx_value_220 == nil {
				var hx_zero_221 float64
				return hx_zero_221
			}
			return hx_value_220.(float64)
		}(value)
		if math.IsNaN(number) {
			_this_4 := self.buf
			_this_4.b = hxrt.StringConcatStringPtr(_this_4.b, hxrt.StringFromLiteral("k"))
		} else {
			if !(!math.IsInf(number, 0) && !math.IsNaN(number)) {
				_this_5 := self.buf
				var hx_if_222 *string
				if number < 0 {
					hx_if_222 = hxrt.StringFromLiteral("m")
				} else {
					hx_if_222 = hxrt.StringFromLiteral("p")
				}
				_this_5.b = hxrt.StringConcatStringPtr(_this_5.b, hx_if_222)
			} else {
				_this_6 := self.buf
				_this_6.b = hxrt.StringConcatStringPtr(_this_6.b, hxrt.StringFromLiteral("d"))
				_this_7 := self.buf
				_this_7.b = hxrt.StringConcatStringPtr(_this_7.b, hxrt.StdString(number))
			}
		}
	case 3:
		boolean := func(hx_value_223 any) bool {
			if hx_value_223 == nil {
				var hx_zero_224 bool
				return hx_zero_224
			}
			return hx_value_223.(bool)
		}(value)
		_this_8 := self.buf
		var hx_if_225 *string
		if boolean {
			hx_if_225 = hxrt.StringFromLiteral("t")
		} else {
			hx_if_225 = hxrt.StringFromLiteral("f")
		}
		_this_8.b = hxrt.StringConcatStringPtr(_this_8.b, hx_if_225)
	case 4:
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *hxrt__TypeClassValue:
				return true
			default:
				return false
			}
		}(any(value)) {
			_this_9 := self.buf
			_this_9.b = hxrt.StringConcatStringPtr(_this_9.b, hxrt.StringFromLiteral("A"))
			self.serializeString(Type_getClassName(value))
		} else {
			if func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt__TypeEnumValue:
					return true
				default:
					return false
				}
			}(any(value)) {
				_this_10 := self.buf
				_this_10.b = hxrt.StringConcatStringPtr(_this_10.b, hxrt.StringFromLiteral("B"))
				self.serializeString(Type_getEnumName(value))
			} else {
				if self.useCache && self.serializeRef(value) {
					return
				}
				_this_11 := self.buf
				_this_11.b = hxrt.StringConcatStringPtr(_this_11.b, hxrt.StringFromLiteral("o"))
				self.serializeFields(value)
			}
		}
	case 5:
		hxrt.Throw(hxrt.StringFromLiteral("Cannot serialize function"))
	case 6:
		var _g_1 any = _g.params[0]
		var declaration any = _g_1
		self.serializeClass(value, declaration)
	case 7:
		var _g_2 any = _g.params[0]
		var declaration_1 any = _g_2
		self.serializeEnum(value, declaration_1)
	case 8:
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot serialize "), hxrt.StdString(value)))
	}
}

func (self *haxe__Serializer) serializeException(value any) {
	_this := self.buf
	_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("x"))
	self.serialize(value)
}

var haxe__Serializer_BASE64 *string = hxrt.StringFromLiteral("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789%:")

var haxe__Serializer_USE_CACHE bool = false

var haxe__Serializer_USE_ENUM_INDEX bool = false

func haxe__Serializer_run(value any) *string {
	serializer := New_haxe__Serializer()
	serializer.serialize(value)
	return serializer.toString()
}
