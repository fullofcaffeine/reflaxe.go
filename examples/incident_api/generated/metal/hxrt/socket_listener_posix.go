//go:build aix || android || darwin || dragonfly || freebsd || hurd || illumos || ios || linux || netbsd || openbsd || solaris

package hxrt

import (
	"errors"
	"net"
	"os"
	"syscall"
)

// socketBoundTCPPOSIX owns a bound descriptor before the listen transition.
//
// What: Retains one reserved IPv4 endpoint that is not yet accepting clients.
// Why: Go's net.Listen combines bind and listen and chooses its own backlog,
// while Haxe exposes those as separate, ordered operations.
// How: Bind through the POSIX socket API, then convert the descriptor to a
// pollable net.TCPListener only after Socket.listen supplies the backlog.
type socketBoundTCPPOSIX struct {
	file    *os.File
	address *net.TCPAddr
}

func socketBindTCPNative(host string, port int) (socketBoundTCP, error) {
	if port < 0 || port > 65535 {
		return nil, errors.New("socket bind port is outside 0...65535")
	}
	ip := net.ParseIP(host).To4()
	if ip == nil {
		return nil, errors.New("socket bind requires a numeric IPv4 address")
	}

	fileDescriptor, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		return nil, err
	}
	syscall.CloseOnExec(fileDescriptor)
	closeOnError := true
	defer func() {
		if closeOnError {
			_ = syscall.Close(fileDescriptor)
		}
	}()
	if err := syscall.SetsockoptInt(fileDescriptor, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		return nil, err
	}
	nativeAddress := &syscall.SockaddrInet4{Port: port}
	copy(nativeAddress.Addr[:], ip)
	if err := syscall.Bind(fileDescriptor, nativeAddress); err != nil {
		return nil, err
	}
	actual, err := syscall.Getsockname(fileDescriptor)
	if err != nil {
		return nil, err
	}
	actualIPv4, ok := actual.(*syscall.SockaddrInet4)
	if !ok {
		return nil, errors.New("bound TCP socket did not return an IPv4 address")
	}

	file := os.NewFile(uintptr(fileDescriptor), "haxe-go-bound-tcp")
	if file == nil {
		return nil, errors.New("could not retain bound TCP descriptor")
	}
	closeOnError = false
	return &socketBoundTCPPOSIX{
		file: file,
		address: &net.TCPAddr{
			IP:   net.IPv4(actualIPv4.Addr[0], actualIPv4.Addr[1], actualIPv4.Addr[2], actualIPv4.Addr[3]),
			Port: actualIPv4.Port,
		},
	}, nil
}

func (bound *socketBoundTCPPOSIX) Addr() net.Addr {
	if bound == nil {
		return nil
	}
	return bound.address
}

func (bound *socketBoundTCPPOSIX) Listen(backlog int) (*net.TCPListener, error) {
	if bound == nil || bound.file == nil {
		return nil, net.ErrClosed
	}
	if err := syscall.Listen(int(bound.file.Fd()), backlog); err != nil {
		return nil, err
	}
	listener, err := net.FileListener(bound.file)
	if err != nil {
		return nil, err
	}
	tcpListener, ok := listener.(*net.TCPListener)
	if !ok {
		_ = listener.Close()
		return nil, errors.New("bound TCP descriptor did not create a TCP listener")
	}
	if err := bound.file.Close(); err != nil {
		_ = tcpListener.Close()
		bound.file = nil
		return nil, err
	}
	bound.file = nil
	return tcpListener, nil
}

func (bound *socketBoundTCPPOSIX) Close() error {
	if bound == nil || bound.file == nil {
		return nil
	}
	err := bound.file.Close()
	bound.file = nil
	return err
}

func socketRelistenTCP(listener socketDeadlineListener, backlog int) error {
	rawListener, ok := listener.(interface {
		SyscallConn() (syscall.RawConn, error)
	})
	if !ok {
		return errors.New("socket listener does not expose POSIX backlog control")
	}
	raw, err := rawListener.SyscallConn()
	if err != nil {
		return err
	}
	var listenErr error
	if err := raw.Control(func(fileDescriptor uintptr) {
		listenErr = syscall.Listen(int(fileDescriptor), backlog)
	}); err != nil {
		return err
	}
	return listenErr
}
