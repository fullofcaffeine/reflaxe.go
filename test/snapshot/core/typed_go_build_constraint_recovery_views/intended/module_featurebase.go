package main

import "snapshot/hxrt"

type I_FeatureBase interface {
}

type FeatureBase struct {
	__hx_this I_FeatureBase
	label     *string
}

func (self *FeatureBase) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_FeatureBase() *FeatureBase {
	return hx_init_FeatureBase(hx_alloc_FeatureBase(nil))
}

func hx_alloc_FeatureBase(owner any) *FeatureBase {
	self := &FeatureBase{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_FeatureBase)
	}
	return self
}

func hx_init_FeatureBase(self *FeatureBase) *FeatureBase {
	self.label = hxrt.StringFromLiteral("ok")
	return self
}
