//go:build unix

package main

import (
	"snapshot/hxrt"
	"syscall"
)

type I_UnixOpenResult interface {
}

type UnixOpenResult struct {
	__hx_this I_UnixOpenResult
	fd        int
	err       *hxrt.NativeErrorValue
}

func (self *UnixOpenResult) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_UnixOpenResult(fd int, err *hxrt.NativeErrorValue) *UnixOpenResult {
	return hx_init_UnixOpenResult(hx_alloc_UnixOpenResult(nil), fd, err)
}

func hx_alloc_UnixOpenResult(owner any) *UnixOpenResult {
	self := &UnixOpenResult{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_UnixOpenResult)
	}
	return self
}

func hx_init_UnixOpenResult(self *UnixOpenResult, fd int, err *hxrt.NativeErrorValue) *UnixOpenResult {
	self.fd = fd
	self.err = err
	return self
}

func (self *UnixOpenResult) HxClassName() string {
	return "UnixOpenResult"
}

type I_UnixSyscallOpen interface {
	openReadOnly(path *string, perm uint32) bool
}

type UnixSyscallOpen struct {
	__hx_this I_UnixSyscallOpen
}

func (self *UnixSyscallOpen) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_UnixSyscallOpen() *UnixSyscallOpen {
	return hx_init_UnixSyscallOpen(hx_alloc_UnixSyscallOpen(nil))
}

func hx_alloc_UnixSyscallOpen(owner any) *UnixSyscallOpen {
	self := &UnixSyscallOpen{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_UnixSyscallOpen)
	}
	return self
}

func hx_init_UnixSyscallOpen(self *UnixSyscallOpen) *UnixSyscallOpen {
	return self
}

func (self *UnixSyscallOpen) openReadOnly(path *string, perm uint32) bool {
	opened := func() *UnixOpenResult {
		hx_tuple_1, hx_tuple_2 := syscall.Open(*hxrt.StdString(path), syscall.O_RDONLY, perm)
		return New_UnixOpenResult(hx_tuple_1, hxrt.ImportNativeError(hx_tuple_2))
	}()
	if opened.err != nil {
		return false
	}
	return (hxrt.ImportNativeError(syscall.Close(opened.fd)) == nil)
}

func (self *UnixSyscallOpen) HxClassName() string {
	return "UnixSyscallOpen"
}

var UnixSyscallOpen_installed bool = SyscallOpen_install(New_UnixSyscallOpen())
