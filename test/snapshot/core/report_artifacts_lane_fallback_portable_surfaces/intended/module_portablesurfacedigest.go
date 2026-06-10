package main

import "snapshot/hxrt"

func PortableSurfaceDigest_compute(seed int) *string {
	list := New_haxe__ds__List()
	list.add(seed)
	list.push(int(int32((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(1)))))
	list.push(int(int32((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(3)))))
	var popValue any = func(hx_value_1 any) any {
		if hx_value_1 == nil {
			return nil
		}
		return hx_value_1.(int)
	}(list.pop())
	listDigest := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(list.length, hxrt.StringFromLiteral("|")), hxrt.StdString(func(hx_value_2 any) any {
		if hx_value_2 == nil {
			return nil
		}
		return hx_value_2.(int)
	}(list.first()))), hxrt.StringFromLiteral("|")), hxrt.StdString(popValue))
	_g := New_haxe__ds__StringMap()
	_g.set(hxrt.StringFromLiteral("a"), seed)
	_g.set(hxrt.StringFromLiteral("bb"), int(int32((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(2)))))
	_g.set(hxrt.StringFromLiteral("ccc"), int(int32((hxrt.Int32Wrap(seed) + hxrt.Int32Wrap(4)))))
	map_ := _g
	mapCount := 0
	mapKeyLen := 0
	mapValueSum := 0
	var entry_map haxe__IMap
	var entry_keys map[string]any
	var map__1 haxe__IMap = map_
	entry_map = map__1
	entry_keys = func(hx_value_3 any) map[string]any {
		if hx_value_3 == nil {
			var hx_zero_4 map[string]any
			return hx_zero_4
		}
		return hx_value_3.(map[string]any)
	}(map__1.keys())
	for func(hx_obj_5 map[string]any) func() bool {
		hx_field_6 := hx_obj_5["hasNext"]
		if hx_field_6 == nil {
			var hx_zero_7 func() bool
			return hx_zero_7
		}
		return hx_field_6.(func() bool)
	}(entry_keys)() {
		entry := func() map[string]any {
			key := func(hx_obj_8 map[string]any) func() *string {
				hx_field_9 := hx_obj_8["next"]
				if hx_field_9 == nil {
					var hx_zero_10 func() *string
					return hx_zero_10
				}
				return hx_field_9.(func() *string)
			}(entry_keys)()
			var value any = func(hx_value_11 any) any {
				if hx_value_11 == nil {
					return nil
				}
				return hx_value_11.(int)
			}(entry_map.get(key))
			hx_obj_12 := map[string]any{}
			hx_obj_12["key"] = any(key)
			hx_obj_12["value"] = any(value)
			return hx_obj_12
		}()
		key_1 := func(hx_obj_13 map[string]any) *string {
			hx_field_14 := hx_obj_13["key"]
			if hx_field_14 == nil {
				var hx_zero_15 *string
				return hx_zero_15
			}
			return hx_field_14.(*string)
		}(entry)
		value_1 := func(hx_obj_16 map[string]any) int {
			hx_field_17 := hx_obj_16["value"]
			if hx_field_17 == nil {
				var hx_zero_18 int
				return hx_zero_18
			}
			return hx_field_17.(int)
		}(entry)
		mapCount = int(int32((mapCount + 1)))
		mapKeyLen = int(int32((hxrt.Int32Wrap(mapKeyLen) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(key_1)))))
		mapValueSum = int(int32((hxrt.Int32Wrap(mapValueSum) + hxrt.Int32Wrap(value_1))))
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
		value_2 := unicodeIter_s
		hx_post_19 := unicodeIter_offset
		unicodeIter_offset = int(int32((unicodeIter_offset + 1)))
		index := hx_post_19
		c := hxrt.StringCharCodeAtStringPtr(value_2, index)
		if ((c >= 55296) && (c <= 56319)) && (unicodeIter_offset < hxrt.StringLengthStringPtr(unicodeIter_s)) {
			c = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(c) - hxrt.Int32Wrap(55232))))) << uint(10))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(func() int {
				value_3 := unicodeIter_s
				hx_post_20 := unicodeIter_offset
				unicodeIter_offset = int(int32((unicodeIter_offset + 1)))
				index_1 := hx_post_20
				return hxrt.StringCharCodeAtStringPtr(value_3, index_1)
			}()) & hxrt.Int32Wrap(1023))))))))
		}
		unicodeSum = int(int32((hxrt.Int32Wrap(unicodeSum) + hxrt.Int32Wrap(c))))
	}
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(listDigest, hxrt.StringFromLiteral("|")), mapCount), hxrt.StringFromLiteral("|")), mapKeyLen), hxrt.StringFromLiteral("|")), mapValueSum), hxrt.StringFromLiteral("|")), unicodeCount), hxrt.StringFromLiteral("|")), unicodeSum)
}
