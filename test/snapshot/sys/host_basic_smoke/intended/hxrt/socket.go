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

// SocketHandle owns one native socket resource behind a typed opaque boundary.
//
// What: Stores a TCP/UDP connection or listener plus buffering and deadline policy.
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
	listener         net.Listener
	deadlineListener socketDeadlineListener
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
	if errors.Is(err, io.EOF) || errors.Is(err, io.ErrClosedPipe) || errors.Is(err, net.ErrClosed) || errors.Is(err, os.ErrClosed) {
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

func socketDeadline(timeout float64) time.Time {
	return time.Now().Add(time.Duration(timeout * float64(time.Second)))
}

func (handle *SocketHandle) configuredDeadlineLocked() time.Time {
	if !handle.blocking {
		return time.Now()
	}
	if handle.hasTimeout {
		return socketDeadline(handle.timeout)
	}
	return time.Time{}
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

func (handle *SocketHandle) applyFastSendLocked() error {
	if handle.conn == nil {
		return nil
	}
	tcpConn, ok := handle.conn.(*net.TCPConn)
	if !ok {
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
	oldListener := handle.listener
	handle.conn = conn
	handle.listener = nil
	handle.deadlineListener = nil
	handle.reader = bufio.NewReader(conn)
	fastErr := handle.applyFastSendLocked()
	deadlineErr := handle.applyConnDeadlineLocked()
	handle.stateMu.Unlock()
	if oldConn != nil && oldConn != conn {
		_ = oldConn.Close()
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

func (handle *SocketHandle) installListener(listener net.Listener, deadlineListener socketDeadlineListener) {
	if handle == nil || listener == nil {
		return
	}
	handle.stateMu.Lock()
	oldConn := handle.conn
	oldListener := handle.listener
	handle.conn = nil
	handle.reader = nil
	handle.listener = listener
	handle.deadlineListener = deadlineListener
	deadlineErr := handle.applyListenerDeadlineLocked()
	handle.stateMu.Unlock()
	if oldConn != nil {
		_ = oldConn.Close()
	}
	if oldListener != nil && oldListener != listener {
		_ = oldListener.Close()
	}
	if deadlineErr != nil {
		socketThrow(deadlineErr)
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
	listener := handle.listener
	handle.conn = nil
	handle.listener = nil
	handle.deadlineListener = nil
	handle.reader = nil
	handle.stateMu.Unlock()

	var closeErrors []error
	if conn != nil {
		if err := conn.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
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
	conn, err := net.Dial("tcp4", net.JoinHostPort(*StdString(host), strconv.Itoa(port)))
	if err != nil {
		socketThrow(err)
		return
	}
	handle.installConn(conn)
}

// SocketBindTCP creates the listener immediately; SocketListen retains the
// upstream backlog-shaped API but Go owns the listen transition atomically.
func SocketBindTCP(handle *SocketHandle, host *string, port int) {
	if handle == nil || host == nil {
		socketThrow(errors.New("socket bind requires host"))
		return
	}
	listener, err := net.Listen("tcp4", net.JoinHostPort(*StdString(host), strconv.Itoa(port)))
	if err != nil {
		socketThrow(err)
		return
	}
	deadlineListener, _ := listener.(socketDeadlineListener)
	handle.installListener(listener, deadlineListener)
}

// SocketListen preserves the Haxe API after BindTCP has created Go's listener.
func SocketListen(_ *SocketHandle, _ int) {}

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

// SocketShutdown performs half-close when the connection supports it.
func SocketShutdown(handle *SocketHandle, read bool, write bool) {
	if handle == nil || (!read && !write) {
		return
	}
	conn := handle.snapshotConn()
	if conn == nil {
		return
	}
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		if read {
			if err := tcpConn.CloseRead(); err != nil && !errors.Is(err, net.ErrClosed) {
				socketThrow(err)
			}
		}
		if write {
			if err := tcpConn.CloseWrite(); err != nil && !errors.Is(err, net.ErrClosed) {
				socketThrow(err)
			}
		}
		return
	}
	if read && write {
		SocketClose(handle)
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
	listener := handle.listener
	handle.stateMu.Unlock()
	if conn != nil {
		return socketAddress(conn.LocalAddr())
	}
	if listener != nil {
		return socketAddress(listener.Addr())
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
	if connErr != nil && !errors.Is(connErr, net.ErrClosed) {
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
	if connErr != nil && !errors.Is(connErr, net.ErrClosed) {
		socketThrow(connErr)
	}
	if listenerErr != nil && !errors.Is(listenerErr, net.ErrClosed) {
		socketThrow(listenerErr)
	}
}

// SocketSetFastSend applies TCP_NODELAY where the connection exposes TCP control.
func SocketSetFastSend(handle *SocketHandle, fastSend bool) {
	if handle == nil {
		return
	}
	handle.stateMu.Lock()
	handle.fastSend = fastSend
	err := handle.applyFastSendLocked()
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
	if deadlineErr != nil && !errors.Is(deadlineErr, net.ErrClosed) {
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
	if deadlineErr != nil && !errors.Is(deadlineErr, net.ErrClosed) {
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

func (handle *SocketHandle) pollRead() (ready bool, exceptional bool) {
	if handle == nil {
		return false, true
	}
	handle.readMu.Lock()
	defer handle.readMu.Unlock()
	handle.stateMu.Lock()
	reader := handle.reader
	conn := handle.conn
	if reader == nil || conn == nil {
		handle.stateMu.Unlock()
		return false, true
	}
	if reader.Buffered() > 0 {
		handle.stateMu.Unlock()
		return true, false
	}
	// A deadline a fraction in the future lets an already-scheduled packet win
	// the race with the timer. An exactly-now deadline can report timeout even
	// when bytes are already queued on implementations such as net.Pipe.
	_ = conn.SetReadDeadline(time.Now().Add(time.Millisecond))
	handle.stateMu.Unlock()
	_, err := reader.Peek(1)
	handle.stateMu.Lock()
	_ = handle.applyConnDeadlineLocked()
	handle.stateMu.Unlock()
	if err == nil {
		return true, false
	}
	if socketErrorStatus(err) == SocketIOBlocked {
		return false, false
	}
	if socketErrorStatus(err) == SocketIOEOF {
		return true, false
	}
	return false, true
}

func socketConnected(handle *SocketHandle) bool {
	if handle == nil {
		return false
	}
	handle.stateMu.Lock()
	defer handle.stateMu.Unlock()
	return handle.conn != nil
}

// SocketSelect polls typed handles and returns indexes into the caller's arrays.
func SocketSelect(read []*SocketHandle, write []*SocketHandle, others []*SocketHandle, timeout float64, hasTimeout bool) *SocketSelectResult {
	started := time.Now()
	for {
		result := &SocketSelectResult{Read: []int{}, Write: []int{}, Others: []int{}}
		for index, handle := range read {
			ready, exceptional := handle.pollRead()
			// A closed or otherwise exceptional stream is readable as EOF from the
			// Haxe caller's perspective. Returning it also prevents waitForRead from
			// blocking forever after a concurrent or prior close.
			if ready || exceptional {
				result.Read = append(result.Read, index)
			}
		}
		for index, handle := range write {
			if socketConnected(handle) {
				result.Write = append(result.Write, index)
			}
		}
		for index, handle := range others {
			if !socketConnected(handle) {
				result.Others = append(result.Others, index)
			}
		}
		if len(result.Read) > 0 || len(result.Write) > 0 || len(result.Others) > 0 {
			return result
		}
		if hasTimeout && (timeout <= 0 || time.Since(started) >= time.Duration(timeout*float64(time.Second))) {
			return result
		}
		time.Sleep(time.Millisecond)
	}
}

// SocketWaitForRead blocks until the handle becomes readable or exceptional.
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
	handle.conn = udpConn
	handle.listener = nil
	handle.deadlineListener = nil
	handle.reader = bufio.NewReader(udpConn)
	deadlineErr := handle.applyConnDeadlineLocked()
	handle.stateMu.Unlock()
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
	if deadlineErr != nil && !errors.Is(deadlineErr, net.ErrClosed) {
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
	if deadlineErr != nil && !errors.Is(deadlineErr, net.ErrClosed) {
		socketThrow(deadlineErr)
		return &SocketDatagramResult{Values: []int{}, Status: SocketIOEOF}
	}
	raw := make([]byte, length)
	count, remote, err := conn.ReadFromUDP(raw)
	if count > 0 {
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
