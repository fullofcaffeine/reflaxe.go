//go:build linux || darwin

package hxrt

import (
	"crypto/tls"
	"io"
	"net"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type socketResourceServer struct {
	listener          *net.TCPListener
	activeConnections atomic.Int64
	handlers          sync.WaitGroup
	acceptDone        chan struct{}
	handle            func(*net.TCPConn)
}

func newSocketResourceServer(t *testing.T, handle func(*net.TCPConn)) *socketResourceServer {
	t.Helper()
	listener, err := net.ListenTCP("tcp4", &net.TCPAddr{IP: net.ParseIP(socketTestHost), Port: 0})
	if err != nil {
		t.Fatal(err)
	}
	server := &socketResourceServer{
		listener:   listener,
		acceptDone: make(chan struct{}),
		handle:     handle,
	}
	go func() {
		defer close(server.acceptDone)
		for {
			connection, acceptErr := listener.AcceptTCP()
			if acceptErr != nil {
				return
			}
			server.activeConnections.Add(1)
			server.handlers.Add(1)
			go func() {
				defer server.handlers.Done()
				defer server.activeConnections.Add(-1)
				defer connection.Close()
				_ = connection.SetDeadline(time.Now().Add(2 * time.Second))
				handle(connection)
			}()
		}
	}()
	t.Cleanup(func() {
		_ = listener.Close()
		<-server.acceptDone
		server.handlers.Wait()
	})
	return server
}

func (server *socketResourceServer) port() int {
	return server.listener.Addr().(*net.TCPAddr).Port
}

func socketResourceOpenFDCount() int {
	entries, err := os.ReadDir("/proc/self/fd")
	if err == nil {
		return len(entries)
	}
	entries, err = os.ReadDir("/dev/fd")
	if err == nil {
		return len(entries)
	}
	return -1
}

func socketResourceActiveConnections(servers ...*socketResourceServer) int64 {
	var active int64
	for _, server := range servers {
		active += server.activeConnections.Load()
	}
	return active
}

func socketResourceAwaitActiveZero(t *testing.T, servers ...*socketResourceServer) {
	t.Helper()
	deadline := time.Now().Add(3 * time.Second)
	for socketResourceActiveConnections(servers...) != 0 {
		if time.Now().After(deadline) {
			t.Fatalf(
				"socket server connections did not quiesce: active=%d",
				socketResourceActiveConnections(servers...),
			)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func socketResourceAwaitConvergence(
	t *testing.T,
	baselineGoroutines int,
	baselineFDs int,
	servers ...*socketResourceServer,
) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for {
		runtime.GC()
		active := socketResourceActiveConnections(servers...)
		currentGoroutines := runtime.NumGoroutine()
		currentFDs := socketResourceOpenFDCount()
		goroutinesConverged := currentGoroutines <= baselineGoroutines+4
		fdsConverged := baselineFDs < 0 || currentFDs <= baselineFDs+2
		if active == 0 && goroutinesConverged && fdsConverged {
			t.Logf(
				"socket resources converged: active=%d goroutines=%d->%d fds=%d->%d",
				active,
				baselineGoroutines,
				currentGoroutines,
				baselineFDs,
				currentFDs,
			)
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf(
				"socket resources did not converge: active=%d, goroutines baseline=%d current=%d, fds baseline=%d current=%d",
				active,
				baselineGoroutines,
				currentGoroutines,
				baselineFDs,
				currentFDs,
			)
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func socketResourceTLSPair(t *testing.T, serverConfig *tls.Config) (*SocketHandle, *SocketHandle) {
	t.Helper()
	listener, err := net.Listen("tcp4", net.JoinHostPort(socketTestHost, "0"))
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	serverRaw := make(chan net.Conn, 1)
	serverErr := make(chan error, 1)
	go func() {
		connection, acceptErr := listener.Accept()
		if acceptErr != nil {
			serverErr <- acceptErr
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
	case acceptErr := <-serverErr:
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
	return client, server
}

func socketResourceExercise(
	t *testing.T,
	echo *socketResourceServer,
	stall *socketResourceServer,
	reset *socketResourceServer,
	serverTLS *tls.Config,
	attempt int,
) {
	t.Helper()

	client := SocketNewTCP()
	SocketConnectTCP(client, socketTestString(socketTestHost), echo.port())
	if write := SocketWriteValues(client, []int{'e'}); write.Status != SocketIOReady || write.Count != 1 {
		t.Fatalf("attempt %d TCP echo write = %#v", attempt, write)
	}
	if read := SocketReadValues(client, 1); read.Status != SocketIOReady || !socketTestIntsEqual(read.Values, 'e') {
		t.Fatalf("attempt %d TCP echo read = %#v", attempt, read)
	}
	SocketClose(client)
	SocketClose(client)

	timed := SocketNewTCP()
	SocketSetTimeout(timed, 0.01)
	SocketConnectTCP(timed, socketTestString(socketTestHost), stall.port())
	if write := SocketWriteValues(timed, []int{'t'}); write.Status != SocketIOReady || write.Count != 1 {
		t.Fatalf("attempt %d timeout write = %#v", attempt, write)
	}
	if read := SocketReadValues(timed, 1); read.Status != SocketIOBlocked {
		t.Fatalf("attempt %d timeout read = %#v, want blocked", attempt, read)
	}
	SocketClose(timed)

	resetClient := SocketNewTCP()
	SocketSetTimeout(resetClient, 0.5)
	resetConnectRecovered := socketTestRecovered(func() {
		SocketConnectTCP(resetClient, socketTestString(socketTestHost), reset.port())
	})
	if resetConnectRecovered != nil {
		if _, ok := resetConnectRecovered.(HaxeException); !ok {
			t.Fatalf("attempt %d reset connect recovered %#v", attempt, resetConnectRecovered)
		}
		if installed := resetClient.snapshotConn(); installed != nil {
			t.Fatalf("attempt %d failed reset connection remained installed as %T", attempt, installed)
		}
	} else {
		var resetRead *SocketIOResult
		resetReadRecovered := socketTestRecovered(func() {
			resetRead = SocketReadValues(resetClient, 1)
		})
		if resetReadRecovered == nil && (resetRead == nil || resetRead.Status != SocketIOEOF) {
			t.Fatalf("attempt %d reset read = %#v, want EOF or a Haxe socket error", attempt, resetRead)
		}
		if resetReadRecovered != nil {
			if _, ok := resetReadRecovered.(HaxeException); !ok {
				t.Fatalf("attempt %d reset read recovered %#v", attempt, resetReadRecovered)
			}
		}
	}
	SocketClose(resetClient)

	stalledTLS := SocketNewTCP()
	SocketSetTimeout(stalledTLS, 0.01)
	stalledRecovered := socketTestRecovered(func() {
		SslSocketConnect(
			stalledTLS,
			SocketEndpointNew(socketTestString(socketTestHost), socketTestString(socketTestHost)),
			stall.port(),
			false,
			nil,
			nil,
			nil,
			nil,
		)
	})
	if _, ok := stalledRecovered.(HaxeException); !ok {
		t.Fatalf("attempt %d stalled TLS recovered %#v, want HaxeException", attempt, stalledRecovered)
	}
	SocketClose(stalledTLS)

	udpServer := SocketNewUDP()
	udpClient := SocketNewUDP()
	SocketUdpBind(udpServer, socketTestString(socketTestHost), 0)
	SocketSetTimeout(udpServer, 0.5)
	udpAddress := SocketHost(udpServer)
	if sent := SocketUdpSendTo(udpClient, []int{}, udpAddress.Host, udpAddress.Port); sent.Status != SocketIOReady || sent.Count != 0 {
		t.Fatalf("attempt %d empty UDP send = %#v", attempt, sent)
	}
	if received := SocketUdpReadFrom(udpServer, 1); received.Status != SocketIOReady || received.Count != 0 || received.Port <= 0 {
		t.Fatalf("attempt %d empty UDP receive = %#v", attempt, received)
	}
	if sent := SocketUdpSendTo(udpClient, []int{'u'}, udpAddress.Host, udpAddress.Port); sent.Status != SocketIOReady || sent.Count != 1 {
		t.Fatalf("attempt %d UDP send = %#v", attempt, sent)
	}
	if received := SocketUdpReadFrom(udpServer, 1); received.Status != SocketIOReady || !socketTestIntsEqual(received.Values, 'u') {
		t.Fatalf("attempt %d UDP receive = %#v", attempt, received)
	}
	SocketClose(udpClient)
	SocketClose(udpServer)

	listener, readyClient, accepted := socketTestTCPPair(t)
	selected := SocketSelect(nil, []*SocketHandle{readyClient}, nil, 0, true)
	if selected == nil || !socketTestIntsEqual(selected.Write, 0) {
		t.Fatalf("attempt %d readiness = %#v, want write index 0", attempt, selected)
	}
	readDone := make(chan *SocketIOResult, 1)
	go func() {
		readDone <- SocketReadValues(readyClient, 1)
	}()
	time.Sleep(time.Millisecond)
	var closeGroup sync.WaitGroup
	for closer := 0; closer < 8; closer++ {
		closeGroup.Add(1)
		go func() {
			defer closeGroup.Done()
			SocketClose(readyClient)
		}()
	}
	closeGroup.Wait()
	select {
	case result := <-readDone:
		if result.Status != SocketIOEOF && result.Status != SocketIOBlocked {
			t.Fatalf("attempt %d canceled read = %#v", attempt, result)
		}
	case <-time.After(time.Second):
		t.Fatalf("attempt %d concurrent close did not unblock read", attempt)
	}
	SocketClose(accepted)
	SocketClose(listener)

	tlsClient, tlsServer := socketResourceTLSPair(t, serverTLS)
	SocketSetTimeout(tlsClient, 1)
	SocketSetTimeout(tlsServer, 1)
	if write := SocketWriteValues(tlsClient, []int{'s'}); write.Status != SocketIOReady || write.Count != 1 {
		t.Fatalf("attempt %d TLS write = %#v", attempt, write)
	}
	SocketShutdown(tlsClient, false, true)
	if read := SocketReadValues(tlsServer, 1); read.Status != SocketIOReady || !socketTestIntsEqual(read.Values, 's') {
		t.Fatalf("attempt %d TLS server read = %#v", attempt, read)
	}
	if eof := SocketReadValues(tlsServer, 1); eof.Status != SocketIOEOF {
		t.Fatalf("attempt %d TLS close-notify read = %#v", attempt, eof)
	}
	if write := SocketWriteValues(tlsServer, []int{'r'}); write.Status != SocketIOReady || write.Count != 1 {
		t.Fatalf("attempt %d TLS response write = %#v", attempt, write)
	}
	if read := SocketReadValues(tlsClient, 1); read.Status != SocketIOReady || !socketTestIntsEqual(read.Values, 'r') {
		t.Fatalf("attempt %d TLS response read = %#v", attempt, read)
	}
	SocketClose(tlsClient)
	SocketClose(tlsServer)
}

func TestSocketOperationResourcesConvergeAcrossFailureModes(t *testing.T) {
	echo := newSocketResourceServer(t, func(connection *net.TCPConn) {
		var payload [1]byte
		if _, err := io.ReadFull(connection, payload[:]); err == nil {
			_, _ = connection.Write(payload[:])
		}
	})
	stall := newSocketResourceServer(t, func(connection *net.TCPConn) {
		_, _ = io.Copy(io.Discard, connection)
	})
	reset := newSocketResourceServer(t, func(connection *net.TCPConn) {
		_ = connection.SetLinger(0)
	})
	serverTLS, _ := socketTestCertificate(t, socketTestHost)

	// Warm lazy netpoll, TLS, and crypto paths before measuring retained state.
	socketResourceExercise(t, echo, stall, reset, serverTLS, -1)
	socketResourceAwaitActiveZero(t, echo, stall, reset)
	runtime.GC()
	time.Sleep(30 * time.Millisecond)

	baselineGoroutines := runtime.NumGoroutine()
	baselineFDs := socketResourceOpenFDCount()
	if runtime.GOOS == "linux" && baselineFDs < 0 {
		t.Fatal("Linux file-descriptor baseline is unavailable")
	}

	const attempts = 20
	for attempt := 0; attempt < attempts; attempt++ {
		socketResourceExercise(t, echo, stall, reset, serverTLS, attempt)
	}
	socketResourceAwaitConvergence(
		t,
		baselineGoroutines,
		baselineFDs,
		echo,
		stall,
		reset,
	)
}
