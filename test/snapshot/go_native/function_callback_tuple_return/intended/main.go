package main

import (
	"context"
	"snapshot/hxrt"
)

type I_CancelPair interface {
}

type CancelPair struct {
	__hx_this I_CancelPair
	context   context.Context
	cancel    *hxrt.Function[func()]
}

func (self *CancelPair) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_CancelPair(context context.Context, cancel *hxrt.Function[func()]) *CancelPair {
	self := &CancelPair{}
	self.__hx_this = self
	self.context = context
	self.cancel = cancel
	return self
}

func main() {
	pair := func() *CancelPair {
		hx_tuple_1, hx_tuple_2 := context.WithCancel(context.Background())
		return New_CancelPair(hx_tuple_1, func() *hxrt.Function[func()] {
			hx_extern_method_3 := hx_tuple_2
			if hx_extern_method_3 == nil {
				return nil
			}
			return hxrt.NewFunction(func() {
				hx_extern_method_3()
			})
		}())
	}()
	var v any = any((func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(pair.context.Err()) == nil))
	hxrt.Println(v)
	cancel := pair.cancel
	alias := cancel
	func() {
		var hx_callable_4 *hxrt.Function[func()] = cancel
		hxrt.FunctionCall(hx_callable_4)()
	}()
	func() {
		var hx_callable_5 *hxrt.Function[func()] = alias
		hxrt.FunctionCall(hx_callable_5)()
	}()
	var v_1 any = any((func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(pair.context.Err()) != nil))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.ReflectIsFunction(any(cancel)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.ReflectCompareMethods(any(cancel), any(alias)))
	hxrt.Println(v_3)
}
