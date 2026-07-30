package main

import "snapshot/hxrt"

type I_haxe__io__Output interface {
	writeByte(value int)
	writeBytes(bytes *haxe__io__Bytes, pos int, len int) int
	flush()
	close()
	set_bigEndian(value bool) bool
	write(bytes *haxe__io__Bytes)
	writeFullBytes(bytes *haxe__io__Bytes, pos int, len int)
	writeFloat(value float64)
	writeDouble(value float64)
	writeInt8(value int)
	writeInt16(value int)
	writeUInt16(value int)
	writeInt24(value int)
	writeUInt24(value int)
	writeInt32(value int)
	prepare(nbytes int)
	writeInput(input *haxe__io__Input, bufsize any)
	writeString(value *string, encoding *haxe__io__Encoding)
}

type haxe__io__Output struct {
	__hx_this I_haxe__io__Output
	bigEndian bool
}

func New_haxe__io__Output() *haxe__io__Output {
	self := &haxe__io__Output{}
	self.__hx_this = self
	return self
}

func (self *haxe__io__Output) writeByte(value int) {
	hxrt.Throw(New_haxe__exceptions__NotImplementedException(nil, nil, func() map[string]any {
		hx_obj_48 := map[string]any{}
		hx_obj_48["fileName"] = hxrt.StringFromLiteral("haxe/io/Output.hx")
		hx_obj_48["lineNumber"] = 17
		hx_obj_48["className"] = hxrt.StringFromLiteral("haxe.io.Output")
		hx_obj_48["methodName"] = hxrt.StringFromLiteral("writeByte")
		return hx_obj_48
	}()))
}

func (self *haxe__io__Output) writeBytes(bytes *haxe__io__Bytes, pos int, len int) int {
	if ((pos < 0) || (len < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(len)))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	total := len
	for len > 0 {
		self.__hx_this.writeByte(bytes.b[pos])
		pos = int(int32((pos + 1)))
		len = int(int32((len - 1)))
	}
	return total
}

func (self *haxe__io__Output) flush() {
}

func (self *haxe__io__Output) close() {
}

func (self *haxe__io__Output) set_bigEndian(value bool) bool {
	self.bigEndian = value
	return value
}

func (self *haxe__io__Output) write(bytes *haxe__io__Bytes) {
	self.__hx_this.writeFullBytes(bytes, 0, bytes.length)
}

func (self *haxe__io__Output) writeFullBytes(bytes *haxe__io__Bytes, pos int, len int) {
	for len > 0 {
		count := self.__hx_this.writeBytes(bytes, pos, len)
		if count == 0 {
			hxrt.Throw(haxe__io__Error_Blocked)
		}
		pos = int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(count))))
		len = int(int32((hxrt.Int32Wrap(len) - hxrt.Int32Wrap(count))))
	}
}

func (self *haxe__io__Output) writeFloat(value float64) {
	self.__hx_this.writeInt32(haxe__io__FPHelper_floatToI32(value))
}

func (self *haxe__io__Output) writeDouble(value float64) {
	bits := haxe__io__FPHelper_doubleToI64(value)
	if self.bigEndian {
		self.__hx_this.writeInt32(bits.high)
		self.__hx_this.writeInt32(bits.low)
	} else {
		self.__hx_this.writeInt32(bits.low)
		self.__hx_this.writeInt32(bits.high)
	}
}

func (self *haxe__io__Output) writeInt8(value int) {
	if (value < -128) || (value >= 128) {
		hxrt.Throw(haxe__io__Error_Overflow)
	}
	self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))))
}

func (self *haxe__io__Output) writeInt16(value int) {
	if (value < -32768) || (value >= 32768) {
		hxrt.Throw(haxe__io__Error_Overflow)
	}
	self.__hx_this.writeUInt16(int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(65535)))))
}

func (self *haxe__io__Output) writeUInt16(value int) {
	if (value < 0) || (value >= 65536) {
		hxrt.Throw(haxe__io__Error_Overflow)
	}
	if self.bigEndian {
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) >> uint(8)))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))))
	} else {
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) >> uint(8)))))
	}
}

func (self *haxe__io__Output) writeInt24(value int) {
	if (value < -8388608) || (value >= 8388608) {
		hxrt.Throw(haxe__io__Error_Overflow)
	}
	self.__hx_this.writeUInt24(int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(16777215)))))
}

func (self *haxe__io__Output) writeUInt24(value int) {
	if (value < 0) || (value >= 16777216) {
		hxrt.Throw(haxe__io__Error_Overflow)
	}
	if self.bigEndian {
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) >> uint(16)))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255)))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))))
	} else {
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255)))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) >> uint(16)))))
	}
}

func (self *haxe__io__Output) writeInt32(value int) {
	if self.bigEndian {
		self.__hx_this.writeByte(int(int32(int32((uint32(hxrt.Int32Wrap(value)) >> uint(24))))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(16))))) & hxrt.Int32Wrap(255)))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255)))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))))
	} else {
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255)))))
		self.__hx_this.writeByte(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(16))))) & hxrt.Int32Wrap(255)))))
		self.__hx_this.writeByte(int(int32(int32((uint32(hxrt.Int32Wrap(value)) >> uint(24))))))
	}
}

func (self *haxe__io__Output) prepare(nbytes int) {
}

func (self *haxe__io__Output) writeInput(input *haxe__io__Input, bufsize any) {
	if bufsize == nil {
		bufsize = 4096
	}
	buffer := haxe__io__Bytes_alloc(hxrt.IntFromNullableAny(bufsize.(int)))
	hxrt.TryCatch(func() {
		for true {
			count := input.__hx_this.readBytes(buffer, 0, hxrt.IntFromNullableAny(bufsize.(int)))
			if count == 0 {
				hxrt.Throw(haxe__io__Error_Blocked)
			}
			self.__hx_this.writeFullBytes(buffer, 0, count)
		}
	}, func(hx_caught_49 any) {
		switch hx_typed_50 := hx_caught_49.(type) {
		case *haxe__io__Eof:
			hx_tmp := hx_typed_50
			_ = hx_tmp
		default:
			hxrt.Throw(hx_caught_49)
		}
	})
}

func (self *haxe__io__Output) writeString(value *string, encoding *haxe__io__Encoding) {
	bytes := haxe__io__Bytes_ofString(value, encoding)
	self.__hx_this.writeFullBytes(bytes, 0, bytes.length)
}
