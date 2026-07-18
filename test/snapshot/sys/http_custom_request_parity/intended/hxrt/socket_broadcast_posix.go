//go:build aix || android || darwin || dragonfly || freebsd || hurd || illumos || ios || linux || netbsd || openbsd || solaris

package hxrt

import "syscall"

// socketSetBroadcast adapts RawConn's uintptr to the POSIX descriptor type.
func socketSetBroadcast(fileDescriptor uintptr, value int) error {
	return syscall.SetsockoptInt(int(fileDescriptor), syscall.SOL_SOCKET, syscall.SO_BROADCAST, value)
}
