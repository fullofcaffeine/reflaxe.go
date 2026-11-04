package main

import "snapshot/hxrt"

type I_fixture__ClassWithToStringChild interface {
}

type fixture__ClassWithToStringChild struct {
	*fixture__Base
	__hx_this I_fixture__ClassWithToStringChild
}

func (self *fixture__ClassWithToStringChild) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_fixture__ClassWithToStringChild() *fixture__ClassWithToStringChild {
	return hx_init_fixture__ClassWithToStringChild(hx_alloc_fixture__ClassWithToStringChild(nil))
}

func hx_alloc_fixture__ClassWithToStringChild(owner any) *fixture__ClassWithToStringChild {
	self := &fixture__ClassWithToStringChild{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_fixture__ClassWithToStringChild)
	}
	self.fixture__Base = hx_alloc_fixture__Base(self.__hx_this)
	return self
}

func hx_init_fixture__ClassWithToStringChild(self *fixture__ClassWithToStringChild) *fixture__ClassWithToStringChild {
	hx_init_fixture__Base(self.fixture__Base)
	return self
}

type I_fixture__ClassWithToStringChild2 interface {
}

type fixture__ClassWithToStringChild2 struct {
	*fixture__Base
	__hx_this I_fixture__ClassWithToStringChild2
}

func (self *fixture__ClassWithToStringChild2) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_fixture__ClassWithToStringChild2() *fixture__ClassWithToStringChild2 {
	return hx_init_fixture__ClassWithToStringChild2(hx_alloc_fixture__ClassWithToStringChild2(nil))
}

func hx_alloc_fixture__ClassWithToStringChild2(owner any) *fixture__ClassWithToStringChild2 {
	self := &fixture__ClassWithToStringChild2{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_fixture__ClassWithToStringChild2)
	}
	self.fixture__Base = hx_alloc_fixture__Base(self.__hx_this)
	return self
}

func hx_init_fixture__ClassWithToStringChild2(self *fixture__ClassWithToStringChild2) *fixture__ClassWithToStringChild2 {
	hx_init_fixture__Base(self.fixture__Base)
	return self
}

type I_fixture__KeptAuxiliary interface {
	value() *string
}

type fixture__KeptAuxiliary struct {
	__hx_this I_fixture__KeptAuxiliary
}

func (self *fixture__KeptAuxiliary) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_fixture__KeptAuxiliary() *fixture__KeptAuxiliary {
	return hx_init_fixture__KeptAuxiliary(hx_alloc_fixture__KeptAuxiliary(nil))
}

func hx_alloc_fixture__KeptAuxiliary(owner any) *fixture__KeptAuxiliary {
	self := &fixture__KeptAuxiliary{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_fixture__KeptAuxiliary)
	}
	return self
}

func hx_init_fixture__KeptAuxiliary(self *fixture__KeptAuxiliary) *fixture__KeptAuxiliary {
	return self
}

func (self *fixture__KeptAuxiliary) value() *string {
	return hxrt.StringFromLiteral("kept")
}

type fixture__TestSpecification struct {
}
