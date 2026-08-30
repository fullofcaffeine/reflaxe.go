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
	entry := func(hx_value_3 any) map[string]any {
		if hx_value_3 == nil {
			var hx_zero_4 map[string]any
			return hx_zero_4
		}
		return hx_value_3.(map[string]any)
	}(map_.__hx_this.keyValueIterator())
	for func(hx_obj_5 map[string]any) func() bool {
		hx_field_6 := hx_obj_5["hasNext"]
		if hx_field_6 == nil {
			var hx_zero_7 func() bool
			return hx_zero_7
		}
		return hx_field_6.(func() bool)
	}(entry)() {
		entry_1 := func(hx_obj_8 map[string]any) func() map[string]any {
			hx_field_9 := hx_obj_8["next"]
			if hx_field_9 == nil {
				var hx_zero_10 func() map[string]any
				return hx_zero_10
			}
			return hx_field_9.(func() map[string]any)
		}(entry)()
		key := func(hx_obj_11 map[string]any) *string {
			hx_field_12 := hx_obj_11["key"]
			if hx_field_12 == nil {
				var hx_zero_13 *string
				return hx_zero_13
			}
			return hx_field_12.(*string)
		}(entry_1)
		value := func(hx_obj_14 map[string]any) int {
			hx_field_15 := hx_obj_14["value"]
			if hx_field_15 == nil {
				var hx_zero_16 int
				return hx_zero_16
			}
			return hx_field_15.(int)
		}(entry_1)
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
		hx_post_17 := unicodeIter_offset
		unicodeIter_offset = int(int32((unicodeIter_offset + 1)))
		index := hx_post_17
		c := hxrt.StringCharCodeAtStringPtr(value_1, index)
		if ((c >= 55296) && (c <= 56319)) && (unicodeIter_offset < hxrt.StringLengthStringPtr(unicodeIter_s)) {
			c = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(c) - hxrt.Int32Wrap(55232)))) << uint(10)))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(func() int {
				value_2 := unicodeIter_s
				hx_post_18 := unicodeIter_offset
				unicodeIter_offset = int(int32((unicodeIter_offset + 1)))
				index_1 := hx_post_18
				return hxrt.StringCharCodeAtStringPtr(value_2, index_1)
			}()) & hxrt.Int32Wrap(1023))))))
		}
		unicodeSum = int((hxrt.Int32Wrap(unicodeSum) + hxrt.Int32Wrap(c)))
	}
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(listDigest, hxrt.StringFromLiteral("|")), mapCount), hxrt.StringFromLiteral("|")), mapKeyLen), hxrt.StringFromLiteral("|")), mapValueSum), hxrt.StringFromLiteral("|")), unicodeCount), hxrt.StringFromLiteral("|")), unicodeSum)
}
