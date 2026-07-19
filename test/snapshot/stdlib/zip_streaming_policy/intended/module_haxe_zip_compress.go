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
	outputLimit := int(int32((hxrt.Int32Wrap(dst.length) - hxrt.Int32Wrap(dstPos))))
	if outputLimit == 0 {
		hx_obj_56 := map[string]any{}
		hx_obj_56["done"] = false
		hx_obj_56["read"] = 0
		hx_obj_56["write"] = 0
		return hx_obj_56
	}
	step := hxrt.ZipDeflateExecute(self.handle, haxe__zip__Compress_toValuesFrom(src, srcPos), outputLimit, self.flushMode)
	write := haxe__zip__Compress_writeValues(dst, dstPos, step.Values)
	hx_obj_57 := map[string]any{}
	hx_obj_57["done"] = step.Done
	hx_obj_57["read"] = step.Read
	hx_obj_57["write"] = write
	return hx_obj_57
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
	var hx_switch_58 int
	switch mode.tag {
	case 0:
		hx_switch_58 = hxrt.ZipFlushNo
	case 1:
		hx_switch_58 = hxrt.ZipFlushSync
	case 2:
		hx_switch_58 = func() int {
			hxrt.Throw(hxrt.StringFromLiteral("haxe.zip.FlushMode.FULL is not supported by Go's standard compressor"))
			var hx_throw_zero_59 int
			return hx_throw_zero_59
		}()
	case 3:
		hx_switch_58 = hxrt.ZipFlushFinish
	case 4:
		hx_switch_58 = func() int {
			hxrt.Throw(hxrt.StringFromLiteral("haxe.zip.FlushMode.BLOCK is not supported by Go's standard compressor"))
			var hx_throw_zero_60 int
			return hx_throw_zero_60
		}()
	}
	return hx_switch_58
}

func haxe__zip__Compress_fromValues(values []int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_61 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_61
		bytes.b[index] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
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
		hx_post_62 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_62
		values.Push(bytes.b[index])
	}
	return func(hx_lambda_raw_64 []any) []int {
		hx_lambda_out_65 := make([]int, 0, len(hx_lambda_raw_64))
		for _, hx_lambda_item_66 := range hx_lambda_raw_64 {
			hx_lambda_out_65 = append(hx_lambda_out_65, func(hx_value_67 any) int {
				if hx_value_67 == nil {
					var hx_zero_68 int
					return hx_zero_68
				}
				return hx_value_67.(int)
			}(hx_lambda_item_66))
		}
		return hx_lambda_out_65
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
		hx_post_69 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_69
		destination.b[int(int32((hxrt.Int32Wrap(position) + hxrt.Int32Wrap(index))))] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
		destination.__hx_rawValid = false
	}
	return len(values)
}
