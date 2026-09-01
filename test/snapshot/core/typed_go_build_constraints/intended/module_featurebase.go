package main

import "snapshot/hxrt"

type I_FeatureBase interface {
	label() *string
}

type FeatureBase struct {
	__hx_this I_FeatureBase
	name      *string
}

func New_FeatureBase() *FeatureBase {
	self := &FeatureBase{}
	self.__hx_this = self
	self.name = hxrt.StringFromLiteral("common")
	return self
}

func (self *FeatureBase) label() *string {
	return self.name
}
