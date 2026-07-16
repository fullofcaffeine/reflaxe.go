package main

import "snapshot/hxrt"

func haxe__crypto__Md5_encode(value *string) *string {
	return hxrt.StdString(hxrt.CryptoMd5String(value))
}

func haxe__crypto__Md5_fromValues(values []int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_97 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_97
		bytes.b[index] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
	}
	return bytes
}

func haxe__crypto__Md5_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	return haxe__crypto__Md5_fromValues(hxrt.CryptoMd5Values(haxe__crypto__Md5_toValues(value)))
}

func haxe__crypto__Md5_toValues(bytes *haxe__io__Bytes) []int {
	values := []int{}
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_98 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_98
		values = append(values, bytes.b[index])
	}
	return values
}
