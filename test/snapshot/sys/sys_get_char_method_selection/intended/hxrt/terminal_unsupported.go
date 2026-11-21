//go:build !linux && !darwin && !windows

package hxrt

import (
	"fmt"
	"os"
	"runtime"
)

func enterTerminalCharacterMode(file *os.File) (func() error, error) {
	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("inspect standard input for Sys.getChar: %w", err)
	}
	if info.Mode()&os.ModeCharDevice == 0 {
		return nil, nil
	}
	return nil, fmt.Errorf("Sys.getChar terminal control is unsupported on %s", runtime.GOOS)
}
