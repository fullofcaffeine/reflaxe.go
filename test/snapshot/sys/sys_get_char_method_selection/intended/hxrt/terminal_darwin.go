//go:build darwin

package hxrt

import "syscall"

const (
	terminalReadTermiosRequest  uintptr = syscall.TIOCGETA
	terminalWriteTermiosRequest uintptr = syscall.TIOCSETA
)
