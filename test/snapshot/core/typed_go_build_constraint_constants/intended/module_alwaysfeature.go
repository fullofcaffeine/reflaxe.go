package main

type I_AlwaysFeature interface {
}

type AlwaysFeature struct {
	__hx_this I_AlwaysFeature
}

func (self *AlwaysFeature) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_AlwaysFeature() *AlwaysFeature {
	return hx_init_AlwaysFeature(hx_alloc_AlwaysFeature(nil))
}

func hx_alloc_AlwaysFeature(owner any) *AlwaysFeature {
	self := &AlwaysFeature{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_AlwaysFeature)
	}
	return self
}

func hx_init_AlwaysFeature(self *AlwaysFeature) *AlwaysFeature {
	return self
}
