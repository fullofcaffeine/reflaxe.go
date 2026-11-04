package main

import "snapshot/hxrt"

type I_Base interface {
	asBase() *Base
	baseId() int
}

type Base struct {
	__hx_this I_Base
	id        int
}

func (self *Base) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Base(id int) *Base {
	return hx_init_Base(hx_alloc_Base(nil), id)
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

func hx_init_Base(self *Base, id int) *Base {
	self.id = id
	return self
}

func (self *Base) asBase() *Base {
	return self
}

func (self *Base) baseId() int {
	return self.id
}

type I_Child interface {
	asBase() *Base
	baseId() int
	superId() int
	superAsBase() *Base
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
	hx_init_Base(self.Base, 7)
	return self
}

func (self *Child) superId() int {
	return self.Base.baseId()
}

func (self *Child) superAsBase() *Base {
	return self.Base.asBase()
}

func main() {
	child := New_Child()
	asBase := child.superAsBase()
	var v any = any(hxrt.StdString(child.superId()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(asBase.id))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StdString(nil))
	hxrt.Println(v_2)
	hxrt.Println(any(hxrt.StringFromLiteral("3")))
	var v_3 any = any(hxrt.StdString(1.5))
	hxrt.Println(v_3)
	hxrt.Println(any(hxrt.StringFromLiteral("true")))
	hxrt.Println(any(hxrt.StringFromLiteral("ok")))
}
