package main

import "snapshot/hxrt"

type I_StringBuf interface {
}

type StringBuf struct {
	__hx_this I_StringBuf
	b         *string
}

func (self *StringBuf) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_StringBuf() *StringBuf {
	return hx_init_StringBuf(hx_alloc_StringBuf(nil))
}

func hx_alloc_StringBuf(owner any) *StringBuf {
	self := &StringBuf{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_StringBuf)
	}
	return self
}

func hx_init_StringBuf(self *StringBuf) *StringBuf {
	self.b = hxrt.StringFromLiteral("")
	return self
}
