package main

import "snapshot/hxrt"

type I_haxe__zip__Compress interface {
	execute(src *haxe__io__Bytes, srcPos int, dst *haxe__io__Bytes, dstPos int) map[string]any
	setFlushMode(f *haxe__zip__FlushMode)
	close()
	ensureOpen()
}

type haxe__zip__Compress struct {
	__hx_this I_haxe__zip__Compress
	handle    *hxrt.ZipDeflateHandle
	flushMode int
	closed    bool
}

func New_haxe__zip__Compress(level int) *haxe__zip__Compress {
	self := &haxe__zip__Compress{}
	self.__hx_this = self
	haxe__zip__Compress_validateLevel(level)
	self.handle = hxrt.ZipDeflateCreate(level)
	self.flushMode = hxrt.ZipFlushNo
	self.closed = false
	return self
}

func (self *haxe__zip__Compress) execute(src *haxe__io__Bytes, srcPos int, dst *haxe__io__Bytes, dstPos int) map[string]any {
	self.__hx_this.ensureOpen()
	haxe__zip__Compress_validatePosition(srcPos, src.length)
	haxe__zip__Compress_validatePosition(dstPos, dst.length)
	outputLimit := int((hxrt.Int32Wrap(dst.length) - hxrt.Int32Wrap(dstPos)))
	if outputLimit == 0 {
		hx_obj_1 := map[string]any{}
		hx_obj_1["done"] = false
		hx_obj_1["read"] = 0
		hx_obj_1["write"] = 0
		return hx_obj_1
	}
	step := hxrt.ZipDeflateExecute(self.handle, haxe__zip__Compress_toValuesFrom(src, srcPos), outputLimit, self.flushMode)
	write := haxe__zip__Compress_writeValues(dst, dstPos, step.Values)
	hx_obj_2 := map[string]any{}
	hx_obj_2["done"] = step.Done
	hx_obj_2["read"] = step.Read
	hx_obj_2["write"] = write
	return hx_obj_2
}

func (self *haxe__zip__Compress) setFlushMode(f *haxe__zip__FlushMode) {
	self.__hx_this.ensureOpen()
	self.flushMode = haxe__zip__Compress_flushModeCode(f)
}

func (self *haxe__zip__Compress) close() {
	if self.closed {
		return
	}
	self.closed = true
	hxrt.ZipDeflateClose(self.handle)
}

func (self *haxe__zip__Compress) ensureOpen() {
	if self.closed {
		hxrt.Throw(hxrt.StringFromLiteral("haxe.zip.Compress is closed"))
	}
}

func haxe__zip__Compress_flushModeCode(mode *haxe__zip__FlushMode) int {
	var hx_switch_3 int
	switch mode.tag {
	case 0:
		hx_switch_3 = hxrt.ZipFlushNo
	case 1:
		hx_switch_3 = hxrt.ZipFlushSync
	case 2:
		hx_switch_3 = func() int {
			hxrt.Throw(hxrt.StringFromLiteral("haxe.zip.FlushMode.FULL is not supported by Go's standard compressor"))
			var hx_throw_zero_4 int
			return hx_throw_zero_4
		}()
	case 3:
		hx_switch_3 = hxrt.ZipFlushFinish
	case 4:
		hx_switch_3 = func() int {
			hxrt.Throw(hxrt.StringFromLiteral("haxe.zip.FlushMode.BLOCK is not supported by Go's standard compressor"))
			var hx_throw_zero_5 int
			return hx_throw_zero_5
		}()
	}
	return hx_switch_3
}

func haxe__zip__Compress_fromValues(values []int) *haxe__io__Bytes {
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

func haxe__zip__Compress_run(s *haxe__io__Bytes, level int) *haxe__io__Bytes {
	haxe__zip__Compress_validateLevel(level)
	return haxe__zip__Compress_fromValues(hxrt.ZipCompress(haxe__zip__Compress_toValues(s), level))
}

func haxe__zip__Compress_toValues(bytes *haxe__io__Bytes) []int {
	return haxe__zip__Compress_toValuesFrom(bytes, 0)
}

func haxe__zip__Compress_toValuesFrom(bytes *haxe__io__Bytes, position int) []int {
	values := hxrt.NewArray()
	_g := position
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_7 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_7
		values.Push(bytes.b[index])
	}
	return func(hx_lambda_raw_9 []any) []int {
		hx_lambda_out_10 := make([]int, 0, len(hx_lambda_raw_9))
		for _, hx_lambda_item_11 := range hx_lambda_raw_9 {
			hx_lambda_out_10 = append(hx_lambda_out_10, func(hx_value_12 any) int {
				if hx_value_12 == nil {
					var hx_zero_13 int
					return hx_zero_13
				}
				return hx_value_12.(int)
			}(hx_lambda_item_11))
		}
		return hx_lambda_out_10
	}(values.Values())
}

func haxe__zip__Compress_validateLevel(level int) {
	if (level < -1) || (level > 9) {
		hxrt.Throw(hxrt.StringConcatAny(hxrt.StringFromLiteral("Invalid zlib compression level: "), level))
	}
}

func haxe__zip__Compress_validatePosition(position int, length int) {
	if (position < 0) || (position > length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
}

func haxe__zip__Compress_writeValues(destination *haxe__io__Bytes, position int, values []int) int {
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_14 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_14
		destination.b[int((hxrt.Int32Wrap(position) + hxrt.Int32Wrap(index)))] = int((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255)))
		destination.__hx_rawValid = false
	}
	return len(values)
}
