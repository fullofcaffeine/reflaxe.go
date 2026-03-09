package main

import "snapshot/hxrt"

func haxe__io__GoIoHelpers_bytesOutputGetBytes(out *haxe__io__BytesOutput) *haxe__io__Bytes {
	return out.getBytes()
}

func haxe__io__GoIoHelpers_inputRead(self haxe__io__Input, nbytes int) *haxe__io__Bytes {
	out := haxe__io__Bytes_alloc(nbytes)
	haxe__io__GoIoHelpers_inputReadFullBytes(self, out, 0, nbytes)
	return out
}

func haxe__io__GoIoHelpers_inputReadAll(self haxe__io__Input, bufsize int) *haxe__io__Bytes {
	if self == nil {
		return haxe__io__Bytes_alloc(0)
	}
	buf := haxe__io__Bytes_alloc(bufsize)
	_ = buf
	total := New_haxe__io__BytesOutput()
	_ = total
	for true {
		done := false
		hxrt.TryCatch(func() {
			chunk := self.readBytes(buf, 0, bufsize)
			_ = chunk
			if chunk == 0 {
				hxrt.Throw(haxe__io__Error_Blocked)
			}
			_ = func() int { total.writeFullBytes(buf, 0, chunk); return 0 }()
		}, func(hx_caught_1 any) {
			switch hx_typed_2 := hx_caught_1.(type) {
			case *haxe__io__Eof:
				hx_tmp := hx_typed_2
				_ = hx_tmp
				done = true
			default:
				hxrt.Throw(hx_caught_1)
			}
		})
		if done {
			break
		}
	}
	return total.getBytes()
}

func haxe__io__GoIoHelpers_inputReadBytes(self haxe__io__Input, buf *haxe__io__Bytes, pos int, len int) int {
	return self.readBytes(buf, pos, len)
}

func haxe__io__GoIoHelpers_inputReadFullBytes(self haxe__io__Input, s *haxe__io__Bytes, pos int, len int) {
	if self == nil {
		hxrt.Throw(haxe__io__Error_Blocked)
	}
	for len > 0 {
		read := self.readBytes(s, pos, len)
		if read == 0 {
			hxrt.Throw(haxe__io__Error_Blocked)
		}
		pos = int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(read))))
		len = int(int32((hxrt.Int32Wrap(len) - hxrt.Int32Wrap(read))))
	}
}

func haxe__io__GoIoHelpers_inputReadLine(self haxe__io__Input) *string {
	buf := New_haxe__io__BytesOutput()
	_ = buf
	for true {
		last := 0
		_ = last
		ended := false
		hx_try_return_3 := false
		var hx_try_value_4 *string
		hxrt.TryCatch(func() {
			last = self.readByte()
			if last == 10 {
				ended = true
			} else {
				_ = func() int { buf.writeByte(last); return 0 }()
			}
		}, func(hx_caught_5 any) {
			switch hx_typed_6 := hx_caught_5.(type) {
			case *haxe__io__Eof:
				e := hx_typed_6
				partial := buf.getBytes().toString()
				if hxrt.StringLengthStringPtr(partial) == 0 {
					hxrt.Throw(e)
				}
				hx_try_value_4 = partial
				hx_try_return_3 = true
				return
			default:
				hxrt.Throw(hx_caught_5)
			}
		})
		if hx_try_return_3 {
			return hx_try_value_4
		}
		if ended {
			break
		}
	}
	out := buf.getBytes().toString()
	if (hxrt.StringLengthStringPtr(out) > 0) && (hxrt.StringCharCodeAtAnyStringPtr(out, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(out))-hxrt.Int32Wrap(1))))) == 13) {
		out = hxrt.StringSubstrStringPtr(out, 0, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(out)) - hxrt.Int32Wrap(1)))), true)
	}
	return out
}

func haxe__io__GoIoHelpers_inputReadUntil(self haxe__io__Input, end int) *string {
	buf := New_haxe__io__BytesOutput()
	_ = buf
	for true {
		last := self.readByte()
		_ = last
		if last == end {
			break
		}
		_ = func() int { buf.writeByte(last); return 0 }()
	}
	return buf.getBytes().toString()
}

func haxe__io__GoIoHelpers_outputWrite(self haxe__io__Output, s *haxe__io__Bytes) {
	if (self == nil) || (s == nil) {
		return
	}
	remaining := s.length
	_ = remaining
	pos := 0
	_ = pos
	for remaining > 0 {
		wrote := self.writeBytes(s, pos, remaining)
		if wrote == 0 {
			hxrt.Throw(haxe__io__Error_Blocked)
		}
		pos = int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(wrote))))
		remaining = int(int32((hxrt.Int32Wrap(remaining) - hxrt.Int32Wrap(wrote))))
	}
}

func haxe__io__GoIoHelpers_outputWriteBytes(self haxe__io__Output, buf *haxe__io__Bytes, pos int, len int) int {
	return self.writeBytes(buf, pos, len)
}

func haxe__io__GoIoHelpers_outputWriteFullBytes(self haxe__io__Output, s *haxe__io__Bytes, pos int, len int) {
	for len > 0 {
		wrote := self.writeBytes(s, pos, len)
		if wrote == 0 {
			hxrt.Throw(haxe__io__Error_Blocked)
		}
		pos = int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(wrote))))
		len = int(int32((hxrt.Int32Wrap(len) - hxrt.Int32Wrap(wrote))))
	}
}

func haxe__io__GoIoHelpers_outputWriteInput(self haxe__io__Output, i haxe__io__Input, bufsize int) {
	if (self == nil) || (i == nil) {
		return
	}
	buf := haxe__io__Bytes_alloc(bufsize)
	_ = buf
	for true {
		done := false
		hxrt.TryCatch(func() {
			lenRead := i.readBytes(buf, 0, bufsize)
			if lenRead == 0 {
				hxrt.Throw(haxe__io__Error_Blocked)
			}
			haxe__io__GoIoHelpers_outputWriteFullBytes(self, buf, 0, lenRead)
		}, func(hx_caught_7 any) {
			switch hx_typed_8 := hx_caught_7.(type) {
			case *haxe__io__Eof:
				hx_tmp := hx_typed_8
				_ = hx_tmp
				done = true
			default:
				hxrt.Throw(hx_caught_7)
			}
		})
		if done {
			break
		}
	}
}

func haxe__io__GoIoHelpers_outputWriteString(self haxe__io__Output, s *string, encoding *haxe__io__Encoding) {
	if hxrt.StringEqualStringPtr(s, nil) {
		s = hxrt.StringFromLiteral("")
	}
	var hx_if_9 *haxe__io__Bytes
	if encoding == nil {
		hx_if_9 = haxe__io__Bytes_ofString(s)
	} else {
		hx_if_9 = haxe__io__Bytes_ofString(s, encoding)
	}
	bytes := hx_if_9
	haxe__io__GoIoHelpers_outputWriteFullBytes(self, bytes, 0, bytes.length)
}
