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
	var hx_if_43 int
	if length == nil {
		hx_if_43 = int(int32((hxrt.Int32Wrap(bytes.length) - hxrt.Int32Wrap(pos))))
	} else {
		hx_if_43 = length.(int)
	}
	resolvedLength := hx_if_43
	if ((pos < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(resolvedLength)))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
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
	var hx_if_44 int
	if length == nil {
		hx_if_44 = int(int32((hxrt.Int32Wrap(self.byteLength) - hxrt.Int32Wrap(begin))))
	} else {
		hx_if_44 = length.(int)
	}
	resolvedLength := hx_if_44
	if ((begin < 0) || (resolvedLength < 0)) || (int(int32((hxrt.Int32Wrap(begin) + hxrt.Int32Wrap(resolvedLength)))) > self.byteLength) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	return New_haxe__io__ArrayBufferViewImpl(self.bytes, int(int32((hxrt.Int32Wrap(self.byteOffset) + hxrt.Int32Wrap(begin)))), resolvedLength)
}

func (self *haxe__io__ArrayBufferViewImpl) subarray(begin any, end any) *haxe__io__ArrayBufferViewImpl {
	var hx_if_45 int
	if begin == nil {
		hx_if_45 = 0
	} else {
		hx_if_45 = begin.(int)
	}
	resolvedBegin := hx_if_45
	var hx_if_46 int
	if end == nil {
		hx_if_46 = int(int32((hxrt.Int32Wrap(self.byteLength) - hxrt.Int32Wrap(resolvedBegin))))
	} else {
		hx_if_46 = end.(int)
	}
	resolvedEnd := hx_if_46
	return self.sub(resolvedBegin, int(int32((hxrt.Int32Wrap(resolvedEnd) - hxrt.Int32Wrap(resolvedBegin)))))
}
