package main

import "snapshot/hxrt"

type I_haxe__ds__WeakMap interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
	remove(key any) bool
	keys() hxrt.Anonymous
	iterator() hxrt.Anonymous
	keyValueIterator() hxrt.Anonymous
	copy() *haxe__ds__WeakMap
	toString() *string
	clear()
}

type haxe__ds__WeakMap struct {
	__hx_this I_haxe__ds__WeakMap
}

func (self *haxe__ds__WeakMap) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__WeakMap() *haxe__ds__WeakMap {
	return hx_init_haxe__ds__WeakMap(hx_alloc_haxe__ds__WeakMap(nil))
}

func hx_alloc_haxe__ds__WeakMap(owner any) *haxe__ds__WeakMap {
	self := &haxe__ds__WeakMap{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__ds__WeakMap)
	}
	return self
}

func hx_init_haxe__ds__WeakMap(self *haxe__ds__WeakMap) *haxe__ds__WeakMap {
	hxrt.Throw(haxe__Exception_thrown(New_haxe__exceptions__NotImplementedException(hxrt.StringFromLiteral("Not implemented for this platform"), nil, func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("haxe/ds/WeakMap.hx")
		hx_obj_1["lineNumber"] = 39
		hx_obj_1["className"] = hxrt.StringFromLiteral("haxe.ds.WeakMap")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("new")
		return hxrt.AnonymousFromMap(hx_obj_1)
	}())))
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

func (self *haxe__ds__WeakMap) keys() hxrt.Anonymous {
	return nil
}

func (self *haxe__ds__WeakMap) iterator() hxrt.Anonymous {
	return nil
}

func (self *haxe__ds__WeakMap) keyValueIterator() hxrt.Anonymous {
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
