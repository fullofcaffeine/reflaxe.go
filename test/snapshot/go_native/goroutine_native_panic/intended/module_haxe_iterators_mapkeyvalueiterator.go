package main

type I_haxe__iterators__MapKeyValueIterator interface {
	hasNext() bool
	next() map[string]any
}

type haxe__iterators__MapKeyValueIterator struct {
	__hx_this I_haxe__iterators__MapKeyValueIterator
	map_      haxe__IMap
	keys      map[string]any
}

func New_haxe__iterators__MapKeyValueIterator(map_ haxe__IMap) *haxe__iterators__MapKeyValueIterator {
	self := &haxe__iterators__MapKeyValueIterator{}
	self.__hx_this = self
	self.map_ = map_
	self.keys = func(hx_value_10 any) map[string]any {
		if hx_value_10 == nil {
			var hx_zero_11 map[string]any
			return hx_zero_11
		}
		return hx_value_10.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_12 map[string]any) func() bool {
		hx_field_13 := hx_obj_12["hasNext"]
		if hx_field_13 == nil {
			var hx_zero_14 func() bool
			return hx_zero_14
		}
		return hx_field_13.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_15 map[string]any) func() any {
		hx_field_16 := hx_obj_15["next"]
		if hx_field_16 == nil {
			var hx_zero_17 func() any
			return hx_zero_17
		}
		return hx_field_16.(func() any)
	}(self.keys)()
	hx_obj_18 := map[string]any{}
	hx_obj_18["key"] = key
	hx_obj_18["value"] = self.map_.getIMap(key)
	return hx_obj_18
}
