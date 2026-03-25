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

func haxe__io___UInt8Array__UInt8Array_Impl__fromArray(a []int, pos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_24 int
	if hxrt.AnyEqualsNull(length) {
		hx_if_24 = int(int32((hxrt.Int32Wrap(len(a)) - hxrt.Int32Wrap(pos))))
	} else {
		hx_if_24 = hxrt.IntFromNullableAny(func(hx_value_22 any) int {
			if hx_value_22 == nil {
				var hx_zero_23 int
				return hx_zero_23
			}
			return hx_value_22.(int)
		}(length))
	}
	resolvedLength := hx_if_24
	if ((pos < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength)))) > len(a)) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		var hx_throw_zero_25 *haxe__io__ArrayBufferViewImpl
		return hx_throw_zero_25
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
		hx_post_26 := _g
		_g = int(int32((_g + 1)))
		idx := hx_post_26
		value := a[int(int32((hxrt.Int32Wrap(idx) + hxrt.Int32Wrap(pos))))]
		if (idx >= 0) && (idx < out.byteLength) {
			_this := out.bytes
			pos_1 := int(int32((hxrt.Int32Wrap(idx) + hxrt.Int32Wrap(out.byteOffset))))
			_this.b[pos_1] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
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
	return this1.bytes.b[int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(this1.byteOffset))))]
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
		pos := int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(this1.byteOffset))))
		_this.b[pos] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
		return value
	}
	return 0
}

func haxe__io___UInt8Array__UInt8Array_Impl__sub(this1 *haxe__io__ArrayBufferViewImpl, begin int, length any) *haxe__io__ArrayBufferViewImpl {
	return haxe__io___UInt8Array__UInt8Array_Impl__fromData(this1.sub(begin, length))
}

func haxe__io___UInt8Array__UInt8Array_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	return haxe__io___UInt8Array__UInt8Array_Impl__fromData(this1.subarray(begin, end))
}

var haxe__io___UInt8Array__UInt8Array_Impl__view *haxe__io__ArrayBufferViewImpl
