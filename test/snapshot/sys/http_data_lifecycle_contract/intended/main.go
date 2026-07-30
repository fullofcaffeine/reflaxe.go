package main

import "snapshot/hxrt"

func main() {
	request := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,hello%20world"))
	output := New__Main__DataOutput()
	request.onStatus = func(status int) {
		hx_arr_1 := output.events
		hx_arr_1.Push(hxrt.StringConcatAny(hxrt.StringFromLiteral("status:"), status))
	}
	request.onError = func(message *string) {
		hx_arr_2 := output.events
		hx_arr_2.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error:"), message))
	}
	request.__hx_this.customRequest(false, output.haxe__io__BytesOutput.haxe__io__Output, nil, nil)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("events="), hxrt.StringJoinAny(output.events.Values(), hxrt.StringFromLiteral(">"))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("body="), output.getBytes().__hx_this.toString()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("closeCount="), output.closeCount))
	hxrt.Println(v_2)
}

type I__Main__DataOutput interface {
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
	prepare(size int)
	writeInput(input *haxe__io__Input, bufsize any)
	writeString(value *string, encoding *haxe__io__Encoding)
	get_length() int
	getBytes() *haxe__io__Bytes
}

type _Main__DataOutput struct {
	*haxe__io__BytesOutput
	__hx_this  I__Main__DataOutput
	events     *hxrt.Array
	closeCount int
}

func New__Main__DataOutput() *_Main__DataOutput {
	self := &_Main__DataOutput{}
	self.haxe__io__BytesOutput = New_haxe__io__BytesOutput()
	self.haxe__io__BytesOutput.haxe__io__Output.__hx_this = self
	self.haxe__io__BytesOutput.__hx_this = self
	self.__hx_this = self
	self.closeCount = 0
	self.events = hxrt.NewArray()
	return self
}

func (self *_Main__DataOutput) prepare(size int) {
	hx_arr_3 := self.events
	hx_arr_3.Push(hxrt.StringConcatAny(hxrt.StringFromLiteral("prepare:"), size))
	self.haxe__io__Output.prepare(size)
}

func (self *_Main__DataOutput) writeBytes(bytes *haxe__io__Bytes, pos int, len int) int {
	hx_arr_4 := self.events
	hx_arr_4.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("write:"), bytes.__hx_this.sub(pos, len).__hx_this.toString()))
	return self.haxe__io__BytesOutput.writeBytes(bytes, pos, len)
}

func (self *_Main__DataOutput) close() {
	self.closeCount = int(int32((self.closeCount + 1)))
	hx_arr_5 := self.events
	hx_arr_5.Push(hxrt.StringFromLiteral("close"))
	self.haxe__io__Output.close()
}
