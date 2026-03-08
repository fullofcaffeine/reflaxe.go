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
	hxrt.Println(greeter.greet(hxrt.StringFromLiteral("world"), hxrt.StringFromLiteral("!")))
	hxrt.Println(greeter.greet(hxrt.StringFromLiteral("Go"), hxrt.StringFromLiteral("!")))
	hxrt.Println(greeter.greet(hxrt.StringFromLiteral("Go"), hxrt.StringFromLiteral("?")))
	hxrt.Println(greeter.wrap(hxrt.StringFromLiteral("["), hxrt.StringFromLiteral("]")))
	hxrt.Println(greeter.wrap(hxrt.StringFromLiteral("<"), hxrt.StringFromLiteral("]")))
	hxrt.Println(greeter.wrap(hxrt.StringFromLiteral("<"), hxrt.StringFromLiteral(">")))
}
