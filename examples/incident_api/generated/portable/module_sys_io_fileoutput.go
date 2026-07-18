package main

import "examples_incident_api_portable/hxrt"

type I_sys__io__FileOutput interface {
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
	seek(p int, pos *sys__io__FileSeek)
	tell() int
}

type sys__io__FileOutput struct {
	*haxe__io__Output
	__hx_this I_sys__io__FileOutput
	handle    *hxrt.FileOutput
}

func New_sys__io__FileOutput(handle *hxrt.FileOutput) *sys__io__FileOutput {
	self := &sys__io__FileOutput{}
	self.haxe__io__Output = New_haxe__io__Output()
	self.haxe__io__Output.__hx_this = self
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
		hx_post_131 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_131
		values.Push(bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))])
	}
	return hxrt.FileOutputWriteValues(self.handle, func(hx_lambda_raw_133 []any) []int {
		hx_lambda_out_134 := make([]int, 0, len(hx_lambda_raw_133))
		for _, hx_lambda_item_135 := range hx_lambda_raw_133 {
			hx_lambda_out_134 = append(hx_lambda_out_134, func(hx_value_136 any) int {
				if hx_value_136 == nil {
					var hx_zero_137 int
					return hx_zero_137
				}
				return hx_value_136.(int)
			}(hx_lambda_item_135))
		}
		return hx_lambda_out_134
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
