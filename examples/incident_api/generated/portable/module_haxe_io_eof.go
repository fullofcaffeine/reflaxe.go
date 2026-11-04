package main

import "examples_incident_api_portable/hxrt"

type I_haxe__io__Eof interface {
	toString() *string
}

type haxe__io__Eof struct {
	__hx_this I_haxe__io__Eof
}

func (self *haxe__io__Eof) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__io__Eof() *haxe__io__Eof {
	return hx_init_haxe__io__Eof(hx_alloc_haxe__io__Eof(nil))
}

func hx_alloc_haxe__io__Eof(owner any) *haxe__io__Eof {
	self := &haxe__io__Eof{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__io__Eof)
	}
	return self
}

func hx_init_haxe__io__Eof(self *haxe__io__Eof) *haxe__io__Eof {
	return self
}

func (self *haxe__io__Eof) toString() *string {
	return hxrt.StringFromLiteral("Eof")
}

func (self *haxe__io__Eof) String() string {
	return *self.__hx_this.toString()
}
