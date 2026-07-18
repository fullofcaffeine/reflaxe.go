package main

import "snapshot/hxrt"

type I_Base interface {
	who() *string
	callWho() *string
}

type Base struct {
	__hx_this I_Base
}

func New_Base() *Base {
	self := &Base{}
	self.__hx_this = self
	return self
}

func (self *Base) who() *string {
	return hxrt.StringFromLiteral("base")
}

func (self *Base) callWho() *string {
	return self.__hx_this.who()
}

type I_Leaf interface {
	who() *string
	callWho() *string
}

type Leaf struct {
	*Middle
	__hx_this     I_Leaf
	constructedAs *string
}

func New_Leaf() *Leaf {
	self := &Leaf{}
	self.Middle = New_Middle()
	self.Middle.Base.__hx_this = self
	self.Middle.__hx_this = self
	self.__hx_this = self
	self.constructedAs = self.__hx_this.callWho()
	return self
}

func (self *Leaf) who() *string {
	return hxrt.StringFromLiteral("leaf")
}

func main() {
	leaf := New_Leaf()
	middle := leaf.Middle
	base := leaf.Middle.Base
	boundBaseMethod := base.callWho
	var v any = any(leaf.constructedAs)
	hxrt.Println(v)
	var v_1 any = any(leaf.who())
	hxrt.Println(v_1)
	var v_2 any = any(middle.__hx_this.who())
	hxrt.Println(v_2)
	var v_3 any = any(base.__hx_this.who())
	hxrt.Println(v_3)
	var v_4 any = any(base.__hx_this.callWho())
	hxrt.Println(v_4)
	var v_5 any = any(boundBaseMethod())
	hxrt.Println(v_5)
}

type I_Middle interface {
	who() *string
	callWho() *string
}

type Middle struct {
	*Base
	__hx_this I_Middle
}

func New_Middle() *Middle {
	self := &Middle{}
	self.Base = New_Base()
	self.Base.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *Middle) who() *string {
	return hxrt.StringFromLiteral("middle")
}
