package main

import "examples_tui_todo_portable/hxrt"

type I_haxe__io__Eof interface {
	toString() *string
}

type haxe__io__Eof struct {
	__hx_this I_haxe__io__Eof
}

func New_haxe__io__Eof() *haxe__io__Eof {
	self := &haxe__io__Eof{}
	self.__hx_this = self
	return self
}

func (self *haxe__io__Eof) toString() *string {
	return hxrt.StringFromLiteral("Eof")
}

func (self *haxe__io__Eof) String() string {
	return *self.__hx_this.toString()
}
