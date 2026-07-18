//go:build !aix && !android && !darwin && !dragonfly && !freebsd && !hurd && !illumos && !ios && !linux && !netbsd && !openbsd && !solaris && !windows

package hxrt

import "errors"

// socketSetBroadcast keeps unsupported Go platforms buildable and explicit.
func socketSetBroadcast(_ uintptr, _ int) error {
	return errors.New("UDP broadcast socket options are unavailable on this platform")
}
