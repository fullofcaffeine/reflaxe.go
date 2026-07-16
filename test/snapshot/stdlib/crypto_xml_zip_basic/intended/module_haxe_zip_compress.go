package main

import "snapshot/hxrt"

type I_haxe__zip__Compress interface {
	execute(src *haxe__io__Bytes, srcPos int, dst *haxe__io__Bytes, dstPos int) map[string]any
	setFlushMode(f *haxe__zip__FlushMode)
	close()
}

type haxe__zip__Compress struct {
	__hx_this I_haxe__zip__Compress
	level     int
}

func New_haxe__zip__Compress(level int) *haxe__zip__Compress {
	self := &haxe__zip__Compress{}
	self.__hx_this = self
	haxe__zip__Compress_validateLevel(level)
	self.level = level
	return self
}

func (self *haxe__zip__Compress) execute(src *haxe__io__Bytes, srcPos int, dst *haxe__io__Bytes, dstPos int) map[string]any {
	input := src.sub(srcPos, int(int32((hxrt.Int32Wrap(src.length) - hxrt.Int32Wrap(srcPos)))))
	data := haxe__zip__Compress_fromValues(hxrt.ZipCompress(haxe__zip__Compress_toValues(input), self.level))
	dst.blit(dstPos, data, 0, data.length)
	hx_obj_12 := map[string]any{}
	hx_obj_12["done"] = true
	hx_obj_12["read"] = input.length
	hx_obj_12["write"] = data.length
	return hx_obj_12
}

func (self *haxe__zip__Compress) setFlushMode(f *haxe__zip__FlushMode) {
}

func (self *haxe__zip__Compress) close() {
}

func haxe__zip__Compress_fromValues(values []int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_13 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_13
		bytes.b[index] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
	}
	return bytes
}

func haxe__zip__Compress_run(s *haxe__io__Bytes, level int) *haxe__io__Bytes {
	haxe__zip__Compress_validateLevel(level)
	return haxe__zip__Compress_fromValues(hxrt.ZipCompress(haxe__zip__Compress_toValues(s), level))
}

func haxe__zip__Compress_toValues(bytes *haxe__io__Bytes) []int {
	values := []int{}
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_14 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_14
		values = append(values, bytes.b[index])
	}
	return values
}

func haxe__zip__Compress_validateLevel(level int) {
	if (level < -1) || (level > 9) {
		hxrt.Throw(hxrt.StringConcatAny(hxrt.StringFromLiteral("Invalid zlib compression level: "), level))
	}
}
