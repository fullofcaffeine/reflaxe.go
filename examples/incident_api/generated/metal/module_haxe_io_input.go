package main

import "examples_incident_api_metal/hxrt"

type I_haxe__io__Input interface {
	readByte() int
	readBytes(bytes *haxe__io__Bytes, pos int, len int) int
	close()
	set_bigEndian(value bool) bool
	readAll(bufsize any) *haxe__io__Bytes
	readFullBytes(bytes *haxe__io__Bytes, pos int, len int)
	read(nbytes int) *haxe__io__Bytes
	readUntil(end int) *string
	readLine() *string
	readFloat() float64
	readDouble() float64
	readInt8() int
	readInt16() int
	readUInt16() int
	readInt24() int
	readUInt24() int
	readInt32() int
	readString(len int, encoding *haxe__io__Encoding) *string
}

type haxe__io__Input struct {
	__hx_this I_haxe__io__Input
	bigEndian bool
}

func New_haxe__io__Input() *haxe__io__Input {
	self := &haxe__io__Input{}
	self.__hx_this = self
	return self
}

func (self *haxe__io__Input) readByte() int {
	return func() int {
		hxrt.Throw(New_haxe__exceptions__NotImplementedException(nil, nil, func() map[string]any {
			hx_obj_1 := map[string]any{}
			hx_obj_1["fileName"] = hxrt.StringFromLiteral("haxe/io/Input.hx")
			hx_obj_1["lineNumber"] = 18
			hx_obj_1["className"] = hxrt.StringFromLiteral("haxe.io.Input")
			hx_obj_1["methodName"] = hxrt.StringFromLiteral("readByte")
			return hx_obj_1
		}()))
		var hx_throw_zero_2 int
		return hx_throw_zero_2
	}()
}

func (self *haxe__io__Input) readBytes(bytes *haxe__io__Bytes, pos int, len int) int {
	if ((pos < 0) || (len < 0)) || (int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(len))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	remaining := len
	hxrt.TryCatch(func() {
		for remaining > 0 {
			value := self.__hx_this.readByte()
			bytes.b[pos] = int((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))
			bytes.__hx_rawValid = false
			pos = int(int32((pos + 1)))
			remaining = int(int32((remaining - 1)))
		}
	}, func(hx_caught_3 any) {
		switch hx_typed_4 := hx_caught_3.(type) {
		case *haxe__io__Eof:
			hx_tmp := hx_typed_4
			_ = hx_tmp
		default:
			hxrt.Throw(hx_caught_3)
		}
	})
	return int((hxrt.Int32Wrap(len) - hxrt.Int32Wrap(remaining)))
}

func (self *haxe__io__Input) close() {
}

func (self *haxe__io__Input) set_bigEndian(value bool) bool {
	self.bigEndian = value
	return value
}

func (self *haxe__io__Input) readAll(bufsize any) *haxe__io__Bytes {
	if bufsize == nil {
		bufsize = 16384
	}
	buffer := haxe__io__Bytes_alloc(hxrt.IntFromNullableAny(bufsize.(int)))
	total := New_haxe__io__BytesBuffer()
	hxrt.TryCatch(func() {
		for true {
			count := self.__hx_this.readBytes(buffer, 0, hxrt.IntFromNullableAny(bufsize.(int)))
			if count == 0 {
				hxrt.Throw(haxe__io__Error_Blocked)
			}
			if (count < 0) || (count > buffer.length) {
				hxrt.Throw(haxe__io__Error_OutsideBounds)
			}
			total.b = hxrt.BytesBufferAddSlice(total.b, buffer.__hx_this.getData(), 0, count)
		}
	}, func(hx_caught_5 any) {
		switch hx_typed_6 := hx_caught_5.(type) {
		case *haxe__io__Eof:
			hx_tmp := hx_typed_6
			_ = hx_tmp
		default:
			hxrt.Throw(hx_caught_5)
		}
	})
	return total.__hx_this.getBytes()
}

func (self *haxe__io__Input) readFullBytes(bytes *haxe__io__Bytes, pos int, len int) {
	for len > 0 {
		count := self.__hx_this.readBytes(bytes, pos, len)
		if count == 0 {
			hxrt.Throw(haxe__io__Error_Blocked)
		}
		pos = int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(count)))
		len = int((hxrt.Int32Wrap(len) - hxrt.Int32Wrap(count)))
	}
}

func (self *haxe__io__Input) read(nbytes int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(nbytes)
	self.__hx_this.readFullBytes(bytes, 0, nbytes)
	return bytes
}

func (self *haxe__io__Input) readUntil(end int) *string {
	buffer := New_haxe__io__BytesBuffer()
	var value int
	for func() int {
		value = self.__hx_this.readByte()
		return value
	}() != end {
		buffer.b = hxrt.BytesBufferAddByte(buffer.b, value)
	}
	return buffer.__hx_this.getBytes().__hx_this.toString()
}

func (self *haxe__io__Input) readLine() *string {
	buffer := New_haxe__io__BytesBuffer()
	var value int
	var result *string
	hxrt.TryCatch(func() {
		for func() int {
			value = self.__hx_this.readByte()
			return value
		}() != 10 {
			buffer.b = hxrt.BytesBufferAddByte(buffer.b, value)
		}
		result = buffer.__hx_this.getBytes().__hx_this.toString()
		if (hxrt.StringLengthStringPtr(result) > 0) && (hxrt.StringCharCodeAtAnyStringPtr(result, int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(result))-hxrt.Int32Wrap(1)))) == 13) {
			result = hxrt.StringSubstrStringPtr(result, 0, -1, true)
		}
	}, func(hx_caught_7 any) {
		switch hx_typed_8 := hx_caught_7.(type) {
		case *haxe__io__Eof:
			error := hx_typed_8
			result = buffer.__hx_this.getBytes().__hx_this.toString()
			if hxrt.StringLengthStringPtr(result) == 0 {
				hxrt.Throw(error)
			}
		default:
			hxrt.Throw(hx_caught_7)
		}
	})
	return result
}

func (self *haxe__io__Input) readFloat() float64 {
	return haxe__io__FPHelper_i32ToFloat(self.__hx_this.readInt32())
}

func (self *haxe__io__Input) readDouble() float64 {
	first := self.__hx_this.readInt32()
	second := self.__hx_this.readInt32()
	var hx_if_9 float64
	if self.bigEndian {
		hx_if_9 = haxe__io__FPHelper_i64ToDouble(second, first)
	} else {
		hx_if_9 = haxe__io__FPHelper_i64ToDouble(first, second)
	}
	return hx_if_9
}

func (self *haxe__io__Input) readInt8() int {
	value := self.__hx_this.readByte()
	var hx_if_10 int
	if value >= 128 {
		hx_if_10 = int((hxrt.Int32Wrap(value) - hxrt.Int32Wrap(256)))
	} else {
		hx_if_10 = value
	}
	return hx_if_10
}

func (self *haxe__io__Input) readInt16() int {
	first := self.__hx_this.readByte()
	second := self.__hx_this.readByte()
	var hx_if_11 int
	if self.bigEndian {
		hx_if_11 = int((hxrt.Int32Wrap(second) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(first) << uint(8))))))
	} else {
		hx_if_11 = int((hxrt.Int32Wrap(first) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(second) << uint(8))))))
	}
	value := hx_if_11
	var hx_if_12 int
	if int((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(32768))) != 0 {
		hx_if_12 = int((hxrt.Int32Wrap(value) - hxrt.Int32Wrap(65536)))
	} else {
		hx_if_12 = value
	}
	return hx_if_12
}

func (self *haxe__io__Input) readUInt16() int {
	first := self.__hx_this.readByte()
	second := self.__hx_this.readByte()
	var hx_if_13 int
	if self.bigEndian {
		hx_if_13 = int((hxrt.Int32Wrap(second) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(first) << uint(8))))))
	} else {
		hx_if_13 = int((hxrt.Int32Wrap(first) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(second) << uint(8))))))
	}
	return hx_if_13
}

func (self *haxe__io__Input) readInt24() int {
	first := self.__hx_this.readByte()
	second := self.__hx_this.readByte()
	third := self.__hx_this.readByte()
	var hx_if_14 int
	if self.bigEndian {
		hx_if_14 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(third) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(second) << uint(8))))))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(first) << uint(16))))))
	} else {
		hx_if_14 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(first) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(second) << uint(8))))))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(third) << uint(16))))))
	}
	value := hx_if_14
	var hx_if_15 int
	if int((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(8388608))) != 0 {
		hx_if_15 = int((hxrt.Int32Wrap(value) - hxrt.Int32Wrap(16777216)))
	} else {
		hx_if_15 = value
	}
	return hx_if_15
}

func (self *haxe__io__Input) readUInt24() int {
	first := self.__hx_this.readByte()
	second := self.__hx_this.readByte()
	third := self.__hx_this.readByte()
	var hx_if_16 int
	if self.bigEndian {
		hx_if_16 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(third) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(second) << uint(8))))))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(first) << uint(16))))))
	} else {
		hx_if_16 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(first) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(second) << uint(8))))))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(third) << uint(16))))))
	}
	return hx_if_16
}

func (self *haxe__io__Input) readInt32() int {
	first := self.__hx_this.readByte()
	second := self.__hx_this.readByte()
	third := self.__hx_this.readByte()
	fourth := self.__hx_this.readByte()
	var hx_if_17 int
	if self.bigEndian {
		hx_if_17 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(fourth) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(third) << uint(8))))))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(second) << uint(16))))))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(first) << uint(24))))))
	} else {
		hx_if_17 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(first) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(second) << uint(8))))))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(third) << uint(16))))))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(fourth) << uint(24))))))
	}
	return hx_if_17
}

func (self *haxe__io__Input) readString(len int, encoding *haxe__io__Encoding) *string {
	bytes := haxe__io__Bytes_alloc(len)
	self.__hx_this.readFullBytes(bytes, 0, len)
	return bytes.__hx_this.getString(0, len, encoding)
}
