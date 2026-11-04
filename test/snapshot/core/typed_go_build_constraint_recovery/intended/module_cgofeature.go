//go:build cgo

package main

import "snapshot/hxrt"

type I_CgoFeature interface {
}

type CgoFeature struct {
	__hx_this I_CgoFeature
	label     *string
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
	return self
}

func hx_init_CgoFeature(self *CgoFeature) *CgoFeature {
	self.label = hxrt.StringFromLiteral("ok")
	return self
}

func CgoFeature_check() *string {
	var erased any = New_CgoFeature()
	restored := hx_recover_class_0(erased)
	source := hxrt.NewFunction(func(value *CgoFeature) *string {
		return value.label
	})
	var callback any = source
	recovered := hx_recover_callable_0(callback)
	return func() *string {
		var hx_callable_1 *hxrt.Function[func(*CgoFeature) *string] = recovered
		var hx_argument_2 *CgoFeature = restored
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
}

var CgoFeature_installed bool = Registry_install(CgoFeature_check())
