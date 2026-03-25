package main

import "snapshot/hxrt"

var haxe__io___Int32Array__Int32Array_Impl__BYTES_PER_ELEMENT int = 4

func haxe__io___Int32Array__Int32Array_Impl___new(elements int) *haxe__io__ArrayBufferViewImpl {
	var this1 *haxe__io__ArrayBufferViewImpl
	size := int(int32((hxrt.Int32Wrap(elements) * hxrt.Int32Wrap(4))))
	var this1_2 *haxe__io__ArrayBufferViewImpl
	this1_2 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(size), 0, size)
	this1_1 := this1_2
	this1 = this1_1
	return this1
}

func haxe__io___Int32Array__Int32Array_Impl__fromArray(a []int, pos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_63 int
	if hxrt.AnyEqualsNull(length) {
		hx_if_63 = int(int32((hxrt.Int32Wrap(len(a)) - hxrt.Int32Wrap(pos))))
	} else {
		hx_if_63 = hxrt.IntFromNullableAny(func(hx_value_61 any) int {
			if hx_value_61 == nil {
				var hx_zero_62 int
				return hx_zero_62
			}
			return hx_value_61.(int)
		}(length))
	}
	resolvedLength := hx_if_63
	if ((pos < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength)))) > len(a)) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		var hx_throw_zero_64 *haxe__io__ArrayBufferViewImpl
		return hx_throw_zero_64
	}
	var this1 *haxe__io__ArrayBufferViewImpl
	size := int(int32((hxrt.Int32Wrap(resolvedLength) * hxrt.Int32Wrap(4))))
	var this1_2 *haxe__io__ArrayBufferViewImpl
	this1_2 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(size), 0, size)
	this1_1 := this1_2
	this1 = this1_1
	out := this1
	_g := 0
	_g1 := resolvedLength
	for _g < _g1 {
		hx_post_65 := _g
		_g = int(int32((_g + 1)))
		idx := hx_post_65
		value := a[int(int32((hxrt.Int32Wrap(idx) + hxrt.Int32Wrap(pos))))]
		if (idx >= 0) && (idx < int(int32((hxrt.Int32Wrap(out.byteLength) >> uint(2))))) {
			_this := out.bytes
			pos_1 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(idx) << uint(2))))) + hxrt.Int32Wrap(out.byteOffset))))
			_this.b[pos_1] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
			_this.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255))))
			_this.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(16))))) & hxrt.Int32Wrap(255))))
			_this.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
		} else {
			_ = 0
		}
	}
	return out
}

func haxe__io___Int32Array__Int32Array_Impl__fromBytes(bytes *haxe__io__Bytes, bytePos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_68 any
	if hxrt.AnyEqualsNull(length) {
		hx_if_68 = nil
	} else {
		hx_if_68 = int(int32((hxrt.Int32Wrap(func(hx_value_66 any) int {
			if hx_value_66 == nil {
				var hx_zero_67 int
				return hx_zero_67
			}
			return hx_value_66.(int)
		}(length)) << uint(2))))
	}
	var resolvedLength any = hx_if_68
	return haxe__io___Int32Array__Int32Array_Impl__fromData(func() *haxe__io__ArrayBufferViewImpl {
		this1 := haxe__io___ArrayBufferView__ArrayBufferView_Impl__fromBytes(bytes, bytePos, resolvedLength)
		return this1
	}())
}

func haxe__io___Int32Array__Int32Array_Impl__fromData(d *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return d
}

func haxe__io___Int32Array__Int32Array_Impl__get(this1 *haxe__io__ArrayBufferViewImpl, index int) int {
	_this := this1.bytes
	pos := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(index) << uint(2))))) + hxrt.Int32Wrap(this1.byteOffset))))
	return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[pos]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
}

func haxe__io___Int32Array__Int32Array_Impl__getData(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

func haxe__io___Int32Array__Int32Array_Impl__get_length(this1 *haxe__io__ArrayBufferViewImpl) int {
	return int(int32((hxrt.Int32Wrap(this1.byteLength) >> uint(2))))
}

func haxe__io___Int32Array__Int32Array_Impl__get_view(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

var haxe__io___Int32Array__Int32Array_Impl__length int

func haxe__io___Int32Array__Int32Array_Impl__set(this1 *haxe__io__ArrayBufferViewImpl, index int, value int) int {
	if (index >= 0) && (index < int(int32((hxrt.Int32Wrap(this1.byteLength) >> uint(2))))) {
		_this := this1.bytes
		pos := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(index) << uint(2))))) + hxrt.Int32Wrap(this1.byteOffset))))
		_this.b[pos] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
		_this.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255))))
		_this.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(16))))) & hxrt.Int32Wrap(255))))
		_this.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
		return value
	}
	return 0
}

func haxe__io___Int32Array__Int32Array_Impl__sub(this1 *haxe__io__ArrayBufferViewImpl, begin int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_71 any
	if hxrt.AnyEqualsNull(length) {
		hx_if_71 = nil
	} else {
		hx_if_71 = int(int32((hxrt.Int32Wrap(func(hx_value_69 any) int {
			if hx_value_69 == nil {
				var hx_zero_70 int
				return hx_zero_70
			}
			return hx_value_69.(int)
		}(length)) << uint(2))))
	}
	var scaledLength any = hx_if_71
	return haxe__io___Int32Array__Int32Array_Impl__fromData(this1.sub(int(int32((hxrt.Int32Wrap(begin) << uint(2)))), scaledLength))
}

func haxe__io___Int32Array__Int32Array_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_74 any
	if hxrt.AnyEqualsNull(begin) {
		hx_if_74 = nil
	} else {
		hx_if_74 = int(int32((hxrt.Int32Wrap(func(hx_value_72 any) int {
			if hx_value_72 == nil {
				var hx_zero_73 int
				return hx_zero_73
			}
			return hx_value_72.(int)
		}(begin)) << uint(2))))
	}
	var scaledBegin any = hx_if_74
	var hx_if_77 any
	if hxrt.AnyEqualsNull(end) {
		hx_if_77 = nil
	} else {
		hx_if_77 = int(int32((hxrt.Int32Wrap(func(hx_value_75 any) int {
			if hx_value_75 == nil {
				var hx_zero_76 int
				return hx_zero_76
			}
			return hx_value_75.(int)
		}(end)) << uint(2))))
	}
	var scaledEnd any = hx_if_77
	return haxe__io___Int32Array__Int32Array_Impl__fromData(this1.subarray(scaledBegin, scaledEnd))
}

var haxe__io___Int32Array__Int32Array_Impl__view *haxe__io__ArrayBufferViewImpl
