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
	self.keys = func(hx_value_284 any) map[string]any {
		if hx_value_284 == nil {
			var hx_zero_285 map[string]any
			return hx_zero_285
		}
		return hx_value_284.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_286 map[string]any) func() bool {
		hx_field_287 := hx_obj_286["hasNext"]
		if hx_field_287 == nil {
			var hx_zero_288 func() bool
			return hx_zero_288
		}
		return hx_field_287.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_289 map[string]any) func() any {
		hx_field_290 := hx_obj_289["next"]
		if hx_field_290 == nil {
			var hx_zero_291 func() any
			return hx_zero_291
		}
		return hx_field_290.(func() any)
	}(self.keys)()
	hx_obj_292 := map[string]any{}
	hx_obj_292["key"] = key
	hx_obj_292["value"] = self.map_.getIMap(key)
	return hx_obj_292
}
