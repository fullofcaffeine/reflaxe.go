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
	server := New_sys__net__UdpSocket()
	client := New_sys__net__UdpSocket()
	var failure any = nil
	hxrt.TryCatch(func() {
		server.bind(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), 0)
		bound := server.host()
		if (bound == nil) || (func(hx_obj_3 map[string]any) int {
			hx_field_4 := hx_obj_3["port"]
			if hx_field_4 == nil {
				var hx_zero_5 int
				return hx_zero_5
			}
			return hx_field_4.(int)
		}(bound) <= 0) {
			hxrt.Throw(hxrt.StringFromLiteral("missing bound udp port"))
		}
		server.setBlocking(true)
		client.setBroadcast(false)
		sent := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("udp-ping"))
		target := New_sys__net__Address()
		target.host = New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")).ip
		target.port = func(hx_obj_6 map[string]any) int {
			hx_field_7 := hx_obj_6["port"]
			if hx_field_7 == nil {
				var hx_zero_8 int
				return hx_zero_8
			}
			return hx_field_7.(int)
		}(bound)
		wrote := client.sendTo(sent, 0, sent.length, target)
		recv := haxe__io__Bytes_alloc(32)
		remote := New_sys__net__Address()
		read := server.readFrom(recv, 0, recv.length, remote)
		hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bound.host="), func(hx_obj_9 map[string]any) *sys__net__Host {
			hx_field_10 := hx_obj_9["host"]
			if hx_field_10 == nil {
				var hx_zero_11 *sys__net__Host
				return hx_zero_11
			}
			return hx_field_10.(*sys__net__Host)
		}(bound).toString()))
		hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bound.port.positive="), hxrt.StdString((func(hx_obj_12 map[string]any) int {
			hx_field_13 := hx_obj_12["port"]
			if hx_field_13 == nil {
				var hx_zero_14 int
				return hx_zero_14
			}
			return hx_field_13.(int)
		}(bound) > 0))))
		hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("wrote="), wrote))
		hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("read="), read))
		hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("payload="), recv.sub(0, read).toString()))
		hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remote.port.positive="), hxrt.StdString((remote.port > 0))))
		hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remote.host="), remote.getHost().toString()))
	}, func(hx_caught_1 any) {
		error := hx_caught_1
		failure = error
	})
	safeClose(client)
	safeClose(server)
	if !hxrt.AnyEqualsNull(failure) {
		hxrt.Throw(failure)
	}
}

func safeClose(socket *sys__net__UdpSocket) {
	if socket == nil {
		return
	}
	hxrt.TryCatch(func() {
		socket.close()
	}, func(hx_caught_15 any) {
		hx_tmp := hx_caught_15
		_ = hx_tmp
	})
}

type haxe__io__Encoding struct {
	tag int
}

type haxe__io__Input interface {
	get_bigEndian() bool
	set_bigEndian(e bool) bool
	readByte() int
	readBytes(buf *haxe__io__Bytes, pos int, len int) int
	close()
}

type haxe__io__Output interface {
	get_bigEndian() bool
	set_bigEndian(e bool) bool
	writeByte(c int)
	writeBytes(s *haxe__io__Bytes, pos int, len int) int
	flush()
	close()
}

type haxe__io__Eof struct {
}

type haxe__io__Error struct {
	tag    int
	params []any
}

type haxe__io__Bytes struct {
	b             []int
	length        int
	__hx_raw      []byte
	__hx_rawValid bool
}

type haxe__io__BytesBuffer struct {
	b []int
}

type haxe__io__BytesInput struct {
	bigEndian bool
	b         []int
	pos       int
	len       int
	totlen    int
}

type haxe__io__BytesOutput struct {
	bigEndian bool
	b         *haxe__io__BytesBuffer
}

func New_haxe__io__Input() haxe__io__Input {
	return New_haxe__io__BytesInput(&haxe__io__Bytes{b: []int{}, length: 0})
}

func New_haxe__io__Output() haxe__io__Output {
	return New_haxe__io__BytesOutput()
}

func New_haxe__io__Eof() *haxe__io__Eof {
	return &haxe__io__Eof{}
}

var haxe__io__Encoding_UTF8 *haxe__io__Encoding = &haxe__io__Encoding{tag: 0}

var haxe__io__Encoding_RawNative *haxe__io__Encoding = &haxe__io__Encoding{tag: 1}

func (self *haxe__io__Encoding) String() string {
	if self == nil {
		return "null"
	}
	switch self.tag {
	case 0:
		return "UTF8"
	case 1:
		return "RawNative"
	default:
		return "Encoding"
	}
}

func (self *haxe__io__Encoding) toString() *string {
	return hxrt.StringFromLiteral(self.String())
}

func haxe__io__resolveEncodingTag(encoding ...*haxe__io__Encoding) int {
	if len(encoding) == 0 || encoding[0] == nil {
		return 0
	}
	return encoding[0].tag
}

func haxe__io__bytes_fromStringRawNativeUTF16LE(value *string) *haxe__io__Bytes {
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func haxe__io__bytes_toStringRawNativeUTF16LE(value []int) *string {
	return hxrt.BytesToString(value)
}

func (self *haxe__io__Eof) toString() *string {
	return hxrt.StringFromLiteral("Eof")
}

var haxe__io__Error_Blocked *haxe__io__Error = &haxe__io__Error{tag: 0}

var haxe__io__Error_Overflow *haxe__io__Error = &haxe__io__Error{tag: 1}

var haxe__io__Error_OutsideBounds *haxe__io__Error = &haxe__io__Error{tag: 2}

func haxe__io__Error_Custom(e any) *haxe__io__Error {
	return &haxe__io__Error{tag: 3, params: []any{e}}
}

func (self *haxe__io__Error) String() string {
	if self == nil {
		return "null"
	}
	switch self.tag {
	case 0:
		return "Blocked"
	case 1:
		return "Overflow"
	case 2:
		return "OutsideBounds"
	case 3:
		if len(self.params) == 0 {
			return "Custom(null)"
		}
		return "Custom(" + *hxrt.StdString(self.params[0]) + ")"
	default:
		return "Error"
	}
}

func (self *haxe__io__Error) toString() *string {
	return hxrt.StringFromLiteral(self.String())
}

func New_haxe__io__Bytes(length int, b []int) *haxe__io__Bytes {
	if b == nil {
		b = make([]int, length)
	}
	return &haxe__io__Bytes{b: b, length: len(b)}
}

func haxe__io__Bytes_alloc(length int) *haxe__io__Bytes {
	return &haxe__io__Bytes{b: make([]int, length), length: length}
}

func haxe__io__Bytes_ofString(value *string, encoding ...*haxe__io__Encoding) *haxe__io__Bytes {
	if haxe__io__resolveEncodingTag(encoding...) == 1 {
		return haxe__io__bytes_fromStringRawNativeUTF16LE(value)
	}
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func haxe__io__Bytes_ofData(b []int) *haxe__io__Bytes {
	if b == nil {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return &haxe__io__Bytes{b: b, length: len(b)}
}

func haxe__io__Bytes_ofHex(s *string) *haxe__io__Bytes {
	decoded := hxrt.BytesOfHex(s)
	return &haxe__io__Bytes{b: decoded, length: len(decoded)}
}

func (self *haxe__io__Bytes) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToString(self.b)
}

func (self *haxe__io__Bytes) toHex() *string {
	if self == nil || self.length == 0 {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToHex(self.b, self.length)
}

func (self *haxe__io__Bytes) getData() []int {
	if self == nil {
		return []int{}
	}
	return self.b
}

func (self *haxe__io__Bytes) getString(pos int, len int, encoding ...*haxe__io__Encoding) *string {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return hxrt.StringFromLiteral("")
	}
	slice := self.b[pos : pos+len]
	if haxe__io__resolveEncodingTag(encoding...) == 1 {
		return haxe__io__bytes_toStringRawNativeUTF16LE(slice)
	}
	return hxrt.BytesToString(slice)
}

func (self *haxe__io__Bytes) readString(pos int, len int) *string {
	return self.getString(pos, len)
}

func (self *haxe__io__Bytes) get(pos int) int {
	return self.b[pos]
}

func (self *haxe__io__Bytes) set(pos int, value int) {
	self.b[pos] = value
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) blit(pos int, src *haxe__io__Bytes, srcpos int, len int) {
	if self == nil || src == nil || pos < 0 || srcpos < 0 || len < 0 || pos+len > self.length || srcpos+len > src.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	if len == 0 {
		return
	}
	if self == src && pos > srcpos {
		for i := len - 1; i >= 0; i-- {
			self.b[pos+i] = src.b[srcpos+i]
		}
	} else {
		for i := 0; i < len; i++ {
			self.b[pos+i] = src.b[srcpos+i]
		}
	}
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) fill(pos int, len int, value int) {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	masked := value & 255
	for i := 0; i < len; i++ {
		self.b[pos+i] = masked
	}
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) sub(pos int, len int) *haxe__io__Bytes {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	if len == 0 {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	copied := make([]int, len)
	copy(copied, self.b[pos:pos+len])
	return &haxe__io__Bytes{b: copied, length: len}
}

func (self *haxe__io__Bytes) compare(other *haxe__io__Bytes) int {
	if self == nil && other == nil {
		return 0
	}
	if self == nil {
		return -1
	}
	if other == nil {
		return 1
	}
	limit := self.length
	if other.length < limit {
		limit = other.length
	}
	for i := 0; i < limit; i++ {
		if self.b[i] < other.b[i] {
			return -1
		}
		if self.b[i] > other.b[i] {
			return 1
		}
	}
	if self.length < other.length {
		return -1
	}
	if self.length > other.length {
		return 1
	}
	return 0
}

func New_haxe__io__BytesBuffer() *haxe__io__BytesBuffer {
	return &haxe__io__BytesBuffer{b: []int{}}
}

func (self *haxe__io__BytesBuffer) addByte(value int) {
	self.b = hxrt.BytesBufferAddByte(self.b, value)
}

func (self *haxe__io__BytesBuffer) add(src *haxe__io__Bytes) {
	if src == nil {
		return
	}
	self.b = hxrt.BytesBufferAdd(self.b, src.b)
}

func (self *haxe__io__BytesBuffer) addBytes(src *haxe__io__Bytes, pos int, len int) {
	if src == nil || pos < 0 || len < 0 || pos+len > src.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	if len == 0 {
		return
	}
	self.b = hxrt.BytesBufferAddSlice(self.b, src.b, pos, len)
}

func (self *haxe__io__BytesBuffer) addString(value *string, encoding ...*haxe__io__Encoding) {
	self.add(haxe__io__Bytes_ofString(value))
}

func (self *haxe__io__BytesBuffer) getBytes() *haxe__io__Bytes {
	copied := hxrt.BytesClone(self.b)
	return &haxe__io__Bytes{b: copied, length: len(copied)}
}

func (self *haxe__io__BytesBuffer) get_length() int {
	return hxrt.BytesBufferLength(self.b)
}

func New_haxe__io__BytesInput(b *haxe__io__Bytes, opts ...int) *haxe__io__BytesInput {
	if b == nil {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__BytesInput{}
	}
	start := 0
	if len(opts) > 0 {
		start = opts[0]
	}
	sliceLen := (b.length - start)
	if len(opts) > 1 {
		sliceLen = opts[1]
	}
	if start < 0 || sliceLen < 0 || start+sliceLen > b.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__BytesInput{}
	}
	return &haxe__io__BytesInput{b: b.b, pos: start, len: sliceLen, totlen: sliceLen}
}

func (self *haxe__io__BytesInput) get_position() int {
	return self.pos
}

func (self *haxe__io__BytesInput) set_position(p int) int {
	if p < 0 {
		p = 0
	} else {
		if p > self.totlen {
			p = self.totlen
		}
	}
	self.len = (self.totlen - p)
	self.pos = p
	return p
}

func (self *haxe__io__BytesInput) get_length() int {
	return self.totlen
}

func (self *haxe__io__BytesInput) readByte() int {
	if self == nil || self.len == 0 {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	self.len = (self.len - 1)
	value := self.b[self.pos]
	self.pos = (self.pos + 1)
	return value
}

func (self *haxe__io__BytesInput) readBytes(buf *haxe__io__Bytes, pos int, len int) int {
	if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if len > 0 && (self == nil || self.len == 0) {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	if self == nil {
		return 0
	}
	if self.len < len {
		len = self.len
	}
	for i := 0; i < len; i++ {
		buf.b[pos+i] = self.b[self.pos+i]
	}
	self.pos += len
	self.len -= len
	return len
}

func (self *haxe__io__BytesInput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.bigEndian
}

func (self *haxe__io__BytesInput) set_bigEndian(e bool) bool {
	if self != nil {
		self.bigEndian = e
	}
	return e
}

func (self *haxe__io__BytesInput) close() {
	_ = self
}

func New_haxe__io__BytesOutput() *haxe__io__BytesOutput {
	return &haxe__io__BytesOutput{b: &haxe__io__BytesBuffer{b: []int{}}}
}

func (self *haxe__io__BytesOutput) get_length() int {
	if self == nil || self.b == nil {
		return 0
	}
	return self.b.get_length()
}

func (self *haxe__io__BytesOutput) writeByte(c int) {
	if self == nil || self.b == nil {
		return
	}
	self.b.addByte(c)
}

func (self *haxe__io__BytesOutput) writeBytes(buf *haxe__io__Bytes, pos int, len int) int {
	if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if self == nil || self.b == nil {
		return 0
	}
	self.b.addBytes(buf, pos, len)
	return len
}

func (self *haxe__io__BytesOutput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.bigEndian
}

func (self *haxe__io__BytesOutput) set_bigEndian(e bool) bool {
	if self != nil {
		self.bigEndian = e
	}
	return e
}

func (self *haxe__io__BytesOutput) flush() {
	_ = self
}

func (self *haxe__io__BytesOutput) close() {
	_ = self
}

func (self *haxe__io__BytesOutput) getBytes() *haxe__io__Bytes {
	if self == nil || self.b == nil {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return self.b.getBytes()
}

type sys__net__Host struct {
	host     *string
	ip       int
	resolved *string
}

func hxrt__host_empty() *sys__net__Host {
	return &sys__net__Host{host: hxrt.StringFromLiteral(""), ip: 0, resolved: hxrt.StringFromLiteral("")}
}

func hxrt__host_ipv4Int(ip net.IP) int {
	if ip == nil {
		return 0
	}
	v4 := ip.To4()
	if v4 == nil {
		return 0
	}
	return int(v4[0])<<24 | int(v4[1])<<16 | int(v4[2])<<8 | int(v4[3])
}

func hxrt__host_ipv4String(value int) *string {
	u := uint32(value)
	rendered := net.IPv4(byte((u>>24)&0xff), byte((u>>16)&0xff), byte((u>>8)&0xff), byte(u&0xff)).String()
	return hxrt.StringFromLiteral(rendered)
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
	ip := hxrt__host_ipv4Int(selected)
	return &sys__net__Host{host: name, ip: ip, resolved: resolved}
}

func (self *sys__net__Host) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	if self.resolved != nil && *self.resolved != "" {
		return self.resolved
	}
	if self.ip != 0 {
		return hxrt__host_ipv4String(self.ip)
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

type sys__net__UdpSocket struct {
	*sys__net__Socket
}

func New_sys__net__Socket() *sys__net__Socket {
	return &sys__net__Socket{input: &sys__net__SocketInput{}, output: &sys__net__SocketOutput{}, blocking: true}
}

func New_sys__net__UdpSocket() *sys__net__UdpSocket {
	return &sys__net__UdpSocket{sys__net__Socket: New_sys__net__Socket()}
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

func (self *sys__net__Socket) hxrt__socket_conn() net.Conn {
	if self == nil {
		return nil
	}
	return self.conn
}

func (self *sys__net__UdpSocket) hxrt__udp_socket_conn(create bool) *net.UDPConn {
	if self == nil || self.sys__net__Socket == nil {
		hxrt.Throw(hxrt.StringFromLiteral("udp socket is nil"))
		return nil
	}
	if self.conn != nil {
		udpConn, ok := self.conn.(*net.UDPConn)
		if !ok {
			hxrt.Throw(hxrt.StringFromLiteral("udp socket expects UDP connection"))
			return nil
		}
		return udpConn
	}
	if !create {
		return nil
	}
	conn, err := net.ListenUDP("udp4", nil)
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	self.hxrt__socket_setConn(conn)
	return conn
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

func (self *sys__net__UdpSocket) bind(host *sys__net__Host, port int) {
	if self == nil || host == nil || self.sys__net__Socket == nil {
		hxrt.Throw(hxrt.StringFromLiteral("udp bind requires host"))
		return
	}
	resolvedHost := host.toString()
	if resolvedHost == nil {
		hxrt.Throw(hxrt.StringFromLiteral("udp bind requires host"))
		return
	}
	udpAddr, err := net.ResolveUDPAddr("udp4", net.JoinHostPort(*hxrt.StdString(resolvedHost), strconv.Itoa(port)))
	if err != nil {
		hxrt.Throw(err)
		return
	}
	conn, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		hxrt.Throw(err)
		return
	}
	self.close()
	self.hxrt__socket_setConn(conn)
}

func (self *sys__net__UdpSocket) setBroadcast(enabled bool) {
	_ = self
	_ = enabled
}

func (self *sys__net__UdpSocket) sendTo(buf *haxe__io__Bytes, pos int, length int, addr *sys__net__Address) int {
	if buf == nil || addr == nil {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if pos < 0 || length < 0 || pos+length > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	conn := self.hxrt__udp_socket_conn(true)
	if conn == nil {
		return 0
	}
	self.hxrt__socket_applyConnDeadline()
	raw := make([]byte, buf.length)
	for i := 0; i < buf.length; i++ {
		raw[i] = byte(buf.b[i])
	}
	target := &net.UDPAddr{IP: net.IPv4(byte((uint32(addr.host)>>24)&0xff), byte((uint32(addr.host)>>16)&0xff), byte((uint32(addr.host)>>8)&0xff), byte(uint32(addr.host)&0xff)), Port: addr.port}
	written, err := conn.WriteToUDP(raw[pos:pos+length], target)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			hxrt.Throw(haxe__io__Error_Blocked)
			return 0
		}
		hxrt.Throw(err)
		return 0
	}
	return written
}

func (self *sys__net__UdpSocket) readFrom(buf *haxe__io__Bytes, pos int, length int, addr *sys__net__Address) int {
	if buf == nil || addr == nil {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if pos < 0 || length < 0 || pos+length > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	conn := self.hxrt__udp_socket_conn(false)
	if conn == nil {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	self.hxrt__socket_applyConnDeadline()
	raw := make([]byte, length)
	read, remote, err := conn.ReadFromUDP(raw)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			hxrt.Throw(haxe__io__Error_Blocked)
			return 0
		}
		hxrt.Throw(err)
		return 0
	}
	if read <= 0 {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	for i := 0; i < read; i++ {
		buf.b[pos+i] = int(raw[i])
	}
	buf.__hx_rawValid = false
	addr.host = hxrt__host_ipv4Int(remote.IP)
	addr.port = remote.Port
	return read
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
