package main

import "examples_portable_beta/hxrt"

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
	var hx_if_58 *string
	if c1 < c2 {
		var hx_if_56 *string
		if c2 != int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1)))) {
			hx_if_56 = hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("\\"))
		} else {
			hx_if_56 = path
		}
		hx_if_58 = hx_if_56
	} else {
		var hx_if_57 *string
		if c1 != int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(path)) - hxrt.Int32Wrap(1)))) {
			hx_if_57 = hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("/"))
		} else {
			hx_if_57 = path
		}
		hx_if_58 = hx_if_57
	}
	return hx_if_58
}

func haxe__io__Path_join(paths *hxrt.Array) *string {
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := paths
	for _g1 < _g2.Len() {
		v := func(hx_value_59 any) *string {
			if hx_value_59 == nil {
				var hx_zero_60 *string
				return hx_zero_60
			}
			return hx_value_59.(*string)
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
	path := func(hx_value_62 any) *string {
		if hx_value_62 == nil {
			var hx_zero_63 *string
			return hx_zero_63
		}
		return hx_value_62.(*string)
	}(paths_1.Get(0))
	_g_1 := 1
	_g1_1 := paths_1.Len()
	for _g_1 < _g1_1 {
		hx_post_64 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		i := hx_post_64
		path = haxe__io__Path_addTrailingSlash(path)
		path = hxrt.StringConcatAny(path, paths_1.Get(i))
	}
	return haxe__io__Path_normalize(path)
}

func haxe__io__Path_normalize(path *string) *string {
	slash := hxrt.StringFromLiteral("/")
	path = hxrt.StringJoinAny(hxrt.ArrayFromValues(func(hx_sort_src_65 []*string) []any {
		hx_sort_out_67 := make([]any, 0, len(hx_sort_src_65))
		for _, hx_sort_item_66 := range hx_sort_src_65 {
			hx_sort_out_67 = append(hx_sort_out_67, hx_sort_item_66)
		}
		return hx_sort_out_67
	}(hxrt.StringSplitStringPtr(path, hxrt.StringFromLiteral("\\")))).Values(), slash)
	if hxrt.StringEqualStringPtr(path, slash) {
		return slash
	}
	target := hxrt.NewArray()
	_g := 0
	_g1 := hxrt.ArrayFromValues(func(hx_sort_src_68 []*string) []any {
		hx_sort_out_70 := make([]any, 0, len(hx_sort_src_68))
		for _, hx_sort_item_69 := range hx_sort_src_68 {
			hx_sort_out_70 = append(hx_sort_out_70, hx_sort_item_69)
		}
		return hx_sort_out_70
	}(hxrt.StringSplitStringPtr(path, slash)))
	for _g < _g1.Len() {
		token := func(hx_value_71 any) *string {
			if hx_value_71 == nil {
				var hx_zero_72 *string
				return hx_zero_72
			}
			return hx_value_71.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if (hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral("..")) && (target.Len() > 0)) && !hxrt.StringEqualAny(target.Get(int(int32((hxrt.Int32Wrap(target.Len())-hxrt.Int32Wrap(1))))), hxrt.StringFromLiteral("..")) {
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
		hx_post_76 := _g_offset
		_g_offset = int(int32((_g_offset + 1)))
		index := hx_post_76
		c_1 := hxrt.StringCharCodeAtStringPtr(value, index)
		if ((c_1 >= 55296) && (c_1 <= 56319)) && (_g_offset < hxrt.StringLengthStringPtr(_g_s)) {
			c_1 = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(c_1) - hxrt.Int32Wrap(55232))))) << uint(10))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(func() int {
				value_1 := _g_s
				hx_post_77 := _g_offset
				_g_offset = int(int32((_g_offset + 1)))
				index_1 := hx_post_77
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
