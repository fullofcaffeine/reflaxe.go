//go:build windows

package hxrt

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

const (
	windowsEnableLineInput = uint32(0x0002)
	windowsEnableEchoInput = uint32(0x0004)
	windowsInvalidHandle   = syscall.Errno(6)
)

var windowsSetConsoleMode = syscall.NewLazyDLL("kernel32.dll").NewProc("SetConsoleMode")

func setWindowsConsoleMode(handle syscall.Handle, mode uint32) error {
	result, _, callErr := windowsSetConsoleMode.Call(uintptr(handle), uintptr(mode))
	if result != 0 {
		return nil
	}
	if callErr != nil && callErr != syscall.Errno(0) {
		return callErr
	}
	return syscall.EINVAL
}

func enterTerminalCharacterMode(file *os.File) (func() error, error) {
	handle := syscall.Handle(file.Fd())
	var original uint32
	if err := syscall.GetConsoleMode(handle, &original); err != nil {
		if errors.Is(err, windowsInvalidHandle) {
			return nil, nil
		}
		return nil, fmt.Errorf("inspect standard input console mode: %w", err)
	}

	characterMode := original &^ (windowsEnableLineInput | windowsEnableEchoInput)
	if err := setWindowsConsoleMode(handle, characterMode); err != nil {
		return nil, fmt.Errorf("enter standard input character mode: %w", err)
	}
	return func() error {
		if err := setWindowsConsoleMode(handle, original); err != nil {
			return fmt.Errorf("restore standard input console mode: %w", err)
		}
		return nil
	}, nil
}
