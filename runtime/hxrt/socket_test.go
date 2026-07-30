package hxrt

import (
	"bytes"
	"io"
	"net"
	"sync"
	"testing"
	"time"
)

const socketTestHost = "127.0.0.1"

func socketTestString(value string) *string {
	return &value
}

func socketTestIntsEqual(actual []int, expected ...int) bool {
	if len(actual) != len(expected) {
		return false
	}
	for index := range actual {
		if actual[index] != expected[index] {
			return false
		}
	}
	return true
}

type socketTestPartialWriteConn struct {
	bytes.Buffer
	maxWrite int
}

func (connection *socketTestPartialWriteConn) Read(_ []byte) (int, error) {
	return 0, io.EOF
}

func (connection *socketTestPartialWriteConn) Write(raw []byte) (int, error) {
	limit := len(raw)
	if limit > connection.maxWrite {
		limit = connection.maxWrite
	}
	return connection.Buffer.Write(raw[:limit])
}

func (connection *socketTestPartialWriteConn) Close() error                      { return nil }
func (connection *socketTestPartialWriteConn) LocalAddr() net.Addr               { return nil }
func (connection *socketTestPartialWriteConn) RemoteAddr() net.Addr              { return nil }
func (connection *socketTestPartialWriteConn) SetDeadline(_ time.Time) error     { return nil }
func (connection *socketTestPartialWriteConn) SetReadDeadline(_ time.Time) error { return nil }
func (connection *socketTestPartialWriteConn) SetWriteDeadline(_ time.Time) error {
	return nil
}

func socketTestTCPPair(t *testing.T) (*SocketHandle, *SocketHandle, *SocketHandle) {
	t.Helper()
	listener := SocketNewTCP()
	SocketBindTCP(listener, socketTestString(socketTestHost), 0)
	SocketListen(listener, 1)
	bound := SocketHost(listener)
	if bound == nil || bound.Port <= 0 {
		t.Fatalf("SocketHost(listener) = %#v, want a positive port", bound)
	}

	acceptedResult := make(chan *SocketAcceptResult, 1)
	go func() {
		acceptedResult <- SocketAccept(listener)
	}()
	client := SocketNewTCP()
	SocketConnectTCP(client, socketTestString(socketTestHost), bound.Port)

	select {
	case result := <-acceptedResult:
		if result == nil || result.Status != SocketIOReady || result.Handle == nil {
			t.Fatalf("SocketAccept = %#v, want a ready typed handle", result)
		}
		return listener, client, result.Handle
	case <-time.After(2 * time.Second):
		SocketClose(client)
		SocketClose(listener)
		t.Fatal("SocketAccept did not complete")
		return nil, nil, nil
	}
}

func TestSocketTypedTCPHandleRoundTrip(t *testing.T) {
	listener, client, accepted := socketTestTCPPair(t)
	defer SocketClose(listener)
	defer SocketClose(client)
	defer SocketClose(accepted)

	written := SocketWriteValues(client, []int{'p', 'i', 'n', 'g'})
	if written.Status != SocketIOReady || written.Count != 4 {
		t.Fatalf("SocketWriteValues = %#v, want four ready bytes", written)
	}
	read := SocketReadValues(accepted, 4)
	if read.Status != SocketIOReady || !socketTestIntsEqual(read.Values, 'p', 'i', 'n', 'g') {
		t.Fatalf("SocketReadValues = %#v, want ping", read)
	}

	peer := SocketPeer(client)
	local := SocketHost(client)
	if peer == nil || peer.Host == 0 || peer.Port <= 0 {
		t.Fatalf("SocketPeer(client) = %#v, want a loopback peer", peer)
	}
	if local == nil || local.Host == 0 || local.Port <= 0 {
		t.Fatalf("SocketHost(client) = %#v, want a loopback address", local)
	}
}

func TestSocketWriteReportsPartialProgressForSourceOwnedWriteFullBytes(t *testing.T) {
	connection := &socketTestPartialWriteConn{maxWrite: 2}
	handle := SocketNewTCP()
	handle.installConn(connection)
	defer SocketClose(handle)

	first := SocketWriteValues(handle, []int{'a', 'b', 'c', 'd'})
	if first.Status != SocketIOReady || first.Count != 2 {
		t.Fatalf("first partial write = %#v, want two ready bytes", first)
	}
	second := SocketWriteValues(handle, []int{'c', 'd'})
	if second.Status != SocketIOReady || second.Count != 2 {
		t.Fatalf("second partial write = %#v, want two ready bytes", second)
	}
	if got := connection.String(); got != "abcd" {
		t.Fatalf("written bytes = %q, want abcd", got)
	}
}

func TestSocketTimeoutAndReadinessAreExplicit(t *testing.T) {
	listener, client, accepted := socketTestTCPPair(t)
	defer SocketClose(listener)
	defer SocketClose(client)
	defer SocketClose(accepted)

	SocketSetTimeout(client, 0.02)
	started := time.Now()
	read := SocketReadValues(client, 1)
	if read.Status != SocketIOBlocked {
		t.Fatalf("timed read = %#v, want SocketIOBlocked", read)
	}
	if elapsed := time.Since(started); elapsed < 5*time.Millisecond || elapsed > time.Second {
		t.Fatalf("timed read elapsed %s, want a bounded timeout", elapsed)
	}

	SocketSetTimeout(client, -1)
	go func() {
		time.Sleep(10 * time.Millisecond)
		SocketWriteValues(accepted, []int{'x'})
	}()
	selected := SocketSelect([]*SocketHandle{client}, nil, nil, 1, true)
	if selected == nil || !socketTestIntsEqual(selected.Read, 0) {
		t.Fatalf("SocketSelect = %#v, want read index 0", selected)
	}
	if value := SocketReadByteValue(client); value != int('x') {
		t.Fatalf("SocketReadByteValue = %d, want x", value)
	}
}

func TestSocketAcceptAndUDPReadHonorConfiguredTimeouts(t *testing.T) {
	listener := SocketNewTCP()
	defer SocketClose(listener)
	SocketBindTCP(listener, socketTestString(socketTestHost), 0)
	SocketListen(listener, 1)
	SocketSetTimeout(listener, 0.02)
	acceptStarted := time.Now()
	accepted := SocketAccept(listener)
	if accepted.Status != SocketIOBlocked || accepted.Handle != nil {
		t.Fatalf("timed accept = %#v, want blocked without a handle", accepted)
	}
	if elapsed := time.Since(acceptStarted); elapsed < 5*time.Millisecond || elapsed > time.Second {
		t.Fatalf("timed accept took %s, want configured bounded timeout", elapsed)
	}

	udp := SocketNewUDP()
	defer SocketClose(udp)
	SocketUdpBind(udp, socketTestString(socketTestHost), 0)
	SocketSetTimeout(udp, 0.02)
	readStarted := time.Now()
	read := SocketUdpReadFrom(udp, 8)
	if read.Status != SocketIOBlocked {
		t.Fatalf("timed UDP read = %#v, want blocked", read)
	}
	if elapsed := time.Since(readStarted); elapsed < 5*time.Millisecond || elapsed > time.Second {
		t.Fatalf("timed UDP read took %s, want configured bounded timeout", elapsed)
	}
}

func TestSslSocketConnectHonorsTheConfiguredHandshakeTimeout(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	accepted := make(chan net.Conn, 1)
	go func() {
		connection, acceptErr := listener.Accept()
		if acceptErr == nil {
			accepted <- connection
		}
	}()

	address := listener.Addr().(*net.TCPAddr)
	handle := SocketNewTCP()
	SocketSetTimeout(handle, 0.02)
	recovered := make(chan any, 1)
	started := time.Now()
	go func() {
		defer func() {
			recovered <- recover()
		}()
		SslSocketConnect(handle, socketTestString(socketTestHost), address.Port, false, nil, nil, nil, nil)
	}()

	select {
	case value := <-recovered:
		if _, ok := value.(HaxeException); !ok {
			t.Fatalf("stalled TLS handshake recovered %#v, want HaxeException timeout", value)
		}
		if elapsed := time.Since(started); elapsed < 5*time.Millisecond || elapsed > time.Second {
			t.Fatalf("stalled TLS handshake took %s, want configured bounded timeout", elapsed)
		}
	case connection := <-accepted:
		defer connection.Close()
		select {
		case value := <-recovered:
			if _, ok := value.(HaxeException); !ok {
				t.Fatalf("stalled TLS handshake recovered %#v, want HaxeException timeout", value)
			}
			if elapsed := time.Since(started); elapsed < 5*time.Millisecond || elapsed > time.Second {
				t.Fatalf("stalled TLS handshake took %s, want configured bounded timeout", elapsed)
			}
		case <-time.After(200 * time.Millisecond):
			_ = connection.Close()
			<-recovered
			t.Fatal("stalled TLS handshake ignored Socket.setTimeout")
		}
	case <-time.After(time.Second):
		t.Fatal("TLS client did not connect to the local stalled peer")
	}
}

func TestSocketPeerClosePreservesPartialReadThenReportsEOF(t *testing.T) {
	local, peer := net.Pipe()
	handle := SocketNewTCP()
	handle.installConn(local)
	defer SocketClose(handle)

	peerClosed := make(chan struct{})
	go func() {
		_, _ = peer.Write([]byte("xy"))
		_ = peer.Close()
		close(peerClosed)
	}()

	first := SocketReadValues(handle, 8)
	if first.Status != SocketIOReady || first.Count != 2 || !socketTestIntsEqual(first.Values, 'x', 'y') {
		t.Fatalf("partial peer-close read = %#v, want two ready bytes", first)
	}
	<-peerClosed
	second := SocketReadValues(handle, 8)
	if second.Status != SocketIOEOF || second.Count != 0 {
		t.Fatalf("read after peer close = %#v, want EOF", second)
	}
}

func TestSocketTypedUDPHandleRoundTripAndBroadcast(t *testing.T) {
	server := SocketNewUDP()
	client := SocketNewUDP()
	defer SocketClose(server)
	defer SocketClose(client)

	SocketUdpBind(server, socketTestString(socketTestHost), 0)
	bound := SocketHost(server)
	if bound == nil || bound.Port <= 0 {
		t.Fatalf("SocketHost(udp) = %#v, want a positive port", bound)
	}
	SocketUdpSetBroadcast(client, true)
	written := SocketUdpSendTo(client, []int{'u', 'd', 'p'}, HostResolve(socketTestString(socketTestHost)), bound.Port)
	if written.Status != SocketIOReady || written.Count != 3 {
		t.Fatalf("SocketUdpSendTo = %#v, want three ready bytes", written)
	}
	read := SocketUdpReadFrom(server, 16)
	if read.Status != SocketIOReady || !socketTestIntsEqual(read.Values, 'u', 'd', 'p') {
		t.Fatalf("SocketUdpReadFrom = %#v, want udp payload", read)
	}
	if read.Host == 0 || read.Port <= 0 {
		t.Fatalf("SocketUdpReadFrom peer = %d:%d, want loopback", read.Host, read.Port)
	}
}

func TestSocketConcurrentCloseIsIdempotentAndUnblocksRead(t *testing.T) {
	left, right := net.Pipe()
	defer right.Close()
	handle := SocketNewTCP()
	handle.installConn(left)

	readDone := make(chan *SocketIOResult, 1)
	go func() {
		readDone <- SocketReadValues(handle, 1)
	}()

	var group sync.WaitGroup
	for index := 0; index < 32; index++ {
		group.Add(1)
		go func(index int) {
			defer group.Done()
			SocketSetBlocking(handle, index%2 == 0)
			SocketSetTimeout(handle, float64(index%3)/1000)
			SocketSetFastSend(handle, index%2 == 1)
			SocketClose(handle)
		}(index)
	}
	group.Wait()

	select {
	case result := <-readDone:
		if result.Status != SocketIOEOF && result.Status != SocketIOBlocked {
			t.Fatalf("read after close = %#v, want EOF or blocked", result)
		}
	case <-time.After(time.Second):
		t.Fatal("SocketClose did not unblock a pending read")
	}
	SocketClose(handle)
}

func TestSocketWaitForReadReturnsForClosedHandle(t *testing.T) {
	handle := SocketNewTCP()
	SocketClose(handle)
	done := make(chan struct{})
	go func() {
		SocketWaitForRead(handle)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("SocketWaitForRead blocked forever after the handle was closed")
	}
}

func TestSocketConcurrentUDPLazyInitializationUsesOneConnection(t *testing.T) {
	handle := SocketNewUDP()
	defer SocketClose(handle)

	const workers = 64
	start := make(chan struct{})
	connections := make(chan *net.UDPConn, workers)
	var group sync.WaitGroup
	for index := 0; index < workers; index++ {
		group.Add(1)
		go func() {
			defer group.Done()
			<-start
			connections <- handle.udpConn(true)
		}()
	}
	close(start)
	group.Wait()
	close(connections)

	var first *net.UDPConn
	for connection := range connections {
		if connection == nil {
			t.Fatal("udpConn(true) returned nil")
		}
		if first == nil {
			first = connection
			continue
		}
		if connection != first {
			t.Fatalf("udpConn(true) returned multiple native connections: %p and %p", first, connection)
		}
	}
}
