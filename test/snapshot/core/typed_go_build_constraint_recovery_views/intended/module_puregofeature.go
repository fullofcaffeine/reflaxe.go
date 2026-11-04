//go:build !cgo

package main

import "snapshot/hxrt"

type I_PureGoFeature interface {
}

type PureGoFeature struct {
	*FeatureBase
	__hx_this I_PureGoFeature
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
	return self
}

func PureGoFeature_check() *string {
	var erased any = New_PureGoFeature()
	restored := hx_recover_class_0(erased)
	view := hx_structural_recover(hx_structural_recover(erased))
	func() *string {
		hx_structural_target_1 := view
		hx_structural_value_3 := hxrt.StringFromLiteral("changed")
		hxrt.AnonymousSet(hx_structural_target_1, "label", hx_structural_value_3)
		return hx_structural_value_3
	}()
	if !hxrt.StringEqualStringPtr(restored.label, hxrt.StringFromLiteral("changed")) {
		hxrt.Throw(hxrt.StringFromLiteral("structural recovery copied source"))
	}
	restored.label = hxrt.StringFromLiteral("ok")
	if !hxrt.StringEqualStringPtr(func(hx_value_4 any) *string {
		if hx_value_4 == nil {
			var hx_zero_5 *string
			return hx_zero_5
		}
		return hx_value_4.(*string)
	}(hxrt.AnonymousGet(view, "label")), hxrt.StringFromLiteral("ok")) {
		hxrt.Throw(hxrt.StringFromLiteral("base recovery copied source"))
	}
	source := hxrt.NewFunction(func(value *PureGoFeature) *string {
		return value.label
	})
	var callback any = source
	recovered := hx_recover_callable_0(callback)
	if !hxrt.StringEqualStringPtr(func() *string {
		var hx_callable_6 *hxrt.Function[func(any) *string] = recovered
		var hx_argument_7 any = erased
		return hxrt.FunctionCall(hx_callable_6)(hx_argument_7)
	}(), hxrt.StringFromLiteral("ok")) {
		hxrt.Throw(hxrt.StringFromLiteral("conditional callable projection failed"))
	}
	if !hxrt.StringEqualAny(reflaxe__go___internal__CompilerReflect_callMethod(callback, []any{erased}), hxrt.StringFromLiteral("ok")) {
		hxrt.Throw(hxrt.StringFromLiteral("conditional erased invocation failed"))
	}
	return restored.label
}

var PureGoFeature_installed bool = Registry_install(PureGoFeature_check())
