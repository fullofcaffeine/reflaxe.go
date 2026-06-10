package main

import "snapshot/hxrt"

var haxe__io___UInt16Array__UInt16Array_Impl__BYTES_PER_ELEMENT int = 2

func haxe__io___UInt16Array__UInt16Array_Impl___new(elements int) *haxe__io__ArrayBufferViewImpl {
	var this1 *haxe__io__ArrayBufferViewImpl
	size := int(int32((hxrt.Int32Wrap(elements) * hxrt.Int32Wrap(2))))
	var this1_2 *haxe__io__ArrayBufferViewImpl
	this1_2 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(size), 0, size)
	this1_1 := this1_2
	this1 = this1_1
	return this1
}

func haxe__io___UInt16Array__UInt16Array_Impl__fromArray(a []int, pos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_17 int
	if length == nil {
		hx_if_17 = int(int32((hxrt.Int32Wrap(len(a)) - hxrt.Int32Wrap(pos))))
	} else {
		hx_if_17 = length.(int)
	}
	resolvedLength := hx_if_17
	if ((pos < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength)))) > len(a)) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		var hx_throw_zero_18 *haxe__io__ArrayBufferViewImpl
		return hx_throw_zero_18
	}
	var this1 *haxe__io__ArrayBufferViewImpl
	size := int(int32((hxrt.Int32Wrap(resolvedLength) * hxrt.Int32Wrap(2))))
	var this1_2 *haxe__io__ArrayBufferViewImpl
	this1_2 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(size), 0, size)
	this1_1 := this1_2
	this1 = this1_1
	out := this1
	_g := 0
	_g1 := resolvedLength
	for _g < _g1 {
		hx_post_19 := _g
		_g = int(int32((_g + 1)))
		idx := hx_post_19
		value := a[int(int32((hxrt.Int32Wrap(idx) + hxrt.Int32Wrap(pos))))]
		if (idx >= 0) && (idx < int(int32((hxrt.Int32Wrap(out.byteLength) >> uint(1))))) {
			_this := out.bytes
			pos_1 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(idx) << uint(1))))) + hxrt.Int32Wrap(out.byteOffset))))
			_this.b[pos_1] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
			_this.b[int(int32((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255))))
		} else {
			_ = 0
		}
	}
	return out
}

func haxe__io___UInt16Array__UInt16Array_Impl__fromBytes(bytes *haxe__io__Bytes, bytePos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_20 any
	if length == nil {
		hx_if_20 = nil
	} else {
		hx_if_20 = int(int32((hxrt.Int32Wrap(hxrt.IntFromNullableAny(length.(int))) << uint(1))))
	}
	var resolvedLength any = hx_if_20
	return haxe__io___UInt16Array__UInt16Array_Impl__fromData(func() *haxe__io__ArrayBufferViewImpl {
		this1 := haxe__io___ArrayBufferView__ArrayBufferView_Impl__fromBytes(bytes, bytePos, resolvedLength)
		return this1
	}())
}

func haxe__io___UInt16Array__UInt16Array_Impl__fromData(d *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return d
}

func haxe__io___UInt16Array__UInt16Array_Impl__get(this1 *haxe__io__ArrayBufferViewImpl, index int) int {
	_this := this1.bytes
	pos := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(index) << uint(1))))) + hxrt.Int32Wrap(this1.byteOffset))))
	return int(int32((hxrt.Int32Wrap(_this.b[pos]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1))))]) << uint(8))))))))
}

func haxe__io___UInt16Array__UInt16Array_Impl__getData(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

func haxe__io___UInt16Array__UInt16Array_Impl__get_length(this1 *haxe__io__ArrayBufferViewImpl) int {
	return int(int32((hxrt.Int32Wrap(this1.byteLength) >> uint(1))))
}

func haxe__io___UInt16Array__UInt16Array_Impl__get_view(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

var haxe__io___UInt16Array__UInt16Array_Impl__length int

func haxe__io___UInt16Array__UInt16Array_Impl__set(this1 *haxe__io__ArrayBufferViewImpl, index int, value int) int {
	if (index >= 0) && (index < int(int32((hxrt.Int32Wrap(this1.byteLength) >> uint(1))))) {
		_this := this1.bytes
		pos := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(index) << uint(1))))) + hxrt.Int32Wrap(this1.byteOffset))))
		_this.b[pos] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
		_this.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255))))
		return value
	}
	return 0
}

func haxe__io___UInt16Array__UInt16Array_Impl__sub(this1 *haxe__io__ArrayBufferViewImpl, begin int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_21 any
	if length == nil {
		hx_if_21 = nil
	} else {
		hx_if_21 = int(int32((hxrt.Int32Wrap(hxrt.IntFromNullableAny(length.(int))) << uint(1))))
	}
	var scaledLength any = hx_if_21
	return haxe__io___UInt16Array__UInt16Array_Impl__fromData(this1.sub(int(int32((hxrt.Int32Wrap(begin) << uint(1)))), scaledLength))
}

func haxe__io___UInt16Array__UInt16Array_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_22 any
	if begin == nil {
		hx_if_22 = nil
	} else {
		hx_if_22 = int(int32((hxrt.Int32Wrap(hxrt.IntFromNullableAny(begin.(int))) << uint(1))))
	}
	var scaledBegin any = hx_if_22
	var hx_if_23 any
	if end == nil {
		hx_if_23 = nil
	} else {
		hx_if_23 = int(int32((hxrt.Int32Wrap(hxrt.IntFromNullableAny(end.(int))) << uint(1))))
	}
	var scaledEnd any = hx_if_23
	return haxe__io___UInt16Array__UInt16Array_Impl__fromData(this1.subarray(scaledBegin, scaledEnd))
}

var haxe__io___UInt16Array__UInt16Array_Impl__view *haxe__io__ArrayBufferViewImpl
