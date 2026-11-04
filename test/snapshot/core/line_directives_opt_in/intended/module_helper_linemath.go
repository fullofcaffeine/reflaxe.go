package main

import "snapshot/hxrt"

type I_helper__LineMath interface {
	doubleIt(value int) int
}

type helper__LineMath struct {
	__hx_this I_helper__LineMath
}

func (self *helper__LineMath) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_helper__LineMath() *helper__LineMath {
	return hx_init_helper__LineMath(hx_alloc_helper__LineMath(nil))
}

func hx_alloc_helper__LineMath(owner any) *helper__LineMath {
	self := &helper__LineMath{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_helper__LineMath)
	}
	return self
}

func hx_init_helper__LineMath(self *helper__LineMath) *helper__LineMath {
	//line helper/LineMath.hx:4
	return self
}

func (self *helper__LineMath) doubleIt(value int) int {
	//line helper/LineMath.hx:6
	twice := int((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(2)))
	return twice
}
