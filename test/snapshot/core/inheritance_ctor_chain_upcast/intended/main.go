package main

import "snapshot/hxrt"

type I_Base interface {
	read() int
}

type Base struct {
	__hx_this I_Base
	value     int
}

func (self *Base) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Base(value int) *Base {
	return hx_init_Base(hx_alloc_Base(nil), value)
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

func hx_init_Base(self *Base, value int) *Base {
	self.value = value
	return self
}

func (self *Base) read() int {
	return self.value
}

type I_Child interface {
	read() int
}

type Child struct {
	*Base
	__hx_this I_Child
}

func (self *Child) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Child(value int) *Child {
	return hx_init_Child(hx_alloc_Child(nil), value)
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

func hx_init_Child(self *Child, value int) *Child {
	hx_init_Base(self.Base, int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(1))))
	return self
}

func main() {
	child := New_Child(4)
	base := func(hx_value *Child) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Base
	}(child)
	show(func(hx_value *Child) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Base
	}(child))
	var v any = any(base.__hx_this.read())
	hxrt.Println(v)
}

func show(base *Base) {
	var v any = any(base.__hx_this.read())
	hxrt.Println(v)
}
