package main

import "snapshot/hxrt"

type I_sys__io__FileInput interface {
	readByte() int
	readBytes(bytes *haxe__io__Bytes, pos int, length int) int
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
	seek(p int, pos *sys__io__FileSeek)
	tell() int
	eof() bool
}

type sys__io__FileInput struct {
	*haxe__io__Input
	__hx_this I_sys__io__FileInput
	handle    *hxrt.FileInput
}

func New_sys__io__FileInput(handle *hxrt.FileInput) *sys__io__FileInput {
	self := &sys__io__FileInput{}
	self.haxe__io__Input = New_haxe__io__Input()
	self.haxe__io__Input.__hx_this = self
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
		hx_post_62 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_62
		bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
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
