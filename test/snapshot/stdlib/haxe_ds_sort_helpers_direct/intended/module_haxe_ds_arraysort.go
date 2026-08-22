package main

import "snapshot/hxrt"

func haxe__ds__ArraySort_doMerge(a *hxrt.Array, cmp func(any, any) int, from int, pivot int, to int, len1 int, len2 int) {
	var first_cut int
	var second_cut int
	var len11 int
	var len22 int
	var new_mid int
	if (len1 == 0) || (len2 == 0) {
		return
	}
	if int(int32((hxrt.Int32Wrap(len1) + hxrt.Int32Wrap(len2)))) == 2 {
		if cmp(a.Get(pivot), a.Get(from)) < 0 {
			haxe__ds__ArraySort_swap(a, pivot, from)
		}
		return
	}
	if len1 > len2 {
		len11 = int(int32((hxrt.Int32Wrap(len1) >> uint(1))))
		first_cut = int(int32((hxrt.Int32Wrap(from) + hxrt.Int32Wrap(len11))))
		second_cut = haxe__ds__ArraySort_lower(a, cmp, pivot, to, first_cut)
		len22 = int(int32((hxrt.Int32Wrap(second_cut) - hxrt.Int32Wrap(pivot))))
	} else {
		len22 = int(int32((hxrt.Int32Wrap(len2) >> uint(1))))
		second_cut = int(int32((hxrt.Int32Wrap(pivot) + hxrt.Int32Wrap(len22))))
		first_cut = haxe__ds__ArraySort_upper(a, cmp, from, pivot, second_cut)
		len11 = int(int32((hxrt.Int32Wrap(first_cut) - hxrt.Int32Wrap(from))))
	}
	haxe__ds__ArraySort_rotate(a, cmp, first_cut, pivot, second_cut)
	new_mid = int(int32((hxrt.Int32Wrap(first_cut) + hxrt.Int32Wrap(len22))))
	haxe__ds__ArraySort_doMerge(a, cmp, from, first_cut, new_mid, len11, len22)
	haxe__ds__ArraySort_doMerge(a, cmp, new_mid, second_cut, to, int(int32((hxrt.Int32Wrap(len1) - hxrt.Int32Wrap(len11)))), int(int32((hxrt.Int32Wrap(len2) - hxrt.Int32Wrap(len22)))))
}

func haxe__ds__ArraySort_gcd(m int, n int) int {
	for n != 0 {
		t := int(int32((hxrt.Int32Wrap(m) % hxrt.Int32Wrap(n))))
		m = n
		n = t
	}
	return m
}

func haxe__ds__ArraySort_lower(a *hxrt.Array, cmp func(any, any) int, from int, to int, val int) int {
	len := int(int32((hxrt.Int32Wrap(to) - hxrt.Int32Wrap(from))))
	var half int
	var mid int
	for len > 0 {
		half = int(int32((hxrt.Int32Wrap(len) >> uint(1))))
		mid = int(int32((hxrt.Int32Wrap(from) + hxrt.Int32Wrap(half))))
		if cmp(a.Get(mid), a.Get(val)) < 0 {
			from = int(int32((hxrt.Int32Wrap(mid) + hxrt.Int32Wrap(1))))
			len = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(len) - hxrt.Int32Wrap(half))))) - hxrt.Int32Wrap(1))))
		} else {
			len = half
		}
	}
	return from
}

func haxe__ds__ArraySort_rec(a *hxrt.Array, cmp func(any, any) int, from int, to int) {
	middle := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(from) + hxrt.Int32Wrap(to))))) >> uint(1))))
	if int(int32((hxrt.Int32Wrap(to) - hxrt.Int32Wrap(from)))) < 12 {
		if to <= from {
			return
		}
		_g := int(int32((hxrt.Int32Wrap(from) + hxrt.Int32Wrap(1))))
		_g1 := to
		for _g < _g1 {
			hx_post_1 := _g
			_g = int(int32((_g + 1)))
			i := hx_post_1
			j := i
			for j > from {
				if cmp(a.Get(j), a.Get(int(int32((hxrt.Int32Wrap(j)-hxrt.Int32Wrap(1)))))) < 0 {
					haxe__ds__ArraySort_swap(a, int(int32((hxrt.Int32Wrap(j) - hxrt.Int32Wrap(1)))), j)
				} else {
					break
				}
				j = int(int32((j - 1)))
			}
		}
		return
	}
	haxe__ds__ArraySort_rec(a, cmp, from, middle)
	haxe__ds__ArraySort_rec(a, cmp, middle, to)
	haxe__ds__ArraySort_doMerge(a, cmp, from, middle, to, int(int32((hxrt.Int32Wrap(middle) - hxrt.Int32Wrap(from)))), int(int32((hxrt.Int32Wrap(to) - hxrt.Int32Wrap(middle)))))
}

func haxe__ds__ArraySort_rotate(a *hxrt.Array, cmp func(any, any) int, from int, mid int, to int) {
	var n int
	if (from == mid) || (mid == to) {
		return
	}
	n = haxe__ds__ArraySort_gcd(int(int32((hxrt.Int32Wrap(to) - hxrt.Int32Wrap(from)))), int(int32((hxrt.Int32Wrap(mid) - hxrt.Int32Wrap(from)))))
	for func() int {
		hx_post_2 := n
		n = int(int32((n - 1)))
		return hx_post_2
	}() != 0 {
		var val any = a.Get(int(int32((hxrt.Int32Wrap(from) + hxrt.Int32Wrap(n)))))
		shift := int(int32((hxrt.Int32Wrap(mid) - hxrt.Int32Wrap(from))))
		p1 := int(int32((hxrt.Int32Wrap(from) + hxrt.Int32Wrap(n))))
		p2 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(from) + hxrt.Int32Wrap(n))))) + hxrt.Int32Wrap(shift))))
		for p2 != int(int32((hxrt.Int32Wrap(from) + hxrt.Int32Wrap(n)))) {
			hx_array_target_3 := a
			hx_array_index_4 := p1
			hx_array_target_3.Set(hx_array_index_4, a.Get(p2))
			p1 = p2
			if int(int32((hxrt.Int32Wrap(to) - hxrt.Int32Wrap(p2)))) > shift {
				p2 = int(int32((hxrt.Int32Wrap(p2) + hxrt.Int32Wrap(shift))))
			} else {
				p2 = int(int32((hxrt.Int32Wrap(from) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(shift) - hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(to) - hxrt.Int32Wrap(p2))))))))))))
			}
		}
		hx_array_target_5 := a
		hx_array_index_6 := p1
		hx_array_target_5.Set(hx_array_index_6, val)
	}
}

func haxe__ds__ArraySort_sort(a *hxrt.Array, cmp func(any, any) int) {
	haxe__ds__ArraySort_rec(a, cmp, 0, a.Len())
}

func haxe__ds__ArraySort_swap(a *hxrt.Array, i int, j int) {
	var tmp any = a.Get(i)
	hx_array_target_7 := a
	hx_array_index_8 := i
	hx_array_target_7.Set(hx_array_index_8, a.Get(j))
	hx_array_target_9 := a
	hx_array_index_10 := j
	hx_array_target_9.Set(hx_array_index_10, tmp)
}

func haxe__ds__ArraySort_upper(a *hxrt.Array, cmp func(any, any) int, from int, to int, val int) int {
	len := int(int32((hxrt.Int32Wrap(to) - hxrt.Int32Wrap(from))))
	var half int
	var mid int
	for len > 0 {
		half = int(int32((hxrt.Int32Wrap(len) >> uint(1))))
		mid = int(int32((hxrt.Int32Wrap(from) + hxrt.Int32Wrap(half))))
		if cmp(a.Get(val), a.Get(mid)) < 0 {
			len = half
		} else {
			from = int(int32((hxrt.Int32Wrap(mid) + hxrt.Int32Wrap(1))))
			len = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(len) - hxrt.Int32Wrap(half))))) - hxrt.Int32Wrap(1))))
		}
	}
	return from
}
