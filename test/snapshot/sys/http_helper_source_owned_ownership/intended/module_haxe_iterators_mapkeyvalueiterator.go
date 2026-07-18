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
	self.keys = func(hx_value_310 any) map[string]any {
		if hx_value_310 == nil {
			var hx_zero_311 map[string]any
			return hx_zero_311
		}
		return hx_value_310.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_312 map[string]any) func() bool {
		hx_field_313 := hx_obj_312["hasNext"]
		if hx_field_313 == nil {
			var hx_zero_314 func() bool
			return hx_zero_314
		}
		return hx_field_313.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_315 map[string]any) func() any {
		hx_field_316 := hx_obj_315["next"]
		if hx_field_316 == nil {
			var hx_zero_317 func() any
			return hx_zero_317
		}
		return hx_field_316.(func() any)
	}(self.keys)()
	hx_obj_318 := map[string]any{}
	hx_obj_318["key"] = key
	hx_obj_318["value"] = self.map_.getIMap(key)
	return hx_obj_318
}
