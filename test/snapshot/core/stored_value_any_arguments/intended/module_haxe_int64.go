package main

type I_haxe___Int64_____Int64 interface {
}

type haxe___Int64_____Int64 struct {
	__hx_this I_haxe___Int64_____Int64
	high      int
	low       int
}

func (self *haxe___Int64_____Int64) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe___Int64_____Int64(high int, low int) *haxe___Int64_____Int64 {
	return hx_init_haxe___Int64_____Int64(hx_alloc_haxe___Int64_____Int64(nil), high, low)
}

func hx_alloc_haxe___Int64_____Int64(owner any) *haxe___Int64_____Int64 {
	self := &haxe___Int64_____Int64{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe___Int64_____Int64)
	}
	return self
}

func hx_init_haxe___Int64_____Int64(self *haxe___Int64_____Int64, high int, low int) *haxe___Int64_____Int64 {
	self.high = high
	self.low = low
	return self
}
