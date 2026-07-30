package main

import "snapshot/hxrt"

type I_sys__net__SocketOutput interface {
	writeByte(value int)
	writeBytes(bytes *haxe__io__Bytes, pos int, length int) int
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

type sys__net__SocketOutput struct {
	*haxe__io__Output
	__hx_this I_sys__net__SocketOutput
	handle    *hxrt.SocketHandle
}

func New_sys__net__SocketOutput(handle *hxrt.SocketHandle) *sys__net__SocketOutput {
	self := &sys__net__SocketOutput{}
	self.haxe__io__Output = New_haxe__io__Output()
	self.haxe__io__Output.__hx_this = self
	self.__hx_this = self
	self.handle = handle
	return self
}

func (self *sys__net__SocketOutput) writeByte(value int) {
	result := hxrt.SocketWriteValues(self.handle, []int{value})
	sys__net__SocketOutput_translateWriteStatus(result)
}

func (self *sys__net__SocketOutput) writeBytes(bytes *haxe__io__Bytes, pos int, length int) int {
	if ((pos < 0) || (length < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(length)))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	if length == 0 {
		return 0
	}
	values := hxrt.NewArray()
	_g := 0
	_g1 := length
	for _g < _g1 {
		hx_post_37 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_37
		values.Push(bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))])
	}
	result := hxrt.SocketWriteValues(self.handle, func(hx_lambda_raw_39 []any) []int {
		hx_lambda_out_40 := make([]int, 0, len(hx_lambda_raw_39))
		for _, hx_lambda_item_41 := range hx_lambda_raw_39 {
			hx_lambda_out_40 = append(hx_lambda_out_40, func(hx_value_42 any) int {
				if hx_value_42 == nil {
					var hx_zero_43 int
					return hx_zero_43
				}
				return hx_value_42.(int)
			}(hx_lambda_item_41))
		}
		return hx_lambda_out_40
	}(values.Values()))
	sys__net__SocketOutput_translateWriteStatus(result)
	return result.Count
}

func (self *sys__net__SocketOutput) flush() {
	hxrt.SocketFlush(self.handle)
}

func (self *sys__net__SocketOutput) close() {
	hxrt.SocketClose(self.handle)
}

func sys__net__SocketOutput_translateWriteStatus(result *hxrt.SocketIOResult) {
	if result.Status == hxrt.SocketIOBlocked {
		hxrt.Throw(haxe__io__Error_Blocked)
	}
	if result.Status == hxrt.SocketIOEOF {
		hxrt.Throw(New_haxe__io__Eof())
	}
}

type I_sys__net__SocketInput interface {
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
}

type sys__net__SocketInput struct {
	*haxe__io__Input
	__hx_this I_sys__net__SocketInput
	handle    *hxrt.SocketHandle
}

func New_sys__net__SocketInput(handle *hxrt.SocketHandle) *sys__net__SocketInput {
	self := &sys__net__SocketInput{}
	self.haxe__io__Input = New_haxe__io__Input()
	self.haxe__io__Input.__hx_this = self
	self.__hx_this = self
	self.handle = handle
	return self
}

func (self *sys__net__SocketInput) readByte() int {
	value := hxrt.SocketReadByteValue(self.handle)
	if value == hxrt.SocketReadBlocked {
		hxrt.Throw(haxe__io__Error_Blocked)
	}
	if value == hxrt.SocketReadEOF {
		hxrt.Throw(New_haxe__io__Eof())
	}
	return value
}

func (self *sys__net__SocketInput) readBytes(bytes *haxe__io__Bytes, pos int, length int) int {
	if ((pos < 0) || (length < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(length)))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	if length == 0 {
		return 0
	}
	result := hxrt.SocketReadValues(self.handle, length)
	sys__net__SocketInput_translateReadStatus(result)
	_g := 0
	_g1 := result.Count
	for _g < _g1 {
		hx_post_44 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_44
		value := result.Values[index]
		bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	return result.Count
}

func (self *sys__net__SocketInput) close() {
	hxrt.SocketClose(self.handle)
}

func sys__net__SocketInput_translateReadStatus(result *hxrt.SocketIOResult) {
	if result.Status == hxrt.SocketIOBlocked {
		hxrt.Throw(haxe__io__Error_Blocked)
	}
	if result.Status == hxrt.SocketIOEOF {
		hxrt.Throw(New_haxe__io__Eof())
	}
}
