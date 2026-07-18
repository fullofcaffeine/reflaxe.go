package main

import "examples_incident_api_portable/hxrt"

type I_haxe__io__BytesBuffer interface {
	get_length() int
	addByte(value int)
	add(source *haxe__io__Bytes)
	addString(value *string, encoding *haxe__io__Encoding)
	addInt32(value int)
	addInt64(value *haxe___Int64_____Int64)
	addFloat(value float64)
	addDouble(value float64)
	addBytes(source *haxe__io__Bytes, pos int, len int)
	getBytes() *haxe__io__Bytes
}

type haxe__io__BytesBuffer struct {
	__hx_this I_haxe__io__BytesBuffer
	b         []int
	length    int
}

func New_haxe__io__BytesBuffer() *haxe__io__BytesBuffer {
	self := &haxe__io__BytesBuffer{}
	self.__hx_this = self
	self.b = hxrt.BytesAllocValues(0)
	return self
}

func (self *haxe__io__BytesBuffer) get_length() int {
	return len(self.b)
}

func (self *haxe__io__BytesBuffer) addByte(value int) {
	self.b = hxrt.BytesBufferAddByte(self.b, value)
}

func (self *haxe__io__BytesBuffer) add(source *haxe__io__Bytes) {
	self.b = hxrt.BytesBufferAdd(self.b, source.__hx_this.getData())
}

func (self *haxe__io__BytesBuffer) addString(value *string, encoding *haxe__io__Encoding) {
	source := haxe__io__Bytes_ofString(value, encoding)
	self.b = hxrt.BytesBufferAdd(self.b, source.__hx_this.getData())
}

func (self *haxe__io__BytesBuffer) addInt32(value int) {
	self.b = hxrt.BytesBufferAddByte(self.b, int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))))
	self.b = hxrt.BytesBufferAddByte(self.b, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255)))))
	self.b = hxrt.BytesBufferAddByte(self.b, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(16))))) & hxrt.Int32Wrap(255)))))
	self.b = hxrt.BytesBufferAddByte(self.b, int(int32(int32((uint32(hxrt.Int32Wrap(value)) >> uint(24))))))
}

func (self *haxe__io__BytesBuffer) addInt64(value *haxe___Int64_____Int64) {
	self.__hx_this.addInt32(value.low)
	self.__hx_this.addInt32(value.high)
}

func (self *haxe__io__BytesBuffer) addFloat(value float64) {
	self.__hx_this.addInt32(haxe__io__FPHelper_floatToI32(value))
}

func (self *haxe__io__BytesBuffer) addDouble(value float64) {
	self.__hx_this.addInt64(haxe__io__FPHelper_doubleToI64(value))
}

func (self *haxe__io__BytesBuffer) addBytes(source *haxe__io__Bytes, pos int, len int) {
	if ((pos < 0) || (len < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(len)))) > source.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	self.b = hxrt.BytesBufferAddSlice(self.b, source.__hx_this.getData(), pos, len)
}

func (self *haxe__io__BytesBuffer) getBytes() *haxe__io__Bytes {
	return haxe__io__Bytes_ofData(hxrt.BytesClone(self.b))
}
