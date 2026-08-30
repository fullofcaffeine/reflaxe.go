package main

import "snapshot/hxrt"

type I_haxe__io__BufferInput interface {
	readByte() int
	readBytes(bytes *haxe__io__Bytes, targetPos int, len int) int
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
	refill()
}

type haxe__io__BufferInput struct {
	*haxe__io__Input
	__hx_this I_haxe__io__BufferInput
	i         *haxe__io__Input
	buf       *haxe__io__Bytes
	available int
	pos       int
}

func New_haxe__io__BufferInput(input *haxe__io__Input, buffer *haxe__io__Bytes, pos int, available int) *haxe__io__BufferInput {
	self := &haxe__io__BufferInput{}
	self.haxe__io__Input = New_haxe__io__Input()
	self.haxe__io__Input.__hx_this = self
	self.__hx_this = self
	self.i = input
	self.buf = buffer
	self.pos = pos
	self.available = available
	return self
}

func (self *haxe__io__BufferInput) refill() {
	if self.pos > 0 {
		self.buf.__hx_this.blit(0, self.buf, self.pos, self.available)
		self.pos = 0
	}
	self.available = int((hxrt.Int32Wrap(self.available) + hxrt.Int32Wrap(self.i.__hx_this.readBytes(self.buf, self.available, int((hxrt.Int32Wrap(self.buf.length)-hxrt.Int32Wrap(self.available)))))))
}

func (self *haxe__io__BufferInput) readByte() int {
	if self.available == 0 {
		self.__hx_this.refill()
	}
	value := self.buf.b[self.pos]
	self.pos = int(int32((self.pos + 1)))
	self.available = int(int32((self.available - 1)))
	return value
}

func (self *haxe__io__BufferInput) readBytes(bytes *haxe__io__Bytes, targetPos int, len int) int {
	if self.available == 0 {
		self.__hx_this.refill()
	}
	var hx_if_1 int
	if len > self.available {
		hx_if_1 = self.available
	} else {
		hx_if_1 = len
	}
	size := hx_if_1
	bytes.__hx_this.blit(targetPos, self.buf, self.pos, size)
	self.pos = int((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(size)))
	self.available = int((hxrt.Int32Wrap(self.available) - hxrt.Int32Wrap(size)))
	return size
}
