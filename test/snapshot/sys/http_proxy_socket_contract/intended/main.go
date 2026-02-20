package main

import (
	"bytes"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"snapshot/hxrt"
	"strings"
	"time"
)

func main() {
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy0="), sys__Http_hxrt_proxyDescriptor()))
	hx_obj_1 := map[string]any{}
	hx_obj_1["host"] = hxrt.StringFromLiteral("proxy.local")
	hx_obj_1["port"] = 3128
	hx_obj_2 := map[string]any{}
	hx_obj_2["user"] = hxrt.StringFromLiteral("scott")
	hx_obj_2["pass"] = hxrt.StringFromLiteral("tiger")
	hx_obj_1["auth"] = hx_obj_2
	sys__Http_PROXY = hx_obj_1
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy1="), sys__Http_hxrt_proxyDescriptor()))
	hx_obj_3 := map[string]any{}
	hx_obj_3["host"] = hxrt.StringFromLiteral("proxy.local")
	hx_obj_3["port"] = 80
	hx_obj_4 := map[string]any{}
	hx_obj_4["user"] = hxrt.StringFromLiteral("scott")
	hx_obj_4["pass"] = nil
	hx_obj_3["auth"] = hx_obj_4
	sys__Http_PROXY = hx_obj_3
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy2="), sys__Http_hxrt_proxyDescriptor()))
	hx_obj_5 := map[string]any{}
	hx_obj_5["host"] = hxrt.StringFromLiteral("proxy.local:9000")
	hx_obj_5["port"] = 3128
	hx_obj_5["auth"] = nil
	sys__Http_PROXY = hx_obj_5
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy3="), sys__Http_hxrt_proxyDescriptor()))
	hx_obj_6 := map[string]any{}
	hx_obj_6["host"] = nil
	hx_obj_6["port"] = 3128
	hx_obj_6["auth"] = nil
	sys__Http_PROXY = hx_obj_6
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy4="), sys__Http_hxrt_proxyDescriptor()))
	sys__Http_PROXY = nil
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,body"))
	_ = http
	sink := New_haxe__io__BytesBuffer()
	_ = sink
	http.customRequest(false, sink, func() map[string]any {
		hx_obj_7 := map[string]any{}
		hx_obj_7["marker"] = hxrt.StringFromLiteral("sock")
		return hx_obj_7
	}(), hxrt.StringFromLiteral("PATCH"))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("methodSock="), sink.getBytes().toString()))
}

type haxe__io__Encoding struct {
}

type haxe__io__Input interface {
	get_bigEndian() bool
	set_bigEndian(e bool) bool
	readByte() int
	readBytes(buf *haxe__io__Bytes, pos int, len int) int
	close()
	readAll(bufsize ...int) *haxe__io__Bytes
	readFullBytes(s *haxe__io__Bytes, pos int, len int)
	read(nbytes int) *haxe__io__Bytes
	readUntil(end int) *string
	readLine() *string
	readFloat() float64
	readDouble() float64
	readInt8() int
	readInt16() int
	readUInt16() int
	readInt24() int
	readUInt24() int
	readInt32() int
	readString(len int, encoding ...*haxe__io__Encoding) *string
}

type haxe__io__Output interface {
	get_bigEndian() bool
	set_bigEndian(e bool) bool
	writeByte(c int)
	writeBytes(s *haxe__io__Bytes, pos int, len int) int
	flush()
	close()
	write(s *haxe__io__Bytes)
	writeFullBytes(s *haxe__io__Bytes, pos int, len int)
	writeFloat(x float64)
	writeDouble(x float64)
	writeInt8(x int)
	writeInt16(x int)
	writeUInt16(x int)
	writeInt24(x int)
	writeUInt24(x int)
	writeInt32(x int)
	prepare(nbytes int)
	writeInput(i haxe__io__Input, bufsize ...int)
	writeString(s *string, encoding ...*haxe__io__Encoding)
}

type haxe__io__Eof struct {
}

type haxe__io__Error struct {
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

func (self *haxe__io__Eof) toString() *string {
	return hxrt.StringFromLiteral("Eof")
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
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func (self *haxe__io__Bytes) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToString(self.b)
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
		hxrt.Throw(hxrt.StringFromLiteral("OutsideBounds"))
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
		hxrt.Throw(hxrt.StringFromLiteral("OutsideBounds"))
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
		hxrt.Throw(hxrt.StringFromLiteral("OutsideBounds"))
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
	self.b = append(self.b, (value & 255))
}

func (self *haxe__io__BytesBuffer) add(src *haxe__io__Bytes) {
	if src == nil {
		return
	}
	self.b = append(self.b, src.b...)
}

func (self *haxe__io__BytesBuffer) addBytes(src *haxe__io__Bytes, pos int, len int) {
	if src == nil || pos < 0 || len < 0 || pos+len > src.length {
		hxrt.Throw(hxrt.StringFromLiteral("OutsideBounds"))
		return
	}
	if len == 0 {
		return
	}
	self.b = append(self.b, src.b[pos:pos+len]...)
}

func (self *haxe__io__BytesBuffer) addString(value *string, encoding ...*haxe__io__Encoding) {
	self.add(haxe__io__Bytes_ofString(value))
}

func (self *haxe__io__BytesBuffer) getBytes() *haxe__io__Bytes {
	copied := hxrt.BytesClone(self.b)
	return &haxe__io__Bytes{b: copied, length: len(copied)}
}

func (self *haxe__io__BytesBuffer) get_length() int {
	return len(self.b)
}

func haxe__io__input_isEof(value any) bool {
	_, ok := value.(*haxe__io__Eof)
	return ok
}

func haxe__io__input_readAll(self haxe__io__Input, bufsize ...int) *haxe__io__Bytes {
	if self == nil {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	resolved := 1 << 14
	if len(bufsize) > 0 {
		resolved = bufsize[0]
	}
	buf := haxe__io__Bytes_alloc(resolved)
	total := New_haxe__io__BytesBuffer()
	for {
		chunk := 0
		threw := false
		var thrown any
		func() {
			defer func() {
				if recovered := recover(); recovered != nil {
					threw = true
					thrown = hxrt.UnwrapException(recovered)
				}
			}()
			chunk = self.readBytes(buf, 0, resolved)
		}()
		if threw {
			if haxe__io__input_isEof(thrown) {
				break
			}
			hxrt.Throw(thrown)
			return &haxe__io__Bytes{b: []int{}, length: 0}
		}
		if chunk == 0 {
			hxrt.Throw(hxrt.StringFromLiteral("Blocked"))
			return &haxe__io__Bytes{b: []int{}, length: 0}
		}
		total.addBytes(buf, 0, chunk)
	}
	return total.getBytes()
}

func haxe__io__input_readFullBytes(self haxe__io__Input, s *haxe__io__Bytes, pos int, len int) {
	if self == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Blocked"))
		return
	}
	for len > 0 {
		k := self.readBytes(s, pos, len)
		if k == 0 {
			hxrt.Throw(hxrt.StringFromLiteral("Blocked"))
			return
		}
		pos += k
		len -= k
	}
}

func haxe__io__input_read(self haxe__io__Input, nbytes int) *haxe__io__Bytes {
	s := haxe__io__Bytes_alloc(nbytes)
	p := 0
	for nbytes > 0 {
		k := self.readBytes(s, p, nbytes)
		if k == 0 {
			hxrt.Throw(hxrt.StringFromLiteral("Blocked"))
			return &haxe__io__Bytes{b: []int{}, length: 0}
		}
		p += k
		nbytes -= k
	}
	return s
}

func haxe__io__input_readUntil(self haxe__io__Input, end int) *string {
	buf := New_haxe__io__BytesBuffer()
	for {
		last := self.readByte()
		if last == end {
			break
		}
		buf.addByte(last)
	}
	return buf.getBytes().toString()
}

func haxe__io__input_readLine(self haxe__io__Input) *string {
	buf := New_haxe__io__BytesBuffer()
	for {
		last := 0
		threw := false
		var thrown any
		func() {
			defer func() {
				if recovered := recover(); recovered != nil {
					threw = true
					thrown = hxrt.UnwrapException(recovered)
				}
			}()
			last = self.readByte()
		}()
		if threw {
			if haxe__io__input_isEof(thrown) {
				s := buf.getBytes().toString()
				raw := *hxrt.StdString(s)
				if len(raw) == 0 {
					hxrt.Throw(thrown)
					return hxrt.StringFromLiteral("")
				}
				return s
			}
			hxrt.Throw(thrown)
			return hxrt.StringFromLiteral("")
		}
		if last == 10 {
			break
		}
		buf.addByte(last)
	}
	s := buf.getBytes().toString()
	raw := *hxrt.StdString(s)
	if len(raw) > 0 && raw[len(raw)-1] == 13 {
		raw = raw[:len(raw)-1]
	}
	return hxrt.StringFromLiteral(raw)
}

func haxe__io__input_readFloat(self haxe__io__Input) float64 {
	bits := uint32(self.readInt32())
	return float64(math.Float32frombits(bits))
}

func haxe__io__input_readDouble(self haxe__io__Input) float64 {
	i1 := self.readInt32()
	i2 := self.readInt32()
	if self.get_bigEndian() {
		return math.Float64frombits((uint64(uint32(i1)) << 32) | uint64(uint32(i2)))
	}
	return math.Float64frombits((uint64(uint32(i2)) << 32) | uint64(uint32(i1)))
}

func haxe__io__input_readInt8(self haxe__io__Input) int {
	n := self.readByte()
	if n >= 128 {
		return n - 256
	}
	return n
}

func haxe__io__input_readInt16(self haxe__io__Input) int {
	ch1 := self.readByte()
	ch2 := self.readByte()
	n := 0
	if self.get_bigEndian() {
		n = ch2 | (ch1 << 8)
	} else {
		n = ch1 | (ch2 << 8)
	}
	if (n & 0x8000) != 0 {
		return n - 0x10000
	}
	return n
}

func haxe__io__input_readUInt16(self haxe__io__Input) int {
	ch1 := self.readByte()
	ch2 := self.readByte()
	if self.get_bigEndian() {
		return ch2 | (ch1 << 8)
	}
	return ch1 | (ch2 << 8)
}

func haxe__io__input_readInt24(self haxe__io__Input) int {
	ch1 := self.readByte()
	ch2 := self.readByte()
	ch3 := self.readByte()
	n := 0
	if self.get_bigEndian() {
		n = ch3 | (ch2 << 8) | (ch1 << 16)
	} else {
		n = ch1 | (ch2 << 8) | (ch3 << 16)
	}
	if (n & 0x800000) != 0 {
		return n - 0x1000000
	}
	return n
}

func haxe__io__input_readUInt24(self haxe__io__Input) int {
	ch1 := self.readByte()
	ch2 := self.readByte()
	ch3 := self.readByte()
	if self.get_bigEndian() {
		return ch3 | (ch2 << 8) | (ch1 << 16)
	}
	return ch1 | (ch2 << 8) | (ch3 << 16)
}

func haxe__io__input_readInt32(self haxe__io__Input) int {
	ch1 := self.readByte()
	ch2 := self.readByte()
	ch3 := self.readByte()
	ch4 := self.readByte()
	if self.get_bigEndian() {
		return ch4 | (ch3 << 8) | (ch2 << 16) | (ch1 << 24)
	}
	return ch1 | (ch2 << 8) | (ch3 << 16) | (ch4 << 24)
}

func haxe__io__input_readString(self haxe__io__Input, len int, encoding ...*haxe__io__Encoding) *string {
	b := haxe__io__Bytes_alloc(len)
	haxe__io__input_readFullBytes(self, b, 0, len)
	return b.toString()
}

func haxe__io__output_write(self haxe__io__Output, s *haxe__io__Bytes) {
	if self == nil || s == nil {
		return
	}
	remaining := s.length
	p := 0
	for remaining > 0 {
		k := self.writeBytes(s, p, remaining)
		if k == 0 {
			hxrt.Throw(hxrt.StringFromLiteral("Blocked"))
			return
		}
		p += k
		remaining -= k
	}
}

func haxe__io__output_writeFullBytes(self haxe__io__Output, s *haxe__io__Bytes, pos int, len int) {
	for len > 0 {
		k := self.writeBytes(s, pos, len)
		pos += k
		len -= k
	}
}

func haxe__io__output_writeFloat(self haxe__io__Output, x float64) {
	self.writeInt32(int(math.Float32bits(float32(x))))
}

func haxe__io__output_writeDouble(self haxe__io__Output, x float64) {
	bits := math.Float64bits(x)
	low := int(uint32(bits))
	high := int(uint32(bits >> 32))
	if self.get_bigEndian() {
		self.writeInt32(high)
		self.writeInt32(low)
		return
	}
	self.writeInt32(low)
	self.writeInt32(high)
}

func haxe__io__output_writeInt8(self haxe__io__Output, x int) {
	if x < -0x80 || x >= 0x80 {
		hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		return
	}
	self.writeByte(x & 0xFF)
}

func haxe__io__output_writeInt16(self haxe__io__Output, x int) {
	if x < -0x8000 || x >= 0x8000 {
		hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		return
	}
	self.writeUInt16(x & 0xFFFF)
}

func haxe__io__output_writeUInt16(self haxe__io__Output, x int) {
	if x < 0 || x >= 0x10000 {
		hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		return
	}
	if self.get_bigEndian() {
		self.writeByte(x >> 8)
		self.writeByte(x & 0xFF)
		return
	}
	self.writeByte(x & 0xFF)
	self.writeByte(x >> 8)
}

func haxe__io__output_writeInt24(self haxe__io__Output, x int) {
	if x < -0x800000 || x >= 0x800000 {
		hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		return
	}
	self.writeUInt24(x & 0xFFFFFF)
}

func haxe__io__output_writeUInt24(self haxe__io__Output, x int) {
	if x < 0 || x >= 0x1000000 {
		hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		return
	}
	if self.get_bigEndian() {
		self.writeByte(x >> 16)
		self.writeByte((x >> 8) & 0xFF)
		self.writeByte(x & 0xFF)
		return
	}
	self.writeByte(x & 0xFF)
	self.writeByte((x >> 8) & 0xFF)
	self.writeByte(x >> 16)
}

func haxe__io__output_writeInt32(self haxe__io__Output, x int) {
	if self.get_bigEndian() {
		self.writeByte(int(uint(x) >> 24))
		self.writeByte((x >> 16) & 0xFF)
		self.writeByte((x >> 8) & 0xFF)
		self.writeByte(x & 0xFF)
		return
	}
	self.writeByte(x & 0xFF)
	self.writeByte((x >> 8) & 0xFF)
	self.writeByte((x >> 16) & 0xFF)
	self.writeByte(int(uint(x) >> 24))
}

func haxe__io__output_writeInput(self haxe__io__Output, i haxe__io__Input, bufsize ...int) {
	if self == nil || i == nil {
		return
	}
	resolved := 4096
	if len(bufsize) > 0 {
		resolved = bufsize[0]
	}
	buf := haxe__io__Bytes_alloc(resolved)
	for {
		lenRead := 0
		threw := false
		var thrown any
		func() {
			defer func() {
				if recovered := recover(); recovered != nil {
					threw = true
					thrown = hxrt.UnwrapException(recovered)
				}
			}()
			lenRead = i.readBytes(buf, 0, resolved)
		}()
		if threw {
			if haxe__io__input_isEof(thrown) {
				break
			}
			hxrt.Throw(thrown)
			return
		}
		if lenRead == 0 {
			hxrt.Throw(hxrt.StringFromLiteral("Blocked"))
			return
		}
		p := 0
		for lenRead > 0 {
			k := self.writeBytes(buf, p, lenRead)
			if k == 0 {
				hxrt.Throw(hxrt.StringFromLiteral("Blocked"))
				return
			}
			p += k
			lenRead -= k
		}
	}
}

func haxe__io__output_writeString(self haxe__io__Output, s *string, encoding ...*haxe__io__Encoding) {
	if s == nil {
		s = hxrt.StringFromLiteral("")
	}
	b := haxe__io__Bytes_ofString(s, encoding...)
	self.writeFullBytes(b, 0, b.length)
}

func New_haxe__io__BytesInput(b *haxe__io__Bytes, opts ...int) *haxe__io__BytesInput {
	if b == nil {
		hxrt.Throw(hxrt.StringFromLiteral("OutsideBounds"))
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
		hxrt.Throw(hxrt.StringFromLiteral("OutsideBounds"))
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
		hxrt.Throw(hxrt.StringFromLiteral("OutsideBounds"))
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

func (self *haxe__io__BytesInput) readAll(bufsize ...int) *haxe__io__Bytes {
	return haxe__io__input_readAll(self, bufsize...)
}

func (self *haxe__io__BytesInput) readFullBytes(s *haxe__io__Bytes, pos int, len int) {
	haxe__io__input_readFullBytes(self, s, pos, len)
}

func (self *haxe__io__BytesInput) read(nbytes int) *haxe__io__Bytes {
	return haxe__io__input_read(self, nbytes)
}

func (self *haxe__io__BytesInput) readUntil(end int) *string {
	return haxe__io__input_readUntil(self, end)
}

func (self *haxe__io__BytesInput) readLine() *string {
	return haxe__io__input_readLine(self)
}

func (self *haxe__io__BytesInput) readFloat() float64 {
	return haxe__io__input_readFloat(self)
}

func (self *haxe__io__BytesInput) readDouble() float64 {
	return haxe__io__input_readDouble(self)
}

func (self *haxe__io__BytesInput) readInt8() int {
	return haxe__io__input_readInt8(self)
}

func (self *haxe__io__BytesInput) readInt16() int {
	return haxe__io__input_readInt16(self)
}

func (self *haxe__io__BytesInput) readUInt16() int {
	return haxe__io__input_readUInt16(self)
}

func (self *haxe__io__BytesInput) readInt24() int {
	return haxe__io__input_readInt24(self)
}

func (self *haxe__io__BytesInput) readUInt24() int {
	return haxe__io__input_readUInt24(self)
}

func (self *haxe__io__BytesInput) readInt32() int {
	return haxe__io__input_readInt32(self)
}

func (self *haxe__io__BytesInput) readString(len int, encoding ...*haxe__io__Encoding) *string {
	return haxe__io__input_readString(self, len, encoding...)
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
		hxrt.Throw(hxrt.StringFromLiteral("OutsideBounds"))
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

func (self *haxe__io__BytesOutput) write(s *haxe__io__Bytes) {
	haxe__io__output_write(self, s)
}

func (self *haxe__io__BytesOutput) writeFullBytes(s *haxe__io__Bytes, pos int, len int) {
	haxe__io__output_writeFullBytes(self, s, pos, len)
}

func (self *haxe__io__BytesOutput) writeFloat(x float64) {
	haxe__io__output_writeFloat(self, x)
}

func (self *haxe__io__BytesOutput) writeDouble(x float64) {
	haxe__io__output_writeDouble(self, x)
}

func (self *haxe__io__BytesOutput) writeInt8(x int) {
	haxe__io__output_writeInt8(self, x)
}

func (self *haxe__io__BytesOutput) writeInt16(x int) {
	haxe__io__output_writeInt16(self, x)
}

func (self *haxe__io__BytesOutput) writeUInt16(x int) {
	haxe__io__output_writeUInt16(self, x)
}

func (self *haxe__io__BytesOutput) writeInt24(x int) {
	haxe__io__output_writeInt24(self, x)
}

func (self *haxe__io__BytesOutput) writeUInt24(x int) {
	haxe__io__output_writeUInt24(self, x)
}

func (self *haxe__io__BytesOutput) writeInt32(x int) {
	haxe__io__output_writeInt32(self, x)
}

func (self *haxe__io__BytesOutput) prepare(nbytes int) {
	_ = self
	_ = nbytes
}

func (self *haxe__io__BytesOutput) writeInput(i haxe__io__Input, bufsize ...int) {
	haxe__io__output_writeInput(self, i, bufsize...)
}

func (self *haxe__io__BytesOutput) writeString(s *string, encoding ...*haxe__io__Encoding) {
	haxe__io__output_writeString(self, s, encoding...)
}

func (self *haxe__io__BytesOutput) getBytes() *haxe__io__Bytes {
	if self == nil || self.b == nil {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return self.b.getBytes()
}

type haxe__ds__IntMap struct {
	h map[int]any
}

type haxe__ds__StringMap struct {
	h map[string]any
}

type haxe__ds__ObjectMap struct {
	h map[any]any
}

type haxe__ds__EnumValueMap struct {
	h map[any]any
}

type haxe__ds__List struct {
	items  []any
	length int
}

func New_haxe__ds__IntMap() *haxe__ds__IntMap {
	return &haxe__ds__IntMap{h: map[int]any{}}
}

func (self *haxe__ds__IntMap) set(key int, value any) {
	self.h[key] = value
}

func (self *haxe__ds__IntMap) get(key int) any {
	value := self.h[key]
	return value
}

func (self *haxe__ds__IntMap) exists(key int) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__IntMap) remove(key int) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return &haxe__ds__StringMap{h: map[string]any{}}
}

func (self *haxe__ds__StringMap) set(key *string, value any) {
	self.h[*hxrt.StdString(key)] = value
}

func (self *haxe__ds__StringMap) get(key *string) any {
	value := self.h[*hxrt.StdString(key)]
	return value
}

func (self *haxe__ds__StringMap) exists(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	return ok
}

func (self *haxe__ds__StringMap) remove(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	delete(self.h, *hxrt.StdString(key))
	return ok
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	return &haxe__ds__ObjectMap{h: map[any]any{}}
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	return &haxe__ds__EnumValueMap{h: map[any]any{}}
}

func (self *haxe__ds__EnumValueMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__List() *haxe__ds__List {
	return &haxe__ds__List{items: []any{}, length: 0}
}

func (self *haxe__ds__List) add(item any) {
	self.items = append(self.items, item)
	self.length = len(self.items)
}

func (self *haxe__ds__List) push(item any) {
	self.items = append([]any{item}, self.items...)
	self.length = len(self.items)
}

func (self *haxe__ds__List) pop() any {
	if len(self.items) == 0 {
		return nil
	}
	head := self.items[0]
	self.items = self.items[1:]
	self.length = len(self.items)
	return head
}

func (self *haxe__ds__List) first() any {
	if len(self.items) == 0 {
		return nil
	}
	return self.items[0]
}

func (self *haxe__ds__List) last() any {
	size := len(self.items)
	if size == 0 {
		return nil
	}
	return self.items[(size - 1)]
}

type hxrt__http__Pair struct {
	name  *string
	value *string
}

type hxrt__http__FileUpload struct {
	param    *string
	filename *string
	size     int
	mimeType *string
	fileRef  any
}

var sys__Http_PROXY any = nil

type sys__Http struct {
	url                    *string
	responseAsString       *string
	responseBytes          *haxe__io__Bytes
	postData               *string
	postBytes              *haxe__io__Bytes
	headers                []hxrt__http__Pair
	params                 []hxrt__http__Pair
	onData                 func(*string)
	onBytes                func(*haxe__io__Bytes)
	onError                func(*string)
	onStatus               func(int)
	noShutdown             bool
	cnxTimeout             float64
	responseHeaders        *haxe__ds__StringMap
	responseHeadersSameKey map[string][]*string
	fileUpload             *hxrt__http__FileUpload
}

func New_sys__Http(url *string) *sys__Http {
	self := &sys__Http{url: url, headers: []hxrt__http__Pair{}, params: []hxrt__http__Pair{}, cnxTimeout: 10, responseHeaders: New_haxe__ds__StringMap(), responseHeadersSameKey: map[string][]*string{}}
	self.onData = func(data *string) {}
	self.onBytes = func(data *haxe__io__Bytes) {}
	self.onError = func(msg *string) {}
	self.onStatus = func(status int) {}
	return self
}

func (self *sys__Http) setHeader(name *string, value *string) {
	if self == nil {
		return
	}
	for i := 0; i < len(self.headers); i++ {
		if *hxrt.StdString(self.headers[i].name) == *hxrt.StdString(name) {
			self.headers[i] = hxrt__http__Pair{name: name, value: value}
			return
		}
	}
	self.headers = append(self.headers, hxrt__http__Pair{name: name, value: value})
}

func (self *sys__Http) addHeader(header *string, value *string) {
	if self == nil {
		return
	}
	self.headers = append(self.headers, hxrt__http__Pair{name: header, value: value})
}

func (self *sys__Http) setParameter(name *string, value *string) {
	if self == nil {
		return
	}
	for i := 0; i < len(self.params); i++ {
		if *hxrt.StdString(self.params[i].name) == *hxrt.StdString(name) {
			self.params[i] = hxrt__http__Pair{name: name, value: value}
			return
		}
	}
	self.params = append(self.params, hxrt__http__Pair{name: name, value: value})
}

func (self *sys__Http) addParameter(name *string, value *string) {
	if self == nil {
		return
	}
	self.params = append(self.params, hxrt__http__Pair{name: name, value: value})
}

func (self *sys__Http) setPostData(data *string) {
	if self == nil {
		return
	}
	self.postData = data
	self.postBytes = nil
}

func (self *sys__Http) setPostBytes(data *haxe__io__Bytes) {
	if self == nil {
		return
	}
	self.postBytes = data
	self.postData = nil
}

func (self *sys__Http) fileTransfer(argname *string, filename *string, file any, size int, mimeType ...*string) {
	if self == nil {
		return
	}
	resolvedMime := hxrt.StringFromLiteral("application/octet-stream")
	if len(mimeType) > 0 {
		if mimeType[0] != nil {
			resolvedMime = mimeType[0]
		}
	}
	self.fileUpload = &hxrt__http__FileUpload{param: argname, filename: filename, size: size, mimeType: resolvedMime, fileRef: file}
}

func (self *sys__Http) fileTransfert(argname *string, filename *string, file any, size int, mimeType ...*string) {
	self.fileTransfer(argname, filename, file, size, mimeType...)
}

func (self *sys__Http) getResponseHeaderValues(key *string) []*string {
	if self == nil {
		return nil
	}
	rawKey := *hxrt.StdString(key)
	normalized := strings.ToLower(rawKey)
	if self.responseHeadersSameKey != nil {
		if values, ok := self.responseHeadersSameKey[rawKey]; ok {
			return values
		}
		if values, ok := self.responseHeadersSameKey[normalized]; ok {
			return values
		}
	}
	if self.responseHeaders == nil {
		return nil
	}
	single := self.responseHeaders.get(hxrt.StringFromLiteral(rawKey))
	if single == nil && rawKey != normalized {
		single = self.responseHeaders.get(hxrt.StringFromLiteral(normalized))
	}
	if single == nil {
		return nil
	}
	return []*string{hxrt.StdString(single)}
}

func (self *sys__Http) get_responseData() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	if self.responseAsString == nil && self.responseBytes != nil {
		self.responseAsString = self.responseBytes.toString()
	}
	return self.responseAsString
}

func (self *sys__Http) customRequest(post bool, api any, rest ...any) {
	var socketOverride any = nil
	var methodOverride *string = nil
	if len(rest) >= 1 {
		switch candidate := rest[0].(type) {
		case string:
			if len(rest) == 1 {
				methodOverride = hxrt.StringFromLiteral(candidate)
			}
		case *string:
			if len(rest) == 1 {
				methodOverride = candidate
			}
		default:
			socketOverride = candidate
		}
	}
	if len(rest) >= 2 {
		switch candidate := rest[1].(type) {
		case *string:
			methodOverride = candidate
		case string:
			methodOverride = hxrt.StringFromLiteral(candidate)
		}
	}
	self.hxrt__http__requestWith(post, methodOverride, api, socketOverride)
}

func (self *sys__Http) request(post ...bool) {
	if self == nil {
		return
	}
	isPost := false
	if len(post) > 0 {
		isPost = post[0]
	}
	if self.postData != nil || self.postBytes != nil || self.fileUpload != nil {
		isPost = true
	}
	self.hxrt__http__requestWith(isPost, nil, nil, nil)
}

func (self *sys__Http) hxrt__http__requestWith(post bool, methodOverride *string, api any, sock any) {
	self.responseAsString = nil
	self.responseBytes = nil
	self.responseHeaders = New_haxe__ds__StringMap()
	self.responseHeadersSameKey = map[string][]*string{}
	rawUrl := *hxrt.StdString(self.url)
	parsedURL, err := url.Parse(rawUrl)
	if err != nil || parsedURL == nil {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral("Invalid URL"))
		}
		return
	}
	query := parsedURL.Query()
	for _, param := range self.params {
		query.Set(*hxrt.StdString(param.name), *hxrt.StdString(param.value))
	}
	var bodyReader io.Reader = nil
	var contentTypeOverride *string = nil
	if post {
		if self.fileUpload != nil {
			multipartPayload := ""
			for _, param := range self.params {
				multipartPayload += "--hxrt-go-boundary\r\n"
				multipartPayload += "Content-Disposition: form-data; name=\"" + *hxrt.StdString(param.name) + "\"\r\n\r\n"
				multipartPayload += *hxrt.StdString(param.value) + "\r\n"
			}
			multipartPayload += "--hxrt-go-boundary\r\n"
			multipartPayload += "Content-Disposition: form-data; name=\"" + *hxrt.StdString(self.fileUpload.param) + "\"; filename=\"" + *hxrt.StdString(self.fileUpload.filename) + "\"\r\n"
			multipartPayload += "Content-Type: " + *hxrt.StdString(self.fileUpload.mimeType) + "\r\n\r\n"
			multipartPayload += "[uploaded-bytes=" + *hxrt.StdString(self.fileUpload.size) + "]\r\n"
			multipartPayload += "--hxrt-go-boundary--\r\n"
			bodyReader = strings.NewReader(multipartPayload)
			contentTypeOverride = hxrt.StringFromLiteral("multipart/form-data; boundary=hxrt-go-boundary")
		} else if self.postBytes != nil {
			rawBody := make([]byte, len(self.postBytes.b))
			for i := 0; i < len(self.postBytes.b); i++ {
				rawBody[i] = byte(self.postBytes.b[i])
			}
			bodyReader = bytes.NewReader(rawBody)
		} else if self.postData != nil {
			bodyReader = strings.NewReader(*hxrt.StdString(self.postData))
		} else {
			encoded := query.Encode()
			bodyReader = strings.NewReader(encoded)
			hasContentType := false
			for _, header := range self.headers {
				if strings.EqualFold(*hxrt.StdString(header.name), "Content-Type") {
					hasContentType = true
					break
				}
			}
			if !hasContentType {
				contentTypeOverride = hxrt.StringFromLiteral("application/x-www-form-urlencoded")
			}
		}
	} else {
		parsedURL.RawQuery = query.Encode()
	}
	if parsedURL.Scheme == "data" {
		payload := parsedURL.Opaque
		mediaType := "text/plain"
		commaIndex := strings.Index(payload, ",")
		if commaIndex >= 0 {
			if commaIndex > 0 {
				mediaType = payload[:commaIndex]
			}
			payload = payload[commaIndex+1:]
		}
		if post {
			if self.fileUpload != nil {
				payload = "multipart file=" + *hxrt.StdString(self.fileUpload.filename) + ";mime=" + *hxrt.StdString(self.fileUpload.mimeType) + ";size=" + *hxrt.StdString(self.fileUpload.size)
			} else if bodyReader != nil {
				rawBody, readErr := io.ReadAll(bodyReader)
				if readErr == nil {
					payload = string(rawBody)
				}
			}
		}
		decoded, decodeErr := url.QueryUnescape(payload)
		if decodeErr == nil {
			payload = decoded
		}
		if methodOverride != nil {
			methodToken := strings.ToUpper(*hxrt.StdString(methodOverride))
			if methodToken != "" && methodToken != "NULL" {
				payload = methodToken + " " + payload
			}
		}
		rawPayload := []byte(payload)
		intPayload := make([]int, len(rawPayload))
		for i := 0; i < len(rawPayload); i++ {
			intPayload[i] = int(rawPayload[i])
		}
		self.responseBytes = &haxe__io__Bytes{b: intPayload, length: len(intPayload)}
		self.responseAsString = hxrt.StringFromLiteral(payload)
		self.responseHeaders = New_haxe__ds__StringMap()
		self.responseHeaders.set(hxrt.StringFromLiteral("content-type"), hxrt.StringFromLiteral(mediaType))
		self.responseHeaders.set(hxrt.StringFromLiteral("Content-Type"), hxrt.StringFromLiteral(mediaType))
		self.responseHeadersSameKey = map[string][]*string{}
		hxrt__http__captureApi(api, self.responseBytes)
		if self.onStatus != nil {
			self.onStatus(200)
		}
		if self.onData != nil {
			self.onData(self.responseAsString)
		}
		if self.onBytes != nil {
			self.onBytes(self.responseBytes)
		}
		return
	}
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral("Invalid URL"))
		}
		return
	}
	method := "GET"
	if post {
		method = "POST"
	}
	if methodOverride != nil {
		methodToken := strings.ToUpper(*hxrt.StdString(methodOverride))
		if methodToken != "" && methodToken != "NULL" {
			method = methodToken
		}
	}
	request, err := http.NewRequest(method, parsedURL.String(), bodyReader)
	if err != nil {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral(err.Error()))
		}
		return
	}
	for _, header := range self.headers {
		request.Header.Set(*hxrt.StdString(header.name), *hxrt.StdString(header.value))
	}
	if contentTypeOverride != nil && request.Header.Get("Content-Type") == "" {
		request.Header.Set("Content-Type", *hxrt.StdString(contentTypeOverride))
	}
	transport := &http.Transport{}
	proxyURL := hxrt__http__proxyURL()
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	var socketAdapter interface {
		hxrt__socket_conn() net.Conn
		hxrt__socket_setConn(net.Conn)
		close()
	}
	if candidate, ok := sock.(interface {
		hxrt__socket_conn() net.Conn
		hxrt__socket_setConn(net.Conn)
		close()
	}); ok {
		socketAdapter = candidate
		transport.DisableKeepAlives = true
		request.Close = true
		socketConsumed := false
		transport.Dial = func(network string, addr string) (net.Conn, error) {
			if socketConsumed {
				return nil, io.EOF
			}
			socketConsumed = true
			conn := socketAdapter.hxrt__socket_conn()
			if conn == nil {
				dialConn, dialErr := net.Dial(network, addr)
				if dialErr != nil {
					return nil, dialErr
				}
				socketAdapter.hxrt__socket_setConn(dialConn)
				conn = dialConn
			}
			return conn, nil
		}
		defer socketAdapter.close()
	}
	timeout := time.Duration(self.cnxTimeout * float64(time.Second))
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	client := &http.Client{Transport: transport, Timeout: timeout}
	response, err := client.Do(request)
	if err != nil {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral(err.Error()))
		}
		return
	}
	defer response.Body.Close()
	self.responseHeaders = New_haxe__ds__StringMap()
	self.responseHeadersSameKey = map[string][]*string{}
	for name, values := range response.Header {
		if len(values) == 0 {
			continue
		}
		lowerKey := strings.ToLower(name)
		lastValue := hxrt.StringFromLiteral(values[len(values)-1])
		self.responseHeaders.set(hxrt.StringFromLiteral(name), lastValue)
		if lowerKey != name {
			self.responseHeaders.set(hxrt.StringFromLiteral(lowerKey), lastValue)
		}
		if len(values) > 1 {
			allValues := make([]*string, 0, len(values))
			for _, rawValue := range values {
				allValues = append(allValues, hxrt.StringFromLiteral(rawValue))
			}
			self.responseHeadersSameKey[name] = allValues
			if lowerKey != name {
				self.responseHeadersSameKey[lowerKey] = allValues
			}
		}
	}
	if self.onStatus != nil {
		self.onStatus(response.StatusCode)
	}
	rawPayload, err := io.ReadAll(response.Body)
	if err != nil {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral(err.Error()))
		}
		return
	}
	intPayload := make([]int, len(rawPayload))
	for i := 0; i < len(rawPayload); i++ {
		intPayload[i] = int(rawPayload[i])
	}
	self.responseBytes = &haxe__io__Bytes{b: intPayload, length: len(intPayload)}
	self.responseAsString = hxrt.StringFromLiteral(string(rawPayload))
	hxrt__http__captureApi(api, self.responseBytes)
	if response.StatusCode >= 400 {
		if self.onError != nil {
			self.onError(hxrt.StringConcatAny(hxrt.StringFromLiteral("Http Error #"), response.StatusCode))
		}
		return
	}
	if self.onData != nil {
		self.onData(self.responseAsString)
	}
	if self.onBytes != nil {
		self.onBytes(self.responseBytes)
	}
}

func hxrt__http__captureApi(api any, payload *haxe__io__Bytes) {
	if api == nil || payload == nil {
		return
	}
	switch out := api.(type) {
	case *haxe__io__BytesBuffer:
		out.add(payload)
	case interface{ add(*haxe__io__Bytes) }:
		out.add(payload)
	case interface {
		writeBytes(*haxe__io__Bytes, int, int) int
	}:
		out.writeBytes(payload, 0, payload.length)
	case interface {
		writeFullBytes(*haxe__io__Bytes, int, int)
	}:
		out.writeFullBytes(payload, 0, payload.length)
	case interface{ writeString(*string) }:
		out.writeString(payload.toString())
	}
}

func hxrt__http__proxyURL() *url.URL {
	if sys__Http_PROXY == nil {
		return nil
	}
	config, ok := sys__Http_PROXY.(map[string]any)
	if !ok {
		return nil
	}
	host := *hxrt.StdString(config["host"])
	if host == "" {
		return nil
	}
	if host == "null" {
		return nil
	}
	port := *hxrt.StdString(config["port"])
	hostPort := host
	if port != "" && port != "null" && !strings.Contains(hostPort, ":") {
		hostPort = hostPort + ":" + port
	}
	proxyURL, err := url.Parse("http://" + hostPort)
	if err != nil {
		return nil
	}
	if authValue, ok := config["auth"]; ok {
		if authMap, ok := authValue.(map[string]any); ok {
			user := *hxrt.StdString(authMap["user"])
			pass := *hxrt.StdString(authMap["pass"])
			if user != "" && user != "null" {
				if pass == "null" {
					pass = ""
				}
				proxyURL.User = url.UserPassword(user, pass)
			}
		}
	}
	return proxyURL
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxyURL := hxrt__http__proxyURL()
	if proxyURL == nil {
		return hxrt.StringFromLiteral("null")
	}
	return hxrt.StringFromLiteral(proxyURL.String())
}

func sys__Http_requestUrl(url *string) *string {
	self := New_sys__Http(url)
	result := hxrt.StringFromLiteral("")
	self.onData = func(data *string) { result = data }
	self.onError = func(msg *string) { result = msg }
	self.request()
	return result
}
