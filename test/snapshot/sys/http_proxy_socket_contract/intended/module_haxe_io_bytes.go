package main

import "snapshot/hxrt"

type I_haxe__io__Bytes interface {
	get(pos int) int
	set(pos int, value int)
	blit(pos int, src *haxe__io__Bytes, srcpos int, len int)
	fill(pos int, len int, value int)
	sub(pos int, len int) *haxe__io__Bytes
	compare(other *haxe__io__Bytes) int
	getDouble(pos int) float64
	getFloat(pos int) float64
	setDouble(pos int, value float64)
	setFloat(pos int, value float64)
	getUInt16(pos int) int
	setUInt16(pos int, value int)
	getInt32(pos int) int
	getInt64(pos int) *haxe___Int64_____Int64
	setInt32(pos int, value int)
	setInt64(pos int, value *haxe___Int64_____Int64)
	getString(pos int, len int, encoding *haxe__io__Encoding) *string
	readString(pos int, len int) *string
	toString() *string
	toHex() *string
	getData() []int
	__hx_nativeView() *hxrt.ByteView
}

type haxe__io__Bytes struct {
	__hx_this        I_haxe__io__Bytes
	length           int
	b                []int
	__hx_raw         *hxrt.ByteView
	__hx_rawValid    bool
	__hx_dataExposed bool
}

func New_haxe__io__Bytes(length int, b []int, raw *hxrt.ByteView) *haxe__io__Bytes {
	self := &haxe__io__Bytes{}
	self.__hx_this = self
	self.length = length
	self.b = b
	self.__hx_raw = raw
	self.__hx_rawValid = (raw != nil)
	self.__hx_dataExposed = false
	return self
}

func (self *haxe__io__Bytes) get(pos int) int {
	return self.b[pos]
}

func (self *haxe__io__Bytes) set(pos int, value int) {
	self.b[pos] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) blit(pos int, src *haxe__io__Bytes, srcpos int, len int) {
	if ((((pos < 0) || (srcpos < 0)) || (len < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(len)))) > self.length)) || (int(int32((hxrt.Int32Wrap(srcpos) + hxrt.Int32Wrap(len)))) > src.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	if len == 0 {
		return
	}
	hxrt.BytesBlitValues(self.b, pos, src.b, srcpos, len)
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) fill(pos int, len int, value int) {
	if ((pos < 0) || (len < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(len)))) > self.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	masked := int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
	_g := 0
	_g1 := len
	for _g < _g1 {
		hx_post_224 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_224
		self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))] = masked
	}
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) sub(pos int, len int) *haxe__io__Bytes {
	if ((pos < 0) || (len < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(len)))) > self.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	out := haxe__io__Bytes_alloc(len)
	_g := 0
	_g1 := len
	for _g < _g1 {
		hx_post_225 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_225
		out.b[index] = self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))]
	}
	return out
}

func (self *haxe__io__Bytes) compare(other *haxe__io__Bytes) int {
	var hx_if_226 int
	if self.length < other.length {
		hx_if_226 = self.length
	} else {
		hx_if_226 = other.length
	}
	limit := hx_if_226
	_g := 0
	_g1 := limit
	for _g < _g1 {
		hx_post_227 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_227
		if self.b[index] < other.b[index] {
			return -1
		}
		if self.b[index] > other.b[index] {
			return 1
		}
	}
	var hx_if_229 int
	if self.length < other.length {
		hx_if_229 = -1
	} else {
		var hx_if_228 int
		if self.length > other.length {
			hx_if_228 = 1
		} else {
			hx_if_228 = 0
		}
		hx_if_229 = hx_if_228
	}
	return hx_if_229
}

func (self *haxe__io__Bytes) getDouble(pos int) float64 {
	return haxe__io__FPHelper_i64ToDouble(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[pos]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(3))))]) << uint(24)))))))), func() int {
		pos_1 := int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(4))))
		return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[pos_1]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos_1)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos_1)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos_1)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
	}())
}

func (self *haxe__io__Bytes) getFloat(pos int) float64 {
	return haxe__io__FPHelper_i32ToFloat(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[pos]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(3))))]) << uint(24)))))))))
}

func (self *haxe__io__Bytes) setDouble(pos int, value float64) {
	bits := haxe__io__FPHelper_doubleToI64(value)
	value_1 := bits.low
	self.b[pos] = int(int32((hxrt.Int32Wrap(value_1) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value_1) >> uint(8))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value_1) >> uint(16))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value_1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	pos_1 := int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(4))))
	value_2 := bits.high
	self.b[pos_1] = int(int32((hxrt.Int32Wrap(value_2) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value_2) >> uint(8))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value_2) >> uint(16))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value_2)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) setFloat(pos int, value float64) {
	value_1 := haxe__io__FPHelper_floatToI32(value)
	self.b[pos] = int(int32((hxrt.Int32Wrap(value_1) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value_1) >> uint(8))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value_1) >> uint(16))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value_1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) getUInt16(pos int) int {
	return int(int32((hxrt.Int32Wrap(self.b[pos]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1))))]) << uint(8))))))))
}

func (self *haxe__io__Bytes) setUInt16(pos int, value int) {
	self.b[pos] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) getInt32(pos int) int {
	return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[pos]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
}

func (self *haxe__io__Bytes) getInt64(pos int) *haxe___Int64_____Int64 {
	pos_1 := int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(4))))
	high := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[pos_1]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos_1)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos_1)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos_1)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
	low := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[pos]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(self.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
	x := New_haxe___Int64_____Int64(high, low)
	var this1 *haxe___Int64_____Int64
	this1 = x
	return this1
}

func (self *haxe__io__Bytes) setInt32(pos int, value int) {
	self.b[pos] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(16))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) setInt64(pos int, value *haxe___Int64_____Int64) {
	value1 := value.low
	self.b[pos] = int(int32((hxrt.Int32Wrap(value1) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value1) >> uint(8))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value1) >> uint(16))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	pos_1 := int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(4))))
	value_1 := value.high
	self.b[pos_1] = int(int32((hxrt.Int32Wrap(value_1) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value_1) >> uint(8))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value_1) >> uint(16))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
	self.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value_1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) getString(pos int, len int, encoding *haxe__io__Encoding) *string {
	if ((pos < 0) || (len < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(len)))) > self.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	return hxrt.StdString(hxrt.BytesStringFromView(self.__hx_this.__hx_nativeView(), pos, len, ((encoding == haxe__io__Encoding_RawNative) && false)))
}

func (self *haxe__io__Bytes) readString(pos int, len int) *string {
	return self.__hx_this.getString(pos, len, nil)
}

func (self *haxe__io__Bytes) toString() *string {
	return self.__hx_this.getString(0, self.length, nil)
}

func (self *haxe__io__Bytes) toHex() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	digits := hxrt.StringFromLiteral("0123456789abcdef")
	_g := 0
	_g1 := self.length
	for _g < _g1 {
		hx_post_230 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_230
		value := self.b[index]
		c := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(digits, int(int32((hxrt.Int32Wrap(value) >> uint(4))))))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromCharCode(c))
		c_1 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(digits, int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(15))))))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromCharCode(c_1))
	}
	return out_b
}

func (self *haxe__io__Bytes) getData() []int {
	self.__hx_dataExposed = true
	return self.b
}

func (self *haxe__io__Bytes) __hx_nativeView() *hxrt.ByteView {
	if !self.__hx_rawValid || (self.__hx_dataExposed && !hxrt.BytesViewMatchesValues(self.__hx_raw, self.b)) {
		self.__hx_raw = hxrt.BytesViewFromValues(self.b)
		self.__hx_rawValid = true
	}
	return self.__hx_raw
}

func (self *haxe__io__Bytes) String() string {
	return *self.__hx_this.toString()
}

func haxe__io__Bytes___hx_fromNativeView(view *hxrt.ByteView) *haxe__io__Bytes {
	return New_haxe__io__Bytes(hxrt.BytesViewLength(view), hxrt.BytesValuesFromView(view), view)
}

func haxe__io__Bytes_alloc(length int) *haxe__io__Bytes {
	return New_haxe__io__Bytes(length, hxrt.BytesAllocValues(length), nil)
}

func haxe__io__Bytes_fastGet(data []int, pos int) int {
	return data[pos]
}

func haxe__io__Bytes_ofData(data []int) *haxe__io__Bytes {
	if data == nil {
		data = hxrt.BytesAllocValues(0)
	}
	bytes := New_haxe__io__Bytes(len(data), data, nil)
	bytes.__hx_dataExposed = true
	return bytes
}

func haxe__io__Bytes_ofHex(value *string) *haxe__io__Bytes {
	textLength := hxrt.StringLengthStringPtr(value)
	if int(int32((hxrt.Int32Wrap(textLength) & hxrt.Int32Wrap(1)))) != 0 {
		hxrt.Throw(hxrt.StringFromLiteral("Not a hex string (odd number of digits)"))
	}
	out := haxe__io__Bytes_alloc(int(int32((hxrt.Int32Wrap(textLength) >> uint(1)))))
	_g := 0
	_g1 := out.length
	for _g < _g1 {
		hx_post_231 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_231
		var c any = hxrt.StringCharCodeAtAnyStringPtr(value, int(int32((hxrt.Int32Wrap(index) * hxrt.Int32Wrap(2)))))
		var hx_if_232 int
		if c == nil {
			hx_if_232 = -1
		} else {
			hx_if_232 = c.(int)
		}
		high := hx_if_232
		var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(value, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(index) * hxrt.Int32Wrap(2))))) + hxrt.Int32Wrap(1)))))
		var hx_if_233 int
		if c_1 == nil {
			hx_if_233 = -1
		} else {
			hx_if_233 = c_1.(int)
		}
		low := hx_if_233
		high = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(high) & hxrt.Int32Wrap(15))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(high) & hxrt.Int32Wrap(64))))) >> uint(6))))) * hxrt.Int32Wrap(9))))))))
		low = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(low) & hxrt.Int32Wrap(15))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(low) & hxrt.Int32Wrap(64))))) >> uint(6))))) * hxrt.Int32Wrap(9))))))))
		out.b[index] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(high) << uint(4))))) | hxrt.Int32Wrap(low))))) & hxrt.Int32Wrap(255))))) & hxrt.Int32Wrap(255))))
		out.__hx_rawValid = false
	}
	return out
}

func haxe__io__Bytes_ofString(value *string, encoding *haxe__io__Encoding) *haxe__io__Bytes {
	view := hxrt.BytesViewFromString(value, ((encoding == haxe__io__Encoding_RawNative) && false))
	return haxe__io__Bytes___hx_fromNativeView(view)
}

func haxe__io__Bytes_rawNativeUsesUtf16LE() bool {
	return false
}
