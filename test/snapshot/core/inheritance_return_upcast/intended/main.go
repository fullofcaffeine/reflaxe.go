package main

import "snapshot/hxrt"

type I_Base interface {
	tag() int
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

func (self *Base) tag() int {
	return 1
}

type I_Child interface {
	tag() int
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

func (self *Child) tag() int {
	return 2
}

func Factory_makeBase(flag bool) *Base {
	if flag {
		return func(hx_value *Child) *Base {
			if hx_value == nil {
				return nil
			}
			return hx_value.Base
		}(New_Child())
	}
	return New_Base()
}

func main() {
	asBase := Factory_makeBase(true)
	var v any = any(asBase.__hx_this.tag())
	hxrt.Println(v)
}
