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
	self.keys = func(hx_value_289 any) map[string]any {
		if hx_value_289 == nil {
			var hx_zero_290 map[string]any
			return hx_zero_290
		}
		return hx_value_289.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_291 map[string]any) func() bool {
		hx_field_292 := hx_obj_291["hasNext"]
		if hx_field_292 == nil {
			var hx_zero_293 func() bool
			return hx_zero_293
		}
		return hx_field_292.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_294 map[string]any) func() any {
		hx_field_295 := hx_obj_294["next"]
		if hx_field_295 == nil {
			var hx_zero_296 func() any
			return hx_zero_296
		}
		return hx_field_295.(func() any)
	}(self.keys)()
	hx_obj_297 := map[string]any{}
	hx_obj_297["key"] = key
	hx_obj_297["value"] = self.map_.getIMap(key)
	return hx_obj_297
}
