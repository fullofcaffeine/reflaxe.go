package main

import "snapshot/hxrt"

type I_sys__net__Host interface {
	toString() *string
	reverse() *string
}

type sys__net__Host struct {
	__hx_this I_sys__net__Host
	host      *string
	ip        int
}

func (self *sys__net__Host) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__net__Host(name *string) *sys__net__Host {
	return hx_init_sys__net__Host(hx_alloc_sys__net__Host(nil), name)
}

func hx_alloc_sys__net__Host(owner any) *sys__net__Host {
	self := &sys__net__Host{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__net__Host)
	}
	return self
}

func hx_init_sys__net__Host(self *sys__net__Host, name *string) *sys__net__Host {
	self.host = name
	self.ip = hxrt.HostResolve(name)
	return self
}

func (self *sys__net__Host) toString() *string {
	return hxrt.StdString(hxrt.HostToString(self.ip))
}

func (self *sys__net__Host) reverse() *string {
	return hxrt.StdString(hxrt.HostReverse(self.ip))
}

func (self *sys__net__Host) String() string {
	return *self.__hx_this.toString()
}

func sys__net__Host_fromIPv4(value int) *sys__net__Host {
	result := New_sys__net__Host(hxrt.StdString(hxrt.HostToString(value)))
	result.ip = value
	return result
}

func sys__net__Host_localhost() *string {
	return hxrt.StdString(hxrt.HostLocal())
}
