package hxrt

import (
	"bufio"
	"errors"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	// SocketIOReady reports a successful native operation.
	SocketIOReady = iota
	// SocketIOEOF reports an orderly close or an unavailable socket resource.
	SocketIOEOF
	// SocketIOBlocked reports a timeout or nonblocking operation that made no progress.
	SocketIOBlocked
)

const (
	// SocketReadEOF is the scalar read sentinel translated to haxe.io.Eof by staged source.
	SocketReadEOF = -1
	// SocketReadBlocked is the scalar read sentinel translated to haxe.io.Error.Blocked.
	SocketReadBlocked = -2
)

const socketNonblockingProbeWindow = time.Millisecond

// SocketAddress is the typed native address carrier consumed by staged sys.net.
//
// What: Carries an IPv4 address in Haxe's network-order Int representation and a port.
// Why: Returning map[string]any would erase the native boundary and couple hxrt to
// generated anonymous-object layouts.
// How: Native operations populate this concrete carrier; staged Haxe constructs the
// public `{host: Host, port: Int}` object.
type SocketAddress struct {
	Host int
	Port int
}

// SocketEndpoint keeps a network address separate from its logical host identity.
//
// What: Carries the resolved address used for a native dial and the original
// source hostname used by protocols such as TLS.
// Why: Host.toString() intentionally renders the resolved IPv4 address, but
// using that value for certificate verification or SNI discards the caller's
// security identity.
// How: Staged Host code supplies both strings through SocketEndpointNew; native
// transports consume this opaque typed carrier without inspecting generated
// Haxe object layouts.
type SocketEndpoint struct {
	NetworkAddress string
	LogicalHost    string
}

// SocketEndpointNew constructs one typed endpoint at the staged/native boundary.
func SocketEndpointNew(networkAddress *string, logicalHost *string) *SocketEndpoint {
	if networkAddress == nil || *StdString(networkAddress) == "" {
		socketThrow(errors.New("socket endpoint requires a network address"))
		return &SocketEndpoint{}
	}
	address := *StdString(networkAddress)
	logical := address
	if logicalHost != nil && *StdString(logicalHost) != "" {
		logical = *StdString(logicalHost)
	}
	return &SocketEndpoint{NetworkAddress: address, LogicalHost: logical}
}

// SocketIOResult separates byte progress from EOF and blocked states.
//
// What: Returns byte values, their count, and one explicit status code.
// Why: EOF and Error.Blocked are public Haxe values that hxrt must not synthesize as
// generated target types.
// How: Native reads and writes classify Go errors; staged stream wrappers perform the
// final Haxe exception translation.
type SocketIOResult struct {
	Values []int
	Count  int
	Status int
}

// SocketDatagramResult adds the peer address to one typed UDP read result.
type SocketDatagramResult struct {
	Values []int
	Count  int
	Status int
	Host   int
	Port   int
}

// SocketAcceptResult separates an accepted handle from timeout/nonblocking state.
type SocketAcceptResult struct {
	Handle *SocketHandle
	Status int
}

// SocketSelectResult carries ready indexes rather than generated Socket objects.
// Staged Haxe maps each index back to the exact input object, preserving custom data
// and identity without putting public collection policy in hxrt.
type SocketSelectResult struct {
	Read   []int
	Write  []int
	Others []int
}

type socketDeadlineListener interface {
	SetDeadline(time.Time) error
}

type socketBoundTCP interface {
	Addr() net.Addr
	Listen(backlog int) (*net.TCPListener, error)
	Close() error
}

type socketListenerWrapper func(net.Listener) net.Listener

type socketSyscallResource interface {
	SyscallConn() (syscall.RawConn, error)
}

type socketNestedConnection interface {
	NetConn() net.Conn
}

type socketCloseReader interface {
	CloseRead() error
}

type socketCloseWriter interface {
	CloseWrite() error
}

type socketReadinessSnapshot struct {
	descriptor    uintptr
	hasDescriptor bool
	buffered      bool
	eof           bool
}

func (snapshot *socketReadinessSnapshot) release() {
	if snapshot == nil || !snapshot.hasDescriptor {
		return
	}
	_ = socketCloseDescriptor(snapshot.descriptor)
	snapshot.hasDescriptor = false
}

type socketSelectMode uint8

const (
	socketSelectRead socketSelectMode = iota
	socketSelectWrite
	socketSelectOthers
)

type socketSelectEntry struct {
	index         int
	descriptor    uintptr
	hasDescriptor bool
	immediate     bool
}

type socketNativeSelectRequest struct {
	Read    []uintptr
	Write   []uintptr
	Others  []uintptr
	Timeout time.Duration
}

type socketNativeSelectResult struct {
	Read   map[uintptr]struct{}
	Write  map[uintptr]struct{}
	Others map[uintptr]struct{}
}

func newSocketNativeSelectResult() *socketNativeSelectResult {
	return &socketNativeSelectResult{
		Read:   make(map[uintptr]struct{}),
		Write:  make(map[uintptr]struct{}),
		Others: make(map[uintptr]struct{}),
	}
}

// SocketHandle owns one native socket resource behind a typed opaque boundary.
//
// What: Stores a TCP/UDP connection, pre-listen TCP endpoint, or listener plus
// buffering and deadline policy.
// Why: OS resources cannot be represented as portable Haxe data, while exposing
// net.Conn as Dynamic would erase lifecycle ownership and make concurrent close racy.
// How: stateMu protects replaceable resources and policy, while readMu/writeMu
// serialize their respective stream operations. Close detaches resources before
// closing them so blocked operations are safely interrupted without holding stateMu.
type SocketHandle struct {
	stateMu sync.Mutex
	readMu  sync.Mutex
	writeMu sync.Mutex

	conn             net.Conn
	boundTCP         socketBoundTCP
	listener         net.Listener
	deadlineListener socketDeadlineListener
	listenerWrapper  socketListenerWrapper
	reader           *bufio.Reader

	timeout    float64
	hasTimeout bool
	blocking   bool
	fastSend   bool
	broadcast  bool
}

func newSocketHandle() *SocketHandle {
	return &SocketHandle{blocking: true}
}

// SocketNewTCP creates one unconnected typed TCP handle.
func SocketNewTCP() *SocketHandle {
	return newSocketHandle()
}

// SocketNewUDP creates one unbound typed UDP handle.
func SocketNewUDP() *SocketHandle {
	return newSocketHandle()
}

func socketThrow(err error) {
	if err != nil {
		Throw(StringFromLiteral(err.Error()))
	}
}

func socketErrorStatus(err error) int {
	if err == nil {
		return SocketIOReady
	}
	if socketErrorIsClosed(err) {
		return SocketIOEOF
	}
	var netError net.Error
	if errors.As(err, &netError) && netError.Timeout() {
		return SocketIOBlocked
	}
	if errors.Is(err, os.ErrDeadlineExceeded) {
		return SocketIOBlocked
	}
	return -1
}

func socketErrorIsClosed(err error) bool {
	return errors.Is(err, io.EOF) ||
		errors.Is(err, io.ErrClosedPipe) ||
		errors.Is(err, net.ErrClosed) ||
		errors.Is(err, os.ErrClosed)
}

func socketDeadline(timeout float64) time.Time {
	return time.Now().Add(time.Duration(timeout * float64(time.Second)))
}

func (handle *SocketHandle) configuredDeadlineLocked() time.Time {
	if !handle.blocking {
		// Go checks an expired deadline before attempting the underlying syscall,
		// which can report Blocked even when the descriptor is already ready. A
		// one-millisecond probe remains bounded while allowing that ready syscall
		// to make progress.
		return time.Now().Add(socketNonblockingProbeWindow)
	}
	if handle.hasTimeout {
		return socketDeadline(handle.timeout)
	}
	return time.Time{}
}

// dialer snapshots the staged blocking/timeout policy before connection setup.
//
// What: Builds one native dialer whose deadline covers TCP establishment and,
// through tls.DialWithDialer, the TLS handshake.
// Why: Applying SocketHandle deadlines only after installConn leaves connect and
// handshake able to block past Socket.setTimeout.
// How: Preserve an unlimited dial only for blocking handles without a timeout;
// zero-timeout and nonblocking handles receive an immediate deadline.
func (handle *SocketHandle) dialer() *net.Dialer {
	dialer := &net.Dialer{}
	if handle == nil {
		return dialer
	}
	handle.stateMu.Lock()
	blocking := handle.blocking
	hasTimeout := handle.hasTimeout
	timeout := handle.timeout
	handle.stateMu.Unlock()
	if !blocking || (hasTimeout && timeout <= 0) {
		dialer.Deadline = time.Now()
	} else if hasTimeout {
		dialer.Timeout = time.Duration(timeout * float64(time.Second))
	}
	return dialer
}

func (handle *SocketHandle) applyConnDeadlineLocked() error {
	if handle.conn == nil {
		return nil
	}
	return handle.conn.SetDeadline(handle.configuredDeadlineLocked())
}

func (handle *SocketHandle) applyListenerDeadlineLocked() error {
	if handle.deadlineListener == nil {
		return nil
	}
	return handle.deadlineListener.SetDeadline(handle.configuredDeadlineLocked())
}

// socketTCPConnection finds TCP control through typed connection wrappers.
//
// What: Follows NetConn links until it reaches the TCP transport.
// Why: TLS owns protocol framing but setFastSend controls the TCP socket below
// it; a direct *net.TCPConn assertion silently ignored the public request.
// How: Traverse a bounded chain so a broken self-referential wrapper cannot
// hang the runtime, returning nil when no TCP transport is exposed.
func socketTCPConnection(conn net.Conn) *net.TCPConn {
	for depth := 0; conn != nil && depth < 16; depth++ {
		if tcpConn, ok := conn.(*net.TCPConn); ok {
			return tcpConn
		}
		nested, ok := conn.(socketNestedConnection)
		if !ok {
			return nil
		}
		conn = nested.NetConn()
	}
	return nil
}

func (handle *SocketHandle) applyFastSendLocked(requireSupport bool) error {
	if handle.conn == nil {
		return nil
	}
	tcpConn := socketTCPConnection(handle.conn)
	if tcpConn == nil {
		if requireSupport {
			return errors.New("socket fast-send requires a TCP connection")
		}
		return nil
	}
	return tcpConn.SetNoDelay(handle.fastSend)
}

func (handle *SocketHandle) installConn(conn net.Conn) {
	if handle == nil || conn == nil {
		return
	}
	handle.stateMu.Lock()
	oldConn := handle.conn
	oldBound := handle.boundTCP
	oldListener := handle.listener
	handle.conn = conn
	handle.boundTCP = nil
	handle.listener = nil
	handle.deadlineListener = nil
	handle.listenerWrapper = nil
	handle.reader = bufio.NewReader(conn)
	fastErr := handle.applyFastSendLocked(false)
	deadlineErr := handle.applyConnDeadlineLocked()
	handle.stateMu.Unlock()
	if oldConn != nil && oldConn != conn {
		_ = oldConn.Close()
	}
	if oldBound != nil {
		_ = oldBound.Close()
	}
	if oldListener != nil {
		_ = oldListener.Close()
	}
	if fastErr != nil {
		socketThrow(fastErr)
	}
	if deadlineErr != nil {
		socketThrow(deadlineErr)
	}
}

func (handle *SocketHandle) installBoundTCP(bound socketBoundTCP, wrapper socketListenerWrapper) {
	if handle == nil || bound == nil {
		return
	}
	handle.stateMu.Lock()
	oldConn := handle.conn
	oldBound := handle.boundTCP
	oldListener := handle.listener
	handle.conn = nil
	handle.boundTCP = bound
	handle.listener = nil
	handle.deadlineListener = nil
	handle.listenerWrapper = wrapper
	handle.reader = nil
	handle.stateMu.Unlock()
	if oldConn != nil {
		_ = oldConn.Close()
	}
	if oldBound != nil && oldBound != bound {
		_ = oldBound.Close()
	}
	if oldListener != nil {
		_ = oldListener.Close()
	}
}

func (handle *SocketHandle) snapshotConn() net.Conn {
	if handle == nil {
		return nil
	}
	handle.stateMu.Lock()
	defer handle.stateMu.Unlock()
	return handle.conn
}

func (handle *SocketHandle) close() error {
	if handle == nil {
		return nil
	}
	handle.stateMu.Lock()
	conn := handle.conn
	bound := handle.boundTCP
	listener := handle.listener
	handle.conn = nil
	handle.boundTCP = nil
	handle.listener = nil
	handle.deadlineListener = nil
	handle.listenerWrapper = nil
	handle.reader = nil
	handle.stateMu.Unlock()

	var closeErrors []error
	if conn != nil {
		if err := conn.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
			closeErrors = append(closeErrors, err)
		}
	}
	if bound != nil {
		if err := bound.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
			closeErrors = append(closeErrors, err)
		}
	}
	if listener != nil {
		if err := listener.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
			closeErrors = append(closeErrors, err)
		}
	}
	return errors.Join(closeErrors...)
}

func socketIPv4Int(ip net.IP) int {
	v4 := ip.To4()
	if v4 == nil {
		return 0
	}
	return int(v4[0])<<24 | int(v4[1])<<16 | int(v4[2])<<8 | int(v4[3])
}

func socketIPv4(value int) net.IP {
	raw := uint32(value)
	return net.IPv4(byte(raw>>24), byte(raw>>16), byte(raw>>8), byte(raw))
}

func socketResolveIPv4(name string) (net.IP, error) {
	if parsed := net.ParseIP(name); parsed != nil {
		if v4 := parsed.To4(); v4 != nil {
			return v4, nil
		}
	}
	addresses, err := net.LookupIP(name)
	if err != nil {
		return nil, err
	}
	for _, address := range addresses {
		if v4 := address.To4(); v4 != nil {
			return v4, nil
		}
	}
	return nil, errors.New("Could not resolve host")
}

// HostResolve resolves one hostname to the network-order IPv4 Int used by Haxe.
func HostResolve(name *string) int {
	if name == nil {
		socketThrow(errors.New("Could not resolve host"))
		return 0
	}
	ip, err := socketResolveIPv4(*StdString(name))
	if err != nil {
		socketThrow(errors.New("Could not resolve host"))
		return 0
	}
	return socketIPv4Int(ip)
}

// HostToString renders a network-order IPv4 Int as a dotted quad.
func HostToString(value int) *string {
	return StringFromLiteral(socketIPv4(value).String())
}

// HostReverse performs reverse DNS for one network-order IPv4 Int.
func HostReverse(value int) *string {
	names, err := net.LookupAddr(socketIPv4(value).String())
	if err != nil || len(names) == 0 {
		socketThrow(errors.New("Could not reverse host"))
		return StringFromLiteral("")
	}
	return StringFromLiteral(strings.TrimSuffix(names[0], "."))
}

// HostLocal returns the host OS name with the portable localhost fallback.
func HostLocal() *string {
	name, err := os.Hostname()
	if err != nil || name == "" {
		name = "localhost"
	}
	return StringFromLiteral(name)
}

func socketAddress(address net.Addr) *SocketAddress {
	if address == nil {
		return &SocketAddress{}
	}
	switch typed := address.(type) {
	case *net.TCPAddr:
		return &SocketAddress{Host: socketIPv4Int(typed.IP), Port: typed.Port}
	case *net.UDPAddr:
		return &SocketAddress{Host: socketIPv4Int(typed.IP), Port: typed.Port}
	}
	host, portText, err := net.SplitHostPort(address.String())
	if err != nil {
		return &SocketAddress{}
	}
	port, _ := strconv.Atoi(portText)
	ip, resolveErr := socketResolveIPv4(host)
	if resolveErr != nil {
		return &SocketAddress{Port: port}
	}
	return &SocketAddress{Host: socketIPv4Int(ip), Port: port}
}

// SocketConnectTCP connects one typed handle to a TCP endpoint.
func SocketConnectTCP(handle *SocketHandle, host *string, port int) {
	if handle == nil || host == nil {
		socketThrow(errors.New("socket connect requires host"))
		return
	}
	conn, err := handle.dialer().Dial("tcp4", net.JoinHostPort(*StdString(host), strconv.Itoa(port)))
	if err != nil {
		socketThrow(err)
		return
	}
	handle.installConn(conn)
}

func socketBindTCP(handle *SocketHandle, host *string, port int, wrapper socketListenerWrapper) {
	if handle == nil || host == nil {
		socketThrow(errors.New("socket bind requires host"))
		return
	}
	bound, err := socketBindTCPNative(*StdString(host), port)
	if err != nil {
		socketThrow(err)
		return
	}
	handle.installBoundTCP(bound, wrapper)
}

// SocketBindTCP reserves one TCP endpoint without starting to listen.
func SocketBindTCP(handle *SocketHandle, host *string, port int) {
	socketBindTCP(handle, host, port, nil)
}

// SocketListen starts accepting connections with the requested OS backlog.
func SocketListen(handle *SocketHandle, backlog int) {
	if handle == nil {
		socketThrow(errors.New("socket listen requires a bound socket"))
		return
	}
	if backlog < 0 {
		socketThrow(errors.New("socket listen backlog cannot be negative"))
		return
	}

	handle.stateMu.Lock()
	if handle.boundTCP == nil {
		deadlineListener := handle.deadlineListener
		if handle.listener == nil || deadlineListener == nil {
			handle.stateMu.Unlock()
			socketThrow(errors.New("socket listen requires a bound socket"))
			return
		}
		err := socketRelistenTCP(deadlineListener, backlog)
		handle.stateMu.Unlock()
		if err != nil {
			socketThrow(err)
		}
		return
	}

	bound := handle.boundTCP
	tcpListener, err := bound.Listen(backlog)
	if err != nil {
		handle.stateMu.Unlock()
		socketThrow(err)
		return
	}
	listener := net.Listener(tcpListener)
	if handle.listenerWrapper != nil {
		listener = handle.listenerWrapper(listener)
	}
	handle.boundTCP = nil
	handle.listener = listener
	handle.deadlineListener = tcpListener
	handle.listenerWrapper = nil
	deadlineErr := handle.applyListenerDeadlineLocked()
	handle.stateMu.Unlock()
	if deadlineErr != nil {
		_ = listener.Close()
		socketThrow(deadlineErr)
	}
}

// SocketAccept accepts one native connection and inherits the listener policy.
func SocketAccept(handle *SocketHandle) *SocketAcceptResult {
	if handle == nil {
		return &SocketAcceptResult{Status: SocketIOEOF}
	}
	handle.stateMu.Lock()
	listener := handle.listener
	deadlineErr := handle.applyListenerDeadlineLocked()
	timeout := handle.timeout
	hasTimeout := handle.hasTimeout
	blocking := handle.blocking
	fastSend := handle.fastSend
	handle.stateMu.Unlock()
	if deadlineErr != nil {
		socketThrow(deadlineErr)
		return &SocketAcceptResult{Status: SocketIOEOF}
	}
	if listener == nil {
		return &SocketAcceptResult{Status: SocketIOEOF}
	}
	conn, err := listener.Accept()
	if err != nil {
		status := socketErrorStatus(err)
		if status >= 0 {
			return &SocketAcceptResult{Status: status}
		}
		socketThrow(err)
		return &SocketAcceptResult{Status: SocketIOEOF}
	}
	accepted := SocketNewTCP()
	accepted.timeout = timeout
	accepted.hasTimeout = hasTimeout
	accepted.blocking = blocking
	accepted.fastSend = fastSend
	accepted.installConn(conn)
	return &SocketAcceptResult{Handle: accepted, Status: SocketIOReady}
}

// SocketClose releases the current connection/listener and is idempotent.
func SocketClose(handle *SocketHandle) {
	if err := handle.close(); err != nil {
		socketThrow(err)
	}
}

// SocketShutdown performs protocol-aware half-close where it is truthful.
//
// What: Preserves TCP CloseRead/CloseWrite, sends TLS close_notify for a TLS
// write shutdown, and rejects TLS read-only shutdown explicitly.
// Why: Calling TCP methods directly cannot see through TLS, while silently
// ignoring inherited TLS controls makes successful-looking calls lie.
// How: A wrapped connection with CloseWrite receives the protocol-level write
// close; full wrapped shutdown releases the handle, while unsupported halves
// become Haxe-visible errors instead of no-ops.
func SocketShutdown(handle *SocketHandle, read bool, write bool) {
	if handle == nil || (!read && !write) {
		return
	}
	conn := handle.snapshotConn()
	if conn == nil {
		return
	}
	_, wrapped := conn.(socketNestedConnection)
	if wrapped {
		if read && write {
			SocketClose(handle)
			return
		}
		if read {
			socketThrow(errors.New("TLS read-only shutdown is unsupported"))
			return
		}
		if write {
			closeWriter, ok := conn.(socketCloseWriter)
			if !ok {
				socketThrow(errors.New("wrapped socket write shutdown is unsupported"))
				return
			}
			if err := closeWriter.CloseWrite(); err != nil && !socketErrorIsClosed(err) {
				socketThrow(err)
			}
		}
		return
	}

	if read {
		closeReader, ok := conn.(socketCloseReader)
		if !ok {
			if read && write {
				SocketClose(handle)
				return
			}
			socketThrow(errors.New("socket read shutdown is unsupported"))
			return
		}
		if err := closeReader.CloseRead(); err != nil && !socketErrorIsClosed(err) {
			socketThrow(err)
		}
	}
	if write {
		closeWriter, ok := conn.(socketCloseWriter)
		if !ok {
			if read && write {
				SocketClose(handle)
				return
			}
			socketThrow(errors.New("socket write shutdown is unsupported"))
			return
		}
		if err := closeWriter.CloseWrite(); err != nil && !socketErrorIsClosed(err) {
			socketThrow(err)
		}
	}
}

// SocketPeer returns the remote address of a connected handle.
func SocketPeer(handle *SocketHandle) *SocketAddress {
	conn := handle.snapshotConn()
	if conn == nil {
		return &SocketAddress{}
	}
	return socketAddress(conn.RemoteAddr())
}

// SocketHost returns the local connection or listener address.
func SocketHost(handle *SocketHandle) *SocketAddress {
	if handle == nil {
		return &SocketAddress{}
	}
	handle.stateMu.Lock()
	conn := handle.conn
	bound := handle.boundTCP
	listener := handle.listener
	handle.stateMu.Unlock()
	if conn != nil {
		return socketAddress(conn.LocalAddr())
	}
	if listener != nil {
		return socketAddress(listener.Addr())
	}
	if bound != nil {
		return socketAddress(bound.Addr())
	}
	return &SocketAddress{}
}

// SocketSetTimeout updates the deadline policy; a negative value clears it.
func SocketSetTimeout(handle *SocketHandle, timeout float64) {
	if handle == nil {
		return
	}
	handle.stateMu.Lock()
	if timeout < 0 {
		handle.timeout = 0
		handle.hasTimeout = false
	} else {
		handle.timeout = timeout
		handle.hasTimeout = true
	}
	connErr := handle.applyConnDeadlineLocked()
	listenerErr := handle.applyListenerDeadlineLocked()
	handle.stateMu.Unlock()
	if connErr != nil && !socketErrorIsClosed(connErr) {
		socketThrow(connErr)
	}
	if listenerErr != nil && !errors.Is(listenerErr, net.ErrClosed) {
		socketThrow(listenerErr)
	}
}

// SocketSetBlocking switches between normal deadline policy and immediate polls.
func SocketSetBlocking(handle *SocketHandle, blocking bool) {
	if handle == nil {
		return
	}
	handle.stateMu.Lock()
	handle.blocking = blocking
	connErr := handle.applyConnDeadlineLocked()
	listenerErr := handle.applyListenerDeadlineLocked()
	handle.stateMu.Unlock()
	if connErr != nil && !socketErrorIsClosed(connErr) {
		socketThrow(connErr)
	}
	if listenerErr != nil && !errors.Is(listenerErr, net.ErrClosed) {
		socketThrow(listenerErr)
	}
}

// SocketSetFastSend applies TCP_NODELAY through direct or wrapped TCP connections.
func SocketSetFastSend(handle *SocketHandle, fastSend bool) {
	if handle == nil {
		return
	}
	handle.stateMu.Lock()
	handle.fastSend = fastSend
	err := handle.applyFastSendLocked(true)
	handle.stateMu.Unlock()
	if err != nil && !errors.Is(err, net.ErrClosed) {
		socketThrow(err)
	}
}

// SocketReadValues reads up to length bytes and preserves EOF/blocked separately.
func SocketReadValues(handle *SocketHandle, length int) *SocketIOResult {
	if length <= 0 {
		return &SocketIOResult{Values: []int{}, Status: SocketIOReady}
	}
	if handle == nil {
		return &SocketIOResult{Values: []int{}, Status: SocketIOEOF}
	}
	handle.readMu.Lock()
	defer handle.readMu.Unlock()
	handle.stateMu.Lock()
	reader := handle.reader
	deadlineErr := handle.applyConnDeadlineLocked()
	handle.stateMu.Unlock()
	if deadlineErr != nil {
		if socketErrorIsClosed(deadlineErr) {
			return &SocketIOResult{Values: []int{}, Status: SocketIOEOF}
		}
		socketThrow(deadlineErr)
		return &SocketIOResult{Values: []int{}, Status: SocketIOEOF}
	}
	if reader == nil {
		return &SocketIOResult{Values: []int{}, Status: SocketIOEOF}
	}
	raw := make([]byte, length)
	count, err := reader.Read(raw)
	if count > 0 {
		values := make([]int, count)
		for index := 0; index < count; index++ {
			values[index] = int(raw[index])
		}
		return &SocketIOResult{Values: values, Count: count, Status: SocketIOReady}
	}
	status := socketErrorStatus(err)
	if status < 0 {
		socketThrow(err)
		status = SocketIOEOF
	}
	return &SocketIOResult{Values: []int{}, Status: status}
}

// SocketReadByteValue reads one byte or returns a documented negative sentinel.
func SocketReadByteValue(handle *SocketHandle) int {
	result := SocketReadValues(handle, 1)
	if result.Status == SocketIOBlocked {
		return SocketReadBlocked
	}
	if result.Status != SocketIOReady || result.Count == 0 {
		return SocketReadEOF
	}
	return result.Values[0]
}

// SocketWriteValues writes one native byte slice and reports partial progress.
func SocketWriteValues(handle *SocketHandle, values []int) *SocketIOResult {
	if len(values) == 0 {
		return &SocketIOResult{Values: []int{}, Status: SocketIOReady}
	}
	if handle == nil {
		return &SocketIOResult{Values: []int{}, Status: SocketIOEOF}
	}
	raw := make([]byte, len(values))
	for index, value := range values {
		raw[index] = byte(value)
	}
	handle.writeMu.Lock()
	defer handle.writeMu.Unlock()
	handle.stateMu.Lock()
	conn := handle.conn
	deadlineErr := handle.applyConnDeadlineLocked()
	handle.stateMu.Unlock()
	if deadlineErr != nil {
		if socketErrorIsClosed(deadlineErr) {
			return &SocketIOResult{Values: []int{}, Status: SocketIOEOF}
		}
		socketThrow(deadlineErr)
		return &SocketIOResult{Values: []int{}, Status: SocketIOEOF}
	}
	if conn == nil {
		return &SocketIOResult{Values: []int{}, Status: SocketIOEOF}
	}
	count, err := conn.Write(raw)
	if count > 0 {
		return &SocketIOResult{Count: count, Status: SocketIOReady}
	}
	status := socketErrorStatus(err)
	if status < 0 {
		socketThrow(err)
		status = SocketIOEOF
	}
	return &SocketIOResult{Status: status}
}

// SocketFlush is a typed no-op because SocketWriteValues writes directly to net.Conn.
func SocketFlush(_ *SocketHandle) {}

func socketConnectionSyscallResource(conn net.Conn) (socketSyscallResource, bool) {
	for conn != nil {
		if resource, ok := conn.(socketSyscallResource); ok {
			return resource, true
		}
		nested, ok := conn.(socketNestedConnection)
		if !ok {
			return nil, false
		}
		next := nested.NetConn()
		if next == nil || next == conn {
			return nil, false
		}
		conn = next
	}
	return nil, false
}

func socketResourceDescriptor(resource socketSyscallResource) (uintptr, error) {
	raw, err := resource.SyscallConn()
	if err != nil {
		return 0, err
	}
	var descriptor uintptr
	var duplicateErr error
	if err := raw.Control(func(value uintptr) {
		descriptor, duplicateErr = socketDuplicateDescriptor(value)
	}); err != nil {
		if duplicateErr == nil {
			_ = socketCloseDescriptor(descriptor)
		}
		return 0, err
	}
	if duplicateErr != nil {
		return 0, duplicateErr
	}
	return descriptor, nil
}

func (handle *SocketHandle) readinessSnapshot() (socketReadinessSnapshot, error) {
	if handle == nil {
		return socketReadinessSnapshot{eof: true}, nil
	}
	handle.readMu.Lock()
	handle.stateMu.Lock()
	reader := handle.reader
	conn := handle.conn
	listener := handle.listener
	deadlineListener := handle.deadlineListener
	buffered := reader != nil && reader.Buffered() > 0
	if conn == nil && listener == nil {
		handle.stateMu.Unlock()
		handle.readMu.Unlock()
		return socketReadinessSnapshot{buffered: buffered, eof: true}, nil
	}

	var resource socketSyscallResource
	if conn != nil {
		var ok bool
		resource, ok = socketConnectionSyscallResource(conn)
		if !ok {
			handle.stateMu.Unlock()
			handle.readMu.Unlock()
			return socketReadinessSnapshot{}, errors.New("socket resource does not expose native readiness")
		}
	} else {
		var ok bool
		resource, ok = deadlineListener.(socketSyscallResource)
		if !ok {
			handle.stateMu.Unlock()
			handle.readMu.Unlock()
			return socketReadinessSnapshot{}, errors.New("socket listener does not expose native readiness")
		}
	}
	descriptor, err := socketResourceDescriptor(resource)
	handle.stateMu.Unlock()
	handle.readMu.Unlock()
	if err != nil {
		if socketErrorIsClosed(err) {
			return socketReadinessSnapshot{buffered: buffered, eof: true}, nil
		}
		return socketReadinessSnapshot{}, err
	}
	return socketReadinessSnapshot{
		descriptor:    descriptor,
		hasDescriptor: true,
		buffered:      buffered,
	}, nil
}

func socketSelectEntries(
	handles []*SocketHandle,
	mode socketSelectMode,
	cache map[*SocketHandle]socketReadinessSnapshot,
) ([]socketSelectEntry, error) {
	entries := make([]socketSelectEntry, 0, len(handles))
	for index, handle := range handles {
		snapshot, ok := cache[handle]
		if !ok {
			var err error
			snapshot, err = handle.readinessSnapshot()
			if err != nil {
				return nil, err
			}
			cache[handle] = snapshot
		}
		entry := socketSelectEntry{index: index}
		switch mode {
		case socketSelectRead:
			entry.immediate = snapshot.buffered || snapshot.eof
		case socketSelectWrite, socketSelectOthers:
		default:
			return nil, errors.New("unknown socket readiness mode")
		}
		if snapshot.hasDescriptor {
			entry.descriptor = snapshot.descriptor
			entry.hasDescriptor = true
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func socketSelectDescriptors(entries []socketSelectEntry) []uintptr {
	descriptors := make([]uintptr, 0, len(entries))
	for _, entry := range entries {
		if entry.hasDescriptor {
			descriptors = append(descriptors, entry.descriptor)
		}
	}
	return descriptors
}

func socketSelectIndexes(
	entries []socketSelectEntry,
	ready map[uintptr]struct{},
	includeImmediate bool,
) []int {
	indexes := make([]int, 0, len(entries))
	for _, entry := range entries {
		_, descriptorReady := ready[entry.descriptor]
		descriptorReady = entry.hasDescriptor && descriptorReady
		if (includeImmediate && entry.immediate) || descriptorReady {
			indexes = append(indexes, entry.index)
		}
	}
	return indexes
}

func socketReleaseReadinessSnapshots(cache map[*SocketHandle]socketReadinessSnapshot) {
	for handle, snapshot := range cache {
		snapshot.release()
		cache[handle] = snapshot
	}
}

func socketSelectOnce(
	read []*SocketHandle,
	write []*SocketHandle,
	others []*SocketHandle,
	wait time.Duration,
) (*SocketSelectResult, error) {
	cache := make(map[*SocketHandle]socketReadinessSnapshot)
	defer socketReleaseReadinessSnapshots(cache)

	readEntries, err := socketSelectEntries(read, socketSelectRead, cache)
	if err != nil {
		return nil, err
	}
	writeEntries, err := socketSelectEntries(write, socketSelectWrite, cache)
	if err != nil {
		return nil, err
	}
	otherEntries, err := socketSelectEntries(others, socketSelectOthers, cache)
	if err != nil {
		return nil, err
	}
	for _, entry := range readEntries {
		if entry.immediate {
			wait = 0
			break
		}
	}

	native, err := socketSelectNative(socketNativeSelectRequest{
		Read:    socketSelectDescriptors(readEntries),
		Write:   socketSelectDescriptors(writeEntries),
		Others:  socketSelectDescriptors(otherEntries),
		Timeout: wait,
	})
	if err != nil {
		return nil, err
	}
	return &SocketSelectResult{
		Read:   socketSelectIndexes(readEntries, native.Read, true),
		Write:  socketSelectIndexes(writeEntries, native.Write, false),
		Others: socketSelectIndexes(otherEntries, native.Others, false),
	}, nil
}

// SocketSelect waits on real OS read, write, and exceptional readiness.
//
// What: Returns original caller-array indexes whose native descriptors are ready.
// Why: Connection presence is not write readiness, and resource absence is not an
// exceptional socket condition.
// How: Preserve already-buffered read bytes, poll duplicated descriptors through
// build-tagged native fd sets in bounded slices, then close each duplicate before
// resnapshotting so concurrent source close and descriptor reuse remain safe.
func SocketSelect(read []*SocketHandle, write []*SocketHandle, others []*SocketHandle, timeout float64, hasTimeout bool) *SocketSelectResult {
	const maximumWaitSlice = 10 * time.Millisecond
	finite := hasTimeout && timeout >= 0
	var deadline time.Time
	if finite {
		deadline = time.Now().Add(time.Duration(timeout * float64(time.Second)))
	}
	retriedBadDescriptor := false
	for {
		wait := maximumWaitSlice
		if finite {
			remaining := time.Until(deadline)
			if remaining <= 0 {
				wait = 0
			} else if remaining < wait {
				wait = remaining
			}
		}
		result, err := socketSelectOnce(read, write, others, wait)
		if err == nil {
			if len(result.Read) > 0 || len(result.Write) > 0 || len(result.Others) > 0 {
				return result
			}
			if finite && !time.Now().Before(deadline) {
				return result
			}
			retriedBadDescriptor = false
			continue
		}
		if errors.Is(err, syscall.EBADF) && !retriedBadDescriptor {
			retriedBadDescriptor = true
			continue
		}
		socketThrow(err)
		return &SocketSelectResult{Read: []int{}, Write: []int{}, Others: []int{}}
	}
}

// SocketWaitForRead blocks until bytes, EOF, or a read-side error is observable.
func SocketWaitForRead(handle *SocketHandle) {
	_ = SocketSelect([]*SocketHandle{handle}, nil, nil, 0, false)
}

func (handle *SocketHandle) udpConn(create bool) *net.UDPConn {
	if handle == nil {
		return nil
	}
	handle.stateMu.Lock()
	conn := handle.conn
	if conn != nil {
		handle.stateMu.Unlock()
		udpConn, ok := conn.(*net.UDPConn)
		if !ok {
			socketThrow(errors.New("udp socket expects UDP connection"))
			return nil
		}
		return udpConn
	}
	if !create {
		handle.stateMu.Unlock()
		return nil
	}
	// Keep lazy creation under the state lock so concurrent first sends all
	// observe the same connection instead of installing and closing one another's
	// ephemeral sockets.
	udpConn, err := net.ListenUDP("udp4", nil)
	if err != nil {
		handle.stateMu.Unlock()
		socketThrow(err)
		return nil
	}
	oldBound := handle.boundTCP
	handle.conn = udpConn
	handle.boundTCP = nil
	handle.listener = nil
	handle.deadlineListener = nil
	handle.listenerWrapper = nil
	handle.reader = bufio.NewReader(udpConn)
	deadlineErr := handle.applyConnDeadlineLocked()
	handle.stateMu.Unlock()
	if oldBound != nil {
		_ = oldBound.Close()
	}
	if deadlineErr != nil {
		handle.stateMu.Lock()
		if handle.conn == udpConn {
			handle.conn = nil
			handle.reader = nil
		}
		handle.stateMu.Unlock()
		_ = udpConn.Close()
		socketThrow(deadlineErr)
		return nil
	}
	if err := handle.applyBroadcast(); err != nil {
		socketThrow(err)
	}
	return udpConn
}

func (handle *SocketHandle) applyBroadcast() error {
	udpConn := handle.udpConn(false)
	if udpConn == nil {
		return nil
	}
	handle.stateMu.Lock()
	enabled := handle.broadcast
	handle.stateMu.Unlock()
	rawConn, err := udpConn.SyscallConn()
	if err != nil {
		return err
	}
	value := 0
	if enabled {
		value = 1
	}
	var optionErr error
	controlErr := rawConn.Control(func(fileDescriptor uintptr) {
		optionErr = socketSetBroadcast(fileDescriptor, value)
	})
	if controlErr != nil {
		return controlErr
	}
	return optionErr
}

// SocketUdpBind binds one typed UDP handle to an IPv4 endpoint.
func SocketUdpBind(handle *SocketHandle, host *string, port int) {
	if handle == nil || host == nil {
		socketThrow(errors.New("udp bind requires host"))
		return
	}
	address, err := net.ResolveUDPAddr("udp4", net.JoinHostPort(*StdString(host), strconv.Itoa(port)))
	if err != nil {
		socketThrow(err)
		return
	}
	conn, err := net.ListenUDP("udp4", address)
	if err != nil {
		socketThrow(err)
		return
	}
	handle.installConn(conn)
	if err := handle.applyBroadcast(); err != nil {
		socketThrow(err)
	}
}

// SocketUdpSetBroadcast installs SO_BROADCAST on the current or lazily created UDP socket.
func SocketUdpSetBroadcast(handle *SocketHandle, enabled bool) {
	if handle == nil {
		socketThrow(errors.New("udp socket is nil"))
		return
	}
	handle.stateMu.Lock()
	handle.broadcast = enabled
	handle.stateMu.Unlock()
	if handle.udpConn(true) == nil {
		return
	}
	if err := handle.applyBroadcast(); err != nil {
		socketThrow(err)
	}
}

// SocketUdpSendTo writes one complete datagram to a network-order IPv4 peer.
func SocketUdpSendTo(handle *SocketHandle, values []int, host int, port int) *SocketIOResult {
	conn := handle.udpConn(true)
	if conn == nil {
		return &SocketIOResult{Status: SocketIOEOF}
	}
	raw := make([]byte, len(values))
	for index, value := range values {
		raw[index] = byte(value)
	}
	handle.writeMu.Lock()
	defer handle.writeMu.Unlock()
	handle.stateMu.Lock()
	deadlineErr := handle.applyConnDeadlineLocked()
	handle.stateMu.Unlock()
	if deadlineErr != nil {
		if socketErrorIsClosed(deadlineErr) {
			return &SocketIOResult{Status: SocketIOEOF}
		}
		socketThrow(deadlineErr)
		return &SocketIOResult{Status: SocketIOEOF}
	}
	count, err := conn.WriteToUDP(raw, &net.UDPAddr{IP: socketIPv4(host), Port: port})
	if count > 0 {
		return &SocketIOResult{Count: count, Status: SocketIOReady}
	}
	status := socketErrorStatus(err)
	if status < 0 {
		socketThrow(err)
		status = SocketIOEOF
	}
	return &SocketIOResult{Status: status}
}

// SocketUdpReadFrom reads one datagram and returns its typed peer address.
//
// What: Preserves the sender for both payload-bearing and zero-byte datagrams.
// Why: UDP permits an empty datagram, so count == 0 does not mean that no
// packet arrived or that its source address may be discarded.
// How: Treat a nil read error as the success authority, copy however many
// payload bytes were returned, and always populate the non-nil UDP peer.
func SocketUdpReadFrom(handle *SocketHandle, length int) *SocketDatagramResult {
	if length <= 0 {
		return &SocketDatagramResult{Values: []int{}, Status: SocketIOReady}
	}
	conn := handle.udpConn(false)
	if conn == nil {
		return &SocketDatagramResult{Values: []int{}, Status: SocketIOEOF}
	}
	handle.readMu.Lock()
	defer handle.readMu.Unlock()
	handle.stateMu.Lock()
	deadlineErr := handle.applyConnDeadlineLocked()
	handle.stateMu.Unlock()
	if deadlineErr != nil {
		if socketErrorIsClosed(deadlineErr) {
			return &SocketDatagramResult{Values: []int{}, Status: SocketIOEOF}
		}
		socketThrow(deadlineErr)
		return &SocketDatagramResult{Values: []int{}, Status: SocketIOEOF}
	}
	raw := make([]byte, length)
	count, remote, err := conn.ReadFromUDP(raw)
	if err == nil {
		if remote == nil {
			socketThrow(errors.New("udp read completed without a peer address"))
			return &SocketDatagramResult{Values: []int{}, Status: SocketIOEOF}
		}
		values := make([]int, count)
		for index := 0; index < count; index++ {
			values[index] = int(raw[index])
		}
		return &SocketDatagramResult{
			Values: values,
			Count:  count,
			Status: SocketIOReady,
			Host:   socketIPv4Int(remote.IP),
			Port:   remote.Port,
		}
	}
	status := socketErrorStatus(err)
	if status < 0 {
		socketThrow(err)
		status = SocketIOEOF
	}
	return &SocketDatagramResult{Values: []int{}, Status: status}
}
