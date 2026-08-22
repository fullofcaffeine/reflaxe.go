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

func haxe__io___Float64Array__Float64Array_Impl__fromArray(a *hxrt.Array, pos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_1 int
	if length == nil {
		hx_if_1 = int(int32((hxrt.Int32Wrap(a.Len()) - hxrt.Int32Wrap(pos))))
	} else {
		hx_if_1 = length.(int)
	}
	resolvedLength := hx_if_1
	if ((pos < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength)))) > a.Len()) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
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
		hx_post_2 := _g
		_g = int(int32((_g + 1)))
		idx := hx_post_2
		value := func(hx_value_3 any) float64 {
			if hx_value_3 == nil {
				var hx_zero_4 float64
				return hx_zero_4
			}
			return hx_value_3.(float64)
		}(a.Get(int(int32((hxrt.Int32Wrap(idx) + hxrt.Int32Wrap(pos))))))
		if (idx >= 0) && (idx < int(int32((hxrt.Int32Wrap(out.byteLength) >> uint(3))))) {
			pos_1 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(idx) << uint(3))))) + hxrt.Int32Wrap(out.byteOffset))))
			bits := haxe__io__FPHelper_doubleToI64(value)
			_this := out.bytes
			value1 := bits.low
			_this.b[pos_1] = int(int32((hxrt.Int32Wrap(value1) & hxrt.Int32Wrap(255))))
			_this.__hx_rawValid = false
			_this.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value1) >> uint(8))))) & hxrt.Int32Wrap(255))))
			_this.__hx_rawValid = false
			_this.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value1) >> uint(16))))) & hxrt.Int32Wrap(255))))
			_this.__hx_rawValid = false
			_this.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
			_this.__hx_rawValid = false
			_this_1 := out.bytes
			pos_2 := int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(4))))
			value1_1 := bits.high
			_this_1.b[pos_2] = int(int32((hxrt.Int32Wrap(value1_1) & hxrt.Int32Wrap(255))))
			_this_1.__hx_rawValid = false
			_this_1.b[int(int32((hxrt.Int32Wrap(pos_2) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value1_1) >> uint(8))))) & hxrt.Int32Wrap(255))))
			_this_1.__hx_rawValid = false
			_this_1.b[int(int32((hxrt.Int32Wrap(pos_2) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value1_1) >> uint(16))))) & hxrt.Int32Wrap(255))))
			_this_1.__hx_rawValid = false
			_this_1.b[int(int32((hxrt.Int32Wrap(pos_2) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value1_1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
			_this_1.__hx_rawValid = false
		} else {
			_ = 0
		}
	}
	return out
}

func haxe__io___Float64Array__Float64Array_Impl__fromBytes(bytes *haxe__io__Bytes, bytePos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_5 any
	if length == nil {
		hx_if_5 = nil
	} else {
		hx_if_5 = int(int32((hxrt.Int32Wrap(length.(int)) << uint(3))))
	}
	var resolvedLength any = hx_if_5
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
		value1 := bits.low
		_this.b[pos] = int(int32((hxrt.Int32Wrap(value1) & hxrt.Int32Wrap(255))))
		_this.__hx_rawValid = false
		_this.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value1) >> uint(8))))) & hxrt.Int32Wrap(255))))
		_this.__hx_rawValid = false
		_this.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value1) >> uint(16))))) & hxrt.Int32Wrap(255))))
		_this.__hx_rawValid = false
		_this.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
		_this.__hx_rawValid = false
		_this_1 := this1.bytes
		pos_1 := int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(4))))
		value1_1 := bits.high
		_this_1.b[pos_1] = int(int32((hxrt.Int32Wrap(value1_1) & hxrt.Int32Wrap(255))))
		_this_1.__hx_rawValid = false
		_this_1.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value1_1) >> uint(8))))) & hxrt.Int32Wrap(255))))
		_this_1.__hx_rawValid = false
		_this_1.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value1_1) >> uint(16))))) & hxrt.Int32Wrap(255))))
		_this_1.__hx_rawValid = false
		_this_1.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value1_1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
		_this_1.__hx_rawValid = false
		return value
	}
	return 0
}

func haxe__io___Float64Array__Float64Array_Impl__sub(this1 *haxe__io__ArrayBufferViewImpl, begin int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_6 any
	if length == nil {
		hx_if_6 = nil
	} else {
		hx_if_6 = int(int32((hxrt.Int32Wrap(length.(int)) << uint(3))))
	}
	var scaledLength any = hx_if_6
	return haxe__io___Float64Array__Float64Array_Impl__fromData(this1.__hx_this.sub(int(int32((hxrt.Int32Wrap(begin) << uint(3)))), scaledLength))
}

func haxe__io___Float64Array__Float64Array_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_7 any
	if begin == nil {
		hx_if_7 = nil
	} else {
		hx_if_7 = int(int32((hxrt.Int32Wrap(begin.(int)) << uint(3))))
	}
	var scaledBegin any = hx_if_7
	var hx_if_8 any
	if end == nil {
		hx_if_8 = nil
	} else {
		hx_if_8 = int(int32((hxrt.Int32Wrap(end.(int)) << uint(3))))
	}
	var scaledEnd any = hx_if_8
	return haxe__io___Float64Array__Float64Array_Impl__fromData(this1.__hx_this.subarray(scaledBegin, scaledEnd))
}

var haxe__io___Float64Array__Float64Array_Impl__view *haxe__io__ArrayBufferViewImpl
