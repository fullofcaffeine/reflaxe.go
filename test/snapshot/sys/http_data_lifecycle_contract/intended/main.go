package main

import "snapshot/hxrt"

func main() {
	request := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,hello%20world"))
	output := New__Main__DataOutput()
	request.onStatus = hxrt.NewFunction(func(status int) {
		hx_arr_1 := output.events
		hx_arr_1.Push(hxrt.StringConcatAny(hxrt.StringFromLiteral("status:"), status))
	})
	request.onError = hxrt.NewFunction(func(message *string) {
		hx_arr_2 := output.events
		hx_arr_2.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error:"), message))
	})
	request.__hx_this.customRequest(false, func(hx_value *_Main__DataOutput) *haxe__io__Output {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__BytesOutput.haxe__io__Output
	}(output), nil, nil)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("events="), hxrt.StringJoinAny(output.events.Values(), hxrt.StringFromLiteral(">"))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("body="), output.getBytes().__hx_this.toString()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("closeCount="), output.closeCount))
	hxrt.Println(v_2)
	status := 101
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("status."), status), hxrt.StringFromLiteral("=")), sys__Http_hxrt_statusError(status)))
	hxrt.Println(v_3)
	status_1 := 199
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("status."), status_1), hxrt.StringFromLiteral("=")), sys__Http_hxrt_statusError(status_1)))
	hxrt.Println(v_4)
	status_2 := 200
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("status."), status_2), hxrt.StringFromLiteral("=")), sys__Http_hxrt_statusError(status_2)))
	hxrt.Println(v_5)
	status_3 := 399
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("status."), status_3), hxrt.StringFromLiteral("=")), sys__Http_hxrt_statusError(status_3)))
	hxrt.Println(v_6)
	status_4 := 400
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("status."), status_4), hxrt.StringFromLiteral("=")), sys__Http_hxrt_statusError(status_4)))
	hxrt.Println(v_7)
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

func (self *_Main__DataOutput) HxFunctionReceiver() any {
	return self.__hx_this
}

func New__Main__DataOutput() *_Main__DataOutput {
	return hx_init__Main__DataOutput(hx_alloc__Main__DataOutput(nil))
}

func hx_alloc__Main__DataOutput(owner any) *_Main__DataOutput {
	self := &_Main__DataOutput{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I__Main__DataOutput)
	}
	self.haxe__io__BytesOutput = hx_alloc_haxe__io__BytesOutput(self.__hx_this)
	return self
}

func hx_init__Main__DataOutput(self *_Main__DataOutput) *_Main__DataOutput {
	hx_init_haxe__io__BytesOutput(self.haxe__io__BytesOutput)
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

func (self *_Main__DataOutput) HxClassName() string {
	return "_Main.DataOutput"
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *_Main__DataOutput:
		return func(hx_structural_source_1 *_Main__DataOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "events":
					return hxrt.AnonymousField{Value: hx_structural_source_1.events, Found: true}
				case "closeCount":
					return hxrt.AnonymousField{Value: hx_structural_source_1.closeCount, Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_2 := hx_structural_source_1
						return hxrt.BoundFunction(func(size int) {
							hx_receiver_2.__hx_this.prepare(size)
						}, func() any {
							return hx_receiver_2.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_3 := hx_structural_source_1
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_3.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_3.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_4 := hx_structural_source_1
						return hxrt.BoundFunction(func() {
							hx_receiver_4.__hx_this.close()
						}, func() any {
							return hx_receiver_4.__hx_this
						}, "close")
					}(), Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_1.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_5 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_5.__hx_this.get_length()
						}, func() any {
							return hx_receiver_5.__hx_this
						}, "get_length")
					}(), Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_6 := hx_structural_source_1
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_6.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_6.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_7 := hx_structural_source_1
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_7.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_7.__hx_this
						}, "getBytes")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_1.bigEndian, Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_8 := hx_structural_source_1
						return hxrt.BoundFunction(func() {
							hx_receiver_8.__hx_this.flush()
						}, func() any {
							return hx_receiver_8.__hx_this
						}, "flush")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_9 := hx_structural_source_1
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_9.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_9.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_10 := hx_structural_source_1
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_10.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_10.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_11 := hx_structural_source_1
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_11.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_12 := hx_structural_source_1
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_12.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_13 := hx_structural_source_1
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_13.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_13.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_14 := hx_structural_source_1
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_14.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_14.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_15 := hx_structural_source_1
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_15.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_15.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_16 := hx_structural_source_1
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_16.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_16.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_17 := hx_structural_source_1
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_17.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_18 := hx_structural_source_1
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_18.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_19 := hx_structural_source_1
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_19.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_20 := hx_structural_source_1
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_20.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_20.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_21 := hx_structural_source_1
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_21.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_1.b = hx_recover_class_haxe__io__BytesBuffer(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_22 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_22, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_22.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_22.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_22.high = func(hx_value_23 any) int {
						if hx_value_23 == nil {
							var hx_zero_24 int
							return hx_zero_24
						}
						return hx_value_23.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_22.low = func(hx_value_25 any) int {
						if hx_value_25 == nil {
							var hx_zero_26 int
							return hx_zero_26
						}
						return hx_value_25.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketOutput:
		return func(hx_structural_source_27 *sys__net__SocketOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_27, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_27.handle, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_30 := hx_structural_source_27
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_30.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_30.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_31 := hx_structural_source_27
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_31.__hx_this.writeBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_32 := hx_structural_source_27
						return hxrt.BoundFunction(func() {
							hx_receiver_32.__hx_this.flush()
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_33 := hx_structural_source_27
						return hxrt.BoundFunction(func() {
							hx_receiver_33.__hx_this.close()
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_27.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_34 := hx_structural_source_27
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_34.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_35 := hx_structural_source_27
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_35.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_36 := hx_structural_source_27
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_36.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_36.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_37 := hx_structural_source_27
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_37.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_37.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_38 := hx_structural_source_27
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_38.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_38.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_39 := hx_structural_source_27
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_39.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_39.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_40 := hx_structural_source_27
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_40.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_40.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_41 := hx_structural_source_27
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_41.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_41.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_42 := hx_structural_source_27
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_42.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_42.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_43 := hx_structural_source_27
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_43.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_43.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_44 := hx_structural_source_27
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_44.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_44.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_45 := hx_structural_source_27
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_45.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_45.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_46 := hx_structural_source_27
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_46.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_46.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_47 := hx_structural_source_27
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_47.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_47.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_27.handle = func(hx_value_28 any) *hxrt.SocketHandle {
						if hx_value_28 == nil {
							var hx_zero_29 *hxrt.SocketHandle
							return hx_zero_29
						}
						return hx_value_28.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketInput:
		return func(hx_structural_source_48 *sys__net__SocketInput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_48, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_48.handle, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_51 := hx_structural_source_48
						return hxrt.BoundFunction(func() int {
							return hx_receiver_51.__hx_this.readByte()
						}, func() any {
							return hx_receiver_51.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_52 := hx_structural_source_48
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_52.__hx_this.readBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_52.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_53 := hx_structural_source_48
						return hxrt.BoundFunction(func() {
							hx_receiver_53.__hx_this.close()
						}, func() any {
							return hx_receiver_53.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_48.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_54 := hx_structural_source_48
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_54.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_54.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_55 := hx_structural_source_48
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_55.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_55.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_56 := hx_structural_source_48
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_56.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_56.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_57 := hx_structural_source_48
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_57.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_57.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_58 := hx_structural_source_48
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_58.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_58.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_59 := hx_structural_source_48
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_59.__hx_this.readLine()
						}, func() any {
							return hx_receiver_59.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_60 := hx_structural_source_48
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_60.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_60.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_61 := hx_structural_source_48
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_61.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_61.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_62 := hx_structural_source_48
						return hxrt.BoundFunction(func() int {
							return hx_receiver_62.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_62.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_63 := hx_structural_source_48
						return hxrt.BoundFunction(func() int {
							return hx_receiver_63.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_63.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_64 := hx_structural_source_48
						return hxrt.BoundFunction(func() int {
							return hx_receiver_64.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_64.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_65 := hx_structural_source_48
						return hxrt.BoundFunction(func() int {
							return hx_receiver_65.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_65.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_66 := hx_structural_source_48
						return hxrt.BoundFunction(func() int {
							return hx_receiver_66.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_66.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_67 := hx_structural_source_48
						return hxrt.BoundFunction(func() int {
							return hx_receiver_67.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_67.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_68 := hx_structural_source_48
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_68.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_68.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_48.handle = func(hx_value_49 any) *hxrt.SocketHandle {
						if hx_value_49 == nil {
							var hx_zero_50 *hxrt.SocketHandle
							return hx_zero_50
						}
						return hx_value_49.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Socket:
		return func(hx_structural_source_69 *sys__net__Socket) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_69, func(name string) hxrt.AnonymousField {
				switch name {
				case "input":
					return hxrt.AnonymousField{Value: hx_structural_source_69.input, Found: true}
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_69.output, Found: true}
				case "custom":
					return hxrt.AnonymousField{Value: hx_structural_source_69.custom, Found: true}
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_69.handle, Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_72 := hx_structural_source_69
						return hxrt.BoundFunction(func() {
							hx_receiver_72.__hx_this.close()
						}, func() any {
							return hx_receiver_72.__hx_this
						}, "close")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_73 := hx_structural_source_69
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_73.__hx_this.read()
						}, func() any {
							return hx_receiver_73.__hx_this
						}, "read")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_74 := hx_structural_source_69
						return hxrt.BoundFunction(func(content *string) {
							hx_receiver_74.__hx_this.write(content)
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "write")
					}(), Found: true}
				case "connect":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_75 := hx_structural_source_69
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_75.__hx_this.connect(host, port)
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "connect")
					}(), Found: true}
				case "listen":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_76 := hx_structural_source_69
						return hxrt.BoundFunction(func(connections int) {
							hx_receiver_76.__hx_this.listen(connections)
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "listen")
					}(), Found: true}
				case "shutdown":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, bool)] {
						hx_receiver_77 := hx_structural_source_69
						return hxrt.BoundFunction(func(read bool, write bool) {
							hx_receiver_77.__hx_this.shutdown(read, write)
						}, func() any {
							return hx_receiver_77.__hx_this
						}, "shutdown")
					}(), Found: true}
				case "bind":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_78 := hx_structural_source_69
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_78.__hx_this.bind(host, port)
						}, func() any {
							return hx_receiver_78.__hx_this
						}, "bind")
					}(), Found: true}
				case "accept":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *sys__net__Socket] {
						hx_receiver_79 := hx_structural_source_69
						return hxrt.BoundFunction(func() *sys__net__Socket {
							return hx_receiver_79.__hx_this.accept()
						}, func() any {
							return hx_receiver_79.__hx_this
						}, "accept")
					}(), Found: true}
				case "peer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_80 := hx_structural_source_69
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_80.__hx_this.peer()
						}, func() any {
							return hx_receiver_80.__hx_this
						}, "peer")
					}(), Found: true}
				case "host":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_81 := hx_structural_source_69
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_81.__hx_this.host()
						}, func() any {
							return hx_receiver_81.__hx_this
						}, "host")
					}(), Found: true}
				case "setTimeout":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_82 := hx_structural_source_69
						return hxrt.BoundFunction(func(timeout float64) {
							hx_receiver_82.__hx_this.setTimeout(timeout)
						}, func() any {
							return hx_receiver_82.__hx_this
						}, "setTimeout")
					}(), Found: true}
				case "waitForRead":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_83 := hx_structural_source_69
						return hxrt.BoundFunction(func() {
							hx_receiver_83.__hx_this.waitForRead()
						}, func() any {
							return hx_receiver_83.__hx_this
						}, "waitForRead")
					}(), Found: true}
				case "setBlocking":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_84 := hx_structural_source_69
						return hxrt.BoundFunction(func(blocking bool) {
							hx_receiver_84.__hx_this.setBlocking(blocking)
						}, func() any {
							return hx_receiver_84.__hx_this
						}, "setBlocking")
					}(), Found: true}
				case "setFastSend":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_85 := hx_structural_source_69
						return hxrt.BoundFunction(func(fastSend bool) {
							hx_receiver_85.__hx_this.setFastSend(fastSend)
						}, func() any {
							return hx_receiver_85.__hx_this
						}, "setFastSend")
					}(), Found: true}
				case "replaceHandle":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.SocketHandle)] {
						hx_receiver_86 := hx_structural_source_69
						return hxrt.BoundFunction(func(next *hxrt.SocketHandle) {
							hx_receiver_86.__hx_this.replaceHandle(next)
						}, func() any {
							return hx_receiver_86.__hx_this
						}, "replaceHandle")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "custom":
					hx_structural_source_69.custom = incoming
					return true
				case "handle":
					hx_structural_source_69.handle = func(hx_value_70 any) *hxrt.SocketHandle {
						if hx_value_70 == nil {
							var hx_zero_71 *hxrt.SocketHandle
							return hx_zero_71
						}
						return hx_value_70.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Host:
		return func(hx_structural_source_87 *sys__net__Host) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_87, func(name string) hxrt.AnonymousField {
				switch name {
				case "host":
					return hxrt.AnonymousField{Value: hx_structural_source_87.host, Found: true}
				case "ip":
					return hxrt.AnonymousField{Value: hx_structural_source_87.ip, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_88 := hx_structural_source_87
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_88.__hx_this.toString()
						}, func() any {
							return hx_receiver_88.__hx_this
						}, "toString")
					}(), Found: true}
				case "reverse":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_89 := hx_structural_source_87
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_89.__hx_this.reverse()
						}, func() any {
							return hx_receiver_89.__hx_this
						}, "reverse")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *sys__Http:
		return func(hx_structural_source_90 *sys__Http) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_90, func(name string) hxrt.AnonymousField {
				switch name {
				case "noShutdown":
					return hxrt.AnonymousField{Value: hx_structural_source_90.noShutdown, Found: true}
				case "cnxTimeout":
					return hxrt.AnonymousField{Value: hx_structural_source_90.cnxTimeout, Found: true}
				case "responseHeaders":
					return hxrt.AnonymousField{Value: hx_structural_source_90.responseHeaders, Found: true}
				case "responseHeadersSameKey":
					return hxrt.AnonymousField{Value: hx_structural_source_90.responseHeadersSameKey, Found: true}
				case "file":
					return hxrt.AnonymousField{Value: hx_structural_source_90.file, Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_97 := hx_structural_source_90
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_97.__hx_this.request(post)
						}, func() any {
							return hx_receiver_97.__hx_this
						}, "request")
					}(), Found: true}
				case "fileTransfert":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_98 := hx_structural_source_90
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_98.__hx_this.fileTransfert(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_98.__hx_this
						}, "fileTransfert")
					}(), Found: true}
				case "fileTransfer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_99 := hx_structural_source_90
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_99.__hx_this.fileTransfer(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_99.__hx_this
						}, "fileTransfer")
					}(), Found: true}
				case "customRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_100 := hx_structural_source_90
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_100.__hx_this.customRequest(post, api, sock, method)
						}, func() any {
							return hx_receiver_100.__hx_this
						}, "customRequest")
					}(), Found: true}
				case "getResponseHeaderValues":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_101 := hx_structural_source_90
						return hxrt.BoundFunction(func(key *string) *hxrt.Array {
							return hx_receiver_101.__hx_this.getResponseHeaderValues(key)
						}, func() any {
							return hx_receiver_101.__hx_this
						}, "getResponseHeaderValues")
					}(), Found: true}
				case "requestWith":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_102 := hx_structural_source_90
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_102.__hx_this.requestWith(post, api, sock, method)
						}, func() any {
							return hx_receiver_102.__hx_this
						}, "requestWith")
					}(), Found: true}
				case "pumpUpload":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange, hxrt.Anonymous) hxrt.Anonymous] {
						hx_receiver_103 := hx_structural_source_90
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange, upload hxrt.Anonymous) hxrt.Anonymous {
							return hx_receiver_103.__hx_this.pumpUpload(exchange, upload)
						}, func() any {
							return hx_receiver_103.__hx_this
						}, "pumpUpload")
					}(), Found: true}
				case "handleDataRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *string)] {
						hx_receiver_104 := hx_structural_source_90
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, method *string) {
							hx_receiver_104.__hx_this.handleDataRequest(post, api, method)
						}, func() any {
							return hx_receiver_104.__hx_this
						}, "handleDataRequest")
					}(), Found: true}
				case "recordResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange)] {
						hx_receiver_105 := hx_structural_source_90
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange) {
							hx_receiver_105.__hx_this.recordResponseHeaders(exchange)
						}, func() any {
							return hx_receiver_105.__hx_this
						}, "recordResponseHeaders")
					}(), Found: true}
				case "resetResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_106 := hx_structural_source_90
						return hxrt.BoundFunction(func() {
							hx_receiver_106.__hx_this.resetResponseHeaders()
						}, func() any {
							return hx_receiver_106.__hx_this
						}, "resetResponseHeaders")
					}(), Found: true}
				case "encodedParameters":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_107 := hx_structural_source_90
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_107.__hx_this.encodedParameters()
						}, func() any {
							return hx_receiver_107.__hx_this
						}, "encodedParameters")
					}(), Found: true}
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_90.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_90.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_90.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_90.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_90.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_90.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_90.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_90.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_118 := hx_structural_source_90
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_118.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_118.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_119 := hx_structural_source_90
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_119.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_119.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_120 := hx_structural_source_90
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_120.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_120.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_121 := hx_structural_source_90
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_121.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_121.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_122 := hx_structural_source_90
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_122.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_122.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_123 := hx_structural_source_90
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_123.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_123.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_90.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_90.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_90.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_90.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_124 := hx_structural_source_90
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_124.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_124.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_125 := hx_structural_source_90
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_125.__hx_this.success(data)
						}, func() any {
							return hx_receiver_125.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_126 := hx_structural_source_90
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_126.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_126.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "noShutdown":
					hx_structural_source_90.noShutdown = func(hx_value_91 any) bool {
						if hx_value_91 == nil {
							var hx_zero_92 bool
							return hx_zero_92
						}
						return hx_value_91.(bool)
					}(incoming)
					return true
				case "cnxTimeout":
					hx_structural_source_90.cnxTimeout = func(hx_value_93 any) float64 {
						if hx_value_93 == nil {
							var hx_zero_94 float64
							return hx_zero_94
						}
						return hx_value_93.(float64)
					}(incoming)
					return true
				case "responseHeaders":
					hx_structural_source_90.responseHeaders = func(hx_value_95 any) *haxe__ds__StringMap {
						if hx_value_95 == nil {
							var hx_zero_96 *haxe__ds__StringMap
							return hx_zero_96
						}
						return hx_value_95.(*haxe__ds__StringMap)
					}(incoming)
					return true
				case "responseHeadersSameKey":
					hx_structural_source_90.responseHeadersSameKey = hx_recover_class_haxe__ds__StringMap(incoming)
					return true
				case "file":
					hx_structural_source_90.file = hx_structural_recover(incoming)
					return true
				case "url":
					hx_structural_source_90.url = func(hx_value_108 any) *string {
						if hx_value_108 == nil {
							var hx_zero_109 *string
							return hx_zero_109
						}
						return hx_value_108.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_90.responseAsString = func(hx_value_110 any) *string {
						if hx_value_110 == nil {
							var hx_zero_111 *string
							return hx_zero_111
						}
						return hx_value_110.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_90.postData = func(hx_value_112 any) *string {
						if hx_value_112 == nil {
							var hx_zero_113 *string
							return hx_zero_113
						}
						return hx_value_112.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_90.postBytes = hx_recover_class_haxe__io__Bytes(incoming)
					return true
				case "headers":
					hx_structural_source_90.headers = func(hx_value_114 any) *hxrt.Array {
						if hx_value_114 == nil {
							var hx_zero_115 *hxrt.Array
							return hx_zero_115
						}
						return hx_value_114.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_90.params = func(hx_value_116 any) *hxrt.Array {
						if hx_value_116 == nil {
							var hx_zero_117 *hxrt.Array
							return hx_zero_117
						}
						return hx_value_116.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_90.onData = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onBytes":
					hx_structural_source_90.onBytes = hx_recover_callable_c5f1b802(incoming)
					return true
				case "onError":
					hx_structural_source_90.onError = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onStatus":
					hx_structural_source_90.onStatus = hx_recover_callable_d8be58e9(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Output:
		return func(hx_structural_source_127 *haxe__io__Output) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_127, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_127.bigEndian, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_128 := hx_structural_source_127
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_128.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_128.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_129 := hx_structural_source_127
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_129.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_129.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_130 := hx_structural_source_127
						return hxrt.BoundFunction(func() {
							hx_receiver_130.__hx_this.flush()
						}, func() any {
							return hx_receiver_130.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_131 := hx_structural_source_127
						return hxrt.BoundFunction(func() {
							hx_receiver_131.__hx_this.close()
						}, func() any {
							return hx_receiver_131.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_132 := hx_structural_source_127
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_132.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_132.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_133 := hx_structural_source_127
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_133.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_133.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_134 := hx_structural_source_127
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_134.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_134.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_135 := hx_structural_source_127
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_135.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_135.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_136 := hx_structural_source_127
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_136.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_136.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_137 := hx_structural_source_127
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_137.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_137.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_138 := hx_structural_source_127
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_138.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_138.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_139 := hx_structural_source_127
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_139.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_139.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_140 := hx_structural_source_127
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_140.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_140.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_141 := hx_structural_source_127
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_141.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_141.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_142 := hx_structural_source_127
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_142.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_142.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_143 := hx_structural_source_127
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_143.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_143.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_144 := hx_structural_source_127
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_144.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_144.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_145 := hx_structural_source_127
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_145.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_145.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *haxe__io__Input:
		return func(hx_structural_source_146 *haxe__io__Input) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_146, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_146.bigEndian, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_147 := hx_structural_source_146
						return hxrt.BoundFunction(func() int {
							return hx_receiver_147.__hx_this.readByte()
						}, func() any {
							return hx_receiver_147.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_148 := hx_structural_source_146
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_148.__hx_this.readBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_148.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_149 := hx_structural_source_146
						return hxrt.BoundFunction(func() {
							hx_receiver_149.__hx_this.close()
						}, func() any {
							return hx_receiver_149.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_150 := hx_structural_source_146
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_150.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_150.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_151 := hx_structural_source_146
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_151.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_151.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_152 := hx_structural_source_146
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_152.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_152.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_153 := hx_structural_source_146
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_153.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_153.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_154 := hx_structural_source_146
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_154.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_154.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_155 := hx_structural_source_146
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_155.__hx_this.readLine()
						}, func() any {
							return hx_receiver_155.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_156 := hx_structural_source_146
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_156.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_156.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_157 := hx_structural_source_146
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_157.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_157.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_158 := hx_structural_source_146
						return hxrt.BoundFunction(func() int {
							return hx_receiver_158.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_158.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_159 := hx_structural_source_146
						return hxrt.BoundFunction(func() int {
							return hx_receiver_159.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_159.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_160 := hx_structural_source_146
						return hxrt.BoundFunction(func() int {
							return hx_receiver_160.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_160.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_161 := hx_structural_source_146
						return hxrt.BoundFunction(func() int {
							return hx_receiver_161.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_161.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_162 := hx_structural_source_146
						return hxrt.BoundFunction(func() int {
							return hx_receiver_162.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_162.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_163 := hx_structural_source_146
						return hxrt.BoundFunction(func() int {
							return hx_receiver_163.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_163.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_164 := hx_structural_source_146
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_164.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_164.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *haxe__io__Eof:
		return func(hx_structural_source_165 *haxe__io__Eof) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_165, func(name string) hxrt.AnonymousField {
				switch name {
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_166 := hx_structural_source_165
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_166.__hx_this.toString()
						}, func() any {
							return hx_receiver_166.__hx_this
						}, "toString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *haxe__io__BytesOutput:
		return func(hx_structural_source_167 *haxe__io__BytesOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_167, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_167.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_168 := hx_structural_source_167
						return hxrt.BoundFunction(func() int {
							return hx_receiver_168.__hx_this.get_length()
						}, func() any {
							return hx_receiver_168.__hx_this
						}, "get_length")
					}(), Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_169 := hx_structural_source_167
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_169.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_169.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_170 := hx_structural_source_167
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_170.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_170.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_171 := hx_structural_source_167
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_171.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_171.__hx_this
						}, "getBytes")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_167.bigEndian, Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_172 := hx_structural_source_167
						return hxrt.BoundFunction(func() {
							hx_receiver_172.__hx_this.flush()
						}, func() any {
							return hx_receiver_172.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_173 := hx_structural_source_167
						return hxrt.BoundFunction(func() {
							hx_receiver_173.__hx_this.close()
						}, func() any {
							return hx_receiver_173.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_174 := hx_structural_source_167
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_174.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_174.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_175 := hx_structural_source_167
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_175.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_175.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_176 := hx_structural_source_167
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_176.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_176.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_177 := hx_structural_source_167
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_177.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_177.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_178 := hx_structural_source_167
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_178.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_178.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_179 := hx_structural_source_167
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_179.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_179.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_180 := hx_structural_source_167
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_180.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_180.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_181 := hx_structural_source_167
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_181.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_181.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_182 := hx_structural_source_167
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_182.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_182.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_183 := hx_structural_source_167
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_183.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_183.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_184 := hx_structural_source_167
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_184.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_184.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_185 := hx_structural_source_167
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_185.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_185.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_186 := hx_structural_source_167
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_186.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_186.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_187 := hx_structural_source_167
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_187.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_187.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_167.b = hx_recover_class_haxe__io__BytesBuffer(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__BytesBuffer:
		return func(hx_structural_source_188 *haxe__io__BytesBuffer) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_188, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_188.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_191 := hx_structural_source_188
						return hxrt.BoundFunction(func() int {
							return hx_receiver_191.__hx_this.get_length()
						}, func() any {
							return hx_receiver_191.__hx_this
						}, "get_length")
					}(), Found: true}
				case "addByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_192 := hx_structural_source_188
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_192.__hx_this.addByte(value)
						}, func() any {
							return hx_receiver_192.__hx_this
						}, "addByte")
					}(), Found: true}
				case "add":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_193 := hx_structural_source_188
						return hxrt.BoundFunction(func(source *haxe__io__Bytes) {
							hx_receiver_193.__hx_this.add(source)
						}, func() any {
							return hx_receiver_193.__hx_this
						}, "add")
					}(), Found: true}
				case "addString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_194 := hx_structural_source_188
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_194.__hx_this.addString(value, encoding)
						}, func() any {
							return hx_receiver_194.__hx_this
						}, "addString")
					}(), Found: true}
				case "addInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_195 := hx_structural_source_188
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_195.__hx_this.addInt32(value)
						}, func() any {
							return hx_receiver_195.__hx_this
						}, "addInt32")
					}(), Found: true}
				case "addInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Int64_____Int64)] {
						hx_receiver_196 := hx_structural_source_188
						return hxrt.BoundFunction(func(value *haxe___Int64_____Int64) {
							hx_receiver_196.__hx_this.addInt64(value)
						}, func() any {
							return hx_receiver_196.__hx_this
						}, "addInt64")
					}(), Found: true}
				case "addFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_197 := hx_structural_source_188
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_197.__hx_this.addFloat(value)
						}, func() any {
							return hx_receiver_197.__hx_this
						}, "addFloat")
					}(), Found: true}
				case "addDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_198 := hx_structural_source_188
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_198.__hx_this.addDouble(value)
						}, func() any {
							return hx_receiver_198.__hx_this
						}, "addDouble")
					}(), Found: true}
				case "addBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_199 := hx_structural_source_188
						return hxrt.BoundFunction(func(source *haxe__io__Bytes, pos int, len int) {
							hx_receiver_199.__hx_this.addBytes(source, pos, len)
						}, func() any {
							return hx_receiver_199.__hx_this
						}, "addBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_200 := hx_structural_source_188
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_200.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_200.__hx_this
						}, "getBytes")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_188.b = func(hx_value_189 any) []int {
						if hx_value_189 == nil {
							var hx_zero_190 []int
							return hx_zero_190
						}
						return hx_value_189.([]int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_201 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_201, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_201.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_201.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_201.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_201.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_201.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_210 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_210.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_210.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_211 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_211.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_211.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_212 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
							hx_receiver_212.__hx_this.blit(pos, src, srcpos, len)
						}, func() any {
							return hx_receiver_212.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_213 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, len int, value int) {
							hx_receiver_213.__hx_this.fill(pos, len, value)
						}, func() any {
							return hx_receiver_213.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_214 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
							return hx_receiver_214.__hx_this.sub(pos, len)
						}, func() any {
							return hx_receiver_214.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_215 := hx_structural_source_201
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_215.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_215.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_216 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_216.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_216.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_217 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_217.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_217.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_218 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_218.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_218.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_219 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_219.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_219.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_220 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_220.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_220.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_221 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_221.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_221.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_222 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_222.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_222.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_223 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_223.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_223.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_224 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_224.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_224.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_225 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_225.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_225.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_226 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_226.__hx_this.getString(pos, len, encoding)
						}, func() any {
							return hx_receiver_226.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_227 := hx_structural_source_201
						return hxrt.BoundFunction(func(pos int, len int) *string {
							return hx_receiver_227.__hx_this.readString(pos, len)
						}, func() any {
							return hx_receiver_227.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_228 := hx_structural_source_201
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_228.__hx_this.toString()
						}, func() any {
							return hx_receiver_228.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_229 := hx_structural_source_201
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_229.__hx_this.toHex()
						}, func() any {
							return hx_receiver_229.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_230 := hx_structural_source_201
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_230.__hx_this.getData()
						}, func() any {
							return hx_receiver_230.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_231 := hx_structural_source_201
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_231.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_231.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_201.b = func(hx_value_202 any) []int {
						if hx_value_202 == nil {
							var hx_zero_203 []int
							return hx_zero_203
						}
						return hx_value_202.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_201.__hx_raw = func(hx_value_204 any) *hxrt.ByteView {
						if hx_value_204 == nil {
							var hx_zero_205 *hxrt.ByteView
							return hx_zero_205
						}
						return hx_value_204.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_201.__hx_rawValid = func(hx_value_206 any) bool {
						if hx_value_206 == nil {
							var hx_zero_207 bool
							return hx_zero_207
						}
						return hx_value_206.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_201.__hx_dataExposed = func(hx_value_208 any) bool {
						if hx_value_208 == nil {
							var hx_zero_209 bool
							return hx_zero_209
						}
						return hx_value_208.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__http__HttpBase:
		return func(hx_structural_source_232 *haxe__http__HttpBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_232, func(name string) hxrt.AnonymousField {
				switch name {
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_232.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_232.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_232.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_232.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_232.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_232.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_232.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_232.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_243 := hx_structural_source_232
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_243.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_243.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_244 := hx_structural_source_232
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_244.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_244.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_245 := hx_structural_source_232
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_245.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_245.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_246 := hx_structural_source_232
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_246.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_246.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_247 := hx_structural_source_232
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_247.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_247.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_248 := hx_structural_source_232
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_248.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_248.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_249 := hx_structural_source_232
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_249.__hx_this.request(post)
						}, func() any {
							return hx_receiver_249.__hx_this
						}, "request")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_232.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_232.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_232.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_232.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_250 := hx_structural_source_232
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_250.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_250.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_251 := hx_structural_source_232
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_251.__hx_this.success(data)
						}, func() any {
							return hx_receiver_251.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_252 := hx_structural_source_232
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_252.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_252.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "url":
					hx_structural_source_232.url = func(hx_value_233 any) *string {
						if hx_value_233 == nil {
							var hx_zero_234 *string
							return hx_zero_234
						}
						return hx_value_233.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_232.responseAsString = func(hx_value_235 any) *string {
						if hx_value_235 == nil {
							var hx_zero_236 *string
							return hx_zero_236
						}
						return hx_value_235.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_232.postData = func(hx_value_237 any) *string {
						if hx_value_237 == nil {
							var hx_zero_238 *string
							return hx_zero_238
						}
						return hx_value_237.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_232.postBytes = hx_recover_class_haxe__io__Bytes(incoming)
					return true
				case "headers":
					hx_structural_source_232.headers = func(hx_value_239 any) *hxrt.Array {
						if hx_value_239 == nil {
							var hx_zero_240 *hxrt.Array
							return hx_zero_240
						}
						return hx_value_239.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_232.params = func(hx_value_241 any) *hxrt.Array {
						if hx_value_241 == nil {
							var hx_zero_242 *hxrt.Array
							return hx_zero_242
						}
						return hx_value_241.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_232.onData = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onBytes":
					hx_structural_source_232.onBytes = hx_recover_callable_c5f1b802(incoming)
					return true
				case "onError":
					hx_structural_source_232.onError = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onStatus":
					hx_structural_source_232.onStatus = hx_recover_callable_d8be58e9(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__NotImplementedException:
		return func(hx_structural_source_253 *haxe__exceptions__NotImplementedException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_253, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_253.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_254 := hx_structural_source_253
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_254.__hx_this.toString()
						}, func() any {
							return hx_receiver_254.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_253.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_253.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_253.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_253.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_257 := hx_structural_source_253
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_257.__hx_this.get_message()
						}, func() any {
							return hx_receiver_257.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_258 := hx_structural_source_253
						return hxrt.BoundFunction(func() any {
							return hx_receiver_258.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_258.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_259 := hx_structural_source_253
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_259.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_259.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_260 := hx_structural_source_253
						return hxrt.BoundFunction(func() any {
							return hx_receiver_260.__hx_this.get_native()
						}, func() any {
							return hx_receiver_260.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_261 := hx_structural_source_253
						return hxrt.BoundFunction(func() any {
							return hx_receiver_261.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_261.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_262 := hx_structural_source_253
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_262.__hx_this.details()
						}, func() any {
							return hx_receiver_262.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_253.savedMessage = func(hx_value_255 any) *string {
						if hx_value_255 == nil {
							var hx_zero_256 *string
							return hx_zero_256
						}
						return hx_value_255.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_253.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_253.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_253.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__StringMap:
		return func(hx_structural_source_263 *haxe__ds__StringMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_263, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_263.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, any)] {
						hx_receiver_266 := hx_structural_source_263
						return hxrt.BoundFunction(func(key *string, value any) {
							hx_receiver_266.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_266.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_267 := hx_structural_source_263
						return hxrt.BoundFunction(func(key *string) any {
							return hx_receiver_267.__hx_this.get(key)
						}, func() any {
							return hx_receiver_267.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_268 := hx_structural_source_263
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_268.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_268.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_269 := hx_structural_source_263
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_269.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_269.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_270 := hx_structural_source_263
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_270.__hx_this.keys()
						}, func() any {
							return hx_receiver_270.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_271 := hx_structural_source_263
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_271.__hx_this.iterator()
						}, func() any {
							return hx_receiver_271.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_272 := hx_structural_source_263
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_272.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_272.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_273 := hx_structural_source_263
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_273.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_273.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_274 := hx_structural_source_263
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_274.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_274.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_275 := hx_structural_source_263
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_275.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_275.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_276 := hx_structural_source_263
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_276.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_276.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_277 := hx_structural_source_263
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_277.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_277.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__StringMap] {
						hx_receiver_278 := hx_structural_source_263
						return hxrt.BoundFunction(func() *haxe__ds__StringMap {
							return hx_receiver_278.__hx_this.copy()
						}, func() any {
							return hx_receiver_278.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_279 := hx_structural_source_263
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_279.__hx_this.toString()
						}, func() any {
							return hx_receiver_279.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_280 := hx_structural_source_263
						return hxrt.BoundFunction(func() {
							hx_receiver_280.__hx_this.clear()
						}, func() any {
							return hx_receiver_280.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_263.h = func(hx_value_264 any) *hxrt.StringMapCell {
						if hx_value_264 == nil {
							var hx_zero_265 *hxrt.StringMapCell
							return hx_zero_265
						}
						return hx_value_264.(*hxrt.StringMapCell)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_281 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_281, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_281.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_281.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_281.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_281.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_284 := hx_structural_source_281
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_284.__hx_this.get_message()
						}, func() any {
							return hx_receiver_284.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_285 := hx_structural_source_281
						return hxrt.BoundFunction(func() any {
							return hx_receiver_285.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_285.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_286 := hx_structural_source_281
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_286.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_286.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_287 := hx_structural_source_281
						return hxrt.BoundFunction(func() any {
							return hx_receiver_287.__hx_this.get_native()
						}, func() any {
							return hx_receiver_287.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_288 := hx_structural_source_281
						return hxrt.BoundFunction(func() any {
							return hx_receiver_288.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_288.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_289 := hx_structural_source_281
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_289.__hx_this.toString()
						}, func() any {
							return hx_receiver_289.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_290 := hx_structural_source_281
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_290.__hx_this.details()
						}, func() any {
							return hx_receiver_290.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_281.savedMessage = func(hx_value_282 any) *string {
						if hx_value_282 == nil {
							var hx_zero_283 *string
							return hx_zero_283
						}
						return hx_value_282.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_281.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_281.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_281.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_291 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_291, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_291.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_291.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_296 := hx_structural_source_291
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_296.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_296.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_297 := hx_structural_source_291
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_297.__hx_this.next()
						}, func() any {
							return hx_receiver_297.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_291.offset = func(hx_value_292 any) int {
						if hx_value_292 == nil {
							var hx_zero_293 int
							return hx_zero_293
						}
						return hx_value_292.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_291.s = func(hx_value_294 any) *string {
						if hx_value_294 == nil {
							var hx_zero_295 *string
							return hx_zero_295
						}
						return hx_value_294.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_298 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_298, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_298.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_298.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_303 := hx_structural_source_298
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_303.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_303.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_304 := hx_structural_source_298
						return hxrt.BoundFunction(func() int {
							return hx_receiver_304.__hx_this.next()
						}, func() any {
							return hx_receiver_304.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_298.offset = func(hx_value_299 any) int {
						if hx_value_299 == nil {
							var hx_zero_300 int
							return hx_zero_300
						}
						return hx_value_299.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_298.s = func(hx_value_301 any) *string {
						if hx_value_301 == nil {
							var hx_zero_302 *string
							return hx_zero_302
						}
						return hx_value_301.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_305 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_305, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_305.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_305.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_308 := hx_structural_source_305
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_308.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_308.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_309 := hx_structural_source_305
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_309.__hx_this.next()
						}, func() any {
							return hx_receiver_309.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_305.map_ = func(hx_value_306 any) haxe__IMap {
						if hx_value_306 == nil {
							var hx_zero_307 haxe__IMap
							return hx_zero_307
						}
						return hx_value_306.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_305.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_310 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_310, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_310.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_311 := hx_structural_source_310
						return hxrt.BoundFunction(func() any {
							return hx_receiver_311.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_311.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_310.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_310.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_310.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_310.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_314 := hx_structural_source_310
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_314.__hx_this.get_message()
						}, func() any {
							return hx_receiver_314.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_315 := hx_structural_source_310
						return hxrt.BoundFunction(func() any {
							return hx_receiver_315.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_315.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_316 := hx_structural_source_310
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_316.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_316.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_317 := hx_structural_source_310
						return hxrt.BoundFunction(func() any {
							return hx_receiver_317.__hx_this.get_native()
						}, func() any {
							return hx_receiver_317.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_318 := hx_structural_source_310
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_318.__hx_this.toString()
						}, func() any {
							return hx_receiver_318.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_319 := hx_structural_source_310
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_319.__hx_this.details()
						}, func() any {
							return hx_receiver_319.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_310.savedMessage = func(hx_value_312 any) *string {
						if hx_value_312 == nil {
							var hx_zero_313 *string
							return hx_zero_313
						}
						return hx_value_312.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_310.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_310.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_310.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__PosException:
		return func(hx_structural_source_320 *haxe__exceptions__PosException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_320, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_320.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_321 := hx_structural_source_320
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_321.__hx_this.toString()
						}, func() any {
							return hx_receiver_321.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_320.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_320.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_320.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_320.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_324 := hx_structural_source_320
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_324.__hx_this.get_message()
						}, func() any {
							return hx_receiver_324.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_325 := hx_structural_source_320
						return hxrt.BoundFunction(func() any {
							return hx_receiver_325.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_325.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_326 := hx_structural_source_320
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_326.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_326.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_327 := hx_structural_source_320
						return hxrt.BoundFunction(func() any {
							return hx_receiver_327.__hx_this.get_native()
						}, func() any {
							return hx_receiver_327.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_328 := hx_structural_source_320
						return hxrt.BoundFunction(func() any {
							return hx_receiver_328.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_328.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_329 := hx_structural_source_320
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_329.__hx_this.details()
						}, func() any {
							return hx_receiver_329.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_320.savedMessage = func(hx_value_322 any) *string {
						if hx_value_322 == nil {
							var hx_zero_323 *string
							return hx_zero_323
						}
						return hx_value_322.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_320.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_320.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_320.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_callable_9da3ad96(value any) *hxrt.Function[func(*string)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*string)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string)]:
		return source
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_5 any) float64 {
				if hx_value_5 == nil {
					var hx_zero_6 float64
					return hx_zero_6
				}
				return hx_value_5.(float64)
			}(arg_0))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_11 any) bool {
				if hx_value_11 == nil {
					var hx_zero_12 bool
					return hx_zero_12
				}
				return hx_value_11.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_13 any) *hxrt.SocketHandle {
				if hx_value_13 == nil {
					var hx_zero_14 *hxrt.SocketHandle
					return hx_zero_14
				}
				return hx_value_13.(*hxrt.SocketHandle)
			}(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_15 any) *hxrt.HttpExchange {
				if hx_value_15 == nil {
					var hx_zero_16 *hxrt.HttpExchange
					return hx_zero_16
				}
				return hx_value_15.(*hxrt.HttpExchange)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe___Int64_____Int64(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_17 any) int {
				if hx_value_17 == nil {
					var hx_zero_18 int
					return hx_zero_18
				}
				return hx_value_17.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_19 any) int {
				if hx_value_19 == nil {
					var hx_zero_20 int
					return hx_zero_20
				}
				return hx_value_19.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_21 any) int {
				if hx_value_21 == nil {
					var hx_zero_22 int
					return hx_zero_22
				}
				return hx_value_21.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_b9d17ba0(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_19 any) *string {
				if hx_value_19 == nil {
					var hx_zero_20 *string
					return hx_zero_20
				}
				return hx_value_19.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_21 any) *string {
				if hx_value_21 == nil {
					var hx_zero_22 *string
					return hx_zero_22
				}
				return hx_value_21.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_25 any) *string {
				if hx_value_25 == nil {
					var hx_zero_26 *string
					return hx_zero_26
				}
				return hx_value_25.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_c5309b43(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_c5f1b802(value any) *hxrt.Function[func(*haxe__io__Bytes)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return source
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) float64 {
				if hx_value_7 == nil {
					var hx_zero_8 float64
					return hx_zero_8
				}
				return hx_value_7.(float64)
			}(arg_0))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_11 any) int {
				if hx_value_11 == nil {
					var hx_zero_12 int
					return hx_zero_12
				}
				return hx_value_11.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_13 any) bool {
				if hx_value_13 == nil {
					var hx_zero_14 bool
					return hx_zero_14
				}
				return hx_value_13.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_15 any) *hxrt.SocketHandle {
				if hx_value_15 == nil {
					var hx_zero_16 *hxrt.SocketHandle
					return hx_zero_16
				}
				return hx_value_15.(*hxrt.SocketHandle)
			}(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_19 any) *hxrt.HttpExchange {
				if hx_value_19 == nil {
					var hx_zero_20 *hxrt.HttpExchange
					return hx_zero_20
				}
				return hx_value_19.(*hxrt.HttpExchange)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe___Int64_____Int64(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_21 any) int {
				if hx_value_21 == nil {
					var hx_zero_22 int
					return hx_zero_22
				}
				return hx_value_21.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_23 any) int {
				if hx_value_23 == nil {
					var hx_zero_24 int
					return hx_zero_24
				}
				return hx_value_23.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_25 any) int {
				if hx_value_25 == nil {
					var hx_zero_26 int
					return hx_zero_26
				}
				return hx_value_25.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_27 any) *string {
				if hx_value_27 == nil {
					var hx_zero_28 *string
					return hx_zero_28
				}
				return hx_value_27.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_29 any) *string {
				if hx_value_29 == nil {
					var hx_zero_30 *string
					return hx_zero_30
				}
				return hx_value_29.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_cf6928b1(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_9 any) bool {
				if hx_value_9 == nil {
					var hx_zero_10 bool
					return hx_zero_10
				}
				return hx_value_9.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_11 any) bool {
				if hx_value_11 == nil {
					var hx_zero_12 bool
					return hx_zero_12
				}
				return hx_value_11.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_13 any) bool {
				if hx_value_13 == nil {
					var hx_zero_14 bool
					return hx_zero_14
				}
				return hx_value_13.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_15 any) bool {
				if hx_value_15 == nil {
					var hx_zero_16 bool
					return hx_zero_16
				}
				return hx_value_15.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_17 any) bool {
				if hx_value_17 == nil {
					var hx_zero_18 bool
					return hx_zero_18
				}
				return hx_value_17.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_19 any) bool {
				if hx_value_19 == nil {
					var hx_zero_20 bool
					return hx_zero_20
				}
				return hx_value_19.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_21 any) bool {
				if hx_value_21 == nil {
					var hx_zero_22 bool
					return hx_zero_22
				}
				return hx_value_21.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_23 any) bool {
				if hx_value_23 == nil {
					var hx_zero_24 bool
					return hx_zero_24
				}
				return hx_value_23.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_25 any) bool {
				if hx_value_25 == nil {
					var hx_zero_26 bool
					return hx_zero_26
				}
				return hx_value_25.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_d8be58e9(value any) *hxrt.Function[func(int)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int)]:
		return source
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(float64(arg_0))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) *hxrt.SocketHandle {
				if hx_value_7 == nil {
					var hx_zero_8 *hxrt.SocketHandle
					return hx_zero_8
				}
				return hx_value_7.(*hxrt.SocketHandle)
			}(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_11 any) *hxrt.HttpExchange {
				if hx_value_11 == nil {
					var hx_zero_12 *hxrt.HttpExchange
					return hx_zero_12
				}
				return hx_value_11.(*hxrt.HttpExchange)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe___Int64_____Int64(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_haxe___Int64_____Int64(value any) *haxe___Int64_____Int64 {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return source
	default:
		return value.(*haxe___Int64_____Int64)
	}
}

func hx_recover_class_haxe__ds__StringMap(value any) *haxe__ds__StringMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return source
	default:
		return value.(*haxe__ds__StringMap)
	}
}

func hx_recover_class_haxe__io__Bytes(value any) *haxe__io__Bytes {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return source
	default:
		return value.(*haxe__io__Bytes)
	}
}

func hx_recover_class_haxe__io__BytesBuffer(value any) *haxe__io__BytesBuffer {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__BytesBuffer:
		return source
	default:
		return value.(*haxe__io__BytesBuffer)
	}
}

func hx_recover_class_haxe__io__Eof(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_recover_class_haxe__io__Input(value any) *haxe__io__Input {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return source.haxe__io__Input
	case *haxe__io__Input:
		return source
	default:
		return value.(*haxe__io__Input)
	}
}

func hx_recover_class_sys__net__Socket(value any) *sys__net__Socket {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__Socket:
		return source
	default:
		return value.(*sys__net__Socket)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Eof(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}
