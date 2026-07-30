//go:build !aix && !android && !darwin && !dragonfly && !freebsd && !hurd && !illumos && !ios && !linux && !netbsd && !openbsd && !solaris && !windows

package hxrt

import "errors"

func socketBindTCPNative(_ string, _ int) (socketBoundTCP, error) {
	return nil, errors.New("separate TCP bind/listen is unavailable on this platform")
}

func socketRelistenTCP(_ socketDeadlineListener, _ int) error {
	return errors.New("TCP backlog control is unavailable on this platform")
}
