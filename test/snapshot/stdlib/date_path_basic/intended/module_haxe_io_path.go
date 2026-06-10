package main

import "snapshot/hxrt"

type I_haxe__io__Path interface {
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

func haxe__io__Path_addTrailingSlash(path *string) *string {
	if hxrt.StringLengthStringPtr(path) == 0 {
		return hxrt.StringFromLiteral("/")
	}
	c1 := hxrt.StringLastIndexOfStringPtr(path, hxrt.StringFromLiteral("/"), 0, false)
	c2 := hxrt.StringLastIndexOfStringPtr(path, hxrt.StringFromLiteral("\\"), 0, false)
	var hx_if_3 *string
	if c1 < c2 {
		var hx_if_1 *string
		if c2 != int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1)))) {
			hx_if_1 = hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("\\"))
		} else {
			hx_if_1 = path
		}
		hx_if_3 = hx_if_1
	} else {
		var hx_if_2 *string
		if c1 != int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1)))) {
			hx_if_2 = hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("/"))
		} else {
			hx_if_2 = path
		}
		hx_if_3 = hx_if_2
	}
	return hx_if_3
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
		hx_post_5 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		i := hx_post_5
		path = haxe__io__Path_addTrailingSlash(path)
		path = hxrt.StringConcatStringPtr(path, paths_1[i])
	}
	return haxe__io__Path_normalize(path)
}

func haxe__io__Path_normalize(path *string) *string {
	slash := hxrt.StringFromLiteral("/")
	path = hxrt.StringJoinAny(func(hx_sort_src_6 []*string) []any {
		hx_sort_out_8 := make([]any, 0, len(hx_sort_src_6))
		for _, hx_sort_item_7 := range hx_sort_src_6 {
			hx_sort_out_8 = append(hx_sort_out_8, hx_sort_item_7)
		}
		return hx_sort_out_8
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
	tmp := hxrt.StringJoinAny(func(hx_sort_src_12 []*string) []any {
		hx_sort_out_14 := make([]any, 0, len(hx_sort_src_12))
		for _, hx_sort_item_13 := range hx_sort_src_12 {
			hx_sort_out_14 = append(hx_sort_out_14, hx_sort_item_13)
		}
		return hx_sort_out_14
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
		hx_post_15 := _g_offset
		_g_offset = int(int32((_g_offset + 1)))
		index := hx_post_15
		c_1 := hxrt.StringCharCodeAtStringPtr(value, index)
		if ((c_1 >= 55296) && (c_1 <= 56319)) && (_g_offset < hxrt.StringLengthStringPtr(_g_s)) {
			c_1 = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(c_1) - hxrt.Int32Wrap(55232))))) << uint(10))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(func() int {
				value_1 := _g_s
				hx_post_16 := _g_offset
				_g_offset = int(int32((_g_offset + 1)))
				index_1 := hx_post_16
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
