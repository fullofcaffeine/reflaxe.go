package main

import "snapshot/hxrt"

type I_Base interface {
	who() int
	callWho() int
}

type Base struct {
	__hx_this I_Base
}

func (self *Base) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Base() *Base {
	return hx_init_Base(hx_alloc_Base(nil))
}

func hx_alloc_Base(owner any) *Base {
	self := &Base{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Base)
	}
	return self
}

func hx_init_Base(self *Base) *Base {
	return self
}

func (self *Base) who() int {
	return 1
}

func (self *Base) callWho() int {
	return self.__hx_this.who()
}

type I_Child interface {
	who() int
	callWho() int
}

type Child struct {
	*Base
	__hx_this I_Child
}

func (self *Child) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Child() *Child {
	return hx_init_Child(hx_alloc_Child(nil))
}

func hx_alloc_Child(owner any) *Child {
	self := &Child{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Child)
	}
	self.Base = hx_alloc_Base(self.__hx_this)
	return self
}

func hx_init_Child(self *Child) *Child {
	hx_init_Base(self.Base)
	return self
}

func (self *Child) who() int {
	return 2
}

func main() {
	child := New_Child()
	base := func(hx_value *Child) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Base
	}(child)
	var v any = any(base.__hx_this.callWho())
	hxrt.Println(v)
}
