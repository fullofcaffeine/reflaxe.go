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

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *_Main__DataOutput:
		return func(hx_structural_source_450 *_Main__DataOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_450, func(name string) hxrt.AnonymousField {
				switch name {
				case "events":
					return hxrt.AnonymousField{Value: hx_structural_source_450.events, Found: true}
				case "closeCount":
					return hxrt.AnonymousField{Value: hx_structural_source_450.closeCount, Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_451 := hx_structural_source_450
						return hxrt.BoundFunction(func(size int) {
							hx_receiver_451.__hx_this.prepare(size)
						}, func() any {
							return hx_receiver_451.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_452 := hx_structural_source_450
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_452.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_452.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_453 := hx_structural_source_450
						return hxrt.BoundFunction(func() {
							hx_receiver_453.__hx_this.close()
						}, func() any {
							return hx_receiver_453.__hx_this
						}, "close")
					}(), Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_450.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_454 := hx_structural_source_450
						return hxrt.BoundFunction(func() int {
							return hx_receiver_454.__hx_this.get_length()
						}, func() any {
							return hx_receiver_454.__hx_this
						}, "get_length")
					}(), Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_455 := hx_structural_source_450
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_455.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_455.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_456 := hx_structural_source_450
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_456.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_456.__hx_this
						}, "getBytes")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_450.bigEndian, Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_457 := hx_structural_source_450
						return hxrt.BoundFunction(func() {
							hx_receiver_457.__hx_this.flush()
						}, func() any {
							return hx_receiver_457.__hx_this
						}, "flush")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_458 := hx_structural_source_450
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_458.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_458.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_459 := hx_structural_source_450
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_459.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_459.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_460 := hx_structural_source_450
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_460.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_460.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_461 := hx_structural_source_450
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_461.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_461.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_462 := hx_structural_source_450
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_462.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_462.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_463 := hx_structural_source_450
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_463.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_463.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_464 := hx_structural_source_450
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_464.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_464.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_465 := hx_structural_source_450
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_465.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_465.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_466 := hx_structural_source_450
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_466.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_466.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_467 := hx_structural_source_450
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_467.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_467.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_468 := hx_structural_source_450
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_468.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_468.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_469 := hx_structural_source_450
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_469.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_469.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_470 := hx_structural_source_450
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_470.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_470.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_450.b = hx_recover_class_6(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_471 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_471, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_471.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_471.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_471.high = func(hx_value_472 any) int {
						if hx_value_472 == nil {
							var hx_zero_473 int
							return hx_zero_473
						}
						return hx_value_472.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_471.low = func(hx_value_474 any) int {
						if hx_value_474 == nil {
							var hx_zero_475 int
							return hx_zero_475
						}
						return hx_value_474.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketOutput:
		return func(hx_structural_source_476 *sys__net__SocketOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_476, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_476.handle, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_479 := hx_structural_source_476
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_479.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_479.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_480 := hx_structural_source_476
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_480.__hx_this.writeBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_480.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_481 := hx_structural_source_476
						return hxrt.BoundFunction(func() {
							hx_receiver_481.__hx_this.flush()
						}, func() any {
							return hx_receiver_481.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_482 := hx_structural_source_476
						return hxrt.BoundFunction(func() {
							hx_receiver_482.__hx_this.close()
						}, func() any {
							return hx_receiver_482.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_476.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_483 := hx_structural_source_476
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_483.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_483.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_484 := hx_structural_source_476
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_484.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_484.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_485 := hx_structural_source_476
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_485.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_485.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_486 := hx_structural_source_476
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_486.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_486.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_487 := hx_structural_source_476
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_487.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_487.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_488 := hx_structural_source_476
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_488.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_488.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_489 := hx_structural_source_476
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_489.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_489.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_490 := hx_structural_source_476
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_490.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_490.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_491 := hx_structural_source_476
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_491.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_491.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_492 := hx_structural_source_476
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_492.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_492.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_493 := hx_structural_source_476
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_493.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_493.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_494 := hx_structural_source_476
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_494.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_494.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_495 := hx_structural_source_476
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_495.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_495.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_496 := hx_structural_source_476
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_496.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_496.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_476.handle = func(hx_value_477 any) *hxrt.SocketHandle {
						if hx_value_477 == nil {
							var hx_zero_478 *hxrt.SocketHandle
							return hx_zero_478
						}
						return hx_value_477.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketInput:
		return func(hx_structural_source_497 *sys__net__SocketInput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_497, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_497.handle, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_500 := hx_structural_source_497
						return hxrt.BoundFunction(func() int {
							return hx_receiver_500.__hx_this.readByte()
						}, func() any {
							return hx_receiver_500.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_501 := hx_structural_source_497
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_501.__hx_this.readBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_501.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_502 := hx_structural_source_497
						return hxrt.BoundFunction(func() {
							hx_receiver_502.__hx_this.close()
						}, func() any {
							return hx_receiver_502.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_497.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_503 := hx_structural_source_497
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_503.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_503.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_504 := hx_structural_source_497
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_504.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_504.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_505 := hx_structural_source_497
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_505.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_505.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_506 := hx_structural_source_497
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_506.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_506.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_507 := hx_structural_source_497
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_507.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_507.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_508 := hx_structural_source_497
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_508.__hx_this.readLine()
						}, func() any {
							return hx_receiver_508.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_509 := hx_structural_source_497
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_509.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_509.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_510 := hx_structural_source_497
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_510.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_510.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_511 := hx_structural_source_497
						return hxrt.BoundFunction(func() int {
							return hx_receiver_511.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_511.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_512 := hx_structural_source_497
						return hxrt.BoundFunction(func() int {
							return hx_receiver_512.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_512.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_513 := hx_structural_source_497
						return hxrt.BoundFunction(func() int {
							return hx_receiver_513.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_513.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_514 := hx_structural_source_497
						return hxrt.BoundFunction(func() int {
							return hx_receiver_514.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_514.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_515 := hx_structural_source_497
						return hxrt.BoundFunction(func() int {
							return hx_receiver_515.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_515.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_516 := hx_structural_source_497
						return hxrt.BoundFunction(func() int {
							return hx_receiver_516.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_516.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_517 := hx_structural_source_497
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_517.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_517.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_497.handle = func(hx_value_498 any) *hxrt.SocketHandle {
						if hx_value_498 == nil {
							var hx_zero_499 *hxrt.SocketHandle
							return hx_zero_499
						}
						return hx_value_498.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Socket:
		return func(hx_structural_source_518 *sys__net__Socket) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_518, func(name string) hxrt.AnonymousField {
				switch name {
				case "input":
					return hxrt.AnonymousField{Value: hx_structural_source_518.input, Found: true}
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_518.output, Found: true}
				case "custom":
					return hxrt.AnonymousField{Value: hx_structural_source_518.custom, Found: true}
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_518.handle, Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_521 := hx_structural_source_518
						return hxrt.BoundFunction(func() {
							hx_receiver_521.__hx_this.close()
						}, func() any {
							return hx_receiver_521.__hx_this
						}, "close")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_522 := hx_structural_source_518
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_522.__hx_this.read()
						}, func() any {
							return hx_receiver_522.__hx_this
						}, "read")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_523 := hx_structural_source_518
						return hxrt.BoundFunction(func(content *string) {
							hx_receiver_523.__hx_this.write(content)
						}, func() any {
							return hx_receiver_523.__hx_this
						}, "write")
					}(), Found: true}
				case "connect":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_524 := hx_structural_source_518
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_524.__hx_this.connect(host, port)
						}, func() any {
							return hx_receiver_524.__hx_this
						}, "connect")
					}(), Found: true}
				case "listen":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_525 := hx_structural_source_518
						return hxrt.BoundFunction(func(connections int) {
							hx_receiver_525.__hx_this.listen(connections)
						}, func() any {
							return hx_receiver_525.__hx_this
						}, "listen")
					}(), Found: true}
				case "shutdown":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, bool)] {
						hx_receiver_526 := hx_structural_source_518
						return hxrt.BoundFunction(func(read bool, write bool) {
							hx_receiver_526.__hx_this.shutdown(read, write)
						}, func() any {
							return hx_receiver_526.__hx_this
						}, "shutdown")
					}(), Found: true}
				case "bind":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_527 := hx_structural_source_518
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_527.__hx_this.bind(host, port)
						}, func() any {
							return hx_receiver_527.__hx_this
						}, "bind")
					}(), Found: true}
				case "accept":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *sys__net__Socket] {
						hx_receiver_528 := hx_structural_source_518
						return hxrt.BoundFunction(func() *sys__net__Socket {
							return hx_receiver_528.__hx_this.accept()
						}, func() any {
							return hx_receiver_528.__hx_this
						}, "accept")
					}(), Found: true}
				case "peer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_529 := hx_structural_source_518
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_529.__hx_this.peer()
						}, func() any {
							return hx_receiver_529.__hx_this
						}, "peer")
					}(), Found: true}
				case "host":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_530 := hx_structural_source_518
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_530.__hx_this.host()
						}, func() any {
							return hx_receiver_530.__hx_this
						}, "host")
					}(), Found: true}
				case "setTimeout":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_531 := hx_structural_source_518
						return hxrt.BoundFunction(func(timeout float64) {
							hx_receiver_531.__hx_this.setTimeout(timeout)
						}, func() any {
							return hx_receiver_531.__hx_this
						}, "setTimeout")
					}(), Found: true}
				case "waitForRead":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_532 := hx_structural_source_518
						return hxrt.BoundFunction(func() {
							hx_receiver_532.__hx_this.waitForRead()
						}, func() any {
							return hx_receiver_532.__hx_this
						}, "waitForRead")
					}(), Found: true}
				case "setBlocking":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_533 := hx_structural_source_518
						return hxrt.BoundFunction(func(blocking bool) {
							hx_receiver_533.__hx_this.setBlocking(blocking)
						}, func() any {
							return hx_receiver_533.__hx_this
						}, "setBlocking")
					}(), Found: true}
				case "setFastSend":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_534 := hx_structural_source_518
						return hxrt.BoundFunction(func(fastSend bool) {
							hx_receiver_534.__hx_this.setFastSend(fastSend)
						}, func() any {
							return hx_receiver_534.__hx_this
						}, "setFastSend")
					}(), Found: true}
				case "replaceHandle":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.SocketHandle)] {
						hx_receiver_535 := hx_structural_source_518
						return hxrt.BoundFunction(func(next *hxrt.SocketHandle) {
							hx_receiver_535.__hx_this.replaceHandle(next)
						}, func() any {
							return hx_receiver_535.__hx_this
						}, "replaceHandle")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "custom":
					hx_structural_source_518.custom = incoming
					return true
				case "handle":
					hx_structural_source_518.handle = func(hx_value_519 any) *hxrt.SocketHandle {
						if hx_value_519 == nil {
							var hx_zero_520 *hxrt.SocketHandle
							return hx_zero_520
						}
						return hx_value_519.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Host:
		return func(hx_structural_source_536 *sys__net__Host) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_536, func(name string) hxrt.AnonymousField {
				switch name {
				case "host":
					return hxrt.AnonymousField{Value: hx_structural_source_536.host, Found: true}
				case "ip":
					return hxrt.AnonymousField{Value: hx_structural_source_536.ip, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_537 := hx_structural_source_536
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_537.__hx_this.toString()
						}, func() any {
							return hx_receiver_537.__hx_this
						}, "toString")
					}(), Found: true}
				case "reverse":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_538 := hx_structural_source_536
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_538.__hx_this.reverse()
						}, func() any {
							return hx_receiver_538.__hx_this
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
		return func(hx_structural_source_539 *sys__Http) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_539, func(name string) hxrt.AnonymousField {
				switch name {
				case "noShutdown":
					return hxrt.AnonymousField{Value: hx_structural_source_539.noShutdown, Found: true}
				case "cnxTimeout":
					return hxrt.AnonymousField{Value: hx_structural_source_539.cnxTimeout, Found: true}
				case "responseHeaders":
					return hxrt.AnonymousField{Value: hx_structural_source_539.responseHeaders, Found: true}
				case "responseHeadersSameKey":
					return hxrt.AnonymousField{Value: hx_structural_source_539.responseHeadersSameKey, Found: true}
				case "file":
					return hxrt.AnonymousField{Value: hx_structural_source_539.file, Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_546 := hx_structural_source_539
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_546.__hx_this.request(post)
						}, func() any {
							return hx_receiver_546.__hx_this
						}, "request")
					}(), Found: true}
				case "fileTransfert":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_547 := hx_structural_source_539
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_547.__hx_this.fileTransfert(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_547.__hx_this
						}, "fileTransfert")
					}(), Found: true}
				case "fileTransfer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_548 := hx_structural_source_539
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_548.__hx_this.fileTransfer(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_548.__hx_this
						}, "fileTransfer")
					}(), Found: true}
				case "customRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_549 := hx_structural_source_539
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_549.__hx_this.customRequest(post, api, sock, method)
						}, func() any {
							return hx_receiver_549.__hx_this
						}, "customRequest")
					}(), Found: true}
				case "getResponseHeaderValues":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_550 := hx_structural_source_539
						return hxrt.BoundFunction(func(key *string) *hxrt.Array {
							return hx_receiver_550.__hx_this.getResponseHeaderValues(key)
						}, func() any {
							return hx_receiver_550.__hx_this
						}, "getResponseHeaderValues")
					}(), Found: true}
				case "requestWith":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_551 := hx_structural_source_539
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_551.__hx_this.requestWith(post, api, sock, method)
						}, func() any {
							return hx_receiver_551.__hx_this
						}, "requestWith")
					}(), Found: true}
				case "pumpUpload":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange, hxrt.Anonymous) hxrt.Anonymous] {
						hx_receiver_552 := hx_structural_source_539
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange, upload hxrt.Anonymous) hxrt.Anonymous {
							return hx_receiver_552.__hx_this.pumpUpload(exchange, upload)
						}, func() any {
							return hx_receiver_552.__hx_this
						}, "pumpUpload")
					}(), Found: true}
				case "handleDataRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *string)] {
						hx_receiver_553 := hx_structural_source_539
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, method *string) {
							hx_receiver_553.__hx_this.handleDataRequest(post, api, method)
						}, func() any {
							return hx_receiver_553.__hx_this
						}, "handleDataRequest")
					}(), Found: true}
				case "recordResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange)] {
						hx_receiver_554 := hx_structural_source_539
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange) {
							hx_receiver_554.__hx_this.recordResponseHeaders(exchange)
						}, func() any {
							return hx_receiver_554.__hx_this
						}, "recordResponseHeaders")
					}(), Found: true}
				case "resetResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_555 := hx_structural_source_539
						return hxrt.BoundFunction(func() {
							hx_receiver_555.__hx_this.resetResponseHeaders()
						}, func() any {
							return hx_receiver_555.__hx_this
						}, "resetResponseHeaders")
					}(), Found: true}
				case "encodedParameters":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_556 := hx_structural_source_539
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_556.__hx_this.encodedParameters()
						}, func() any {
							return hx_receiver_556.__hx_this
						}, "encodedParameters")
					}(), Found: true}
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_539.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_539.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_539.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_539.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_539.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_539.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_539.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_539.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_567 := hx_structural_source_539
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_567.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_567.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_568 := hx_structural_source_539
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_568.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_568.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_569 := hx_structural_source_539
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_569.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_569.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_570 := hx_structural_source_539
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_570.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_570.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_571 := hx_structural_source_539
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_571.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_571.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_572 := hx_structural_source_539
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_572.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_572.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_539.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_539.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_539.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_539.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_573 := hx_structural_source_539
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_573.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_573.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_574 := hx_structural_source_539
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_574.__hx_this.success(data)
						}, func() any {
							return hx_receiver_574.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_575 := hx_structural_source_539
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_575.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_575.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "noShutdown":
					hx_structural_source_539.noShutdown = func(hx_value_540 any) bool {
						if hx_value_540 == nil {
							var hx_zero_541 bool
							return hx_zero_541
						}
						return hx_value_540.(bool)
					}(incoming)
					return true
				case "cnxTimeout":
					hx_structural_source_539.cnxTimeout = func(hx_value_542 any) float64 {
						if hx_value_542 == nil {
							var hx_zero_543 float64
							return hx_zero_543
						}
						return hx_value_542.(float64)
					}(incoming)
					return true
				case "responseHeaders":
					hx_structural_source_539.responseHeaders = func(hx_value_544 any) *haxe__ds__StringMap {
						if hx_value_544 == nil {
							var hx_zero_545 *haxe__ds__StringMap
							return hx_zero_545
						}
						return hx_value_544.(*haxe__ds__StringMap)
					}(incoming)
					return true
				case "responseHeadersSameKey":
					hx_structural_source_539.responseHeadersSameKey = hx_recover_class_4(incoming)
					return true
				case "file":
					hx_structural_source_539.file = hx_structural_recover(incoming)
					return true
				case "url":
					hx_structural_source_539.url = func(hx_value_557 any) *string {
						if hx_value_557 == nil {
							var hx_zero_558 *string
							return hx_zero_558
						}
						return hx_value_557.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_539.responseAsString = func(hx_value_559 any) *string {
						if hx_value_559 == nil {
							var hx_zero_560 *string
							return hx_zero_560
						}
						return hx_value_559.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_539.postData = func(hx_value_561 any) *string {
						if hx_value_561 == nil {
							var hx_zero_562 *string
							return hx_zero_562
						}
						return hx_value_561.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_539.postBytes = hx_recover_class_7(incoming)
					return true
				case "headers":
					hx_structural_source_539.headers = func(hx_value_563 any) *hxrt.Array {
						if hx_value_563 == nil {
							var hx_zero_564 *hxrt.Array
							return hx_zero_564
						}
						return hx_value_563.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_539.params = func(hx_value_565 any) *hxrt.Array {
						if hx_value_565 == nil {
							var hx_zero_566 *hxrt.Array
							return hx_zero_566
						}
						return hx_value_565.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_539.onData = hx_recover_callable_3(incoming)
					return true
				case "onBytes":
					hx_structural_source_539.onBytes = hx_recover_callable_4(incoming)
					return true
				case "onError":
					hx_structural_source_539.onError = hx_recover_callable_3(incoming)
					return true
				case "onStatus":
					hx_structural_source_539.onStatus = hx_recover_callable_5(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Output:
		return func(hx_structural_source_576 *haxe__io__Output) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_576, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_576.bigEndian, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_577 := hx_structural_source_576
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_577.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_577.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_578 := hx_structural_source_576
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_578.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_578.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_579 := hx_structural_source_576
						return hxrt.BoundFunction(func() {
							hx_receiver_579.__hx_this.flush()
						}, func() any {
							return hx_receiver_579.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_580 := hx_structural_source_576
						return hxrt.BoundFunction(func() {
							hx_receiver_580.__hx_this.close()
						}, func() any {
							return hx_receiver_580.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_581 := hx_structural_source_576
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_581.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_581.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_582 := hx_structural_source_576
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_582.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_582.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_583 := hx_structural_source_576
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_583.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_583.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_584 := hx_structural_source_576
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_584.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_584.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_585 := hx_structural_source_576
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_585.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_585.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_586 := hx_structural_source_576
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_586.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_586.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_587 := hx_structural_source_576
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_587.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_587.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_588 := hx_structural_source_576
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_588.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_588.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_589 := hx_structural_source_576
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_589.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_589.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_590 := hx_structural_source_576
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_590.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_590.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_591 := hx_structural_source_576
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_591.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_591.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_592 := hx_structural_source_576
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_592.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_592.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_593 := hx_structural_source_576
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_593.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_593.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_594 := hx_structural_source_576
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_594.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_594.__hx_this
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
		return func(hx_structural_source_595 *haxe__io__Input) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_595, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_595.bigEndian, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_596 := hx_structural_source_595
						return hxrt.BoundFunction(func() int {
							return hx_receiver_596.__hx_this.readByte()
						}, func() any {
							return hx_receiver_596.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_597 := hx_structural_source_595
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_597.__hx_this.readBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_597.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_598 := hx_structural_source_595
						return hxrt.BoundFunction(func() {
							hx_receiver_598.__hx_this.close()
						}, func() any {
							return hx_receiver_598.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_599 := hx_structural_source_595
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_599.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_599.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_600 := hx_structural_source_595
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_600.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_600.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_601 := hx_structural_source_595
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_601.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_601.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_602 := hx_structural_source_595
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_602.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_602.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_603 := hx_structural_source_595
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_603.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_603.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_604 := hx_structural_source_595
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_604.__hx_this.readLine()
						}, func() any {
							return hx_receiver_604.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_605 := hx_structural_source_595
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_605.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_605.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_606 := hx_structural_source_595
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_606.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_606.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_607 := hx_structural_source_595
						return hxrt.BoundFunction(func() int {
							return hx_receiver_607.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_607.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_608 := hx_structural_source_595
						return hxrt.BoundFunction(func() int {
							return hx_receiver_608.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_608.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_609 := hx_structural_source_595
						return hxrt.BoundFunction(func() int {
							return hx_receiver_609.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_609.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_610 := hx_structural_source_595
						return hxrt.BoundFunction(func() int {
							return hx_receiver_610.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_610.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_611 := hx_structural_source_595
						return hxrt.BoundFunction(func() int {
							return hx_receiver_611.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_611.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_612 := hx_structural_source_595
						return hxrt.BoundFunction(func() int {
							return hx_receiver_612.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_612.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_613 := hx_structural_source_595
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_613.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_613.__hx_this
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
		return func(hx_structural_source_614 *haxe__io__Eof) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_614, func(name string) hxrt.AnonymousField {
				switch name {
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_615 := hx_structural_source_614
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_615.__hx_this.toString()
						}, func() any {
							return hx_receiver_615.__hx_this
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
		return func(hx_structural_source_616 *haxe__io__BytesOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_616, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_616.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_617 := hx_structural_source_616
						return hxrt.BoundFunction(func() int {
							return hx_receiver_617.__hx_this.get_length()
						}, func() any {
							return hx_receiver_617.__hx_this
						}, "get_length")
					}(), Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_618 := hx_structural_source_616
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_618.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_618.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_619 := hx_structural_source_616
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_619.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_619.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_620 := hx_structural_source_616
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_620.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_620.__hx_this
						}, "getBytes")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_616.bigEndian, Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_621 := hx_structural_source_616
						return hxrt.BoundFunction(func() {
							hx_receiver_621.__hx_this.flush()
						}, func() any {
							return hx_receiver_621.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_622 := hx_structural_source_616
						return hxrt.BoundFunction(func() {
							hx_receiver_622.__hx_this.close()
						}, func() any {
							return hx_receiver_622.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_623 := hx_structural_source_616
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_623.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_623.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_624 := hx_structural_source_616
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_624.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_624.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_625 := hx_structural_source_616
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_625.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_625.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_626 := hx_structural_source_616
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_626.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_626.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_627 := hx_structural_source_616
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_627.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_627.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_628 := hx_structural_source_616
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_628.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_628.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_629 := hx_structural_source_616
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_629.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_629.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_630 := hx_structural_source_616
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_630.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_630.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_631 := hx_structural_source_616
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_631.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_631.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_632 := hx_structural_source_616
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_632.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_632.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_633 := hx_structural_source_616
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_633.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_633.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_634 := hx_structural_source_616
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_634.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_634.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_635 := hx_structural_source_616
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_635.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_635.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_636 := hx_structural_source_616
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_636.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_636.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_616.b = hx_recover_class_6(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__BytesBuffer:
		return func(hx_structural_source_637 *haxe__io__BytesBuffer) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_637, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_637.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_640 := hx_structural_source_637
						return hxrt.BoundFunction(func() int {
							return hx_receiver_640.__hx_this.get_length()
						}, func() any {
							return hx_receiver_640.__hx_this
						}, "get_length")
					}(), Found: true}
				case "addByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_641 := hx_structural_source_637
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_641.__hx_this.addByte(value)
						}, func() any {
							return hx_receiver_641.__hx_this
						}, "addByte")
					}(), Found: true}
				case "add":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_642 := hx_structural_source_637
						return hxrt.BoundFunction(func(source *haxe__io__Bytes) {
							hx_receiver_642.__hx_this.add(source)
						}, func() any {
							return hx_receiver_642.__hx_this
						}, "add")
					}(), Found: true}
				case "addString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_643 := hx_structural_source_637
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_643.__hx_this.addString(value, encoding)
						}, func() any {
							return hx_receiver_643.__hx_this
						}, "addString")
					}(), Found: true}
				case "addInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_644 := hx_structural_source_637
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_644.__hx_this.addInt32(value)
						}, func() any {
							return hx_receiver_644.__hx_this
						}, "addInt32")
					}(), Found: true}
				case "addInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Int64_____Int64)] {
						hx_receiver_645 := hx_structural_source_637
						return hxrt.BoundFunction(func(value *haxe___Int64_____Int64) {
							hx_receiver_645.__hx_this.addInt64(value)
						}, func() any {
							return hx_receiver_645.__hx_this
						}, "addInt64")
					}(), Found: true}
				case "addFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_646 := hx_structural_source_637
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_646.__hx_this.addFloat(value)
						}, func() any {
							return hx_receiver_646.__hx_this
						}, "addFloat")
					}(), Found: true}
				case "addDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_647 := hx_structural_source_637
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_647.__hx_this.addDouble(value)
						}, func() any {
							return hx_receiver_647.__hx_this
						}, "addDouble")
					}(), Found: true}
				case "addBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_648 := hx_structural_source_637
						return hxrt.BoundFunction(func(source *haxe__io__Bytes, pos int, len int) {
							hx_receiver_648.__hx_this.addBytes(source, pos, len)
						}, func() any {
							return hx_receiver_648.__hx_this
						}, "addBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_649 := hx_structural_source_637
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_649.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_649.__hx_this
						}, "getBytes")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_637.b = func(hx_value_638 any) []int {
						if hx_value_638 == nil {
							var hx_zero_639 []int
							return hx_zero_639
						}
						return hx_value_638.([]int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_650 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_650, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_650.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_650.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_650.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_650.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_650.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_659 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_659.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_659.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_660 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_660.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_660.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_661 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
							hx_receiver_661.__hx_this.blit(pos, src, srcpos, len)
						}, func() any {
							return hx_receiver_661.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_662 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, len int, value int) {
							hx_receiver_662.__hx_this.fill(pos, len, value)
						}, func() any {
							return hx_receiver_662.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_663 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
							return hx_receiver_663.__hx_this.sub(pos, len)
						}, func() any {
							return hx_receiver_663.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_664 := hx_structural_source_650
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_664.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_664.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_665 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_665.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_665.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_666 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_666.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_666.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_667 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_667.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_667.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_668 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_668.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_668.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_669 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_669.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_669.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_670 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_670.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_670.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_671 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_671.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_671.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_672 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_672.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_672.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_673 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_673.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_673.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_674 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_674.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_674.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_675 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_675.__hx_this.getString(pos, len, encoding)
						}, func() any {
							return hx_receiver_675.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_676 := hx_structural_source_650
						return hxrt.BoundFunction(func(pos int, len int) *string {
							return hx_receiver_676.__hx_this.readString(pos, len)
						}, func() any {
							return hx_receiver_676.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_677 := hx_structural_source_650
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_677.__hx_this.toString()
						}, func() any {
							return hx_receiver_677.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_678 := hx_structural_source_650
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_678.__hx_this.toHex()
						}, func() any {
							return hx_receiver_678.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_679 := hx_structural_source_650
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_679.__hx_this.getData()
						}, func() any {
							return hx_receiver_679.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_680 := hx_structural_source_650
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_680.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_680.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_650.b = func(hx_value_651 any) []int {
						if hx_value_651 == nil {
							var hx_zero_652 []int
							return hx_zero_652
						}
						return hx_value_651.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_650.__hx_raw = func(hx_value_653 any) *hxrt.ByteView {
						if hx_value_653 == nil {
							var hx_zero_654 *hxrt.ByteView
							return hx_zero_654
						}
						return hx_value_653.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_650.__hx_rawValid = func(hx_value_655 any) bool {
						if hx_value_655 == nil {
							var hx_zero_656 bool
							return hx_zero_656
						}
						return hx_value_655.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_650.__hx_dataExposed = func(hx_value_657 any) bool {
						if hx_value_657 == nil {
							var hx_zero_658 bool
							return hx_zero_658
						}
						return hx_value_657.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__http__HttpBase:
		return func(hx_structural_source_681 *haxe__http__HttpBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_681, func(name string) hxrt.AnonymousField {
				switch name {
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_681.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_681.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_681.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_681.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_681.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_681.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_681.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_681.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_692 := hx_structural_source_681
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_692.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_692.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_693 := hx_structural_source_681
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_693.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_693.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_694 := hx_structural_source_681
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_694.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_694.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_695 := hx_structural_source_681
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_695.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_695.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_696 := hx_structural_source_681
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_696.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_696.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_697 := hx_structural_source_681
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_697.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_697.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_698 := hx_structural_source_681
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_698.__hx_this.request(post)
						}, func() any {
							return hx_receiver_698.__hx_this
						}, "request")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_681.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_681.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_681.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_681.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_699 := hx_structural_source_681
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_699.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_699.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_700 := hx_structural_source_681
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_700.__hx_this.success(data)
						}, func() any {
							return hx_receiver_700.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_701 := hx_structural_source_681
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_701.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_701.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "url":
					hx_structural_source_681.url = func(hx_value_682 any) *string {
						if hx_value_682 == nil {
							var hx_zero_683 *string
							return hx_zero_683
						}
						return hx_value_682.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_681.responseAsString = func(hx_value_684 any) *string {
						if hx_value_684 == nil {
							var hx_zero_685 *string
							return hx_zero_685
						}
						return hx_value_684.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_681.postData = func(hx_value_686 any) *string {
						if hx_value_686 == nil {
							var hx_zero_687 *string
							return hx_zero_687
						}
						return hx_value_686.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_681.postBytes = hx_recover_class_7(incoming)
					return true
				case "headers":
					hx_structural_source_681.headers = func(hx_value_688 any) *hxrt.Array {
						if hx_value_688 == nil {
							var hx_zero_689 *hxrt.Array
							return hx_zero_689
						}
						return hx_value_688.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_681.params = func(hx_value_690 any) *hxrt.Array {
						if hx_value_690 == nil {
							var hx_zero_691 *hxrt.Array
							return hx_zero_691
						}
						return hx_value_690.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_681.onData = hx_recover_callable_3(incoming)
					return true
				case "onBytes":
					hx_structural_source_681.onBytes = hx_recover_callable_4(incoming)
					return true
				case "onError":
					hx_structural_source_681.onError = hx_recover_callable_3(incoming)
					return true
				case "onStatus":
					hx_structural_source_681.onStatus = hx_recover_callable_5(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__NotImplementedException:
		return func(hx_structural_source_702 *haxe__exceptions__NotImplementedException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_702, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_702.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_703 := hx_structural_source_702
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_703.__hx_this.toString()
						}, func() any {
							return hx_receiver_703.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_702.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_702.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_702.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_702.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_706 := hx_structural_source_702
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_706.__hx_this.get_message()
						}, func() any {
							return hx_receiver_706.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_707 := hx_structural_source_702
						return hxrt.BoundFunction(func() any {
							return hx_receiver_707.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_707.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_708 := hx_structural_source_702
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_708.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_708.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_709 := hx_structural_source_702
						return hxrt.BoundFunction(func() any {
							return hx_receiver_709.__hx_this.get_native()
						}, func() any {
							return hx_receiver_709.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_710 := hx_structural_source_702
						return hxrt.BoundFunction(func() any {
							return hx_receiver_710.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_710.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_711 := hx_structural_source_702
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_711.__hx_this.details()
						}, func() any {
							return hx_receiver_711.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_702.savedMessage = func(hx_value_704 any) *string {
						if hx_value_704 == nil {
							var hx_zero_705 *string
							return hx_zero_705
						}
						return hx_value_704.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_702.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_702.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_702.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__StringMap:
		return func(hx_structural_source_712 *haxe__ds__StringMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_712, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_712.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, any)] {
						hx_receiver_715 := hx_structural_source_712
						return hxrt.BoundFunction(func(key *string, value any) {
							hx_receiver_715.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_715.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_716 := hx_structural_source_712
						return hxrt.BoundFunction(func(key *string) any {
							return hx_receiver_716.__hx_this.get(key)
						}, func() any {
							return hx_receiver_716.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_717 := hx_structural_source_712
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_717.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_717.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_718 := hx_structural_source_712
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_718.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_718.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_719 := hx_structural_source_712
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_719.__hx_this.keys()
						}, func() any {
							return hx_receiver_719.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_720 := hx_structural_source_712
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_720.__hx_this.iterator()
						}, func() any {
							return hx_receiver_720.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_721 := hx_structural_source_712
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_721.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_721.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_722 := hx_structural_source_712
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_722.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_722.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_723 := hx_structural_source_712
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_723.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_723.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_724 := hx_structural_source_712
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_724.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_724.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_725 := hx_structural_source_712
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_725.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_725.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_726 := hx_structural_source_712
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_726.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_726.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__StringMap] {
						hx_receiver_727 := hx_structural_source_712
						return hxrt.BoundFunction(func() *haxe__ds__StringMap {
							return hx_receiver_727.__hx_this.copy()
						}, func() any {
							return hx_receiver_727.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_728 := hx_structural_source_712
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_728.__hx_this.toString()
						}, func() any {
							return hx_receiver_728.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_729 := hx_structural_source_712
						return hxrt.BoundFunction(func() {
							hx_receiver_729.__hx_this.clear()
						}, func() any {
							return hx_receiver_729.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_712.h = func(hx_value_713 any) *hxrt.StringMapCell {
						if hx_value_713 == nil {
							var hx_zero_714 *hxrt.StringMapCell
							return hx_zero_714
						}
						return hx_value_713.(*hxrt.StringMapCell)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_730 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_730, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_730.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_730.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_730.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_730.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_733 := hx_structural_source_730
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_733.__hx_this.get_message()
						}, func() any {
							return hx_receiver_733.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_734 := hx_structural_source_730
						return hxrt.BoundFunction(func() any {
							return hx_receiver_734.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_734.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_735 := hx_structural_source_730
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_735.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_735.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_736 := hx_structural_source_730
						return hxrt.BoundFunction(func() any {
							return hx_receiver_736.__hx_this.get_native()
						}, func() any {
							return hx_receiver_736.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_737 := hx_structural_source_730
						return hxrt.BoundFunction(func() any {
							return hx_receiver_737.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_737.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_738 := hx_structural_source_730
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_738.__hx_this.toString()
						}, func() any {
							return hx_receiver_738.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_739 := hx_structural_source_730
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_739.__hx_this.details()
						}, func() any {
							return hx_receiver_739.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_730.savedMessage = func(hx_value_731 any) *string {
						if hx_value_731 == nil {
							var hx_zero_732 *string
							return hx_zero_732
						}
						return hx_value_731.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_730.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_730.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_730.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_740 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_740, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_740.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_740.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_745 := hx_structural_source_740
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_745.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_745.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_746 := hx_structural_source_740
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_746.__hx_this.next()
						}, func() any {
							return hx_receiver_746.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_740.offset = func(hx_value_741 any) int {
						if hx_value_741 == nil {
							var hx_zero_742 int
							return hx_zero_742
						}
						return hx_value_741.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_740.s = func(hx_value_743 any) *string {
						if hx_value_743 == nil {
							var hx_zero_744 *string
							return hx_zero_744
						}
						return hx_value_743.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_747 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_747, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_747.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_747.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_752 := hx_structural_source_747
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_752.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_752.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_753 := hx_structural_source_747
						return hxrt.BoundFunction(func() int {
							return hx_receiver_753.__hx_this.next()
						}, func() any {
							return hx_receiver_753.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_747.offset = func(hx_value_748 any) int {
						if hx_value_748 == nil {
							var hx_zero_749 int
							return hx_zero_749
						}
						return hx_value_748.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_747.s = func(hx_value_750 any) *string {
						if hx_value_750 == nil {
							var hx_zero_751 *string
							return hx_zero_751
						}
						return hx_value_750.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_754 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_754, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_754.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_754.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_757 := hx_structural_source_754
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_757.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_757.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_758 := hx_structural_source_754
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_758.__hx_this.next()
						}, func() any {
							return hx_receiver_758.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_754.map_ = func(hx_value_755 any) haxe__IMap {
						if hx_value_755 == nil {
							var hx_zero_756 haxe__IMap
							return hx_zero_756
						}
						return hx_value_755.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_754.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_759 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_759, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_759.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_760 := hx_structural_source_759
						return hxrt.BoundFunction(func() any {
							return hx_receiver_760.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_760.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_759.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_759.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_759.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_759.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_763 := hx_structural_source_759
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_763.__hx_this.get_message()
						}, func() any {
							return hx_receiver_763.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_764 := hx_structural_source_759
						return hxrt.BoundFunction(func() any {
							return hx_receiver_764.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_764.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_765 := hx_structural_source_759
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_765.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_765.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_766 := hx_structural_source_759
						return hxrt.BoundFunction(func() any {
							return hx_receiver_766.__hx_this.get_native()
						}, func() any {
							return hx_receiver_766.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_767 := hx_structural_source_759
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_767.__hx_this.toString()
						}, func() any {
							return hx_receiver_767.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_768 := hx_structural_source_759
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_768.__hx_this.details()
						}, func() any {
							return hx_receiver_768.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_759.savedMessage = func(hx_value_761 any) *string {
						if hx_value_761 == nil {
							var hx_zero_762 *string
							return hx_zero_762
						}
						return hx_value_761.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_759.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_759.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_759.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__PosException:
		return func(hx_structural_source_769 *haxe__exceptions__PosException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_769, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_769.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_770 := hx_structural_source_769
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_770.__hx_this.toString()
						}, func() any {
							return hx_receiver_770.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_769.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_769.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_769.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_769.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_773 := hx_structural_source_769
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_773.__hx_this.get_message()
						}, func() any {
							return hx_receiver_773.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_774 := hx_structural_source_769
						return hxrt.BoundFunction(func() any {
							return hx_receiver_774.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_774.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_775 := hx_structural_source_769
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_775.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_775.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_776 := hx_structural_source_769
						return hxrt.BoundFunction(func() any {
							return hx_receiver_776.__hx_this.get_native()
						}, func() any {
							return hx_receiver_776.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_777 := hx_structural_source_769
						return hxrt.BoundFunction(func() any {
							return hx_receiver_777.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_777.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_778 := hx_structural_source_769
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_778.__hx_this.details()
						}, func() any {
							return hx_receiver_778.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_769.savedMessage = func(hx_value_771 any) *string {
						if hx_value_771 == nil {
							var hx_zero_772 *string
							return hx_zero_772
						}
						return hx_value_771.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_769.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_769.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_769.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() bool] {
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
			return func(hx_value_779 any) bool {
				if hx_value_779 == nil {
					var hx_zero_780 bool
					return hx_zero_780
				}
				return hx_value_779.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_781 any) bool {
				if hx_value_781 == nil {
					var hx_zero_782 bool
					return hx_zero_782
				}
				return hx_value_781.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_783 any) bool {
				if hx_value_783 == nil {
					var hx_zero_784 bool
					return hx_zero_784
				}
				return hx_value_783.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_785 any) bool {
				if hx_value_785 == nil {
					var hx_zero_786 bool
					return hx_zero_786
				}
				return hx_value_785.(bool)
			}(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_787 any) bool {
				if hx_value_787 == nil {
					var hx_zero_788 bool
					return hx_zero_788
				}
				return hx_value_787.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_789 any) bool {
				if hx_value_789 == nil {
					var hx_zero_790 bool
					return hx_zero_790
				}
				return hx_value_789.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_791 any) bool {
				if hx_value_791 == nil {
					var hx_zero_792 bool
					return hx_zero_792
				}
				return hx_value_791.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_793 any) bool {
				if hx_value_793 == nil {
					var hx_zero_794 bool
					return hx_zero_794
				}
				return hx_value_793.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_795 any) bool {
				if hx_value_795 == nil {
					var hx_zero_796 bool
					return hx_zero_796
				}
				return hx_value_795.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_797 any) bool {
				if hx_value_797 == nil {
					var hx_zero_798 bool
					return hx_zero_798
				}
				return hx_value_797.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_799 any) bool {
				if hx_value_799 == nil {
					var hx_zero_800 bool
					return hx_zero_800
				}
				return hx_value_799.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_801 any) bool {
				if hx_value_801 == nil {
					var hx_zero_802 bool
					return hx_zero_802
				}
				return hx_value_801.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_803 any) bool {
				if hx_value_803 == nil {
					var hx_zero_804 bool
					return hx_zero_804
				}
				return hx_value_803.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() *string] {
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
			return func(hx_value_805 any) *string {
				if hx_value_805 == nil {
					var hx_zero_806 *string
					return hx_zero_806
				}
				return hx_value_805.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_807 any) *string {
				if hx_value_807 == nil {
					var hx_zero_808 *string
					return hx_zero_808
				}
				return hx_value_807.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_809 any) *string {
				if hx_value_809 == nil {
					var hx_zero_810 *string
					return hx_zero_810
				}
				return hx_value_809.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_811 any) *string {
				if hx_value_811 == nil {
					var hx_zero_812 *string
					return hx_zero_812
				}
				return hx_value_811.(*string)
			}(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_813 any) *string {
				if hx_value_813 == nil {
					var hx_zero_814 *string
					return hx_zero_814
				}
				return hx_value_813.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_815 any) *string {
				if hx_value_815 == nil {
					var hx_zero_816 *string
					return hx_zero_816
				}
				return hx_value_815.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_817 any) *string {
				if hx_value_817 == nil {
					var hx_zero_818 *string
					return hx_zero_818
				}
				return hx_value_817.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_819 any) *string {
				if hx_value_819 == nil {
					var hx_zero_820 *string
					return hx_zero_820
				}
				return hx_value_819.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_821 any) *string {
				if hx_value_821 == nil {
					var hx_zero_822 *string
					return hx_zero_822
				}
				return hx_value_821.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_823 any) *string {
				if hx_value_823 == nil {
					var hx_zero_824 *string
					return hx_zero_824
				}
				return hx_value_823.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_825 any) *string {
				if hx_value_825 == nil {
					var hx_zero_826 *string
					return hx_zero_826
				}
				return hx_value_825.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_827 any) *string {
				if hx_value_827 == nil {
					var hx_zero_828 *string
					return hx_zero_828
				}
				return hx_value_827.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_829 any) *string {
				if hx_value_829 == nil {
					var hx_zero_830 *string
					return hx_zero_830
				}
				return hx_value_829.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() any] {
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

func hx_recover_callable_3(value any) *hxrt.Function[func(*string)] {
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
			hxrt.FunctionCall(source)(func(hx_value_831 any) int {
				if hx_value_831 == nil {
					var hx_zero_832 int
					return hx_zero_832
				}
				return hx_value_831.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string)]:
		return source
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
		})
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_833 any) bool {
				if hx_value_833 == nil {
					var hx_zero_834 bool
					return hx_zero_834
				}
				return hx_value_833.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_835 any) float64 {
				if hx_value_835 == nil {
					var hx_zero_836 float64
					return hx_zero_836
				}
				return hx_value_835.(float64)
			}(arg_0))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_837 any) int {
				if hx_value_837 == nil {
					var hx_zero_838 int
					return hx_zero_838
				}
				return hx_value_837.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_839 any) int {
				if hx_value_839 == nil {
					var hx_zero_840 int
					return hx_zero_840
				}
				return hx_value_839.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_841 any) bool {
				if hx_value_841 == nil {
					var hx_zero_842 bool
					return hx_zero_842
				}
				return hx_value_841.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_843 any) *hxrt.SocketHandle {
				if hx_value_843 == nil {
					var hx_zero_844 *hxrt.SocketHandle
					return hx_zero_844
				}
				return hx_value_843.(*hxrt.SocketHandle)
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
			hxrt.FunctionCall(source)(func(hx_value_845 any) *hxrt.HttpExchange {
				if hx_value_845 == nil {
					var hx_zero_846 *hxrt.HttpExchange
					return hx_zero_846
				}
				return hx_value_845.(*hxrt.HttpExchange)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_8(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_847 any) int {
				if hx_value_847 == nil {
					var hx_zero_848 int
					return hx_zero_848
				}
				return hx_value_847.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_849 any) int {
				if hx_value_849 == nil {
					var hx_zero_850 int
					return hx_zero_850
				}
				return hx_value_849.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_851 any) int {
				if hx_value_851 == nil {
					var hx_zero_852 int
					return hx_zero_852
				}
				return hx_value_851.(int)
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

func hx_recover_callable_4(value any) *hxrt.Function[func(*haxe__io__Bytes)] {
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
			hxrt.FunctionCall(source)(func(hx_value_853 any) int {
				if hx_value_853 == nil {
					var hx_zero_854 int
					return hx_zero_854
				}
				return hx_value_853.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_855 any) *string {
				if hx_value_855 == nil {
					var hx_zero_856 *string
					return hx_zero_856
				}
				return hx_value_855.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return source
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_857 any) bool {
				if hx_value_857 == nil {
					var hx_zero_858 bool
					return hx_zero_858
				}
				return hx_value_857.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_859 any) float64 {
				if hx_value_859 == nil {
					var hx_zero_860 float64
					return hx_zero_860
				}
				return hx_value_859.(float64)
			}(arg_0))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_861 any) int {
				if hx_value_861 == nil {
					var hx_zero_862 int
					return hx_zero_862
				}
				return hx_value_861.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_863 any) int {
				if hx_value_863 == nil {
					var hx_zero_864 int
					return hx_zero_864
				}
				return hx_value_863.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_865 any) bool {
				if hx_value_865 == nil {
					var hx_zero_866 bool
					return hx_zero_866
				}
				return hx_value_865.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_867 any) *hxrt.SocketHandle {
				if hx_value_867 == nil {
					var hx_zero_868 *hxrt.SocketHandle
					return hx_zero_868
				}
				return hx_value_867.(*hxrt.SocketHandle)
			}(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_869 any) *string {
				if hx_value_869 == nil {
					var hx_zero_870 *string
					return hx_zero_870
				}
				return hx_value_869.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_871 any) *hxrt.HttpExchange {
				if hx_value_871 == nil {
					var hx_zero_872 *hxrt.HttpExchange
					return hx_zero_872
				}
				return hx_value_871.(*hxrt.HttpExchange)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(hx_recover_class_8(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_873 any) int {
				if hx_value_873 == nil {
					var hx_zero_874 int
					return hx_zero_874
				}
				return hx_value_873.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_875 any) int {
				if hx_value_875 == nil {
					var hx_zero_876 int
					return hx_zero_876
				}
				return hx_value_875.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_877 any) int {
				if hx_value_877 == nil {
					var hx_zero_878 int
					return hx_zero_878
				}
				return hx_value_877.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_879 any) *string {
				if hx_value_879 == nil {
					var hx_zero_880 *string
					return hx_zero_880
				}
				return hx_value_879.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_881 any) *string {
				if hx_value_881 == nil {
					var hx_zero_882 *string
					return hx_zero_882
				}
				return hx_value_881.(*string)
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

func hx_recover_callable_5(value any) *hxrt.Function[func(int)] {
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
			hxrt.FunctionCall(source)(func(hx_value_883 any) *string {
				if hx_value_883 == nil {
					var hx_zero_884 *string
					return hx_zero_884
				}
				return hx_value_883.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
		})
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_885 any) bool {
				if hx_value_885 == nil {
					var hx_zero_886 bool
					return hx_zero_886
				}
				return hx_value_885.(bool)
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
			hxrt.FunctionCall(source)(func(hx_value_887 any) bool {
				if hx_value_887 == nil {
					var hx_zero_888 bool
					return hx_zero_888
				}
				return hx_value_887.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_889 any) *hxrt.SocketHandle {
				if hx_value_889 == nil {
					var hx_zero_890 *hxrt.SocketHandle
					return hx_zero_890
				}
				return hx_value_889.(*hxrt.SocketHandle)
			}(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_891 any) *string {
				if hx_value_891 == nil {
					var hx_zero_892 *string
					return hx_zero_892
				}
				return hx_value_891.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_893 any) *hxrt.HttpExchange {
				if hx_value_893 == nil {
					var hx_zero_894 *hxrt.HttpExchange
					return hx_zero_894
				}
				return hx_value_893.(*hxrt.HttpExchange)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_8(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
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
			hxrt.FunctionCall(source)(func(hx_value_895 any) *string {
				if hx_value_895 == nil {
					var hx_zero_896 *string
					return hx_zero_896
				}
				return hx_value_895.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_897 any) *string {
				if hx_value_897 == nil {
					var hx_zero_898 *string
					return hx_zero_898
				}
				return hx_value_897.(*string)
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

func hx_recover_class_0(value any) *sys__net__Socket {
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

func hx_recover_class_1(value any) *haxe__Exception {
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

func hx_recover_class_2(value any) *haxe__io__Eof {
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

func hx_recover_class_3(value any) *haxe__io__Input {
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

func hx_recover_class_4(value any) *haxe__ds__StringMap {
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

func hx_recover_class_5(value any) *haxe__ValueException {
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

func hx_recover_class_6(value any) *haxe__io__BytesBuffer {
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

func hx_recover_class_7(value any) *haxe__io__Bytes {
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

func hx_recover_class_8(value any) *haxe___Int64_____Int64 {
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
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
