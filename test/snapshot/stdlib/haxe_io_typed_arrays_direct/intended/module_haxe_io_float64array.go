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
	var hx_if_31 int
	if length == nil {
		hx_if_31 = int(int32((hxrt.Int32Wrap(len(a)) - hxrt.Int32Wrap(pos))))
	} else {
		hx_if_31 = length.(int)
	}
	resolvedLength := hx_if_31
	if ((pos < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength)))) > len(a)) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		var hx_throw_zero_32 *haxe__io__ArrayBufferViewImpl
		return hx_throw_zero_32
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
		hx_post_33 := _g
		_g = int(int32((_g + 1)))
		idx := hx_post_33
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
	var hx_if_34 any
	if length == nil {
		hx_if_34 = nil
	} else {
		hx_if_34 = int(int32((hxrt.Int32Wrap(length.(int)) << uint(3))))
	}
	var resolvedLength any = hx_if_34
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
	var hx_if_35 any
	if length == nil {
		hx_if_35 = nil
	} else {
		hx_if_35 = int(int32((hxrt.Int32Wrap(length.(int)) << uint(3))))
	}
	var scaledLength any = hx_if_35
	return haxe__io___Float64Array__Float64Array_Impl__fromData(this1.sub(int(int32((hxrt.Int32Wrap(begin) << uint(3)))), scaledLength))
}

func haxe__io___Float64Array__Float64Array_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_36 any
	if begin == nil {
		hx_if_36 = nil
	} else {
		hx_if_36 = int(int32((hxrt.Int32Wrap(begin.(int)) << uint(3))))
	}
	var scaledBegin any = hx_if_36
	var hx_if_37 any
	if end == nil {
		hx_if_37 = nil
	} else {
		hx_if_37 = int(int32((hxrt.Int32Wrap(end.(int)) << uint(3))))
	}
	var scaledEnd any = hx_if_37
	return haxe__io___Float64Array__Float64Array_Impl__fromData(this1.subarray(scaledBegin, scaledEnd))
}

var haxe__io___Float64Array__Float64Array_Impl__view *haxe__io__ArrayBufferViewImpl
