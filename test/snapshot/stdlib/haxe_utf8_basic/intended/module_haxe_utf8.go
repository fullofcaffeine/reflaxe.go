package main

import "snapshot/hxrt"

type I_haxe__Utf8 interface {
	addChar(c int)
	toString() *string
}

type haxe__Utf8 struct {
	__hx_this I_haxe__Utf8
	__b       *string
}

func New_haxe__Utf8(size int) *haxe__Utf8 {
	self := &haxe__Utf8{}
	self.__hx_this = self
	self.__b = hxrt.StringFromLiteral("")
	return self
}

func (self *haxe__Utf8) addChar(c int) {
	self.__b = hxrt.StringConcatStringPtr(self.__b, haxe__Utf8_codePointToString(c))
}

func (self *haxe__Utf8) toString() *string {
	return self.__b
}

func (self *haxe__Utf8) String() string {
	return *self.__hx_this.toString()
}

func haxe__Utf8_charCodeAt(s *string, index int) int {
	return hxrt.StringCharCodeAtStringPtr(s, index)
}

func haxe__Utf8_codePointToString(code int) *string {
	var hx_if_19 *hxrt.Array
	if code < 128 {
		hx_if_19 = hxrt.NewArray(code)
	} else {
		var hx_if_18 *hxrt.Array
		if code < 2048 {
			hx_if_18 = hxrt.NewArray(int(int32((hxrt.Int32Wrap(192) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) >> uint(6)))))))), int(int32((hxrt.Int32Wrap(128) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) & hxrt.Int32Wrap(63)))))))))
		} else {
			var hx_if_17 *hxrt.Array
			if code < 65536 {
				hx_if_17 = hxrt.NewArray(int(int32((hxrt.Int32Wrap(224) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) >> uint(12)))))))), int(int32((hxrt.Int32Wrap(128) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) >> uint(6))))) & hxrt.Int32Wrap(63)))))))), int(int32((hxrt.Int32Wrap(128) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) & hxrt.Int32Wrap(63)))))))))
			} else {
				hx_if_17 = hxrt.NewArray(int(int32((hxrt.Int32Wrap(240) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) >> uint(18)))))))), int(int32((hxrt.Int32Wrap(128) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) >> uint(12))))) & hxrt.Int32Wrap(63)))))))), int(int32((hxrt.Int32Wrap(128) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) >> uint(6))))) & hxrt.Int32Wrap(63)))))))), int(int32((hxrt.Int32Wrap(128) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) & hxrt.Int32Wrap(63)))))))))
			}
			hx_if_18 = hx_if_17
		}
		hx_if_19 = hx_if_18
	}
	raw := hx_if_19
	bytes := haxe__io__Bytes_alloc(raw.Len())
	_g := 0
	_g1 := raw.Len()
	for _g < _g1 {
		hx_post_20 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_20
		bytes.b[index] = int(int32((hxrt.Int32Wrap(hxrt.IntFromNullableAny(raw.Get(index))) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	return bytes.__hx_this.toString()
}

func haxe__Utf8_compare(a *string, b *string) int {
	left := haxe__io__Bytes_ofString(a, nil)
	right := haxe__io__Bytes_ofString(b, nil)
	var hx_if_21 int
	if left.length < right.length {
		hx_if_21 = left.length
	} else {
		hx_if_21 = right.length
	}
	limit := hx_if_21
	_g := 0
	_g1 := limit
	for _g < _g1 {
		hx_post_22 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_22
		l := left.b[index]
		r := right.b[index]
		if l > r {
			return 1
		}
		if l < r {
			return -1
		}
	}
	if left.length > right.length {
		return 1
	}
	if left.length < right.length {
		return -1
	}
	return 0
}

func haxe__Utf8_decode(s *string) *string {
	bytes := haxe__io__Bytes_alloc(hxrt.StringLengthStringPtr(s))
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(s)
	for _g < _g1 {
		hx_post_23 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_23
		var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
		var hx_if_24 int
		if c == nil {
			hx_if_24 = -1
		} else {
			hx_if_24 = c.(int)
		}
		code := hx_if_24
		bytes.b[index] = int(int32((hxrt.Int32Wrap(func() int {
			var hx_if_25 int
			if code < 0 {
				hx_if_25 = 0
			} else {
				hx_if_25 = int(int32((hxrt.Int32Wrap(code) & hxrt.Int32Wrap(255))))
			}
			return hx_if_25
		}()) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	return bytes.__hx_this.toString()
}

func haxe__Utf8_encode(s *string) *string {
	bytes := haxe__io__Bytes_ofString(s, nil)
	out := hxrt.StringFromLiteral("")
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_26 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_26
		out = hxrt.StringConcatStringPtr(out, haxe__Utf8_codePointToString(bytes.b[index]))
	}
	return out
}

func haxe__Utf8_iter(s *string, chars func(int)) {
	var unicode any = s
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(hxrt.StdString(func(hx_value_27 any) *string {
		if hx_value_27 == nil {
			var hx_zero_28 *string
			return hx_zero_28
		}
		return hx_value_27.(*string)
	}(unicode)))
	for _g < _g1 {
		hx_post_29 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_29
		chars(hxrt.StringCharCodeAtStringPtr(s, index))
	}
}

func haxe__Utf8_length(s *string) int {
	return haxe__io__Bytes_ofString(s, nil).length
}

func haxe__Utf8_sub(s *string, pos int, len int) *string {
	var unicode any = s
	return _UnicodeString__UnicodeString_Impl__substr(hxrt.StdString(unicode), pos, len)
}

func haxe__Utf8_validate(s *string) bool {
	return _UnicodeString__UnicodeString_Impl__validate(haxe__io__Bytes_ofString(s, nil), haxe__io__Encoding_UTF8)
}
