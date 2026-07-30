//go:build darwin || ios

package hxrt

import (
	"fmt"
	"syscall"
	"time"
)

const socketFDSetWordBits = 32

func socketDuplicateDescriptor(descriptor uintptr) (uintptr, error) {
	duplicate, _, errno := syscall.Syscall(
		syscall.SYS_FCNTL,
		descriptor,
		uintptr(syscall.F_DUPFD_CLOEXEC),
		0,
	)
	if errno != 0 {
		return 0, errno
	}
	return duplicate, nil
}

func socketCloseDescriptor(descriptor uintptr) error {
	return syscall.Close(int(descriptor))
}

func socketFDSetAdd(set *syscall.FdSet, descriptor uintptr) error {
	if descriptor >= uintptr(len(set.Bits)*socketFDSetWordBits) {
		return fmt.Errorf("socket descriptor %d exceeds select capacity", descriptor)
	}
	index := descriptor / socketFDSetWordBits
	bit := descriptor % socketFDSetWordBits
	set.Bits[index] |= int32(1) << bit
	return nil
}

func socketFDSetContains(set *syscall.FdSet, descriptor uintptr) bool {
	if descriptor >= uintptr(len(set.Bits)*socketFDSetWordBits) {
		return false
	}
	index := descriptor / socketFDSetWordBits
	bit := descriptor % socketFDSetWordBits
	return set.Bits[index]&(int32(1)<<bit) != 0
}

func socketSelectNative(request socketNativeSelectRequest) (*socketNativeSelectResult, error) {
	deadline := time.Now().Add(request.Timeout)
	for {
		var readSet syscall.FdSet
		var writeSet syscall.FdSet
		var otherSet syscall.FdSet
		maxDescriptor := -1
		add := func(set *syscall.FdSet, descriptors []uintptr) error {
			for _, descriptor := range descriptors {
				if err := socketFDSetAdd(set, descriptor); err != nil {
					return err
				}
				if int(descriptor) > maxDescriptor {
					maxDescriptor = int(descriptor)
				}
			}
			return nil
		}
		if err := add(&readSet, request.Read); err != nil {
			return nil, err
		}
		if err := add(&writeSet, request.Write); err != nil {
			return nil, err
		}
		if err := add(&otherSet, request.Others); err != nil {
			return nil, err
		}

		remaining := time.Until(deadline)
		if request.Timeout <= 0 || remaining < 0 {
			remaining = 0
		}
		timeval := syscall.NsecToTimeval(remaining.Nanoseconds())
		err := syscall.Select(maxDescriptor+1, &readSet, &writeSet, &otherSet, &timeval)
		if err != nil {
			if err == syscall.EINTR {
				continue
			}
			return nil, err
		}

		result := newSocketNativeSelectResult()
		for _, descriptor := range request.Read {
			if socketFDSetContains(&readSet, descriptor) {
				result.Read[descriptor] = struct{}{}
			}
		}
		for _, descriptor := range request.Write {
			if socketFDSetContains(&writeSet, descriptor) {
				result.Write[descriptor] = struct{}{}
			}
		}
		for _, descriptor := range request.Others {
			if socketFDSetContains(&otherSet, descriptor) {
				result.Others[descriptor] = struct{}{}
			}
		}
		return result, nil
	}
}
