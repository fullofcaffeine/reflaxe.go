package main

import "snapshot/hxrt"

func main() {
	hxrt.TryCatch(func() {
		New_haxe__ds__WeakMap()
		hxrt.Println(any(hxrt.StringFromLiteral("constructed")))
	}, func(hx_caught_1 any) {
		switch hx_typed_2 := hx_caught_1.(type) {
		case *haxe__exceptions__NotImplementedException:
			err := hx_typed_2
			var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("not_impl="), hxrt.ExceptionMessage(err)))
			hxrt.Println(v)
			var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed="), hxrt.StdString((err != nil))))
			hxrt.Println(v_1)
		default:
			hxrt.Throw(hx_caught_1)
		}
	})
}

type I__Main__WeakMapKey interface {
}

type _Main__WeakMapKey struct {
	__hx_this I__Main__WeakMapKey
}

func New__Main__WeakMapKey() *_Main__WeakMapKey {
	self := &_Main__WeakMapKey{}
	self.__hx_this = self
	return self
}
