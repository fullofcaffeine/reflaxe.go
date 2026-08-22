package main

import "snapshot/hxrt"

type I_sys__io___Process__ProcessOutput interface {
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

type sys__io___Process__ProcessOutput struct {
	*haxe__io__Input
	__hx_this I_sys__io___Process__ProcessOutput
	handle    *hxrt.ProcessOutput
}

func New_sys__io___Process__ProcessOutput(handle *hxrt.ProcessOutput) *sys__io___Process__ProcessOutput {
	self := &sys__io___Process__ProcessOutput{}
	self.haxe__io__Input = New_haxe__io__Input()
	self.haxe__io__Input.__hx_this = self
	self.__hx_this = self
	self.handle = handle
	return self
}

func (self *sys__io___Process__ProcessOutput) readByte() int {
	if self.handle == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Process output is closed"))
	}
	value := hxrt.ProcessOutputReadByteValue(self.handle)
	if value < 0 {
		hxrt.Throw(New_haxe__io__Eof())
	}
	return value
}

func (self *sys__io___Process__ProcessOutput) readBytes(bytes *haxe__io__Bytes, pos int, length int) int {
	if ((pos < 0) || (length < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(length)))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	if length == 0 {
		return 0
	}
	if self.handle == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Process output is closed"))
	}
	values := hxrt.ProcessOutputReadValues(self.handle, length)
	if len(values) == 0 {
		hxrt.Throw(New_haxe__io__Eof())
	}
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_1
		bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	return len(values)
}

func (self *sys__io___Process__ProcessOutput) close() {
	if self.handle == nil {
		return
	}
	hxrt.ProcessOutputClose(self.handle)
	self.handle = nil
}

type I_sys__io___Process__ProcessInput interface {
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

type sys__io___Process__ProcessInput struct {
	*haxe__io__Output
	__hx_this I_sys__io___Process__ProcessInput
	handle    *hxrt.ProcessInput
}

func New_sys__io___Process__ProcessInput(handle *hxrt.ProcessInput) *sys__io___Process__ProcessInput {
	self := &sys__io___Process__ProcessInput{}
	self.haxe__io__Output = New_haxe__io__Output()
	self.haxe__io__Output.__hx_this = self
	self.__hx_this = self
	self.handle = handle
	return self
}

func (self *sys__io___Process__ProcessInput) writeByte(value int) {
	if (self.handle == nil) || !hxrt.ProcessInputWriteByteValue(self.handle, value) {
		hxrt.Throw(New_haxe__io__Eof())
	}
}

func (self *sys__io___Process__ProcessInput) writeBytes(bytes *haxe__io__Bytes, pos int, length int) int {
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
		hx_post_2 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_2
		values.Push(bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))])
	}
	if (self.handle == nil) || !hxrt.ProcessInputWriteValues(self.handle, func(hx_lambda_raw_4 []any) []int {
		hx_lambda_out_5 := make([]int, 0, len(hx_lambda_raw_4))
		for _, hx_lambda_item_6 := range hx_lambda_raw_4 {
			hx_lambda_out_5 = append(hx_lambda_out_5, func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(hx_lambda_item_6))
		}
		return hx_lambda_out_5
	}(values.Values())) {
		hxrt.Throw(New_haxe__io__Eof())
	}
	return length
}

func (self *sys__io___Process__ProcessInput) flush() {
	if self.handle != nil {
		hxrt.ProcessInputFlush(self.handle)
	}
}

func (self *sys__io___Process__ProcessInput) close() {
	if self.handle == nil {
		return
	}
	hxrt.ProcessInputClose(self.handle)
	self.handle = nil
}

type I_sys__io__Process interface {
	getPid() int
	exitCode(block bool) any
	close()
	kill()
	requireHandle() *hxrt.Process
}

type sys__io__Process struct {
	__hx_this I_sys__io__Process
	stdout    *haxe__io__Input
	stderr    *haxe__io__Input
	stdin     *haxe__io__Output
	handle    *hxrt.Process
}

func New_sys__io__Process(cmd *string, args *hxrt.Array, detached bool) *sys__io__Process {
	self := &sys__io__Process{}
	self.__hx_this = self
	if detached {
		hxrt.Throw(hxrt.StringFromLiteral("Detached process is not supported on this platform"))
	}
	self.handle = hxrt.ProcessCreate(cmd, func() []*string {
		var hx_if_14 []*string
		if args == nil {
			hx_if_14 = nil
		} else {
			hx_if_14 = func(hx_lambda_raw_9 []any) []*string {
				hx_lambda_out_10 := make([]*string, 0, len(hx_lambda_raw_9))
				for _, hx_lambda_item_11 := range hx_lambda_raw_9 {
					hx_lambda_out_10 = append(hx_lambda_out_10, func(hx_value_12 any) *string {
						if hx_value_12 == nil {
							var hx_zero_13 *string
							return hx_zero_13
						}
						return hx_value_12.(*string)
					}(hx_lambda_item_11))
				}
				return hx_lambda_out_10
			}(args.Values())
		}
		return hx_if_14
	}())
	self.stdout = New_sys__io___Process__ProcessOutput(hxrt.ProcessStdout(self.handle)).haxe__io__Input
	self.stderr = New_sys__io___Process__ProcessOutput(hxrt.ProcessStderr(self.handle)).haxe__io__Input
	self.stdin = New_sys__io___Process__ProcessInput(hxrt.ProcessStdin(self.handle)).haxe__io__Output
	return self
}

func (self *sys__io__Process) getPid() int {
	return hxrt.ProcessPid(self.__hx_this.requireHandle())
}

func (self *sys__io__Process) exitCode(block bool) any {
	status := hxrt.ProcessExitStatusValue(self.__hx_this.requireHandle(), block)
	var hx_if_15 any
	if status.Available {
		hx_if_15 = status.Code
	} else {
		hx_if_15 = nil
	}
	return hx_if_15
}

func (self *sys__io__Process) close() {
	if self.handle == nil {
		return
	}
	hxrt.ProcessClose(self.handle)
	self.handle = nil
}

func (self *sys__io__Process) kill() {
	hxrt.ProcessKill(self.__hx_this.requireHandle())
}

func (self *sys__io__Process) requireHandle() *hxrt.Process {
	if self.handle == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Process is closed"))
	}
	return self.handle
}
