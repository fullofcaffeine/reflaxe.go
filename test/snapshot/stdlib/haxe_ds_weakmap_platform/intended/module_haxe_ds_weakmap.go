package main

import "snapshot/hxrt"

type I_haxe__ds__WeakMap interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
	remove(key any) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
	copy() *haxe__ds__WeakMap
	toString() *string
	clear()
}

type haxe__ds__WeakMap struct {
	__hx_this I_haxe__ds__WeakMap
}

func New_haxe__ds__WeakMap() *haxe__ds__WeakMap {
	self := &haxe__ds__WeakMap{}
	self.__hx_this = self
	hxrt.Throw(New_haxe__exceptions__NotImplementedException(hxrt.StringFromLiteral("Not implemented for this platform"), nil, func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("haxe/ds/WeakMap.hx")
		hx_obj_1["lineNumber"] = 39
		hx_obj_1["className"] = hxrt.StringFromLiteral("haxe.ds.WeakMap")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("new")
		return hx_obj_1
	}()))
	return self
}

func (self *haxe__ds__WeakMap) set(key any, value any) {
}

func (self *haxe__ds__WeakMap) get(key any) any {
	return nil
}

func (self *haxe__ds__WeakMap) exists(key any) bool {
	return false
}

func (self *haxe__ds__WeakMap) remove(key any) bool {
	return false
}

func (self *haxe__ds__WeakMap) keys() map[string]any {
	return nil
}

func (self *haxe__ds__WeakMap) iterator() map[string]any {
	return nil
}

func (self *haxe__ds__WeakMap) keyValueIterator() map[string]any {
	return nil
}

func (self *haxe__ds__WeakMap) copy() *haxe__ds__WeakMap {
	return nil
}

func (self *haxe__ds__WeakMap) toString() *string {
	return nil
}

func (self *haxe__ds__WeakMap) clear() {
}

func (self *haxe__ds__WeakMap) String() string {
	return *self.__hx_this.toString()
}
