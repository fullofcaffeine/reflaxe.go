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
	self.keys = func(hx_value_298 any) map[string]any {
		if hx_value_298 == nil {
			var hx_zero_299 map[string]any
			return hx_zero_299
		}
		return hx_value_298.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_300 map[string]any) func() bool {
		hx_field_301 := hx_obj_300["hasNext"]
		if hx_field_301 == nil {
			var hx_zero_302 func() bool
			return hx_zero_302
		}
		return hx_field_301.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_303 map[string]any) func() any {
		hx_field_304 := hx_obj_303["next"]
		if hx_field_304 == nil {
			var hx_zero_305 func() any
			return hx_zero_305
		}
		return hx_field_304.(func() any)
	}(self.keys)()
	hx_obj_306 := map[string]any{}
	hx_obj_306["key"] = key
	hx_obj_306["value"] = self.map_.getIMap(key)
	return hx_obj_306
}
