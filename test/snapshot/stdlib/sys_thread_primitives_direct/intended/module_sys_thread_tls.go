package main

import "snapshot/hxrt"

type I_sys__thread__Tls interface {
	currentId() int
	get_value() any
	set_value(v any) any
}

type sys__thread__Tls struct {
	__hx_this I_sys__thread__Tls
	__mutex   *sys__thread__Mutex
	__values  *haxe__ds__IntMap
	value     any
}

func New_sys__thread__Tls() *sys__thread__Tls {
	self := &sys__thread__Tls{}
	self.__hx_this = self
	self.__values = New_haxe__ds__IntMap()
	self.__mutex = New_sys__thread__Mutex()
	return self
}

func (self *sys__thread__Tls) currentId() int {
	return hxrt.ThreadCurrentId()
}

func (self *sys__thread__Tls) get_value() any {
	id := hxrt.ThreadCurrentId()
	self.__mutex.acquire()
	var value any = self.__values.get(id)
	self.__mutex.release()
	return value
}

func (self *sys__thread__Tls) set_value(v any) any {
	id := hxrt.ThreadCurrentId()
	self.__mutex.acquire()
	if hxrt.AnyEqualsNull(v) {
		func(hx_value_19 any) bool {
			if hx_value_19 == nil {
				var hx_zero_20 bool
				return hx_zero_20
			}
			return hx_value_19.(bool)
		}(self.__values.remove(id))
	} else {
		self.__values.set(id, v)
	}
	self.__mutex.release()
	return v
}
