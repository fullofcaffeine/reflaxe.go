package main

import "snapshot/hxrt"

type I_haxe__zip__Uncompress interface {
	execute(src *haxe__io__Bytes, srcPos int, dst *haxe__io__Bytes, dstPos int) map[string]any
	setFlushMode(f *haxe__zip__FlushMode)
	close()
}

type haxe__zip__Uncompress struct {
	__hx_this I_haxe__zip__Uncompress
	raw       bool
}

func New_haxe__zip__Uncompress(windowBits any) *haxe__zip__Uncompress {
	self := &haxe__zip__Uncompress{}
	self.__hx_this = self
	self.raw = ((windowBits != nil) && (hxrt.IntFromNullableAny(windowBits.(int)) < 0))
	return self
}

func (self *haxe__zip__Uncompress) execute(src *haxe__io__Bytes, srcPos int, dst *haxe__io__Bytes, dstPos int) map[string]any {
	input := src.__hx_this.sub(srcPos, int(int32((hxrt.Int32Wrap(src.length) - hxrt.Int32Wrap(srcPos)))))
	bufferSize := int(int32((hxrt.Int32Wrap(dst.length) - hxrt.Int32Wrap(dstPos))))
	if bufferSize <= 0 {
		hx_obj_5 := map[string]any{}
		hx_obj_5["done"] = false
		hx_obj_5["read"] = 0
		hx_obj_5["write"] = 0
		return hx_obj_5
	}
	data := haxe__zip__Uncompress_fromValues(hxrt.ZipUncompress(haxe__zip__Uncompress_toValues(input), self.raw, bufferSize))
	dst.__hx_this.blit(dstPos, data, 0, data.length)
	hx_obj_6 := map[string]any{}
	hx_obj_6["done"] = true
	hx_obj_6["read"] = input.length
	hx_obj_6["write"] = data.length
	return hx_obj_6
}

func (self *haxe__zip__Uncompress) setFlushMode(f *haxe__zip__FlushMode) {
}

func (self *haxe__zip__Uncompress) close() {
}

func haxe__zip__Uncompress_fromValues(values []int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_7 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_7
		bytes.b[index] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	return bytes
}

func haxe__zip__Uncompress_run(src *haxe__io__Bytes, bufsize any) *haxe__io__Bytes {
	var hx_if_8 any
	if bufsize == nil {
		hx_if_8 = 65536
	} else {
		hx_if_8 = bufsize.(int)
	}
	var resolvedBufferSize any = hx_if_8
	if hxrt.IntFromNullableAny(resolvedBufferSize) <= 0 {
		hxrt.Throw(hxrt.StringConcatAny(hxrt.StringFromLiteral("Invalid zlib buffer size: "), resolvedBufferSize))
	}
	return haxe__zip__Uncompress_fromValues(hxrt.ZipUncompress(haxe__zip__Uncompress_toValues(src), false, hxrt.IntFromNullableAny(resolvedBufferSize)))
}

func haxe__zip__Uncompress_toValues(bytes *haxe__io__Bytes) []int {
	values := hxrt.NewArray()
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_9 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_9
		values.Push(bytes.b[index])
	}
	return func(hx_lambda_raw_11 []any) []int {
		hx_lambda_out_12 := make([]int, 0, len(hx_lambda_raw_11))
		for _, hx_lambda_item_13 := range hx_lambda_raw_11 {
			hx_lambda_out_12 = append(hx_lambda_out_12, func(hx_value_14 any) int {
				if hx_value_14 == nil {
					var hx_zero_15 int
					return hx_zero_15
				}
				return hx_value_14.(int)
			}(hx_lambda_item_13))
		}
		return hx_lambda_out_12
	}(values.Values())
}
