//go:build !cgo

package main

import "snapshot/hxrt"

type I_PureGoFeature interface {
}

type PureGoFeature struct {
	__hx_this I_PureGoFeature
	label     *string
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
	return self
}

func hx_init_PureGoFeature(self *PureGoFeature) *PureGoFeature {
	self.label = hxrt.StringFromLiteral("ok")
	return self
}

func PureGoFeature_check() *string {
	var erased any = New_PureGoFeature()
	restored := hx_recover_class_1(erased)
	source := hxrt.NewFunction(func(value *PureGoFeature) *string {
		return value.label
	})
	var callback any = source
	recovered := hx_recover_callable_1(callback)
	return func() *string {
		var hx_callable_1 *hxrt.Function[func(*PureGoFeature) *string] = recovered
		var hx_argument_2 *PureGoFeature = restored
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
}

var PureGoFeature_installed bool = Registry_install(PureGoFeature_check())
