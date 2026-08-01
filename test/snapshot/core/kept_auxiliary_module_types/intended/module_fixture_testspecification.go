package main

import "snapshot/hxrt"

type I_fixture__ClassWithToStringChild interface {
}

type fixture__ClassWithToStringChild struct {
	*fixture__Base
	__hx_this I_fixture__ClassWithToStringChild
}

func New_fixture__ClassWithToStringChild() *fixture__ClassWithToStringChild {
	self := &fixture__ClassWithToStringChild{}
	self.fixture__Base = New_fixture__Base()
	self.fixture__Base.__hx_this = self
	self.__hx_this = self
	return self
}

type I_fixture__ClassWithToStringChild2 interface {
}

type fixture__ClassWithToStringChild2 struct {
	*fixture__Base
	__hx_this I_fixture__ClassWithToStringChild2
}

func New_fixture__ClassWithToStringChild2() *fixture__ClassWithToStringChild2 {
	self := &fixture__ClassWithToStringChild2{}
	self.fixture__Base = New_fixture__Base()
	self.fixture__Base.__hx_this = self
	self.__hx_this = self
	return self
}

type I_fixture__KeptAuxiliary interface {
	value() *string
}

type fixture__KeptAuxiliary struct {
	__hx_this I_fixture__KeptAuxiliary
}

func New_fixture__KeptAuxiliary() *fixture__KeptAuxiliary {
	self := &fixture__KeptAuxiliary{}
	self.__hx_this = self
	return self
}

func (self *fixture__KeptAuxiliary) value() *string {
	return hxrt.StringFromLiteral("kept")
}

type fixture__TestSpecification struct {
}
