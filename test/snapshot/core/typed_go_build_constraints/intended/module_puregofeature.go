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

func New_PureGoFeature() *PureGoFeature {
	self := &PureGoFeature{}
	self.FeatureBase = New_FeatureBase()
	self.FeatureBase.__hx_this = self
	self.__hx_this = self
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
