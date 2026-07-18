//go:build windows

package hxrt

import "syscall"

// socketSetBroadcast adapts RawConn's uintptr to the Windows Handle type.
func socketSetBroadcast(fileDescriptor uintptr, value int) error {
	return syscall.SetsockoptInt(syscall.Handle(fileDescriptor), syscall.SOL_SOCKET, syscall.SO_BROADCAST, value)
}
