package main

import "snapshot/hxrt"

type I_sys__io__FileOutput interface {
	writeByte(value int)
	writeBytes(bytes *haxe__io__Bytes, pos int, length int) int
	seek(p int, pos *sys__io__FileSeek)
	tell() int
	flush()
	close()
}

type sys__io__FileOutput struct {
	__hx_this         I_sys__io__FileOutput
	handle            *hxrt.FileOutput
	__hx_io_bigEndian bool
}

func New_sys__io__FileOutput(handle *hxrt.FileOutput) *sys__io__FileOutput {
	self := &sys__io__FileOutput{}
	self.__hx_this = self
	self.handle = handle
	return self
}

func (self *sys__io__FileOutput) writeByte(value int) {
	hxrt.FileOutputWriteByteValue(self.handle, value)
}

func (self *sys__io__FileOutput) writeBytes(bytes *haxe__io__Bytes, pos int, length int) int {
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
		hx_post_15 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_15
		values.Push(bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))])
	}
	return hxrt.FileOutputWriteValues(self.handle, func(hx_lambda_raw_17 []any) []int {
		hx_lambda_out_18 := make([]int, 0, len(hx_lambda_raw_17))
		for _, hx_lambda_item_19 := range hx_lambda_raw_17 {
			hx_lambda_out_18 = append(hx_lambda_out_18, func(hx_value_20 any) int {
				if hx_value_20 == nil {
					var hx_zero_21 int
					return hx_zero_21
				}
				return hx_value_20.(int)
			}(hx_lambda_item_19))
		}
		return hx_lambda_out_18
	}(values.Values()), 0, length)
}

func (self *sys__io__FileOutput) seek(p int, pos *sys__io__FileSeek) {
	switch pos.tag {
	case 0:
		hxrt.FileOutputSeek(self.handle, p, 0)
	case 1:
		hxrt.FileOutputSeek(self.handle, p, 1)
	case 2:
		hxrt.FileOutputSeek(self.handle, p, 2)
	}
}

func (self *sys__io__FileOutput) tell() int {
	return hxrt.FileOutputTell(self.handle)
}

func (self *sys__io__FileOutput) flush() {
	hxrt.FileOutputFlush(self.handle)
}

func (self *sys__io__FileOutput) close() {
	hxrt.FileOutputClose(self.handle)
}

func (self *sys__io__FileOutput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.__hx_io_bigEndian
}

func (self *sys__io__FileOutput) set_bigEndian(e bool) bool {
	if self != nil {
		self.__hx_io_bigEndian = e
	}
	return e
}

func (self *sys__io__FileOutput) prepare(nbytes int) {
	_ = self
	_ = nbytes
}

func (self *sys__io__FileOutput) write(s *haxe__io__Bytes) {
	haxe__io__output_write(self, s)
}

func (self *sys__io__FileOutput) writeFullBytes(s *haxe__io__Bytes, pos int, len int) {
	haxe__io__output_writeFullBytes(self, s, pos, len)
}

func (self *sys__io__FileOutput) writeFloat(x float64) {
	haxe__io__output_writeFloat(self, x)
}

func (self *sys__io__FileOutput) writeDouble(x float64) {
	haxe__io__output_writeDouble(self, x)
}

func (self *sys__io__FileOutput) writeInt8(x int) {
	haxe__io__output_writeInt8(self, x)
}

func (self *sys__io__FileOutput) writeInt16(x int) {
	haxe__io__output_writeInt16(self, x)
}

func (self *sys__io__FileOutput) writeUInt16(x int) {
	haxe__io__output_writeUInt16(self, x)
}

func (self *sys__io__FileOutput) writeInt24(x int) {
	haxe__io__output_writeInt24(self, x)
}

func (self *sys__io__FileOutput) writeUInt24(x int) {
	haxe__io__output_writeUInt24(self, x)
}

func (self *sys__io__FileOutput) writeInt32(x int) {
	haxe__io__output_writeInt32(self, x)
}

func (self *sys__io__FileOutput) writeInput(i haxe__io__Input, bufsize ...int) {
	haxe__io__output_writeInput(self, i, bufsize...)
}

func (self *sys__io__FileOutput) writeString(s *string, encoding ...*haxe__io__Encoding) {
	haxe__io__output_writeString(self, s, encoding...)
}
