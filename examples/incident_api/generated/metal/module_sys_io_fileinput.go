package main

import "examples_incident_api_metal/hxrt"

type I_sys__io__FileInput interface {
	readByte() int
	readBytes(bytes *haxe__io__Bytes, pos int, length int) int
	seek(p int, pos *sys__io__FileSeek)
	tell() int
	eof() bool
	close()
}

type sys__io__FileInput struct {
	__hx_this         I_sys__io__FileInput
	handle            *hxrt.FileInput
	__hx_io_bigEndian bool
}

func New_sys__io__FileInput(handle *hxrt.FileInput) *sys__io__FileInput {
	self := &sys__io__FileInput{}
	self.__hx_this = self
	self.handle = handle
	return self
}

func (self *sys__io__FileInput) readByte() int {
	value := hxrt.FileInputReadByteValue(self.handle)
	if value < 0 {
		hxrt.Throw(New_haxe__io__Eof())
	}
	return value
}

func (self *sys__io__FileInput) readBytes(bytes *haxe__io__Bytes, pos int, length int) int {
	if ((pos < 0) || (length < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(length)))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	if length == 0 {
		return 0
	}
	values := hxrt.FileInputReadValues(self.handle, length)
	if len(values) == 0 {
		hxrt.Throw(New_haxe__io__Eof())
	}
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_61 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_61
		bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
	}
	return len(values)
}

func (self *sys__io__FileInput) seek(p int, pos *sys__io__FileSeek) {
	switch pos.tag {
	case 0:
		hxrt.FileInputSeek(self.handle, p, 0)
	case 1:
		hxrt.FileInputSeek(self.handle, p, 1)
	case 2:
		hxrt.FileInputSeek(self.handle, p, 2)
	}
}

func (self *sys__io__FileInput) tell() int {
	return hxrt.FileInputTell(self.handle)
}

func (self *sys__io__FileInput) eof() bool {
	return hxrt.FileInputEof(self.handle)
}

func (self *sys__io__FileInput) close() {
	hxrt.FileInputClose(self.handle)
}

func (self *sys__io__FileInput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.__hx_io_bigEndian
}

func (self *sys__io__FileInput) set_bigEndian(e bool) bool {
	if self != nil {
		self.__hx_io_bigEndian = e
	}
	return e
}

func (self *sys__io__FileInput) readAll(bufsize ...int) *haxe__io__Bytes {
	return haxe__io__input_readAll(self, bufsize...)
}

func (self *sys__io__FileInput) readFullBytes(s *haxe__io__Bytes, pos int, len int) {
	haxe__io__input_readFullBytes(self, s, pos, len)
}

func (self *sys__io__FileInput) read(nbytes int) *haxe__io__Bytes {
	return haxe__io__input_read(self, nbytes)
}

func (self *sys__io__FileInput) readUntil(end int) *string {
	return haxe__io__input_readUntil(self, end)
}

func (self *sys__io__FileInput) readLine() *string {
	return haxe__io__input_readLine(self)
}

func (self *sys__io__FileInput) readFloat() float64 {
	return haxe__io__input_readFloat(self)
}

func (self *sys__io__FileInput) readDouble() float64 {
	return haxe__io__input_readDouble(self)
}

func (self *sys__io__FileInput) readInt8() int {
	return haxe__io__input_readInt8(self)
}

func (self *sys__io__FileInput) readInt16() int {
	return haxe__io__input_readInt16(self)
}

func (self *sys__io__FileInput) readUInt16() int {
	return haxe__io__input_readUInt16(self)
}

func (self *sys__io__FileInput) readInt24() int {
	return haxe__io__input_readInt24(self)
}

func (self *sys__io__FileInput) readUInt24() int {
	return haxe__io__input_readUInt24(self)
}

func (self *sys__io__FileInput) readInt32() int {
	return haxe__io__input_readInt32(self)
}

func (self *sys__io__FileInput) readString(len int, encoding ...*haxe__io__Encoding) *string {
	return haxe__io__input_readString(self, len, encoding...)
}
