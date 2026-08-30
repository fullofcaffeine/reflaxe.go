package main

import "snapshot/hxrt"

type I_haxe__io__BytesOutput interface {
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
	get_length() int
	getBytes() *haxe__io__Bytes
}

type haxe__io__BytesOutput struct {
	*haxe__io__Output
	__hx_this I_haxe__io__BytesOutput
	b         *haxe__io__BytesBuffer
	length    int
}

func New_haxe__io__BytesOutput() *haxe__io__BytesOutput {
	self := &haxe__io__BytesOutput{}
	self.haxe__io__Output = New_haxe__io__Output()
	self.haxe__io__Output.__hx_this = self
	self.__hx_this = self
	self.b = New_haxe__io__BytesBuffer()
	return self
}

func (self *haxe__io__BytesOutput) get_length() int {
	return len(self.b.b)
}

func (self *haxe__io__BytesOutput) writeByte(value int) {
	_this := self.b
	_this.b = hxrt.BytesBufferAddByte(_this.b, value)
}

func (self *haxe__io__BytesOutput) writeBytes(bytes *haxe__io__Bytes, pos int, len int) int {
	_this := self.b
	if ((pos < 0) || (len < 0)) || (int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(len))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	_this.b = hxrt.BytesBufferAddSlice(_this.b, bytes.__hx_this.getData(), pos, len)
	return len
}

func (self *haxe__io__BytesOutput) getBytes() *haxe__io__Bytes {
	return self.b.__hx_this.getBytes()
}
