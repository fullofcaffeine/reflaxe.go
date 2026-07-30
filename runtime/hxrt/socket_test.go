package hxrt

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"math/big"
	"net"
	"strconv"
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

func socketTestCertificate(t *testing.T, logicalHost string) (*tls.Config, *SslCertificate) {
	t.Helper()
	now := time.Now()
	rootKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	rootTemplate := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "haxe.go socket test root"},
		NotBefore:             now.Add(-time.Hour),
		NotAfter:              now.Add(time.Hour),
		IsCA:                  true,
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageCertSign,
	}
	rootDER, err := x509.CreateCertificate(rand.Reader, rootTemplate, rootTemplate, &rootKey.PublicKey, rootKey)
	if err != nil {
		t.Fatal(err)
	}
	root, err := x509.ParseCertificate(rootDER)
	if err != nil {
		t.Fatal(err)
	}

	leafKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	leafTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: logicalHost},
		DNSNames:     []string{logicalHost},
		NotBefore:    now.Add(-time.Hour),
		NotAfter:     now.Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	leafDER, err := x509.CreateCertificate(rand.Reader, leafTemplate, root, &leafKey.PublicKey, rootKey)
	if err != nil {
		t.Fatal(err)
	}
	leaf, err := x509.ParseCertificate(leafDER)
	if err != nil {
		t.Fatal(err)
	}
	return &tls.Config{Certificates: []tls.Certificate{{
		Certificate: [][]byte{leafDER, rootDER},
		PrivateKey:  leafKey,
		Leaf:        leaf,
	}}}, newSslCertificate([]*x509.Certificate{root}, nil)
}

func socketTestRecovered(call func()) (recovered any) {
	defer func() {
		recovered = recover()
	}()
	call()
	return nil
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

func TestSocketBindReservesEndpointWithoutListening(t *testing.T) {
	listener := SocketNewTCP()
	defer SocketClose(listener)
	SocketBindTCP(listener, socketTestString(socketTestHost), 0)
	bound := SocketHost(listener)
	if bound == nil || bound.Port <= 0 {
		t.Fatalf("SocketHost(bound) = %#v, want a reserved positive port", bound)
	}
	address := net.JoinHostPort(socketTestHost, strconv.Itoa(bound.Port))

	beforeListen, err := net.DialTimeout("tcp4", address, 20*time.Millisecond)
	if err == nil {
		_ = beforeListen.Close()
		t.Fatal("TCP connection succeeded after bind but before listen")
	}

	SocketListen(listener, 1)
	afterListen, err := net.DialTimeout("tcp4", address, time.Second)
	if err != nil {
		t.Fatalf("TCP connection after listen failed: %v", err)
	}
	_ = afterListen.Close()
}

func TestSocketListenValidatesLifecycleAndBacklog(t *testing.T) {
	unbound := SocketNewTCP()
	defer SocketClose(unbound)
	if recovered := socketTestRecovered(func() {
		SocketListen(unbound, 1)
	}); recovered == nil {
		t.Fatal("SocketListen on an unbound handle did not fail")
	}

	bound := SocketNewTCP()
	defer SocketClose(bound)
	SocketBindTCP(bound, socketTestString(socketTestHost), 0)
	if recovered := socketTestRecovered(func() {
		SocketListen(bound, -1)
	}); recovered == nil {
		t.Fatal("SocketListen accepted a negative backlog")
	}
	if recovered := socketTestRecovered(func() {
		SocketListen(bound, 1)
	}); recovered != nil {
		t.Fatalf("first SocketListen recovered %#v", recovered)
	}
	if recovered := socketTestRecovered(func() {
		SocketListen(bound, 2)
	}); recovered != nil {
		t.Fatalf("duplicate SocketListen recovered %#v", recovered)
	}

	closed := SocketNewTCP()
	SocketBindTCP(closed, socketTestString(socketTestHost), 0)
	closedAddress := SocketHost(closed)
	SocketClose(closed)
	if recovered := socketTestRecovered(func() {
		SocketListen(closed, 1)
	}); recovered == nil {
		t.Fatal("SocketListen after close did not fail")
	}
	rebound, err := net.Listen(
		"tcp4",
		net.JoinHostPort(socketTestHost, strconv.Itoa(closedAddress.Port)),
	)
	if err != nil {
		t.Fatalf("bind then close did not release reserved endpoint: %v", err)
	}
	_ = rebound.Close()
}

func TestSocketAcceptInheritsListenerPolicyAfterListen(t *testing.T) {
	listener := SocketNewTCP()
	client := SocketNewTCP()
	defer SocketClose(listener)
	defer SocketClose(client)

	SocketBindTCP(listener, socketTestString(socketTestHost), 0)
	SocketListen(listener, 1)
	SocketSetTimeout(listener, 0.25)
	SocketSetFastSend(listener, true)
	bound := SocketHost(listener)
	SocketConnectTCP(client, socketTestString(socketTestHost), bound.Port)

	result := SocketAccept(listener)
	if result == nil || result.Status != SocketIOReady || result.Handle == nil {
		t.Fatalf("SocketAccept = %#v, want a ready inherited handle", result)
	}
	defer SocketClose(result.Handle)
	if !result.Handle.blocking {
		t.Fatal("accepted handle did not inherit the listener's blocking policy")
	}
	if !result.Handle.hasTimeout || result.Handle.timeout != 0.25 {
		t.Fatalf(
			"accepted timeout = (%v, %v), want configured 0.25",
			result.Handle.hasTimeout,
			result.Handle.timeout,
		)
	}
	if !result.Handle.fastSend {
		t.Fatal("accepted handle did not inherit fast-send policy")
	}
}

func TestSocketTLSWrapperStartsOnlyAtListen(t *testing.T) {
	serverConfig, _ := socketTestCertificate(t, socketTestHost)
	listener := SocketNewTCP()
	defer SocketClose(listener)
	socketBindTCP(
		listener,
		socketTestString(socketTestHost),
		0,
		func(raw net.Listener) net.Listener {
			return tls.NewListener(raw, serverConfig)
		},
	)
	bound := SocketHost(listener)
	address := net.JoinHostPort(socketTestHost, strconv.Itoa(bound.Port))

	beforeListen, err := net.DialTimeout("tcp4", address, 20*time.Millisecond)
	if err == nil {
		_ = beforeListen.Close()
		t.Fatal("TLS TCP connection succeeded after bind but before listen")
	}
	SocketListen(listener, 1)

	serverResult := make(chan any, 1)
	go func() {
		result := SocketAccept(listener)
		if result == nil || result.Status != SocketIOReady || result.Handle == nil {
			serverResult <- result
			return
		}
		defer SocketClose(result.Handle)
		serverResult <- socketTestRecovered(func() {
			SslSocketHandshake(result.Handle)
		})
	}()

	client, err := tls.Dial(
		"tcp4",
		address,
		&tls.Config{InsecureSkipVerify: true},
	)
	if err != nil {
		t.Fatalf("TLS dial after listen failed: %v", err)
	}
	_ = client.Close()
	select {
	case result := <-serverResult:
		if result != nil {
			t.Fatalf("TLS accept/handshake result = %#v, want nil", result)
		}
	case <-time.After(time.Second):
		t.Fatal("TLS accept/handshake did not complete")
	}
}

func TestSocketConcurrentListenAndCloseHaveDeterministicOutcomes(t *testing.T) {
	for attempt := 0; attempt < 100; attempt++ {
		handle := SocketNewTCP()
		SocketBindTCP(handle, socketTestString(socketTestHost), 0)
		start := make(chan struct{})
		listenResult := make(chan any, 1)
		closeResult := make(chan any, 1)

		go func() {
			<-start
			listenResult <- socketTestRecovered(func() {
				SocketListen(handle, 1)
			})
		}()
		go func() {
			<-start
			closeResult <- socketTestRecovered(func() {
				SocketClose(handle)
			})
		}()
		close(start)

		listenRecovered := <-listenResult
		closeRecovered := <-closeResult
		if closeRecovered != nil {
			t.Fatalf("attempt %d close recovered %#v", attempt, closeRecovered)
		}
		if listenRecovered != nil {
			if _, ok := listenRecovered.(HaxeException); !ok {
				t.Fatalf("attempt %d listen recovered %#v, want HaxeException", attempt, listenRecovered)
			}
		}
		if recovered := socketTestRecovered(func() {
			SocketClose(handle)
		}); recovered != nil {
			t.Fatalf("attempt %d final close recovered %#v", attempt, recovered)
		}
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

func socketTestSaturateTCPWriteBuffer(t *testing.T, sender *SocketHandle, receiver *SocketHandle) {
	t.Helper()
	senderConn, ok := sender.snapshotConn().(*net.TCPConn)
	if !ok {
		t.Fatalf("sender connection = %T, want *net.TCPConn", sender.snapshotConn())
	}
	receiverConn, ok := receiver.snapshotConn().(*net.TCPConn)
	if !ok {
		t.Fatalf("receiver connection = %T, want *net.TCPConn", receiver.snapshotConn())
	}
	if err := senderConn.SetWriteBuffer(4096); err != nil {
		t.Fatal(err)
	}
	if err := receiverConn.SetReadBuffer(4096); err != nil {
		t.Fatal(err)
	}
	socketTestSaturateTCPWriteBufferNative(t, senderConn)
}

func TestSocketSelectUsesActualWriteReadinessAndBecomesReadyAfterDrain(t *testing.T) {
	listener, client, accepted := socketTestTCPPair(t)
	defer SocketClose(listener)
	defer SocketClose(client)
	defer SocketClose(accepted)

	socketTestSaturateTCPWriteBuffer(t, client, accepted)
	saturated := SocketSelect(nil, []*SocketHandle{client}, nil, 0, true)
	if saturated == nil || len(saturated.Write) != 0 {
		t.Fatalf("saturated SocketSelect write = %#v, want no ready index", saturated)
	}
	SocketSetBlocking(client, false)
	nonblockingPayload := make([]int, 64*1024)
	blocked := false
	for attempt := 0; attempt < 64; attempt++ {
		write := SocketWriteValues(client, nonblockingPayload)
		if write.Status == SocketIOBlocked && write.Count == 0 {
			blocked = true
			break
		}
		if write.Status != SocketIOReady || write.Count <= 0 {
			t.Fatalf("saturated nonblocking write attempt %d = %#v", attempt, write)
		}
	}
	if !blocked {
		t.Fatal("saturated nonblocking writes never reported blocked")
	}
	SocketSetBlocking(client, true)

	drained := make(chan struct{})
	go func() {
		defer close(drained)
		buffer := make([]byte, 64*1024)
		connection := accepted.snapshotConn()
		for {
			if _, err := connection.Read(buffer); err != nil {
				return
			}
		}
	}()
	ready := SocketSelect(nil, []*SocketHandle{client}, nil, 1, true)
	if ready == nil || !socketTestIntsEqual(ready.Write, 0) {
		t.Fatalf("drained SocketSelect write = %#v, want index 0", ready)
	}
	SocketClose(client)
	SocketClose(accepted)
	select {
	case <-drained:
	case <-time.After(time.Second):
		t.Fatal("send-buffer drain did not stop after close")
	}
}

func TestSocketNonblockingConnectedReadWriteAndAcceptAreImmediate(t *testing.T) {
	listener, client, accepted := socketTestTCPPair(t)
	defer SocketClose(listener)
	defer SocketClose(client)
	defer SocketClose(accepted)

	SocketSetBlocking(client, false)
	started := time.Now()
	read := SocketReadValues(client, 1)
	if read.Status != SocketIOBlocked || read.Count != 0 {
		t.Fatalf("empty nonblocking read = %#v, want blocked", read)
	}
	if elapsed := time.Since(started); elapsed > 100*time.Millisecond {
		t.Fatalf("empty nonblocking read took %s", elapsed)
	}
	write := SocketWriteValues(client, []int{'n'})
	if write.Status != SocketIOReady || write.Count != 1 {
		t.Fatalf("available nonblocking write = %#v, want one ready byte", write)
	}
	if received := SocketReadByteValue(accepted); received != int('n') {
		t.Fatalf("nonblocking write peer read = %d, want n", received)
	}

	server := SocketNewTCP()
	connecter := SocketNewTCP()
	defer SocketClose(server)
	defer SocketClose(connecter)
	SocketBindTCP(server, socketTestString(socketTestHost), 0)
	SocketListen(server, 1)
	SocketSetBlocking(server, false)
	started = time.Now()
	emptyAccept := SocketAccept(server)
	if emptyAccept.Status != SocketIOBlocked || emptyAccept.Handle != nil {
		t.Fatalf("empty nonblocking accept = %#v, want blocked", emptyAccept)
	}
	if elapsed := time.Since(started); elapsed > 100*time.Millisecond {
		t.Fatalf("empty nonblocking accept took %s", elapsed)
	}
	bound := SocketHost(server)
	SocketConnectTCP(connecter, socketTestString(socketTestHost), bound.Port)
	readyAccept := SocketAccept(server)
	if readyAccept.Status != SocketIOReady || readyAccept.Handle == nil {
		t.Fatalf("queued nonblocking accept = %#v, want ready handle", readyAccept)
	}
	SocketClose(readyAccept.Handle)
}

func TestSocketSelectDoesNotFabricateDisconnectedExceptions(t *testing.T) {
	unconnected := SocketNewTCP()
	closed := SocketNewTCP()
	SocketClose(closed)
	defer SocketClose(unconnected)

	selected := SocketSelect(nil, nil, []*SocketHandle{unconnected, closed}, 0, true)
	if selected == nil || len(selected.Others) != 0 {
		t.Fatalf("disconnected SocketSelect others = %#v, want no fabricated exception", selected)
	}
}

func TestSocketSelectPreservesDuplicateReadIndexesBufferedBytesAndEOF(t *testing.T) {
	listener, client, accepted := socketTestTCPPair(t)
	defer SocketClose(listener)
	defer SocketClose(client)
	defer SocketClose(accepted)

	if written := SocketWriteValues(accepted, []int{'x', 'y'}); written.Status != SocketIOReady || written.Count != 2 {
		t.Fatalf("SocketWriteValues = %#v, want two bytes", written)
	}
	ready := SocketSelect([]*SocketHandle{client, client}, nil, nil, 1, true)
	if ready == nil || !socketTestIntsEqual(ready.Read, 0, 1) {
		t.Fatalf("duplicate SocketSelect read = %#v, want indexes 0 and 1", ready)
	}
	if value := SocketReadByteValue(client); value != int('x') {
		t.Fatalf("first read = %d, want x", value)
	}
	buffered := SocketSelect([]*SocketHandle{client}, nil, nil, 0, true)
	if buffered == nil || !socketTestIntsEqual(buffered.Read, 0) {
		t.Fatalf("buffered SocketSelect read = %#v, want index 0", buffered)
	}
	if value := SocketReadByteValue(client); value != int('y') {
		t.Fatalf("second read = %d, want y", value)
	}

	SocketClose(accepted)
	eof := SocketSelect([]*SocketHandle{client}, nil, nil, 1, true)
	if eof == nil || !socketTestIntsEqual(eof.Read, 0) {
		t.Fatalf("EOF SocketSelect read = %#v, want index 0", eof)
	}
	if value := SocketReadByteValue(client); value != SocketReadEOF {
		t.Fatalf("read after EOF readiness = %d, want SocketReadEOF", value)
	}
}

func TestSocketSelectMakesResetObservableThroughReadReadiness(t *testing.T) {
	listener, client, accepted := socketTestTCPPair(t)
	defer SocketClose(listener)
	defer SocketClose(client)

	acceptedConn, ok := accepted.snapshotConn().(*net.TCPConn)
	if !ok {
		t.Fatalf("accepted connection = %T, want *net.TCPConn", accepted.snapshotConn())
	}
	if err := acceptedConn.SetLinger(0); err != nil {
		t.Fatal(err)
	}
	SocketClose(accepted)

	selected := SocketSelect([]*SocketHandle{client}, nil, []*SocketHandle{client}, 1, true)
	if selected == nil || !socketTestIntsEqual(selected.Read, 0) {
		t.Fatalf("reset SocketSelect = %#v, want read index 0", selected)
	}
	var read *SocketIOResult
	recovered := socketTestRecovered(func() {
		read = SocketReadValues(client, 1)
	})
	if recovered != nil {
		if _, ok := recovered.(HaxeException); !ok {
			t.Fatalf("reset read recovered %#v, want HaxeException", recovered)
		}
		return
	}
	if read == nil || read.Status != SocketIOEOF {
		t.Fatalf("reset read = %#v, want EOF or a typed HaxeException", read)
	}
}

func TestSocketSelectZeroFiniteAndAbsentTimeoutsAreBounded(t *testing.T) {
	listener, client, accepted := socketTestTCPPair(t)
	defer SocketClose(listener)
	defer SocketClose(client)
	defer SocketClose(accepted)

	started := time.Now()
	immediate := SocketSelect([]*SocketHandle{client}, nil, nil, 0, true)
	if immediate == nil || len(immediate.Read) != 0 {
		t.Fatalf("zero-timeout SocketSelect = %#v, want empty", immediate)
	}
	if elapsed := time.Since(started); elapsed > 100*time.Millisecond {
		t.Fatalf("zero-timeout SocketSelect took %s", elapsed)
	}

	started = time.Now()
	finite := SocketSelect([]*SocketHandle{client}, nil, nil, 0.02, true)
	if finite == nil || len(finite.Read) != 0 {
		t.Fatalf("finite-timeout SocketSelect = %#v, want empty", finite)
	}
	if elapsed := time.Since(started); elapsed < 5*time.Millisecond || elapsed > time.Second {
		t.Fatalf("finite-timeout SocketSelect took %s", elapsed)
	}

	go func() {
		time.Sleep(10 * time.Millisecond)
		SocketWriteValues(accepted, []int{'z'})
	}()
	started = time.Now()
	absent := SocketSelect([]*SocketHandle{client}, nil, nil, 0, false)
	if absent == nil || !socketTestIntsEqual(absent.Read, 0) {
		t.Fatalf("absent-timeout SocketSelect = %#v, want read index 0", absent)
	}
	if elapsed := time.Since(started); elapsed < 5*time.Millisecond || elapsed > time.Second {
		t.Fatalf("absent-timeout SocketSelect took %s", elapsed)
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
		SslSocketConnect(handle, SocketEndpointNew(socketTestString(socketTestHost), socketTestString(socketTestHost)), address.Port, false, nil, nil, nil,
			nil)
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

func TestSslSocketConnectPreservesLogicalHostForVerificationAndSNI(t *testing.T) {
	const logicalHost = "logical.test"
	serverConfig, ca := socketTestCertificate(t, logicalHost)
	observedSNI := make(chan string, 4)
	serverConfig.GetConfigForClient = func(hello *tls.ClientHelloInfo) (*tls.Config, error) {
		observedSNI <- hello.ServerName
		return nil, nil
	}
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	serverErrors := make(chan error, 4)
	go func() {
		for index := 0; index < 4; index++ {
			raw, acceptErr := listener.Accept()
			if acceptErr != nil {
				serverErrors <- acceptErr
				continue
			}
			connection := tls.Server(raw, serverConfig)
			serverErrors <- connection.Handshake()
			_ = connection.Close()
		}
	}()

	port := listener.Addr().(*net.TCPAddr).Port
	endpoint := SocketEndpointNew(socketTestString(socketTestHost), socketTestString(logicalHost))
	connect := func(endpoint *SocketEndpoint, verify bool, serverName *string, ca *SslCertificate) any {
		handle := SocketNewTCP()
		defer SocketClose(handle)
		return socketTestRecovered(func() {
			SslSocketConnect(handle, endpoint, port, verify, ca, serverName, nil, nil)
		})
	}
	assertSNI := func(expected string) {
		t.Helper()
		select {
		case actual := <-observedSNI:
			if actual != expected {
				t.Fatalf("ClientHello SNI = %q, want %q", actual, expected)
			}
		case <-time.After(time.Second):
			t.Fatal("TLS server did not observe ClientHello SNI")
		}
	}

	if recovered := connect(endpoint, true, nil, ca); recovered != nil {
		t.Fatalf("default logical-host verification recovered %#v", recovered)
	}
	assertSNI(logicalHost)

	mismatch := SocketEndpointNew(socketTestString(socketTestHost), socketTestString("wrong.test"))
	if recovered := connect(mismatch, true, nil, ca); recovered == nil {
		t.Fatal("mismatched logical-host verification unexpectedly succeeded")
	}
	assertSNI("wrong.test")

	override := socketTestString(logicalHost)
	if recovered := connect(mismatch, true, override, ca); recovered != nil {
		t.Fatalf("explicit hostname override recovered %#v", recovered)
	}
	assertSNI(logicalHost)

	if recovered := connect(mismatch, false, nil, nil); recovered != nil {
		t.Fatalf("verifyCert=false recovered %#v", recovered)
	}
	assertSNI("wrong.test")

	for index := 0; index < 4; index++ {
		select {
		case serverErr := <-serverErrors:
			if index != 1 && serverErr != nil {
				t.Fatalf("TLS server handshake %d failed: %v", index, serverErr)
			}
		case <-time.After(time.Second):
			t.Fatal("TLS server handshake did not finish")
		}
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

func TestSocketZeroByteUDPDatagramPreservesSender(t *testing.T) {
	server := SocketNewUDP()
	client := SocketNewUDP()
	defer SocketClose(server)
	defer SocketClose(client)

	SocketUdpBind(server, socketTestString(socketTestHost), 0)
	bound := SocketHost(server)
	if bound == nil || bound.Port <= 0 {
		t.Fatalf("SocketHost(udp) = %#v, want a positive port", bound)
	}
	written := SocketUdpSendTo(client, []int{}, HostResolve(socketTestString(socketTestHost)), bound.Port)
	if written.Status != SocketIOReady || written.Count != 0 {
		t.Fatalf("zero-byte SocketUdpSendTo = %#v, want ready with zero bytes", written)
	}
	read := SocketUdpReadFrom(server, 1)
	if read.Status != SocketIOReady || read.Count != 0 || len(read.Values) != 0 {
		t.Fatalf("zero-byte SocketUdpReadFrom = %#v, want a ready empty datagram", read)
	}
	if read.Host != HostResolve(socketTestString(socketTestHost)) || read.Port <= 0 {
		t.Fatalf("zero-byte SocketUdpReadFrom peer = %d:%d, want loopback sender", read.Host, read.Port)
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
	recovered := make(chan any, 32)
	for index := 0; index < 32; index++ {
		group.Add(1)
		go func(index int) {
			defer group.Done()
			recovered <- socketTestRecovered(func() {
				SocketSetBlocking(handle, index%2 == 0)
				SocketSetTimeout(handle, float64(index%3)/1000)
				SocketSetFastSend(handle, index%2 == 1)
				SocketClose(handle)
			})
		}(index)
	}
	group.Wait()
	close(recovered)
	for value := range recovered {
		if value == nil {
			continue
		}
		exception, ok := value.(HaxeException)
		if !ok || *StdString(exception.Value) != "socket fast-send requires a TCP connection" {
			t.Fatalf("concurrent synthetic connection control recovered %#v", value)
		}
	}

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
