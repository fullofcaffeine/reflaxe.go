package main

import "snapshot/hxrt"

type I_haxe__io__BytesInput interface {
	readByte() int
	readBytes(bytes *haxe__io__Bytes, targetPos int, requested int) int
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
	get_position() int
	get_length() int
	set_position(value int) int
}

type haxe__io__BytesInput struct {
	*haxe__io__Input
	__hx_this I_haxe__io__BytesInput
	b         []int
	pos       int
	len       int
	totlen    int
	position  int
	length    int
}

func New_haxe__io__BytesInput(bytes *haxe__io__Bytes, pos any, len any) *haxe__io__BytesInput {
	self := &haxe__io__BytesInput{}
	self.haxe__io__Input = New_haxe__io__Input()
	self.haxe__io__Input.__hx_this = self
	self.__hx_this = self
	if pos == nil {
		pos = 0
	}
	if len == nil {
		len = int(int32((hxrt.Int32Wrap(bytes.length) - hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(pos.(int)))))))
	}
	if ((hxrt.IntFromNullableAny(pos.(int)) < 0) || (hxrt.IntFromNullableAny(len.(int)) < 0)) || (int(int32((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(pos.(int)))) + hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(len.(int))))))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	self.b = bytes.__hx_this.getData()
	self.pos = pos.(int)
	self.len = len.(int)
	self.totlen = len.(int)
	return self
}

func (self *haxe__io__BytesInput) get_position() int {
	return self.pos
}

func (self *haxe__io__BytesInput) get_length() int {
	return self.totlen
}

func (self *haxe__io__BytesInput) set_position(value int) int {
	if value < 0 {
		value = 0
	} else {
		if value > self.totlen {
			value = self.totlen
		}
	}
	self.len = int(int32((hxrt.Int32Wrap(self.totlen) - hxrt.Int32Wrap(value))))
	return func() int {
		self.pos = value
		return self.pos
	}()
}

func (self *haxe__io__BytesInput) readByte() int {
	if self.len == 0 {
		hxrt.Throw(New_haxe__io__Eof())
	}
	self.len = int(int32((self.len - 1)))
	hx_post_21 := self.pos
	self.pos = int(int32((self.pos + 1)))
	return self.b[hx_post_21]
}

func (self *haxe__io__BytesInput) readBytes(bytes *haxe__io__Bytes, targetPos int, requested int) int {
	if ((targetPos < 0) || (requested < 0)) || (int(int32((hxrt.Int32Wrap(targetPos) + hxrt.Int32Wrap(requested)))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	if (self.len == 0) && (requested > 0) {
		hxrt.Throw(New_haxe__io__Eof())
	}
	if requested > self.len {
		requested = self.len
	}
	_g := 0
	_g1 := requested
	for _g < _g1 {
		hx_post_22 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_22
		value := self.b[int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(index))))]
		bytes.b[int(int32((hxrt.Int32Wrap(targetPos) + hxrt.Int32Wrap(index))))] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	self.pos = int(int32((hxrt.Int32Wrap(self.pos) + hxrt.Int32Wrap(requested))))
	self.len = int(int32((hxrt.Int32Wrap(self.len) - hxrt.Int32Wrap(requested))))
	return requested
}
