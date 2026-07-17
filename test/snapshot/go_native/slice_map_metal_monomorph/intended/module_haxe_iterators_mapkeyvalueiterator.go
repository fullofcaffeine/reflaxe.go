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
	self.keys = func(hx_value_18 any) map[string]any {
		if hx_value_18 == nil {
			var hx_zero_19 map[string]any
			return hx_zero_19
		}
		return hx_value_18.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_20 map[string]any) func() bool {
		hx_field_21 := hx_obj_20["hasNext"]
		if hx_field_21 == nil {
			var hx_zero_22 func() bool
			return hx_zero_22
		}
		return hx_field_21.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_23 map[string]any) func() any {
		hx_field_24 := hx_obj_23["next"]
		if hx_field_24 == nil {
			var hx_zero_25 func() any
			return hx_zero_25
		}
		return hx_field_24.(func() any)
	}(self.keys)()
	hx_obj_26 := map[string]any{}
	hx_obj_26["key"] = key
	hx_obj_26["value"] = self.map_.getIMap(key)
	return hx_obj_26
}
