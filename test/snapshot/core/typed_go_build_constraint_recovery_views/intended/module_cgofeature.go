//go:build cgo

package main

import "snapshot/hxrt"

type I_CgoFeature interface {
}

type CgoFeature struct {
	*FeatureBase
	__hx_this I_CgoFeature
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
	return self
}

func CgoFeature_check() *string {
	var erased any = New_CgoFeature()
	restored := hx_recover_class_0(erased)
	view := hx_structural_recover(hx_structural_recover(erased))
	func() *string {
		hx_structural_target_1 := view
		hx_structural_value_3 := hxrt.StringFromLiteral("changed")
		hxrt.AnonymousSet(hx_structural_target_1, "label", hx_structural_value_3)
		return hx_structural_value_3
	}()
	if !hxrt.StringEqualStringPtr(restored.label, hxrt.StringFromLiteral("changed")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("structural recovery copied source")))
	}
	restored.label = hxrt.StringFromLiteral("ok")
	if !hxrt.StringEqualStringPtr(func(hx_value_4 any) *string {
		if hx_value_4 == nil {
			var hx_zero_5 *string
			return hx_zero_5
		}
		return hx_value_4.(*string)
	}(hxrt.AnonymousGet(view, "label")), hxrt.StringFromLiteral("ok")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("base recovery copied source")))
	}
	source := hxrt.NewFunction(func(value *CgoFeature) *string {
		return value.label
	})
	var callback any = source
	recovered := hx_recover_callable_0(callback)
	if !hxrt.StringEqualStringPtr(func() *string {
		var hx_callable_6 *hxrt.Function[func(any) *string] = recovered
		var hx_argument_7 any = erased
		return hxrt.FunctionCall(hx_callable_6)(hx_argument_7)
	}(), hxrt.StringFromLiteral("ok")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("conditional callable projection failed")))
	}
	if !hxrt.StringEqualAny(reflaxe__go___internal__CompilerReflect_callMethod(callback, []any{erased}), hxrt.StringFromLiteral("ok")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("conditional erased invocation failed")))
	}
	return restored.label
}

var CgoFeature_installed bool = Registry_install(CgoFeature_check())
