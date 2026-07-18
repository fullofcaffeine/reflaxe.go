package main

import (
	"math"
	"reflect"
	"snapshot/hxrt"
)

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

var CERT_PEM *string = hxrt.StringFromLiteral("-----BEGIN CERTIFICATE-----\nMIIDETCCAfmgAwIBAgIUVftG+IyWIEqIQE+K9ztyXCDYIcswDQYJKoZIhvcNAQEL\nBQAwGDEWMBQGA1UEAwwNZGVmYXVsdC5sb2NhbDAeFw0yNjAzMDcwNjA1NTBaFw0y\nNzAzMDcwNjA1NTBaMBgxFjAUBgNVBAMMDWRlZmF1bHQubG9jYWwwggEiMA0GCSqG\nSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCktatoj06k/dqldeSzjPUnPxeMC/WpprMz\n7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqaQDkLF0Zy+Q2K+GWLRCux\n7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVKT3jjcc3ivv6j/kOsbE7j\njELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLeoT+UU7O9VKsEIQHCGt5P\np7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubpnx+I/wUmLK4Aqh2VuuZz\nmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi9/IqJ3NpAgMBAAGjUzBR\nMB0GA1UdDgQWBBR7eOzK13X2mO5Kpc9MZOFGJ+MdlDAfBgNVHSMEGDAWgBR7eOzK\n13X2mO5Kpc9MZOFGJ+MdlDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUA\nA4IBAQAxilSFs9ZEEI5RFGVpWUTCH8Iuewc4K0JH2LbGqExgLu5MQJF5xsGoEBRe\nmDB3duaMDwnoDY19hdoClIz/Z5IO0wPEcny3hTb582W8+cRDiCQx0Qz5g2NpsfGH\nwZaGyVzSS2xbt0F6TPRQarXtzV97J067j7bRuMbD4fFYb6iqF1GbaaQsP89sA9b4\nx8yT0TGREdx2Dw29lokd4H16b3O9aYX274qZk9qpgE3oYpCeNsjHCSdTWRDw2+pg\ncKP2FsxcbB4qChjPXhhE5I3zDMTI1/4r4wGbnVIud6TmibElpAU6hwbpo1onyxKw\nVCpI190WWaS+Cz+9vGVNBSsRYjgN\n-----END CERTIFICATE-----\n")

var KEY_PEM *string = hxrt.StringFromLiteral("-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCktatoj06k/dql\ndeSzjPUnPxeMC/WpprMz7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqa\nQDkLF0Zy+Q2K+GWLRCux7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVK\nT3jjcc3ivv6j/kOsbE7jjELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLe\noT+UU7O9VKsEIQHCGt5Pp7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubp\nnx+I/wUmLK4Aqh2VuuZzmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi\n9/IqJ3NpAgMBAAECggEAECDG2jvwGjlOO9XtsUQr7C4iHuE76qMLWJo5vI5GfteQ\ndZUHVuA1Fh+NC+3z2N+uHIukuWz9G+wWGuEPhN3AVPk8oX89qDOiaK90XV6Cwt3s\n0ZU6gW/nz+qHmvdXru60nCrLephnEr+cP0TFZFYMMDf+BK5cz4kid2cQUmItbKBy\nE9k1vHvyyLONK18y72IXsMEA572V+AhtvWVar5qUJiY2+dWtxmWh9dzUvWCUJ/Gj\niqyrJdNxHJfXdteT8QpJEndsX19MWNAMZY+e10XOVzs2Rd4RdnlLwUE37HbPz73V\nC1z2h6tshSOYi9ZwlNdYEjfWbuiQimwgljwc9/NagQKBgQDg4JkcchEOmntEXPz0\n/HTEmAd7Sl75V0ozj7xPTYcaSvdoEtLEbUJbGoVfLKJM17jLuEjpisS1IHa29kkG\n2PACbbIRzOLlCA2Spkzm9R2Rni9qbBCwYbnYbdhV+hshcnr6LwUt8KBAFHBCk2lW\nMfiCsBE1KzBeye1JFxIZ9+dyKQKBgQC7gVSvPDIqPRYP8l8sLbONMHTNhbZpMh23\npdD1i0YpYTIuvrxZMXwY6gt71SJcHIwyh/6HEg8mAIMZutdRzpZc4NWQw98yG9Zo\nkpY7S2uLZ8YnARps+NprdWiARkY3PJjIiK7tieNo+dQtfBeuL3Ox7s/XR0nDUXMP\nxkuAnByfQQKBgQCks9twehsEFyExcOnUhRMA6liQdGgbN1OhcCT78EyDdWS/VQoJ\n0/xFvabxjj9RCK7QhqjgZEKuZpiMaNYTrdAb9zv0zZthJATM5ABvKBgAD1urFnsi\ntHDpk4pfbk9wr+hiVQ32F8dHJ7EREeaUuwTIsyvnRTqoMj0Yy0z2uBtMAQKBgQC2\nEnG+7z7vEP4ZYgrUhVQyp3jkERD9uTJuH892f1UT3VOzXHbcTVbpgmrARkflFbt1\nXeTkF78p8ZlcJLfssiQD8DaxKeHTcICUbrL+xM+bQJuDSGj2o/bEHe/pj1OjU24w\nW7kw45I1X1KPEE6WT3GSuAiOTKTtymtmR/EM44pPgQKBgQC52gWDbpJb1UREu9Sb\naIy1zgqGhSxg1DLChu4vz3t5ZBhWoiw6wG7vQDrtrn9Jl61mVtog0rSTSiIj726W\nLDxSrTdYBhcgkbJf/fw1N6Uk+FfJcVS/zuoJwogkiP6tmircL0qA9sYR4RZo2ubE\nAg6A+vHIw54BclYKSczsXnCe+A==\n-----END PRIVATE KEY-----\n")

func main() {
	defer hxrt.ThreadWaitForAll()
	cert := sys__ssl__Certificate_fromString(CERT_PEM)
	key := sys__ssl__Key_readPEM(KEY_PEM, false, nil)
	server := New_sys__ssl__Socket()
	server.setCertificate(cert, key)
	server.bind(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), 0)
	server.listen(1)
	bound := server.host()
	acceptedSockets := New_sys__thread__Deque()
	sys__thread__Thread_create(func() {
		acceptedSockets.add(server.accept())
	})
	client := New_sys__ssl__Socket()
	client.verifyCert = false
	client.setCA(cert)
	client.setHostname(hxrt.StringFromLiteral("localhost"))
	client.connect(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), func(hx_obj_1 map[string]any) int {
		hx_field_2 := hx_obj_1["port"]
		if hx_field_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_field_2.(int)
	}(bound))
	client.handshake()
	accepted := func(hx_value_4 any) *sys__net__Socket {
		if hx_value_4 == nil {
			var hx_zero_5 *sys__net__Socket
			return hx_zero_5
		}
		return hx_value_4.(*sys__net__Socket)
	}(acceptedSockets.pop(true))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("server.accept.ssl="), hxrt.StdString(func(hx_value *sys__net__Socket) bool {
		if hx_value == nil {
			return false
		}
		_, ok := hx_value.__hx_this.(*sys__ssl__Socket)
		return ok
	}(accepted))))
	hxrt.Println(v)
	client.output.writeString(hxrt.StringFromLiteral("ping\n"))
	client.output.flush()
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("server.read="), accepted.input.readLine()))
	hxrt.Println(v_1)
	accepted.output.writeString(hxrt.StringFromLiteral("pong\n"))
	accepted.output.flush()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("client.read="), client.input.readLine()))
	hxrt.Println(v_2)
	peer := client.peerCertificate()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("peer.cn="), peer.get_commonName()))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("peer.issuer.cn="), peer.issuer(hxrt.StringFromLiteral("CN"))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("peer.next="), hxrt.StdString((peer.next() == nil))))
	hxrt.Println(v_5)
	client.close()
	accepted.close()
	server.close()
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

type Std struct {
}

type Type struct {
}

type Reflect struct {
}

func Reflect_compare(a any, b any) int {
	toFloat := func(value any) (float64, bool) {
		switch v := value.(type) {
		case int:
			return float64(v), true
		case int8:
			return float64(v), true
		case int16:
			return float64(v), true
		case int32:
			return float64(v), true
		case int64:
			return float64(v), true
		case uint:
			return float64(v), true
		case uint8:
			return float64(v), true
		case uint16:
			return float64(v), true
		case uint32:
			return float64(v), true
		case uint64:
			return float64(v), true
		case float32:
			return float64(v), true
		case float64:
			return v, true
		default:
			return 0, false
		}
	}
	if af, ok := toFloat(a); ok {
		if bf, okB := toFloat(b); okB {
			if af < bf {
				return -1
			}
			if af > bf {
				return 1
			}
			return 0
		}
	}
	aStr := *hxrt.StdString(a)
	bStr := *hxrt.StdString(b)
	if aStr < bStr {
		return -1
	}
	if aStr > bStr {
		return 1
	}
	return 0
}

func Reflect_compareMethods(a any, b any) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	if !av.IsValid() || !bv.IsValid() {
		return !av.IsValid() && !bv.IsValid()
	}
	if av.Kind() == reflect.Func && bv.Kind() == reflect.Func {
		if av.IsNil() || bv.IsNil() {
			return av.IsNil() && bv.IsNil()
		}
		return av.Pointer() == bv.Pointer()
	}
	return reflect.DeepEqual(a, b)
}

func Reflect_field(obj any, field *string) any {
	if obj == nil {
		return nil
	}
	key := *hxrt.StdString(field)
	if metadataValue, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return metadataValue
	}
	switch value := obj.(type) {
	case map[string]any:
		return value[key]
	case map[any]any:
		return value[key]
	case *map[string]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	case *map[any]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return nil
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if fieldValue := rv.FieldByName(key); fieldValue.IsValid() && fieldValue.CanInterface() {
			return fieldValue.Interface()
		}
	}
	method := reflect.ValueOf(obj).MethodByName(key)
	if method.IsValid() {
		return method.Interface()
	}
	return nil
}

func Reflect_hasField(obj any, field *string) bool {
	if obj == nil {
		return false
	}
	key := *hxrt.StdString(field)
	if _, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return true
	}
	switch value := obj.(type) {
	case map[string]any:
		_, ok := value[key]
		return ok
	case map[any]any:
		_, ok := value[key]
		return ok
	case *map[string]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	case *map[any]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if rv.FieldByName(key).IsValid() {
			return true
		}
	}
	return reflect.ValueOf(obj).MethodByName(key).IsValid()
}

func Reflect_setField(obj any, field *string, value any) {
	if obj == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	key := *hxrt.StdString(field)
	switch target := obj.(type) {
	case map[string]any:
		target[key] = value
		return
	case map[any]any:
		target[key] = value
		return
	case *map[string]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	case *map[any]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() || rv.Kind() != reflect.Pointer {
		return
	}
	if rv.IsNil() {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	fieldValue := rv.FieldByName(key)
	if !fieldValue.IsValid() || !fieldValue.CanSet() {
		return
	}
	if value == nil {
		fieldValue.Set(reflect.Zero(fieldValue.Type()))
		return
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(fieldValue.Type()) {
		fieldValue.Set(incoming)
		return
	}
	if incoming.Type().ConvertibleTo(fieldValue.Type()) {
		fieldValue.Set(incoming.Convert(fieldValue.Type()))
		return
	}
	if fieldValue.Kind() == reflect.Interface {
		fieldValue.Set(incoming)
	}
}

func hxrt_haxeBytesToRaw(value *haxe__io__Bytes) []byte {
	if value == nil {
		return []byte{}
	}
	if value.__hx_rawValid && len(value.__hx_raw) == len(value.b) {
		return value.__hx_raw
	}
	raw := make([]byte, len(value.b))
	for i := 0; i < len(value.b); i++ {
		raw[i] = byte(value.b[i])
	}
	value.__hx_raw = raw
	value.__hx_rawValid = true
	return raw
}

func hxrt_rawToHaxeBytes(value []byte) *haxe__io__Bytes {
	converted := make([]int, len(value))
	for i := 0; i < len(value); i++ {
		converted[i] = int(value[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: value, __hx_rawValid: true}
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}

type ValueType struct {
	tag    int
	params []any
}

var ValueType_TNull *ValueType = &ValueType{tag: 0, params: []any{}}

var ValueType_TInt *ValueType = &ValueType{tag: 1, params: []any{}}

var ValueType_TFloat *ValueType = &ValueType{tag: 2, params: []any{}}

var ValueType_TBool *ValueType = &ValueType{tag: 3, params: []any{}}

var ValueType_TObject *ValueType = &ValueType{tag: 4, params: []any{}}

var ValueType_TFunction *ValueType = &ValueType{tag: 5, params: []any{}}

var ValueType_TUnknown *ValueType = &ValueType{tag: 8, params: []any{}}

func ValueType_TClass(c any) *ValueType {
	return &ValueType{tag: 6, params: []any{c}}
}

func ValueType_TEnum(e any) *ValueType {
	return &ValueType{tag: 7, params: []any{e}}
}

func hxrt_typeCallAny(callable any, args []any) (any, bool) {
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	if callable == nil {
		return nil, false
	}
	fn := reflect.ValueOf(callable)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, false
	}
	fnType := fn.Type()
	if fnType.NumIn() != len(args) {
		return nil, false
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		paramType := fnType.In(i)
		arg := args[i]
		if arg == nil {
			in[i] = reflect.Zero(paramType)
			continue
		}
		v := reflect.ValueOf(arg)
		if v.IsValid() && v.Type().AssignableTo(paramType) {
			in[i] = v
			continue
		}
		if v.IsValid() && v.Type().ConvertibleTo(paramType) {
			in[i] = v.Convert(paramType)
			continue
		}
		if paramType.Kind() == reflect.Interface && v.IsValid() {
			in[i] = v
			continue
		}
		return nil, false
	}
	out := fn.Call(in)
	if len(out) == 0 {
		return nil, true
	}
	first := out[0]
	if !first.IsValid() {
		return nil, true
	}
	result = first.Interface()
	ok = true
	return result, ok
}

func hxrt_typeArrayValues(value *hxrt.Array) []any {
	if value == nil {
		return []any{}
	}
	return value.Values()
}

func hxrt_typeResolvedClassName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeClassValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeClassValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeResolvedEnumName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeEnumValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeEnumValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeCreateClassInstance(className string, args []any) (any, bool) {
	switch className {
	case "Date":
		return hxrt_typeCallAny(New_Date, args)
	case "Main":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.ds.List":
		return hxrt_typeCallAny(New_haxe__ds__List, args)
	case "haxe.ds._List.GoListIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListIterator, args)
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListKeyValueIterator, args)
	case "haxe.io.GoIoHelpers":
		return nil, false
	case "sys.net.Host":
		return hxrt_typeCallAny(New_sys__net__Host, args)
	case "sys.net.Socket":
		return hxrt_typeCallAny(New_sys__net__Socket, args)
	case "sys.net.SocketInput":
		return hxrt_typeCallAny(New_sys__net__SocketInput, args)
	case "sys.net.SocketOutput":
		return hxrt_typeCallAny(New_sys__net__SocketOutput, args)
	case "sys.ssl.Certificate":
		return hxrt_typeCallAny(New_sys__ssl__Certificate, args)
	case "sys.ssl.Key":
		return hxrt_typeCallAny(New_sys__ssl__Key, args)
	case "sys.ssl.Socket":
		return hxrt_typeCallAny(New_sys__ssl__Socket, args)
	case "sys.thread.Deque":
		return hxrt_typeCallAny(New_sys__thread__Deque, args)
	case "sys.thread.EventLoop":
		return hxrt_typeCallAny(New_sys__thread__EventLoop, args)
	case "sys.thread.Lock":
		return hxrt_typeCallAny(New_sys__thread__Lock, args)
	case "sys.thread.Mutex":
		return hxrt_typeCallAny(New_sys__thread__Mutex, args)
	case "sys.thread.NoEventLoopException":
		return hxrt_typeCallAny(New_sys__thread__NoEventLoopException, args)
	case "sys.thread.Thread":
		return hxrt_typeCallAny(New_sys__thread__Thread, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "Date":
		return &Date{}, true
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	case "haxe.ds.List":
		return &haxe__ds__List{}, true
	case "haxe.ds._List.GoListIterator":
		return &haxe__ds___List__GoListIterator{}, true
	case "haxe.ds._List.GoListKeyValueIterator":
		return &haxe__ds___List__GoListKeyValueIterator{}, true
	case "sys.net.Host":
		return &sys__net__Host{}, true
	case "sys.net.Socket":
		return &sys__net__Socket{}, true
	case "sys.net.SocketInput":
		return &sys__net__SocketInput{}, true
	case "sys.net.SocketOutput":
		return &sys__net__SocketOutput{}, true
	case "sys.ssl.Certificate":
		return &sys__ssl__Certificate{}, true
	case "sys.ssl.Key":
		return &sys__ssl__Key{}, true
	case "sys.ssl.Socket":
		return &sys__ssl__Socket{}, true
	case "sys.thread.Deque":
		return &sys__thread__Deque{}, true
	case "sys.thread.EventLoop":
		return &sys__thread__EventLoop{}, true
	case "sys.thread.Lock":
		return &sys__thread__Lock{}, true
	case "sys.thread.Mutex":
		return &sys__thread__Mutex{}, true
	case "sys.thread.NoEventLoopException":
		return &sys__thread__NoEventLoopException{}, true
	case "sys.thread.Thread":
		return &sys__thread__Thread{}, true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "ValueType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TNull, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TInt, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFloat, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TBool, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TObject, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFunction, true
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TClass, args)
			case 7:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TEnum, args)
			case 8:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TUnknown, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TNull":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TNull, true
		case "TInt":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TInt, true
		case "TFloat":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFloat, true
		case "TBool":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TBool, true
		case "TObject":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TObject, true
		case "TFunction":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFunction, true
		case "TClass":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TClass, args)
		case "TEnum":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TEnum, args)
		case "TUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TUnknown, true
		default:
			return nil, false
		}
	case "sys.thread.NextEventTime":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return sys__thread__NextEventTime_Now, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return sys__thread__NextEventTime_Never, true
			case 2:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(sys__thread__NextEventTime_AnyTime, args)
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(sys__thread__NextEventTime_At, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "Now":
			if len(args) != 0 {
				return nil, false
			}
			return sys__thread__NextEventTime_Now, true
		case "Never":
			if len(args) != 0 {
				return nil, false
			}
			return sys__thread__NextEventTime_Never, true
		case "AnyTime":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(sys__thread__NextEventTime_AnyTime, args)
		case "At":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(sys__thread__NextEventTime_At, args)
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeClassValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeClassValue:
		copyValue := value
		return &copyValue
	case *hxrt.Array:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")}
	case *Date:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Date")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe__ds__List:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.List")}
	case *haxe__ds___List__GoListIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._List.GoListIterator")}
	case *haxe__ds___List__GoListKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._List.GoListKeyValueIterator")}
	case *sys__net__Host:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.Host")}
	case *sys__net__Socket:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.Socket")}
	case *sys__net__SocketInput:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.SocketInput")}
	case *sys__net__SocketOutput:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.SocketOutput")}
	case *sys__ssl__Certificate:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.ssl.Certificate")}
	case *sys__ssl__Key:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.ssl.Key")}
	case *sys__ssl__Socket:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.ssl.Socket")}
	case *sys__thread__Deque:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Deque")}
	case *sys__thread__EventLoop:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.EventLoop")}
	case *sys__thread__Lock:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Lock")}
	case *sys__thread__Mutex:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Mutex")}
	case *sys__thread__NoEventLoopException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.NoEventLoopException")}
	case *sys__thread__Thread:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Thread")}
	default:
		return nil
	}
}

func Type_getEnum(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeEnumValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeEnumValue:
		copyValue := value
		return &copyValue
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	case *sys__thread__NextEventTime:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("sys.thread.NextEventTime")}
	default:
		return nil
	}
}

func Type_getSuperClass(c any) any {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	switch className {
	case "Date":
		return nil
	case "Main":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.ds.List":
		return nil
	case "haxe.ds._List.GoListIterator":
		return nil
	case "haxe.ds._List.GoListKeyValueIterator":
		return nil
	case "haxe.io.GoIoHelpers":
		return nil
	case "sys.net.Host":
		return nil
	case "sys.net.Socket":
		return nil
	case "sys.net.SocketInput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Input")}
	case "sys.net.SocketOutput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Output")}
	case "sys.ssl.Certificate":
		return nil
	case "sys.ssl.Key":
		return nil
	case "sys.ssl.Socket":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.Socket")}
	case "sys.thread.Deque":
		return nil
	case "sys.thread.EventLoop":
		return nil
	case "sys.thread.Lock":
		return nil
	case "sys.thread.Mutex":
		return nil
	case "sys.thread.NoEventLoopException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "sys.thread.Thread":
		return nil
	default:
		return nil
	}
}

func Type_getClassName(c any) *string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(className)
}

func Type_getClassFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Date":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromMilliseconds"), hxrt.StringFromLiteral("fromString"), hxrt.StringFromLiteral("fromTime"), hxrt.StringFromLiteral("now"))
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("CERT_PEM"), hxrt.StringFromLiteral("KEY_PEM"), hxrt.StringFromLiteral("main"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("sameValue"))
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray()
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.io.GoIoHelpers":
		return hxrt.NewArray(hxrt.StringFromLiteral("bytesOutputGetBytes"), hxrt.StringFromLiteral("inputRead"), hxrt.StringFromLiteral("inputReadAll"), hxrt.StringFromLiteral("inputReadBytes"), hxrt.StringFromLiteral("inputReadFullBytes"), hxrt.StringFromLiteral("inputReadLine"), hxrt.StringFromLiteral("inputReadUntil"), hxrt.StringFromLiteral("outputWrite"), hxrt.StringFromLiteral("outputWriteBytes"), hxrt.StringFromLiteral("outputWriteFullBytes"), hxrt.StringFromLiteral("outputWriteInput"), hxrt.StringFromLiteral("outputWriteString"))
	case "sys.net.Host":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromIPv4"), hxrt.StringFromLiteral("localhost"))
	case "sys.net.Socket":
		return hxrt.NewArray(hxrt.StringFromLiteral("pick"), hxrt.StringFromLiteral("publicAddress"), hxrt.StringFromLiteral("select"))
	case "sys.net.SocketInput":
		return hxrt.NewArray(hxrt.StringFromLiteral("translateReadStatus"))
	case "sys.net.SocketOutput":
		return hxrt.NewArray(hxrt.StringFromLiteral("translateWriteStatus"))
	case "sys.ssl.Certificate":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromString"), hxrt.StringFromLiteral("loadDefaults"), hxrt.StringFromLiteral("loadFile"), hxrt.StringFromLiteral("loadPath"))
	case "sys.ssl.Key":
		return hxrt.NewArray(hxrt.StringFromLiteral("loadFile"), hxrt.StringFromLiteral("readDER"), hxrt.StringFromLiteral("readPEM"))
	case "sys.ssl.Socket":
		return hxrt.NewArray(hxrt.StringFromLiteral("DEFAULT_CA"), hxrt.StringFromLiteral("DEFAULT_VERIFY_CERT"))
	case "sys.thread.Deque":
		return hxrt.NewArray()
	case "sys.thread.EventLoop":
		return hxrt.NewArray(hxrt.StringFromLiteral("__fromHandle"))
	case "sys.thread.Lock":
		return hxrt.NewArray()
	case "sys.thread.Mutex":
		return hxrt.NewArray()
	case "sys.thread.NoEventLoopException":
		return hxrt.NewArray()
	case "sys.thread.Thread":
		return hxrt.NewArray(hxrt.StringFromLiteral("create"), hxrt.StringFromLiteral("createWithEventLoop"), hxrt.StringFromLiteral("current"), hxrt.StringFromLiteral("processEvents"), hxrt.StringFromLiteral("readMessage"), hxrt.StringFromLiteral("runWithEventLoop"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getInstanceFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Date":
		return hxrt.NewArray(hxrt.StringFromLiteral("getDate"), hxrt.StringFromLiteral("getDay"), hxrt.StringFromLiteral("getFullYear"), hxrt.StringFromLiteral("getHours"), hxrt.StringFromLiteral("getMinutes"), hxrt.StringFromLiteral("getMonth"), hxrt.StringFromLiteral("getSeconds"), hxrt.StringFromLiteral("getTime"), hxrt.StringFromLiteral("getTimezoneOffset"), hxrt.StringFromLiteral("getUTCDate"), hxrt.StringFromLiteral("getUTCDay"), hxrt.StringFromLiteral("getUTCFullYear"), hxrt.StringFromLiteral("getUTCHours"), hxrt.StringFromLiteral("getUTCMinutes"), hxrt.StringFromLiteral("getUTCMonth"), hxrt.StringFromLiteral("getUTCSeconds"), hxrt.StringFromLiteral("localParts"), hxrt.StringFromLiteral("ms"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("utcParts"))
	case "Main":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("filter"), hxrt.StringFromLiteral("first"), hxrt.StringFromLiteral("isEmpty"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("join"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("last"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.io.GoIoHelpers":
		return hxrt.NewArray()
	case "sys.net.Host":
		return hxrt.NewArray(hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("ip"), hxrt.StringFromLiteral("reverse"), hxrt.StringFromLiteral("toString"))
	case "sys.net.Socket":
		return hxrt.NewArray(hxrt.StringFromLiteral("accept"), hxrt.StringFromLiteral("bind"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("connect"), hxrt.StringFromLiteral("custom"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("input"), hxrt.StringFromLiteral("listen"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("peer"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("replaceHandle"), hxrt.StringFromLiteral("setBlocking"), hxrt.StringFromLiteral("setFastSend"), hxrt.StringFromLiteral("setTimeout"), hxrt.StringFromLiteral("shutdown"), hxrt.StringFromLiteral("waitForRead"), hxrt.StringFromLiteral("write"))
	case "sys.net.SocketInput":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("readAll"), hxrt.StringFromLiteral("readByte"), hxrt.StringFromLiteral("readBytes"), hxrt.StringFromLiteral("readDouble"), hxrt.StringFromLiteral("readFloat"), hxrt.StringFromLiteral("readInt32"), hxrt.StringFromLiteral("readLine"), hxrt.StringFromLiteral("set_bigEndian"))
	case "sys.net.SocketOutput":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("flush"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("set_bigEndian"), hxrt.StringFromLiteral("writeByte"), hxrt.StringFromLiteral("writeBytes"), hxrt.StringFromLiteral("writeDouble"), hxrt.StringFromLiteral("writeFloat"), hxrt.StringFromLiteral("writeFullBytes"), hxrt.StringFromLiteral("writeInt32"), hxrt.StringFromLiteral("writeString"))
	case "sys.ssl.Certificate":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("addDER"), hxrt.StringFromLiteral("altNames"), hxrt.StringFromLiteral("commonName"), hxrt.StringFromLiteral("get_altNames"), hxrt.StringFromLiteral("get_commonName"), hxrt.StringFromLiteral("get_notAfter"), hxrt.StringFromLiteral("get_notBefore"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("issuer"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("notAfter"), hxrt.StringFromLiteral("notBefore"), hxrt.StringFromLiteral("subject"))
	case "sys.ssl.Key":
		return hxrt.NewArray(hxrt.StringFromLiteral("handle"))
	case "sys.ssl.Socket":
		return hxrt.NewArray(hxrt.StringFromLiteral("accept"), hxrt.StringFromLiteral("addSNICertificate"), hxrt.StringFromLiteral("bind"), hxrt.StringFromLiteral("caCert"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("connect"), hxrt.StringFromLiteral("custom"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("handshake"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("hostname"), hxrt.StringFromLiteral("input"), hxrt.StringFromLiteral("listen"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("ownCert"), hxrt.StringFromLiteral("ownKey"), hxrt.StringFromLiteral("peer"), hxrt.StringFromLiteral("peerCertificate"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("replaceHandle"), hxrt.StringFromLiteral("setBlocking"), hxrt.StringFromLiteral("setCA"), hxrt.StringFromLiteral("setCertificate"), hxrt.StringFromLiteral("setFastSend"), hxrt.StringFromLiteral("setHostname"), hxrt.StringFromLiteral("setTimeout"), hxrt.StringFromLiteral("shutdown"), hxrt.StringFromLiteral("sniConfig"), hxrt.StringFromLiteral("verifyCert"), hxrt.StringFromLiteral("waitForRead"), hxrt.StringFromLiteral("write"))
	case "sys.thread.Deque":
		return hxrt.NewArray(hxrt.StringFromLiteral("__available"), hxrt.StringFromLiteral("__items"), hxrt.StringFromLiteral("__mutex"), hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"))
	case "sys.thread.EventLoop":
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"), hxrt.StringFromLiteral("cancel"), hxrt.StringFromLiteral("loop"), hxrt.StringFromLiteral("progress"), hxrt.StringFromLiteral("promise"), hxrt.StringFromLiteral("repeat"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("runPromised"), hxrt.StringFromLiteral("wait"))
	case "sys.thread.Lock":
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"), hxrt.StringFromLiteral("release"), hxrt.StringFromLiteral("wait"))
	case "sys.thread.Mutex":
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"), hxrt.StringFromLiteral("acquire"), hxrt.StringFromLiteral("release"), hxrt.StringFromLiteral("tryAcquire"))
	case "sys.thread.NoEventLoopException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "sys.thread.Thread":
		return hxrt.NewArray(hxrt.StringFromLiteral("__id"), hxrt.StringFromLiteral("events"), hxrt.StringFromLiteral("get_events"), hxrt.StringFromLiteral("sendMessage"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getEnumName(e any) *string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(enumName)
}

func Type_resolveClass(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "Date":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.List":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.GoIoHelpers":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.Host":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.Socket":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.SocketInput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.SocketOutput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.ssl.Certificate":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.ssl.Key":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.ssl.Socket":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Deque":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.EventLoop":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Lock":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Mutex":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.NoEventLoopException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Thread":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_resolveEnum(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.NextEventTime":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_createInstance(cl any, args *hxrt.Array) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, hxrt_typeArrayValues(args))
	if !ok {
		return nil
	}
	return instance
}

func Type_createEmptyInstance(cl any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassEmptyInstance(className)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEnum(e any, constr *string, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_enumConstructor(e any) *string {
	if hxrt.AnyEqualsNull(e) {
		return nil
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TNull")
		case 1:
			return hxrt.StringFromLiteral("TInt")
		case 2:
			return hxrt.StringFromLiteral("TFloat")
		case 3:
			return hxrt.StringFromLiteral("TBool")
		case 4:
			return hxrt.StringFromLiteral("TObject")
		case 5:
			return hxrt.StringFromLiteral("TFunction")
		case 6:
			return hxrt.StringFromLiteral("TClass")
		case 7:
			return hxrt.StringFromLiteral("TEnum")
		case 8:
			return hxrt.StringFromLiteral("TUnknown")
		default:
			return nil
		}
	case *sys__thread__NextEventTime:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("Now")
		case 1:
			return hxrt.StringFromLiteral("Never")
		case 2:
			return hxrt.StringFromLiteral("AnyTime")
		case 3:
			return hxrt.StringFromLiteral("At")
		default:
			return nil
		}
	default:
		return nil
	}
}

func Type_enumIndex(e any) int {
	if hxrt.AnyEqualsNull(e) {
		return -1
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	case *sys__thread__NextEventTime:
		if value == nil {
			return -1
		}
		return value.tag
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown"))
	case "sys.thread.NextEventTime":
		return hxrt.NewArray(hxrt.StringFromLiteral("Now"), hxrt.StringFromLiteral("Never"), hxrt.StringFromLiteral("AnyTime"), hxrt.StringFromLiteral("At"))
	default:
		return hxrt.NewArray()
	}
}

func Type_enumParameters(e any) *hxrt.Array {
	if hxrt.AnyEqualsNull(e) {
		return hxrt.NewArray()
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *sys__thread__NextEventTime:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	default:
		return hxrt.NewArray()
	}
}

func Type_allEnums(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown)
	case "sys.thread.NextEventTime":
		return hxrt.NewArray(sys__thread__NextEventTime_Now, sys__thread__NextEventTime_Never)
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) *ValueType {
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
	}
	switch v.(type) {
	case *hxrt__TypeClassValue, hxrt__TypeClassValue, *hxrt__TypeEnumValue, hxrt__TypeEnumValue:
		return ValueType_TObject
	}
	if enumValue := Type_getEnum(v); enumValue != nil {
		return ValueType_TEnum(enumValue)
	}
	if classValue := Type_getClass(v); classValue != nil {
		return ValueType_TClass(classValue)
	}
	switch v.(type) {
	case bool:
		return ValueType_TBool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return ValueType_TInt
	case float32, float64:
		return ValueType_TFloat
	case string, *string:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")})
	case *hxrt.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	}
	ref := reflect.ValueOf(v)
	if !ref.IsValid() {
		return ValueType_TNull
	}
	switch ref.Kind() {
	case reflect.Func:
		return ValueType_TFunction
	case reflect.Slice, reflect.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:
		return ValueType_TObject
	default:
		return ValueType_TUnknown
	}
}

func Type_enumEq(a any, b any) bool {
	return reflect.DeepEqual(a, b)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	className, ok := hxrt_typeResolvedClassName(value)
	if !ok {
		return nil, false
	}
	switch className {
	default:
		return nil, false
	}
}
