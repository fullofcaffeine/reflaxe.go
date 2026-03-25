package main

import "snapshot/hxrt"

type I_sys__net__Address interface {
	getHost() *sys__net__Host
	compare(a *sys__net__Address) int
	clone() *sys__net__Address
}

type sys__net__Address struct {
	__hx_this I_sys__net__Address
	host      int
	port      int
}

func New_sys__net__Address() *sys__net__Address {
	self := &sys__net__Address{}
	self.__hx_this = self
	self.host = 0
	self.port = 0
	return self
}

func (self *sys__net__Address) getHost() *sys__net__Host {
	h := New_sys__net__Host(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		value := self.host
		return hxrt.StdString(int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value)) >> uint(24)))))) & hxrt.Int32Wrap(255)))))
	}(), hxrt.StringFromLiteral(".")), func() *string {
		value_1 := self.host
		return hxrt.StdString(int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value_1)) >> uint(16)))))) & hxrt.Int32Wrap(255)))))
	}()), hxrt.StringFromLiteral(".")), func() *string {
		value_2 := self.host
		return hxrt.StdString(int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value_2)) >> uint(8)))))) & hxrt.Int32Wrap(255)))))
	}()), hxrt.StringFromLiteral(".")), func() *string {
		value_3 := self.host
		return hxrt.StdString(int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value_3)) >> uint(0)))))) & hxrt.Int32Wrap(255)))))
	}()))
	h.ip = self.host
	return h
}

func (self *sys__net__Address) compare(a *sys__net__Address) int {
	dh := int(int32((hxrt.Int32Wrap(a.host) - hxrt.Int32Wrap(self.host))))
	if dh != 0 {
		return dh
	}
	dp := int(int32((hxrt.Int32Wrap(a.port) - hxrt.Int32Wrap(self.port))))
	if dp != 0 {
		return dp
	}
	return 0
}

func (self *sys__net__Address) clone() *sys__net__Address {
	c := New_sys__net__Address()
	c.host = self.host
	c.port = self.port
	return c
}

func sys__net__Address_octet(value int, shift int) *string {
	return hxrt.StdString(int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value)) >> uint(shift)))))) & hxrt.Int32Wrap(255)))))
}
