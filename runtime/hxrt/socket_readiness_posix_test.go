//go:build linux || darwin

package hxrt

import (
	"errors"
	"net"
	"syscall"
	"testing"
	"time"
)

type socketTestControlFailureRawConn struct{}

func (socketTestControlFailureRawConn) Control(_ func(uintptr)) error {
	return errors.New("control failed before callback")
}

func (socketTestControlFailureRawConn) Read(_ func(uintptr) bool) error {
	return errors.New("unexpected read")
}

func (socketTestControlFailureRawConn) Write(_ func(uintptr) bool) error {
	return errors.New("unexpected write")
}

type socketTestControlFailureResource struct{}

func (socketTestControlFailureResource) SyscallConn() (syscall.RawConn, error) {
	return socketTestControlFailureRawConn{}, nil
}

type socketTestCallbackRawConn struct {
	descriptor uintptr
	controlErr error
}

func (connection socketTestCallbackRawConn) Control(callback func(uintptr)) error {
	callback(connection.descriptor)
	return connection.controlErr
}

func (socketTestCallbackRawConn) Read(_ func(uintptr) bool) error {
	return errors.New("unexpected read")
}

func (socketTestCallbackRawConn) Write(_ func(uintptr) bool) error {
	return errors.New("unexpected write")
}

type socketTestCallbackResource struct {
	raw syscall.RawConn
}

func (resource socketTestCallbackResource) SyscallConn() (syscall.RawConn, error) {
	return resource.raw, nil
}

func socketTestSaturateTCPWriteBufferNative(t *testing.T, connection *net.TCPConn) {
	t.Helper()
	raw, err := connection.SyscallConn()
	if err != nil {
		t.Fatal(err)
	}
	payload := make([]byte, 64*1024)
	consecutiveBlocked := 0
	for written := 0; written < 64<<20; {
		count := 0
		var writeErr error
		if err := raw.Control(func(descriptor uintptr) {
			count, writeErr = syscall.Write(int(descriptor), payload)
		}); err != nil {
			t.Fatal(err)
		}
		written += count
		if writeErr == nil {
			consecutiveBlocked = 0
			continue
		}
		if errors.Is(writeErr, syscall.EAGAIN) || errors.Is(writeErr, syscall.EWOULDBLOCK) {
			consecutiveBlocked++
			if consecutiveBlocked == 3 {
				return
			}
			// The kernel may acknowledge bytes into the peer receive window after
			// the first EAGAIN, briefly reopening sender space without an
			// application read. Require three quiescent blocked observations so
			// the following readiness assertion starts from a stable full window.
			time.Sleep(5 * time.Millisecond)
			continue
		}
		t.Fatalf("filling TCP send buffer failed after %d bytes: %v", written, writeErr)
	}
	t.Fatal("TCP send buffer did not reach EAGAIN before the bounded write limit")
}

func TestSocketReadinessSnapshotOwnsStableDescriptor(t *testing.T) {
	listener, client, accepted := socketTestTCPPair(t)
	defer SocketClose(listener)
	defer SocketClose(accepted)

	snapshot, err := client.readinessSnapshot()
	if err != nil {
		t.Fatal(err)
	}
	if !snapshot.hasDescriptor {
		t.Fatal("readiness snapshot has no descriptor")
	}
	flags, _, errno := syscall.Syscall(
		syscall.SYS_FCNTL,
		snapshot.descriptor,
		uintptr(syscall.F_GETFD),
		0,
	)
	if errno != 0 {
		t.Fatalf("reading snapshot descriptor flags failed: %v", errno)
	}
	if flags&syscall.FD_CLOEXEC == 0 {
		t.Fatal("snapshot descriptor is not close-on-exec")
	}
	SocketClose(client)

	var stat syscall.Stat_t
	if err := syscall.Fstat(int(snapshot.descriptor), &stat); err != nil {
		t.Fatalf("snapshot descriptor did not survive source close: %v", err)
	}
	snapshot.release()
	if err := syscall.Fstat(int(snapshot.descriptor), &stat); !errors.Is(err, syscall.EBADF) {
		t.Fatalf("released snapshot descriptor error = %v, want EBADF", err)
	}
}

func TestSocketReadinessControlFailureDoesNotCloseDescriptorZero(t *testing.T) {
	savedStdin, savedStdinErr := syscall.Dup(0)
	pipe := []int{-1, -1}
	if err := syscall.Pipe(pipe); err != nil {
		t.Fatal(err)
	}
	if pipe[0] != 0 {
		if err := syscall.Dup2(pipe[0], 0); err != nil {
			_ = syscall.Close(pipe[0])
			_ = syscall.Close(pipe[1])
			t.Fatal(err)
		}
		_ = syscall.Close(pipe[0])
	}
	defer func() {
		_ = syscall.Close(pipe[1])
		if savedStdinErr == nil {
			_ = syscall.Dup2(savedStdin, 0)
			_ = syscall.Close(savedStdin)
		} else {
			_ = syscall.Close(0)
		}
	}()

	if _, err := socketResourceDescriptor(socketTestControlFailureResource{}); err == nil {
		t.Fatal("readiness descriptor acquisition unexpectedly succeeded")
	}
	var stat syscall.Stat_t
	if err := syscall.Fstat(0, &stat); err != nil {
		t.Fatalf("Control failure closed unrelated descriptor 0: %v", err)
	}
}

func TestSocketReadinessControlFailureClosesOwnedDuplicateZero(t *testing.T) {
	savedStdin, savedStdinErr := syscall.Dup(0)
	pipe := []int{-1, -1}
	if err := syscall.Pipe(pipe); err != nil {
		t.Fatal(err)
	}
	if err := syscall.Close(0); err != nil && !errors.Is(err, syscall.EBADF) {
		t.Fatal(err)
	}
	defer func() {
		_ = syscall.Close(pipe[0])
		_ = syscall.Close(pipe[1])
		if savedStdinErr == nil {
			_ = syscall.Dup2(savedStdin, 0)
			_ = syscall.Close(savedStdin)
		}
	}()

	resource := socketTestCallbackResource{raw: socketTestCallbackRawConn{
		descriptor: uintptr(pipe[0]),
		controlErr: errors.New("control failed after callback"),
	}}
	if _, err := socketResourceDescriptor(resource); err == nil {
		t.Fatal("post-callback Control failure unexpectedly succeeded")
	}
	var stat syscall.Stat_t
	if err := syscall.Fstat(0, &stat); !errors.Is(err, syscall.EBADF) {
		t.Fatalf("owned duplicate descriptor 0 remained open: %v", err)
	}
	if err := syscall.Fstat(pipe[0], &stat); err != nil {
		t.Fatalf("source descriptor was damaged: %v", err)
	}
}

func TestSocketReadinessDuplicationFailureClosesNothingUnrelated(t *testing.T) {
	savedStdin, savedStdinErr := syscall.Dup(0)
	pipe := []int{-1, -1}
	if err := syscall.Pipe(pipe); err != nil {
		t.Fatal(err)
	}
	if pipe[0] != 0 {
		if err := syscall.Dup2(pipe[0], 0); err != nil {
			t.Fatal(err)
		}
		_ = syscall.Close(pipe[0])
	}
	defer func() {
		_ = syscall.Close(pipe[1])
		if savedStdinErr == nil {
			_ = syscall.Dup2(savedStdin, 0)
			_ = syscall.Close(savedStdin)
		} else {
			_ = syscall.Close(0)
		}
	}()

	resource := socketTestCallbackResource{raw: socketTestCallbackRawConn{
		descriptor: ^uintptr(0),
	}}
	if _, err := socketResourceDescriptor(resource); err == nil {
		t.Fatal("invalid source descriptor unexpectedly duplicated")
	}
	var stat syscall.Stat_t
	if err := syscall.Fstat(0, &stat); err != nil {
		t.Fatalf("duplication failure damaged descriptor 0: %v", err)
	}
}

func TestSocketSelectReportsNativeOutOfBandException(t *testing.T) {
	listener, client, accepted := socketTestTCPPair(t)
	defer SocketClose(listener)
	defer SocketClose(client)
	defer SocketClose(accepted)

	sender, ok := accepted.snapshotConn().(*net.TCPConn)
	if !ok {
		t.Fatalf("accepted connection = %T, want *net.TCPConn", accepted.snapshotConn())
	}
	raw, err := sender.SyscallConn()
	if err != nil {
		t.Fatal(err)
	}
	var sendErr error
	if err := raw.Control(func(descriptor uintptr) {
		_, sendErr = syscall.SendmsgN(int(descriptor), []byte{'!'}, nil, nil, syscall.MSG_OOB)
	}); err != nil {
		t.Fatal(err)
	}
	if sendErr != nil {
		t.Fatal(sendErr)
	}

	selected := SocketSelect(nil, nil, []*SocketHandle{client}, 1, true)
	if selected == nil || !socketTestIntsEqual(selected.Others, 0) {
		t.Fatalf("out-of-band SocketSelect others = %#v, want index 0", selected)
	}
}
