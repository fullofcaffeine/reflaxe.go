package main

import "examples_incident_api_portable/hxrt"

var StringTools_MAX_HIGH_SURROGATE_CODE_POINT int = 56319

var StringTools_MIN_HIGH_SURROGATE_CODE_POINT int = 55296

var StringTools_MIN_SURROGATE_CODE_POINT int = 65536

func StringTools_contains(s *string, value *string) bool {
	return StringTools_containsImpl(s, value)
}

func StringTools_containsImpl(s *string, value *string) bool {
	if hxrt.StringLengthStringPtr(value) == 0 {
		return true
	}
	limit := int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s)) - hxrt.Int32Wrap(hxrt.StringLengthStringPtr(value)))))
	index := 0
	for index <= limit {
		if hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(s, index, hxrt.StringLengthStringPtr(value), true), value) {
			return true
		}
		index = int(int32((index + 1)))
	}
	return false
}

func StringTools_endsWith(s *string, end *string) bool {
	return StringTools_endsWithImpl(s, end)
}

func StringTools_endsWithImpl(s *string, end *string) bool {
	elen := hxrt.StringLengthStringPtr(end)
	slen := hxrt.StringLengthStringPtr(s)
	return ((slen >= elen) && hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(s, int(int32((hxrt.Int32Wrap(slen)-hxrt.Int32Wrap(elen)))), elen, true), end))
}

func StringTools_fastCodeAt(s *string, index int) int {
	var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
	var hx_if_184 int
	if c == nil {
		hx_if_184 = -1
	} else {
		hx_if_184 = c.(int)
	}
	return hx_if_184
}

func StringTools_hex(n int, digits any) *string {
	hexChars := hxrt.StringFromLiteral("0123456789ABCDEF")
	value := n
	out := hxrt.StringFromLiteral("")
	hx_do_first_185 := true
	for hx_do_first_185 || (value > 0) {
		hx_do_first_185 = false
		out = hxrt.StringConcatStringPtr(hxrt.StringCharAtStringPtr(hexChars, int(int32((hxrt.Int32Wrap(value)&hxrt.Int32Wrap(15))))), out)
		value = int(int32(int32((uint32(hxrt.Int32Wrap(value)) >> uint(4)))))
	}
	var hx_if_186 int
	if digits == nil {
		hx_if_186 = 0
	} else {
		hx_if_186 = digits.(int)
	}
	resolvedDigits := hx_if_186
	for (resolvedDigits != 0) && (hxrt.StringLengthStringPtr(out) < resolvedDigits) {
		out = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("0"), out)
	}
	return out
}

func StringTools_hexDigitValue(value *string) int {
	if hxrt.StringEqualStringPtr(value, nil) || (hxrt.StringLengthStringPtr(value) == 0) {
		return -1
	}
	var c any = hxrt.StringCharCodeAtAnyStringPtr(value, 0)
	var hx_if_187 int
	if c == nil {
		hx_if_187 = -1
	} else {
		hx_if_187 = c.(int)
	}
	code := hx_if_187
	if code == -1 {
		return -1
	}
	if (code >= 48) && (code <= 57) {
		return int(int32((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(48))))
	}
	if (code >= 65) && (code <= 70) {
		return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(65))))) + hxrt.Int32Wrap(10))))
	}
	if (code >= 97) && (code <= 102) {
		return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(97))))) + hxrt.Int32Wrap(10))))
	}
	return -1
}

func StringTools_htmlEscape(s *string, quotes any) *string {
	s = StringTools_replace(s, hxrt.StringFromLiteral("&"), hxrt.StringFromLiteral("&amp;"))
	s = StringTools_replace(s, hxrt.StringFromLiteral("<"), hxrt.StringFromLiteral("&lt;"))
	s = StringTools_replace(s, hxrt.StringFromLiteral(">"), hxrt.StringFromLiteral("&gt;"))
	if (quotes != nil) && (quotes.(bool) == true) {
		s = StringTools_replace(s, hxrt.StringFromLiteral("\""), hxrt.StringFromLiteral("&quot;"))
		s = StringTools_replace(s, hxrt.StringFromLiteral("'"), hxrt.StringFromLiteral("&#039;"))
	}
	return s
}

func StringTools_htmlUnescape(s *string) *string {
	s = StringTools_replace(s, hxrt.StringFromLiteral("&gt;"), hxrt.StringFromLiteral(">"))
	s = StringTools_replace(s, hxrt.StringFromLiteral("&lt;"), hxrt.StringFromLiteral("<"))
	s = StringTools_replace(s, hxrt.StringFromLiteral("&quot;"), hxrt.StringFromLiteral("\""))
	s = StringTools_replace(s, hxrt.StringFromLiteral("&#039;"), hxrt.StringFromLiteral("'"))
	s = StringTools_replace(s, hxrt.StringFromLiteral("&amp;"), hxrt.StringFromLiteral("&"))
	return s
}

func StringTools_isEof(c int) bool {
	return (c == -1)
}

func StringTools_isSpace(s *string, pos int) bool {
	if ((hxrt.StringLengthStringPtr(s) == 0) || (pos < 0)) || (pos >= hxrt.StringLengthStringPtr(s)) {
		return false
	}
	var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(s, pos)
	var hx_if_188 int
	if c_1 == nil {
		hx_if_188 = -1
	} else {
		hx_if_188 = c_1.(int)
	}
	c := hx_if_188
	return ((c != -1) && (((c > 8) && (c < 14)) || (c == 32)))
}

func StringTools_iterator(s *string) *haxe__iterators__StringIterator {
	return New_haxe__iterators__StringIterator(s)
}

func StringTools_keyValueIterator(s *string) *haxe__iterators__StringKeyValueIterator {
	return New_haxe__iterators__StringKeyValueIterator(s)
}

func StringTools_lpad(s *string, c *string, l int) *string {
	if hxrt.StringLengthStringPtr(c) <= 0 {
		return s
	}
	buf := hxrt.StringFromLiteral("")
	for int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(buf)) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s))))) < l {
		buf = hxrt.StringConcatStringPtr(buf, c)
	}
	return hxrt.StringConcatStringPtr(buf, s)
}

func StringTools_ltrim(s *string) *string {
	r := 0
	for (r < hxrt.StringLengthStringPtr(s)) && StringTools_isSpace(s, r) {
		r = int(int32((r + 1)))
	}
	var hx_if_189 *string
	if r > 0 {
		hx_if_189 = hxrt.StringSubstrStringPtr(s, r, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s)) - hxrt.Int32Wrap(r)))), true)
	} else {
		hx_if_189 = s
	}
	return hx_if_189
}

func StringTools_replace(s *string, sub *string, by *string) *string {
	if hxrt.StringLengthStringPtr(sub) == 0 {
		var every_b *string
		every_b = hxrt.StringFromLiteral("")
		every_b = hxrt.StringConcatStringPtr(every_b, hxrt.StdString(by))
		_g := 0
		_g1 := hxrt.StringLengthStringPtr(s)
		for _g < _g1 {
			hx_post_190 := _g
			_g = int(int32((_g + 1)))
			index := hx_post_190
			x := hxrt.StringSubstrStringPtr(s, index, 1, true)
			every_b = hxrt.StringConcatStringPtr(every_b, hxrt.StdString(x))
			every_b = hxrt.StringConcatStringPtr(every_b, hxrt.StdString(by))
		}
		return every_b
	}
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	index_1 := 0
	for index_1 < hxrt.StringLengthStringPtr(s) {
		if (int(int32((hxrt.Int32Wrap(index_1) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(sub))))) <= hxrt.StringLengthStringPtr(s)) && hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(s, index_1, hxrt.StringLengthStringPtr(sub), true), sub) {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(by))
			index_1 = int(int32((hxrt.Int32Wrap(index_1) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(sub)))))
		} else {
			x_1 := hxrt.StringSubstrStringPtr(s, index_1, 1, true)
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
			index_1 = int(int32((index_1 + 1)))
		}
	}
	return out_b
}

func StringTools_rpad(s *string, c *string, l int) *string {
	if hxrt.StringLengthStringPtr(c) <= 0 {
		return s
	}
	buf := s
	for hxrt.StringLengthStringPtr(buf) < l {
		buf = hxrt.StringConcatStringPtr(buf, c)
	}
	return buf
}

func StringTools_rtrim(s *string) *string {
	r := 0
	for (r < hxrt.StringLengthStringPtr(s)) && StringTools_isSpace(s, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s))-hxrt.Int32Wrap(r)))))-hxrt.Int32Wrap(1))))) {
		r = int(int32((r + 1)))
	}
	var hx_if_191 *string
	if r > 0 {
		hx_if_191 = hxrt.StringSubstrStringPtr(s, 0, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s)) - hxrt.Int32Wrap(r)))), true)
	} else {
		hx_if_191 = s
	}
	return hx_if_191
}

func StringTools_startsWith(s *string, start *string) bool {
	return StringTools_startsWithImpl(s, start)
}

func StringTools_startsWithImpl(s *string, start *string) bool {
	return ((hxrt.StringLengthStringPtr(s) >= hxrt.StringLengthStringPtr(start)) && hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(s, 0, hxrt.StringLengthStringPtr(start), true), start))
}

func StringTools_trim(s *string) *string {
	return StringTools_ltrim(StringTools_rtrim(s))
}

func StringTools_unsafeCodeAt(s *string, index int) int {
	var c any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
	var hx_if_192 int
	if c == nil {
		hx_if_192 = -1
	} else {
		hx_if_192 = c.(int)
	}
	return hx_if_192
}

func StringTools_urlDecode(s *string) *string {
	input := StringTools_replace(s, hxrt.StringFromLiteral("+"), hxrt.StringFromLiteral(" "))
	bytes := hxrt.NewArray()
	index := 0
	for index < hxrt.StringLengthStringPtr(input) {
		c := hxrt.StringSubstrStringPtr(input, index, 1, true)
		if hxrt.StringEqualStringPtr(c, hxrt.StringFromLiteral("%")) && (int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(2)))) < hxrt.StringLengthStringPtr(input)) {
			hi := StringTools_hexDigitValue(hxrt.StringSubstrStringPtr(input, int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(1)))), 1, true))
			lo := StringTools_hexDigitValue(hxrt.StringSubstrStringPtr(input, int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(2)))), 1, true))
			if (hi >= 0) && (lo >= 0) {
				bytes.Push(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(hi) << uint(4))))) | hxrt.Int32Wrap(lo)))))
				index = int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(3))))
				continue
			}
		}
		chunk := haxe__io__Bytes_ofString(hxrt.StringCharAtStringPtr(input, index), nil)
		_g := 0
		_g1 := chunk.length
		for _g < _g1 {
			hx_post_194 := _g
			_g = int(int32((_g + 1)))
			chunkIndex := hx_post_194
			bytes.Push(chunk.b[chunkIndex])
		}
		index = int(int32((index + 1)))
	}
	out := haxe__io__Bytes_alloc(bytes.Len())
	_g_1 := 0
	_g1_1 := bytes.Len()
	for _g_1 < _g1_1 {
		hx_post_196 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		byteIndex := hx_post_196
		out.b[byteIndex] = int(int32((hxrt.Int32Wrap(hxrt.IntFromNullableAny(bytes.Get(byteIndex))) & hxrt.Int32Wrap(255))))
		out.__hx_rawValid = false
	}
	return out.__hx_this.toString()
}

func StringTools_urlEncode(s *string) *string {
	bytes := haxe__io__Bytes_ofString(s, nil)
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	ascii := haxe__io__Bytes_alloc(1)
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_197 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_197
		b := bytes.b[index]
		isUnreserved := ((((((((b >= 65) && (b <= 90)) || ((b >= 97) && (b <= 122))) || ((b >= 48) && (b <= 57))) || (b == 45)) || (b == 95)) || (b == 46)) || (b == 126))
		if isUnreserved {
			ascii.b[0] = int(int32((hxrt.Int32Wrap(b) & hxrt.Int32Wrap(255))))
			ascii.__hx_rawValid = false
			x := ascii.__hx_this.toString()
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		} else {
			if b == 32 {
				out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("%20"))
			} else {
				out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("%"))
				x_1 := StringTools_hex(b, 2)
				out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
			}
		}
	}
	return out_b
}

func StringTools_utf16CodePointAt(s *string, index int) int {
	var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(s, index)
	var hx_if_198 int
	if c_1 == nil {
		hx_if_198 = -1
	} else {
		hx_if_198 = c_1.(int)
	}
	c := hx_if_198
	if (c >= 55296) && (c <= 56319) {
		c = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(c) - hxrt.Int32Wrap(55232))))) << uint(10))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(func() int {
			var c_2 any = hxrt.StringCharCodeAtAnyStringPtr(s, int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(1)))))
			var hx_if_199 int
			if c_2 == nil {
				hx_if_199 = -1
			} else {
				hx_if_199 = c_2.(int)
			}
			return hx_if_199
		}()) & hxrt.Int32Wrap(1023))))))))
	}
	return c
}
