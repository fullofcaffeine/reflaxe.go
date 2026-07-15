package main

import "snapshot/hxrt"

type I_sys__thread__Tls interface {
	get_value() any
	set_value(v any) any
}

type sys__thread__Tls struct {
	__hx_this I_sys__thread__Tls
	__handle  *hxrt.ThreadLocalHandle
	value     any
}

func New_sys__thread__Tls() *sys__thread__Tls {
	self := &sys__thread__Tls{}
	self.__hx_this = self
	self.__handle = hxrt.ThreadLocalNew()
	return self
}

func (self *sys__thread__Tls) get_value() any {
	return hxrt.ThreadLocalGet(self.__handle)
}

func (self *sys__thread__Tls) set_value(v any) any {
	hxrt.ThreadLocalSet(self.__handle, v)
	return v
}
