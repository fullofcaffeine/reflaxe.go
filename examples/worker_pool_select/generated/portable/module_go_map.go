package main

import "examples_worker_pool_select_portable/hxrt"

type I_go___Map interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
}

type go___Map struct {
	__hx_this I_go___Map
	inner     *haxe__ds__StringMap
}

func New_go___Map() *go___Map {
	self := &go___Map{}
	self.__hx_this = self
	self.inner = New_haxe__ds__StringMap()
	return self
}

func (self *go___Map) set(key any, value any) {
	self.inner.set(hxrt.StdString(key), value)
}

func (self *go___Map) get(key any) any {
	return self.inner.get(hxrt.StdString(key))
}

func (self *go___Map) exists(key any) bool {
	return func(hx_value_17 any) bool {
		if hx_value_17 == nil {
			var hx_zero_18 bool
			return hx_zero_18
		}
		return hx_value_17.(bool)
	}(self.inner.exists(hxrt.StdString(key)))
}
