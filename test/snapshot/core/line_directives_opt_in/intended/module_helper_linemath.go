package main

import "snapshot/hxrt"

type I_helper__LineMath interface {
	doubleIt(value int) int
}

type helper__LineMath struct {
	__hx_this I_helper__LineMath
}

func New_helper__LineMath() *helper__LineMath {
	self := &helper__LineMath{}
	self.__hx_this = self
	//line helper/LineMath.hx:4
	return self
}

func (self *helper__LineMath) doubleIt(value int) int {
	//line helper/LineMath.hx:6
	twice := int((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(2)))
	return twice
}
