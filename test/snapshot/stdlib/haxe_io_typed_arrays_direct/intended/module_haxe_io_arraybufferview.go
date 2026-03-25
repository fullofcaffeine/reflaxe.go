package main

import "snapshot/hxrt"

func haxe__io___ArrayBufferView__ArrayBufferView_Impl___new(size int) *haxe__io__ArrayBufferViewImpl {
	var this1 *haxe__io__ArrayBufferViewImpl
	this1 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(size), 0, size)
	return this1
}

var haxe__io___ArrayBufferView__ArrayBufferView_Impl__buffer *haxe__io__Bytes

var haxe__io___ArrayBufferView__ArrayBufferView_Impl__byteLength int

var haxe__io___ArrayBufferView__ArrayBufferView_Impl__byteOffset int

func haxe__io___ArrayBufferView__ArrayBufferView_Impl__fromBytes(bytes *haxe__io__Bytes, pos int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_114 int
	if hxrt.AnyEqualsNull(length) {
		hx_if_114 = int(int32((hxrt.Int32Wrap(bytes.length) - hxrt.Int32Wrap(pos))))
	} else {
		hx_if_114 = hxrt.IntFromNullableAny(func(hx_value_112 any) int {
			if hx_value_112 == nil {
				var hx_zero_113 int
				return hx_zero_113
			}
			return hx_value_112.(int)
		}(length))
	}
	resolvedLength := hx_if_114
	if ((pos < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength)))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		var hx_throw_zero_115 *haxe__io__ArrayBufferViewImpl
		return hx_throw_zero_115
	}
	a := New_haxe__io__ArrayBufferViewImpl(bytes, pos, resolvedLength)
	return a
}

func haxe__io___ArrayBufferView__ArrayBufferView_Impl__fromData(a *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return a
}

func haxe__io___ArrayBufferView__ArrayBufferView_Impl__getData(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__ArrayBufferViewImpl {
	return this1
}

func haxe__io___ArrayBufferView__ArrayBufferView_Impl__get_buffer(this1 *haxe__io__ArrayBufferViewImpl) *haxe__io__Bytes {
	return this1.bytes
}

func haxe__io___ArrayBufferView__ArrayBufferView_Impl__get_byteLength(this1 *haxe__io__ArrayBufferViewImpl) int {
	return this1.byteLength
}

func haxe__io___ArrayBufferView__ArrayBufferView_Impl__get_byteOffset(this1 *haxe__io__ArrayBufferViewImpl) int {
	return this1.byteOffset
}

func haxe__io___ArrayBufferView__ArrayBufferView_Impl__sub(this1 *haxe__io__ArrayBufferViewImpl, begin int, length any) *haxe__io__ArrayBufferViewImpl {
	a := this1.sub(begin, length)
	return a
}

func haxe__io___ArrayBufferView__ArrayBufferView_Impl__subarray(this1 *haxe__io__ArrayBufferViewImpl, begin any, end any) *haxe__io__ArrayBufferViewImpl {
	a := this1.subarray(begin, end)
	return a
}

type I_haxe__io__ArrayBufferViewImpl interface {
	sub(begin int, length any) *haxe__io__ArrayBufferViewImpl
	subarray(begin any, end any) *haxe__io__ArrayBufferViewImpl
}

type haxe__io__ArrayBufferViewImpl struct {
	__hx_this  I_haxe__io__ArrayBufferViewImpl
	bytes      *haxe__io__Bytes
	byteOffset int
	byteLength int
}

func New_haxe__io__ArrayBufferViewImpl(bytes *haxe__io__Bytes, pos int, length int) *haxe__io__ArrayBufferViewImpl {
	self := &haxe__io__ArrayBufferViewImpl{}
	self.__hx_this = self
	self.bytes = bytes
	self.byteOffset = pos
	self.byteLength = length
	return self
}

func (self *haxe__io__ArrayBufferViewImpl) sub(begin int, length any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_118 int
	if hxrt.AnyEqualsNull(length) {
		hx_if_118 = int(int32((hxrt.Int32Wrap(self.byteLength) - hxrt.Int32Wrap(begin))))
	} else {
		hx_if_118 = hxrt.IntFromNullableAny(func(hx_value_116 any) int {
			if hx_value_116 == nil {
				var hx_zero_117 int
				return hx_zero_117
			}
			return hx_value_116.(int)
		}(length))
	}
	resolvedLength := hx_if_118
	if ((begin < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(begin) + hxrt.Int32Wrap(resolvedLength)))) > self.byteLength) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		var hx_throw_zero_119 *haxe__io__ArrayBufferViewImpl
		return hx_throw_zero_119
	}
	return New_haxe__io__ArrayBufferViewImpl(self.bytes, int(int32((hxrt.Int32Wrap(self.byteOffset) + hxrt.Int32Wrap(begin)))), resolvedLength)
}

func (self *haxe__io__ArrayBufferViewImpl) subarray(begin any, end any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_122 int
	if hxrt.AnyEqualsNull(begin) {
		hx_if_122 = 0
	} else {
		hx_if_122 = hxrt.IntFromNullableAny(func(hx_value_120 any) int {
			if hx_value_120 == nil {
				var hx_zero_121 int
				return hx_zero_121
			}
			return hx_value_120.(int)
		}(begin))
	}
	resolvedBegin := hx_if_122
	var hx_if_125 int
	if hxrt.AnyEqualsNull(end) {
		hx_if_125 = int(int32((hxrt.Int32Wrap(self.byteLength) - hxrt.Int32Wrap(resolvedBegin))))
	} else {
		hx_if_125 = hxrt.IntFromNullableAny(func(hx_value_123 any) int {
			if hx_value_123 == nil {
				var hx_zero_124 int
				return hx_zero_124
			}
			return hx_value_123.(int)
		}(end))
	}
	resolvedEnd := hx_if_125
	return self.sub(resolvedBegin, int(int32((hxrt.Int32Wrap(resolvedEnd) - hxrt.Int32Wrap(resolvedBegin)))))
}
