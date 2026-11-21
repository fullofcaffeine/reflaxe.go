//go:build linux

package hxrt

import "syscall"

const (
	terminalReadTermiosRequest  uintptr = syscall.TCGETS
	terminalWriteTermiosRequest uintptr = syscall.TCSETS
)
