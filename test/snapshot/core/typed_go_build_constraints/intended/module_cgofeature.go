//go:build cgo

package main

import "snapshot/hxrt"

type I_CgoFeature interface {
	label() *string
}

type CgoFeature struct {
	*FeatureBase
	__hx_this I_CgoFeature
	mode      *string
}

func (self *CgoFeature) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_CgoFeature() *CgoFeature {
	return hx_init_CgoFeature(hx_alloc_CgoFeature(nil))
}

func hx_alloc_CgoFeature(owner any) *CgoFeature {
	self := &CgoFeature{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_CgoFeature)
	}
	self.FeatureBase = hx_alloc_FeatureBase(self.__hx_this)
	return self
}

func hx_init_CgoFeature(self *CgoFeature) *CgoFeature {
	hx_init_FeatureBase(self.FeatureBase)
	self.mode = hxrt.StringFromLiteral("cgo")
	return self
}

func (self *CgoFeature) label() *string {
	return self.mode
}

var CgoFeature_installed bool = Registry_install(hxrt.StringFromLiteral("cgo"))

func CgoFeature_selected() *string {
	return hxrt.StringFromLiteral("cgo")
}
