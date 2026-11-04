package main

import (
	"io/fs"
	"reflect"
	"snapshot/hxrt"
)

type I_ErrorMatch interface {
}

type ErrorMatch struct {
	__hx_this I_ErrorMatch
	value     any
	matched   bool
}

func (self *ErrorMatch) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_ErrorMatch(value any, matched bool) *ErrorMatch {
	return hx_init_ErrorMatch(hx_alloc_ErrorMatch(nil), value, matched)
}

func hx_alloc_ErrorMatch(owner any) *ErrorMatch {
	self := &ErrorMatch{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_ErrorMatch)
	}
	return self
}

func hx_init_ErrorMatch(self *ErrorMatch, value any, matched bool) *ErrorMatch {
	self.value = value
	self.matched = matched
	return self
}

func main() {
	NativeErrorProbe_check()
	var scalar reflect.Type = reflect.TypeFor[int]()
	var v any = any(hxrt.StdString(scalar.String()))
	hxrt.Println(v)
	var named reflect.Type = reflect.TypeFor[*fs.PathError]()
	var v_1 any = any(hxrt.StdString(named.String()))
	hxrt.Println(v_1)
	factory := hxrt.NewFunction(func() reflect.Type {
		return reflect.TypeFor[string]()
	})
	var v_2 any = any(hxrt.StdString(func() reflect.Type {
		var hx_callable_2 *hxrt.Function[func() reflect.Type] = factory
		return hxrt.FunctionCall(hx_callable_2)()
	}().String()))
	hxrt.Println(v_2)
}
