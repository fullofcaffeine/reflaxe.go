//go:build alpha

package main

import (
	"image"
	"snapshot/hxrt"
)

type I_AlphaFeature interface {
}

type AlphaFeature struct {
	__hx_this I_AlphaFeature
}

func (self *AlphaFeature) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_AlphaFeature() *AlphaFeature {
	return hx_init_AlphaFeature(hx_alloc_AlphaFeature(nil))
}

func hx_alloc_AlphaFeature(owner any) *AlphaFeature {
	self := &AlphaFeature{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_AlphaFeature)
	}
	return self
}

func hx_init_AlphaFeature(self *AlphaFeature) *AlphaFeature {
	return self
}

func AlphaFeature_check() bool {
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
	source := hxrt.NewFunction(func(callback *hxrt.Function[func(*AlphaFeature) *string]) *string {
		return func() *string {
			var hx_callable_3 *hxrt.Function[func(*AlphaFeature) *string] = callback
			var hx_argument_4 *AlphaFeature = New_AlphaFeature()
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

var AlphaFeature_installed bool = AlphaFeature_check()
