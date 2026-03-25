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
	var hx_if_46 int
	if hxrt.AnyEqualsNull(length) {
		hx_if_46 = int(int32((hxrt.Int32Wrap(len(a)) - hxrt.Int32Wrap(pos))))
	} else {
		hx_if_46 = hxrt.IntFromNullableAny(func(hx_value_44 any) int {
			if hx_value_44 == nil {
				var hx_zero_45 int
				return hx_zero_45
			}
			return hx_value_44.(int)
		}(length))
	}
	resolvedLength := hx_if_46
	if ((pos < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength)))) > len(a)) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		var hx_throw_zero_47 *haxe__io__ArrayBufferViewImpl
		return hx_throw_zero_47
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
		hx_post_48 := _g
		_g = int(int32((_g + 1)))
		idx := hx_post_48
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
	var hx_if_51 any
	if hxrt.AnyEqualsNull(length) {
		hx_if_51 = nil
	} else {
		hx_if_51 = int(int32((hxrt.Int32Wrap(func(hx_value_49 any) int {
			if hx_value_49 == nil {
				var hx_zero_50 int
				return hx_zero_50
			}
			return hx_value_49.(int)
		}(length)) << uint(1))))
	}
	var resolvedLength any = hx_if_51
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
	var hx_if_54 any
	if hxrt.AnyEqualsNull(length) {
		hx_if_54 = nil
	} else {
		hx_if_54 = int(int32((hxrt.Int32Wrap(func(hx_value_52 any) int {
			if hx_value_52 == nil {
				var hx_zero_53 int
				return hx_zero_53
			}
			return hx_value_52.(int)
		}(length)) << uint(1))))
	}
	var scaledLength any = hx_if_54
	return haxe__io___UInt16Array__UInt16Array_Impl__fromData(this1.sub(int(int32((hxrt.Int32Wrap(begin) << uint(1)))), scaledLength))
}

func haxe__io___UInt16Array__UInt16Array_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_57 any
	if hxrt.AnyEqualsNull(begin) {
		hx_if_57 = nil
	} else {
		hx_if_57 = int(int32((hxrt.Int32Wrap(func(hx_value_55 any) int {
			if hx_value_55 == nil {
				var hx_zero_56 int
				return hx_zero_56
			}
			return hx_value_55.(int)
		}(begin)) << uint(1))))
	}
	var scaledBegin any = hx_if_57
	var hx_if_60 any
	if hxrt.AnyEqualsNull(end) {
		hx_if_60 = nil
	} else {
		hx_if_60 = int(int32((hxrt.Int32Wrap(func(hx_value_58 any) int {
			if hx_value_58 == nil {
				var hx_zero_59 int
				return hx_zero_59
			}
			return hx_value_58.(int)
		}(end)) << uint(1))))
	}
	var scaledEnd any = hx_if_60
	return haxe__io___UInt16Array__UInt16Array_Impl__fromData(this1.subarray(scaledBegin, scaledEnd))
}

var haxe__io___UInt16Array__UInt16Array_Impl__view *haxe__io__ArrayBufferViewImpl
