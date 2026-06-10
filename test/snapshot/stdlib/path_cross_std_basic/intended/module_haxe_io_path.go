package main

import "snapshot/hxrt"

type I_haxe__io__Path interface {
	toString() *string
}

type haxe__io__Path struct {
	__hx_this I_haxe__io__Path
	dir       *string
	file      *string
	ext       *string
	backslash bool
}

func New_haxe__io__Path(path *string) *haxe__io__Path {
	self := &haxe__io__Path{}
	self.__hx_this = self
	self.dir = nil
	self.file = hxrt.StringFromLiteral("")
	self.ext = nil
	self.backslash = false
	if hxrt.StringEqualStringPtr(path, hxrt.StringFromLiteral(".")) || hxrt.StringEqualStringPtr(path, hxrt.StringFromLiteral("..")) {
		self.dir = path
	} else {
		slashIndex := haxe__io__Path_lastIndexOfCode(path, 47)
		backslashIndex := haxe__io__Path_lastIndexOfCode(path, 92)
		if slashIndex < backslashIndex {
			self.dir = hxrt.StringSubstrStringPtr(path, 0, backslashIndex, true)
			path = hxrt.StringSubstrStringPtr(path, int(int32((hxrt.Int32Wrap(backslashIndex) + hxrt.Int32Wrap(1)))), 0, false)
			self.backslash = true
		} else {
			if backslashIndex < slashIndex {
				self.dir = hxrt.StringSubstrStringPtr(path, 0, slashIndex, true)
				path = hxrt.StringSubstrStringPtr(path, int(int32((hxrt.Int32Wrap(slashIndex) + hxrt.Int32Wrap(1)))), 0, false)
			}
		}
		dotIndex := haxe__io__Path_lastIndexOfCode(path, 46)
		if dotIndex != -1 {
			self.ext = hxrt.StringSubstrStringPtr(path, int(int32((hxrt.Int32Wrap(dotIndex) + hxrt.Int32Wrap(1)))), 0, false)
			self.file = hxrt.StringSubstrStringPtr(path, 0, dotIndex, true)
		} else {
			self.file = path
		}
	}
	return self
}

func (self *haxe__io__Path) toString() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_2 *string
		if hxrt.StringEqualStringPtr(self.dir, nil) {
			hx_if_2 = hxrt.StringFromLiteral("")
		} else {
			hx_if_2 = hxrt.StringConcatStringPtr(self.dir, func() *string {
				var hx_if_1 *string
				if self.backslash {
					hx_if_1 = hxrt.StringFromLiteral("\\")
				} else {
					hx_if_1 = hxrt.StringFromLiteral("/")
				}
				return hx_if_1
			}())
		}
		return hx_if_2
	}(), self.file), func() *string {
		var hx_if_3 *string
		if hxrt.StringEqualStringPtr(self.ext, nil) {
			hx_if_3 = hxrt.StringFromLiteral("")
		} else {
			hx_if_3 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("."), self.ext)
		}
		return hx_if_3
	}())
}

func haxe__io__Path_addTrailingSlash(path *string) *string {
	if hxrt.StringLengthStringPtr(path) == 0 {
		return hxrt.StringFromLiteral("/")
	}
	slashIndex := haxe__io__Path_lastIndexOfCode(path, 47)
	backslashIndex := haxe__io__Path_lastIndexOfCode(path, 92)
	var hx_if_6 *string
	if slashIndex < backslashIndex {
		var hx_if_4 *string
		if backslashIndex != int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1)))) {
			hx_if_4 = hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("\\"))
		} else {
			hx_if_4 = path
		}
		hx_if_6 = hx_if_4
	} else {
		var hx_if_5 *string
		if slashIndex != int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1)))) {
			hx_if_5 = hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("/"))
		} else {
			hx_if_5 = path
		}
		hx_if_6 = hx_if_5
	}
	return hx_if_6
}

func haxe__io__Path_directory(path *string) *string {
	resolved := New_haxe__io__Path(path)
	var hx_if_7 *string
	if hxrt.StringEqualStringPtr(resolved.dir, nil) {
		hx_if_7 = hxrt.StringFromLiteral("")
	} else {
		hx_if_7 = resolved.dir
	}
	return hx_if_7
}

func haxe__io__Path_extension(path *string) *string {
	resolved := New_haxe__io__Path(path)
	var hx_if_8 *string
	if hxrt.StringEqualStringPtr(resolved.ext, nil) {
		hx_if_8 = hxrt.StringFromLiteral("")
	} else {
		hx_if_8 = resolved.ext
	}
	return hx_if_8
}

func haxe__io__Path_isAbsolute(path *string) bool {
	if StringTools_startsWith(path, hxrt.StringFromLiteral("/")) {
		return true
	}
	if (hxrt.StringLengthStringPtr(path) > 1) && (func() int {
		var c any = hxrt.StringCharCodeAtAnyStringPtr(path, 1)
		var hx_if_9 int
		if c == nil {
			hx_if_9 = -1
		} else {
			hx_if_9 = hxrt.IntFromNullableAny(c)
		}
		return hx_if_9
	}() == 58) {
		return true
	}
	if StringTools_startsWith(path, hxrt.StringFromLiteral("\\\\")) {
		return true
	}
	return false
}

func haxe__io__Path_join(paths []*string) *string {
	filtered := []*string{}
	_g := 0
	for _g < len(paths) {
		segment := paths[_g]
		_g = int(int32((_g + 1)))
		if !hxrt.StringEqualStringPtr(segment, nil) && !hxrt.StringEqualStringPtr(segment, hxrt.StringFromLiteral("")) {
			filtered = append(filtered, segment)
		}
	}
	if len(filtered) == 0 {
		return hxrt.StringFromLiteral("")
	}
	path := filtered[0]
	_g_1 := 1
	_g1 := len(filtered)
	for _g_1 < _g1 {
		hx_post_11 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index := hx_post_11
		path = haxe__io__Path_addTrailingSlash(path)
		path = hxrt.StringConcatStringPtr(path, filtered[index])
	}
	return haxe__io__Path_normalize(path)
}

func haxe__io__Path_joinWithSlash(tokens []*string) *string {
	var output_b *string
	output_b = hxrt.StringFromLiteral("")
	_g := 0
	_g1 := len(tokens)
	for _g < _g1 {
		hx_post_12 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_12
		if index > 0 {
			output_b = hxrt.StringConcatStringPtr(output_b, hxrt.StringFromLiteral("/"))
		}
		output_b = hxrt.StringConcatStringPtr(output_b, hxrt.StdString(tokens[index]))
	}
	return output_b
}

func haxe__io__Path_lastIndexOfCode(path *string, code int) int {
	found := -1
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(path)
	for _g < _g1 {
		hx_post_13 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_13
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(path, index)
			var hx_if_14 int
			if c == nil {
				hx_if_14 = -1
			} else {
				hx_if_14 = hxrt.IntFromNullableAny(c)
			}
			return hx_if_14
		}() == code {
			found = index
		}
	}
	return found
}

func haxe__io__Path_normalize(path *string) *string {
	path = StringTools_replace(path, hxrt.StringFromLiteral("\\"), hxrt.StringFromLiteral("/"))
	if hxrt.StringEqualStringPtr(path, hxrt.StringFromLiteral("/")) {
		return hxrt.StringFromLiteral("/")
	}
	target := []*string{}
	absolute := ((hxrt.StringLengthStringPtr(path) > 0) && (func() int {
		var c any = hxrt.StringCharCodeAtAnyStringPtr(path, 0)
		var hx_if_15 int
		if c == nil {
			hx_if_15 = -1
		} else {
			hx_if_15 = hxrt.IntFromNullableAny(c)
		}
		return hx_if_15
	}() == 47))
	_g := 0
	_g1 := haxe__io__Path_splitOnSlash(path)
	for _g < len(_g1) {
		token := _g1[_g]
		_g = int(int32((_g + 1)))
		if (hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral("..")) && (len(target) > 0)) && !hxrt.StringEqualStringPtr(target[int(int32((hxrt.Int32Wrap(len(target))-hxrt.Int32Wrap(1))))], hxrt.StringFromLiteral("..")) {
			if len(target) > 0 {
				target = target[:(len(target) - 1)]
			}
		} else {
			if hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral("")) {
				if (len(target) > 0) || absolute {
					target = append(target, token)
				}
			} else {
				if !hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral(".")) {
					target = append(target, token)
				}
			}
		}
	}
	compact := haxe__io__Path_joinWithSlash(target)
	var output_b *string
	output_b = hxrt.StringFromLiteral("")
	sawColon := false
	sawSlashes := false
	_g_1 := 0
	_g1_1 := hxrt.StringLengthStringPtr(compact)
	for _g_1 < _g1_1 {
		hx_post_19 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index := hx_post_19
		var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(compact, index)
		var hx_if_20 int
		if c_1 == nil {
			hx_if_20 = -1
		} else {
			hx_if_20 = hxrt.IntFromNullableAny(c_1)
		}
		code := hx_if_20
		if code == 58 {
			x := hxrt.StringCharAtStringPtr(compact, index)
			output_b = hxrt.StringConcatStringPtr(output_b, hxrt.StdString(x))
			sawColon = true
		} else {
			if (code == 47) && !sawColon {
				sawSlashes = true
			} else {
				sawColon = false
				if sawSlashes {
					output_b = hxrt.StringConcatStringPtr(output_b, hxrt.StringFromLiteral("/"))
					sawSlashes = false
				}
				x_1 := hxrt.StringCharAtStringPtr(compact, index)
				output_b = hxrt.StringConcatStringPtr(output_b, hxrt.StdString(x_1))
			}
		}
	}
	return output_b
}

func haxe__io__Path_removeTrailingSlashes(path *string) *string {
	for hxrt.StringLengthStringPtr(path) > 0 {
		index := int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1))))
		var c any = hxrt.StringCharCodeAtAnyStringPtr(path, index)
		var hx_if_21 int
		if c == nil {
			hx_if_21 = -1
		} else {
			hx_if_21 = hxrt.IntFromNullableAny(c)
		}
		code := hx_if_21
		if (code == 47) || (code == 92) {
			path = hxrt.StringSubstrStringPtr(path, 0, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1)))), true)
			continue
		}
		return path
	}
	return path
}

func haxe__io__Path_splitOnSlash(path *string) []*string {
	tokens := []*string{}
	start := 0
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(path)
	for _g < _g1 {
		hx_post_22 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_22
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(path, index)
			var hx_if_23 int
			if c == nil {
				hx_if_23 = -1
			} else {
				hx_if_23 = hxrt.IntFromNullableAny(c)
			}
			return hx_if_23
		}() == 47 {
			tokens = append(tokens, hxrt.StringSubstrStringPtr(path, start, int(int32((hxrt.Int32Wrap(index)-hxrt.Int32Wrap(start)))), true))
			start = int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(1))))
		}
	}
	tokens = append(tokens, hxrt.StringSubstrStringPtr(path, start, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path))-hxrt.Int32Wrap(start)))), true))
	return tokens
}

func haxe__io__Path_withExtension(path *string, ext *string) *string {
	resolved := New_haxe__io__Path(path)
	resolved.ext = ext
	return resolved.toString()
}

func haxe__io__Path_withoutDirectory(path *string) *string {
	resolved := New_haxe__io__Path(path)
	resolved.dir = nil
	return resolved.toString()
}

func haxe__io__Path_withoutExtension(path *string) *string {
	resolved := New_haxe__io__Path(path)
	resolved.ext = nil
	return resolved.toString()
}
