package main

import (
	"bufio"
	"net"
	"os"
	"snapshot/hxrt"
	"strconv"
	"strings"
	"time"
)

func main() {
	loop := New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1"))
	_ = loop
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("loop.to_nonempty="), hxrt.StdString(!hxrt.StringEqualAny(loop.toString(), hxrt.StringFromLiteral("")))))
	named := New_sys__net__Host(hxrt.StringFromLiteral("localhost"))
	_ = named
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("named.to_nonempty="), hxrt.StdString(!hxrt.StringEqualAny(named.toString(), hxrt.StringFromLiteral("")))))
	invalidThrows := false
	_ = invalidThrows
	hxrt.TryCatch(func() {
		New_sys__net__Host(hxrt.StringFromLiteral("256.256.256.256"))
	}, func(hx_caught_1 any) {
		hx_tmp := hx_caught_1
		_ = hx_tmp
		invalidThrows = true
	})
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("invalid_throws="), hxrt.StdString(invalidThrows)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("localhost_nonempty="), hxrt.StdString(!hxrt.StringEqualAny(sys__net__Host_localhost(), hxrt.StringFromLiteral("")))))
	hxrt.TryCatch(func() {
		loop.reverse()
	}, func(hx_caught_3 any) {
		hx_tmp_1 := hx_caught_3
		_ = hx_tmp_1
	})
	hxrt.Println(hxrt.StringFromLiteral("reverse_called=true"))
}

type sys__net__Host struct {
	host     *string
	ip       int
	resolved *string
}

func hxrt__host_empty() *sys__net__Host {
	return &sys__net__Host{host: hxrt.StringFromLiteral(""), ip: 0, resolved: hxrt.StringFromLiteral("")}
}

func New_sys__net__Host(name *string) *sys__net__Host {
	if name == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Could not resolve host"))
		return hxrt__host_empty()
	}
	rawName := *hxrt.StdString(name)
	ips, err := net.LookupIP(rawName)
	if err != nil || len(ips) == 0 {
		hxrt.Throw(hxrt.StringFromLiteral("Could not resolve host"))
		return hxrt__host_empty()
	}
	selected := ips[0]
	for _, candidate := range ips {
		if v4 := candidate.To4(); v4 != nil {
			selected = v4
			break
		}
	}
	resolved := hxrt.StringFromLiteral(selected.String())
	return &sys__net__Host{host: name, ip: 0, resolved: resolved}
}

func (self *sys__net__Host) toString() *string {
	if self == nil || self.resolved == nil {
		return hxrt.StringFromLiteral("")
	}
	return self.resolved
}

func (self *sys__net__Host) reverse() *string {
	if self == nil || self.resolved == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Could not reverse host"))
		return hxrt.StringFromLiteral("")
	}
	names, err := net.LookupAddr(*hxrt.StdString(self.resolved))
	if err != nil || len(names) == 0 {
		hxrt.Throw(hxrt.StringFromLiteral("Could not reverse host"))
		return hxrt.StringFromLiteral("")
	}
	resolved := strings.TrimSuffix(names[0], ".")
	return hxrt.StringFromLiteral(resolved)
}

func sys__net__Host_localhost() *string {
	name, err := os.Hostname()
	if err != nil || name == "" {
		return hxrt.StringFromLiteral("localhost")
	}
	return hxrt.StringFromLiteral(name)
}

type sys__net__SocketInput struct {
	reader *bufio.Reader
	socket *sys__net__Socket
}

type sys__net__SocketOutput struct {
	writer *bufio.Writer
	socket *sys__net__Socket
}

type sys__net__Socket struct {
	input      *sys__net__SocketInput
	output     *sys__net__SocketOutput
	custom     any
	conn       net.Conn
	listener   net.Listener
	timeout    float64
	hasTimeout bool
	blocking   bool
	fastSend   bool
}

func New_sys__net__Socket() *sys__net__Socket {
	return &sys__net__Socket{input: &sys__net__SocketInput{}, output: &sys__net__SocketOutput{}, blocking: true}
}

func hxrt__socket_deadline(timeout float64) time.Time {
	duration := time.Duration(timeout * float64(time.Second))
	return time.Now().Add(duration)
}

func (self *sys__net__Socket) hxrt__socket_applyConnDeadline() {
	if self == nil || self.conn == nil {
		return
	}
	if !self.blocking {
		_ = self.conn.SetDeadline(time.Now())
		return
	}
	if self.hasTimeout {
		_ = self.conn.SetDeadline(hxrt__socket_deadline(self.timeout))
		return
	}
	_ = self.conn.SetDeadline(time.Time{})
}

func (self *sys__net__Socket) hxrt__socket_applyListenerDeadline() {
	if self == nil || self.listener == nil {
		return
	}
	tcpListener, ok := self.listener.(*net.TCPListener)
	if !ok {
		return
	}
	if !self.blocking {
		_ = tcpListener.SetDeadline(time.Now())
		return
	}
	if self.hasTimeout {
		_ = tcpListener.SetDeadline(hxrt__socket_deadline(self.timeout))
		return
	}
	_ = tcpListener.SetDeadline(time.Time{})
}

func (self *sys__net__Socket) hxrt__socket_applyFastSend() {
	if self == nil || self.conn == nil {
		return
	}
	tcpConn, ok := self.conn.(*net.TCPConn)
	if !ok {
		return
	}
	if err := tcpConn.SetNoDelay(self.fastSend); err != nil {
		hxrt.Throw(err)
	}
}

func (self *sys__net__Socket) hxrt__socket_setConn(conn net.Conn) {
	if self == nil || conn == nil {
		return
	}
	self.conn = conn
	self.input = &sys__net__SocketInput{reader: bufio.NewReader(conn), socket: self}
	self.output = &sys__net__SocketOutput{writer: bufio.NewWriter(conn), socket: self}
	self.hxrt__socket_applyFastSend()
	self.hxrt__socket_applyConnDeadline()
}

func (self *sys__net__Socket) close() {
	if self == nil {
		return
	}
	if self.conn != nil {
		_ = self.conn.Close()
		self.conn = nil
	}
	if self.listener != nil {
		_ = self.listener.Close()
		self.listener = nil
	}
}

func (self *sys__net__Socket) connect(host *sys__net__Host, port int) {
	if self == nil || host == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket connect requires host"))
		return
	}
	resolvedHost := host.toString()
	if resolvedHost == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket connect requires host"))
		return
	}
	address := net.JoinHostPort(*hxrt.StdString(resolvedHost), strconv.Itoa(port))
	conn, err := net.Dial("tcp", address)
	if err != nil {
		hxrt.Throw(err)
		return
	}
	self.hxrt__socket_setConn(conn)
}

func (self *sys__net__Socket) bind(host *sys__net__Host, port int) {
	if self == nil || host == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket bind requires host"))
		return
	}
	resolvedHost := host.toString()
	if resolvedHost == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket bind requires host"))
		return
	}
	address := net.JoinHostPort(*hxrt.StdString(resolvedHost), strconv.Itoa(port))
	listener, err := net.Listen("tcp", address)
	if err != nil {
		hxrt.Throw(err)
		return
	}
	if self.listener != nil {
		_ = self.listener.Close()
	}
	self.listener = listener
	self.hxrt__socket_applyListenerDeadline()
}

func (self *sys__net__Socket) listen(connections int) {
	_ = connections
}

func (self *sys__net__Socket) accept() *sys__net__Socket {
	if self == nil || self.listener == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket accept requires listener"))
		return New_sys__net__Socket()
	}
	self.hxrt__socket_applyListenerDeadline()
	conn, err := self.listener.Accept()
	if err != nil {
		hxrt.Throw(err)
		return New_sys__net__Socket()
	}
	accepted := New_sys__net__Socket()
	accepted.timeout = self.timeout
	accepted.hasTimeout = self.hasTimeout
	accepted.blocking = self.blocking
	accepted.fastSend = self.fastSend
	accepted.hxrt__socket_setConn(conn)
	return accepted
}

func (self *sys__net__Socket) read() *string {
	if self == nil || self.input == nil {
		return hxrt.StringFromLiteral("")
	}
	return self.input.readLine()
}

func (self *sys__net__Socket) write(content *string) {
	if self == nil || self.output == nil {
		return
	}
	self.output.writeString(content)
	self.output.flush()
}

func (self *sys__net__Socket) shutdown(read bool, write bool) {
	if self == nil || self.conn == nil || (!read && !write) {
		return
	}
	if tcpConn, ok := self.conn.(*net.TCPConn); ok {
		if read {
			if err := tcpConn.CloseRead(); err != nil {
				hxrt.Throw(err)
			}
		}
		if write {
			if err := tcpConn.CloseWrite(); err != nil {
				hxrt.Throw(err)
			}
		}
		return
	}
	if err := self.conn.Close(); err != nil {
		hxrt.Throw(err)
	}
	self.conn = nil
}

func hxrt__socket_addrInfo(addr net.Addr) map[string]any {
	if addr == nil {
		return map[string]any{"host": hxrt__host_empty(), "port": 0}
	}
	rawHost := ""
	rawPort := "0"
	hostPart, portPart, err := net.SplitHostPort(addr.String())
	if err == nil {
		rawHost = hostPart
		rawPort = portPart
	}
	port, _ := strconv.Atoi(rawPort)
	if rawHost == "" {
		return map[string]any{"host": hxrt__host_empty(), "port": port}
	}
	return map[string]any{"host": New_sys__net__Host(hxrt.StringFromLiteral(rawHost)), "port": port}
}

func (self *sys__net__Socket) peer() map[string]any {
	if self == nil || self.conn == nil {
		return map[string]any{"host": hxrt__host_empty(), "port": 0}
	}
	return hxrt__socket_addrInfo(self.conn.RemoteAddr())
}

func (self *sys__net__Socket) host() map[string]any {
	if self == nil {
		return map[string]any{"host": hxrt__host_empty(), "port": 0}
	}
	if self.conn != nil {
		return hxrt__socket_addrInfo(self.conn.LocalAddr())
	}
	if self.listener != nil {
		return hxrt__socket_addrInfo(self.listener.Addr())
	}
	return map[string]any{"host": hxrt__host_empty(), "port": 0}
}

func (self *sys__net__Socket) setTimeout(timeout float64) {
	if self == nil {
		return
	}
	if timeout < 0 {
		self.hasTimeout = false
		self.timeout = 0
	} else {
		self.hasTimeout = true
		self.timeout = timeout
	}
	self.hxrt__socket_applyConnDeadline()
	self.hxrt__socket_applyListenerDeadline()
}

func (self *sys__net__Socket) waitForRead() {
	if self == nil {
		return
	}
	_ = sys__net__Socket_select_([]*sys__net__Socket{self}, []*sys__net__Socket{}, []*sys__net__Socket{}, -1)
}

func (self *sys__net__Socket) setBlocking(b bool) {
	if self == nil {
		return
	}
	self.blocking = b
	self.hxrt__socket_applyConnDeadline()
	self.hxrt__socket_applyListenerDeadline()
}

func (self *sys__net__Socket) setFastSend(b bool) {
	if self == nil {
		return
	}
	self.fastSend = b
	self.hxrt__socket_applyFastSend()
}

func sys__net__Socket_select_(read []*sys__net__Socket, write []*sys__net__Socket, others []*sys__net__Socket, timeout ...float64) map[string]any {
	if read == nil {
		read = []*sys__net__Socket{}
	}
	if write == nil {
		write = []*sys__net__Socket{}
	}
	if others == nil {
		others = []*sys__net__Socket{}
	}
	effectiveTimeout := -1.0
	if len(timeout) > 0 {
		effectiveTimeout = timeout[0]
	}
	readyRead := make([]*sys__net__Socket, 0, len(read))
	readyWrite := make([]*sys__net__Socket, 0, len(write))
	readyOther := make([]*sys__net__Socket, 0, len(others))
	for _, socket := range read {
		if socket == nil || socket.conn == nil || socket.input == nil || socket.input.reader == nil {
			continue
		}
		reader := socket.input.reader
		if reader.Buffered() > 0 {
			readyRead = append(readyRead, socket)
			continue
		}
		if effectiveTimeout >= 0 {
			deadline := time.Now()
			if effectiveTimeout > 0 {
				deadline = time.Now().Add(time.Duration(effectiveTimeout * float64(time.Second)))
			}
			_ = socket.conn.SetReadDeadline(deadline)
		}
		_, err := reader.Peek(1)
		socket.hxrt__socket_applyConnDeadline()
		if err == nil {
			readyRead = append(readyRead, socket)
			continue
		}
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			continue
		}
		readyOther = append(readyOther, socket)
	}
	for _, socket := range write {
		if socket == nil || socket.conn == nil {
			continue
		}
		readyWrite = append(readyWrite, socket)
	}
	for _, socket := range others {
		if socket == nil {
			continue
		}
		readyOther = append(readyOther, socket)
	}
	return map[string]any{"read": readyRead, "write": readyWrite, "others": readyOther}
}

func (self *sys__net__SocketInput) readLine() *string {
	if self == nil || self.reader == nil {
		return hxrt.StringFromLiteral("")
	}
	if self.socket != nil {
		self.socket.hxrt__socket_applyConnDeadline()
	}
	line, err := self.reader.ReadString('\n')
	if err != nil && len(line) == 0 {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(strings.TrimRight(line, "\r\n"))
}

func (self *sys__net__SocketOutput) writeString(value *string) {
	if self == nil || self.writer == nil || value == nil {
		return
	}
	if self.socket != nil {
		self.socket.hxrt__socket_applyConnDeadline()
	}
	if _, err := self.writer.WriteString(*hxrt.StdString(value)); err != nil {
		hxrt.Throw(err)
	}
}

func (self *sys__net__SocketOutput) flush() {
	if self == nil || self.writer == nil {
		return
	}
	if self.socket != nil {
		self.socket.hxrt__socket_applyConnDeadline()
	}
	if err := self.writer.Flush(); err != nil {
		hxrt.Throw(err)
	}
}
