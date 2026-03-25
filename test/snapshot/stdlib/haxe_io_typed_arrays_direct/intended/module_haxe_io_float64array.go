package main

import "snapshot/hxrt"

var haxe__io___Float64Array__Float64Array_Impl__BYTES_PER_ELEMENT int = 8

func haxe__io___Float64Array__Float64Array_Impl___new(elements int) *haxe__io__ArrayBufferViewImpl {
	var this1 *haxe__io__ArrayBufferViewImpl
	size := int(int32((hxrt.Int32Wrap(elements) * hxrt.Int32Wrap(8))))
	var this1_2 *haxe__io__ArrayBufferViewImpl
	this1_2 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(size), 0, size)
	this1_1 := this1_2
	this1 = this1_1
	return this1
}

func haxe__io___Float64Array__Float64Array_Impl__fromArray(a []float64, pos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_80 int
	if hxrt.AnyEqualsNull(length) {
		hx_if_80 = int(int32((hxrt.Int32Wrap(len(a)) - hxrt.Int32Wrap(pos))))
	} else {
		hx_if_80 = hxrt.IntFromNullableAny(func(hx_value_78 any) int {
			if hx_value_78 == nil {
				var hx_zero_79 int
				return hx_zero_79
			}
			return hx_value_78.(int)
		}(length))
	}
	resolvedLength := hx_if_80
	if ((pos < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength)))) > len(a)) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		var hx_throw_zero_81 *haxe__io__ArrayBufferViewImpl
		return hx_throw_zero_81
	}
	var this1 *haxe__io__ArrayBufferViewImpl
	size := int(int32((hxrt.Int32Wrap(resolvedLength) * hxrt.Int32Wrap(8))))
	var this1_2 *haxe__io__ArrayBufferViewImpl
	this1_2 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(size), 0, size)
	this1_1 := this1_2
	this1 = this1_1
	out := this1
	_g := 0
	_g1 := resolvedLength
	for _g < _g1 {
		hx_post_82 := _g
		_g = int(int32((_g + 1)))
		idx := hx_post_82
		value := a[int(int32((hxrt.Int32Wrap(idx) + hxrt.Int32Wrap(pos))))]
		if (idx >= 0) && (idx < int(int32((hxrt.Int32Wrap(out.byteLength) >> uint(3))))) {
			pos_1 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(idx) << uint(3))))) + hxrt.Int32Wrap(out.byteOffset))))
			bits := haxe__io__FPHelper_doubleToI64(value)
			_this := out.bytes
			v := bits.low
			_this.b[pos_1] = int(int32((hxrt.Int32Wrap(v) & hxrt.Int32Wrap(255))))
			_this.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(v) >> uint(8))))) & hxrt.Int32Wrap(255))))
			_this.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(v) >> uint(16))))) & hxrt.Int32Wrap(255))))
			_this.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(v)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
			_this_1 := out.bytes
			pos_2 := int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(4))))
			v_1 := bits.high
			_this_1.b[pos_2] = int(int32((hxrt.Int32Wrap(v_1) & hxrt.Int32Wrap(255))))
			_this_1.b[int(int32((hxrt.Int32Wrap(pos_2) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(v_1) >> uint(8))))) & hxrt.Int32Wrap(255))))
			_this_1.b[int(int32((hxrt.Int32Wrap(pos_2) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(v_1) >> uint(16))))) & hxrt.Int32Wrap(255))))
			_this_1.b[int(int32((hxrt.Int32Wrap(pos_2) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(v_1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
		} else {
			_ = 0
		}
	}
	return out
}

func haxe__io___Float64Array__Float64Array_Impl__fromBytes(bytes *haxe__io__Bytes, bytePos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_85 any
	if hxrt.AnyEqualsNull(length) {
		hx_if_85 = nil
	} else {
		hx_if_85 = int(int32((hxrt.Int32Wrap(func(hx_value_83 any) int {
			if hx_value_83 == nil {
				var hx_zero_84 int
				return hx_zero_84
			}
			return hx_value_83.(int)
		}(length)) << uint(3))))
	}
	var resolvedLength any = hx_if_85
	return haxe__io___Float64Array__Float64Array_Impl__fromData(func() *haxe__io__ArrayBufferViewImpl {
		this1 := haxe__io___ArrayBufferView__ArrayBufferView_Impl__fromBytes(bytes, bytePos, resolvedLength)
		return this1
	}())
}

func haxe__io___Float64Array__Float64Array_Impl__fromData(d *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return d
}

func haxe__io___Float64Array__Float64Array_Impl__get(this1 *haxe__io__ArrayBufferViewImpl, index int) float64 {
	pos := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(index) << uint(3))))) + hxrt.Int32Wrap(this1.byteOffset))))
	_this := this1.bytes
	low := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[pos]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
	_this_1 := this1.bytes
	pos_1 := int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(4))))
	high := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_1.b[pos_1]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_1.b[int(int32((hxrt.Int32Wrap(pos_1)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_1.b[int(int32((hxrt.Int32Wrap(pos_1)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_1.b[int(int32((hxrt.Int32Wrap(pos_1)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
	return haxe__io__FPHelper_i64ToDouble(low, high)
}

func haxe__io___Float64Array__Float64Array_Impl__getData(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

func haxe__io___Float64Array__Float64Array_Impl__get_length(this1 *haxe__io__ArrayBufferViewImpl) int {
	return int(int32((hxrt.Int32Wrap(this1.byteLength) >> uint(3))))
}

func haxe__io___Float64Array__Float64Array_Impl__get_view(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

var haxe__io___Float64Array__Float64Array_Impl__length int

func haxe__io___Float64Array__Float64Array_Impl__set(this1 *haxe__io__ArrayBufferViewImpl, index int, value float64) float64 {
	if (index >= 0) && (index < int(int32((hxrt.Int32Wrap(this1.byteLength) >> uint(3))))) {
		pos := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(index) << uint(3))))) + hxrt.Int32Wrap(this1.byteOffset))))
		bits := haxe__io__FPHelper_doubleToI64(value)
		_this := this1.bytes
		v := bits.low
		_this.b[pos] = int(int32((hxrt.Int32Wrap(v) & hxrt.Int32Wrap(255))))
		_this.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(v) >> uint(8))))) & hxrt.Int32Wrap(255))))
		_this.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(v) >> uint(16))))) & hxrt.Int32Wrap(255))))
		_this.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(v)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
		_this_1 := this1.bytes
		pos_1 := int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(4))))
		v_1 := bits.high
		_this_1.b[pos_1] = int(int32((hxrt.Int32Wrap(v_1) & hxrt.Int32Wrap(255))))
		_this_1.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(v_1) >> uint(8))))) & hxrt.Int32Wrap(255))))
		_this_1.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(v_1) >> uint(16))))) & hxrt.Int32Wrap(255))))
		_this_1.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(v_1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
		return value
	}
	return 0
}

func haxe__io___Float64Array__Float64Array_Impl__sub(this1 *haxe__io__ArrayBufferViewImpl, begin int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_88 any
	if hxrt.AnyEqualsNull(length) {
		hx_if_88 = nil
	} else {
		hx_if_88 = int(int32((hxrt.Int32Wrap(func(hx_value_86 any) int {
			if hx_value_86 == nil {
				var hx_zero_87 int
				return hx_zero_87
			}
			return hx_value_86.(int)
		}(length)) << uint(3))))
	}
	var scaledLength any = hx_if_88
	return haxe__io___Float64Array__Float64Array_Impl__fromData(this1.sub(int(int32((hxrt.Int32Wrap(begin) << uint(3)))), scaledLength))
}

func haxe__io___Float64Array__Float64Array_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_91 any
	if hxrt.AnyEqualsNull(begin) {
		hx_if_91 = nil
	} else {
		hx_if_91 = int(int32((hxrt.Int32Wrap(func(hx_value_89 any) int {
			if hx_value_89 == nil {
				var hx_zero_90 int
				return hx_zero_90
			}
			return hx_value_89.(int)
		}(begin)) << uint(3))))
	}
	var scaledBegin any = hx_if_91
	var hx_if_94 any
	if hxrt.AnyEqualsNull(end) {
		hx_if_94 = nil
	} else {
		hx_if_94 = int(int32((hxrt.Int32Wrap(func(hx_value_92 any) int {
			if hx_value_92 == nil {
				var hx_zero_93 int
				return hx_zero_93
			}
			return hx_value_92.(int)
		}(end)) << uint(3))))
	}
	var scaledEnd any = hx_if_94
	return haxe__io___Float64Array__Float64Array_Impl__fromData(this1.subarray(scaledBegin, scaledEnd))
}

var haxe__io___Float64Array__Float64Array_Impl__view *haxe__io__ArrayBufferViewImpl
