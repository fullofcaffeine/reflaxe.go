package main

import "snapshot/hxrt"

func PortableSurfaceDigest_compute(seed int) *string {
	list := New_haxe__ds__List()
	list.__hx_this.add(seed)
	list.__hx_this.push(int((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(1))))
	list.__hx_this.push(int((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(3))))
	var popValue any = func(hx_value_1 any) any {
		if hx_value_1 == nil {
			return nil
		}
		return hx_value_1.(int)
	}(list.__hx_this.pop())
	listDigest := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(list.length, hxrt.StringFromLiteral("|")), hxrt.StdString(func(hx_value_2 any) any {
		if hx_value_2 == nil {
			return nil
		}
		return hx_value_2.(int)
	}(list.__hx_this.first()))), hxrt.StringFromLiteral("|")), hxrt.StdString(popValue))
	_g := New_haxe__ds__StringMap()
	_g.__hx_this.set(hxrt.StringFromLiteral("a"), seed)
	_g.__hx_this.set(hxrt.StringFromLiteral("bb"), int((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(2))))
	_g.__hx_this.set(hxrt.StringFromLiteral("ccc"), int((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(4))))
	map_ := _g
	mapCount := 0
	mapKeyLen := 0
	mapValueSum := 0
	entry := map_.__hx_this.keyValueIterator()
	for func() bool {
		var hx_callable_3 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(entry, "hasNext"))
		return hxrt.FunctionCall(hx_callable_3)()
	}() {
		entry_1 := func() hxrt.Anonymous {
			var hx_callable_4 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_1(hxrt.AnonymousGet(entry, "next"))
			return hxrt.FunctionCall(hx_callable_4)()
		}()
		key := func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(hxrt.AnonymousGet(entry_1, "key"))
		value := func(hx_value_7 any) int {
			if hx_value_7 == nil {
				var hx_zero_8 int
				return hx_zero_8
			}
			return hx_value_7.(int)
		}(hxrt.AnonymousGet(entry_1, "value"))
		mapCount = int(int32((mapCount + 1)))
		mapKeyLen = int((hxrt.Int32Wrap(mapKeyLen) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(key))))
		mapValueSum = int((hxrt.Int32Wrap(mapValueSum) + hxrt.Int32Wrap(value)))
	}
	unicode := hxrt.StringFromLiteral("A☺B")
	var unicodeIter_s *string
	var unicodeIter_offset int
	unicodeIter_offset = 0
	unicodeIter_s = unicode
	unicodeCount := 0
	unicodeSum := 0
	for unicodeIter_offset < hxrt.StringLengthStringPtr(unicodeIter_s) {
		unicodeCount = int(int32((unicodeCount + 1)))
		value_1 := unicodeIter_s
		hx_post_9 := unicodeIter_offset
		unicodeIter_offset = int(int32((unicodeIter_offset + 1)))
		index := hx_post_9
		c := hxrt.StringCharCodeAtStringPtr(value_1, index)
		if ((c >= 55296) && (c <= 56319)) && (unicodeIter_offset < hxrt.StringLengthStringPtr(unicodeIter_s)) {
			c = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(c) - hxrt.Int32Wrap(55232)))) << uint(10)))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(func() int {
				value_2 := unicodeIter_s
				hx_post_10 := unicodeIter_offset
				unicodeIter_offset = int(int32((unicodeIter_offset + 1)))
				index_1 := hx_post_10
				return hxrt.StringCharCodeAtStringPtr(value_2, index_1)
			}()) & hxrt.Int32Wrap(1023))))))
		}
		unicodeSum = int((hxrt.Int32Wrap(unicodeSum) + hxrt.Int32Wrap(c)))
	}
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(listDigest, hxrt.StringFromLiteral("|")), mapCount), hxrt.StringFromLiteral("|")), mapKeyLen), hxrt.StringFromLiteral("|")), mapValueSum), hxrt.StringFromLiteral("|")), unicodeCount), hxrt.StringFromLiteral("|")), unicodeSum)
}
