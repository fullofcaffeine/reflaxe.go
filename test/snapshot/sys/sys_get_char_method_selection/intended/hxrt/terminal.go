package hxrt

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

var sysReadCharMu sync.Mutex

// readStandardInputChar reads one byte while temporarily placing an interactive
// terminal in non-canonical, no-echo mode.
//
// What: Provides the native character-at-a-time capability needed by staged
// Sys.getChar while retaining redirected byte-stream behavior.
// Why: os.Stdin.Read alone remains line-buffered and host-echoed on a terminal;
// terminal state is native process state and cannot be represented in Haxe.
// How: Serialize terminal-state changes, delegate platform control to the
// build-tagged helper, read exactly one byte, and restore state on every return.
func readStandardInputChar() (value int, eof bool, err error) {
	sysReadCharMu.Lock()
	defer sysReadCharMu.Unlock()

	if os.Stdin == nil {
		return 0, false, fmt.Errorf("read standard input character: %w", os.ErrInvalid)
	}

	restore, err := enterTerminalCharacterMode(os.Stdin)
	if err != nil {
		return 0, false, err
	}
	if restore != nil {
		defer func() {
			if restoreErr := restore(); restoreErr != nil && err == nil {
				err = restoreErr
			}
		}()
	}

	var one [1]byte
	count, readErr := os.Stdin.Read(one[:])
	if count > 0 {
		return int(one[0]), false, nil
	}
	if errors.Is(readErr, io.EOF) || (count == 0 && readErr == nil) {
		return 0, true, nil
	}
	if readErr != nil {
		return 0, false, fmt.Errorf("read standard input character: %w", readErr)
	}
	return 0, true, nil
}

// SysReadCharValue is the Haxe-shaped typed character-input capability.
//
// What: Returns one byte value, or -1 when redirected input reaches EOF.
// Why: Haxe source must own haxe.io.Eof construction and explicit echo policy;
// only terminal state and native read failures belong in hxrt.
// How: Translate native failures through the normal Haxe exception boundary and
// leave EOF as a typed sentinel for the staged Sys implementation.
func SysReadCharValue() int {
	value, eof, err := readStandardInputChar()
	if err != nil {
		Throw(StringFromLiteral(err.Error()))
		return -1
	}
	if eof {
		return -1
	}
	return value
}
