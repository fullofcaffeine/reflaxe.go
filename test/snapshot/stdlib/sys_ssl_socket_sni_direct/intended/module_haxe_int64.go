package main

type I_haxe___Int64_____Int64 interface {
}

type haxe___Int64_____Int64 struct {
	__hx_this I_haxe___Int64_____Int64
	high      int
	low       int
}

func New_haxe___Int64_____Int64(high int, low int) *haxe___Int64_____Int64 {
	self := &haxe___Int64_____Int64{}
	self.__hx_this = self
	self.high = high
	self.low = low
	return self
}
