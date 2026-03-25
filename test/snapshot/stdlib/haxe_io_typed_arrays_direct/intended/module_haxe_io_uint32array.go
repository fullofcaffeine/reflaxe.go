package main

import "snapshot/hxrt"

var haxe__io___UInt32Array__UInt32Array_Impl__BYTES_PER_ELEMENT int = 4

func haxe__io___UInt32Array__UInt32Array_Impl___new(elements int) *haxe__io__ArrayBufferViewImpl {
	var this1 *haxe__io__ArrayBufferViewImpl
	size := int(int32((hxrt.Int32Wrap(elements) * hxrt.Int32Wrap(4))))
	var this1_2 *haxe__io__ArrayBufferViewImpl
	this1_2 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(size), 0, size)
	this1_1 := this1_2
	this1 = this1_1
	return this1
}

func haxe__io___UInt32Array__UInt32Array_Impl__fromArray(a []int, pos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_29 int
	if hxrt.AnyEqualsNull(length) {
		hx_if_29 = int(int32((hxrt.Int32Wrap(len(a)) - hxrt.Int32Wrap(pos))))
	} else {
		hx_if_29 = hxrt.IntFromNullableAny(func(hx_value_27 any) int {
			if hx_value_27 == nil {
				var hx_zero_28 int
				return hx_zero_28
			}
			return hx_value_27.(int)
		}(length))
	}
	resolvedLength := hx_if_29
	if ((pos < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength)))) > len(a)) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		var hx_throw_zero_30 *haxe__io__ArrayBufferViewImpl
		return hx_throw_zero_30
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
		hx_post_31 := _g
		_g = int(int32((_g + 1)))
		idx := hx_post_31
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

func haxe__io___UInt32Array__UInt32Array_Impl__fromBytes(bytes *haxe__io__Bytes, bytePos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_34 any
	if hxrt.AnyEqualsNull(length) {
		hx_if_34 = nil
	} else {
		hx_if_34 = int(int32((hxrt.Int32Wrap(func(hx_value_32 any) int {
			if hx_value_32 == nil {
				var hx_zero_33 int
				return hx_zero_33
			}
			return hx_value_32.(int)
		}(length)) << uint(2))))
	}
	var resolvedLength any = hx_if_34
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
	pos := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(index) << uint(2))))) + hxrt.Int32Wrap(this1.byteOffset))))
	return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[pos]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this.b[int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
}

func haxe__io___UInt32Array__UInt32Array_Impl__getData(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

func haxe__io___UInt32Array__UInt32Array_Impl__get_length(this1 *haxe__io__ArrayBufferViewImpl) int {
	return int(int32((hxrt.Int32Wrap(this1.byteLength) >> uint(2))))
}

func haxe__io___UInt32Array__UInt32Array_Impl__get_view(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

var haxe__io___UInt32Array__UInt32Array_Impl__length int

func haxe__io___UInt32Array__UInt32Array_Impl__set(this1 *haxe__io__ArrayBufferViewImpl, index int, value int) int {
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

func haxe__io___UInt32Array__UInt32Array_Impl__sub(this1 *haxe__io__ArrayBufferViewImpl, begin int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_37 any
	if hxrt.AnyEqualsNull(length) {
		hx_if_37 = nil
	} else {
		hx_if_37 = int(int32((hxrt.Int32Wrap(func(hx_value_35 any) int {
			if hx_value_35 == nil {
				var hx_zero_36 int
				return hx_zero_36
			}
			return hx_value_35.(int)
		}(length)) << uint(2))))
	}
	var scaledLength any = hx_if_37
	return haxe__io___UInt32Array__UInt32Array_Impl__fromData(this1.sub(int(int32((hxrt.Int32Wrap(begin) << uint(2)))), scaledLength))
}

func haxe__io___UInt32Array__UInt32Array_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_40 any
	if hxrt.AnyEqualsNull(begin) {
		hx_if_40 = nil
	} else {
		hx_if_40 = int(int32((hxrt.Int32Wrap(func(hx_value_38 any) int {
			if hx_value_38 == nil {
				var hx_zero_39 int
				return hx_zero_39
			}
			return hx_value_38.(int)
		}(begin)) << uint(2))))
	}
	var scaledBegin any = hx_if_40
	var hx_if_43 any
	if hxrt.AnyEqualsNull(end) {
		hx_if_43 = nil
	} else {
		hx_if_43 = int(int32((hxrt.Int32Wrap(func(hx_value_41 any) int {
			if hx_value_41 == nil {
				var hx_zero_42 int
				return hx_zero_42
			}
			return hx_value_41.(int)
		}(end)) << uint(2))))
	}
	var scaledEnd any = hx_if_43
	return haxe__io___UInt32Array__UInt32Array_Impl__fromData(this1.subarray(scaledBegin, scaledEnd))
}

var haxe__io___UInt32Array__UInt32Array_Impl__view *haxe__io__ArrayBufferViewImpl
