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

func (self *haxe__io__StringInput) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__io__StringInput(value *string) *haxe__io__StringInput {
	return hx_init_haxe__io__StringInput(hx_alloc_haxe__io__StringInput(nil), value)
}

func hx_alloc_haxe__io__StringInput(owner any) *haxe__io__StringInput {
	self := &haxe__io__StringInput{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__io__StringInput)
	}
	self.haxe__io__BytesInput = hx_alloc_haxe__io__BytesInput(self.__hx_this)
	return self
}

func hx_init_haxe__io__StringInput(self *haxe__io__StringInput, value *string) *haxe__io__StringInput {
	hx_init_haxe__io__BytesInput(self.haxe__io__BytesInput, haxe__io__Bytes_ofString(value, nil), nil, nil)
	return self
}
