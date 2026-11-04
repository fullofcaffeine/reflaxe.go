//go:build beta

package main

import (
	"image"
	"snapshot/hxrt"
)

type I_BetaFeature interface {
}

type BetaFeature struct {
	__hx_this I_BetaFeature
}

func (self *BetaFeature) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_BetaFeature() *BetaFeature {
	return hx_init_BetaFeature(hx_alloc_BetaFeature(nil))
}

func hx_alloc_BetaFeature(owner any) *BetaFeature {
	self := &BetaFeature{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_BetaFeature)
	}
	return self
}

func hx_init_BetaFeature(self *BetaFeature) *BetaFeature {
	return self
}

func BetaFeature_check() bool {
	point := &image.Point{}
	point.X = 42
	native := hxrt.NewFunction(func(value *image.Point) int {
		return value.X
	})
	var erased any = native
	recovered := hx_recover_callable_0(erased)
	if func() int {
		var hx_callable_1 *hxrt.Function[func(any) int] = recovered
		var hx_argument_2 any = point
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}() != 42 {
		hxrt.Throw(hxrt.StringFromLiteral("native recovery failed"))
	}
	if !hxrt.HaxeEqual(reflaxe__go___internal__CompilerReflect_callMethod(erased, []any{point}), 42) {
		hxrt.Throw(hxrt.StringFromLiteral("native erased invocation failed"))
	}
	source := hxrt.NewFunction(func(callback *hxrt.Function[func(*BetaFeature) *string]) *string {
		return func() *string {
			var hx_callable_3 *hxrt.Function[func(*BetaFeature) *string] = callback
			var hx_argument_4 *BetaFeature = New_BetaFeature()
			return hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
		}()
	})
	var stored any = source
	nested := hx_recover_callable_1(stored)
	calls := 0
	result := func() *string {
		var hx_callable_5 *hxrt.Function[func(*hxrt.Function[func(any) *string]) *string] = nested
		var hx_argument_6 *hxrt.Function[func(any) *string] = hxrt.NewFunction(func(value any) *string {
			calls = int(int32((calls + 1)))
			return hxrt.StringFromLiteral("ok")
		})
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()
	if !hxrt.StringEqualStringPtr(result, hxrt.StringFromLiteral("ok")) || (calls != 1) {
		hxrt.Throw(hxrt.StringFromLiteral("nested callback recovery changed evaluation"))
	}
	return Registry_install()
}

var BetaFeature_installed bool = BetaFeature_check()
