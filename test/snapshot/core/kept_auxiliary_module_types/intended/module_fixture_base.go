package main

type I_fixture__Base interface {
}

type fixture__Base struct {
	__hx_this I_fixture__Base
}

func (self *fixture__Base) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_fixture__Base() *fixture__Base {
	return hx_init_fixture__Base(hx_alloc_fixture__Base(nil))
}

func hx_alloc_fixture__Base(owner any) *fixture__Base {
	self := &fixture__Base{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_fixture__Base)
	}
	return self
}

func hx_init_fixture__Base(self *fixture__Base) *fixture__Base {
	return self
}
