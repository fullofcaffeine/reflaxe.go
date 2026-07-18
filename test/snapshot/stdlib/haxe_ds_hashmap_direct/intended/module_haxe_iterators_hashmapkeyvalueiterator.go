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
	hashes := func(hx_value_44 any) map[string]any {
		if hx_value_44 == nil {
			var hx_zero_45 map[string]any
			return hx_zero_45
		}
		return hx_value_44.(map[string]any)
	}(map_.keysByHash.__hx_this.keys())
	hx_obj_46 := map[string]any{}
	hx_obj_46["hasNext"] = func() bool {
		return func(hx_obj_47 map[string]any) func() bool {
			hx_field_48 := hx_obj_47["hasNext"]
			if hx_field_48 == nil {
				var hx_zero_49 func() bool
				return hx_zero_49
			}
			return hx_field_48.(func() bool)
		}(hashes)()
	}
	hx_obj_46["next"] = func() interface{ hashCode() int } {
		return func(hx_value_53 any) interface{ hashCode() int } {
			if hx_value_53 == nil {
				var hx_zero_54 interface{ hashCode() int }
				return hx_zero_54
			}
			return hx_value_53.(interface{ hashCode() int })
		}(_gthis.keysByHash.__hx_this.get(func(hx_obj_50 map[string]any) func() int {
			hx_field_51 := hx_obj_50["next"]
			if hx_field_51 == nil {
				var hx_zero_52 func() int
				return hx_zero_52
			}
			return hx_field_51.(func() int)
		}(hashes)()))
	}
	self.keys = hx_obj_46
	return self
}

func (self *haxe__iterators__HashMapKeyValueIterator) hasNext() bool {
	return func(hx_obj_55 map[string]any) func() bool {
		hx_field_56 := hx_obj_55["hasNext"]
		if hx_field_56 == nil {
			var hx_zero_57 func() bool
			return hx_zero_57
		}
		return hx_field_56.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__HashMapKeyValueIterator) next() map[string]any {
	key := func(hx_obj_58 map[string]any) func() interface{ hashCode() int } {
		hx_field_59 := hx_obj_58["next"]
		if hx_field_59 == nil {
			var hx_zero_60 func() interface{ hashCode() int }
			return hx_zero_60
		}
		return hx_field_59.(func() interface{ hashCode() int })
	}(self.keys)()
	_this := self.map_
	var value any = _this.valuesByHash.__hx_this.get(key.hashCode())
	hx_obj_61 := map[string]any{}
	hx_obj_61["key"] = key
	hx_obj_61["value"] = value
	return hx_obj_61
}
