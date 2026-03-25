package main

import "snapshot/hxrt"

type I_StringBuf interface {
}

type StringBuf struct {
	__hx_this I_StringBuf
	b         *string
}

func New_StringBuf() *StringBuf {
	self := &StringBuf{}
	self.__hx_this = self
	self.b = hxrt.StringFromLiteral("")
	return self
}
