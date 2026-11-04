//go:build !cgo

package main

import "snapshot/hxrt"

type I_PureGoFeature interface {
	label() *string
}

type PureGoFeature struct {
	*FeatureBase
	__hx_this I_PureGoFeature
	mode      *string
}

func (self *PureGoFeature) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_PureGoFeature() *PureGoFeature {
	return hx_init_PureGoFeature(hx_alloc_PureGoFeature(nil))
}

func hx_alloc_PureGoFeature(owner any) *PureGoFeature {
	self := &PureGoFeature{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_PureGoFeature)
	}
	self.FeatureBase = hx_alloc_FeatureBase(self.__hx_this)
	return self
}

func hx_init_PureGoFeature(self *PureGoFeature) *PureGoFeature {
	hx_init_FeatureBase(self.FeatureBase)
	self.mode = hxrt.StringFromLiteral("pure-go")
	return self
}

func (self *PureGoFeature) label() *string {
	return self.mode
}

var PureGoFeature_installed bool = Registry_install(hxrt.StringFromLiteral("pure-go"))

func PureGoFeature_selected() *string {
	return hxrt.StringFromLiteral("pure-go")
}
