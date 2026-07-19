package main

import "snapshot/hxrt"

func PortableSurfaceDigest_compute(seed int) *string {
	list := New_haxe__ds__List()
	list.__hx_this.add(seed)
	list.__hx_this.push(int(int32((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(1)))))
	list.__hx_this.push(int(int32((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(3)))))
	var popValue any = func(hx_value_6 any) any {
		if hx_value_6 == nil {
			return nil
		}
		return hx_value_6.(int)
	}(list.__hx_this.pop())
	listDigest := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(list.length, hxrt.StringFromLiteral("|")), hxrt.StdString(func(hx_value_7 any) any {
		if hx_value_7 == nil {
			return nil
		}
		return hx_value_7.(int)
	}(list.__hx_this.first()))), hxrt.StringFromLiteral("|")), hxrt.StdString(popValue))
	_g := New_haxe__ds__StringMap()
	_g.__hx_this.set(hxrt.StringFromLiteral("a"), seed)
	_g.__hx_this.set(hxrt.StringFromLiteral("bb"), int(int32((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(2)))))
	_g.__hx_this.set(hxrt.StringFromLiteral("ccc"), int(int32((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(4)))))
	map_ := _g
	mapCount := 0
	mapKeyLen := 0
	mapValueSum := 0
	entry := func(hx_value_8 any) map[string]any {
		if hx_value_8 == nil {
			var hx_zero_9 map[string]any
			return hx_zero_9
		}
		return hx_value_8.(map[string]any)
	}(map_.__hx_this.keyValueIterator())
	for func(hx_obj_10 map[string]any) func() bool {
		hx_field_11 := hx_obj_10["hasNext"]
		if hx_field_11 == nil {
			var hx_zero_12 func() bool
			return hx_zero_12
		}
		return hx_field_11.(func() bool)
	}(entry)() {
		entry_1 := func(hx_obj_13 map[string]any) func() map[string]any {
			hx_field_14 := hx_obj_13["next"]
			if hx_field_14 == nil {
				var hx_zero_15 func() map[string]any
				return hx_zero_15
			}
			return hx_field_14.(func() map[string]any)
		}(entry)()
		key := func(hx_obj_16 map[string]any) *string {
			hx_field_17 := hx_obj_16["key"]
			if hx_field_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_field_17.(*string)
		}(entry_1)
		value := func(hx_obj_19 map[string]any) int {
			hx_field_20 := hx_obj_19["value"]
			if hx_field_20 == nil {
				var hx_zero_21 int
				return hx_zero_21
			}
			return hx_field_20.(int)
		}(entry_1)
		mapCount = int(int32((mapCount + 1)))
		mapKeyLen = int(int32((hxrt.Int32Wrap(mapKeyLen) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(key)))))
		mapValueSum = int(int32((hxrt.Int32Wrap(mapValueSum) + hxrt.Int32Wrap(value))))
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
		hx_post_22 := unicodeIter_offset
		unicodeIter_offset = int(int32((unicodeIter_offset + 1)))
		index := hx_post_22
		c := hxrt.StringCharCodeAtStringPtr(value_1, index)
		if ((c >= 55296) && (c <= 56319)) && (unicodeIter_offset < hxrt.StringLengthStringPtr(unicodeIter_s)) {
			c = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(c) - hxrt.Int32Wrap(55232))))) << uint(10))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(func() int {
				value_2 := unicodeIter_s
				hx_post_23 := unicodeIter_offset
				unicodeIter_offset = int(int32((unicodeIter_offset + 1)))
				index_1 := hx_post_23
				return hxrt.StringCharCodeAtStringPtr(value_2, index_1)
			}()) & hxrt.Int32Wrap(1023))))))))
		}
		unicodeSum = int(int32((hxrt.Int32Wrap(unicodeSum) + hxrt.Int32Wrap(c))))
	}
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(listDigest, hxrt.StringFromLiteral("|")), mapCount), hxrt.StringFromLiteral("|")), mapKeyLen), hxrt.StringFromLiteral("|")), mapValueSum), hxrt.StringFromLiteral("|")), unicodeCount), hxrt.StringFromLiteral("|")), unicodeSum)
}
