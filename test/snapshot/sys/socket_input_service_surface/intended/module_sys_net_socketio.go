package main

import "snapshot/hxrt"

type I_sys__net__SocketOutput interface {
	writeByte(value int)
	writeBytes(bytes *haxe__io__Bytes, pos int, length int) int
	flush()
	close()
}

type sys__net__SocketOutput struct {
	__hx_this         I_sys__net__SocketOutput
	handle            *hxrt.SocketHandle
	__hx_io_bigEndian bool
}

func New_sys__net__SocketOutput(handle *hxrt.SocketHandle) *sys__net__SocketOutput {
	self := &sys__net__SocketOutput{}
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
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_1
		values.Push(bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))])
	}
	result := hxrt.SocketWriteValues(self.handle, func(hx_lambda_raw_3 []any) []int {
		hx_lambda_out_4 := make([]int, 0, len(hx_lambda_raw_3))
		for _, hx_lambda_item_5 := range hx_lambda_raw_3 {
			hx_lambda_out_4 = append(hx_lambda_out_4, func(hx_value_6 any) int {
				if hx_value_6 == nil {
					var hx_zero_7 int
					return hx_zero_7
				}
				return hx_value_6.(int)
			}(hx_lambda_item_5))
		}
		return hx_lambda_out_4
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

func (self *sys__net__SocketOutput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.__hx_io_bigEndian
}

func (self *sys__net__SocketOutput) set_bigEndian(e bool) bool {
	if self != nil {
		self.__hx_io_bigEndian = e
	}
	return e
}

func (self *sys__net__SocketOutput) prepare(nbytes int) {
	_ = self
	_ = nbytes
}

func (self *sys__net__SocketOutput) write(s *haxe__io__Bytes) {
	haxe__io__output_write(self, s)
}

func (self *sys__net__SocketOutput) writeFullBytes(s *haxe__io__Bytes, pos int, len int) {
	haxe__io__output_writeFullBytes(self, s, pos, len)
}

func (self *sys__net__SocketOutput) writeFloat(x float64) {
	haxe__io__output_writeFloat(self, x)
}

func (self *sys__net__SocketOutput) writeDouble(x float64) {
	haxe__io__output_writeDouble(self, x)
}

func (self *sys__net__SocketOutput) writeInt8(x int) {
	haxe__io__output_writeInt8(self, x)
}

func (self *sys__net__SocketOutput) writeInt16(x int) {
	haxe__io__output_writeInt16(self, x)
}

func (self *sys__net__SocketOutput) writeUInt16(x int) {
	haxe__io__output_writeUInt16(self, x)
}

func (self *sys__net__SocketOutput) writeInt24(x int) {
	haxe__io__output_writeInt24(self, x)
}

func (self *sys__net__SocketOutput) writeUInt24(x int) {
	haxe__io__output_writeUInt24(self, x)
}

func (self *sys__net__SocketOutput) writeInt32(x int) {
	haxe__io__output_writeInt32(self, x)
}

func (self *sys__net__SocketOutput) writeInput(i haxe__io__Input, bufsize ...int) {
	haxe__io__output_writeInput(self, i, bufsize...)
}

func (self *sys__net__SocketOutput) writeString(s *string, encoding ...*haxe__io__Encoding) {
	haxe__io__output_writeString(self, s, encoding...)
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
}

type sys__net__SocketInput struct {
	__hx_this         I_sys__net__SocketInput
	handle            *hxrt.SocketHandle
	__hx_io_bigEndian bool
}

func New_sys__net__SocketInput(handle *hxrt.SocketHandle) *sys__net__SocketInput {
	self := &sys__net__SocketInput{}
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
		hx_post_8 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_8
		v := result.Values[index]
		bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))] = int(int32((hxrt.Int32Wrap(v) & hxrt.Int32Wrap(255))))
	}
	return result.Count
}

func (self *sys__net__SocketInput) close() {
	hxrt.SocketClose(self.handle)
}

func (self *sys__net__SocketInput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.__hx_io_bigEndian
}

func (self *sys__net__SocketInput) set_bigEndian(e bool) bool {
	if self != nil {
		self.__hx_io_bigEndian = e
	}
	return e
}

func (self *sys__net__SocketInput) readAll(bufsize ...int) *haxe__io__Bytes {
	return haxe__io__input_readAll(self, bufsize...)
}

func (self *sys__net__SocketInput) readFullBytes(s *haxe__io__Bytes, pos int, len int) {
	haxe__io__input_readFullBytes(self, s, pos, len)
}

func (self *sys__net__SocketInput) read(nbytes int) *haxe__io__Bytes {
	return haxe__io__input_read(self, nbytes)
}

func (self *sys__net__SocketInput) readUntil(end int) *string {
	return haxe__io__input_readUntil(self, end)
}

func (self *sys__net__SocketInput) readLine() *string {
	return haxe__io__input_readLine(self)
}

func (self *sys__net__SocketInput) readFloat() float64 {
	return haxe__io__input_readFloat(self)
}

func (self *sys__net__SocketInput) readDouble() float64 {
	return haxe__io__input_readDouble(self)
}

func (self *sys__net__SocketInput) readInt8() int {
	return haxe__io__input_readInt8(self)
}

func (self *sys__net__SocketInput) readInt16() int {
	return haxe__io__input_readInt16(self)
}

func (self *sys__net__SocketInput) readUInt16() int {
	return haxe__io__input_readUInt16(self)
}

func (self *sys__net__SocketInput) readInt24() int {
	return haxe__io__input_readInt24(self)
}

func (self *sys__net__SocketInput) readUInt24() int {
	return haxe__io__input_readUInt24(self)
}

func (self *sys__net__SocketInput) readInt32() int {
	return haxe__io__input_readInt32(self)
}

func (self *sys__net__SocketInput) readString(len int, encoding ...*haxe__io__Encoding) *string {
	return haxe__io__input_readString(self, len, encoding...)
}

func sys__net__SocketInput_translateReadStatus(result *hxrt.SocketIOResult) {
	if result.Status == hxrt.SocketIOBlocked {
		hxrt.Throw(haxe__io__Error_Blocked)
	}
	if result.Status == hxrt.SocketIOEOF {
		hxrt.Throw(New_haxe__io__Eof())
	}
}
