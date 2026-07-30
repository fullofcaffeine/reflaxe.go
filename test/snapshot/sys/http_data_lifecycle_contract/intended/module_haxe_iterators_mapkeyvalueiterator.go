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
	self.keys = func(hx_value_303 any) map[string]any {
		if hx_value_303 == nil {
			var hx_zero_304 map[string]any
			return hx_zero_304
		}
		return hx_value_303.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_305 map[string]any) func() bool {
		hx_field_306 := hx_obj_305["hasNext"]
		if hx_field_306 == nil {
			var hx_zero_307 func() bool
			return hx_zero_307
		}
		return hx_field_306.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_308 map[string]any) func() any {
		hx_field_309 := hx_obj_308["next"]
		if hx_field_309 == nil {
			var hx_zero_310 func() any
			return hx_zero_310
		}
		return hx_field_309.(func() any)
	}(self.keys)()
	hx_obj_311 := map[string]any{}
	hx_obj_311["key"] = key
	hx_obj_311["value"] = self.map_.getIMap(key)
	return hx_obj_311
}
