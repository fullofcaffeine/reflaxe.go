//go:build windows

package main

import (
	"snapshot/hxrt"
	"syscall"
)

type I_WindowsFileTypeResult interface {
}

type WindowsFileTypeResult struct {
	__hx_this I_WindowsFileTypeResult
	n         uint32
	err       *hxrt.NativeErrorValue
}

func (self *WindowsFileTypeResult) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_WindowsFileTypeResult(n uint32, err *hxrt.NativeErrorValue) *WindowsFileTypeResult {
	return hx_init_WindowsFileTypeResult(hx_alloc_WindowsFileTypeResult(nil), n, err)
}

func hx_alloc_WindowsFileTypeResult(owner any) *WindowsFileTypeResult {
	self := &WindowsFileTypeResult{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_WindowsFileTypeResult)
	}
	return self
}

func hx_init_WindowsFileTypeResult(self *WindowsFileTypeResult, n uint32, err *hxrt.NativeErrorValue) *WindowsFileTypeResult {
	self.n = n
	self.err = err
	return self
}

func (self *WindowsFileTypeResult) HxClassName() string {
	return "WindowsFileTypeResult"
}

type I_WindowsOpenResult interface {
}

type WindowsOpenResult struct {
	__hx_this I_WindowsOpenResult
	fd        *syscall.Handle
	err       *hxrt.NativeErrorValue
}

func (self *WindowsOpenResult) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_WindowsOpenResult(fd *syscall.Handle, err *hxrt.NativeErrorValue) *WindowsOpenResult {
	return hx_init_WindowsOpenResult(hx_alloc_WindowsOpenResult(nil), fd, err)
}

func hx_alloc_WindowsOpenResult(owner any) *WindowsOpenResult {
	self := &WindowsOpenResult{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_WindowsOpenResult)
	}
	return self
}

func hx_init_WindowsOpenResult(self *WindowsOpenResult, fd *syscall.Handle, err *hxrt.NativeErrorValue) *WindowsOpenResult {
	self.fd = fd
	self.err = err
	return self
}

func (self *WindowsOpenResult) HxClassName() string {
	return "WindowsOpenResult"
}

type I_WindowsSyscallOpen interface {
	openReadOnly(path *string, perm uint32) bool
}

type WindowsSyscallOpen struct {
	__hx_this I_WindowsSyscallOpen
}

func (self *WindowsSyscallOpen) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_WindowsSyscallOpen() *WindowsSyscallOpen {
	return hx_init_WindowsSyscallOpen(hx_alloc_WindowsSyscallOpen(nil))
}

func hx_alloc_WindowsSyscallOpen(owner any) *WindowsSyscallOpen {
	self := &WindowsSyscallOpen{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_WindowsSyscallOpen)
	}
	return self
}

func hx_init_WindowsSyscallOpen(self *WindowsSyscallOpen) *WindowsSyscallOpen {
	return self
}

func (self *WindowsSyscallOpen) openReadOnly(path *string, perm uint32) bool {
	opened := func() *WindowsOpenResult {
		hx_tuple_1, hx_tuple_2 := syscall.Open(*hxrt.StdString(path), syscall.O_RDONLY, perm)
		return New_WindowsOpenResult(&hx_tuple_1, hxrt.ImportNativeError(hx_tuple_2))
	}()
	if opened.err != nil {
		return false
	}
	fileType := func() *WindowsFileTypeResult {
		hx_tuple_3, hx_tuple_4 := syscall.GetFileType(*opened.fd)
		return New_WindowsFileTypeResult(hx_tuple_3, hxrt.ImportNativeError(hx_tuple_4))
	}()
	closed := hxrt.ImportNativeError(syscall.CloseHandle(*opened.fd))
	return (((fileType.err == nil) && (fileType.n == uint32(syscall.FILE_TYPE_DISK))) && (closed == nil))
}

func (self *WindowsSyscallOpen) HxClassName() string {
	return "WindowsSyscallOpen"
}

var WindowsSyscallOpen_installed bool = SyscallOpen_install(New_WindowsSyscallOpen())
