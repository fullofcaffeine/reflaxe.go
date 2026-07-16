//go:build linux || darwin

package hxrt

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

// terminalIoctlTermios is the sole unsafe island in terminal support.
//
// What: Passes a typed syscall.Termios address to the host ioctl operation.
// Why: The frozen Go syscall package exposes Termios and ioctl constants but no
// safe terminal-state wrapper, while adding x/term would either raise the
// generated Go language floor or pin a dependency with a known advisory.
// How: Keep the pointer live for the syscall duration and expose only typed
// state transitions to the rest of hxrt. Race/checkptr gates exercise this path.
func terminalIoctlTermios(fd uintptr, request uintptr, state *syscall.Termios) error {
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		fd,
		request,
		uintptr(unsafe.Pointer(state)),
	)
	runtime.KeepAlive(state)
	if errno != 0 {
		return errno
	}
	return nil
}

func enterTerminalCharacterMode(file *os.File) (func() error, error) {
	state := new(syscall.Termios)
	if err := terminalIoctlTermios(file.Fd(), terminalReadTermiosRequest, state); err != nil {
		if errors.Is(err, syscall.ENOTTY) {
			return nil, nil
		}
		return nil, fmt.Errorf("inspect standard input terminal state: %w", err)
	}

	original := *state
	state.Lflag &^= syscall.ICANON | syscall.ECHO | syscall.ECHONL
	state.Cc[syscall.VMIN] = 1
	state.Cc[syscall.VTIME] = 0
	if err := terminalIoctlTermios(file.Fd(), terminalWriteTermiosRequest, state); err != nil {
		return nil, fmt.Errorf("enter standard input character mode: %w", err)
	}

	return func() error {
		restored := original
		if err := terminalIoctlTermios(file.Fd(), terminalWriteTermiosRequest, &restored); err != nil {
			return fmt.Errorf("restore standard input terminal state: %w", err)
		}
		return nil
	}, nil
}
