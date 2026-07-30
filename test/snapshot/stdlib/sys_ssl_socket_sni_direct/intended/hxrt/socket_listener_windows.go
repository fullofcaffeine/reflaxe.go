//go:build windows

package hxrt

import (
	"errors"
	"net"
	"os"
	"syscall"
)

// socketBoundTCPWindows retains a Winsock handle between bind and listen.
//
// Windows remains compile-only release evidence. This typed adapter mirrors
// the public lifecycle without turning a Winsock handle into generated data.
type socketBoundTCPWindows struct {
	handle  syscall.Handle
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

	handle, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		return nil, err
	}
	closeOnError := true
	defer func() {
		if closeOnError {
			_ = syscall.Closesocket(handle)
		}
	}()
	nativeAddress := &syscall.SockaddrInet4{Port: port}
	copy(nativeAddress.Addr[:], ip)
	if err := syscall.Bind(handle, nativeAddress); err != nil {
		return nil, err
	}
	actual, err := syscall.Getsockname(handle)
	if err != nil {
		return nil, err
	}
	actualIPv4, ok := actual.(*syscall.SockaddrInet4)
	if !ok {
		return nil, errors.New("bound TCP socket did not return an IPv4 address")
	}

	closeOnError = false
	return &socketBoundTCPWindows{
		handle: handle,
		address: &net.TCPAddr{
			IP:   net.IPv4(actualIPv4.Addr[0], actualIPv4.Addr[1], actualIPv4.Addr[2], actualIPv4.Addr[3]),
			Port: actualIPv4.Port,
		},
	}, nil
}

func (bound *socketBoundTCPWindows) Addr() net.Addr {
	if bound == nil {
		return nil
	}
	return bound.address
}

func (bound *socketBoundTCPWindows) Listen(backlog int) (*net.TCPListener, error) {
	if bound == nil || bound.handle == syscall.InvalidHandle {
		return nil, net.ErrClosed
	}
	if err := syscall.Listen(bound.handle, backlog); err != nil {
		return nil, err
	}

	// net.FileListener duplicates the Winsock handle into Go's poller. os.File
	// does not own raw sockets correctly on Windows, so mark that temporary view
	// closed before releasing the original with closesocket.
	file := os.NewFile(uintptr(bound.handle), "haxe-go-bound-tcp")
	if file == nil {
		_ = syscall.Closesocket(bound.handle)
		bound.handle = syscall.InvalidHandle
		return nil, errors.New("could not retain bound TCP handle")
	}
	listener, listenerErr := net.FileListener(file)
	_ = file.Close()
	closeErr := syscall.Closesocket(bound.handle)
	bound.handle = syscall.InvalidHandle
	if listenerErr != nil {
		return nil, listenerErr
	}
	if closeErr != nil {
		_ = listener.Close()
		return nil, closeErr
	}
	tcpListener, ok := listener.(*net.TCPListener)
	if !ok {
		_ = listener.Close()
		return nil, errors.New("bound TCP handle did not create a TCP listener")
	}
	return tcpListener, nil
}

func (bound *socketBoundTCPWindows) Close() error {
	if bound == nil || bound.handle == syscall.InvalidHandle {
		return nil
	}
	err := syscall.Closesocket(bound.handle)
	bound.handle = syscall.InvalidHandle
	return err
}

func socketRelistenTCP(listener socketDeadlineListener, backlog int) error {
	rawListener, ok := listener.(interface {
		SyscallConn() (syscall.RawConn, error)
	})
	if !ok {
		return errors.New("socket listener does not expose Winsock backlog control")
	}
	raw, err := rawListener.SyscallConn()
	if err != nil {
		return err
	}
	var listenErr error
	if err := raw.Control(func(handle uintptr) {
		listenErr = syscall.Listen(syscall.Handle(handle), backlog)
	}); err != nil {
		return err
	}
	return listenErr
}
