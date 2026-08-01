package main

type I_fixture__Base interface {
}

type fixture__Base struct {
	__hx_this I_fixture__Base
}

func New_fixture__Base() *fixture__Base {
	self := &fixture__Base{}
	self.__hx_this = self
	return self
}
