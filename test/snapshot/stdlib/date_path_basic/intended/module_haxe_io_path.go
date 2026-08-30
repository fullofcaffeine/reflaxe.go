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
		path = hxrt.StringSubstrStringPtr(path, int((hxrt.Int32Wrap(c2) + hxrt.Int32Wrap(1))), 0, false)
		self.backslash = true
	} else {
		if c2 < c1 {
			self.dir = hxrt.StringSubstrStringPtr(path, 0, c1, true)
			path = hxrt.StringSubstrStringPtr(path, int((hxrt.Int32Wrap(c1) + hxrt.Int32Wrap(1))), 0, false)
		} else {
			self.dir = nil
		}
	}
	cp := hxrt.StringLastIndexOfStringPtr(path, hxrt.StringFromLiteral("."), 0, false)
	if cp != -1 {
		self.ext = hxrt.StringSubstrStringPtr(path, int((hxrt.Int32Wrap(cp) + hxrt.Int32Wrap(1))), 0, false)
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
		if c2 != int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1))) {
			hx_if_1 = hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("\\"))
		} else {
			hx_if_1 = path
		}
		hx_if_3 = hx_if_1
	} else {
		var hx_if_2 *string
		if c1 != int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1))) {
			hx_if_2 = hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("/"))
		} else {
			hx_if_2 = path
		}
		hx_if_3 = hx_if_2
	}
	return hx_if_3
}

func haxe__io__Path_join(paths *hxrt.Array) *string {
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := paths
	for _g1 < _g2.Len() {
		v := func(hx_value_4 any) *string {
			if hx_value_4 == nil {
				var hx_zero_5 *string
				return hx_zero_5
			}
			return hx_value_4.(*string)
		}(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		if !hxrt.StringEqualStringPtr(v, nil) && !hxrt.StringEqualStringPtr(v, hxrt.StringFromLiteral("")) {
			_g.Push(v)
		}
	}
	paths_1 := _g
	if paths_1.Len() == 0 {
		return hxrt.StringFromLiteral("")
	}
	path := func(hx_value_7 any) *string {
		if hx_value_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_value_7.(*string)
	}(paths_1.Get(0))
	_g_1 := 1
	_g1_1 := paths_1.Len()
	for _g_1 < _g1_1 {
		hx_post_9 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		i := hx_post_9
		path = haxe__io__Path_addTrailingSlash(path)
		path = hxrt.StringConcatAny(path, paths_1.Get(i))
	}
	return haxe__io__Path_normalize(path)
}

func haxe__io__Path_normalize(path *string) *string {
	slash := hxrt.StringFromLiteral("/")
	path = hxrt.StringJoinAny(hxrt.ArrayFromValues(func(hx_sort_src_10 []*string) []any {
		hx_sort_out_12 := make([]any, 0, len(hx_sort_src_10))
		for _, hx_sort_item_11 := range hx_sort_src_10 {
			hx_sort_out_12 = append(hx_sort_out_12, hx_sort_item_11)
		}
		return hx_sort_out_12
	}(hxrt.StringSplitStringPtr(path, hxrt.StringFromLiteral("\\")))).Values(), slash)
	if hxrt.StringEqualStringPtr(path, slash) {
		return slash
	}
	target := hxrt.NewArray()
	_g := 0
	_g1 := hxrt.ArrayFromValues(func(hx_sort_src_13 []*string) []any {
		hx_sort_out_15 := make([]any, 0, len(hx_sort_src_13))
		for _, hx_sort_item_14 := range hx_sort_src_13 {
			hx_sort_out_15 = append(hx_sort_out_15, hx_sort_item_14)
		}
		return hx_sort_out_15
	}(hxrt.StringSplitStringPtr(path, slash)))
	for _g < _g1.Len() {
		token := func(hx_value_16 any) *string {
			if hx_value_16 == nil {
				var hx_zero_17 *string
				return hx_zero_17
			}
			return hx_value_16.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if (hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral("..")) && (target.Len() > 0)) && !hxrt.StringEqualAny(target.Get(int((hxrt.Int32Wrap(target.Len())-hxrt.Int32Wrap(1)))), hxrt.StringFromLiteral("..")) {
			target.Pop()
		} else {
			if hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral("")) {
				if (target.Len() > 0) || (hxrt.StringCharCodeAtAnyStringPtr(path, 0) == 47) {
					target.Push(token)
				}
			} else {
				if !hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral(".")) {
					target.Push(token)
				}
			}
		}
	}
	tmp := hxrt.StringJoinAny(target.Values(), slash)
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
		hx_post_21 := _g_offset
		_g_offset = int(int32((_g_offset + 1)))
		index := hx_post_21
		c_1 := hxrt.StringCharCodeAtStringPtr(value, index)
		if ((c_1 >= 55296) && (c_1 <= 56319)) && (_g_offset < hxrt.StringLengthStringPtr(_g_s)) {
			c_1 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(c_1) - hxrt.Int32Wrap(55232)))) << uint(10)))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(func() int {
				value_1 := _g_s
				hx_post_22 := _g_offset
				_g_offset = int(int32((_g_offset + 1)))
				index_1 := hx_post_22
				return hxrt.StringCharCodeAtStringPtr(value_1, index_1)
			}()) & hxrt.Int32Wrap(1023))))))
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
