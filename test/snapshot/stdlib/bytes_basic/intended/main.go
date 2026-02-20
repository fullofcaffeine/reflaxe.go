package main

import "snapshot/hxrt"

func main() {
	bytes := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("abc"))
	_ = bytes
	bytes.b[1] = 122
	hxrt.Println(bytes.toString())
	hxrt.Println(bytes.length)
	buffer := New_haxe__io__BytesBuffer()
	_ = buffer
	var encoding *haxe__io__Encoding = nil
	_ = encoding
	src := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("Hi"), encoding)
	_ = src
	b1 := buffer.b
	_ = b1
	b2 := src.b
	_ = b2
	_g := 0
	_ = _g
	_g1 := src.length
	_ = _g1
	for _g < _g1 {
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_1
		_ = i
		buffer.b = append(buffer.b, (b2[i] & 255))
	}
	buffer.b = append(buffer.b, (33 & 255))
	out := buffer.getBytes()
	_ = out
	hxrt.Println(out.toString())
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

func (self *haxe__io__Bytes) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToString(self.b)
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
		hxrt.Throw(haxe__io__Error_OutsideBounds)
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
