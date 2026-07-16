package main

import "snapshot/hxrt"

func main() {
	var s any = hxrt.StringFromLiteral("a😀bé")
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("length="), hxrt.StringLengthStringPtr(hxrt.StdString(func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(s)))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("char1="), _UnicodeString__UnicodeString_Impl__charAt(hxrt.StdString(s), 1)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("code1="), _UnicodeString__UnicodeString_Impl__charCodeAt(hxrt.StdString(s), 1)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substr="), _UnicodeString__UnicodeString_Impl__substr(hxrt.StdString(s), 1, 2)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring="), _UnicodeString__UnicodeString_Impl__substring(hxrt.StdString(s), 1, 3)))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring.swap="), _UnicodeString__UnicodeString_Impl__substring(hxrt.StdString(s), 3, 1)))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring.neg="), _UnicodeString__UnicodeString_Impl__substring(hxrt.StdString(s), -2, 2)))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring.omit="), _UnicodeString__UnicodeString_Impl__substring(hxrt.StdString(s), 2, nil)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substr.neglen="), _UnicodeString__UnicodeString_Impl__substr(hxrt.StdString(s), 1, -1)))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substr.negpos="), _UnicodeString__UnicodeString_Impl__substr(hxrt.StdString(s), -2, 2)))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("index="), _UnicodeString__UnicodeString_Impl__indexOf(hxrt.StdString(s), hxrt.StringFromLiteral("bé"), nil)))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("index.empty="), _UnicodeString__UnicodeString_Impl__indexOf(hxrt.StdString(s), hxrt.StringFromLiteral(""), nil)))
	hxrt.Println(v_11)
	var v_12 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("index.startNeg="), _UnicodeString__UnicodeString_Impl__indexOf(hxrt.StdString(s), hxrt.StringFromLiteral("bé"), -2)))
	hxrt.Println(v_12)
	var v_13 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("last="), _UnicodeString__UnicodeString_Impl__lastIndexOf(hxrt.StdString(s), hxrt.StringFromLiteral("a"), nil)))
	hxrt.Println(v_13)
	var v_14 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("last.empty="), _UnicodeString__UnicodeString_Impl__lastIndexOf(hxrt.StdString(s), hxrt.StringFromLiteral(""), nil)))
	hxrt.Println(v_14)
	var v_15 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("last.start="), _UnicodeString__UnicodeString_Impl__lastIndexOf(hxrt.StdString(s), hxrt.StringFromLiteral("bé"), 2)))
	hxrt.Println(v_15)
	var left any = hxrt.StringFromLiteral("a😀")
	var right any = hxrt.StringFromLiteral("bé")
	var v_16 any = func() any {
		var a any = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("operators="), hxrt.StdString(_UnicodeString__UnicodeString_Impl__lessThan(left, right))), hxrt.StringFromLiteral("|")), hxrt.StdString(_UnicodeString__UnicodeString_Impl__lessThanOrEqual(left, right))), hxrt.StringFromLiteral("|")), hxrt.StdString(_UnicodeString__UnicodeString_Impl__greaterThan(right, left))), hxrt.StringFromLiteral("|")), hxrt.StdString(_UnicodeString__UnicodeString_Impl__greaterThanOrEqual(right, left))), hxrt.StringFromLiteral("|")), hxrt.StdString(hxrt.StringEqualStringPtr(func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(left), func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(left)))), hxrt.StringFromLiteral("|")), hxrt.StdString(!hxrt.StringEqualStringPtr(func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(left), func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
		}(right)))), hxrt.StringFromLiteral("|"))
		return hxrt.StringConcatStringPtr(func(hx_value_11 any) *string {
			if hx_value_11 == nil {
				var hx_zero_12 *string
				return hx_zero_12
			}
			return hx_value_11.(*string)
		}(a), func(hx_value_17 any) *string {
			if hx_value_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_value_17.(*string)
		}(any(hxrt.StringConcatStringPtr(func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(left), func(hx_value_15 any) *string {
			if hx_value_15 == nil {
				var hx_zero_16 *string
				return hx_zero_16
			}
			return hx_value_15.(*string)
		}(right)))))
	}()
	hxrt.Println(v_16)
	hxrt.Println(any(hxrt.StringConcatStringPtr(func(hx_value_25 any) *string {
		if hx_value_25 == nil {
			var hx_zero_26 *string
			return hx_zero_26
		}
		return hx_value_25.(*string)
	}(any(hxrt.StringConcatStringPtr(func(hx_value_23 any) *string {
		if hx_value_23 == nil {
			var hx_zero_24 *string
			return hx_zero_24
		}
		return hx_value_23.(*string)
	}(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mixed="), func(hx_value_21 any) *string {
		if hx_value_21 == nil {
			var hx_zero_22 *string
			return hx_zero_22
		}
		return hx_value_21.(*string)
	}(any(hxrt.StringConcatStringPtr(func(hx_value_19 any) *string {
		if hx_value_19 == nil {
			var hx_zero_20 *string
			return hx_zero_20
		}
		return hx_value_19.(*string)
	}(left), hxrt.StringFromLiteral("x"))))))), hxrt.StringFromLiteral("|")))), func(hx_value_29 any) *string {
		if hx_value_29 == nil {
			var hx_zero_30 *string
			return hx_zero_30
		}
		return hx_value_29.(*string)
	}(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("x"), func(hx_value_27 any) *string {
		if hx_value_27 == nil {
			var hx_zero_28 *string
			return hx_zero_28
		}
		return hx_value_27.(*string)
	}(left)))))))
	var assigned any = left
	assigned = any(hxrt.StringConcatStringPtr(func(hx_value_31 any) *string {
		if hx_value_31 == nil {
			var hx_zero_32 *string
			return hx_zero_32
		}
		return hx_value_31.(*string)
	}(assigned), func(hx_value_33 any) *string {
		if hx_value_33 == nil {
			var hx_zero_34 *string
			return hx_zero_34
		}
		return hx_value_33.(*string)
	}(right)))
	assigned = any(hxrt.StringConcatStringPtr(func(hx_value_35 any) *string {
		if hx_value_35 == nil {
			var hx_zero_36 *string
			return hx_zero_36
		}
		return hx_value_35.(*string)
	}(assigned), hxrt.StringFromLiteral("x")))
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("assigned="), func(hx_value_37 any) *string {
		if hx_value_37 == nil {
			var hx_zero_38 *string
			return hx_zero_38
		}
		return hx_value_37.(*string)
	}(assigned))))
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("valid.utf8="), hxrt.StdString(_UnicodeString__UnicodeString_Impl__validate(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("ok")), haxe__io__Encoding_UTF8))))
	hxrt.Println(v_17)
	var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("valid.invalid="), hxrt.StdString(_UnicodeString__UnicodeString_Impl__validate(haxe__io__Bytes_ofHex(hxrt.StringFromLiteral("ff")), haxe__io__Encoding_UTF8))))
	hxrt.Println(v_18)
	hxrt.TryCatch(func() {
		_UnicodeString__UnicodeString_Impl__validate(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("ok")), haxe__io__Encoding_RawNative)
		hxrt.Println(any(hxrt.StringFromLiteral("valid.raw=ok")))
	}, func(hx_caught_39 any) {
		error := hx_caught_39
		var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("valid.raw="), hxrt.StdString(error)))
		hxrt.Println(v_19)
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
