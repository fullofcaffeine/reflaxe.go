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

func New_Base(id int) *Base {
	self := &Base{}
	self.__hx_this = self
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

func New_Child() *Child {
	self := &Child{}
	self.Base = New_Base(7)
	self.Base.__hx_this = self
	self.__hx_this = self
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
	hxrt.Println(hxrt.StdString(child.superId()))
	hxrt.Println(hxrt.StdString(asBase.id))
	hxrt.Println(hxrt.StdString(nil))
	hxrt.Println(hxrt.StringFromLiteral("3"))
	hxrt.Println(hxrt.StdString(1.5))
	hxrt.Println(hxrt.StringFromLiteral("true"))
	hxrt.Println(hxrt.StringFromLiteral("ok"))
}
