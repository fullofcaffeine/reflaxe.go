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
	self.keys = func(hx_value_305 any) map[string]any {
		if hx_value_305 == nil {
			var hx_zero_306 map[string]any
			return hx_zero_306
		}
		return hx_value_305.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_307 map[string]any) func() bool {
		hx_field_308 := hx_obj_307["hasNext"]
		if hx_field_308 == nil {
			var hx_zero_309 func() bool
			return hx_zero_309
		}
		return hx_field_308.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_310 map[string]any) func() any {
		hx_field_311 := hx_obj_310["next"]
		if hx_field_311 == nil {
			var hx_zero_312 func() any
			return hx_zero_312
		}
		return hx_field_311.(func() any)
	}(self.keys)()
	hx_obj_313 := map[string]any{}
	hx_obj_313["key"] = key
	hx_obj_313["value"] = self.map_.getIMap(key)
	return hx_obj_313
}
