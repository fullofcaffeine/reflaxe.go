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

func New_CgoFeature() *CgoFeature {
	self := &CgoFeature{}
	self.FeatureBase = New_FeatureBase()
	self.FeatureBase.__hx_this = self
	self.__hx_this = self
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
