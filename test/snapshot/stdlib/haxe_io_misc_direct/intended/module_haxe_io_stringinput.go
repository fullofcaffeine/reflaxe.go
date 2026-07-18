package main

type I_haxe__io__StringInput interface {
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

type haxe__io__StringInput struct {
	*haxe__io__BytesInput
	__hx_this I_haxe__io__StringInput
}

func New_haxe__io__StringInput(value *string) *haxe__io__StringInput {
	self := &haxe__io__StringInput{}
	self.haxe__io__BytesInput = New_haxe__io__BytesInput(haxe__io__Bytes_ofString(value, nil), nil, nil)
	self.haxe__io__BytesInput.haxe__io__Input.__hx_this = self
	self.haxe__io__BytesInput.__hx_this = self
	self.__hx_this = self
	return self
}
