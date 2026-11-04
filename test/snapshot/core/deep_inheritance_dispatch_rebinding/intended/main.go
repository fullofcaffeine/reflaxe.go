package main

import "snapshot/hxrt"

type I_Base interface {
	who() *string
	callWho() *string
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

func (self *Leaf) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Leaf() *Leaf {
	return hx_init_Leaf(hx_alloc_Leaf(nil))
}

func hx_alloc_Leaf(owner any) *Leaf {
	self := &Leaf{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Leaf)
	}
	self.Middle = hx_alloc_Middle(self.__hx_this)
	return self
}

func hx_init_Leaf(self *Leaf) *Leaf {
	hx_init_Middle(self.Middle)
	self.constructedAs = self.__hx_this.callWho()
	return self
}

func (self *Leaf) who() *string {
	return hxrt.StringFromLiteral("leaf")
}

func main() {
	leaf := New_Leaf()
	middle := func(hx_value *Leaf) *Middle {
		if hx_value == nil {
			return nil
		}
		return hx_value.Middle
	}(leaf)
	base := func(hx_value *Leaf) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Middle.Base
	}(leaf)
	boundBaseMethod := func() *hxrt.Function[func() *string] {
		hx_receiver_1 := base
		return hxrt.BoundFunction(func() *string {
			return hx_receiver_1.__hx_this.callWho()
		}, func() any {
			return hx_receiver_1.__hx_this
		}, "callWho")
	}()
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
	var v_5 any = any(func() *string {
		var hx_callable_2 *hxrt.Function[func() *string] = boundBaseMethod
		return hxrt.FunctionCall(hx_callable_2)()
	}())
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

func (self *Middle) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Middle() *Middle {
	return hx_init_Middle(hx_alloc_Middle(nil))
}

func hx_alloc_Middle(owner any) *Middle {
	self := &Middle{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Middle)
	}
	self.Base = hx_alloc_Base(self.__hx_this)
	return self
}

func hx_init_Middle(self *Middle) *Middle {
	hx_init_Base(self.Base)
	return self
}

func (self *Middle) who() *string {
	return hxrt.StringFromLiteral("middle")
}
