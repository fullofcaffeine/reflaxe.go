package main

import "snapshot/hxrt"

type I_Greeter interface {
	greet(name *string, punct *string) *string
	wrap(prefix *string, suffix *string) *string
}

type Greeter struct {
	__hx_this I_Greeter
}

func New_Greeter() *Greeter {
	self := &Greeter{}
	self.__hx_this = self
	return self
}

func (self *Greeter) greet(name *string, punct *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("hello "), name), punct)
}

func (self *Greeter) wrap(prefix *string, suffix *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(prefix, hxrt.StringFromLiteral("go")), suffix)
}

func main() {
	greeter := New_Greeter()
	var v any = any(greeter.greet(hxrt.StringFromLiteral("world"), hxrt.StringFromLiteral("!")))
	hxrt.Println(v)
	var v_1 any = any(greeter.greet(hxrt.StringFromLiteral("Go"), hxrt.StringFromLiteral("!")))
	hxrt.Println(v_1)
	var v_2 any = any(greeter.greet(hxrt.StringFromLiteral("Go"), hxrt.StringFromLiteral("?")))
	hxrt.Println(v_2)
	var v_3 any = any(greeter.wrap(hxrt.StringFromLiteral("["), hxrt.StringFromLiteral("]")))
	hxrt.Println(v_3)
	var v_4 any = any(greeter.wrap(hxrt.StringFromLiteral("<"), hxrt.StringFromLiteral("]")))
	hxrt.Println(v_4)
	var v_5 any = any(greeter.wrap(hxrt.StringFromLiteral("<"), hxrt.StringFromLiteral(">")))
	hxrt.Println(v_5)
}
