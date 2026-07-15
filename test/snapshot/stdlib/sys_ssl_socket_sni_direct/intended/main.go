package main

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"io"
	"math"
	"net"
	"os"
	"reflect"
	"snapshot/hxrt"
	"strconv"
	"strings"
	"time"
)

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

var DEFAULT_CERT_PEM *string = hxrt.StringFromLiteral("-----BEGIN CERTIFICATE-----\nMIIDETCCAfmgAwIBAgIUVftG+IyWIEqIQE+K9ztyXCDYIcswDQYJKoZIhvcNAQEL\nBQAwGDEWMBQGA1UEAwwNZGVmYXVsdC5sb2NhbDAeFw0yNjAzMDcwNjA1NTBaFw0y\nNzAzMDcwNjA1NTBaMBgxFjAUBgNVBAMMDWRlZmF1bHQubG9jYWwwggEiMA0GCSqG\nSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCktatoj06k/dqldeSzjPUnPxeMC/WpprMz\n7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqaQDkLF0Zy+Q2K+GWLRCux\n7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVKT3jjcc3ivv6j/kOsbE7j\njELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLeoT+UU7O9VKsEIQHCGt5P\np7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubpnx+I/wUmLK4Aqh2VuuZz\nmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi9/IqJ3NpAgMBAAGjUzBR\nMB0GA1UdDgQWBBR7eOzK13X2mO5Kpc9MZOFGJ+MdlDAfBgNVHSMEGDAWgBR7eOzK\n13X2mO5Kpc9MZOFGJ+MdlDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUA\nA4IBAQAxilSFs9ZEEI5RFGVpWUTCH8Iuewc4K0JH2LbGqExgLu5MQJF5xsGoEBRe\nmDB3duaMDwnoDY19hdoClIz/Z5IO0wPEcny3hTb582W8+cRDiCQx0Qz5g2NpsfGH\nwZaGyVzSS2xbt0F6TPRQarXtzV97J067j7bRuMbD4fFYb6iqF1GbaaQsP89sA9b4\nx8yT0TGREdx2Dw29lokd4H16b3O9aYX274qZk9qpgE3oYpCeNsjHCSdTWRDw2+pg\ncKP2FsxcbB4qChjPXhhE5I3zDMTI1/4r4wGbnVIud6TmibElpAU6hwbpo1onyxKw\nVCpI190WWaS+Cz+9vGVNBSsRYjgN\n-----END CERTIFICATE-----\n")

var DEFAULT_KEY_PEM *string = hxrt.StringFromLiteral("-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCktatoj06k/dql\ndeSzjPUnPxeMC/WpprMz7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqa\nQDkLF0Zy+Q2K+GWLRCux7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVK\nT3jjcc3ivv6j/kOsbE7jjELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLe\noT+UU7O9VKsEIQHCGt5Pp7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubp\nnx+I/wUmLK4Aqh2VuuZzmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi\n9/IqJ3NpAgMBAAECggEAECDG2jvwGjlOO9XtsUQr7C4iHuE76qMLWJo5vI5GfteQ\ndZUHVuA1Fh+NC+3z2N+uHIukuWz9G+wWGuEPhN3AVPk8oX89qDOiaK90XV6Cwt3s\n0ZU6gW/nz+qHmvdXru60nCrLephnEr+cP0TFZFYMMDf+BK5cz4kid2cQUmItbKBy\nE9k1vHvyyLONK18y72IXsMEA572V+AhtvWVar5qUJiY2+dWtxmWh9dzUvWCUJ/Gj\niqyrJdNxHJfXdteT8QpJEndsX19MWNAMZY+e10XOVzs2Rd4RdnlLwUE37HbPz73V\nC1z2h6tshSOYi9ZwlNdYEjfWbuiQimwgljwc9/NagQKBgQDg4JkcchEOmntEXPz0\n/HTEmAd7Sl75V0ozj7xPTYcaSvdoEtLEbUJbGoVfLKJM17jLuEjpisS1IHa29kkG\n2PACbbIRzOLlCA2Spkzm9R2Rni9qbBCwYbnYbdhV+hshcnr6LwUt8KBAFHBCk2lW\nMfiCsBE1KzBeye1JFxIZ9+dyKQKBgQC7gVSvPDIqPRYP8l8sLbONMHTNhbZpMh23\npdD1i0YpYTIuvrxZMXwY6gt71SJcHIwyh/6HEg8mAIMZutdRzpZc4NWQw98yG9Zo\nkpY7S2uLZ8YnARps+NprdWiARkY3PJjIiK7tieNo+dQtfBeuL3Ox7s/XR0nDUXMP\nxkuAnByfQQKBgQCks9twehsEFyExcOnUhRMA6liQdGgbN1OhcCT78EyDdWS/VQoJ\n0/xFvabxjj9RCK7QhqjgZEKuZpiMaNYTrdAb9zv0zZthJATM5ABvKBgAD1urFnsi\ntHDpk4pfbk9wr+hiVQ32F8dHJ7EREeaUuwTIsyvnRTqoMj0Yy0z2uBtMAQKBgQC2\nEnG+7z7vEP4ZYgrUhVQyp3jkERD9uTJuH892f1UT3VOzXHbcTVbpgmrARkflFbt1\nXeTkF78p8ZlcJLfssiQD8DaxKeHTcICUbrL+xM+bQJuDSGj2o/bEHe/pj1OjU24w\nW7kw45I1X1KPEE6WT3GSuAiOTKTtymtmR/EM44pPgQKBgQC52gWDbpJb1UREu9Sb\naIy1zgqGhSxg1DLChu4vz3t5ZBhWoiw6wG7vQDrtrn9Jl61mVtog0rSTSiIj726W\nLDxSrTdYBhcgkbJf/fw1N6Uk+FfJcVS/zuoJwogkiP6tmircL0qA9sYR4RZo2ubE\nAg6A+vHIw54BclYKSczsXnCe+A==\n-----END PRIVATE KEY-----\n")

var SNI_CERT_PEM *string = hxrt.StringFromLiteral("-----BEGIN CERTIFICATE-----\nMIIDHzCCAgegAwIBAgIUMCFv2IwHqj9Wy6ycvAibkSutMKQwDQYJKoZIhvcNAQEL\nBQAwFDESMBAGA1UEAwwJc25pLmxvY2FsMB4XDTI2MDYxMDA1MTUwMloXDTI3MDYx\nMDA1MTUwMlowFDESMBAGA1UEAwwJc25pLmxvY2FsMIIBIjANBgkqhkiG9w0BAQEF\nAAOCAQ8AMIIBCgKCAQEAsQVlch80ZQ4hg3Hb3XMPgk+XPsJuoIdE0rn78PGM/O6W\n1iQocq9hYRosp/aT6itBnpvxh+tCt9D7deoEDBzwU6CkJu7N6CGY0yWWrFkits4f\n8KiTbkWZA3qKfreUOiDkiU3DgM+b9GUyBkgUXxBxW1ql+mFAX6l3eRDK5q9dxxme\nmouBL4yQxMa1bgunsYoZQhV/0RoH0gUuDUTPr12+YEA/XyCJC82p3v50mIcURjnX\nHl26xMRv5O73mJDUqshn3Z0SMmXhHfk6JFsacc2sB7f8FCOAAFsjjhSKHEQ8OTTP\nzcXAw6fcqxF1WNMrsIfZJ4Dzjt2kOYO9KxiLR3+0dwIDAQABo2kwZzAdBgNVHQ4E\nFgQUnbhLMtV250tds1NaZBDZyIvDXfwwHwYDVR0jBBgwFoAUnbhLMtV250tds1Na\nZBDZyIvDXfwwDwYDVR0TAQH/BAUwAwEB/zAUBgNVHREEDTALgglzbmkubG9jYWww\nDQYJKoZIhvcNAQELBQADggEBADHJ3F34+3iZrJNgxuFbr5IOsm4GDCZxd8g0BmWD\n+JxabwGYMk+xGOL/dn7jPwhodHcFbJXrYzT2giGMG9EdLLExqwnpQb7Z+uqga4Ej\nfyQVZZ61Np8wSGJQ7C03Yu6A0jYljW8KWM3UeiX6ee/d8/JXK8Nts8qsKXe6SFJO\nF1vblGLBZyb/DZ4GKD5cp3RZdCQxwlvLmbYKrKDwySJYwWQEx8bD5UUMYPr0Y593\nn6ZbEIWl+17PkP0kQRloeGUOxogZaI9jM/fGalKf52SbipDQV4P8NXqsdW0Im/LA\nCW6HU+6ozncjLj3SyLuwLvd54+YEPhv/cuFgP6/WWqyCx0A=\n-----END CERTIFICATE-----\n")

var SNI_KEY_PEM *string = hxrt.StringFromLiteral("-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCxBWVyHzRlDiGD\ncdvdcw+CT5c+wm6gh0TSufvw8Yz87pbWJChyr2FhGiyn9pPqK0Gem/GH60K30Pt1\n6gQMHPBToKQm7s3oIZjTJZasWSK2zh/wqJNuRZkDeop+t5Q6IOSJTcOAz5v0ZTIG\nSBRfEHFbWqX6YUBfqXd5EMrmr13HGZ6ai4EvjJDExrVuC6exihlCFX/RGgfSBS4N\nRM+vXb5gQD9fIIkLzane/nSYhxRGOdceXbrExG/k7veYkNSqyGfdnRIyZeEd+Tok\nWxpxzawHt/wUI4AAWyOOFIocRDw5NM/NxcDDp9yrEXVY0yuwh9kngPOO3aQ5g70r\nGItHf7R3AgMBAAECggEANjMyxmhrgG19MWPRL9Kk8v8vjdW2TYxdNDAhxboPsvnS\nUSqs/8BXDoYXGi5TR6WK5+dTYoxT1zgzZf0K1DKgGtrap9kCTorK4gtmQMrh6Brg\niKz0xxSkLv58HSRUTB/6GVgn/e6TD5dUY7v6EMlWC+SLYUgZj7Cxle3gUhVrnyPk\noaEOMDU8w9iiPoCK0E60DReKAmIcXigGqMNlkFjADjfOtr5GDpercL+mvwDmfLmB\ndzYcXuHM3fm1pgoDXN5zrpO7mFSp0VG/O697nZzY18A25AMMNiBUjuk7/u4KJEcA\nMmOxVAU6yjHAKO+YRS98pKsxkvZKNpj69lXtBv6agQKBgQDVnTDnIudJwbaAvfBf\n7ZZ3l46Ojm2s3qDR0Sp+W0tTc5cVhJiJaShzpnYKPbZP6R8EnDrdtldw8sYRRB6R\nPfM3/DX87z793+pKB5WDJNzevO81ecpMhiYgUxlKzzkz7nfDZC8tBQ3YLLhUX5JL\nv7REeUskkmW+xq9e9+ecK/RpVwKBgQDUJWrvJ2tEf2JJ+qdYZBHJXw5+ajIZeUs+\nCh3cNtOrUGDW/y9SKvrRHzUP0PvRle21O4oUimEb1SqdNG2nFB3f4BVUB8X6DcMz\nM16kyNBdk4q2k3xITw2OdErlYZU2WAsLUrOxurDK9ng6YO1fJwz3zrGiNDjbdebY\n1m4NbyN54QKBgQCE0akvfh9LV/wPDoqgSszs7TpBX0PIYeCitShzynYKnGuLgJeL\nkOwLBKyOb5KlGzEjH7TmWFMEMp9+6tkKu/c3j0VOUL/dANXfU9nd4hTHFbiyiliD\nvkGEhcbLIg/SP2sN/YPrvSG/kQbHx2jiWn9OuBBF3BURSt6N8Rx8mUPuHwKBgQCG\nAqj9L8Jz/5/gKaVCkdwmf5SRSJYjP1rHcu6P6FZntpul1IdY+Wt9ZKBJQHOCXppN\nTLIZ7ZwQT+Teb3sA+xUwEcaHUW2/Wqg/FKkpoOz237fVQ29T4hQnM9EH+0+dh5pa\nacC3eb4qR+2EuyvXWry3YWsWkrSD9YOA4FuewuD/IQKBgDUet4ixooKltLLhW6RT\nB+8+u2hKmlctueypoiB9+yiCt/oEiZEHYX9IwlGaax6bpDg6vO77Ay7JwGaggM+K\nhhIsqAuszWLO7bdjdS4JZ1YshP5YusdPs3m9OGeTjnO/+goL+tECrZmzupwg4BTA\nKOG6OBx2h7HDGnJcWJNLuBbJ\n-----END PRIVATE KEY-----\n")

func connectAndPrint(port int, hostname *string) {
	client := New_sys__ssl__Socket()
	client.verifyCert = false
	client.setHostname(hostname)
	client.connect(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), port)
	client.handshake()
	peer := client.peerCertificate()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hostname, hxrt.StringFromLiteral("=")), peer.get_commonName()))
	hxrt.Println(v)
	client.close()
}

func main() {
	defaultCert := sys__ssl__Certificate_fromString(DEFAULT_CERT_PEM)
	defaultKey := sys__ssl__Key_readPEM(DEFAULT_KEY_PEM, false, nil)
	sniCert := sys__ssl__Certificate_fromString(SNI_CERT_PEM)
	sniKey := sys__ssl__Key_readPEM(SNI_KEY_PEM, false, nil)
	server := New_sys__ssl__Socket()
	server.setCertificate(defaultCert, defaultKey)
	server.addSNICertificate(func(name *string) bool {
		return hxrt.StringEqualStringPtr(name, hxrt.StringFromLiteral("sni.local"))
	}, sniCert, sniKey)
	server.bind(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), 0)
	server.listen(2)
	bound := server.host()
	connectAndPrint(func(hx_obj_1 map[string]any) int {
		hx_field_2 := hx_obj_1["port"]
		if hx_field_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_field_2.(int)
	}(bound), hxrt.StringFromLiteral("default.local"))
	connectAndPrint(func(hx_obj_4 map[string]any) int {
		hx_field_5 := hx_obj_4["port"]
		if hx_field_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_field_5.(int)
	}(bound), hxrt.StringFromLiteral("sni.local"))
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

func _UnicodeString__UnicodeString_Impl__get_length(value any) int {
	return len([]rune(*hxrt.StdString(value)))
}

func _UnicodeString__UnicodeString_Impl__charAt(value any, index int) *string {
	if index < 0 {
		return hxrt.StringFromLiteral("")
	}
	runes := []rune(*hxrt.StdString(value))
	if index >= len(runes) {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(string(runes[index]))
}

func _UnicodeString__UnicodeString_Impl__charCodeAt(value any, index int) any {
	if index < 0 {
		return nil
	}
	runes := []rune(*hxrt.StdString(value))
	if index >= len(runes) {
		return nil
	}
	return int(runes[index])
}

func _UnicodeString__UnicodeString_Impl__substring(value any, startIndex int, endIndex ...int) *string {
	runes := []rune(*hxrt.StdString(value))
	end := len(runes)
	if len(endIndex) > 0 {
		end = endIndex[0]
	}
	if startIndex < 0 {
		startIndex = 0
	}
	if end < 0 {
		end = 0
	}
	if startIndex == end {
		return hxrt.StringFromLiteral("")
	}
	if startIndex > end {
		startIndex, end = end, startIndex
	}
	if startIndex > len(runes) {
		return hxrt.StringFromLiteral("")
	}
	if end > len(runes) {
		end = len(runes)
	}
	return hxrt.StringFromLiteral(string(runes[startIndex:end]))
}

func _UnicodeString__UnicodeString_Impl__substr(value any, pos int, lengthArgs ...int) *string {
	runes := []rune(*hxrt.StdString(value))
	unicodeLength := len(runes)
	if pos < 0 {
		pos = unicodeLength + pos
		if pos < 0 {
			pos = 0
		}
	}
	if pos > unicodeLength {
		return hxrt.StringFromLiteral("")
	}
	if len(lengthArgs) == 0 {
		return hxrt.StringFromLiteral(string(runes[pos:]))
	}
	lengthValue := lengthArgs[0]
	end := unicodeLength
	if lengthValue < 0 {
		end = unicodeLength + lengthValue
	} else {
		end = pos + lengthValue
	}
	if end < pos {
		return hxrt.StringFromLiteral("")
	}
	if end > unicodeLength {
		end = unicodeLength
	}
	return hxrt.StringFromLiteral(string(runes[pos:end]))
}

func _UnicodeString__UnicodeString_Impl__indexOf(value any, str *string, startIndex ...int) int {
	runes := []rune(*hxrt.StdString(value))
	needle := []rune(*hxrt.StdString(str))
	start := 0
	if len(startIndex) > 0 {
		start = startIndex[0]
		if start < 0 {
			start = len(runes) + start
		}
	}
	if start < 0 {
		start = 0
	}
	if len(needle) == 0 {
		if start > len(runes) {
			return len(runes)
		}
		return start
	}
	if start > len(runes) || len(needle) > len(runes) {
		return -1
	}
	for i := start; i+len(needle) <= len(runes); i++ {
		matched := true
		for j := 0; j < len(needle); j++ {
			if runes[i+j] != needle[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}

func _UnicodeString__UnicodeString_Impl__lastIndexOf(value any, str *string, startIndex ...int) int {
	runes := []rune(*hxrt.StdString(value))
	needle := []rune(*hxrt.StdString(str))
	if len(needle) == 0 {
		if len(startIndex) == 0 {
			return len(runes)
		}
		start := startIndex[0]
		if start < 0 {
			start = 0
		}
		if start > len(runes) {
			return len(runes)
		}
		return start
	}
	start := len(runes)
	if len(startIndex) > 0 {
		start = startIndex[0]
		if start < 0 {
			start = 0
		}
	}
	limit := start + len(needle)
	if limit > len(runes) {
		limit = len(runes)
	}
	for i := limit - len(needle); i >= 0; i-- {
		matched := true
		for j := 0; j < len(needle); j++ {
			if runes[i+j] != needle[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}

func _UnicodeString__UnicodeString_Impl__validate(value *haxe__io__Bytes, encoding *haxe__io__Encoding) bool {
	if haxe__io__resolveEncodingTag(encoding) == 1 {
		hxrt.Throw(hxrt.StringFromLiteral("UnicodeString.validate: RawNative encoding is not supported"))
		return false
	}
	raw := hxrt_haxeBytesToRaw(value)
	pos := 0
	max := len(raw)
	for pos < max {
		c := int(raw[pos])
		pos++
		if c < 0x80 {
			continue
		} else if c < 0xC2 {
			return false
		} else if c < 0xE0 {
			if pos+1 > max {
				return false
			}
			c2 := int(raw[pos])
			pos++
			if c2 < 0x80 || c2 > 0xBF {
				return false
			}
		} else if c < 0xF0 {
			if pos+2 > max {
				return false
			}
			c2 := int(raw[pos])
			pos++
			if c == 0xE0 {
				if c2 < 0xA0 || c2 > 0xBF {
					return false
				}
			} else if c2 < 0x80 || c2 > 0xBF {
				return false
			}
			c3 := int(raw[pos])
			pos++
			if c3 < 0x80 || c3 > 0xBF {
				return false
			}
			c = (c << 16) | (c2 << 8) | c3
			if 0xEDA080 <= c && c <= 0xEDBFBF {
				return false
			}
		} else if c > 0xF4 {
			return false
		} else {
			if pos+3 > max {
				return false
			}
			c2 := int(raw[pos])
			pos++
			if c == 0xF0 {
				if c2 < 0x90 || c2 > 0xBF {
					return false
				}
			} else if c == 0xF4 {
				if c2 < 0x80 || c2 > 0x8F {
					return false
				}
			} else if c2 < 0x80 || c2 > 0xBF {
				return false
			}
			c3 := int(raw[pos])
			pos++
			if c3 < 0x80 || c3 > 0xBF {
				return false
			}
			c4 := int(raw[pos])
			pos++
			if c4 < 0x80 || c4 > 0xBF {
				return false
			}
		}
	}
	return true
}

type Date struct {
	value time.Time
}

func Date_fromString(source *string) *Date {
	raw := *hxrt.StdString(source)
	parsed, err := time.ParseInLocation("2006-01-02 15:04:05", raw, time.Local)
	if err != nil {
		parsedDateOnly, errDateOnly := time.ParseInLocation("2006-01-02", raw, time.Local)
		if errDateOnly == nil {
			parsed = parsedDateOnly
		} else {
			parsed = time.Unix(0, 0)
		}
	}
	return &Date{value: parsed}
}

func Date_now() *Date {
	return &Date{value: time.Now()}
}

func Date_fromTime(ms float64) *Date {
	nanos := int64(ms * 1e6)
	return &Date{value: time.Unix(0, nanos).In(time.Local)}
}

func (self *Date) getFullYear() int {
	return self.value.Year()
}

func (self *Date) getMonth() int {
	return int(self.value.Month()) - 1
}

func (self *Date) getDate() int {
	return self.value.Day()
}

func (self *Date) getDay() int {
	return int(self.value.Weekday())
}

func (self *Date) getHours() int {
	return self.value.Hour()
}

func (self *Date) getMinutes() int {
	return self.value.Minute()
}

func (self *Date) getSeconds() int {
	return self.value.Second()
}

func (self *Date) getTime() float64 {
	return float64(self.value.UnixNano()) / 1e6
}

type Math struct {
}

func Math_floor(value float64) int {
	return int(math.Floor(value))
}

func Math_ceil(value float64) int {
	return int(math.Ceil(value))
}

func Math_round(value float64) int {
	return int(math.Floor(value + 0.5))
}

func Math_abs(value float64) float64 {
	return math.Abs(value)
}

func Math_isNaN(value float64) bool {
	return math.IsNaN(value)
}

func Math_isFinite(value float64) bool {
	return !math.IsInf(value, 0)
}

func Math_min(a float64, b float64) float64 {
	return math.Min(a, b)
}

func Math_max(a float64, b float64) float64 {
	return math.Max(a, b)
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

var Xml_Element int = 0

var Xml_PCData int = 1

var Xml_CData int = 2

var Xml_Comment int = 3

var Xml_DocType int = 4

var Xml_ProcessingInstruction int = 5

var Xml_Document int = 6

type Xml struct {
	nodeType       int
	nodeName       *string
	nodeValue      *string
	parent         *Xml
	children       []*Xml
	attributeMap   map[string]string
	attributeOrder []string
}

func _Xml__XmlType_Impl__toString(value int) *string {
	switch value {
	case Xml_Element:
		return hxrt.StringFromLiteral("Element")
	case Xml_PCData:
		return hxrt.StringFromLiteral("PCData")
	case Xml_CData:
		return hxrt.StringFromLiteral("CData")
	case Xml_Comment:
		return hxrt.StringFromLiteral("Comment")
	case Xml_DocType:
		return hxrt.StringFromLiteral("DocType")
	case Xml_ProcessingInstruction:
		return hxrt.StringFromLiteral("ProcessingInstruction")
	case Xml_Document:
		return hxrt.StringFromLiteral("Document")
	default:
		return hxrt.StringFromLiteral("XmlType")
	}
}

func New_Xml(nodeType int) *Xml {
	return &Xml{nodeType: nodeType, children: []*Xml{}, attributeMap: map[string]string{}, attributeOrder: []string{}}
}

func Xml_createElement(name *string) *Xml {
	xml := New_Xml(Xml_Element)
	xml.nodeName = name
	return xml
}

func Xml_createPCData(data *string) *Xml {
	xml := New_Xml(Xml_PCData)
	xml.nodeValue = data
	return xml
}

func Xml_createCData(data *string) *Xml {
	xml := New_Xml(Xml_CData)
	xml.nodeValue = data
	return xml
}

func Xml_createComment(data *string) *Xml {
	xml := New_Xml(Xml_Comment)
	xml.nodeValue = data
	return xml
}

func Xml_createDocType(data *string) *Xml {
	xml := New_Xml(Xml_DocType)
	xml.nodeValue = data
	return xml
}

func Xml_createProcessingInstruction(data *string) *Xml {
	xml := New_Xml(Xml_ProcessingInstruction)
	xml.nodeValue = data
	return xml
}

func Xml_createDocument() *Xml {
	return New_Xml(Xml_Document)
}

func Xml_ensureElementType(self *Xml) {
	if self.nodeType != Xml_Document && self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element or Document but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
	}
}

func Xml_parse(source *string) *Xml {
	return haxe__xml__Parser_parse(source)
}

func (self *Xml) get(att *string) *string {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return nil
	}
	key := *hxrt.StdString(att)
	value, ok := self.attributeMap[key]
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(value)
}

func (self *Xml) set(att *string, value *string) {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return
	}
	key := *hxrt.StdString(att)
	if _, ok := self.attributeMap[key]; !ok {
		self.attributeOrder = append(self.attributeOrder, key)
	}
	self.attributeMap[key] = *hxrt.StdString(value)
}

func (self *Xml) remove(att *string) {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return
	}
	key := *hxrt.StdString(att)
	delete(self.attributeMap, key)
	filtered := make([]string, 0, len(self.attributeOrder))
	for _, existing := range self.attributeOrder {
		if existing != key {
			filtered = append(filtered, existing)
		}
	}
	self.attributeOrder = filtered
}

func (self *Xml) exists(att *string) bool {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return false
	}
	_, ok := self.attributeMap[*hxrt.StdString(att)]
	return ok
}

func (self *Xml) attributes() map[string]any {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return map[string]any{}
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(self.attributeOrder) }
	iter["next"] = func() *string { key := self.attributeOrder[index]; index++; return hxrt.StringFromLiteral(key) }
	return iter
}

func (self *Xml) iterator() map[string]any {
	Xml_ensureElementType(self)
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(self.children) }
	iter["next"] = func() *Xml { child := self.children[index]; index++; return child }
	return iter
}

func (self *Xml) elements() map[string]any {
	Xml_ensureElementType(self)
	matches := make([]*Xml, 0, len(self.children))
	for _, child := range self.children {
		if child.nodeType == Xml_Element {
			matches = append(matches, child)
		}
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(matches) }
	iter["next"] = func() *Xml { child := matches[index]; index++; return child }
	return iter
}

func (self *Xml) elementsNamed(name *string) map[string]any {
	Xml_ensureElementType(self)
	wanted := *hxrt.StdString(name)
	matches := make([]*Xml, 0, len(self.children))
	for _, child := range self.children {
		if child.nodeType == Xml_Element && *hxrt.StdString(child.nodeName) == wanted {
			matches = append(matches, child)
		}
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(matches) }
	iter["next"] = func() *Xml { child := matches[index]; index++; return child }
	return iter
}

func (self *Xml) firstChild() *Xml {
	Xml_ensureElementType(self)
	if len(self.children) == 0 {
		return nil
	}
	return self.children[0]
}

func (self *Xml) firstElement() *Xml {
	Xml_ensureElementType(self)
	for _, child := range self.children {
		if child.nodeType == Xml_Element {
			return child
		}
	}
	return nil
}

func (self *Xml) addChild(x *Xml) {
	Xml_ensureElementType(self)
	if x == nil {
		return
	}
	if x.parent != nil {
		x.parent.removeChild(x)
	}
	self.children = append(self.children, x)
	x.parent = self
}

func (self *Xml) removeChild(x *Xml) bool {
	Xml_ensureElementType(self)
	for i, child := range self.children {
		if child == x {
			self.children = append(self.children[:i], self.children[i+1:]...)
			x.parent = nil
			return true
		}
	}
	return false
}

func (self *Xml) insertChild(x *Xml, pos int) {
	Xml_ensureElementType(self)
	if x == nil {
		return
	}
	if x.parent != nil {
		x.parent.removeChild(x)
	}
	if pos < 0 {
		pos = 0
	}
	if pos > len(self.children) {
		pos = len(self.children)
	}
	self.children = append(self.children, nil)
	copy(self.children[pos+1:], self.children[pos:])
	self.children[pos] = x
	x.parent = self
}

func (self *Xml) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return haxe__xml__Printer_print(self)
}

type haxe__crypto__Base64 struct {
}

type haxe__crypto__Md5 struct {
}

type haxe__crypto__Sha1 struct {
}

type haxe__crypto__Sha224 struct {
}

type haxe__crypto__Sha256 struct {
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

func haxe__crypto__Base64_encode(bytes *haxe__io__Bytes, complement ...bool) *string {
	useComplement := true
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	encoded := base64.StdEncoding.EncodeToString(hxrt_haxeBytesToRaw(bytes))
	if !useComplement {
		encoded = strings.TrimRight(encoded, "=")
	}
	return hxrt.StringFromLiteral(encoded)
}

func haxe__crypto__Base64_decode(value *string, complement ...bool) *haxe__io__Bytes {
	useComplement := true
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	rawValue := *hxrt.StdString(value)
	if useComplement {
		rawValue = strings.TrimRight(rawValue, "=")
	}
	decoded, err := base64.RawStdEncoding.DecodeString(rawValue)
	if err != nil {
		decoded, err = base64.StdEncoding.DecodeString(*hxrt.StdString(value))
		if err != nil {
			hxrt.Throw(err)
			return &haxe__io__Bytes{b: []int{}, length: 0}
		}
	}
	return hxrt_rawToHaxeBytes(decoded)
}

func haxe__crypto__Base64_urlEncode(bytes *haxe__io__Bytes, complement ...bool) *string {
	useComplement := false
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	encoded := base64.RawURLEncoding.EncodeToString(hxrt_haxeBytesToRaw(bytes))
	if useComplement {
		missing := len(encoded) % 4
		if missing != 0 {
			encoded = (encoded + strings.Repeat("=", (4-missing)))
		}
	}
	return hxrt.StringFromLiteral(encoded)
}

func haxe__crypto__Base64_urlDecode(value *string, complement ...bool) *haxe__io__Bytes {
	rawValue := *hxrt.StdString(value)
	decoded, err := base64.RawURLEncoding.DecodeString(strings.TrimRight(rawValue, "="))
	if err != nil {
		hxrt.Throw(err)
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return hxrt_rawToHaxeBytes(decoded)
}

func haxe__crypto__Md5_encode(value *string) *string {
	sum := md5.Sum([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Md5_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := md5.Sum(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha1_encode(value *string) *string {
	sum := sha1.Sum([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha1_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha1.Sum(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha224_encode(value *string) *string {
	sum := sha256.Sum224([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha224_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha256.Sum224(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha256_encode(value *string) *string {
	sum := sha256.Sum256([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha256_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha256.Sum256(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}

type haxe__xml__Parser struct {
}

type haxe__xml__Printer struct {
}

func haxe__xml__Parser_parse(source *string, strict ...bool) *Xml {
	raw := *hxrt.StdString(source)
	doc := Xml_createDocument()
	stack := []*Xml{doc}
	decoder := xml.NewDecoder(strings.NewReader(raw))
	for {
		tokenStart := decoder.InputOffset()
		token, err := decoder.Token()
		tokenEnd := decoder.InputOffset()
		if err == io.EOF {
			break
		}
		if err != nil {
			hxrt.Throw(err)
			return Xml_createDocument()
		}
		current := stack[len(stack)-1]
		switch value := token.(type) {
		case xml.StartElement:
			node := Xml_createElement(hxrt.StringFromLiteral(value.Name.Local))
			for _, attr := range value.Attr {
				node.set(hxrt.StringFromLiteral(attr.Name.Local), hxrt.StringFromLiteral(attr.Value))
			}
			current.addChild(node)
			stack = append(stack, node)
		case xml.EndElement:
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			text := string([]byte(value))
			if len(text) != 0 {
				tokenSource := raw[tokenStart:tokenEnd]
				if strings.HasPrefix(tokenSource, "<![CDATA[") && strings.HasSuffix(tokenSource, "]]>") {
					current.addChild(Xml_createCData(hxrt.StringFromLiteral(text)))
				} else {
					current.addChild(Xml_createPCData(hxrt.StringFromLiteral(text)))
				}
			}
		case xml.Comment:
			current.addChild(Xml_createComment(hxrt.StringFromLiteral(string([]byte(value)))))
		case xml.Directive:
			directive := strings.TrimSpace(string([]byte(value)))
			upper := strings.ToUpper(directive)
			if strings.HasPrefix(upper, "DOCTYPE") {
				directive = strings.TrimSpace(directive[len("DOCTYPE"):])
			}
			current.addChild(Xml_createDocType(hxrt.StringFromLiteral(directive)))
		case xml.ProcInst:
			payload := value.Target
			if len(value.Inst) != 0 {
				payload += " " + string(value.Inst)
			}
			current.addChild(Xml_createProcessingInstruction(hxrt.StringFromLiteral(strings.TrimSpace(payload))))
		}
	}
	return doc
}

func haxe__xml__Printer_escapeText(value string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;").Replace(value)
}

func haxe__xml__Printer_escapeAttr(value string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "\"", "&quot;").Replace(value)
}

func haxe__xml__Printer_hasChildren(value *Xml) bool {
	for _, child := range value.children {
		switch child.nodeType {
		case Xml_Element, Xml_PCData:
			return true
		case Xml_CData, Xml_Comment:
			if len(strings.TrimLeft(*hxrt.StdString(child.nodeValue), " \n\r\t")) != 0 {
				return true
			}
		}
	}
	return false
}

func haxe__xml__Printer_writeNode(output *strings.Builder, value *Xml, tabs string, pretty bool) {
	newline := func() {
		if pretty {
			output.WriteString("\n")
		}
	}
	switch value.nodeType {
	case Xml_CData:
		output.WriteString(tabs + "<![CDATA[")
		output.WriteString(*hxrt.StdString(value.nodeValue))
		output.WriteString("]]>")
		newline()
	case Xml_Comment:
		commentContent := strings.NewReplacer("\n", "", "\r", "", "\t", "").Replace(*hxrt.StdString(value.nodeValue))
		output.WriteString(tabs)
		output.WriteString(strings.TrimSpace("<!--" + commentContent + "-->"))
		newline()
	case Xml_Document:
		for _, child := range value.children {
			haxe__xml__Printer_writeNode(output, child, tabs, pretty)
		}
	case Xml_Element:
		output.WriteString(tabs + "<")
		output.WriteString(*hxrt.StdString(value.nodeName))
		for _, attribute := range value.attributeOrder {
			output.WriteString(" " + attribute + "=\"")
			output.WriteString(haxe__xml__Printer_escapeAttr(value.attributeMap[attribute]))
			output.WriteString("\"")
		}
		if haxe__xml__Printer_hasChildren(value) {
			output.WriteString(">")
			newline()
			childTabs := tabs
			if pretty {
				childTabs = tabs + "\t"
			}
			for _, child := range value.children {
				haxe__xml__Printer_writeNode(output, child, childTabs, pretty)
			}
			output.WriteString(tabs + "</")
			output.WriteString(*hxrt.StdString(value.nodeName))
			output.WriteString(">")
			newline()
		} else {
			output.WriteString("/>")
			newline()
		}
	case Xml_PCData:
		nodeValue := *hxrt.StdString(value.nodeValue)
		if len(nodeValue) != 0 {
			output.WriteString(tabs + haxe__xml__Printer_escapeText(nodeValue))
			newline()
		}
	case Xml_ProcessingInstruction:
		output.WriteString("<?" + *hxrt.StdString(value.nodeValue) + "?>")
		newline()
	case Xml_DocType:
		output.WriteString("<!DOCTYPE " + *hxrt.StdString(value.nodeValue) + ">")
		newline()
	}
}

func haxe__xml__Printer_print(value *Xml, pretty ...bool) *string {
	if value == nil {
		return hxrt.StringFromLiteral("")
	}
	usePretty := false
	if len(pretty) > 0 {
		usePretty = pretty[0]
	}
	var output strings.Builder
	haxe__xml__Printer_writeNode(&output, value, "", usePretty)
	return hxrt.StringFromLiteral(output.String())
}

type haxe__zip__Compress struct {
}

type haxe__zip__Uncompress struct {
}

func haxe__zip__Compress_run(src *haxe__io__Bytes, level int) *haxe__io__Bytes {
	raw := hxrt_haxeBytesToRaw(src)
	var buffer bytes.Buffer
	writer, err := zlib.NewWriterLevel(&buffer, level)
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	if _, err := writer.Write(raw); err != nil {
		_ = writer.Close()
		hxrt.Throw(err)
		return nil
	}
	if err := writer.Close(); err != nil {
		hxrt.Throw(err)
		return nil
	}
	return hxrt_rawToHaxeBytes(buffer.Bytes())
}

func haxe__zip__Uncompress_run(src *haxe__io__Bytes, bufsize ...int) *haxe__io__Bytes {
	raw := hxrt_haxeBytesToRaw(src)
	reader, err := zlib.NewReader(bytes.NewReader(raw))
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	defer reader.Close()
	decoded, err := io.ReadAll(reader)
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	return hxrt_rawToHaxeBytes(decoded)
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
	case "haxe.io.GoIoHelpers":
		return nil, false
	case "sys.ssl.Certificate":
		return hxrt_typeCallAny(New_sys__ssl__Certificate, args)
	case "sys.ssl.Key":
		return hxrt_typeCallAny(New_sys__ssl__Key, args)
	case "sys.ssl.Socket":
		return hxrt_typeCallAny(New_sys__ssl__Socket, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	case "sys.ssl.Certificate":
		return &sys__ssl__Certificate{}, true
	case "sys.ssl.Key":
		return &sys__ssl__Key{}, true
	case "sys.ssl.Socket":
		return &sys__ssl__Socket{}, true
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
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
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
	case "haxe.io.GoIoHelpers":
		return nil
	case "sys.ssl.Certificate":
		return nil
	case "sys.ssl.Key":
		return nil
	case "sys.ssl.Socket":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.Socket")}
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

func Type_getClassFields(c any) []*string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return []*string{}
	}
	switch className {
	case "Main":
		return []*string{hxrt.StringFromLiteral("DEFAULT_CERT_PEM"), hxrt.StringFromLiteral("DEFAULT_KEY_PEM"), hxrt.StringFromLiteral("SNI_CERT_PEM"), hxrt.StringFromLiteral("SNI_KEY_PEM"), hxrt.StringFromLiteral("connectAndPrint"), hxrt.StringFromLiteral("main")}
	case "haxe.Int64Helper":
		return []*string{}
	case "haxe._Int32.Int32_Impl_":
		return []*string{}
	case "haxe._Int64.Int64_Impl_":
		return []*string{}
	case "haxe._Int64.___Int64":
		return []*string{}
	case "haxe.io.GoIoHelpers":
		return []*string{hxrt.StringFromLiteral("bytesOutputGetBytes"), hxrt.StringFromLiteral("inputRead"), hxrt.StringFromLiteral("inputReadAll"), hxrt.StringFromLiteral("inputReadBytes"), hxrt.StringFromLiteral("inputReadFullBytes"), hxrt.StringFromLiteral("inputReadLine"), hxrt.StringFromLiteral("inputReadUntil"), hxrt.StringFromLiteral("outputWrite"), hxrt.StringFromLiteral("outputWriteBytes"), hxrt.StringFromLiteral("outputWriteFullBytes"), hxrt.StringFromLiteral("outputWriteInput"), hxrt.StringFromLiteral("outputWriteString")}
	case "sys.ssl.Certificate":
		return []*string{hxrt.StringFromLiteral("fromString"), hxrt.StringFromLiteral("loadDefaults"), hxrt.StringFromLiteral("loadFile"), hxrt.StringFromLiteral("loadPath")}
	case "sys.ssl.Key":
		return []*string{hxrt.StringFromLiteral("loadFile"), hxrt.StringFromLiteral("readDER"), hxrt.StringFromLiteral("readPEM")}
	case "sys.ssl.Socket":
		return []*string{hxrt.StringFromLiteral("DEFAULT_CA"), hxrt.StringFromLiteral("DEFAULT_VERIFY_CERT")}
	default:
		return []*string{}
	}
}

func Type_getInstanceFields(c any) []*string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return []*string{}
	}
	switch className {
	case "Main":
		return []*string{}
	case "haxe.Int64Helper":
		return []*string{}
	case "haxe._Int32.Int32_Impl_":
		return []*string{}
	case "haxe._Int64.Int64_Impl_":
		return []*string{}
	case "haxe._Int64.___Int64":
		return []*string{hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low")}
	case "haxe.io.GoIoHelpers":
		return []*string{}
	case "sys.ssl.Certificate":
		return []*string{hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("addDER"), hxrt.StringFromLiteral("altNames"), hxrt.StringFromLiteral("commonName"), hxrt.StringFromLiteral("get_altNames"), hxrt.StringFromLiteral("get_commonName"), hxrt.StringFromLiteral("get_notAfter"), hxrt.StringFromLiteral("get_notBefore"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("issuer"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("notAfter"), hxrt.StringFromLiteral("notBefore"), hxrt.StringFromLiteral("subject")}
	case "sys.ssl.Key":
		return []*string{hxrt.StringFromLiteral("handle")}
	case "sys.ssl.Socket":
		return []*string{hxrt.StringFromLiteral("accept"), hxrt.StringFromLiteral("addSNICertificate"), hxrt.StringFromLiteral("bind"), hxrt.StringFromLiteral("caCert"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("connect"), hxrt.StringFromLiteral("custom"), hxrt.StringFromLiteral("handshake"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("hostname"), hxrt.StringFromLiteral("input"), hxrt.StringFromLiteral("listen"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("ownCert"), hxrt.StringFromLiteral("ownKey"), hxrt.StringFromLiteral("peer"), hxrt.StringFromLiteral("peerCertificate"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("setBlocking"), hxrt.StringFromLiteral("setCA"), hxrt.StringFromLiteral("setCertificate"), hxrt.StringFromLiteral("setFastSend"), hxrt.StringFromLiteral("setHostname"), hxrt.StringFromLiteral("setTimeout"), hxrt.StringFromLiteral("shutdown"), hxrt.StringFromLiteral("sniConfig"), hxrt.StringFromLiteral("verifyCert"), hxrt.StringFromLiteral("waitForRead"), hxrt.StringFromLiteral("write")}
	default:
		return []*string{}
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
	case "haxe.io.GoIoHelpers":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.ssl.Certificate":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.ssl.Key":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.ssl.Socket":
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
	default:
		return nil
	}
}

func Type_createInstance(cl any, args []any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, args)
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

func Type_createEnum(e any, constr *string, params []any) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, params)
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params []any) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, params)
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
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) []*string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return []*string{}
	}
	switch enumName {
	case "ValueType":
		return []*string{hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown")}
	default:
		return []*string{}
	}
}

func Type_enumParameters(e any) []any {
	if hxrt.AnyEqualsNull(e) {
		return []any{}
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return []any{}
		}
		out := make([]any, len(value.params))
		copy(out, value.params)
		return out
	default:
		return []any{}
	}
}

func Type_allEnums(e any) []any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return []any{}
	}
	switch enumName {
	case "ValueType":
		return []any{ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown}
	default:
		return []any{}
	}
}

func Type_typeof(v any) any {
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
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
