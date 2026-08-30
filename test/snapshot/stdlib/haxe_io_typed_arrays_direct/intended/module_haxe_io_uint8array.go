package main

import "snapshot/hxrt"

var haxe__io___UInt8Array__UInt8Array_Impl__BYTES_PER_ELEMENT int = 1

func haxe__io___UInt8Array__UInt8Array_Impl___new(elements int) *haxe__io__ArrayBufferViewImpl {
	var this1 *haxe__io__ArrayBufferViewImpl
	var this1_2 *haxe__io__ArrayBufferViewImpl
	this1_2 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(elements), 0, elements)
	this1_1 := this1_2
	this1 = this1_1
	return this1
}

func haxe__io___UInt8Array__UInt8Array_Impl__fromArray(a *hxrt.Array, pos int, length any) *haxe__io__ArrayBufferViewImpl {
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
	var this1_2 *haxe__io__ArrayBufferViewImpl
	this1_2 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(resolvedLength), 0, resolvedLength)
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
		if (idx >= 0) && (idx < out.byteLength) {
			_this := out.bytes
			pos_1 := int((hxrt.Int32Wrap(idx) + hxrt.Int32Wrap(out.byteOffset)))
			_this.b[pos_1] = int((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))
			_this.__hx_rawValid = false
		} else {
			_ = 0
		}
	}
	return out
}

func haxe__io___UInt8Array__UInt8Array_Impl__fromBytes(bytes *haxe__io__Bytes, bytePos int, length any) *haxe__io__ArrayBufferViewImpl {
	return haxe__io___UInt8Array__UInt8Array_Impl__fromData(func() *haxe__io__ArrayBufferViewImpl {
		this1 := haxe__io___ArrayBufferView__ArrayBufferView_Impl__fromBytes(bytes, bytePos, length)
		return this1
	}())
}

func haxe__io___UInt8Array__UInt8Array_Impl__fromData(d *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return d
}

func haxe__io___UInt8Array__UInt8Array_Impl__get(this1 *haxe__io__ArrayBufferViewImpl, index int) int {
	return this1.bytes.b[int((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(this1.byteOffset)))]
}

func haxe__io___UInt8Array__UInt8Array_Impl__getData(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

func haxe__io___UInt8Array__UInt8Array_Impl__get_length(this1 *haxe__io__ArrayBufferViewImpl) int {
	return this1.byteLength
}

func haxe__io___UInt8Array__UInt8Array_Impl__get_view(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

var haxe__io___UInt8Array__UInt8Array_Impl__length int

func haxe__io___UInt8Array__UInt8Array_Impl__set(this1 *haxe__io__ArrayBufferViewImpl, index int, value int) int {
	if (index >= 0) && (index < this1.byteLength) {
		_this := this1.bytes
		pos := int((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(this1.byteOffset)))
		_this.b[pos] = int((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255)))
		_this.__hx_rawValid = false
		return value
	}
	return 0
}

func haxe__io___UInt8Array__UInt8Array_Impl__sub(this1 *haxe__io__ArrayBufferViewImpl, begin int, length any) *haxe__io__ArrayBufferViewImpl {
	return haxe__io___UInt8Array__UInt8Array_Impl__fromData(this1.__hx_this.sub(begin, length))
}

func haxe__io___UInt8Array__UInt8Array_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	return haxe__io___UInt8Array__UInt8Array_Impl__fromData(this1.__hx_this.subarray(begin, end))
}

var haxe__io___UInt8Array__UInt8Array_Impl__view *haxe__io__ArrayBufferViewImpl
