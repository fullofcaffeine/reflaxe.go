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
	switch *hxrt.StdString(path) {
	case *hxrt.StdString(hxrt.StringFromLiteral(".")), *hxrt.StdString(hxrt.StringFromLiteral("..")):
		self.dir = path
		self.file = hxrt.StringFromLiteral("")
		return self
	}
	c1 := hxrt.StringLastIndexOfStringPtr(path, hxrt.StringFromLiteral("/"), 0, false)
	c2 := hxrt.StringLastIndexOfStringPtr(path, hxrt.StringFromLiteral("\\"), 0, false)
	if c1 < c2 {
		self.dir = hxrt.StringSubstrStringPtr(path, 0, c2, true)
		path = hxrt.StringSubstrStringPtr(path, int(int32((hxrt.Int32Wrap(c2) + hxrt.Int32Wrap(1)))), 0, false)
		self.backslash = true
	} else {
		if c2 < c1 {
			self.dir = hxrt.StringSubstrStringPtr(path, 0, c1, true)
			path = hxrt.StringSubstrStringPtr(path, int(int32((hxrt.Int32Wrap(c1) + hxrt.Int32Wrap(1)))), 0, false)
		} else {
			self.dir = nil
		}
	}
	cp := hxrt.StringLastIndexOfStringPtr(path, hxrt.StringFromLiteral("."), 0, false)
	if cp != -1 {
		self.ext = hxrt.StringSubstrStringPtr(path, int(int32((hxrt.Int32Wrap(cp) + hxrt.Int32Wrap(1)))), 0, false)
		self.file = hxrt.StringSubstrStringPtr(path, 0, cp, true)
	} else {
		self.ext = nil
		self.file = path
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
	c1 := hxrt.StringLastIndexOfStringPtr(path, hxrt.StringFromLiteral("/"), 0, false)
	c2 := hxrt.StringLastIndexOfStringPtr(path, hxrt.StringFromLiteral("\\"), 0, false)
	var hx_if_6 *string
	if c1 < c2 {
		var hx_if_4 *string
		if c2 != int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1)))) {
			hx_if_4 = hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("\\"))
		} else {
			hx_if_4 = path
		}
		hx_if_6 = hx_if_4
	} else {
		var hx_if_5 *string
		if c1 != int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1)))) {
			hx_if_5 = hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("/"))
		} else {
			hx_if_5 = path
		}
		hx_if_6 = hx_if_5
	}
	return hx_if_6
}

func haxe__io__Path_directory(path *string) *string {
	s := New_haxe__io__Path(path)
	if hxrt.StringEqualStringPtr(s.dir, nil) {
		return hxrt.StringFromLiteral("")
	}
	return s.dir
}

func haxe__io__Path_extension(path *string) *string {
	s := New_haxe__io__Path(path)
	if hxrt.StringEqualStringPtr(s.ext, nil) {
		return hxrt.StringFromLiteral("")
	}
	return s.ext
}

func haxe__io__Path_isAbsolute(path *string) bool {
	if StringTools_startsWith(path, hxrt.StringFromLiteral("/")) {
		return true
	}
	if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(path, 1), hxrt.StringFromLiteral(":")) {
		return true
	}
	if StringTools_startsWith(path, hxrt.StringFromLiteral("\\\\")) {
		return true
	}
	return false
}

func haxe__io__Path_join(paths []*string) *string {
	_g := []*string{}
	_g1 := 0
	_g2 := paths
	for _g1 < len(_g2) {
		v := _g2[_g1]
		_g1 = int(int32((_g1 + 1)))
		if !hxrt.StringEqualStringPtr(v, nil) && !hxrt.StringEqualStringPtr(v, hxrt.StringFromLiteral("")) {
			_g = append(_g, v)
		}
	}
	paths_1 := _g
	if len(paths_1) == 0 {
		return hxrt.StringFromLiteral("")
	}
	path := paths_1[0]
	_g_1 := 1
	_g1_1 := len(paths_1)
	for _g_1 < _g1_1 {
		hx_post_8 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		i := hx_post_8
		path = haxe__io__Path_addTrailingSlash(path)
		path = hxrt.StringConcatStringPtr(path, paths_1[i])
	}
	return haxe__io__Path_normalize(path)
}

func haxe__io__Path_normalize(path *string) *string {
	slash := hxrt.StringFromLiteral("/")
	path = hxrt.StringJoinAny(func(hx_sort_src_9 []*string) []any {
		hx_sort_out_11 := make([]any, 0, len(hx_sort_src_9))
		for _, hx_sort_item_10 := range hx_sort_src_9 {
			hx_sort_out_11 = append(hx_sort_out_11, hx_sort_item_10)
		}
		return hx_sort_out_11
	}(hxrt.StringSplitStringPtr(path, hxrt.StringFromLiteral("\\"))), slash)
	if hxrt.StringEqualStringPtr(path, slash) {
		return slash
	}
	target := []*string{}
	_g := 0
	_g1 := hxrt.StringSplitStringPtr(path, slash)
	for _g < len(_g1) {
		token := _g1[_g]
		_g = int(int32((_g + 1)))
		if (hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral("..")) && (len(target) > 0)) && !hxrt.StringEqualStringPtr(target[int(int32((hxrt.Int32Wrap(len(target))-hxrt.Int32Wrap(1))))], hxrt.StringFromLiteral("..")) {
			if len(target) > 0 {
				target = target[:(len(target) - 1)]
			}
		} else {
			if hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral("")) {
				if (len(target) > 0) || (hxrt.StringCharCodeAtAnyStringPtr(path, 0) == 47) {
					target = append(target, token)
				}
			} else {
				if !hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral(".")) {
					target = append(target, token)
				}
			}
		}
	}
	tmp := hxrt.StringJoinAny(func(hx_sort_src_15 []*string) []any {
		hx_sort_out_17 := make([]any, 0, len(hx_sort_src_15))
		for _, hx_sort_item_16 := range hx_sort_src_15 {
			hx_sort_out_17 = append(hx_sort_out_17, hx_sort_item_16)
		}
		return hx_sort_out_17
	}(target), slash)
	var acc_b *string
	acc_b = hxrt.StringFromLiteral("")
	colon := false
	slashes := false
	var _g_s *string
	var _g_offset int
	_g_offset = 0
	_g_s = tmp
	for _g_offset < hxrt.StringLengthStringPtr(_g_s) {
		value := _g_s
		hx_post_18 := _g_offset
		_g_offset = int(int32((_g_offset + 1)))
		index := hx_post_18
		c_1 := hxrt.StringCharCodeAtStringPtr(value, index)
		if ((c_1 >= 55296) && (c_1 <= 56319)) && (_g_offset < hxrt.StringLengthStringPtr(_g_s)) {
			c_1 = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(c_1) - hxrt.Int32Wrap(55232))))) << uint(10))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(func() int {
				value_1 := _g_s
				hx_post_19 := _g_offset
				_g_offset = int(int32((_g_offset + 1)))
				index_1 := hx_post_19
				return hxrt.StringCharCodeAtStringPtr(value_1, index_1)
			}()) & hxrt.Int32Wrap(1023))))))))
		}
		c := c_1
		switch c {
		case 47:
			if !colon {
				slashes = true
			} else {
				i := c
				colon = false
				if slashes {
					acc_b = hxrt.StringConcatStringPtr(acc_b, hxrt.StringFromLiteral("/"))
					slashes = false
				}
				acc_b = hxrt.StringConcatStringPtr(acc_b, hxrt.StringFromCharCode(i))
			}
		case 58:
			acc_b = hxrt.StringConcatStringPtr(acc_b, hxrt.StringFromLiteral(":"))
			colon = true
		default:
			i_1 := c
			colon = false
			if slashes {
				acc_b = hxrt.StringConcatStringPtr(acc_b, hxrt.StringFromLiteral("/"))
				slashes = false
			}
			acc_b = hxrt.StringConcatStringPtr(acc_b, hxrt.StringFromCharCode(i_1))
		}
	}
	return acc_b
}

func haxe__io__Path_removeTrailingSlashes(path *string) *string {
hx_loop_20:
	for true {
		var _g any = hxrt.StringCharCodeAtAnyStringPtr(path, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1)))))
		if _g == nil {
			break
		} else {
			switch _g.(int) {
			case 47, 92:
				path = hxrt.StringSubstrStringPtr(path, 0, -1, true)
			default:
				break hx_loop_20
			}
		}
	}
	return path
}

func haxe__io__Path_withExtension(path *string, ext *string) *string {
	s := New_haxe__io__Path(path)
	s.ext = ext
	return s.toString()
}

func haxe__io__Path_withoutDirectory(path *string) *string {
	s := New_haxe__io__Path(path)
	s.dir = nil
	return s.toString()
}

func haxe__io__Path_withoutExtension(path *string) *string {
	s := New_haxe__io__Path(path)
	s.ext = nil
	return s.toString()
}
