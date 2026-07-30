//go:build linux || darwin

package hxrt

import (
	"crypto/tls"
	"net"
	"sync"
	"syscall"
	"testing"
)

func socketTestTLSPair(t *testing.T) (*SocketHandle, *SocketHandle) {
	t.Helper()
	serverConfig, _ := socketTestCertificate(t, socketTestHost)
	listener, err := net.Listen("tcp4", net.JoinHostPort(socketTestHost, "0"))
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	serverRaw := make(chan net.Conn, 1)
	serverAcceptErr := make(chan error, 1)
	go func() {
		connection, acceptErr := listener.Accept()
		if acceptErr != nil {
			serverAcceptErr <- acceptErr
			return
		}
		serverRaw <- connection
	}()

	clientRaw, err := net.Dial("tcp4", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	var acceptedRaw net.Conn
	select {
	case acceptedRaw = <-serverRaw:
	case acceptErr := <-serverAcceptErr:
		_ = clientRaw.Close()
		t.Fatal(acceptErr)
	}

	clientTLS := tls.Client(clientRaw, &tls.Config{InsecureSkipVerify: true})
	serverTLS := tls.Server(acceptedRaw, serverConfig)
	serverHandshake := make(chan error, 1)
	go func() {
		serverHandshake <- serverTLS.Handshake()
	}()
	if err := clientTLS.Handshake(); err != nil {
		_ = clientTLS.Close()
		_ = serverTLS.Close()
		t.Fatal(err)
	}
	if err := <-serverHandshake; err != nil {
		_ = clientTLS.Close()
		_ = serverTLS.Close()
		t.Fatal(err)
	}

	client := SocketNewTCP()
	server := SocketNewTCP()
	client.installConn(clientTLS)
	server.installConn(serverTLS)
	t.Cleanup(func() {
		SocketClose(client)
		SocketClose(server)
	})
	return client, server
}

func socketTestReadExact(t *testing.T, handle *SocketHandle, length int) []int {
	t.Helper()
	values := make([]int, 0, length)
	for len(values) < length {
		result := SocketReadValues(handle, length-len(values))
		if result.Status != SocketIOReady || result.Count <= 0 {
			t.Fatalf("SocketReadValues = %#v after %d bytes, want progress", result, len(values))
		}
		values = append(values, result.Values...)
	}
	return values
}

func socketTestTCPConnection(t *testing.T, handle *SocketHandle) *net.TCPConn {
	t.Helper()
	connection := handle.snapshotConn()
	for connection != nil {
		if tcpConnection, ok := connection.(*net.TCPConn); ok {
			return tcpConnection
		}
		nested, ok := connection.(interface{ NetConn() net.Conn })
		if !ok {
			t.Fatalf("connection %T does not expose an underlying TCP socket", connection)
		}
		connection = nested.NetConn()
	}
	t.Fatal("socket has no connection")
	return nil
}

func socketTestNoDelay(t *testing.T, handle *SocketHandle) int {
	t.Helper()
	raw, err := socketTestTCPConnection(t, handle).SyscallConn()
	if err != nil {
		t.Fatal(err)
	}
	value := -1
	var optionErr error
	if err := raw.Control(func(descriptor uintptr) {
		value, optionErr = syscall.GetsockoptInt(
			int(descriptor),
			syscall.IPPROTO_TCP,
			syscall.TCP_NODELAY,
		)
	}); err != nil {
		t.Fatal(err)
	}
	if optionErr != nil {
		t.Fatal(optionErr)
	}
	return value
}

func TestSocketTLSWriteShutdownSignalsEOFAndPreservesReadSide(t *testing.T) {
	client, server := socketTestTLSPair(t)
	SocketSetTimeout(client, 1)
	SocketSetTimeout(server, 1)

	if written := SocketWriteValues(client, []int{'p', 'i', 'n', 'g'}); written.Status != SocketIOReady || written.Count != 4 {
		t.Fatalf("TLS request write = %#v, want four ready bytes", written)
	}
	SocketShutdown(client, false, true)
	if got := socketTestReadExact(t, server, 4); !socketTestIntsEqual(got, 'p', 'i', 'n', 'g') {
		t.Fatalf("TLS request = %v, want ping", got)
	}
	if eof := SocketReadValues(server, 1); eof.Status != SocketIOEOF {
		t.Fatalf("TLS read after peer write shutdown = %#v, want EOF", eof)
	}

	if written := SocketWriteValues(server, []int{'o', 'k'}); written.Status != SocketIOReady || written.Count != 2 {
		t.Fatalf("TLS response after peer write shutdown = %#v, want two ready bytes", written)
	}
	if got := socketTestReadExact(t, client, 2); !socketTestIntsEqual(got, 'o', 'k') {
		t.Fatalf("TLS response = %v, want ok", got)
	}
	if recovered := socketTestRecovered(func() {
		SocketShutdown(client, false, true)
	}); recovered != nil {
		t.Fatalf("repeated TLS write shutdown recovered %#v", recovered)
	}
}

func TestSocketTLSReadOnlyShutdownIsExplicitAndFullShutdownIsIdempotent(t *testing.T) {
	client, server := socketTestTLSPair(t)
	SocketSetTimeout(client, 1)
	SocketSetTimeout(server, 1)

	recovered := socketTestRecovered(func() {
		SocketShutdown(client, true, false)
	})
	exception, ok := recovered.(HaxeException)
	if !ok || *StdString(exception.Value) != "TLS read-only shutdown is unsupported" {
		t.Fatalf("TLS read-only shutdown recovered %#v, want exact HaxeException", recovered)
	}
	if written := SocketWriteValues(client, []int{'x'}); written.Status != SocketIOReady || written.Count != 1 {
		t.Fatalf("TLS write after rejected read shutdown = %#v", written)
	}
	if got := socketTestReadExact(t, server, 1); !socketTestIntsEqual(got, 'x') {
		t.Fatalf("TLS peer read = %v, want x", got)
	}

	SocketShutdown(client, true, true)
	if connection := client.snapshotConn(); connection != nil {
		t.Fatalf("TLS full shutdown retained %T", connection)
	}
	if recovered := socketTestRecovered(func() {
		SocketShutdown(client, true, true)
	}); recovered != nil {
		t.Fatalf("repeated TLS full shutdown recovered %#v", recovered)
	}
	if eof := SocketReadValues(server, 1); eof.Status != SocketIOEOF {
		t.Fatalf("TLS peer read after full shutdown = %#v, want EOF", eof)
	}
}

func TestSocketTLSFastSendReachesUnderlyingTCPAndUnsupportedConnectionsFail(t *testing.T) {
	client, _ := socketTestTLSPair(t)

	if err := socketTestTCPConnection(t, client).SetNoDelay(false); err != nil {
		t.Fatal(err)
	}
	if value := socketTestNoDelay(t, client); value != 0 {
		t.Fatalf("TLS TCP_NODELAY test baseline = %d, want 0", value)
	}
	SocketSetFastSend(client, true)
	if value := socketTestNoDelay(t, client); value == 0 {
		t.Fatal("TLS TCP_NODELAY remained disabled after setFastSend(true)")
	}
	SocketSetFastSend(client, false)
	if value := socketTestNoDelay(t, client); value != 0 {
		t.Fatalf("TLS TCP_NODELAY after setFastSend(false) = %d, want 0", value)
	}

	left, right := net.Pipe()
	defer right.Close()
	unsupported := SocketNewTCP()
	unsupported.installConn(left)
	defer SocketClose(unsupported)
	if recovered := socketTestRecovered(func() {
		SocketSetFastSend(unsupported, true)
	}); recovered == nil {
		t.Fatal("setFastSend silently accepted a connection without TCP control")
	}
}

func TestSocketPlainTCPHalfCloseRemainsBidirectional(t *testing.T) {
	_, client, server := socketTestTCPPair(t)
	SocketSetTimeout(client, 1)
	SocketSetTimeout(server, 1)

	if written := SocketWriteValues(client, []int{'q'}); written.Status != SocketIOReady || written.Count != 1 {
		t.Fatalf("TCP request write = %#v", written)
	}
	SocketShutdown(client, false, true)
	if got := socketTestReadExact(t, server, 1); !socketTestIntsEqual(got, 'q') {
		t.Fatalf("TCP peer read = %v, want q", got)
	}
	if eof := SocketReadValues(server, 1); eof.Status != SocketIOEOF {
		t.Fatalf("TCP peer read after write shutdown = %#v, want EOF", eof)
	}
	if written := SocketWriteValues(server, []int{'r'}); written.Status != SocketIOReady || written.Count != 1 {
		t.Fatalf("TCP response write = %#v", written)
	}
	if got := socketTestReadExact(t, client, 1); !socketTestIntsEqual(got, 'r') {
		t.Fatalf("TCP response = %v, want r", got)
	}
}

func TestSocketTLSShutdownAndCloseAreConcurrentSafe(t *testing.T) {
	client, _ := socketTestTLSPair(t)

	const workers = 24
	recovered := make(chan any, workers)
	var group sync.WaitGroup
	for index := 0; index < workers; index++ {
		group.Add(1)
		go func(index int) {
			defer group.Done()
			recovered <- socketTestRecovered(func() {
				if index%2 == 0 {
					SocketShutdown(client, false, true)
				} else {
					SocketClose(client)
				}
			})
		}(index)
	}
	group.Wait()
	close(recovered)
	for value := range recovered {
		if value != nil {
			t.Fatalf("concurrent TLS shutdown recovered %#v", value)
		}
	}
	SocketClose(client)
}
