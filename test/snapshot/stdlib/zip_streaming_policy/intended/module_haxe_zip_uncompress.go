package main

import "snapshot/hxrt"

type I_haxe__zip__Uncompress interface {
	execute(src *haxe__io__Bytes, srcPos int, dst *haxe__io__Bytes, dstPos int) map[string]any
	setFlushMode(f *haxe__zip__FlushMode)
	close()
	ensureOpen()
}

type haxe__zip__Uncompress struct {
	__hx_this I_haxe__zip__Uncompress
	raw       bool
	handle    *hxrt.ZipInflateHandle
	flushMode int
	closed    bool
}

func New_haxe__zip__Uncompress(windowBits any) *haxe__zip__Uncompress {
	self := &haxe__zip__Uncompress{}
	self.__hx_this = self
	self.raw = ((windowBits != nil) && (hxrt.IntFromNullableAny(windowBits.(int)) < 0))
	self.handle = hxrt.ZipInflateCreate(self.raw)
	self.flushMode = hxrt.ZipFlushNo
	self.closed = false
	return self
}

func (self *haxe__zip__Uncompress) execute(src *haxe__io__Bytes, srcPos int, dst *haxe__io__Bytes, dstPos int) map[string]any {
	self.__hx_this.ensureOpen()
	haxe__zip__Uncompress_validatePosition(srcPos, src.length)
	haxe__zip__Uncompress_validatePosition(dstPos, dst.length)
	bufferSize := int((hxrt.Int32Wrap(dst.length) - hxrt.Int32Wrap(dstPos)))
	if bufferSize == 0 {
		hx_obj_1 := map[string]any{}
		hx_obj_1["done"] = false
		hx_obj_1["read"] = 0
		hx_obj_1["write"] = 0
		return hx_obj_1
	}
	step := hxrt.ZipInflateExecute(self.handle, haxe__zip__Uncompress_toValuesFrom(src, srcPos), bufferSize, self.flushMode)
	write := haxe__zip__Uncompress_writeValues(dst, dstPos, step.Values)
	hx_obj_2 := map[string]any{}
	hx_obj_2["done"] = step.Done
	hx_obj_2["read"] = step.Read
	hx_obj_2["write"] = write
	return hx_obj_2
}

func (self *haxe__zip__Uncompress) setFlushMode(f *haxe__zip__FlushMode) {
	self.__hx_this.ensureOpen()
	self.flushMode = haxe__zip__Uncompress_flushModeCode(f)
}

func (self *haxe__zip__Uncompress) close() {
	if self.closed {
		return
	}
	self.closed = true
	hxrt.ZipInflateClose(self.handle)
}

func (self *haxe__zip__Uncompress) ensureOpen() {
	if self.closed {
		hxrt.Throw(hxrt.StringFromLiteral("haxe.zip.Uncompress is closed"))
	}
}

func haxe__zip__Uncompress_flushModeCode(mode *haxe__zip__FlushMode) int {
	var hx_switch_3 int
	switch mode.tag {
	case 0:
		hx_switch_3 = hxrt.ZipFlushNo
	case 1:
		hx_switch_3 = hxrt.ZipFlushSync
	case 2:
		hx_switch_3 = func() int {
			hxrt.Throw(hxrt.StringFromLiteral("haxe.zip.FlushMode.FULL is not supported by Go's standard inflater"))
			var hx_throw_zero_4 int
			return hx_throw_zero_4
		}()
	case 3:
		hx_switch_3 = hxrt.ZipFlushFinish
	case 4:
		hx_switch_3 = func() int {
			hxrt.Throw(hxrt.StringFromLiteral("haxe.zip.FlushMode.BLOCK is not supported by Go's standard inflater"))
			var hx_throw_zero_5 int
			return hx_throw_zero_5
		}()
	}
	return hx_switch_3
}

func haxe__zip__Uncompress_fromValues(values []int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_6 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_6
		bytes.b[index] = int((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255)))
		bytes.__hx_rawValid = false
	}
	return bytes
}

func haxe__zip__Uncompress_run(src *haxe__io__Bytes, bufsize any) *haxe__io__Bytes {
	var hx_if_7 any
	if bufsize == nil {
		hx_if_7 = 65536
	} else {
		hx_if_7 = bufsize.(int)
	}
	var resolvedBufferSize any = hx_if_7
	if hxrt.IntFromNullableAny(resolvedBufferSize) <= 0 {
		hxrt.Throw(hxrt.StringConcatAny(hxrt.StringFromLiteral("Invalid zlib buffer size: "), resolvedBufferSize))
	}
	return haxe__zip__Uncompress_fromValues(hxrt.ZipUncompress(haxe__zip__Uncompress_toValues(src), false, hxrt.IntFromNullableAny(resolvedBufferSize)))
}

func haxe__zip__Uncompress_toValues(bytes *haxe__io__Bytes) []int {
	return haxe__zip__Uncompress_toValuesFrom(bytes, 0)
}

func haxe__zip__Uncompress_toValuesFrom(bytes *haxe__io__Bytes, position int) []int {
	values := hxrt.NewArray()
	_g := position
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_8 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_8
		values.Push(bytes.b[index])
	}
	return func(hx_lambda_raw_10 []any) []int {
		hx_lambda_out_11 := make([]int, 0, len(hx_lambda_raw_10))
		for _, hx_lambda_item_12 := range hx_lambda_raw_10 {
			hx_lambda_out_11 = append(hx_lambda_out_11, func(hx_value_13 any) int {
				if hx_value_13 == nil {
					var hx_zero_14 int
					return hx_zero_14
				}
				return hx_value_13.(int)
			}(hx_lambda_item_12))
		}
		return hx_lambda_out_11
	}(values.Values())
}

func haxe__zip__Uncompress_validatePosition(position int, length int) {
	if (position < 0) || (position > length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
}

func haxe__zip__Uncompress_writeValues(destination *haxe__io__Bytes, position int, values []int) int {
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_15 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_15
		destination.b[int((hxrt.Int32Wrap(position) + hxrt.Int32Wrap(index)))] = int((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255)))
		destination.__hx_rawValid = false
	}
	return len(values)
}
