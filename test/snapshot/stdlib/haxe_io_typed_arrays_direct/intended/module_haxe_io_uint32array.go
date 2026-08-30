package main

import "snapshot/hxrt"

var haxe__io___UInt32Array__UInt32Array_Impl__BYTES_PER_ELEMENT int = 4

func haxe__io___UInt32Array__UInt32Array_Impl___new(elements int) *haxe__io__ArrayBufferViewImpl {
	var this1 *haxe__io__ArrayBufferViewImpl
	size := int((hxrt.Int32Wrap(elements) * hxrt.Int32Wrap(4)))
	var this1_2 *haxe__io__ArrayBufferViewImpl
	this1_2 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(size), 0, size)
	this1_1 := this1_2
	this1 = this1_1
	return this1
}

func haxe__io___UInt32Array__UInt32Array_Impl__fromArray(a *hxrt.Array, pos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_1 int
	if length == nil {
		hx_if_1 = int((hxrt.Int32Wrap(a.Len()) - hxrt.Int32Wrap(pos)))
	} else {
		hx_if_1 = length.(int)
	}
	resolvedLength := hx_if_1
	if ((pos < 0) || (resolvedLength < 0)) || (int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength))) > a.Len()) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	var this1 *haxe__io__ArrayBufferViewImpl
	size := int((hxrt.Int32Wrap(resolvedLength) * hxrt.Int32Wrap(4)))
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
		value := hxrt.IntFromNullableAny(a.Get(int((hxrt.Int32Wrap(idx) + hxrt.Int32Wrap(pos)))))
		if (idx >= 0) && (idx < int((hxrt.Int32Wrap(out.byteLength) >> uint(2)))) {
			_this := out.bytes
			pos_1 := int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(idx) << uint(2)))) + hxrt.Int32Wrap(out.byteOffset)))
			_this.b[pos_1] = int((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))
			_this.__hx_rawValid = false
			_this.b[int((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(1)))] = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(value) >> uint(8)))) & hxrt.Int32Wrap(255)))
			_this.__hx_rawValid = false
			_this.b[int((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(2)))] = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(value) >> uint(16)))) & hxrt.Int32Wrap(255)))
			_this.__hx_rawValid = false
			_this.b[int((hxrt.Int32Wrap(pos_1) + hxrt.Int32Wrap(3)))] = int((hxrt.Int32Wrap(int(int32((uint32(hxrt.Int32Wrap(value)) >> uint(24))))) & hxrt.Int32Wrap(255)))
			_this.__hx_rawValid = false
		} else {
			_ = 0
		}
	}
	return out
}

func haxe__io___UInt32Array__UInt32Array_Impl__fromBytes(bytes *haxe__io__Bytes, bytePos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_3 any
	if length == nil {
		hx_if_3 = nil
	} else {
		hx_if_3 = int((hxrt.Int32Wrap(length.(int)) << uint(2)))
	}
	var resolvedLength any = hx_if_3
	return haxe__io___UInt32Array__UInt32Array_Impl__fromData(func() *haxe__io__ArrayBufferViewImpl {
		this1 := haxe__io___ArrayBufferView__ArrayBufferView_Impl__fromBytes(bytes, bytePos, resolvedLength)
		return this1
	}())
}

func haxe__io___UInt32Array__UInt32Array_Impl__fromData(d *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return d
}

func haxe__io___UInt32Array__UInt32Array_Impl__get(this1 *haxe__io__ArrayBufferViewImpl, index int) int {
	_this := this1.bytes
	pos := int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(index) << uint(2)))) + hxrt.Int32Wrap(this1.byteOffset)))
	return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(_this.b[pos]) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(_this.b[int((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1)))]) << uint(8))))))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(_this.b[int((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(2)))]) << uint(16))))))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(_this.b[int((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(3)))]) << uint(24))))))
}

func haxe__io___UInt32Array__UInt32Array_Impl__getData(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

func haxe__io___UInt32Array__UInt32Array_Impl__get_length(this1 *haxe__io__ArrayBufferViewImpl) int {
	return int((hxrt.Int32Wrap(this1.byteLength) >> uint(2)))
}

func haxe__io___UInt32Array__UInt32Array_Impl__get_view(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

var haxe__io___UInt32Array__UInt32Array_Impl__length int

func haxe__io___UInt32Array__UInt32Array_Impl__set(this1 *haxe__io__ArrayBufferViewImpl, index int, value int) int {
	if (index >= 0) && (index < int((hxrt.Int32Wrap(this1.byteLength) >> uint(2)))) {
		_this := this1.bytes
		pos := int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(index) << uint(2)))) + hxrt.Int32Wrap(this1.byteOffset)))
		_this.b[pos] = int((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))
		_this.__hx_rawValid = false
		_this.b[int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1)))] = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(value) >> uint(8)))) & hxrt.Int32Wrap(255)))
		_this.__hx_rawValid = false
		_this.b[int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2)))] = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(value) >> uint(16)))) & hxrt.Int32Wrap(255)))
		_this.__hx_rawValid = false
		_this.b[int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(3)))] = int((hxrt.Int32Wrap(int(int32((uint32(hxrt.Int32Wrap(value)) >> uint(24))))) & hxrt.Int32Wrap(255)))
		_this.__hx_rawValid = false
		return value
	}
	return 0
}

func haxe__io___UInt32Array__UInt32Array_Impl__sub(this1 *haxe__io__ArrayBufferViewImpl, begin int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_4 any
	if length == nil {
		hx_if_4 = nil
	} else {
		hx_if_4 = int((hxrt.Int32Wrap(length.(int)) << uint(2)))
	}
	var scaledLength any = hx_if_4
	return haxe__io___UInt32Array__UInt32Array_Impl__fromData(this1.__hx_this.sub(int((hxrt.Int32Wrap(begin) << uint(2))), scaledLength))
}

func haxe__io___UInt32Array__UInt32Array_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_5 any
	if begin == nil {
		hx_if_5 = nil
	} else {
		hx_if_5 = int((hxrt.Int32Wrap(begin.(int)) << uint(2)))
	}
	var scaledBegin any = hx_if_5
	var hx_if_6 any
	if end == nil {
		hx_if_6 = nil
	} else {
		hx_if_6 = int((hxrt.Int32Wrap(end.(int)) << uint(2)))
	}
	var scaledEnd any = hx_if_6
	return haxe__io___UInt32Array__UInt32Array_Impl__fromData(this1.__hx_this.subarray(scaledBegin, scaledEnd))
}

var haxe__io___UInt32Array__UInt32Array_Impl__view *haxe__io__ArrayBufferViewImpl
