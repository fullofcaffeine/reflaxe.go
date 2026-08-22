package main

type I_haxe__iterators__HashMapKeyValueIterator interface {
	hasNext() bool
	next() map[string]any
}

type haxe__iterators__HashMapKeyValueIterator struct {
	__hx_this I_haxe__iterators__HashMapKeyValueIterator
	map_      *haxe__ds__HashMap
	keys      map[string]any
}

func New_haxe__iterators__HashMapKeyValueIterator(map_ *haxe__ds__HashMap) *haxe__iterators__HashMapKeyValueIterator {
	self := &haxe__iterators__HashMapKeyValueIterator{}
	self.__hx_this = self
	self.map_ = map_
	_gthis := map_
	hashes := func(hx_value_1 any) map[string]any {
		if hx_value_1 == nil {
			var hx_zero_2 map[string]any
			return hx_zero_2
		}
		return hx_value_1.(map[string]any)
	}(map_.keysByHash.__hx_this.keys())
	hx_obj_3 := map[string]any{}
	hx_obj_3["hasNext"] = func() bool {
		return func(hx_obj_4 map[string]any) func() bool {
			hx_field_5 := hx_obj_4["hasNext"]
			if hx_field_5 == nil {
				var hx_zero_6 func() bool
				return hx_zero_6
			}
			return hx_field_5.(func() bool)
		}(hashes)()
	}
	hx_obj_3["next"] = func() interface{ hashCode() int } {
		return func(hx_value_10 any) interface{ hashCode() int } {
			if hx_value_10 == nil {
				var hx_zero_11 interface{ hashCode() int }
				return hx_zero_11
			}
			return hx_value_10.(interface{ hashCode() int })
		}(_gthis.keysByHash.__hx_this.get(func(hx_obj_7 map[string]any) func() int {
			hx_field_8 := hx_obj_7["next"]
			if hx_field_8 == nil {
				var hx_zero_9 func() int
				return hx_zero_9
			}
			return hx_field_8.(func() int)
		}(hashes)()))
	}
	self.keys = hx_obj_3
	return self
}

func (self *haxe__iterators__HashMapKeyValueIterator) hasNext() bool {
	return func(hx_obj_12 map[string]any) func() bool {
		hx_field_13 := hx_obj_12["hasNext"]
		if hx_field_13 == nil {
			var hx_zero_14 func() bool
			return hx_zero_14
		}
		return hx_field_13.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__HashMapKeyValueIterator) next() map[string]any {
	key := func(hx_obj_15 map[string]any) func() interface{ hashCode() int } {
		hx_field_16 := hx_obj_15["next"]
		if hx_field_16 == nil {
			var hx_zero_17 func() interface{ hashCode() int }
			return hx_zero_17
		}
		return hx_field_16.(func() interface{ hashCode() int })
	}(self.keys)()
	_this := self.map_
	var value any = _this.valuesByHash.__hx_this.get(key.hashCode())
	hx_obj_18 := map[string]any{}
	hx_obj_18["key"] = key
	hx_obj_18["value"] = value
	return hx_obj_18
}
