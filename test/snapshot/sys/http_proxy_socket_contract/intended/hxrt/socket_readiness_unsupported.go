//go:build !linux && !darwin && !ios

package hxrt

import "errors"

func socketDuplicateDescriptor(_ uintptr) (uintptr, error) {
	return 0, errors.New("native socket readiness is unavailable on this platform")
}

func socketCloseDescriptor(_ uintptr) error {
	return nil
}

func socketSelectNative(_ socketNativeSelectRequest) (*socketNativeSelectResult, error) {
	return nil, errors.New("native socket readiness is unavailable on this platform")
}
