package main

import (
	"bufio"
	"math"
	"net"
	"os"
	"snapshot/hxrt"
	"strconv"
	"strings"
	"time"
)

func main() {
	host := New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1"))
	addr := New_sys__net__Address()
	addr.host = host.ip
	addr.port = 3210
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("snapshot.host="), addr.getHost().toString()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("snapshot.compare="), addr.compare(addr.clone())))
	hxrt.Println(v_1)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("snapshot.alg="), any(hxrt.StringFromLiteral("SHA224"))), hxrt.StringFromLiteral(",")), any(hxrt.StringFromLiteral("SHA384"))), hxrt.StringFromLiteral(",")), any(hxrt.StringFromLiteral("RIPEMD160")))))
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

func (self *haxe__io__Eof) String() string {
	_ = self
	return "Eof"
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

func haxe__io__input_isEof(value any) bool {
	_, ok := value.(*haxe__io__Eof)
	return ok
}

func haxe__io__input_readAll(self haxe__io__Input, bufsize ...int) *haxe__io__Bytes {
	resolved := 1 << 14
	if len(bufsize) > 0 {
		resolved = bufsize[0]
	}
	return haxe__io__GoIoHelpers_inputReadAll(self, resolved)
}

func haxe__io__input_readFullBytes(self haxe__io__Input, s *haxe__io__Bytes, pos int, len int) {
	haxe__io__GoIoHelpers_inputReadFullBytes(self, s, pos, len)
}

func haxe__io__input_read(self haxe__io__Input, nbytes int) *haxe__io__Bytes {
	return haxe__io__GoIoHelpers_inputRead(self, nbytes)
}

func haxe__io__input_readUntil(self haxe__io__Input, end int) *string {
	return haxe__io__GoIoHelpers_inputReadUntil(self, end)
}

func haxe__io__input_readLine(self haxe__io__Input) *string {
	return haxe__io__GoIoHelpers_inputReadLine(self)
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
	return b.getString(0, len, encoding...)
}

func haxe__io__output_write(self haxe__io__Output, s *haxe__io__Bytes) {
	haxe__io__GoIoHelpers_outputWrite(self, s)
}

func haxe__io__output_writeFullBytes(self haxe__io__Output, s *haxe__io__Bytes, pos int, len int) {
	haxe__io__GoIoHelpers_outputWriteFullBytes(self, s, pos, len)
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
		hxrt.Throw(haxe__io__Error_Overflow)
		return
	}
	self.writeByte(x & 0xFF)
}

func haxe__io__output_writeInt16(self haxe__io__Output, x int) {
	if x < -0x8000 || x >= 0x8000 {
		hxrt.Throw(haxe__io__Error_Overflow)
		return
	}
	self.writeUInt16(x & 0xFFFF)
}

func haxe__io__output_writeUInt16(self haxe__io__Output, x int) {
	if x < 0 || x >= 0x10000 {
		hxrt.Throw(haxe__io__Error_Overflow)
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
		hxrt.Throw(haxe__io__Error_Overflow)
		return
	}
	self.writeUInt24(x & 0xFFFFFF)
}

func haxe__io__output_writeUInt24(self haxe__io__Output, x int) {
	if x < 0 || x >= 0x1000000 {
		hxrt.Throw(haxe__io__Error_Overflow)
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
	resolved := 4096
	if len(bufsize) > 0 {
		resolved = bufsize[0]
	}
	haxe__io__GoIoHelpers_outputWriteInput(self, i, resolved)
}

func haxe__io__output_writeString(self haxe__io__Output, s *string, encoding ...*haxe__io__Encoding) {
	var resolved *haxe__io__Encoding
	if len(encoding) > 0 {
		resolved = encoding[0]
	}
	haxe__io__GoIoHelpers_outputWriteString(self, s, resolved)
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
	reader    *bufio.Reader
	socket    *sys__net__Socket
	bigEndian bool
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

func (self *sys__net__Socket) hxrt__socket_conn() net.Conn {
	if self == nil {
		return nil
	}
	return self.conn
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
	_ = sys__net__Socket_select_(hxrt.NewArray(self), hxrt.NewArray(), hxrt.NewArray(), -1)
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

func sys__net__Socket_select_(read *hxrt.Array, write *hxrt.Array, others *hxrt.Array, timeout ...float64) map[string]any {
	if read == nil {
		read = hxrt.NewArray()
	}
	if write == nil {
		write = hxrt.NewArray()
	}
	if others == nil {
		others = hxrt.NewArray()
	}
	effectiveTimeout := -1.0
	if len(timeout) > 0 {
		effectiveTimeout = timeout[0]
	}
	readyRead := hxrt.NewArray()
	readyWrite := hxrt.NewArray()
	readyOther := hxrt.NewArray()
	for _, rawSocket := range read.Values() {
		socket, ok := rawSocket.(*sys__net__Socket)
		if !ok || socket == nil || socket.conn == nil || socket.input == nil || socket.input.reader == nil {
			continue
		}
		reader := socket.input.reader
		if reader.Buffered() > 0 {
			readyRead.Push(socket)
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
			readyRead.Push(socket)
			continue
		}
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			continue
		}
		readyOther.Push(socket)
	}
	for _, rawSocket := range write.Values() {
		socket, ok := rawSocket.(*sys__net__Socket)
		if !ok || socket == nil || socket.conn == nil {
			continue
		}
		readyWrite.Push(socket)
	}
	for _, rawSocket := range others.Values() {
		socket, ok := rawSocket.(*sys__net__Socket)
		if !ok || socket == nil {
			continue
		}
		readyOther.Push(socket)
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

func (self *sys__net__SocketInput) readByte() int {
	if self == nil || self.reader == nil {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	if self.socket != nil {
		self.socket.hxrt__socket_applyConnDeadline()
	}
	value, err := self.reader.ReadByte()
	if err != nil {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	return int(value)
}

func (self *sys__net__SocketInput) readBytes(buf *haxe__io__Bytes, pos int, len int) int {
	if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if len == 0 {
		return 0
	}
	if self == nil || self.reader == nil {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	if self.socket != nil {
		self.socket.hxrt__socket_applyConnDeadline()
	}
	tmp := make([]byte, len)
	read, err := self.reader.Read(tmp)
	if err != nil && read == 0 {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	for i := 0; i < read; i++ {
		buf.b[pos+i] = int(tmp[i])
	}
	return read
}

func (self *sys__net__SocketInput) readAll(bufsize ...int) *haxe__io__Bytes {
	return haxe__io__input_readAll(self, bufsize...)
}

func (self *sys__net__SocketInput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.bigEndian
}

func (self *sys__net__SocketInput) set_bigEndian(e bool) bool {
	if self != nil {
		self.bigEndian = e
	}
	return e
}

func (self *sys__net__SocketInput) close() {
	if self != nil && self.socket != nil {
		self.socket.close()
	}
}

func (self *sys__net__SocketInput) readFullBytes(s *haxe__io__Bytes, pos int, len int) {
	haxe__io__input_readFullBytes(self, s, pos, len)
}

func (self *sys__net__SocketInput) read(nbytes int) *haxe__io__Bytes {
	return haxe__io__input_read(self, nbytes)
}

func (self *sys__net__SocketInput) readUntil(end int) *string {
	return haxe__io__input_readUntil(self, end)
}

func (self *sys__net__SocketInput) readFloat() float64 {
	return haxe__io__input_readFloat(self)
}

func (self *sys__net__SocketInput) readDouble() float64 {
	return haxe__io__input_readDouble(self)
}

func (self *sys__net__SocketInput) readInt8() int {
	return haxe__io__input_readInt8(self)
}

func (self *sys__net__SocketInput) readInt16() int {
	return haxe__io__input_readInt16(self)
}

func (self *sys__net__SocketInput) readUInt16() int {
	return haxe__io__input_readUInt16(self)
}

func (self *sys__net__SocketInput) readInt24() int {
	return haxe__io__input_readInt24(self)
}

func (self *sys__net__SocketInput) readUInt24() int {
	return haxe__io__input_readUInt24(self)
}

func (self *sys__net__SocketInput) readInt32() int {
	return haxe__io__input_readInt32(self)
}

func (self *sys__net__SocketInput) readString(len int, encoding ...*haxe__io__Encoding) *string {
	return haxe__io__input_readString(self, len, encoding...)
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
