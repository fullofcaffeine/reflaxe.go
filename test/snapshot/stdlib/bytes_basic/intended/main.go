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
		_g = (_g + 1)
		i := hx_post_1
		_ = i
		buffer.b = append(buffer.b, b2[i])
	}
	buffer.b = append(buffer.b, 33)
	out := buffer.getBytes()
	_ = out
	hxrt.Println(out.toString())
}

type haxe__io__Encoding struct {
}

type haxe__io__Input struct {
}

type haxe__io__Output struct {
}

type haxe__io__Bytes struct {
	b      []int
	length int
}

type haxe__io__BytesBuffer struct {
	b []int
}

func New_haxe__io__Input() *haxe__io__Input {
	return &haxe__io__Input{}
}

func New_haxe__io__Output() *haxe__io__Output {
	return &haxe__io__Output{}
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
	raw := hxrt.BytesFromString(value)
	return &haxe__io__Bytes{b: raw, length: len(raw)}
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
}

func New_haxe__io__BytesBuffer() *haxe__io__BytesBuffer {
	return &haxe__io__BytesBuffer{b: []int{}}
}

func (self *haxe__io__BytesBuffer) addByte(value int) {
	self.b = append(self.b, value)
}

func (self *haxe__io__BytesBuffer) add(src *haxe__io__Bytes) {
	if src == nil {
		return
	}
	self.b = append(self.b, src.b...)
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
