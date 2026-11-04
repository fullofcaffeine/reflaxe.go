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

func (self *sys__thread__Tls) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__thread__Tls() *sys__thread__Tls {
	return hx_init_sys__thread__Tls(hx_alloc_sys__thread__Tls(nil))
}

func hx_alloc_sys__thread__Tls(owner any) *sys__thread__Tls {
	self := &sys__thread__Tls{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__thread__Tls)
	}
	return self
}

func hx_init_sys__thread__Tls(self *sys__thread__Tls) *sys__thread__Tls {
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
