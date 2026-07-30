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
	self.keys = func(hx_value_279 any) map[string]any {
		if hx_value_279 == nil {
			var hx_zero_280 map[string]any
			return hx_zero_280
		}
		return hx_value_279.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_281 map[string]any) func() bool {
		hx_field_282 := hx_obj_281["hasNext"]
		if hx_field_282 == nil {
			var hx_zero_283 func() bool
			return hx_zero_283
		}
		return hx_field_282.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_284 map[string]any) func() any {
		hx_field_285 := hx_obj_284["next"]
		if hx_field_285 == nil {
			var hx_zero_286 func() any
			return hx_zero_286
		}
		return hx_field_285.(func() any)
	}(self.keys)()
	hx_obj_287 := map[string]any{}
	hx_obj_287["key"] = key
	hx_obj_287["value"] = self.map_.getIMap(key)
	return hx_obj_287
}
