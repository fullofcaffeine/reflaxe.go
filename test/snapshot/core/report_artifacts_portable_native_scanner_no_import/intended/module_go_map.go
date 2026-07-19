package main

import "snapshot/hxrt"

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
	self.inner.__hx_this.set(hxrt.StdString(key), value)
}

func (self *go___Map) get(key any) any {
	return self.inner.__hx_this.get(hxrt.StdString(key))
}

func (self *go___Map) exists(key any) bool {
	return func(hx_value_4 any) bool {
		if hx_value_4 == nil {
			var hx_zero_5 bool
			return hx_zero_5
		}
		return hx_value_4.(bool)
	}(self.inner.__hx_this.exists(hxrt.StdString(key)))
}
